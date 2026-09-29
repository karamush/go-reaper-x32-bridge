# Changelog

All notable changes to this project are documented in this file. The format is
based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and the
project uses [semantic versioning](https://semver.org/).

The Go modules are `x32emu` (the console, `emulator/`) and `x32bridge` (the
bridge and its commands, `bridge/`); they are versioned together.

## [0.1.0] - 2026-09-30

First public release: the X32 emulator, the REAPER bridge and the full protocol
documentation.

### Emulator (`emulator/`, commands `x32emu` and `x32probe`)

* Plain OSC 1.0 over UDP on port 10023: discovery (`/xinfo`), `/info`, `/status`,
  `/showdump`, the complete command tree (`/ch`, `/auxin`, `/fxrtn`, `/bus`,
  `/mtx`, `/dca`, `/fx`, `/outputs`, `/headamp`, `/config`, `/-prefs`, `/-stat`,
  `/-show`, `/-libs`, `/insert`, `/preamp`, …) with 20 694 table entries in 221
  tables, `/node` bulk and single reads, the string-path `/` command, `/copy`,
  `/save`, `/delete`, `/add`, `/load`, `lock`, `/renew`, `/unsubscribe`.
* OSC encoding/decoding details that clients depend on: 4 byte padding,
  big-endian `,i`/`,f`, empty tag strings, the "old" address notation, `k`
  suffixed numbers, and the meter frame layout (little-endian float values).
* `/xremote` notification model: up to four clients, an 11 second window, both
  notification shapes (echo of the request packet, one message per parameter),
  change detection with an epsilon, and complete silence when nothing changed.
* `/meters` subscriptions for all 17 groups with the reference frame sizes and
  the `50 ms × timefactor` cadence, plus a **level feed** for an outside source.
* `/-stat/userpar/NN/value` (user assign controls): an intentional extension over
  the reference, used by control surface bridges to receive button events.
* `.X32res.rc` state persistence, byte compatible with the reference emulator.
* `emu` package: the console as a library (`New`, `Run`, `Stop`, `Address`,
  `SetMeterLevel`), used by the all-in-one `x32reaper` command.
* Counters and reports: byte-for-byte parity with the reference for the whole
  test corpus, apart from the documented deviations.

### Bridge (`bridge/`, commands `x32reaper`, `x32bridge` and `x32listen`)

* Two way OSC bridge modelled on `X32ReaperW` v2.86: faders, pans, mutes
  (inverted), solos, selects, names, sends, master, DCA ranges, mute groups and
  the Bank C transport remote.
* Channel banks of 8/16/32 DAW tracks (`-bank`, `-bankof`, `-trkmin/-trkmax`)
  with instant recall from a per track state cache.
* REAPER's **track and master meters** are fed into the console meter frames, so
  X32-Edit and Mixing Station show DAW levels on all 32 channels (`-meters`).
* `/device/track/count` announcement at start up: REAPER reports every track
  instead of the eight of the factory OSC template.
* Filter masks (`-tox32`, `-toreaper`) for both directions and 30+ flags instead
  of the original's Windows GUI.
* `x32reaper` runs the console and the bridge in one process (and names the
  console `REAPER`, `-name` changes it).

### Documentation and tools

* `docs/X32_OSC_PROTOCOL.md`: full protocol reference for re-implementing the
  emulator in any language (transport, dispatch, command tree, GET/SET
  semantics, notifications, `/node`, `/meters`, `/xremote`, value scales,
  persistence, deviations, acceptance checklist).
* `docs/X32_REAPER_BRIDGE.md`: the bridge reference (address tables both ways,
  mapping, REAPER setup, what REAPER sends and when, meters, deviations).
* Generated catalogue: `docs/x32_commands.json`, `docs/x32_enums.json` and
  `docs/APPENDIX_COMMANDS.md` (221 tables, 20 694 entries, 149 enum arrays).
* `tools/gen_tables.ps1`, `tools/gen_appendix.ps1`: catalogue and Go table
  generators; `tools/difftest.ps1`: differential test against the reference
  emulator binary; `tools/build.ps1`: build all commands into `bin/`.
* `Makefile`: universal build entry point (`make build`, `make test`,
  `make check`, `make snapshot`).
* `.goreleaser.yml` and `.github/workflows/release.yml`: a `v*` tag builds the
  five commands for Linux (386/amd64/arm/arm64), macOS (amd64/arm64) and Windows
  (386/amd64/arm64) and publishes one archive per platform - the binaries, the
  two protocol documents, the licence and a `checksums.txt` - with semver
  pre-releases detected automatically. CI builds the same archives on every push
  and leaves them as downloadable workflow artifacts.
* Every command takes `-version` and prints the version, commit and build date
  injected at link time.
