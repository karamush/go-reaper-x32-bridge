// Command x32bridge is a two way OSC bridge between an X32 console (or the
// emulator of this workspace) and REAPER. It is the Go equivalent of
// X32ReaperW's engine, without the Windows GUI: everything is configured with
// command line flags.
//
// Typical use (REAPER: Control/OSC/web, local listen port 8000, send to 9000):
//
//	x32bridge -x32 127.0.0.1:10023 -host 127.0.0.1 -port 8000 -listen 0.0.0.0:9000
//
// With a real console the -x32 address is the console itself. When the console
// and the bridge are the same machine, x32reaper of this module runs both in one
// process (and feeds REAPER's levels into the console meters).
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

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
	fs := flag.NewFlagSet("x32bridge", flag.ExitOnError)
	cfg := bridge.DefaultConfig()
	extra := cliflags.Register(fs, &cfg)
	showVersion := fs.Bool("version", false, "print the version and exit")
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "x32bridge %s - OSC bridge between an X32 and REAPER\n\n", versionString())
		fs.PrintDefaults()
	}
	_ = fs.Parse(os.Args[1:])
	if *showVersion {
		fmt.Printf("x32bridge %s\n", versionString())
		return
	}

	if err := extra.Apply(&cfg); err != nil {
		fmt.Fprintln(os.Stderr, "x32bridge:", err)
		os.Exit(2)
	}
	b, err := bridge.New(cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "x32bridge:", err)
		os.Exit(2)
	}
	fmt.Printf("x32bridge: X32 %s, REAPER %s (listen %s), bank %d channels at offset %d, tracks %d..%d\n",
		cfg.X32Addr, net.JoinHostPort(cfg.Host, cfg.Port), cfg.Listen,
		cfg.BankSize, cfg.BankOffset, cfg.TrkMin, cfg.TrkMax)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sig
		b.Stop()
	}()
	if err := b.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "x32bridge:", err)
		os.Exit(1)
	}
}
