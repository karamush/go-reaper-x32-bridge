# x32bridge / x32reaper — two way OSC bridge between an X32 and REAPER

A Go re-implementation of the engine of **X32ReaperW** (v2.86, Patrick-Gilles
Maillot; upstream project: <https://github.com/pmaillot/X32-Behringer>): the desk
is a control surface for REAPER, REAPER moves the desk, and a bank of 8/16/32
channels follows a block of DAW tracks. No GUI: everything is a command line flag.

Two commands are built from this module:

* **`x32reaper`** — the console *and* the bridge in one process: it runs the
  emulator of `../emulator`, names it `REAPER`, and feeds the console meter frames
  with REAPER's `/track/N/vu` levels (the reference emulator reports zeros).
* **`x32bridge`** — the bridge alone, for a real console or for an emulator started
  separately.

The full reference (address tables, mapping rules, REAPER setup, what REAPER
sends and when, deviations from the C tool) is in
[`../docs/X32_REAPER_BRIDGE.md`](../docs/X32_REAPER_BRIDGE.md).

## Quick start (REAPER on the same machine)

```bash
# from the repository root:
make build
./bin/x32reaper -listen 0.0.0.0:9000     # console on 0.0.0.0:10023 + bridge
```

REAPER: Preferences, Control/OSC/web, add an OSC surface with **local listen port
8000** and **send to 127.0.0.1:9000** (pattern config `Default.ReaperOSC`). The
bridge announces `/device/track/count 32` at startup, so REAPER reports every
track instead of the eight of the factory template. X32-Edit or Mixing Station
connect to the same machine on port 10023 and will show a console named `REAPER`
whose channel meters follow the DAW.

With a real console, or with the two programs started separately:

```bash
./bin/x32emu                                            # console on 0.0.0.0:10023
./bin/x32bridge -x32 127.0.0.1:10023 -host 127.0.0.1 -port 8000 -listen 0.0.0.0:9000
```

Flags: see `x32reaper -h` / `x32bridge -h`, or §5 of the reference document. The
most useful ones:

```
-x32 127.0.0.1:10023    console address (with the built-in console: its own)
-host 127.0.0.1 -port 8000   where REAPER listens
-listen 0.0.0.0:9000    our socket (REAPER "send to")
-bank 32 -bankof 0      channel bank size and initial bank
-trkmin 1 -trkmax 32    REAPER tracks the channel bank can reach
-auxmin/-auxmax, -fxrmin/-fxrmax, -busmin/-busmax, -rdca 1-4,5-8,...
-tox32/-toreaper all|none|fader,pan,name,...   filter masks
-transport -chbank -master    Bank C modes
-meters all|ch|off      feed REAPER's levels into the console meters (x32reaper)
-emu 0                  use a real console instead of the built-in one
-name REAPER -i 0.0.0.0 -x32port 10023 -ev      built-in console options
-v                      log every datagram (off by default)
```

## Tools and tests

```bash
make test                                        # unit tests, incl. meters end to end
X32BRIDGE_LIVE=1 go test ./internal/bridge -run Live -v   # against a real REAPER
./bin/x32listen -listen 0.0.0.0:9000 -t 5s                # dump what REAPER sends
./bin/x32listen -listen 0.0.0.0:9100 -send 127.0.0.1:10023 -msg "/ch/01/mix/fader 0.5"
./bin/x32listen -listen 127.0.0.1:9100 -send 127.0.0.1:10023 -meters 1 -t 3s
```

## Layout

| path | content |
| --- | --- |
| `internal/bridge/config.go` | configuration, filter masks, strip <-> track mapping |
| `internal/bridge/reaper.go` | REAPER side: bundle splitting, `/track/...`, `/master/...`, transport, meters |
| `internal/bridge/x32.go` | console side: `/ch/...`, `/-stat/...`, Bank C transport remote |
| `internal/bridge/msg.go` | message builders for both sides |
| `internal/bridge/bridge.go` | sockets, main loop, handshake, bank recall |
| `internal/bridge/bank.go` | per track state cache (what makes a bank change instant) |
| `internal/cliflags` | the flags shared by the two commands |
| `cmd/x32reaper` | console + bridge in one process (the recommended entry point) |
| `cmd/x32bridge` | the bridge alone |
| `cmd/x32listen` | OSC dump tool, with meter frame decoding (`-meters <n>`) |

## License

GNU GPL version 3 or later: this module re-implements the engine of the GPLv3
`X32ReaperW.c` and links the `x32emu` module (itself a port of the GPLv3 `X32.c`).
See `../LICENSE` for the full text and `../NOTICE` for the attribution.
