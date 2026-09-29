// Package main is the Go port of the X32 emulator of the upstream X32-Behringer
// project (X32.c v0.88 by Patrick-Gilles Maillot,
// https://github.com/pmaillot/X32-Behringer).
//
// The emulator is only a console here: it has no REAPER side. The x32reaper
// command of the bridge module runs the same emulator together with the bridge,
// which is what one wants for a studio (see docs/X32_REAPER_BRIDGE.md).
package main

import (
	"flag"
	"fmt"
	"os"

	"x32emu/emu"
	"x32emu/internal/x32"
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
	showVersion := flag.Bool("version", false, "print the version and exit")
	ip := flag.String("i", "0.0.0.0", "IP address to bind (0.0.0.0 = every interface)")
	port := flag.String("port", x32.UDPPort, "UDP port to listen on")
	name := flag.String("name", x32.DefaultName, "console name (/-prefs/name), empty to keep the resource file value")
	debug := flag.Int("d", 0, "debug option (0/1)")
	verbose := flag.Int("v", 0, "log the protocol (0/1) - the C emulator defaults to 1, this port to 0")
	verbX := flag.Int("x", 0, "echo incoming /xremote (0/1)")
	verbB := flag.Int("b", 0, "echo incoming /batchsubscribe (0/1)")
	verbF := flag.Int("f", 0, "echo incoming /formatsubscribe (0/1)")
	verbR := flag.Int("r", 0, "echo incoming /renew (0/1)")
	verbM := flag.Int("m", 0, "echo incoming /meters (0/1)")
	res := flag.String("res", x32.ResourceFileName, "resource file (default .X32res.rc)")
	flag.Parse()
	if *showVersion {
		fmt.Printf("x32emu %s\n", versionString())
		return
	}

	em, err := emu.New(emu.Config{
		BindIP:     *ip,
		Port:       *port,
		Name:       *name,
		Resource:   *res,
		Verbose:    *verbose != 0,
		Debug:      *debug != 0,
		VerbRemote: *verbX != 0,
		VerbBatch:  *verbB != 0,
		VerbFormat: *verbF != 0,
		VerbRenew:  *verbR != 0,
		VerbMeter:  *verbM != 0,
	})
	if err != nil {
		fmt.Printf("Error on IP address: %s - cannot run (%v)\n", *ip, err)
		os.Exit(1)
	}
	defer em.Close()

	fmt.Printf("X32 - Go port of X32.c v0.88 - An X32 Emulator - (c)2014-2019 Patrick-Gilles Maillot\n")
	if em.Fresh {
		fmt.Printf("X32 resource file does not exist, create one with '/shutdown' command\n")
	}
	// "0.0.0.0" means "listen on every interface": the socket is bound to the
	// wildcard address while /info, /xinfo and /status advertise a usable IPv4 of
	// the system, so that clients can reach us.
	fmt.Printf("Listening to port: %s, X32 IP = %s (bound to %s)\n", em.Port(), em.IP(), em.Bound())
	if err := em.Run(); err != nil {
		fmt.Printf("Error while receiving: %v\n", err)
		os.Exit(1)
	}
}
