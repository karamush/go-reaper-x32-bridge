# X32 <-> REAPER bridge (reference and Go implementation)

This document describes the two way OSC bridge between a Behringer X32 and
REAPER: what the original `X32ReaperW` (v2.86, Patrick-Gilles Maillot, of the
upstream X32-Behringer project, <https://github.com/pmaillot/X32-Behringer>) does,
how the Go implementation in `bridge/` reproduces it, and what was verified
against a real REAPER.

```
 X32 console (or x32emu)                                REAPER
      10023  <── OSC ──┐                        ┌── OSC ──>  8000 (local listen port)
                       │                        │
                  ┌────┴────────────────────────┴────┐
                  │        x32bridge (Go)            │
                  │  console side: unbound UDP socket │
                  │  REAPER side: bound to 9000       │
                  └──────────────────────────────────┘
                       REAPER sends to 9000 (send-to)
```

## 1. The bridge in one paragraph

The desk is used as a control surface for REAPER: moving a fader on the console
moves the matching REAPER track, and REAPER's state is mirrored back on the desk
(faders, pans, mutes, solos, names, send levels, master, transport). One "bank"
of 8/16/32 channels follows a block of REAPER tracks; the Bank C buttons and
encoders of the console are programmed as a transport and marker remote.

Nothing is guessed: REAPER itself is told how many tracks the surface displays
(`/device/track/count`), which makes it report the whole range of tracks instead
of the eight of the factory `.ReaperOSC` template.

## 2. Addresses, in both directions

### 2.1 Console -> REAPER (what the bridge listens to on the desk)

| console address | REAPER message | note |
| --- | --- | --- |
| `/ch/NN/mix/fader` | `/track/T/volume` | 1:1 |
| `/ch/NN/mix/pan` | `/track/T/pan` | 1:1 |
| `/ch/NN/mix/on` | `/track/T/mute` | **inverted**: `on 1` means "not muted" |
| `/ch/NN/mix/MM/level` | `/track/T/send/BB/volume` | `BB = MM + SendOffset` |
| `/ch/NN/config/name` | `/track/T/name` | |
| `/auxin/NN/...`, `/fxrtn/NN/...`, `/bus/NN/...`, `/mtx/NN/...` | same `/track/T/...` | their own track ranges |
| `/-stat/selidx` | `/action/40297` + `/track/T/select 1.0` | the desk reports a zero based strip |
| `/-stat/solosw/NN` | `/track/T/solo 0/1` | for a DCA strip: every track of its range |
| `/dca/N/fader` | `/track/T/volume` | one DCA can drive several tracks |
| `/dca/N/on` | `/track/T/mute` | inverted |
| `/dca/N/config/name` | `/track/T/name` | |
| `/main/st/mix/fader`, `/main/st/mix/pan` | `/master/volume`, `/master/pan` | |
| `/-stat/userpar/17..24/value` | transport actions (see 2.4) | the value `0` is the button release |
| `/-stat/userpar/33..36/value` | scrub actions (see 2.4) | `< 64` = backwards, `>= 64` = forwards |

### 2.2 REAPER -> console

| REAPER message | console address |
| --- | --- |
| `/track/T/volume`, `/pan` | `/ch/NN/mix/fader`, `/ch/NN/mix/pan` |
| `/track/T/mute` | `/ch/NN/mix/on` (inverted) |
| `/track/T/solo` | `/-stat/solosw/NN` |
| `/track/T/select` | `/-stat/selidx` (zero based) |
| `/track/T/name` | `/ch/NN/config/name` |
| `/track/T/send/BB/volume` | `/ch/NN/mix/MM/level` (`MM = BB - SendOffset`) |
| `/master/volume`, `/master/pan` | `/main/st/mix/fader`, `/main/st/mix/pan` |
| `/play`, `/pause`, `/stop`, `/record`, `/repeat` | `/-stat/userpar/18, 19, 23, 24, 22/value` |
| `/device/...`, `/reaper/...` | ignored (device state announcements) |
| `/track/T/vu[/L/R]`, `/master/vu` | **not used yet** - see §6 |

`T` is the **absolute** REAPER track number: the bridge does the bank arithmetic
itself and never uses REAPER's own device bank, so the channel bank, the auxin /
fxrtn / bus ranges and the DCA ranges can all be addressed at the same time.

### 2.3 Startup handshake

1. `/info` to the console (a console answers exactly `/info`) - connectivity check;
2. `/xremote` (and every 9 s afterwards: a console drops a client that has been
   silent for 11 s);
3. `/-stat/selidx` - the currently selected strip;
4. Bank C setup: `/config/userctrl/C/enc/1..4` and `/btn/5..12` names,
   `/-stat/userpar/33..36 = 64`, `/-stat/userpar/17..24 = 0`,
   `/config/userctrl/C/color`, `/-stat/userbank = 2`;
5. `/device/track/count <N>` and `/device/send/count 16` to REAPER, where `N`
   covers every track any section of the desk can reach (default 32);
6. `/action/41743` ("refresh the control surface") so that REAPER reports its
   whole state, which fills the desk.

### 2.4 Bank C (transport remote)

Buttons fire on the **release** (`/-stat/userpar/NN/value ,i 0`), as in
X32ReaperW ("ignore non 0 value"); the encoders are endless and are re-centred
(`= 64`) after every turn.

| userpar | console button | REAPER |
| --- | --- | --- |
| 17 | C 1 | `/action/40042` (go to start) |
| 18 | C 2 | `/play 1.0` |
| 19 | C 3 | `/pause 1.0` |
| 20 | C 4 | `/action/40043` (go to end) |
| 21 | C 5 | loop toggle: `/action/40222` / `/action/40223` |
| 22 | C 6 | `/repeat 1.0` |
| 23 | C 7 | `/stop 1.0` |
| 24 | C 8 | `/record 1.0` |
| 33 | encoder 1 | `/action/40842` / `/action/40841` (beat back/forward) |
| 34 | encoder 2 | `/action/40840` / `/action/40839` (measure back/forward) |
| 35 | encoder 3 | `/action/40172` / `/action/40173` (marker back/forward) |
| 36 | encoder 4 | `/action/40318` / `/action/40319` (item left/right) |

With `-chbank` the buttons 21/22 (configurable: `-bankup`, `-bankdn`) move the
channel bank instead of driving the loop and repeat, and with `-transport=false`
the button `-markerbtn` (default 7) inserts a marker (`/action/40157`).

## 3. Track mapping and channel banks

| console strip | REAPER track |
| --- | --- |
| channel `1..BankSize` | `TrkMin + BankOffset*BankSize + strip - 1` |
| auxin `33..40` | `AuxMin + strip - 33` |
| fxrtn `41..48` | `FxrMin + strip - 41` |
| bus `49..64` | `BusMin + strip - 49` |
| DCA `73..80` | the track range `RdcaMin[n]..RdcaMax[n]` |
| master `/main/st` | REAPER's master track |

A section is active only when its range is configured (`-auxmin/-auxmax`,
`-fxrmin/-fxrmax`, `-busmin/-busmax`, `-rdca`); by default only the channel bank
(tracks 1..32) and the master are mapped.

The bridge caches the state of **every track REAPER reports**, even the ones the
desk is not showing. When the bank moves (Bank C up/down, or `-bankof` at
startup) it

1. re-selects the new offset,
2. pushes the cached fader, pan, mute, name and solo of each strip of the new
   bank to the console (a console cannot know about invisible tracks),
3. asks REAPER for a full refresh (`/device/track/count` + `/action/41743`),
   which corrects anything the cache did not know.

This is what makes a bank change look instant. X32ReaperW does the same thing
with its `XMbanktracks` array; the difference is that the Go bridge takes the
initial values from REAPER instead of pushing zeros.

## 4. Filters (masks)

`-tox32` (REAPER -> console) and `-toreaper` (console -> REAPER) accept `all`
(the default), `none` or a list of classes:

`pan, fader, name, mute, select, send, solo, fx, mpan, mfader`

Blocking a class is the usual way to break a feedback fight: for instance
`-tox32 fader,pan,mpan,mfader` lets the desk move REAPER, but REAPER never
touches the faders of the desk (X32ReaperW's `Xxsend`/`Xrsend` bits).

## 5. Command line

For a studio where the console runs on the same machine, one binary does
everything:

```
x32reaper -listen 0.0.0.0:9000
```

`x32reaper` (in `bridge/cmd/x32reaper`) starts the console **and** the bridge in
one process. The bridge then talks to that console on the loopback address, and
the console is named `REAPER` so that clients display it (`-name` changes it,
`-emu=0` uses a real console on `-x32` instead).

Running the two programs side by side stays possible (and is what a real console
needs):

```
x32emu  -i 0.0.0.0                 # the console, port 10023
x32bridge -x32 127.0.0.1:10023 -host 127.0.0.1 -port 8000 -listen 0.0.0.0:9000
```

Both commands accept the same bridge flags (`-x32`, `-listen`, `-host`, `-port`,
…) and `x32reaper` adds the console ones:

| flag | default | meaning |
| --- | --- | --- |
| `-x32` | `127.0.0.1:10023` | console (or emulator) to talk to; with the built-in console it defaults to it |
| `-host`, `-port` | `127.0.0.1`, `8000` | where REAPER listens (REAPER "local listen port") |
| `-listen` | `0.0.0.0:9000` | local socket; REAPER "send to" address |
| `-localport` | `0` | pin the local port of the console side socket (0 = ephemeral, like X32ReaperW) |
| `-bank`, `-bankof` | `32`, `0` | channel bank size and initial bank |
| `-sendoff` | `0` | TrackSendOffset (REAPER renumbers sends when a track also feeds a hardware output) |
| `-trkmin/-trkmax` | `1`, `32` | REAPER tracks the channel bank can reach |
| `-auxmin/-auxmax`, `-fxrmin/-fxrmax`, `-busmin/-busmax` | `0` (off) | auxin, fxrtn and bus track ranges |
| `-dcamin`, `-dcamax`, `-rdca 1-4,5-8,...` | `0` (off) | DCA track ranges |
| `-transport`, `-chbank`, `-master` | `true`, `false`, `true` | Bank C modes and master mirroring |
| `-bankup`, `-bankdn`, `-markerbtn` | `9`, `10`, `7` | Bank C button numbers (5..12) |
| `-tox32`, `-toreaper` | `all` | filter masks (§4) |
| `-delayb`, `-delayg` | `0` | ms between the messages of a bank recall / ordinary messages |
| `-v` | `false` | log every datagram, with the same dump format as `Xfdump`; bundles are split so that the log shows messages |
| `-meters` | `all` | feed REAPER's levels into the console meters: `off`, `ch` (channels only) or `all` (§6) |
| `-emu` | `true` | run the built-in console (`0`: use the one named by `-x32`) |
| `-i`, `-x32port` | `0.0.0.0`, `10023` | where the built-in console binds |
| `-name` | `REAPER` | console name (`/-prefs/name`), what X32-Edit and Mixing Station display |
| `-res` | `.X32res.rc` | resource file of the built-in console |
| `-ev`, `-ed` | `false` | console protocol log / hexadecimal dump (the `-v` of `x32emu`) |

`bridge/cmd/x32listen` is a small OSC dump tool for both sides (it splits
bundles, which REAPER uses for its feedback, and decodes meter frames):

```bash
./bin/x32listen -listen 0.0.0.0:9000 -t 5s                        # what REAPER sends
./bin/x32listen -listen 0.0.0.0:9000 -send 127.0.0.1:8000 -msg "/action/41743"
./bin/x32listen -listen 127.0.0.1:9100 -send 127.0.0.1:10023 -meters 1  # decode /meters/1
```

## 6. Meters (audio levels) — implemented

REAPER sends its meters, and the bridge now feeds them into the console:

```
REAPER  -> /track/N/vu[/L/R] ,f 0.5243   /master/vu[/L/R] ,f 0.8945
bridge  -> console.SetMeterLevel()       (one call per track and per update)
console -> /meters/1 frames of 96 values, channel ch level at index ch - 1
```

Measured on a running REAPER (project playing, `x32reaper` + `x32listen`):

```
R-> <,  24 B: /track/15/vu/L~~,f~~[0.5243]
R-> <,  24 B: /track/15/vu/R~~,f~~[0.5483]
R-> <,  20 B: /master/vu~~,f~~[0.8945]
R-> <,  24 B: /master/vu/L~~~~,f~~[0.8967]
  /meters/1: 96 values, 1 non-zero: 14=0.5483      <- track 15 on console channel 15
```

* REAPER's OSC surface emits about 20 bundles per second, each carrying
  `/track/N/vu`, `/track/N/vu/L`, `/track/N/vu/R` for the tracks that are moving
  plus `/master/vu[/L/R]`, `/track/vu[/L/R]` (the device's own track) and the
  `/time`, `/beat`, `/samples` clocks. X32ReaperW ignores the meters entirely,
  which is why the reference tool bridges no levels.
* a console channel meter is a single bar, so the bridge uses the **loudest side**
  of `/vu/L` and `/vu/R` (the mono `/track/N/vu` value is used as it comes).
* the main bus is fed as two levels (left and right).
* a level is only pushed to the console when it changed, and the cached level is
  re-sent when the channel bank changes, so the bars of the new bank are correct
  immediately.
* `-meters` selects how much is fed: `all` (default: channels, aux inputs, effect
  returns, buses and the main bus), `ch` (channels only, the group whose layout is
  confirmed by `X32Automix.c`) or `off` (the console then reports zeros, exactly
  like the reference emulator).

Nothing changes for a **real** console: `/meters/NN` is a subscription there (the
console meters its own audio), so the feed only exists in the built-in emulator —
which is one of the reasons `x32reaper` runs both halves in one process.

Verify it yourself:

```bash
./bin/x32reaper                                     # console + bridge
./bin/x32listen -listen 127.0.0.1:9100 -send 127.0.0.1:10023 -meters 1 -t 3s
```

## 7. REAPER setup

1. Preferences, Control/OSC/web, add an **OSC (Open Sound Control)** control
   surface, mode "Local port", and set

   | field | value |
   | --- | --- |
   | Local listen port | 8000 (the `-port` of the bridge) |
   | Send to | 127.0.0.1, port 9000 (`-listen` of the bridge) |
   | Pattern config | `Default.ReaperOSC` (a copy you can edit) |

2. The patterns the bridge uses are part of the default template; make sure they
   are not commented out (the file lives in
   `%APPDATA%\REAPER\OSC\Default.ReaperOSC`):

```
TRACK_NAME        s/track/name s/track/@/name
TRACK_VOLUME      n/track/volume n/track/@/volume
TRACK_PAN         n/track/pan n/track/@/pan
TRACK_MUTE        b/track/mute b/track/@/mute
TRACK_SOLO        b/track/solo b/track/@/solo
TRACK_SELECT      b/track/select b/track/@/select
TRACK_SEND_VOLUME n/track/send/@/volume n/track/@/send/@/volume
MASTER_VOLUME     n/master/volume
MASTER_PAN        n/master/pan
ACTION            i/action s/action/str t/action/@ f/action/@/cc
PLAY              t/play
PAUSE             t/pause
STOP              t/stop
RECORD            t/record
REPEAT            t/repeat
DEVICE_TRACK_COUNT i/device/track/count t/device/track/count/@
DEVICE_SEND_COUNT  i/device/send/count t/device/send/count/@
```

(names quoted from a stock `Default.ReaperOSC`; the action is
`TRACK_SEND_VOLUME`, with an underscore.)

   The bridge sends its numeric parameters as **floats** (like X32ReaperW) and its
   actions as bare `/action/<id>` addresses; REAPER accepts both.

3. `DEVICE_TRACK_COUNT 8` in a template limits the feedback to eight tracks; the
   bridge does not rely on the template value, it sends `/device/track/count` at
   startup (32 by default). Verified: with a 16 track project REAPER then reports
   every track (16 x 16 sends, all names, pans, mutes, solos).

### What REAPER sends, and when

Measured with `x32listen -listen 0.0.0.0:9000`:

* continuously, in bundles, about **20 times per second**: `/track/N/vu[/L/R]`,
  `/master/vu[/L/R]`, `/track/vu[/L/R]`, `/time`, `/beat/str`, `/samples/str`,
  `/frames/str`;
* on **change**: the parameters of the track that changed (volume, pan, mute,
  solo, name, send volumes, fx parameters of the selected track);
* on demand: everything, after a **refresh** - `/action/41743` (this is what
  X32ReaperW means by "force REAPER track refresh to ensure sync", and it is what
  the bridge sends at startup and after a bank change).

A fresh surface starts with only the meter/time stream: the bridge therefore
always asks for a refresh, otherwise the desk would stay empty until something
changes in REAPER.

## 8. Verified behaviour (live, against a real REAPER)

`bridge/internal/bridge/live_test.go` (`X32BRIDGE_LIVE=1 go test -run Live -v`)
runs the bridge against a real REAPER without touching the project:

```
REAPER -> desk: 368 distinct addresses
  bank C setup   25
  faders         17       (16 tracks + master)
  pans           17
  mutes          16
  solo           16
  sends          256      (16 tracks x 16 sends)
  names          17
  /ch/01/mix/fader     = 0.6300
  /ch/01/config/name   = Drums
desk -> REAPER: ->R >,   24 B: /track/1/volume~,f~~[0.6300]
```

The unit tests cover the mapping rules, the mute inversion, the send offset, the
select/solo paths, the master, the DCA ranges, the Bank C transport, the encoder
re-centring, the filter masks and the bundle splitting.

### End to end: emulator -> bridge -> REAPER -> bridge -> emulator

`x32emu` + `x32bridge` + a real REAPER, a fader moved on the desk (the console
side of the emulator was driven with `x32listen`):

```
desk  -> emulator : /ch/01/mix/fader ,f 0.65
emulator -> bridge: /ch/01/mix/fader~~~~,f~~[0.6500]
bridge -> REAPER  : /track/1/volume~,f~~[0.6500]
REAPER -> bridge  : /track/1/volume~,f~~[0.6300]     (after the value was put back)
bridge -> desk    : /ch/01/mix/fader ,f 0.6300
```

The starting value (0.6300 for track 1 of the test project) had been pushed to the
desk by the bridge when REAPER reported it, so the whole loop is closed in both
directions.

### End to end: REAPER meters -> console

With `x32reaper` (console and bridge in one process) and a project playing:

```
REAPER -> bridge   : /track/15/vu/L ,f [0.5243]   /track/15/vu/R ,f [0.5483]
bridge -> console  : meter level 0.5483 on console channel 15
console -> client  : /meters/1 frame, 96 values, index 14 = 0.5483
```

The value index is the console channel minus one, exactly what the reference
decoder of `X32Automix.c` reads (see section 7.6 of `X32_OSC_PROTOCOL.md`).
This is the one thing the C emulator cannot do: it has no audio and reports
zeros for every meter, so a DAW can only be metered through this feed.

## 9. Differences from X32ReaperW

| topic | X32ReaperW | Go bridge |
| --- | --- | --- |
| platform | Windows GUI (wWinMain), Linux CLI version separate | console tool, no GUI, everything on the command line |
| track addressing | absolute track numbers, but `/ch/NN` uses the track number itself for the REAPER -> desk direction (only correct when `Xtrk_min = 1`) | one consistent rule in both directions: `strip = track - TrkMin + 1 - BankOffset*BankSize` |
| OSC parsing | fixed byte offsets in the datagram, two digit numbers only | address parsing, any number of digits, truncated datagrams are rejected |
| REAPER feedback | relies on `DEVICE_TRACK_COUNT 8` of the template (only 8 tracks) | announces `/device/track/count N` at startup, so every mapped track reports |
| bank change | pushes its zero initialised `XMbanktracks` array | caches every track REAPER reports, pushes that cache, then asks REAPER to refresh |
| encoders | `userpar` re-centred to 64 (same) but the "scrub while playing" dance of `X32ReaperW` (`/action/40073` around the move) is not reproduced | simple scrub, no play/pause dance |
| markers | optional marker button depending on the mode | same (`-transport=false` + `-markerbtn`) |
| EQ / DYN mirroring | implemented (ReaEQ / ReaComp with conversion curves) | **not implemented** (phase B): `/ch/NN/eq/*` and `/ch/NN/dyn/*` are ignored |
| meters | received and ignored (the reference tool bridges no level) | fed into the built-in console: REAPER's `/track/N/vu` reaches the `/meters/NN` frames, so X32-Edit and Mixing Station show the DAW levels (section 6) |
| packaging | GUI application, the console is always a separate program | `x32reaper` runs the console **and** the bridge in one process (a real console still works with `-emu=0`) |
| resource file | `.X32Reaper.ini` + presets | command line flags (the ini is not read) |
| everything else | same model: `/xremote` every 9 s, `/info` probe, mute inversion, send offset, masks, Bank C layout, REAPER actions |

## 10. Not implemented (next steps)

1. **EQ and dynamics mirroring** (phase B of the plan): `/ch/NN/eq/on`,
   `eq/x/{f,g,q}` <-> ReaEQ `fxparam`, `/ch/NN/dyn/*` <-> ReaComp, with the
   conversion curves of `X32ReaperW` (`X32_EQF`, `X32_DYNTHR`, ...) and the
   `-eqcmp` mode with `-reqindex`/`-rcindex`.
2. `/-stat/screen/CHAN/page` handling (desk shows the EQ page when REAPER opens
   an effect UI).
3. Reading `.X32Reaper.ini` for drop-in compatibility with the C tool.
4. Meter **groups other than the channels** are filled from the layout table of
   section 7.6 of `X32_OSC_PROTOCOL.md`, where only `/meters/1` is confirmed by
   a reference client (`X32Automix.c`): the bus, aux, matrix and main groups are
   inferred from the group sizes. `-meters ch` restricts the feed to the
   confirmed group, and `x32listen -meters <n>` shows what a client receives.

## 11. Build, run, validate

```bash
# build the five commands (bin/)
make build

# one binary: console + bridge
./bin/x32reaper -listen 0.0.0.0:9000

# or the two programs side by side (also for a real console)
./bin/x32emu
./bin/x32bridge -x32 127.0.0.1:10023 -host 127.0.0.1 -port 8000 -listen 0.0.0.0:9000 -v

# tests
make test                                        # unit tests (no network)
X32BRIDGE_LIVE=1 go test ./bridge/internal/bridge -run Live -v   # with a real REAPER

# what REAPER sends
./bin/x32listen -listen 0.0.0.0:9000 -t 5s

# what a console client receives (the meter feed)
./bin/x32listen -listen 127.0.0.1:9100 -send 127.0.0.1:10023 -meters 1 -t 3s
```

A quick manual check of the desk -> REAPER direction (this moves the fader of the
first track of the project, so use a project you can disturb). With `x32emu`
running and the bridge connected:

```bash
./bin/x32listen -listen 0.0.0.0:9100 -send 127.0.0.1:10023 -msg "/ch/01/mix/fader 0.5"
```

The bridge then prints `->R /track/1/volume ,f [0.5]`, REAPER moves the fader and
answers with its own state, which the bridge forwards back to the desk
(`/ch/01/mix/fader ,f [0.5]`).


