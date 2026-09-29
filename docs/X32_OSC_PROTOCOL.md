# X32 OSC protocol — reference for implementing an X32 emulator

This document describes, at implementation level, **how the Behringer X32 console
speaks OSC** and **how the reference emulator implements it**, so that the same
emulator can be written in another language.

Sources of truth:

| source | what it is |
| --- | --- |
| `X32.c` v0.88 of the X32-Behringer project (5661 lines) | the reference emulator, "X32 – v0.88 – An X32 Emulator" (Patrick-Gilles Maillot, GPLv3), <https://github.com/pmaillot/X32-Behringer>. **Not part of this repository**: only its behaviour is documented here |
| the `X32*.h` of that project (14 files, ~1.9 MB) | the command tables: 20 694 entries / 221 tables / ~17 200 leaf commands |
| the `X32lib/*.c` of that project | OSC serialisation (`Xsprint`), text→OSC conversion (`Xcparse`), hex/decoded dumps (`Xdump`) |
| the compiled `X32-emulator.exe` of that project | the oracle of the differential tests in §14 (not distributed with this repository) |
| `emulator/` of this repository | the Go port of the emulator, validated byte-for-byte against that oracle |
| `docs/APPENDIX_COMMANDS.md` | generated catalogue: every table, every command, its type and enum list |
| `docs/x32_commands.json`, `docs/x32_enums.json` | the same catalogue in machine readable form |

There is **no custom protocol**: this is plain OSC 1.0 over UDP plus the
Behringer address tree (`/ch/…`, `/node`, `/xremote`, `/meters`, `/-stat/lock`, …).
"Reference emulator" (or "the C emulator") below means `X32.c` /
`X32-emulator.exe` of that project; "the port" means `emulator/` (Go, in this
repository).

---

## 1. Transport and connection model

* **UDP, port 10023** (`X32.c`: `strcpy(Xport_str, "10023")`; 10024 is used by the
  XAir series, so a port flag must stay available).
* The emulator **binds** to `<IP>:10023`, where `<IP>` comes from `-i` or is the
  first IP of the system (`getmyIP()`).
* Replies are sent to the **source address of the request** (`recvfrom()` fills
  `Client_ip_pt`), so both unicast and broadcast discovery work.
* Buffers: 512 bytes (`BSIZE`) for receive and send. The C code writes a
  terminating `\0` at `r_buf[r_len]`.
* Main loop (single threaded): `select(socket, timeout = 10 ms)` → `recvfrom()` →
  dispatch → `Xsend()` → send pending meter frames. Unrecognised packets produce
  **no answer at all**.
* **Clients** (`X32Client[4]`): registered by `/xremote`, identified by the pair
  (UDP port, IP) — the C code compares `sockaddr.sa_data` with `strcmp()`, which
  in practice compares the network-order port followed by the IP address.
* Each client holds an expiry time: `time(NULL) + 11` seconds (`XREMOTE_TIME`).
  A client that does not re-send `/xremote` within ~11 s stops receiving
  notifications and its slot may be reused. X32-Edit and Mixing Station send
  `/xremote` every 10 s.
* At most 4 clients; extra registrations are silently dropped.

## 2. OSC encoding on the wire

The emulator implements exactly this subset:

* **Address**: `"<addr>\0"` padded with `\0` to a 4 byte boundary, i.e. padding
  length `pad4(len(addr) + 1) = (len(addr) + 4) & ~3`.
* **Type tag string**: optional. When present: `"," + tags + "\0"` + padding.
  When **absent altogether** the message is a *GET* ("old" OSC notation used by
  many clients, e.g. `"/ch/01/mix/fader"` = 20 bytes with no tag string).
  An **empty** tag string (`","` followed by three NULs) is also accepted and
  means "no arguments" — again a GET.
* Supported tags: **`i`**, **`f`**, **`s`** and (partially) **`b`**:
  * `i`, `f`: 4 bytes, **big endian**;
  * `s`: NUL terminated string, padded to 4 bytes;
  * `b`: big-endian 32 bit size, then that many bytes, padded to 4 bytes
    (used by meter frames, §7.6; blob *replies* to GET requests are not
    implemented — `X32.c` marks them "todo").
* Not supported: bundles, 64 bit types (`d`, `h`, `t`), `T`/`F`/`I`/`N`, several
  messages in one datagram.
* Number parsing (`Xr_float()`) also accepts the `12k5` notation (`12k5 == 12500`)
  as well as plain `12` and `12.5`.

Byte-level example (this trace is in the project README and is reproduced by the
Go port unit tests):

```
request   -> X,   20 B: /ch/01/mix/fader~~~~
answer    X ->,   28 B: /ch/01/mix/fader~~~~,f~~[1.0000]
request   -> X,   28 B: /ch/01/mix/fader~~~~,f~~[0.5000]
```

* request: 16 address bytes + NUL + 3 padding = 20 bytes (GET, old notation);
* answer: 16 + NUL + 3 padding = 20 bytes, then `,f\0\0` = 4 bytes, then the
  float = 28 bytes.


## 3. Dispatch: headers, subsystem handlers, tables

Dispatch happens in three steps.

### 3.1 Header table (`Xheader[]`, 37 entries)

The **first four bytes** of the message are compared (as a 32 bit integer) with
the table below. The comparison is byte-exact: the `/` command needs three NUL
bytes after the slash.

| key | handler | key | handler |
| --- | --- | --- | --- |
| `/shu` | `/shutdown` (non Behringer) | `/hea` | `/headamp/…` |
| `/inf` | `/info` | `/met` | `/meters` |
| `/xin` | `/xinfo` | `/-ha` | misc (`/-ha/…`) |
| `/sta` | `/status` | `/ins` | misc (no entries) |
| `/xre` | `/xremote` | `/-sh` | `/-show/…` |
| `/nod` | `/node` | `/ren` | `/renew` (silent) |
| `"/\0\0\0"` | `/` (string-path command) | `/cop` | `/copy` |
| `/con` | `/config/…` | `/add` | `/add` |
| `/mai` | `/main/…` | `/loa` | `/load` |
| `/-pr` | `/-prefs/…` | `/sav` | `/save` |
| `/-st` | `/-stat/…` | `/del` | `/delete` |
| `/-ur` | `/-urec/…` | `/uns` | `/unsubscribe` |
| `/ch/` | `/ch/NN/…` | `/-us` | misc (`/-usb/…`) |
| `/aux` | `/auxin/NN/…` | `/und` | `/undo` (dummy, prints) |
| `/fxr` | `/fxrtn/NN/…` | `/-ac` | `/-action/…` |
| `/bus` | `/bus/NN/…` | `/-li` | `/-libs` (silent) |
| `/mtx` | `/mtx/NN/…` | `/sho` | `/showdump` |
| `/dca` | `/dca/N/…` | | |
| `/fx/` | `/fx/N/…` | | |
| `/out` | `/outputs/…` | | |

