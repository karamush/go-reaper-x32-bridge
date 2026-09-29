// Command x32reaper is the all-in-one command of this workspace: it runs the X32
// emulator and the bridge to REAPER in one process, so that one program is enough
// for a studio (a real console can still be used instead with -emu=0).
//
// What it adds over running x32emu and x32bridge side by side:
//
//   - the bridge talks to the built-in console on the loopback address, no
//     separate program and no -x32 flag to keep in sync;
//   - REAPER's track meters are fed into the console meter frames, so X32-Edit
//     and Mixing Station show the levels of the DAW (the C emulator shows zeros);
//   - the console answers to the name "REAPER" (-name changes it), which is what
//     clients display.
//
// Typical use (REAPER: Control/OSC/web, local listen port 8000, send to 9000):
//
//	x32reaper -listen 0.0.0.0:9000
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"x32emu/emu"

	"x32bridge/internal/bridge"
	"x32bridge/internal/cliflags"
)

// version, commit and date are injected by the linker in release builds (see
// .goreleaser.yml); a plain "go build" leaves them at these defaults.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func versionString() string {
	return fmt.Sprintf("%s (commit %s, built %s)", version, commit, date)
}

func main() {
	fs := flag.NewFlagSet("x32reaper", flag.ExitOnError)
	cfg := bridge.DefaultConfig()
	extra := cliflags.Register(fs, &cfg)
	showVersion := fs.Bool("version", false, "print the version and exit")

	// console (emulator) options
	emuOn := fs.Bool("emu", true, "run the built-in X32 emulator (0 to use a real console on -x32)")
	emuIP := fs.String("i", "0.0.0.0", "IP address the emulator binds (0.0.0.0 = every interface)")
	emuPort := fs.String("x32port", emu.DefaultPort, "UDP port of the built-in emulator")
	emuName := fs.String("name", emu.DefaultName, "console name (/-prefs/name); empty keeps the resource file value")
	emuRes := fs.String("res", emu.DefaultResource, "resource file of the emulator")
	emuVerbose := fs.Bool("ev", false, "log the console protocol (what -v does in x32emu)")
	emuDebug := fs.Bool("ed", false, "dump the console datagrams in hexadecimal")
	meters := fs.String("meters", "all", "feed REAPER's levels into the console meters: off, ch (channels only) or all")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "x32reaper %s - X32 emulator and REAPER bridge in one process\n\n", versionString())
		fmt.Fprintf(os.Stderr, `Usage: x32reaper [options]

REAPER side:  -host/-port is where we send, -listen is where REAPER sends.
Console side: the built-in emulator listens on -i:-x32port unless -emu=0, in
              which case -x32 names the console to talk to.

`)
		fs.PrintDefaults()
	}
	_ = fs.Parse(os.Args[1:])
	if *showVersion {
		fmt.Printf("x32reaper %s\n", versionString())
		return
	}

	if err := extra.Apply(&cfg); err != nil {
		fmt.Fprintln(os.Stderr, "x32reaper:", err)
		os.Exit(2)
	}

	// Start the console first: the bridge needs its address.
	var console *emu.Emulator
	if *emuOn {
		e, err := emu.New(emu.Config{
			BindIP:   *emuIP,
			Port:     *emuPort,
			Name:     *emuName,
			Resource: *emuRes,
			Verbose:  *emuVerbose,
			Debug:    *emuDebug,
		})
		if err != nil {
			fmt.Fprintln(os.Stderr, "x32reaper: cannot start the console:", err)
			os.Exit(1)
		}
		console = e
		if !cliflags.WasSet(fs, "x32") {
			cfg.X32Addr = console.Address()
		}
	}

	b, err := bridge.New(cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "x32reaper:", err)
		os.Exit(2)
	}
	if console != nil {
		sink, err := meterSink(*meters, console)
		if err != nil {
			fmt.Fprintln(os.Stderr, "x32reaper:", err)
			os.Exit(2)
		}
		b.SetMeters(sink)
	} else if *meters != "off" {
		fmt.Fprintln(os.Stderr, "x32reaper: -meters needs the built-in console (-emu=1): a real X32 meters its own signals")
	}

	fmt.Printf("x32reaper: console %s (%s), REAPER %s (listen %s), bank %d channels at offset %d, tracks %d..%d, meters %s\n",
		cfg.X32Addr, *emuName, net.JoinHostPort(cfg.Host, cfg.Port), cfg.Listen,
		cfg.BankSize, cfg.BankOffset, cfg.TrkMin, cfg.TrkMax, *meters)

	// The console runs in its own goroutine: it stops when a client sends
	// /shutdown (as X32-Edit does when it closes a session) or on a signal, and
	// then the bridge follows.
	consoleDone := make(chan struct{})
	if console != nil {
		go func() {
			if err := console.Run(); err != nil {
				fmt.Fprintln(os.Stderr, "x32reaper: console:", err)
			}
			close(consoleDone)
		}()
		defer console.Close()
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		select {
		case <-sig:
		case <-consoleDone:
			fmt.Println("x32reaper: the console stopped, leaving")
		}
		b.Stop()
		if console != nil {
			console.Stop()
		}
	}()

	if err := b.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "x32reaper:", err)
		os.Exit(1)
	}
}

// meterSink turns a -meters setting into the sink the bridge calls for the levels
// REAPER reports. Only the built-in console can be fed: a real X32 meters its own
// signals, and its meter frames cannot be written from outside.
func meterSink(setting string, console *emu.Emulator) (bridge.MeterSink, error) {
	switch setting {
	case "off":
		return bridge.MeterSink{}, nil
	case "ch", "channels":
		return bridge.MeterSink{
			Track: func(section string, number int, v float32) {
				if section == "ch" {
					console.SetMeterLevel(emu.MeterCh, number-1, v)
				}
			},
		}, nil
	case "all":
		return bridge.MeterSink{
			Track: func(section string, number int, v float32) {
				kind, ok := meterKind(section)
				if !ok {
					return
				}
				console.SetMeterLevel(kind, number-1, v)
			},
			Master: func(left, right float32) {
				console.SetMeterLevel(emu.MeterMain, 0, left)
				console.SetMeterLevel(emu.MeterMain, 1, right)
			},
		}, nil
	}
	return bridge.MeterSink{}, fmt.Errorf("-meters %q is not usable (off, ch or all)", setting)
}

// meterKind maps a console section name onto the meter family of the emulator.
func meterKind(section string) (emu.MeterKind, bool) {
	switch section {
	case "ch":
		return emu.MeterCh, true
	case "auxin":
		return emu.MeterAuxin, true
	case "fxrtn":
		return emu.MeterFxrtn, true
	case "bus":
		return emu.MeterBus, true
	}
	return 0, false
}
