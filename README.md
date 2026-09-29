[![CI](https://github.com/karamush/go-reaper-x32-bridge/actions/workflows/ci.yml/badge.svg)](https://github.com/karamush/go-reaper-x32-bridge/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/karamush/go-reaper-x32-bridge.svg)](https://github.com/karamush/go-reaper-x32-bridge/releases/latest)
[![Downloads](https://img.shields.io/github/downloads/karamush/go-reaper-x32-bridge/total.svg)](https://github.com/karamush/go-reaper-x32-bridge/releases)

# X32 OSC emulator + X32 to REAPER bridge

A **Behringer X32 console emulator** and a two way **X32 <-> REAPER bridge**, in Go.
Standard library only, no dependencies, one process, command line flags only.

* **One binary does both.** `x32reaper` is a virtual X32 that X32-Edit and Mixing
  Station connect to on UDP port 10023, *and* the bridge that makes that console a
  control surface for REAPER - including REAPER's track metres on the console.
* **Byte for byte compatible** with the reference X32 emulator it was ported from
  (`X32.c` v0.88): same replies, same `/node` texts, same notifications, same
  meter frames, same silence when nothing changed. Verified by a differential
  test, and by connecting real clients (X32-Edit, Mixing Station).
* **The protocol is documented** for re-implementation in another language:
  `docs/X32_OSC_PROTOCOL.md` plus a generated catalogue of all 20 694 command
  table entries.

```
  X32-Edit / Mixing Station          x32reaper                    REAPER
        phone, tablet                (this project)            (Windows/macOS/Linux)
             │                            │                            │
             └──────── OSC/UDP ──────────>│   console side: 0.0.0.0:10023
                                          │   bridge  side: 0.0.0.0:9000 <─── OSC ───┐
                                          └──────── OSC/UDP ────────────────────────>│ 127.0.0.1:8000
```

## Features

Emulator (`x32emu`, also built into `x32reaper`):

| area | what is implemented |
| --- | --- |
| transport | OSC 1.0 over UDP, port 10023 (`-port`), bind address `-i` (0.0.0.0 by default), real host IP in `/info`, `/xinfo`, `/status` |
| discovery | `/xinfo`, `/info`, `/status`, `/-prefs/name` (console name), `/shutdown` |
| command tree | 221 tables / 20 694 entries: `/ch`, `/auxin`, `/fxrtn`, `/bus`, `/mtx`, `/dca`, `/fx`, `/outputs`, `/headamp`, `/config`, `/-prefs`, `/-stat`, `/-show`, `/-libs`, `/insert`, `/preamp`, `/main`, … |
| reads | GET of a leaf (with `,i` / `,f` / `,s` / no tag / empty tag), `/node` group and single parameter replies, `/showdump` |
| writes | SET with change detection, echo to the other clients, silence when nothing changed, the string-path `/` command with all its value scales |
| multi client | `/xremote` (four clients, 11 s window), `/renew`, `/unsubscribe`, the two notification shapes a real desk sends |
| meters | `/meters/NN` for all 17 groups, reference frame layout and cadence, plus a level feed (`SetMeterLevel`) used by the bridge |
| library | `/copy`, `/save`, `/delete`, `/add`, `/load` with the scene/snippet notifications |
| state | `.X32res.rc` resource file, byte compatible with the reference emulator |
| extension | `/-stat/userpar/NN/value` user assign controls, the addresses `X32ReaperW` uses to receive transport commands from a desk |

Bridge (`x32reaper`, `x32bridge`, `x32listen`):

| area | what is implemented |
| --- | --- |
| console to REAPER | faders, pans, mutes (inverted on the dish), solos, track selects, track names, send levels, master, DCA ranges, mute groups, Bank C transport remote |
| REAPER to console | fader/pan/mute/solo/select/name/send/master feedback, transport state, plus REAPER's **track and master metres** into the console meter frames |
| banks | channel banks of 8/16/32 tracks (`-bank`, `-bankof`, `-trkmin`, `-trkmax`), instant recall from a per track cache |
| filters | `-tox32` / `-toreaper` masks (fader, pan, name, on, solo, select, send, DCA, transport, meters) |
| layout | `-auxmin/-auxmax`, `-fxrmin/-fxrmax`, `-busmin/-busmax`, `-rdca` and friends map console sections onto REAPER track ranges |
| tools | `x32listen` dumps and decodes OSC on either side (`-meters <n>` prints meter values) |

## Install

**From a release** (no Go needed): download the archive for your platform from the
[releases page](https://github.com/karamush/go-reaper-x32-bridge/releases/latest)
(`tar.gz` for Linux/macOS, `zip` for Windows) and unpack it. Each archive holds the
five commands, this README, the two protocol documents, the licence and a
`checksums.txt`.

```bash
tar xzf x32-reaper-bridge_0.1.0_linux_amd64.tar.gz
cd x32-reaper-bridge_0.1.0_linux_amd64
./x32reaper -h
```

**From source** (needs Go 1.23 or newer, nothing else):

```bash
make build     # or: go build ./...  - the five commands land in bin/
```

Archives are built for Linux (386/amd64/arm/arm64), macOS (amd64/arm64) and
Windows (386/amd64/arm64); a single command is trivial to build for anything else:

```bash
GOOS=linux GOARCH=arm64 go build -o bin/x32reaper ./bridge/cmd/x32reaper
```

## Quick start (REAPER on the same machine)

```bash
# 1. build; the binaries land in bin/
make build

# 2. run: console on 0.0.0.0:10023 + bridge
./bin/x32reaper -listen 0.0.0.0:9000

# 3. REAPER: Preferences > Control/OSC/web > add "OSC (Open Sound Control)"
#      mode                     : Configure device IP+local port
#      local listen port        : 8000
#      device IP / port         : 127.0.0.1 / 9000
#      pattern config           : Default.ReaperOSC
#    (this is the "send to" side: our -listen address)

# 4. connect X32-Edit or Mixing Station to <this machine>:10023
#    Channel faders move REAPER's tracks, REAPER's levels move the channel metres.
```

With a **real console** instead of the built in one, or with the two halves in
separate processes:

```bash
./bin/x32reaper -emu 0 -x32 192.168.1.50:10023 -listen 0.0.0.0:9000

./bin/x32emu -i 0.0.0.0                        # console only, port 10023
./bin/x32bridge -x32 127.0.0.1:10023 -host 127.0.0.1 -port 8000 -listen 0.0.0.0:9000
```

Every flag is documented by `-h`, in `docs/X32_REAPER_BRIDGE.md` section 5 (bridge)
and in `emulator/README.md` (console). Nothing is logged unless asked: `-v` on the
bridge, `-v 1` on the console.

## Build and test

Requirements: **Go 1.23 or newer** (nothing else; the code uses the standard
library only). The repository is a Go workspace (`go.work`) over the two modules
`emulator/` and `bridge/`.

```bash
make build      # the five commands into bin/
make test       # unit tests of both modules
make check      # gofmt + go vet + go test
make snapshot   # local release dry run into dist/ (needs goreleaser)
make help       # list the targets
```

Without `make`:

```bash
go build ./emulator/... ./bridge/...   # from the repository root
go test  ./emulator/... ./bridge/...   # unit tests of both modules
go vet   ./emulator/... ./bridge/...
```

On Windows, `tools\build.ps1` builds the same five binaries (`-Release` strips
the symbols, `-TargetOS linux -Arch arm64` cross compiles).

Tools that need something extra (all optional):

```bash
# what REAPER actually sends on port 9000 (any OSC source will do, not only REAPER)
./bin/x32listen -listen 0.0.0.0:9000

# what a console client receives, with the meter frames decoded
./bin/x32listen -listen 127.0.0.1:9100 -send 127.0.0.1:10023 -meters 1 -t 3s
```

Two helper scripts are PowerShell on purpose (they drive Windows binaries and
parse the C sources) and `pwsh` runs them on Linux and macOS too:

```bash
# differential test against the reference emulator binary (see "How it was validated")
pwsh -File tools/difftest.ps1 -Oracle /path/to/X32-emulator.exe

# regenerate the command catalogue and the Go tables from the upstream C tables
pwsh -File tools/gen_tables.ps1 -Src /path/to/X32-Behringer
pwsh -File tools/gen_appendix.ps1
```

## Releases and CI

* **Every push** runs CI ([`.github/workflows/ci.yml`](.github/workflows/ci.yml)):
  `gofmt`, `go vet` and `go test` on Ubuntu and Windows, then a
  [GoReleaser](https://goreleaser.com) snapshot build which uploads the
  per-platform archives as workflow artifacts - downloadable from the run page
  without publishing anything.
* **A `v*` tag** runs [`release.yml`](.github/workflows/release.yml): the same
  build, but the archives and `checksums.txt` are attached to a GitHub release
  (marked as a pre-release automatically for tags such as `v0.2.0-rc1`):

  ```bash
  git tag v0.1.0
  git push origin v0.1.0
  ```

* **Locally**: `make snapshot` (or `goreleaser release --clean --snapshot`)
  writes `dist/` without touching GitHub. The configuration is
  [`.goreleaser.yml`](.goreleaser.yml): five commands, Linux/macOS/Windows,
  linker-injected version (`x32reaper -version`).

## Repository layout

| path | what it is |
| --- | --- |
| `emulator/` | **the console** (`x32emu`), plus the `emu` library package used by the all-in-one command |
| `bridge/` | **the bridge** (`x32reaper`, `x32bridge`, `x32listen`) |
| `docs/X32_OSC_PROTOCOL.md` | the protocol reference: transport, OSC encoding, dispatch, command tree, GET/SET and notification semantics, `/node`, `/meters`, `/xremote`, value scales, persistence, deviations, acceptance checklist |
| `docs/X32_REAPER_BRIDGE.md` | the bridge reference: address tables both ways, mapping rules, REAPER setup, what REAPER sends and when, meters, deviations from the C tool |
| `docs/APPENDIX_COMMANDS.md` | generated: every table, every command, its type and enum list |
| `docs/x32_commands.json`, `docs/x32_enums.json` | the same catalogue in machine readable form (for code generation in any language) |
| `tools/` | `build.ps1` (the Windows build script), `gen_tables.ps1`, `gen_appendix.ps1`, `difftest.ps1` |
| `Makefile` | the universal build entry point: `make build`, `make test`, `make check`, `make snapshot` |
| `.goreleaser.yml` | release configuration: the five commands for Linux/macOS/Windows |
| `.github/workflows/` | `ci.yml` (fmt/vet/test on Ubuntu and Windows + snapshot archives) and `release.yml` (releases on `v*` tags) |
| `NOTICE`, `LICENSE` | attribution to the upstream project and the GPLv3 text |

## How it was validated

* **Differential test.** The same ~135 step scripted OSC session (discovery, GETs
  of every subsystem, SETs with and without changes, `/node` group and single
  queries, `/` commands, `/copy`, `/save`, `/delete`, `/add`, `/load`,
  `/showdump`, meter subscriptions, `lock`, `/unsubscribe`, unknown commands) is
  replayed against the reference emulator and against this port; every datagram is
  captured in hex and compared. Result: two logs of ~440 lines and **14 protocol
  differences**, all of them intentional and listed in `docs/X32_OSC_PROTOCOL.md`
  section 14.2 (for example the `/status` IP field and the notifications the
  reference forgets for string-path single parameter writes). Everything else -
  every GET reply, `/node` text, echo, meter frame and every silence - is byte
  identical.
* **Real clients.** X32-Edit and Mixing Station connect, pull the `/node`
  snapshots, subscribe to `/meters` and stay in sync with each other and with
  changes made from the desk.
* **Real REAPER.** 368 distinct console addresses were observed coming from
  REAPER's feedback (16 tracks x 16 sends, names, pans, mutes, solos, master) and
  the desk to REAPER path was confirmed end to end (`/ch/01/mix/fader` ->
  `/track/1/volume`). REAPER's `/track/N/vu` levels (about 20 bundles per second)
  are fed into the console meter frames: measured, `/track/15/vu/R 0.5483`
  arrives as value 14 of `/meters/1`, the index `X32Automix.c` reads for channel
  15.
* Unit tests cover the OSC layer, the protocol semantics, the 11 s `/xremote`
  window, `/node` rendering, the meter frame layout and the level feed, the
  resource file round trip, and the bridge mapping (both modules run in CI on
  Ubuntu and Windows).

## Credits and license

This project is a **derivative work of the X32-Behringer project by
Patrick-Gilles Maillot** (<https://github.com/pmaillot/X32-Behringer>), released
under the GNU GPL v3 or later:

* `emulator/` is a port of `X32.c` v0.88 ("X32 - An X32 Emulator") and of the OSC
  helpers in `X32lib/`;
* the command catalogue (`emulator/internal/x32/*_gen.go`, `docs/APPENDIX_COMMANDS.md`,
  `docs/x32_commands.json`, `docs/x32_enums.json`) is mechanically extracted from
  that project's command table headers;
* `bridge/` is a re-implementation of the engine of `X32ReaperW.c` v2.86.

The upstream C sources and its prebuilt `X32-emulator.exe` are **not** part of
this repository. To build them (and to run the differential test), clone the
upstream project - it builds with `make` (MinGW/MSVC on Windows) - and point the
generator at it with `tools\gen_tables.ps1 -Src <checkout>`.

Because it is derived from GPLv3 code, **this project is licensed under the
GNU GPL version 3 or later** - see `LICENSE` (full text) and `NOTICE`
(attribution and what exactly is derived from what). There is no warranty.