Addresses matching no key are ignored silently. That is why `/batchsubscribe`,
`/formatsubscribe` and `/subscribe` (very common in clients) produce **no reply**:
they are not implemented at all.

### 3.2 Subsystem handler: exact address lookup

Each handler searches its table for an **exact address match** (`strcmp`) and then
calls the common parameter engine `funct_params()` (§6).

Numbered subsystems extract their index from fixed byte offsets of the packet, so
the address shape is part of the protocol:

| subsystem | index bytes | mapping |
| --- | --- | --- |
| `/ch/NN/…` | 4,5 | `NN` − 1 → `Xchannelset[0..31]` |
| `/auxin/NN/…` | 7,8 | `NN` − 1 → `Xauxinset[0..7]` |
| `/fxrtn/NN/…` | 7,8 | `NN` − 1 → `Xfxrtnset[0..7]` |
| `/bus/NN/…` | 5,6 | `NN` − 1 → `Xbusset[0..15]` |
| `/mtx/NN/…` | 5,6 | `NN` − 1 → `Xmtxset[0..5]` |
| `/fx/N/…` | 4 | `N` − 1 → `Xfxset[0..7]` |
| `/headamp/NNN/…` | 9,10,11 | `NNN` → `Xheadmpset[0..127]` (**0 based**, no −1) |

### 3.3 Node table (`Xnode[]`)

`/node` and the string-path `/` command first map the path to a table through a
**prefix** table (order matters: specific paths first):

```
conf, main, -pre, -sta,
ch/01 … ch/32, ch, auxin/01 … auxin/08, auxin, fxrtn/01 … fxrtn/08, fxrtn,
fx/1 … fx/8, fx, bus/01 … bus/16, bus, mtx/01 … mtx/06, mtx, dca,
outputs/main/01, outputs/main, outputs, headamp/000 … headamp/127, headamp,
-ha, -usb, undo, -action, -show/showfile/snippet, -show/showfile/scene, -show,
-urec, -libs/fx, -libs/r, -libs
```

Entries written without an index (`ch`, `fx`, `-pre`…) refer to the first table of
the series (used by generic queries).

## 4. Command tables: types, flags, enum lists

```c
typedef struct X32command {
    char* command;              // "/ch/01/mix/fader"
    union { int typ; char* str; } format;   // value/command type
    int flags;                  // F_GET, F_SET, F_XET, F_FND (F_NPR is unused)
    union { int ii; float ff; char* str; void* dta; } value;  // live value
    char** node;                // enum string list (values start with a space)
} X32command;
```

* **The tables are the console state.** Reading a parameter reads the current
  value of that entry; writing stores into the same entry.
* Flags: `F_GET 0x1`, `F_SET 0x2`, `F_XET = F_GET|F_SET`, `F_FND 0x8`
  (group header / first of a series), `F_NPR 0x4` (declared, never used).
* Formats and how the value is stored:
  * `I32` int, `F32` float (levels are normalised 0…1), `S32` string, `E32` enum
    (stored as the int index into the enum list), `P32` bitmap/int, `B32` blob
    (unused by the tables), `FX32` "type depends on the effect in the slot" (§11);
  * **composed types** used by group headers (`CHCO`, `CHDE`, `CHPR`, `CHGA`,
    `CHGF`, `CHDY`, `CHDF`, `CHIN`, `CHEQ`, `CHMX`, `CHMO`, `CHME`, `CHGRP`,
    `CHAMIX`, `AXPR`, `BSCO`, `MXPR`, `MXDY`, `MSMX`, `FXTYP1/2`, `FXSRC`,
    `FXPAR1/2`, `OMAIN`, `OMAIN2`, `OP16`, `OMAIND`, `HAMP`, `PREFS`, `PIR`,
    `PIQ`, `PCARD`, `PRTA`, `PIP`, `PADDR`, `PMASK`, `PGWAY`, `PKEY`, `STAT`,
    `SSCREEN`, `SCHA`, `SMET`, `SROU`, `SSET`, `SLIB`, `SFX`, `SMON`, `SUSB`,
    `SSCE`, `SASS`, `SSOLOSW`, `SAES`, `STAPE`, `SOSC`, `STALK`, `USB`, `SNAM`,
    `SCUE`, `SSCN`, `SSNP`, `HA`, `ACTION`, `UREC`, `SLIBS`, `D48`, `D48A`,
    `D48G`, `UROUO`, `UROUI`, `OFFON`, `CMONO`, `CSOLO`, `CTALK`, `CTALKAB`,
    `COSC`, `CROUTSW`, `CROUTIN`, `CROUTAC`, `CROUTOT`, `CROUTPLAY`, `CCTRL`,
    `CENC`, `CTAPE`, `CMIX`). They carry no single value: they drive the `/node`
    renderer (§8) and the string-path `/` command (§9).
* **Enum lists** are string arrays whose entries start with a space and end with
  an empty string, e.g. `{" OFF"," ON",""}`:
  * `/node` output appends the string **with** its leading space;
  * parsing compares with the leading space stripped (`node[j]+1` in `X32.c`), so
    clients send `ON`, never ` ON`.
* Group headers keep the number of their items in `value.ii` (`n=` in
  `docs/APPENDIX_COMMANDS.md`); that count is used by the generic
  `OFFON`/`SSOLOSW`/`CROUTAC`/`CENC` renderers only.

Table statistics (measured, `docs/x32_commands.json`):

