# x32emu — X32 emulator (Go port of X32-Behringer's `X32.c`)

A Go implementation of the X32 emulator of the
[X32-Behringer](https://github.com/pmaillot/X32-Behringer) repository
(`X32.c` v0.88, "X32 – An X32 Emulator" by Patrick-Gilles Maillot, GPLv3).

It speaks plain OSC 1.0 over UDP on port 10023 and behaves like a real console as
seen by X32-Edit, Mixing Station, `X32_Command` and any other OSC client:
discovery (`/xinfo`), `/info`, `/status`, the whole command tree
(`/ch`, `/auxin`, `/fxrtn`, `/bus`, `/mtx`, `/dca`, `/fx`, `/outputs`,
`/headamp`, `/config`, `/-prefs`, `/-stat`, `/-show`, `/-libs`, …), `/node`
bulk reads, the string-path `/` command, `/xremote` notifications for up to four
clients, `/meters` subscriptions, the user assign controls
(`/-stat/userpar/NN/value`, used by control surface bridges) and `.X32res.rc`
state persistence.

No audio engine: like the reference emulator, meters report zeros — unless an
outside source feeds them (`emu.Emulator.SetMeterLevel`, used by the all-in-one
`x32reaper` command of the `bridge/` module, which pipes REAPER's `/track/N/vu`
levels into the frames: see `docs/X32_REAPER_BRIDGE.md` §6).

The console is named `REAPER` by default (`-name` changes it), so that clients
such as X32-Edit and Mixing Station display the studio setup; a name stored in
`/-prefs/name` still wins.

## Build, run, test

```bash
# from the repository root (make build puts them in bin/):
make build && ./bin/x32emu -i 0.0.0.0    # listen on every interface (default)
./bin/x32emu -i 127.0.0.1                # or bind one address only
./bin/x32emu -name "Studio"              # console name

# or work in this directory:
go test ./...                            # OSC layer + protocol behaviour
go run ./cmd/x32emu -i 0.0.0.0
```

Flags (mirroring `X32.c`): `-i <ip>` bind address — `0.0.0.0` (the default)
listens everywhere and still advertises a usable IP in `/info`, `/xinfo`,
`/status`; `-d 0/1` debug dump, `-v 0/1` log the protocol (**`0`** here, `1` in
the C emulator), `-x/-b/-f/-r/-m 0/1` echo of
`/xremote`, `/batchsubscribe`, `/formatsubscribe`, `/renew`, `/meters`,
`-res <file>` resource file (default `.X32res.rc`), `-name <name>` console name
(default `REAPER`), `-port <n>` UDP port (default `10023`).
Non-Behringer command: `/shutdown` saves the resource file and exits.

## Layout

| path | content |
| --- | --- |
| `emu/` | the console as a library (`New`, `Run`, `Stop`, `Address`, `SetMeterLevel`), used by `x32reaper` of the bridge module and by the tests |
| `internal/osc` | OSC encode/decode/dump, "old" notation, meter blob frames |
| `internal/x32/types.go` | flags, command/table types, formatting and scale helpers |
| `internal/x32/dispatch.go` | `Xheader`/`Xnode` tables, subsystem handlers, main packet entry |
| `internal/x32/command.go` | `funct_params()`: GET/SET, change detection, notification policy |
| `internal/x32/setters.go`, `slash.go`, `slash2.go` | value parsers and the string-path `/` command |
| `internal/x32/noderender.go`, `nodetypes*.go`, `nodes.go` | `/node` rendering for every group type |
| `internal/x32/special.go`, `copysave.go` | `/info`, `/xinfo`, `/status`, `/showdump`, `/copy`, `/save`, … |
| `internal/x32/meters.go`, `meterfeed.go` | meter subscriptions, frames, the layout table and the level feed |
| `internal/x32/userpar.go` | user assign controls `/-stat/userpar/NN/value` (extension over the reference) |
| `internal/x32/persist.go` | `.X32res.rc` (format compatible with `X32.c`) |
| `internal/x32/server.go`, `state.go` | UDP server, clients, notifications |
| `internal/x32/*_gen.go` | generated: types, enum arrays, 221 command tables (20 694 entries) |

The generated files come from the command tables of the upstream X32-Behringer
project, which is not part of this repository (clone
<https://github.com/pmaillot/X32-Behringer> and point `-Src` at that checkout):

```bash
pwsh -File ../tools/gen_tables.ps1 -Src /src/X32-Behringer   # docs/*.json + *_gen.go
pwsh -File ../tools/gen_appendix.ps1                         # docs/APPENDIX_COMMANDS.md
```

The generator runs `gofmt -w` on the Go files it writes, so regenerating the
tables cannot break the CI format check.

## Validation

* `go test ./...` — unit tests: OSC padding/tag handling, the README dialog trace
  (byte exact), GET/SET/echo/notification semantics, `/node` replies, the meter
  frame layout and the level feed (offsets and little-endian values), the console
  name, the user assign controls and persistence round trip. `emu/emu_test.go`
  runs a real socket: `/xinfo`, then a meter subscription filled through
  `SetMeterLevel`.
* `pwsh -File ../tools/difftest.ps1 -Oracle <X32-emulator.exe>` — differential test: the same ~135 step
  OSC session is replayed against the reference `X32-emulator.exe` and against this
  port; every datagram is compared in hex. Current result: **435 lines from the C
  emulator, 446 from the port, 14 protocol differences**, all of them intentional
  deviations documented in `docs/X32_OSC_PROTOCOL.md` §14.2 (`/status` IP field,
  notifications for string-path single parameter sets); asynchronous meter frames
  are counted apart (they are not answers to a request, so one frame more or less
  is a timing artifact). The oracle binary is not part of this repository: the
  script looks for `X32-emulator.exe` next to the repository root or in
  `X32-Behringer\`, `-Oracle <path>` points it anywhere, and without a binary it
  skips itself with an explanation. The script starts the port with
  `-name "X32 Emulator"` because its own default is `REAPER` (deviation 7 of
  §14.2).
* `go run ./cmd/x32probe -userpar` — the same session plus the user assign
  scenario; it shows the `/-stat/userpar/18/value ,i` notifications the port
  pushes on a button event (the C emulator cannot answer these addresses).

## Known gaps

* FX parameter scaling (`SetFxPar1`/`GetFxPar1` in `X32.c`): parameters are stored
  and reported normalised.
* Meter values are zeros unless something feeds them (the built-in console of
  `x32reaper` does, from REAPER's track meters); blob GET replies (`B32`) are not
  implemented, and everything the reference does not implement stays
  unimplemented (`/batchsubscribe`, `/formatsubscribe`, `/-libs`, `/undo`, scene
  payloads, recorder, RTA data).

## Running it together with REAPER

One binary does both halves (see `docs/X32_REAPER_BRIDGE.md`):

```bash
# from the repository root:
make build
./bin/x32reaper -listen 0.0.0.0:9000
```

Then connect X32-Edit or Mixing Station to the machine (port 10023) and set
REAPER's OSC surface to local listen port 8000 / send to 127.0.0.1:9000.

## License

GNU GPL version 3 or later: this package is a port of GPLv3 code. See `../LICENSE`
for the full text and `../NOTICE` for the attribution (upstream X32-Behringer
project by Patrick-Gilles Maillot).