```
tables 221   entries 20 694   enum arrays 149
leaf commands: 17 168   ->  I32 6526  F32 3468  S32 1799  E32 4005  P32 950  FX32 512
7218 /-show/...   1804 /-libs/...   5120 /ch/...   1584 /bus/...   912 /auxin/...
880 /fxrtn/...    635 /-prefs|/-stat|/-action|/-urec   517 /config+/main
414 /mtx/...      385 /headamp/...  311 misc        301 /outputs/...  57 /dca/...
```

## 5. Console state

* All values live in the tables (RAM only). A fresh process starts with **all
  values zero / empty strings**, unless a resource file is loaded (§13).
* There is no audio engine: meters are all zero, FX work on normalised values,
  scenes/snippets have no payload.
* Channels 1…32 use **32 independent tables**, all with the same 160 entry layout
  (config, grp, preamp, delay, insert, gate, gate filter, dyn, dyn filter, 4 EQ
  bands, 16 sends, automix); the same applies to buses (16 × 99), aux inputs
  (8 × 114), FX returns (8 × 110), matrix (6 × 69), DCA (8 × 7), head amps
  (128 × 3 or 4) and FX slots (8, see §11).
* A group header (flag `F_FND`) is stored in the table as the entry `i`; its data
  entries are `i+1 …` (composed types know how many they use).

## 6. Parameter engine: GET, SET, change detection, notifications

`funct_params(table, i)` is the heart of the emulator. `i` is the index of the
entry whose address matched the request.

### 6.1 Which form is it?

```
tagBlock = ((len(addr) + 4) & ~3) + 1     // offset of the first tag character
SET if: (len(packet) - 4 > len(addr)) && packet[tagBlock] != 0
GET otherwise
```

So a request with no tag string at all, or with an empty one (`","`), is a GET.

### 6.2 GET

1. If the matched entry is a group header (`F_FND`), advance to the first leaf
   (`i++`) — this is how `"/ch/01/config"` answers with `/ch/01/config/name`.
2. If the entry has `F_GET`, build the answer:
   * `I32`, `E32`, `P32` → `"<addr>" + ",i" + int32`
   * `F32` → `"<addr>" + ",f" + float32`
   * `S32` → `"<addr>" + ",s" + string` (an unset string is a single NUL byte,
     i.e. an empty string, not `NULL`)
   * `FX32` → the tag is decided by the effect type in the slot (§11)
   * `B32` → not implemented
   * composed types → only the address is sent (no value)
3. The answer goes **only to the requesting client** (`S_SND`).

### 6.3 SET

1. Skip a group header (`i++`), require `F_SET`, then read the tag string:
   `nTags` characters starting at `tagBlock`; the data block starts at
   `pad4(tagBlock + nTags + 1)`.
2. For every tag, consume the next value and store it into **the next table
   entry** (`i` is incremented per parameter — that is why a multi parameter SET
   is a sequence of adjacent commands, e.g. `/ch/01/config ,siii "name" icon
   color source`):
   * `i` → int32 (big endian)
   * `f` → float32
   * `s` → NUL terminated string (empty string clears the value)
   * `b` → accepted and ignored
3. **Change detection**: a value is considered changed if
   * int: `new != old`,
   * float: `new != old` (exact comparison here; the `/` command setters use a
     `1e-4` epsilon instead),
   * string: different content, or clearing a previously set string.
4. If at least one value changed:
   * the **original request packet is echoed verbatim** to all registered
     `/xremote` clients **except the sender** (`S_REM`),
   * the sender receives **nothing** (this is what real consoles do as well),
   * and the *whole* packet is echoed, even if only one parameter of a multi
     parameter SET changed.
5. If nothing changed, the emulator answers **nothing at all** (no ACK, no echo).
   Clients must be written accordingly.

### 6.4 Client notifications

* Notifications go to the clients registered through `/xremote` whose slot has not
  expired (11 s).
* The sender of the change never receives its own echo.
* Two different notification shapes exist, and both are needed for compatibility:
  1. **direct command** (`/ch/01/mix/fader ,f 0.5`) → the raw request packet is
     echoed (as described in §6.3), one message can carry several parameters;
  2. **string-path `/` command and the group setters** (§9) → one single-parameter
     message per changed value:
     `"<full address>" + ",i|f|s" + value`, e.g.
     `/ch/01/mix/fader\0\0\0\0,f\0\0` + float.
* The port also notifies on shape (1) for string-path single-parameter sets, which
  the C reference forgets to do (deviation §14.2) — a real console notifies on
  every change.


## 7. Life cycle commands (answers are byte exact)

### 7.1 `/info`

```
/info\0\0\0,ssss\0\0\0 "V2.07" "<console name>" "X32" "4.06"
```

The four strings are: `"V2.07"`, the console name (`/-prefs/name`, default
`"X32 Emulator"`), `"X32"` and the firmware version the emulator announces
(`"4.06"`, `XVERSION` in `X32.c`). The answer goes to the requester only.

The name is the one stored in `/-prefs/name` (resource file or a client renaming
the console); when it is empty, the reference reports the hard coded
`"X32 Emulator"` while this port reports `"REAPER"` (its own default, `-name`
changes it — deviation §14.2).

### 7.2 `/xinfo` (discovery)

Same layout, but the first string is the console **IP address**:

```
/xinfo ,ssss  "<ip>" "<name>" "X32" "4.06"
```

`X32_Command`, X32-Edit and Mixing Station use `/xinfo` to find a console, so this
command is mandatory for discoverability.

### 7.3 `/status`

```
/status ,sss  "active" "<ip>" "<name>"
```

Note: `X32.c` fills `<ip>` with `getmyIP()` (the first host address), not with the
bound address; the port uses the bound address instead (deviation §14.2, visible
only when the emulator binds something else than the host address).

### 7.4 `/xremote`

Registers (or refreshes) the sender as a notification client for 11 seconds. It
**never answers**. Clients send it every 10 s.

### 7.5 `/unsubscribe`, `/renew`, silence

* `/unsubscribe` removes the sender from the client list, no answer.
* `/renew` produces no datagram at all (the emulator returns an empty buffer).
* `/batchsubscribe`, `/formatsubscribe`, `/subscribe`: not implemented, no answer.
* Unknown addresses: no answer.
* `/-libs/…`: no answer (the tables exist but the handler does nothing).
* `/undo`: handler exists and only prints a message; no answer.
* `/shutdown` (non-Behringer extension) and `/-stat/lock ,i 2`: save the resource
  file (§13) and terminate the process; no answer.

### 7.6 `/meters`

Request (as X32-Edit/Mixing Station send it):

```
/meters\0\0\0\0,siii\0\0\0 "/meters/<n>"\0\0\0\0 <int> <int> <timefactor>
```

* the emulator subscribes the **sender** to meter group `<n>` (0…16) for **10 s**;
* the frame is then sent every `50 ms × timefactor` (a factor < 1 or > 99 is
  clamped to 1);
* one client per meter group (the last requester wins);
* frame layout:

```
"/meters/<n>" + pad   ",b\0\0"   BE int32 blob size = (l+1)*4   LE int32 l   4*l value bytes
```

| group | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| values `l` | 70 | 96 | 49 | 22 | 82 | 27 | 4 | 16 | 6 | 32 | 32 | 5 | 4 | 48 | 80 | 50 | 48 |

  Group 15 is the RTA (16 bit samples, divided by 256 by `Xdump.c`), group 16 the
  gain reduction group (also 16 bit). The blob starts with a **little-endian**
  count word, then the values, which are **little-endian 32 bit floats** on the
  scale of §12.1 (`0` = -oo, `1` = 0 dBFS, up to `8.0` = +18 dB). See
  `X32lib/Xdump.c` for the reference decoder; `X32Automix.c` reads
  `"/meters/1"` as the channel levels.
* frame size is `24 + 4*l` bytes (408 for `/meters/1`, 224 for `/meters/15`);
* the reference emulator sends **zeros** for every value (no audio); a useful port
  keeps the same sizes and timing and only fills the values;
* time factor offset: `packet[k+24]` with `k = 0`, `k = 12` for `/meters/5` and
  `k = 8` for `/meters/6` (`0.88` fix). For all other groups that offset lands in
  the address padding, so the factor becomes 0 → clamped to 1.

#### Value layout of the groups

The values of a group are **block-wise**: the level of element *n* of a family
sits at index *n* of its block, then the following blocks carry the other values
of the same elements (gain reductions, which an emulator cannot know and which
stay at zero). `X32Automix.c` proves it for the channels: it reads

```c
if (strncmp(r_buf, Meters + 16, 9) == 0)                   // "/meters/1"
    for (ch = chstart - 1; ch < chstop; ch++) {
        ff = (float *)(r_buf + 24 + 4*ch);                 // channel ch level
        if (*ff > X32sensitivity) { ... }
```

so the level of channel `ch` (0 based) is the value at index `ch`, i.e. the
**first 32 values** of `/meters/1` are the 32 channel levels and the remaining 64
are the two gain reduction blocks.

| group | values | content | confidence |
| --- | --- | --- | --- |
| 0 | 70 | 32 channels, 8 aux inputs, 8 effect returns, 16 buses, 6 matrices (32+8+8+16+6) | exact sum |
| 1 | 96 | 32 channel levels, then 32+32 gain reductions | **verified** (`X32Automix.c`) |
| 2 | 49 | 16 bus levels, then gain reductions and one spare (16×3+1) | inferred |
| 3 | 22 | 8 aux inputs, 8 effect returns, 6 matrices (8+8+6) | exact sum |
| 4 | 82 | unknown, left at zero | — |
| 5 | 27 | 8 aux input levels, then the main bus (8×3+3) | inferred |
| 6 | 4 | main bus L, R, mono (+1 spare) | inferred |
| 7 | 16 | 16 bus levels | inferred |
| 8 | 6 | 6 matrix levels | inferred |
| 9, 10 | 32 | 32 channel levels | inferred |
| 11, 12 | 5, 4 | main bus | inferred |
| 13 | 48 | 16 bus levels, then gain reductions (16×3) | inferred |
| 14 | 80 | every fader: 32 ch, 8 aux, 8 fx returns, 16 buses, 6 matrices, main L/R, 8 DCA (no meters) | exact sum |
| 15 | 50 | RTA: 100 × 16 bit samples | `Xdump.c` |
| 16 | 48 | gain reductions: 96 × 16 bit samples | `Xdump.c` |

### 7.7 `/-stat/userpar/NN/value` — user assign controls

A real console reports every User Assign control through

```
"/-stat/userpar/NN/value" ,i <0…127>
```

with `NN` numbered as on the hardware: **buttons** `1…8` (bank A), `9…16` (B),
`17…24` (C), `25…32` (D) and **encoders** `33…40`. Control surface bridges such as
`X32ReaperW` use the Bank C group as a transport remote: they assign names to the
controls (`/config/userctrl/C/…`, §6), initialise them
(`/-stat/userpar/17…24 = 0`, `/-stat/userpar/33…36 = 64`), then listen for the
console's own reports.

Semantics (important, and not symmetric):

* the value is **pushed by the console**, not polled: it is the physical control
  that generates the message — `127` when the button goes down, `0` when it comes
  back up (`X32ReaperW` acts on the **release**, `== 0`, "button up transition");
* an encoder reports its absolute position `0…127`, centered = `64`;
* a GET of the same address returns the current value as `,i`.

The **reference emulator has no such addresses** (they are missing from its
tables, so a SET is silently ignored) — §14.2 deviation 6. The port implements the
Bank C range (17…24 and 33…36) as a hand written, non-persisted table, so that
`x32emu` can act as the desk side of such a bridge; `PressUserPar()` generates the
physical events (a GUI, a web front end or a script has to, since an emulator has
no buttons).

Byte layout of the notification — clients parse fixed offsets, so nothing may
change here:

```
"/-stat/userpar/18/value" (23 bytes) NUL pad   "," "i" NUL   BE int32
 0                    22   23   24 25  26 27   28       31
```

i.e. a 32 byte packet with the integer at offset **28** in every case (the address
is always 23 characters long for a two digit index).


## 8. `/node` — bulk state read

`/node ,s "<path>"` has two completely different behaviours.

### 8.1 Group query

If `<path>` matches a group header (first `F_FND` entry whose address *starts
with* the path, in table order), the answer is:

```
"node\0\0\0\0,s\0\0" + "/<group path> <value1> <value2> …\n" + "\0" + padding
```

One string argument with the group path followed by all its values in the order
defined by the renderer for that composed type, terminated by a line feed.

Rendering helpers (all of `X32.c`):

| helper | output |
| --- | --- |
| `Slevel(f)` | `" -oo"` for 0, else `" %+.1f"` dB using the 4 segment curve of §12 |
| `Slinf(f,min,max,dec)`, `Slinfs(...)` | `min + (max-min)*f`, `" %.Nf"` / `" %+.Nf"` |
| `Slogf(f,min,max,dec)` | `exp(f*ln(max/min) + ln(min))`, `" %.Nf"` |
| `Sint(i)` | `" %d"` |
| `Sbitmp(i,bits)` | `" %"` followed by `bits` bits, MSB first |
| on/off | `" ON"` / `" OFF"` |
| strings | `" \"text\""` (quoted); empty → `" \"\""` (a few renderers use `" \" \""`) |
| enums | the raw enum string, i.e. **with** its leading space (`" OFF"`) |

The complete per-type order is transcribed in the port in
`emulator/internal/x32/nodetypesa.go`, `nodetypesb.go`, `nodetypesc.go`. Two
details matter: some renderers skip an entry (`MXDY` skips `i+13`, `MSMX` skips
`i+3`) — clients parse positionally; and `CHEQ`/`CHGF`/`CHDF`/`CHPR` use
precomputed display tables (`f101`, `f121`, `f201`) that live in `X32.c`.

### 8.2 Single parameter query

If `<path>` designates a leaf (or a group path that is not found), the emulator
rewrites the request into the equivalent direct command
(`/node ,s -prefs/rta/visibility` → `/-prefs/rta/visibility`), lets the normal
engine answer it, and wraps that answer:

```
"node\0\0\0\0,s\0\0" + "/<address>" + <value> + "\n" + "\0" + padding
```

with `<value>`:

* enum entry (when the command has an enum list) → the raw string with its
  leading space (`" OFF"`);
* `I32` → `" %d"`; `F32` → `" %f"` (6 decimals); `P32` → `" %"` + bits from the
  highest set bit down to bit 0;
* `S32` → the string **without** any separator (reference quirk:
  `node ,s /ch/01/config/nameMyChannel`);
* any other type → `" TODO"`.

### 8.3 `/showdump`

Any address starting with `/sho` (clients send `/showdump`) answers with

```
node ,s  "/-show/showfile/show \"<show name>\" 0 0 0 0 0 0 0 0 0 0 \"4.06\""
```


## 9. The `/` command (string path)

Address `/` with one string argument: `"<path> [<value> [<value> …]]"`, e.g.

```
/ ,s  "/ch/01/mix/fader 0.5"
/ ,s  "/ch/01/config/name \"Lead\""
```

Semantics:

1. The path is looked up in the node table (prefix), then matched **exactly**
   against the addresses of that subsystem (address without its leading `/`).
2. If the matched entry is a **leaf**, exactly one value is consumed according to
   the entry type: `I32`/`P32` int, `E32` enum name, `F32` float, `S32` string
   (`"…"`, `'…'`, or an unquoted token that stops at the first space).
3. If the matched entry is a **group header**, the emulator walks to the last
   consecutive `F_FND` entry and applies the values with the parser chain of that
   composed type (table below). Every parser consumes exactly one value.
4. The **whole request is echoed back to the sender** (unlike direct commands),
   and every changed parameter is pushed to the other `/xremote` clients as
   described in §6.4(2).
5. Failure handling (reference behaviour): if the **first** value cannot be
   parsed, the command is answered with nothing at all; if a **later** value
   fails, the already parsed values are applied and the request is still echoed.

Value parsers (`XslashSet*` of `X32.c`):

| parser | meaning |
| --- | --- |
| `XslashSetInt` | plain integer |
| `XslashSetPerInt` | integer or `%0101` bitmap |
| `XslashSetList` | enum name (the leading space of the table entry is ignored) |
| `XslashSetLinf(min, range, step)` | physical value → `(v-min)/range`, rounded to `step` |
| `XslashSetLogf(min, ln(max/min), steps)` | physical value → `ln(v/min)/ln(max/min)`, rounded to `steps` |
| `XslashSetLevl(steps)` | dB → normalised with the 4 segment level curve, rounded to `steps` |
| `XslashSetString` | `"…"`, `'…'` or unquoted token |

Composed types and their parameter chains (exact numbers from `X32.c`; the port is
in `emulator/internal/x32/slash.go` + `slash2.go`):

```
CHCO name icon colour input                 CHDE on time(0.3..500 step .1)
CHPR trim(-18..36 .25) invert hpon hpslope hpf(log 20..400, 100 steps)
CHGA on mode thr(-80..80 .5) range(3..57) attack(0..120) hold(log .02..2000,100) release(log 5..4000,100) keysrc
CHGF/CHDF on type f(log 20..20000, 200)     CHIN on pos sel
CHEQ type f(log 20..20000,200) g(-15..30 .25) q(log 10..0.3, 71)
CHDY on mode det env thr(-60..60 .5) ratio knee mgain attack hold release pos keysrc mix auto
CHMX on level(1023) pan panFollow tap tapFollow    CHMO on level(1023) pan type panFollow
CHME on level(1023)                        CHGRP %dca %mute
CHAMIX group weight(-12..24 .5)            AXPR trim(-18..36 .25) invert
BSCO name icon colour                      MXDY 13 values (like CHDY minus 2)
MXPR invert                                MSMX on level(1023) pan(-100..200 step 2)
FXTYP1/2 type                              FXSRC two enum values
FXPAR1/FXPAR2 effect parameters (§11)      OMAIN src pos invert / OMAIN2 src pos
OP16 3 enums + int                         OMAIND on time(0.3..500 .1)
HAMP gain(-12..72 .5) phantom              PREFS name + 24 values
PIR 4 / PIQ 3 / PCARD 13 / PRTA 11 / PIP 1 / PKEY 1 / PADDR|PMASK|PGWAY 4 ints
STAT 24 values                             SSCREEN 3 / SCHA|SMET|SROU|SSET|SLIB|SFX|SMON|SUSB|SSCE|SASS 1 enum
SSOLOSW 1 + 79 enums                       SOSC 1 / STALK 2 / USB 2 strings
SNAM 11 / SCUE 9 / SSCN 4 / SSNP 6         UREC 16 / SLIBS 5
D48 2 / D48A 48 ints / D48G 12 names       UROUO 48 ints / UROUI 32 ints
```

## 10. Library, scene and file commands

### 10.1 `/copy ,siii "libchan" <src> <dst> <mask>`

Only the channel library is implemented (like the reference). The mask selects
blocks, which are copied **inside the channel tables** (indices are table
positions, not user visible parameters):

| mask | block | channel table indices |
| --- | --- | --- |
| 0x01 | head amps | 8 … 19 |
| 0x02 | config | 0 … 7 |
| 0x04 | gate | 20 … 32 |
| 0x08 | dynamics | 33 … 52 |
| 0x10 | EQ | 53 … 74 |
| 0x20 | sends/mixes | 75 … 145 |

Answer: `/copy ,si "libchan" 1` (or `0` when `src`/`dst` are out of range, or when
the type is not `libchan`, which then produces no answer at all).
The reference copies only the int member of its value union, which corrupts string
entries (names); the port copies according to the type (deviation §14.2).

### 10.2 `/save ,siss "<type>" <number> "<name>" "<note>"`

* `scene`: writes `/-show/showfile/scene/NNN/name` and `.../notes`, sets
  `.../hasdata` to 1, notifies the other clients for each changed entry and
  answers `/save ,si scene 1`; an unknown scene number → `/save ,si scene 0`.
  **Scene payloads are not stored** ("hasdata" is only a flag).
* `snippet`: same with `/-show/showfile/snippet/NNN/name` + `hasdata`
  (answer `/save ,si snippet 1`).
* `libchan`: nothing is saved, answer `/save ,si libchan 1`.
* any other type: no answer.

### 10.3 `/delete ,si "<type>" <number>`

Clears the scene/snippet name, notes and `hasdata`, notifies the clients and
answers `/delete ,si scene|snippet 1` (or `0` when not found). `libchan` is a stub
that always answers 1.

### 10.4 `/add`, `/load`

Stubs. `/add ,sis "cue" …` answers `/add ,si cue 1` (only for `cue`), `/load …`
answers `/load ,si libchan 1`. Nothing is added or loaded.

### 10.5 Not implemented

`/batchsubscribe`, `/formatsubscribe`, `/subscribe`, `/-libs/…`, `/undo`,
`/-action/…` (values are stored and propagated like normal parameters, but no
action is executed — e.g. `/-action/goscene` does **not** recall a scene),
`/renew`, RTA data, recorder, X-Live.


## 11. FX subsystem

* 8 slots: `/fx/1 … /fx/4` use the `FXTYP1`/`FXPAR1` layouts, `/fx/5 … /fx/8` the
  `FXTYP2`/`FXPAR2` ones.
* `/fx/N/type` (`E32`): the effect loaded in the slot, from `Sfxtyp1[]` (60
  entries: `HALL, AMBI, RPLT, …, PIT`) or `Sfxtyp2[]` (34 entries:
  `GEQ2, TEQ2, GEQ, …, SUB`).
* `/fx/N/par/PP`, PP = 01…63: 63 normalised parameters. The **type of each
  parameter** (int / float / enum / string) is given by the string
  `Sflookup[slot type]` (slots 1…4) or `Sflookup2[slot type]` (slots 5…8): one
  character per parameter, `-` for unused. That is `FXc_lookup()` in `X32.c`: a
  GET answers `,i` or `,f` depending on that character, so a client can resolve
  the type without knowing the effect.
* The physical scaling of each parameter (linear/log/level/enum per effect type)
  lives in `SetFxPar1()` / `GetFxPar1()` in `X32.c` (~1500 lines of `switch`).
  The Go port consumes the number of parameters given by the lookup string and
  stores the normalised values; porting the per-type scales is the main remaining
  task for full parity (§14.3).
* `/node ,s fx/N` renders the slot (type, source and all parameters); the port
  renders normalised values until the scales are ported.

## 12. Value scales

### 12.1 Levels (faders, sends)

Normalised value `v ∈ [0,1]` ↔ dB, 4 linear segments:

| segment | range | formula |
| --- | --- | --- |
| 1 | `v ≤ 0.0625` | `dB = v/0.0625*30 - 90` (−∞ … −60; `v = 0` prints `-oo`) |
| 2 | `v ≤ 0.25` | `dB = (v-0.0625)/(0.25-0.0625)*30 - 60` |
| 3 | `v < 0.5` | `dB = (v-0.25)/(0.5-0.25)*20 - 30` |
| 4 | `v ≥ 0.5` | `dB = (v-0.5)/(1-0.5)*20 - 10` |

Input direction (`XslashSetLevl`) is the inverse, quantised on `steps` (1023 for
faders/sends, 160 for one send variant) with `f = int(f*steps + 0.5)/steps`; values
below −60 dB are clamped to ≥ 0 and `-oo` maps to 0.

### 12.2 Logarithmic and linear mapping

```
normalised = ln(v / min) / ln(max / min)      (rounded to `steps`)
normalised = (v - min) / (max - min)          (rounded to `step`)
```

Both clamp to `[0,1]` and force `0` instead of `-0.0`.

### 12.3 Bitmaps

`P32` values are rendered by `/node` as `" %0110"` (MSB first). When parsing, `%…`
is decoded as a bitmap; the reference has a bug here (it shifts the old value
instead of the accumulator and keeps bit 0 only). The port implements the inverse
of `Sbitmp()` (deviation §14.2), which is what clients and real consoles expect.

## 13. Persistence (`.X32res.rc`)

* The non-Behringer command `/shutdown` (or `/-stat/lock ,i 2`) writes
  `.X32res.rc` in the working directory and terminates the emulator; the file is
  read at start-up (`X32Init`).
* Format, in a **fixed table order** (`Xconfig, Xmain, Xprefs, Xstat`, all 32
  channels, auxin, fxrtn, bus, mtx, dca, fx, outputs, headamp, misc, urec, libsc,
  libsr, libsf):

```
<type> <int32>\n                 for every non string entry
<type> <length> <string>\n       for S32 entries (spaces are kept)
```

  Floats are stored as **the bit pattern of the float reinterpreted as int32**
  (`fprintf("%d", value.ii)` in C) and read back the same way, so the round trip
  is exact.
* **Not** persisted: show files, scenes, snippets, `/-action`, meters, `/-show`.
* Without the file all values keep their zero defaults, and the main program
  copies its own IP address into `/-prefs/ip/addr/0..3`.
* The port keeps the same format (and adds nothing), so resource files are
  interchangeable between the C and the Go emulator.


## 14. Reference emulator vs. real console vs. Go port

### 14.1 Differential test (how the port was validated)

`tools/difftest.ps1` starts the reference `X32-emulator.exe` (which is *not* shipped
with this repository: build it from the upstream C sources, or point the script at a
prebuilt binary with `-Oracle <path>`; without one the script skips itself with an
explanation) and then `build/x32emu.exe` alone on `127.0.0.1:10023`, and drives both
with the same scripted session
(`emulator/cmd/x32probe`, ~130 steps: discovery, GETs of every subsystem, SETs with
and without changes, `/node` group and single queries, `/` commands, `/copy`,
`/save`, `/delete`, `/add`, `/load`, `/showdump`, meter subscriptions, `lock`,
`/unsubscribe`, unknown commands). Every datagram is captured in hex plus its
decoded dump, then the two logs are compared line by line.

Result on the shipped corpus: **435 lines captured from the C emulator, 447 from the
port, 14 protocol differences**, all explained by the intentional deviations below —
every other GET reply, `/node` text, echo, meter frame and silence is byte
identical (13 of the lines are answers the port adds, 1 is the `/status` IP field,
where the two emulators print different addresses by design).

`tools/difftest.ps1` separates the asynchronous meter frames from the protocol
differences (a `/meters` frame is delivered on a 10 s window instead of answering
a request, so one frame more or less is a timing artifact): a run reports 14
protocol differences, and sometimes one extra meter frame.

The client that receives notifications re-registers `/xremote` every 20 steps:
a console drops a client that has been silent for 11 s, so without this refresh
the last steps of the session would be compared in a state where the subscription
has already expired (while the corpus was young this made the number of differing
lines fluctuate between 9 and 10). The 11 s timeout itself is verified separately
with a timed test.

### 14.2 Intentional deviations of the port

| # | topic | reference behaviour | port behaviour | why |
| --- | --- | --- | --- | --- |
| 1 | `/status` IP field | `getmyIP()`: the first host address (may differ from the bound one) | the bound address | reports the address that actually answers |
| 2 | notification for a string-path single parameter set (`/ ,s "/ch/01/mix/fader 0.5"`) | applied and echoed to the sender, but **not** pushed to other `/xremote` clients | also pushed as a one-parameter message | a real console notifies every change; clients stay in sync |
| 3 | `%bitmap` parsing | buggy (keeps bit 0 of the last character) | correct inverse of `Sbitmp()` | clients send real bitmaps |
| 4 | `/copy libchan` | copies the int member of every entry (corrupts names) | copies according to the entry type | keeps names intact |
| 5 | setter/renderer off-by-ones (`AXPR` invert written into the trim entry, `PIQ` third value written into the next block, `SSCREEN` writing three times into the same entry) | buggy | sane parameter targets | the reference corrupts neighbouring parameters |
| 6 | `/-stat/userpar/NN/value` (user assign controls, §7.6) | the address is absent from the tables: a SET is silently ignored, nothing can be read | implemented for the Bank C controls (buttons 17…24, encoders 33…36): `,i` replies and console-like notifications | it is how control surface bridges (X32ReaperW) receive transport and marker commands from a desk; without it an emulator cannot drive REAPER |
| 7 | console name when `/-prefs/name` is empty (§7.1) | hard coded `"X32 Emulator"` | `"REAPER"` (constant `DefaultName`), `-name <name>` changes it | this port usually *is* the console of the REAPER bridge, and clients display that name; a name stored in `/-prefs/name` (resource file or a client) still wins |
| 8 | default verbosity | `-v 1`: every datagram is logged | `-v 0`: silent unless asked | the port also runs as a library (and as the daemon `x32reaper`); `-v 1` restores the reference behaviour |

Two differences were **fixed** instead of being kept, once the differential test
became deterministic: the port now notifies the other clients for the scene
`safes` and snippet `eventtyp` flags on `/save` and `/delete` (X32.c does it with
its "hasdata" comment pointing at the wrong index), and for a single parameter
changed through a `/ ,s` command (deviation 2).

### 14.3 Known gaps of the port (to be completed)

* FX parameter **scales** (`SetFxPar1`/`GetFxPar1`, §11): values are stored
  normalised and `/node ,s fx/N` prints normalised values.
* Meter **values** are zeros unless an outside source feeds them: the reference
  (and this port without a feed) reports zeros, which is what keeps the
  differential test byte identical. The `emu` package exposes the feed
  (`SetMeterLevel`, §14.4), `docs/X32_REAPER_BRIDGE.md` §6 describes how REAPER's
  `/track/N/vu[/L/R]` levels reach the frames, and §7.6 lists which groups carry
  levels (`/meters/15`, the RTA, and `/meters/16`, the gain reductions, are always
  zero: they are not levels).
* Blob GET replies (`B32`) are not implemented (same as the reference).
* The user assign table (deviation 6) is not rendered by `/node ,s -stat`: the
  reference has no such address at all and adding it there would break the byte
  parity of the `/node` replies.
* Everything the reference does not implement stays unimplemented:
  `/batchsubscribe`, `/formatsubscribe`, `/-libs`, `/undo`, scene payloads,
  recorder, RTA data, X-Live.

### 14.4 Extensions of the Go port (beyond the reference)

| entry point | purpose |
| --- | --- |
| `x32emu -i 0.0.0.0` (default) | bind every interface while `/info`, `/xinfo` and `/status` advertise the first usable IPv4 of the machine |
| `x32emu -res <file>` | resource file path (`.X32res.rc` by default) |
| `x32emu -name <name>` | console name used when `/-prefs/name` is empty (default `REAPER`); `-name ""` keeps the stored one |
| `x32emu -port <n>` | UDP port (default `10023`); `0` asks the system for a free one |
| `emulator/emu` package | the console as a library: `New(Config)`, `Run`, `Stop`, `Close`, `Address`, `IP`, `Port`, `SetMeterLevel`, `MeterLevel`, `DefaultName/DefaultPort/DefaultResource`, `FirstIPv4` |
| `(*State).SetMeterLevel(kind, i, v)` | store one meter level (`MeterCh/Auxin/Fxrtn/Bus/Mtx/Main`); the frames of every subscribed group then carry it (layout in §7.6). `SetMeterLevels` replaces the whole table, `MeterLevel` reads one back. Safe from another goroutine than the server loop |
| `(*State).PressUserPar(i, v)` | simulate a physical Bank C button/encoder event: store the value and push `/-stat/userpar/NN/value ,i v` to every `/xremote` client (no exclusion, there is no originating client) |
| `(*State).NotifyAll(b)` | push a datagram to every registered `/xremote` client (used by console side events) |
| `(*State).SetName(name)` | set the console name used when `/-prefs/name` is empty |
| `x32probe -userpar` | opt-in extension of the differential corpus for the user assign controls (the C emulator cannot answer them) |
| `x32reaper` (bridge module) | the all-in-one command: this console plus the REAPER bridge in one process, feeding the meter frames from REAPER's `/track/N/vu` |
| `x32listen -meters <n>` (bridge module) | subscribe to a meter group and print the non-zero values of every frame (the tool used to verify the feed) |


## 15. The Go port

```
emulator/
  go.mod                       module x32emu (no external dependencies)
  cmd/x32emu/main.go           the emulator: reference flags (-i -d -x -b -f -r -m) plus -name, -port, -res
  cmd/x32probe/main.go         scripted OSC session used by the differential test
  emu/emu.go                   the console as a library (New/Run/Stop/Address/SetMeterLevel)
  emu/emu_test.go              library tests: real socket, /xinfo, meter feed
  osc/                         OSC encode / decode / dump (osc_test.go)
  internal/x32/
    types.go setters.go command.go dispatch.go slash.go slash2.go
    noderender.go nodetypesa.go nodetypesb.go nodetypesc.go nodes.go
    special.go copysave.go meters.go meters_table.go meterfeed.go
    fx.go persist.go server.go state.go userpar.go
    x32_test.go userpar_test.go meterfeed_test.go
    types_gen.go enums_gen.go tables_gen.go sets_gen.go      <- generated
```

The bridge module adds the all-in-one command:

```
bridge/
  cmd/x32reaper/main.go        X32 emulator + REAPER bridge in one process
  cmd/x32bridge/main.go        the bridge alone (real console, or x32emu started separately)
  cmd/x32listen/main.go        OSC dump tool (also decodes /meters frames with -meters)
  internal/bridge/             the bridge engine (config, bank, x32, reaper, meters)
  internal/cliflags/           the flags shared by the two commands
```

A second module, `bridge/` (`module x32bridge`, `replace x32emu => ../emulator`),
reuses `x32emu/osc` for the X32 <-> REAPER bridge described in
`docs/X32_REAPER_BRIDGE.md`.

Build / run / test:

```bash
# from the repository root:
make build                        # the five commands into bin/
make test                         # unit tests (OSC layer + protocol behaviour)
./bin/x32emu -i 127.0.0.1         # run the emulator

pwsh -File tools/difftest.ps1 -Oracle <X32-emulator.exe>   # differential test
```

Design notes for other languages:

* one table entry = one live value; the tables **are** the console state (in Go
  they are package level, like the C globals — one emulator per process);
* the tables are **generated** from the C headers (`tools/gen_tables.ps1` →
  `tables_gen.go`), never typed by hand: 20 694 entries, 221 tables;
* dispatch = 4 byte prefix table + exact address match + one generic parameter
  engine (`funct_params`), exactly like the reference;
* notifications must reproduce **both** shapes of §6.4, otherwise multi client
  setups (X32-Edit + Mixing Station + control apps) get out of sync;
* `/node` value order matters more than elegance: clients parse it positionally;
* value scales (§12) and the 11 s `/xremote` window are part of the protocol, not
  implementation details.

## 16. Acceptance checklist

Protocol level (covered by unit tests and the differential test):

- [x] `/xinfo`, `/info`, `/status` answered; discovery works.
- [x] GET of a leaf without a tag string, with `,i`/`,f`/`,s`, and with an empty
      tag string.
- [x] SET applies the value, echoes the request to the other `/xremote` clients
      and stays completely silent when nothing changed.
- [x] `/xremote` registration, refresh, expiry (11 s) and `/unsubscribe`.
- [x] `/node` group replies for every implemented composed type, byte identical.
- [x] `/node` single parameter replies (including the `S32` no-separator quirk).
- [x] `/` commands: leaves and groups, scales, quoting, echo, silence on a bad
      first value.
- [x] `/meters` subscription, frame layout and sizes for all 17 groups.
- [x] `/copy`, `/save`, `/delete`, `/add`, `/load`, `/showdump`, `lock`, shutdown.
- [x] `.X32res.rc` written and read back exactly (floats bit exact).

Client level (manual, needs the real applications):

- [ ] X32-Edit connects, shows the console name from `/info`/`/xinfo`, syncs
      fader/EQ/dyn/gate/sends and sees changes made by another client.
- [ ] Mixing Station connects from a phone, pulls the `/node` snapshots, uses
      `/meters`, and its changes appear in X32-Edit.
- [ ] Two clients at once stay in sync, including after one of them pauses for
      more than 11 seconds and re-sends `/xremote`.
- [ ] `/shutdown` creates `.X32res.rc` and a restart restores the state.

## Appendices

* `docs/APPENDIX_COMMANDS.md` — every table and every command with its type,
  flags and enum list (generated, 845 kB).
* `docs/x32_commands.json`, `docs/x32_enums.json` — machine readable catalogue,
  suitable for code generation in any language.
* `tools/gen_tables.ps1` — regenerates the catalogue and the Go tables from the C
  command tables of the upstream project (`-Src <X32-Behringer checkout>`); it runs
  `gofmt -w` on the files it writes.
* `tools/gen_appendix.ps1` — regenerates the Markdown appendix.
* `tools/difftest.ps1` — differential test against the reference emulator binary
  (`-Oracle <path>`; it skips itself when the binary is not there).
* `tools/build.ps1` — builds every command into `bin/`.

