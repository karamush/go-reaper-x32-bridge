# Appendix: X32 command catalogue

Generated from `docs/x32_commands.json`, itself extracted by `tools/gen_tables.ps1`
from the command table headers of the upstream X32-Behringer project
(<https://github.com/pmaillot/X32-Behringer>, Patrick-Gilles Maillot, GPLv3).

- total tables: 221
- total entries: 20694 (leaf commands + group headers)
- enum string arrays: 149

## Entries per type

| type | entries |
| --- | --- |
| ACTION | 1 |
| AXPR | 8 |
| BSCO | 103 |
| CCTRL | 3 |
| CENC | 6 |
| CHAMIX | 32 |
| CHCO | 120 |
| CHDE | 32 |
| CHDF | 56 |
| CHDY | 48 |
| CHEQ | 336 |
| CHGA | 32 |
| CHGF | 32 |
| CHGRP | 72 |
| CHIN | 56 |
| CHME | 454 |
| CHMO | 438 |
| CHMX | 64 |
| CHPR | 32 |
| CMIX | 1 |
| CMONO | 1 |
| COSC | 1 |
| CROUTAC | 3 |
| CROUTIN | 2 |
| CROUTOT | 1 |
| CROUTSW | 1 |
| CSOLO | 1 |
| CTALK | 1 |
| CTALKAB | 2 |
| CTAPE | 1 |
| D48 | 1 |
| D48A | 1 |
| D48G | 1 |
| E32 | 4005 |
| F32 | 3468 |
| FX32 | 512 |
| FXPAR1 | 4 |
| FXPAR2 | 4 |
| FXSRC | 4 |
| FXTYP1 | 12 |
| FXTYP2 | 4 |
| HA | 2 |
| HAMP | 129 |
| I32 | 6526 |
| MSMX | 1 |
| MXDY | 8 |
| MXPR | 6 |
| OFFON | 80 |
| OMAIN | 44 |
| OMAIN2 | 3 |
| OMAIND | 16 |
| OP16 | 16 |
| P32 | 950 |
| PADDR | 3 |
| PCARD | 1 |
| PIP | 1 |
| PIQ | 17 |
| PIR | 1 |
| PKEY | 1 |
| PREFS | 1 |
| PRTA | 1 |
| S32 | 1799 |
| SAES | 1 |
| SASS | 1 |
| SCHA | 1 |
| SCUE | 601 |
| SFX | 1 |
| SLIB | 1 |
| SLIBS | 304 |
| SMET | 1 |
| SMON | 1 |
| SNAM | 3 |
| SOSC | 1 |
| SROU | 1 |
| SSCE | 1 |
| SSCN | 101 |
| SSCREEN | 1 |
| SSET | 1 |
| SSNP | 101 |
| SSOLOSW | 1 |
| STALK | 1 |
| STAPE | 1 |
| STAT | 1 |
| SUSB | 1 |
| UREC | 1 |
| UROUI | 1 |
| UROUO | 2 |
| USB | 1 |

## Entries per table

| table | source file | entries |
| --- | --- | --- |
| Xchannel01 | X32Channel.h | 160 |
| Xchannel02 | X32Channel.h | 160 |
| Xchannel03 | X32Channel.h | 160 |
| Xchannel04 | X32Channel.h | 160 |
| Xchannel05 | X32Channel.h | 160 |
| Xchannel06 | X32Channel.h | 160 |
| Xchannel07 | X32Channel.h | 160 |
| Xchannel08 | X32Channel.h | 160 |
| Xchannel09 | X32Channel.h | 160 |
| Xchannel10 | X32Channel.h | 160 |
| Xchannel11 | X32Channel.h | 160 |
| Xchannel12 | X32Channel.h | 160 |
| Xchannel13 | X32Channel.h | 160 |
| Xchannel14 | X32Channel.h | 160 |
| Xchannel15 | X32Channel.h | 160 |
| Xchannel16 | X32Channel.h | 160 |
| Xchannel17 | X32Channel.h | 160 |
| Xchannel18 | X32Channel.h | 160 |
| Xchannel19 | X32Channel.h | 160 |
| Xchannel20 | X32Channel.h | 160 |
| Xchannel21 | X32Channel.h | 160 |
| Xchannel22 | X32Channel.h | 160 |
| Xchannel23 | X32Channel.h | 160 |
| Xchannel24 | X32Channel.h | 160 |
| Xchannel25 | X32Channel.h | 160 |
| Xchannel26 | X32Channel.h | 160 |
| Xchannel27 | X32Channel.h | 160 |
| Xchannel28 | X32Channel.h | 160 |
| Xchannel29 | X32Channel.h | 160 |
| Xchannel30 | X32Channel.h | 160 |
| Xchannel31 | X32Channel.h | 160 |
| Xchannel32 | X32Channel.h | 160 |
| Xconfig | X32CfgMain.h | 335 |
| Xmain | X32CfgMain.h | 182 |
| Xprefs | X32PrefStat.h | 239 |
| Xstat | X32PrefStat.h | 147 |
| Xaction | X32PrefStat.h | 30 |
| Xurec | X32PrefStat.h | 219 |
| Xauxin01 | X32Auxin.h | 114 |
| Xauxin02 | X32Auxin.h | 114 |
| Xauxin03 | X32Auxin.h | 114 |
| Xauxin04 | X32Auxin.h | 114 |
| Xauxin05 | X32Auxin.h | 114 |
| Xauxin06 | X32Auxin.h | 114 |
| Xauxin07 | X32Auxin.h | 114 |
| Xauxin08 | X32Auxin.h | 114 |
| Xfxrtn01 | X32Fxrtn.h | 110 |
| Xfxrtn02 | X32Fxrtn.h | 110 |
| Xfxrtn03 | X32Fxrtn.h | 110 |
| Xfxrtn04 | X32Fxrtn.h | 110 |
| Xfxrtn05 | X32Fxrtn.h | 110 |
| Xfxrtn06 | X32Fxrtn.h | 110 |
| Xfxrtn07 | X32Fxrtn.h | 110 |
| Xfxrtn08 | X32Fxrtn.h | 110 |
| Xbus01 | X32Bus.h | 99 |
| Xbus02 | X32Bus.h | 99 |
| Xbus03 | X32Bus.h | 99 |
| Xbus04 | X32Bus.h | 99 |
| Xbus05 | X32Bus.h | 99 |
| Xbus06 | X32Bus.h | 99 |
| Xbus07 | X32Bus.h | 99 |
| Xbus08 | X32Bus.h | 99 |
| Xbus09 | X32Bus.h | 99 |
| Xbus10 | X32Bus.h | 99 |
| Xbus11 | X32Bus.h | 99 |
| Xbus12 | X32Bus.h | 99 |
| Xbus13 | X32Bus.h | 99 |
| Xbus14 | X32Bus.h | 99 |
| Xbus15 | X32Bus.h | 99 |
| Xbus16 | X32Bus.h | 99 |
| Xmtx01 | X32Mtx.h | 69 |
| Xmtx02 | X32Mtx.h | 69 |
| Xmtx03 | X32Mtx.h | 69 |
| Xmtx04 | X32Mtx.h | 69 |
| Xmtx05 | X32Mtx.h | 69 |
| Xmtx06 | X32Mtx.h | 69 |
| Xdca | X32Dca.h | 57 |
| Xfx1 | X32Fx.h | 71 |
| Xfx2 | X32Fx.h | 71 |
| Xfx3 | X32Fx.h | 71 |
| Xfx4 | X32Fx.h | 71 |
| Xfx5 | X32Fx.h | 68 |
| Xfx6 | X32Fx.h | 68 |
| Xfx7 | X32Fx.h | 68 |
| Xfx8 | X32Fx.h | 68 |
| Xoutput | X32Output.h | 301 |
| Xheadamp | X32Headamp.h | 4 |
| Xheadamp001 | X32Headamp.h | 3 |
| Xheadamp002 | X32Headamp.h | 3 |
| Xheadamp003 | X32Headamp.h | 3 |
| Xheadamp004 | X32Headamp.h | 3 |
| Xheadamp005 | X32Headamp.h | 3 |
| Xheadamp006 | X32Headamp.h | 3 |
| Xheadamp007 | X32Headamp.h | 3 |
| Xheadamp008 | X32Headamp.h | 3 |
| Xheadamp009 | X32Headamp.h | 3 |
| Xheadamp010 | X32Headamp.h | 3 |
| Xheadamp011 | X32Headamp.h | 3 |
| Xheadamp012 | X32Headamp.h | 3 |
| Xheadamp013 | X32Headamp.h | 3 |
| Xheadamp014 | X32Headamp.h | 3 |
| Xheadamp015 | X32Headamp.h | 3 |
| Xheadamp016 | X32Headamp.h | 3 |
| Xheadamp017 | X32Headamp.h | 3 |
| Xheadamp018 | X32Headamp.h | 3 |
| Xheadamp019 | X32Headamp.h | 3 |
| Xheadamp020 | X32Headamp.h | 3 |
| Xheadamp021 | X32Headamp.h | 3 |
| Xheadamp022 | X32Headamp.h | 3 |
| Xheadamp023 | X32Headamp.h | 3 |
| Xheadamp024 | X32Headamp.h | 3 |
| Xheadamp025 | X32Headamp.h | 3 |
| Xheadamp026 | X32Headamp.h | 3 |
| Xheadamp027 | X32Headamp.h | 3 |
| Xheadamp028 | X32Headamp.h | 3 |
| Xheadamp029 | X32Headamp.h | 3 |
| Xheadamp030 | X32Headamp.h | 3 |
| Xheadamp031 | X32Headamp.h | 3 |
| Xheadamp032 | X32Headamp.h | 3 |
| Xheadamp033 | X32Headamp.h | 3 |
| Xheadamp034 | X32Headamp.h | 3 |
| Xheadamp035 | X32Headamp.h | 3 |
| Xheadamp036 | X32Headamp.h | 3 |
| Xheadamp037 | X32Headamp.h | 3 |
| Xheadamp038 | X32Headamp.h | 3 |
| Xheadamp039 | X32Headamp.h | 3 |
| Xheadamp040 | X32Headamp.h | 3 |
| Xheadamp041 | X32Headamp.h | 3 |
| Xheadamp042 | X32Headamp.h | 3 |
| Xheadamp043 | X32Headamp.h | 3 |
| Xheadamp044 | X32Headamp.h | 3 |
| Xheadamp045 | X32Headamp.h | 3 |
| Xheadamp046 | X32Headamp.h | 3 |
| Xheadamp047 | X32Headamp.h | 3 |
| Xheadamp048 | X32Headamp.h | 3 |
| Xheadamp049 | X32Headamp.h | 3 |
| Xheadamp050 | X32Headamp.h | 3 |
| Xheadamp051 | X32Headamp.h | 3 |
| Xheadamp052 | X32Headamp.h | 3 |
| Xheadamp053 | X32Headamp.h | 3 |
| Xheadamp054 | X32Headamp.h | 3 |
| Xheadamp055 | X32Headamp.h | 3 |
| Xheadamp056 | X32Headamp.h | 3 |
| Xheadamp057 | X32Headamp.h | 3 |
| Xheadamp058 | X32Headamp.h | 3 |
| Xheadamp059 | X32Headamp.h | 3 |
| Xheadamp060 | X32Headamp.h | 3 |
| Xheadamp061 | X32Headamp.h | 3 |
| Xheadamp062 | X32Headamp.h | 3 |
| Xheadamp063 | X32Headamp.h | 3 |
| Xheadamp064 | X32Headamp.h | 3 |
| Xheadamp065 | X32Headamp.h | 3 |
| Xheadamp066 | X32Headamp.h | 3 |
| Xheadamp067 | X32Headamp.h | 3 |
| Xheadamp068 | X32Headamp.h | 3 |
| Xheadamp069 | X32Headamp.h | 3 |
| Xheadamp070 | X32Headamp.h | 3 |
| Xheadamp071 | X32Headamp.h | 3 |
| Xheadamp072 | X32Headamp.h | 3 |
| Xheadamp073 | X32Headamp.h | 3 |
| Xheadamp074 | X32Headamp.h | 3 |
| Xheadamp075 | X32Headamp.h | 3 |
| Xheadamp076 | X32Headamp.h | 3 |
| Xheadamp077 | X32Headamp.h | 3 |
| Xheadamp078 | X32Headamp.h | 3 |
| Xheadamp079 | X32Headamp.h | 3 |
| Xheadamp080 | X32Headamp.h | 3 |
| Xheadamp081 | X32Headamp.h | 3 |
| Xheadamp082 | X32Headamp.h | 3 |
| Xheadamp083 | X32Headamp.h | 3 |
| Xheadamp084 | X32Headamp.h | 3 |
| Xheadamp085 | X32Headamp.h | 3 |
| Xheadamp086 | X32Headamp.h | 3 |
| Xheadamp087 | X32Headamp.h | 3 |
| Xheadamp088 | X32Headamp.h | 3 |
| Xheadamp089 | X32Headamp.h | 3 |
| Xheadamp090 | X32Headamp.h | 3 |
| Xheadamp091 | X32Headamp.h | 3 |
| Xheadamp092 | X32Headamp.h | 3 |
| Xheadamp093 | X32Headamp.h | 3 |
| Xheadamp094 | X32Headamp.h | 3 |
| Xheadamp095 | X32Headamp.h | 3 |
| Xheadamp096 | X32Headamp.h | 3 |
| Xheadamp097 | X32Headamp.h | 3 |
| Xheadamp098 | X32Headamp.h | 3 |
| Xheadamp099 | X32Headamp.h | 3 |
| Xheadamp100 | X32Headamp.h | 3 |
| Xheadamp101 | X32Headamp.h | 3 |
| Xheadamp102 | X32Headamp.h | 3 |
| Xheadamp103 | X32Headamp.h | 3 |
| Xheadamp104 | X32Headamp.h | 3 |
| Xheadamp105 | X32Headamp.h | 3 |
| Xheadamp106 | X32Headamp.h | 3 |
| Xheadamp107 | X32Headamp.h | 3 |
| Xheadamp108 | X32Headamp.h | 3 |
| Xheadamp109 | X32Headamp.h | 3 |
| Xheadamp110 | X32Headamp.h | 3 |
| Xheadamp111 | X32Headamp.h | 3 |
| Xheadamp112 | X32Headamp.h | 3 |
| Xheadamp113 | X32Headamp.h | 3 |
| Xheadamp114 | X32Headamp.h | 3 |
| Xheadamp115 | X32Headamp.h | 3 |
| Xheadamp116 | X32Headamp.h | 3 |
| Xheadamp117 | X32Headamp.h | 3 |
| Xheadamp118 | X32Headamp.h | 3 |
| Xheadamp119 | X32Headamp.h | 3 |
| Xheadamp120 | X32Headamp.h | 3 |
| Xheadamp121 | X32Headamp.h | 3 |
| Xheadamp122 | X32Headamp.h | 3 |
| Xheadamp123 | X32Headamp.h | 3 |
| Xheadamp124 | X32Headamp.h | 3 |
| Xheadamp125 | X32Headamp.h | 3 |
| Xheadamp126 | X32Headamp.h | 3 |
| Xheadamp127 | X32Headamp.h | 3 |
| Xshow | X32Show.h | 6016 |
| Xscene | X32Show.h | 501 |
| Xsnippet | X32Show.h | 701 |
| Xmisc | X32Misc.h | 311 |
| Xlibsc | X32Libs.h | 602 |
| Xlibsr | X32Libs.h | 601 |
| Xlibsf | X32Libs.h | 601 |

## Command tables

Group headers are marked with `F_FND`; `n=` is the item count stored in the
table (used by the `/node` renderer when it is not zero). Leaves carry their
value type and their enum array name (if any).

### Xchannel01 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/01  <CHCO> n=0
/ch/01/config  <CHCO> n=0
    /ch/01/config/name  S32 F_XET
    /ch/01/config/icon  I32 F_XET
    /ch/01/config/color  E32 F_XET enum=Xcolors
    /ch/01/config/source  I32 F_XET
/ch/01/grp  <CHGRP> n=0
    /ch/01/grp/dca  P32 F_XET
    /ch/01/grp/mute  P32 F_XET
/ch/01/preamp  <CHPR> n=0
    /ch/01/preamp/trim  F32 F_XET
    /ch/01/preamp/invert  E32 F_XET enum=OffOn
    /ch/01/preamp/hpon  E32 F_XET enum=OffOn
    /ch/01/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/01/preamp/hpf  F32 F_XET
/ch/01/delay  <CHDE> n=0
    /ch/01/delay/on  E32 F_XET enum=OffOn
    /ch/01/delay/time  F32 F_XET
/ch/01/insert  <CHIN> n=0
    /ch/01/insert/on  E32 F_XET enum=OffOn
    /ch/01/insert/pos  E32 F_XET enum=Xdyppos
    /ch/01/insert/sel  E32 F_XET enum=Xisel
/ch/01/gate  <CHGA> n=0
    /ch/01/gate/on  E32 F_XET enum=OffOn
    /ch/01/gate/mode  E32 F_XET enum=Xgmode
    /ch/01/gate/thr  F32 F_XET
    /ch/01/gate/range  F32 F_XET
    /ch/01/gate/attack  F32 F_XET
    /ch/01/gate/hold  F32 F_XET
    /ch/01/gate/release  F32 F_XET
    /ch/01/gate/keysrc  I32 F_XET
/ch/01/gate/filter  <CHGF> n=0
    /ch/01/gate/filter/on  E32 F_XET enum=OffOn
    /ch/01/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/01/gate/filter/f  F32 F_XET
/ch/01/dyn  <CHDY> n=0
    /ch/01/dyn/on  E32 F_XET enum=OffOn
    /ch/01/dyn/mode  E32 F_XET enum=Xdymode
    /ch/01/dyn/det  E32 F_XET enum=Xdydet
    /ch/01/dyn/env  E32 F_XET enum=Xdyenv
    /ch/01/dyn/thr  F32 F_XET
    /ch/01/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/01/dyn/knee  F32 F_XET
    /ch/01/dyn/mgain  F32 F_XET
    /ch/01/dyn/attack  F32 F_XET
    /ch/01/dyn/hold  F32 F_XET
    /ch/01/dyn/release  F32 F_XET
    /ch/01/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/01/dyn/keysrc  I32 F_XET
    /ch/01/dyn/mix  F32 F_XET
    /ch/01/dyn/auto  E32 F_XET enum=OffOn
/ch/01/dyn/filter  <CHDF> n=0
    /ch/01/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/01/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/01/dyn/filter/f  F32 F_XET
/ch/01/eq  <OFFON> n=1
    /ch/01/eq/on  E32 F_XET enum=OffOn
/ch/01/eq/1  <CHEQ> n=0
    /ch/01/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/01/eq/1/f  F32 F_XET
    /ch/01/eq/1/g  F32 F_XET
    /ch/01/eq/1/q  F32 F_XET
/ch/01/eq/2  <CHEQ> n=0
    /ch/01/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/01/eq/2/f  F32 F_XET
    /ch/01/eq/2/g  F32 F_XET
    /ch/01/eq/2/q  F32 F_XET
/ch/01/eq/3  <CHEQ> n=0
    /ch/01/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/01/eq/3/f  F32 F_XET
    /ch/01/eq/3/g  F32 F_XET
    /ch/01/eq/3/q  F32 F_XET
/ch/01/eq/4  <CHEQ> n=0
    /ch/01/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/01/eq/4/f  F32 F_XET
    /ch/01/eq/4/g  F32 F_XET
    /ch/01/eq/4/q  F32 F_XET
/ch/01/mix  <CHMX> n=0
    /ch/01/mix/on  E32 F_XET enum=OffOn
    /ch/01/mix/fader  F32 F_XET
    /ch/01/mix/st  E32 F_XET enum=OffOn
    /ch/01/mix/pan  F32 F_XET
    /ch/01/mix/mono  E32 F_XET enum=OffOn
    /ch/01/mix/mlevel  F32 F_XET
/ch/01/mix/01  <CHMO> n=0
    /ch/01/mix/01/on  E32 F_XET enum=OffOn
    /ch/01/mix/01/level  F32 F_XET
    /ch/01/mix/01/pan  F32 F_XET
    /ch/01/mix/01/type  E32 F_XET enum=Xmtype
    /ch/01/mix/01/panFollow  E32 F_XET
/ch/01/mix/02  <CHME> n=0
    /ch/01/mix/02/on  E32 F_XET enum=OffOn
    /ch/01/mix/02/level  F32 F_XET
/ch/01/mix/03  <CHMO> n=0
    /ch/01/mix/03/on  E32 F_XET enum=OffOn
    /ch/01/mix/03/level  F32 F_XET
    /ch/01/mix/03/pan  F32 F_XET
    /ch/01/mix/03/type  E32 F_XET enum=Xmtype
    /ch/01/mix/03/panFollow  E32 F_XET
/ch/01/mix/04  <CHME> n=0
    /ch/01/mix/04/on  E32 F_XET enum=OffOn
    /ch/01/mix/04/level  F32 F_XET
/ch/01/mix/05  <CHMO> n=0
    /ch/01/mix/05/on  E32 F_XET enum=OffOn
    /ch/01/mix/05/level  F32 F_XET
    /ch/01/mix/05/pan  F32 F_XET
    /ch/01/mix/05/type  E32 F_XET enum=Xmtype
    /ch/01/mix/05/panFollow  E32 F_XET
/ch/01/mix/06  <CHME> n=0
    /ch/01/mix/06/on  E32 F_XET enum=OffOn
    /ch/01/mix/06/level  F32 F_XET
/ch/01/mix/07  <CHMO> n=0
    /ch/01/mix/07/on  E32 F_XET enum=OffOn
    /ch/01/mix/07/level  F32 F_XET
    /ch/01/mix/07/pan  F32 F_XET
    /ch/01/mix/07/type  E32 F_XET enum=Xmtype
    /ch/01/mix/07/panFollow  E32 F_XET
/ch/01/mix/08  <CHME> n=0
    /ch/01/mix/08/on  E32 F_XET enum=OffOn
    /ch/01/mix/08/level  F32 F_XET
/ch/01/mix/09  <CHMO> n=0
    /ch/01/mix/09/on  E32 F_XET enum=OffOn
    /ch/01/mix/09/level  F32 F_XET
    /ch/01/mix/09/pan  F32 F_XET
    /ch/01/mix/09/type  E32 F_XET enum=Xmtype
    /ch/01/mix/09/panFollow  E32 F_XET
/ch/01/mix/10  <CHME> n=0
    /ch/01/mix/10/on  E32 F_XET enum=OffOn
    /ch/01/mix/10/level  F32 F_XET
/ch/01/mix/11  <CHMO> n=0
    /ch/01/mix/11/on  E32 F_XET enum=OffOn
    /ch/01/mix/11/level  F32 F_XET
    /ch/01/mix/11/pan  F32 F_XET
    /ch/01/mix/11/type  E32 F_XET enum=Xmtype
    /ch/01/mix/11/panFollow  E32 F_XET
/ch/01/mix/12  <CHME> n=0
    /ch/01/mix/12/on  E32 F_XET enum=OffOn
    /ch/01/mix/12/level  F32 F_XET
/ch/01/mix/13  <CHMO> n=0
    /ch/01/mix/13/on  E32 F_XET enum=OffOn
    /ch/01/mix/13/level  F32 F_XET
    /ch/01/mix/13/pan  F32 F_XET
    /ch/01/mix/13/type  E32 F_XET enum=Xmtype
    /ch/01/mix/13/panFollow  E32 F_XET
/ch/01/mix/14  <CHME> n=0
    /ch/01/mix/14/on  E32 F_XET enum=OffOn
    /ch/01/mix/14/level  F32 F_XET
/ch/01/mix/15  <CHMO> n=0
    /ch/01/mix/15/on  E32 F_XET enum=OffOn
    /ch/01/mix/15/level  F32 F_XET
    /ch/01/mix/15/pan  F32 F_XET
    /ch/01/mix/15/type  E32 F_XET enum=Xmtype
    /ch/01/mix/15/panFollow  E32 F_XET
/ch/01/mix/16  <CHME> n=0
    /ch/01/mix/16/on  E32 F_XET enum=OffOn
    /ch/01/mix/16/level  F32 F_XET
/ch/01/automix  <CHAMIX> n=0
    /ch/01/automix/group  E32 F_XET enum=Xamxgrp
    /ch/01/automix/weight  F32 F_XET
```

### Xchannel02 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/02  <CHCO> n=0
/ch/02/config  <CHCO> n=0
    /ch/02/config/name  S32 F_XET
    /ch/02/config/icon  I32 F_XET
    /ch/02/config/color  E32 F_XET enum=Xcolors
    /ch/02/config/source  I32 F_XET
/ch/02/grp  <CHGRP> n=0
    /ch/02/grp/dca  P32 F_XET
    /ch/02/grp/mute  P32 F_XET
/ch/02/preamp  <CHPR> n=0
    /ch/02/preamp/trim  F32 F_XET
    /ch/02/preamp/invert  E32 F_XET enum=OffOn
    /ch/02/preamp/hpon  E32 F_XET enum=OffOn
    /ch/02/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/02/preamp/hpf  F32 F_XET
/ch/02/delay  <CHDE> n=0
    /ch/02/delay/on  E32 F_XET enum=OffOn
    /ch/02/delay/time  F32 F_XET
/ch/02/insert  <CHIN> n=0
    /ch/02/insert/on  E32 F_XET enum=OffOn
    /ch/02/insert/pos  E32 F_XET enum=Xdyppos
    /ch/02/insert/sel  E32 F_XET enum=Xisel
/ch/02/gate  <CHGA> n=0
    /ch/02/gate/on  E32 F_XET enum=OffOn
    /ch/02/gate/mode  E32 F_XET enum=Xgmode
    /ch/02/gate/thr  F32 F_XET
    /ch/02/gate/range  F32 F_XET
    /ch/02/gate/attack  F32 F_XET
    /ch/02/gate/hold  F32 F_XET
    /ch/02/gate/release  F32 F_XET
    /ch/02/gate/keysrc  I32 F_XET
/ch/02/gate/filter  <CHGF> n=0
    /ch/02/gate/filter/on  E32 F_XET enum=OffOn
    /ch/02/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/02/gate/filter/f  F32 F_XET
/ch/02/dyn  <CHDY> n=0
    /ch/02/dyn/on  E32 F_XET enum=OffOn
    /ch/02/dyn/mode  E32 F_XET enum=Xdymode
    /ch/02/dyn/det  E32 F_XET enum=Xdydet
    /ch/02/dyn/env  E32 F_XET enum=Xdyenv
    /ch/02/dyn/thr  F32 F_XET
    /ch/02/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/02/dyn/knee  F32 F_XET
    /ch/02/dyn/mgain  F32 F_XET
    /ch/02/dyn/attack  F32 F_XET
    /ch/02/dyn/hold  F32 F_XET
    /ch/02/dyn/release  F32 F_XET
    /ch/02/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/02/dyn/keysrc  I32 F_XET
    /ch/02/dyn/mix  F32 F_XET
    /ch/02/dyn/auto  E32 F_XET enum=OffOn
/ch/02/dyn/filter  <CHDF> n=0
    /ch/02/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/02/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/02/dyn/filter/f  F32 F_XET
/ch/02/eq  <OFFON> n=1
    /ch/02/eq/on  E32 F_XET enum=OffOn
/ch/02/eq/1  <CHEQ> n=0
    /ch/02/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/02/eq/1/f  F32 F_XET
    /ch/02/eq/1/g  F32 F_XET
    /ch/02/eq/1/q  F32 F_XET
/ch/02/eq/2  <CHEQ> n=0
    /ch/02/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/02/eq/2/f  F32 F_XET
    /ch/02/eq/2/g  F32 F_XET
    /ch/02/eq/2/q  F32 F_XET
/ch/02/eq/3  <CHEQ> n=0
    /ch/02/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/02/eq/3/f  F32 F_XET
    /ch/02/eq/3/g  F32 F_XET
    /ch/02/eq/3/q  F32 F_XET
/ch/02/eq/4  <CHEQ> n=0
    /ch/02/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/02/eq/4/f  F32 F_XET
    /ch/02/eq/4/g  F32 F_XET
    /ch/02/eq/4/q  F32 F_XET
/ch/02/mix  <CHMX> n=0
    /ch/02/mix/on  E32 F_XET enum=OffOn
    /ch/02/mix/fader  F32 F_XET
    /ch/02/mix/st  E32 F_XET enum=OffOn
    /ch/02/mix/pan  F32 F_XET
    /ch/02/mix/mono  E32 F_XET enum=OffOn
    /ch/02/mix/mlevel  F32 F_XET
/ch/02/mix/01  <CHMO> n=0
    /ch/02/mix/01/on  E32 F_XET enum=OffOn
    /ch/02/mix/01/level  F32 F_XET
    /ch/02/mix/01/pan  F32 F_XET
    /ch/02/mix/01/type  E32 F_XET enum=Xmtype
    /ch/02/mix/01/panFollow  E32 F_XET
/ch/02/mix/02  <CHME> n=0
    /ch/02/mix/02/on  E32 F_XET enum=OffOn
    /ch/02/mix/02/level  F32 F_XET
/ch/02/mix/03  <CHMO> n=0
    /ch/02/mix/03/on  E32 F_XET enum=OffOn
    /ch/02/mix/03/level  F32 F_XET
    /ch/02/mix/03/pan  F32 F_XET
    /ch/02/mix/03/type  E32 F_XET enum=Xmtype
    /ch/02/mix/03/panFollow  E32 F_XET
/ch/02/mix/04  <CHME> n=0
    /ch/02/mix/04/on  E32 F_XET enum=OffOn
    /ch/02/mix/04/level  F32 F_XET
/ch/02/mix/05  <CHMO> n=0
    /ch/02/mix/05/on  E32 F_XET enum=OffOn
    /ch/02/mix/05/level  F32 F_XET
    /ch/02/mix/05/pan  F32 F_XET
    /ch/02/mix/05/type  E32 F_XET enum=Xmtype
    /ch/02/mix/05/panFollow  E32 F_XET
/ch/02/mix/06  <CHME> n=0
    /ch/02/mix/06/on  E32 F_XET enum=OffOn
    /ch/02/mix/06/level  F32 F_XET
/ch/02/mix/07  <CHMO> n=0
    /ch/02/mix/07/on  E32 F_XET enum=OffOn
    /ch/02/mix/07/level  F32 F_XET
    /ch/02/mix/07/pan  F32 F_XET
    /ch/02/mix/07/type  E32 F_XET enum=Xmtype
    /ch/02/mix/07/panFollow  E32 F_XET
/ch/02/mix/08  <CHME> n=0
    /ch/02/mix/08/on  E32 F_XET enum=OffOn
    /ch/02/mix/08/level  F32 F_XET
/ch/02/mix/09  <CHMO> n=0
    /ch/02/mix/09/on  E32 F_XET enum=OffOn
    /ch/02/mix/09/level  F32 F_XET
    /ch/02/mix/09/pan  F32 F_XET
    /ch/02/mix/09/type  E32 F_XET enum=Xmtype
    /ch/02/mix/09/panFollow  E32 F_XET
/ch/02/mix/10  <CHME> n=0
    /ch/02/mix/10/on  E32 F_XET enum=OffOn
    /ch/02/mix/10/level  F32 F_XET
/ch/02/mix/11  <CHMO> n=0
    /ch/02/mix/11/on  E32 F_XET enum=OffOn
    /ch/02/mix/11/level  F32 F_XET
    /ch/02/mix/11/pan  F32 F_XET
    /ch/02/mix/11/type  E32 F_XET enum=Xmtype
    /ch/02/mix/11/panFollow  E32 F_XET
/ch/02/mix/12  <CHME> n=0
    /ch/02/mix/12/on  E32 F_XET enum=OffOn
    /ch/02/mix/12/level  F32 F_XET
/ch/02/mix/13  <CHMO> n=0
    /ch/02/mix/13/on  E32 F_XET enum=OffOn
    /ch/02/mix/13/level  F32 F_XET
    /ch/02/mix/13/pan  F32 F_XET
    /ch/02/mix/13/type  E32 F_XET enum=Xmtype
    /ch/02/mix/13/panFollow  E32 F_XET
/ch/02/mix/14  <CHME> n=0
    /ch/02/mix/14/on  E32 F_XET enum=OffOn
    /ch/02/mix/14/level  F32 F_XET
/ch/02/mix/15  <CHMO> n=0
    /ch/02/mix/15/on  E32 F_XET enum=OffOn
    /ch/02/mix/15/level  F32 F_XET
    /ch/02/mix/15/pan  F32 F_XET
    /ch/02/mix/15/type  E32 F_XET enum=Xmtype
    /ch/02/mix/15/panFollow  E32 F_XET
/ch/02/mix/16  <CHME> n=0
    /ch/02/mix/16/on  E32 F_XET enum=OffOn
    /ch/02/mix/16/level  F32 F_XET
/ch/02/automix  <CHAMIX> n=0
    /ch/02/automix/group  E32 F_XET enum=Xamxgrp
    /ch/02/automix/weight  F32 F_XET
```

### Xchannel03 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/03  <CHCO> n=0
/ch/03/config  <CHCO> n=0
    /ch/03/config/name  S32 F_XET
    /ch/03/config/icon  I32 F_XET
    /ch/03/config/color  E32 F_XET enum=Xcolors
    /ch/03/config/source  I32 F_XET
/ch/03/grp  <CHGRP> n=0
    /ch/03/grp/dca  P32 F_XET
    /ch/03/grp/mute  P32 F_XET
/ch/03/preamp  <CHPR> n=0
    /ch/03/preamp/trim  F32 F_XET
    /ch/03/preamp/invert  E32 F_XET enum=OffOn
    /ch/03/preamp/hpon  E32 F_XET enum=OffOn
    /ch/03/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/03/preamp/hpf  F32 F_XET
/ch/03/delay  <CHDE> n=0
    /ch/03/delay/on  E32 F_XET enum=OffOn
    /ch/03/delay/time  F32 F_XET
/ch/03/insert  <CHIN> n=0
    /ch/03/insert/on  E32 F_XET enum=OffOn
    /ch/03/insert/pos  E32 F_XET enum=Xdyppos
    /ch/03/insert/sel  E32 F_XET enum=Xisel
/ch/03/gate  <CHGA> n=0
    /ch/03/gate/on  E32 F_XET enum=OffOn
    /ch/03/gate/mode  E32 F_XET enum=Xgmode
    /ch/03/gate/thr  F32 F_XET
    /ch/03/gate/range  F32 F_XET
    /ch/03/gate/attack  F32 F_XET
    /ch/03/gate/hold  F32 F_XET
    /ch/03/gate/release  F32 F_XET
    /ch/03/gate/keysrc  I32 F_XET
/ch/03/gate/filter  <CHGF> n=0
    /ch/03/gate/filter/on  E32 F_XET enum=OffOn
    /ch/03/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/03/gate/filter/f  F32 F_XET
/ch/03/dyn  <CHDY> n=0
    /ch/03/dyn/on  E32 F_XET enum=OffOn
    /ch/03/dyn/mode  E32 F_XET enum=Xdymode
    /ch/03/dyn/det  E32 F_XET enum=Xdydet
    /ch/03/dyn/env  E32 F_XET enum=Xdyenv
    /ch/03/dyn/thr  F32 F_XET
    /ch/03/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/03/dyn/knee  F32 F_XET
    /ch/03/dyn/mgain  F32 F_XET
    /ch/03/dyn/attack  F32 F_XET
    /ch/03/dyn/hold  F32 F_XET
    /ch/03/dyn/release  F32 F_XET
    /ch/03/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/03/dyn/keysrc  I32 F_XET
    /ch/03/dyn/mix  F32 F_XET
    /ch/03/dyn/auto  E32 F_XET enum=OffOn
/ch/03/dyn/filter  <CHDF> n=0
    /ch/03/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/03/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/03/dyn/filter/f  F32 F_XET
/ch/03/eq  <OFFON> n=1
    /ch/03/eq/on  E32 F_XET enum=OffOn
/ch/03/eq/1  <CHEQ> n=0
    /ch/03/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/03/eq/1/f  F32 F_XET
    /ch/03/eq/1/g  F32 F_XET
    /ch/03/eq/1/q  F32 F_XET
/ch/03/eq/2  <CHEQ> n=0
    /ch/03/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/03/eq/2/f  F32 F_XET
    /ch/03/eq/2/g  F32 F_XET
    /ch/03/eq/2/q  F32 F_XET
/ch/03/eq/3  <CHEQ> n=0
    /ch/03/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/03/eq/3/f  F32 F_XET
    /ch/03/eq/3/g  F32 F_XET
    /ch/03/eq/3/q  F32 F_XET
/ch/03/eq/4  <CHEQ> n=0
    /ch/03/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/03/eq/4/f  F32 F_XET
    /ch/03/eq/4/g  F32 F_XET
    /ch/03/eq/4/q  F32 F_XET
/ch/03/mix  <CHMX> n=0
    /ch/03/mix/on  E32 F_XET enum=OffOn
    /ch/03/mix/fader  F32 F_XET
    /ch/03/mix/st  E32 F_XET enum=OffOn
    /ch/03/mix/pan  F32 F_XET
    /ch/03/mix/mono  E32 F_XET enum=OffOn
    /ch/03/mix/mlevel  F32 F_XET
/ch/03/mix/01  <CHMO> n=0
    /ch/03/mix/01/on  E32 F_XET enum=OffOn
    /ch/03/mix/01/level  F32 F_XET
    /ch/03/mix/01/pan  F32 F_XET
    /ch/03/mix/01/type  E32 F_XET enum=Xmtype
    /ch/03/mix/01/panFollow  E32 F_XET
/ch/03/mix/02  <CHME> n=0
    /ch/03/mix/02/on  E32 F_XET enum=OffOn
    /ch/03/mix/02/level  F32 F_XET
/ch/03/mix/03  <CHMO> n=0
    /ch/03/mix/03/on  E32 F_XET enum=OffOn
    /ch/03/mix/03/level  F32 F_XET
    /ch/03/mix/03/pan  F32 F_XET
    /ch/03/mix/03/type  E32 F_XET enum=Xmtype
    /ch/03/mix/03/panFollow  E32 F_XET
/ch/03/mix/04  <CHME> n=0
    /ch/03/mix/04/on  E32 F_XET enum=OffOn
    /ch/03/mix/04/level  F32 F_XET
/ch/03/mix/05  <CHMO> n=0
    /ch/03/mix/05/on  E32 F_XET enum=OffOn
    /ch/03/mix/05/level  F32 F_XET
    /ch/03/mix/05/pan  F32 F_XET
    /ch/03/mix/05/type  E32 F_XET enum=Xmtype
    /ch/03/mix/05/panFollow  E32 F_XET
/ch/03/mix/06  <CHME> n=0
    /ch/03/mix/06/on  E32 F_XET enum=OffOn
    /ch/03/mix/06/level  F32 F_XET
/ch/03/mix/07  <CHMO> n=0
    /ch/03/mix/07/on  E32 F_XET enum=OffOn
    /ch/03/mix/07/level  F32 F_XET
    /ch/03/mix/07/pan  F32 F_XET
    /ch/03/mix/07/type  E32 F_XET enum=Xmtype
    /ch/03/mix/07/panFollow  E32 F_XET
/ch/03/mix/08  <CHME> n=0
    /ch/03/mix/08/on  E32 F_XET enum=OffOn
    /ch/03/mix/08/level  F32 F_XET
/ch/03/mix/09  <CHMO> n=0
    /ch/03/mix/09/on  E32 F_XET enum=OffOn
    /ch/03/mix/09/level  F32 F_XET
    /ch/03/mix/09/pan  F32 F_XET
    /ch/03/mix/09/type  E32 F_XET enum=Xmtype
    /ch/03/mix/09/panFollow  E32 F_XET
/ch/03/mix/10  <CHME> n=0
    /ch/03/mix/10/on  E32 F_XET enum=OffOn
    /ch/03/mix/10/level  F32 F_XET
/ch/03/mix/11  <CHMO> n=0
    /ch/03/mix/11/on  E32 F_XET enum=OffOn
    /ch/03/mix/11/level  F32 F_XET
    /ch/03/mix/11/pan  F32 F_XET
    /ch/03/mix/11/type  E32 F_XET enum=Xmtype
    /ch/03/mix/11/panFollow  E32 F_XET
/ch/03/mix/12  <CHME> n=0
    /ch/03/mix/12/on  E32 F_XET enum=OffOn
    /ch/03/mix/12/level  F32 F_XET
/ch/03/mix/13  <CHMO> n=0
    /ch/03/mix/13/on  E32 F_XET enum=OffOn
    /ch/03/mix/13/level  F32 F_XET
    /ch/03/mix/13/pan  F32 F_XET
    /ch/03/mix/13/type  E32 F_XET enum=Xmtype
    /ch/03/mix/13/panFollow  E32 F_XET
/ch/03/mix/14  <CHME> n=0
    /ch/03/mix/14/on  E32 F_XET enum=OffOn
    /ch/03/mix/14/level  F32 F_XET
/ch/03/mix/15  <CHMO> n=0
    /ch/03/mix/15/on  E32 F_XET enum=OffOn
    /ch/03/mix/15/level  F32 F_XET
    /ch/03/mix/15/pan  F32 F_XET
    /ch/03/mix/15/type  E32 F_XET enum=Xmtype
    /ch/03/mix/15/panFollow  E32 F_XET
/ch/03/mix/16  <CHME> n=0
    /ch/03/mix/16/on  E32 F_XET enum=OffOn
    /ch/03/mix/16/level  F32 F_XET
/ch/03/automix  <CHAMIX> n=0
    /ch/03/automix/group  E32 F_XET enum=Xamxgrp
    /ch/03/automix/weight  F32 F_XET
```

### Xchannel04 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/04  <CHCO> n=0
/ch/04/config  <CHCO> n=0
    /ch/04/config/name  S32 F_XET
    /ch/04/config/icon  I32 F_XET
    /ch/04/config/color  E32 F_XET enum=Xcolors
    /ch/04/config/source  I32 F_XET
/ch/04/grp  <CHGRP> n=0
    /ch/04/grp/dca  P32 F_XET
    /ch/04/grp/mute  P32 F_XET
/ch/04/preamp  <CHPR> n=0
    /ch/04/preamp/trim  F32 F_XET
    /ch/04/preamp/invert  E32 F_XET enum=OffOn
    /ch/04/preamp/hpon  E32 F_XET enum=OffOn
    /ch/04/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/04/preamp/hpf  F32 F_XET
/ch/04/delay  <CHDE> n=0
    /ch/04/delay/on  E32 F_XET enum=OffOn
    /ch/04/delay/time  F32 F_XET
/ch/04/insert  <CHIN> n=0
    /ch/04/insert/on  E32 F_XET enum=OffOn
    /ch/04/insert/pos  E32 F_XET enum=Xdyppos
    /ch/04/insert/sel  E32 F_XET enum=Xisel
/ch/04/gate  <CHGA> n=0
    /ch/04/gate/on  E32 F_XET enum=OffOn
    /ch/04/gate/mode  E32 F_XET enum=Xgmode
    /ch/04/gate/thr  F32 F_XET
    /ch/04/gate/range  F32 F_XET
    /ch/04/gate/attack  F32 F_XET
    /ch/04/gate/hold  F32 F_XET
    /ch/04/gate/release  F32 F_XET
    /ch/04/gate/keysrc  I32 F_XET
/ch/04/gate/filter  <CHGF> n=0
    /ch/04/gate/filter/on  E32 F_XET enum=OffOn
    /ch/04/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/04/gate/filter/f  F32 F_XET
/ch/04/dyn  <CHDY> n=0
    /ch/04/dyn/on  E32 F_XET enum=OffOn
    /ch/04/dyn/mode  E32 F_XET enum=Xdymode
    /ch/04/dyn/det  E32 F_XET enum=Xdydet
    /ch/04/dyn/env  E32 F_XET enum=Xdyenv
    /ch/04/dyn/thr  F32 F_XET
    /ch/04/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/04/dyn/knee  F32 F_XET
    /ch/04/dyn/mgain  F32 F_XET
    /ch/04/dyn/attack  F32 F_XET
    /ch/04/dyn/hold  F32 F_XET
    /ch/04/dyn/release  F32 F_XET
    /ch/04/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/04/dyn/keysrc  I32 F_XET
    /ch/04/dyn/mix  F32 F_XET
    /ch/04/dyn/auto  E32 F_XET enum=OffOn
/ch/04/dyn/filter  <CHDF> n=0
    /ch/04/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/04/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/04/dyn/filter/f  F32 F_XET
/ch/04/eq  <OFFON> n=1
    /ch/04/eq/on  E32 F_XET enum=OffOn
/ch/04/eq/1  <CHEQ> n=0
    /ch/04/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/04/eq/1/f  F32 F_XET
    /ch/04/eq/1/g  F32 F_XET
    /ch/04/eq/1/q  F32 F_XET
/ch/04/eq/2  <CHEQ> n=0
    /ch/04/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/04/eq/2/f  F32 F_XET
    /ch/04/eq/2/g  F32 F_XET
    /ch/04/eq/2/q  F32 F_XET
/ch/04/eq/3  <CHEQ> n=0
    /ch/04/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/04/eq/3/f  F32 F_XET
    /ch/04/eq/3/g  F32 F_XET
    /ch/04/eq/3/q  F32 F_XET
/ch/04/eq/4  <CHEQ> n=0
    /ch/04/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/04/eq/4/f  F32 F_XET
    /ch/04/eq/4/g  F32 F_XET
    /ch/04/eq/4/q  F32 F_XET
/ch/04/mix  <CHMX> n=0
    /ch/04/mix/on  E32 F_XET enum=OffOn
    /ch/04/mix/fader  F32 F_XET
    /ch/04/mix/st  E32 F_XET enum=OffOn
    /ch/04/mix/pan  F32 F_XET
    /ch/04/mix/mono  E32 F_XET enum=OffOn
    /ch/04/mix/mlevel  F32 F_XET
/ch/04/mix/01  <CHMO> n=0
    /ch/04/mix/01/on  E32 F_XET enum=OffOn
    /ch/04/mix/01/level  F32 F_XET
    /ch/04/mix/01/pan  F32 F_XET
    /ch/04/mix/01/type  E32 F_XET enum=Xmtype
    /ch/04/mix/01/panFollow  E32 F_XET
/ch/04/mix/02  <CHME> n=0
    /ch/04/mix/02/on  E32 F_XET enum=OffOn
    /ch/04/mix/02/level  F32 F_XET
/ch/04/mix/03  <CHMO> n=0
    /ch/04/mix/03/on  E32 F_XET enum=OffOn
    /ch/04/mix/03/level  F32 F_XET
    /ch/04/mix/03/pan  F32 F_XET
    /ch/04/mix/03/type  E32 F_XET enum=Xmtype
    /ch/04/mix/03/panFollow  E32 F_XET
/ch/04/mix/04  <CHME> n=0
    /ch/04/mix/04/on  E32 F_XET enum=OffOn
    /ch/04/mix/04/level  F32 F_XET
/ch/04/mix/05  <CHMO> n=0
    /ch/04/mix/05/on  E32 F_XET enum=OffOn
    /ch/04/mix/05/level  F32 F_XET
    /ch/04/mix/05/pan  F32 F_XET
    /ch/04/mix/05/type  E32 F_XET enum=Xmtype
    /ch/04/mix/05/panFollow  E32 F_XET
/ch/04/mix/06  <CHME> n=0
    /ch/04/mix/06/on  E32 F_XET enum=OffOn
    /ch/04/mix/06/level  F32 F_XET
/ch/04/mix/07  <CHMO> n=0
    /ch/04/mix/07/on  E32 F_XET enum=OffOn
    /ch/04/mix/07/level  F32 F_XET
    /ch/04/mix/07/pan  F32 F_XET
    /ch/04/mix/07/type  E32 F_XET enum=Xmtype
    /ch/04/mix/07/panFollow  E32 F_XET
/ch/04/mix/08  <CHME> n=0
    /ch/04/mix/08/on  E32 F_XET enum=OffOn
    /ch/04/mix/08/level  F32 F_XET
/ch/04/mix/09  <CHMO> n=0
    /ch/04/mix/09/on  E32 F_XET enum=OffOn
    /ch/04/mix/09/level  F32 F_XET
    /ch/04/mix/09/pan  F32 F_XET
    /ch/04/mix/09/type  E32 F_XET enum=Xmtype
    /ch/04/mix/09/panFollow  E32 F_XET
/ch/04/mix/10  <CHME> n=0
    /ch/04/mix/10/on  E32 F_XET enum=OffOn
    /ch/04/mix/10/level  F32 F_XET
/ch/04/mix/11  <CHMO> n=0
    /ch/04/mix/11/on  E32 F_XET enum=OffOn
    /ch/04/mix/11/level  F32 F_XET
    /ch/04/mix/11/pan  F32 F_XET
    /ch/04/mix/11/type  E32 F_XET enum=Xmtype
    /ch/04/mix/11/panFollow  E32 F_XET
/ch/04/mix/12  <CHME> n=0
    /ch/04/mix/12/on  E32 F_XET enum=OffOn
    /ch/04/mix/12/level  F32 F_XET
/ch/04/mix/13  <CHMO> n=0
    /ch/04/mix/13/on  E32 F_XET enum=OffOn
    /ch/04/mix/13/level  F32 F_XET
    /ch/04/mix/13/pan  F32 F_XET
    /ch/04/mix/13/type  E32 F_XET enum=Xmtype
    /ch/04/mix/13/panFollow  E32 F_XET
/ch/04/mix/14  <CHME> n=0
    /ch/04/mix/14/on  E32 F_XET enum=OffOn
    /ch/04/mix/14/level  F32 F_XET
/ch/04/mix/15  <CHMO> n=0
    /ch/04/mix/15/on  E32 F_XET enum=OffOn
    /ch/04/mix/15/level  F32 F_XET
    /ch/04/mix/15/pan  F32 F_XET
    /ch/04/mix/15/type  E32 F_XET enum=Xmtype
    /ch/04/mix/15/panFollow  E32 F_XET
/ch/04/mix/16  <CHME> n=0
    /ch/04/mix/16/on  E32 F_XET enum=OffOn
    /ch/04/mix/16/level  F32 F_XET
/ch/04/automix  <CHAMIX> n=0
    /ch/04/automix/group  E32 F_XET enum=Xamxgrp
    /ch/04/automix/weight  F32 F_XET
```

### Xchannel05 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/05  <CHCO> n=0
/ch/05/config  <CHCO> n=0
    /ch/05/config/name  S32 F_XET
    /ch/05/config/icon  I32 F_XET
    /ch/05/config/color  E32 F_XET enum=Xcolors
    /ch/05/config/source  I32 F_XET
/ch/05/grp  <CHGRP> n=0
    /ch/05/grp/dca  P32 F_XET
    /ch/05/grp/mute  P32 F_XET
/ch/05/preamp  <CHPR> n=0
    /ch/05/preamp/trim  F32 F_XET
    /ch/05/preamp/invert  E32 F_XET enum=OffOn
    /ch/05/preamp/hpon  E32 F_XET enum=OffOn
    /ch/05/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/05/preamp/hpf  F32 F_XET
/ch/05/delay  <CHDE> n=0
    /ch/05/delay/on  E32 F_XET enum=OffOn
    /ch/05/delay/time  F32 F_XET
/ch/05/insert  <CHIN> n=0
    /ch/05/insert/on  E32 F_XET enum=OffOn
    /ch/05/insert/pos  E32 F_XET enum=Xdyppos
    /ch/05/insert/sel  E32 F_XET enum=Xisel
/ch/05/gate  <CHGA> n=0
    /ch/05/gate/on  E32 F_XET enum=OffOn
    /ch/05/gate/mode  E32 F_XET enum=Xgmode
    /ch/05/gate/thr  F32 F_XET
    /ch/05/gate/range  F32 F_XET
    /ch/05/gate/attack  F32 F_XET
    /ch/05/gate/hold  F32 F_XET
    /ch/05/gate/release  F32 F_XET
    /ch/05/gate/keysrc  I32 F_XET
/ch/05/gate/filter  <CHGF> n=0
    /ch/05/gate/filter/on  E32 F_XET enum=OffOn
    /ch/05/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/05/gate/filter/f  F32 F_XET
/ch/05/dyn  <CHDY> n=0
    /ch/05/dyn/on  E32 F_XET enum=OffOn
    /ch/05/dyn/mode  E32 F_XET enum=Xdymode
    /ch/05/dyn/det  E32 F_XET enum=Xdydet
    /ch/05/dyn/env  E32 F_XET enum=Xdyenv
    /ch/05/dyn/thr  F32 F_XET
    /ch/05/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/05/dyn/knee  F32 F_XET
    /ch/05/dyn/mgain  F32 F_XET
    /ch/05/dyn/attack  F32 F_XET
    /ch/05/dyn/hold  F32 F_XET
    /ch/05/dyn/release  F32 F_XET
    /ch/05/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/05/dyn/keysrc  I32 F_XET
    /ch/05/dyn/mix  F32 F_XET
    /ch/05/dyn/auto  E32 F_XET enum=OffOn
/ch/05/dyn/filter  <CHDF> n=0
    /ch/05/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/05/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/05/dyn/filter/f  F32 F_XET
/ch/05/eq  <OFFON> n=1
    /ch/05/eq/on  E32 F_XET enum=OffOn
/ch/05/eq/1  <CHEQ> n=0
    /ch/05/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/05/eq/1/f  F32 F_XET
    /ch/05/eq/1/g  F32 F_XET
    /ch/05/eq/1/q  F32 F_XET
/ch/05/eq/2  <CHEQ> n=0
    /ch/05/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/05/eq/2/f  F32 F_XET
    /ch/05/eq/2/g  F32 F_XET
    /ch/05/eq/2/q  F32 F_XET
/ch/05/eq/3  <CHEQ> n=0
    /ch/05/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/05/eq/3/f  F32 F_XET
    /ch/05/eq/3/g  F32 F_XET
    /ch/05/eq/3/q  F32 F_XET
/ch/05/eq/4  <CHEQ> n=0
    /ch/05/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/05/eq/4/f  F32 F_XET
    /ch/05/eq/4/g  F32 F_XET
    /ch/05/eq/4/q  F32 F_XET
/ch/05/mix  <CHMX> n=0
    /ch/05/mix/on  E32 F_XET enum=OffOn
    /ch/05/mix/fader  F32 F_XET
    /ch/05/mix/st  E32 F_XET enum=OffOn
    /ch/05/mix/pan  F32 F_XET
    /ch/05/mix/mono  E32 F_XET enum=OffOn
    /ch/05/mix/mlevel  F32 F_XET
/ch/05/mix/01  <CHMO> n=0
    /ch/05/mix/01/on  E32 F_XET enum=OffOn
    /ch/05/mix/01/level  F32 F_XET
    /ch/05/mix/01/pan  F32 F_XET
    /ch/05/mix/01/type  E32 F_XET enum=Xmtype
    /ch/05/mix/01/panFollow  E32 F_XET
/ch/05/mix/02  <CHME> n=0
    /ch/05/mix/02/on  E32 F_XET enum=OffOn
    /ch/05/mix/02/level  F32 F_XET
/ch/05/mix/03  <CHMO> n=0
    /ch/05/mix/03/on  E32 F_XET enum=OffOn
    /ch/05/mix/03/level  F32 F_XET
    /ch/05/mix/03/pan  F32 F_XET
    /ch/05/mix/03/type  E32 F_XET enum=Xmtype
    /ch/05/mix/03/panFollow  E32 F_XET
/ch/05/mix/04  <CHME> n=0
    /ch/05/mix/04/on  E32 F_XET enum=OffOn
    /ch/05/mix/04/level  F32 F_XET
/ch/05/mix/05  <CHMO> n=0
    /ch/05/mix/05/on  E32 F_XET enum=OffOn
    /ch/05/mix/05/level  F32 F_XET
    /ch/05/mix/05/pan  F32 F_XET
    /ch/05/mix/05/type  E32 F_XET enum=Xmtype
    /ch/05/mix/05/panFollow  E32 F_XET
/ch/05/mix/06  <CHME> n=0
    /ch/05/mix/06/on  E32 F_XET enum=OffOn
    /ch/05/mix/06/level  F32 F_XET
/ch/05/mix/07  <CHMO> n=0
    /ch/05/mix/07/on  E32 F_XET enum=OffOn
    /ch/05/mix/07/level  F32 F_XET
    /ch/05/mix/07/pan  F32 F_XET
    /ch/05/mix/07/type  E32 F_XET enum=Xmtype
    /ch/05/mix/07/panFollow  E32 F_XET
/ch/05/mix/08  <CHME> n=0
    /ch/05/mix/08/on  E32 F_XET enum=OffOn
    /ch/05/mix/08/level  F32 F_XET
/ch/05/mix/09  <CHMO> n=0
    /ch/05/mix/09/on  E32 F_XET enum=OffOn
    /ch/05/mix/09/level  F32 F_XET
    /ch/05/mix/09/pan  F32 F_XET
    /ch/05/mix/09/type  E32 F_XET enum=Xmtype
    /ch/05/mix/09/panFollow  E32 F_XET
/ch/05/mix/10  <CHME> n=0
    /ch/05/mix/10/on  E32 F_XET enum=OffOn
    /ch/05/mix/10/level  F32 F_XET
/ch/05/mix/11  <CHMO> n=0
    /ch/05/mix/11/on  E32 F_XET enum=OffOn
    /ch/05/mix/11/level  F32 F_XET
    /ch/05/mix/11/pan  F32 F_XET
    /ch/05/mix/11/type  E32 F_XET enum=Xmtype
    /ch/05/mix/11/panFollow  E32 F_XET
/ch/05/mix/12  <CHME> n=0
    /ch/05/mix/12/on  E32 F_XET enum=OffOn
    /ch/05/mix/12/level  F32 F_XET
/ch/05/mix/13  <CHMO> n=0
    /ch/05/mix/13/on  E32 F_XET enum=OffOn
    /ch/05/mix/13/level  F32 F_XET
    /ch/05/mix/13/pan  F32 F_XET
    /ch/05/mix/13/type  E32 F_XET enum=Xmtype
    /ch/05/mix/13/panFollow  E32 F_XET
/ch/05/mix/14  <CHME> n=0
    /ch/05/mix/14/on  E32 F_XET enum=OffOn
    /ch/05/mix/14/level  F32 F_XET
/ch/05/mix/15  <CHMO> n=0
    /ch/05/mix/15/on  E32 F_XET enum=OffOn
    /ch/05/mix/15/level  F32 F_XET
    /ch/05/mix/15/pan  F32 F_XET
    /ch/05/mix/15/type  E32 F_XET enum=Xmtype
    /ch/05/mix/15/panFollow  E32 F_XET
/ch/05/mix/16  <CHME> n=0
    /ch/05/mix/16/on  E32 F_XET enum=OffOn
    /ch/05/mix/16/level  F32 F_XET
/ch/05/automix  <CHAMIX> n=0
    /ch/05/automix/group  E32 F_XET enum=Xamxgrp
    /ch/05/automix/weight  F32 F_XET
```

### Xchannel06 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/06  <CHCO> n=0
/ch/06/config  <CHCO> n=0
    /ch/06/config/name  S32 F_XET
    /ch/06/config/icon  I32 F_XET
    /ch/06/config/color  E32 F_XET enum=Xcolors
    /ch/06/config/source  I32 F_XET
/ch/06/grp  <CHGRP> n=0
    /ch/06/grp/dca  P32 F_XET
    /ch/06/grp/mute  P32 F_XET
/ch/06/preamp  <CHPR> n=0
    /ch/06/preamp/trim  F32 F_XET
    /ch/06/preamp/invert  E32 F_XET enum=OffOn
    /ch/06/preamp/hpon  E32 F_XET enum=OffOn
    /ch/06/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/06/preamp/hpf  F32 F_XET
/ch/06/delay  <CHDE> n=0
    /ch/06/delay/on  E32 F_XET enum=OffOn
    /ch/06/delay/time  F32 F_XET
/ch/06/insert  <CHIN> n=0
    /ch/06/insert/on  E32 F_XET enum=OffOn
    /ch/06/insert/pos  E32 F_XET enum=Xdyppos
    /ch/06/insert/sel  E32 F_XET enum=Xisel
/ch/06/gate  <CHGA> n=0
    /ch/06/gate/on  E32 F_XET enum=OffOn
    /ch/06/gate/mode  E32 F_XET enum=Xgmode
    /ch/06/gate/thr  F32 F_XET
    /ch/06/gate/range  F32 F_XET
    /ch/06/gate/attack  F32 F_XET
    /ch/06/gate/hold  F32 F_XET
    /ch/06/gate/release  F32 F_XET
    /ch/06/gate/keysrc  I32 F_XET
/ch/06/gate/filter  <CHGF> n=0
    /ch/06/gate/filter/on  E32 F_XET enum=OffOn
    /ch/06/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/06/gate/filter/f  F32 F_XET
/ch/06/dyn  <CHDY> n=0
    /ch/06/dyn/on  E32 F_XET enum=OffOn
    /ch/06/dyn/mode  E32 F_XET enum=Xdymode
    /ch/06/dyn/det  E32 F_XET enum=Xdydet
    /ch/06/dyn/env  E32 F_XET enum=Xdyenv
    /ch/06/dyn/thr  F32 F_XET
    /ch/06/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/06/dyn/knee  F32 F_XET
    /ch/06/dyn/mgain  F32 F_XET
    /ch/06/dyn/attack  F32 F_XET
    /ch/06/dyn/hold  F32 F_XET
    /ch/06/dyn/release  F32 F_XET
    /ch/06/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/06/dyn/keysrc  I32 F_XET
    /ch/06/dyn/mix  F32 F_XET
    /ch/06/dyn/auto  E32 F_XET enum=OffOn
/ch/06/dyn/filter  <CHDF> n=0
    /ch/06/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/06/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/06/dyn/filter/f  F32 F_XET
/ch/06/eq  <OFFON> n=1
    /ch/06/eq/on  E32 F_XET enum=OffOn
/ch/06/eq/1  <CHEQ> n=0
    /ch/06/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/06/eq/1/f  F32 F_XET
    /ch/06/eq/1/g  F32 F_XET
    /ch/06/eq/1/q  F32 F_XET
/ch/06/eq/2  <CHEQ> n=0
    /ch/06/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/06/eq/2/f  F32 F_XET
    /ch/06/eq/2/g  F32 F_XET
    /ch/06/eq/2/q  F32 F_XET
/ch/06/eq/3  <CHEQ> n=0
    /ch/06/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/06/eq/3/f  F32 F_XET
    /ch/06/eq/3/g  F32 F_XET
    /ch/06/eq/3/q  F32 F_XET
/ch/06/eq/4  <CHEQ> n=0
    /ch/06/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/06/eq/4/f  F32 F_XET
    /ch/06/eq/4/g  F32 F_XET
    /ch/06/eq/4/q  F32 F_XET
/ch/06/mix  <CHMX> n=0
    /ch/06/mix/on  E32 F_XET enum=OffOn
    /ch/06/mix/fader  F32 F_XET
    /ch/06/mix/st  E32 F_XET enum=OffOn
    /ch/06/mix/pan  F32 F_XET
    /ch/06/mix/mono  E32 F_XET enum=OffOn
    /ch/06/mix/mlevel  F32 F_XET
/ch/06/mix/01  <CHMO> n=0
    /ch/06/mix/01/on  E32 F_XET enum=OffOn
    /ch/06/mix/01/level  F32 F_XET
    /ch/06/mix/01/pan  F32 F_XET
    /ch/06/mix/01/type  E32 F_XET enum=Xmtype
    /ch/06/mix/01/panFollow  E32 F_XET
/ch/06/mix/02  <CHME> n=0
    /ch/06/mix/02/on  E32 F_XET enum=OffOn
    /ch/06/mix/02/level  F32 F_XET
/ch/06/mix/03  <CHMO> n=0
    /ch/06/mix/03/on  E32 F_XET enum=OffOn
    /ch/06/mix/03/level  F32 F_XET
    /ch/06/mix/03/pan  F32 F_XET
    /ch/06/mix/03/type  E32 F_XET enum=Xmtype
    /ch/06/mix/03/panFollow  E32 F_XET
/ch/06/mix/04  <CHME> n=0
    /ch/06/mix/04/on  E32 F_XET enum=OffOn
    /ch/06/mix/04/level  F32 F_XET
/ch/06/mix/05  <CHMO> n=0
    /ch/06/mix/05/on  E32 F_XET enum=OffOn
    /ch/06/mix/05/level  F32 F_XET
    /ch/06/mix/05/pan  F32 F_XET
    /ch/06/mix/05/type  E32 F_XET enum=Xmtype
    /ch/06/mix/05/panFollow  E32 F_XET
/ch/06/mix/06  <CHME> n=0
    /ch/06/mix/06/on  E32 F_XET enum=OffOn
    /ch/06/mix/06/level  F32 F_XET
/ch/06/mix/07  <CHMO> n=0
    /ch/06/mix/07/on  E32 F_XET enum=OffOn
    /ch/06/mix/07/level  F32 F_XET
    /ch/06/mix/07/pan  F32 F_XET
    /ch/06/mix/07/type  E32 F_XET enum=Xmtype
    /ch/06/mix/07/panFollow  E32 F_XET
/ch/06/mix/08  <CHME> n=0
    /ch/06/mix/08/on  E32 F_XET enum=OffOn
    /ch/06/mix/08/level  F32 F_XET
/ch/06/mix/09  <CHMO> n=0
    /ch/06/mix/09/on  E32 F_XET enum=OffOn
    /ch/06/mix/09/level  F32 F_XET
    /ch/06/mix/09/pan  F32 F_XET
    /ch/06/mix/09/type  E32 F_XET enum=Xmtype
    /ch/06/mix/09/panFollow  E32 F_XET
/ch/06/mix/10  <CHME> n=0
    /ch/06/mix/10/on  E32 F_XET enum=OffOn
    /ch/06/mix/10/level  F32 F_XET
/ch/06/mix/11  <CHMO> n=0
    /ch/06/mix/11/on  E32 F_XET enum=OffOn
    /ch/06/mix/11/level  F32 F_XET
    /ch/06/mix/11/pan  F32 F_XET
    /ch/06/mix/11/type  E32 F_XET enum=Xmtype
    /ch/06/mix/11/panFollow  E32 F_XET
/ch/06/mix/12  <CHME> n=0
    /ch/06/mix/12/on  E32 F_XET enum=OffOn
    /ch/06/mix/12/level  F32 F_XET
/ch/06/mix/13  <CHMO> n=0
    /ch/06/mix/13/on  E32 F_XET enum=OffOn
    /ch/06/mix/13/level  F32 F_XET
    /ch/06/mix/13/pan  F32 F_XET
    /ch/06/mix/13/type  E32 F_XET enum=Xmtype
    /ch/06/mix/13/panFollow  E32 F_XET
/ch/06/mix/14  <CHME> n=0
    /ch/06/mix/14/on  E32 F_XET enum=OffOn
    /ch/06/mix/14/level  F32 F_XET
/ch/06/mix/15  <CHMO> n=0
    /ch/06/mix/15/on  E32 F_XET enum=OffOn
    /ch/06/mix/15/level  F32 F_XET
    /ch/06/mix/15/pan  F32 F_XET
    /ch/06/mix/15/type  E32 F_XET enum=Xmtype
    /ch/06/mix/15/panFollow  E32 F_XET
/ch/06/mix/16  <CHME> n=0
    /ch/06/mix/16/on  E32 F_XET enum=OffOn
    /ch/06/mix/16/level  F32 F_XET
/ch/06/automix  <CHAMIX> n=0
    /ch/06/automix/group  E32 F_XET enum=Xamxgrp
    /ch/06/automix/weight  F32 F_XET
```

### Xchannel07 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/07  <CHCO> n=0
/ch/07/config  <CHCO> n=0
    /ch/07/config/name  S32 F_XET
    /ch/07/config/icon  I32 F_XET
    /ch/07/config/color  E32 F_XET enum=Xcolors
    /ch/07/config/source  I32 F_XET
/ch/07/grp  <CHGRP> n=0
    /ch/07/grp/dca  P32 F_XET
    /ch/07/grp/mute  P32 F_XET
/ch/07/preamp  <CHPR> n=0
    /ch/07/preamp/trim  F32 F_XET
    /ch/07/preamp/invert  E32 F_XET enum=OffOn
    /ch/07/preamp/hpon  E32 F_XET enum=OffOn
    /ch/07/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/07/preamp/hpf  F32 F_XET
/ch/07/delay  <CHDE> n=0
    /ch/07/delay/on  E32 F_XET enum=OffOn
    /ch/07/delay/time  F32 F_XET
/ch/07/insert  <CHIN> n=0
    /ch/07/insert/on  E32 F_XET enum=OffOn
    /ch/07/insert/pos  E32 F_XET enum=Xdyppos
    /ch/07/insert/sel  E32 F_XET enum=Xisel
/ch/07/gate  <CHGA> n=0
    /ch/07/gate/on  E32 F_XET enum=OffOn
    /ch/07/gate/mode  E32 F_XET enum=Xgmode
    /ch/07/gate/thr  F32 F_XET
    /ch/07/gate/range  F32 F_XET
    /ch/07/gate/attack  F32 F_XET
    /ch/07/gate/hold  F32 F_XET
    /ch/07/gate/release  F32 F_XET
    /ch/07/gate/keysrc  I32 F_XET
/ch/07/gate/filter  <CHGF> n=0
    /ch/07/gate/filter/on  E32 F_XET enum=OffOn
    /ch/07/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/07/gate/filter/f  F32 F_XET
/ch/07/dyn  <CHDY> n=0
    /ch/07/dyn/on  E32 F_XET enum=OffOn
    /ch/07/dyn/mode  E32 F_XET enum=Xdymode
    /ch/07/dyn/det  E32 F_XET enum=Xdydet
    /ch/07/dyn/env  E32 F_XET enum=Xdyenv
    /ch/07/dyn/thr  F32 F_XET
    /ch/07/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/07/dyn/knee  F32 F_XET
    /ch/07/dyn/mgain  F32 F_XET
    /ch/07/dyn/attack  F32 F_XET
    /ch/07/dyn/hold  F32 F_XET
    /ch/07/dyn/release  F32 F_XET
    /ch/07/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/07/dyn/keysrc  I32 F_XET
    /ch/07/dyn/mix  F32 F_XET
    /ch/07/dyn/auto  E32 F_XET enum=OffOn
/ch/07/dyn/filter  <CHDF> n=0
    /ch/07/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/07/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/07/dyn/filter/f  F32 F_XET
/ch/07/eq  <OFFON> n=1
    /ch/07/eq/on  E32 F_XET enum=OffOn
/ch/07/eq/1  <CHEQ> n=0
    /ch/07/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/07/eq/1/f  F32 F_XET
    /ch/07/eq/1/g  F32 F_XET
    /ch/07/eq/1/q  F32 F_XET
/ch/07/eq/2  <CHEQ> n=0
    /ch/07/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/07/eq/2/f  F32 F_XET
    /ch/07/eq/2/g  F32 F_XET
    /ch/07/eq/2/q  F32 F_XET
/ch/07/eq/3  <CHEQ> n=0
    /ch/07/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/07/eq/3/f  F32 F_XET
    /ch/07/eq/3/g  F32 F_XET
    /ch/07/eq/3/q  F32 F_XET
/ch/07/eq/4  <CHEQ> n=0
    /ch/07/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/07/eq/4/f  F32 F_XET
    /ch/07/eq/4/g  F32 F_XET
    /ch/07/eq/4/q  F32 F_XET
/ch/07/mix  <CHMX> n=0
    /ch/07/mix/on  E32 F_XET enum=OffOn
    /ch/07/mix/fader  F32 F_XET
    /ch/07/mix/st  E32 F_XET enum=OffOn
    /ch/07/mix/pan  F32 F_XET
    /ch/07/mix/mono  E32 F_XET enum=OffOn
    /ch/07/mix/mlevel  F32 F_XET
/ch/07/mix/01  <CHMO> n=0
    /ch/07/mix/01/on  E32 F_XET enum=OffOn
    /ch/07/mix/01/level  F32 F_XET
    /ch/07/mix/01/pan  F32 F_XET
    /ch/07/mix/01/type  E32 F_XET enum=Xmtype
    /ch/07/mix/01/panFollow  E32 F_XET
/ch/07/mix/02  <CHME> n=0
    /ch/07/mix/02/on  E32 F_XET enum=OffOn
    /ch/07/mix/02/level  F32 F_XET
/ch/07/mix/03  <CHMO> n=0
    /ch/07/mix/03/on  E32 F_XET enum=OffOn
    /ch/07/mix/03/level  F32 F_XET
    /ch/07/mix/03/pan  F32 F_XET
    /ch/07/mix/03/type  E32 F_XET enum=Xmtype
    /ch/07/mix/03/panFollow  E32 F_XET
/ch/07/mix/04  <CHME> n=0
    /ch/07/mix/04/on  E32 F_XET enum=OffOn
    /ch/07/mix/04/level  F32 F_XET
/ch/07/mix/05  <CHMO> n=0
    /ch/07/mix/05/on  E32 F_XET enum=OffOn
    /ch/07/mix/05/level  F32 F_XET
    /ch/07/mix/05/pan  F32 F_XET
    /ch/07/mix/05/type  E32 F_XET enum=Xmtype
    /ch/07/mix/05/panFollow  E32 F_XET
/ch/07/mix/06  <CHME> n=0
    /ch/07/mix/06/on  E32 F_XET enum=OffOn
    /ch/07/mix/06/level  F32 F_XET
/ch/07/mix/07  <CHMO> n=0
    /ch/07/mix/07/on  E32 F_XET enum=OffOn
    /ch/07/mix/07/level  F32 F_XET
    /ch/07/mix/07/pan  F32 F_XET
    /ch/07/mix/07/type  E32 F_XET enum=Xmtype
    /ch/07/mix/07/panFollow  E32 F_XET
/ch/07/mix/08  <CHME> n=0
    /ch/07/mix/08/on  E32 F_XET enum=OffOn
    /ch/07/mix/08/level  F32 F_XET
/ch/07/mix/09  <CHMO> n=0
    /ch/07/mix/09/on  E32 F_XET enum=OffOn
    /ch/07/mix/09/level  F32 F_XET
    /ch/07/mix/09/pan  F32 F_XET
    /ch/07/mix/09/type  E32 F_XET enum=Xmtype
    /ch/07/mix/09/panFollow  E32 F_XET
/ch/07/mix/10  <CHME> n=0
    /ch/07/mix/10/on  E32 F_XET enum=OffOn
    /ch/07/mix/10/level  F32 F_XET
/ch/07/mix/11  <CHMO> n=0
    /ch/07/mix/11/on  E32 F_XET enum=OffOn
    /ch/07/mix/11/level  F32 F_XET
    /ch/07/mix/11/pan  F32 F_XET
    /ch/07/mix/11/type  E32 F_XET enum=Xmtype
    /ch/07/mix/11/panFollow  E32 F_XET
/ch/07/mix/12  <CHME> n=0
    /ch/07/mix/12/on  E32 F_XET enum=OffOn
    /ch/07/mix/12/level  F32 F_XET
/ch/07/mix/13  <CHMO> n=0
    /ch/07/mix/13/on  E32 F_XET enum=OffOn
    /ch/07/mix/13/level  F32 F_XET
    /ch/07/mix/13/pan  F32 F_XET
    /ch/07/mix/13/type  E32 F_XET enum=Xmtype
    /ch/07/mix/13/panFollow  E32 F_XET
/ch/07/mix/14  <CHME> n=0
    /ch/07/mix/14/on  E32 F_XET enum=OffOn
    /ch/07/mix/14/level  F32 F_XET
/ch/07/mix/15  <CHMO> n=0
    /ch/07/mix/15/on  E32 F_XET enum=OffOn
    /ch/07/mix/15/level  F32 F_XET
    /ch/07/mix/15/pan  F32 F_XET
    /ch/07/mix/15/type  E32 F_XET enum=Xmtype
    /ch/07/mix/15/panFollow  E32 F_XET
/ch/07/mix/16  <CHME> n=0
    /ch/07/mix/16/on  E32 F_XET enum=OffOn
    /ch/07/mix/16/level  F32 F_XET
/ch/07/automix  <CHAMIX> n=0
    /ch/07/automix/group  E32 F_XET enum=Xamxgrp
    /ch/07/automix/weight  F32 F_XET
```

### Xchannel08 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/08  <CHCO> n=0
/ch/08/config  <CHCO> n=0
    /ch/08/config/name  S32 F_XET
    /ch/08/config/icon  I32 F_XET
    /ch/08/config/color  E32 F_XET enum=Xcolors
    /ch/08/config/source  I32 F_XET
/ch/08/grp  <CHGRP> n=0
    /ch/08/grp/dca  P32 F_XET
    /ch/08/grp/mute  P32 F_XET
/ch/08/preamp  <CHPR> n=0
    /ch/08/preamp/trim  F32 F_XET
    /ch/08/preamp/invert  E32 F_XET enum=OffOn
    /ch/08/preamp/hpon  E32 F_XET enum=OffOn
    /ch/08/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/08/preamp/hpf  F32 F_XET
/ch/08/delay  <CHDE> n=0
    /ch/08/delay/on  E32 F_XET enum=OffOn
    /ch/08/delay/time  F32 F_XET
/ch/08/insert  <CHIN> n=0
    /ch/08/insert/on  E32 F_XET enum=OffOn
    /ch/08/insert/pos  E32 F_XET enum=Xdyppos
    /ch/08/insert/sel  E32 F_XET enum=Xisel
/ch/08/gate  <CHGA> n=0
    /ch/08/gate/on  E32 F_XET enum=OffOn
    /ch/08/gate/mode  E32 F_XET enum=Xgmode
    /ch/08/gate/thr  F32 F_XET
    /ch/08/gate/range  F32 F_XET
    /ch/08/gate/attack  F32 F_XET
    /ch/08/gate/hold  F32 F_XET
    /ch/08/gate/release  F32 F_XET
    /ch/08/gate/keysrc  I32 F_XET
/ch/08/gate/filter  <CHGF> n=0
    /ch/08/gate/filter/on  E32 F_XET enum=OffOn
    /ch/08/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/08/gate/filter/f  F32 F_XET
/ch/08/dyn  <CHDY> n=0
    /ch/08/dyn/on  E32 F_XET enum=OffOn
    /ch/08/dyn/mode  E32 F_XET enum=Xdymode
    /ch/08/dyn/det  E32 F_XET enum=Xdydet
    /ch/08/dyn/env  E32 F_XET enum=Xdyenv
    /ch/08/dyn/thr  F32 F_XET
    /ch/08/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/08/dyn/knee  F32 F_XET
    /ch/08/dyn/mgain  F32 F_XET
    /ch/08/dyn/attack  F32 F_XET
    /ch/08/dyn/hold  F32 F_XET
    /ch/08/dyn/release  F32 F_XET
    /ch/08/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/08/dyn/keysrc  I32 F_XET
    /ch/08/dyn/mix  F32 F_XET
    /ch/08/dyn/auto  E32 F_XET enum=OffOn
/ch/08/dyn/filter  <CHDF> n=0
    /ch/08/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/08/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/08/dyn/filter/f  F32 F_XET
/ch/08/eq  <OFFON> n=1
    /ch/08/eq/on  E32 F_XET enum=OffOn
/ch/08/eq/1  <CHEQ> n=0
    /ch/08/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/08/eq/1/f  F32 F_XET
    /ch/08/eq/1/g  F32 F_XET
    /ch/08/eq/1/q  F32 F_XET
/ch/08/eq/2  <CHEQ> n=0
    /ch/08/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/08/eq/2/f  F32 F_XET
    /ch/08/eq/2/g  F32 F_XET
    /ch/08/eq/2/q  F32 F_XET
/ch/08/eq/3  <CHEQ> n=0
    /ch/08/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/08/eq/3/f  F32 F_XET
    /ch/08/eq/3/g  F32 F_XET
    /ch/08/eq/3/q  F32 F_XET
/ch/08/eq/4  <CHEQ> n=0
    /ch/08/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/08/eq/4/f  F32 F_XET
    /ch/08/eq/4/g  F32 F_XET
    /ch/08/eq/4/q  F32 F_XET
/ch/08/mix  <CHMX> n=0
    /ch/08/mix/on  E32 F_XET enum=OffOn
    /ch/08/mix/fader  F32 F_XET
    /ch/08/mix/st  E32 F_XET enum=OffOn
    /ch/08/mix/pan  F32 F_XET
    /ch/08/mix/mono  E32 F_XET enum=OffOn
    /ch/08/mix/mlevel  F32 F_XET
/ch/08/mix/01  <CHMO> n=0
    /ch/08/mix/01/on  E32 F_XET enum=OffOn
    /ch/08/mix/01/level  F32 F_XET
    /ch/08/mix/01/pan  F32 F_XET
    /ch/08/mix/01/type  E32 F_XET enum=Xmtype
    /ch/08/mix/01/panFollow  E32 F_XET
/ch/08/mix/02  <CHME> n=0
    /ch/08/mix/02/on  E32 F_XET enum=OffOn
    /ch/08/mix/02/level  F32 F_XET
/ch/08/mix/03  <CHMO> n=0
    /ch/08/mix/03/on  E32 F_XET enum=OffOn
    /ch/08/mix/03/level  F32 F_XET
    /ch/08/mix/03/pan  F32 F_XET
    /ch/08/mix/03/type  E32 F_XET enum=Xmtype
    /ch/08/mix/03/panFollow  E32 F_XET
/ch/08/mix/04  <CHME> n=0
    /ch/08/mix/04/on  E32 F_XET enum=OffOn
    /ch/08/mix/04/level  F32 F_XET
/ch/08/mix/05  <CHMO> n=0
    /ch/08/mix/05/on  E32 F_XET enum=OffOn
    /ch/08/mix/05/level  F32 F_XET
    /ch/08/mix/05/pan  F32 F_XET
    /ch/08/mix/05/type  E32 F_XET enum=Xmtype
    /ch/08/mix/05/panFollow  E32 F_XET
/ch/08/mix/06  <CHME> n=0
    /ch/08/mix/06/on  E32 F_XET enum=OffOn
    /ch/08/mix/06/level  F32 F_XET
/ch/08/mix/07  <CHMO> n=0
    /ch/08/mix/07/on  E32 F_XET enum=OffOn
    /ch/08/mix/07/level  F32 F_XET
    /ch/08/mix/07/pan  F32 F_XET
    /ch/08/mix/07/type  E32 F_XET enum=Xmtype
    /ch/08/mix/07/panFollow  E32 F_XET
/ch/08/mix/08  <CHME> n=0
    /ch/08/mix/08/on  E32 F_XET enum=OffOn
    /ch/08/mix/08/level  F32 F_XET
/ch/08/mix/09  <CHMO> n=0
    /ch/08/mix/09/on  E32 F_XET enum=OffOn
    /ch/08/mix/09/level  F32 F_XET
    /ch/08/mix/09/pan  F32 F_XET
    /ch/08/mix/09/type  E32 F_XET enum=Xmtype
    /ch/08/mix/09/panFollow  E32 F_XET
/ch/08/mix/10  <CHME> n=0
    /ch/08/mix/10/on  E32 F_XET enum=OffOn
    /ch/08/mix/10/level  F32 F_XET
/ch/08/mix/11  <CHMO> n=0
    /ch/08/mix/11/on  E32 F_XET enum=OffOn
    /ch/08/mix/11/level  F32 F_XET
    /ch/08/mix/11/pan  F32 F_XET
    /ch/08/mix/11/type  E32 F_XET enum=Xmtype
    /ch/08/mix/11/panFollow  E32 F_XET
/ch/08/mix/12  <CHME> n=0
    /ch/08/mix/12/on  E32 F_XET enum=OffOn
    /ch/08/mix/12/level  F32 F_XET
/ch/08/mix/13  <CHMO> n=0
    /ch/08/mix/13/on  E32 F_XET enum=OffOn
    /ch/08/mix/13/level  F32 F_XET
    /ch/08/mix/13/pan  F32 F_XET
    /ch/08/mix/13/type  E32 F_XET enum=Xmtype
    /ch/08/mix/13/panFollow  E32 F_XET
/ch/08/mix/14  <CHME> n=0
    /ch/08/mix/14/on  E32 F_XET enum=OffOn
    /ch/08/mix/14/level  F32 F_XET
/ch/08/mix/15  <CHMO> n=0
    /ch/08/mix/15/on  E32 F_XET enum=OffOn
    /ch/08/mix/15/level  F32 F_XET
    /ch/08/mix/15/pan  F32 F_XET
    /ch/08/mix/15/type  E32 F_XET enum=Xmtype
    /ch/08/mix/15/panFollow  E32 F_XET
/ch/08/mix/16  <CHME> n=0
    /ch/08/mix/16/on  E32 F_XET enum=OffOn
    /ch/08/mix/16/level  F32 F_XET
/ch/08/automix  <CHAMIX> n=0
    /ch/08/automix/group  E32 F_XET enum=Xamxgrp
    /ch/08/automix/weight  F32 F_XET
```

### Xchannel09 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/09  <CHCO> n=0
/ch/09/config  <CHCO> n=0
    /ch/09/config/name  S32 F_XET
    /ch/09/config/icon  I32 F_XET
    /ch/09/config/color  E32 F_XET enum=Xcolors
    /ch/09/config/source  I32 F_XET
/ch/09/grp  <CHGRP> n=0
    /ch/09/grp/dca  P32 F_XET
    /ch/09/grp/mute  P32 F_XET
/ch/09/preamp  <CHPR> n=0
    /ch/09/preamp/trim  F32 F_XET
    /ch/09/preamp/invert  E32 F_XET enum=OffOn
    /ch/09/preamp/hpon  E32 F_XET enum=OffOn
    /ch/09/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/09/preamp/hpf  F32 F_XET
/ch/09/delay  <CHDE> n=0
    /ch/09/delay/on  E32 F_XET enum=OffOn
    /ch/09/delay/time  F32 F_XET
/ch/09/insert  <CHIN> n=0
    /ch/09/insert/on  E32 F_XET enum=OffOn
    /ch/09/insert/pos  E32 F_XET enum=Xdyppos
    /ch/09/insert/sel  E32 F_XET enum=Xisel
/ch/09/gate  <CHGA> n=0
    /ch/09/gate/on  E32 F_XET enum=OffOn
    /ch/09/gate/mode  E32 F_XET enum=Xgmode
    /ch/09/gate/thr  F32 F_XET
    /ch/09/gate/range  F32 F_XET
    /ch/09/gate/attack  F32 F_XET
    /ch/09/gate/hold  F32 F_XET
    /ch/09/gate/release  F32 F_XET
    /ch/09/gate/keysrc  I32 F_XET
/ch/09/gate/filter  <CHGF> n=0
    /ch/09/gate/filter/on  E32 F_XET enum=OffOn
    /ch/09/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/09/gate/filter/f  F32 F_XET
/ch/09/dyn  <CHDY> n=0
    /ch/09/dyn/on  E32 F_XET enum=OffOn
    /ch/09/dyn/mode  E32 F_XET enum=Xdymode
    /ch/09/dyn/det  E32 F_XET enum=Xdydet
    /ch/09/dyn/env  E32 F_XET enum=Xdyenv
    /ch/09/dyn/thr  F32 F_XET
    /ch/09/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/09/dyn/knee  F32 F_XET
    /ch/09/dyn/mgain  F32 F_XET
    /ch/09/dyn/attack  F32 F_XET
    /ch/09/dyn/hold  F32 F_XET
    /ch/09/dyn/release  F32 F_XET
    /ch/09/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/09/dyn/keysrc  I32 F_XET
    /ch/09/dyn/mix  F32 F_XET
    /ch/09/dyn/auto  E32 F_XET enum=OffOn
/ch/09/dyn/filter  <CHDF> n=0
    /ch/09/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/09/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/09/dyn/filter/f  F32 F_XET
/ch/09/eq  <OFFON> n=1
    /ch/09/eq/on  E32 F_XET enum=OffOn
/ch/09/eq/1  <CHEQ> n=0
    /ch/09/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/09/eq/1/f  F32 F_XET
    /ch/09/eq/1/g  F32 F_XET
    /ch/09/eq/1/q  F32 F_XET
/ch/09/eq/2  <CHEQ> n=0
    /ch/09/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/09/eq/2/f  F32 F_XET
    /ch/09/eq/2/g  F32 F_XET
    /ch/09/eq/2/q  F32 F_XET
/ch/09/eq/3  <CHEQ> n=0
    /ch/09/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/09/eq/3/f  F32 F_XET
    /ch/09/eq/3/g  F32 F_XET
    /ch/09/eq/3/q  F32 F_XET
/ch/09/eq/4  <CHEQ> n=0
    /ch/09/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/09/eq/4/f  F32 F_XET
    /ch/09/eq/4/g  F32 F_XET
    /ch/09/eq/4/q  F32 F_XET
/ch/09/mix  <CHMX> n=0
    /ch/09/mix/on  E32 F_XET enum=OffOn
    /ch/09/mix/fader  F32 F_XET
    /ch/09/mix/st  E32 F_XET enum=OffOn
    /ch/09/mix/pan  F32 F_XET
    /ch/09/mix/mono  E32 F_XET enum=OffOn
    /ch/09/mix/mlevel  F32 F_XET
/ch/09/mix/01  <CHMO> n=0
    /ch/09/mix/01/on  E32 F_XET enum=OffOn
    /ch/09/mix/01/level  F32 F_XET
    /ch/09/mix/01/pan  F32 F_XET
    /ch/09/mix/01/type  E32 F_XET enum=Xmtype
    /ch/09/mix/01/panFollow  E32 F_XET
/ch/09/mix/02  <CHME> n=0
    /ch/09/mix/02/on  E32 F_XET enum=OffOn
    /ch/09/mix/02/level  F32 F_XET
/ch/09/mix/03  <CHMO> n=0
    /ch/09/mix/03/on  E32 F_XET enum=OffOn
    /ch/09/mix/03/level  F32 F_XET
    /ch/09/mix/03/pan  F32 F_XET
    /ch/09/mix/03/type  E32 F_XET enum=Xmtype
    /ch/09/mix/03/panFollow  E32 F_XET
/ch/09/mix/04  <CHME> n=0
    /ch/09/mix/04/on  E32 F_XET enum=OffOn
    /ch/09/mix/04/level  F32 F_XET
/ch/09/mix/05  <CHMO> n=0
    /ch/09/mix/05/on  E32 F_XET enum=OffOn
    /ch/09/mix/05/level  F32 F_XET
    /ch/09/mix/05/pan  F32 F_XET
    /ch/09/mix/05/type  E32 F_XET enum=Xmtype
    /ch/09/mix/05/panFollow  E32 F_XET
/ch/09/mix/06  <CHME> n=0
    /ch/09/mix/06/on  E32 F_XET enum=OffOn
    /ch/09/mix/06/level  F32 F_XET
/ch/09/mix/07  <CHMO> n=0
    /ch/09/mix/07/on  E32 F_XET enum=OffOn
    /ch/09/mix/07/level  F32 F_XET
    /ch/09/mix/07/pan  F32 F_XET
    /ch/09/mix/07/type  E32 F_XET enum=Xmtype
    /ch/09/mix/07/panFollow  E32 F_XET
/ch/09/mix/08  <CHME> n=0
    /ch/09/mix/08/on  E32 F_XET enum=OffOn
    /ch/09/mix/08/level  F32 F_XET
/ch/09/mix/09  <CHMO> n=0
    /ch/09/mix/09/on  E32 F_XET enum=OffOn
    /ch/09/mix/09/level  F32 F_XET
    /ch/09/mix/09/pan  F32 F_XET
    /ch/09/mix/09/type  E32 F_XET enum=Xmtype
    /ch/09/mix/09/panFollow  E32 F_XET
/ch/09/mix/10  <CHME> n=0
    /ch/09/mix/10/on  E32 F_XET enum=OffOn
    /ch/09/mix/10/level  F32 F_XET
/ch/09/mix/11  <CHMO> n=0
    /ch/09/mix/11/on  E32 F_XET enum=OffOn
    /ch/09/mix/11/level  F32 F_XET
    /ch/09/mix/11/pan  F32 F_XET
    /ch/09/mix/11/type  E32 F_XET enum=Xmtype
    /ch/09/mix/11/panFollow  E32 F_XET
/ch/09/mix/12  <CHME> n=0
    /ch/09/mix/12/on  E32 F_XET enum=OffOn
    /ch/09/mix/12/level  F32 F_XET
/ch/09/mix/13  <CHMO> n=0
    /ch/09/mix/13/on  E32 F_XET enum=OffOn
    /ch/09/mix/13/level  F32 F_XET
    /ch/09/mix/13/pan  F32 F_XET
    /ch/09/mix/13/type  E32 F_XET enum=Xmtype
    /ch/09/mix/13/panFollow  E32 F_XET
/ch/09/mix/14  <CHME> n=0
    /ch/09/mix/14/on  E32 F_XET enum=OffOn
    /ch/09/mix/14/level  F32 F_XET
/ch/09/mix/15  <CHMO> n=0
    /ch/09/mix/15/on  E32 F_XET enum=OffOn
    /ch/09/mix/15/level  F32 F_XET
    /ch/09/mix/15/pan  F32 F_XET
    /ch/09/mix/15/type  E32 F_XET enum=Xmtype
    /ch/09/mix/15/panFollow  E32 F_XET
/ch/09/mix/16  <CHME> n=0
    /ch/09/mix/16/on  E32 F_XET enum=OffOn
    /ch/09/mix/16/level  F32 F_XET
/ch/09/automix  <CHAMIX> n=0
    /ch/09/automix/group  E32 F_XET enum=Xamxgrp
    /ch/09/automix/weight  F32 F_XET
```

### Xchannel10 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/10  <CHCO> n=0
/ch/10/config  <CHCO> n=0
    /ch/10/config/name  S32 F_XET
    /ch/10/config/icon  I32 F_XET
    /ch/10/config/color  E32 F_XET enum=Xcolors
    /ch/10/config/source  I32 F_XET
/ch/10/grp  <CHGRP> n=0
    /ch/10/grp/dca  P32 F_XET
    /ch/10/grp/mute  P32 F_XET
/ch/10/preamp  <CHPR> n=0
    /ch/10/preamp/trim  F32 F_XET
    /ch/10/preamp/invert  E32 F_XET enum=OffOn
    /ch/10/preamp/hpon  E32 F_XET enum=OffOn
    /ch/10/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/10/preamp/hpf  F32 F_XET
/ch/10/delay  <CHDE> n=0
    /ch/10/delay/on  E32 F_XET enum=OffOn
    /ch/10/delay/time  F32 F_XET
/ch/10/insert  <CHIN> n=0
    /ch/10/insert/on  E32 F_XET enum=OffOn
    /ch/10/insert/pos  E32 F_XET enum=Xdyppos
    /ch/10/insert/sel  E32 F_XET enum=Xisel
/ch/10/gate  <CHGA> n=0
    /ch/10/gate/on  E32 F_XET enum=OffOn
    /ch/10/gate/mode  E32 F_XET enum=Xgmode
    /ch/10/gate/thr  F32 F_XET
    /ch/10/gate/range  F32 F_XET
    /ch/10/gate/attack  F32 F_XET
    /ch/10/gate/hold  F32 F_XET
    /ch/10/gate/release  F32 F_XET
    /ch/10/gate/keysrc  I32 F_XET
/ch/10/gate/filter  <CHGF> n=0
    /ch/10/gate/filter/on  E32 F_XET enum=OffOn
    /ch/10/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/10/gate/filter/f  F32 F_XET
/ch/10/dyn  <CHDY> n=0
    /ch/10/dyn/on  E32 F_XET enum=OffOn
    /ch/10/dyn/mode  E32 F_XET enum=Xdymode
    /ch/10/dyn/det  E32 F_XET enum=Xdydet
    /ch/10/dyn/env  E32 F_XET enum=Xdyenv
    /ch/10/dyn/thr  F32 F_XET
    /ch/10/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/10/dyn/knee  F32 F_XET
    /ch/10/dyn/mgain  F32 F_XET
    /ch/10/dyn/attack  F32 F_XET
    /ch/10/dyn/hold  F32 F_XET
    /ch/10/dyn/release  F32 F_XET
    /ch/10/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/10/dyn/keysrc  I32 F_XET
    /ch/10/dyn/mix  F32 F_XET
    /ch/10/dyn/auto  E32 F_XET enum=OffOn
/ch/10/dyn/filter  <CHDF> n=0
    /ch/10/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/10/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/10/dyn/filter/f  F32 F_XET
/ch/10/eq  <OFFON> n=1
    /ch/10/eq/on  E32 F_XET enum=OffOn
/ch/10/eq/1  <CHEQ> n=0
    /ch/10/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/10/eq/1/f  F32 F_XET
    /ch/10/eq/1/g  F32 F_XET
    /ch/10/eq/1/q  F32 F_XET
/ch/10/eq/2  <CHEQ> n=0
    /ch/10/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/10/eq/2/f  F32 F_XET
    /ch/10/eq/2/g  F32 F_XET
    /ch/10/eq/2/q  F32 F_XET
/ch/10/eq/3  <CHEQ> n=0
    /ch/10/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/10/eq/3/f  F32 F_XET
    /ch/10/eq/3/g  F32 F_XET
    /ch/10/eq/3/q  F32 F_XET
/ch/10/eq/4  <CHEQ> n=0
    /ch/10/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/10/eq/4/f  F32 F_XET
    /ch/10/eq/4/g  F32 F_XET
    /ch/10/eq/4/q  F32 F_XET
/ch/10/mix  <CHMX> n=0
    /ch/10/mix/on  E32 F_XET enum=OffOn
    /ch/10/mix/fader  F32 F_XET
    /ch/10/mix/st  E32 F_XET enum=OffOn
    /ch/10/mix/pan  F32 F_XET
    /ch/10/mix/mono  E32 F_XET enum=OffOn
    /ch/10/mix/mlevel  F32 F_XET
/ch/10/mix/01  <CHMO> n=0
    /ch/10/mix/01/on  E32 F_XET enum=OffOn
    /ch/10/mix/01/level  F32 F_XET
    /ch/10/mix/01/pan  F32 F_XET
    /ch/10/mix/01/type  E32 F_XET enum=Xmtype
    /ch/10/mix/01/panFollow  E32 F_XET
/ch/10/mix/02  <CHME> n=0
    /ch/10/mix/02/on  E32 F_XET enum=OffOn
    /ch/10/mix/02/level  F32 F_XET
/ch/10/mix/03  <CHMO> n=0
    /ch/10/mix/03/on  E32 F_XET enum=OffOn
    /ch/10/mix/03/level  F32 F_XET
    /ch/10/mix/03/pan  F32 F_XET
    /ch/10/mix/03/type  E32 F_XET enum=Xmtype
    /ch/10/mix/03/panFollow  E32 F_XET
/ch/10/mix/04  <CHME> n=0
    /ch/10/mix/04/on  E32 F_XET enum=OffOn
    /ch/10/mix/04/level  F32 F_XET
/ch/10/mix/05  <CHMO> n=0
    /ch/10/mix/05/on  E32 F_XET enum=OffOn
    /ch/10/mix/05/level  F32 F_XET
    /ch/10/mix/05/pan  F32 F_XET
    /ch/10/mix/05/type  E32 F_XET enum=Xmtype
    /ch/10/mix/05/panFollow  E32 F_XET
/ch/10/mix/06  <CHME> n=0
    /ch/10/mix/06/on  E32 F_XET enum=OffOn
    /ch/10/mix/06/level  F32 F_XET
/ch/10/mix/07  <CHMO> n=0
    /ch/10/mix/07/on  E32 F_XET enum=OffOn
    /ch/10/mix/07/level  F32 F_XET
    /ch/10/mix/07/pan  F32 F_XET
    /ch/10/mix/07/type  E32 F_XET enum=Xmtype
    /ch/10/mix/07/panFollow  E32 F_XET
/ch/10/mix/08  <CHME> n=0
    /ch/10/mix/08/on  E32 F_XET enum=OffOn
    /ch/10/mix/08/level  F32 F_XET
/ch/10/mix/09  <CHMO> n=0
    /ch/10/mix/09/on  E32 F_XET enum=OffOn
    /ch/10/mix/09/level  F32 F_XET
    /ch/10/mix/09/pan  F32 F_XET
    /ch/10/mix/09/type  E32 F_XET enum=Xmtype
    /ch/10/mix/09/panFollow  E32 F_XET
/ch/10/mix/10  <CHME> n=0
    /ch/10/mix/10/on  E32 F_XET enum=OffOn
    /ch/10/mix/10/level  F32 F_XET
/ch/10/mix/11  <CHMO> n=0
    /ch/10/mix/11/on  E32 F_XET enum=OffOn
    /ch/10/mix/11/level  F32 F_XET
    /ch/10/mix/11/pan  F32 F_XET
    /ch/10/mix/11/type  E32 F_XET enum=Xmtype
    /ch/10/mix/11/panFollow  E32 F_XET
/ch/10/mix/12  <CHME> n=0
    /ch/10/mix/12/on  E32 F_XET enum=OffOn
    /ch/10/mix/12/level  F32 F_XET
/ch/10/mix/13  <CHMO> n=0
    /ch/10/mix/13/on  E32 F_XET enum=OffOn
    /ch/10/mix/13/level  F32 F_XET
    /ch/10/mix/13/pan  F32 F_XET
    /ch/10/mix/13/type  E32 F_XET enum=Xmtype
    /ch/10/mix/13/panFollow  E32 F_XET
/ch/10/mix/14  <CHME> n=0
    /ch/10/mix/14/on  E32 F_XET enum=OffOn
    /ch/10/mix/14/level  F32 F_XET
/ch/10/mix/15  <CHMO> n=0
    /ch/10/mix/15/on  E32 F_XET enum=OffOn
    /ch/10/mix/15/level  F32 F_XET
    /ch/10/mix/15/pan  F32 F_XET
    /ch/10/mix/15/type  E32 F_XET enum=Xmtype
    /ch/10/mix/15/panFollow  E32 F_XET
/ch/10/mix/16  <CHME> n=0
    /ch/10/mix/16/on  E32 F_XET enum=OffOn
    /ch/10/mix/16/level  F32 F_XET
/ch/10/automix  <CHAMIX> n=0
    /ch/10/automix/group  E32 F_XET enum=Xamxgrp
    /ch/10/automix/weight  F32 F_XET
```

### Xchannel11 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/11  <CHCO> n=0
/ch/11/config  <CHCO> n=0
    /ch/11/config/name  S32 F_XET
    /ch/11/config/icon  I32 F_XET
    /ch/11/config/color  E32 F_XET enum=Xcolors
    /ch/11/config/source  I32 F_XET
/ch/11/grp  <CHGRP> n=0
    /ch/11/grp/dca  P32 F_XET
    /ch/11/grp/mute  P32 F_XET
/ch/11/preamp  <CHPR> n=0
    /ch/11/preamp/trim  F32 F_XET
    /ch/11/preamp/invert  E32 F_XET enum=OffOn
    /ch/11/preamp/hpon  E32 F_XET enum=OffOn
    /ch/11/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/11/preamp/hpf  F32 F_XET
/ch/11/delay  <CHDE> n=0
    /ch/11/delay/on  E32 F_XET enum=OffOn
    /ch/11/delay/time  F32 F_XET
/ch/11/insert  <CHIN> n=0
    /ch/11/insert/on  E32 F_XET enum=OffOn
    /ch/11/insert/pos  E32 F_XET enum=Xdyppos
    /ch/11/insert/sel  E32 F_XET enum=Xisel
/ch/11/gate  <CHGA> n=0
    /ch/11/gate/on  E32 F_XET enum=OffOn
    /ch/11/gate/mode  E32 F_XET enum=Xgmode
    /ch/11/gate/thr  F32 F_XET
    /ch/11/gate/range  F32 F_XET
    /ch/11/gate/attack  F32 F_XET
    /ch/11/gate/hold  F32 F_XET
    /ch/11/gate/release  F32 F_XET
    /ch/11/gate/keysrc  I32 F_XET
/ch/11/gate/filter  <CHGF> n=0
    /ch/11/gate/filter/on  E32 F_XET enum=OffOn
    /ch/11/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/11/gate/filter/f  F32 F_XET
/ch/11/dyn  <CHDY> n=0
    /ch/11/dyn/on  E32 F_XET enum=OffOn
    /ch/11/dyn/mode  E32 F_XET enum=Xdymode
    /ch/11/dyn/det  E32 F_XET enum=Xdydet
    /ch/11/dyn/env  E32 F_XET enum=Xdyenv
    /ch/11/dyn/thr  F32 F_XET
    /ch/11/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/11/dyn/knee  F32 F_XET
    /ch/11/dyn/mgain  F32 F_XET
    /ch/11/dyn/attack  F32 F_XET
    /ch/11/dyn/hold  F32 F_XET
    /ch/11/dyn/release  F32 F_XET
    /ch/11/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/11/dyn/keysrc  I32 F_XET
    /ch/11/dyn/mix  F32 F_XET
    /ch/11/dyn/auto  E32 F_XET enum=OffOn
/ch/11/dyn/filter  <CHDF> n=0
    /ch/11/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/11/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/11/dyn/filter/f  F32 F_XET
/ch/11/eq  <OFFON> n=1
    /ch/11/eq/on  E32 F_XET enum=OffOn
/ch/11/eq/1  <CHEQ> n=0
    /ch/11/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/11/eq/1/f  F32 F_XET
    /ch/11/eq/1/g  F32 F_XET
    /ch/11/eq/1/q  F32 F_XET
/ch/11/eq/2  <CHEQ> n=0
    /ch/11/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/11/eq/2/f  F32 F_XET
    /ch/11/eq/2/g  F32 F_XET
    /ch/11/eq/2/q  F32 F_XET
/ch/11/eq/3  <CHEQ> n=0
    /ch/11/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/11/eq/3/f  F32 F_XET
    /ch/11/eq/3/g  F32 F_XET
    /ch/11/eq/3/q  F32 F_XET
/ch/11/eq/4  <CHEQ> n=0
    /ch/11/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/11/eq/4/f  F32 F_XET
    /ch/11/eq/4/g  F32 F_XET
    /ch/11/eq/4/q  F32 F_XET
/ch/11/mix  <CHMX> n=0
    /ch/11/mix/on  E32 F_XET enum=OffOn
    /ch/11/mix/fader  F32 F_XET
    /ch/11/mix/st  E32 F_XET enum=OffOn
    /ch/11/mix/pan  F32 F_XET
    /ch/11/mix/mono  E32 F_XET enum=OffOn
    /ch/11/mix/mlevel  F32 F_XET
/ch/11/mix/01  <CHMO> n=0
    /ch/11/mix/01/on  E32 F_XET enum=OffOn
    /ch/11/mix/01/level  F32 F_XET
    /ch/11/mix/01/pan  F32 F_XET
    /ch/11/mix/01/type  E32 F_XET enum=Xmtype
    /ch/11/mix/01/panFollow  E32 F_XET
/ch/11/mix/02  <CHME> n=0
    /ch/11/mix/02/on  E32 F_XET enum=OffOn
    /ch/11/mix/02/level  F32 F_XET
/ch/11/mix/03  <CHMO> n=0
    /ch/11/mix/03/on  E32 F_XET enum=OffOn
    /ch/11/mix/03/level  F32 F_XET
    /ch/11/mix/03/pan  F32 F_XET
    /ch/11/mix/03/type  E32 F_XET enum=Xmtype
    /ch/11/mix/03/panFollow  E32 F_XET
/ch/11/mix/04  <CHME> n=0
    /ch/11/mix/04/on  E32 F_XET enum=OffOn
    /ch/11/mix/04/level  F32 F_XET
/ch/11/mix/05  <CHMO> n=0
    /ch/11/mix/05/on  E32 F_XET enum=OffOn
    /ch/11/mix/05/level  F32 F_XET
    /ch/11/mix/05/pan  F32 F_XET
    /ch/11/mix/05/type  E32 F_XET enum=Xmtype
    /ch/11/mix/05/panFollow  E32 F_XET
/ch/11/mix/06  <CHME> n=0
    /ch/11/mix/06/on  E32 F_XET enum=OffOn
    /ch/11/mix/06/level  F32 F_XET
/ch/11/mix/07  <CHMO> n=0
    /ch/11/mix/07/on  E32 F_XET enum=OffOn
    /ch/11/mix/07/level  F32 F_XET
    /ch/11/mix/07/pan  F32 F_XET
    /ch/11/mix/07/type  E32 F_XET enum=Xmtype
    /ch/11/mix/07/panFollow  E32 F_XET
/ch/11/mix/08  <CHME> n=0
    /ch/11/mix/08/on  E32 F_XET enum=OffOn
    /ch/11/mix/08/level  F32 F_XET
/ch/11/mix/09  <CHMO> n=0
    /ch/11/mix/09/on  E32 F_XET enum=OffOn
    /ch/11/mix/09/level  F32 F_XET
    /ch/11/mix/09/pan  F32 F_XET
    /ch/11/mix/09/type  E32 F_XET enum=Xmtype
    /ch/11/mix/09/panFollow  E32 F_XET
/ch/11/mix/10  <CHME> n=0
    /ch/11/mix/10/on  E32 F_XET enum=OffOn
    /ch/11/mix/10/level  F32 F_XET
/ch/11/mix/11  <CHMO> n=0
    /ch/11/mix/11/on  E32 F_XET enum=OffOn
    /ch/11/mix/11/level  F32 F_XET
    /ch/11/mix/11/pan  F32 F_XET
    /ch/11/mix/11/type  E32 F_XET enum=Xmtype
    /ch/11/mix/11/panFollow  E32 F_XET
/ch/11/mix/12  <CHME> n=0
    /ch/11/mix/12/on  E32 F_XET enum=OffOn
    /ch/11/mix/12/level  F32 F_XET
/ch/11/mix/13  <CHMO> n=0
    /ch/11/mix/13/on  E32 F_XET enum=OffOn
    /ch/11/mix/13/level  F32 F_XET
    /ch/11/mix/13/pan  F32 F_XET
    /ch/11/mix/13/type  E32 F_XET enum=Xmtype
    /ch/11/mix/13/panFollow  E32 F_XET
/ch/11/mix/14  <CHME> n=0
    /ch/11/mix/14/on  E32 F_XET enum=OffOn
    /ch/11/mix/14/level  F32 F_XET
/ch/11/mix/15  <CHMO> n=0
    /ch/11/mix/15/on  E32 F_XET enum=OffOn
    /ch/11/mix/15/level  F32 F_XET
    /ch/11/mix/15/pan  F32 F_XET
    /ch/11/mix/15/type  E32 F_XET enum=Xmtype
    /ch/11/mix/15/panFollow  E32 F_XET
/ch/11/mix/16  <CHME> n=0
    /ch/11/mix/16/on  E32 F_XET enum=OffOn
    /ch/11/mix/16/level  F32 F_XET
/ch/11/automix  <CHAMIX> n=0
    /ch/11/automix/group  E32 F_XET enum=Xamxgrp
    /ch/11/automix/weight  F32 F_XET
```

### Xchannel12 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/12  <CHCO> n=0
/ch/12/config  <CHCO> n=0
    /ch/12/config/name  S32 F_XET
    /ch/12/config/icon  I32 F_XET
    /ch/12/config/color  E32 F_XET enum=Xcolors
    /ch/12/config/source  I32 F_XET
/ch/12/grp  <CHGRP> n=0
    /ch/12/grp/dca  P32 F_XET
    /ch/12/grp/mute  P32 F_XET
/ch/12/preamp  <CHPR> n=0
    /ch/12/preamp/trim  F32 F_XET
    /ch/12/preamp/invert  E32 F_XET enum=OffOn
    /ch/12/preamp/hpon  E32 F_XET enum=OffOn
    /ch/12/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/12/preamp/hpf  F32 F_XET
/ch/12/delay  <CHDE> n=0
    /ch/12/delay/on  E32 F_XET enum=OffOn
    /ch/12/delay/time  F32 F_XET
/ch/12/insert  <CHIN> n=0
    /ch/12/insert/on  E32 F_XET enum=OffOn
    /ch/12/insert/pos  E32 F_XET enum=Xdyppos
    /ch/12/insert/sel  E32 F_XET enum=Xisel
/ch/12/gate  <CHGA> n=0
    /ch/12/gate/on  E32 F_XET enum=OffOn
    /ch/12/gate/mode  E32 F_XET enum=Xgmode
    /ch/12/gate/thr  F32 F_XET
    /ch/12/gate/range  F32 F_XET
    /ch/12/gate/attack  F32 F_XET
    /ch/12/gate/hold  F32 F_XET
    /ch/12/gate/release  F32 F_XET
    /ch/12/gate/keysrc  I32 F_XET
/ch/12/gate/filter  <CHGF> n=0
    /ch/12/gate/filter/on  E32 F_XET enum=OffOn
    /ch/12/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/12/gate/filter/f  F32 F_XET
/ch/12/dyn  <CHDY> n=0
    /ch/12/dyn/on  E32 F_XET enum=OffOn
    /ch/12/dyn/mode  E32 F_XET enum=Xdymode
    /ch/12/dyn/det  E32 F_XET enum=Xdydet
    /ch/12/dyn/env  E32 F_XET enum=Xdyenv
    /ch/12/dyn/thr  F32 F_XET
    /ch/12/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/12/dyn/knee  F32 F_XET
    /ch/12/dyn/mgain  F32 F_XET
    /ch/12/dyn/attack  F32 F_XET
    /ch/12/dyn/hold  F32 F_XET
    /ch/12/dyn/release  F32 F_XET
    /ch/12/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/12/dyn/keysrc  I32 F_XET
    /ch/12/dyn/mix  F32 F_XET
    /ch/12/dyn/auto  E32 F_XET enum=OffOn
/ch/12/dyn/filter  <CHDF> n=0
    /ch/12/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/12/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/12/dyn/filter/f  F32 F_XET
/ch/12/eq  <OFFON> n=1
    /ch/12/eq/on  E32 F_XET enum=OffOn
/ch/12/eq/1  <CHEQ> n=0
    /ch/12/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/12/eq/1/f  F32 F_XET
    /ch/12/eq/1/g  F32 F_XET
    /ch/12/eq/1/q  F32 F_XET
/ch/12/eq/2  <CHEQ> n=0
    /ch/12/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/12/eq/2/f  F32 F_XET
    /ch/12/eq/2/g  F32 F_XET
    /ch/12/eq/2/q  F32 F_XET
/ch/12/eq/3  <CHEQ> n=0
    /ch/12/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/12/eq/3/f  F32 F_XET
    /ch/12/eq/3/g  F32 F_XET
    /ch/12/eq/3/q  F32 F_XET
/ch/12/eq/4  <CHEQ> n=0
    /ch/12/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/12/eq/4/f  F32 F_XET
    /ch/12/eq/4/g  F32 F_XET
    /ch/12/eq/4/q  F32 F_XET
/ch/12/mix  <CHMX> n=0
    /ch/12/mix/on  E32 F_XET enum=OffOn
    /ch/12/mix/fader  F32 F_XET
    /ch/12/mix/st  E32 F_XET enum=OffOn
    /ch/12/mix/pan  F32 F_XET
    /ch/12/mix/mono  E32 F_XET enum=OffOn
    /ch/12/mix/mlevel  F32 F_XET
/ch/12/mix/01  <CHMO> n=0
    /ch/12/mix/01/on  E32 F_XET enum=OffOn
    /ch/12/mix/01/level  F32 F_XET
    /ch/12/mix/01/pan  F32 F_XET
    /ch/12/mix/01/type  E32 F_XET enum=Xmtype
    /ch/12/mix/01/panFollow  E32 F_XET
/ch/12/mix/02  <CHME> n=0
    /ch/12/mix/02/on  E32 F_XET enum=OffOn
    /ch/12/mix/02/level  F32 F_XET
/ch/12/mix/03  <CHMO> n=0
    /ch/12/mix/03/on  E32 F_XET enum=OffOn
    /ch/12/mix/03/level  F32 F_XET
    /ch/12/mix/03/pan  F32 F_XET
    /ch/12/mix/03/type  E32 F_XET enum=Xmtype
    /ch/12/mix/03/panFollow  E32 F_XET
/ch/12/mix/04  <CHME> n=0
    /ch/12/mix/04/on  E32 F_XET enum=OffOn
    /ch/12/mix/04/level  F32 F_XET
/ch/12/mix/05  <CHMO> n=0
    /ch/12/mix/05/on  E32 F_XET enum=OffOn
    /ch/12/mix/05/level  F32 F_XET
    /ch/12/mix/05/pan  F32 F_XET
    /ch/12/mix/05/type  E32 F_XET enum=Xmtype
    /ch/12/mix/05/panFollow  E32 F_XET
/ch/12/mix/06  <CHME> n=0
    /ch/12/mix/06/on  E32 F_XET enum=OffOn
    /ch/12/mix/06/level  F32 F_XET
/ch/12/mix/07  <CHMO> n=0
    /ch/12/mix/07/on  E32 F_XET enum=OffOn
    /ch/12/mix/07/level  F32 F_XET
    /ch/12/mix/07/pan  F32 F_XET
    /ch/12/mix/07/type  E32 F_XET enum=Xmtype
    /ch/12/mix/07/panFollow  E32 F_XET
/ch/12/mix/08  <CHME> n=0
    /ch/12/mix/08/on  E32 F_XET enum=OffOn
    /ch/12/mix/08/level  F32 F_XET
/ch/12/mix/09  <CHMO> n=0
    /ch/12/mix/09/on  E32 F_XET enum=OffOn
    /ch/12/mix/09/level  F32 F_XET
    /ch/12/mix/09/pan  F32 F_XET
    /ch/12/mix/09/type  E32 F_XET enum=Xmtype
    /ch/12/mix/09/panFollow  E32 F_XET
/ch/12/mix/10  <CHME> n=0
    /ch/12/mix/10/on  E32 F_XET enum=OffOn
    /ch/12/mix/10/level  F32 F_XET
/ch/12/mix/11  <CHMO> n=0
    /ch/12/mix/11/on  E32 F_XET enum=OffOn
    /ch/12/mix/11/level  F32 F_XET
    /ch/12/mix/11/pan  F32 F_XET
    /ch/12/mix/11/type  E32 F_XET enum=Xmtype
    /ch/12/mix/11/panFollow  E32 F_XET
/ch/12/mix/12  <CHME> n=0
    /ch/12/mix/12/on  E32 F_XET enum=OffOn
    /ch/12/mix/12/level  F32 F_XET
/ch/12/mix/13  <CHMO> n=0
    /ch/12/mix/13/on  E32 F_XET enum=OffOn
    /ch/12/mix/13/level  F32 F_XET
    /ch/12/mix/13/pan  F32 F_XET
    /ch/12/mix/13/type  E32 F_XET enum=Xmtype
    /ch/12/mix/13/panFollow  E32 F_XET
/ch/12/mix/14  <CHME> n=0
    /ch/12/mix/14/on  E32 F_XET enum=OffOn
    /ch/12/mix/14/level  F32 F_XET
/ch/12/mix/15  <CHMO> n=0
    /ch/12/mix/15/on  E32 F_XET enum=OffOn
    /ch/12/mix/15/level  F32 F_XET
    /ch/12/mix/15/pan  F32 F_XET
    /ch/12/mix/15/type  E32 F_XET enum=Xmtype
    /ch/12/mix/15/panFollow  E32 F_XET
/ch/12/mix/16  <CHME> n=0
    /ch/12/mix/16/on  E32 F_XET enum=OffOn
    /ch/12/mix/16/level  F32 F_XET
/ch/12/automix  <CHAMIX> n=0
    /ch/12/automix/group  E32 F_XET enum=Xamxgrp
    /ch/12/automix/weight  F32 F_XET
```

### Xchannel13 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/13  <CHCO> n=0
/ch/13/config  <CHCO> n=0
    /ch/13/config/name  S32 F_XET
    /ch/13/config/icon  I32 F_XET
    /ch/13/config/color  E32 F_XET enum=Xcolors
    /ch/13/config/source  I32 F_XET
/ch/13/grp  <CHGRP> n=0
    /ch/13/grp/dca  P32 F_XET
    /ch/13/grp/mute  P32 F_XET
/ch/13/preamp  <CHPR> n=0
    /ch/13/preamp/trim  F32 F_XET
    /ch/13/preamp/invert  E32 F_XET enum=OffOn
    /ch/13/preamp/hpon  E32 F_XET enum=OffOn
    /ch/13/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/13/preamp/hpf  F32 F_XET
/ch/13/delay  <CHDE> n=0
    /ch/13/delay/on  E32 F_XET enum=OffOn
    /ch/13/delay/time  F32 F_XET
/ch/13/insert  <CHIN> n=0
    /ch/13/insert/on  E32 F_XET enum=OffOn
    /ch/13/insert/pos  E32 F_XET enum=Xdyppos
    /ch/13/insert/sel  E32 F_XET enum=Xisel
/ch/13/gate  <CHGA> n=0
    /ch/13/gate/on  E32 F_XET enum=OffOn
    /ch/13/gate/mode  E32 F_XET enum=Xgmode
    /ch/13/gate/thr  F32 F_XET
    /ch/13/gate/range  F32 F_XET
    /ch/13/gate/attack  F32 F_XET
    /ch/13/gate/hold  F32 F_XET
    /ch/13/gate/release  F32 F_XET
    /ch/13/gate/keysrc  I32 F_XET
/ch/13/gate/filter  <CHGF> n=0
    /ch/13/gate/filter/on  E32 F_XET enum=OffOn
    /ch/13/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/13/gate/filter/f  F32 F_XET
/ch/13/dyn  <CHDY> n=0
    /ch/13/dyn/on  E32 F_XET enum=OffOn
    /ch/13/dyn/mode  E32 F_XET enum=Xdymode
    /ch/13/dyn/det  E32 F_XET enum=Xdydet
    /ch/13/dyn/env  E32 F_XET enum=Xdyenv
    /ch/13/dyn/thr  F32 F_XET
    /ch/13/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/13/dyn/knee  F32 F_XET
    /ch/13/dyn/mgain  F32 F_XET
    /ch/13/dyn/attack  F32 F_XET
    /ch/13/dyn/hold  F32 F_XET
    /ch/13/dyn/release  F32 F_XET
    /ch/13/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/13/dyn/keysrc  I32 F_XET
    /ch/13/dyn/mix  F32 F_XET
    /ch/13/dyn/auto  E32 F_XET enum=OffOn
/ch/13/dyn/filter  <CHDF> n=0
    /ch/13/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/13/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/13/dyn/filter/f  F32 F_XET
/ch/13/eq  <OFFON> n=1
    /ch/13/eq/on  E32 F_XET enum=OffOn
/ch/13/eq/1  <CHEQ> n=0
    /ch/13/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/13/eq/1/f  F32 F_XET
    /ch/13/eq/1/g  F32 F_XET
    /ch/13/eq/1/q  F32 F_XET
/ch/13/eq/2  <CHEQ> n=0
    /ch/13/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/13/eq/2/f  F32 F_XET
    /ch/13/eq/2/g  F32 F_XET
    /ch/13/eq/2/q  F32 F_XET
/ch/13/eq/3  <CHEQ> n=0
    /ch/13/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/13/eq/3/f  F32 F_XET
    /ch/13/eq/3/g  F32 F_XET
    /ch/13/eq/3/q  F32 F_XET
/ch/13/eq/4  <CHEQ> n=0
    /ch/13/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/13/eq/4/f  F32 F_XET
    /ch/13/eq/4/g  F32 F_XET
    /ch/13/eq/4/q  F32 F_XET
/ch/13/mix  <CHMX> n=0
    /ch/13/mix/on  E32 F_XET enum=OffOn
    /ch/13/mix/fader  F32 F_XET
    /ch/13/mix/st  E32 F_XET enum=OffOn
    /ch/13/mix/pan  F32 F_XET
    /ch/13/mix/mono  E32 F_XET enum=OffOn
    /ch/13/mix/mlevel  F32 F_XET
/ch/13/mix/01  <CHMO> n=0
    /ch/13/mix/01/on  E32 F_XET enum=OffOn
    /ch/13/mix/01/level  F32 F_XET
    /ch/13/mix/01/pan  F32 F_XET
    /ch/13/mix/01/type  E32 F_XET enum=Xmtype
    /ch/13/mix/01/panFollow  E32 F_XET
/ch/13/mix/02  <CHME> n=0
    /ch/13/mix/02/on  E32 F_XET enum=OffOn
    /ch/13/mix/02/level  F32 F_XET
/ch/13/mix/03  <CHMO> n=0
    /ch/13/mix/03/on  E32 F_XET enum=OffOn
    /ch/13/mix/03/level  F32 F_XET
    /ch/13/mix/03/pan  F32 F_XET
    /ch/13/mix/03/type  E32 F_XET enum=Xmtype
    /ch/13/mix/03/panFollow  E32 F_XET
/ch/13/mix/04  <CHME> n=0
    /ch/13/mix/04/on  E32 F_XET enum=OffOn
    /ch/13/mix/04/level  F32 F_XET
/ch/13/mix/05  <CHMO> n=0
    /ch/13/mix/05/on  E32 F_XET enum=OffOn
    /ch/13/mix/05/level  F32 F_XET
    /ch/13/mix/05/pan  F32 F_XET
    /ch/13/mix/05/type  E32 F_XET enum=Xmtype
    /ch/13/mix/05/panFollow  E32 F_XET
/ch/13/mix/06  <CHME> n=0
    /ch/13/mix/06/on  E32 F_XET enum=OffOn
    /ch/13/mix/06/level  F32 F_XET
/ch/13/mix/07  <CHMO> n=0
    /ch/13/mix/07/on  E32 F_XET enum=OffOn
    /ch/13/mix/07/level  F32 F_XET
    /ch/13/mix/07/pan  F32 F_XET
    /ch/13/mix/07/type  E32 F_XET enum=Xmtype
    /ch/13/mix/07/panFollow  E32 F_XET
/ch/13/mix/08  <CHME> n=0
    /ch/13/mix/08/on  E32 F_XET enum=OffOn
    /ch/13/mix/08/level  F32 F_XET
/ch/13/mix/09  <CHMO> n=0
    /ch/13/mix/09/on  E32 F_XET enum=OffOn
    /ch/13/mix/09/level  F32 F_XET
    /ch/13/mix/09/pan  F32 F_XET
    /ch/13/mix/09/type  E32 F_XET enum=Xmtype
    /ch/13/mix/09/panFollow  E32 F_XET
/ch/13/mix/10  <CHME> n=0
    /ch/13/mix/10/on  E32 F_XET enum=OffOn
    /ch/13/mix/10/level  F32 F_XET
/ch/13/mix/11  <CHMO> n=0
    /ch/13/mix/11/on  E32 F_XET enum=OffOn
    /ch/13/mix/11/level  F32 F_XET
    /ch/13/mix/11/pan  F32 F_XET
    /ch/13/mix/11/type  E32 F_XET enum=Xmtype
    /ch/13/mix/11/panFollow  E32 F_XET
/ch/13/mix/12  <CHME> n=0
    /ch/13/mix/12/on  E32 F_XET enum=OffOn
    /ch/13/mix/12/level  F32 F_XET
/ch/13/mix/13  <CHMO> n=0
    /ch/13/mix/13/on  E32 F_XET enum=OffOn
    /ch/13/mix/13/level  F32 F_XET
    /ch/13/mix/13/pan  F32 F_XET
    /ch/13/mix/13/type  E32 F_XET enum=Xmtype
    /ch/13/mix/13/panFollow  E32 F_XET
/ch/13/mix/14  <CHME> n=0
    /ch/13/mix/14/on  E32 F_XET enum=OffOn
    /ch/13/mix/14/level  F32 F_XET
/ch/13/mix/15  <CHMO> n=0
    /ch/13/mix/15/on  E32 F_XET enum=OffOn
    /ch/13/mix/15/level  F32 F_XET
    /ch/13/mix/15/pan  F32 F_XET
    /ch/13/mix/15/type  E32 F_XET enum=Xmtype
    /ch/13/mix/15/panFollow  E32 F_XET
/ch/13/mix/16  <CHME> n=0
    /ch/13/mix/16/on  E32 F_XET enum=OffOn
    /ch/13/mix/16/level  F32 F_XET
/ch/13/automix  <CHAMIX> n=0
    /ch/13/automix/group  E32 F_XET enum=Xamxgrp
    /ch/13/automix/weight  F32 F_XET
```

### Xchannel14 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/14  <CHCO> n=0
/ch/14/config  <CHCO> n=0
    /ch/14/config/name  S32 F_XET
    /ch/14/config/icon  I32 F_XET
    /ch/14/config/color  E32 F_XET enum=Xcolors
    /ch/14/config/source  I32 F_XET
/ch/14/grp  <CHGRP> n=0
    /ch/14/grp/dca  P32 F_XET
    /ch/14/grp/mute  P32 F_XET
/ch/14/preamp  <CHPR> n=0
    /ch/14/preamp/trim  F32 F_XET
    /ch/14/preamp/invert  E32 F_XET enum=OffOn
    /ch/14/preamp/hpon  E32 F_XET enum=OffOn
    /ch/14/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/14/preamp/hpf  F32 F_XET
/ch/14/delay  <CHDE> n=0
    /ch/14/delay/on  E32 F_XET enum=OffOn
    /ch/14/delay/time  F32 F_XET
/ch/14/insert  <CHIN> n=0
    /ch/14/insert/on  E32 F_XET enum=OffOn
    /ch/14/insert/pos  E32 F_XET enum=Xdyppos
    /ch/14/insert/sel  E32 F_XET enum=Xisel
/ch/14/gate  <CHGA> n=0
    /ch/14/gate/on  E32 F_XET enum=OffOn
    /ch/14/gate/mode  E32 F_XET enum=Xgmode
    /ch/14/gate/thr  F32 F_XET
    /ch/14/gate/range  F32 F_XET
    /ch/14/gate/attack  F32 F_XET
    /ch/14/gate/hold  F32 F_XET
    /ch/14/gate/release  F32 F_XET
    /ch/14/gate/keysrc  I32 F_XET
/ch/14/gate/filter  <CHGF> n=0
    /ch/14/gate/filter/on  E32 F_XET enum=OffOn
    /ch/14/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/14/gate/filter/f  F32 F_XET
/ch/14/dyn  <CHDY> n=0
    /ch/14/dyn/on  E32 F_XET enum=OffOn
    /ch/14/dyn/mode  E32 F_XET enum=Xdymode
    /ch/14/dyn/det  E32 F_XET enum=Xdydet
    /ch/14/dyn/env  E32 F_XET enum=Xdyenv
    /ch/14/dyn/thr  F32 F_XET
    /ch/14/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/14/dyn/knee  F32 F_XET
    /ch/14/dyn/mgain  F32 F_XET
    /ch/14/dyn/attack  F32 F_XET
    /ch/14/dyn/hold  F32 F_XET
    /ch/14/dyn/release  F32 F_XET
    /ch/14/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/14/dyn/keysrc  I32 F_XET
    /ch/14/dyn/mix  F32 F_XET
    /ch/14/dyn/auto  E32 F_XET enum=OffOn
/ch/14/dyn/filter  <CHDF> n=0
    /ch/14/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/14/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/14/dyn/filter/f  F32 F_XET
/ch/14/eq  <OFFON> n=1
    /ch/14/eq/on  E32 F_XET enum=OffOn
/ch/14/eq/1  <CHEQ> n=0
    /ch/14/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/14/eq/1/f  F32 F_XET
    /ch/14/eq/1/g  F32 F_XET
    /ch/14/eq/1/q  F32 F_XET
/ch/14/eq/2  <CHEQ> n=0
    /ch/14/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/14/eq/2/f  F32 F_XET
    /ch/14/eq/2/g  F32 F_XET
    /ch/14/eq/2/q  F32 F_XET
/ch/14/eq/3  <CHEQ> n=0
    /ch/14/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/14/eq/3/f  F32 F_XET
    /ch/14/eq/3/g  F32 F_XET
    /ch/14/eq/3/q  F32 F_XET
/ch/14/eq/4  <CHEQ> n=0
    /ch/14/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/14/eq/4/f  F32 F_XET
    /ch/14/eq/4/g  F32 F_XET
    /ch/14/eq/4/q  F32 F_XET
/ch/14/mix  <CHMX> n=0
    /ch/14/mix/on  E32 F_XET enum=OffOn
    /ch/14/mix/fader  F32 F_XET
    /ch/14/mix/st  E32 F_XET enum=OffOn
    /ch/14/mix/pan  F32 F_XET
    /ch/14/mix/mono  E32 F_XET enum=OffOn
    /ch/14/mix/mlevel  F32 F_XET
/ch/14/mix/01  <CHMO> n=0
    /ch/14/mix/01/on  E32 F_XET enum=OffOn
    /ch/14/mix/01/level  F32 F_XET
    /ch/14/mix/01/pan  F32 F_XET
    /ch/14/mix/01/type  E32 F_XET enum=Xmtype
    /ch/14/mix/01/panFollow  E32 F_XET
/ch/14/mix/02  <CHME> n=0
    /ch/14/mix/02/on  E32 F_XET enum=OffOn
    /ch/14/mix/02/level  F32 F_XET
/ch/14/mix/03  <CHMO> n=0
    /ch/14/mix/03/on  E32 F_XET enum=OffOn
    /ch/14/mix/03/level  F32 F_XET
    /ch/14/mix/03/pan  F32 F_XET
    /ch/14/mix/03/type  E32 F_XET enum=Xmtype
    /ch/14/mix/03/panFollow  E32 F_XET
/ch/14/mix/04  <CHME> n=0
    /ch/14/mix/04/on  E32 F_XET enum=OffOn
    /ch/14/mix/04/level  F32 F_XET
/ch/14/mix/05  <CHMO> n=0
    /ch/14/mix/05/on  E32 F_XET enum=OffOn
    /ch/14/mix/05/level  F32 F_XET
    /ch/14/mix/05/pan  F32 F_XET
    /ch/14/mix/05/type  E32 F_XET enum=Xmtype
    /ch/14/mix/05/panFollow  E32 F_XET
/ch/14/mix/06  <CHME> n=0
    /ch/14/mix/06/on  E32 F_XET enum=OffOn
    /ch/14/mix/06/level  F32 F_XET
/ch/14/mix/07  <CHMO> n=0
    /ch/14/mix/07/on  E32 F_XET enum=OffOn
    /ch/14/mix/07/level  F32 F_XET
    /ch/14/mix/07/pan  F32 F_XET
    /ch/14/mix/07/type  E32 F_XET enum=Xmtype
    /ch/14/mix/07/panFollow  E32 F_XET
/ch/14/mix/08  <CHME> n=0
    /ch/14/mix/08/on  E32 F_XET enum=OffOn
    /ch/14/mix/08/level  F32 F_XET
/ch/14/mix/09  <CHMO> n=0
    /ch/14/mix/09/on  E32 F_XET enum=OffOn
    /ch/14/mix/09/level  F32 F_XET
    /ch/14/mix/09/pan  F32 F_XET
    /ch/14/mix/09/type  E32 F_XET enum=Xmtype
    /ch/14/mix/09/panFollow  E32 F_XET
/ch/14/mix/10  <CHME> n=0
    /ch/14/mix/10/on  E32 F_XET enum=OffOn
    /ch/14/mix/10/level  F32 F_XET
/ch/14/mix/11  <CHMO> n=0
    /ch/14/mix/11/on  E32 F_XET enum=OffOn
    /ch/14/mix/11/level  F32 F_XET
    /ch/14/mix/11/pan  F32 F_XET
    /ch/14/mix/11/type  E32 F_XET enum=Xmtype
    /ch/14/mix/11/panFollow  E32 F_XET
/ch/14/mix/12  <CHME> n=0
    /ch/14/mix/12/on  E32 F_XET enum=OffOn
    /ch/14/mix/12/level  F32 F_XET
/ch/14/mix/13  <CHMO> n=0
    /ch/14/mix/13/on  E32 F_XET enum=OffOn
    /ch/14/mix/13/level  F32 F_XET
    /ch/14/mix/13/pan  F32 F_XET
    /ch/14/mix/13/type  E32 F_XET enum=Xmtype
    /ch/14/mix/13/panFollow  E32 F_XET
/ch/14/mix/14  <CHME> n=0
    /ch/14/mix/14/on  E32 F_XET enum=OffOn
    /ch/14/mix/14/level  F32 F_XET
/ch/14/mix/15  <CHMO> n=0
    /ch/14/mix/15/on  E32 F_XET enum=OffOn
    /ch/14/mix/15/level  F32 F_XET
    /ch/14/mix/15/pan  F32 F_XET
    /ch/14/mix/15/type  E32 F_XET enum=Xmtype
    /ch/14/mix/15/panFollow  E32 F_XET
/ch/14/mix/16  <CHME> n=0
    /ch/14/mix/16/on  E32 F_XET enum=OffOn
    /ch/14/mix/16/level  F32 F_XET
/ch/14/automix  <CHAMIX> n=0
    /ch/14/automix/group  E32 F_XET enum=Xamxgrp
    /ch/14/automix/weight  F32 F_XET
```

### Xchannel15 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/15  <CHCO> n=0
/ch/15/config  <CHCO> n=0
    /ch/15/config/name  S32 F_XET
    /ch/15/config/icon  I32 F_XET
    /ch/15/config/color  E32 F_XET enum=Xcolors
    /ch/15/config/source  I32 F_XET
/ch/15/grp  <CHGRP> n=0
    /ch/15/grp/dca  P32 F_XET
    /ch/15/grp/mute  P32 F_XET
/ch/15/preamp  <CHPR> n=0
    /ch/15/preamp/trim  F32 F_XET
    /ch/15/preamp/invert  E32 F_XET enum=OffOn
    /ch/15/preamp/hpon  E32 F_XET enum=OffOn
    /ch/15/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/15/preamp/hpf  F32 F_XET
/ch/15/delay  <CHDE> n=0
    /ch/15/delay/on  E32 F_XET enum=OffOn
    /ch/15/delay/time  F32 F_XET
/ch/15/insert  <CHIN> n=0
    /ch/15/insert/on  E32 F_XET enum=OffOn
    /ch/15/insert/pos  E32 F_XET enum=Xdyppos
    /ch/15/insert/sel  E32 F_XET enum=Xisel
/ch/15/gate  <CHGA> n=0
    /ch/15/gate/on  E32 F_XET enum=OffOn
    /ch/15/gate/mode  E32 F_XET enum=Xgmode
    /ch/15/gate/thr  F32 F_XET
    /ch/15/gate/range  F32 F_XET
    /ch/15/gate/attack  F32 F_XET
    /ch/15/gate/hold  F32 F_XET
    /ch/15/gate/release  F32 F_XET
    /ch/15/gate/keysrc  I32 F_XET
/ch/15/gate/filter  <CHGF> n=0
    /ch/15/gate/filter/on  E32 F_XET enum=OffOn
    /ch/15/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/15/gate/filter/f  F32 F_XET
/ch/15/dyn  <CHDY> n=0
    /ch/15/dyn/on  E32 F_XET enum=OffOn
    /ch/15/dyn/mode  E32 F_XET enum=Xdymode
    /ch/15/dyn/det  E32 F_XET enum=Xdydet
    /ch/15/dyn/env  E32 F_XET enum=Xdyenv
    /ch/15/dyn/thr  F32 F_XET
    /ch/15/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/15/dyn/knee  F32 F_XET
    /ch/15/dyn/mgain  F32 F_XET
    /ch/15/dyn/attack  F32 F_XET
    /ch/15/dyn/hold  F32 F_XET
    /ch/15/dyn/release  F32 F_XET
    /ch/15/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/15/dyn/keysrc  I32 F_XET
    /ch/15/dyn/mix  F32 F_XET
    /ch/15/dyn/auto  E32 F_XET enum=OffOn
/ch/15/dyn/filter  <CHDF> n=0
    /ch/15/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/15/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/15/dyn/filter/f  F32 F_XET
/ch/15/eq  <OFFON> n=1
    /ch/15/eq/on  E32 F_XET enum=OffOn
/ch/15/eq/1  <CHEQ> n=0
    /ch/15/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/15/eq/1/f  F32 F_XET
    /ch/15/eq/1/g  F32 F_XET
    /ch/15/eq/1/q  F32 F_XET
/ch/15/eq/2  <CHEQ> n=0
    /ch/15/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/15/eq/2/f  F32 F_XET
    /ch/15/eq/2/g  F32 F_XET
    /ch/15/eq/2/q  F32 F_XET
/ch/15/eq/3  <CHEQ> n=0
    /ch/15/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/15/eq/3/f  F32 F_XET
    /ch/15/eq/3/g  F32 F_XET
    /ch/15/eq/3/q  F32 F_XET
/ch/15/eq/4  <CHEQ> n=0
    /ch/15/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/15/eq/4/f  F32 F_XET
    /ch/15/eq/4/g  F32 F_XET
    /ch/15/eq/4/q  F32 F_XET
/ch/15/mix  <CHMX> n=0
    /ch/15/mix/on  E32 F_XET enum=OffOn
    /ch/15/mix/fader  F32 F_XET
    /ch/15/mix/st  E32 F_XET enum=OffOn
    /ch/15/mix/pan  F32 F_XET
    /ch/15/mix/mono  E32 F_XET enum=OffOn
    /ch/15/mix/mlevel  F32 F_XET
/ch/15/mix/01  <CHMO> n=0
    /ch/15/mix/01/on  E32 F_XET enum=OffOn
    /ch/15/mix/01/level  F32 F_XET
    /ch/15/mix/01/pan  F32 F_XET
    /ch/15/mix/01/type  E32 F_XET enum=Xmtype
    /ch/15/mix/01/panFollow  E32 F_XET
/ch/15/mix/02  <CHME> n=0
    /ch/15/mix/02/on  E32 F_XET enum=OffOn
    /ch/15/mix/02/level  F32 F_XET
/ch/15/mix/03  <CHMO> n=0
    /ch/15/mix/03/on  E32 F_XET enum=OffOn
    /ch/15/mix/03/level  F32 F_XET
    /ch/15/mix/03/pan  F32 F_XET
    /ch/15/mix/03/type  E32 F_XET enum=Xmtype
    /ch/15/mix/03/panFollow  E32 F_XET
/ch/15/mix/04  <CHME> n=0
    /ch/15/mix/04/on  E32 F_XET enum=OffOn
    /ch/15/mix/04/level  F32 F_XET
/ch/15/mix/05  <CHMO> n=0
    /ch/15/mix/05/on  E32 F_XET enum=OffOn
    /ch/15/mix/05/level  F32 F_XET
    /ch/15/mix/05/pan  F32 F_XET
    /ch/15/mix/05/type  E32 F_XET enum=Xmtype
    /ch/15/mix/05/panFollow  E32 F_XET
/ch/15/mix/06  <CHME> n=0
    /ch/15/mix/06/on  E32 F_XET enum=OffOn
    /ch/15/mix/06/level  F32 F_XET
/ch/15/mix/07  <CHMO> n=0
    /ch/15/mix/07/on  E32 F_XET enum=OffOn
    /ch/15/mix/07/level  F32 F_XET
    /ch/15/mix/07/pan  F32 F_XET
    /ch/15/mix/07/type  E32 F_XET enum=Xmtype
    /ch/15/mix/07/panFollow  E32 F_XET
/ch/15/mix/08  <CHME> n=0
    /ch/15/mix/08/on  E32 F_XET enum=OffOn
    /ch/15/mix/08/level  F32 F_XET
/ch/15/mix/09  <CHMO> n=0
    /ch/15/mix/09/on  E32 F_XET enum=OffOn
    /ch/15/mix/09/level  F32 F_XET
    /ch/15/mix/09/pan  F32 F_XET
    /ch/15/mix/09/type  E32 F_XET enum=Xmtype
    /ch/15/mix/09/panFollow  E32 F_XET
/ch/15/mix/10  <CHME> n=0
    /ch/15/mix/10/on  E32 F_XET enum=OffOn
    /ch/15/mix/10/level  F32 F_XET
/ch/15/mix/11  <CHMO> n=0
    /ch/15/mix/11/on  E32 F_XET enum=OffOn
    /ch/15/mix/11/level  F32 F_XET
    /ch/15/mix/11/pan  F32 F_XET
    /ch/15/mix/11/type  E32 F_XET enum=Xmtype
    /ch/15/mix/11/panFollow  E32 F_XET
/ch/15/mix/12  <CHME> n=0
    /ch/15/mix/12/on  E32 F_XET enum=OffOn
    /ch/15/mix/12/level  F32 F_XET
/ch/15/mix/13  <CHMO> n=0
    /ch/15/mix/13/on  E32 F_XET enum=OffOn
    /ch/15/mix/13/level  F32 F_XET
    /ch/15/mix/13/pan  F32 F_XET
    /ch/15/mix/13/type  E32 F_XET enum=Xmtype
    /ch/15/mix/13/panFollow  E32 F_XET
/ch/15/mix/14  <CHME> n=0
    /ch/15/mix/14/on  E32 F_XET enum=OffOn
    /ch/15/mix/14/level  F32 F_XET
/ch/15/mix/15  <CHMO> n=0
    /ch/15/mix/15/on  E32 F_XET enum=OffOn
    /ch/15/mix/15/level  F32 F_XET
    /ch/15/mix/15/pan  F32 F_XET
    /ch/15/mix/15/type  E32 F_XET enum=Xmtype
    /ch/15/mix/15/panFollow  E32 F_XET
/ch/15/mix/16  <CHME> n=0
    /ch/15/mix/16/on  E32 F_XET enum=OffOn
    /ch/15/mix/16/level  F32 F_XET
/ch/15/automix  <CHAMIX> n=0
    /ch/15/automix/group  E32 F_XET enum=Xamxgrp
    /ch/15/automix/weight  F32 F_XET
```

### Xchannel16 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/16  <CHCO> n=0
/ch/16/config  <CHCO> n=0
    /ch/16/config/name  S32 F_XET
    /ch/16/config/icon  I32 F_XET
    /ch/16/config/color  E32 F_XET enum=Xcolors
    /ch/16/config/source  I32 F_XET
/ch/16/grp  <CHGRP> n=0
    /ch/16/grp/dca  P32 F_XET
    /ch/16/grp/mute  P32 F_XET
/ch/16/preamp  <CHPR> n=0
    /ch/16/preamp/trim  F32 F_XET
    /ch/16/preamp/invert  E32 F_XET enum=OffOn
    /ch/16/preamp/hpon  E32 F_XET enum=OffOn
    /ch/16/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/16/preamp/hpf  F32 F_XET
/ch/16/delay  <CHDE> n=0
    /ch/16/delay/on  E32 F_XET enum=OffOn
    /ch/16/delay/time  F32 F_XET
/ch/16/insert  <CHIN> n=0
    /ch/16/insert/on  E32 F_XET enum=OffOn
    /ch/16/insert/pos  E32 F_XET enum=Xdyppos
    /ch/16/insert/sel  E32 F_XET enum=Xisel
/ch/16/gate  <CHGA> n=0
    /ch/16/gate/on  E32 F_XET enum=OffOn
    /ch/16/gate/mode  E32 F_XET enum=Xgmode
    /ch/16/gate/thr  F32 F_XET
    /ch/16/gate/range  F32 F_XET
    /ch/16/gate/attack  F32 F_XET
    /ch/16/gate/hold  F32 F_XET
    /ch/16/gate/release  F32 F_XET
    /ch/16/gate/keysrc  I32 F_XET
/ch/16/gate/filter  <CHGF> n=0
    /ch/16/gate/filter/on  E32 F_XET enum=OffOn
    /ch/16/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/16/gate/filter/f  F32 F_XET
/ch/16/dyn  <CHDY> n=0
    /ch/16/dyn/on  E32 F_XET enum=OffOn
    /ch/16/dyn/mode  E32 F_XET enum=Xdymode
    /ch/16/dyn/det  E32 F_XET enum=Xdydet
    /ch/16/dyn/env  E32 F_XET enum=Xdyenv
    /ch/16/dyn/thr  F32 F_XET
    /ch/16/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/16/dyn/knee  F32 F_XET
    /ch/16/dyn/mgain  F32 F_XET
    /ch/16/dyn/attack  F32 F_XET
    /ch/16/dyn/hold  F32 F_XET
    /ch/16/dyn/release  F32 F_XET
    /ch/16/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/16/dyn/keysrc  I32 F_XET
    /ch/16/dyn/mix  F32 F_XET
    /ch/16/dyn/auto  E32 F_XET enum=OffOn
/ch/16/dyn/filter  <CHDF> n=0
    /ch/16/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/16/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/16/dyn/filter/f  F32 F_XET
/ch/16/eq  <OFFON> n=1
    /ch/16/eq/on  E32 F_XET enum=OffOn
/ch/16/eq/1  <CHEQ> n=0
    /ch/16/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/16/eq/1/f  F32 F_XET
    /ch/16/eq/1/g  F32 F_XET
    /ch/16/eq/1/q  F32 F_XET
/ch/16/eq/2  <CHEQ> n=0
    /ch/16/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/16/eq/2/f  F32 F_XET
    /ch/16/eq/2/g  F32 F_XET
    /ch/16/eq/2/q  F32 F_XET
/ch/16/eq/3  <CHEQ> n=0
    /ch/16/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/16/eq/3/f  F32 F_XET
    /ch/16/eq/3/g  F32 F_XET
    /ch/16/eq/3/q  F32 F_XET
/ch/16/eq/4  <CHEQ> n=0
    /ch/16/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/16/eq/4/f  F32 F_XET
    /ch/16/eq/4/g  F32 F_XET
    /ch/16/eq/4/q  F32 F_XET
/ch/16/mix  <CHMX> n=0
    /ch/16/mix/on  E32 F_XET enum=OffOn
    /ch/16/mix/fader  F32 F_XET
    /ch/16/mix/st  E32 F_XET enum=OffOn
    /ch/16/mix/pan  F32 F_XET
    /ch/16/mix/mono  E32 F_XET enum=OffOn
    /ch/16/mix/mlevel  F32 F_XET
/ch/16/mix/01  <CHMO> n=0
    /ch/16/mix/01/on  E32 F_XET enum=OffOn
    /ch/16/mix/01/level  F32 F_XET
    /ch/16/mix/01/pan  F32 F_XET
    /ch/16/mix/01/type  E32 F_XET enum=Xmtype
    /ch/16/mix/01/panFollow  E32 F_XET
/ch/16/mix/02  <CHME> n=0
    /ch/16/mix/02/on  E32 F_XET enum=OffOn
    /ch/16/mix/02/level  F32 F_XET
/ch/16/mix/03  <CHMO> n=0
    /ch/16/mix/03/on  E32 F_XET enum=OffOn
    /ch/16/mix/03/level  F32 F_XET
    /ch/16/mix/03/pan  F32 F_XET
    /ch/16/mix/03/type  E32 F_XET enum=Xmtype
    /ch/16/mix/03/panFollow  E32 F_XET
/ch/16/mix/04  <CHME> n=0
    /ch/16/mix/04/on  E32 F_XET enum=OffOn
    /ch/16/mix/04/level  F32 F_XET
/ch/16/mix/05  <CHMO> n=0
    /ch/16/mix/05/on  E32 F_XET enum=OffOn
    /ch/16/mix/05/level  F32 F_XET
    /ch/16/mix/05/pan  F32 F_XET
    /ch/16/mix/05/type  E32 F_XET enum=Xmtype
    /ch/16/mix/05/panFollow  E32 F_XET
/ch/16/mix/06  <CHME> n=0
    /ch/16/mix/06/on  E32 F_XET enum=OffOn
    /ch/16/mix/06/level  F32 F_XET
/ch/16/mix/07  <CHMO> n=0
    /ch/16/mix/07/on  E32 F_XET enum=OffOn
    /ch/16/mix/07/level  F32 F_XET
    /ch/16/mix/07/pan  F32 F_XET
    /ch/16/mix/07/type  E32 F_XET enum=Xmtype
    /ch/16/mix/07/panFollow  E32 F_XET
/ch/16/mix/08  <CHME> n=0
    /ch/16/mix/08/on  E32 F_XET enum=OffOn
    /ch/16/mix/08/level  F32 F_XET
/ch/16/mix/09  <CHMO> n=0
    /ch/16/mix/09/on  E32 F_XET enum=OffOn
    /ch/16/mix/09/level  F32 F_XET
    /ch/16/mix/09/pan  F32 F_XET
    /ch/16/mix/09/type  E32 F_XET enum=Xmtype
    /ch/16/mix/09/panFollow  E32 F_XET
/ch/16/mix/10  <CHME> n=0
    /ch/16/mix/10/on  E32 F_XET enum=OffOn
    /ch/16/mix/10/level  F32 F_XET
/ch/16/mix/11  <CHMO> n=0
    /ch/16/mix/11/on  E32 F_XET enum=OffOn
    /ch/16/mix/11/level  F32 F_XET
    /ch/16/mix/11/pan  F32 F_XET
    /ch/16/mix/11/type  E32 F_XET enum=Xmtype
    /ch/16/mix/11/panFollow  E32 F_XET
/ch/16/mix/12  <CHME> n=0
    /ch/16/mix/12/on  E32 F_XET enum=OffOn
    /ch/16/mix/12/level  F32 F_XET
/ch/16/mix/13  <CHMO> n=0
    /ch/16/mix/13/on  E32 F_XET enum=OffOn
    /ch/16/mix/13/level  F32 F_XET
    /ch/16/mix/13/pan  F32 F_XET
    /ch/16/mix/13/type  E32 F_XET enum=Xmtype
    /ch/16/mix/13/panFollow  E32 F_XET
/ch/16/mix/14  <CHME> n=0
    /ch/16/mix/14/on  E32 F_XET enum=OffOn
    /ch/16/mix/14/level  F32 F_XET
/ch/16/mix/15  <CHMO> n=0
    /ch/16/mix/15/on  E32 F_XET enum=OffOn
    /ch/16/mix/15/level  F32 F_XET
    /ch/16/mix/15/pan  F32 F_XET
    /ch/16/mix/15/type  E32 F_XET enum=Xmtype
    /ch/16/mix/15/panFollow  E32 F_XET
/ch/16/mix/16  <CHME> n=0
    /ch/16/mix/16/on  E32 F_XET enum=OffOn
    /ch/16/mix/16/level  F32 F_XET
/ch/16/automix  <CHAMIX> n=0
    /ch/16/automix/group  E32 F_XET enum=Xamxgrp
    /ch/16/automix/weight  F32 F_XET
```

### Xchannel17 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/17  <CHCO> n=0
/ch/17/config  <CHCO> n=0
    /ch/17/config/name  S32 F_XET
    /ch/17/config/icon  I32 F_XET
    /ch/17/config/color  E32 F_XET enum=Xcolors
    /ch/17/config/source  I32 F_XET
/ch/17/grp  <CHGRP> n=0
    /ch/17/grp/dca  P32 F_XET
    /ch/17/grp/mute  P32 F_XET
/ch/17/preamp  <CHPR> n=0
    /ch/17/preamp/trim  F32 F_XET
    /ch/17/preamp/invert  E32 F_XET enum=OffOn
    /ch/17/preamp/hpon  E32 F_XET enum=OffOn
    /ch/17/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/17/preamp/hpf  F32 F_XET
/ch/17/delay  <CHDE> n=0
    /ch/17/delay/on  E32 F_XET enum=OffOn
    /ch/17/delay/time  F32 F_XET
/ch/17/insert  <CHIN> n=0
    /ch/17/insert/on  E32 F_XET enum=OffOn
    /ch/17/insert/pos  E32 F_XET enum=Xdyppos
    /ch/17/insert/sel  E32 F_XET enum=Xisel
/ch/17/gate  <CHGA> n=0
    /ch/17/gate/on  E32 F_XET enum=OffOn
    /ch/17/gate/mode  E32 F_XET enum=Xgmode
    /ch/17/gate/thr  F32 F_XET
    /ch/17/gate/range  F32 F_XET
    /ch/17/gate/attack  F32 F_XET
    /ch/17/gate/hold  F32 F_XET
    /ch/17/gate/release  F32 F_XET
    /ch/17/gate/keysrc  I32 F_XET
/ch/17/gate/filter  <CHGF> n=0
    /ch/17/gate/filter/on  E32 F_XET enum=OffOn
    /ch/17/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/17/gate/filter/f  F32 F_XET
/ch/17/dyn  <CHDY> n=0
    /ch/17/dyn/on  E32 F_XET enum=OffOn
    /ch/17/dyn/mode  E32 F_XET enum=Xdymode
    /ch/17/dyn/det  E32 F_XET enum=Xdydet
    /ch/17/dyn/env  E32 F_XET enum=Xdyenv
    /ch/17/dyn/thr  F32 F_XET
    /ch/17/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/17/dyn/knee  F32 F_XET
    /ch/17/dyn/mgain  F32 F_XET
    /ch/17/dyn/attack  F32 F_XET
    /ch/17/dyn/hold  F32 F_XET
    /ch/17/dyn/release  F32 F_XET
    /ch/17/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/17/dyn/keysrc  I32 F_XET
    /ch/17/dyn/mix  F32 F_XET
    /ch/17/dyn/auto  E32 F_XET enum=OffOn
/ch/17/dyn/filter  <CHDF> n=0
    /ch/17/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/17/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/17/dyn/filter/f  F32 F_XET
/ch/17/eq  <OFFON> n=1
    /ch/17/eq/on  E32 F_XET enum=OffOn
/ch/17/eq/1  <CHEQ> n=0
    /ch/17/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/17/eq/1/f  F32 F_XET
    /ch/17/eq/1/g  F32 F_XET
    /ch/17/eq/1/q  F32 F_XET
/ch/17/eq/2  <CHEQ> n=0
    /ch/17/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/17/eq/2/f  F32 F_XET
    /ch/17/eq/2/g  F32 F_XET
    /ch/17/eq/2/q  F32 F_XET
/ch/17/eq/3  <CHEQ> n=0
    /ch/17/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/17/eq/3/f  F32 F_XET
    /ch/17/eq/3/g  F32 F_XET
    /ch/17/eq/3/q  F32 F_XET
/ch/17/eq/4  <CHEQ> n=0
    /ch/17/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/17/eq/4/f  F32 F_XET
    /ch/17/eq/4/g  F32 F_XET
    /ch/17/eq/4/q  F32 F_XET
/ch/17/mix  <CHMX> n=0
    /ch/17/mix/on  E32 F_XET enum=OffOn
    /ch/17/mix/fader  F32 F_XET
    /ch/17/mix/st  E32 F_XET enum=OffOn
    /ch/17/mix/pan  F32 F_XET
    /ch/17/mix/mono  E32 F_XET enum=OffOn
    /ch/17/mix/mlevel  F32 F_XET
/ch/17/mix/01  <CHMO> n=0
    /ch/17/mix/01/on  E32 F_XET enum=OffOn
    /ch/17/mix/01/level  F32 F_XET
    /ch/17/mix/01/pan  F32 F_XET
    /ch/17/mix/01/type  E32 F_XET enum=Xmtype
    /ch/17/mix/01/panFollow  E32 F_XET
/ch/17/mix/02  <CHME> n=0
    /ch/17/mix/02/on  E32 F_XET enum=OffOn
    /ch/17/mix/02/level  F32 F_XET
/ch/17/mix/03  <CHMO> n=0
    /ch/17/mix/03/on  E32 F_XET enum=OffOn
    /ch/17/mix/03/level  F32 F_XET
    /ch/17/mix/03/pan  F32 F_XET
    /ch/17/mix/03/type  E32 F_XET enum=Xmtype
    /ch/17/mix/03/panFollow  E32 F_XET
/ch/17/mix/04  <CHME> n=0
    /ch/17/mix/04/on  E32 F_XET enum=OffOn
    /ch/17/mix/04/level  F32 F_XET
/ch/17/mix/05  <CHMO> n=0
    /ch/17/mix/05/on  E32 F_XET enum=OffOn
    /ch/17/mix/05/level  F32 F_XET
    /ch/17/mix/05/pan  F32 F_XET
    /ch/17/mix/05/type  E32 F_XET enum=Xmtype
    /ch/17/mix/05/panFollow  E32 F_XET
/ch/17/mix/06  <CHME> n=0
    /ch/17/mix/06/on  E32 F_XET enum=OffOn
    /ch/17/mix/06/level  F32 F_XET
/ch/17/mix/07  <CHMO> n=0
    /ch/17/mix/07/on  E32 F_XET enum=OffOn
    /ch/17/mix/07/level  F32 F_XET
    /ch/17/mix/07/pan  F32 F_XET
    /ch/17/mix/07/type  E32 F_XET enum=Xmtype
    /ch/17/mix/07/panFollow  E32 F_XET
/ch/17/mix/08  <CHME> n=0
    /ch/17/mix/08/on  E32 F_XET enum=OffOn
    /ch/17/mix/08/level  F32 F_XET
/ch/17/mix/09  <CHMO> n=0
    /ch/17/mix/09/on  E32 F_XET enum=OffOn
    /ch/17/mix/09/level  F32 F_XET
    /ch/17/mix/09/pan  F32 F_XET
    /ch/17/mix/09/type  E32 F_XET enum=Xmtype
    /ch/17/mix/09/panFollow  E32 F_XET
/ch/17/mix/10  <CHME> n=0
    /ch/17/mix/10/on  E32 F_XET enum=OffOn
    /ch/17/mix/10/level  F32 F_XET
/ch/17/mix/11  <CHMO> n=0
    /ch/17/mix/11/on  E32 F_XET enum=OffOn
    /ch/17/mix/11/level  F32 F_XET
    /ch/17/mix/11/pan  F32 F_XET
    /ch/17/mix/11/type  E32 F_XET enum=Xmtype
    /ch/17/mix/11/panFollow  E32 F_XET
/ch/17/mix/12  <CHME> n=0
    /ch/17/mix/12/on  E32 F_XET enum=OffOn
    /ch/17/mix/12/level  F32 F_XET
/ch/17/mix/13  <CHMO> n=0
    /ch/17/mix/13/on  E32 F_XET enum=OffOn
    /ch/17/mix/13/level  F32 F_XET
    /ch/17/mix/13/pan  F32 F_XET
    /ch/17/mix/13/type  E32 F_XET enum=Xmtype
    /ch/17/mix/13/panFollow  E32 F_XET
/ch/17/mix/14  <CHME> n=0
    /ch/17/mix/14/on  E32 F_XET enum=OffOn
    /ch/17/mix/14/level  F32 F_XET
/ch/17/mix/15  <CHMO> n=0
    /ch/17/mix/15/on  E32 F_XET enum=OffOn
    /ch/17/mix/15/level  F32 F_XET
    /ch/17/mix/15/pan  F32 F_XET
    /ch/17/mix/15/type  E32 F_XET enum=Xmtype
    /ch/17/mix/15/panFollow  E32 F_XET
/ch/17/mix/16  <CHME> n=0
    /ch/17/mix/16/on  E32 F_XET enum=OffOn
    /ch/17/mix/16/level  F32 F_XET
/ch/17/automix  <CHAMIX> n=0
    /ch/17/automix/group  E32 F_XET enum=Xamxgrp
    /ch/17/automix/weight  F32 F_XET
```

### Xchannel18 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/18  <CHCO> n=0
/ch/18/config  <CHCO> n=0
    /ch/18/config/name  S32 F_XET
    /ch/18/config/icon  I32 F_XET
    /ch/18/config/color  E32 F_XET enum=Xcolors
    /ch/18/config/source  I32 F_XET
/ch/18/grp  <CHGRP> n=0
    /ch/18/grp/dca  P32 F_XET
    /ch/18/grp/mute  P32 F_XET
/ch/18/preamp  <CHPR> n=0
    /ch/18/preamp/trim  F32 F_XET
    /ch/18/preamp/invert  E32 F_XET enum=OffOn
    /ch/18/preamp/hpon  E32 F_XET enum=OffOn
    /ch/18/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/18/preamp/hpf  F32 F_XET
/ch/18/delay  <CHDE> n=0
    /ch/18/delay/on  E32 F_XET enum=OffOn
    /ch/18/delay/time  F32 F_XET
/ch/18/insert  <CHIN> n=0
    /ch/18/insert/on  E32 F_XET enum=OffOn
    /ch/18/insert/pos  E32 F_XET enum=Xdyppos
    /ch/18/insert/sel  E32 F_XET enum=Xisel
/ch/18/gate  <CHGA> n=0
    /ch/18/gate/on  E32 F_XET enum=OffOn
    /ch/18/gate/mode  E32 F_XET enum=Xgmode
    /ch/18/gate/thr  F32 F_XET
    /ch/18/gate/range  F32 F_XET
    /ch/18/gate/attack  F32 F_XET
    /ch/18/gate/hold  F32 F_XET
    /ch/18/gate/release  F32 F_XET
    /ch/18/gate/keysrc  I32 F_XET
/ch/18/gate/filter  <CHGF> n=0
    /ch/18/gate/filter/on  E32 F_XET enum=OffOn
    /ch/18/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/18/gate/filter/f  F32 F_XET
/ch/18/dyn  <CHDY> n=0
    /ch/18/dyn/on  E32 F_XET enum=OffOn
    /ch/18/dyn/mode  E32 F_XET enum=Xdymode
    /ch/18/dyn/det  E32 F_XET enum=Xdydet
    /ch/18/dyn/env  E32 F_XET enum=Xdyenv
    /ch/18/dyn/thr  F32 F_XET
    /ch/18/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/18/dyn/knee  F32 F_XET
    /ch/18/dyn/mgain  F32 F_XET
    /ch/18/dyn/attack  F32 F_XET
    /ch/18/dyn/hold  F32 F_XET
    /ch/18/dyn/release  F32 F_XET
    /ch/18/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/18/dyn/keysrc  I32 F_XET
    /ch/18/dyn/mix  F32 F_XET
    /ch/18/dyn/auto  E32 F_XET enum=OffOn
/ch/18/dyn/filter  <CHDF> n=0
    /ch/18/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/18/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/18/dyn/filter/f  F32 F_XET
/ch/18/eq  <OFFON> n=1
    /ch/18/eq/on  E32 F_XET enum=OffOn
/ch/18/eq/1  <CHEQ> n=0
    /ch/18/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/18/eq/1/f  F32 F_XET
    /ch/18/eq/1/g  F32 F_XET
    /ch/18/eq/1/q  F32 F_XET
/ch/18/eq/2  <CHEQ> n=0
    /ch/18/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/18/eq/2/f  F32 F_XET
    /ch/18/eq/2/g  F32 F_XET
    /ch/18/eq/2/q  F32 F_XET
/ch/18/eq/3  <CHEQ> n=0
    /ch/18/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/18/eq/3/f  F32 F_XET
    /ch/18/eq/3/g  F32 F_XET
    /ch/18/eq/3/q  F32 F_XET
/ch/18/eq/4  <CHEQ> n=0
    /ch/18/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/18/eq/4/f  F32 F_XET
    /ch/18/eq/4/g  F32 F_XET
    /ch/18/eq/4/q  F32 F_XET
/ch/18/mix  <CHMX> n=0
    /ch/18/mix/on  E32 F_XET enum=OffOn
    /ch/18/mix/fader  F32 F_XET
    /ch/18/mix/st  E32 F_XET enum=OffOn
    /ch/18/mix/pan  F32 F_XET
    /ch/18/mix/mono  E32 F_XET enum=OffOn
    /ch/18/mix/mlevel  F32 F_XET
/ch/18/mix/01  <CHMO> n=0
    /ch/18/mix/01/on  E32 F_XET enum=OffOn
    /ch/18/mix/01/level  F32 F_XET
    /ch/18/mix/01/pan  F32 F_XET
    /ch/18/mix/01/type  E32 F_XET enum=Xmtype
    /ch/18/mix/01/panFollow  E32 F_XET
/ch/18/mix/02  <CHME> n=0
    /ch/18/mix/02/on  E32 F_XET enum=OffOn
    /ch/18/mix/02/level  F32 F_XET
/ch/18/mix/03  <CHMO> n=0
    /ch/18/mix/03/on  E32 F_XET enum=OffOn
    /ch/18/mix/03/level  F32 F_XET
    /ch/18/mix/03/pan  F32 F_XET
    /ch/18/mix/03/type  E32 F_XET enum=Xmtype
    /ch/18/mix/03/panFollow  E32 F_XET
/ch/18/mix/04  <CHME> n=0
    /ch/18/mix/04/on  E32 F_XET enum=OffOn
    /ch/18/mix/04/level  F32 F_XET
/ch/18/mix/05  <CHMO> n=0
    /ch/18/mix/05/on  E32 F_XET enum=OffOn
    /ch/18/mix/05/level  F32 F_XET
    /ch/18/mix/05/pan  F32 F_XET
    /ch/18/mix/05/type  E32 F_XET enum=Xmtype
    /ch/18/mix/05/panFollow  E32 F_XET
/ch/18/mix/06  <CHME> n=0
    /ch/18/mix/06/on  E32 F_XET enum=OffOn
    /ch/18/mix/06/level  F32 F_XET
/ch/18/mix/07  <CHMO> n=0
    /ch/18/mix/07/on  E32 F_XET enum=OffOn
    /ch/18/mix/07/level  F32 F_XET
    /ch/18/mix/07/pan  F32 F_XET
    /ch/18/mix/07/type  E32 F_XET enum=Xmtype
    /ch/18/mix/07/panFollow  E32 F_XET
/ch/18/mix/08  <CHME> n=0
    /ch/18/mix/08/on  E32 F_XET enum=OffOn
    /ch/18/mix/08/level  F32 F_XET
/ch/18/mix/09  <CHMO> n=0
    /ch/18/mix/09/on  E32 F_XET enum=OffOn
    /ch/18/mix/09/level  F32 F_XET
    /ch/18/mix/09/pan  F32 F_XET
    /ch/18/mix/09/type  E32 F_XET enum=Xmtype
    /ch/18/mix/09/panFollow  E32 F_XET
/ch/18/mix/10  <CHME> n=0
    /ch/18/mix/10/on  E32 F_XET enum=OffOn
    /ch/18/mix/10/level  F32 F_XET
/ch/18/mix/11  <CHMO> n=0
    /ch/18/mix/11/on  E32 F_XET enum=OffOn
    /ch/18/mix/11/level  F32 F_XET
    /ch/18/mix/11/pan  F32 F_XET
    /ch/18/mix/11/type  E32 F_XET enum=Xmtype
    /ch/18/mix/11/panFollow  E32 F_XET
/ch/18/mix/12  <CHME> n=0
    /ch/18/mix/12/on  E32 F_XET enum=OffOn
    /ch/18/mix/12/level  F32 F_XET
/ch/18/mix/13  <CHMO> n=0
    /ch/18/mix/13/on  E32 F_XET enum=OffOn
    /ch/18/mix/13/level  F32 F_XET
    /ch/18/mix/13/pan  F32 F_XET
    /ch/18/mix/13/type  E32 F_XET enum=Xmtype
    /ch/18/mix/13/panFollow  E32 F_XET
/ch/18/mix/14  <CHME> n=0
    /ch/18/mix/14/on  E32 F_XET enum=OffOn
    /ch/18/mix/14/level  F32 F_XET
/ch/18/mix/15  <CHMO> n=0
    /ch/18/mix/15/on  E32 F_XET enum=OffOn
    /ch/18/mix/15/level  F32 F_XET
    /ch/18/mix/15/pan  F32 F_XET
    /ch/18/mix/15/type  E32 F_XET enum=Xmtype
    /ch/18/mix/15/panFollow  E32 F_XET
/ch/18/mix/16  <CHME> n=0
    /ch/18/mix/16/on  E32 F_XET enum=OffOn
    /ch/18/mix/16/level  F32 F_XET
/ch/18/automix  <CHAMIX> n=0
    /ch/18/automix/group  E32 F_XET enum=Xamxgrp
    /ch/18/automix/weight  F32 F_XET
```

### Xchannel19 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/19  <CHCO> n=0
/ch/19/config  <CHCO> n=0
    /ch/19/config/name  S32 F_XET
    /ch/19/config/icon  I32 F_XET
    /ch/19/config/color  E32 F_XET enum=Xcolors
    /ch/19/config/source  I32 F_XET
/ch/19/grp  <CHGRP> n=0
    /ch/19/grp/dca  P32 F_XET
    /ch/19/grp/mute  P32 F_XET
/ch/19/preamp  <CHPR> n=0
    /ch/19/preamp/trim  F32 F_XET
    /ch/19/preamp/invert  E32 F_XET enum=OffOn
    /ch/19/preamp/hpon  E32 F_XET enum=OffOn
    /ch/19/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/19/preamp/hpf  F32 F_XET
/ch/19/delay  <CHDE> n=0
    /ch/19/delay/on  E32 F_XET enum=OffOn
    /ch/19/delay/time  F32 F_XET
/ch/19/insert  <CHIN> n=0
    /ch/19/insert/on  E32 F_XET enum=OffOn
    /ch/19/insert/pos  E32 F_XET enum=Xdyppos
    /ch/19/insert/sel  E32 F_XET enum=Xisel
/ch/19/gate  <CHGA> n=0
    /ch/19/gate/on  E32 F_XET enum=OffOn
    /ch/19/gate/mode  E32 F_XET enum=Xgmode
    /ch/19/gate/thr  F32 F_XET
    /ch/19/gate/range  F32 F_XET
    /ch/19/gate/attack  F32 F_XET
    /ch/19/gate/hold  F32 F_XET
    /ch/19/gate/release  F32 F_XET
    /ch/19/gate/keysrc  I32 F_XET
/ch/19/gate/filter  <CHGF> n=0
    /ch/19/gate/filter/on  E32 F_XET enum=OffOn
    /ch/19/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/19/gate/filter/f  F32 F_XET
/ch/19/dyn  <CHDY> n=0
    /ch/19/dyn/on  E32 F_XET enum=OffOn
    /ch/19/dyn/mode  E32 F_XET enum=Xdymode
    /ch/19/dyn/det  E32 F_XET enum=Xdydet
    /ch/19/dyn/env  E32 F_XET enum=Xdyenv
    /ch/19/dyn/thr  F32 F_XET
    /ch/19/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/19/dyn/knee  F32 F_XET
    /ch/19/dyn/mgain  F32 F_XET
    /ch/19/dyn/attack  F32 F_XET
    /ch/19/dyn/hold  F32 F_XET
    /ch/19/dyn/release  F32 F_XET
    /ch/19/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/19/dyn/keysrc  I32 F_XET
    /ch/19/dyn/mix  F32 F_XET
    /ch/19/dyn/auto  E32 F_XET enum=OffOn
/ch/19/dyn/filter  <CHDF> n=0
    /ch/19/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/19/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/19/dyn/filter/f  F32 F_XET
/ch/19/eq  <OFFON> n=1
    /ch/19/eq/on  E32 F_XET enum=OffOn
/ch/19/eq/1  <CHEQ> n=0
    /ch/19/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/19/eq/1/f  F32 F_XET
    /ch/19/eq/1/g  F32 F_XET
    /ch/19/eq/1/q  F32 F_XET
/ch/19/eq/2  <CHEQ> n=0
    /ch/19/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/19/eq/2/f  F32 F_XET
    /ch/19/eq/2/g  F32 F_XET
    /ch/19/eq/2/q  F32 F_XET
/ch/19/eq/3  <CHEQ> n=0
    /ch/19/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/19/eq/3/f  F32 F_XET
    /ch/19/eq/3/g  F32 F_XET
    /ch/19/eq/3/q  F32 F_XET
/ch/19/eq/4  <CHEQ> n=0
    /ch/19/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/19/eq/4/f  F32 F_XET
    /ch/19/eq/4/g  F32 F_XET
    /ch/19/eq/4/q  F32 F_XET
/ch/19/mix  <CHMX> n=0
    /ch/19/mix/on  E32 F_XET enum=OffOn
    /ch/19/mix/fader  F32 F_XET
    /ch/19/mix/st  E32 F_XET enum=OffOn
    /ch/19/mix/pan  F32 F_XET
    /ch/19/mix/mono  E32 F_XET enum=OffOn
    /ch/19/mix/mlevel  F32 F_XET
/ch/19/mix/01  <CHMO> n=0
    /ch/19/mix/01/on  E32 F_XET enum=OffOn
    /ch/19/mix/01/level  F32 F_XET
    /ch/19/mix/01/pan  F32 F_XET
    /ch/19/mix/01/type  E32 F_XET enum=Xmtype
    /ch/19/mix/01/panFollow  E32 F_XET
/ch/19/mix/02  <CHME> n=0
    /ch/19/mix/02/on  E32 F_XET enum=OffOn
    /ch/19/mix/02/level  F32 F_XET
/ch/19/mix/03  <CHMO> n=0
    /ch/19/mix/03/on  E32 F_XET enum=OffOn
    /ch/19/mix/03/level  F32 F_XET
    /ch/19/mix/03/pan  F32 F_XET
    /ch/19/mix/03/type  E32 F_XET enum=Xmtype
    /ch/19/mix/03/panFollow  E32 F_XET
/ch/19/mix/04  <CHME> n=0
    /ch/19/mix/04/on  E32 F_XET enum=OffOn
    /ch/19/mix/04/level  F32 F_XET
/ch/19/mix/05  <CHMO> n=0
    /ch/19/mix/05/on  E32 F_XET enum=OffOn
    /ch/19/mix/05/level  F32 F_XET
    /ch/19/mix/05/pan  F32 F_XET
    /ch/19/mix/05/type  E32 F_XET enum=Xmtype
    /ch/19/mix/05/panFollow  E32 F_XET
/ch/19/mix/06  <CHME> n=0
    /ch/19/mix/06/on  E32 F_XET enum=OffOn
    /ch/19/mix/06/level  F32 F_XET
/ch/19/mix/07  <CHMO> n=0
    /ch/19/mix/07/on  E32 F_XET enum=OffOn
    /ch/19/mix/07/level  F32 F_XET
    /ch/19/mix/07/pan  F32 F_XET
    /ch/19/mix/07/type  E32 F_XET enum=Xmtype
    /ch/19/mix/07/panFollow  E32 F_XET
/ch/19/mix/08  <CHME> n=0
    /ch/19/mix/08/on  E32 F_XET enum=OffOn
    /ch/19/mix/08/level  F32 F_XET
/ch/19/mix/09  <CHMO> n=0
    /ch/19/mix/09/on  E32 F_XET enum=OffOn
    /ch/19/mix/09/level  F32 F_XET
    /ch/19/mix/09/pan  F32 F_XET
    /ch/19/mix/09/type  E32 F_XET enum=Xmtype
    /ch/19/mix/09/panFollow  E32 F_XET
/ch/19/mix/10  <CHME> n=0
    /ch/19/mix/10/on  E32 F_XET enum=OffOn
    /ch/19/mix/10/level  F32 F_XET
/ch/19/mix/11  <CHMO> n=0
    /ch/19/mix/11/on  E32 F_XET enum=OffOn
    /ch/19/mix/11/level  F32 F_XET
    /ch/19/mix/11/pan  F32 F_XET
    /ch/19/mix/11/type  E32 F_XET enum=Xmtype
    /ch/19/mix/11/panFollow  E32 F_XET
/ch/19/mix/12  <CHME> n=0
    /ch/19/mix/12/on  E32 F_XET enum=OffOn
    /ch/19/mix/12/level  F32 F_XET
/ch/19/mix/13  <CHMO> n=0
    /ch/19/mix/13/on  E32 F_XET enum=OffOn
    /ch/19/mix/13/level  F32 F_XET
    /ch/19/mix/13/pan  F32 F_XET
    /ch/19/mix/13/type  E32 F_XET enum=Xmtype
    /ch/19/mix/13/panFollow  E32 F_XET
/ch/19/mix/14  <CHME> n=0
    /ch/19/mix/14/on  E32 F_XET enum=OffOn
    /ch/19/mix/14/level  F32 F_XET
/ch/19/mix/15  <CHMO> n=0
    /ch/19/mix/15/on  E32 F_XET enum=OffOn
    /ch/19/mix/15/level  F32 F_XET
    /ch/19/mix/15/pan  F32 F_XET
    /ch/19/mix/15/type  E32 F_XET enum=Xmtype
    /ch/19/mix/15/panFollow  E32 F_XET
/ch/19/mix/16  <CHME> n=0
    /ch/19/mix/16/on  E32 F_XET enum=OffOn
    /ch/19/mix/16/level  F32 F_XET
/ch/19/automix  <CHAMIX> n=0
    /ch/19/automix/group  E32 F_XET enum=Xamxgrp
    /ch/19/automix/weight  F32 F_XET
```

### Xchannel20 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/20  <CHCO> n=0
/ch/20/config  <CHCO> n=0
    /ch/20/config/name  S32 F_XET
    /ch/20/config/icon  I32 F_XET
    /ch/20/config/color  E32 F_XET enum=Xcolors
    /ch/20/config/source  I32 F_XET
/ch/20/grp  <CHGRP> n=0
    /ch/20/grp/dca  P32 F_XET
    /ch/20/grp/mute  P32 F_XET
/ch/20/preamp  <CHPR> n=0
    /ch/20/preamp/trim  F32 F_XET
    /ch/20/preamp/invert  E32 F_XET enum=OffOn
    /ch/20/preamp/hpon  E32 F_XET enum=OffOn
    /ch/20/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/20/preamp/hpf  F32 F_XET
/ch/20/delay  <CHDE> n=0
    /ch/20/delay/on  E32 F_XET enum=OffOn
    /ch/20/delay/time  F32 F_XET
/ch/20/insert  <CHIN> n=0
    /ch/20/insert/on  E32 F_XET enum=OffOn
    /ch/20/insert/pos  E32 F_XET enum=Xdyppos
    /ch/20/insert/sel  E32 F_XET enum=Xisel
/ch/20/gate  <CHGA> n=0
    /ch/20/gate/on  E32 F_XET enum=OffOn
    /ch/20/gate/mode  E32 F_XET enum=Xgmode
    /ch/20/gate/thr  F32 F_XET
    /ch/20/gate/range  F32 F_XET
    /ch/20/gate/attack  F32 F_XET
    /ch/20/gate/hold  F32 F_XET
    /ch/20/gate/release  F32 F_XET
    /ch/20/gate/keysrc  I32 F_XET
/ch/20/gate/filter  <CHGF> n=0
    /ch/20/gate/filter/on  E32 F_XET enum=OffOn
    /ch/20/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/20/gate/filter/f  F32 F_XET
/ch/20/dyn  <CHDY> n=0
    /ch/20/dyn/on  E32 F_XET enum=OffOn
    /ch/20/dyn/mode  E32 F_XET enum=Xdymode
    /ch/20/dyn/det  E32 F_XET enum=Xdydet
    /ch/20/dyn/env  E32 F_XET enum=Xdyenv
    /ch/20/dyn/thr  F32 F_XET
    /ch/20/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/20/dyn/knee  F32 F_XET
    /ch/20/dyn/mgain  F32 F_XET
    /ch/20/dyn/attack  F32 F_XET
    /ch/20/dyn/hold  F32 F_XET
    /ch/20/dyn/release  F32 F_XET
    /ch/20/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/20/dyn/keysrc  I32 F_XET
    /ch/20/dyn/mix  F32 F_XET
    /ch/20/dyn/auto  E32 F_XET enum=OffOn
/ch/20/dyn/filter  <CHDF> n=0
    /ch/20/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/20/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/20/dyn/filter/f  F32 F_XET
/ch/20/eq  <OFFON> n=1
    /ch/20/eq/on  E32 F_XET enum=OffOn
/ch/20/eq/1  <CHEQ> n=0
    /ch/20/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/20/eq/1/f  F32 F_XET
    /ch/20/eq/1/g  F32 F_XET
    /ch/20/eq/1/q  F32 F_XET
/ch/20/eq/2  <CHEQ> n=0
    /ch/20/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/20/eq/2/f  F32 F_XET
    /ch/20/eq/2/g  F32 F_XET
    /ch/20/eq/2/q  F32 F_XET
/ch/20/eq/3  <CHEQ> n=0
    /ch/20/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/20/eq/3/f  F32 F_XET
    /ch/20/eq/3/g  F32 F_XET
    /ch/20/eq/3/q  F32 F_XET
/ch/20/eq/4  <CHEQ> n=0
    /ch/20/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/20/eq/4/f  F32 F_XET
    /ch/20/eq/4/g  F32 F_XET
    /ch/20/eq/4/q  F32 F_XET
/ch/20/mix  <CHMX> n=0
    /ch/20/mix/on  E32 F_XET enum=OffOn
    /ch/20/mix/fader  F32 F_XET
    /ch/20/mix/st  E32 F_XET enum=OffOn
    /ch/20/mix/pan  F32 F_XET
    /ch/20/mix/mono  E32 F_XET enum=OffOn
    /ch/20/mix/mlevel  F32 F_XET
/ch/20/mix/01  <CHMO> n=0
    /ch/20/mix/01/on  E32 F_XET enum=OffOn
    /ch/20/mix/01/level  F32 F_XET
    /ch/20/mix/01/pan  F32 F_XET
    /ch/20/mix/01/type  E32 F_XET enum=Xmtype
    /ch/20/mix/01/panFollow  E32 F_XET
/ch/20/mix/02  <CHME> n=0
    /ch/20/mix/02/on  E32 F_XET enum=OffOn
    /ch/20/mix/02/level  F32 F_XET
/ch/20/mix/03  <CHMO> n=0
    /ch/20/mix/03/on  E32 F_XET enum=OffOn
    /ch/20/mix/03/level  F32 F_XET
    /ch/20/mix/03/pan  F32 F_XET
    /ch/20/mix/03/type  E32 F_XET enum=Xmtype
    /ch/20/mix/03/panFollow  E32 F_XET
/ch/20/mix/04  <CHME> n=0
    /ch/20/mix/04/on  E32 F_XET enum=OffOn
    /ch/20/mix/04/level  F32 F_XET
/ch/20/mix/05  <CHMO> n=0
    /ch/20/mix/05/on  E32 F_XET enum=OffOn
    /ch/20/mix/05/level  F32 F_XET
    /ch/20/mix/05/pan  F32 F_XET
    /ch/20/mix/05/type  E32 F_XET enum=Xmtype
    /ch/20/mix/05/panFollow  E32 F_XET
/ch/20/mix/06  <CHME> n=0
    /ch/20/mix/06/on  E32 F_XET enum=OffOn
    /ch/20/mix/06/level  F32 F_XET
/ch/20/mix/07  <CHMO> n=0
    /ch/20/mix/07/on  E32 F_XET enum=OffOn
    /ch/20/mix/07/level  F32 F_XET
    /ch/20/mix/07/pan  F32 F_XET
    /ch/20/mix/07/type  E32 F_XET enum=Xmtype
    /ch/20/mix/07/panFollow  E32 F_XET
/ch/20/mix/08  <CHME> n=0
    /ch/20/mix/08/on  E32 F_XET enum=OffOn
    /ch/20/mix/08/level  F32 F_XET
/ch/20/mix/09  <CHMO> n=0
    /ch/20/mix/09/on  E32 F_XET enum=OffOn
    /ch/20/mix/09/level  F32 F_XET
    /ch/20/mix/09/pan  F32 F_XET
    /ch/20/mix/09/type  E32 F_XET enum=Xmtype
    /ch/20/mix/09/panFollow  E32 F_XET
/ch/20/mix/10  <CHME> n=0
    /ch/20/mix/10/on  E32 F_XET enum=OffOn
    /ch/20/mix/10/level  F32 F_XET
/ch/20/mix/11  <CHMO> n=0
    /ch/20/mix/11/on  E32 F_XET enum=OffOn
    /ch/20/mix/11/level  F32 F_XET
    /ch/20/mix/11/pan  F32 F_XET
    /ch/20/mix/11/type  E32 F_XET enum=Xmtype
    /ch/20/mix/11/panFollow  E32 F_XET
/ch/20/mix/12  <CHME> n=0
    /ch/20/mix/12/on  E32 F_XET enum=OffOn
    /ch/20/mix/12/level  F32 F_XET
/ch/20/mix/13  <CHMO> n=0
    /ch/20/mix/13/on  E32 F_XET enum=OffOn
    /ch/20/mix/13/level  F32 F_XET
    /ch/20/mix/13/pan  F32 F_XET
    /ch/20/mix/13/type  E32 F_XET enum=Xmtype
    /ch/20/mix/13/panFollow  E32 F_XET
/ch/20/mix/14  <CHME> n=0
    /ch/20/mix/14/on  E32 F_XET enum=OffOn
    /ch/20/mix/14/level  F32 F_XET
/ch/20/mix/15  <CHMO> n=0
    /ch/20/mix/15/on  E32 F_XET enum=OffOn
    /ch/20/mix/15/level  F32 F_XET
    /ch/20/mix/15/pan  F32 F_XET
    /ch/20/mix/15/type  E32 F_XET enum=Xmtype
    /ch/20/mix/15/panFollow  E32 F_XET
/ch/20/mix/16  <CHME> n=0
    /ch/20/mix/16/on  E32 F_XET enum=OffOn
    /ch/20/mix/16/level  F32 F_XET
/ch/20/automix  <CHAMIX> n=0
    /ch/20/automix/group  E32 F_XET enum=Xamxgrp
    /ch/20/automix/weight  F32 F_XET
```

### Xchannel21 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/21  <CHCO> n=0
/ch/21/config  <CHCO> n=0
    /ch/21/config/name  S32 F_XET
    /ch/21/config/icon  I32 F_XET
    /ch/21/config/color  E32 F_XET enum=Xcolors
    /ch/21/config/source  I32 F_XET
/ch/21/grp  <CHGRP> n=0
    /ch/21/grp/dca  P32 F_XET
    /ch/21/grp/mute  P32 F_XET
/ch/21/preamp  <CHPR> n=0
    /ch/21/preamp/trim  F32 F_XET
    /ch/21/preamp/invert  E32 F_XET enum=OffOn
    /ch/21/preamp/hpon  E32 F_XET enum=OffOn
    /ch/21/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/21/preamp/hpf  F32 F_XET
/ch/21/delay  <CHDE> n=0
    /ch/21/delay/on  E32 F_XET enum=OffOn
    /ch/21/delay/time  F32 F_XET
/ch/21/insert  <CHIN> n=0
    /ch/21/insert/on  E32 F_XET enum=OffOn
    /ch/21/insert/pos  E32 F_XET enum=Xdyppos
    /ch/21/insert/sel  E32 F_XET enum=Xisel
/ch/21/gate  <CHGA> n=0
    /ch/21/gate/on  E32 F_XET enum=OffOn
    /ch/21/gate/mode  E32 F_XET enum=Xgmode
    /ch/21/gate/thr  F32 F_XET
    /ch/21/gate/range  F32 F_XET
    /ch/21/gate/attack  F32 F_XET
    /ch/21/gate/hold  F32 F_XET
    /ch/21/gate/release  F32 F_XET
    /ch/21/gate/keysrc  I32 F_XET
/ch/21/gate/filter  <CHGF> n=0
    /ch/21/gate/filter/on  E32 F_XET enum=OffOn
    /ch/21/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/21/gate/filter/f  F32 F_XET
/ch/21/dyn  <CHDY> n=0
    /ch/21/dyn/on  E32 F_XET enum=OffOn
    /ch/21/dyn/mode  E32 F_XET enum=Xdymode
    /ch/21/dyn/det  E32 F_XET enum=Xdydet
    /ch/21/dyn/env  E32 F_XET enum=Xdyenv
    /ch/21/dyn/thr  F32 F_XET
    /ch/21/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/21/dyn/knee  F32 F_XET
    /ch/21/dyn/mgain  F32 F_XET
    /ch/21/dyn/attack  F32 F_XET
    /ch/21/dyn/hold  F32 F_XET
    /ch/21/dyn/release  F32 F_XET
    /ch/21/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/21/dyn/keysrc  I32 F_XET
    /ch/21/dyn/mix  F32 F_XET
    /ch/21/dyn/auto  E32 F_XET enum=OffOn
/ch/21/dyn/filter  <CHDF> n=0
    /ch/21/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/21/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/21/dyn/filter/f  F32 F_XET
/ch/21/eq  <OFFON> n=1
    /ch/21/eq/on  E32 F_XET enum=OffOn
/ch/21/eq/1  <CHEQ> n=0
    /ch/21/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/21/eq/1/f  F32 F_XET
    /ch/21/eq/1/g  F32 F_XET
    /ch/21/eq/1/q  F32 F_XET
/ch/21/eq/2  <CHEQ> n=0
    /ch/21/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/21/eq/2/f  F32 F_XET
    /ch/21/eq/2/g  F32 F_XET
    /ch/21/eq/2/q  F32 F_XET
/ch/21/eq/3  <CHEQ> n=0
    /ch/21/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/21/eq/3/f  F32 F_XET
    /ch/21/eq/3/g  F32 F_XET
    /ch/21/eq/3/q  F32 F_XET
/ch/21/eq/4  <CHEQ> n=0
    /ch/21/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/21/eq/4/f  F32 F_XET
    /ch/21/eq/4/g  F32 F_XET
    /ch/21/eq/4/q  F32 F_XET
/ch/21/mix  <CHMX> n=0
    /ch/21/mix/on  E32 F_XET enum=OffOn
    /ch/21/mix/fader  F32 F_XET
    /ch/21/mix/st  E32 F_XET enum=OffOn
    /ch/21/mix/pan  F32 F_XET
    /ch/21/mix/mono  E32 F_XET enum=OffOn
    /ch/21/mix/mlevel  F32 F_XET
/ch/21/mix/01  <CHMO> n=0
    /ch/21/mix/01/on  E32 F_XET enum=OffOn
    /ch/21/mix/01/level  F32 F_XET
    /ch/21/mix/01/pan  F32 F_XET
    /ch/21/mix/01/type  E32 F_XET enum=Xmtype
    /ch/21/mix/01/panFollow  E32 F_XET
/ch/21/mix/02  <CHME> n=0
    /ch/21/mix/02/on  E32 F_XET enum=OffOn
    /ch/21/mix/02/level  F32 F_XET
/ch/21/mix/03  <CHMO> n=0
    /ch/21/mix/03/on  E32 F_XET enum=OffOn
    /ch/21/mix/03/level  F32 F_XET
    /ch/21/mix/03/pan  F32 F_XET
    /ch/21/mix/03/type  E32 F_XET enum=Xmtype
    /ch/21/mix/03/panFollow  E32 F_XET
/ch/21/mix/04  <CHME> n=0
    /ch/21/mix/04/on  E32 F_XET enum=OffOn
    /ch/21/mix/04/level  F32 F_XET
/ch/21/mix/05  <CHMO> n=0
    /ch/21/mix/05/on  E32 F_XET enum=OffOn
    /ch/21/mix/05/level  F32 F_XET
    /ch/21/mix/05/pan  F32 F_XET
    /ch/21/mix/05/type  E32 F_XET enum=Xmtype
    /ch/21/mix/05/panFollow  E32 F_XET
/ch/21/mix/06  <CHME> n=0
    /ch/21/mix/06/on  E32 F_XET enum=OffOn
    /ch/21/mix/06/level  F32 F_XET
/ch/21/mix/07  <CHMO> n=0
    /ch/21/mix/07/on  E32 F_XET enum=OffOn
    /ch/21/mix/07/level  F32 F_XET
    /ch/21/mix/07/pan  F32 F_XET
    /ch/21/mix/07/type  E32 F_XET enum=Xmtype
    /ch/21/mix/07/panFollow  E32 F_XET
/ch/21/mix/08  <CHME> n=0
    /ch/21/mix/08/on  E32 F_XET enum=OffOn
    /ch/21/mix/08/level  F32 F_XET
/ch/21/mix/09  <CHMO> n=0
    /ch/21/mix/09/on  E32 F_XET enum=OffOn
    /ch/21/mix/09/level  F32 F_XET
    /ch/21/mix/09/pan  F32 F_XET
    /ch/21/mix/09/type  E32 F_XET enum=Xmtype
    /ch/21/mix/09/panFollow  E32 F_XET
/ch/21/mix/10  <CHME> n=0
    /ch/21/mix/10/on  E32 F_XET enum=OffOn
    /ch/21/mix/10/level  F32 F_XET
/ch/21/mix/11  <CHMO> n=0
    /ch/21/mix/11/on  E32 F_XET enum=OffOn
    /ch/21/mix/11/level  F32 F_XET
    /ch/21/mix/11/pan  F32 F_XET
    /ch/21/mix/11/type  E32 F_XET enum=Xmtype
    /ch/21/mix/11/panFollow  E32 F_XET
/ch/21/mix/12  <CHME> n=0
    /ch/21/mix/12/on  E32 F_XET enum=OffOn
    /ch/21/mix/12/level  F32 F_XET
/ch/21/mix/13  <CHMO> n=0
    /ch/21/mix/13/on  E32 F_XET enum=OffOn
    /ch/21/mix/13/level  F32 F_XET
    /ch/21/mix/13/pan  F32 F_XET
    /ch/21/mix/13/type  E32 F_XET enum=Xmtype
    /ch/21/mix/13/panFollow  E32 F_XET
/ch/21/mix/14  <CHME> n=0
    /ch/21/mix/14/on  E32 F_XET enum=OffOn
    /ch/21/mix/14/level  F32 F_XET
/ch/21/mix/15  <CHMO> n=0
    /ch/21/mix/15/on  E32 F_XET enum=OffOn
    /ch/21/mix/15/level  F32 F_XET
    /ch/21/mix/15/pan  F32 F_XET
    /ch/21/mix/15/type  E32 F_XET enum=Xmtype
    /ch/21/mix/15/panFollow  E32 F_XET
/ch/21/mix/16  <CHME> n=0
    /ch/21/mix/16/on  E32 F_XET enum=OffOn
    /ch/21/mix/16/level  F32 F_XET
/ch/21/automix  <CHAMIX> n=0
    /ch/21/automix/group  E32 F_XET enum=Xamxgrp
    /ch/21/automix/weight  F32 F_XET
```

### Xchannel22 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/22  <CHCO> n=0
/ch/22/config  <CHCO> n=0
    /ch/22/config/name  S32 F_XET
    /ch/22/config/icon  I32 F_XET
    /ch/22/config/color  E32 F_XET enum=Xcolors
    /ch/22/config/source  I32 F_XET
/ch/22/grp  <CHGRP> n=0
    /ch/22/grp/dca  P32 F_XET
    /ch/22/grp/mute  P32 F_XET
/ch/22/preamp  <CHPR> n=0
    /ch/22/preamp/trim  F32 F_XET
    /ch/22/preamp/invert  E32 F_XET enum=OffOn
    /ch/22/preamp/hpon  E32 F_XET enum=OffOn
    /ch/22/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/22/preamp/hpf  F32 F_XET
/ch/22/delay  <CHDE> n=0
    /ch/22/delay/on  E32 F_XET enum=OffOn
    /ch/22/delay/time  F32 F_XET
/ch/22/insert  <CHIN> n=0
    /ch/22/insert/on  E32 F_XET enum=OffOn
    /ch/22/insert/pos  E32 F_XET enum=Xdyppos
    /ch/22/insert/sel  E32 F_XET enum=Xisel
/ch/22/gate  <CHGA> n=0
    /ch/22/gate/on  E32 F_XET enum=OffOn
    /ch/22/gate/mode  E32 F_XET enum=Xgmode
    /ch/22/gate/thr  F32 F_XET
    /ch/22/gate/range  F32 F_XET
    /ch/22/gate/attack  F32 F_XET
    /ch/22/gate/hold  F32 F_XET
    /ch/22/gate/release  F32 F_XET
    /ch/22/gate/keysrc  I32 F_XET
/ch/22/gate/filter  <CHGF> n=0
    /ch/22/gate/filter/on  E32 F_XET enum=OffOn
    /ch/22/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/22/gate/filter/f  F32 F_XET
/ch/22/dyn  <CHDY> n=0
    /ch/22/dyn/on  E32 F_XET enum=OffOn
    /ch/22/dyn/mode  E32 F_XET enum=Xdymode
    /ch/22/dyn/det  E32 F_XET enum=Xdydet
    /ch/22/dyn/env  E32 F_XET enum=Xdyenv
    /ch/22/dyn/thr  F32 F_XET
    /ch/22/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/22/dyn/knee  F32 F_XET
    /ch/22/dyn/mgain  F32 F_XET
    /ch/22/dyn/attack  F32 F_XET
    /ch/22/dyn/hold  F32 F_XET
    /ch/22/dyn/release  F32 F_XET
    /ch/22/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/22/dyn/keysrc  I32 F_XET
    /ch/22/dyn/mix  F32 F_XET
    /ch/22/dyn/auto  E32 F_XET enum=OffOn
/ch/22/dyn/filter  <CHDF> n=0
    /ch/22/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/22/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/22/dyn/filter/f  F32 F_XET
/ch/22/eq  <OFFON> n=1
    /ch/22/eq/on  E32 F_XET enum=OffOn
/ch/22/eq/1  <CHEQ> n=0
    /ch/22/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/22/eq/1/f  F32 F_XET
    /ch/22/eq/1/g  F32 F_XET
    /ch/22/eq/1/q  F32 F_XET
/ch/22/eq/2  <CHEQ> n=0
    /ch/22/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/22/eq/2/f  F32 F_XET
    /ch/22/eq/2/g  F32 F_XET
    /ch/22/eq/2/q  F32 F_XET
/ch/22/eq/3  <CHEQ> n=0
    /ch/22/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/22/eq/3/f  F32 F_XET
    /ch/22/eq/3/g  F32 F_XET
    /ch/22/eq/3/q  F32 F_XET
/ch/22/eq/4  <CHEQ> n=0
    /ch/22/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/22/eq/4/f  F32 F_XET
    /ch/22/eq/4/g  F32 F_XET
    /ch/22/eq/4/q  F32 F_XET
/ch/22/mix  <CHMX> n=0
    /ch/22/mix/on  E32 F_XET enum=OffOn
    /ch/22/mix/fader  F32 F_XET
    /ch/22/mix/st  E32 F_XET enum=OffOn
    /ch/22/mix/pan  F32 F_XET
    /ch/22/mix/mono  E32 F_XET enum=OffOn
    /ch/22/mix/mlevel  F32 F_XET
/ch/22/mix/01  <CHMO> n=0
    /ch/22/mix/01/on  E32 F_XET enum=OffOn
    /ch/22/mix/01/level  F32 F_XET
    /ch/22/mix/01/pan  F32 F_XET
    /ch/22/mix/01/type  E32 F_XET enum=Xmtype
    /ch/22/mix/01/panFollow  E32 F_XET
/ch/22/mix/02  <CHME> n=0
    /ch/22/mix/02/on  E32 F_XET enum=OffOn
    /ch/22/mix/02/level  F32 F_XET
/ch/22/mix/03  <CHMO> n=0
    /ch/22/mix/03/on  E32 F_XET enum=OffOn
    /ch/22/mix/03/level  F32 F_XET
    /ch/22/mix/03/pan  F32 F_XET
    /ch/22/mix/03/type  E32 F_XET enum=Xmtype
    /ch/22/mix/03/panFollow  E32 F_XET
/ch/22/mix/04  <CHME> n=0
    /ch/22/mix/04/on  E32 F_XET enum=OffOn
    /ch/22/mix/04/level  F32 F_XET
/ch/22/mix/05  <CHMO> n=0
    /ch/22/mix/05/on  E32 F_XET enum=OffOn
    /ch/22/mix/05/level  F32 F_XET
    /ch/22/mix/05/pan  F32 F_XET
    /ch/22/mix/05/type  E32 F_XET enum=Xmtype
    /ch/22/mix/05/panFollow  E32 F_XET
/ch/22/mix/06  <CHME> n=0
    /ch/22/mix/06/on  E32 F_XET enum=OffOn
    /ch/22/mix/06/level  F32 F_XET
/ch/22/mix/07  <CHMO> n=0
    /ch/22/mix/07/on  E32 F_XET enum=OffOn
    /ch/22/mix/07/level  F32 F_XET
    /ch/22/mix/07/pan  F32 F_XET
    /ch/22/mix/07/type  E32 F_XET enum=Xmtype
    /ch/22/mix/07/panFollow  E32 F_XET
/ch/22/mix/08  <CHME> n=0
    /ch/22/mix/08/on  E32 F_XET enum=OffOn
    /ch/22/mix/08/level  F32 F_XET
/ch/22/mix/09  <CHMO> n=0
    /ch/22/mix/09/on  E32 F_XET enum=OffOn
    /ch/22/mix/09/level  F32 F_XET
    /ch/22/mix/09/pan  F32 F_XET
    /ch/22/mix/09/type  E32 F_XET enum=Xmtype
    /ch/22/mix/09/panFollow  E32 F_XET
/ch/22/mix/10  <CHME> n=0
    /ch/22/mix/10/on  E32 F_XET enum=OffOn
    /ch/22/mix/10/level  F32 F_XET
/ch/22/mix/11  <CHMO> n=0
    /ch/22/mix/11/on  E32 F_XET enum=OffOn
    /ch/22/mix/11/level  F32 F_XET
    /ch/22/mix/11/pan  F32 F_XET
    /ch/22/mix/11/type  E32 F_XET enum=Xmtype
    /ch/22/mix/11/panFollow  E32 F_XET
/ch/22/mix/12  <CHME> n=0
    /ch/22/mix/12/on  E32 F_XET enum=OffOn
    /ch/22/mix/12/level  F32 F_XET
/ch/22/mix/13  <CHMO> n=0
    /ch/22/mix/13/on  E32 F_XET enum=OffOn
    /ch/22/mix/13/level  F32 F_XET
    /ch/22/mix/13/pan  F32 F_XET
    /ch/22/mix/13/type  E32 F_XET enum=Xmtype
    /ch/22/mix/13/panFollow  E32 F_XET
/ch/22/mix/14  <CHME> n=0
    /ch/22/mix/14/on  E32 F_XET enum=OffOn
    /ch/22/mix/14/level  F32 F_XET
/ch/22/mix/15  <CHMO> n=0
    /ch/22/mix/15/on  E32 F_XET enum=OffOn
    /ch/22/mix/15/level  F32 F_XET
    /ch/22/mix/15/pan  F32 F_XET
    /ch/22/mix/15/type  E32 F_XET enum=Xmtype
    /ch/22/mix/15/panFollow  E32 F_XET
/ch/22/mix/16  <CHME> n=0
    /ch/22/mix/16/on  E32 F_XET enum=OffOn
    /ch/22/mix/16/level  F32 F_XET
/ch/22/automix  <CHAMIX> n=0
    /ch/22/automix/group  E32 F_XET enum=Xamxgrp
    /ch/22/automix/weight  F32 F_XET
```

### Xchannel23 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/23  <CHCO> n=0
/ch/23/config  <CHCO> n=0
    /ch/23/config/name  S32 F_XET
    /ch/23/config/icon  I32 F_XET
    /ch/23/config/color  E32 F_XET enum=Xcolors
    /ch/23/config/source  I32 F_XET
/ch/23/grp  <CHGRP> n=0
    /ch/23/grp/dca  P32 F_XET
    /ch/23/grp/mute  P32 F_XET
/ch/23/preamp  <CHPR> n=0
    /ch/23/preamp/trim  F32 F_XET
    /ch/23/preamp/invert  E32 F_XET enum=OffOn
    /ch/23/preamp/hpon  E32 F_XET enum=OffOn
    /ch/23/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/23/preamp/hpf  F32 F_XET
/ch/23/delay  <CHDE> n=0
    /ch/23/delay/on  E32 F_XET enum=OffOn
    /ch/23/delay/time  F32 F_XET
/ch/23/insert  <CHIN> n=0
    /ch/23/insert/on  E32 F_XET enum=OffOn
    /ch/23/insert/pos  E32 F_XET enum=Xdyppos
    /ch/23/insert/sel  E32 F_XET enum=Xisel
/ch/23/gate  <CHGA> n=0
    /ch/23/gate/on  E32 F_XET enum=OffOn
    /ch/23/gate/mode  E32 F_XET enum=Xgmode
    /ch/23/gate/thr  F32 F_XET
    /ch/23/gate/range  F32 F_XET
    /ch/23/gate/attack  F32 F_XET
    /ch/23/gate/hold  F32 F_XET
    /ch/23/gate/release  F32 F_XET
    /ch/23/gate/keysrc  I32 F_XET
/ch/23/gate/filter  <CHGF> n=0
    /ch/23/gate/filter/on  E32 F_XET enum=OffOn
    /ch/23/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/23/gate/filter/f  F32 F_XET
/ch/23/dyn  <CHDY> n=0
    /ch/23/dyn/on  E32 F_XET enum=OffOn
    /ch/23/dyn/mode  E32 F_XET enum=Xdymode
    /ch/23/dyn/det  E32 F_XET enum=Xdydet
    /ch/23/dyn/env  E32 F_XET enum=Xdyenv
    /ch/23/dyn/thr  F32 F_XET
    /ch/23/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/23/dyn/knee  F32 F_XET
    /ch/23/dyn/mgain  F32 F_XET
    /ch/23/dyn/attack  F32 F_XET
    /ch/23/dyn/hold  F32 F_XET
    /ch/23/dyn/release  F32 F_XET
    /ch/23/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/23/dyn/keysrc  I32 F_XET
    /ch/23/dyn/mix  F32 F_XET
    /ch/23/dyn/auto  E32 F_XET enum=OffOn
/ch/23/dyn/filter  <CHDF> n=0
    /ch/23/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/23/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/23/dyn/filter/f  F32 F_XET
/ch/23/eq  <OFFON> n=1
    /ch/23/eq/on  E32 F_XET enum=OffOn
/ch/23/eq/1  <CHEQ> n=0
    /ch/23/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/23/eq/1/f  F32 F_XET
    /ch/23/eq/1/g  F32 F_XET
    /ch/23/eq/1/q  F32 F_XET
/ch/23/eq/2  <CHEQ> n=0
    /ch/23/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/23/eq/2/f  F32 F_XET
    /ch/23/eq/2/g  F32 F_XET
    /ch/23/eq/2/q  F32 F_XET
/ch/23/eq/3  <CHEQ> n=0
    /ch/23/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/23/eq/3/f  F32 F_XET
    /ch/23/eq/3/g  F32 F_XET
    /ch/23/eq/3/q  F32 F_XET
/ch/23/eq/4  <CHEQ> n=0
    /ch/23/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/23/eq/4/f  F32 F_XET
    /ch/23/eq/4/g  F32 F_XET
    /ch/23/eq/4/q  F32 F_XET
/ch/23/mix  <CHMX> n=0
    /ch/23/mix/on  E32 F_XET enum=OffOn
    /ch/23/mix/fader  F32 F_XET
    /ch/23/mix/st  E32 F_XET enum=OffOn
    /ch/23/mix/pan  F32 F_XET
    /ch/23/mix/mono  E32 F_XET enum=OffOn
    /ch/23/mix/mlevel  F32 F_XET
/ch/23/mix/01  <CHMO> n=0
    /ch/23/mix/01/on  E32 F_XET enum=OffOn
    /ch/23/mix/01/level  F32 F_XET
    /ch/23/mix/01/pan  F32 F_XET
    /ch/23/mix/01/type  E32 F_XET enum=Xmtype
    /ch/23/mix/01/panFollow  E32 F_XET
/ch/23/mix/02  <CHME> n=0
    /ch/23/mix/02/on  E32 F_XET enum=OffOn
    /ch/23/mix/02/level  F32 F_XET
/ch/23/mix/03  <CHMO> n=0
    /ch/23/mix/03/on  E32 F_XET enum=OffOn
    /ch/23/mix/03/level  F32 F_XET
    /ch/23/mix/03/pan  F32 F_XET
    /ch/23/mix/03/type  E32 F_XET enum=Xmtype
    /ch/23/mix/03/panFollow  E32 F_XET
/ch/23/mix/04  <CHME> n=0
    /ch/23/mix/04/on  E32 F_XET enum=OffOn
    /ch/23/mix/04/level  F32 F_XET
/ch/23/mix/05  <CHMO> n=0
    /ch/23/mix/05/on  E32 F_XET enum=OffOn
    /ch/23/mix/05/level  F32 F_XET
    /ch/23/mix/05/pan  F32 F_XET
    /ch/23/mix/05/type  E32 F_XET enum=Xmtype
    /ch/23/mix/05/panFollow  E32 F_XET
/ch/23/mix/06  <CHME> n=0
    /ch/23/mix/06/on  E32 F_XET enum=OffOn
    /ch/23/mix/06/level  F32 F_XET
/ch/23/mix/07  <CHMO> n=0
    /ch/23/mix/07/on  E32 F_XET enum=OffOn
    /ch/23/mix/07/level  F32 F_XET
    /ch/23/mix/07/pan  F32 F_XET
    /ch/23/mix/07/type  E32 F_XET enum=Xmtype
    /ch/23/mix/07/panFollow  E32 F_XET
/ch/23/mix/08  <CHME> n=0
    /ch/23/mix/08/on  E32 F_XET enum=OffOn
    /ch/23/mix/08/level  F32 F_XET
/ch/23/mix/09  <CHMO> n=0
    /ch/23/mix/09/on  E32 F_XET enum=OffOn
    /ch/23/mix/09/level  F32 F_XET
    /ch/23/mix/09/pan  F32 F_XET
    /ch/23/mix/09/type  E32 F_XET enum=Xmtype
    /ch/23/mix/09/panFollow  E32 F_XET
/ch/23/mix/10  <CHME> n=0
    /ch/23/mix/10/on  E32 F_XET enum=OffOn
    /ch/23/mix/10/level  F32 F_XET
/ch/23/mix/11  <CHMO> n=0
    /ch/23/mix/11/on  E32 F_XET enum=OffOn
    /ch/23/mix/11/level  F32 F_XET
    /ch/23/mix/11/pan  F32 F_XET
    /ch/23/mix/11/type  E32 F_XET enum=Xmtype
    /ch/23/mix/11/panFollow  E32 F_XET
/ch/23/mix/12  <CHME> n=0
    /ch/23/mix/12/on  E32 F_XET enum=OffOn
    /ch/23/mix/12/level  F32 F_XET
/ch/23/mix/13  <CHMO> n=0
    /ch/23/mix/13/on  E32 F_XET enum=OffOn
    /ch/23/mix/13/level  F32 F_XET
    /ch/23/mix/13/pan  F32 F_XET
    /ch/23/mix/13/type  E32 F_XET enum=Xmtype
    /ch/23/mix/13/panFollow  E32 F_XET
/ch/23/mix/14  <CHME> n=0
    /ch/23/mix/14/on  E32 F_XET enum=OffOn
    /ch/23/mix/14/level  F32 F_XET
/ch/23/mix/15  <CHMO> n=0
    /ch/23/mix/15/on  E32 F_XET enum=OffOn
    /ch/23/mix/15/level  F32 F_XET
    /ch/23/mix/15/pan  F32 F_XET
    /ch/23/mix/15/type  E32 F_XET enum=Xmtype
    /ch/23/mix/15/panFollow  E32 F_XET
/ch/23/mix/16  <CHME> n=0
    /ch/23/mix/16/on  E32 F_XET enum=OffOn
    /ch/23/mix/16/level  F32 F_XET
/ch/23/automix  <CHAMIX> n=0
    /ch/23/automix/group  E32 F_XET enum=Xamxgrp
    /ch/23/automix/weight  F32 F_XET
```

### Xchannel24 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/24  <CHCO> n=0
/ch/24/config  <CHCO> n=0
    /ch/24/config/name  S32 F_XET
    /ch/24/config/icon  I32 F_XET
    /ch/24/config/color  E32 F_XET enum=Xcolors
    /ch/24/config/source  I32 F_XET
/ch/24/grp  <CHGRP> n=0
    /ch/24/grp/dca  P32 F_XET
    /ch/24/grp/mute  P32 F_XET
/ch/24/preamp  <CHPR> n=0
    /ch/24/preamp/trim  F32 F_XET
    /ch/24/preamp/invert  E32 F_XET enum=OffOn
    /ch/24/preamp/hpon  E32 F_XET enum=OffOn
    /ch/24/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/24/preamp/hpf  F32 F_XET
/ch/24/delay  <CHDE> n=0
    /ch/24/delay/on  E32 F_XET enum=OffOn
    /ch/24/delay/time  F32 F_XET
/ch/24/insert  <CHIN> n=0
    /ch/24/insert/on  E32 F_XET enum=OffOn
    /ch/24/insert/pos  E32 F_XET enum=Xdyppos
    /ch/24/insert/sel  E32 F_XET enum=Xisel
/ch/24/gate  <CHGA> n=0
    /ch/24/gate/on  E32 F_XET enum=OffOn
    /ch/24/gate/mode  E32 F_XET enum=Xgmode
    /ch/24/gate/thr  F32 F_XET
    /ch/24/gate/range  F32 F_XET
    /ch/24/gate/attack  F32 F_XET
    /ch/24/gate/hold  F32 F_XET
    /ch/24/gate/release  F32 F_XET
    /ch/24/gate/keysrc  I32 F_XET
/ch/24/gate/filter  <CHGF> n=0
    /ch/24/gate/filter/on  E32 F_XET enum=OffOn
    /ch/24/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/24/gate/filter/f  F32 F_XET
/ch/24/dyn  <CHDY> n=0
    /ch/24/dyn/on  E32 F_XET enum=OffOn
    /ch/24/dyn/mode  E32 F_XET enum=Xdymode
    /ch/24/dyn/det  E32 F_XET enum=Xdydet
    /ch/24/dyn/env  E32 F_XET enum=Xdyenv
    /ch/24/dyn/thr  F32 F_XET
    /ch/24/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/24/dyn/knee  F32 F_XET
    /ch/24/dyn/mgain  F32 F_XET
    /ch/24/dyn/attack  F32 F_XET
    /ch/24/dyn/hold  F32 F_XET
    /ch/24/dyn/release  F32 F_XET
    /ch/24/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/24/dyn/keysrc  I32 F_XET
    /ch/24/dyn/mix  F32 F_XET
    /ch/24/dyn/auto  E32 F_XET enum=OffOn
/ch/24/dyn/filter  <CHDF> n=0
    /ch/24/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/24/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/24/dyn/filter/f  F32 F_XET
/ch/24/eq  <OFFON> n=1
    /ch/24/eq/on  E32 F_XET enum=OffOn
/ch/24/eq/1  <CHEQ> n=0
    /ch/24/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/24/eq/1/f  F32 F_XET
    /ch/24/eq/1/g  F32 F_XET
    /ch/24/eq/1/q  F32 F_XET
/ch/24/eq/2  <CHEQ> n=0
    /ch/24/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/24/eq/2/f  F32 F_XET
    /ch/24/eq/2/g  F32 F_XET
    /ch/24/eq/2/q  F32 F_XET
/ch/24/eq/3  <CHEQ> n=0
    /ch/24/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/24/eq/3/f  F32 F_XET
    /ch/24/eq/3/g  F32 F_XET
    /ch/24/eq/3/q  F32 F_XET
/ch/24/eq/4  <CHEQ> n=0
    /ch/24/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/24/eq/4/f  F32 F_XET
    /ch/24/eq/4/g  F32 F_XET
    /ch/24/eq/4/q  F32 F_XET
/ch/24/mix  <CHMX> n=0
    /ch/24/mix/on  E32 F_XET enum=OffOn
    /ch/24/mix/fader  F32 F_XET
    /ch/24/mix/st  E32 F_XET enum=OffOn
    /ch/24/mix/pan  F32 F_XET
    /ch/24/mix/mono  E32 F_XET enum=OffOn
    /ch/24/mix/mlevel  F32 F_XET
/ch/24/mix/01  <CHMO> n=0
    /ch/24/mix/01/on  E32 F_XET enum=OffOn
    /ch/24/mix/01/level  F32 F_XET
    /ch/24/mix/01/pan  F32 F_XET
    /ch/24/mix/01/type  E32 F_XET enum=Xmtype
    /ch/24/mix/01/panFollow  E32 F_XET
/ch/24/mix/02  <CHME> n=0
    /ch/24/mix/02/on  E32 F_XET enum=OffOn
    /ch/24/mix/02/level  F32 F_XET
/ch/24/mix/03  <CHMO> n=0
    /ch/24/mix/03/on  E32 F_XET enum=OffOn
    /ch/24/mix/03/level  F32 F_XET
    /ch/24/mix/03/pan  F32 F_XET
    /ch/24/mix/03/type  E32 F_XET enum=Xmtype
    /ch/24/mix/03/panFollow  E32 F_XET
/ch/24/mix/04  <CHME> n=0
    /ch/24/mix/04/on  E32 F_XET enum=OffOn
    /ch/24/mix/04/level  F32 F_XET
/ch/24/mix/05  <CHMO> n=0
    /ch/24/mix/05/on  E32 F_XET enum=OffOn
    /ch/24/mix/05/level  F32 F_XET
    /ch/24/mix/05/pan  F32 F_XET
    /ch/24/mix/05/type  E32 F_XET enum=Xmtype
    /ch/24/mix/05/panFollow  E32 F_XET
/ch/24/mix/06  <CHME> n=0
    /ch/24/mix/06/on  E32 F_XET enum=OffOn
    /ch/24/mix/06/level  F32 F_XET
/ch/24/mix/07  <CHMO> n=0
    /ch/24/mix/07/on  E32 F_XET enum=OffOn
    /ch/24/mix/07/level  F32 F_XET
    /ch/24/mix/07/pan  F32 F_XET
    /ch/24/mix/07/type  E32 F_XET enum=Xmtype
    /ch/24/mix/07/panFollow  E32 F_XET
/ch/24/mix/08  <CHME> n=0
    /ch/24/mix/08/on  E32 F_XET enum=OffOn
    /ch/24/mix/08/level  F32 F_XET
/ch/24/mix/09  <CHMO> n=0
    /ch/24/mix/09/on  E32 F_XET enum=OffOn
    /ch/24/mix/09/level  F32 F_XET
    /ch/24/mix/09/pan  F32 F_XET
    /ch/24/mix/09/type  E32 F_XET enum=Xmtype
    /ch/24/mix/09/panFollow  E32 F_XET
/ch/24/mix/10  <CHME> n=0
    /ch/24/mix/10/on  E32 F_XET enum=OffOn
    /ch/24/mix/10/level  F32 F_XET
/ch/24/mix/11  <CHMO> n=0
    /ch/24/mix/11/on  E32 F_XET enum=OffOn
    /ch/24/mix/11/level  F32 F_XET
    /ch/24/mix/11/pan  F32 F_XET
    /ch/24/mix/11/type  E32 F_XET enum=Xmtype
    /ch/24/mix/11/panFollow  E32 F_XET
/ch/24/mix/12  <CHME> n=0
    /ch/24/mix/12/on  E32 F_XET enum=OffOn
    /ch/24/mix/12/level  F32 F_XET
/ch/24/mix/13  <CHMO> n=0
    /ch/24/mix/13/on  E32 F_XET enum=OffOn
    /ch/24/mix/13/level  F32 F_XET
    /ch/24/mix/13/pan  F32 F_XET
    /ch/24/mix/13/type  E32 F_XET enum=Xmtype
    /ch/24/mix/13/panFollow  E32 F_XET
/ch/24/mix/14  <CHME> n=0
    /ch/24/mix/14/on  E32 F_XET enum=OffOn
    /ch/24/mix/14/level  F32 F_XET
/ch/24/mix/15  <CHMO> n=0
    /ch/24/mix/15/on  E32 F_XET enum=OffOn
    /ch/24/mix/15/level  F32 F_XET
    /ch/24/mix/15/pan  F32 F_XET
    /ch/24/mix/15/type  E32 F_XET enum=Xmtype
    /ch/24/mix/15/panFollow  E32 F_XET
/ch/24/mix/16  <CHME> n=0
    /ch/24/mix/16/on  E32 F_XET enum=OffOn
    /ch/24/mix/16/level  F32 F_XET
/ch/24/automix  <CHAMIX> n=0
    /ch/24/automix/group  E32 F_XET enum=Xamxgrp
    /ch/24/automix/weight  F32 F_XET
```

### Xchannel25 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/25  <CHCO> n=0
/ch/25/config  <CHCO> n=0
    /ch/25/config/name  S32 F_XET
    /ch/25/config/icon  I32 F_XET
    /ch/25/config/color  E32 F_XET enum=Xcolors
    /ch/25/config/source  I32 F_XET
/ch/25/grp  <CHGRP> n=0
    /ch/25/grp/dca  P32 F_XET
    /ch/25/grp/mute  P32 F_XET
/ch/25/preamp  <CHPR> n=0
    /ch/25/preamp/trim  F32 F_XET
    /ch/25/preamp/invert  E32 F_XET enum=OffOn
    /ch/25/preamp/hpon  E32 F_XET enum=OffOn
    /ch/25/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/25/preamp/hpf  F32 F_XET
/ch/25/delay  <CHDE> n=0
    /ch/25/delay/on  E32 F_XET enum=OffOn
    /ch/25/delay/time  F32 F_XET
/ch/25/insert  <CHIN> n=0
    /ch/25/insert/on  E32 F_XET enum=OffOn
    /ch/25/insert/pos  E32 F_XET enum=Xdyppos
    /ch/25/insert/sel  E32 F_XET enum=Xisel
/ch/25/gate  <CHGA> n=0
    /ch/25/gate/on  E32 F_XET enum=OffOn
    /ch/25/gate/mode  E32 F_XET enum=Xgmode
    /ch/25/gate/thr  F32 F_XET
    /ch/25/gate/range  F32 F_XET
    /ch/25/gate/attack  F32 F_XET
    /ch/25/gate/hold  F32 F_XET
    /ch/25/gate/release  F32 F_XET
    /ch/25/gate/keysrc  I32 F_XET
/ch/25/gate/filter  <CHGF> n=0
    /ch/25/gate/filter/on  E32 F_XET enum=OffOn
    /ch/25/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/25/gate/filter/f  F32 F_XET
/ch/25/dyn  <CHDY> n=0
    /ch/25/dyn/on  E32 F_XET enum=OffOn
    /ch/25/dyn/mode  E32 F_XET enum=Xdymode
    /ch/25/dyn/det  E32 F_XET enum=Xdydet
    /ch/25/dyn/env  E32 F_XET enum=Xdyenv
    /ch/25/dyn/thr  F32 F_XET
    /ch/25/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/25/dyn/knee  F32 F_XET
    /ch/25/dyn/mgain  F32 F_XET
    /ch/25/dyn/attack  F32 F_XET
    /ch/25/dyn/hold  F32 F_XET
    /ch/25/dyn/release  F32 F_XET
    /ch/25/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/25/dyn/keysrc  I32 F_XET
    /ch/25/dyn/mix  F32 F_XET
    /ch/25/dyn/auto  E32 F_XET enum=OffOn
/ch/25/dyn/filter  <CHDF> n=0
    /ch/25/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/25/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/25/dyn/filter/f  F32 F_XET
/ch/25/eq  <OFFON> n=1
    /ch/25/eq/on  E32 F_XET enum=OffOn
/ch/25/eq/1  <CHEQ> n=0
    /ch/25/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/25/eq/1/f  F32 F_XET
    /ch/25/eq/1/g  F32 F_XET
    /ch/25/eq/1/q  F32 F_XET
/ch/25/eq/2  <CHEQ> n=0
    /ch/25/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/25/eq/2/f  F32 F_XET
    /ch/25/eq/2/g  F32 F_XET
    /ch/25/eq/2/q  F32 F_XET
/ch/25/eq/3  <CHEQ> n=0
    /ch/25/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/25/eq/3/f  F32 F_XET
    /ch/25/eq/3/g  F32 F_XET
    /ch/25/eq/3/q  F32 F_XET
/ch/25/eq/4  <CHEQ> n=0
    /ch/25/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/25/eq/4/f  F32 F_XET
    /ch/25/eq/4/g  F32 F_XET
    /ch/25/eq/4/q  F32 F_XET
/ch/25/mix  <CHMX> n=0
    /ch/25/mix/on  E32 F_XET enum=OffOn
    /ch/25/mix/fader  F32 F_XET
    /ch/25/mix/st  E32 F_XET enum=OffOn
    /ch/25/mix/pan  F32 F_XET
    /ch/25/mix/mono  E32 F_XET enum=OffOn
    /ch/25/mix/mlevel  F32 F_XET
/ch/25/mix/01  <CHMO> n=0
    /ch/25/mix/01/on  E32 F_XET enum=OffOn
    /ch/25/mix/01/level  F32 F_XET
    /ch/25/mix/01/pan  F32 F_XET
    /ch/25/mix/01/type  E32 F_XET enum=Xmtype
    /ch/25/mix/01/panFollow  E32 F_XET
/ch/25/mix/02  <CHME> n=0
    /ch/25/mix/02/on  E32 F_XET enum=OffOn
    /ch/25/mix/02/level  F32 F_XET
/ch/25/mix/03  <CHMO> n=0
    /ch/25/mix/03/on  E32 F_XET enum=OffOn
    /ch/25/mix/03/level  F32 F_XET
    /ch/25/mix/03/pan  F32 F_XET
    /ch/25/mix/03/type  E32 F_XET enum=Xmtype
    /ch/25/mix/03/panFollow  E32 F_XET
/ch/25/mix/04  <CHME> n=0
    /ch/25/mix/04/on  E32 F_XET enum=OffOn
    /ch/25/mix/04/level  F32 F_XET
/ch/25/mix/05  <CHMO> n=0
    /ch/25/mix/05/on  E32 F_XET enum=OffOn
    /ch/25/mix/05/level  F32 F_XET
    /ch/25/mix/05/pan  F32 F_XET
    /ch/25/mix/05/type  E32 F_XET enum=Xmtype
    /ch/25/mix/05/panFollow  E32 F_XET
/ch/25/mix/06  <CHME> n=0
    /ch/25/mix/06/on  E32 F_XET enum=OffOn
    /ch/25/mix/06/level  F32 F_XET
/ch/25/mix/07  <CHMO> n=0
    /ch/25/mix/07/on  E32 F_XET enum=OffOn
    /ch/25/mix/07/level  F32 F_XET
    /ch/25/mix/07/pan  F32 F_XET
    /ch/25/mix/07/type  E32 F_XET enum=Xmtype
    /ch/25/mix/07/panFollow  E32 F_XET
/ch/25/mix/08  <CHME> n=0
    /ch/25/mix/08/on  E32 F_XET enum=OffOn
    /ch/25/mix/08/level  F32 F_XET
/ch/25/mix/09  <CHMO> n=0
    /ch/25/mix/09/on  E32 F_XET enum=OffOn
    /ch/25/mix/09/level  F32 F_XET
    /ch/25/mix/09/pan  F32 F_XET
    /ch/25/mix/09/type  E32 F_XET enum=Xmtype
    /ch/25/mix/09/panFollow  E32 F_XET
/ch/25/mix/10  <CHME> n=0
    /ch/25/mix/10/on  E32 F_XET enum=OffOn
    /ch/25/mix/10/level  F32 F_XET
/ch/25/mix/11  <CHMO> n=0
    /ch/25/mix/11/on  E32 F_XET enum=OffOn
    /ch/25/mix/11/level  F32 F_XET
    /ch/25/mix/11/pan  F32 F_XET
    /ch/25/mix/11/type  E32 F_XET enum=Xmtype
    /ch/25/mix/11/panFollow  E32 F_XET
/ch/25/mix/12  <CHME> n=0
    /ch/25/mix/12/on  E32 F_XET enum=OffOn
    /ch/25/mix/12/level  F32 F_XET
/ch/25/mix/13  <CHMO> n=0
    /ch/25/mix/13/on  E32 F_XET enum=OffOn
    /ch/25/mix/13/level  F32 F_XET
    /ch/25/mix/13/pan  F32 F_XET
    /ch/25/mix/13/type  E32 F_XET enum=Xmtype
    /ch/25/mix/13/panFollow  E32 F_XET
/ch/25/mix/14  <CHME> n=0
    /ch/25/mix/14/on  E32 F_XET enum=OffOn
    /ch/25/mix/14/level  F32 F_XET
/ch/25/mix/15  <CHMO> n=0
    /ch/25/mix/15/on  E32 F_XET enum=OffOn
    /ch/25/mix/15/level  F32 F_XET
    /ch/25/mix/15/pan  F32 F_XET
    /ch/25/mix/15/type  E32 F_XET enum=Xmtype
    /ch/25/mix/15/panFollow  E32 F_XET
/ch/25/mix/16  <CHME> n=0
    /ch/25/mix/16/on  E32 F_XET enum=OffOn
    /ch/25/mix/16/level  F32 F_XET
/ch/25/automix  <CHAMIX> n=0
    /ch/25/automix/group  E32 F_XET enum=Xamxgrp
    /ch/25/automix/weight  F32 F_XET
```

### Xchannel26 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/26  <CHCO> n=0
/ch/26/config  <CHCO> n=0
    /ch/26/config/name  S32 F_XET
    /ch/26/config/icon  I32 F_XET
    /ch/26/config/color  E32 F_XET enum=Xcolors
    /ch/26/config/source  I32 F_XET
/ch/26/grp  <CHGRP> n=0
    /ch/26/grp/dca  P32 F_XET
    /ch/26/grp/mute  P32 F_XET
/ch/26/preamp  <CHPR> n=0
    /ch/26/preamp/trim  F32 F_XET
    /ch/26/preamp/invert  E32 F_XET enum=OffOn
    /ch/26/preamp/hpon  E32 F_XET enum=OffOn
    /ch/26/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/26/preamp/hpf  F32 F_XET
/ch/26/delay  <CHDE> n=0
    /ch/26/delay/on  E32 F_XET enum=OffOn
    /ch/26/delay/time  F32 F_XET
/ch/26/insert  <CHIN> n=0
    /ch/26/insert/on  E32 F_XET enum=OffOn
    /ch/26/insert/pos  E32 F_XET enum=Xdyppos
    /ch/26/insert/sel  E32 F_XET enum=Xisel
/ch/26/gate  <CHGA> n=0
    /ch/26/gate/on  E32 F_XET enum=OffOn
    /ch/26/gate/mode  E32 F_XET enum=Xgmode
    /ch/26/gate/thr  F32 F_XET
    /ch/26/gate/range  F32 F_XET
    /ch/26/gate/attack  F32 F_XET
    /ch/26/gate/hold  F32 F_XET
    /ch/26/gate/release  F32 F_XET
    /ch/26/gate/keysrc  I32 F_XET
/ch/26/gate/filter  <CHGF> n=0
    /ch/26/gate/filter/on  E32 F_XET enum=OffOn
    /ch/26/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/26/gate/filter/f  F32 F_XET
/ch/26/dyn  <CHDY> n=0
    /ch/26/dyn/on  E32 F_XET enum=OffOn
    /ch/26/dyn/mode  E32 F_XET enum=Xdymode
    /ch/26/dyn/det  E32 F_XET enum=Xdydet
    /ch/26/dyn/env  E32 F_XET enum=Xdyenv
    /ch/26/dyn/thr  F32 F_XET
    /ch/26/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/26/dyn/knee  F32 F_XET
    /ch/26/dyn/mgain  F32 F_XET
    /ch/26/dyn/attack  F32 F_XET
    /ch/26/dyn/hold  F32 F_XET
    /ch/26/dyn/release  F32 F_XET
    /ch/26/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/26/dyn/keysrc  I32 F_XET
    /ch/26/dyn/mix  F32 F_XET
    /ch/26/dyn/auto  E32 F_XET enum=OffOn
/ch/26/dyn/filter  <CHDF> n=0
    /ch/26/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/26/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/26/dyn/filter/f  F32 F_XET
/ch/26/eq  <OFFON> n=1
    /ch/26/eq/on  E32 F_XET enum=OffOn
/ch/26/eq/1  <CHEQ> n=0
    /ch/26/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/26/eq/1/f  F32 F_XET
    /ch/26/eq/1/g  F32 F_XET
    /ch/26/eq/1/q  F32 F_XET
/ch/26/eq/2  <CHEQ> n=0
    /ch/26/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/26/eq/2/f  F32 F_XET
    /ch/26/eq/2/g  F32 F_XET
    /ch/26/eq/2/q  F32 F_XET
/ch/26/eq/3  <CHEQ> n=0
    /ch/26/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/26/eq/3/f  F32 F_XET
    /ch/26/eq/3/g  F32 F_XET
    /ch/26/eq/3/q  F32 F_XET
/ch/26/eq/4  <CHEQ> n=0
    /ch/26/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/26/eq/4/f  F32 F_XET
    /ch/26/eq/4/g  F32 F_XET
    /ch/26/eq/4/q  F32 F_XET
/ch/26/mix  <CHMX> n=0
    /ch/26/mix/on  E32 F_XET enum=OffOn
    /ch/26/mix/fader  F32 F_XET
    /ch/26/mix/st  E32 F_XET enum=OffOn
    /ch/26/mix/pan  F32 F_XET
    /ch/26/mix/mono  E32 F_XET enum=OffOn
    /ch/26/mix/mlevel  F32 F_XET
/ch/26/mix/01  <CHMO> n=0
    /ch/26/mix/01/on  E32 F_XET enum=OffOn
    /ch/26/mix/01/level  F32 F_XET
    /ch/26/mix/01/pan  F32 F_XET
    /ch/26/mix/01/type  E32 F_XET enum=Xmtype
    /ch/26/mix/01/panFollow  E32 F_XET
/ch/26/mix/02  <CHME> n=0
    /ch/26/mix/02/on  E32 F_XET enum=OffOn
    /ch/26/mix/02/level  F32 F_XET
/ch/26/mix/03  <CHMO> n=0
    /ch/26/mix/03/on  E32 F_XET enum=OffOn
    /ch/26/mix/03/level  F32 F_XET
    /ch/26/mix/03/pan  F32 F_XET
    /ch/26/mix/03/type  E32 F_XET enum=Xmtype
    /ch/26/mix/03/panFollow  E32 F_XET
/ch/26/mix/04  <CHME> n=0
    /ch/26/mix/04/on  E32 F_XET enum=OffOn
    /ch/26/mix/04/level  F32 F_XET
/ch/26/mix/05  <CHMO> n=0
    /ch/26/mix/05/on  E32 F_XET enum=OffOn
    /ch/26/mix/05/level  F32 F_XET
    /ch/26/mix/05/pan  F32 F_XET
    /ch/26/mix/05/type  E32 F_XET enum=Xmtype
    /ch/26/mix/05/panFollow  E32 F_XET
/ch/26/mix/06  <CHME> n=0
    /ch/26/mix/06/on  E32 F_XET enum=OffOn
    /ch/26/mix/06/level  F32 F_XET
/ch/26/mix/07  <CHMO> n=0
    /ch/26/mix/07/on  E32 F_XET enum=OffOn
    /ch/26/mix/07/level  F32 F_XET
    /ch/26/mix/07/pan  F32 F_XET
    /ch/26/mix/07/type  E32 F_XET enum=Xmtype
    /ch/26/mix/07/panFollow  E32 F_XET
/ch/26/mix/08  <CHME> n=0
    /ch/26/mix/08/on  E32 F_XET enum=OffOn
    /ch/26/mix/08/level  F32 F_XET
/ch/26/mix/09  <CHMO> n=0
    /ch/26/mix/09/on  E32 F_XET enum=OffOn
    /ch/26/mix/09/level  F32 F_XET
    /ch/26/mix/09/pan  F32 F_XET
    /ch/26/mix/09/type  E32 F_XET enum=Xmtype
    /ch/26/mix/09/panFollow  E32 F_XET
/ch/26/mix/10  <CHME> n=0
    /ch/26/mix/10/on  E32 F_XET enum=OffOn
    /ch/26/mix/10/level  F32 F_XET
/ch/26/mix/11  <CHMO> n=0
    /ch/26/mix/11/on  E32 F_XET enum=OffOn
    /ch/26/mix/11/level  F32 F_XET
    /ch/26/mix/11/pan  F32 F_XET
    /ch/26/mix/11/type  E32 F_XET enum=Xmtype
    /ch/26/mix/11/panFollow  E32 F_XET
/ch/26/mix/12  <CHME> n=0
    /ch/26/mix/12/on  E32 F_XET enum=OffOn
    /ch/26/mix/12/level  F32 F_XET
/ch/26/mix/13  <CHMO> n=0
    /ch/26/mix/13/on  E32 F_XET enum=OffOn
    /ch/26/mix/13/level  F32 F_XET
    /ch/26/mix/13/pan  F32 F_XET
    /ch/26/mix/13/type  E32 F_XET enum=Xmtype
    /ch/26/mix/13/panFollow  E32 F_XET
/ch/26/mix/14  <CHME> n=0
    /ch/26/mix/14/on  E32 F_XET enum=OffOn
    /ch/26/mix/14/level  F32 F_XET
/ch/26/mix/15  <CHMO> n=0
    /ch/26/mix/15/on  E32 F_XET enum=OffOn
    /ch/26/mix/15/level  F32 F_XET
    /ch/26/mix/15/pan  F32 F_XET
    /ch/26/mix/15/type  E32 F_XET enum=Xmtype
    /ch/26/mix/15/panFollow  E32 F_XET
/ch/26/mix/16  <CHME> n=0
    /ch/26/mix/16/on  E32 F_XET enum=OffOn
    /ch/26/mix/16/level  F32 F_XET
/ch/26/automix  <CHAMIX> n=0
    /ch/26/automix/group  E32 F_XET enum=Xamxgrp
    /ch/26/automix/weight  F32 F_XET
```

### Xchannel27 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/27  <CHCO> n=0
/ch/27/config  <CHCO> n=0
    /ch/27/config/name  S32 F_XET
    /ch/27/config/icon  I32 F_XET
    /ch/27/config/color  E32 F_XET enum=Xcolors
    /ch/27/config/source  I32 F_XET
/ch/27/grp  <CHGRP> n=0
    /ch/27/grp/dca  P32 F_XET
    /ch/27/grp/mute  P32 F_XET
/ch/27/preamp  <CHPR> n=0
    /ch/27/preamp/trim  F32 F_XET
    /ch/27/preamp/invert  E32 F_XET enum=OffOn
    /ch/27/preamp/hpon  E32 F_XET enum=OffOn
    /ch/27/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/27/preamp/hpf  F32 F_XET
/ch/27/delay  <CHDE> n=0
    /ch/27/delay/on  E32 F_XET enum=OffOn
    /ch/27/delay/time  F32 F_XET
/ch/27/insert  <CHIN> n=0
    /ch/27/insert/on  E32 F_XET enum=OffOn
    /ch/27/insert/pos  E32 F_XET enum=Xdyppos
    /ch/27/insert/sel  E32 F_XET enum=Xisel
/ch/27/gate  <CHGA> n=0
    /ch/27/gate/on  E32 F_XET enum=OffOn
    /ch/27/gate/mode  E32 F_XET enum=Xgmode
    /ch/27/gate/thr  F32 F_XET
    /ch/27/gate/range  F32 F_XET
    /ch/27/gate/attack  F32 F_XET
    /ch/27/gate/hold  F32 F_XET
    /ch/27/gate/release  F32 F_XET
    /ch/27/gate/keysrc  I32 F_XET
/ch/27/gate/filter  <CHGF> n=0
    /ch/27/gate/filter/on  E32 F_XET enum=OffOn
    /ch/27/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/27/gate/filter/f  F32 F_XET
/ch/27/dyn  <CHDY> n=0
    /ch/27/dyn/on  E32 F_XET enum=OffOn
    /ch/27/dyn/mode  E32 F_XET enum=Xdymode
    /ch/27/dyn/det  E32 F_XET enum=Xdydet
    /ch/27/dyn/env  E32 F_XET enum=Xdyenv
    /ch/27/dyn/thr  F32 F_XET
    /ch/27/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/27/dyn/knee  F32 F_XET
    /ch/27/dyn/mgain  F32 F_XET
    /ch/27/dyn/attack  F32 F_XET
    /ch/27/dyn/hold  F32 F_XET
    /ch/27/dyn/release  F32 F_XET
    /ch/27/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/27/dyn/keysrc  I32 F_XET
    /ch/27/dyn/mix  F32 F_XET
    /ch/27/dyn/auto  E32 F_XET enum=OffOn
/ch/27/dyn/filter  <CHDF> n=0
    /ch/27/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/27/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/27/dyn/filter/f  F32 F_XET
/ch/27/eq  <OFFON> n=1
    /ch/27/eq/on  E32 F_XET enum=OffOn
/ch/27/eq/1  <CHEQ> n=0
    /ch/27/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/27/eq/1/f  F32 F_XET
    /ch/27/eq/1/g  F32 F_XET
    /ch/27/eq/1/q  F32 F_XET
/ch/27/eq/2  <CHEQ> n=0
    /ch/27/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/27/eq/2/f  F32 F_XET
    /ch/27/eq/2/g  F32 F_XET
    /ch/27/eq/2/q  F32 F_XET
/ch/27/eq/3  <CHEQ> n=0
    /ch/27/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/27/eq/3/f  F32 F_XET
    /ch/27/eq/3/g  F32 F_XET
    /ch/27/eq/3/q  F32 F_XET
/ch/27/eq/4  <CHEQ> n=0
    /ch/27/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/27/eq/4/f  F32 F_XET
    /ch/27/eq/4/g  F32 F_XET
    /ch/27/eq/4/q  F32 F_XET
/ch/27/mix  <CHMX> n=0
    /ch/27/mix/on  E32 F_XET enum=OffOn
    /ch/27/mix/fader  F32 F_XET
    /ch/27/mix/st  E32 F_XET enum=OffOn
    /ch/27/mix/pan  F32 F_XET
    /ch/27/mix/mono  E32 F_XET enum=OffOn
    /ch/27/mix/mlevel  F32 F_XET
/ch/27/mix/01  <CHMO> n=0
    /ch/27/mix/01/on  E32 F_XET enum=OffOn
    /ch/27/mix/01/level  F32 F_XET
    /ch/27/mix/01/pan  F32 F_XET
    /ch/27/mix/01/type  E32 F_XET enum=Xmtype
    /ch/27/mix/01/panFollow  E32 F_XET
/ch/27/mix/02  <CHME> n=0
    /ch/27/mix/02/on  E32 F_XET enum=OffOn
    /ch/27/mix/02/level  F32 F_XET
/ch/27/mix/03  <CHMO> n=0
    /ch/27/mix/03/on  E32 F_XET enum=OffOn
    /ch/27/mix/03/level  F32 F_XET
    /ch/27/mix/03/pan  F32 F_XET
    /ch/27/mix/03/type  E32 F_XET enum=Xmtype
    /ch/27/mix/03/panFollow  E32 F_XET
/ch/27/mix/04  <CHME> n=0
    /ch/27/mix/04/on  E32 F_XET enum=OffOn
    /ch/27/mix/04/level  F32 F_XET
/ch/27/mix/05  <CHMO> n=0
    /ch/27/mix/05/on  E32 F_XET enum=OffOn
    /ch/27/mix/05/level  F32 F_XET
    /ch/27/mix/05/pan  F32 F_XET
    /ch/27/mix/05/type  E32 F_XET enum=Xmtype
    /ch/27/mix/05/panFollow  E32 F_XET
/ch/27/mix/06  <CHME> n=0
    /ch/27/mix/06/on  E32 F_XET enum=OffOn
    /ch/27/mix/06/level  F32 F_XET
/ch/27/mix/07  <CHMO> n=0
    /ch/27/mix/07/on  E32 F_XET enum=OffOn
    /ch/27/mix/07/level  F32 F_XET
    /ch/27/mix/07/pan  F32 F_XET
    /ch/27/mix/07/type  E32 F_XET enum=Xmtype
    /ch/27/mix/07/panFollow  E32 F_XET
/ch/27/mix/08  <CHME> n=0
    /ch/27/mix/08/on  E32 F_XET enum=OffOn
    /ch/27/mix/08/level  F32 F_XET
/ch/27/mix/09  <CHMO> n=0
    /ch/27/mix/09/on  E32 F_XET enum=OffOn
    /ch/27/mix/09/level  F32 F_XET
    /ch/27/mix/09/pan  F32 F_XET
    /ch/27/mix/09/type  E32 F_XET enum=Xmtype
    /ch/27/mix/09/panFollow  E32 F_XET
/ch/27/mix/10  <CHME> n=0
    /ch/27/mix/10/on  E32 F_XET enum=OffOn
    /ch/27/mix/10/level  F32 F_XET
/ch/27/mix/11  <CHMO> n=0
    /ch/27/mix/11/on  E32 F_XET enum=OffOn
    /ch/27/mix/11/level  F32 F_XET
    /ch/27/mix/11/pan  F32 F_XET
    /ch/27/mix/11/type  E32 F_XET enum=Xmtype
    /ch/27/mix/11/panFollow  E32 F_XET
/ch/27/mix/12  <CHME> n=0
    /ch/27/mix/12/on  E32 F_XET enum=OffOn
    /ch/27/mix/12/level  F32 F_XET
/ch/27/mix/13  <CHMO> n=0
    /ch/27/mix/13/on  E32 F_XET enum=OffOn
    /ch/27/mix/13/level  F32 F_XET
    /ch/27/mix/13/pan  F32 F_XET
    /ch/27/mix/13/type  E32 F_XET enum=Xmtype
    /ch/27/mix/13/panFollow  E32 F_XET
/ch/27/mix/14  <CHME> n=0
    /ch/27/mix/14/on  E32 F_XET enum=OffOn
    /ch/27/mix/14/level  F32 F_XET
/ch/27/mix/15  <CHMO> n=0
    /ch/27/mix/15/on  E32 F_XET enum=OffOn
    /ch/27/mix/15/level  F32 F_XET
    /ch/27/mix/15/pan  F32 F_XET
    /ch/27/mix/15/type  E32 F_XET enum=Xmtype
    /ch/27/mix/15/panFollow  E32 F_XET
/ch/27/mix/16  <CHME> n=0
    /ch/27/mix/16/on  E32 F_XET enum=OffOn
    /ch/27/mix/16/level  F32 F_XET
/ch/27/automix  <CHAMIX> n=0
    /ch/27/automix/group  E32 F_XET enum=Xamxgrp
    /ch/27/automix/weight  F32 F_XET
```

### Xchannel28 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/28  <CHCO> n=0
/ch/28/config  <CHCO> n=0
    /ch/28/config/name  S32 F_XET
    /ch/28/config/icon  I32 F_XET
    /ch/28/config/color  E32 F_XET enum=Xcolors
    /ch/28/config/source  I32 F_XET
/ch/28/grp  <CHGRP> n=0
    /ch/28/grp/dca  P32 F_XET
    /ch/28/grp/mute  P32 F_XET
/ch/28/preamp  <CHPR> n=0
    /ch/28/preamp/trim  F32 F_XET
    /ch/28/preamp/invert  E32 F_XET enum=OffOn
    /ch/28/preamp/hpon  E32 F_XET enum=OffOn
    /ch/28/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/28/preamp/hpf  F32 F_XET
/ch/28/delay  <CHDE> n=0
    /ch/28/delay/on  E32 F_XET enum=OffOn
    /ch/28/delay/time  F32 F_XET
/ch/28/insert  <CHIN> n=0
    /ch/28/insert/on  E32 F_XET enum=OffOn
    /ch/28/insert/pos  E32 F_XET enum=Xdyppos
    /ch/28/insert/sel  E32 F_XET enum=Xisel
/ch/28/gate  <CHGA> n=0
    /ch/28/gate/on  E32 F_XET enum=OffOn
    /ch/28/gate/mode  E32 F_XET enum=Xgmode
    /ch/28/gate/thr  F32 F_XET
    /ch/28/gate/range  F32 F_XET
    /ch/28/gate/attack  F32 F_XET
    /ch/28/gate/hold  F32 F_XET
    /ch/28/gate/release  F32 F_XET
    /ch/28/gate/keysrc  I32 F_XET
/ch/28/gate/filter  <CHGF> n=0
    /ch/28/gate/filter/on  E32 F_XET enum=OffOn
    /ch/28/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/28/gate/filter/f  F32 F_XET
/ch/28/dyn  <CHDY> n=0
    /ch/28/dyn/on  E32 F_XET enum=OffOn
    /ch/28/dyn/mode  E32 F_XET enum=Xdymode
    /ch/28/dyn/det  E32 F_XET enum=Xdydet
    /ch/28/dyn/env  E32 F_XET enum=Xdyenv
    /ch/28/dyn/thr  F32 F_XET
    /ch/28/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/28/dyn/knee  F32 F_XET
    /ch/28/dyn/mgain  F32 F_XET
    /ch/28/dyn/attack  F32 F_XET
    /ch/28/dyn/hold  F32 F_XET
    /ch/28/dyn/release  F32 F_XET
    /ch/28/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/28/dyn/keysrc  I32 F_XET
    /ch/28/dyn/mix  F32 F_XET
    /ch/28/dyn/auto  E32 F_XET enum=OffOn
/ch/28/dyn/filter  <CHDF> n=0
    /ch/28/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/28/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/28/dyn/filter/f  F32 F_XET
/ch/28/eq  <OFFON> n=1
    /ch/28/eq/on  E32 F_XET enum=OffOn
/ch/28/eq/1  <CHEQ> n=0
    /ch/28/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/28/eq/1/f  F32 F_XET
    /ch/28/eq/1/g  F32 F_XET
    /ch/28/eq/1/q  F32 F_XET
/ch/28/eq/2  <CHEQ> n=0
    /ch/28/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/28/eq/2/f  F32 F_XET
    /ch/28/eq/2/g  F32 F_XET
    /ch/28/eq/2/q  F32 F_XET
/ch/28/eq/3  <CHEQ> n=0
    /ch/28/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/28/eq/3/f  F32 F_XET
    /ch/28/eq/3/g  F32 F_XET
    /ch/28/eq/3/q  F32 F_XET
/ch/28/eq/4  <CHEQ> n=0
    /ch/28/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/28/eq/4/f  F32 F_XET
    /ch/28/eq/4/g  F32 F_XET
    /ch/28/eq/4/q  F32 F_XET
/ch/28/mix  <CHMX> n=0
    /ch/28/mix/on  E32 F_XET enum=OffOn
    /ch/28/mix/fader  F32 F_XET
    /ch/28/mix/st  E32 F_XET enum=OffOn
    /ch/28/mix/pan  F32 F_XET
    /ch/28/mix/mono  E32 F_XET enum=OffOn
    /ch/28/mix/mlevel  F32 F_XET
/ch/28/mix/01  <CHMO> n=0
    /ch/28/mix/01/on  E32 F_XET enum=OffOn
    /ch/28/mix/01/level  F32 F_XET
    /ch/28/mix/01/pan  F32 F_XET
    /ch/28/mix/01/type  E32 F_XET enum=Xmtype
    /ch/28/mix/01/panFollow  E32 F_XET
/ch/28/mix/02  <CHME> n=0
    /ch/28/mix/02/on  E32 F_XET enum=OffOn
    /ch/28/mix/02/level  F32 F_XET
/ch/28/mix/03  <CHMO> n=0
    /ch/28/mix/03/on  E32 F_XET enum=OffOn
    /ch/28/mix/03/level  F32 F_XET
    /ch/28/mix/03/pan  F32 F_XET
    /ch/28/mix/03/type  E32 F_XET enum=Xmtype
    /ch/28/mix/03/panFollow  E32 F_XET
/ch/28/mix/04  <CHME> n=0
    /ch/28/mix/04/on  E32 F_XET enum=OffOn
    /ch/28/mix/04/level  F32 F_XET
/ch/28/mix/05  <CHMO> n=0
    /ch/28/mix/05/on  E32 F_XET enum=OffOn
    /ch/28/mix/05/level  F32 F_XET
    /ch/28/mix/05/pan  F32 F_XET
    /ch/28/mix/05/type  E32 F_XET enum=Xmtype
    /ch/28/mix/05/panFollow  E32 F_XET
/ch/28/mix/06  <CHME> n=0
    /ch/28/mix/06/on  E32 F_XET enum=OffOn
    /ch/28/mix/06/level  F32 F_XET
/ch/28/mix/07  <CHMO> n=0
    /ch/28/mix/07/on  E32 F_XET enum=OffOn
    /ch/28/mix/07/level  F32 F_XET
    /ch/28/mix/07/pan  F32 F_XET
    /ch/28/mix/07/type  E32 F_XET enum=Xmtype
    /ch/28/mix/07/panFollow  E32 F_XET
/ch/28/mix/08  <CHME> n=0
    /ch/28/mix/08/on  E32 F_XET enum=OffOn
    /ch/28/mix/08/level  F32 F_XET
/ch/28/mix/09  <CHMO> n=0
    /ch/28/mix/09/on  E32 F_XET enum=OffOn
    /ch/28/mix/09/level  F32 F_XET
    /ch/28/mix/09/pan  F32 F_XET
    /ch/28/mix/09/type  E32 F_XET enum=Xmtype
    /ch/28/mix/09/panFollow  E32 F_XET
/ch/28/mix/10  <CHME> n=0
    /ch/28/mix/10/on  E32 F_XET enum=OffOn
    /ch/28/mix/10/level  F32 F_XET
/ch/28/mix/11  <CHMO> n=0
    /ch/28/mix/11/on  E32 F_XET enum=OffOn
    /ch/28/mix/11/level  F32 F_XET
    /ch/28/mix/11/pan  F32 F_XET
    /ch/28/mix/11/type  E32 F_XET enum=Xmtype
    /ch/28/mix/11/panFollow  E32 F_XET
/ch/28/mix/12  <CHME> n=0
    /ch/28/mix/12/on  E32 F_XET enum=OffOn
    /ch/28/mix/12/level  F32 F_XET
/ch/28/mix/13  <CHMO> n=0
    /ch/28/mix/13/on  E32 F_XET enum=OffOn
    /ch/28/mix/13/level  F32 F_XET
    /ch/28/mix/13/pan  F32 F_XET
    /ch/28/mix/13/type  E32 F_XET enum=Xmtype
    /ch/28/mix/13/panFollow  E32 F_XET
/ch/28/mix/14  <CHME> n=0
    /ch/28/mix/14/on  E32 F_XET enum=OffOn
    /ch/28/mix/14/level  F32 F_XET
/ch/28/mix/15  <CHMO> n=0
    /ch/28/mix/15/on  E32 F_XET enum=OffOn
    /ch/28/mix/15/level  F32 F_XET
    /ch/28/mix/15/pan  F32 F_XET
    /ch/28/mix/15/type  E32 F_XET enum=Xmtype
    /ch/28/mix/15/panFollow  E32 F_XET
/ch/28/mix/16  <CHME> n=0
    /ch/28/mix/16/on  E32 F_XET enum=OffOn
    /ch/28/mix/16/level  F32 F_XET
/ch/28/automix  <CHAMIX> n=0
    /ch/28/automix/group  E32 F_XET enum=Xamxgrp
    /ch/28/automix/weight  F32 F_XET
```

### Xchannel29 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/29  <CHCO> n=0
/ch/29/config  <CHCO> n=0
    /ch/29/config/name  S32 F_XET
    /ch/29/config/icon  I32 F_XET
    /ch/29/config/color  E32 F_XET enum=Xcolors
    /ch/29/config/source  I32 F_XET
/ch/29/grp  <CHGRP> n=0
    /ch/29/grp/dca  P32 F_XET
    /ch/29/grp/mute  P32 F_XET
/ch/29/preamp  <CHPR> n=0
    /ch/29/preamp/trim  F32 F_XET
    /ch/29/preamp/invert  E32 F_XET enum=OffOn
    /ch/29/preamp/hpon  E32 F_XET enum=OffOn
    /ch/29/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/29/preamp/hpf  F32 F_XET
/ch/29/delay  <CHDE> n=0
    /ch/29/delay/on  E32 F_XET enum=OffOn
    /ch/29/delay/time  F32 F_XET
/ch/29/insert  <CHIN> n=0
    /ch/29/insert/on  E32 F_XET enum=OffOn
    /ch/29/insert/pos  E32 F_XET enum=Xdyppos
    /ch/29/insert/sel  E32 F_XET enum=Xisel
/ch/29/gate  <CHGA> n=0
    /ch/29/gate/on  E32 F_XET enum=OffOn
    /ch/29/gate/mode  E32 F_XET enum=Xgmode
    /ch/29/gate/thr  F32 F_XET
    /ch/29/gate/range  F32 F_XET
    /ch/29/gate/attack  F32 F_XET
    /ch/29/gate/hold  F32 F_XET
    /ch/29/gate/release  F32 F_XET
    /ch/29/gate/keysrc  I32 F_XET
/ch/29/gate/filter  <CHGF> n=0
    /ch/29/gate/filter/on  E32 F_XET enum=OffOn
    /ch/29/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/29/gate/filter/f  F32 F_XET
/ch/29/dyn  <CHDY> n=0
    /ch/29/dyn/on  E32 F_XET enum=OffOn
    /ch/29/dyn/mode  E32 F_XET enum=Xdymode
    /ch/29/dyn/det  E32 F_XET enum=Xdydet
    /ch/29/dyn/env  E32 F_XET enum=Xdyenv
    /ch/29/dyn/thr  F32 F_XET
    /ch/29/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/29/dyn/knee  F32 F_XET
    /ch/29/dyn/mgain  F32 F_XET
    /ch/29/dyn/attack  F32 F_XET
    /ch/29/dyn/hold  F32 F_XET
    /ch/29/dyn/release  F32 F_XET
    /ch/29/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/29/dyn/keysrc  I32 F_XET
    /ch/29/dyn/mix  F32 F_XET
    /ch/29/dyn/auto  E32 F_XET enum=OffOn
/ch/29/dyn/filter  <CHDF> n=0
    /ch/29/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/29/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/29/dyn/filter/f  F32 F_XET
/ch/29/eq  <OFFON> n=1
    /ch/29/eq/on  E32 F_XET enum=OffOn
/ch/29/eq/1  <CHEQ> n=0
    /ch/29/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/29/eq/1/f  F32 F_XET
    /ch/29/eq/1/g  F32 F_XET
    /ch/29/eq/1/q  F32 F_XET
/ch/29/eq/2  <CHEQ> n=0
    /ch/29/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/29/eq/2/f  F32 F_XET
    /ch/29/eq/2/g  F32 F_XET
    /ch/29/eq/2/q  F32 F_XET
/ch/29/eq/3  <CHEQ> n=0
    /ch/29/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/29/eq/3/f  F32 F_XET
    /ch/29/eq/3/g  F32 F_XET
    /ch/29/eq/3/q  F32 F_XET
/ch/29/eq/4  <CHEQ> n=0
    /ch/29/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/29/eq/4/f  F32 F_XET
    /ch/29/eq/4/g  F32 F_XET
    /ch/29/eq/4/q  F32 F_XET
/ch/29/mix  <CHMX> n=0
    /ch/29/mix/on  E32 F_XET enum=OffOn
    /ch/29/mix/fader  F32 F_XET
    /ch/29/mix/st  E32 F_XET enum=OffOn
    /ch/29/mix/pan  F32 F_XET
    /ch/29/mix/mono  E32 F_XET enum=OffOn
    /ch/29/mix/mlevel  F32 F_XET
/ch/29/mix/01  <CHMO> n=0
    /ch/29/mix/01/on  E32 F_XET enum=OffOn
    /ch/29/mix/01/level  F32 F_XET
    /ch/29/mix/01/pan  F32 F_XET
    /ch/29/mix/01/type  E32 F_XET enum=Xmtype
    /ch/29/mix/01/panFollow  E32 F_XET
/ch/29/mix/02  <CHME> n=0
    /ch/29/mix/02/on  E32 F_XET enum=OffOn
    /ch/29/mix/02/level  F32 F_XET
/ch/29/mix/03  <CHMO> n=0
    /ch/29/mix/03/on  E32 F_XET enum=OffOn
    /ch/29/mix/03/level  F32 F_XET
    /ch/29/mix/03/pan  F32 F_XET
    /ch/29/mix/03/type  E32 F_XET enum=Xmtype
    /ch/29/mix/03/panFollow  E32 F_XET
/ch/29/mix/04  <CHME> n=0
    /ch/29/mix/04/on  E32 F_XET enum=OffOn
    /ch/29/mix/04/level  F32 F_XET
/ch/29/mix/05  <CHMO> n=0
    /ch/29/mix/05/on  E32 F_XET enum=OffOn
    /ch/29/mix/05/level  F32 F_XET
    /ch/29/mix/05/pan  F32 F_XET
    /ch/29/mix/05/type  E32 F_XET enum=Xmtype
    /ch/29/mix/05/panFollow  E32 F_XET
/ch/29/mix/06  <CHME> n=0
    /ch/29/mix/06/on  E32 F_XET enum=OffOn
    /ch/29/mix/06/level  F32 F_XET
/ch/29/mix/07  <CHMO> n=0
    /ch/29/mix/07/on  E32 F_XET enum=OffOn
    /ch/29/mix/07/level  F32 F_XET
    /ch/29/mix/07/pan  F32 F_XET
    /ch/29/mix/07/type  E32 F_XET enum=Xmtype
    /ch/29/mix/07/panFollow  E32 F_XET
/ch/29/mix/08  <CHME> n=0
    /ch/29/mix/08/on  E32 F_XET enum=OffOn
    /ch/29/mix/08/level  F32 F_XET
/ch/29/mix/09  <CHMO> n=0
    /ch/29/mix/09/on  E32 F_XET enum=OffOn
    /ch/29/mix/09/level  F32 F_XET
    /ch/29/mix/09/pan  F32 F_XET
    /ch/29/mix/09/type  E32 F_XET enum=Xmtype
    /ch/29/mix/09/panFollow  E32 F_XET
/ch/29/mix/10  <CHME> n=0
    /ch/29/mix/10/on  E32 F_XET enum=OffOn
    /ch/29/mix/10/level  F32 F_XET
/ch/29/mix/11  <CHMO> n=0
    /ch/29/mix/11/on  E32 F_XET enum=OffOn
    /ch/29/mix/11/level  F32 F_XET
    /ch/29/mix/11/pan  F32 F_XET
    /ch/29/mix/11/type  E32 F_XET enum=Xmtype
    /ch/29/mix/11/panFollow  E32 F_XET
/ch/29/mix/12  <CHME> n=0
    /ch/29/mix/12/on  E32 F_XET enum=OffOn
    /ch/29/mix/12/level  F32 F_XET
/ch/29/mix/13  <CHMO> n=0
    /ch/29/mix/13/on  E32 F_XET enum=OffOn
    /ch/29/mix/13/level  F32 F_XET
    /ch/29/mix/13/pan  F32 F_XET
    /ch/29/mix/13/type  E32 F_XET enum=Xmtype
    /ch/29/mix/13/panFollow  E32 F_XET
/ch/29/mix/14  <CHME> n=0
    /ch/29/mix/14/on  E32 F_XET enum=OffOn
    /ch/29/mix/14/level  F32 F_XET
/ch/29/mix/15  <CHMO> n=0
    /ch/29/mix/15/on  E32 F_XET enum=OffOn
    /ch/29/mix/15/level  F32 F_XET
    /ch/29/mix/15/pan  F32 F_XET
    /ch/29/mix/15/type  E32 F_XET enum=Xmtype
    /ch/29/mix/15/panFollow  E32 F_XET
/ch/29/mix/16  <CHME> n=0
    /ch/29/mix/16/on  E32 F_XET enum=OffOn
    /ch/29/mix/16/level  F32 F_XET
/ch/29/automix  <CHAMIX> n=0
    /ch/29/automix/group  E32 F_XET enum=Xamxgrp
    /ch/29/automix/weight  F32 F_XET
```

### Xchannel30 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/30  <CHCO> n=0
/ch/30/config  <CHCO> n=0
    /ch/30/config/name  S32 F_XET
    /ch/30/config/icon  I32 F_XET
    /ch/30/config/color  E32 F_XET enum=Xcolors
    /ch/30/config/source  I32 F_XET
/ch/30/grp  <CHGRP> n=0
    /ch/30/grp/dca  P32 F_XET
    /ch/30/grp/mute  P32 F_XET
/ch/30/preamp  <CHPR> n=0
    /ch/30/preamp/trim  F32 F_XET
    /ch/30/preamp/invert  E32 F_XET enum=OffOn
    /ch/30/preamp/hpon  E32 F_XET enum=OffOn
    /ch/30/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/30/preamp/hpf  F32 F_XET
/ch/30/delay  <CHDE> n=0
    /ch/30/delay/on  E32 F_XET enum=OffOn
    /ch/30/delay/time  F32 F_XET
/ch/30/insert  <CHIN> n=0
    /ch/30/insert/on  E32 F_XET enum=OffOn
    /ch/30/insert/pos  E32 F_XET enum=Xdyppos
    /ch/30/insert/sel  E32 F_XET enum=Xisel
/ch/30/gate  <CHGA> n=0
    /ch/30/gate/on  E32 F_XET enum=OffOn
    /ch/30/gate/mode  E32 F_XET enum=Xgmode
    /ch/30/gate/thr  F32 F_XET
    /ch/30/gate/range  F32 F_XET
    /ch/30/gate/attack  F32 F_XET
    /ch/30/gate/hold  F32 F_XET
    /ch/30/gate/release  F32 F_XET
    /ch/30/gate/keysrc  I32 F_XET
/ch/30/gate/filter  <CHGF> n=0
    /ch/30/gate/filter/on  E32 F_XET enum=OffOn
    /ch/30/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/30/gate/filter/f  F32 F_XET
/ch/30/dyn  <CHDY> n=0
    /ch/30/dyn/on  E32 F_XET enum=OffOn
    /ch/30/dyn/mode  E32 F_XET enum=Xdymode
    /ch/30/dyn/det  E32 F_XET enum=Xdydet
    /ch/30/dyn/env  E32 F_XET enum=Xdyenv
    /ch/30/dyn/thr  F32 F_XET
    /ch/30/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/30/dyn/knee  F32 F_XET
    /ch/30/dyn/mgain  F32 F_XET
    /ch/30/dyn/attack  F32 F_XET
    /ch/30/dyn/hold  F32 F_XET
    /ch/30/dyn/release  F32 F_XET
    /ch/30/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/30/dyn/keysrc  I32 F_XET
    /ch/30/dyn/mix  F32 F_XET
    /ch/30/dyn/auto  E32 F_XET enum=OffOn
/ch/30/dyn/filter  <CHDF> n=0
    /ch/30/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/30/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/30/dyn/filter/f  F32 F_XET
/ch/30/eq  <OFFON> n=1
    /ch/30/eq/on  E32 F_XET enum=OffOn
/ch/30/eq/1  <CHEQ> n=0
    /ch/30/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/30/eq/1/f  F32 F_XET
    /ch/30/eq/1/g  F32 F_XET
    /ch/30/eq/1/q  F32 F_XET
/ch/30/eq/2  <CHEQ> n=0
    /ch/30/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/30/eq/2/f  F32 F_XET
    /ch/30/eq/2/g  F32 F_XET
    /ch/30/eq/2/q  F32 F_XET
/ch/30/eq/3  <CHEQ> n=0
    /ch/30/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/30/eq/3/f  F32 F_XET
    /ch/30/eq/3/g  F32 F_XET
    /ch/30/eq/3/q  F32 F_XET
/ch/30/eq/4  <CHEQ> n=0
    /ch/30/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/30/eq/4/f  F32 F_XET
    /ch/30/eq/4/g  F32 F_XET
    /ch/30/eq/4/q  F32 F_XET
/ch/30/mix  <CHMX> n=0
    /ch/30/mix/on  E32 F_XET enum=OffOn
    /ch/30/mix/fader  F32 F_XET
    /ch/30/mix/st  E32 F_XET enum=OffOn
    /ch/30/mix/pan  F32 F_XET
    /ch/30/mix/mono  E32 F_XET enum=OffOn
    /ch/30/mix/mlevel  F32 F_XET
/ch/30/mix/01  <CHMO> n=0
    /ch/30/mix/01/on  E32 F_XET enum=OffOn
    /ch/30/mix/01/level  F32 F_XET
    /ch/30/mix/01/pan  F32 F_XET
    /ch/30/mix/01/type  E32 F_XET enum=Xmtype
    /ch/30/mix/01/panFollow  E32 F_XET
/ch/30/mix/02  <CHME> n=0
    /ch/30/mix/02/on  E32 F_XET enum=OffOn
    /ch/30/mix/02/level  F32 F_XET
/ch/30/mix/03  <CHMO> n=0
    /ch/30/mix/03/on  E32 F_XET enum=OffOn
    /ch/30/mix/03/level  F32 F_XET
    /ch/30/mix/03/pan  F32 F_XET
    /ch/30/mix/03/type  E32 F_XET enum=Xmtype
    /ch/30/mix/03/panFollow  E32 F_XET
/ch/30/mix/04  <CHME> n=0
    /ch/30/mix/04/on  E32 F_XET enum=OffOn
    /ch/30/mix/04/level  F32 F_XET
/ch/30/mix/05  <CHMO> n=0
    /ch/30/mix/05/on  E32 F_XET enum=OffOn
    /ch/30/mix/05/level  F32 F_XET
    /ch/30/mix/05/pan  F32 F_XET
    /ch/30/mix/05/type  E32 F_XET enum=Xmtype
    /ch/30/mix/05/panFollow  E32 F_XET
/ch/30/mix/06  <CHME> n=0
    /ch/30/mix/06/on  E32 F_XET enum=OffOn
    /ch/30/mix/06/level  F32 F_XET
/ch/30/mix/07  <CHMO> n=0
    /ch/30/mix/07/on  E32 F_XET enum=OffOn
    /ch/30/mix/07/level  F32 F_XET
    /ch/30/mix/07/pan  F32 F_XET
    /ch/30/mix/07/type  E32 F_XET enum=Xmtype
    /ch/30/mix/07/panFollow  E32 F_XET
/ch/30/mix/08  <CHME> n=0
    /ch/30/mix/08/on  E32 F_XET enum=OffOn
    /ch/30/mix/08/level  F32 F_XET
/ch/30/mix/09  <CHMO> n=0
    /ch/30/mix/09/on  E32 F_XET enum=OffOn
    /ch/30/mix/09/level  F32 F_XET
    /ch/30/mix/09/pan  F32 F_XET
    /ch/30/mix/09/type  E32 F_XET enum=Xmtype
    /ch/30/mix/09/panFollow  E32 F_XET
/ch/30/mix/10  <CHME> n=0
    /ch/30/mix/10/on  E32 F_XET enum=OffOn
    /ch/30/mix/10/level  F32 F_XET
/ch/30/mix/11  <CHMO> n=0
    /ch/30/mix/11/on  E32 F_XET enum=OffOn
    /ch/30/mix/11/level  F32 F_XET
    /ch/30/mix/11/pan  F32 F_XET
    /ch/30/mix/11/type  E32 F_XET enum=Xmtype
    /ch/30/mix/11/panFollow  E32 F_XET
/ch/30/mix/12  <CHME> n=0
    /ch/30/mix/12/on  E32 F_XET enum=OffOn
    /ch/30/mix/12/level  F32 F_XET
/ch/30/mix/13  <CHMO> n=0
    /ch/30/mix/13/on  E32 F_XET enum=OffOn
    /ch/30/mix/13/level  F32 F_XET
    /ch/30/mix/13/pan  F32 F_XET
    /ch/30/mix/13/type  E32 F_XET enum=Xmtype
    /ch/30/mix/13/panFollow  E32 F_XET
/ch/30/mix/14  <CHME> n=0
    /ch/30/mix/14/on  E32 F_XET enum=OffOn
    /ch/30/mix/14/level  F32 F_XET
/ch/30/mix/15  <CHMO> n=0
    /ch/30/mix/15/on  E32 F_XET enum=OffOn
    /ch/30/mix/15/level  F32 F_XET
    /ch/30/mix/15/pan  F32 F_XET
    /ch/30/mix/15/type  E32 F_XET enum=Xmtype
    /ch/30/mix/15/panFollow  E32 F_XET
/ch/30/mix/16  <CHME> n=0
    /ch/30/mix/16/on  E32 F_XET enum=OffOn
    /ch/30/mix/16/level  F32 F_XET
/ch/30/automix  <CHAMIX> n=0
    /ch/30/automix/group  E32 F_XET enum=Xamxgrp
    /ch/30/automix/weight  F32 F_XET
```

### Xchannel31 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/31  <CHCO> n=0
/ch/31/config  <CHCO> n=0
    /ch/31/config/name  S32 F_XET
    /ch/31/config/icon  I32 F_XET
    /ch/31/config/color  E32 F_XET enum=Xcolors
    /ch/31/config/source  I32 F_XET
/ch/31/grp  <CHGRP> n=0
    /ch/31/grp/dca  P32 F_XET
    /ch/31/grp/mute  P32 F_XET
/ch/31/preamp  <CHPR> n=0
    /ch/31/preamp/trim  F32 F_XET
    /ch/31/preamp/invert  E32 F_XET enum=OffOn
    /ch/31/preamp/hpon  E32 F_XET enum=OffOn
    /ch/31/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/31/preamp/hpf  F32 F_XET
/ch/31/delay  <CHDE> n=0
    /ch/31/delay/on  E32 F_XET enum=OffOn
    /ch/31/delay/time  F32 F_XET
/ch/31/insert  <CHIN> n=0
    /ch/31/insert/on  E32 F_XET enum=OffOn
    /ch/31/insert/pos  E32 F_XET enum=Xdyppos
    /ch/31/insert/sel  E32 F_XET enum=Xisel
/ch/31/gate  <CHGA> n=0
    /ch/31/gate/on  E32 F_XET enum=OffOn
    /ch/31/gate/mode  E32 F_XET enum=Xgmode
    /ch/31/gate/thr  F32 F_XET
    /ch/31/gate/range  F32 F_XET
    /ch/31/gate/attack  F32 F_XET
    /ch/31/gate/hold  F32 F_XET
    /ch/31/gate/release  F32 F_XET
    /ch/31/gate/keysrc  I32 F_XET
/ch/31/gate/filter  <CHGF> n=0
    /ch/31/gate/filter/on  E32 F_XET enum=OffOn
    /ch/31/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/31/gate/filter/f  F32 F_XET
/ch/31/dyn  <CHDY> n=0
    /ch/31/dyn/on  E32 F_XET enum=OffOn
    /ch/31/dyn/mode  E32 F_XET enum=Xdymode
    /ch/31/dyn/det  E32 F_XET enum=Xdydet
    /ch/31/dyn/env  E32 F_XET enum=Xdyenv
    /ch/31/dyn/thr  F32 F_XET
    /ch/31/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/31/dyn/knee  F32 F_XET
    /ch/31/dyn/mgain  F32 F_XET
    /ch/31/dyn/attack  F32 F_XET
    /ch/31/dyn/hold  F32 F_XET
    /ch/31/dyn/release  F32 F_XET
    /ch/31/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/31/dyn/keysrc  I32 F_XET
    /ch/31/dyn/mix  F32 F_XET
    /ch/31/dyn/auto  E32 F_XET enum=OffOn
/ch/31/dyn/filter  <CHDF> n=0
    /ch/31/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/31/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/31/dyn/filter/f  F32 F_XET
/ch/31/eq  <OFFON> n=1
    /ch/31/eq/on  E32 F_XET enum=OffOn
/ch/31/eq/1  <CHEQ> n=0
    /ch/31/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/31/eq/1/f  F32 F_XET
    /ch/31/eq/1/g  F32 F_XET
    /ch/31/eq/1/q  F32 F_XET
/ch/31/eq/2  <CHEQ> n=0
    /ch/31/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/31/eq/2/f  F32 F_XET
    /ch/31/eq/2/g  F32 F_XET
    /ch/31/eq/2/q  F32 F_XET
/ch/31/eq/3  <CHEQ> n=0
    /ch/31/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/31/eq/3/f  F32 F_XET
    /ch/31/eq/3/g  F32 F_XET
    /ch/31/eq/3/q  F32 F_XET
/ch/31/eq/4  <CHEQ> n=0
    /ch/31/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/31/eq/4/f  F32 F_XET
    /ch/31/eq/4/g  F32 F_XET
    /ch/31/eq/4/q  F32 F_XET
/ch/31/mix  <CHMX> n=0
    /ch/31/mix/on  E32 F_XET enum=OffOn
    /ch/31/mix/fader  F32 F_XET
    /ch/31/mix/st  E32 F_XET enum=OffOn
    /ch/31/mix/pan  F32 F_XET
    /ch/31/mix/mono  E32 F_XET enum=OffOn
    /ch/31/mix/mlevel  F32 F_XET
/ch/31/mix/01  <CHMO> n=0
    /ch/31/mix/01/on  E32 F_XET enum=OffOn
    /ch/31/mix/01/level  F32 F_XET
    /ch/31/mix/01/pan  F32 F_XET
    /ch/31/mix/01/type  E32 F_XET enum=Xmtype
    /ch/31/mix/01/panFollow  E32 F_XET
/ch/31/mix/02  <CHME> n=0
    /ch/31/mix/02/on  E32 F_XET enum=OffOn
    /ch/31/mix/02/level  F32 F_XET
/ch/31/mix/03  <CHMO> n=0
    /ch/31/mix/03/on  E32 F_XET enum=OffOn
    /ch/31/mix/03/level  F32 F_XET
    /ch/31/mix/03/pan  F32 F_XET
    /ch/31/mix/03/type  E32 F_XET enum=Xmtype
    /ch/31/mix/03/panFollow  E32 F_XET
/ch/31/mix/04  <CHME> n=0
    /ch/31/mix/04/on  E32 F_XET enum=OffOn
    /ch/31/mix/04/level  F32 F_XET
/ch/31/mix/05  <CHMO> n=0
    /ch/31/mix/05/on  E32 F_XET enum=OffOn
    /ch/31/mix/05/level  F32 F_XET
    /ch/31/mix/05/pan  F32 F_XET
    /ch/31/mix/05/type  E32 F_XET enum=Xmtype
    /ch/31/mix/05/panFollow  E32 F_XET
/ch/31/mix/06  <CHME> n=0
    /ch/31/mix/06/on  E32 F_XET enum=OffOn
    /ch/31/mix/06/level  F32 F_XET
/ch/31/mix/07  <CHMO> n=0
    /ch/31/mix/07/on  E32 F_XET enum=OffOn
    /ch/31/mix/07/level  F32 F_XET
    /ch/31/mix/07/pan  F32 F_XET
    /ch/31/mix/07/type  E32 F_XET enum=Xmtype
    /ch/31/mix/07/panFollow  E32 F_XET
/ch/31/mix/08  <CHME> n=0
    /ch/31/mix/08/on  E32 F_XET enum=OffOn
    /ch/31/mix/08/level  F32 F_XET
/ch/31/mix/09  <CHMO> n=0
    /ch/31/mix/09/on  E32 F_XET enum=OffOn
    /ch/31/mix/09/level  F32 F_XET
    /ch/31/mix/09/pan  F32 F_XET
    /ch/31/mix/09/type  E32 F_XET enum=Xmtype
    /ch/31/mix/09/panFollow  E32 F_XET
/ch/31/mix/10  <CHME> n=0
    /ch/31/mix/10/on  E32 F_XET enum=OffOn
    /ch/31/mix/10/level  F32 F_XET
/ch/31/mix/11  <CHMO> n=0
    /ch/31/mix/11/on  E32 F_XET enum=OffOn
    /ch/31/mix/11/level  F32 F_XET
    /ch/31/mix/11/pan  F32 F_XET
    /ch/31/mix/11/type  E32 F_XET enum=Xmtype
    /ch/31/mix/11/panFollow  E32 F_XET
/ch/31/mix/12  <CHME> n=0
    /ch/31/mix/12/on  E32 F_XET enum=OffOn
    /ch/31/mix/12/level  F32 F_XET
/ch/31/mix/13  <CHMO> n=0
    /ch/31/mix/13/on  E32 F_XET enum=OffOn
    /ch/31/mix/13/level  F32 F_XET
    /ch/31/mix/13/pan  F32 F_XET
    /ch/31/mix/13/type  E32 F_XET enum=Xmtype
    /ch/31/mix/13/panFollow  E32 F_XET
/ch/31/mix/14  <CHME> n=0
    /ch/31/mix/14/on  E32 F_XET enum=OffOn
    /ch/31/mix/14/level  F32 F_XET
/ch/31/mix/15  <CHMO> n=0
    /ch/31/mix/15/on  E32 F_XET enum=OffOn
    /ch/31/mix/15/level  F32 F_XET
    /ch/31/mix/15/pan  F32 F_XET
    /ch/31/mix/15/type  E32 F_XET enum=Xmtype
    /ch/31/mix/15/panFollow  E32 F_XET
/ch/31/mix/16  <CHME> n=0
    /ch/31/mix/16/on  E32 F_XET enum=OffOn
    /ch/31/mix/16/level  F32 F_XET
/ch/31/automix  <CHAMIX> n=0
    /ch/31/automix/group  E32 F_XET enum=Xamxgrp
    /ch/31/automix/weight  F32 F_XET
```

### Xchannel32 (X32Channel.h, 160 entries)

```
/ch  <CHCO> n=0
/ch/32  <CHCO> n=0
/ch/32/config  <CHCO> n=0
    /ch/32/config/name  S32 F_XET
    /ch/32/config/icon  I32 F_XET
    /ch/32/config/color  E32 F_XET enum=Xcolors
    /ch/32/config/source  I32 F_XET
/ch/32/grp  <CHGRP> n=0
    /ch/32/grp/dca  P32 F_XET
    /ch/32/grp/mute  P32 F_XET
/ch/32/preamp  <CHPR> n=0
    /ch/32/preamp/trim  F32 F_XET
    /ch/32/preamp/invert  E32 F_XET enum=OffOn
    /ch/32/preamp/hpon  E32 F_XET enum=OffOn
    /ch/32/preamp/hpslope  E32 F_XET enum=Xhslop
    /ch/32/preamp/hpf  F32 F_XET
/ch/32/delay  <CHDE> n=0
    /ch/32/delay/on  E32 F_XET enum=OffOn
    /ch/32/delay/time  F32 F_XET
/ch/32/insert  <CHIN> n=0
    /ch/32/insert/on  E32 F_XET enum=OffOn
    /ch/32/insert/pos  E32 F_XET enum=Xdyppos
    /ch/32/insert/sel  E32 F_XET enum=Xisel
/ch/32/gate  <CHGA> n=0
    /ch/32/gate/on  E32 F_XET enum=OffOn
    /ch/32/gate/mode  E32 F_XET enum=Xgmode
    /ch/32/gate/thr  F32 F_XET
    /ch/32/gate/range  F32 F_XET
    /ch/32/gate/attack  F32 F_XET
    /ch/32/gate/hold  F32 F_XET
    /ch/32/gate/release  F32 F_XET
    /ch/32/gate/keysrc  I32 F_XET
/ch/32/gate/filter  <CHGF> n=0
    /ch/32/gate/filter/on  E32 F_XET enum=OffOn
    /ch/32/gate/filter/type  E32 F_XET enum=Xdyftyp
    /ch/32/gate/filter/f  F32 F_XET
/ch/32/dyn  <CHDY> n=0
    /ch/32/dyn/on  E32 F_XET enum=OffOn
    /ch/32/dyn/mode  E32 F_XET enum=Xdymode
    /ch/32/dyn/det  E32 F_XET enum=Xdydet
    /ch/32/dyn/env  E32 F_XET enum=Xdyenv
    /ch/32/dyn/thr  F32 F_XET
    /ch/32/dyn/ratio  E32 F_XET enum=Xdyrat
    /ch/32/dyn/knee  F32 F_XET
    /ch/32/dyn/mgain  F32 F_XET
    /ch/32/dyn/attack  F32 F_XET
    /ch/32/dyn/hold  F32 F_XET
    /ch/32/dyn/release  F32 F_XET
    /ch/32/dyn/pos  E32 F_XET enum=Xdyppos
    /ch/32/dyn/keysrc  I32 F_XET
    /ch/32/dyn/mix  F32 F_XET
    /ch/32/dyn/auto  E32 F_XET enum=OffOn
/ch/32/dyn/filter  <CHDF> n=0
    /ch/32/dyn/filter/on  E32 F_XET enum=OffOn
    /ch/32/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /ch/32/dyn/filter/f  F32 F_XET
/ch/32/eq  <OFFON> n=1
    /ch/32/eq/on  E32 F_XET enum=OffOn
/ch/32/eq/1  <CHEQ> n=0
    /ch/32/eq/1/type  E32 F_XET enum=Xeqty1
    /ch/32/eq/1/f  F32 F_XET
    /ch/32/eq/1/g  F32 F_XET
    /ch/32/eq/1/q  F32 F_XET
/ch/32/eq/2  <CHEQ> n=0
    /ch/32/eq/2/type  E32 F_XET enum=Xeqty1
    /ch/32/eq/2/f  F32 F_XET
    /ch/32/eq/2/g  F32 F_XET
    /ch/32/eq/2/q  F32 F_XET
/ch/32/eq/3  <CHEQ> n=0
    /ch/32/eq/3/type  E32 F_XET enum=Xeqty1
    /ch/32/eq/3/f  F32 F_XET
    /ch/32/eq/3/g  F32 F_XET
    /ch/32/eq/3/q  F32 F_XET
/ch/32/eq/4  <CHEQ> n=0
    /ch/32/eq/4/type  E32 F_XET enum=Xeqty1
    /ch/32/eq/4/f  F32 F_XET
    /ch/32/eq/4/g  F32 F_XET
    /ch/32/eq/4/q  F32 F_XET
/ch/32/mix  <CHMX> n=0
    /ch/32/mix/on  E32 F_XET enum=OffOn
    /ch/32/mix/fader  F32 F_XET
    /ch/32/mix/st  E32 F_XET enum=OffOn
    /ch/32/mix/pan  F32 F_XET
    /ch/32/mix/mono  E32 F_XET enum=OffOn
    /ch/32/mix/mlevel  F32 F_XET
/ch/32/mix/01  <CHMO> n=0
    /ch/32/mix/01/on  E32 F_XET enum=OffOn
    /ch/32/mix/01/level  F32 F_XET
    /ch/32/mix/01/pan  F32 F_XET
    /ch/32/mix/01/type  E32 F_XET enum=Xmtype
    /ch/32/mix/01/panFollow  E32 F_XET
/ch/32/mix/02  <CHME> n=0
    /ch/32/mix/02/on  E32 F_XET enum=OffOn
    /ch/32/mix/02/level  F32 F_XET
/ch/32/mix/03  <CHMO> n=0
    /ch/32/mix/03/on  E32 F_XET enum=OffOn
    /ch/32/mix/03/level  F32 F_XET
    /ch/32/mix/03/pan  F32 F_XET
    /ch/32/mix/03/type  E32 F_XET enum=Xmtype
    /ch/32/mix/03/panFollow  E32 F_XET
/ch/32/mix/04  <CHME> n=0
    /ch/32/mix/04/on  E32 F_XET enum=OffOn
    /ch/32/mix/04/level  F32 F_XET
/ch/32/mix/05  <CHMO> n=0
    /ch/32/mix/05/on  E32 F_XET enum=OffOn
    /ch/32/mix/05/level  F32 F_XET
    /ch/32/mix/05/pan  F32 F_XET
    /ch/32/mix/05/type  E32 F_XET enum=Xmtype
    /ch/32/mix/05/panFollow  E32 F_XET
/ch/32/mix/06  <CHME> n=0
    /ch/32/mix/06/on  E32 F_XET enum=OffOn
    /ch/32/mix/06/level  F32 F_XET
/ch/32/mix/07  <CHMO> n=0
    /ch/32/mix/07/on  E32 F_XET enum=OffOn
    /ch/32/mix/07/level  F32 F_XET
    /ch/32/mix/07/pan  F32 F_XET
    /ch/32/mix/07/type  E32 F_XET enum=Xmtype
    /ch/32/mix/07/panFollow  E32 F_XET
/ch/32/mix/08  <CHME> n=0
    /ch/32/mix/08/on  E32 F_XET enum=OffOn
    /ch/32/mix/08/level  F32 F_XET
/ch/32/mix/09  <CHMO> n=0
    /ch/32/mix/09/on  E32 F_XET enum=OffOn
    /ch/32/mix/09/level  F32 F_XET
    /ch/32/mix/09/pan  F32 F_XET
    /ch/32/mix/09/type  E32 F_XET enum=Xmtype
    /ch/32/mix/09/panFollow  E32 F_XET
/ch/32/mix/10  <CHME> n=0
    /ch/32/mix/10/on  E32 F_XET enum=OffOn
    /ch/32/mix/10/level  F32 F_XET
/ch/32/mix/11  <CHMO> n=0
    /ch/32/mix/11/on  E32 F_XET enum=OffOn
    /ch/32/mix/11/level  F32 F_XET
    /ch/32/mix/11/pan  F32 F_XET
    /ch/32/mix/11/type  E32 F_XET enum=Xmtype
    /ch/32/mix/11/panFollow  E32 F_XET
/ch/32/mix/12  <CHME> n=0
    /ch/32/mix/12/on  E32 F_XET enum=OffOn
    /ch/32/mix/12/level  F32 F_XET
/ch/32/mix/13  <CHMO> n=0
    /ch/32/mix/13/on  E32 F_XET enum=OffOn
    /ch/32/mix/13/level  F32 F_XET
    /ch/32/mix/13/pan  F32 F_XET
    /ch/32/mix/13/type  E32 F_XET enum=Xmtype
    /ch/32/mix/13/panFollow  E32 F_XET
/ch/32/mix/14  <CHME> n=0
    /ch/32/mix/14/on  E32 F_XET enum=OffOn
    /ch/32/mix/14/level  F32 F_XET
/ch/32/mix/15  <CHMO> n=0
    /ch/32/mix/15/on  E32 F_XET enum=OffOn
    /ch/32/mix/15/level  F32 F_XET
    /ch/32/mix/15/pan  F32 F_XET
    /ch/32/mix/15/type  E32 F_XET enum=Xmtype
    /ch/32/mix/15/panFollow  E32 F_XET
/ch/32/mix/16  <CHME> n=0
    /ch/32/mix/16/on  E32 F_XET enum=OffOn
    /ch/32/mix/16/level  F32 F_XET
/ch/32/automix  <CHAMIX> n=0
    /ch/32/automix/group  E32 F_XET enum=Xamxgrp
    /ch/32/automix/weight  F32 F_XET
```

### Xconfig (X32CfgMain.h, 335 entries)

```
/config  <OFFON> n=16
/config/chlink  <OFFON> n=16
    /config/chlink/1-2  E32 F_XET enum=OffOn
    /config/chlink/3-4  E32 F_XET enum=OffOn
    /config/chlink/5-6  E32 F_XET enum=OffOn
    /config/chlink/7-8  E32 F_XET enum=OffOn
    /config/chlink/9-10  E32 F_XET enum=OffOn
    /config/chlink/11-12  E32 F_XET enum=OffOn
    /config/chlink/13-14  E32 F_XET enum=OffOn
    /config/chlink/15-16  E32 F_XET enum=OffOn
    /config/chlink/17-18  E32 F_XET enum=OffOn
    /config/chlink/19-20  E32 F_XET enum=OffOn
    /config/chlink/21-22  E32 F_XET enum=OffOn
    /config/chlink/23-24  E32 F_XET enum=OffOn
    /config/chlink/25-26  E32 F_XET enum=OffOn
    /config/chlink/27-28  E32 F_XET enum=OffOn
    /config/chlink/29-30  E32 F_XET enum=OffOn
    /config/chlink/31-32  E32 F_XET enum=OffOn
/config/auxlink  <OFFON> n=4
    /config/auxlink/1-2  E32 F_XET enum=OffOn
    /config/auxlink/3-4  E32 F_XET enum=OffOn
    /config/auxlink/5-6  E32 F_XET enum=OffOn
    /config/auxlink/7-8  E32 F_XET enum=OffOn
/config/fxlink  <OFFON> n=4
    /config/fxlink/1-2  E32 F_XET enum=OffOn
    /config/fxlink/3-4  E32 F_XET enum=OffOn
    /config/fxlink/5-6  E32 F_XET enum=OffOn
    /config/fxlink/7-8  E32 F_XET enum=OffOn
/config/buslink  <OFFON> n=8
    /config/buslink/1-2  E32 F_XET enum=OffOn
    /config/buslink/3-4  E32 F_XET enum=OffOn
    /config/buslink/5-6  E32 F_XET enum=OffOn
    /config/buslink/7-8  E32 F_XET enum=OffOn
    /config/buslink/9-10  E32 F_XET enum=OffOn
    /config/buslink/11-12  E32 F_XET enum=OffOn
    /config/buslink/13-14  E32 F_XET enum=OffOn
    /config/buslink/15-16  E32 F_XET enum=OffOn
/config/mtxlink  <OFFON> n=3
    /config/mtxlink/1-2  E32 F_XET enum=OffOn
    /config/mtxlink/3-4  E32 F_XET enum=OffOn
    /config/mtxlink/5-6  E32 F_XET enum=OffOn
/config/mute  <OFFON> n=6
    /config/mute/1  E32 F_XET enum=OffOn
    /config/mute/2  E32 F_XET enum=OffOn
    /config/mute/3  E32 F_XET enum=OffOn
    /config/mute/4  E32 F_XET enum=OffOn
    /config/mute/5  E32 F_XET enum=OffOn
    /config/mute/6  E32 F_XET enum=OffOn
/config/linkcfg  <OFFON> n=4
    /config/linkcfg/hadly  E32 F_XET enum=OffOn
    /config/linkcfg/eq  E32 F_XET enum=OffOn
    /config/linkcfg/dyn  E32 F_XET enum=OffOn
    /config/linkcfg/fdrmute  E32 F_XET enum=OffOn
/config/mono  <CMONO> n=0
    /config/mono/mode  E32 F_XET enum=Xmnmode
    /config/mono/link  E32 F_XET enum=OffOn
/config/solo  <CSOLO> n=0
    /config/solo/level  F32 F_XET
    /config/solo/source  I32 F_XET enum=XSsourc
    /config/solo/sourcetrim  F32 F_XET
    /config/solo/chmode  E32 F_XET enum=Xchmode
    /config/solo/busmode  E32 F_XET enum=Xchmode
    /config/solo/dcamode  E32 F_XET enum=Xchmode
    /config/solo/exclusive  E32 F_XET enum=OffOn
    /config/solo/followsel  E32 F_XET enum=OffOn
    /config/solo/followsolo  E32 F_XET enum=OffOn
    /config/solo/dimatt  F32 F_XET
    /config/solo/dim  E32 F_XET enum=OffOn
    /config/solo/mono  E32 F_XET enum=OffOn
    /config/solo/delay  E32 F_XET enum=OffOn
    /config/solo/delaytime  F32 F_XET
    /config/solo/masterctrl  E32 F_XET enum=OffOn
    /config/solo/mute  E32 F_XET enum=OffOn
    /config/solo/dimpfl  E32 F_XET enum=OffOn
/config/talk  <CTALK> n=0
    /config/talk/enable  E32 F_XET enum=OffOn
    /config/talk/source  E32 F_XET enum=XTsourc
/config/talk/A  <CTALKAB> n=0
    /config/talk/A/level  F32 F_XET
    /config/talk/A/dim  E32 F_XET enum=OffOn
    /config/talk/A/latch  E32 F_XET enum=OffOn
    /config/talk/A/destmap  P32 F_XET
/config/talk/B  <CTALKAB> n=0
    /config/talk/B/level  F32 F_XET
    /config/talk/B/dim  E32 F_XET enum=OffOn
    /config/talk/B/latch  E32 F_XET enum=OffOn
    /config/talk/B/destmap  P32 F_XET
/config/osc  <COSC> n=0
    /config/osc/level  F32 F_XET
    /config/osc/f1  F32 F_XET
    /config/osc/f2  F32 F_XET
    /config/osc/fsel  E32 F_XET enum=XOscsel
    /config/osc/type  E32 F_XET enum=XOsctyp
    /config/osc/dest  E32 F_XET
/config/userrout  <UROUO> n=0
/config/userrout/out  <UROUO> n=0
    /config/userrout/out/01  I32 F_XET
    /config/userrout/out/02  I32 F_XET
    /config/userrout/out/03  I32 F_XET
    /config/userrout/out/04  I32 F_XET
    /config/userrout/out/05  I32 F_XET
    /config/userrout/out/06  I32 F_XET
    /config/userrout/out/07  I32 F_XET
    /config/userrout/out/08  I32 F_XET
    /config/userrout/out/09  I32 F_XET
    /config/userrout/out/10  I32 F_XET
    /config/userrout/out/11  I32 F_XET
    /config/userrout/out/12  I32 F_XET
    /config/userrout/out/13  I32 F_XET
    /config/userrout/out/14  I32 F_XET
    /config/userrout/out/15  I32 F_XET
    /config/userrout/out/16  I32 F_XET
    /config/userrout/out/17  I32 F_XET
    /config/userrout/out/18  I32 F_XET
    /config/userrout/out/19  I32 F_XET
    /config/userrout/out/20  I32 F_XET
    /config/userrout/out/21  I32 F_XET
    /config/userrout/out/22  I32 F_XET
    /config/userrout/out/23  I32 F_XET
    /config/userrout/out/24  I32 F_XET
    /config/userrout/out/25  I32 F_XET
    /config/userrout/out/26  I32 F_XET
    /config/userrout/out/27  I32 F_XET
    /config/userrout/out/28  I32 F_XET
    /config/userrout/out/29  I32 F_XET
    /config/userrout/out/30  I32 F_XET
    /config/userrout/out/31  I32 F_XET
    /config/userrout/out/32  I32 F_XET
    /config/userrout/out/33  I32 F_XET
    /config/userrout/out/34  I32 F_XET
    /config/userrout/out/35  I32 F_XET
    /config/userrout/out/36  I32 F_XET
    /config/userrout/out/37  I32 F_XET
    /config/userrout/out/38  I32 F_XET
    /config/userrout/out/39  I32 F_XET
    /config/userrout/out/40  I32 F_XET
    /config/userrout/out/41  I32 F_XET
    /config/userrout/out/42  I32 F_XET
    /config/userrout/out/43  I32 F_XET
    /config/userrout/out/44  I32 F_XET
    /config/userrout/out/45  I32 F_XET
    /config/userrout/out/46  I32 F_XET
    /config/userrout/out/47  I32 F_XET
    /config/userrout/out/48  I32 F_XET
/config/userrout/in  <UROUI> n=0
    /config/userrout/in/01  I32 F_XET
    /config/userrout/in/02  I32 F_XET
    /config/userrout/in/03  I32 F_XET
    /config/userrout/in/04  I32 F_XET
    /config/userrout/in/05  I32 F_XET
    /config/userrout/in/06  I32 F_XET
    /config/userrout/in/07  I32 F_XET
    /config/userrout/in/08  I32 F_XET
    /config/userrout/in/09  I32 F_XET
    /config/userrout/in/10  I32 F_XET
    /config/userrout/in/11  I32 F_XET
    /config/userrout/in/12  I32 F_XET
    /config/userrout/in/13  I32 F_XET
    /config/userrout/in/14  I32 F_XET
    /config/userrout/in/15  I32 F_XET
    /config/userrout/in/16  I32 F_XET
    /config/userrout/in/17  I32 F_XET
    /config/userrout/in/18  I32 F_XET
    /config/userrout/in/19  I32 F_XET
    /config/userrout/in/20  I32 F_XET
    /config/userrout/in/21  I32 F_XET
    /config/userrout/in/22  I32 F_XET
    /config/userrout/in/23  I32 F_XET
    /config/userrout/in/24  I32 F_XET
    /config/userrout/in/25  I32 F_XET
    /config/userrout/in/26  I32 F_XET
    /config/userrout/in/27  I32 F_XET
    /config/userrout/in/28  I32 F_XET
    /config/userrout/in/29  I32 F_XET
    /config/userrout/in/30  I32 F_XET
    /config/userrout/in/31  I32 F_XET
    /config/userrout/in/32  I32 F_XET
/config/routing  <CROUTSW> n=1
    /config/routing/routswitch  E32 F_XET enum=XCFrsw
/config/routing/IN  <CROUTIN> n=5
    /config/routing/IN/1-8  E32 F_XET enum=XRtgin
    /config/routing/IN/9-16  E32 F_XET enum=XRtgin
    /config/routing/IN/17-24  E32 F_XET enum=XRtgin
    /config/routing/IN/25-32  E32 F_XET enum=XRtgin
    /config/routing/IN/AUX  E32 F_XET enum=XRtina
/config/routing/AES50A  <CROUTAC> n=6
    /config/routing/AES50A/1-8  E32 F_XET enum=XRtaea
    /config/routing/AES50A/9-16  E32 F_XET enum=XRtaea
    /config/routing/AES50A/17-24  E32 F_XET enum=XRtaea
    /config/routing/AES50A/25-32  E32 F_XET enum=XRtaea
    /config/routing/AES50A/33-40  E32 F_XET enum=XRtaea
    /config/routing/AES50A/41-48  E32 F_XET enum=XRtaea
/config/routing/AES50B  <CROUTAC> n=6
    /config/routing/AES50B/1-8  E32 F_XET enum=XRtaea
    /config/routing/AES50B/9-16  E32 F_XET enum=XRtaea
    /config/routing/AES50B/17-24  E32 F_XET enum=XRtaea
    /config/routing/AES50B/25-32  E32 F_XET enum=XRtaea
    /config/routing/AES50B/33-40  E32 F_XET enum=XRtaea
    /config/routing/AES50B/41-48  E32 F_XET enum=XRtaea
/config/routing/CARD  <CROUTAC> n=4
    /config/routing/CARD/1-8  E32 F_XET enum=XRtaea
    /config/routing/CARD/9-16  E32 F_XET enum=XRtaea
    /config/routing/CARD/17-24  E32 F_XET enum=XRtaea
    /config/routing/CARD/25-32  E32 F_XET enum=XRtaea
/config/routing/OUT  <CROUTOT> n=0
    /config/routing/OUT/1-4  E32 F_XET enum=XRout1
    /config/routing/OUT/9-12  E32 F_XET enum=XRout1
    /config/routing/OUT/5-8  E32 F_XET enum=XRout5
    /config/routing/OUT/13-16  E32 F_XET enum=XRout5
/config/routing/PLAY  <CROUTIN> n=0
    /config/routing/PLAY/1-8  E32 F_XET enum=XRtgin
    /config/routing/PLAY/9-16  E32 F_XET enum=XRtgin
    /config/routing/PLAY/17-24  E32 F_XET enum=XRtgin
    /config/routing/PLAY/25-32  E32 F_XET enum=XRtgin
    /config/routing/PLAY/AUX  E32 F_XET enum=XRtina
/config/userctrl/A  <CCTRL> n=0
    /config/userctrl/A/color  E32 F_XET enum=Xcolors
/config/userctrl/A/enc  <CENC> n=4
    /config/userctrl/A/enc/1  S32 F_XET
    /config/userctrl/A/enc/2  S32 F_XET
    /config/userctrl/A/enc/3  S32 F_XET
    /config/userctrl/A/enc/4  S32 F_XET
/config/userctrl/A/btn  <CENC> n=8
    /config/userctrl/A/btn/5  S32 F_XET
    /config/userctrl/A/btn/6  S32 F_XET
    /config/userctrl/A/btn/7  S32 F_XET
    /config/userctrl/A/btn/8  S32 F_XET
    /config/userctrl/A/btn/9  S32 F_XET
    /config/userctrl/A/btn/10  S32 F_XET
    /config/userctrl/A/btn/11  S32 F_XET
    /config/userctrl/A/btn/12  S32 F_XET
/config/userctrl/B  <CCTRL> n=0
    /config/userctrl/B/color  E32 F_XET enum=Xcolors
/config/userctrl/B/enc  <CENC> n=4
    /config/userctrl/B/enc/1  S32 F_XET
    /config/userctrl/B/enc/2  S32 F_XET
    /config/userctrl/B/enc/3  S32 F_XET
    /config/userctrl/B/enc/4  S32 F_XET
/config/userctrl/B/btn  <CENC> n=8
    /config/userctrl/B/btn/5  S32 F_XET
    /config/userctrl/B/btn/6  S32 F_XET
    /config/userctrl/B/btn/7  S32 F_XET
    /config/userctrl/B/btn/8  S32 F_XET
    /config/userctrl/B/btn/9  S32 F_XET
    /config/userctrl/B/btn/10  S32 F_XET
    /config/userctrl/B/btn/11  S32 F_XET
    /config/userctrl/B/btn/12  S32 F_XET
/config/userctrl/C  <CCTRL> n=0
    /config/userctrl/C/color  E32 F_XET enum=Xcolors
/config/userctrl/C/enc  <CENC> n=4
    /config/userctrl/C/enc/1  S32 F_XET
    /config/userctrl/C/enc/2  S32 F_XET
    /config/userctrl/C/enc/3  S32 F_XET
    /config/userctrl/C/enc/4  S32 F_XET
/config/userctrl/C/btn  <CENC> n=8
    /config/userctrl/C/btn/5  S32 F_XET
    /config/userctrl/C/btn/6  S32 F_XET
    /config/userctrl/C/btn/7  S32 F_XET
    /config/userctrl/C/btn/8  S32 F_XET
    /config/userctrl/C/btn/9  S32 F_XET
    /config/userctrl/C/btn/10  S32 F_XET
    /config/userctrl/C/btn/11  S32 F_XET
    /config/userctrl/C/btn/12  S32 F_XET
/config/tape  <CTAPE> n=0
    /config/tape/gainL  F32 F_XET
    /config/tape/gainR  F32 F_XET
    /config/tape/autoplay  E32 F_XET enum=OffOn
/config/amixenable  <CMIX> n=0
    /config/amixenable/X  E32 F_XET enum=OffOn
    /config/amixenable/Y  E32 F_XET enum=OffOn
/config/dp48  <D48> n=0
    /config/dp48/scope  I32 F_XET
    /config/dp48/broadcast  I32 F_XET
/config/dp48/assign  <D48A> n=0
    /config/dp48/assign/01  I32 F_XET
    /config/dp48/assign/02  I32 F_XET
    /config/dp48/assign/03  I32 F_XET
    /config/dp48/assign/04  I32 F_XET
    /config/dp48/assign/05  I32 F_XET
    /config/dp48/assign/06  I32 F_XET
    /config/dp48/assign/07  I32 F_XET
    /config/dp48/assign/08  I32 F_XET
    /config/dp48/assign/09  I32 F_XET
    /config/dp48/assign/10  I32 F_XET
    /config/dp48/assign/11  I32 F_XET
    /config/dp48/assign/12  I32 F_XET
    /config/dp48/assign/13  I32 F_XET
    /config/dp48/assign/14  I32 F_XET
    /config/dp48/assign/15  I32 F_XET
    /config/dp48/assign/16  I32 F_XET
    /config/dp48/assign/17  I32 F_XET
    /config/dp48/assign/18  I32 F_XET
    /config/dp48/assign/19  I32 F_XET
    /config/dp48/assign/20  I32 F_XET
    /config/dp48/assign/21  I32 F_XET
    /config/dp48/assign/22  I32 F_XET
    /config/dp48/assign/23  I32 F_XET
    /config/dp48/assign/24  I32 F_XET
    /config/dp48/assign/25  I32 F_XET
    /config/dp48/assign/26  I32 F_XET
    /config/dp48/assign/27  I32 F_XET
    /config/dp48/assign/28  I32 F_XET
    /config/dp48/assign/29  I32 F_XET
    /config/dp48/assign/30  I32 F_XET
    /config/dp48/assign/31  I32 F_XET
    /config/dp48/assign/32  I32 F_XET
    /config/dp48/assign/33  I32 F_XET
    /config/dp48/assign/34  I32 F_XET
    /config/dp48/assign/35  I32 F_XET
    /config/dp48/assign/36  I32 F_XET
    /config/dp48/assign/37  I32 F_XET
    /config/dp48/assign/38  I32 F_XET
    /config/dp48/assign/39  I32 F_XET
    /config/dp48/assign/40  I32 F_XET
    /config/dp48/assign/41  I32 F_XET
    /config/dp48/assign/42  I32 F_XET
    /config/dp48/assign/43  I32 F_XET
    /config/dp48/assign/44  I32 F_XET
    /config/dp48/assign/45  I32 F_XET
    /config/dp48/assign/46  I32 F_XET
    /config/dp48/assign/47  I32 F_XET
    /config/dp48/assign/48  I32 F_XET
/config/dp48/grpname  <D48G> n=0
    /config/dp48/grpname/01  S32 F_XET
    /config/dp48/grpname/02  S32 F_XET
    /config/dp48/grpname/03  S32 F_XET
    /config/dp48/grpname/04  S32 F_XET
    /config/dp48/grpname/05  S32 F_XET
    /config/dp48/grpname/06  S32 F_XET
    /config/dp48/grpname/07  S32 F_XET
    /config/dp48/grpname/08  S32 F_XET
    /config/dp48/grpname/09  S32 F_XET
    /config/dp48/grpname/10  S32 F_XET
    /config/dp48/grpname/11  S32 F_XET
    /config/dp48/grpname/12  S32 F_XET
```

### Xmain (X32CfgMain.h, 182 entries)

```
/main  <BSCO> n=0
/main/st  <BSCO> n=0
/main/st/config  <BSCO> n=0
    /main/st/config/name  S32 F_XET
    /main/st/config/icon  I32 F_XET
    /main/st/config/color  E32 F_XET enum=Xcolors
/main/st/dyn  <MXDY> n=0
    /main/st/dyn/on  E32 F_XET enum=OffOn
    /main/st/dyn/mode  E32 F_XET enum=Xdymode
    /main/st/dyn/det  E32 F_XET enum=Xdydet
    /main/st/dyn/env  E32 F_XET enum=Xdyenv
    /main/st/dyn/thr  F32 F_XET
    /main/st/dyn/ratio  E32 F_XET enum=Xdyrat
    /main/st/dyn/knee  F32 F_XET
    /main/st/dyn/mgain  F32 F_XET
    /main/st/dyn/attack  F32 F_XET
    /main/st/dyn/hold  F32 F_XET
    /main/st/dyn/release  F32 F_XET
    /main/st/dyn/pos  E32 F_XET enum=Xdyppos
    /main/st/dyn/mix  F32 F_XET
    /main/st/dyn/auto  E32 F_XET enum=OffOn
/main/st/dyn/filter  <CHDF> n=0
    /main/st/dyn/filter/on  E32 F_XET enum=OffOn
    /main/st/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /main/st/dyn/filter/f  F32 F_XET
/main/st/insert  <CHIN> n=0
    /main/st/insert/on  E32 F_XET enum=OffOn
    /main/st/insert/pos  E32 F_XET enum=Xdyppos
    /main/st/insert/sel  E32 F_XET enum=Xisel
/main/st/eq  <OFFON> n=1
    /main/st/eq/on  E32 F_XET enum=OffOn
/main/st/eq/1  <CHEQ> n=0
    /main/st/eq/1/type  E32 F_XET enum=Xeqty2
    /main/st/eq/1/f  F32 F_XET
    /main/st/eq/1/g  F32 F_XET
    /main/st/eq/1/q  F32 F_XET
/main/st/eq/2  <CHEQ> n=0
    /main/st/eq/2/type  E32 F_XET enum=Xeqty2
    /main/st/eq/2/f  F32 F_XET
    /main/st/eq/2/g  F32 F_XET
    /main/st/eq/2/q  F32 F_XET
/main/st/eq/3  <CHEQ> n=0
    /main/st/eq/3/type  E32 F_XET enum=Xeqty2
    /main/st/eq/3/f  F32 F_XET
    /main/st/eq/3/g  F32 F_XET
    /main/st/eq/3/q  F32 F_XET
/main/st/eq/4  <CHEQ> n=0
    /main/st/eq/4/type  E32 F_XET enum=Xeqty2
    /main/st/eq/4/f  F32 F_XET
    /main/st/eq/4/g  F32 F_XET
    /main/st/eq/4/q  F32 F_XET
/main/st/eq/5  <CHEQ> n=0
    /main/st/eq/5/type  E32 F_XET enum=Xeqty2
    /main/st/eq/5/f  F32 F_XET
    /main/st/eq/5/g  F32 F_XET
    /main/st/eq/5/q  F32 F_XET
/main/st/eq/6  <CHEQ> n=0
    /main/st/eq/6/type  E32 F_XET enum=Xeqty2
    /main/st/eq/6/f  F32 F_XET
    /main/st/eq/6/g  F32 F_XET
    /main/st/eq/6/q  F32 F_XET
/main/st/mix  <MSMX> n=0
    /main/st/mix/on  E32 F_XET enum=OffOn
    /main/st/mix/fader  F32 F_XET
    /main/st/mix/pan  F32 F_XET
/main/st/mix/01  <CHMO> n=0
    /main/st/mix/01/on  E32 F_XET enum=OffOn
    /main/st/mix/01/level  F32 F_XET
    /main/st/mix/01/pan  F32 F_XET
    /main/st/mix/01/type  E32 F_XET enum=Xmtype
/main/st/mix/02  <CHME> n=0
    /main/st/mix/02/on  E32 F_XET enum=OffOn
    /main/st/mix/02/level  F32 F_XET
/main/st/mix/03  <CHMO> n=0
    /main/st/mix/03/on  E32 F_XET enum=OffOn
    /main/st/mix/03/level  F32 F_XET
    /main/st/mix/03/pan  F32 F_XET
    /main/st/mix/03/type  E32 F_XET enum=Xmtype
/main/st/mix/04  <CHME> n=0
    /main/st/mix/04/on  E32 F_XET enum=OffOn
    /main/st/mix/04/level  F32 F_XET
/main/st/mix/05  <CHMO> n=0
    /main/st/mix/05/on  E32 F_XET enum=OffOn
    /main/st/mix/05/level  F32 F_XET
    /main/st/mix/05/pan  F32 F_XET
    /main/st/mix/05/type  E32 F_XET enum=Xmtype
/main/st/mix/06  <CHME> n=0
    /main/st/mix/06/on  E32 F_XET enum=OffOn
    /main/st/mix/06/level  F32 F_XET
/main/st/grp  <CHGRP> n=0
    /main/st/grp/dca  P32 F_XET
    /main/st/grp/mute  P32 F_XET
/main/m  <BSCO> n=0
/main/m/config  <BSCO> n=0
    /main/m/config/name  S32 F_XET
    /main/m/config/icon  I32 F_XET
    /main/m/config/color  E32 F_XET enum=Xcolors
/main/m/dyn  <MXDY> n=0
    /main/m/dyn/on  E32 F_XET enum=OffOn
    /main/m/dyn/mode  E32 F_XET enum=Xdymode
    /main/m/dyn/det  E32 F_XET enum=Xdydet
    /main/m/dyn/env  E32 F_XET enum=Xdyenv
    /main/m/dyn/thr  F32 F_XET
    /main/m/dyn/ratio  E32 F_XET enum=Xdyrat
    /main/m/dyn/knee  F32 F_XET
    /main/m/dyn/mgain  F32 F_XET
    /main/m/dyn/attack  F32 F_XET
    /main/m/dyn/hold  F32 F_XET
    /main/m/dyn/release  F32 F_XET
    /main/m/dyn/pos  E32 F_XET enum=Xdyppos
    /main/m/dyn/mix  F32 F_XET
    /main/m/dyn/auto  E32 F_XET enum=OffOn
/main/m/dyn/filter  <CHDF> n=0
    /main/m/dyn/filter/on  E32 F_XET enum=OffOn
    /main/m/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /main/m/dyn/filter/f  F32 F_XET
/main/m/insert  <CHIN> n=0
    /main/m/insert/on  E32 F_XET enum=OffOn
    /main/m/insert/pos  E32 F_XET enum=Xdyppos
    /main/m/insert/sel  E32 F_XET enum=Xisel
/main/m/eq  <OFFON> n=1
    /main/m/eq/on  E32 F_XET enum=OffOn
/main/m/eq/1  <CHEQ> n=0
    /main/m/eq/1/type  E32 F_XET enum=Xeqty2
    /main/m/eq/1/f  F32 F_XET
    /main/m/eq/1/g  F32 F_XET
    /main/m/eq/1/q  F32 F_XET
/main/m/eq/2  <CHEQ> n=0
    /main/m/eq/2/type  E32 F_XET enum=Xeqty2
    /main/m/eq/2/f  F32 F_XET
    /main/m/eq/2/g  F32 F_XET
    /main/m/eq/2/q  F32 F_XET
/main/m/eq/3  <CHEQ> n=0
    /main/m/eq/3/type  E32 F_XET enum=Xeqty2
    /main/m/eq/3/f  F32 F_XET
    /main/m/eq/3/g  F32 F_XET
    /main/m/eq/3/q  F32 F_XET
/main/m/eq/4  <CHEQ> n=0
    /main/m/eq/4/type  E32 F_XET enum=Xeqty2
    /main/m/eq/4/f  F32 F_XET
    /main/m/eq/4/g  F32 F_XET
    /main/m/eq/4/q  F32 F_XET
/main/m/eq/5  <CHEQ> n=0
    /main/m/eq/5/type  E32 F_XET enum=Xeqty2
    /main/m/eq/5/f  F32 F_XET
    /main/m/eq/5/g  F32 F_XET
    /main/m/eq/5/q  F32 F_XET
/main/m/eq/6  <CHEQ> n=0
    /main/m/eq/6/type  E32 F_XET enum=Xeqty2
    /main/m/eq/6/f  F32 F_XET
    /main/m/eq/6/g  F32 F_XET
    /main/m/eq/6/q  F32 F_XET
/main/m/mix  <CHME> n=0
    /main/m/mix/on  E32 F_XET enum=OffOn
    /main/m/mix/fader  F32 F_XET
/main/m/mix/01  <CHMO> n=0
    /main/m/mix/01/on  E32 F_XET enum=OffOn
    /main/m/mix/01/level  F32 F_XET
    /main/m/mix/01/pan  F32 F_XET
    /main/m/mix/01/type  E32 F_XET enum=Xmtype
/main/m/mix/02  <CHME> n=0
    /main/m/mix/02/on  E32 F_XET enum=OffOn
    /main/m/mix/02/level  F32 F_XET
/main/m/mix/03  <CHMO> n=0
    /main/m/mix/03/on  E32 F_XET enum=OffOn
    /main/m/mix/03/level  F32 F_XET
    /main/m/mix/03/pan  F32 F_XET
    /main/m/mix/03/type  E32 F_XET enum=Xmtype
/main/m/mix/04  <CHME> n=0
    /main/m/mix/04/on  E32 F_XET enum=OffOn
    /main/m/mix/04/level  F32 F_XET
/main/m/mix/05  <CHMO> n=0
    /main/m/mix/05/on  E32 F_XET enum=OffOn
    /main/m/mix/05/level  F32 F_XET
    /main/m/mix/05/pan  F32 F_XET
    /main/m/mix/05/type  E32 F_XET enum=Xmtype
/main/m/mix/06  <CHME> n=0
    /main/m/mix/06/on  E32 F_XET enum=OffOn
    /main/m/mix/06/level  F32 F_XET
/main/m/grp  <CHGRP> n=0
    /main/m/grp/dca  P32 F_XET
    /main/m/grp/mute  P32 F_XET
```

### Xprefs (X32PrefStat.h, 239 entries)

```
/-prefs  <PREFS> n=0
    /-prefs/style  S32 F_XET
    /-prefs/bright  F32 F_XET
    /-prefs/lcdcont  F32 F_XET
    /-prefs/ledbright  F32 F_XET
    /-prefs/lamp  F32 F_XET
    /-prefs/lampon  E32 F_XET enum=OffOn
    /-prefs/clockrate  E32 F_XET enum=PRrate
    /-prefs/clocksource  E32 F_XET enum=Psource
    /-prefs/confirm_general  E32 F_XET enum=OffOn
    /-prefs/confirm_overwrite  E32 F_XET enum=OffOn
    /-prefs/confirm_sceneload  E32 F_XET enum=OffOn
    /-prefs/viewrtn  E32 F_XET enum=OffOn
    /-prefs/selfollowsbank  E32 F_XET enum=OffOn
    /-prefs/scene_advance  E32 F_XET enum=OffOn
    /-prefs/safe_masterlevels  E32 F_XET enum=OffOn
    /-prefs/haflags  P32 F_XET
    /-prefs/autosel  E32 F_XET enum=OffOn
    /-prefs/show_control  E32 F_XET enum=PSCont
    /-prefs/clockmode  E32 F_XET enum=Pclkmod
    /-prefs/hardmute  E32 F_XET enum=OffOn
    /-prefs/dcamute  E32 F_XET enum=OffOn
    /-prefs/invertmutes  E32 F_XET enum=Pinvmut
    /-prefs/name  S32 F_XET
    /-prefs/rec_control  E32 F_XET enum=Purrctl
/-prefs/remote  <PIR> n=0
    /-prefs/remote/enable  E32 F_XET enum=OffOn
    /-prefs/remote/protocol  E32 F_XET enum=PRpro
    /-prefs/remote/port  E32 F_XET enum=PRport
    /-prefs/remote/ioenable  P32 F_XET
/-prefs/iQ  <PIQ> n=0
/-prefs/iQ/01  <PIQ> n=0
    /-prefs/iQ/01/iQmodel  E32 F_XET enum=XiQspk
    /-prefs/iQ/01/IQeqset  E32 F_XET enum=XiQeq
    /-prefs/iQ/01/IQsound  I32 F_XET
/-prefs/iQ/02  <PIQ> n=0
    /-prefs/iQ/02/iQmodel  E32 F_XET enum=XiQspk
    /-prefs/iQ/02/IQeqset  E32 F_XET enum=XiQeq
    /-prefs/iQ/02/IQsound  I32 F_XET
/-prefs/iQ/03  <PIQ> n=0
    /-prefs/iQ/03/iQmodel  E32 F_XET enum=XiQspk
    /-prefs/iQ/03/IQeqset  E32 F_XET enum=XiQeq
    /-prefs/iQ/03/IQsound  I32 F_XET
/-prefs/iQ/04  <PIQ> n=0
    /-prefs/iQ/04/iQmodel  E32 F_XET enum=XiQspk
    /-prefs/iQ/04/IQeqset  E32 F_XET enum=XiQeq
    /-prefs/iQ/04/IQsound  I32 F_XET
/-prefs/iQ/05  <PIQ> n=0
    /-prefs/iQ/05/iQmodel  E32 F_XET enum=XiQspk
    /-prefs/iQ/05/IQeqset  E32 F_XET enum=XiQeq
    /-prefs/iQ/05/IQsound  I32 F_XET
/-prefs/iQ/06  <PIQ> n=0
    /-prefs/iQ/06/iQmodel  E32 F_XET enum=XiQspk
    /-prefs/iQ/06/IQeqset  E32 F_XET enum=XiQeq
    /-prefs/iQ/06/IQsound  I32 F_XET
/-prefs/iQ/07  <PIQ> n=0
    /-prefs/iQ/07/iQmodel  E32 F_XET enum=XiQspk
    /-prefs/iQ/07/IQeqset  E32 F_XET enum=XiQeq
    /-prefs/iQ/07/IQsound  I32 F_XET
/-prefs/iQ/08  <PIQ> n=0
    /-prefs/iQ/08/iQmodel  E32 F_XET enum=XiQspk
    /-prefs/iQ/08/IQeqset  E32 F_XET enum=XiQeq
    /-prefs/iQ/08/IQsound  I32 F_XET
/-prefs/iQ/09  <PIQ> n=0
    /-prefs/iQ/09/iQmodel  E32 F_XET enum=XiQspk
    /-prefs/iQ/09/IQeqset  E32 F_XET enum=XiQeq
    /-prefs/iQ/09/IQsound  I32 F_XET
/-prefs/iQ/10  <PIQ> n=0
    /-prefs/iQ/10/iQmodel  E32 F_XET enum=XiQspk
    /-prefs/iQ/10/IQeqset  E32 F_XET enum=XiQeq
    /-prefs/iQ/10/IQsound  I32 F_XET
/-prefs/iQ/11  <PIQ> n=0
    /-prefs/iQ/11/iQmodel  E32 F_XET enum=XiQspk
    /-prefs/iQ/11/IQeqset  E32 F_XET enum=XiQeq
    /-prefs/iQ/11/IQsound  I32 F_XET
/-prefs/iQ/12  <PIQ> n=0
    /-prefs/iQ/12/iQmodel  E32 F_XET enum=XiQspk
    /-prefs/iQ/12/IQeqset  E32 F_XET enum=XiQeq
    /-prefs/iQ/12/IQsound  I32 F_XET
/-prefs/iQ/13  <PIQ> n=0
    /-prefs/iQ/13/iQmodel  E32 F_XET enum=XiQspk
    /-prefs/iQ/13/IQeqset  E32 F_XET enum=XiQeq
    /-prefs/iQ/13/IQsound  I32 F_XET
/-prefs/iQ/14  <PIQ> n=0
    /-prefs/iQ/14/iQmodel  E32 F_XET enum=XiQspk
    /-prefs/iQ/14/IQeqset  E32 F_XET enum=XiQeq
    /-prefs/iQ/14/IQsound  I32 F_XET
/-prefs/iQ/15  <PIQ> n=0
    /-prefs/iQ/15/iQmodel  E32 F_XET enum=XiQspk
    /-prefs/iQ/15/IQeqset  E32 F_XET enum=XiQeq
    /-prefs/iQ/15/IQsound  I32 F_XET
/-prefs/iQ/16  <PIQ> n=0
    /-prefs/iQ/16/iQmodel  E32 F_XET enum=XiQspk
    /-prefs/iQ/16/IQeqset  E32 F_XET enum=XiQeq
    /-prefs/iQ/16/IQsound  I32 F_XET
/-prefs/card  <PCARD> n=0
    /-prefs/card/UFifc  E32 F_XET enum=Pctype
    /-prefs/card/UFmode  E32 F_XET enum=Pufmode
    /-prefs/card/USBmode  E32 F_XET enum=Pusbmod
    /-prefs/card/ADATwc  E32 F_XET enum=Pcaw
    /-prefs/card/ADATsync  E32 F_XET enum=Pcas
    /-prefs/card/MADImode  E32 F_XET enum=Pmdmode
    /-prefs/card/MADIin  E32 F_XET enum=Pcmadi
    /-prefs/card/MADIout  E32 F_XET enum=Pcmado
    /-prefs/card/MADIsrc  E32 F_XET enum=Pmadsrc
    /-prefs/card/URECtracks  E32 F_XET enum=Purectk
    /-prefs/card/URECplayb  E32 F_XET enum=Purplbk
    /-prefs/card/URECrout  E32 F_XET enum=Purerpa
    /-prefs/card/URECsdsel  E32 F_XET enum=Pursdsl
/-prefs/rta  <PRTA> n=0
    /-prefs/rta/visibility  E32 F_XET enum=Prtavis
    /-prefs/rta/gain  F32 F_XET
    /-prefs/rta/autogain  E32 F_XET enum=OffOn
    /-prefs/rta/source  I32 F_XET
    /-prefs/rta/pos  E32 F_XET enum=PRpos
    /-prefs/rta/mode  E32 F_XET enum=PRmode
    /-prefs/rta/options  P32 F_XET
    /-prefs/rta/det  E32 F_XET enum=PRdet
    /-prefs/rta/decay  F32 F_XET
    /-prefs/rta/peakhold  E32 F_XET enum=Prtaph
/-prefs/ip  <PIP> n=0
    /-prefs/ip/dhcp  E32 F_XET enum=OffOn
/-prefs/ip/addr  <PADDR> n=0
    /-prefs/ip/addr/0  I32 F_XET
    /-prefs/ip/addr/1  I32 F_XET
    /-prefs/ip/addr/2  I32 F_XET
    /-prefs/ip/addr/3  I32 F_XET
/-prefs/ip/mask  <PADDR> n=0
    /-prefs/ip/mask/0  I32 F_XET
    /-prefs/ip/mask/1  I32 F_XET
    /-prefs/ip/mask/2  I32 F_XET
    /-prefs/ip/mask/3  I32 F_XET
/-prefs/ip/gateway  <PADDR> n=0
    /-prefs/ip/gateway/0  I32 F_XET
    /-prefs/ip/gateway/1  I32 F_XET
    /-prefs/ip/gateway/2  I32 F_XET
    /-prefs/ip/gateway/3  I32 F_XET
/-prefs/key  <PKEY> n=0
    /-prefs/key/layout  E32 F_XET
    /-prefs/key/00  S32 F_XET
    /-prefs/key/01  S32 F_XET
    /-prefs/key/02  S32 F_XET
    /-prefs/key/03  S32 F_XET
    /-prefs/key/04  S32 F_XET
    /-prefs/key/05  S32 F_XET
    /-prefs/key/06  S32 F_XET
    /-prefs/key/07  S32 F_XET
    /-prefs/key/08  S32 F_XET
    /-prefs/key/09  S32 F_XET
    /-prefs/key/10  S32 F_XET
    /-prefs/key/11  S32 F_XET
    /-prefs/key/12  S32 F_XET
    /-prefs/key/13  S32 F_XET
    /-prefs/key/14  S32 F_XET
    /-prefs/key/15  S32 F_XET
    /-prefs/key/16  S32 F_XET
    /-prefs/key/17  S32 F_XET
    /-prefs/key/18  S32 F_XET
    /-prefs/key/19  S32 F_XET
    /-prefs/key/20  S32 F_XET
    /-prefs/key/21  S32 F_XET
    /-prefs/key/22  S32 F_XET
    /-prefs/key/23  S32 F_XET
    /-prefs/key/24  S32 F_XET
    /-prefs/key/25  S32 F_XET
    /-prefs/key/26  S32 F_XET
    /-prefs/key/27  S32 F_XET
    /-prefs/key/28  S32 F_XET
    /-prefs/key/29  S32 F_XET
    /-prefs/key/30  S32 F_XET
    /-prefs/key/31  S32 F_XET
    /-prefs/key/32  S32 F_XET
    /-prefs/key/33  S32 F_XET
    /-prefs/key/34  S32 F_XET
    /-prefs/key/35  S32 F_XET
    /-prefs/key/36  S32 F_XET
    /-prefs/key/37  S32 F_XET
    /-prefs/key/38  S32 F_XET
    /-prefs/key/39  S32 F_XET
    /-prefs/key/40  S32 F_XET
    /-prefs/key/41  S32 F_XET
    /-prefs/key/42  S32 F_XET
    /-prefs/key/43  S32 F_XET
    /-prefs/key/44  S32 F_XET
    /-prefs/key/45  S32 F_XET
    /-prefs/key/46  S32 F_XET
    /-prefs/key/47  S32 F_XET
    /-prefs/key/48  S32 F_XET
    /-prefs/key/49  S32 F_XET
    /-prefs/key/50  S32 F_XET
    /-prefs/key/51  S32 F_XET
    /-prefs/key/52  S32 F_XET
    /-prefs/key/53  S32 F_XET
    /-prefs/key/54  S32 F_XET
    /-prefs/key/55  S32 F_XET
    /-prefs/key/56  S32 F_XET
    /-prefs/key/57  S32 F_XET
    /-prefs/key/58  S32 F_XET
    /-prefs/key/59  S32 F_XET
    /-prefs/key/60  S32 F_XET
    /-prefs/key/61  S32 F_XET
    /-prefs/key/62  S32 F_XET
    /-prefs/key/63  S32 F_XET
    /-prefs/key/64  S32 F_XET
    /-prefs/key/65  S32 F_XET
    /-prefs/key/66  S32 F_XET
    /-prefs/key/67  S32 F_XET
    /-prefs/key/68  S32 F_XET
    /-prefs/key/69  S32 F_XET
    /-prefs/key/70  S32 F_XET
    /-prefs/key/71  S32 F_XET
    /-prefs/key/72  S32 F_XET
    /-prefs/key/73  S32 F_XET
    /-prefs/key/74  S32 F_XET
    /-prefs/key/75  S32 F_XET
    /-prefs/key/76  S32 F_XET
    /-prefs/key/77  S32 F_XET
    /-prefs/key/78  S32 F_XET
    /-prefs/key/79  S32 F_XET
    /-prefs/key/80  S32 F_XET
    /-prefs/key/81  S32 F_XET
    /-prefs/key/82  S32 F_XET
    /-prefs/key/83  S32 F_XET
    /-prefs/key/84  S32 F_XET
    /-prefs/key/85  S32 F_XET
    /-prefs/key/86  S32 F_XET
    /-prefs/key/87  S32 F_XET
    /-prefs/key/88  S32 F_XET
    /-prefs/key/89  S32 F_XET
    /-prefs/key/90  S32 F_XET
    /-prefs/key/91  S32 F_XET
    /-prefs/key/92  S32 F_XET
    /-prefs/key/93  S32 F_XET
    /-prefs/key/94  S32 F_XET
    /-prefs/key/95  S32 F_XET
    /-prefs/key/96  S32 F_XET
    /-prefs/key/97  S32 F_XET
    /-prefs/key/98  S32 F_XET
    /-prefs/key/99  S32 F_XET
```

### Xstat (X32PrefStat.h, 147 entries)

```
/-stat  <STAT> n=0
    /-stat/selidx  E32 F_XET enum=Sselidx
    /-stat/chfaderbank  I32 F_XET
    /-stat/grpfaderbank  I32 F_XET
    /-stat/sendsonfader  E32 F_XET enum=OffOn
    /-stat/bussendbank  I32 F_XET
    /-stat/eqband  I32 F_XET
    /-stat/solo  E32 F_XET enum=OffOn
    /-stat/keysolo  E32 F_XET enum=OffOn
    /-stat/userbank  I32 F_XET
    /-stat/autosave  E32 F_XET enum=OffOn
    /-stat/lock  I32 F_XET
    /-stat/usbmounted  E32 F_XET enum=OffOn
    /-stat/remote  E32 F_XET enum=OffOn
    /-stat/rtamodeeq  E32 F_XET enum=PRmode
    /-stat/rtamodegeq  E32 F_XET enum=PRmode
    /-stat/rtaeqpre  E32 F_XET enum=OffOn
    /-stat/rtageqpost  E32 F_XET enum=OffOn
    /-stat/rtasource  I32 F_XET
    /-stat/xcardtype  I32 F_XET
    /-stat/xcardsync  E32 F_XET enum=OffOn
    /-stat/geqonfdr  E32 F_XET enum=OffOn
    /-stat/geqpos  I32 F_XET
    /-stat/dcaspill  I32 F_XET
/-stat/screen  <SSCREEN> n=0
    /-stat/screen/screen  E32 F_XET enum=Sscrn
    /-stat/screen/mutegrp  E32 F_XET enum=OffOn
    /-stat/screen/utils  E32 F_XET enum=OffOn
/-stat/screen/CHAN  <SCHA> n=0
    /-stat/screen/CHAN/page  E32 F_XET enum=Schal
/-stat/screen/METER  <SMET> n=0
    /-stat/screen/METER/page  E32 F_XET enum=Smetl
/-stat/screen/ROUTE  <SROU> n=0
    /-stat/screen/ROUTE/page  E32 F_XET enum=Sroul
/-stat/screen/SETUP  <SSET> n=0
    /-stat/screen/SETUP/page  E32 F_XET enum=Ssetl
/-stat/screen/LIB  <SLIB> n=0
    /-stat/screen/LIB/page  E32 F_XET enum=Slibl
/-stat/screen/FX  <SFX> n=0
    /-stat/screen/FX/page  E32 F_XET enum=Sfxl
/-stat/screen/MON  <SMON> n=0
    /-stat/screen/MON/page  E32 F_XET enum=Smonl
/-stat/screen/USB  <SUSB> n=0
    /-stat/screen/USB/page  E32 F_XET enum=Susbl
/-stat/screen/SCENE  <SSCE> n=0
    /-stat/screen/SCENE/page  E32 F_XET enum=Sscel
/-stat/screen/ASSIGN  <SASS> n=0
    /-stat/screen/ASSIGN/page  E32 F_XET enum=Sassl
/-stat/solosw  <SSOLOSW> n=80
    /-stat/solosw/01  E32 F_XET enum=OffOn
    /-stat/solosw/02  E32 F_XET enum=OffOn
    /-stat/solosw/03  E32 F_XET enum=OffOn
    /-stat/solosw/04  E32 F_XET enum=OffOn
    /-stat/solosw/05  E32 F_XET enum=OffOn
    /-stat/solosw/06  E32 F_XET enum=OffOn
    /-stat/solosw/07  E32 F_XET enum=OffOn
    /-stat/solosw/08  E32 F_XET enum=OffOn
    /-stat/solosw/09  E32 F_XET enum=OffOn
    /-stat/solosw/10  E32 F_XET enum=OffOn
    /-stat/solosw/11  E32 F_XET enum=OffOn
    /-stat/solosw/12  E32 F_XET enum=OffOn
    /-stat/solosw/13  E32 F_XET enum=OffOn
    /-stat/solosw/14  E32 F_XET enum=OffOn
    /-stat/solosw/15  E32 F_XET enum=OffOn
    /-stat/solosw/16  E32 F_XET enum=OffOn
    /-stat/solosw/17  E32 F_XET enum=OffOn
    /-stat/solosw/18  E32 F_XET enum=OffOn
    /-stat/solosw/19  E32 F_XET enum=OffOn
    /-stat/solosw/20  E32 F_XET enum=OffOn
    /-stat/solosw/21  E32 F_XET enum=OffOn
    /-stat/solosw/22  E32 F_XET enum=OffOn
    /-stat/solosw/23  E32 F_XET enum=OffOn
    /-stat/solosw/24  E32 F_XET enum=OffOn
    /-stat/solosw/25  E32 F_XET enum=OffOn
    /-stat/solosw/26  E32 F_XET enum=OffOn
    /-stat/solosw/27  E32 F_XET enum=OffOn
    /-stat/solosw/28  E32 F_XET enum=OffOn
    /-stat/solosw/29  E32 F_XET enum=OffOn
    /-stat/solosw/30  E32 F_XET enum=OffOn
    /-stat/solosw/31  E32 F_XET enum=OffOn
    /-stat/solosw/32  E32 F_XET enum=OffOn
    /-stat/solosw/33  E32 F_XET enum=OffOn
    /-stat/solosw/34  E32 F_XET enum=OffOn
    /-stat/solosw/35  E32 F_XET enum=OffOn
    /-stat/solosw/36  E32 F_XET enum=OffOn
    /-stat/solosw/37  E32 F_XET enum=OffOn
    /-stat/solosw/38  E32 F_XET enum=OffOn
    /-stat/solosw/39  E32 F_XET enum=OffOn
    /-stat/solosw/40  E32 F_XET enum=OffOn
    /-stat/solosw/41  E32 F_XET enum=OffOn
    /-stat/solosw/42  E32 F_XET enum=OffOn
    /-stat/solosw/43  E32 F_XET enum=OffOn
    /-stat/solosw/44  E32 F_XET enum=OffOn
    /-stat/solosw/45  E32 F_XET enum=OffOn
    /-stat/solosw/46  E32 F_XET enum=OffOn
    /-stat/solosw/47  E32 F_XET enum=OffOn
    /-stat/solosw/48  E32 F_XET enum=OffOn
    /-stat/solosw/49  E32 F_XET enum=OffOn
    /-stat/solosw/50  E32 F_XET enum=OffOn
    /-stat/solosw/51  E32 F_XET enum=OffOn
    /-stat/solosw/52  E32 F_XET enum=OffOn
    /-stat/solosw/53  E32 F_XET enum=OffOn
    /-stat/solosw/54  E32 F_XET enum=OffOn
    /-stat/solosw/55  E32 F_XET enum=OffOn
    /-stat/solosw/56  E32 F_XET enum=OffOn
    /-stat/solosw/57  E32 F_XET enum=OffOn
    /-stat/solosw/58  E32 F_XET enum=OffOn
    /-stat/solosw/59  E32 F_XET enum=OffOn
    /-stat/solosw/60  E32 F_XET enum=OffOn
    /-stat/solosw/61  E32 F_XET enum=OffOn
    /-stat/solosw/62  E32 F_XET enum=OffOn
    /-stat/solosw/63  E32 F_XET enum=OffOn
    /-stat/solosw/64  E32 F_XET enum=OffOn
    /-stat/solosw/65  E32 F_XET enum=OffOn
    /-stat/solosw/66  E32 F_XET enum=OffOn
    /-stat/solosw/67  E32 F_XET enum=OffOn
    /-stat/solosw/68  E32 F_XET enum=OffOn
    /-stat/solosw/69  E32 F_XET enum=OffOn
    /-stat/solosw/70  E32 F_XET enum=OffOn
    /-stat/solosw/71  E32 F_XET enum=OffOn
    /-stat/solosw/72  E32 F_XET enum=OffOn
    /-stat/solosw/73  E32 F_XET enum=OffOn
    /-stat/solosw/74  E32 F_XET enum=OffOn
    /-stat/solosw/75  E32 F_XET enum=OffOn
    /-stat/solosw/76  E32 F_XET enum=OffOn
    /-stat/solosw/77  E32 F_XET enum=OffOn
    /-stat/solosw/78  E32 F_XET enum=OffOn
    /-stat/solosw/79  E32 F_XET enum=OffOn
    /-stat/solosw/80  E32 F_XET enum=OffOn
/-stat/aes50  <SAES> n=0
    /-stat/aes50/A  S32 F_XET
    /-stat/aes50/B  S32 F_XET
    /-stat/aes50/state  P32 F_XET
/-stat/tape  <STAPE> n=0
    /-stat/tape/state  E32 F_XET enum=Stapl
    /-stat/tape/file  S32 F_XET
    /-stat/tape/etime  E32 F_XET
    /-stat/tape/rtime  E32 F_XET
/-stat/osc  <SOSC> n=0
    /-stat/osc/on  E32 F_XET enum=OffOn
/-stat/talk  <STALK> n=0
    /-stat/talk/A  E32 F_XET enum=OffOn
    /-stat/talk/B  E32 F_XET enum=OffOn
    /-stat/urec  E32 F_XET
    /-stat/urec/state  E32 F_XET
    /-stat/urec/etime  E32 F_XET
    /-stat/urec/rtime  E32 F_XET
```

### Xaction (X32PrefStat.h, 30 entries)

```
/-action  <ACTION> n=0
    /-action/setip  I32 F_XET
    /-action/setclock  S32 F_XET
    /-action/initall  I32 F_XET
    /-action/initlib  I32 F_XET
    /-action/initshow  I32 F_XET
    /-action/savestate  I32 F_XET
    /-action/undopt  I32 F_XET
    /-action/doundo  I32 F_XET
    /-action/platrack  I32 F_XET
    /-action/newscreen  I32 F_XET
    /-action/clearsolo  I32 F_XET
    /-action/setprebus  I32 F_XET
    /-action/setsrate  I32 F_XET
    /-action/setrtasrc  I32 F_XET
    /-action/newscreen  I32 F_XET
    /-action/recselect  I32 F_XET
    /-action/gocue  I32 F_XET
    /-action/goscene  I32 F_XET
    /-action/undopt  I32 F_XET
    /-action/gosnippet  I32 F_XET
    /-action/selsession  I32 F_XET
    /-action/delsession  I32 F_XET
    /-action/selmarker  I32 F_XET
    /-action/delmarker  I32 F_XET
    /-action/savemarker  I32 F_XET
    /-action/addmarker  I32 F_XET
    /-action/selposition  I32 F_XET
    /-action/clearalert  I32 F_XET
    /-action/formatcard  I32 F_XET
```

### Xurec (X32PrefStat.h, 219 entries)

```
/-urec  <UREC> n=0
    /-urec/sessionmax  I32 F_XET
    /-urec/markermax  I32 F_XET
    /-urec/sessionlen  I32 F_XET
    /-urec/sessionpos  I32 F_XET
    /-urec/markerpos  I32 F_XET
    /-urec/batterystate  E32 F_XET enum=Ubat
    /-urec/srate  I32 F_XET
    /-urec/tracks  I32 F_XET
    /-urec/sessionspan  I32 F_XET
    /-urec/sessionoffs  I32 F_XET
    /-urec/sd1state  E32 F_XET enum=Usdc
    /-urec/sd2state  E32 F_XET enum=Usdc
    /-urec/sd1info  S32 F_XET
    /-urec/sd2info  S32 F_XET
    /-urec/errormessage  S32 F_XET
    /-urec/errorcode  I32 F_XET
/-urec/session  <S32> n=0
    /-urec/session/001/name  S32 F_XET
    /-urec/session/002/name  S32 F_XET
    /-urec/session/003/name  S32 F_XET
    /-urec/session/004/name  S32 F_XET
    /-urec/session/005/name  S32 F_XET
    /-urec/session/006/name  S32 F_XET
    /-urec/session/007/name  S32 F_XET
    /-urec/session/008/name  S32 F_XET
    /-urec/session/009/name  S32 F_XET
    /-urec/session/010/name  S32 F_XET
    /-urec/session/011/name  S32 F_XET
    /-urec/session/012/name  S32 F_XET
    /-urec/session/013/name  S32 F_XET
    /-urec/session/014/name  S32 F_XET
    /-urec/session/015/name  S32 F_XET
    /-urec/session/016/name  S32 F_XET
    /-urec/session/017/name  S32 F_XET
    /-urec/session/018/name  S32 F_XET
    /-urec/session/019/name  S32 F_XET
    /-urec/session/020/name  S32 F_XET
    /-urec/session/021/name  S32 F_XET
    /-urec/session/022/name  S32 F_XET
    /-urec/session/023/name  S32 F_XET
    /-urec/session/024/name  S32 F_XET
    /-urec/session/025/name  S32 F_XET
    /-urec/session/026/name  S32 F_XET
    /-urec/session/027/name  S32 F_XET
    /-urec/session/028/name  S32 F_XET
    /-urec/session/029/name  S32 F_XET
    /-urec/session/030/name  S32 F_XET
    /-urec/session/031/name  S32 F_XET
    /-urec/session/032/name  S32 F_XET
    /-urec/session/033/name  S32 F_XET
    /-urec/session/034/name  S32 F_XET
    /-urec/session/035/name  S32 F_XET
    /-urec/session/036/name  S32 F_XET
    /-urec/session/037/name  S32 F_XET
    /-urec/session/038/name  S32 F_XET
    /-urec/session/039/name  S32 F_XET
    /-urec/session/040/name  S32 F_XET
    /-urec/session/041/name  S32 F_XET
    /-urec/session/042/name  S32 F_XET
    /-urec/session/043/name  S32 F_XET
    /-urec/session/044/name  S32 F_XET
    /-urec/session/045/name  S32 F_XET
    /-urec/session/046/name  S32 F_XET
    /-urec/session/047/name  S32 F_XET
    /-urec/session/048/name  S32 F_XET
    /-urec/session/049/name  S32 F_XET
    /-urec/session/050/name  S32 F_XET
    /-urec/session/051/name  S32 F_XET
    /-urec/session/052/name  S32 F_XET
    /-urec/session/053/name  S32 F_XET
    /-urec/session/054/name  S32 F_XET
    /-urec/session/055/name  S32 F_XET
    /-urec/session/056/name  S32 F_XET
    /-urec/session/057/name  S32 F_XET
    /-urec/session/058/name  S32 F_XET
    /-urec/session/059/name  S32 F_XET
    /-urec/session/060/name  S32 F_XET
    /-urec/session/061/name  S32 F_XET
    /-urec/session/062/name  S32 F_XET
    /-urec/session/063/name  S32 F_XET
    /-urec/session/064/name  S32 F_XET
    /-urec/session/065/name  S32 F_XET
    /-urec/session/066/name  S32 F_XET
    /-urec/session/067/name  S32 F_XET
    /-urec/session/068/name  S32 F_XET
    /-urec/session/069/name  S32 F_XET
    /-urec/session/070/name  S32 F_XET
    /-urec/session/071/name  S32 F_XET
    /-urec/session/072/name  S32 F_XET
    /-urec/session/073/name  S32 F_XET
    /-urec/session/074/name  S32 F_XET
    /-urec/session/075/name  S32 F_XET
    /-urec/session/076/name  S32 F_XET
    /-urec/session/077/name  S32 F_XET
    /-urec/session/078/name  S32 F_XET
    /-urec/session/079/name  S32 F_XET
    /-urec/session/080/name  S32 F_XET
    /-urec/session/081/name  S32 F_XET
    /-urec/session/082/name  S32 F_XET
    /-urec/session/083/name  S32 F_XET
    /-urec/session/084/name  S32 F_XET
    /-urec/session/085/name  S32 F_XET
    /-urec/session/086/name  S32 F_XET
    /-urec/session/087/name  S32 F_XET
    /-urec/session/088/name  S32 F_XET
    /-urec/session/089/name  S32 F_XET
    /-urec/session/090/name  S32 F_XET
    /-urec/session/091/name  S32 F_XET
    /-urec/session/092/name  S32 F_XET
    /-urec/session/093/name  S32 F_XET
    /-urec/session/094/name  S32 F_XET
    /-urec/session/095/name  S32 F_XET
    /-urec/session/096/name  S32 F_XET
    /-urec/session/097/name  S32 F_XET
    /-urec/session/098/name  S32 F_XET
    /-urec/session/099/name  S32 F_XET
    /-urec/session/100/name  S32 F_XET
/-urec/marker  <I32> n=0
    /-urec/marker/001/time  I32 F_XET
    /-urec/marker/002/time  I32 F_XET
    /-urec/marker/003/time  I32 F_XET
    /-urec/marker/004/time  I32 F_XET
    /-urec/marker/005/time  I32 F_XET
    /-urec/marker/006/time  I32 F_XET
    /-urec/marker/007/time  I32 F_XET
    /-urec/marker/008/time  I32 F_XET
    /-urec/marker/009/time  I32 F_XET
    /-urec/marker/010/time  I32 F_XET
    /-urec/marker/011/time  I32 F_XET
    /-urec/marker/012/time  I32 F_XET
    /-urec/marker/013/time  I32 F_XET
    /-urec/marker/014/time  I32 F_XET
    /-urec/marker/015/time  I32 F_XET
    /-urec/marker/016/time  I32 F_XET
    /-urec/marker/017/time  I32 F_XET
    /-urec/marker/018/time  I32 F_XET
    /-urec/marker/019/time  I32 F_XET
    /-urec/marker/020/time  I32 F_XET
    /-urec/marker/021/time  I32 F_XET
    /-urec/marker/022/time  I32 F_XET
    /-urec/marker/023/time  I32 F_XET
    /-urec/marker/024/time  I32 F_XET
    /-urec/marker/025/time  I32 F_XET
    /-urec/marker/026/time  I32 F_XET
    /-urec/marker/027/time  I32 F_XET
    /-urec/marker/028/time  I32 F_XET
    /-urec/marker/029/time  I32 F_XET
    /-urec/marker/030/time  I32 F_XET
    /-urec/marker/031/time  I32 F_XET
    /-urec/marker/032/time  I32 F_XET
    /-urec/marker/033/time  I32 F_XET
    /-urec/marker/034/time  I32 F_XET
    /-urec/marker/035/time  I32 F_XET
    /-urec/marker/036/time  I32 F_XET
    /-urec/marker/037/time  I32 F_XET
    /-urec/marker/038/time  I32 F_XET
    /-urec/marker/039/time  I32 F_XET
    /-urec/marker/040/time  I32 F_XET
    /-urec/marker/041/time  I32 F_XET
    /-urec/marker/042/time  I32 F_XET
    /-urec/marker/043/time  I32 F_XET
    /-urec/marker/044/time  I32 F_XET
    /-urec/marker/045/time  I32 F_XET
    /-urec/marker/046/time  I32 F_XET
    /-urec/marker/047/time  I32 F_XET
    /-urec/marker/048/time  I32 F_XET
    /-urec/marker/049/time  I32 F_XET
    /-urec/marker/050/time  I32 F_XET
    /-urec/marker/051/time  I32 F_XET
    /-urec/marker/052/time  I32 F_XET
    /-urec/marker/053/time  I32 F_XET
    /-urec/marker/054/time  I32 F_XET
    /-urec/marker/055/time  I32 F_XET
    /-urec/marker/056/time  I32 F_XET
    /-urec/marker/057/time  I32 F_XET
    /-urec/marker/058/time  I32 F_XET
    /-urec/marker/059/time  I32 F_XET
    /-urec/marker/060/time  I32 F_XET
    /-urec/marker/061/time  I32 F_XET
    /-urec/marker/062/time  I32 F_XET
    /-urec/marker/063/time  I32 F_XET
    /-urec/marker/064/time  I32 F_XET
    /-urec/marker/065/time  I32 F_XET
    /-urec/marker/066/time  I32 F_XET
    /-urec/marker/067/time  I32 F_XET
    /-urec/marker/068/time  I32 F_XET
    /-urec/marker/069/time  I32 F_XET
    /-urec/marker/070/time  I32 F_XET
    /-urec/marker/071/time  I32 F_XET
    /-urec/marker/072/time  I32 F_XET
    /-urec/marker/073/time  I32 F_XET
    /-urec/marker/074/time  I32 F_XET
    /-urec/marker/075/time  I32 F_XET
    /-urec/marker/076/time  I32 F_XET
    /-urec/marker/077/time  I32 F_XET
    /-urec/marker/078/time  I32 F_XET
    /-urec/marker/079/time  I32 F_XET
    /-urec/marker/080/time  I32 F_XET
    /-urec/marker/081/time  I32 F_XET
    /-urec/marker/082/time  I32 F_XET
    /-urec/marker/083/time  I32 F_XET
    /-urec/marker/084/time  I32 F_XET
    /-urec/marker/085/time  I32 F_XET
    /-urec/marker/086/time  I32 F_XET
    /-urec/marker/087/time  I32 F_XET
    /-urec/marker/088/time  I32 F_XET
    /-urec/marker/089/time  I32 F_XET
    /-urec/marker/090/time  I32 F_XET
    /-urec/marker/091/time  I32 F_XET
    /-urec/marker/092/time  I32 F_XET
    /-urec/marker/093/time  I32 F_XET
    /-urec/marker/094/time  I32 F_XET
    /-urec/marker/095/time  I32 F_XET
    /-urec/marker/096/time  I32 F_XET
    /-urec/marker/097/time  I32 F_XET
    /-urec/marker/098/time  I32 F_XET
    /-urec/marker/099/time  I32 F_XET
    /-urec/marker/100/time  I32 F_XET
```

### Xauxin01 (X32Auxin.h, 114 entries)

```
/auxin  <CHCO> n=0
/auxin/01  <CHCO> n=0
/auxin/01/config  <CHCO> n=0
    /auxin/01/config/name  S32 F_XET
    /auxin/01/config/icon  I32 F_XET
    /auxin/01/config/color  E32 F_XET enum=Xcolors
    /auxin/01/config/source  I32 F_XET
/auxin/01/preamp  <AXPR> n=0
    /auxin/01/preamp/trim  F32 F_XET
    /auxin/01/preamp/invert  E32 F_XET enum=OffOn
/auxin/01/eq  <OFFON> n=1
    /auxin/01/eq/on  E32 F_XET enum=OffOn
/auxin/01/eq/1  <CHEQ> n=0
    /auxin/01/eq/1/type  E32 F_XET enum=Xeqty1
    /auxin/01/eq/1/f  F32 F_XET
    /auxin/01/eq/1/g  F32 F_XET
    /auxin/01/eq/1/q  F32 F_XET
/auxin/01/eq/2  <CHEQ> n=0
    /auxin/01/eq/2/type  E32 F_XET enum=Xeqty1
    /auxin/01/eq/2/f  F32 F_XET
    /auxin/01/eq/2/g  F32 F_XET
    /auxin/01/eq/2/q  F32 F_XET
/auxin/01/eq/3  <CHEQ> n=0
    /auxin/01/eq/3/type  E32 F_XET enum=Xeqty1
    /auxin/01/eq/3/f  F32 F_XET
    /auxin/01/eq/3/g  F32 F_XET
    /auxin/01/eq/3/q  F32 F_XET
/auxin/01/eq/4  <CHEQ> n=0
    /auxin/01/eq/4/type  E32 F_XET enum=Xeqty1
    /auxin/01/eq/4/f  F32 F_XET
    /auxin/01/eq/4/g  F32 F_XET
    /auxin/01/eq/4/q  F32 F_XET
/auxin/01/mix  <CHMX> n=0
    /auxin/01/mix/on  E32 F_XET enum=OffOn
    /auxin/01/mix/fader  F32 F_XET
    /auxin/01/mix/st  E32 F_XET enum=OffOn
    /auxin/01/mix/pan  F32 F_XET
    /auxin/01/mix/mono  E32 F_XET enum=OffOn
    /auxin/01/mix/mlevel  F32 F_XET
/auxin/01/mix/01  <CHMO> n=0
    /auxin/01/mix/01/on  E32 F_XET enum=OffOn
    /auxin/01/mix/01/level  F32 F_XET
    /auxin/01/mix/01/pan  F32 F_XET
    /auxin/01/mix/01/type  E32 F_XET enum=Xmtype
    /auxin/01/mix/01/panFollow  E32 F_XET
/auxin/01/mix/02  <CHME> n=0
    /auxin/01/mix/02/on  E32 F_XET enum=OffOn
    /auxin/01/mix/02/level  F32 F_XET
/auxin/01/mix/03  <CHMO> n=0
    /auxin/01/mix/03/on  E32 F_XET enum=OffOn
    /auxin/01/mix/03/level  F32 F_XET
    /auxin/01/mix/03/pan  F32 F_XET
    /auxin/01/mix/03/type  E32 F_XET enum=Xmtype
    /auxin/01/mix/03/panFollow  E32 F_XET
/auxin/01/mix/04  <CHME> n=0
    /auxin/01/mix/04/on  E32 F_XET enum=OffOn
    /auxin/01/mix/04/level  F32 F_XET
/auxin/01/mix/05  <CHMO> n=0
    /auxin/01/mix/05/on  E32 F_XET enum=OffOn
    /auxin/01/mix/05/level  F32 F_XET
    /auxin/01/mix/05/pan  F32 F_XET
    /auxin/01/mix/05/type  E32 F_XET enum=Xmtype
    /auxin/01/mix/05/panFollow  E32 F_XET
/auxin/01/mix/06  <CHME> n=0
    /auxin/01/mix/06/on  E32 F_XET enum=OffOn
    /auxin/01/mix/06/level  F32 F_XET
/auxin/01/mix/07  <CHMO> n=0
    /auxin/01/mix/07/on  E32 F_XET enum=OffOn
    /auxin/01/mix/07/level  F32 F_XET
    /auxin/01/mix/07/pan  F32 F_XET
    /auxin/01/mix/07/type  E32 F_XET enum=Xmtype
    /auxin/01/mix/07/panFollow  E32 F_XET
/auxin/01/mix/08  <CHME> n=0
    /auxin/01/mix/08/on  E32 F_XET enum=OffOn
    /auxin/01/mix/08/level  F32 F_XET
/auxin/01/mix/09  <CHMO> n=0
    /auxin/01/mix/09/on  E32 F_XET enum=OffOn
    /auxin/01/mix/09/level  F32 F_XET
    /auxin/01/mix/09/pan  F32 F_XET
    /auxin/01/mix/09/type  E32 F_XET enum=Xmtype
    /auxin/01/mix/09/panFollow  E32 F_XET
/auxin/01/mix/10  <CHME> n=0
    /auxin/01/mix/10/on  E32 F_XET enum=OffOn
    /auxin/01/mix/10/level  F32 F_XET
/auxin/01/mix/11  <CHMO> n=0
    /auxin/01/mix/11/on  E32 F_XET enum=OffOn
    /auxin/01/mix/11/level  F32 F_XET
    /auxin/01/mix/11/pan  F32 F_XET
    /auxin/01/mix/11/type  E32 F_XET enum=Xmtype
    /auxin/01/mix/11/panFollow  E32 F_XET
/auxin/01/mix/12  <CHME> n=0
    /auxin/01/mix/12/on  E32 F_XET enum=OffOn
    /auxin/01/mix/12/level  F32 F_XET
/auxin/01/mix/13  <CHMO> n=0
    /auxin/01/mix/13/on  E32 F_XET enum=OffOn
    /auxin/01/mix/13/level  F32 F_XET
    /auxin/01/mix/13/pan  F32 F_XET
    /auxin/01/mix/13/type  E32 F_XET enum=Xmtype
    /auxin/01/mix/13/panFollow  E32 F_XET
/auxin/01/mix/14  <CHME> n=0
    /auxin/01/mix/14/on  E32 F_XET enum=OffOn
    /auxin/01/mix/14/level  F32 F_XET
/auxin/01/mix/15  <CHMO> n=0
    /auxin/01/mix/15/on  E32 F_XET enum=OffOn
    /auxin/01/mix/15/level  F32 F_XET
    /auxin/01/mix/15/pan  F32 F_XET
    /auxin/01/mix/15/type  E32 F_XET enum=Xmtype
    /auxin/01/mix/15/panFollow  E32 F_XET
/auxin/01/mix/16  <CHME> n=0
    /auxin/01/mix/16/on  E32 F_XET enum=OffOn
    /auxin/01/mix/16/level  F32 F_XET
/auxin/01/grp  <CHGRP> n=0
    /auxin/01/grp/dca  P32 F_XET
    /auxin/01/grp/mute  P32 F_XET
```

### Xauxin02 (X32Auxin.h, 114 entries)

```
/auxin  <CHCO> n=0
/auxin/02  <CHCO> n=0
/auxin/02/config  <CHCO> n=0
    /auxin/02/config/name  S32 F_XET
    /auxin/02/config/icon  I32 F_XET
    /auxin/02/config/color  E32 F_XET enum=Xcolors
    /auxin/02/config/source  I32 F_XET
/auxin/02/preamp  <AXPR> n=0
    /auxin/02/preamp/trim  F32 F_XET
    /auxin/02/preamp/invert  E32 F_XET enum=OffOn
/auxin/02/eq  <OFFON> n=1
    /auxin/02/eq/on  E32 F_XET enum=OffOn
/auxin/02/eq/1  <CHEQ> n=0
    /auxin/02/eq/1/type  E32 F_XET enum=Xeqty1
    /auxin/02/eq/1/f  F32 F_XET
    /auxin/02/eq/1/g  F32 F_XET
    /auxin/02/eq/1/q  F32 F_XET
/auxin/02/eq/2  <CHEQ> n=0
    /auxin/02/eq/2/type  E32 F_XET enum=Xeqty1
    /auxin/02/eq/2/f  F32 F_XET
    /auxin/02/eq/2/g  F32 F_XET
    /auxin/02/eq/2/q  F32 F_XET
/auxin/02/eq/3  <CHEQ> n=0
    /auxin/02/eq/3/type  E32 F_XET enum=Xeqty1
    /auxin/02/eq/3/f  F32 F_XET
    /auxin/02/eq/3/g  F32 F_XET
    /auxin/02/eq/3/q  F32 F_XET
/auxin/02/eq/4  <CHEQ> n=0
    /auxin/02/eq/4/type  E32 F_XET enum=Xeqty1
    /auxin/02/eq/4/f  F32 F_XET
    /auxin/02/eq/4/g  F32 F_XET
    /auxin/02/eq/4/q  F32 F_XET
/auxin/02/mix  <CHMX> n=0
    /auxin/02/mix/on  E32 F_XET enum=OffOn
    /auxin/02/mix/fader  F32 F_XET
    /auxin/02/mix/st  E32 F_XET enum=OffOn
    /auxin/02/mix/pan  F32 F_XET
    /auxin/02/mix/mono  E32 F_XET enum=OffOn
    /auxin/02/mix/mlevel  F32 F_XET
/auxin/02/mix/01  <CHMO> n=0
    /auxin/02/mix/01/on  E32 F_XET enum=OffOn
    /auxin/02/mix/01/level  F32 F_XET
    /auxin/02/mix/01/pan  F32 F_XET
    /auxin/02/mix/01/type  E32 F_XET enum=Xmtype
    /auxin/02/mix/01/panFollow  E32 F_XET
/auxin/02/mix/02  <CHME> n=0
    /auxin/02/mix/02/on  E32 F_XET enum=OffOn
    /auxin/02/mix/02/level  F32 F_XET
/auxin/02/mix/03  <CHMO> n=0
    /auxin/02/mix/03/on  E32 F_XET enum=OffOn
    /auxin/02/mix/03/level  F32 F_XET
    /auxin/02/mix/03/pan  F32 F_XET
    /auxin/02/mix/03/type  E32 F_XET enum=Xmtype
    /auxin/02/mix/03/panFollow  E32 F_XET
/auxin/02/mix/04  <CHME> n=0
    /auxin/02/mix/04/on  E32 F_XET enum=OffOn
    /auxin/02/mix/04/level  F32 F_XET
/auxin/02/mix/05  <CHMO> n=0
    /auxin/02/mix/05/on  E32 F_XET enum=OffOn
    /auxin/02/mix/05/level  F32 F_XET
    /auxin/02/mix/05/pan  F32 F_XET
    /auxin/02/mix/05/type  E32 F_XET enum=Xmtype
    /auxin/02/mix/05/panFollow  E32 F_XET
/auxin/02/mix/06  <CHME> n=0
    /auxin/02/mix/06/on  E32 F_XET enum=OffOn
    /auxin/02/mix/06/level  F32 F_XET
/auxin/02/mix/07  <CHMO> n=0
    /auxin/02/mix/07/on  E32 F_XET enum=OffOn
    /auxin/02/mix/07/level  F32 F_XET
    /auxin/02/mix/07/pan  F32 F_XET
    /auxin/02/mix/07/type  E32 F_XET enum=Xmtype
    /auxin/02/mix/07/panFollow  E32 F_XET
/auxin/02/mix/08  <CHME> n=0
    /auxin/02/mix/08/on  E32 F_XET enum=OffOn
    /auxin/02/mix/08/level  F32 F_XET
/auxin/02/mix/09  <CHMO> n=0
    /auxin/02/mix/09/on  E32 F_XET enum=OffOn
    /auxin/02/mix/09/level  F32 F_XET
    /auxin/02/mix/09/pan  F32 F_XET
    /auxin/02/mix/09/type  E32 F_XET enum=Xmtype
    /auxin/02/mix/09/panFollow  E32 F_XET
/auxin/02/mix/10  <CHME> n=0
    /auxin/02/mix/10/on  E32 F_XET enum=OffOn
    /auxin/02/mix/10/level  F32 F_XET
/auxin/02/mix/11  <CHMO> n=0
    /auxin/02/mix/11/on  E32 F_XET enum=OffOn
    /auxin/02/mix/11/level  F32 F_XET
    /auxin/02/mix/11/pan  F32 F_XET
    /auxin/02/mix/11/type  E32 F_XET enum=Xmtype
    /auxin/02/mix/11/panFollow  E32 F_XET
/auxin/02/mix/12  <CHME> n=0
    /auxin/02/mix/12/on  E32 F_XET enum=OffOn
    /auxin/02/mix/12/level  F32 F_XET
/auxin/02/mix/13  <CHMO> n=0
    /auxin/02/mix/13/on  E32 F_XET enum=OffOn
    /auxin/02/mix/13/level  F32 F_XET
    /auxin/02/mix/13/pan  F32 F_XET
    /auxin/02/mix/13/type  E32 F_XET enum=Xmtype
    /auxin/02/mix/13/panFollow  E32 F_XET
/auxin/02/mix/14  <CHME> n=0
    /auxin/02/mix/14/on  E32 F_XET enum=OffOn
    /auxin/02/mix/14/level  F32 F_XET
/auxin/02/mix/15  <CHMO> n=0
    /auxin/02/mix/15/on  E32 F_XET enum=OffOn
    /auxin/02/mix/15/level  F32 F_XET
    /auxin/02/mix/15/pan  F32 F_XET
    /auxin/02/mix/15/type  E32 F_XET enum=Xmtype
    /auxin/02/mix/15/panFollow  E32 F_XET
/auxin/02/mix/16  <CHME> n=0
    /auxin/02/mix/16/on  E32 F_XET enum=OffOn
    /auxin/02/mix/16/level  F32 F_XET
/auxin/02/grp  <CHGRP> n=0
    /auxin/02/grp/dca  P32 F_XET
    /auxin/02/grp/mute  P32 F_XET
```

### Xauxin03 (X32Auxin.h, 114 entries)

```
/auxin  <CHCO> n=0
/auxin/03  <CHCO> n=0
/auxin/03/config  <CHCO> n=0
    /auxin/03/config/name  S32 F_XET
    /auxin/03/config/icon  I32 F_XET
    /auxin/03/config/color  E32 F_XET enum=Xcolors
    /auxin/03/config/source  I32 F_XET
/auxin/03/preamp  <AXPR> n=0
    /auxin/03/preamp/trim  F32 F_XET
    /auxin/03/preamp/invert  E32 F_XET enum=OffOn
/auxin/03/eq  <OFFON> n=1
    /auxin/03/eq/on  E32 F_XET enum=OffOn
/auxin/03/eq/1  <CHEQ> n=0
    /auxin/03/eq/1/type  E32 F_XET enum=Xeqty1
    /auxin/03/eq/1/f  F32 F_XET
    /auxin/03/eq/1/g  F32 F_XET
    /auxin/03/eq/1/q  F32 F_XET
/auxin/03/eq/2  <CHEQ> n=0
    /auxin/03/eq/2/type  E32 F_XET enum=Xeqty1
    /auxin/03/eq/2/f  F32 F_XET
    /auxin/03/eq/2/g  F32 F_XET
    /auxin/03/eq/2/q  F32 F_XET
/auxin/03/eq/3  <CHEQ> n=0
    /auxin/03/eq/3/type  E32 F_XET enum=Xeqty1
    /auxin/03/eq/3/f  F32 F_XET
    /auxin/03/eq/3/g  F32 F_XET
    /auxin/03/eq/3/q  F32 F_XET
/auxin/03/eq/4  <CHEQ> n=0
    /auxin/03/eq/4/type  E32 F_XET enum=Xeqty1
    /auxin/03/eq/4/f  F32 F_XET
    /auxin/03/eq/4/g  F32 F_XET
    /auxin/03/eq/4/q  F32 F_XET
/auxin/03/mix  <CHMX> n=0
    /auxin/03/mix/on  E32 F_XET enum=OffOn
    /auxin/03/mix/fader  F32 F_XET
    /auxin/03/mix/st  E32 F_XET enum=OffOn
    /auxin/03/mix/pan  F32 F_XET
    /auxin/03/mix/mono  E32 F_XET enum=OffOn
    /auxin/03/mix/mlevel  F32 F_XET
/auxin/03/mix/01  <CHMO> n=0
    /auxin/03/mix/01/on  E32 F_XET enum=OffOn
    /auxin/03/mix/01/level  F32 F_XET
    /auxin/03/mix/01/pan  F32 F_XET
    /auxin/03/mix/01/type  E32 F_XET enum=Xmtype
    /auxin/03/mix/01/panFollow  E32 F_XET
/auxin/03/mix/02  <CHME> n=0
    /auxin/03/mix/02/on  E32 F_XET enum=OffOn
    /auxin/03/mix/02/level  F32 F_XET
/auxin/03/mix/03  <CHMO> n=0
    /auxin/03/mix/03/on  E32 F_XET enum=OffOn
    /auxin/03/mix/03/level  F32 F_XET
    /auxin/03/mix/03/pan  F32 F_XET
    /auxin/03/mix/03/type  E32 F_XET enum=Xmtype
    /auxin/03/mix/03/panFollow  E32 F_XET
/auxin/03/mix/04  <CHME> n=0
    /auxin/03/mix/04/on  E32 F_XET enum=OffOn
    /auxin/03/mix/04/level  F32 F_XET
/auxin/03/mix/05  <CHMO> n=0
    /auxin/03/mix/05/on  E32 F_XET enum=OffOn
    /auxin/03/mix/05/level  F32 F_XET
    /auxin/03/mix/05/pan  F32 F_XET
    /auxin/03/mix/05/type  E32 F_XET enum=Xmtype
    /auxin/03/mix/05/panFollow  E32 F_XET
/auxin/03/mix/06  <CHME> n=0
    /auxin/03/mix/06/on  E32 F_XET enum=OffOn
    /auxin/03/mix/06/level  F32 F_XET
/auxin/03/mix/07  <CHMO> n=0
    /auxin/03/mix/07/on  E32 F_XET enum=OffOn
    /auxin/03/mix/07/level  F32 F_XET
    /auxin/03/mix/07/pan  F32 F_XET
    /auxin/03/mix/07/type  E32 F_XET enum=Xmtype
    /auxin/03/mix/07/panFollow  E32 F_XET
/auxin/03/mix/08  <CHME> n=0
    /auxin/03/mix/08/on  E32 F_XET enum=OffOn
    /auxin/03/mix/08/level  F32 F_XET
/auxin/03/mix/09  <CHMO> n=0
    /auxin/03/mix/09/on  E32 F_XET enum=OffOn
    /auxin/03/mix/09/level  F32 F_XET
    /auxin/03/mix/09/pan  F32 F_XET
    /auxin/03/mix/09/type  E32 F_XET enum=Xmtype
    /auxin/03/mix/09/panFollow  E32 F_XET
/auxin/03/mix/10  <CHME> n=0
    /auxin/03/mix/10/on  E32 F_XET enum=OffOn
    /auxin/03/mix/10/level  F32 F_XET
/auxin/03/mix/11  <CHMO> n=0
    /auxin/03/mix/11/on  E32 F_XET enum=OffOn
    /auxin/03/mix/11/level  F32 F_XET
    /auxin/03/mix/11/pan  F32 F_XET
    /auxin/03/mix/11/type  E32 F_XET enum=Xmtype
    /auxin/03/mix/11/panFollow  E32 F_XET
/auxin/03/mix/12  <CHME> n=0
    /auxin/03/mix/12/on  E32 F_XET enum=OffOn
    /auxin/03/mix/12/level  F32 F_XET
/auxin/03/mix/13  <CHMO> n=0
    /auxin/03/mix/13/on  E32 F_XET enum=OffOn
    /auxin/03/mix/13/level  F32 F_XET
    /auxin/03/mix/13/pan  F32 F_XET
    /auxin/03/mix/13/type  E32 F_XET enum=Xmtype
    /auxin/03/mix/13/panFollow  E32 F_XET
/auxin/03/mix/14  <CHME> n=0
    /auxin/03/mix/14/on  E32 F_XET enum=OffOn
    /auxin/03/mix/14/level  F32 F_XET
/auxin/03/mix/15  <CHMO> n=0
    /auxin/03/mix/15/on  E32 F_XET enum=OffOn
    /auxin/03/mix/15/level  F32 F_XET
    /auxin/03/mix/15/pan  F32 F_XET
    /auxin/03/mix/15/type  E32 F_XET enum=Xmtype
    /auxin/03/mix/15/panFollow  E32 F_XET
/auxin/03/mix/16  <CHME> n=0
    /auxin/03/mix/16/on  E32 F_XET enum=OffOn
    /auxin/03/mix/16/level  F32 F_XET
/auxin/03/grp  <CHGRP> n=0
    /auxin/03/grp/dca  P32 F_XET
    /auxin/03/grp/mute  P32 F_XET
```

### Xauxin04 (X32Auxin.h, 114 entries)

```
/auxin  <CHCO> n=0
/auxin/04  <CHCO> n=0
/auxin/04/config  <CHCO> n=0
    /auxin/04/config/name  S32 F_XET
    /auxin/04/config/icon  I32 F_XET
    /auxin/04/config/color  E32 F_XET enum=Xcolors
    /auxin/04/config/source  I32 F_XET
/auxin/04/preamp  <AXPR> n=0
    /auxin/04/preamp/trim  F32 F_XET
    /auxin/04/preamp/invert  E32 F_XET enum=OffOn
/auxin/04/eq  <OFFON> n=1
    /auxin/04/eq/on  E32 F_XET enum=OffOn
/auxin/04/eq/1  <CHEQ> n=0
    /auxin/04/eq/1/type  E32 F_XET enum=Xeqty1
    /auxin/04/eq/1/f  F32 F_XET
    /auxin/04/eq/1/g  F32 F_XET
    /auxin/04/eq/1/q  F32 F_XET
/auxin/04/eq/2  <CHEQ> n=0
    /auxin/04/eq/2/type  E32 F_XET enum=Xeqty1
    /auxin/04/eq/2/f  F32 F_XET
    /auxin/04/eq/2/g  F32 F_XET
    /auxin/04/eq/2/q  F32 F_XET
/auxin/04/eq/3  <CHEQ> n=0
    /auxin/04/eq/3/type  E32 F_XET enum=Xeqty1
    /auxin/04/eq/3/f  F32 F_XET
    /auxin/04/eq/3/g  F32 F_XET
    /auxin/04/eq/3/q  F32 F_XET
/auxin/04/eq/4  <CHEQ> n=0
    /auxin/04/eq/4/type  E32 F_XET enum=Xeqty1
    /auxin/04/eq/4/f  F32 F_XET
    /auxin/04/eq/4/g  F32 F_XET
    /auxin/04/eq/4/q  F32 F_XET
/auxin/04/mix  <CHMX> n=0
    /auxin/04/mix/on  E32 F_XET enum=OffOn
    /auxin/04/mix/fader  F32 F_XET
    /auxin/04/mix/st  E32 F_XET enum=OffOn
    /auxin/04/mix/pan  F32 F_XET
    /auxin/04/mix/mono  E32 F_XET enum=OffOn
    /auxin/04/mix/mlevel  F32 F_XET
/auxin/04/mix/01  <CHMO> n=0
    /auxin/04/mix/01/on  E32 F_XET enum=OffOn
    /auxin/04/mix/01/level  F32 F_XET
    /auxin/04/mix/01/pan  F32 F_XET
    /auxin/04/mix/01/type  E32 F_XET enum=Xmtype
    /auxin/04/mix/01/panFollow  E32 F_XET
/auxin/04/mix/02  <CHME> n=0
    /auxin/04/mix/02/on  E32 F_XET enum=OffOn
    /auxin/04/mix/02/level  F32 F_XET
/auxin/04/mix/03  <CHMO> n=0
    /auxin/04/mix/03/on  E32 F_XET enum=OffOn
    /auxin/04/mix/03/level  F32 F_XET
    /auxin/04/mix/03/pan  F32 F_XET
    /auxin/04/mix/03/type  E32 F_XET enum=Xmtype
    /auxin/04/mix/03/panFollow  E32 F_XET
/auxin/04/mix/04  <CHME> n=0
    /auxin/04/mix/04/on  E32 F_XET enum=OffOn
    /auxin/04/mix/04/level  F32 F_XET
/auxin/04/mix/05  <CHMO> n=0
    /auxin/04/mix/05/on  E32 F_XET enum=OffOn
    /auxin/04/mix/05/level  F32 F_XET
    /auxin/04/mix/05/pan  F32 F_XET
    /auxin/04/mix/05/type  E32 F_XET enum=Xmtype
    /auxin/04/mix/05/panFollow  E32 F_XET
/auxin/04/mix/06  <CHME> n=0
    /auxin/04/mix/06/on  E32 F_XET enum=OffOn
    /auxin/04/mix/06/level  F32 F_XET
/auxin/04/mix/07  <CHMO> n=0
    /auxin/04/mix/07/on  E32 F_XET enum=OffOn
    /auxin/04/mix/07/level  F32 F_XET
    /auxin/04/mix/07/pan  F32 F_XET
    /auxin/04/mix/07/type  E32 F_XET enum=Xmtype
    /auxin/04/mix/07/panFollow  E32 F_XET
/auxin/04/mix/08  <CHME> n=0
    /auxin/04/mix/08/on  E32 F_XET enum=OffOn
    /auxin/04/mix/08/level  F32 F_XET
/auxin/04/mix/09  <CHMO> n=0
    /auxin/04/mix/09/on  E32 F_XET enum=OffOn
    /auxin/04/mix/09/level  F32 F_XET
    /auxin/04/mix/09/pan  F32 F_XET
    /auxin/04/mix/09/type  E32 F_XET enum=Xmtype
    /auxin/04/mix/09/panFollow  E32 F_XET
/auxin/04/mix/10  <CHME> n=0
    /auxin/04/mix/10/on  E32 F_XET enum=OffOn
    /auxin/04/mix/10/level  F32 F_XET
/auxin/04/mix/11  <CHMO> n=0
    /auxin/04/mix/11/on  E32 F_XET enum=OffOn
    /auxin/04/mix/11/level  F32 F_XET
    /auxin/04/mix/11/pan  F32 F_XET
    /auxin/04/mix/11/type  E32 F_XET enum=Xmtype
    /auxin/04/mix/11/panFollow  E32 F_XET
/auxin/04/mix/12  <CHME> n=0
    /auxin/04/mix/12/on  E32 F_XET enum=OffOn
    /auxin/04/mix/12/level  F32 F_XET
/auxin/04/mix/13  <CHMO> n=0
    /auxin/04/mix/13/on  E32 F_XET enum=OffOn
    /auxin/04/mix/13/level  F32 F_XET
    /auxin/04/mix/13/pan  F32 F_XET
    /auxin/04/mix/13/type  E32 F_XET enum=Xmtype
    /auxin/04/mix/13/panFollow  E32 F_XET
/auxin/04/mix/14  <CHME> n=0
    /auxin/04/mix/14/on  E32 F_XET enum=OffOn
    /auxin/04/mix/14/level  F32 F_XET
/auxin/04/mix/15  <CHMO> n=0
    /auxin/04/mix/15/on  E32 F_XET enum=OffOn
    /auxin/04/mix/15/level  F32 F_XET
    /auxin/04/mix/15/pan  F32 F_XET
    /auxin/04/mix/15/type  E32 F_XET enum=Xmtype
    /auxin/04/mix/15/panFollow  E32 F_XET
/auxin/04/mix/16  <CHME> n=0
    /auxin/04/mix/16/on  E32 F_XET enum=OffOn
    /auxin/04/mix/16/level  F32 F_XET
/auxin/04/grp  <CHGRP> n=0
    /auxin/04/grp/dca  P32 F_XET
    /auxin/04/grp/mute  P32 F_XET
```

### Xauxin05 (X32Auxin.h, 114 entries)

```
/auxin  <CHCO> n=0
/auxin/05  <CHCO> n=0
/auxin/05/config  <CHCO> n=0
    /auxin/05/config/name  S32 F_XET
    /auxin/05/config/icon  I32 F_XET
    /auxin/05/config/color  E32 F_XET enum=Xcolors
    /auxin/05/config/source  I32 F_XET
/auxin/05/preamp  <AXPR> n=0
    /auxin/05/preamp/trim  F32 F_XET
    /auxin/05/preamp/invert  E32 F_XET enum=OffOn
/auxin/05/eq  <OFFON> n=1
    /auxin/05/eq/on  E32 F_XET enum=OffOn
/auxin/05/eq/1  <CHEQ> n=0
    /auxin/05/eq/1/type  E32 F_XET enum=Xeqty1
    /auxin/05/eq/1/f  F32 F_XET
    /auxin/05/eq/1/g  F32 F_XET
    /auxin/05/eq/1/q  F32 F_XET
/auxin/05/eq/2  <CHEQ> n=0
    /auxin/05/eq/2/type  E32 F_XET enum=Xeqty1
    /auxin/05/eq/2/f  F32 F_XET
    /auxin/05/eq/2/g  F32 F_XET
    /auxin/05/eq/2/q  F32 F_XET
/auxin/05/eq/3  <CHEQ> n=0
    /auxin/05/eq/3/type  E32 F_XET enum=Xeqty1
    /auxin/05/eq/3/f  F32 F_XET
    /auxin/05/eq/3/g  F32 F_XET
    /auxin/05/eq/3/q  F32 F_XET
/auxin/05/eq/4  <CHEQ> n=0
    /auxin/05/eq/4/type  E32 F_XET enum=Xeqty1
    /auxin/05/eq/4/f  F32 F_XET
    /auxin/05/eq/4/g  F32 F_XET
    /auxin/05/eq/4/q  F32 F_XET
/auxin/05/mix  <CHMX> n=0
    /auxin/05/mix/on  E32 F_XET enum=OffOn
    /auxin/05/mix/fader  F32 F_XET
    /auxin/05/mix/st  E32 F_XET enum=OffOn
    /auxin/05/mix/pan  F32 F_XET
    /auxin/05/mix/mono  E32 F_XET enum=OffOn
    /auxin/05/mix/mlevel  F32 F_XET
/auxin/05/mix/01  <CHMO> n=0
    /auxin/05/mix/01/on  E32 F_XET enum=OffOn
    /auxin/05/mix/01/level  F32 F_XET
    /auxin/05/mix/01/pan  F32 F_XET
    /auxin/05/mix/01/type  E32 F_XET enum=Xmtype
    /auxin/05/mix/01/panFollow  E32 F_XET
/auxin/05/mix/02  <CHME> n=0
    /auxin/05/mix/02/on  E32 F_XET enum=OffOn
    /auxin/05/mix/02/level  F32 F_XET
/auxin/05/mix/03  <CHMO> n=0
    /auxin/05/mix/03/on  E32 F_XET enum=OffOn
    /auxin/05/mix/03/level  F32 F_XET
    /auxin/05/mix/03/pan  F32 F_XET
    /auxin/05/mix/03/type  E32 F_XET enum=Xmtype
    /auxin/05/mix/03/panFollow  E32 F_XET
/auxin/05/mix/04  <CHME> n=0
    /auxin/05/mix/04/on  E32 F_XET enum=OffOn
    /auxin/05/mix/04/level  F32 F_XET
/auxin/05/mix/05  <CHMO> n=0
    /auxin/05/mix/05/on  E32 F_XET enum=OffOn
    /auxin/05/mix/05/level  F32 F_XET
    /auxin/05/mix/05/pan  F32 F_XET
    /auxin/05/mix/05/type  E32 F_XET enum=Xmtype
    /auxin/05/mix/05/panFollow  E32 F_XET
/auxin/05/mix/06  <CHME> n=0
    /auxin/05/mix/06/on  E32 F_XET enum=OffOn
    /auxin/05/mix/06/level  F32 F_XET
/auxin/05/mix/07  <CHMO> n=0
    /auxin/05/mix/07/on  E32 F_XET enum=OffOn
    /auxin/05/mix/07/level  F32 F_XET
    /auxin/05/mix/07/pan  F32 F_XET
    /auxin/05/mix/07/type  E32 F_XET enum=Xmtype
    /auxin/05/mix/07/panFollow  E32 F_XET
/auxin/05/mix/08  <CHME> n=0
    /auxin/05/mix/08/on  E32 F_XET enum=OffOn
    /auxin/05/mix/08/level  F32 F_XET
/auxin/05/mix/09  <CHMO> n=0
    /auxin/05/mix/09/on  E32 F_XET enum=OffOn
    /auxin/05/mix/09/level  F32 F_XET
    /auxin/05/mix/09/pan  F32 F_XET
    /auxin/05/mix/09/type  E32 F_XET enum=Xmtype
    /auxin/05/mix/09/panFollow  E32 F_XET
/auxin/05/mix/10  <CHME> n=0
    /auxin/05/mix/10/on  E32 F_XET enum=OffOn
    /auxin/05/mix/10/level  F32 F_XET
/auxin/05/mix/11  <CHMO> n=0
    /auxin/05/mix/11/on  E32 F_XET enum=OffOn
    /auxin/05/mix/11/level  F32 F_XET
    /auxin/05/mix/11/pan  F32 F_XET
    /auxin/05/mix/11/type  E32 F_XET enum=Xmtype
    /auxin/05/mix/11/panFollow  E32 F_XET
/auxin/05/mix/12  <CHME> n=0
    /auxin/05/mix/12/on  E32 F_XET enum=OffOn
    /auxin/05/mix/12/level  F32 F_XET
/auxin/05/mix/13  <CHMO> n=0
    /auxin/05/mix/13/on  E32 F_XET enum=OffOn
    /auxin/05/mix/13/level  F32 F_XET
    /auxin/05/mix/13/pan  F32 F_XET
    /auxin/05/mix/13/type  E32 F_XET enum=Xmtype
    /auxin/05/mix/13/panFollow  E32 F_XET
/auxin/05/mix/14  <CHME> n=0
    /auxin/05/mix/14/on  E32 F_XET enum=OffOn
    /auxin/05/mix/14/level  F32 F_XET
/auxin/05/mix/15  <CHMO> n=0
    /auxin/05/mix/15/on  E32 F_XET enum=OffOn
    /auxin/05/mix/15/level  F32 F_XET
    /auxin/05/mix/15/pan  F32 F_XET
    /auxin/05/mix/15/type  E32 F_XET enum=Xmtype
    /auxin/05/mix/15/panFollow  E32 F_XET
/auxin/05/mix/16  <CHME> n=0
    /auxin/05/mix/16/on  E32 F_XET enum=OffOn
    /auxin/05/mix/16/level  F32 F_XET
/auxin/05/grp  <CHGRP> n=0
    /auxin/05/grp/dca  P32 F_XET
    /auxin/05/grp/mute  P32 F_XET
```

### Xauxin06 (X32Auxin.h, 114 entries)

```
/auxin  <CHCO> n=0
/auxin/06  <CHCO> n=0
/auxin/06/config  <CHCO> n=0
    /auxin/06/config/name  S32 F_XET
    /auxin/06/config/icon  I32 F_XET
    /auxin/06/config/color  E32 F_XET enum=Xcolors
    /auxin/06/config/source  I32 F_XET
/auxin/06/preamp  <AXPR> n=0
    /auxin/06/preamp/trim  F32 F_XET
    /auxin/06/preamp/invert  E32 F_XET enum=OffOn
/auxin/06/eq  <OFFON> n=1
    /auxin/06/eq/on  E32 F_XET enum=OffOn
/auxin/06/eq/1  <CHEQ> n=0
    /auxin/06/eq/1/type  E32 F_XET enum=Xeqty1
    /auxin/06/eq/1/f  F32 F_XET
    /auxin/06/eq/1/g  F32 F_XET
    /auxin/06/eq/1/q  F32 F_XET
/auxin/06/eq/2  <CHEQ> n=0
    /auxin/06/eq/2/type  E32 F_XET enum=Xeqty1
    /auxin/06/eq/2/f  F32 F_XET
    /auxin/06/eq/2/g  F32 F_XET
    /auxin/06/eq/2/q  F32 F_XET
/auxin/06/eq/3  <CHEQ> n=0
    /auxin/06/eq/3/type  E32 F_XET enum=Xeqty1
    /auxin/06/eq/3/f  F32 F_XET
    /auxin/06/eq/3/g  F32 F_XET
    /auxin/06/eq/3/q  F32 F_XET
/auxin/06/eq/4  <CHEQ> n=0
    /auxin/06/eq/4/type  E32 F_XET enum=Xeqty1
    /auxin/06/eq/4/f  F32 F_XET
    /auxin/06/eq/4/g  F32 F_XET
    /auxin/06/eq/4/q  F32 F_XET
/auxin/06/mix  <CHMX> n=0
    /auxin/06/mix/on  E32 F_XET enum=OffOn
    /auxin/06/mix/fader  F32 F_XET
    /auxin/06/mix/st  E32 F_XET enum=OffOn
    /auxin/06/mix/pan  F32 F_XET
    /auxin/06/mix/mono  E32 F_XET enum=OffOn
    /auxin/06/mix/mlevel  F32 F_XET
/auxin/06/mix/01  <CHMO> n=0
    /auxin/06/mix/01/on  E32 F_XET enum=OffOn
    /auxin/06/mix/01/level  F32 F_XET
    /auxin/06/mix/01/pan  F32 F_XET
    /auxin/06/mix/01/type  E32 F_XET enum=Xmtype
    /auxin/06/mix/01/panFollow  E32 F_XET
/auxin/06/mix/02  <CHME> n=0
    /auxin/06/mix/02/on  E32 F_XET enum=OffOn
    /auxin/06/mix/02/level  F32 F_XET
/auxin/06/mix/03  <CHMO> n=0
    /auxin/06/mix/03/on  E32 F_XET enum=OffOn
    /auxin/06/mix/03/level  F32 F_XET
    /auxin/06/mix/03/pan  F32 F_XET
    /auxin/06/mix/03/type  E32 F_XET enum=Xmtype
    /auxin/06/mix/03/panFollow  E32 F_XET
/auxin/06/mix/04  <CHME> n=0
    /auxin/06/mix/04/on  E32 F_XET enum=OffOn
    /auxin/06/mix/04/level  F32 F_XET
/auxin/06/mix/05  <CHMO> n=0
    /auxin/06/mix/05/on  E32 F_XET enum=OffOn
    /auxin/06/mix/05/level  F32 F_XET
    /auxin/06/mix/05/pan  F32 F_XET
    /auxin/06/mix/05/type  E32 F_XET enum=Xmtype
    /auxin/06/mix/05/panFollow  E32 F_XET
/auxin/06/mix/06  <CHME> n=0
    /auxin/06/mix/06/on  E32 F_XET enum=OffOn
    /auxin/06/mix/06/level  F32 F_XET
/auxin/06/mix/07  <CHMO> n=0
    /auxin/06/mix/07/on  E32 F_XET enum=OffOn
    /auxin/06/mix/07/level  F32 F_XET
    /auxin/06/mix/07/pan  F32 F_XET
    /auxin/06/mix/07/type  E32 F_XET enum=Xmtype
    /auxin/06/mix/07/panFollow  E32 F_XET
/auxin/06/mix/08  <CHME> n=0
    /auxin/06/mix/08/on  E32 F_XET enum=OffOn
    /auxin/06/mix/08/level  F32 F_XET
/auxin/06/mix/09  <CHMO> n=0
    /auxin/06/mix/09/on  E32 F_XET enum=OffOn
    /auxin/06/mix/09/level  F32 F_XET
    /auxin/06/mix/09/pan  F32 F_XET
    /auxin/06/mix/09/type  E32 F_XET enum=Xmtype
    /auxin/06/mix/09/panFollow  E32 F_XET
/auxin/06/mix/10  <CHME> n=0
    /auxin/06/mix/10/on  E32 F_XET enum=OffOn
    /auxin/06/mix/10/level  F32 F_XET
/auxin/06/mix/11  <CHMO> n=0
    /auxin/06/mix/11/on  E32 F_XET enum=OffOn
    /auxin/06/mix/11/level  F32 F_XET
    /auxin/06/mix/11/pan  F32 F_XET
    /auxin/06/mix/11/type  E32 F_XET enum=Xmtype
    /auxin/06/mix/11/panFollow  E32 F_XET
/auxin/06/mix/12  <CHME> n=0
    /auxin/06/mix/12/on  E32 F_XET enum=OffOn
    /auxin/06/mix/12/level  F32 F_XET
/auxin/06/mix/13  <CHMO> n=0
    /auxin/06/mix/13/on  E32 F_XET enum=OffOn
    /auxin/06/mix/13/level  F32 F_XET
    /auxin/06/mix/13/pan  F32 F_XET
    /auxin/06/mix/13/type  E32 F_XET enum=Xmtype
    /auxin/06/mix/13/panFollow  E32 F_XET
/auxin/06/mix/14  <CHME> n=0
    /auxin/06/mix/14/on  E32 F_XET enum=OffOn
    /auxin/06/mix/14/level  F32 F_XET
/auxin/06/mix/15  <CHMO> n=0
    /auxin/06/mix/15/on  E32 F_XET enum=OffOn
    /auxin/06/mix/15/level  F32 F_XET
    /auxin/06/mix/15/pan  F32 F_XET
    /auxin/06/mix/15/type  E32 F_XET enum=Xmtype
    /auxin/06/mix/15/panFollow  E32 F_XET
/auxin/06/mix/16  <CHME> n=0
    /auxin/06/mix/16/on  E32 F_XET enum=OffOn
    /auxin/06/mix/16/level  F32 F_XET
/auxin/06/grp  <CHGRP> n=0
    /auxin/06/grp/dca  P32 F_XET
    /auxin/06/grp/mute  P32 F_XET
```

### Xauxin07 (X32Auxin.h, 114 entries)

```
/auxin  <CHCO> n=0
/auxin/07  <CHCO> n=0
/auxin/07/config  <CHCO> n=0
    /auxin/07/config/name  S32 F_XET
    /auxin/07/config/icon  I32 F_XET
    /auxin/07/config/color  E32 F_XET enum=Xcolors
    /auxin/07/config/source  I32 F_XET
/auxin/07/preamp  <AXPR> n=0
    /auxin/07/preamp/trim  F32 F_XET
    /auxin/07/preamp/invert  E32 F_XET enum=OffOn
/auxin/07/eq  <OFFON> n=1
    /auxin/07/eq/on  E32 F_XET enum=OffOn
/auxin/07/eq/1  <CHEQ> n=0
    /auxin/07/eq/1/type  E32 F_XET enum=Xeqty1
    /auxin/07/eq/1/f  F32 F_XET
    /auxin/07/eq/1/g  F32 F_XET
    /auxin/07/eq/1/q  F32 F_XET
/auxin/07/eq/2  <CHEQ> n=0
    /auxin/07/eq/2/type  E32 F_XET enum=Xeqty1
    /auxin/07/eq/2/f  F32 F_XET
    /auxin/07/eq/2/g  F32 F_XET
    /auxin/07/eq/2/q  F32 F_XET
/auxin/07/eq/3  <CHEQ> n=0
    /auxin/07/eq/3/type  E32 F_XET enum=Xeqty1
    /auxin/07/eq/3/f  F32 F_XET
    /auxin/07/eq/3/g  F32 F_XET
    /auxin/07/eq/3/q  F32 F_XET
/auxin/07/eq/4  <CHEQ> n=0
    /auxin/07/eq/4/type  E32 F_XET enum=Xeqty1
    /auxin/07/eq/4/f  F32 F_XET
    /auxin/07/eq/4/g  F32 F_XET
    /auxin/07/eq/4/q  F32 F_XET
/auxin/07/mix  <CHMX> n=0
    /auxin/07/mix/on  E32 F_XET enum=OffOn
    /auxin/07/mix/fader  F32 F_XET
    /auxin/07/mix/st  E32 F_XET enum=OffOn
    /auxin/07/mix/pan  F32 F_XET
    /auxin/07/mix/mono  E32 F_XET enum=OffOn
    /auxin/07/mix/mlevel  F32 F_XET
/auxin/07/mix/01  <CHMO> n=0
    /auxin/07/mix/01/on  E32 F_XET enum=OffOn
    /auxin/07/mix/01/level  F32 F_XET
    /auxin/07/mix/01/pan  F32 F_XET
    /auxin/07/mix/01/type  E32 F_XET enum=Xmtype
    /auxin/07/mix/01/panFollow  E32 F_XET
/auxin/07/mix/02  <CHME> n=0
    /auxin/07/mix/02/on  E32 F_XET enum=OffOn
    /auxin/07/mix/02/level  F32 F_XET
/auxin/07/mix/03  <CHMO> n=0
    /auxin/07/mix/03/on  E32 F_XET enum=OffOn
    /auxin/07/mix/03/level  F32 F_XET
    /auxin/07/mix/03/pan  F32 F_XET
    /auxin/07/mix/03/type  E32 F_XET enum=Xmtype
    /auxin/07/mix/03/panFollow  E32 F_XET
/auxin/07/mix/04  <CHME> n=0
    /auxin/07/mix/04/on  E32 F_XET enum=OffOn
    /auxin/07/mix/04/level  F32 F_XET
/auxin/07/mix/05  <CHMO> n=0
    /auxin/07/mix/05/on  E32 F_XET enum=OffOn
    /auxin/07/mix/05/level  F32 F_XET
    /auxin/07/mix/05/pan  F32 F_XET
    /auxin/07/mix/05/type  E32 F_XET enum=Xmtype
    /auxin/07/mix/05/panFollow  E32 F_XET
/auxin/07/mix/06  <CHME> n=0
    /auxin/07/mix/06/on  E32 F_XET enum=OffOn
    /auxin/07/mix/06/level  F32 F_XET
/auxin/07/mix/07  <CHMO> n=0
    /auxin/07/mix/07/on  E32 F_XET enum=OffOn
    /auxin/07/mix/07/level  F32 F_XET
    /auxin/07/mix/07/pan  F32 F_XET
    /auxin/07/mix/07/type  E32 F_XET enum=Xmtype
    /auxin/07/mix/07/panFollow  E32 F_XET
/auxin/07/mix/08  <CHME> n=0
    /auxin/07/mix/08/on  E32 F_XET enum=OffOn
    /auxin/07/mix/08/level  F32 F_XET
/auxin/07/mix/09  <CHMO> n=0
    /auxin/07/mix/09/on  E32 F_XET enum=OffOn
    /auxin/07/mix/09/level  F32 F_XET
    /auxin/07/mix/09/pan  F32 F_XET
    /auxin/07/mix/09/type  E32 F_XET enum=Xmtype
    /auxin/07/mix/09/panFollow  E32 F_XET
/auxin/07/mix/10  <CHME> n=0
    /auxin/07/mix/10/on  E32 F_XET enum=OffOn
    /auxin/07/mix/10/level  F32 F_XET
/auxin/07/mix/11  <CHMO> n=0
    /auxin/07/mix/11/on  E32 F_XET enum=OffOn
    /auxin/07/mix/11/level  F32 F_XET
    /auxin/07/mix/11/pan  F32 F_XET
    /auxin/07/mix/11/type  E32 F_XET enum=Xmtype
    /auxin/07/mix/11/panFollow  E32 F_XET
/auxin/07/mix/12  <CHME> n=0
    /auxin/07/mix/12/on  E32 F_XET enum=OffOn
    /auxin/07/mix/12/level  F32 F_XET
/auxin/07/mix/13  <CHMO> n=0
    /auxin/07/mix/13/on  E32 F_XET enum=OffOn
    /auxin/07/mix/13/level  F32 F_XET
    /auxin/07/mix/13/pan  F32 F_XET
    /auxin/07/mix/13/type  E32 F_XET enum=Xmtype
    /auxin/07/mix/13/panFollow  E32 F_XET
/auxin/07/mix/14  <CHME> n=0
    /auxin/07/mix/14/on  E32 F_XET enum=OffOn
    /auxin/07/mix/14/level  F32 F_XET
/auxin/07/mix/15  <CHMO> n=0
    /auxin/07/mix/15/on  E32 F_XET enum=OffOn
    /auxin/07/mix/15/level  F32 F_XET
    /auxin/07/mix/15/pan  F32 F_XET
    /auxin/07/mix/15/type  E32 F_XET enum=Xmtype
    /auxin/07/mix/15/panFollow  E32 F_XET
/auxin/07/mix/16  <CHME> n=0
    /auxin/07/mix/16/on  E32 F_XET enum=OffOn
    /auxin/07/mix/16/level  F32 F_XET
/auxin/07/grp  <CHGRP> n=0
    /auxin/07/grp/dca  P32 F_XET
    /auxin/07/grp/mute  P32 F_XET
```

### Xauxin08 (X32Auxin.h, 114 entries)

```
/auxin  <CHCO> n=0
/auxin/08  <CHCO> n=0
/auxin/08/config  <CHCO> n=0
    /auxin/08/config/name  S32 F_XET
    /auxin/08/config/icon  I32 F_XET
    /auxin/08/config/color  E32 F_XET enum=Xcolors
    /auxin/08/config/source  I32 F_XET
/auxin/08/preamp  <AXPR> n=0
    /auxin/08/preamp/trim  F32 F_XET
    /auxin/08/preamp/invert  E32 F_XET enum=OffOn
/auxin/08/eq  <OFFON> n=1
    /auxin/08/eq/on  E32 F_XET enum=OffOn
/auxin/08/eq/1  <CHEQ> n=0
    /auxin/08/eq/1/type  E32 F_XET enum=Xeqty1
    /auxin/08/eq/1/f  F32 F_XET
    /auxin/08/eq/1/g  F32 F_XET
    /auxin/08/eq/1/q  F32 F_XET
/auxin/08/eq/2  <CHEQ> n=0
    /auxin/08/eq/2/type  E32 F_XET enum=Xeqty1
    /auxin/08/eq/2/f  F32 F_XET
    /auxin/08/eq/2/g  F32 F_XET
    /auxin/08/eq/2/q  F32 F_XET
/auxin/08/eq/3  <CHEQ> n=0
    /auxin/08/eq/3/type  E32 F_XET enum=Xeqty1
    /auxin/08/eq/3/f  F32 F_XET
    /auxin/08/eq/3/g  F32 F_XET
    /auxin/08/eq/3/q  F32 F_XET
/auxin/08/eq/4  <CHEQ> n=0
    /auxin/08/eq/4/type  E32 F_XET enum=Xeqty1
    /auxin/08/eq/4/f  F32 F_XET
    /auxin/08/eq/4/g  F32 F_XET
    /auxin/08/eq/4/q  F32 F_XET
/auxin/08/mix  <CHMX> n=0
    /auxin/08/mix/on  E32 F_XET enum=OffOn
    /auxin/08/mix/fader  F32 F_XET
    /auxin/08/mix/st  E32 F_XET enum=OffOn
    /auxin/08/mix/pan  F32 F_XET
    /auxin/08/mix/mono  E32 F_XET enum=OffOn
    /auxin/08/mix/mlevel  F32 F_XET
/auxin/08/mix/01  <CHMO> n=0
    /auxin/08/mix/01/on  E32 F_XET enum=OffOn
    /auxin/08/mix/01/level  F32 F_XET
    /auxin/08/mix/01/pan  F32 F_XET
    /auxin/08/mix/01/type  E32 F_XET enum=Xmtype
    /auxin/08/mix/01/panFollow  E32 F_XET
/auxin/08/mix/02  <CHME> n=0
    /auxin/08/mix/02/on  E32 F_XET enum=OffOn
    /auxin/08/mix/02/level  F32 F_XET
/auxin/08/mix/03  <CHMO> n=0
    /auxin/08/mix/03/on  E32 F_XET enum=OffOn
    /auxin/08/mix/03/level  F32 F_XET
    /auxin/08/mix/03/pan  F32 F_XET
    /auxin/08/mix/03/type  E32 F_XET enum=Xmtype
    /auxin/08/mix/03/panFollow  E32 F_XET
/auxin/08/mix/04  <CHME> n=0
    /auxin/08/mix/04/on  E32 F_XET enum=OffOn
    /auxin/08/mix/04/level  F32 F_XET
/auxin/08/mix/05  <CHMO> n=0
    /auxin/08/mix/05/on  E32 F_XET enum=OffOn
    /auxin/08/mix/05/level  F32 F_XET
    /auxin/08/mix/05/pan  F32 F_XET
    /auxin/08/mix/05/type  E32 F_XET enum=Xmtype
    /auxin/08/mix/05/panFollow  E32 F_XET
/auxin/08/mix/06  <CHME> n=0
    /auxin/08/mix/06/on  E32 F_XET enum=OffOn
    /auxin/08/mix/06/level  F32 F_XET
/auxin/08/mix/07  <CHMO> n=0
    /auxin/08/mix/07/on  E32 F_XET enum=OffOn
    /auxin/08/mix/07/level  F32 F_XET
    /auxin/08/mix/07/pan  F32 F_XET
    /auxin/08/mix/07/type  E32 F_XET enum=Xmtype
    /auxin/08/mix/07/panFollow  E32 F_XET
/auxin/08/mix/08  <CHME> n=0
    /auxin/08/mix/08/on  E32 F_XET enum=OffOn
    /auxin/08/mix/08/level  F32 F_XET
/auxin/08/mix/09  <CHMO> n=0
    /auxin/08/mix/09/on  E32 F_XET enum=OffOn
    /auxin/08/mix/09/level  F32 F_XET
    /auxin/08/mix/09/pan  F32 F_XET
    /auxin/08/mix/09/type  E32 F_XET enum=Xmtype
    /auxin/08/mix/09/panFollow  E32 F_XET
/auxin/08/mix/10  <CHME> n=0
    /auxin/08/mix/10/on  E32 F_XET enum=OffOn
    /auxin/08/mix/10/level  F32 F_XET
/auxin/08/mix/11  <CHMO> n=0
    /auxin/08/mix/11/on  E32 F_XET enum=OffOn
    /auxin/08/mix/11/level  F32 F_XET
    /auxin/08/mix/11/pan  F32 F_XET
    /auxin/08/mix/11/type  E32 F_XET enum=Xmtype
    /auxin/08/mix/11/panFollow  E32 F_XET
/auxin/08/mix/12  <CHME> n=0
    /auxin/08/mix/12/on  E32 F_XET enum=OffOn
    /auxin/08/mix/12/level  F32 F_XET
/auxin/08/mix/13  <CHMO> n=0
    /auxin/08/mix/13/on  E32 F_XET enum=OffOn
    /auxin/08/mix/13/level  F32 F_XET
    /auxin/08/mix/13/pan  F32 F_XET
    /auxin/08/mix/13/type  E32 F_XET enum=Xmtype
    /auxin/08/mix/13/panFollow  E32 F_XET
/auxin/08/mix/14  <CHME> n=0
    /auxin/08/mix/14/on  E32 F_XET enum=OffOn
    /auxin/08/mix/14/level  F32 F_XET
/auxin/08/mix/15  <CHMO> n=0
    /auxin/08/mix/15/on  E32 F_XET enum=OffOn
    /auxin/08/mix/15/level  F32 F_XET
    /auxin/08/mix/15/pan  F32 F_XET
    /auxin/08/mix/15/type  E32 F_XET enum=Xmtype
    /auxin/08/mix/15/panFollow  E32 F_XET
/auxin/08/mix/16  <CHME> n=0
    /auxin/08/mix/16/on  E32 F_XET enum=OffOn
    /auxin/08/mix/16/level  F32 F_XET
/auxin/08/grp  <CHGRP> n=0
    /auxin/08/grp/dca  P32 F_XET
    /auxin/08/grp/mute  P32 F_XET
```

### Xfxrtn01 (X32Fxrtn.h, 110 entries)

```
/fxrtn  <BSCO> n=0
/fxrtn/01  <BSCO> n=0
/fxrtn/01/config  <BSCO> n=0
    /fxrtn/01/config/name  S32 F_XET
    /fxrtn/01/config/icon  I32 F_XET
    /fxrtn/01/config/color  E32 F_XET enum=Xcolors
/fxrtn/01/eq  <OFFON> n=1
    /fxrtn/01/eq/on  E32 F_XET enum=OffOn
/fxrtn/01/eq/1  <CHEQ> n=0
    /fxrtn/01/eq/1/type  E32 F_XET enum=Xeqty1
    /fxrtn/01/eq/1/f  F32 F_XET
    /fxrtn/01/eq/1/g  F32 F_XET
    /fxrtn/01/eq/1/q  F32 F_XET
/fxrtn/01/eq/2  <CHEQ> n=0
    /fxrtn/01/eq/2/type  E32 F_XET enum=Xeqty1
    /fxrtn/01/eq/2/f  F32 F_XET
    /fxrtn/01/eq/2/g  F32 F_XET
    /fxrtn/01/eq/2/q  F32 F_XET
/fxrtn/01/eq/3  <CHEQ> n=0
    /fxrtn/01/eq/3/type  E32 F_XET enum=Xeqty1
    /fxrtn/01/eq/3/f  F32 F_XET
    /fxrtn/01/eq/3/g  F32 F_XET
    /fxrtn/01/eq/3/q  F32 F_XET
/fxrtn/01/eq/4  <CHEQ> n=0
    /fxrtn/01/eq/4/type  E32 F_XET enum=Xeqty1
    /fxrtn/01/eq/4/f  F32 F_XET
    /fxrtn/01/eq/4/g  F32 F_XET
    /fxrtn/01/eq/4/q  F32 F_XET
/fxrtn/01/mix  <CHMX> n=0
    /fxrtn/01/mix/on  E32 F_XET enum=OffOn
    /fxrtn/01/mix/fader  F32 F_XET
    /fxrtn/01/mix/st  E32 F_XET enum=OffOn
    /fxrtn/01/mix/pan  F32 F_XET
    /fxrtn/01/mix/mono  E32 F_XET enum=OffOn
    /fxrtn/01/mix/mlevel  F32 F_XET
/fxrtn/01/mix/01  <CHMO> n=0
    /fxrtn/01/mix/01/on  E32 F_XET enum=OffOn
    /fxrtn/01/mix/01/level  F32 F_XET
    /fxrtn/01/mix/01/pan  F32 F_XET
    /fxrtn/01/mix/01/type  E32 F_XET enum=Xmtype
    /fxrtn/01/mix/01/panFollow  E32 F_XET
/fxrtn/01/mix/02  <CHME> n=0
    /fxrtn/01/mix/02/on  E32 F_XET enum=OffOn
    /fxrtn/01/mix/02/level  F32 F_XET
/fxrtn/01/mix/03  <CHMO> n=0
    /fxrtn/01/mix/03/on  E32 F_XET enum=OffOn
    /fxrtn/01/mix/03/level  F32 F_XET
    /fxrtn/01/mix/03/pan  F32 F_XET
    /fxrtn/01/mix/03/type  E32 F_XET enum=Xmtype
    /fxrtn/01/mix/03/panFollow  E32 F_XET
/fxrtn/01/mix/04  <CHME> n=0
    /fxrtn/01/mix/04/on  E32 F_XET enum=OffOn
    /fxrtn/01/mix/04/level  F32 F_XET
/fxrtn/01/mix/05  <CHMO> n=0
    /fxrtn/01/mix/05/on  E32 F_XET enum=OffOn
    /fxrtn/01/mix/05/level  F32 F_XET
    /fxrtn/01/mix/05/pan  F32 F_XET
    /fxrtn/01/mix/05/type  E32 F_XET enum=Xmtype
    /fxrtn/01/mix/05/panFollow  E32 F_XET
/fxrtn/01/mix/06  <CHME> n=0
    /fxrtn/01/mix/06/on  E32 F_XET enum=OffOn
    /fxrtn/01/mix/06/level  F32 F_XET
/fxrtn/01/mix/07  <CHMO> n=0
    /fxrtn/01/mix/07/on  E32 F_XET enum=OffOn
    /fxrtn/01/mix/07/level  F32 F_XET
    /fxrtn/01/mix/07/pan  F32 F_XET
    /fxrtn/01/mix/07/type  E32 F_XET enum=Xmtype
    /fxrtn/01/mix/07/panFollow  E32 F_XET
/fxrtn/01/mix/08  <CHME> n=0
    /fxrtn/01/mix/08/on  E32 F_XET enum=OffOn
    /fxrtn/01/mix/08/level  F32 F_XET
/fxrtn/01/mix/09  <CHMO> n=0
    /fxrtn/01/mix/09/on  E32 F_XET enum=OffOn
    /fxrtn/01/mix/09/level  F32 F_XET
    /fxrtn/01/mix/09/pan  F32 F_XET
    /fxrtn/01/mix/09/type  E32 F_XET enum=Xmtype
    /fxrtn/01/mix/09/panFollow  E32 F_XET
/fxrtn/01/mix/10  <CHME> n=0
    /fxrtn/01/mix/10/on  E32 F_XET enum=OffOn
    /fxrtn/01/mix/10/level  F32 F_XET
/fxrtn/01/mix/11  <CHMO> n=0
    /fxrtn/01/mix/11/on  E32 F_XET enum=OffOn
    /fxrtn/01/mix/11/level  F32 F_XET
    /fxrtn/01/mix/11/pan  F32 F_XET
    /fxrtn/01/mix/11/type  E32 F_XET enum=Xmtype
    /fxrtn/01/mix/11/panFollow  E32 F_XET
/fxrtn/01/mix/12  <CHME> n=0
    /fxrtn/01/mix/12/on  E32 F_XET enum=OffOn
    /fxrtn/01/mix/12/level  F32 F_XET
/fxrtn/01/mix/13  <CHMO> n=0
    /fxrtn/01/mix/13/on  E32 F_XET enum=OffOn
    /fxrtn/01/mix/13/level  F32 F_XET
    /fxrtn/01/mix/13/pan  F32 F_XET
    /fxrtn/01/mix/13/type  E32 F_XET enum=Xmtype
    /fxrtn/01/mix/13/panFollow  E32 F_XET
/fxrtn/01/mix/14  <CHME> n=0
    /fxrtn/01/mix/14/on  E32 F_XET enum=OffOn
    /fxrtn/01/mix/14/level  F32 F_XET
/fxrtn/01/mix/15  <CHMO> n=0
    /fxrtn/01/mix/15/on  E32 F_XET enum=OffOn
    /fxrtn/01/mix/15/level  F32 F_XET
    /fxrtn/01/mix/15/pan  F32 F_XET
    /fxrtn/01/mix/15/type  E32 F_XET enum=Xmtype
    /fxrtn/01/mix/15/panFollow  E32 F_XET
/fxrtn/01/mix/16  <CHME> n=0
    /fxrtn/01/mix/16/on  E32 F_XET enum=OffOn
    /fxrtn/01/mix/16/level  F32 F_XET
/fxrtn/01/grp  <CHGRP> n=0
    /fxrtn/01/grp/dca  P32 F_XET
    /fxrtn/01/grp/mute  P32 F_XET
```

### Xfxrtn02 (X32Fxrtn.h, 110 entries)

```
/fxrtn  <BSCO> n=0
/fxrtn/02  <BSCO> n=0
/fxrtn/02/config  <BSCO> n=0
    /fxrtn/02/config/name  S32 F_XET
    /fxrtn/02/config/icon  I32 F_XET
    /fxrtn/02/config/color  E32 F_XET enum=Xcolors
/fxrtn/02/eq  <OFFON> n=1
    /fxrtn/02/eq/on  E32 F_XET enum=OffOn
/fxrtn/02/eq/1  <CHEQ> n=0
    /fxrtn/02/eq/1/type  E32 F_XET enum=Xeqty1
    /fxrtn/02/eq/1/f  F32 F_XET
    /fxrtn/02/eq/1/g  F32 F_XET
    /fxrtn/02/eq/1/q  F32 F_XET
/fxrtn/02/eq/2  <CHEQ> n=0
    /fxrtn/02/eq/2/type  E32 F_XET enum=Xeqty1
    /fxrtn/02/eq/2/f  F32 F_XET
    /fxrtn/02/eq/2/g  F32 F_XET
    /fxrtn/02/eq/2/q  F32 F_XET
/fxrtn/02/eq/3  <CHEQ> n=0
    /fxrtn/02/eq/3/type  E32 F_XET enum=Xeqty1
    /fxrtn/02/eq/3/f  F32 F_XET
    /fxrtn/02/eq/3/g  F32 F_XET
    /fxrtn/02/eq/3/q  F32 F_XET
/fxrtn/02/eq/4  <CHEQ> n=0
    /fxrtn/02/eq/4/type  E32 F_XET enum=Xeqty1
    /fxrtn/02/eq/4/f  F32 F_XET
    /fxrtn/02/eq/4/g  F32 F_XET
    /fxrtn/02/eq/4/q  F32 F_XET
/fxrtn/02/mix  <CHMX> n=0
    /fxrtn/02/mix/on  E32 F_XET enum=OffOn
    /fxrtn/02/mix/fader  F32 F_XET
    /fxrtn/02/mix/st  E32 F_XET enum=OffOn
    /fxrtn/02/mix/pan  F32 F_XET
    /fxrtn/02/mix/mono  E32 F_XET enum=OffOn
    /fxrtn/02/mix/mlevel  F32 F_XET
/fxrtn/02/mix/01  <CHMO> n=0
    /fxrtn/02/mix/01/on  E32 F_XET enum=OffOn
    /fxrtn/02/mix/01/level  F32 F_XET
    /fxrtn/02/mix/01/pan  F32 F_XET
    /fxrtn/02/mix/01/type  E32 F_XET enum=Xmtype
    /fxrtn/02/mix/01/panFollow  E32 F_XET
/fxrtn/02/mix/02  <CHME> n=0
    /fxrtn/02/mix/02/on  E32 F_XET enum=OffOn
    /fxrtn/02/mix/02/level  F32 F_XET
/fxrtn/02/mix/03  <CHMO> n=0
    /fxrtn/02/mix/03/on  E32 F_XET enum=OffOn
    /fxrtn/02/mix/03/level  F32 F_XET
    /fxrtn/02/mix/03/pan  F32 F_XET
    /fxrtn/02/mix/03/type  E32 F_XET enum=Xmtype
    /fxrtn/02/mix/03/panFollow  E32 F_XET
/fxrtn/02/mix/04  <CHME> n=0
    /fxrtn/02/mix/04/on  E32 F_XET enum=OffOn
    /fxrtn/02/mix/04/level  F32 F_XET
/fxrtn/02/mix/05  <CHMO> n=0
    /fxrtn/02/mix/05/on  E32 F_XET enum=OffOn
    /fxrtn/02/mix/05/level  F32 F_XET
    /fxrtn/02/mix/05/pan  F32 F_XET
    /fxrtn/02/mix/05/type  E32 F_XET enum=Xmtype
    /fxrtn/02/mix/05/panFollow  E32 F_XET
/fxrtn/02/mix/06  <CHME> n=0
    /fxrtn/02/mix/06/on  E32 F_XET enum=OffOn
    /fxrtn/02/mix/06/level  F32 F_XET
/fxrtn/02/mix/07  <CHMO> n=0
    /fxrtn/02/mix/07/on  E32 F_XET enum=OffOn
    /fxrtn/02/mix/07/level  F32 F_XET
    /fxrtn/02/mix/07/pan  F32 F_XET
    /fxrtn/02/mix/07/type  E32 F_XET enum=Xmtype
    /fxrtn/02/mix/07/panFollow  E32 F_XET
/fxrtn/02/mix/08  <CHME> n=0
    /fxrtn/02/mix/08/on  E32 F_XET enum=OffOn
    /fxrtn/02/mix/08/level  F32 F_XET
/fxrtn/02/mix/09  <CHMO> n=0
    /fxrtn/02/mix/09/on  E32 F_XET enum=OffOn
    /fxrtn/02/mix/09/level  F32 F_XET
    /fxrtn/02/mix/09/pan  F32 F_XET
    /fxrtn/02/mix/09/type  E32 F_XET enum=Xmtype
    /fxrtn/02/mix/09/panFollow  E32 F_XET
/fxrtn/02/mix/10  <CHME> n=0
    /fxrtn/02/mix/10/on  E32 F_XET enum=OffOn
    /fxrtn/02/mix/10/level  F32 F_XET
/fxrtn/02/mix/11  <CHMO> n=0
    /fxrtn/02/mix/11/on  E32 F_XET enum=OffOn
    /fxrtn/02/mix/11/level  F32 F_XET
    /fxrtn/02/mix/11/pan  F32 F_XET
    /fxrtn/02/mix/11/type  E32 F_XET enum=Xmtype
    /fxrtn/02/mix/11/panFollow  E32 F_XET
/fxrtn/02/mix/12  <CHME> n=0
    /fxrtn/02/mix/12/on  E32 F_XET enum=OffOn
    /fxrtn/02/mix/12/level  F32 F_XET
/fxrtn/02/mix/13  <CHMO> n=0
    /fxrtn/02/mix/13/on  E32 F_XET enum=OffOn
    /fxrtn/02/mix/13/level  F32 F_XET
    /fxrtn/02/mix/13/pan  F32 F_XET
    /fxrtn/02/mix/13/type  E32 F_XET enum=Xmtype
    /fxrtn/02/mix/13/panFollow  E32 F_XET
/fxrtn/02/mix/14  <CHME> n=0
    /fxrtn/02/mix/14/on  E32 F_XET enum=OffOn
    /fxrtn/02/mix/14/level  F32 F_XET
/fxrtn/02/mix/15  <CHMO> n=0
    /fxrtn/02/mix/15/on  E32 F_XET enum=OffOn
    /fxrtn/02/mix/15/level  F32 F_XET
    /fxrtn/02/mix/15/pan  F32 F_XET
    /fxrtn/02/mix/15/type  E32 F_XET enum=Xmtype
    /fxrtn/02/mix/15/panFollow  E32 F_XET
/fxrtn/02/mix/16  <CHME> n=0
    /fxrtn/02/mix/16/on  E32 F_XET enum=OffOn
    /fxrtn/02/mix/16/level  F32 F_XET
/fxrtn/02/grp  <CHGRP> n=0
    /fxrtn/02/grp/dca  P32 F_XET
    /fxrtn/02/grp/mute  P32 F_XET
```

### Xfxrtn03 (X32Fxrtn.h, 110 entries)

```
/fxrtn  <BSCO> n=0
/fxrtn/03  <BSCO> n=0
/fxrtn/03/config  <BSCO> n=0
    /fxrtn/03/config/name  S32 F_XET
    /fxrtn/03/config/icon  I32 F_XET
    /fxrtn/03/config/color  E32 F_XET enum=Xcolors
/fxrtn/03/eq  <OFFON> n=1
    /fxrtn/03/eq/on  E32 F_XET enum=OffOn
/fxrtn/03/eq/1  <CHEQ> n=0
    /fxrtn/03/eq/1/type  E32 F_XET enum=Xeqty1
    /fxrtn/03/eq/1/f  F32 F_XET
    /fxrtn/03/eq/1/g  F32 F_XET
    /fxrtn/03/eq/1/q  F32 F_XET
/fxrtn/03/eq/2  <CHEQ> n=0
    /fxrtn/03/eq/2/type  E32 F_XET enum=Xeqty1
    /fxrtn/03/eq/2/f  F32 F_XET
    /fxrtn/03/eq/2/g  F32 F_XET
    /fxrtn/03/eq/2/q  F32 F_XET
/fxrtn/03/eq/3  <CHEQ> n=0
    /fxrtn/03/eq/3/type  E32 F_XET enum=Xeqty1
    /fxrtn/03/eq/3/f  F32 F_XET
    /fxrtn/03/eq/3/g  F32 F_XET
    /fxrtn/03/eq/3/q  F32 F_XET
/fxrtn/03/eq/4  <CHEQ> n=0
    /fxrtn/03/eq/4/type  E32 F_XET enum=Xeqty1
    /fxrtn/03/eq/4/f  F32 F_XET
    /fxrtn/03/eq/4/g  F32 F_XET
    /fxrtn/03/eq/4/q  F32 F_XET
/fxrtn/03/mix  <CHMX> n=0
    /fxrtn/03/mix/on  E32 F_XET enum=OffOn
    /fxrtn/03/mix/fader  F32 F_XET
    /fxrtn/03/mix/st  E32 F_XET enum=OffOn
    /fxrtn/03/mix/pan  F32 F_XET
    /fxrtn/03/mix/mono  E32 F_XET enum=OffOn
    /fxrtn/03/mix/mlevel  F32 F_XET
/fxrtn/03/mix/01  <CHMO> n=0
    /fxrtn/03/mix/01/on  E32 F_XET enum=OffOn
    /fxrtn/03/mix/01/level  F32 F_XET
    /fxrtn/03/mix/01/pan  F32 F_XET
    /fxrtn/03/mix/01/type  E32 F_XET enum=Xmtype
    /fxrtn/03/mix/01/panFollow  E32 F_XET
/fxrtn/03/mix/02  <CHME> n=0
    /fxrtn/03/mix/02/on  E32 F_XET enum=OffOn
    /fxrtn/03/mix/02/level  F32 F_XET
/fxrtn/03/mix/03  <CHMO> n=0
    /fxrtn/03/mix/03/on  E32 F_XET enum=OffOn
    /fxrtn/03/mix/03/level  F32 F_XET
    /fxrtn/03/mix/03/pan  F32 F_XET
    /fxrtn/03/mix/03/type  E32 F_XET enum=Xmtype
    /fxrtn/03/mix/03/panFollow  E32 F_XET
/fxrtn/03/mix/04  <CHME> n=0
    /fxrtn/03/mix/04/on  E32 F_XET enum=OffOn
    /fxrtn/03/mix/04/level  F32 F_XET
/fxrtn/03/mix/05  <CHMO> n=0
    /fxrtn/03/mix/05/on  E32 F_XET enum=OffOn
    /fxrtn/03/mix/05/level  F32 F_XET
    /fxrtn/03/mix/05/pan  F32 F_XET
    /fxrtn/03/mix/05/type  E32 F_XET enum=Xmtype
    /fxrtn/03/mix/05/panFollow  E32 F_XET
/fxrtn/03/mix/06  <CHME> n=0
    /fxrtn/03/mix/06/on  E32 F_XET enum=OffOn
    /fxrtn/03/mix/06/level  F32 F_XET
/fxrtn/03/mix/07  <CHMO> n=0
    /fxrtn/03/mix/07/on  E32 F_XET enum=OffOn
    /fxrtn/03/mix/07/level  F32 F_XET
    /fxrtn/03/mix/07/pan  F32 F_XET
    /fxrtn/03/mix/07/type  E32 F_XET enum=Xmtype
    /fxrtn/03/mix/07/panFollow  E32 F_XET
/fxrtn/03/mix/08  <CHME> n=0
    /fxrtn/03/mix/08/on  E32 F_XET enum=OffOn
    /fxrtn/03/mix/08/level  F32 F_XET
/fxrtn/03/mix/09  <CHMO> n=0
    /fxrtn/03/mix/09/on  E32 F_XET enum=OffOn
    /fxrtn/03/mix/09/level  F32 F_XET
    /fxrtn/03/mix/09/pan  F32 F_XET
    /fxrtn/03/mix/09/type  E32 F_XET enum=Xmtype
    /fxrtn/03/mix/09/panFollow  E32 F_XET
/fxrtn/03/mix/10  <CHME> n=0
    /fxrtn/03/mix/10/on  E32 F_XET enum=OffOn
    /fxrtn/03/mix/10/level  F32 F_XET
/fxrtn/03/mix/11  <CHMO> n=0
    /fxrtn/03/mix/11/on  E32 F_XET enum=OffOn
    /fxrtn/03/mix/11/level  F32 F_XET
    /fxrtn/03/mix/11/pan  F32 F_XET
    /fxrtn/03/mix/11/type  E32 F_XET enum=Xmtype
    /fxrtn/03/mix/11/panFollow  E32 F_XET
/fxrtn/03/mix/12  <CHME> n=0
    /fxrtn/03/mix/12/on  E32 F_XET enum=OffOn
    /fxrtn/03/mix/12/level  F32 F_XET
/fxrtn/03/mix/13  <CHMO> n=0
    /fxrtn/03/mix/13/on  E32 F_XET enum=OffOn
    /fxrtn/03/mix/13/level  F32 F_XET
    /fxrtn/03/mix/13/pan  F32 F_XET
    /fxrtn/03/mix/13/type  E32 F_XET enum=Xmtype
    /fxrtn/03/mix/13/panFollow  E32 F_XET
/fxrtn/03/mix/14  <CHME> n=0
    /fxrtn/03/mix/14/on  E32 F_XET enum=OffOn
    /fxrtn/03/mix/14/level  F32 F_XET
/fxrtn/03/mix/15  <CHMO> n=0
    /fxrtn/03/mix/15/on  E32 F_XET enum=OffOn
    /fxrtn/03/mix/15/level  F32 F_XET
    /fxrtn/03/mix/15/pan  F32 F_XET
    /fxrtn/03/mix/15/type  E32 F_XET enum=Xmtype
    /fxrtn/03/mix/15/panFollow  E32 F_XET
/fxrtn/03/mix/16  <CHME> n=0
    /fxrtn/03/mix/16/on  E32 F_XET enum=OffOn
    /fxrtn/03/mix/16/level  F32 F_XET
/fxrtn/03/grp  <CHGRP> n=0
    /fxrtn/03/grp/dca  P32 F_XET
    /fxrtn/03/grp/mute  P32 F_XET
```

### Xfxrtn04 (X32Fxrtn.h, 110 entries)

```
/fxrtn  <BSCO> n=0
/fxrtn/04  <BSCO> n=0
/fxrtn/04/config  <BSCO> n=0
    /fxrtn/04/config/name  S32 F_XET
    /fxrtn/04/config/icon  I32 F_XET
    /fxrtn/04/config/color  E32 F_XET enum=Xcolors
/fxrtn/04/eq  <OFFON> n=1
    /fxrtn/04/eq/on  E32 F_XET enum=OffOn
/fxrtn/04/eq/1  <CHEQ> n=0
    /fxrtn/04/eq/1/type  E32 F_XET enum=Xeqty1
    /fxrtn/04/eq/1/f  F32 F_XET
    /fxrtn/04/eq/1/g  F32 F_XET
    /fxrtn/04/eq/1/q  F32 F_XET
/fxrtn/04/eq/2  <CHEQ> n=0
    /fxrtn/04/eq/2/type  E32 F_XET enum=Xeqty1
    /fxrtn/04/eq/2/f  F32 F_XET
    /fxrtn/04/eq/2/g  F32 F_XET
    /fxrtn/04/eq/2/q  F32 F_XET
/fxrtn/04/eq/3  <CHEQ> n=0
    /fxrtn/04/eq/3/type  E32 F_XET enum=Xeqty1
    /fxrtn/04/eq/3/f  F32 F_XET
    /fxrtn/04/eq/3/g  F32 F_XET
    /fxrtn/04/eq/3/q  F32 F_XET
/fxrtn/04/eq/4  <CHEQ> n=0
    /fxrtn/04/eq/4/type  E32 F_XET enum=Xeqty1
    /fxrtn/04/eq/4/f  F32 F_XET
    /fxrtn/04/eq/4/g  F32 F_XET
    /fxrtn/04/eq/4/q  F32 F_XET
/fxrtn/04/mix  <CHMX> n=0
    /fxrtn/04/mix/on  E32 F_XET enum=OffOn
    /fxrtn/04/mix/fader  F32 F_XET
    /fxrtn/04/mix/st  E32 F_XET enum=OffOn
    /fxrtn/04/mix/pan  F32 F_XET
    /fxrtn/04/mix/mono  E32 F_XET enum=OffOn
    /fxrtn/04/mix/mlevel  F32 F_XET
/fxrtn/04/mix/01  <CHMO> n=0
    /fxrtn/04/mix/01/on  E32 F_XET enum=OffOn
    /fxrtn/04/mix/01/level  F32 F_XET
    /fxrtn/04/mix/01/pan  F32 F_XET
    /fxrtn/04/mix/01/type  E32 F_XET enum=Xmtype
    /fxrtn/04/mix/01/panFollow  E32 F_XET
/fxrtn/04/mix/02  <CHME> n=0
    /fxrtn/04/mix/02/on  E32 F_XET enum=OffOn
    /fxrtn/04/mix/02/level  F32 F_XET
/fxrtn/04/mix/03  <CHMO> n=0
    /fxrtn/04/mix/03/on  E32 F_XET enum=OffOn
    /fxrtn/04/mix/03/level  F32 F_XET
    /fxrtn/04/mix/03/pan  F32 F_XET
    /fxrtn/04/mix/03/type  E32 F_XET enum=Xmtype
    /fxrtn/04/mix/03/panFollow  E32 F_XET
/fxrtn/04/mix/04  <CHME> n=0
    /fxrtn/04/mix/04/on  E32 F_XET enum=OffOn
    /fxrtn/04/mix/04/level  F32 F_XET
/fxrtn/04/mix/05  <CHMO> n=0
    /fxrtn/04/mix/05/on  E32 F_XET enum=OffOn
    /fxrtn/04/mix/05/level  F32 F_XET
    /fxrtn/04/mix/05/pan  F32 F_XET
    /fxrtn/04/mix/05/type  E32 F_XET enum=Xmtype
    /fxrtn/04/mix/05/panFollow  E32 F_XET
/fxrtn/04/mix/06  <CHME> n=0
    /fxrtn/04/mix/06/on  E32 F_XET enum=OffOn
    /fxrtn/04/mix/06/level  F32 F_XET
/fxrtn/04/mix/07  <CHMO> n=0
    /fxrtn/04/mix/07/on  E32 F_XET enum=OffOn
    /fxrtn/04/mix/07/level  F32 F_XET
    /fxrtn/04/mix/07/pan  F32 F_XET
    /fxrtn/04/mix/07/type  E32 F_XET enum=Xmtype
    /fxrtn/04/mix/07/panFollow  E32 F_XET
/fxrtn/04/mix/08  <CHME> n=0
    /fxrtn/04/mix/08/on  E32 F_XET enum=OffOn
    /fxrtn/04/mix/08/level  F32 F_XET
/fxrtn/04/mix/09  <CHMO> n=0
    /fxrtn/04/mix/09/on  E32 F_XET enum=OffOn
    /fxrtn/04/mix/09/level  F32 F_XET
    /fxrtn/04/mix/09/pan  F32 F_XET
    /fxrtn/04/mix/09/type  E32 F_XET enum=Xmtype
    /fxrtn/04/mix/09/panFollow  E32 F_XET
/fxrtn/04/mix/10  <CHME> n=0
    /fxrtn/04/mix/10/on  E32 F_XET enum=OffOn
    /fxrtn/04/mix/10/level  F32 F_XET
/fxrtn/04/mix/11  <CHMO> n=0
    /fxrtn/04/mix/11/on  E32 F_XET enum=OffOn
    /fxrtn/04/mix/11/level  F32 F_XET
    /fxrtn/04/mix/11/pan  F32 F_XET
    /fxrtn/04/mix/11/type  E32 F_XET enum=Xmtype
    /fxrtn/04/mix/11/panFollow  E32 F_XET
/fxrtn/04/mix/12  <CHME> n=0
    /fxrtn/04/mix/12/on  E32 F_XET enum=OffOn
    /fxrtn/04/mix/12/level  F32 F_XET
/fxrtn/04/mix/13  <CHMO> n=0
    /fxrtn/04/mix/13/on  E32 F_XET enum=OffOn
    /fxrtn/04/mix/13/level  F32 F_XET
    /fxrtn/04/mix/13/pan  F32 F_XET
    /fxrtn/04/mix/13/type  E32 F_XET enum=Xmtype
    /fxrtn/04/mix/13/panFollow  E32 F_XET
/fxrtn/04/mix/14  <CHME> n=0
    /fxrtn/04/mix/14/on  E32 F_XET enum=OffOn
    /fxrtn/04/mix/14/level  F32 F_XET
/fxrtn/04/mix/15  <CHMO> n=0
    /fxrtn/04/mix/15/on  E32 F_XET enum=OffOn
    /fxrtn/04/mix/15/level  F32 F_XET
    /fxrtn/04/mix/15/pan  F32 F_XET
    /fxrtn/04/mix/15/type  E32 F_XET enum=Xmtype
    /fxrtn/04/mix/15/panFollow  E32 F_XET
/fxrtn/04/mix/16  <CHME> n=0
    /fxrtn/04/mix/16/on  E32 F_XET enum=OffOn
    /fxrtn/04/mix/16/level  F32 F_XET
/fxrtn/04/grp  <CHGRP> n=0
    /fxrtn/04/grp/dca  P32 F_XET
    /fxrtn/04/grp/mute  P32 F_XET
```

### Xfxrtn05 (X32Fxrtn.h, 110 entries)

```
/fxrtn  <BSCO> n=0
/fxrtn/05  <BSCO> n=0
/fxrtn/05/config  <BSCO> n=0
    /fxrtn/05/config/name  S32 F_XET
    /fxrtn/05/config/icon  I32 F_XET
    /fxrtn/05/config/color  E32 F_XET enum=Xcolors
/fxrtn/05/eq  <OFFON> n=1
    /fxrtn/05/eq/on  E32 F_XET enum=OffOn
/fxrtn/05/eq/1  <CHEQ> n=0
    /fxrtn/05/eq/1/type  E32 F_XET enum=Xeqty1
    /fxrtn/05/eq/1/f  F32 F_XET
    /fxrtn/05/eq/1/g  F32 F_XET
    /fxrtn/05/eq/1/q  F32 F_XET
/fxrtn/05/eq/2  <CHEQ> n=0
    /fxrtn/05/eq/2/type  E32 F_XET enum=Xeqty1
    /fxrtn/05/eq/2/f  F32 F_XET
    /fxrtn/05/eq/2/g  F32 F_XET
    /fxrtn/05/eq/2/q  F32 F_XET
/fxrtn/05/eq/3  <CHEQ> n=0
    /fxrtn/05/eq/3/type  E32 F_XET enum=Xeqty1
    /fxrtn/05/eq/3/f  F32 F_XET
    /fxrtn/05/eq/3/g  F32 F_XET
    /fxrtn/05/eq/3/q  F32 F_XET
/fxrtn/05/eq/4  <CHEQ> n=0
    /fxrtn/05/eq/4/type  E32 F_XET enum=Xeqty1
    /fxrtn/05/eq/4/f  F32 F_XET
    /fxrtn/05/eq/4/g  F32 F_XET
    /fxrtn/05/eq/4/q  F32 F_XET
/fxrtn/05/mix  <CHMX> n=0
    /fxrtn/05/mix/on  E32 F_XET enum=OffOn
    /fxrtn/05/mix/fader  F32 F_XET
    /fxrtn/05/mix/st  E32 F_XET enum=OffOn
    /fxrtn/05/mix/pan  F32 F_XET
    /fxrtn/05/mix/mono  E32 F_XET enum=OffOn
    /fxrtn/05/mix/mlevel  F32 F_XET
/fxrtn/05/mix/01  <CHMO> n=0
    /fxrtn/05/mix/01/on  E32 F_XET enum=OffOn
    /fxrtn/05/mix/01/level  F32 F_XET
    /fxrtn/05/mix/01/pan  F32 F_XET
    /fxrtn/05/mix/01/type  E32 F_XET enum=Xmtype
    /fxrtn/05/mix/01/panFollow  E32 F_XET
/fxrtn/05/mix/02  <CHME> n=0
    /fxrtn/05/mix/02/on  E32 F_XET enum=OffOn
    /fxrtn/05/mix/02/level  F32 F_XET
/fxrtn/05/mix/03  <CHMO> n=0
    /fxrtn/05/mix/03/on  E32 F_XET enum=OffOn
    /fxrtn/05/mix/03/level  F32 F_XET
    /fxrtn/05/mix/03/pan  F32 F_XET
    /fxrtn/05/mix/03/type  E32 F_XET enum=Xmtype
    /fxrtn/05/mix/03/panFollow  E32 F_XET
/fxrtn/05/mix/04  <CHME> n=0
    /fxrtn/05/mix/04/on  E32 F_XET enum=OffOn
    /fxrtn/05/mix/04/level  F32 F_XET
/fxrtn/05/mix/05  <CHMO> n=0
    /fxrtn/05/mix/05/on  E32 F_XET enum=OffOn
    /fxrtn/05/mix/05/level  F32 F_XET
    /fxrtn/05/mix/05/pan  F32 F_XET
    /fxrtn/05/mix/05/type  E32 F_XET enum=Xmtype
    /fxrtn/05/mix/05/panFollow  E32 F_XET
/fxrtn/05/mix/06  <CHME> n=0
    /fxrtn/05/mix/06/on  E32 F_XET enum=OffOn
    /fxrtn/05/mix/06/level  F32 F_XET
/fxrtn/05/mix/07  <CHMO> n=0
    /fxrtn/05/mix/07/on  E32 F_XET enum=OffOn
    /fxrtn/05/mix/07/level  F32 F_XET
    /fxrtn/05/mix/07/pan  F32 F_XET
    /fxrtn/05/mix/07/type  E32 F_XET enum=Xmtype
    /fxrtn/05/mix/07/panFollow  E32 F_XET
/fxrtn/05/mix/08  <CHME> n=0
    /fxrtn/05/mix/08/on  E32 F_XET enum=OffOn
    /fxrtn/05/mix/08/level  F32 F_XET
/fxrtn/05/mix/09  <CHMO> n=0
    /fxrtn/05/mix/09/on  E32 F_XET enum=OffOn
    /fxrtn/05/mix/09/level  F32 F_XET
    /fxrtn/05/mix/09/pan  F32 F_XET
    /fxrtn/05/mix/09/type  E32 F_XET enum=Xmtype
    /fxrtn/05/mix/09/panFollow  E32 F_XET
/fxrtn/05/mix/10  <CHME> n=0
    /fxrtn/05/mix/10/on  E32 F_XET enum=OffOn
    /fxrtn/05/mix/10/level  F32 F_XET
/fxrtn/05/mix/11  <CHMO> n=0
    /fxrtn/05/mix/11/on  E32 F_XET enum=OffOn
    /fxrtn/05/mix/11/level  F32 F_XET
    /fxrtn/05/mix/11/pan  F32 F_XET
    /fxrtn/05/mix/11/type  E32 F_XET enum=Xmtype
    /fxrtn/05/mix/11/panFollow  E32 F_XET
/fxrtn/05/mix/12  <CHME> n=0
    /fxrtn/05/mix/12/on  E32 F_XET enum=OffOn
    /fxrtn/05/mix/12/level  F32 F_XET
/fxrtn/05/mix/13  <CHMO> n=0
    /fxrtn/05/mix/13/on  E32 F_XET enum=OffOn
    /fxrtn/05/mix/13/level  F32 F_XET
    /fxrtn/05/mix/13/pan  F32 F_XET
    /fxrtn/05/mix/13/type  E32 F_XET enum=Xmtype
    /fxrtn/05/mix/13/panFollow  E32 F_XET
/fxrtn/05/mix/14  <CHME> n=0
    /fxrtn/05/mix/14/on  E32 F_XET enum=OffOn
    /fxrtn/05/mix/14/level  F32 F_XET
/fxrtn/05/mix/15  <CHMO> n=0
    /fxrtn/05/mix/15/on  E32 F_XET enum=OffOn
    /fxrtn/05/mix/15/level  F32 F_XET
    /fxrtn/05/mix/15/pan  F32 F_XET
    /fxrtn/05/mix/15/type  E32 F_XET enum=Xmtype
    /fxrtn/05/mix/15/panFollow  E32 F_XET
/fxrtn/05/mix/16  <CHME> n=0
    /fxrtn/05/mix/16/on  E32 F_XET enum=OffOn
    /fxrtn/05/mix/16/level  F32 F_XET
/fxrtn/05/grp  <CHGRP> n=0
    /fxrtn/05/grp/dca  P32 F_XET
    /fxrtn/05/grp/mute  P32 F_XET
```

### Xfxrtn06 (X32Fxrtn.h, 110 entries)

```
/fxrtn  <BSCO> n=0
/fxrtn/06  <BSCO> n=0
/fxrtn/06/config  <BSCO> n=0
    /fxrtn/06/config/name  S32 F_XET
    /fxrtn/06/config/icon  I32 F_XET
    /fxrtn/06/config/color  E32 F_XET enum=Xcolors
/fxrtn/06/eq  <OFFON> n=1
    /fxrtn/06/eq/on  E32 F_XET enum=OffOn
/fxrtn/06/eq/1  <CHEQ> n=0
    /fxrtn/06/eq/1/type  E32 F_XET enum=Xeqty1
    /fxrtn/06/eq/1/f  F32 F_XET
    /fxrtn/06/eq/1/g  F32 F_XET
    /fxrtn/06/eq/1/q  F32 F_XET
/fxrtn/06/eq/2  <CHEQ> n=0
    /fxrtn/06/eq/2/type  E32 F_XET enum=Xeqty1
    /fxrtn/06/eq/2/f  F32 F_XET
    /fxrtn/06/eq/2/g  F32 F_XET
    /fxrtn/06/eq/2/q  F32 F_XET
/fxrtn/06/eq/3  <CHEQ> n=0
    /fxrtn/06/eq/3/type  E32 F_XET enum=Xeqty1
    /fxrtn/06/eq/3/f  F32 F_XET
    /fxrtn/06/eq/3/g  F32 F_XET
    /fxrtn/06/eq/3/q  F32 F_XET
/fxrtn/06/eq/4  <CHEQ> n=0
    /fxrtn/06/eq/4/type  E32 F_XET enum=Xeqty1
    /fxrtn/06/eq/4/f  F32 F_XET
    /fxrtn/06/eq/4/g  F32 F_XET
    /fxrtn/06/eq/4/q  F32 F_XET
/fxrtn/06/mix  <CHMX> n=0
    /fxrtn/06/mix/on  E32 F_XET enum=OffOn
    /fxrtn/06/mix/fader  F32 F_XET
    /fxrtn/06/mix/st  E32 F_XET enum=OffOn
    /fxrtn/06/mix/pan  F32 F_XET
    /fxrtn/06/mix/mono  E32 F_XET enum=OffOn
    /fxrtn/06/mix/mlevel  F32 F_XET
/fxrtn/06/mix/01  <CHMO> n=0
    /fxrtn/06/mix/01/on  E32 F_XET enum=OffOn
    /fxrtn/06/mix/01/level  F32 F_XET
    /fxrtn/06/mix/01/pan  F32 F_XET
    /fxrtn/06/mix/01/type  E32 F_XET enum=Xmtype
    /fxrtn/06/mix/01/panFollow  E32 F_XET
/fxrtn/06/mix/02  <CHME> n=0
    /fxrtn/06/mix/02/on  E32 F_XET enum=OffOn
    /fxrtn/06/mix/02/level  F32 F_XET
/fxrtn/06/mix/03  <CHMO> n=0
    /fxrtn/06/mix/03/on  E32 F_XET enum=OffOn
    /fxrtn/06/mix/03/level  F32 F_XET
    /fxrtn/06/mix/03/pan  F32 F_XET
    /fxrtn/06/mix/03/type  E32 F_XET enum=Xmtype
    /fxrtn/06/mix/03/panFollow  E32 F_XET
/fxrtn/06/mix/04  <CHME> n=0
    /fxrtn/06/mix/04/on  E32 F_XET enum=OffOn
    /fxrtn/06/mix/04/level  F32 F_XET
/fxrtn/06/mix/05  <CHMO> n=0
    /fxrtn/06/mix/05/on  E32 F_XET enum=OffOn
    /fxrtn/06/mix/05/level  F32 F_XET
    /fxrtn/06/mix/05/pan  F32 F_XET
    /fxrtn/06/mix/05/type  E32 F_XET enum=Xmtype
    /fxrtn/06/mix/05/panFollow  E32 F_XET
/fxrtn/06/mix/06  <CHME> n=0
    /fxrtn/06/mix/06/on  E32 F_XET enum=OffOn
    /fxrtn/06/mix/06/level  F32 F_XET
/fxrtn/06/mix/07  <CHMO> n=0
    /fxrtn/06/mix/07/on  E32 F_XET enum=OffOn
    /fxrtn/06/mix/07/level  F32 F_XET
    /fxrtn/06/mix/07/pan  F32 F_XET
    /fxrtn/06/mix/07/type  E32 F_XET enum=Xmtype
    /fxrtn/06/mix/07/panFollow  E32 F_XET
/fxrtn/06/mix/08  <CHME> n=0
    /fxrtn/06/mix/08/on  E32 F_XET enum=OffOn
    /fxrtn/06/mix/08/level  F32 F_XET
/fxrtn/06/mix/09  <CHMO> n=0
    /fxrtn/06/mix/09/on  E32 F_XET enum=OffOn
    /fxrtn/06/mix/09/level  F32 F_XET
    /fxrtn/06/mix/09/pan  F32 F_XET
    /fxrtn/06/mix/09/type  E32 F_XET enum=Xmtype
    /fxrtn/06/mix/09/panFollow  E32 F_XET
/fxrtn/06/mix/10  <CHME> n=0
    /fxrtn/06/mix/10/on  E32 F_XET enum=OffOn
    /fxrtn/06/mix/10/level  F32 F_XET
/fxrtn/06/mix/11  <CHMO> n=0
    /fxrtn/06/mix/11/on  E32 F_XET enum=OffOn
    /fxrtn/06/mix/11/level  F32 F_XET
    /fxrtn/06/mix/11/pan  F32 F_XET
    /fxrtn/06/mix/11/type  E32 F_XET enum=Xmtype
    /fxrtn/06/mix/11/panFollow  E32 F_XET
/fxrtn/06/mix/12  <CHME> n=0
    /fxrtn/06/mix/12/on  E32 F_XET enum=OffOn
    /fxrtn/06/mix/12/level  F32 F_XET
/fxrtn/06/mix/13  <CHMO> n=0
    /fxrtn/06/mix/13/on  E32 F_XET enum=OffOn
    /fxrtn/06/mix/13/level  F32 F_XET
    /fxrtn/06/mix/13/pan  F32 F_XET
    /fxrtn/06/mix/13/type  E32 F_XET enum=Xmtype
    /fxrtn/06/mix/13/panFollow  E32 F_XET
/fxrtn/06/mix/14  <CHME> n=0
    /fxrtn/06/mix/14/on  E32 F_XET enum=OffOn
    /fxrtn/06/mix/14/level  F32 F_XET
/fxrtn/06/mix/15  <CHMO> n=0
    /fxrtn/06/mix/15/on  E32 F_XET enum=OffOn
    /fxrtn/06/mix/15/level  F32 F_XET
    /fxrtn/06/mix/15/pan  F32 F_XET
    /fxrtn/06/mix/15/type  E32 F_XET enum=Xmtype
    /fxrtn/06/mix/15/panFollow  E32 F_XET
/fxrtn/06/mix/16  <CHME> n=0
    /fxrtn/06/mix/16/on  E32 F_XET enum=OffOn
    /fxrtn/06/mix/16/level  F32 F_XET
/fxrtn/06/grp  <CHGRP> n=0
    /fxrtn/06/grp/dca  P32 F_XET
    /fxrtn/06/grp/mute  P32 F_XET
```

### Xfxrtn07 (X32Fxrtn.h, 110 entries)

```
/fxrtn  <BSCO> n=0
/fxrtn/07  <BSCO> n=0
/fxrtn/07/config  <BSCO> n=0
    /fxrtn/07/config/name  S32 F_XET
    /fxrtn/07/config/icon  I32 F_XET
    /fxrtn/07/config/color  E32 F_XET enum=Xcolors
/fxrtn/07/eq  <OFFON> n=1
    /fxrtn/07/eq/on  E32 F_XET enum=OffOn
/fxrtn/07/eq/1  <CHEQ> n=0
    /fxrtn/07/eq/1/type  E32 F_XET enum=Xeqty1
    /fxrtn/07/eq/1/f  F32 F_XET
    /fxrtn/07/eq/1/g  F32 F_XET
    /fxrtn/07/eq/1/q  F32 F_XET
/fxrtn/07/eq/2  <CHEQ> n=0
    /fxrtn/07/eq/2/type  E32 F_XET enum=Xeqty1
    /fxrtn/07/eq/2/f  F32 F_XET
    /fxrtn/07/eq/2/g  F32 F_XET
    /fxrtn/07/eq/2/q  F32 F_XET
/fxrtn/07/eq/3  <CHEQ> n=0
    /fxrtn/07/eq/3/type  E32 F_XET enum=Xeqty1
    /fxrtn/07/eq/3/f  F32 F_XET
    /fxrtn/07/eq/3/g  F32 F_XET
    /fxrtn/07/eq/3/q  F32 F_XET
/fxrtn/07/eq/4  <CHEQ> n=0
    /fxrtn/07/eq/4/type  E32 F_XET enum=Xeqty1
    /fxrtn/07/eq/4/f  F32 F_XET
    /fxrtn/07/eq/4/g  F32 F_XET
    /fxrtn/07/eq/4/q  F32 F_XET
/fxrtn/07/mix  <CHMX> n=0
    /fxrtn/07/mix/on  E32 F_XET enum=OffOn
    /fxrtn/07/mix/fader  F32 F_XET
    /fxrtn/07/mix/st  E32 F_XET enum=OffOn
    /fxrtn/07/mix/pan  F32 F_XET
    /fxrtn/07/mix/mono  E32 F_XET enum=OffOn
    /fxrtn/07/mix/mlevel  F32 F_XET
/fxrtn/07/mix/01  <CHMO> n=0
    /fxrtn/07/mix/01/on  E32 F_XET enum=OffOn
    /fxrtn/07/mix/01/level  F32 F_XET
    /fxrtn/07/mix/01/pan  F32 F_XET
    /fxrtn/07/mix/01/type  E32 F_XET enum=Xmtype
    /fxrtn/07/mix/01/panFollow  E32 F_XET
/fxrtn/07/mix/02  <CHME> n=0
    /fxrtn/07/mix/02/on  E32 F_XET enum=OffOn
    /fxrtn/07/mix/02/level  F32 F_XET
/fxrtn/07/mix/03  <CHMO> n=0
    /fxrtn/07/mix/03/on  E32 F_XET enum=OffOn
    /fxrtn/07/mix/03/level  F32 F_XET
    /fxrtn/07/mix/03/pan  F32 F_XET
    /fxrtn/07/mix/03/type  E32 F_XET enum=Xmtype
    /fxrtn/07/mix/03/panFollow  E32 F_XET
/fxrtn/07/mix/04  <CHME> n=0
    /fxrtn/07/mix/04/on  E32 F_XET enum=OffOn
    /fxrtn/07/mix/04/level  F32 F_XET
/fxrtn/07/mix/05  <CHMO> n=0
    /fxrtn/07/mix/05/on  E32 F_XET enum=OffOn
    /fxrtn/07/mix/05/level  F32 F_XET
    /fxrtn/07/mix/05/pan  F32 F_XET
    /fxrtn/07/mix/05/type  E32 F_XET enum=Xmtype
    /fxrtn/07/mix/05/panFollow  E32 F_XET
/fxrtn/07/mix/06  <CHME> n=0
    /fxrtn/07/mix/06/on  E32 F_XET enum=OffOn
    /fxrtn/07/mix/06/level  F32 F_XET
/fxrtn/07/mix/07  <CHMO> n=0
    /fxrtn/07/mix/07/on  E32 F_XET enum=OffOn
    /fxrtn/07/mix/07/level  F32 F_XET
    /fxrtn/07/mix/07/pan  F32 F_XET
    /fxrtn/07/mix/07/type  E32 F_XET enum=Xmtype
    /fxrtn/07/mix/07/panFollow  E32 F_XET
/fxrtn/07/mix/08  <CHME> n=0
    /fxrtn/07/mix/08/on  E32 F_XET enum=OffOn
    /fxrtn/07/mix/08/level  F32 F_XET
/fxrtn/07/mix/09  <CHMO> n=0
    /fxrtn/07/mix/09/on  E32 F_XET enum=OffOn
    /fxrtn/07/mix/09/level  F32 F_XET
    /fxrtn/07/mix/09/pan  F32 F_XET
    /fxrtn/07/mix/09/type  E32 F_XET enum=Xmtype
    /fxrtn/07/mix/09/panFollow  E32 F_XET
/fxrtn/07/mix/10  <CHME> n=0
    /fxrtn/07/mix/10/on  E32 F_XET enum=OffOn
    /fxrtn/07/mix/10/level  F32 F_XET
/fxrtn/07/mix/11  <CHMO> n=0
    /fxrtn/07/mix/11/on  E32 F_XET enum=OffOn
    /fxrtn/07/mix/11/level  F32 F_XET
    /fxrtn/07/mix/11/pan  F32 F_XET
    /fxrtn/07/mix/11/type  E32 F_XET enum=Xmtype
    /fxrtn/07/mix/11/panFollow  E32 F_XET
/fxrtn/07/mix/12  <CHME> n=0
    /fxrtn/07/mix/12/on  E32 F_XET enum=OffOn
    /fxrtn/07/mix/12/level  F32 F_XET
/fxrtn/07/mix/13  <CHMO> n=0
    /fxrtn/07/mix/13/on  E32 F_XET enum=OffOn
    /fxrtn/07/mix/13/level  F32 F_XET
    /fxrtn/07/mix/13/pan  F32 F_XET
    /fxrtn/07/mix/13/type  E32 F_XET enum=Xmtype
    /fxrtn/07/mix/13/panFollow  E32 F_XET
/fxrtn/07/mix/14  <CHME> n=0
    /fxrtn/07/mix/14/on  E32 F_XET enum=OffOn
    /fxrtn/07/mix/14/level  F32 F_XET
/fxrtn/07/mix/15  <CHMO> n=0
    /fxrtn/07/mix/15/on  E32 F_XET enum=OffOn
    /fxrtn/07/mix/15/level  F32 F_XET
    /fxrtn/07/mix/15/pan  F32 F_XET
    /fxrtn/07/mix/15/type  E32 F_XET enum=Xmtype
    /fxrtn/07/mix/15/panFollow  E32 F_XET
/fxrtn/07/mix/16  <CHME> n=0
    /fxrtn/07/mix/16/on  E32 F_XET enum=OffOn
    /fxrtn/07/mix/16/level  F32 F_XET
/fxrtn/07/grp  <CHGRP> n=0
    /fxrtn/07/grp/dca  P32 F_XET
    /fxrtn/07/grp/mute  P32 F_XET
```

### Xfxrtn08 (X32Fxrtn.h, 110 entries)

```
/fxrtn  <BSCO> n=0
/fxrtn/08  <BSCO> n=0
/fxrtn/08/config  <BSCO> n=0
    /fxrtn/08/config/name  S32 F_XET
    /fxrtn/08/config/icon  I32 F_XET
    /fxrtn/08/config/color  E32 F_XET enum=Xcolors
/fxrtn/08/eq  <OFFON> n=1
    /fxrtn/08/eq/on  E32 F_XET enum=OffOn
/fxrtn/08/eq/1  <CHEQ> n=0
    /fxrtn/08/eq/1/type  E32 F_XET enum=Xeqty1
    /fxrtn/08/eq/1/f  F32 F_XET
    /fxrtn/08/eq/1/g  F32 F_XET
    /fxrtn/08/eq/1/q  F32 F_XET
/fxrtn/08/eq/2  <CHEQ> n=0
    /fxrtn/08/eq/2/type  E32 F_XET enum=Xeqty1
    /fxrtn/08/eq/2/f  F32 F_XET
    /fxrtn/08/eq/2/g  F32 F_XET
    /fxrtn/08/eq/2/q  F32 F_XET
/fxrtn/08/eq/3  <CHEQ> n=0
    /fxrtn/08/eq/3/type  E32 F_XET enum=Xeqty1
    /fxrtn/08/eq/3/f  F32 F_XET
    /fxrtn/08/eq/3/g  F32 F_XET
    /fxrtn/08/eq/3/q  F32 F_XET
/fxrtn/08/eq/4  <CHEQ> n=0
    /fxrtn/08/eq/4/type  E32 F_XET enum=Xeqty1
    /fxrtn/08/eq/4/f  F32 F_XET
    /fxrtn/08/eq/4/g  F32 F_XET
    /fxrtn/08/eq/4/q  F32 F_XET
/fxrtn/08/mix  <CHMX> n=0
    /fxrtn/08/mix/on  E32 F_XET enum=OffOn
    /fxrtn/08/mix/fader  F32 F_XET
    /fxrtn/08/mix/st  E32 F_XET enum=OffOn
    /fxrtn/08/mix/pan  F32 F_XET
    /fxrtn/08/mix/mono  E32 F_XET enum=OffOn
    /fxrtn/08/mix/mlevel  F32 F_XET
/fxrtn/08/mix/01  <CHMO> n=0
    /fxrtn/08/mix/01/on  E32 F_XET enum=OffOn
    /fxrtn/08/mix/01/level  F32 F_XET
    /fxrtn/08/mix/01/pan  F32 F_XET
    /fxrtn/08/mix/01/type  E32 F_XET enum=Xmtype
    /fxrtn/08/mix/01/panFollow  E32 F_XET
/fxrtn/08/mix/02  <CHME> n=0
    /fxrtn/08/mix/02/on  E32 F_XET enum=OffOn
    /fxrtn/08/mix/02/level  F32 F_XET
/fxrtn/08/mix/03  <CHMO> n=0
    /fxrtn/08/mix/03/on  E32 F_XET enum=OffOn
    /fxrtn/08/mix/03/level  F32 F_XET
    /fxrtn/08/mix/03/pan  F32 F_XET
    /fxrtn/08/mix/03/type  E32 F_XET enum=Xmtype
    /fxrtn/08/mix/03/panFollow  E32 F_XET
/fxrtn/08/mix/04  <CHME> n=0
    /fxrtn/08/mix/04/on  E32 F_XET enum=OffOn
    /fxrtn/08/mix/04/level  F32 F_XET
/fxrtn/08/mix/05  <CHMO> n=0
    /fxrtn/08/mix/05/on  E32 F_XET enum=OffOn
    /fxrtn/08/mix/05/level  F32 F_XET
    /fxrtn/08/mix/05/pan  F32 F_XET
    /fxrtn/08/mix/05/type  E32 F_XET enum=Xmtype
    /fxrtn/08/mix/05/panFollow  E32 F_XET
/fxrtn/08/mix/06  <CHME> n=0
    /fxrtn/08/mix/06/on  E32 F_XET enum=OffOn
    /fxrtn/08/mix/06/level  F32 F_XET
/fxrtn/08/mix/07  <CHMO> n=0
    /fxrtn/08/mix/07/on  E32 F_XET enum=OffOn
    /fxrtn/08/mix/07/level  F32 F_XET
    /fxrtn/08/mix/07/pan  F32 F_XET
    /fxrtn/08/mix/07/type  E32 F_XET enum=Xmtype
    /fxrtn/08/mix/07/panFollow  E32 F_XET
/fxrtn/08/mix/08  <CHME> n=0
    /fxrtn/08/mix/08/on  E32 F_XET enum=OffOn
    /fxrtn/08/mix/08/level  F32 F_XET
/fxrtn/08/mix/09  <CHMO> n=0
    /fxrtn/08/mix/09/on  E32 F_XET enum=OffOn
    /fxrtn/08/mix/09/level  F32 F_XET
    /fxrtn/08/mix/09/pan  F32 F_XET
    /fxrtn/08/mix/09/type  E32 F_XET enum=Xmtype
    /fxrtn/08/mix/09/panFollow  E32 F_XET
/fxrtn/08/mix/10  <CHME> n=0
    /fxrtn/08/mix/10/on  E32 F_XET enum=OffOn
    /fxrtn/08/mix/10/level  F32 F_XET
/fxrtn/08/mix/11  <CHMO> n=0
    /fxrtn/08/mix/11/on  E32 F_XET enum=OffOn
    /fxrtn/08/mix/11/level  F32 F_XET
    /fxrtn/08/mix/11/pan  F32 F_XET
    /fxrtn/08/mix/11/type  E32 F_XET enum=Xmtype
    /fxrtn/08/mix/11/panFollow  E32 F_XET
/fxrtn/08/mix/12  <CHME> n=0
    /fxrtn/08/mix/12/on  E32 F_XET enum=OffOn
    /fxrtn/08/mix/12/level  F32 F_XET
/fxrtn/08/mix/13  <CHMO> n=0
    /fxrtn/08/mix/13/on  E32 F_XET enum=OffOn
    /fxrtn/08/mix/13/level  F32 F_XET
    /fxrtn/08/mix/13/pan  F32 F_XET
    /fxrtn/08/mix/13/type  E32 F_XET enum=Xmtype
    /fxrtn/08/mix/13/panFollow  E32 F_XET
/fxrtn/08/mix/14  <CHME> n=0
    /fxrtn/08/mix/14/on  E32 F_XET enum=OffOn
    /fxrtn/08/mix/14/level  F32 F_XET
/fxrtn/08/mix/15  <CHMO> n=0
    /fxrtn/08/mix/15/on  E32 F_XET enum=OffOn
    /fxrtn/08/mix/15/level  F32 F_XET
    /fxrtn/08/mix/15/pan  F32 F_XET
    /fxrtn/08/mix/15/type  E32 F_XET enum=Xmtype
    /fxrtn/08/mix/15/panFollow  E32 F_XET
/fxrtn/08/mix/16  <CHME> n=0
    /fxrtn/08/mix/16/on  E32 F_XET enum=OffOn
    /fxrtn/08/mix/16/level  F32 F_XET
/fxrtn/08/grp  <CHGRP> n=0
    /fxrtn/08/grp/dca  P32 F_XET
    /fxrtn/08/grp/mute  P32 F_XET
```

### Xbus01 (X32Bus.h, 99 entries)

```
/bus  <BSCO> n=0
/bus/01  <BSCO> n=0
/bus/01/config  <BSCO> n=0
    /bus/01/config/name  S32 F_XET
    /bus/01/config/icon  I32 F_XET
    /bus/01/config/color  E32 F_XET enum=Xcolors
/bus/01/dyn  <CHDY> n=0
    /bus/01/dyn/on  E32 F_XET enum=OffOn
    /bus/01/dyn/mode  E32 F_XET enum=Xdymode
    /bus/01/dyn/det  E32 F_XET enum=Xdydet
    /bus/01/dyn/env  E32 F_XET enum=Xdyenv
    /bus/01/dyn/thr  F32 F_XET
    /bus/01/dyn/ratio  E32 F_XET enum=OffOn
    /bus/01/dyn/knee  F32 F_XET
    /bus/01/dyn/mgain  F32 F_XET
    /bus/01/dyn/attack  F32 F_XET
    /bus/01/dyn/hold  F32 F_XET
    /bus/01/dyn/release  F32 F_XET
    /bus/01/dyn/pos  E32 F_XET enum=Xdyppos
    /bus/01/dyn/keysrc  I32 F_XET
    /bus/01/dyn/mix  F32 F_XET enum=OffOn
    /bus/01/dyn/auto  E32 F_XET enum=OffOn
/bus/01/dyn/filter  <CHDF> n=0
    /bus/01/dyn/filter/on  E32 F_XET enum=OffOn
    /bus/01/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /bus/01/dyn/filter/f  F32 F_XET
/bus/01/insert  <CHIN> n=0
    /bus/01/insert/on  E32 F_XET enum=OffOn
    /bus/01/insert/pos  E32 F_XET enum=Xdyppos
    /bus/01/insert/sel  E32 F_XET enum=Xisel
/bus/01/eq  <OFFON> n=1
    /bus/01/eq/on  E32 F_XET enum=OffOn
/bus/01/eq/1  <CHEQ> n=0
    /bus/01/eq/1/type  E32 F_XET enum=Xeqty1
    /bus/01/eq/1/f  F32 F_XET
    /bus/01/eq/1/g  F32 F_XET
    /bus/01/eq/1/q  F32 F_XET
/bus/01/eq/2  <CHEQ> n=0
    /bus/01/eq/2/type  E32 F_XET enum=Xeqty1
    /bus/01/eq/2/f  F32 F_XET
    /bus/01/eq/2/g  F32 F_XET
    /bus/01/eq/2/q  F32 F_XET
/bus/01/eq/3  <CHEQ> n=0
    /bus/01/eq/3/type  E32 F_XET enum=Xeqty1
    /bus/01/eq/3/f  F32 F_XET
    /bus/01/eq/3/g  F32 F_XET
    /bus/01/eq/3/q  F32 F_XET
/bus/01/eq/4  <CHEQ> n=0
    /bus/01/eq/4/type  E32 F_XET enum=Xeqty1
    /bus/01/eq/4/f  F32 F_XET
    /bus/01/eq/4/g  F32 F_XET
    /bus/01/eq/4/q  F32 F_XET
/bus/01/eq/5  <CHEQ> n=0
    /bus/01/eq/5/type  E32 F_XET enum=Xeqty1
    /bus/01/eq/5/f  F32 F_XET
    /bus/01/eq/5/g  F32 F_XET
    /bus/01/eq/5/q  F32 F_XET
/bus/01/eq/6  <CHEQ> n=0
    /bus/01/eq/6/type  E32 F_XET enum=Xeqty1
    /bus/01/eq/6/f  F32 F_XET
    /bus/01/eq/6/g  F32 F_XET
    /bus/01/eq/6/q  F32 F_XET
/bus/01/mix  <CHMX> n=0
    /bus/01/mix/on  E32 F_XET enum=OffOn
    /bus/01/mix/fader  F32 F_XET
    /bus/01/mix/st  E32 F_XET enum=OffOn
    /bus/01/mix/pan  F32 F_XET
    /bus/01/mix/mono  E32 F_XET enum=OffOn
    /bus/01/mix/mlevel  F32 F_XET
/bus/01/mix/01  <CHMO> n=0
    /bus/01/mix/01/on  E32 F_XET enum=OffOn
    /bus/01/mix/01/level  F32 F_XET
    /bus/01/mix/01/pan  F32 F_XET
    /bus/01/mix/01/type  E32 F_XET enum=OffOn
    /bus/01/mix/01/panFollow  E32 F_XET
/bus/01/mix/02  <CHME> n=0
    /bus/01/mix/02/on  E32 F_XET enum=OffOn
    /bus/01/mix/02/level  F32 F_XET
/bus/01/mix/03  <CHMO> n=0
    /bus/01/mix/03/on  E32 F_XET enum=OffOn
    /bus/01/mix/03/level  F32 F_XET
    /bus/01/mix/03/pan  F32 F_XET
    /bus/01/mix/03/type  E32 F_XET enum=OffOn
    /bus/01/mix/03/panFollow  E32 F_XET
/bus/01/mix/04  <CHME> n=0
    /bus/01/mix/04/on  E32 F_XET enum=OffOn
    /bus/01/mix/04/level  F32 F_XET
/bus/01/mix/05  <CHMO> n=0
    /bus/01/mix/05/on  E32 F_XET enum=OffOn
    /bus/01/mix/05/level  F32 F_XET
    /bus/01/mix/05/pan  F32 F_XET
    /bus/01/mix/05/type  E32 F_XET enum=OffOn
    /bus/01/mix/05/panFollow  E32 F_XET
/bus/01/mix/06  <CHME> n=0
    /bus/01/mix/06/on  E32 F_XET enum=OffOn
    /bus/01/mix/06/level  F32 F_XET
/bus/01/grp  <CHGRP> n=0
    /bus/01/grp/dca  P32 F_XET
    /bus/01/grp/mute  P32 F_XET
```

### Xbus02 (X32Bus.h, 99 entries)

```
/bus  <BSCO> n=0
/bus/02  <BSCO> n=0
/bus/02/config  <BSCO> n=0
    /bus/02/config/name  S32 F_XET
    /bus/02/config/icon  I32 F_XET
    /bus/02/config/color  E32 F_XET enum=Xcolors
/bus/02/dyn  <CHDY> n=0
    /bus/02/dyn/on  E32 F_XET enum=OffOn
    /bus/02/dyn/mode  E32 F_XET enum=Xdymode
    /bus/02/dyn/det  E32 F_XET enum=Xdydet
    /bus/02/dyn/env  E32 F_XET enum=Xdyenv
    /bus/02/dyn/thr  F32 F_XET
    /bus/02/dyn/ratio  E32 F_XET enum=OffOn
    /bus/02/dyn/knee  F32 F_XET
    /bus/02/dyn/mgain  F32 F_XET
    /bus/02/dyn/attack  F32 F_XET
    /bus/02/dyn/hold  F32 F_XET
    /bus/02/dyn/release  F32 F_XET
    /bus/02/dyn/pos  E32 F_XET enum=Xdyppos
    /bus/02/dyn/keysrc  I32 F_XET
    /bus/02/dyn/mix  F32 F_XET enum=OffOn
    /bus/02/dyn/auto  E32 F_XET enum=OffOn
/bus/02/dyn/filter  <CHDF> n=0
    /bus/02/dyn/filter/on  E32 F_XET enum=OffOn
    /bus/02/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /bus/02/dyn/filter/f  F32 F_XET
/bus/02/insert  <CHIN> n=0
    /bus/02/insert/on  E32 F_XET enum=OffOn
    /bus/02/insert/pos  E32 F_XET enum=Xdyppos
    /bus/02/insert/sel  E32 F_XET enum=Xisel
/bus/02/eq  <OFFON> n=1
    /bus/02/eq/on  E32 F_XET enum=OffOn
/bus/02/eq/1  <CHEQ> n=0
    /bus/02/eq/1/type  E32 F_XET enum=Xeqty1
    /bus/02/eq/1/f  F32 F_XET
    /bus/02/eq/1/g  F32 F_XET
    /bus/02/eq/1/q  F32 F_XET
/bus/02/eq/2  <CHEQ> n=0
    /bus/02/eq/2/type  E32 F_XET enum=Xeqty1
    /bus/02/eq/2/f  F32 F_XET
    /bus/02/eq/2/g  F32 F_XET
    /bus/02/eq/2/q  F32 F_XET
/bus/02/eq/3  <CHEQ> n=0
    /bus/02/eq/3/type  E32 F_XET enum=Xeqty1
    /bus/02/eq/3/f  F32 F_XET
    /bus/02/eq/3/g  F32 F_XET
    /bus/02/eq/3/q  F32 F_XET
/bus/02/eq/4  <CHEQ> n=0
    /bus/02/eq/4/type  E32 F_XET enum=Xeqty1
    /bus/02/eq/4/f  F32 F_XET
    /bus/02/eq/4/g  F32 F_XET
    /bus/02/eq/4/q  F32 F_XET
/bus/02/eq/5  <CHEQ> n=0
    /bus/02/eq/5/type  E32 F_XET enum=Xeqty1
    /bus/02/eq/5/f  F32 F_XET
    /bus/02/eq/5/g  F32 F_XET
    /bus/02/eq/5/q  F32 F_XET
/bus/02/eq/6  <CHEQ> n=0
    /bus/02/eq/6/type  E32 F_XET enum=Xeqty1
    /bus/02/eq/6/f  F32 F_XET
    /bus/02/eq/6/g  F32 F_XET
    /bus/02/eq/6/q  F32 F_XET
/bus/02/mix  <CHMX> n=0
    /bus/02/mix/on  E32 F_XET enum=OffOn
    /bus/02/mix/fader  F32 F_XET
    /bus/02/mix/st  E32 F_XET enum=OffOn
    /bus/02/mix/pan  F32 F_XET
    /bus/02/mix/mono  E32 F_XET enum=OffOn
    /bus/02/mix/mlevel  F32 F_XET
/bus/02/mix/01  <CHMO> n=0
    /bus/02/mix/01/on  E32 F_XET enum=OffOn
    /bus/02/mix/01/level  F32 F_XET
    /bus/02/mix/01/pan  F32 F_XET
    /bus/02/mix/01/type  E32 F_XET enum=OffOn
    /bus/02/mix/01/panFollow  E32 F_XET
/bus/02/mix/02  <CHME> n=0
    /bus/02/mix/02/on  E32 F_XET enum=OffOn
    /bus/02/mix/02/level  F32 F_XET
/bus/02/mix/03  <CHMO> n=0
    /bus/02/mix/03/on  E32 F_XET enum=OffOn
    /bus/02/mix/03/level  F32 F_XET
    /bus/02/mix/03/pan  F32 F_XET
    /bus/02/mix/03/type  E32 F_XET enum=OffOn
    /bus/02/mix/03/panFollow  E32 F_XET
/bus/02/mix/04  <CHME> n=0
    /bus/02/mix/04/on  E32 F_XET enum=OffOn
    /bus/02/mix/04/level  F32 F_XET
/bus/02/mix/05  <CHMO> n=0
    /bus/02/mix/05/on  E32 F_XET enum=OffOn
    /bus/02/mix/05/level  F32 F_XET
    /bus/02/mix/05/pan  F32 F_XET
    /bus/02/mix/05/type  E32 F_XET enum=OffOn
    /bus/02/mix/05/panFollow  E32 F_XET
/bus/02/mix/06  <CHME> n=0
    /bus/02/mix/06/on  E32 F_XET enum=OffOn
    /bus/02/mix/06/level  F32 F_XET
/bus/02/grp  <CHGRP> n=0
    /bus/02/grp/dca  P32 F_XET
    /bus/02/grp/mute  P32 F_XET
```

### Xbus03 (X32Bus.h, 99 entries)

```
/bus  <BSCO> n=0
/bus/03  <BSCO> n=0
/bus/03/config  <BSCO> n=0
    /bus/03/config/name  S32 F_XET
    /bus/03/config/icon  I32 F_XET
    /bus/03/config/color  E32 F_XET enum=Xcolors
/bus/03/dyn  <CHDY> n=0
    /bus/03/dyn/on  E32 F_XET enum=OffOn
    /bus/03/dyn/mode  E32 F_XET enum=Xdymode
    /bus/03/dyn/det  E32 F_XET enum=Xdydet
    /bus/03/dyn/env  E32 F_XET enum=Xdyenv
    /bus/03/dyn/thr  F32 F_XET
    /bus/03/dyn/ratio  E32 F_XET enum=OffOn
    /bus/03/dyn/knee  F32 F_XET
    /bus/03/dyn/mgain  F32 F_XET
    /bus/03/dyn/attack  F32 F_XET
    /bus/03/dyn/hold  F32 F_XET
    /bus/03/dyn/release  F32 F_XET
    /bus/03/dyn/pos  E32 F_XET enum=Xdyppos
    /bus/03/dyn/keysrc  I32 F_XET
    /bus/03/dyn/mix  F32 F_XET enum=OffOn
    /bus/03/dyn/auto  E32 F_XET enum=OffOn
/bus/03/dyn/filter  <CHDF> n=0
    /bus/03/dyn/filter/on  E32 F_XET enum=OffOn
    /bus/03/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /bus/03/dyn/filter/f  F32 F_XET
/bus/03/insert  <CHIN> n=0
    /bus/03/insert/on  E32 F_XET enum=OffOn
    /bus/03/insert/pos  E32 F_XET enum=Xdyppos
    /bus/03/insert/sel  E32 F_XET enum=Xisel
/bus/03/eq  <OFFON> n=1
    /bus/03/eq/on  E32 F_XET enum=OffOn
/bus/03/eq/1  <CHEQ> n=0
    /bus/03/eq/1/type  E32 F_XET enum=Xeqty1
    /bus/03/eq/1/f  F32 F_XET
    /bus/03/eq/1/g  F32 F_XET
    /bus/03/eq/1/q  F32 F_XET
/bus/03/eq/2  <CHEQ> n=0
    /bus/03/eq/2/type  E32 F_XET enum=Xeqty1
    /bus/03/eq/2/f  F32 F_XET
    /bus/03/eq/2/g  F32 F_XET
    /bus/03/eq/2/q  F32 F_XET
/bus/03/eq/3  <CHEQ> n=0
    /bus/03/eq/3/type  E32 F_XET enum=Xeqty1
    /bus/03/eq/3/f  F32 F_XET
    /bus/03/eq/3/g  F32 F_XET
    /bus/03/eq/3/q  F32 F_XET
/bus/03/eq/4  <CHEQ> n=0
    /bus/03/eq/4/type  E32 F_XET enum=Xeqty1
    /bus/03/eq/4/f  F32 F_XET
    /bus/03/eq/4/g  F32 F_XET
    /bus/03/eq/4/q  F32 F_XET
/bus/03/eq/5  <CHEQ> n=0
    /bus/03/eq/5/type  E32 F_XET enum=Xeqty1
    /bus/03/eq/5/f  F32 F_XET
    /bus/03/eq/5/g  F32 F_XET
    /bus/03/eq/5/q  F32 F_XET
/bus/03/eq/6  <CHEQ> n=0
    /bus/03/eq/6/type  E32 F_XET enum=Xeqty1
    /bus/03/eq/6/f  F32 F_XET
    /bus/03/eq/6/g  F32 F_XET
    /bus/03/eq/6/q  F32 F_XET
/bus/03/mix  <CHMX> n=0
    /bus/03/mix/on  E32 F_XET enum=OffOn
    /bus/03/mix/fader  F32 F_XET
    /bus/03/mix/st  E32 F_XET enum=OffOn
    /bus/03/mix/pan  F32 F_XET
    /bus/03/mix/mono  E32 F_XET enum=OffOn
    /bus/03/mix/mlevel  F32 F_XET
/bus/03/mix/01  <CHMO> n=0
    /bus/03/mix/01/on  E32 F_XET enum=OffOn
    /bus/03/mix/01/level  F32 F_XET
    /bus/03/mix/01/pan  F32 F_XET
    /bus/03/mix/01/type  E32 F_XET enum=OffOn
    /bus/03/mix/01/panFollow  E32 F_XET
/bus/03/mix/02  <CHME> n=0
    /bus/03/mix/02/on  E32 F_XET enum=OffOn
    /bus/03/mix/02/level  F32 F_XET
/bus/03/mix/03  <CHMO> n=0
    /bus/03/mix/03/on  E32 F_XET enum=OffOn
    /bus/03/mix/03/level  F32 F_XET
    /bus/03/mix/03/pan  F32 F_XET
    /bus/03/mix/03/type  E32 F_XET enum=OffOn
    /bus/03/mix/03/panFollow  E32 F_XET
/bus/03/mix/04  <CHME> n=0
    /bus/03/mix/04/on  E32 F_XET enum=OffOn
    /bus/03/mix/04/level  F32 F_XET
/bus/03/mix/05  <CHMO> n=0
    /bus/03/mix/05/on  E32 F_XET enum=OffOn
    /bus/03/mix/05/level  F32 F_XET
    /bus/03/mix/05/pan  F32 F_XET
    /bus/03/mix/05/type  E32 F_XET enum=OffOn
    /bus/03/mix/05/panFollow  E32 F_XET
/bus/03/mix/06  <CHME> n=0
    /bus/03/mix/06/on  E32 F_XET enum=OffOn
    /bus/03/mix/06/level  F32 F_XET
/bus/03/grp  <CHGRP> n=0
    /bus/03/grp/dca  P32 F_XET
    /bus/03/grp/mute  P32 F_XET
```

### Xbus04 (X32Bus.h, 99 entries)

```
/bus  <BSCO> n=0
/bus/04  <BSCO> n=0
/bus/04/config  <BSCO> n=0
    /bus/04/config/name  S32 F_XET
    /bus/04/config/icon  I32 F_XET
    /bus/04/config/color  E32 F_XET enum=Xcolors
/bus/04/dyn  <CHDY> n=0
    /bus/04/dyn/on  E32 F_XET enum=OffOn
    /bus/04/dyn/mode  E32 F_XET enum=Xdymode
    /bus/04/dyn/det  E32 F_XET enum=Xdydet
    /bus/04/dyn/env  E32 F_XET enum=Xdyenv
    /bus/04/dyn/thr  F32 F_XET
    /bus/04/dyn/ratio  E32 F_XET enum=OffOn
    /bus/04/dyn/knee  F32 F_XET
    /bus/04/dyn/mgain  F32 F_XET
    /bus/04/dyn/attack  F32 F_XET
    /bus/04/dyn/hold  F32 F_XET
    /bus/04/dyn/release  F32 F_XET
    /bus/04/dyn/pos  E32 F_XET enum=Xdyppos
    /bus/04/dyn/keysrc  I32 F_XET
    /bus/04/dyn/mix  F32 F_XET enum=OffOn
    /bus/04/dyn/auto  E32 F_XET enum=OffOn
/bus/04/dyn/filter  <CHDF> n=0
    /bus/04/dyn/filter/on  E32 F_XET enum=OffOn
    /bus/04/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /bus/04/dyn/filter/f  F32 F_XET
/bus/04/insert  <CHIN> n=0
    /bus/04/insert/on  E32 F_XET enum=OffOn
    /bus/04/insert/pos  E32 F_XET enum=Xdyppos
    /bus/04/insert/sel  E32 F_XET enum=Xisel
/bus/04/eq  <OFFON> n=1
    /bus/04/eq/on  E32 F_XET enum=OffOn
/bus/04/eq/1  <CHEQ> n=0
    /bus/04/eq/1/type  E32 F_XET enum=Xeqty1
    /bus/04/eq/1/f  F32 F_XET
    /bus/04/eq/1/g  F32 F_XET
    /bus/04/eq/1/q  F32 F_XET
/bus/04/eq/2  <CHEQ> n=0
    /bus/04/eq/2/type  E32 F_XET enum=Xeqty1
    /bus/04/eq/2/f  F32 F_XET
    /bus/04/eq/2/g  F32 F_XET
    /bus/04/eq/2/q  F32 F_XET
/bus/04/eq/3  <CHEQ> n=0
    /bus/04/eq/3/type  E32 F_XET enum=Xeqty1
    /bus/04/eq/3/f  F32 F_XET
    /bus/04/eq/3/g  F32 F_XET
    /bus/04/eq/3/q  F32 F_XET
/bus/04/eq/4  <CHEQ> n=0
    /bus/04/eq/4/type  E32 F_XET enum=Xeqty1
    /bus/04/eq/4/f  F32 F_XET
    /bus/04/eq/4/g  F32 F_XET
    /bus/04/eq/4/q  F32 F_XET
/bus/04/eq/5  <CHEQ> n=0
    /bus/04/eq/5/type  E32 F_XET enum=Xeqty1
    /bus/04/eq/5/f  F32 F_XET
    /bus/04/eq/5/g  F32 F_XET
    /bus/04/eq/5/q  F32 F_XET
/bus/04/eq/6  <CHEQ> n=0
    /bus/04/eq/6/type  E32 F_XET enum=Xeqty1
    /bus/04/eq/6/f  F32 F_XET
    /bus/04/eq/6/g  F32 F_XET
    /bus/04/eq/6/q  F32 F_XET
/bus/04/mix  <CHMX> n=0
    /bus/04/mix/on  E32 F_XET enum=OffOn
    /bus/04/mix/fader  F32 F_XET
    /bus/04/mix/st  E32 F_XET enum=OffOn
    /bus/04/mix/pan  F32 F_XET
    /bus/04/mix/mono  E32 F_XET enum=OffOn
    /bus/04/mix/mlevel  F32 F_XET
/bus/04/mix/01  <CHMO> n=0
    /bus/04/mix/01/on  E32 F_XET enum=OffOn
    /bus/04/mix/01/level  F32 F_XET
    /bus/04/mix/01/pan  F32 F_XET
    /bus/04/mix/01/type  E32 F_XET enum=OffOn
    /bus/04/mix/01/panFollow  E32 F_XET
/bus/04/mix/02  <CHME> n=0
    /bus/04/mix/02/on  E32 F_XET enum=OffOn
    /bus/04/mix/02/level  F32 F_XET
/bus/04/mix/03  <CHMO> n=0
    /bus/04/mix/03/on  E32 F_XET enum=OffOn
    /bus/04/mix/03/level  F32 F_XET
    /bus/04/mix/03/pan  F32 F_XET
    /bus/04/mix/03/type  E32 F_XET enum=OffOn
    /bus/04/mix/03/panFollow  E32 F_XET
/bus/04/mix/04  <CHME> n=0
    /bus/04/mix/04/on  E32 F_XET enum=OffOn
    /bus/04/mix/04/level  F32 F_XET
/bus/04/mix/05  <CHMO> n=0
    /bus/04/mix/05/on  E32 F_XET enum=OffOn
    /bus/04/mix/05/level  F32 F_XET
    /bus/04/mix/05/pan  F32 F_XET
    /bus/04/mix/05/type  E32 F_XET enum=OffOn
    /bus/04/mix/05/panFollow  E32 F_XET
/bus/04/mix/06  <CHME> n=0
    /bus/04/mix/06/on  E32 F_XET enum=OffOn
    /bus/04/mix/06/level  F32 F_XET
/bus/04/grp  <CHGRP> n=0
    /bus/04/grp/dca  P32 F_XET
    /bus/04/grp/mute  P32 F_XET
```

### Xbus05 (X32Bus.h, 99 entries)

```
/bus  <BSCO> n=0
/bus/05  <BSCO> n=0
/bus/05/config  <BSCO> n=0
    /bus/05/config/name  S32 F_XET
    /bus/05/config/icon  I32 F_XET
    /bus/05/config/color  E32 F_XET enum=Xcolors
/bus/05/dyn  <CHDY> n=0
    /bus/05/dyn/on  E32 F_XET enum=OffOn
    /bus/05/dyn/mode  E32 F_XET enum=Xdymode
    /bus/05/dyn/det  E32 F_XET enum=Xdydet
    /bus/05/dyn/env  E32 F_XET enum=Xdyenv
    /bus/05/dyn/thr  F32 F_XET
    /bus/05/dyn/ratio  E32 F_XET enum=OffOn
    /bus/05/dyn/knee  F32 F_XET
    /bus/05/dyn/mgain  F32 F_XET
    /bus/05/dyn/attack  F32 F_XET
    /bus/05/dyn/hold  F32 F_XET
    /bus/05/dyn/release  F32 F_XET
    /bus/05/dyn/pos  E32 F_XET enum=Xdyppos
    /bus/05/dyn/keysrc  I32 F_XET
    /bus/05/dyn/mix  F32 F_XET enum=OffOn
    /bus/05/dyn/auto  E32 F_XET enum=OffOn
/bus/05/dyn/filter  <CHDF> n=0
    /bus/05/dyn/filter/on  E32 F_XET enum=OffOn
    /bus/05/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /bus/05/dyn/filter/f  F32 F_XET
/bus/05/insert  <CHIN> n=0
    /bus/05/insert/on  E32 F_XET enum=OffOn
    /bus/05/insert/pos  E32 F_XET enum=Xdyppos
    /bus/05/insert/sel  E32 F_XET enum=Xisel
/bus/05/eq  <OFFON> n=1
    /bus/05/eq/on  E32 F_XET enum=OffOn
/bus/05/eq/1  <CHEQ> n=0
    /bus/05/eq/1/type  E32 F_XET enum=Xeqty1
    /bus/05/eq/1/f  F32 F_XET
    /bus/05/eq/1/g  F32 F_XET
    /bus/05/eq/1/q  F32 F_XET
/bus/05/eq/2  <CHEQ> n=0
    /bus/05/eq/2/type  E32 F_XET enum=Xeqty1
    /bus/05/eq/2/f  F32 F_XET
    /bus/05/eq/2/g  F32 F_XET
    /bus/05/eq/2/q  F32 F_XET
/bus/05/eq/3  <CHEQ> n=0
    /bus/05/eq/3/type  E32 F_XET enum=Xeqty1
    /bus/05/eq/3/f  F32 F_XET
    /bus/05/eq/3/g  F32 F_XET
    /bus/05/eq/3/q  F32 F_XET
/bus/05/eq/4  <CHEQ> n=0
    /bus/05/eq/4/type  E32 F_XET enum=Xeqty1
    /bus/05/eq/4/f  F32 F_XET
    /bus/05/eq/4/g  F32 F_XET
    /bus/05/eq/4/q  F32 F_XET
/bus/05/eq/5  <CHEQ> n=0
    /bus/05/eq/5/type  E32 F_XET enum=Xeqty1
    /bus/05/eq/5/f  F32 F_XET
    /bus/05/eq/5/g  F32 F_XET
    /bus/05/eq/5/q  F32 F_XET
/bus/05/eq/6  <CHEQ> n=0
    /bus/05/eq/6/type  E32 F_XET enum=Xeqty1
    /bus/05/eq/6/f  F32 F_XET
    /bus/05/eq/6/g  F32 F_XET
    /bus/05/eq/6/q  F32 F_XET
/bus/05/mix  <CHMX> n=0
    /bus/05/mix/on  E32 F_XET enum=OffOn
    /bus/05/mix/fader  F32 F_XET
    /bus/05/mix/st  E32 F_XET enum=OffOn
    /bus/05/mix/pan  F32 F_XET
    /bus/05/mix/mono  E32 F_XET enum=OffOn
    /bus/05/mix/mlevel  F32 F_XET
/bus/05/mix/01  <CHMO> n=0
    /bus/05/mix/01/on  E32 F_XET enum=OffOn
    /bus/05/mix/01/level  F32 F_XET
    /bus/05/mix/01/pan  F32 F_XET
    /bus/05/mix/01/type  E32 F_XET enum=OffOn
    /bus/05/mix/01/panFollow  E32 F_XET
/bus/05/mix/02  <CHME> n=0
    /bus/05/mix/02/on  E32 F_XET enum=OffOn
    /bus/05/mix/02/level  F32 F_XET
/bus/05/mix/03  <CHMO> n=0
    /bus/05/mix/03/on  E32 F_XET enum=OffOn
    /bus/05/mix/03/level  F32 F_XET
    /bus/05/mix/03/pan  F32 F_XET
    /bus/05/mix/03/type  E32 F_XET enum=OffOn
    /bus/05/mix/03/panFollow  E32 F_XET
/bus/05/mix/04  <CHME> n=0
    /bus/05/mix/04/on  E32 F_XET enum=OffOn
    /bus/05/mix/04/level  F32 F_XET
/bus/05/mix/05  <CHMO> n=0
    /bus/05/mix/05/on  E32 F_XET enum=OffOn
    /bus/05/mix/05/level  F32 F_XET
    /bus/05/mix/05/pan  F32 F_XET
    /bus/05/mix/05/type  E32 F_XET enum=OffOn
    /bus/05/mix/05/panFollow  E32 F_XET
/bus/05/mix/06  <CHME> n=0
    /bus/05/mix/06/on  E32 F_XET enum=OffOn
    /bus/05/mix/06/level  F32 F_XET
/bus/05/grp  <CHGRP> n=0
    /bus/05/grp/dca  P32 F_XET
    /bus/05/grp/mute  P32 F_XET
```

### Xbus06 (X32Bus.h, 99 entries)

```
/bus  <BSCO> n=0
/bus/06  <BSCO> n=0
/bus/06/config  <BSCO> n=0
    /bus/06/config/name  S32 F_XET
    /bus/06/config/icon  I32 F_XET
    /bus/06/config/color  E32 F_XET enum=Xcolors
/bus/06/dyn  <CHDY> n=0
    /bus/06/dyn/on  E32 F_XET enum=OffOn
    /bus/06/dyn/mode  E32 F_XET enum=Xdymode
    /bus/06/dyn/det  E32 F_XET enum=Xdydet
    /bus/06/dyn/env  E32 F_XET enum=Xdyenv
    /bus/06/dyn/thr  F32 F_XET
    /bus/06/dyn/ratio  E32 F_XET enum=OffOn
    /bus/06/dyn/knee  F32 F_XET
    /bus/06/dyn/mgain  F32 F_XET
    /bus/06/dyn/attack  F32 F_XET
    /bus/06/dyn/hold  F32 F_XET
    /bus/06/dyn/release  F32 F_XET
    /bus/06/dyn/pos  E32 F_XET enum=Xdyppos
    /bus/06/dyn/keysrc  I32 F_XET
    /bus/06/dyn/mix  F32 F_XET enum=OffOn
    /bus/06/dyn/auto  E32 F_XET enum=OffOn
/bus/06/dyn/filter  <CHDF> n=0
    /bus/06/dyn/filter/on  E32 F_XET enum=OffOn
    /bus/06/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /bus/06/dyn/filter/f  F32 F_XET
/bus/06/insert  <CHIN> n=0
    /bus/06/insert/on  E32 F_XET enum=OffOn
    /bus/06/insert/pos  E32 F_XET enum=Xdyppos
    /bus/06/insert/sel  E32 F_XET enum=Xisel
/bus/06/eq  <OFFON> n=1
    /bus/06/eq/on  E32 F_XET enum=OffOn
/bus/06/eq/1  <CHEQ> n=0
    /bus/06/eq/1/type  E32 F_XET enum=Xeqty1
    /bus/06/eq/1/f  F32 F_XET
    /bus/06/eq/1/g  F32 F_XET
    /bus/06/eq/1/q  F32 F_XET
/bus/06/eq/2  <CHEQ> n=0
    /bus/06/eq/2/type  E32 F_XET enum=Xeqty1
    /bus/06/eq/2/f  F32 F_XET
    /bus/06/eq/2/g  F32 F_XET
    /bus/06/eq/2/q  F32 F_XET
/bus/06/eq/3  <CHEQ> n=0
    /bus/06/eq/3/type  E32 F_XET enum=Xeqty1
    /bus/06/eq/3/f  F32 F_XET
    /bus/06/eq/3/g  F32 F_XET
    /bus/06/eq/3/q  F32 F_XET
/bus/06/eq/4  <CHEQ> n=0
    /bus/06/eq/4/type  E32 F_XET enum=Xeqty1
    /bus/06/eq/4/f  F32 F_XET
    /bus/06/eq/4/g  F32 F_XET
    /bus/06/eq/4/q  F32 F_XET
/bus/06/eq/5  <CHEQ> n=0
    /bus/06/eq/5/type  E32 F_XET enum=Xeqty1
    /bus/06/eq/5/f  F32 F_XET
    /bus/06/eq/5/g  F32 F_XET
    /bus/06/eq/5/q  F32 F_XET
/bus/06/eq/6  <CHEQ> n=0
    /bus/06/eq/6/type  E32 F_XET enum=Xeqty1
    /bus/06/eq/6/f  F32 F_XET
    /bus/06/eq/6/g  F32 F_XET
    /bus/06/eq/6/q  F32 F_XET
/bus/06/mix  <CHMX> n=0
    /bus/06/mix/on  E32 F_XET enum=OffOn
    /bus/06/mix/fader  F32 F_XET
    /bus/06/mix/st  E32 F_XET enum=OffOn
    /bus/06/mix/pan  F32 F_XET
    /bus/06/mix/mono  E32 F_XET enum=OffOn
    /bus/06/mix/mlevel  F32 F_XET
/bus/06/mix/01  <CHMO> n=0
    /bus/06/mix/01/on  E32 F_XET enum=OffOn
    /bus/06/mix/01/level  F32 F_XET
    /bus/06/mix/01/pan  F32 F_XET
    /bus/06/mix/01/type  E32 F_XET enum=OffOn
    /bus/06/mix/01/panFollow  E32 F_XET
/bus/06/mix/02  <CHME> n=0
    /bus/06/mix/02/on  E32 F_XET enum=OffOn
    /bus/06/mix/02/level  F32 F_XET
/bus/06/mix/03  <CHMO> n=0
    /bus/06/mix/03/on  E32 F_XET enum=OffOn
    /bus/06/mix/03/level  F32 F_XET
    /bus/06/mix/03/pan  F32 F_XET
    /bus/06/mix/03/type  E32 F_XET enum=OffOn
    /bus/06/mix/03/panFollow  E32 F_XET
/bus/06/mix/04  <CHME> n=0
    /bus/06/mix/04/on  E32 F_XET enum=OffOn
    /bus/06/mix/04/level  F32 F_XET
/bus/06/mix/05  <CHMO> n=0
    /bus/06/mix/05/on  E32 F_XET enum=OffOn
    /bus/06/mix/05/level  F32 F_XET
    /bus/06/mix/05/pan  F32 F_XET
    /bus/06/mix/05/type  E32 F_XET enum=OffOn
    /bus/06/mix/05/panFollow  E32 F_XET
/bus/06/mix/06  <CHME> n=0
    /bus/06/mix/06/on  E32 F_XET enum=OffOn
    /bus/06/mix/06/level  F32 F_XET
/bus/06/grp  <CHGRP> n=0
    /bus/06/grp/dca  P32 F_XET
    /bus/06/grp/mute  P32 F_XET
```

### Xbus07 (X32Bus.h, 99 entries)

```
/bus  <BSCO> n=0
/bus/07  <BSCO> n=0
/bus/07/config  <BSCO> n=0
    /bus/07/config/name  S32 F_XET
    /bus/07/config/icon  I32 F_XET
    /bus/07/config/color  E32 F_XET enum=Xcolors
/bus/07/dyn  <CHDY> n=0
    /bus/07/dyn/on  E32 F_XET enum=OffOn
    /bus/07/dyn/mode  E32 F_XET enum=Xdymode
    /bus/07/dyn/det  E32 F_XET enum=Xdydet
    /bus/07/dyn/env  E32 F_XET enum=Xdyenv
    /bus/07/dyn/thr  F32 F_XET
    /bus/07/dyn/ratio  E32 F_XET enum=OffOn
    /bus/07/dyn/knee  F32 F_XET
    /bus/07/dyn/mgain  F32 F_XET
    /bus/07/dyn/attack  F32 F_XET
    /bus/07/dyn/hold  F32 F_XET
    /bus/07/dyn/release  F32 F_XET
    /bus/07/dyn/pos  E32 F_XET enum=Xdyppos
    /bus/07/dyn/keysrc  I32 F_XET
    /bus/07/dyn/mix  F32 F_XET enum=OffOn
    /bus/07/dyn/auto  E32 F_XET enum=OffOn
/bus/07/dyn/filter  <CHDF> n=0
    /bus/07/dyn/filter/on  E32 F_XET enum=OffOn
    /bus/07/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /bus/07/dyn/filter/f  F32 F_XET
/bus/07/insert  <CHIN> n=0
    /bus/07/insert/on  E32 F_XET enum=OffOn
    /bus/07/insert/pos  E32 F_XET enum=Xdyppos
    /bus/07/insert/sel  E32 F_XET enum=Xisel
/bus/07/eq  <OFFON> n=1
    /bus/07/eq/on  E32 F_XET enum=OffOn
/bus/07/eq/1  <CHEQ> n=0
    /bus/07/eq/1/type  E32 F_XET enum=Xeqty1
    /bus/07/eq/1/f  F32 F_XET
    /bus/07/eq/1/g  F32 F_XET
    /bus/07/eq/1/q  F32 F_XET
/bus/07/eq/2  <CHEQ> n=0
    /bus/07/eq/2/type  E32 F_XET enum=Xeqty1
    /bus/07/eq/2/f  F32 F_XET
    /bus/07/eq/2/g  F32 F_XET
    /bus/07/eq/2/q  F32 F_XET
/bus/07/eq/3  <CHEQ> n=0
    /bus/07/eq/3/type  E32 F_XET enum=Xeqty1
    /bus/07/eq/3/f  F32 F_XET
    /bus/07/eq/3/g  F32 F_XET
    /bus/07/eq/3/q  F32 F_XET
/bus/07/eq/4  <CHEQ> n=0
    /bus/07/eq/4/type  E32 F_XET enum=Xeqty1
    /bus/07/eq/4/f  F32 F_XET
    /bus/07/eq/4/g  F32 F_XET
    /bus/07/eq/4/q  F32 F_XET
/bus/07/eq/5  <CHEQ> n=0
    /bus/07/eq/5/type  E32 F_XET enum=Xeqty1
    /bus/07/eq/5/f  F32 F_XET
    /bus/07/eq/5/g  F32 F_XET
    /bus/07/eq/5/q  F32 F_XET
/bus/07/eq/6  <CHEQ> n=0
    /bus/07/eq/6/type  E32 F_XET enum=Xeqty1
    /bus/07/eq/6/f  F32 F_XET
    /bus/07/eq/6/g  F32 F_XET
    /bus/07/eq/6/q  F32 F_XET
/bus/07/mix  <CHMX> n=0
    /bus/07/mix/on  E32 F_XET enum=OffOn
    /bus/07/mix/fader  F32 F_XET
    /bus/07/mix/st  E32 F_XET enum=OffOn
    /bus/07/mix/pan  F32 F_XET
    /bus/07/mix/mono  E32 F_XET enum=OffOn
    /bus/07/mix/mlevel  F32 F_XET
/bus/07/mix/01  <CHMO> n=0
    /bus/07/mix/01/on  E32 F_XET enum=OffOn
    /bus/07/mix/01/level  F32 F_XET
    /bus/07/mix/01/pan  F32 F_XET
    /bus/07/mix/01/type  E32 F_XET enum=OffOn
    /bus/07/mix/01/panFollow  E32 F_XET
/bus/07/mix/02  <CHME> n=0
    /bus/07/mix/02/on  E32 F_XET enum=OffOn
    /bus/07/mix/02/level  F32 F_XET
/bus/07/mix/03  <CHMO> n=0
    /bus/07/mix/03/on  E32 F_XET enum=OffOn
    /bus/07/mix/03/level  F32 F_XET
    /bus/07/mix/03/pan  F32 F_XET
    /bus/07/mix/03/type  E32 F_XET enum=OffOn
    /bus/07/mix/03/panFollow  E32 F_XET
/bus/07/mix/04  <CHME> n=0
    /bus/07/mix/04/on  E32 F_XET enum=OffOn
    /bus/07/mix/04/level  F32 F_XET
/bus/07/mix/05  <CHMO> n=0
    /bus/07/mix/05/on  E32 F_XET enum=OffOn
    /bus/07/mix/05/level  F32 F_XET
    /bus/07/mix/05/pan  F32 F_XET
    /bus/07/mix/05/type  E32 F_XET enum=OffOn
    /bus/07/mix/05/panFollow  E32 F_XET
/bus/07/mix/06  <CHME> n=0
    /bus/07/mix/06/on  E32 F_XET enum=OffOn
    /bus/07/mix/06/level  F32 F_XET
/bus/07/grp  <CHGRP> n=0
    /bus/07/grp/dca  P32 F_XET
    /bus/07/grp/mute  P32 F_XET
```

### Xbus08 (X32Bus.h, 99 entries)

```
/bus  <BSCO> n=0
/bus/08  <BSCO> n=0
/bus/08/config  <BSCO> n=0
    /bus/08/config/name  S32 F_XET
    /bus/08/config/icon  I32 F_XET
    /bus/08/config/color  E32 F_XET enum=Xcolors
/bus/08/dyn  <CHDY> n=0
    /bus/08/dyn/on  E32 F_XET enum=OffOn
    /bus/08/dyn/mode  E32 F_XET enum=Xdymode
    /bus/08/dyn/det  E32 F_XET enum=Xdydet
    /bus/08/dyn/env  E32 F_XET enum=Xdyenv
    /bus/08/dyn/thr  F32 F_XET
    /bus/08/dyn/ratio  E32 F_XET enum=OffOn
    /bus/08/dyn/knee  F32 F_XET
    /bus/08/dyn/mgain  F32 F_XET
    /bus/08/dyn/attack  F32 F_XET
    /bus/08/dyn/hold  F32 F_XET
    /bus/08/dyn/release  F32 F_XET
    /bus/08/dyn/pos  E32 F_XET enum=Xdyppos
    /bus/08/dyn/keysrc  I32 F_XET
    /bus/08/dyn/mix  F32 F_XET enum=OffOn
    /bus/08/dyn/auto  E32 F_XET enum=OffOn
/bus/08/dyn/filter  <CHDF> n=0
    /bus/08/dyn/filter/on  E32 F_XET enum=OffOn
    /bus/08/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /bus/08/dyn/filter/f  F32 F_XET
/bus/08/insert  <CHIN> n=0
    /bus/08/insert/on  E32 F_XET enum=OffOn
    /bus/08/insert/pos  E32 F_XET enum=Xdyppos
    /bus/08/insert/sel  E32 F_XET enum=Xisel
/bus/08/eq  <OFFON> n=1
    /bus/08/eq/on  E32 F_XET enum=OffOn
/bus/08/eq/1  <CHEQ> n=0
    /bus/08/eq/1/type  E32 F_XET enum=Xeqty1
    /bus/08/eq/1/f  F32 F_XET
    /bus/08/eq/1/g  F32 F_XET
    /bus/08/eq/1/q  F32 F_XET
/bus/08/eq/2  <CHEQ> n=0
    /bus/08/eq/2/type  E32 F_XET enum=Xeqty1
    /bus/08/eq/2/f  F32 F_XET
    /bus/08/eq/2/g  F32 F_XET
    /bus/08/eq/2/q  F32 F_XET
/bus/08/eq/3  <CHEQ> n=0
    /bus/08/eq/3/type  E32 F_XET enum=Xeqty1
    /bus/08/eq/3/f  F32 F_XET
    /bus/08/eq/3/g  F32 F_XET
    /bus/08/eq/3/q  F32 F_XET
/bus/08/eq/4  <CHEQ> n=0
    /bus/08/eq/4/type  E32 F_XET enum=Xeqty1
    /bus/08/eq/4/f  F32 F_XET
    /bus/08/eq/4/g  F32 F_XET
    /bus/08/eq/4/q  F32 F_XET
/bus/08/eq/5  <CHEQ> n=0
    /bus/08/eq/5/type  E32 F_XET enum=Xeqty1
    /bus/08/eq/5/f  F32 F_XET
    /bus/08/eq/5/g  F32 F_XET
    /bus/08/eq/5/q  F32 F_XET
/bus/08/eq/6  <CHEQ> n=0
    /bus/08/eq/6/type  E32 F_XET enum=Xeqty1
    /bus/08/eq/6/f  F32 F_XET
    /bus/08/eq/6/g  F32 F_XET
    /bus/08/eq/6/q  F32 F_XET
/bus/08/mix  <CHMX> n=0
    /bus/08/mix/on  E32 F_XET enum=OffOn
    /bus/08/mix/fader  F32 F_XET
    /bus/08/mix/st  E32 F_XET enum=OffOn
    /bus/08/mix/pan  F32 F_XET
    /bus/08/mix/mono  E32 F_XET enum=OffOn
    /bus/08/mix/mlevel  F32 F_XET
/bus/08/mix/01  <CHMO> n=0
    /bus/08/mix/01/on  E32 F_XET enum=OffOn
    /bus/08/mix/01/level  F32 F_XET
    /bus/08/mix/01/pan  F32 F_XET
    /bus/08/mix/01/type  E32 F_XET enum=OffOn
    /bus/08/mix/01/panFollow  E32 F_XET
/bus/08/mix/02  <CHME> n=0
    /bus/08/mix/02/on  E32 F_XET enum=OffOn
    /bus/08/mix/02/level  F32 F_XET
/bus/08/mix/03  <CHMO> n=0
    /bus/08/mix/03/on  E32 F_XET enum=OffOn
    /bus/08/mix/03/level  F32 F_XET
    /bus/08/mix/03/pan  F32 F_XET
    /bus/08/mix/03/type  E32 F_XET enum=OffOn
    /bus/08/mix/03/panFollow  E32 F_XET
/bus/08/mix/04  <CHME> n=0
    /bus/08/mix/04/on  E32 F_XET enum=OffOn
    /bus/08/mix/04/level  F32 F_XET
/bus/08/mix/05  <CHMO> n=0
    /bus/08/mix/05/on  E32 F_XET enum=OffOn
    /bus/08/mix/05/level  F32 F_XET
    /bus/08/mix/05/pan  F32 F_XET
    /bus/08/mix/05/type  E32 F_XET enum=OffOn
    /bus/08/mix/05/panFollow  E32 F_XET
/bus/08/mix/06  <CHME> n=0
    /bus/08/mix/06/on  E32 F_XET enum=OffOn
    /bus/08/mix/06/level  F32 F_XET
/bus/08/grp  <CHGRP> n=0
    /bus/08/grp/dca  P32 F_XET
    /bus/08/grp/mute  P32 F_XET
```

### Xbus09 (X32Bus.h, 99 entries)

```
/bus  <BSCO> n=0
/bus/09  <BSCO> n=0
/bus/09/config  <BSCO> n=0
    /bus/09/config/name  S32 F_XET
    /bus/09/config/icon  I32 F_XET
    /bus/09/config/color  E32 F_XET enum=Xcolors
/bus/09/dyn  <CHDY> n=0
    /bus/09/dyn/on  E32 F_XET enum=OffOn
    /bus/09/dyn/mode  E32 F_XET enum=Xdymode
    /bus/09/dyn/det  E32 F_XET enum=Xdydet
    /bus/09/dyn/env  E32 F_XET enum=Xdyenv
    /bus/09/dyn/thr  F32 F_XET
    /bus/09/dyn/ratio  E32 F_XET enum=OffOn
    /bus/09/dyn/knee  F32 F_XET
    /bus/09/dyn/mgain  F32 F_XET
    /bus/09/dyn/attack  F32 F_XET
    /bus/09/dyn/hold  F32 F_XET
    /bus/09/dyn/release  F32 F_XET
    /bus/09/dyn/pos  E32 F_XET enum=Xdyppos
    /bus/09/dyn/keysrc  I32 F_XET
    /bus/09/dyn/mix  F32 F_XET enum=OffOn
    /bus/09/dyn/auto  E32 F_XET enum=OffOn
/bus/09/dyn/filter  <CHDF> n=0
    /bus/09/dyn/filter/on  E32 F_XET enum=OffOn
    /bus/09/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /bus/09/dyn/filter/f  F32 F_XET
/bus/09/insert  <CHIN> n=0
    /bus/09/insert/on  E32 F_XET enum=OffOn
    /bus/09/insert/pos  E32 F_XET enum=Xdyppos
    /bus/09/insert/sel  E32 F_XET enum=Xisel
/bus/09/eq  <OFFON> n=1
    /bus/09/eq/on  E32 F_XET enum=OffOn
/bus/09/eq/1  <CHEQ> n=0
    /bus/09/eq/1/type  E32 F_XET enum=Xeqty1
    /bus/09/eq/1/f  F32 F_XET
    /bus/09/eq/1/g  F32 F_XET
    /bus/09/eq/1/q  F32 F_XET
/bus/09/eq/2  <CHEQ> n=0
    /bus/09/eq/2/type  E32 F_XET enum=Xeqty1
    /bus/09/eq/2/f  F32 F_XET
    /bus/09/eq/2/g  F32 F_XET
    /bus/09/eq/2/q  F32 F_XET
/bus/09/eq/3  <CHEQ> n=0
    /bus/09/eq/3/type  E32 F_XET enum=Xeqty1
    /bus/09/eq/3/f  F32 F_XET
    /bus/09/eq/3/g  F32 F_XET
    /bus/09/eq/3/q  F32 F_XET
/bus/09/eq/4  <CHEQ> n=0
    /bus/09/eq/4/type  E32 F_XET enum=Xeqty1
    /bus/09/eq/4/f  F32 F_XET
    /bus/09/eq/4/g  F32 F_XET
    /bus/09/eq/4/q  F32 F_XET
/bus/09/eq/5  <CHEQ> n=0
    /bus/09/eq/5/type  E32 F_XET enum=Xeqty1
    /bus/09/eq/5/f  F32 F_XET
    /bus/09/eq/5/g  F32 F_XET
    /bus/09/eq/5/q  F32 F_XET
/bus/09/eq/6  <CHEQ> n=0
    /bus/09/eq/6/type  E32 F_XET enum=Xeqty1
    /bus/09/eq/6/f  F32 F_XET
    /bus/09/eq/6/g  F32 F_XET
    /bus/09/eq/6/q  F32 F_XET
/bus/09/mix  <CHMX> n=0
    /bus/09/mix/on  E32 F_XET enum=OffOn
    /bus/09/mix/fader  F32 F_XET
    /bus/09/mix/st  E32 F_XET enum=OffOn
    /bus/09/mix/pan  F32 F_XET
    /bus/09/mix/mono  E32 F_XET enum=OffOn
    /bus/09/mix/mlevel  F32 F_XET
/bus/09/mix/01  <CHMO> n=0
    /bus/09/mix/01/on  E32 F_XET enum=OffOn
    /bus/09/mix/01/level  F32 F_XET
    /bus/09/mix/01/pan  F32 F_XET
    /bus/09/mix/01/type  E32 F_XET enum=OffOn
    /bus/09/mix/01/panFollow  E32 F_XET
/bus/09/mix/02  <CHME> n=0
    /bus/09/mix/02/on  E32 F_XET enum=OffOn
    /bus/09/mix/02/level  F32 F_XET
/bus/09/mix/03  <CHMO> n=0
    /bus/09/mix/03/on  E32 F_XET enum=OffOn
    /bus/09/mix/03/level  F32 F_XET
    /bus/09/mix/03/pan  F32 F_XET
    /bus/09/mix/03/type  E32 F_XET enum=OffOn
    /bus/09/mix/03/panFollow  E32 F_XET
/bus/09/mix/04  <CHME> n=0
    /bus/09/mix/04/on  E32 F_XET enum=OffOn
    /bus/09/mix/04/level  F32 F_XET
/bus/09/mix/05  <CHMO> n=0
    /bus/09/mix/05/on  E32 F_XET enum=OffOn
    /bus/09/mix/05/level  F32 F_XET
    /bus/09/mix/05/pan  F32 F_XET
    /bus/09/mix/05/type  E32 F_XET enum=OffOn
    /bus/09/mix/05/panFollow  E32 F_XET
/bus/09/mix/06  <CHME> n=0
    /bus/09/mix/06/on  E32 F_XET enum=OffOn
    /bus/09/mix/06/level  F32 F_XET
/bus/09/grp  <CHGRP> n=0
    /bus/09/grp/dca  P32 F_XET
    /bus/09/grp/mute  P32 F_XET
```

### Xbus10 (X32Bus.h, 99 entries)

```
/bus  <BSCO> n=0
/bus/10  <BSCO> n=0
/bus/10/config  <BSCO> n=0
    /bus/10/config/name  S32 F_XET
    /bus/10/config/icon  I32 F_XET
    /bus/10/config/color  E32 F_XET enum=Xcolors
/bus/10/dyn  <CHDY> n=0
    /bus/10/dyn/on  E32 F_XET enum=OffOn
    /bus/10/dyn/mode  E32 F_XET enum=Xdymode
    /bus/10/dyn/det  E32 F_XET enum=Xdydet
    /bus/10/dyn/env  E32 F_XET enum=Xdyenv
    /bus/10/dyn/thr  F32 F_XET
    /bus/10/dyn/ratio  E32 F_XET enum=OffOn
    /bus/10/dyn/knee  F32 F_XET
    /bus/10/dyn/mgain  F32 F_XET
    /bus/10/dyn/attack  F32 F_XET
    /bus/10/dyn/hold  F32 F_XET
    /bus/10/dyn/release  F32 F_XET
    /bus/10/dyn/pos  E32 F_XET enum=Xdyppos
    /bus/10/dyn/keysrc  I32 F_XET
    /bus/10/dyn/mix  F32 F_XET enum=OffOn
    /bus/10/dyn/auto  E32 F_XET enum=OffOn
/bus/10/dyn/filter  <CHDF> n=0
    /bus/10/dyn/filter/on  E32 F_XET enum=OffOn
    /bus/10/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /bus/10/dyn/filter/f  F32 F_XET
/bus/10/insert  <CHIN> n=0
    /bus/10/insert/on  E32 F_XET enum=OffOn
    /bus/10/insert/pos  E32 F_XET enum=Xdyppos
    /bus/10/insert/sel  E32 F_XET enum=Xisel
/bus/10/eq  <OFFON> n=1
    /bus/10/eq/on  E32 F_XET enum=OffOn
/bus/10/eq/1  <CHEQ> n=0
    /bus/10/eq/1/type  E32 F_XET enum=Xeqty1
    /bus/10/eq/1/f  F32 F_XET
    /bus/10/eq/1/g  F32 F_XET
    /bus/10/eq/1/q  F32 F_XET
/bus/10/eq/2  <CHEQ> n=0
    /bus/10/eq/2/type  E32 F_XET enum=Xeqty1
    /bus/10/eq/2/f  F32 F_XET
    /bus/10/eq/2/g  F32 F_XET
    /bus/10/eq/2/q  F32 F_XET
/bus/10/eq/3  <CHEQ> n=0
    /bus/10/eq/3/type  E32 F_XET enum=Xeqty1
    /bus/10/eq/3/f  F32 F_XET
    /bus/10/eq/3/g  F32 F_XET
    /bus/10/eq/3/q  F32 F_XET
/bus/10/eq/4  <CHEQ> n=0
    /bus/10/eq/4/type  E32 F_XET enum=Xeqty1
    /bus/10/eq/4/f  F32 F_XET
    /bus/10/eq/4/g  F32 F_XET
    /bus/10/eq/4/q  F32 F_XET
/bus/10/eq/5  <CHEQ> n=0
    /bus/10/eq/5/type  E32 F_XET enum=Xeqty1
    /bus/10/eq/5/f  F32 F_XET
    /bus/10/eq/5/g  F32 F_XET
    /bus/10/eq/5/q  F32 F_XET
/bus/10/eq/6  <CHEQ> n=0
    /bus/10/eq/6/type  E32 F_XET enum=Xeqty1
    /bus/10/eq/6/f  F32 F_XET
    /bus/10/eq/6/g  F32 F_XET
    /bus/10/eq/6/q  F32 F_XET
/bus/10/mix  <CHMX> n=0
    /bus/10/mix/on  E32 F_XET enum=OffOn
    /bus/10/mix/fader  F32 F_XET
    /bus/10/mix/st  E32 F_XET enum=OffOn
    /bus/10/mix/pan  F32 F_XET
    /bus/10/mix/mono  E32 F_XET enum=OffOn
    /bus/10/mix/mlevel  F32 F_XET
/bus/10/mix/01  <CHMO> n=0
    /bus/10/mix/01/on  E32 F_XET enum=OffOn
    /bus/10/mix/01/level  F32 F_XET
    /bus/10/mix/01/pan  F32 F_XET
    /bus/10/mix/01/type  E32 F_XET enum=OffOn
    /bus/10/mix/01/panFollow  E32 F_XET
/bus/10/mix/02  <CHME> n=0
    /bus/10/mix/02/on  E32 F_XET enum=OffOn
    /bus/10/mix/02/level  F32 F_XET
/bus/10/mix/03  <CHMO> n=0
    /bus/10/mix/03/on  E32 F_XET enum=OffOn
    /bus/10/mix/03/level  F32 F_XET
    /bus/10/mix/03/pan  F32 F_XET
    /bus/10/mix/03/type  E32 F_XET enum=OffOn
    /bus/10/mix/03/panFollow  E32 F_XET
/bus/10/mix/04  <CHME> n=0
    /bus/10/mix/04/on  E32 F_XET enum=OffOn
    /bus/10/mix/04/level  F32 F_XET
/bus/10/mix/05  <CHMO> n=0
    /bus/10/mix/05/on  E32 F_XET enum=OffOn
    /bus/10/mix/05/level  F32 F_XET
    /bus/10/mix/05/pan  F32 F_XET
    /bus/10/mix/05/type  E32 F_XET enum=OffOn
    /bus/10/mix/05/panFollow  E32 F_XET
/bus/10/mix/06  <CHME> n=0
    /bus/10/mix/06/on  E32 F_XET enum=OffOn
    /bus/10/mix/06/level  F32 F_XET
/bus/10/grp  <CHGRP> n=0
    /bus/10/grp/dca  P32 F_XET
    /bus/10/grp/mute  P32 F_XET
```

### Xbus11 (X32Bus.h, 99 entries)

```
/bus  <BSCO> n=0
/bus/11  <BSCO> n=0
/bus/11/config  <BSCO> n=0
    /bus/11/config/name  S32 F_XET
    /bus/11/config/icon  I32 F_XET
    /bus/11/config/color  E32 F_XET enum=Xcolors
/bus/11/dyn  <CHDY> n=0
    /bus/11/dyn/on  E32 F_XET enum=OffOn
    /bus/11/dyn/mode  E32 F_XET enum=Xdymode
    /bus/11/dyn/det  E32 F_XET enum=Xdydet
    /bus/11/dyn/env  E32 F_XET enum=Xdyenv
    /bus/11/dyn/thr  F32 F_XET
    /bus/11/dyn/ratio  E32 F_XET enum=OffOn
    /bus/11/dyn/knee  F32 F_XET
    /bus/11/dyn/mgain  F32 F_XET
    /bus/11/dyn/attack  F32 F_XET
    /bus/11/dyn/hold  F32 F_XET
    /bus/11/dyn/release  F32 F_XET
    /bus/11/dyn/pos  E32 F_XET enum=Xdyppos
    /bus/11/dyn/keysrc  I32 F_XET
    /bus/11/dyn/mix  F32 F_XET enum=OffOn
    /bus/11/dyn/auto  E32 F_XET enum=OffOn
/bus/11/dyn/filter  <CHDF> n=0
    /bus/11/dyn/filter/on  E32 F_XET enum=OffOn
    /bus/11/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /bus/11/dyn/filter/f  F32 F_XET
/bus/11/insert  <CHIN> n=0
    /bus/11/insert/on  E32 F_XET enum=OffOn
    /bus/11/insert/pos  E32 F_XET enum=Xdyppos
    /bus/11/insert/sel  E32 F_XET enum=Xisel
/bus/11/eq  <OFFON> n=1
    /bus/11/eq/on  E32 F_XET enum=OffOn
/bus/11/eq/1  <CHEQ> n=0
    /bus/11/eq/1/type  E32 F_XET enum=Xeqty1
    /bus/11/eq/1/f  F32 F_XET
    /bus/11/eq/1/g  F32 F_XET
    /bus/11/eq/1/q  F32 F_XET
/bus/11/eq/2  <CHEQ> n=0
    /bus/11/eq/2/type  E32 F_XET enum=Xeqty1
    /bus/11/eq/2/f  F32 F_XET
    /bus/11/eq/2/g  F32 F_XET
    /bus/11/eq/2/q  F32 F_XET
/bus/11/eq/3  <CHEQ> n=0
    /bus/11/eq/3/type  E32 F_XET enum=Xeqty1
    /bus/11/eq/3/f  F32 F_XET
    /bus/11/eq/3/g  F32 F_XET
    /bus/11/eq/3/q  F32 F_XET
/bus/11/eq/4  <CHEQ> n=0
    /bus/11/eq/4/type  E32 F_XET enum=Xeqty1
    /bus/11/eq/4/f  F32 F_XET
    /bus/11/eq/4/g  F32 F_XET
    /bus/11/eq/4/q  F32 F_XET
/bus/11/eq/5  <CHEQ> n=0
    /bus/11/eq/5/type  E32 F_XET enum=Xeqty1
    /bus/11/eq/5/f  F32 F_XET
    /bus/11/eq/5/g  F32 F_XET
    /bus/11/eq/5/q  F32 F_XET
/bus/11/eq/6  <CHEQ> n=0
    /bus/11/eq/6/type  E32 F_XET enum=Xeqty1
    /bus/11/eq/6/f  F32 F_XET
    /bus/11/eq/6/g  F32 F_XET
    /bus/11/eq/6/q  F32 F_XET
/bus/11/mix  <CHMX> n=0
    /bus/11/mix/on  E32 F_XET enum=OffOn
    /bus/11/mix/fader  F32 F_XET
    /bus/11/mix/st  E32 F_XET enum=OffOn
    /bus/11/mix/pan  F32 F_XET
    /bus/11/mix/mono  E32 F_XET enum=OffOn
    /bus/11/mix/mlevel  F32 F_XET
/bus/11/mix/01  <CHMO> n=0
    /bus/11/mix/01/on  E32 F_XET enum=OffOn
    /bus/11/mix/01/level  F32 F_XET
    /bus/11/mix/01/pan  F32 F_XET
    /bus/11/mix/01/type  E32 F_XET enum=OffOn
    /bus/11/mix/01/panFollow  E32 F_XET
/bus/11/mix/02  <CHME> n=0
    /bus/11/mix/02/on  E32 F_XET enum=OffOn
    /bus/11/mix/02/level  F32 F_XET
/bus/11/mix/03  <CHMO> n=0
    /bus/11/mix/03/on  E32 F_XET enum=OffOn
    /bus/11/mix/03/level  F32 F_XET
    /bus/11/mix/03/pan  F32 F_XET
    /bus/11/mix/03/type  E32 F_XET enum=OffOn
    /bus/11/mix/03/panFollow  E32 F_XET
/bus/11/mix/04  <CHME> n=0
    /bus/11/mix/04/on  E32 F_XET enum=OffOn
    /bus/11/mix/04/level  F32 F_XET
/bus/11/mix/05  <CHMO> n=0
    /bus/11/mix/05/on  E32 F_XET enum=OffOn
    /bus/11/mix/05/level  F32 F_XET
    /bus/11/mix/05/pan  F32 F_XET
    /bus/11/mix/05/type  E32 F_XET enum=OffOn
    /bus/11/mix/05/panFollow  E32 F_XET
/bus/11/mix/06  <CHME> n=0
    /bus/11/mix/06/on  E32 F_XET enum=OffOn
    /bus/11/mix/06/level  F32 F_XET
/bus/11/grp  <CHGRP> n=0
    /bus/11/grp/dca  P32 F_XET
    /bus/11/grp/mute  P32 F_XET
```

### Xbus12 (X32Bus.h, 99 entries)

```
/bus  <BSCO> n=0
/bus/12  <BSCO> n=0
/bus/12/config  <BSCO> n=0
    /bus/12/config/name  S32 F_XET
    /bus/12/config/icon  I32 F_XET
    /bus/12/config/color  E32 F_XET enum=Xcolors
/bus/12/dyn  <CHDY> n=0
    /bus/12/dyn/on  E32 F_XET enum=OffOn
    /bus/12/dyn/mode  E32 F_XET enum=Xdymode
    /bus/12/dyn/det  E32 F_XET enum=Xdydet
    /bus/12/dyn/env  E32 F_XET enum=Xdyenv
    /bus/12/dyn/thr  F32 F_XET
    /bus/12/dyn/ratio  E32 F_XET enum=OffOn
    /bus/12/dyn/knee  F32 F_XET
    /bus/12/dyn/mgain  F32 F_XET
    /bus/12/dyn/attack  F32 F_XET
    /bus/12/dyn/hold  F32 F_XET
    /bus/12/dyn/release  F32 F_XET
    /bus/12/dyn/pos  E32 F_XET enum=Xdyppos
    /bus/12/dyn/keysrc  I32 F_XET
    /bus/12/dyn/mix  F32 F_XET enum=OffOn
    /bus/12/dyn/auto  E32 F_XET enum=OffOn
/bus/12/dyn/filter  <CHDF> n=0
    /bus/12/dyn/filter/on  E32 F_XET enum=OffOn
    /bus/12/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /bus/12/dyn/filter/f  F32 F_XET
/bus/12/insert  <CHIN> n=0
    /bus/12/insert/on  E32 F_XET enum=OffOn
    /bus/12/insert/pos  E32 F_XET enum=Xdyppos
    /bus/12/insert/sel  E32 F_XET enum=Xisel
/bus/12/eq  <OFFON> n=1
    /bus/12/eq/on  E32 F_XET enum=OffOn
/bus/12/eq/1  <CHEQ> n=0
    /bus/12/eq/1/type  E32 F_XET enum=Xeqty1
    /bus/12/eq/1/f  F32 F_XET
    /bus/12/eq/1/g  F32 F_XET
    /bus/12/eq/1/q  F32 F_XET
/bus/12/eq/2  <CHEQ> n=0
    /bus/12/eq/2/type  E32 F_XET enum=Xeqty1
    /bus/12/eq/2/f  F32 F_XET
    /bus/12/eq/2/g  F32 F_XET
    /bus/12/eq/2/q  F32 F_XET
/bus/12/eq/3  <CHEQ> n=0
    /bus/12/eq/3/type  E32 F_XET enum=Xeqty1
    /bus/12/eq/3/f  F32 F_XET
    /bus/12/eq/3/g  F32 F_XET
    /bus/12/eq/3/q  F32 F_XET
/bus/12/eq/4  <CHEQ> n=0
    /bus/12/eq/4/type  E32 F_XET enum=Xeqty1
    /bus/12/eq/4/f  F32 F_XET
    /bus/12/eq/4/g  F32 F_XET
    /bus/12/eq/4/q  F32 F_XET
/bus/12/eq/5  <CHEQ> n=0
    /bus/12/eq/5/type  E32 F_XET enum=Xeqty1
    /bus/12/eq/5/f  F32 F_XET
    /bus/12/eq/5/g  F32 F_XET
    /bus/12/eq/5/q  F32 F_XET
/bus/12/eq/6  <CHEQ> n=0
    /bus/12/eq/6/type  E32 F_XET enum=Xeqty1
    /bus/12/eq/6/f  F32 F_XET
    /bus/12/eq/6/g  F32 F_XET
    /bus/12/eq/6/q  F32 F_XET
/bus/12/mix  <CHMX> n=0
    /bus/12/mix/on  E32 F_XET enum=OffOn
    /bus/12/mix/fader  F32 F_XET
    /bus/12/mix/st  E32 F_XET enum=OffOn
    /bus/12/mix/pan  F32 F_XET
    /bus/12/mix/mono  E32 F_XET enum=OffOn
    /bus/12/mix/mlevel  F32 F_XET
/bus/12/mix/01  <CHMO> n=0
    /bus/12/mix/01/on  E32 F_XET enum=OffOn
    /bus/12/mix/01/level  F32 F_XET
    /bus/12/mix/01/pan  F32 F_XET
    /bus/12/mix/01/type  E32 F_XET enum=OffOn
    /bus/12/mix/01/panFollow  E32 F_XET
/bus/12/mix/02  <CHME> n=0
    /bus/12/mix/02/on  E32 F_XET enum=OffOn
    /bus/12/mix/02/level  F32 F_XET
/bus/12/mix/03  <CHMO> n=0
    /bus/12/mix/03/on  E32 F_XET enum=OffOn
    /bus/12/mix/03/level  F32 F_XET
    /bus/12/mix/03/pan  F32 F_XET
    /bus/12/mix/03/type  E32 F_XET enum=OffOn
    /bus/12/mix/03/panFollow  E32 F_XET
/bus/12/mix/04  <CHME> n=0
    /bus/12/mix/04/on  E32 F_XET enum=OffOn
    /bus/12/mix/04/level  F32 F_XET
/bus/12/mix/05  <CHMO> n=0
    /bus/12/mix/05/on  E32 F_XET enum=OffOn
    /bus/12/mix/05/level  F32 F_XET
    /bus/12/mix/05/pan  F32 F_XET
    /bus/12/mix/05/type  E32 F_XET enum=OffOn
    /bus/12/mix/05/panFollow  E32 F_XET
/bus/12/mix/06  <CHME> n=0
    /bus/12/mix/06/on  E32 F_XET enum=OffOn
    /bus/12/mix/06/level  F32 F_XET
/bus/12/grp  <CHGRP> n=0
    /bus/12/grp/dca  P32 F_XET
    /bus/12/grp/mute  P32 F_XET
```

### Xbus13 (X32Bus.h, 99 entries)

```
/bus  <BSCO> n=0
/bus/13  <BSCO> n=0
/bus/13/config  <BSCO> n=0
    /bus/13/config/name  S32 F_XET
    /bus/13/config/icon  I32 F_XET
    /bus/13/config/color  E32 F_XET enum=Xcolors
/bus/13/dyn  <CHDY> n=0
    /bus/13/dyn/on  E32 F_XET enum=OffOn
    /bus/13/dyn/mode  E32 F_XET enum=Xdymode
    /bus/13/dyn/det  E32 F_XET enum=Xdydet
    /bus/13/dyn/env  E32 F_XET enum=Xdyenv
    /bus/13/dyn/thr  F32 F_XET
    /bus/13/dyn/ratio  E32 F_XET enum=OffOn
    /bus/13/dyn/knee  F32 F_XET
    /bus/13/dyn/mgain  F32 F_XET
    /bus/13/dyn/attack  F32 F_XET
    /bus/13/dyn/hold  F32 F_XET
    /bus/13/dyn/release  F32 F_XET
    /bus/13/dyn/pos  E32 F_XET enum=Xdyppos
    /bus/13/dyn/keysrc  I32 F_XET
    /bus/13/dyn/mix  F32 F_XET enum=OffOn
    /bus/13/dyn/auto  E32 F_XET enum=OffOn
/bus/13/dyn/filter  <CHDF> n=0
    /bus/13/dyn/filter/on  E32 F_XET enum=OffOn
    /bus/13/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /bus/13/dyn/filter/f  F32 F_XET
/bus/13/insert  <CHIN> n=0
    /bus/13/insert/on  E32 F_XET enum=OffOn
    /bus/13/insert/pos  E32 F_XET enum=Xdyppos
    /bus/13/insert/sel  E32 F_XET enum=Xisel
/bus/13/eq  <OFFON> n=1
    /bus/13/eq/on  E32 F_XET enum=OffOn
/bus/13/eq/1  <CHEQ> n=0
    /bus/13/eq/1/type  E32 F_XET enum=Xeqty1
    /bus/13/eq/1/f  F32 F_XET
    /bus/13/eq/1/g  F32 F_XET
    /bus/13/eq/1/q  F32 F_XET
/bus/13/eq/2  <CHEQ> n=0
    /bus/13/eq/2/type  E32 F_XET enum=Xeqty1
    /bus/13/eq/2/f  F32 F_XET
    /bus/13/eq/2/g  F32 F_XET
    /bus/13/eq/2/q  F32 F_XET
/bus/13/eq/3  <CHEQ> n=0
    /bus/13/eq/3/type  E32 F_XET enum=Xeqty1
    /bus/13/eq/3/f  F32 F_XET
    /bus/13/eq/3/g  F32 F_XET
    /bus/13/eq/3/q  F32 F_XET
/bus/13/eq/4  <CHEQ> n=0
    /bus/13/eq/4/type  E32 F_XET enum=Xeqty1
    /bus/13/eq/4/f  F32 F_XET
    /bus/13/eq/4/g  F32 F_XET
    /bus/13/eq/4/q  F32 F_XET
/bus/13/eq/5  <CHEQ> n=0
    /bus/13/eq/5/type  E32 F_XET enum=Xeqty1
    /bus/13/eq/5/f  F32 F_XET
    /bus/13/eq/5/g  F32 F_XET
    /bus/13/eq/5/q  F32 F_XET
/bus/13/eq/6  <CHEQ> n=0
    /bus/13/eq/6/type  E32 F_XET enum=Xeqty1
    /bus/13/eq/6/f  F32 F_XET
    /bus/13/eq/6/g  F32 F_XET
    /bus/13/eq/6/q  F32 F_XET
/bus/13/mix  <CHMX> n=0
    /bus/13/mix/on  E32 F_XET enum=OffOn
    /bus/13/mix/fader  F32 F_XET
    /bus/13/mix/st  E32 F_XET enum=OffOn
    /bus/13/mix/pan  F32 F_XET
    /bus/13/mix/mono  E32 F_XET enum=OffOn
    /bus/13/mix/mlevel  F32 F_XET
/bus/13/mix/01  <CHMO> n=0
    /bus/13/mix/01/on  E32 F_XET enum=OffOn
    /bus/13/mix/01/level  F32 F_XET
    /bus/13/mix/01/pan  F32 F_XET
    /bus/13/mix/01/type  E32 F_XET enum=OffOn
    /bus/13/mix/01/panFollow  E32 F_XET
/bus/13/mix/02  <CHME> n=0
    /bus/13/mix/02/on  E32 F_XET enum=OffOn
    /bus/13/mix/02/level  F32 F_XET
/bus/13/mix/03  <CHMO> n=0
    /bus/13/mix/03/on  E32 F_XET enum=OffOn
    /bus/13/mix/03/level  F32 F_XET
    /bus/13/mix/03/pan  F32 F_XET
    /bus/13/mix/03/type  E32 F_XET enum=OffOn
    /bus/13/mix/03/panFollow  E32 F_XET
/bus/13/mix/04  <CHME> n=0
    /bus/13/mix/04/on  E32 F_XET enum=OffOn
    /bus/13/mix/04/level  F32 F_XET
/bus/13/mix/05  <CHMO> n=0
    /bus/13/mix/05/on  E32 F_XET enum=OffOn
    /bus/13/mix/05/level  F32 F_XET
    /bus/13/mix/05/pan  F32 F_XET
    /bus/13/mix/05/type  E32 F_XET enum=OffOn
    /bus/13/mix/05/panFollow  E32 F_XET
/bus/13/mix/06  <CHME> n=0
    /bus/13/mix/06/on  E32 F_XET enum=OffOn
    /bus/13/mix/06/level  F32 F_XET
/bus/13/grp  <CHGRP> n=0
    /bus/13/grp/dca  P32 F_XET
    /bus/13/grp/mute  P32 F_XET
```

### Xbus14 (X32Bus.h, 99 entries)

```
/bus  <BSCO> n=0
/bus/14  <BSCO> n=0
/bus/14/config  <BSCO> n=0
    /bus/14/config/name  S32 F_XET
    /bus/14/config/icon  I32 F_XET
    /bus/14/config/color  E32 F_XET enum=Xcolors
/bus/14/dyn  <CHDY> n=0
    /bus/14/dyn/on  E32 F_XET enum=OffOn
    /bus/14/dyn/mode  E32 F_XET enum=Xdymode
    /bus/14/dyn/det  E32 F_XET enum=Xdydet
    /bus/14/dyn/env  E32 F_XET enum=Xdyenv
    /bus/14/dyn/thr  F32 F_XET
    /bus/14/dyn/ratio  E32 F_XET enum=OffOn
    /bus/14/dyn/knee  F32 F_XET
    /bus/14/dyn/mgain  F32 F_XET
    /bus/14/dyn/attack  F32 F_XET
    /bus/14/dyn/hold  F32 F_XET
    /bus/14/dyn/release  F32 F_XET
    /bus/14/dyn/pos  E32 F_XET enum=Xdyppos
    /bus/14/dyn/keysrc  I32 F_XET
    /bus/14/dyn/mix  F32 F_XET enum=OffOn
    /bus/14/dyn/auto  E32 F_XET enum=OffOn
/bus/14/dyn/filter  <CHDF> n=0
    /bus/14/dyn/filter/on  E32 F_XET enum=OffOn
    /bus/14/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /bus/14/dyn/filter/f  F32 F_XET
/bus/14/insert  <CHIN> n=0
    /bus/14/insert/on  E32 F_XET enum=OffOn
    /bus/14/insert/pos  E32 F_XET enum=Xdyppos
    /bus/14/insert/sel  E32 F_XET enum=Xisel
/bus/14/eq  <OFFON> n=1
    /bus/14/eq/on  E32 F_XET enum=OffOn
/bus/14/eq/1  <CHEQ> n=0
    /bus/14/eq/1/type  E32 F_XET enum=Xeqty1
    /bus/14/eq/1/f  F32 F_XET
    /bus/14/eq/1/g  F32 F_XET
    /bus/14/eq/1/q  F32 F_XET
/bus/14/eq/2  <CHEQ> n=0
    /bus/14/eq/2/type  E32 F_XET enum=Xeqty1
    /bus/14/eq/2/f  F32 F_XET
    /bus/14/eq/2/g  F32 F_XET
    /bus/14/eq/2/q  F32 F_XET
/bus/14/eq/3  <CHEQ> n=0
    /bus/14/eq/3/type  E32 F_XET enum=Xeqty1
    /bus/14/eq/3/f  F32 F_XET
    /bus/14/eq/3/g  F32 F_XET
    /bus/14/eq/3/q  F32 F_XET
/bus/14/eq/4  <CHEQ> n=0
    /bus/14/eq/4/type  E32 F_XET enum=Xeqty1
    /bus/14/eq/4/f  F32 F_XET
    /bus/14/eq/4/g  F32 F_XET
    /bus/14/eq/4/q  F32 F_XET
/bus/14/eq/5  <CHEQ> n=0
    /bus/14/eq/5/type  E32 F_XET enum=Xeqty1
    /bus/14/eq/5/f  F32 F_XET
    /bus/14/eq/5/g  F32 F_XET
    /bus/14/eq/5/q  F32 F_XET
/bus/14/eq/6  <CHEQ> n=0
    /bus/14/eq/6/type  E32 F_XET enum=Xeqty1
    /bus/14/eq/6/f  F32 F_XET
    /bus/14/eq/6/g  F32 F_XET
    /bus/14/eq/6/q  F32 F_XET
/bus/14/mix  <CHMX> n=0
    /bus/14/mix/on  E32 F_XET enum=OffOn
    /bus/14/mix/fader  F32 F_XET
    /bus/14/mix/st  E32 F_XET enum=OffOn
    /bus/14/mix/pan  F32 F_XET
    /bus/14/mix/mono  E32 F_XET enum=OffOn
    /bus/14/mix/mlevel  F32 F_XET
/bus/14/mix/01  <CHMO> n=0
    /bus/14/mix/01/on  E32 F_XET enum=OffOn
    /bus/14/mix/01/level  F32 F_XET
    /bus/14/mix/01/pan  F32 F_XET
    /bus/14/mix/01/type  E32 F_XET enum=OffOn
    /bus/14/mix/01/panFollow  E32 F_XET
/bus/14/mix/02  <CHME> n=0
    /bus/14/mix/02/on  E32 F_XET enum=OffOn
    /bus/14/mix/02/level  F32 F_XET
/bus/14/mix/03  <CHMO> n=0
    /bus/14/mix/03/on  E32 F_XET enum=OffOn
    /bus/14/mix/03/level  F32 F_XET
    /bus/14/mix/03/pan  F32 F_XET
    /bus/14/mix/03/type  E32 F_XET enum=OffOn
    /bus/14/mix/03/panFollow  E32 F_XET
/bus/14/mix/04  <CHME> n=0
    /bus/14/mix/04/on  E32 F_XET enum=OffOn
    /bus/14/mix/04/level  F32 F_XET
/bus/14/mix/05  <CHMO> n=0
    /bus/14/mix/05/on  E32 F_XET enum=OffOn
    /bus/14/mix/05/level  F32 F_XET
    /bus/14/mix/05/pan  F32 F_XET
    /bus/14/mix/05/type  E32 F_XET enum=OffOn
    /bus/14/mix/05/panFollow  E32 F_XET
/bus/14/mix/06  <CHME> n=0
    /bus/14/mix/06/on  E32 F_XET enum=OffOn
    /bus/14/mix/06/level  F32 F_XET
/bus/14/grp  <CHGRP> n=0
    /bus/14/grp/dca  P32 F_XET
    /bus/14/grp/mute  P32 F_XET
```

### Xbus15 (X32Bus.h, 99 entries)

```
/bus  <BSCO> n=0
/bus/15  <BSCO> n=0
/bus/15/config  <BSCO> n=0
    /bus/15/config/name  S32 F_XET
    /bus/15/config/icon  I32 F_XET
    /bus/15/config/color  E32 F_XET enum=Xcolors
/bus/15/dyn  <CHDY> n=0
    /bus/15/dyn/on  E32 F_XET enum=OffOn
    /bus/15/dyn/mode  E32 F_XET enum=Xdymode
    /bus/15/dyn/det  E32 F_XET enum=Xdydet
    /bus/15/dyn/env  E32 F_XET enum=Xdyenv
    /bus/15/dyn/thr  F32 F_XET
    /bus/15/dyn/ratio  E32 F_XET enum=OffOn
    /bus/15/dyn/knee  F32 F_XET
    /bus/15/dyn/mgain  F32 F_XET
    /bus/15/dyn/attack  F32 F_XET
    /bus/15/dyn/hold  F32 F_XET
    /bus/15/dyn/release  F32 F_XET
    /bus/15/dyn/pos  E32 F_XET enum=Xdyppos
    /bus/15/dyn/keysrc  I32 F_XET
    /bus/15/dyn/mix  F32 F_XET enum=OffOn
    /bus/15/dyn/auto  E32 F_XET enum=OffOn
/bus/15/dyn/filter  <CHDF> n=0
    /bus/15/dyn/filter/on  E32 F_XET enum=OffOn
    /bus/15/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /bus/15/dyn/filter/f  F32 F_XET
/bus/15/insert  <CHIN> n=0
    /bus/15/insert/on  E32 F_XET enum=OffOn
    /bus/15/insert/pos  E32 F_XET enum=Xdyppos
    /bus/15/insert/sel  E32 F_XET enum=Xisel
/bus/15/eq  <OFFON> n=1
    /bus/15/eq/on  E32 F_XET enum=OffOn
/bus/15/eq/1  <CHEQ> n=0
    /bus/15/eq/1/type  E32 F_XET enum=Xeqty1
    /bus/15/eq/1/f  F32 F_XET
    /bus/15/eq/1/g  F32 F_XET
    /bus/15/eq/1/q  F32 F_XET
/bus/15/eq/2  <CHEQ> n=0
    /bus/15/eq/2/type  E32 F_XET enum=Xeqty1
    /bus/15/eq/2/f  F32 F_XET
    /bus/15/eq/2/g  F32 F_XET
    /bus/15/eq/2/q  F32 F_XET
/bus/15/eq/3  <CHEQ> n=0
    /bus/15/eq/3/type  E32 F_XET enum=Xeqty1
    /bus/15/eq/3/f  F32 F_XET
    /bus/15/eq/3/g  F32 F_XET
    /bus/15/eq/3/q  F32 F_XET
/bus/15/eq/4  <CHEQ> n=0
    /bus/15/eq/4/type  E32 F_XET enum=Xeqty1
    /bus/15/eq/4/f  F32 F_XET
    /bus/15/eq/4/g  F32 F_XET
    /bus/15/eq/4/q  F32 F_XET
/bus/15/eq/5  <CHEQ> n=0
    /bus/15/eq/5/type  E32 F_XET enum=Xeqty1
    /bus/15/eq/5/f  F32 F_XET
    /bus/15/eq/5/g  F32 F_XET
    /bus/15/eq/5/q  F32 F_XET
/bus/15/eq/6  <CHEQ> n=0
    /bus/15/eq/6/type  E32 F_XET enum=Xeqty1
    /bus/15/eq/6/f  F32 F_XET
    /bus/15/eq/6/g  F32 F_XET
    /bus/15/eq/6/q  F32 F_XET
/bus/15/mix  <CHMX> n=0
    /bus/15/mix/on  E32 F_XET enum=OffOn
    /bus/15/mix/fader  F32 F_XET
    /bus/15/mix/st  E32 F_XET enum=OffOn
    /bus/15/mix/pan  F32 F_XET
    /bus/15/mix/mono  E32 F_XET enum=OffOn
    /bus/15/mix/mlevel  F32 F_XET
/bus/15/mix/01  <CHMO> n=0
    /bus/15/mix/01/on  E32 F_XET enum=OffOn
    /bus/15/mix/01/level  F32 F_XET
    /bus/15/mix/01/pan  F32 F_XET
    /bus/15/mix/01/type  E32 F_XET enum=OffOn
    /bus/15/mix/01/panFollow  E32 F_XET
/bus/15/mix/02  <CHME> n=0
    /bus/15/mix/02/on  E32 F_XET enum=OffOn
    /bus/15/mix/02/level  F32 F_XET
/bus/15/mix/03  <CHMO> n=0
    /bus/15/mix/03/on  E32 F_XET enum=OffOn
    /bus/15/mix/03/level  F32 F_XET
    /bus/15/mix/03/pan  F32 F_XET
    /bus/15/mix/03/type  E32 F_XET enum=OffOn
    /bus/15/mix/03/panFollow  E32 F_XET
/bus/15/mix/04  <CHME> n=0
    /bus/15/mix/04/on  E32 F_XET enum=OffOn
    /bus/15/mix/04/level  F32 F_XET
/bus/15/mix/05  <CHMO> n=0
    /bus/15/mix/05/on  E32 F_XET enum=OffOn
    /bus/15/mix/05/level  F32 F_XET
    /bus/15/mix/05/pan  F32 F_XET
    /bus/15/mix/05/type  E32 F_XET enum=OffOn
    /bus/15/mix/05/panFollow  E32 F_XET
/bus/15/mix/06  <CHME> n=0
    /bus/15/mix/06/on  E32 F_XET enum=OffOn
    /bus/15/mix/06/level  F32 F_XET
/bus/15/grp  <CHGRP> n=0
    /bus/15/grp/dca  P32 F_XET
    /bus/15/grp/mute  P32 F_XET
```

### Xbus16 (X32Bus.h, 99 entries)

```
/bus  <BSCO> n=0
/bus/16  <BSCO> n=0
/bus/16/config  <BSCO> n=0
    /bus/16/config/name  S32 F_XET
    /bus/16/config/icon  I32 F_XET
    /bus/16/config/color  E32 F_XET enum=Xcolors
/bus/16/dyn  <CHDY> n=0
    /bus/16/dyn/on  E32 F_XET enum=OffOn
    /bus/16/dyn/mode  E32 F_XET enum=Xdymode
    /bus/16/dyn/det  E32 F_XET enum=Xdydet
    /bus/16/dyn/env  E32 F_XET enum=Xdyenv
    /bus/16/dyn/thr  F32 F_XET
    /bus/16/dyn/ratio  E32 F_XET enum=OffOn
    /bus/16/dyn/knee  F32 F_XET
    /bus/16/dyn/mgain  F32 F_XET
    /bus/16/dyn/attack  F32 F_XET
    /bus/16/dyn/hold  F32 F_XET
    /bus/16/dyn/release  F32 F_XET
    /bus/16/dyn/pos  E32 F_XET enum=Xdyppos
    /bus/16/dyn/keysrc  I32 F_XET
    /bus/16/dyn/mix  F32 F_XET enum=OffOn
    /bus/16/dyn/auto  E32 F_XET enum=OffOn
/bus/16/dyn/filter  <CHDF> n=0
    /bus/16/dyn/filter/on  E32 F_XET enum=OffOn
    /bus/16/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /bus/16/dyn/filter/f  F32 F_XET
/bus/16/insert  <CHIN> n=0
    /bus/16/insert/on  E32 F_XET enum=OffOn
    /bus/16/insert/pos  E32 F_XET enum=Xdyppos
    /bus/16/insert/sel  E32 F_XET enum=Xisel
/bus/16/eq  <OFFON> n=1
    /bus/16/eq/on  E32 F_XET enum=OffOn
/bus/16/eq/1  <CHEQ> n=0
    /bus/16/eq/1/type  E32 F_XET enum=Xeqty1
    /bus/16/eq/1/f  F32 F_XET
    /bus/16/eq/1/g  F32 F_XET
    /bus/16/eq/1/q  F32 F_XET
/bus/16/eq/2  <CHEQ> n=0
    /bus/16/eq/2/type  E32 F_XET enum=Xeqty1
    /bus/16/eq/2/f  F32 F_XET
    /bus/16/eq/2/g  F32 F_XET
    /bus/16/eq/2/q  F32 F_XET
/bus/16/eq/3  <CHEQ> n=0
    /bus/16/eq/3/type  E32 F_XET enum=Xeqty1
    /bus/16/eq/3/f  F32 F_XET
    /bus/16/eq/3/g  F32 F_XET
    /bus/16/eq/3/q  F32 F_XET
/bus/16/eq/4  <CHEQ> n=0
    /bus/16/eq/4/type  E32 F_XET enum=Xeqty1
    /bus/16/eq/4/f  F32 F_XET
    /bus/16/eq/4/g  F32 F_XET
    /bus/16/eq/4/q  F32 F_XET
/bus/16/eq/5  <CHEQ> n=0
    /bus/16/eq/5/type  E32 F_XET enum=Xeqty1
    /bus/16/eq/5/f  F32 F_XET
    /bus/16/eq/5/g  F32 F_XET
    /bus/16/eq/5/q  F32 F_XET
/bus/16/eq/6  <CHEQ> n=0
    /bus/16/eq/6/type  E32 F_XET enum=Xeqty1
    /bus/16/eq/6/f  F32 F_XET
    /bus/16/eq/6/g  F32 F_XET
    /bus/16/eq/6/q  F32 F_XET
/bus/16/mix  <CHMX> n=0
    /bus/16/mix/on  E32 F_XET enum=OffOn
    /bus/16/mix/fader  F32 F_XET
    /bus/16/mix/st  E32 F_XET enum=OffOn
    /bus/16/mix/pan  F32 F_XET
    /bus/16/mix/mono  E32 F_XET enum=OffOn
    /bus/16/mix/mlevel  F32 F_XET
/bus/16/mix/01  <CHMO> n=0
    /bus/16/mix/01/on  E32 F_XET enum=OffOn
    /bus/16/mix/01/level  F32 F_XET
    /bus/16/mix/01/pan  F32 F_XET
    /bus/16/mix/01/type  E32 F_XET enum=OffOn
    /bus/16/mix/01/panFollow  E32 F_XET
/bus/16/mix/02  <CHME> n=0
    /bus/16/mix/02/on  E32 F_XET enum=OffOn
    /bus/16/mix/02/level  F32 F_XET
/bus/16/mix/03  <CHMO> n=0
    /bus/16/mix/03/on  E32 F_XET enum=OffOn
    /bus/16/mix/03/level  F32 F_XET
    /bus/16/mix/03/pan  F32 F_XET
    /bus/16/mix/03/type  E32 F_XET enum=OffOn
    /bus/16/mix/03/panFollow  E32 F_XET
/bus/16/mix/04  <CHME> n=0
    /bus/16/mix/04/on  E32 F_XET enum=OffOn
    /bus/16/mix/04/level  F32 F_XET
/bus/16/mix/05  <CHMO> n=0
    /bus/16/mix/05/on  E32 F_XET enum=OffOn
    /bus/16/mix/05/level  F32 F_XET
    /bus/16/mix/05/pan  F32 F_XET
    /bus/16/mix/05/type  E32 F_XET enum=OffOn
    /bus/16/mix/05/panFollow  E32 F_XET
/bus/16/mix/06  <CHME> n=0
    /bus/16/mix/06/on  E32 F_XET enum=OffOn
    /bus/16/mix/06/level  F32 F_XET
/bus/16/grp  <CHGRP> n=0
    /bus/16/grp/dca  P32 F_XET
    /bus/16/grp/mute  P32 F_XET
```

### Xmtx01 (X32Mtx.h, 69 entries)

```
/mtx  <BSCO> n=0
/mtx/01  <BSCO> n=0
/mtx/01/config  <BSCO> n=0
    /mtx/01/config/name  S32 F_XET
    /mtx/01/config/icon  I32 F_XET
    /mtx/01/config/color  E32 F_XET enum=Xcolors
/mtx/01/preamp  <MXPR> n=0
    /mtx/01/preamp/invert  E32 F_XET enum=OffOn
/mtx/01/dyn  <MXDY> n=0
    /mtx/01/dyn/on  E32 F_XET enum=OffOn
    /mtx/01/dyn/mode  E32 F_XET enum=Xdymode
    /mtx/01/dyn/det  E32 F_XET enum=Xdydet
    /mtx/01/dyn/env  E32 F_XET enum=Xdyenv
    /mtx/01/dyn/thr  F32 F_XET
    /mtx/01/dyn/ratio  E32 F_XET enum=Xdyrat
    /mtx/01/dyn/knee  F32 F_XET
    /mtx/01/dyn/mgain  F32 F_XET
    /mtx/01/dyn/attack  F32 F_XET
    /mtx/01/dyn/hold  F32 F_XET
    /mtx/01/dyn/release  F32 F_XET
    /mtx/01/dyn/pos  E32 F_XET enum=Xdyppos
    /mtx/01/dyn/mix  F32 F_XET
    /mtx/01/dyn/auto  E32 F_XET enum=OffOn
/mtx/01/dyn/filter  <CHDF> n=0
    /mtx/01/dyn/filter/on  E32 F_XET enum=OffOn
    /mtx/01/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /mtx/01/dyn/filter/f  F32 F_XET
/mtx/01/insert  <CHIN> n=0
    /mtx/01/insert/on  E32 F_XET enum=OffOn
    /mtx/01/insert/pos  E32 F_XET enum=Xdyppos
    /mtx/01/insert/sel  E32 F_XET enum=Xisel
/mtx/01/eq  <OFFON> n=1
    /mtx/01/eq/on  E32 F_XET enum=OffOn
/mtx/01/eq/1  <CHEQ> n=0
    /mtx/01/eq/1/type  E32 F_XET enum=Xeqty2
    /mtx/01/eq/1/f  F32 F_XET
    /mtx/01/eq/1/g  F32 F_XET
    /mtx/01/eq/1/q  F32 F_XET
/mtx/01/eq/2  <CHEQ> n=0
    /mtx/01/eq/2/type  E32 F_XET enum=Xeqty2
    /mtx/01/eq/2/f  F32 F_XET
    /mtx/01/eq/2/g  F32 F_XET
    /mtx/01/eq/2/q  F32 F_XET
/mtx/01/eq/3  <CHEQ> n=0
    /mtx/01/eq/3/type  E32 F_XET enum=Xeqty2
    /mtx/01/eq/3/f  F32 F_XET
    /mtx/01/eq/3/g  F32 F_XET
    /mtx/01/eq/3/q  F32 F_XET
/mtx/01/eq/4  <CHEQ> n=0
    /mtx/01/eq/4/type  E32 F_XET enum=Xeqty2
    /mtx/01/eq/4/f  F32 F_XET
    /mtx/01/eq/4/g  F32 F_XET
    /mtx/01/eq/4/q  F32 F_XET
/mtx/01/eq/5  <CHEQ> n=0
    /mtx/01/eq/5/type  E32 F_XET enum=Xeqty2
    /mtx/01/eq/5/f  F32 F_XET
    /mtx/01/eq/5/g  F32 F_XET
    /mtx/01/eq/5/q  F32 F_XET
/mtx/01/eq/6  <CHEQ> n=0
    /mtx/01/eq/6/type  E32 F_XET enum=Xeqty2
    /mtx/01/eq/6/f  F32 F_XET
    /mtx/01/eq/6/g  F32 F_XET
    /mtx/01/eq/6/q  F32 F_XET
/mtx/01/mix  <CHME> n=0
    /mtx/01/mix/on  E32 F_XET enum=OffOn
    /mtx/01/mix/fader  F32 F_XET
/mtx/01/grp  <CHGRP> n=0
    /mtx/01/grp/dca  P32 F_XET
    /mtx/01/grp/mute  P32 F_XET
```

### Xmtx02 (X32Mtx.h, 69 entries)

```
/mtx  <BSCO> n=0
/mtx/02  <BSCO> n=0
/mtx/02/config  <BSCO> n=0
    /mtx/02/config/name  S32 F_XET
    /mtx/02/config/icon  I32 F_XET
    /mtx/02/config/color  E32 F_XET enum=Xcolors
/mtx/02/preamp  <MXPR> n=0
    /mtx/02/preamp/invert  E32 F_XET enum=OffOn
/mtx/02/dyn  <MXDY> n=0
    /mtx/02/dyn/on  E32 F_XET enum=OffOn
    /mtx/02/dyn/mode  E32 F_XET enum=Xdymode
    /mtx/02/dyn/det  E32 F_XET enum=Xdydet
    /mtx/02/dyn/env  E32 F_XET enum=Xdyenv
    /mtx/02/dyn/thr  F32 F_XET
    /mtx/02/dyn/ratio  E32 F_XET enum=Xdyrat
    /mtx/02/dyn/knee  F32 F_XET
    /mtx/02/dyn/mgain  F32 F_XET
    /mtx/02/dyn/attack  F32 F_XET
    /mtx/02/dyn/hold  F32 F_XET
    /mtx/02/dyn/release  F32 F_XET
    /mtx/02/dyn/pos  E32 F_XET enum=Xdyppos
    /mtx/02/dyn/mix  F32 F_XET
    /mtx/02/dyn/auto  E32 F_XET enum=OffOn
/mtx/02/dyn/filter  <CHDF> n=0
    /mtx/02/dyn/filter/on  E32 F_XET enum=OffOn
    /mtx/02/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /mtx/02/dyn/filter/f  F32 F_XET
/mtx/02/insert  <CHIN> n=0
    /mtx/02/insert/on  E32 F_XET enum=OffOn
    /mtx/02/insert/pos  E32 F_XET enum=Xdyppos
    /mtx/02/insert/sel  E32 F_XET enum=Xisel
/mtx/02/eq  <OFFON> n=1
    /mtx/02/eq/on  E32 F_XET enum=OffOn
/mtx/02/eq/1  <CHEQ> n=0
    /mtx/02/eq/1/type  E32 F_XET enum=Xeqty2
    /mtx/02/eq/1/f  F32 F_XET
    /mtx/02/eq/1/g  F32 F_XET
    /mtx/02/eq/1/q  F32 F_XET
/mtx/02/eq/2  <CHEQ> n=0
    /mtx/02/eq/2/type  E32 F_XET enum=Xeqty2
    /mtx/02/eq/2/f  F32 F_XET
    /mtx/02/eq/2/g  F32 F_XET
    /mtx/02/eq/2/q  F32 F_XET
/mtx/02/eq/3  <CHEQ> n=0
    /mtx/02/eq/3/type  E32 F_XET enum=Xeqty2
    /mtx/02/eq/3/f  F32 F_XET
    /mtx/02/eq/3/g  F32 F_XET
    /mtx/02/eq/3/q  F32 F_XET
/mtx/02/eq/4  <CHEQ> n=0
    /mtx/02/eq/4/type  E32 F_XET enum=Xeqty2
    /mtx/02/eq/4/f  F32 F_XET
    /mtx/02/eq/4/g  F32 F_XET
    /mtx/02/eq/4/q  F32 F_XET
/mtx/02/eq/5  <CHEQ> n=0
    /mtx/02/eq/5/type  E32 F_XET enum=Xeqty2
    /mtx/02/eq/5/f  F32 F_XET
    /mtx/02/eq/5/g  F32 F_XET
    /mtx/02/eq/5/q  F32 F_XET
/mtx/02/eq/6  <CHEQ> n=0
    /mtx/02/eq/6/type  E32 F_XET enum=Xeqty2
    /mtx/02/eq/6/f  F32 F_XET
    /mtx/02/eq/6/g  F32 F_XET
    /mtx/02/eq/6/q  F32 F_XET
/mtx/02/mix  <CHME> n=0
    /mtx/02/mix/on  E32 F_XET enum=OffOn
    /mtx/02/mix/fader  F32 F_XET
/mtx/02/grp  <CHGRP> n=0
    /mtx/02/grp/dca  P32 F_XET
    /mtx/02/grp/mute  P32 F_XET
```

### Xmtx03 (X32Mtx.h, 69 entries)

```
/mtx  <BSCO> n=0
/mtx/03  <BSCO> n=0
/mtx/03/config  <BSCO> n=0
    /mtx/03/config/name  S32 F_XET
    /mtx/03/config/icon  I32 F_XET
    /mtx/03/config/color  E32 F_XET enum=Xcolors
/mtx/03/preamp  <MXPR> n=0
    /mtx/03/preamp/invert  E32 F_XET enum=OffOn
/mtx/03/dyn  <MXDY> n=0
    /mtx/03/dyn/on  E32 F_XET enum=OffOn
    /mtx/03/dyn/mode  E32 F_XET enum=Xdymode
    /mtx/03/dyn/det  E32 F_XET enum=Xdydet
    /mtx/03/dyn/env  E32 F_XET enum=Xdyenv
    /mtx/03/dyn/thr  F32 F_XET
    /mtx/03/dyn/ratio  E32 F_XET enum=Xdyrat
    /mtx/03/dyn/knee  F32 F_XET
    /mtx/03/dyn/mgain  F32 F_XET
    /mtx/03/dyn/attack  F32 F_XET
    /mtx/03/dyn/hold  F32 F_XET
    /mtx/03/dyn/release  F32 F_XET
    /mtx/03/dyn/pos  E32 F_XET enum=Xdyppos
    /mtx/03/dyn/mix  F32 F_XET
    /mtx/03/dyn/auto  E32 F_XET enum=OffOn
/mtx/03/dyn/filter  <CHDF> n=0
    /mtx/03/dyn/filter/on  E32 F_XET enum=OffOn
    /mtx/03/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /mtx/03/dyn/filter/f  F32 F_XET
/mtx/03/insert  <CHIN> n=0
    /mtx/03/insert/on  E32 F_XET enum=OffOn
    /mtx/03/insert/pos  E32 F_XET enum=Xdyppos
    /mtx/03/insert/sel  E32 F_XET enum=Xisel
/mtx/03/eq  <OFFON> n=1
    /mtx/03/eq/on  E32 F_XET enum=OffOn
/mtx/03/eq/1  <CHEQ> n=0
    /mtx/03/eq/1/type  E32 F_XET enum=Xeqty2
    /mtx/03/eq/1/f  F32 F_XET
    /mtx/03/eq/1/g  F32 F_XET
    /mtx/03/eq/1/q  F32 F_XET
/mtx/03/eq/2  <CHEQ> n=0
    /mtx/03/eq/2/type  E32 F_XET enum=Xeqty2
    /mtx/03/eq/2/f  F32 F_XET
    /mtx/03/eq/2/g  F32 F_XET
    /mtx/03/eq/2/q  F32 F_XET
/mtx/03/eq/3  <CHEQ> n=0
    /mtx/03/eq/3/type  E32 F_XET enum=Xeqty2
    /mtx/03/eq/3/f  F32 F_XET
    /mtx/03/eq/3/g  F32 F_XET
    /mtx/03/eq/3/q  F32 F_XET
/mtx/03/eq/4  <CHEQ> n=0
    /mtx/03/eq/4/type  E32 F_XET enum=Xeqty2
    /mtx/03/eq/4/f  F32 F_XET
    /mtx/03/eq/4/g  F32 F_XET
    /mtx/03/eq/4/q  F32 F_XET
/mtx/03/eq/5  <CHEQ> n=0
    /mtx/03/eq/5/type  E32 F_XET enum=Xeqty2
    /mtx/03/eq/5/f  F32 F_XET
    /mtx/03/eq/5/g  F32 F_XET
    /mtx/03/eq/5/q  F32 F_XET
/mtx/03/eq/6  <CHEQ> n=0
    /mtx/03/eq/6/type  E32 F_XET enum=Xeqty2
    /mtx/03/eq/6/f  F32 F_XET
    /mtx/03/eq/6/g  F32 F_XET
    /mtx/03/eq/6/q  F32 F_XET
/mtx/03/mix  <CHME> n=0
    /mtx/03/mix/on  E32 F_XET enum=OffOn
    /mtx/03/mix/fader  F32 F_XET
/mtx/03/grp  <CHGRP> n=0
    /mtx/03/grp/dca  P32 F_XET
    /mtx/03/grp/mute  P32 F_XET
```

### Xmtx04 (X32Mtx.h, 69 entries)

```
/mtx  <BSCO> n=0
/mtx/04  <BSCO> n=0
/mtx/04/config  <BSCO> n=0
    /mtx/04/config/name  S32 F_XET
    /mtx/04/config/icon  I32 F_XET
    /mtx/04/config/color  E32 F_XET enum=Xcolors
/mtx/04/preamp  <MXPR> n=0
    /mtx/04/preamp/invert  E32 F_XET enum=OffOn
/mtx/04/dyn  <MXDY> n=0
    /mtx/04/dyn/on  E32 F_XET enum=OffOn
    /mtx/04/dyn/mode  E32 F_XET enum=Xdymode
    /mtx/04/dyn/det  E32 F_XET enum=Xdydet
    /mtx/04/dyn/env  E32 F_XET enum=Xdyenv
    /mtx/04/dyn/thr  F32 F_XET
    /mtx/04/dyn/ratio  E32 F_XET enum=Xdyrat
    /mtx/04/dyn/knee  F32 F_XET
    /mtx/04/dyn/mgain  F32 F_XET
    /mtx/04/dyn/attack  F32 F_XET
    /mtx/04/dyn/hold  F32 F_XET
    /mtx/04/dyn/release  F32 F_XET
    /mtx/04/dyn/pos  E32 F_XET enum=Xdyppos
    /mtx/04/dyn/mix  F32 F_XET
    /mtx/04/dyn/auto  E32 F_XET enum=OffOn
/mtx/04/dyn/filter  <CHDF> n=0
    /mtx/04/dyn/filter/on  E32 F_XET enum=OffOn
    /mtx/04/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /mtx/04/dyn/filter/f  F32 F_XET
/mtx/04/insert  <CHIN> n=0
    /mtx/04/insert/on  E32 F_XET enum=OffOn
    /mtx/04/insert/pos  E32 F_XET enum=Xdyppos
    /mtx/04/insert/sel  E32 F_XET enum=Xisel
/mtx/04/eq  <OFFON> n=1
    /mtx/04/eq/on  E32 F_XET enum=OffOn
/mtx/04/eq/1  <CHEQ> n=0
    /mtx/04/eq/1/type  E32 F_XET enum=Xeqty2
    /mtx/04/eq/1/f  F32 F_XET
    /mtx/04/eq/1/g  F32 F_XET
    /mtx/04/eq/1/q  F32 F_XET
/mtx/04/eq/2  <CHEQ> n=0
    /mtx/04/eq/2/type  E32 F_XET enum=Xeqty2
    /mtx/04/eq/2/f  F32 F_XET
    /mtx/04/eq/2/g  F32 F_XET
    /mtx/04/eq/2/q  F32 F_XET
/mtx/04/eq/3  <CHEQ> n=0
    /mtx/04/eq/3/type  E32 F_XET enum=Xeqty2
    /mtx/04/eq/3/f  F32 F_XET
    /mtx/04/eq/3/g  F32 F_XET
    /mtx/04/eq/3/q  F32 F_XET
/mtx/04/eq/4  <CHEQ> n=0
    /mtx/04/eq/4/type  E32 F_XET enum=Xeqty2
    /mtx/04/eq/4/f  F32 F_XET
    /mtx/04/eq/4/g  F32 F_XET
    /mtx/04/eq/4/q  F32 F_XET
/mtx/04/eq/5  <CHEQ> n=0
    /mtx/04/eq/5/type  E32 F_XET enum=Xeqty2
    /mtx/04/eq/5/f  F32 F_XET
    /mtx/04/eq/5/g  F32 F_XET
    /mtx/04/eq/5/q  F32 F_XET
/mtx/04/eq/6  <CHEQ> n=0
    /mtx/04/eq/6/type  E32 F_XET enum=Xeqty2
    /mtx/04/eq/6/f  F32 F_XET
    /mtx/04/eq/6/g  F32 F_XET
    /mtx/04/eq/6/q  F32 F_XET
/mtx/04/mix  <CHME> n=0
    /mtx/04/mix/on  E32 F_XET enum=OffOn
    /mtx/04/mix/fader  F32 F_XET
/mtx/04/grp  <CHGRP> n=0
    /mtx/04/grp/dca  P32 F_XET
    /mtx/04/grp/mute  P32 F_XET
```

### Xmtx05 (X32Mtx.h, 69 entries)

```
/mtx  <BSCO> n=0
/mtx/05  <BSCO> n=0
/mtx/05/config  <BSCO> n=0
    /mtx/05/config/name  S32 F_XET
    /mtx/05/config/icon  I32 F_XET
    /mtx/05/config/color  E32 F_XET enum=Xcolors
/mtx/05/preamp  <MXPR> n=0
    /mtx/05/preamp/invert  E32 F_XET enum=OffOn
/mtx/05/dyn  <MXDY> n=0
    /mtx/05/dyn/on  E32 F_XET enum=OffOn
    /mtx/05/dyn/mode  E32 F_XET enum=Xdymode
    /mtx/05/dyn/det  E32 F_XET enum=Xdydet
    /mtx/05/dyn/env  E32 F_XET enum=Xdyenv
    /mtx/05/dyn/thr  F32 F_XET
    /mtx/05/dyn/ratio  E32 F_XET enum=Xdyrat
    /mtx/05/dyn/knee  F32 F_XET
    /mtx/05/dyn/mgain  F32 F_XET
    /mtx/05/dyn/attack  F32 F_XET
    /mtx/05/dyn/hold  F32 F_XET
    /mtx/05/dyn/release  F32 F_XET
    /mtx/05/dyn/pos  E32 F_XET enum=Xdyppos
    /mtx/05/dyn/mix  F32 F_XET
    /mtx/05/dyn/auto  E32 F_XET enum=OffOn
/mtx/05/dyn/filter  <CHDF> n=0
    /mtx/05/dyn/filter/on  E32 F_XET enum=OffOn
    /mtx/05/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /mtx/05/dyn/filter/f  F32 F_XET
/mtx/05/insert  <CHIN> n=0
    /mtx/05/insert/on  E32 F_XET enum=OffOn
    /mtx/05/insert/pos  E32 F_XET enum=Xdyppos
    /mtx/05/insert/sel  E32 F_XET enum=Xisel
/mtx/05/eq  <OFFON> n=1
    /mtx/05/eq/on  E32 F_XET enum=OffOn
/mtx/05/eq/1  <CHEQ> n=0
    /mtx/05/eq/1/type  E32 F_XET enum=Xeqty2
    /mtx/05/eq/1/f  F32 F_XET
    /mtx/05/eq/1/g  F32 F_XET
    /mtx/05/eq/1/q  F32 F_XET
/mtx/05/eq/2  <CHEQ> n=0
    /mtx/05/eq/2/type  E32 F_XET enum=Xeqty2
    /mtx/05/eq/2/f  F32 F_XET
    /mtx/05/eq/2/g  F32 F_XET
    /mtx/05/eq/2/q  F32 F_XET
/mtx/05/eq/3  <CHEQ> n=0
    /mtx/05/eq/3/type  E32 F_XET enum=Xeqty2
    /mtx/05/eq/3/f  F32 F_XET
    /mtx/05/eq/3/g  F32 F_XET
    /mtx/05/eq/3/q  F32 F_XET
/mtx/05/eq/4  <CHEQ> n=0
    /mtx/05/eq/4/type  E32 F_XET enum=Xeqty2
    /mtx/05/eq/4/f  F32 F_XET
    /mtx/05/eq/4/g  F32 F_XET
    /mtx/05/eq/4/q  F32 F_XET
/mtx/05/eq/5  <CHEQ> n=0
    /mtx/05/eq/5/type  E32 F_XET enum=Xeqty2
    /mtx/05/eq/5/f  F32 F_XET
    /mtx/05/eq/5/g  F32 F_XET
    /mtx/05/eq/5/q  F32 F_XET
/mtx/05/eq/6  <CHEQ> n=0
    /mtx/05/eq/6/type  E32 F_XET enum=Xeqty2
    /mtx/05/eq/6/f  F32 F_XET
    /mtx/05/eq/6/g  F32 F_XET
    /mtx/05/eq/6/q  F32 F_XET
/mtx/05/mix  <CHME> n=0
    /mtx/05/mix/on  E32 F_XET enum=OffOn
    /mtx/05/mix/fader  F32 F_XET
/mtx/05/grp  <CHGRP> n=0
    /mtx/05/grp/dca  P32 F_XET
    /mtx/05/grp/mute  P32 F_XET
```

### Xmtx06 (X32Mtx.h, 69 entries)

```
/mtx  <BSCO> n=0
/mtx/06  <BSCO> n=0
/mtx/06/config  <BSCO> n=0
    /mtx/06/config/name  S32 F_XET
    /mtx/06/config/icon  I32 F_XET
    /mtx/06/config/color  E32 F_XET enum=Xcolors
/mtx/06/preamp  <MXPR> n=0
    /mtx/06/preamp/invert  E32 F_XET enum=OffOn
/mtx/06/dyn  <MXDY> n=0
    /mtx/06/dyn/on  E32 F_XET enum=OffOn
    /mtx/06/dyn/mode  E32 F_XET enum=Xdymode
    /mtx/06/dyn/det  E32 F_XET enum=Xdydet
    /mtx/06/dyn/env  E32 F_XET enum=Xdyenv
    /mtx/06/dyn/thr  F32 F_XET
    /mtx/06/dyn/ratio  E32 F_XET enum=Xdyrat
    /mtx/06/dyn/knee  F32 F_XET
    /mtx/06/dyn/mgain  F32 F_XET
    /mtx/06/dyn/attack  F32 F_XET
    /mtx/06/dyn/hold  F32 F_XET
    /mtx/06/dyn/release  F32 F_XET
    /mtx/06/dyn/pos  E32 F_XET enum=Xdyppos
    /mtx/06/dyn/mix  F32 F_XET
    /mtx/06/dyn/auto  E32 F_XET enum=OffOn
/mtx/06/dyn/filter  <CHDF> n=0
    /mtx/06/dyn/filter/on  E32 F_XET enum=OffOn
    /mtx/06/dyn/filter/type  E32 F_XET enum=Xdyftyp
    /mtx/06/dyn/filter/f  F32 F_XET
/mtx/06/insert  <CHIN> n=0
    /mtx/06/insert/on  E32 F_XET enum=OffOn
    /mtx/06/insert/pos  E32 F_XET enum=Xdyppos
    /mtx/06/insert/sel  E32 F_XET enum=Xisel
/mtx/06/eq  <OFFON> n=1
    /mtx/06/eq/on  E32 F_XET enum=OffOn
/mtx/06/eq/1  <CHEQ> n=0
    /mtx/06/eq/1/type  E32 F_XET enum=Xeqty2
    /mtx/06/eq/1/f  F32 F_XET
    /mtx/06/eq/1/g  F32 F_XET
    /mtx/06/eq/1/q  F32 F_XET
/mtx/06/eq/2  <CHEQ> n=0
    /mtx/06/eq/2/type  E32 F_XET enum=Xeqty2
    /mtx/06/eq/2/f  F32 F_XET
    /mtx/06/eq/2/g  F32 F_XET
    /mtx/06/eq/2/q  F32 F_XET
/mtx/06/eq/3  <CHEQ> n=0
    /mtx/06/eq/3/type  E32 F_XET enum=Xeqty2
    /mtx/06/eq/3/f  F32 F_XET
    /mtx/06/eq/3/g  F32 F_XET
    /mtx/06/eq/3/q  F32 F_XET
/mtx/06/eq/4  <CHEQ> n=0
    /mtx/06/eq/4/type  E32 F_XET enum=Xeqty2
    /mtx/06/eq/4/f  F32 F_XET
    /mtx/06/eq/4/g  F32 F_XET
    /mtx/06/eq/4/q  F32 F_XET
/mtx/06/eq/5  <CHEQ> n=0
    /mtx/06/eq/5/type  E32 F_XET enum=Xeqty2
    /mtx/06/eq/5/f  F32 F_XET
    /mtx/06/eq/5/g  F32 F_XET
    /mtx/06/eq/5/q  F32 F_XET
/mtx/06/eq/6  <CHEQ> n=0
    /mtx/06/eq/6/type  E32 F_XET enum=Xeqty2
    /mtx/06/eq/6/f  F32 F_XET
    /mtx/06/eq/6/g  F32 F_XET
    /mtx/06/eq/6/q  F32 F_XET
/mtx/06/mix  <CHME> n=0
    /mtx/06/mix/on  E32 F_XET enum=OffOn
    /mtx/06/mix/fader  F32 F_XET
/mtx/06/grp  <CHGRP> n=0
    /mtx/06/grp/dca  P32 F_XET
    /mtx/06/grp/mute  P32 F_XET
```

### Xdca (X32Dca.h, 57 entries)

```
/dca  <CHME> n=0
/dca/1  <CHME> n=0
    /dca/1/on  E32 F_XET enum=OffOn
    /dca/1/fader  F32 F_XET
/dca/1/config  <BSCO> n=0
    /dca/1/config/name  S32 F_XET
    /dca/1/config/icon  I32 F_XET
    /dca/1/config/color  E32 F_XET enum=Xcolors
/dca/2  <CHME> n=0
    /dca/2/on  E32 F_XET enum=OffOn
    /dca/2/fader  F32 F_XET
/dca/2/config  <BSCO> n=0
    /dca/2/config/name  S32 F_XET
    /dca/2/config/icon  I32 F_XET
    /dca/2/config/color  E32 F_XET enum=Xcolors
/dca/3  <CHME> n=0
    /dca/3/on  E32 F_XET enum=OffOn
    /dca/3/fader  F32 F_XET
/dca/3/config  <BSCO> n=0
    /dca/3/config/name  S32 F_XET
    /dca/3/config/icon  I32 F_XET
    /dca/3/config/color  E32 F_XET enum=Xcolors
/dca/4  <CHME> n=0
    /dca/4/on  E32 F_XET enum=OffOn
    /dca/4/fader  F32 F_XET
/dca/4/config  <BSCO> n=0
    /dca/4/config/name  S32 F_XET
    /dca/4/config/icon  I32 F_XET
    /dca/4/config/color  E32 F_XET enum=Xcolors
/dca/5  <CHME> n=0
    /dca/5/on  E32 F_XET enum=OffOn
    /dca/5/fader  F32 F_XET
/dca/5/config  <BSCO> n=0
    /dca/5/config/name  S32 F_XET
    /dca/5/config/icon  I32 F_XET
    /dca/5/config/color  E32 F_XET enum=Xcolors
/dca/6  <CHME> n=0
    /dca/6/on  E32 F_XET enum=OffOn
    /dca/6/fader  F32 F_XET
/dca/6/config  <BSCO> n=0
    /dca/6/config/name  S32 F_XET
    /dca/6/config/icon  I32 F_XET
    /dca/6/config/color  E32 F_XET enum=Xcolors
/dca/7  <CHME> n=0
    /dca/7/on  E32 F_XET enum=OffOn
    /dca/7/fader  F32 F_XET
/dca/7/config  <BSCO> n=0
    /dca/7/config/name  S32 F_XET
    /dca/7/config/icon  I32 F_XET
    /dca/7/config/color  E32 F_XET enum=Xcolors
/dca/8  <CHME> n=0
    /dca/8/on  E32 F_XET enum=OffOn
    /dca/8/fader  F32 F_XET
/dca/8/config  <BSCO> n=0
    /dca/8/config/name  S32 F_XET
    /dca/8/config/icon  I32 F_XET
    /dca/8/config/color  E32 F_XET enum=Xcolors
```

### Xfx1 (X32Fx.h, 71 entries)

```
/fx  <FXTYP1> n=0
/fx/1  <FXTYP1> n=0
    /fx/1/type  E32 F_XET enum=Sfxtyp1
/fx/1/source  <FXSRC> n=0
    /fx/1/source/l  E32 F_XET enum=Sfxsrc
    /fx/1/source/r  E32 F_XET enum=Sfxsrc
/fx/1/par  <FXPAR1> n=0
    /fx/1/par/01  FX32 F_XET
    /fx/1/par/02  FX32 F_XET
    /fx/1/par/03  FX32 F_XET
    /fx/1/par/04  FX32 F_XET
    /fx/1/par/05  FX32 F_XET
    /fx/1/par/06  FX32 F_XET
    /fx/1/par/07  FX32 F_XET
    /fx/1/par/08  FX32 F_XET
    /fx/1/par/09  FX32 F_XET
    /fx/1/par/10  FX32 F_XET
    /fx/1/par/11  FX32 F_XET
    /fx/1/par/12  FX32 F_XET
    /fx/1/par/13  FX32 F_XET
    /fx/1/par/14  FX32 F_XET
    /fx/1/par/15  FX32 F_XET
    /fx/1/par/16  FX32 F_XET
    /fx/1/par/17  FX32 F_XET
    /fx/1/par/18  FX32 F_XET
    /fx/1/par/19  FX32 F_XET
    /fx/1/par/20  FX32 F_XET
    /fx/1/par/21  FX32 F_XET
    /fx/1/par/22  FX32 F_XET
    /fx/1/par/23  FX32 F_XET
    /fx/1/par/24  FX32 F_XET
    /fx/1/par/25  FX32 F_XET
    /fx/1/par/26  FX32 F_XET
    /fx/1/par/27  FX32 F_XET
    /fx/1/par/28  FX32 F_XET
    /fx/1/par/29  FX32 F_XET
    /fx/1/par/30  FX32 F_XET
    /fx/1/par/31  FX32 F_XET
    /fx/1/par/32  FX32 F_XET
    /fx/1/par/33  FX32 F_XET
    /fx/1/par/34  FX32 F_XET
    /fx/1/par/35  FX32 F_XET
    /fx/1/par/36  FX32 F_XET
    /fx/1/par/37  FX32 F_XET
    /fx/1/par/38  FX32 F_XET
    /fx/1/par/39  FX32 F_XET
    /fx/1/par/40  FX32 F_XET
    /fx/1/par/41  FX32 F_XET
    /fx/1/par/42  FX32 F_XET
    /fx/1/par/43  FX32 F_XET
    /fx/1/par/44  FX32 F_XET
    /fx/1/par/45  FX32 F_XET
    /fx/1/par/46  FX32 F_XET
    /fx/1/par/47  FX32 F_XET
    /fx/1/par/48  FX32 F_XET
    /fx/1/par/49  FX32 F_XET
    /fx/1/par/50  FX32 F_XET
    /fx/1/par/51  FX32 F_XET
    /fx/1/par/52  FX32 F_XET
    /fx/1/par/53  FX32 F_XET
    /fx/1/par/54  FX32 F_XET
    /fx/1/par/55  FX32 F_XET
    /fx/1/par/56  FX32 F_XET
    /fx/1/par/57  FX32 F_XET
    /fx/1/par/58  FX32 F_XET
    /fx/1/par/59  FX32 F_XET
    /fx/1/par/60  FX32 F_XET
    /fx/1/par/61  FX32 F_XET
    /fx/1/par/62  FX32 F_XET
    /fx/1/par/63  FX32 F_XET
    /fx/1/par/64  FX32 F_XET
```

### Xfx2 (X32Fx.h, 71 entries)

```
/fx  <FXTYP1> n=0
/fx/2  <FXTYP1> n=0
    /fx/2/type  E32 F_XET enum=Sfxtyp1
/fx/2/source  <FXSRC> n=0
    /fx/2/source/l  E32 F_XET enum=Sfxsrc
    /fx/2/source/r  E32 F_XET enum=Sfxsrc
/fx/2/par  <FXPAR1> n=0
    /fx/2/par/01  FX32 F_XET
    /fx/2/par/02  FX32 F_XET
    /fx/2/par/03  FX32 F_XET
    /fx/2/par/04  FX32 F_XET
    /fx/2/par/05  FX32 F_XET
    /fx/2/par/06  FX32 F_XET
    /fx/2/par/07  FX32 F_XET
    /fx/2/par/08  FX32 F_XET
    /fx/2/par/09  FX32 F_XET
    /fx/2/par/10  FX32 F_XET
    /fx/2/par/11  FX32 F_XET
    /fx/2/par/12  FX32 F_XET
    /fx/2/par/13  FX32 F_XET
    /fx/2/par/14  FX32 F_XET
    /fx/2/par/15  FX32 F_XET
    /fx/2/par/16  FX32 F_XET
    /fx/2/par/17  FX32 F_XET
    /fx/2/par/18  FX32 F_XET
    /fx/2/par/19  FX32 F_XET
    /fx/2/par/20  FX32 F_XET
    /fx/2/par/21  FX32 F_XET
    /fx/2/par/22  FX32 F_XET
    /fx/2/par/23  FX32 F_XET
    /fx/2/par/24  FX32 F_XET
    /fx/2/par/25  FX32 F_XET
    /fx/2/par/26  FX32 F_XET
    /fx/2/par/27  FX32 F_XET
    /fx/2/par/28  FX32 F_XET
    /fx/2/par/29  FX32 F_XET
    /fx/2/par/30  FX32 F_XET
    /fx/2/par/31  FX32 F_XET
    /fx/2/par/32  FX32 F_XET
    /fx/2/par/33  FX32 F_XET
    /fx/2/par/34  FX32 F_XET
    /fx/2/par/35  FX32 F_XET
    /fx/2/par/36  FX32 F_XET
    /fx/2/par/37  FX32 F_XET
    /fx/2/par/38  FX32 F_XET
    /fx/2/par/39  FX32 F_XET
    /fx/2/par/40  FX32 F_XET
    /fx/2/par/41  FX32 F_XET
    /fx/2/par/42  FX32 F_XET
    /fx/2/par/43  FX32 F_XET
    /fx/2/par/44  FX32 F_XET
    /fx/2/par/45  FX32 F_XET
    /fx/2/par/46  FX32 F_XET
    /fx/2/par/47  FX32 F_XET
    /fx/2/par/48  FX32 F_XET
    /fx/2/par/49  FX32 F_XET
    /fx/2/par/50  FX32 F_XET
    /fx/2/par/51  FX32 F_XET
    /fx/2/par/52  FX32 F_XET
    /fx/2/par/53  FX32 F_XET
    /fx/2/par/54  FX32 F_XET
    /fx/2/par/55  FX32 F_XET
    /fx/2/par/56  FX32 F_XET
    /fx/2/par/57  FX32 F_XET
    /fx/2/par/58  FX32 F_XET
    /fx/2/par/59  FX32 F_XET
    /fx/2/par/60  FX32 F_XET
    /fx/2/par/61  FX32 F_XET
    /fx/2/par/62  FX32 F_XET
    /fx/2/par/63  FX32 F_XET
    /fx/2/par/64  FX32 F_XET
```

### Xfx3 (X32Fx.h, 71 entries)

```
/fx  <FXTYP1> n=0
/fx/3  <FXTYP1> n=0
    /fx/3/type  E32 F_XET enum=Sfxtyp1
/fx/3/source  <FXSRC> n=0
    /fx/3/source/l  E32 F_XET enum=Sfxsrc
    /fx/3/source/r  E32 F_XET enum=Sfxsrc
/fx/3/par  <FXPAR1> n=0
    /fx/3/par/01  FX32 F_XET
    /fx/3/par/02  FX32 F_XET
    /fx/3/par/03  FX32 F_XET
    /fx/3/par/04  FX32 F_XET
    /fx/3/par/05  FX32 F_XET
    /fx/3/par/06  FX32 F_XET
    /fx/3/par/07  FX32 F_XET
    /fx/3/par/08  FX32 F_XET
    /fx/3/par/09  FX32 F_XET
    /fx/3/par/10  FX32 F_XET
    /fx/3/par/11  FX32 F_XET
    /fx/3/par/12  FX32 F_XET
    /fx/3/par/13  FX32 F_XET
    /fx/3/par/14  FX32 F_XET
    /fx/3/par/15  FX32 F_XET
    /fx/3/par/16  FX32 F_XET
    /fx/3/par/17  FX32 F_XET
    /fx/3/par/18  FX32 F_XET
    /fx/3/par/19  FX32 F_XET
    /fx/3/par/20  FX32 F_XET
    /fx/3/par/21  FX32 F_XET
    /fx/3/par/22  FX32 F_XET
    /fx/3/par/23  FX32 F_XET
    /fx/3/par/24  FX32 F_XET
    /fx/3/par/25  FX32 F_XET
    /fx/3/par/26  FX32 F_XET
    /fx/3/par/27  FX32 F_XET
    /fx/3/par/28  FX32 F_XET
    /fx/3/par/29  FX32 F_XET
    /fx/3/par/30  FX32 F_XET
    /fx/3/par/31  FX32 F_XET
    /fx/3/par/32  FX32 F_XET
    /fx/3/par/33  FX32 F_XET
    /fx/3/par/34  FX32 F_XET
    /fx/3/par/35  FX32 F_XET
    /fx/3/par/36  FX32 F_XET
    /fx/3/par/37  FX32 F_XET
    /fx/3/par/38  FX32 F_XET
    /fx/3/par/39  FX32 F_XET
    /fx/3/par/40  FX32 F_XET
    /fx/3/par/41  FX32 F_XET
    /fx/3/par/42  FX32 F_XET
    /fx/3/par/43  FX32 F_XET
    /fx/3/par/44  FX32 F_XET
    /fx/3/par/45  FX32 F_XET
    /fx/3/par/46  FX32 F_XET
    /fx/3/par/47  FX32 F_XET
    /fx/3/par/48  FX32 F_XET
    /fx/3/par/49  FX32 F_XET
    /fx/3/par/50  FX32 F_XET
    /fx/3/par/51  FX32 F_XET
    /fx/3/par/52  FX32 F_XET
    /fx/3/par/53  FX32 F_XET
    /fx/3/par/54  FX32 F_XET
    /fx/3/par/55  FX32 F_XET
    /fx/3/par/56  FX32 F_XET
    /fx/3/par/57  FX32 F_XET
    /fx/3/par/58  FX32 F_XET
    /fx/3/par/59  FX32 F_XET
    /fx/3/par/60  FX32 F_XET
    /fx/3/par/61  FX32 F_XET
    /fx/3/par/62  FX32 F_XET
    /fx/3/par/63  FX32 F_XET
    /fx/3/par/64  FX32 F_XET
```

### Xfx4 (X32Fx.h, 71 entries)

```
/fx  <FXTYP1> n=0
/fx/4  <FXTYP1> n=0
    /fx/4/type  E32 F_XET enum=Sfxtyp1
/fx/4/source  <FXSRC> n=0
    /fx/4/source/l  E32 F_XET enum=Sfxsrc
    /fx/4/source/r  E32 F_XET enum=Sfxsrc
/fx/4/par  <FXPAR1> n=0
    /fx/4/par/01  FX32 F_XET
    /fx/4/par/02  FX32 F_XET
    /fx/4/par/03  FX32 F_XET
    /fx/4/par/04  FX32 F_XET
    /fx/4/par/05  FX32 F_XET
    /fx/4/par/06  FX32 F_XET
    /fx/4/par/07  FX32 F_XET
    /fx/4/par/08  FX32 F_XET
    /fx/4/par/09  FX32 F_XET
    /fx/4/par/10  FX32 F_XET
    /fx/4/par/11  FX32 F_XET
    /fx/4/par/12  FX32 F_XET
    /fx/4/par/13  FX32 F_XET
    /fx/4/par/14  FX32 F_XET
    /fx/4/par/15  FX32 F_XET
    /fx/4/par/16  FX32 F_XET
    /fx/4/par/17  FX32 F_XET
    /fx/4/par/18  FX32 F_XET
    /fx/4/par/19  FX32 F_XET
    /fx/4/par/20  FX32 F_XET
    /fx/4/par/21  FX32 F_XET
    /fx/4/par/22  FX32 F_XET
    /fx/4/par/23  FX32 F_XET
    /fx/4/par/24  FX32 F_XET
    /fx/4/par/25  FX32 F_XET
    /fx/4/par/26  FX32 F_XET
    /fx/4/par/27  FX32 F_XET
    /fx/4/par/28  FX32 F_XET
    /fx/4/par/29  FX32 F_XET
    /fx/4/par/30  FX32 F_XET
    /fx/4/par/31  FX32 F_XET
    /fx/4/par/32  FX32 F_XET
    /fx/4/par/33  FX32 F_XET
    /fx/4/par/34  FX32 F_XET
    /fx/4/par/35  FX32 F_XET
    /fx/4/par/36  FX32 F_XET
    /fx/4/par/37  FX32 F_XET
    /fx/4/par/38  FX32 F_XET
    /fx/4/par/39  FX32 F_XET
    /fx/4/par/40  FX32 F_XET
    /fx/4/par/41  FX32 F_XET
    /fx/4/par/42  FX32 F_XET
    /fx/4/par/43  FX32 F_XET
    /fx/4/par/44  FX32 F_XET
    /fx/4/par/45  FX32 F_XET
    /fx/4/par/46  FX32 F_XET
    /fx/4/par/47  FX32 F_XET
    /fx/4/par/48  FX32 F_XET
    /fx/4/par/49  FX32 F_XET
    /fx/4/par/50  FX32 F_XET
    /fx/4/par/51  FX32 F_XET
    /fx/4/par/52  FX32 F_XET
    /fx/4/par/53  FX32 F_XET
    /fx/4/par/54  FX32 F_XET
    /fx/4/par/55  FX32 F_XET
    /fx/4/par/56  FX32 F_XET
    /fx/4/par/57  FX32 F_XET
    /fx/4/par/58  FX32 F_XET
    /fx/4/par/59  FX32 F_XET
    /fx/4/par/60  FX32 F_XET
    /fx/4/par/61  FX32 F_XET
    /fx/4/par/62  FX32 F_XET
    /fx/4/par/63  FX32 F_XET
    /fx/4/par/64  FX32 F_XET
```

### Xfx5 (X32Fx.h, 68 entries)

```
/fx  <FXTYP1> n=0
/fx/5  <FXTYP2> n=0
    /fx/5/type  E32 F_XET enum=Sfxtyp2
/fx/5/par  <FXPAR2> n=0
    /fx/5/par/01  FX32 F_XET
    /fx/5/par/02  FX32 F_XET
    /fx/5/par/03  FX32 F_XET
    /fx/5/par/04  FX32 F_XET
    /fx/5/par/05  FX32 F_XET
    /fx/5/par/06  FX32 F_XET
    /fx/5/par/07  FX32 F_XET
    /fx/5/par/08  FX32 F_XET
    /fx/5/par/09  FX32 F_XET
    /fx/5/par/10  FX32 F_XET
    /fx/5/par/11  FX32 F_XET
    /fx/5/par/12  FX32 F_XET
    /fx/5/par/13  FX32 F_XET
    /fx/5/par/14  FX32 F_XET
    /fx/5/par/15  FX32 F_XET
    /fx/5/par/16  FX32 F_XET
    /fx/5/par/17  FX32 F_XET
    /fx/5/par/18  FX32 F_XET
    /fx/5/par/19  FX32 F_XET
    /fx/5/par/20  FX32 F_XET
    /fx/5/par/21  FX32 F_XET
    /fx/5/par/22  FX32 F_XET
    /fx/5/par/23  FX32 F_XET
    /fx/5/par/24  FX32 F_XET
    /fx/5/par/25  FX32 F_XET
    /fx/5/par/26  FX32 F_XET
    /fx/5/par/27  FX32 F_XET
    /fx/5/par/28  FX32 F_XET
    /fx/5/par/29  FX32 F_XET
    /fx/5/par/30  FX32 F_XET
    /fx/5/par/31  FX32 F_XET
    /fx/5/par/32  FX32 F_XET
    /fx/5/par/33  FX32 F_XET
    /fx/5/par/34  FX32 F_XET
    /fx/5/par/35  FX32 F_XET
    /fx/5/par/36  FX32 F_XET
    /fx/5/par/37  FX32 F_XET
    /fx/5/par/38  FX32 F_XET
    /fx/5/par/39  FX32 F_XET
    /fx/5/par/40  FX32 F_XET
    /fx/5/par/41  FX32 F_XET
    /fx/5/par/42  FX32 F_XET
    /fx/5/par/43  FX32 F_XET
    /fx/5/par/44  FX32 F_XET
    /fx/5/par/45  FX32 F_XET
    /fx/5/par/46  FX32 F_XET
    /fx/5/par/47  FX32 F_XET
    /fx/5/par/48  FX32 F_XET
    /fx/5/par/49  FX32 F_XET
    /fx/5/par/50  FX32 F_XET
    /fx/5/par/51  FX32 F_XET
    /fx/5/par/52  FX32 F_XET
    /fx/5/par/53  FX32 F_XET
    /fx/5/par/54  FX32 F_XET
    /fx/5/par/55  FX32 F_XET
    /fx/5/par/56  FX32 F_XET
    /fx/5/par/57  FX32 F_XET
    /fx/5/par/58  FX32 F_XET
    /fx/5/par/59  FX32 F_XET
    /fx/5/par/60  FX32 F_XET
    /fx/5/par/61  FX32 F_XET
    /fx/5/par/62  FX32 F_XET
    /fx/5/par/63  FX32 F_XET
    /fx/5/par/64  FX32 F_XET
```

### Xfx6 (X32Fx.h, 68 entries)

```
/fx  <FXTYP1> n=0
/fx/6  <FXTYP2> n=0
    /fx/6/type  E32 F_XET enum=Sfxtyp2
/fx/6/par  <FXPAR2> n=0
    /fx/6/par/01  FX32 F_XET
    /fx/6/par/02  FX32 F_XET
    /fx/6/par/03  FX32 F_XET
    /fx/6/par/04  FX32 F_XET
    /fx/6/par/05  FX32 F_XET
    /fx/6/par/06  FX32 F_XET
    /fx/6/par/07  FX32 F_XET
    /fx/6/par/08  FX32 F_XET
    /fx/6/par/09  FX32 F_XET
    /fx/6/par/10  FX32 F_XET
    /fx/6/par/11  FX32 F_XET
    /fx/6/par/12  FX32 F_XET
    /fx/6/par/13  FX32 F_XET
    /fx/6/par/14  FX32 F_XET
    /fx/6/par/15  FX32 F_XET
    /fx/6/par/16  FX32 F_XET
    /fx/6/par/17  FX32 F_XET
    /fx/6/par/18  FX32 F_XET
    /fx/6/par/19  FX32 F_XET
    /fx/6/par/20  FX32 F_XET
    /fx/6/par/21  FX32 F_XET
    /fx/6/par/22  FX32 F_XET
    /fx/6/par/23  FX32 F_XET
    /fx/6/par/24  FX32 F_XET
    /fx/6/par/25  FX32 F_XET
    /fx/6/par/26  FX32 F_XET
    /fx/6/par/27  FX32 F_XET
    /fx/6/par/28  FX32 F_XET
    /fx/6/par/29  FX32 F_XET
    /fx/6/par/30  FX32 F_XET
    /fx/6/par/31  FX32 F_XET
    /fx/6/par/32  FX32 F_XET
    /fx/6/par/33  FX32 F_XET
    /fx/6/par/34  FX32 F_XET
    /fx/6/par/35  FX32 F_XET
    /fx/6/par/36  FX32 F_XET
    /fx/6/par/37  FX32 F_XET
    /fx/6/par/38  FX32 F_XET
    /fx/6/par/39  FX32 F_XET
    /fx/6/par/40  FX32 F_XET
    /fx/6/par/41  FX32 F_XET
    /fx/6/par/42  FX32 F_XET
    /fx/6/par/43  FX32 F_XET
    /fx/6/par/44  FX32 F_XET
    /fx/6/par/45  FX32 F_XET
    /fx/6/par/46  FX32 F_XET
    /fx/6/par/47  FX32 F_XET
    /fx/6/par/48  FX32 F_XET
    /fx/6/par/49  FX32 F_XET
    /fx/6/par/50  FX32 F_XET
    /fx/6/par/51  FX32 F_XET
    /fx/6/par/52  FX32 F_XET
    /fx/6/par/53  FX32 F_XET
    /fx/6/par/54  FX32 F_XET
    /fx/6/par/55  FX32 F_XET
    /fx/6/par/56  FX32 F_XET
    /fx/6/par/57  FX32 F_XET
    /fx/6/par/58  FX32 F_XET
    /fx/6/par/59  FX32 F_XET
    /fx/6/par/60  FX32 F_XET
    /fx/6/par/61  FX32 F_XET
    /fx/6/par/62  FX32 F_XET
    /fx/6/par/63  FX32 F_XET
    /fx/6/par/64  FX32 F_XET
```

### Xfx7 (X32Fx.h, 68 entries)

```
/fx  <FXTYP1> n=0
/fx/7  <FXTYP2> n=0
    /fx/7/type  E32 F_XET enum=Sfxtyp2
/fx/7/par  <FXPAR2> n=0
    /fx/7/par/01  FX32 F_XET
    /fx/7/par/02  FX32 F_XET
    /fx/7/par/03  FX32 F_XET
    /fx/7/par/04  FX32 F_XET
    /fx/7/par/05  FX32 F_XET
    /fx/7/par/06  FX32 F_XET
    /fx/7/par/07  FX32 F_XET
    /fx/7/par/08  FX32 F_XET
    /fx/7/par/09  FX32 F_XET
    /fx/7/par/10  FX32 F_XET
    /fx/7/par/11  FX32 F_XET
    /fx/7/par/12  FX32 F_XET
    /fx/7/par/13  FX32 F_XET
    /fx/7/par/14  FX32 F_XET
    /fx/7/par/15  FX32 F_XET
    /fx/7/par/16  FX32 F_XET
    /fx/7/par/17  FX32 F_XET
    /fx/7/par/18  FX32 F_XET
    /fx/7/par/19  FX32 F_XET
    /fx/7/par/20  FX32 F_XET
    /fx/7/par/21  FX32 F_XET
    /fx/7/par/22  FX32 F_XET
    /fx/7/par/23  FX32 F_XET
    /fx/7/par/24  FX32 F_XET
    /fx/7/par/25  FX32 F_XET
    /fx/7/par/26  FX32 F_XET
    /fx/7/par/27  FX32 F_XET
    /fx/7/par/28  FX32 F_XET
    /fx/7/par/29  FX32 F_XET
    /fx/7/par/30  FX32 F_XET
    /fx/7/par/31  FX32 F_XET
    /fx/7/par/32  FX32 F_XET
    /fx/7/par/33  FX32 F_XET
    /fx/7/par/34  FX32 F_XET
    /fx/7/par/35  FX32 F_XET
    /fx/7/par/36  FX32 F_XET
    /fx/7/par/37  FX32 F_XET
    /fx/7/par/38  FX32 F_XET
    /fx/7/par/39  FX32 F_XET
    /fx/7/par/40  FX32 F_XET
    /fx/7/par/41  FX32 F_XET
    /fx/7/par/42  FX32 F_XET
    /fx/7/par/43  FX32 F_XET
    /fx/7/par/44  FX32 F_XET
    /fx/7/par/45  FX32 F_XET
    /fx/7/par/46  FX32 F_XET
    /fx/7/par/47  FX32 F_XET
    /fx/7/par/48  FX32 F_XET
    /fx/7/par/49  FX32 F_XET
    /fx/7/par/50  FX32 F_XET
    /fx/7/par/51  FX32 F_XET
    /fx/7/par/52  FX32 F_XET
    /fx/7/par/53  FX32 F_XET
    /fx/7/par/54  FX32 F_XET
    /fx/7/par/55  FX32 F_XET
    /fx/7/par/56  FX32 F_XET
    /fx/7/par/57  FX32 F_XET
    /fx/7/par/58  FX32 F_XET
    /fx/7/par/59  FX32 F_XET
    /fx/7/par/60  FX32 F_XET
    /fx/7/par/61  FX32 F_XET
    /fx/7/par/62  FX32 F_XET
    /fx/7/par/63  FX32 F_XET
    /fx/7/par/64  FX32 F_XET
```

### Xfx8 (X32Fx.h, 68 entries)

```
/fx  <FXTYP1> n=0
/fx/8  <FXTYP2> n=0
    /fx/8/type  E32 F_XET enum=Sfxtyp2
/fx/8/par  <FXPAR2> n=0
    /fx/8/par/01  FX32 F_XET
    /fx/8/par/02  FX32 F_XET
    /fx/8/par/03  FX32 F_XET
    /fx/8/par/04  FX32 F_XET
    /fx/8/par/05  FX32 F_XET
    /fx/8/par/06  FX32 F_XET
    /fx/8/par/07  FX32 F_XET
    /fx/8/par/08  FX32 F_XET
    /fx/8/par/09  FX32 F_XET
    /fx/8/par/10  FX32 F_XET
    /fx/8/par/11  FX32 F_XET
    /fx/8/par/12  FX32 F_XET
    /fx/8/par/13  FX32 F_XET
    /fx/8/par/14  FX32 F_XET
    /fx/8/par/15  FX32 F_XET
    /fx/8/par/16  FX32 F_XET
    /fx/8/par/17  FX32 F_XET
    /fx/8/par/18  FX32 F_XET
    /fx/8/par/19  FX32 F_XET
    /fx/8/par/20  FX32 F_XET
    /fx/8/par/21  FX32 F_XET
    /fx/8/par/22  FX32 F_XET
    /fx/8/par/23  FX32 F_XET
    /fx/8/par/24  FX32 F_XET
    /fx/8/par/25  FX32 F_XET
    /fx/8/par/26  FX32 F_XET
    /fx/8/par/27  FX32 F_XET
    /fx/8/par/28  FX32 F_XET
    /fx/8/par/29  FX32 F_XET
    /fx/8/par/30  FX32 F_XET
    /fx/8/par/31  FX32 F_XET
    /fx/8/par/32  FX32 F_XET
    /fx/8/par/33  FX32 F_XET
    /fx/8/par/34  FX32 F_XET
    /fx/8/par/35  FX32 F_XET
    /fx/8/par/36  FX32 F_XET
    /fx/8/par/37  FX32 F_XET
    /fx/8/par/38  FX32 F_XET
    /fx/8/par/39  FX32 F_XET
    /fx/8/par/40  FX32 F_XET
    /fx/8/par/41  FX32 F_XET
    /fx/8/par/42  FX32 F_XET
    /fx/8/par/43  FX32 F_XET
    /fx/8/par/44  FX32 F_XET
    /fx/8/par/45  FX32 F_XET
    /fx/8/par/46  FX32 F_XET
    /fx/8/par/47  FX32 F_XET
    /fx/8/par/48  FX32 F_XET
    /fx/8/par/49  FX32 F_XET
    /fx/8/par/50  FX32 F_XET
    /fx/8/par/51  FX32 F_XET
    /fx/8/par/52  FX32 F_XET
    /fx/8/par/53  FX32 F_XET
    /fx/8/par/54  FX32 F_XET
    /fx/8/par/55  FX32 F_XET
    /fx/8/par/56  FX32 F_XET
    /fx/8/par/57  FX32 F_XET
    /fx/8/par/58  FX32 F_XET
    /fx/8/par/59  FX32 F_XET
    /fx/8/par/60  FX32 F_XET
    /fx/8/par/61  FX32 F_XET
    /fx/8/par/62  FX32 F_XET
    /fx/8/par/63  FX32 F_XET
    /fx/8/par/64  FX32 F_XET
```

### Xoutput (X32Output.h, 301 entries)

```
/outputs  <OMAIN> n=0
/outputs/main  <OMAIN> n=0
/outputs/main/01  <OMAIN> n=0
    /outputs/main/01/src  I32 F_XET
    /outputs/main/01/pos  E32 F_XET enum=Xotpos
    /outputs/main/01/invert  E32 F_XET enum=OffOn
/outputs/main/01/delay  <OMAIND> n=0
    /outputs/main/01/delay/on  E32 F_XET enum=OffOn
    /outputs/main/01/delay/time  F32 F_XET
/outputs/main/02  <OMAIN> n=0
    /outputs/main/02/src  I32 F_XET
    /outputs/main/02/pos  E32 F_XET enum=Xotpos
    /outputs/main/02/invert  E32 F_XET enum=OffOn
/outputs/main/02/delay  <OMAIND> n=0
    /outputs/main/02/delay/on  E32 F_XET enum=OffOn
    /outputs/main/02/delay/time  F32 F_XET
/outputs/main/03  <OMAIN> n=0
    /outputs/main/03/src  I32 F_XET
    /outputs/main/03/pos  E32 F_XET enum=Xotpos
    /outputs/main/03/invert  E32 F_XET enum=OffOn
/outputs/main/03/delay  <OMAIND> n=0
    /outputs/main/03/delay/on  E32 F_XET enum=OffOn
    /outputs/main/03/delay/time  F32 F_XET
/outputs/main/04  <OMAIN> n=0
    /outputs/main/04/src  I32 F_XET
    /outputs/main/04/pos  E32 F_XET enum=Xotpos
    /outputs/main/04/invert  E32 F_XET enum=OffOn
/outputs/main/04/delay  <OMAIND> n=0
    /outputs/main/04/delay/on  E32 F_XET
    /outputs/main/04/delay/time  F32 F_XET
/outputs/main/05  <OMAIN> n=0
    /outputs/main/05/src  I32 F_XET
    /outputs/main/05/pos  E32 F_XET enum=Xotpos
    /outputs/main/05/invert  E32 F_XET enum=OffOn
/outputs/main/05/delay  <OMAIND> n=0
    /outputs/main/05/delay/on  E32 F_XET enum=OffOn
    /outputs/main/05/delay/time  F32 F_XET
/outputs/main/06  <OMAIN> n=0
    /outputs/main/06/src  I32 F_XET
    /outputs/main/06/pos  E32 F_XET enum=Xotpos
    /outputs/main/06/invert  E32 F_XET enum=OffOn
/outputs/main/06/delay  <OMAIND> n=0
    /outputs/main/06/delay/on  E32 F_XET enum=OffOn
    /outputs/main/06/delay/time  F32 F_XET
/outputs/main/07  <OMAIN> n=0
    /outputs/main/07/src  I32 F_XET
    /outputs/main/07/pos  E32 F_XET enum=Xotpos
    /outputs/main/07/invert  E32 F_XET enum=OffOn
/outputs/main/07/delay  <OMAIND> n=0
    /outputs/main/07/delay/on  E32 F_XET enum=OffOn
    /outputs/main/07/delay/time  F32 F_XET
/outputs/main/08  <OMAIN> n=0
    /outputs/main/08/src  I32 F_XET
    /outputs/main/08/pos  E32 F_XET enum=Xotpos
    /outputs/main/08/invert  E32 F_XET enum=OffOn
/outputs/main/08/delay  <OMAIND> n=0
    /outputs/main/08/delay/on  E32 F_XET enum=OffOn
    /outputs/main/08/delay/time  F32 F_XET
/outputs/main/09  <OMAIN> n=0
    /outputs/main/09/src  I32 F_XET
    /outputs/main/09/pos  E32 F_XET enum=Xotpos
    /outputs/main/09/invert  E32 F_XET enum=OffOn
/outputs/main/09/delay  <OMAIND> n=0
    /outputs/main/09/delay/on  E32 F_XET enum=OffOn
    /outputs/main/09/delay/time  F32 F_XET
/outputs/main/10  <OMAIN> n=0
    /outputs/main/10/src  I32 F_XET
    /outputs/main/10/pos  E32 F_XET enum=Xotpos
    /outputs/main/10/invert  E32 F_XET enum=OffOn
/outputs/main/10/delay  <OMAIND> n=0
    /outputs/main/10/delay/on  E32 F_XET enum=OffOn
    /outputs/main/10/delay/time  F32 F_XET
/outputs/main/11  <OMAIN> n=0
    /outputs/main/11/src  I32 F_XET
    /outputs/main/11/pos  E32 F_XET enum=Xotpos
    /outputs/main/11/invert  E32 F_XET enum=OffOn
/outputs/main/11/delay  <OMAIND> n=0
    /outputs/main/11/delay/on  E32 F_XET enum=OffOn
    /outputs/main/11/delay/time  F32 F_XET
/outputs/main/12  <OMAIN> n=0
    /outputs/main/12/src  I32 F_XET
    /outputs/main/12/pos  E32 F_XET enum=Xotpos
    /outputs/main/12/invert  E32 F_XET enum=OffOn
/outputs/main/12/delay  <OMAIND> n=0
    /outputs/main/12/delay/on  E32 F_XET enum=OffOn
    /outputs/main/12/delay/time  F32 F_XET
/outputs/main/13  <OMAIN> n=0
    /outputs/main/13/src  I32 F_XET
    /outputs/main/13/pos  E32 F_XET enum=Xotpos
    /outputs/main/13/invert  E32 F_XET enum=OffOn
/outputs/main/13/delay  <OMAIND> n=0
    /outputs/main/13/delay/on  E32 F_XET enum=OffOn
    /outputs/main/13/delay/time  F32 F_XET
/outputs/main/14  <OMAIN> n=0
    /outputs/main/14/src  I32 F_XET
    /outputs/main/14/pos  E32 F_XET enum=Xotpos
    /outputs/main/14/invert  E32 F_XET enum=OffOn
/outputs/main/14/delay  <OMAIND> n=0
    /outputs/main/14/delay/on  E32 F_XET enum=OffOn
    /outputs/main/14/delay/time  F32 F_XET
/outputs/main/15  <OMAIN> n=0
    /outputs/main/15/src  I32 F_XET
    /outputs/main/15/pos  E32 F_XET enum=Xotpos
    /outputs/main/15/invert  E32 F_XET enum=OffOn
/outputs/main/15/delay  <OMAIND> n=0
    /outputs/main/15/delay/on  E32 F_XET
    /outputs/main/15/delay/time  F32 F_XET
/outputs/main/16  <OMAIN> n=0
    /outputs/main/16/src  I32 F_XET
    /outputs/main/16/pos  E32 F_XET enum=Xotpos
    /outputs/main/16/invert  E32 F_XET enum=OffOn
/outputs/main/16/delay  <OMAIND> n=0
    /outputs/main/16/delay/on  E32 F_XET enum=OffOn
    /outputs/main/16/delay/time  F32 F_XET
/outputs/aux/01  <OMAIN> n=0
    /outputs/aux/01/src  I32 F_XET
    /outputs/aux/01/pos  E32 F_XET enum=Xotpos
    /outputs/aux/01/invert  E32 F_XET enum=OffOn
/outputs/aux/02  <OMAIN> n=0
    /outputs/aux/02/src  I32 F_XET
    /outputs/aux/02/pos  E32 F_XET enum=Xotpos
    /outputs/aux/02/invert  E32 F_XET enum=OffOn
/outputs/aux/03  <OMAIN> n=0
    /outputs/aux/03/src  I32 F_XET
    /outputs/aux/03/pos  E32 F_XET enum=Xotpos
    /outputs/aux/03/invert  E32 F_XET enum=OffOn
/outputs/aux/04  <OMAIN> n=0
    /outputs/aux/04/src  I32 F_XET
    /outputs/aux/04/pos  E32 F_XET enum=Xotpos
    /outputs/aux/04/invert  E32 F_XET enum=OffOn
/outputs/aux/05  <OMAIN> n=0
    /outputs/aux/05/src  I32 F_XET
    /outputs/aux/05/pos  E32 F_XET enum=Xotpos
    /outputs/aux/05/invert  E32 F_XET enum=OffOn
/outputs/aux/06  <OMAIN> n=0
    /outputs/aux/06/src  I32 F_XET
    /outputs/aux/06/pos  E32 F_XET
    /outputs/aux/06/invert  E32 F_XET enum=OffOn
/outputs/p16  <OMAIN> n=0
/outputs/p16/01  <OMAIN> n=0
    /outputs/p16/01/src  I32 F_XET
    /outputs/p16/01/pos  E32 F_XET enum=Xotpos
    /outputs/p16/01/invert  E32 F_XET enum=OffOn
/outputs/p16/01/iQ  <OP16> n=0
    /outputs/p16/01/iQ/group  E32 F_XET enum=XiQgrp
    /outputs/p16/01/iQ/speaker  E32 F_XET enum=XiQspk
    /outputs/p16/01/iQ/eq  E32 F_XET enum=XiQeq
    /outputs/p16/01/iQ/model  I32 F_XET
/outputs/p16/02  <OMAIN> n=0
    /outputs/p16/02/src  I32 F_XET
    /outputs/p16/02/pos  E32 F_XET enum=Xotpos
    /outputs/p16/02/invert  E32 F_XET enum=OffOn
/outputs/p16/02/iQ  <OP16> n=0
    /outputs/p16/02/iQ/group  E32 F_XET enum=XiQgrp
    /outputs/p16/02/iQ/speaker  E32 F_XET enum=XiQspk
    /outputs/p16/02/iQ/eq  E32 F_XET enum=XiQeq
    /outputs/p16/02/iQ/model  I32 F_XET
/outputs/p16/03  <OMAIN> n=0
    /outputs/p16/03/src  I32 F_XET
    /outputs/p16/03/pos  E32 F_XET enum=Xotpos
    /outputs/p16/03/invert  E32 F_XET enum=OffOn
/outputs/p16/03/iQ  <OP16> n=0
    /outputs/p16/03/iQ/group  E32 F_XET enum=XiQgrp
    /outputs/p16/03/iQ/speaker  E32 F_XET enum=XiQspk
    /outputs/p16/03/iQ/eq  E32 F_XET enum=XiQeq
    /outputs/p16/03/iQ/model  I32 F_XET
/outputs/p16/04  <OMAIN> n=0
    /outputs/p16/04/src  I32 F_XET
    /outputs/p16/04/pos  E32 F_XET enum=Xotpos
    /outputs/p16/04/invert  E32 F_XET enum=OffOn
/outputs/p16/04/iQ  <OP16> n=0
    /outputs/p16/04/iQ/group  E32 F_XET enum=XiQgrp
    /outputs/p16/04/iQ/speaker  E32 F_XET enum=XiQspk
    /outputs/p16/04/iQ/eq  E32 F_XET enum=XiQeq
    /outputs/p16/04/iQ/model  I32 F_XET
/outputs/p16/05  <OMAIN> n=0
    /outputs/p16/05/src  I32 F_XET
    /outputs/p16/05/pos  E32 F_XET enum=Xotpos
    /outputs/p16/05/invert  E32 F_XET enum=OffOn
/outputs/p16/05/iQ  <OP16> n=0
    /outputs/p16/05/iQ/group  E32 F_XET enum=XiQgrp
    /outputs/p16/05/iQ/speaker  E32 F_XET enum=XiQspk
    /outputs/p16/05/iQ/eq  E32 F_XET enum=XiQeq
    /outputs/p16/05/iQ/model  I32 F_XET
/outputs/p16/06  <OMAIN> n=0
    /outputs/p16/06/src  I32 F_XET
    /outputs/p16/06/pos  E32 F_XET enum=Xotpos
    /outputs/p16/06/invert  E32 F_XET enum=OffOn
/outputs/p16/06/iQ  <OP16> n=0
    /outputs/p16/06/iQ/group  E32 F_XET enum=XiQgrp
    /outputs/p16/06/iQ/speaker  E32 F_XET enum=XiQspk
    /outputs/p16/06/iQ/eq  E32 F_XET enum=XiQeq
    /outputs/p16/06/iQ/model  I32 F_XET
/outputs/p16/07  <OMAIN> n=0
    /outputs/p16/07/src  I32 F_XET
    /outputs/p16/07/pos  E32 F_XET enum=Xotpos
    /outputs/p16/07/invert  E32 F_XET enum=OffOn
/outputs/p16/07/iQ  <OP16> n=0
    /outputs/p16/07/iQ/group  E32 F_XET enum=XiQgrp
    /outputs/p16/07/iQ/speaker  E32 F_XET enum=XiQspk
    /outputs/p16/07/iQ/eq  E32 F_XET enum=XiQeq
    /outputs/p16/07/iQ/model  I32 F_XET
/outputs/p16/08  <OMAIN> n=0
    /outputs/p16/08/src  I32 F_XET
    /outputs/p16/08/pos  E32 F_XET enum=Xotpos
    /outputs/p16/08/invert  E32 F_XET enum=OffOn
/outputs/p16/08/iQ  <OP16> n=0
    /outputs/p16/08/iQ/group  E32 F_XET enum=XiQgrp
    /outputs/p16/08/iQ/speaker  E32 F_XET enum=XiQspk
    /outputs/p16/08/iQ/eq  E32 F_XET enum=XiQeq
    /outputs/p16/08/iQ/model  I32 F_XET
/outputs/p16/09  <OMAIN> n=0
    /outputs/p16/09/src  I32 F_XET
    /outputs/p16/09/pos  E32 F_XET enum=Xotpos
    /outputs/p16/09/invert  E32 F_XET enum=OffOn
/outputs/p16/09/iQ  <OP16> n=0
    /outputs/p16/09/iQ/group  E32 F_XET enum=XiQgrp
    /outputs/p16/09/iQ/speaker  E32 F_XET enum=XiQspk
    /outputs/p16/09/iQ/eq  E32 F_XET enum=XiQeq
    /outputs/p16/09/iQ/model  I32 F_XET
/outputs/p16/10  <OMAIN> n=0
    /outputs/p16/10/src  I32 F_XET
    /outputs/p16/10/pos  E32 F_XET enum=Xotpos
    /outputs/p16/10/invert  E32 F_XET enum=OffOn
/outputs/p16/10/iQ  <OP16> n=0
    /outputs/p16/10/iQ/group  E32 F_XET enum=XiQgrp
    /outputs/p16/10/iQ/speaker  E32 F_XET enum=XiQspk
    /outputs/p16/10/iQ/eq  E32 F_XET enum=XiQeq
    /outputs/p16/10/iQ/model  I32 F_XET
/outputs/p16/11  <OMAIN> n=0
    /outputs/p16/11/src  I32 F_XET
    /outputs/p16/11/pos  E32 F_XET enum=Xotpos
    /outputs/p16/11/invert  E32 F_XET enum=OffOn
/outputs/p16/11/iQ  <OP16> n=0
    /outputs/p16/11/iQ/group  E32 F_XET enum=XiQgrp
    /outputs/p16/11/iQ/speaker  E32 F_XET enum=XiQspk
    /outputs/p16/11/iQ/eq  E32 F_XET enum=XiQeq
    /outputs/p16/11/iQ/model  I32 F_XET
/outputs/p16/12  <OMAIN> n=0
    /outputs/p16/12/src  I32 F_XET
    /outputs/p16/12/pos  E32 F_XET enum=Xotpos
    /outputs/p16/12/invert  E32 F_XET enum=OffOn
/outputs/p16/12/iQ  <OP16> n=0
    /outputs/p16/12/iQ/group  E32 F_XET enum=XiQgrp
    /outputs/p16/12/iQ/speaker  E32 F_XET enum=XiQspk
    /outputs/p16/12/iQ/eq  E32 F_XET enum=XiQeq
    /outputs/p16/12/iQ/model  I32 F_XET
/outputs/p16/13  <OMAIN> n=0
    /outputs/p16/13/src  I32 F_XET
    /outputs/p16/13/pos  E32 F_XET enum=Xotpos
    /outputs/p16/13/invert  E32 F_XET enum=OffOn
/outputs/p16/13/iQ  <OP16> n=0
    /outputs/p16/13/iQ/group  E32 F_XET enum=XiQgrp
    /outputs/p16/13/iQ/speaker  E32 F_XET enum=XiQspk
    /outputs/p16/13/iQ/eq  E32 F_XET enum=XiQeq
    /outputs/p16/13/iQ/model  I32 F_XET
/outputs/p16/14  <OMAIN> n=0
    /outputs/p16/14/src  I32 F_XET
    /outputs/p16/14/pos  E32 F_XET enum=Xotpos
    /outputs/p16/14/invert  E32 F_XET enum=OffOn
/outputs/p16/14/iQ  <OP16> n=0
    /outputs/p16/14/iQ/group  E32 F_XET enum=XiQgrp
    /outputs/p16/14/iQ/speaker  E32 F_XET enum=XiQspk
    /outputs/p16/14/iQ/eq  E32 F_XET enum=XiQeq
    /outputs/p16/14/iQ/model  I32 F_XET
/outputs/p16/15  <OMAIN> n=0
    /outputs/p16/15/src  I32 F_XET
    /outputs/p16/15/pos  E32 F_XET enum=Xotpos
    /outputs/p16/15/invert  E32 F_XET enum=OffOn
/outputs/p16/15/iQ  <OP16> n=0
    /outputs/p16/15/iQ/group  E32 F_XET enum=XiQgrp
    /outputs/p16/15/iQ/speaker  E32 F_XET enum=XiQspk
    /outputs/p16/15/iQ/eq  E32 F_XET enum=XiQeq
    /outputs/p16/15/iQ/model  I32 F_XET
/outputs/p16/16  <OMAIN> n=0
    /outputs/p16/16/src  I32 F_XET
    /outputs/p16/16/pos  E32 F_XET enum=Xotpos
    /outputs/p16/16/invert  E32 F_XET enum=OffOn
/outputs/p16/16/iQ  <OP16> n=0
    /outputs/p16/16/iQ/group  E32 F_XET enum=XiQgrp
    /outputs/p16/16/iQ/speaker  E32 F_XET enum=XiQspk
    /outputs/p16/16/iQ/eq  E32 F_XET enum=XiQeq
    /outputs/p16/16/iQ/model  I32 F_XET
/outputs/aes  <OMAIN> n=0
/outputs/aes/01  <OMAIN> n=0
    /outputs/aes/01/src  I32 F_XET
    /outputs/aes/01/pos  E32 F_XET enum=Xotpos
    /outputs/aes/01/invert  E32 F_XET enum=OffOn
/outputs/aes/02  <OMAIN> n=0
    /outputs/aes/02/src  I32 F_XET
    /outputs/aes/02/pos  E32 F_XET enum=Xotpos
    /outputs/aes/02/invert  E32 F_XET enum=OffOn
/outputs/rec  <OMAIN2> n=0
/outputs/rec/01  <OMAIN2> n=0
    /outputs/rec/01/src  I32 F_XET
    /outputs/rec/01/pos  E32 F_XET enum=Xotpos
    /outputs/rec/01/invert  E32 F_XET enum=OffOn
/outputs/rec/02  <OMAIN2> n=0
    /outputs/rec/02/src  I32 F_XET
    /outputs/rec/02/pos  E32 F_XET enum=Xotpos
    /outputs/rec/02/invert  E32 F_XET enum=OffOn
```

### Xheadamp (X32Headamp.h, 4 entries)

```
/headamp  <HAMP> n=0
/headamp/000  <HAMP> n=0
    /headamp/000/gain  F32 F_XET
    /headamp/000/phantom  E32 F_XET enum=OffOn
```

### Xheadamp001 (X32Headamp.h, 3 entries)

```
/headamp/001  <HAMP> n=0
    /headamp/001/gain  F32 F_XET
    /headamp/001/phantom  E32 F_XET enum=OffOn
```

### Xheadamp002 (X32Headamp.h, 3 entries)

```
/headamp/002  <HAMP> n=0
    /headamp/002/gain  F32 F_XET
    /headamp/002/phantom  E32 F_XET enum=OffOn
```

### Xheadamp003 (X32Headamp.h, 3 entries)

```
/headamp/003  <HAMP> n=0
    /headamp/003/gain  F32 F_XET
    /headamp/003/phantom  E32 F_XET enum=OffOn
```

### Xheadamp004 (X32Headamp.h, 3 entries)

```
/headamp/004  <HAMP> n=0
    /headamp/004/gain  F32 F_XET
    /headamp/004/phantom  E32 F_XET enum=OffOn
```

### Xheadamp005 (X32Headamp.h, 3 entries)

```
/headamp/005  <HAMP> n=0
    /headamp/005/gain  F32 F_XET
    /headamp/005/phantom  E32 F_XET enum=OffOn
```

### Xheadamp006 (X32Headamp.h, 3 entries)

```
/headamp/006  <HAMP> n=0
    /headamp/006/gain  F32 F_XET
    /headamp/006/phantom  E32 F_XET enum=OffOn
```

### Xheadamp007 (X32Headamp.h, 3 entries)

```
/headamp/007  <HAMP> n=0
    /headamp/007/gain  F32 F_XET
    /headamp/007/phantom  E32 F_XET enum=OffOn
```

### Xheadamp008 (X32Headamp.h, 3 entries)

```
/headamp/008  <HAMP> n=0
    /headamp/008/gain  F32 F_XET
    /headamp/008/phantom  E32 F_XET enum=OffOn
```

### Xheadamp009 (X32Headamp.h, 3 entries)

```
/headamp/009  <HAMP> n=0
    /headamp/009/gain  F32 F_XET
    /headamp/009/phantom  E32 F_XET enum=OffOn
```

### Xheadamp010 (X32Headamp.h, 3 entries)

```
/headamp/010  <HAMP> n=0
    /headamp/010/gain  F32 F_XET
    /headamp/010/phantom  E32 F_XET enum=OffOn
```

### Xheadamp011 (X32Headamp.h, 3 entries)

```
/headamp/011  <HAMP> n=0
    /headamp/011/gain  F32 F_XET
    /headamp/011/phantom  E32 F_XET enum=OffOn
```

### Xheadamp012 (X32Headamp.h, 3 entries)

```
/headamp/012  <HAMP> n=0
    /headamp/012/gain  F32 F_XET
    /headamp/012/phantom  E32 F_XET enum=OffOn
```

### Xheadamp013 (X32Headamp.h, 3 entries)

```
/headamp/013  <HAMP> n=0
    /headamp/013/gain  F32 F_XET
    /headamp/013/phantom  E32 F_XET enum=OffOn
```

### Xheadamp014 (X32Headamp.h, 3 entries)

```
/headamp/014  <HAMP> n=0
    /headamp/014/gain  F32 F_XET
    /headamp/014/phantom  E32 F_XET enum=OffOn
```

### Xheadamp015 (X32Headamp.h, 3 entries)

```
/headamp/015  <HAMP> n=0
    /headamp/015/gain  F32 F_XET
    /headamp/015/phantom  E32 F_XET enum=OffOn
```

### Xheadamp016 (X32Headamp.h, 3 entries)

```
/headamp/016  <HAMP> n=0
    /headamp/016/gain  F32 F_XET
    /headamp/016/phantom  E32 F_XET enum=OffOn
```

### Xheadamp017 (X32Headamp.h, 3 entries)

```
/headamp/017  <HAMP> n=0
    /headamp/017/gain  F32 F_XET
    /headamp/017/phantom  E32 F_XET enum=OffOn
```

### Xheadamp018 (X32Headamp.h, 3 entries)

```
/headamp/018  <HAMP> n=0
    /headamp/018/gain  F32 F_XET
    /headamp/018/phantom  E32 F_XET enum=OffOn
```

### Xheadamp019 (X32Headamp.h, 3 entries)

```
/headamp/019  <HAMP> n=0
    /headamp/019/gain  F32 F_XET
    /headamp/019/phantom  E32 F_XET enum=OffOn
```

### Xheadamp020 (X32Headamp.h, 3 entries)

```
/headamp/020  <HAMP> n=0
    /headamp/020/gain  F32 F_XET
    /headamp/020/phantom  E32 F_XET enum=OffOn
```

### Xheadamp021 (X32Headamp.h, 3 entries)

```
/headamp/021  <HAMP> n=0
    /headamp/021/gain  F32 F_XET
    /headamp/021/phantom  E32 F_XET enum=OffOn
```

### Xheadamp022 (X32Headamp.h, 3 entries)

```
/headamp/022  <HAMP> n=0
    /headamp/022/gain  F32 F_XET
    /headamp/022/phantom  E32 F_XET enum=OffOn
```

### Xheadamp023 (X32Headamp.h, 3 entries)

```
/headamp/023  <HAMP> n=0
    /headamp/023/gain  F32 F_XET
    /headamp/023/phantom  E32 F_XET enum=OffOn
```

### Xheadamp024 (X32Headamp.h, 3 entries)

```
/headamp/024  <HAMP> n=0
    /headamp/024/gain  F32 F_XET
    /headamp/024/phantom  E32 F_XET enum=OffOn
```

### Xheadamp025 (X32Headamp.h, 3 entries)

```
/headamp/025  <HAMP> n=0
    /headamp/025/gain  F32 F_XET
    /headamp/025/phantom  E32 F_XET enum=OffOn
```

### Xheadamp026 (X32Headamp.h, 3 entries)

```
/headamp/026  <HAMP> n=0
    /headamp/026/gain  F32 F_XET
    /headamp/026/phantom  E32 F_XET enum=OffOn
```

### Xheadamp027 (X32Headamp.h, 3 entries)

```
/headamp/027  <HAMP> n=0
    /headamp/027/gain  F32 F_XET
    /headamp/027/phantom  E32 F_XET enum=OffOn
```

### Xheadamp028 (X32Headamp.h, 3 entries)

```
/headamp/028  <HAMP> n=0
    /headamp/028/gain  F32 F_XET
    /headamp/028/phantom  E32 F_XET enum=OffOn
```

### Xheadamp029 (X32Headamp.h, 3 entries)

```
/headamp/029  <HAMP> n=0
    /headamp/029/gain  F32 F_XET
    /headamp/029/phantom  E32 F_XET enum=OffOn
```

### Xheadamp030 (X32Headamp.h, 3 entries)

```
/headamp/030  <HAMP> n=0
    /headamp/030/gain  F32 F_XET
    /headamp/030/phantom  E32 F_XET enum=OffOn
```

### Xheadamp031 (X32Headamp.h, 3 entries)

```
/headamp/031  <HAMP> n=0
    /headamp/031/gain  F32 F_XET
    /headamp/031/phantom  E32 F_XET enum=OffOn
```

### Xheadamp032 (X32Headamp.h, 3 entries)

```
/headamp/032  <HAMP> n=0
    /headamp/032/gain  F32 F_XET
    /headamp/032/phantom  E32 F_XET enum=OffOn
```

### Xheadamp033 (X32Headamp.h, 3 entries)

```
/headamp/033  <HAMP> n=0
    /headamp/033/gain  F32 F_XET
    /headamp/033/phantom  E32 F_XET enum=OffOn
```

### Xheadamp034 (X32Headamp.h, 3 entries)

```
/headamp/034  <HAMP> n=0
    /headamp/034/gain  F32 F_XET
    /headamp/034/phantom  E32 F_XET enum=OffOn
```

### Xheadamp035 (X32Headamp.h, 3 entries)

```
/headamp/035  <HAMP> n=0
    /headamp/035/gain  F32 F_XET
    /headamp/035/phantom  E32 F_XET enum=OffOn
```

### Xheadamp036 (X32Headamp.h, 3 entries)

```
/headamp/036  <HAMP> n=0
    /headamp/036/gain  F32 F_XET
    /headamp/036/phantom  E32 F_XET enum=OffOn
```

### Xheadamp037 (X32Headamp.h, 3 entries)

```
/headamp/037  <HAMP> n=0
    /headamp/037/gain  F32 F_XET
    /headamp/037/phantom  E32 F_XET enum=OffOn
```

### Xheadamp038 (X32Headamp.h, 3 entries)

```
/headamp/038  <HAMP> n=0
    /headamp/038/gain  F32 F_XET
    /headamp/038/phantom  E32 F_XET enum=OffOn
```

### Xheadamp039 (X32Headamp.h, 3 entries)

```
/headamp/039  <HAMP> n=0
    /headamp/039/gain  F32 F_XET
    /headamp/039/phantom  E32 F_XET enum=OffOn
```

### Xheadamp040 (X32Headamp.h, 3 entries)

```
/headamp/040  <HAMP> n=0
    /headamp/040/gain  F32 F_XET
    /headamp/040/phantom  E32 F_XET enum=OffOn
```

### Xheadamp041 (X32Headamp.h, 3 entries)

```
/headamp/041  <HAMP> n=0
    /headamp/041/gain  F32 F_XET
    /headamp/041/phantom  E32 F_XET enum=OffOn
```

### Xheadamp042 (X32Headamp.h, 3 entries)

```
/headamp/042  <HAMP> n=0
    /headamp/042/gain  F32 F_XET
    /headamp/042/phantom  E32 F_XET enum=OffOn
```

### Xheadamp043 (X32Headamp.h, 3 entries)

```
/headamp/043  <HAMP> n=0
    /headamp/043/gain  F32 F_XET
    /headamp/043/phantom  E32 F_XET enum=OffOn
```

### Xheadamp044 (X32Headamp.h, 3 entries)

```
/headamp/044  <HAMP> n=0
    /headamp/044/gain  F32 F_XET
    /headamp/044/phantom  E32 F_XET enum=OffOn
```

### Xheadamp045 (X32Headamp.h, 3 entries)

```
/headamp/045  <HAMP> n=0
    /headamp/045/gain  F32 F_XET
    /headamp/045/phantom  E32 F_XET enum=OffOn
```

### Xheadamp046 (X32Headamp.h, 3 entries)

```
/headamp/046  <HAMP> n=0
    /headamp/046/gain  F32 F_XET
    /headamp/046/phantom  E32 F_XET enum=OffOn
```

### Xheadamp047 (X32Headamp.h, 3 entries)

```
/headamp/047  <HAMP> n=0
    /headamp/047/gain  F32 F_XET
    /headamp/047/phantom  E32 F_XET enum=OffOn
```

### Xheadamp048 (X32Headamp.h, 3 entries)

```
/headamp/048  <HAMP> n=0
    /headamp/048/gain  F32 F_XET
    /headamp/048/phantom  E32 F_XET enum=OffOn
```

### Xheadamp049 (X32Headamp.h, 3 entries)

```
/headamp/049  <HAMP> n=0
    /headamp/049/gain  F32 F_XET
    /headamp/049/phantom  E32 F_XET enum=OffOn
```

### Xheadamp050 (X32Headamp.h, 3 entries)

```
/headamp/050  <HAMP> n=0
    /headamp/050/gain  F32 F_XET
    /headamp/050/phantom  E32 F_XET enum=OffOn
```

### Xheadamp051 (X32Headamp.h, 3 entries)

```
/headamp/051  <HAMP> n=0
    /headamp/051/gain  F32 F_XET
    /headamp/051/phantom  E32 F_XET enum=OffOn
```

### Xheadamp052 (X32Headamp.h, 3 entries)

```
/headamp/052  <HAMP> n=0
    /headamp/052/gain  F32 F_XET
    /headamp/052/phantom  E32 F_XET enum=OffOn
```

### Xheadamp053 (X32Headamp.h, 3 entries)

```
/headamp/053  <HAMP> n=0
    /headamp/053/gain  F32 F_XET
    /headamp/053/phantom  E32 F_XET enum=OffOn
```

### Xheadamp054 (X32Headamp.h, 3 entries)

```
/headamp/054  <HAMP> n=0
    /headamp/054/gain  F32 F_XET
    /headamp/054/phantom  E32 F_XET enum=OffOn
```

### Xheadamp055 (X32Headamp.h, 3 entries)

```
/headamp/055  <HAMP> n=0
    /headamp/055/gain  F32 F_XET
    /headamp/055/phantom  E32 F_XET enum=OffOn
```

### Xheadamp056 (X32Headamp.h, 3 entries)

```
/headamp/056  <HAMP> n=0
    /headamp/056/gain  F32 F_XET
    /headamp/056/phantom  E32 F_XET enum=OffOn
```

### Xheadamp057 (X32Headamp.h, 3 entries)

```
/headamp/057  <HAMP> n=0
    /headamp/057/gain  F32 F_XET
    /headamp/057/phantom  E32 F_XET enum=OffOn
```

### Xheadamp058 (X32Headamp.h, 3 entries)

```
/headamp/058  <HAMP> n=0
    /headamp/058/gain  F32 F_XET
    /headamp/058/phantom  E32 F_XET enum=OffOn
```

### Xheadamp059 (X32Headamp.h, 3 entries)

```
/headamp/059  <HAMP> n=0
    /headamp/059/gain  F32 F_XET
    /headamp/059/phantom  E32 F_XET enum=OffOn
```

### Xheadamp060 (X32Headamp.h, 3 entries)

```
/headamp/060  <HAMP> n=0
    /headamp/060/gain  F32 F_XET
    /headamp/060/phantom  E32 F_XET enum=OffOn
```

### Xheadamp061 (X32Headamp.h, 3 entries)

```
/headamp/061  <HAMP> n=0
    /headamp/061/gain  F32 F_XET
    /headamp/061/phantom  E32 F_XET enum=OffOn
```

### Xheadamp062 (X32Headamp.h, 3 entries)

```
/headamp/062  <HAMP> n=0
    /headamp/062/gain  F32 F_XET
    /headamp/062/phantom  E32 F_XET enum=OffOn
```

### Xheadamp063 (X32Headamp.h, 3 entries)

```
/headamp/063  <HAMP> n=0
    /headamp/063/gain  F32 F_XET
    /headamp/063/phantom  E32 F_XET enum=OffOn
```

### Xheadamp064 (X32Headamp.h, 3 entries)

```
/headamp/064  <HAMP> n=0
    /headamp/064/gain  F32 F_XET
    /headamp/064/phantom  E32 F_XET enum=OffOn
```

### Xheadamp065 (X32Headamp.h, 3 entries)

```
/headamp/065  <HAMP> n=0
    /headamp/065/gain  F32 F_XET
    /headamp/065/phantom  E32 F_XET enum=OffOn
```

### Xheadamp066 (X32Headamp.h, 3 entries)

```
/headamp/066  <HAMP> n=0
    /headamp/066/gain  F32 F_XET
    /headamp/066/phantom  E32 F_XET enum=OffOn
```

### Xheadamp067 (X32Headamp.h, 3 entries)

```
/headamp/067  <HAMP> n=0
    /headamp/067/gain  F32 F_XET
    /headamp/067/phantom  E32 F_XET enum=OffOn
```

### Xheadamp068 (X32Headamp.h, 3 entries)

```
/headamp/068  <HAMP> n=0
    /headamp/068/gain  F32 F_XET
    /headamp/068/phantom  E32 F_XET enum=OffOn
```

### Xheadamp069 (X32Headamp.h, 3 entries)

```
/headamp/069  <HAMP> n=0
    /headamp/069/gain  F32 F_XET
    /headamp/069/phantom  E32 F_XET enum=OffOn
```

### Xheadamp070 (X32Headamp.h, 3 entries)

```
/headamp/070  <HAMP> n=0
    /headamp/070/gain  F32 F_XET
    /headamp/070/phantom  E32 F_XET enum=OffOn
```

### Xheadamp071 (X32Headamp.h, 3 entries)

```
/headamp/071  <HAMP> n=0
    /headamp/071/gain  F32 F_XET
    /headamp/071/phantom  E32 F_XET enum=OffOn
```

### Xheadamp072 (X32Headamp.h, 3 entries)

```
/headamp/072  <HAMP> n=0
    /headamp/072/gain  F32 F_XET
    /headamp/072/phantom  E32 F_XET enum=OffOn
```

### Xheadamp073 (X32Headamp.h, 3 entries)

```
/headamp/073  <HAMP> n=0
    /headamp/073/gain  F32 F_XET
    /headamp/073/phantom  E32 F_XET enum=OffOn
```

### Xheadamp074 (X32Headamp.h, 3 entries)

```
/headamp/074  <HAMP> n=0
    /headamp/074/gain  F32 F_XET
    /headamp/074/phantom  E32 F_XET enum=OffOn
```

### Xheadamp075 (X32Headamp.h, 3 entries)

```
/headamp/075  <HAMP> n=0
    /headamp/075/gain  F32 F_XET
    /headamp/075/phantom  E32 F_XET enum=OffOn
```

### Xheadamp076 (X32Headamp.h, 3 entries)

```
/headamp/076  <HAMP> n=0
    /headamp/076/gain  F32 F_XET
    /headamp/076/phantom  E32 F_XET enum=OffOn
```

### Xheadamp077 (X32Headamp.h, 3 entries)

```
/headamp/077  <HAMP> n=0
    /headamp/077/gain  F32 F_XET
    /headamp/077/phantom  E32 F_XET enum=OffOn
```

### Xheadamp078 (X32Headamp.h, 3 entries)

```
/headamp/078  <HAMP> n=0
    /headamp/078/gain  F32 F_XET
    /headamp/078/phantom  E32 F_XET enum=OffOn
```

### Xheadamp079 (X32Headamp.h, 3 entries)

```
/headamp/079  <HAMP> n=0
    /headamp/079/gain  F32 F_XET
    /headamp/079/phantom  E32 F_XET enum=OffOn
```

### Xheadamp080 (X32Headamp.h, 3 entries)

```
/headamp/080  <HAMP> n=0
    /headamp/080/gain  F32 F_XET
    /headamp/080/phantom  E32 F_XET enum=OffOn
```

### Xheadamp081 (X32Headamp.h, 3 entries)

```
/headamp/081  <HAMP> n=0
    /headamp/081/gain  F32 F_XET
    /headamp/081/phantom  E32 F_XET enum=OffOn
```

### Xheadamp082 (X32Headamp.h, 3 entries)

```
/headamp/082  <HAMP> n=0
    /headamp/082/gain  F32 F_XET
    /headamp/082/phantom  E32 F_XET enum=OffOn
```

### Xheadamp083 (X32Headamp.h, 3 entries)

```
/headamp/083  <HAMP> n=0
    /headamp/083/gain  F32 F_XET
    /headamp/083/phantom  E32 F_XET enum=OffOn
```

### Xheadamp084 (X32Headamp.h, 3 entries)

```
/headamp/084  <HAMP> n=0
    /headamp/084/gain  F32 F_XET
    /headamp/084/phantom  E32 F_XET enum=OffOn
```

### Xheadamp085 (X32Headamp.h, 3 entries)

```
/headamp/085  <HAMP> n=0
    /headamp/085/gain  F32 F_XET
    /headamp/085/phantom  E32 F_XET enum=OffOn
```

### Xheadamp086 (X32Headamp.h, 3 entries)

```
/headamp/086  <HAMP> n=0
    /headamp/086/gain  F32 F_XET
    /headamp/086/phantom  E32 F_XET enum=OffOn
```

### Xheadamp087 (X32Headamp.h, 3 entries)

```
/headamp/087  <HAMP> n=0
    /headamp/087/gain  F32 F_XET
    /headamp/087/phantom  E32 F_XET enum=OffOn
```

### Xheadamp088 (X32Headamp.h, 3 entries)

```
/headamp/088  <HAMP> n=0
    /headamp/088/gain  F32 F_XET
    /headamp/088/phantom  E32 F_XET enum=OffOn
```

### Xheadamp089 (X32Headamp.h, 3 entries)

```
/headamp/089  <HAMP> n=0
    /headamp/089/gain  F32 F_XET
    /headamp/089/phantom  E32 F_XET enum=OffOn
```

### Xheadamp090 (X32Headamp.h, 3 entries)

```
/headamp/090  <HAMP> n=0
    /headamp/090/gain  F32 F_XET
    /headamp/090/phantom  E32 F_XET enum=OffOn
```

### Xheadamp091 (X32Headamp.h, 3 entries)

```
/headamp/091  <HAMP> n=0
    /headamp/091/gain  F32 F_XET
    /headamp/091/phantom  E32 F_XET enum=OffOn
```

### Xheadamp092 (X32Headamp.h, 3 entries)

```
/headamp/092  <HAMP> n=0
    /headamp/092/gain  F32 F_XET
    /headamp/092/phantom  E32 F_XET enum=OffOn
```

### Xheadamp093 (X32Headamp.h, 3 entries)

```
/headamp/093  <HAMP> n=0
    /headamp/093/gain  F32 F_XET
    /headamp/093/phantom  E32 F_XET enum=OffOn
```

### Xheadamp094 (X32Headamp.h, 3 entries)

```
/headamp/094  <HAMP> n=0
    /headamp/094/gain  F32 F_XET
    /headamp/094/phantom  E32 F_XET enum=OffOn
```

### Xheadamp095 (X32Headamp.h, 3 entries)

```
/headamp/095  <HAMP> n=0
    /headamp/095/gain  F32 F_XET
    /headamp/095/phantom  E32 F_XET enum=OffOn
```

### Xheadamp096 (X32Headamp.h, 3 entries)

```
/headamp/096  <HAMP> n=0
    /headamp/096/gain  F32 F_XET
    /headamp/096/phantom  E32 F_XET enum=OffOn
```

### Xheadamp097 (X32Headamp.h, 3 entries)

```
/headamp/097  <HAMP> n=0
    /headamp/097/gain  F32 F_XET
    /headamp/097/phantom  E32 F_XET enum=OffOn
```

### Xheadamp098 (X32Headamp.h, 3 entries)

```
/headamp/098  <HAMP> n=0
    /headamp/098/gain  F32 F_XET
    /headamp/098/phantom  E32 F_XET enum=OffOn
```

### Xheadamp099 (X32Headamp.h, 3 entries)

```
/headamp/099  <HAMP> n=0
    /headamp/099/gain  F32 F_XET
    /headamp/099/phantom  E32 F_XET enum=OffOn
```

### Xheadamp100 (X32Headamp.h, 3 entries)

```
/headamp/100  <HAMP> n=0
    /headamp/100/gain  F32 F_XET
    /headamp/100/phantom  E32 F_XET enum=OffOn
```

### Xheadamp101 (X32Headamp.h, 3 entries)

```
/headamp/101  <HAMP> n=0
    /headamp/101/gain  F32 F_XET
    /headamp/101/phantom  E32 F_XET enum=OffOn
```

### Xheadamp102 (X32Headamp.h, 3 entries)

```
/headamp/102  <HAMP> n=0
    /headamp/102/gain  F32 F_XET
    /headamp/102/phantom  E32 F_XET enum=OffOn
```

### Xheadamp103 (X32Headamp.h, 3 entries)

```
/headamp/103  <HAMP> n=0
    /headamp/103/gain  F32 F_XET
    /headamp/103/phantom  E32 F_XET enum=OffOn
```

### Xheadamp104 (X32Headamp.h, 3 entries)

```
/headamp/104  <HAMP> n=0
    /headamp/104/gain  F32 F_XET
    /headamp/104/phantom  E32 F_XET enum=OffOn
```

### Xheadamp105 (X32Headamp.h, 3 entries)

```
/headamp/105  <HAMP> n=0
    /headamp/105/gain  F32 F_XET
    /headamp/105/phantom  E32 F_XET enum=OffOn
```

### Xheadamp106 (X32Headamp.h, 3 entries)

```
/headamp/106  <HAMP> n=0
    /headamp/106/gain  F32 F_XET
    /headamp/106/phantom  E32 F_XET enum=OffOn
```

### Xheadamp107 (X32Headamp.h, 3 entries)

```
/headamp/107  <HAMP> n=0
    /headamp/107/gain  F32 F_XET
    /headamp/107/phantom  E32 F_XET enum=OffOn
```

### Xheadamp108 (X32Headamp.h, 3 entries)

```
/headamp/108  <HAMP> n=0
    /headamp/108/gain  F32 F_XET
    /headamp/108/phantom  E32 F_XET enum=OffOn
```

### Xheadamp109 (X32Headamp.h, 3 entries)

```
/headamp/109  <HAMP> n=0
    /headamp/109/gain  F32 F_XET
    /headamp/109/phantom  E32 F_XET enum=OffOn
```

### Xheadamp110 (X32Headamp.h, 3 entries)

```
/headamp/110  <HAMP> n=0
    /headamp/110/gain  F32 F_XET
    /headamp/110/phantom  E32 F_XET enum=OffOn
```

### Xheadamp111 (X32Headamp.h, 3 entries)

```
/headamp/111  <HAMP> n=0
    /headamp/111/gain  F32 F_XET
    /headamp/111/phantom  E32 F_XET enum=OffOn
```

### Xheadamp112 (X32Headamp.h, 3 entries)

```
/headamp/112  <HAMP> n=0
    /headamp/112/gain  F32 F_XET
    /headamp/112/phantom  E32 F_XET enum=OffOn
```

### Xheadamp113 (X32Headamp.h, 3 entries)

```
/headamp/113  <HAMP> n=0
    /headamp/113/gain  F32 F_XET
    /headamp/113/phantom  E32 F_XET enum=OffOn
```

### Xheadamp114 (X32Headamp.h, 3 entries)

```
/headamp/114  <HAMP> n=0
    /headamp/114/gain  F32 F_XET
    /headamp/114/phantom  E32 F_XET enum=OffOn
```

### Xheadamp115 (X32Headamp.h, 3 entries)

```
/headamp/115  <HAMP> n=0
    /headamp/115/gain  F32 F_XET
    /headamp/115/phantom  E32 F_XET enum=OffOn
```

### Xheadamp116 (X32Headamp.h, 3 entries)

```
/headamp/116  <HAMP> n=0
    /headamp/116/gain  F32 F_XET
    /headamp/116/phantom  E32 F_XET enum=OffOn
```

### Xheadamp117 (X32Headamp.h, 3 entries)

```
/headamp/117  <HAMP> n=0
    /headamp/117/gain  F32 F_XET
    /headamp/117/phantom  E32 F_XET enum=OffOn
```

### Xheadamp118 (X32Headamp.h, 3 entries)

```
/headamp/118  <HAMP> n=0
    /headamp/118/gain  F32 F_XET
    /headamp/118/phantom  E32 F_XET enum=OffOn
```

### Xheadamp119 (X32Headamp.h, 3 entries)

```
/headamp/119  <HAMP> n=0
    /headamp/119/gain  F32 F_XET
    /headamp/119/phantom  E32 F_XET enum=OffOn
```

### Xheadamp120 (X32Headamp.h, 3 entries)

```
/headamp/120  <HAMP> n=0
    /headamp/120/gain  F32 F_XET
    /headamp/120/phantom  E32 F_XET enum=OffOn
```

### Xheadamp121 (X32Headamp.h, 3 entries)

```
/headamp/121  <HAMP> n=0
    /headamp/121/gain  F32 F_XET
    /headamp/121/phantom  E32 F_XET enum=OffOn
```

### Xheadamp122 (X32Headamp.h, 3 entries)

```
/headamp/122  <HAMP> n=0
    /headamp/122/gain  F32 F_XET
    /headamp/122/phantom  E32 F_XET enum=OffOn
```

### Xheadamp123 (X32Headamp.h, 3 entries)

```
/headamp/123  <HAMP> n=0
    /headamp/123/gain  F32 F_XET
    /headamp/123/phantom  E32 F_XET enum=OffOn
```

### Xheadamp124 (X32Headamp.h, 3 entries)

```
/headamp/124  <HAMP> n=0
    /headamp/124/gain  F32 F_XET
    /headamp/124/phantom  E32 F_XET enum=OffOn
```

### Xheadamp125 (X32Headamp.h, 3 entries)

```
/headamp/125  <HAMP> n=0
    /headamp/125/gain  F32 F_XET
    /headamp/125/phantom  E32 F_XET enum=OffOn
```

### Xheadamp126 (X32Headamp.h, 3 entries)

```
/headamp/126  <HAMP> n=0
    /headamp/126/gain  F32 F_XET
    /headamp/126/phantom  E32 F_XET enum=OffOn
```

### Xheadamp127 (X32Headamp.h, 3 entries)

```
/headamp/127  <HAMP> n=0
    /headamp/127/gain  F32 F_XET
    /headamp/127/phantom  E32 F_XET enum=OffOn
```

### Xshow (X32Show.h, 6016 entries)

```
    /-show/prepos/current  I32 F_XET
/-show  <SNAM> n=0
/-show/showfile  <SNAM> n=0
/-show/showfile/show  <SNAM> n=0
    /-show/showfile/show/name  S32 F_XET
    /-show/showfile/show/inputs  I32 F_XET
    /-show/showfile/show/mxsends  I32 F_XET
    /-show/showfile/show/mxbuses  I32 F_XET
    /-show/showfile/show/console  I32 F_XET
    /-show/showfile/show/chan16  I32 F_XET
    /-show/showfile/show/chan32  I32 F_XET
    /-show/showfile/show/return  I32 F_XET
    /-show/showfile/show/buses  I32 F_XET
    /-show/showfile/show/lrmtxdca  I32 F_XET
    /-show/showfile/show/effects  I32 F_XET
/-show/showfile/cue  <SCUE> n=0
/-show/showfile/cue/000  <SCUE> n=0
    /-show/showfile/cue/000/numb  I32 F_XET
    /-show/showfile/cue/000/name  S32 F_XET
    /-show/showfile/cue/000/skip  I32 F_XET
    /-show/showfile/cue/000/scene  I32 F_XET
    /-show/showfile/cue/000/bit  I32 F_XET
    /-show/showfile/cue/000/miditype  I32 F_XET
    /-show/showfile/cue/000/midichan  I32 F_XET
    /-show/showfile/cue/000/midipara1  I32 F_XET
    /-show/showfile/cue/000/midipara2  I32 F_XET
/-show/showfile/cue/001  <SCUE> n=0
    /-show/showfile/cue/001/numb  I32 F_XET
    /-show/showfile/cue/001/name  S32 F_XET
    /-show/showfile/cue/001/skip  I32 F_XET
    /-show/showfile/cue/001/scene  I32 F_XET
    /-show/showfile/cue/001/bit  I32 F_XET
    /-show/showfile/cue/001/miditype  I32 F_XET
    /-show/showfile/cue/001/midichan  I32 F_XET
    /-show/showfile/cue/001/midipara1  I32 F_XET
    /-show/showfile/cue/001/midipara2  I32 F_XET
/-show/showfile/cue/002  <SCUE> n=0
    /-show/showfile/cue/002/numb  I32 F_XET
    /-show/showfile/cue/002/name  S32 F_XET
    /-show/showfile/cue/002/skip  I32 F_XET
    /-show/showfile/cue/002/scene  I32 F_XET
    /-show/showfile/cue/002/bit  I32 F_XET
    /-show/showfile/cue/002/miditype  I32 F_XET
    /-show/showfile/cue/002/midichan  I32 F_XET
    /-show/showfile/cue/002/midipara1  I32 F_XET
    /-show/showfile/cue/002/midipara2  I32 F_XET
/-show/showfile/cue/003  <SCUE> n=0
    /-show/showfile/cue/003/numb  I32 F_XET
    /-show/showfile/cue/003/name  S32 F_XET
    /-show/showfile/cue/003/skip  I32 F_XET
    /-show/showfile/cue/003/scene  I32 F_XET
    /-show/showfile/cue/003/bit  I32 F_XET
    /-show/showfile/cue/003/miditype  I32 F_XET
    /-show/showfile/cue/003/midichan  I32 F_XET
    /-show/showfile/cue/003/midipara1  I32 F_XET
    /-show/showfile/cue/003/midipara2  I32 F_XET
/-show/showfile/cue/004  <SCUE> n=0
    /-show/showfile/cue/004/numb  I32 F_XET
    /-show/showfile/cue/004/name  S32 F_XET
    /-show/showfile/cue/004/skip  I32 F_XET
    /-show/showfile/cue/004/scene  I32 F_XET
    /-show/showfile/cue/004/bit  I32 F_XET
    /-show/showfile/cue/004/miditype  I32 F_XET
    /-show/showfile/cue/004/midichan  I32 F_XET
    /-show/showfile/cue/004/midipara1  I32 F_XET
    /-show/showfile/cue/004/midipara2  I32 F_XET
/-show/showfile/cue/005  <SCUE> n=0
    /-show/showfile/cue/005/numb  I32 F_XET
    /-show/showfile/cue/005/name  S32 F_XET
    /-show/showfile/cue/005/skip  I32 F_XET
    /-show/showfile/cue/005/scene  I32 F_XET
    /-show/showfile/cue/005/bit  I32 F_XET
    /-show/showfile/cue/005/miditype  I32 F_XET
    /-show/showfile/cue/005/midichan  I32 F_XET
    /-show/showfile/cue/005/midipara1  I32 F_XET
    /-show/showfile/cue/005/midipara2  I32 F_XET
/-show/showfile/cue/006  <SCUE> n=0
    /-show/showfile/cue/006/numb  I32 F_XET
    /-show/showfile/cue/006/name  S32 F_XET
    /-show/showfile/cue/006/skip  I32 F_XET
    /-show/showfile/cue/006/scene  I32 F_XET
    /-show/showfile/cue/006/bit  I32 F_XET
    /-show/showfile/cue/006/miditype  I32 F_XET
    /-show/showfile/cue/006/midichan  I32 F_XET
    /-show/showfile/cue/006/midipara1  I32 F_XET
    /-show/showfile/cue/006/midipara2  I32 F_XET
/-show/showfile/cue/007  <SCUE> n=0
    /-show/showfile/cue/007/numb  I32 F_XET
    /-show/showfile/cue/007/name  S32 F_XET
    /-show/showfile/cue/007/skip  I32 F_XET
    /-show/showfile/cue/007/scene  I32 F_XET
    /-show/showfile/cue/007/bit  I32 F_XET
    /-show/showfile/cue/007/miditype  I32 F_XET
    /-show/showfile/cue/007/midichan  I32 F_XET
    /-show/showfile/cue/007/midipara1  I32 F_XET
    /-show/showfile/cue/007/midipara2  I32 F_XET
/-show/showfile/cue/008  <SCUE> n=0
    /-show/showfile/cue/008/numb  I32 F_XET
    /-show/showfile/cue/008/name  S32 F_XET
    /-show/showfile/cue/008/skip  I32 F_XET
    /-show/showfile/cue/008/scene  I32 F_XET
    /-show/showfile/cue/008/bit  I32 F_XET
    /-show/showfile/cue/008/miditype  I32 F_XET
    /-show/showfile/cue/008/midichan  I32 F_XET
    /-show/showfile/cue/008/midipara1  I32 F_XET
    /-show/showfile/cue/008/midipara2  I32 F_XET
/-show/showfile/cue/009  <SCUE> n=0
    /-show/showfile/cue/009/numb  I32 F_XET
    /-show/showfile/cue/009/name  S32 F_XET
    /-show/showfile/cue/009/skip  I32 F_XET
    /-show/showfile/cue/009/scene  I32 F_XET
    /-show/showfile/cue/009/bit  I32 F_XET
    /-show/showfile/cue/009/miditype  I32 F_XET
    /-show/showfile/cue/009/midichan  I32 F_XET
    /-show/showfile/cue/009/midipara1  I32 F_XET
    /-show/showfile/cue/009/midipara2  I32 F_XET
/-show/showfile/cue/010  <SCUE> n=0
    /-show/showfile/cue/010/numb  I32 F_XET
    /-show/showfile/cue/010/name  S32 F_XET
    /-show/showfile/cue/010/skip  I32 F_XET
    /-show/showfile/cue/010/scene  I32 F_XET
    /-show/showfile/cue/010/bit  I32 F_XET
    /-show/showfile/cue/010/miditype  I32 F_XET
    /-show/showfile/cue/010/midichan  I32 F_XET
    /-show/showfile/cue/010/midipara1  I32 F_XET
    /-show/showfile/cue/010/midipara2  I32 F_XET
/-show/showfile/cue/011  <SCUE> n=0
    /-show/showfile/cue/011/numb  I32 F_XET
    /-show/showfile/cue/011/name  S32 F_XET
    /-show/showfile/cue/011/skip  I32 F_XET
    /-show/showfile/cue/011/scene  I32 F_XET
    /-show/showfile/cue/011/bit  I32 F_XET
    /-show/showfile/cue/011/miditype  I32 F_XET
    /-show/showfile/cue/011/midichan  I32 F_XET
    /-show/showfile/cue/011/midipara1  I32 F_XET
    /-show/showfile/cue/011/midipara2  I32 F_XET
/-show/showfile/cue/012  <SCUE> n=0
    /-show/showfile/cue/012/numb  I32 F_XET
    /-show/showfile/cue/012/name  S32 F_XET
    /-show/showfile/cue/012/skip  I32 F_XET
    /-show/showfile/cue/012/scene  I32 F_XET
    /-show/showfile/cue/012/bit  I32 F_XET
    /-show/showfile/cue/012/miditype  I32 F_XET
    /-show/showfile/cue/012/midichan  I32 F_XET
    /-show/showfile/cue/012/midipara1  I32 F_XET
    /-show/showfile/cue/012/midipara2  I32 F_XET
/-show/showfile/cue/013  <SCUE> n=0
    /-show/showfile/cue/013/numb  I32 F_XET
    /-show/showfile/cue/013/name  S32 F_XET
    /-show/showfile/cue/013/skip  I32 F_XET
    /-show/showfile/cue/013/scene  I32 F_XET
    /-show/showfile/cue/013/bit  I32 F_XET
    /-show/showfile/cue/013/miditype  I32 F_XET
    /-show/showfile/cue/013/midichan  I32 F_XET
    /-show/showfile/cue/013/midipara1  I32 F_XET
    /-show/showfile/cue/013/midipara2  I32 F_XET
/-show/showfile/cue/014  <SCUE> n=0
    /-show/showfile/cue/014/numb  I32 F_XET
    /-show/showfile/cue/014/name  S32 F_XET
    /-show/showfile/cue/014/skip  I32 F_XET
    /-show/showfile/cue/014/scene  I32 F_XET
    /-show/showfile/cue/014/bit  I32 F_XET
    /-show/showfile/cue/014/miditype  I32 F_XET
    /-show/showfile/cue/014/midichan  I32 F_XET
    /-show/showfile/cue/014/midipara1  I32 F_XET
    /-show/showfile/cue/014/midipara2  I32 F_XET
/-show/showfile/cue/015  <SCUE> n=0
    /-show/showfile/cue/015/numb  I32 F_XET
    /-show/showfile/cue/015/name  S32 F_XET
    /-show/showfile/cue/015/skip  I32 F_XET
    /-show/showfile/cue/015/scene  I32 F_XET
    /-show/showfile/cue/015/bit  I32 F_XET
    /-show/showfile/cue/015/miditype  I32 F_XET
    /-show/showfile/cue/015/midichan  I32 F_XET
    /-show/showfile/cue/015/midipara1  I32 F_XET
    /-show/showfile/cue/015/midipara2  I32 F_XET
/-show/showfile/cue/016  <SCUE> n=0
    /-show/showfile/cue/016/numb  I32 F_XET
    /-show/showfile/cue/016/name  S32 F_XET
    /-show/showfile/cue/016/skip  I32 F_XET
    /-show/showfile/cue/016/scene  I32 F_XET
    /-show/showfile/cue/016/bit  I32 F_XET
    /-show/showfile/cue/016/miditype  I32 F_XET
    /-show/showfile/cue/016/midichan  I32 F_XET
    /-show/showfile/cue/016/midipara1  I32 F_XET
    /-show/showfile/cue/016/midipara2  I32 F_XET
/-show/showfile/cue/017  <SCUE> n=0
    /-show/showfile/cue/017/numb  I32 F_XET
    /-show/showfile/cue/017/name  S32 F_XET
    /-show/showfile/cue/017/skip  I32 F_XET
    /-show/showfile/cue/017/scene  I32 F_XET
    /-show/showfile/cue/017/bit  I32 F_XET
    /-show/showfile/cue/017/miditype  I32 F_XET
    /-show/showfile/cue/017/midichan  I32 F_XET
    /-show/showfile/cue/017/midipara1  I32 F_XET
    /-show/showfile/cue/017/midipara2  I32 F_XET
/-show/showfile/cue/018  <SCUE> n=0
    /-show/showfile/cue/018/numb  I32 F_XET
    /-show/showfile/cue/018/name  S32 F_XET
    /-show/showfile/cue/018/skip  I32 F_XET
    /-show/showfile/cue/018/scene  I32 F_XET
    /-show/showfile/cue/018/bit  I32 F_XET
    /-show/showfile/cue/018/miditype  I32 F_XET
    /-show/showfile/cue/018/midichan  I32 F_XET
    /-show/showfile/cue/018/midipara1  I32 F_XET
    /-show/showfile/cue/018/midipara2  I32 F_XET
/-show/showfile/cue/019  <SCUE> n=0
    /-show/showfile/cue/019/numb  I32 F_XET
    /-show/showfile/cue/019/name  S32 F_XET
    /-show/showfile/cue/019/skip  I32 F_XET
    /-show/showfile/cue/019/scene  I32 F_XET
    /-show/showfile/cue/019/bit  I32 F_XET
    /-show/showfile/cue/019/miditype  I32 F_XET
    /-show/showfile/cue/019/midichan  I32 F_XET
    /-show/showfile/cue/019/midipara1  I32 F_XET
    /-show/showfile/cue/019/midipara2  I32 F_XET
/-show/showfile/cue/020  <SCUE> n=0
    /-show/showfile/cue/020/numb  I32 F_XET
    /-show/showfile/cue/020/name  S32 F_XET
    /-show/showfile/cue/020/skip  I32 F_XET
    /-show/showfile/cue/020/scene  I32 F_XET
    /-show/showfile/cue/020/bit  I32 F_XET
    /-show/showfile/cue/020/miditype  I32 F_XET
    /-show/showfile/cue/020/midichan  I32 F_XET
    /-show/showfile/cue/020/midipara1  I32 F_XET
    /-show/showfile/cue/020/midipara2  I32 F_XET
/-show/showfile/cue/021  <SCUE> n=0
    /-show/showfile/cue/021/numb  I32 F_XET
    /-show/showfile/cue/021/name  S32 F_XET
    /-show/showfile/cue/021/skip  I32 F_XET
    /-show/showfile/cue/021/scene  I32 F_XET
    /-show/showfile/cue/021/bit  I32 F_XET
    /-show/showfile/cue/021/miditype  I32 F_XET
    /-show/showfile/cue/021/midichan  I32 F_XET
    /-show/showfile/cue/021/midipara1  I32 F_XET
    /-show/showfile/cue/021/midipara2  I32 F_XET
/-show/showfile/cue/022  <SCUE> n=0
    /-show/showfile/cue/022/numb  I32 F_XET
    /-show/showfile/cue/022/name  S32 F_XET
    /-show/showfile/cue/022/skip  I32 F_XET
    /-show/showfile/cue/022/scene  I32 F_XET
    /-show/showfile/cue/022/bit  I32 F_XET
    /-show/showfile/cue/022/miditype  I32 F_XET
    /-show/showfile/cue/022/midichan  I32 F_XET
    /-show/showfile/cue/022/midipara1  I32 F_XET
    /-show/showfile/cue/022/midipara2  I32 F_XET
/-show/showfile/cue/023  <SCUE> n=0
    /-show/showfile/cue/023/numb  I32 F_XET
    /-show/showfile/cue/023/name  S32 F_XET
    /-show/showfile/cue/023/skip  I32 F_XET
    /-show/showfile/cue/023/scene  I32 F_XET
    /-show/showfile/cue/023/bit  I32 F_XET
    /-show/showfile/cue/023/miditype  I32 F_XET
    /-show/showfile/cue/023/midichan  I32 F_XET
    /-show/showfile/cue/023/midipara1  I32 F_XET
    /-show/showfile/cue/023/midipara2  I32 F_XET
/-show/showfile/cue/024  <SCUE> n=0
    /-show/showfile/cue/024/numb  I32 F_XET
    /-show/showfile/cue/024/name  S32 F_XET
    /-show/showfile/cue/024/skip  I32 F_XET
    /-show/showfile/cue/024/scene  I32 F_XET
    /-show/showfile/cue/024/bit  I32 F_XET
    /-show/showfile/cue/024/miditype  I32 F_XET
    /-show/showfile/cue/024/midichan  I32 F_XET
    /-show/showfile/cue/024/midipara1  I32 F_XET
    /-show/showfile/cue/024/midipara2  I32 F_XET
/-show/showfile/cue/025  <SCUE> n=0
    /-show/showfile/cue/025/numb  I32 F_XET
    /-show/showfile/cue/025/name  S32 F_XET
    /-show/showfile/cue/025/skip  I32 F_XET
    /-show/showfile/cue/025/scene  I32 F_XET
    /-show/showfile/cue/025/bit  I32 F_XET
    /-show/showfile/cue/025/miditype  I32 F_XET
    /-show/showfile/cue/025/midichan  I32 F_XET
    /-show/showfile/cue/025/midipara1  I32 F_XET
    /-show/showfile/cue/025/midipara2  I32 F_XET
/-show/showfile/cue/026  <SCUE> n=0
    /-show/showfile/cue/026/numb  I32 F_XET
    /-show/showfile/cue/026/name  S32 F_XET
    /-show/showfile/cue/026/skip  I32 F_XET
    /-show/showfile/cue/026/scene  I32 F_XET
    /-show/showfile/cue/026/bit  I32 F_XET
    /-show/showfile/cue/026/miditype  I32 F_XET
    /-show/showfile/cue/026/midichan  I32 F_XET
    /-show/showfile/cue/026/midipara1  I32 F_XET
    /-show/showfile/cue/026/midipara2  I32 F_XET
/-show/showfile/cue/027  <SCUE> n=0
    /-show/showfile/cue/027/numb  I32 F_XET
    /-show/showfile/cue/027/name  S32 F_XET
    /-show/showfile/cue/027/skip  I32 F_XET
    /-show/showfile/cue/027/scene  I32 F_XET
    /-show/showfile/cue/027/bit  I32 F_XET
    /-show/showfile/cue/027/miditype  I32 F_XET
    /-show/showfile/cue/027/midichan  I32 F_XET
    /-show/showfile/cue/027/midipara1  I32 F_XET
    /-show/showfile/cue/027/midipara2  I32 F_XET
/-show/showfile/cue/028  <SCUE> n=0
    /-show/showfile/cue/028/numb  I32 F_XET
    /-show/showfile/cue/028/name  S32 F_XET
    /-show/showfile/cue/028/skip  I32 F_XET
    /-show/showfile/cue/028/scene  I32 F_XET
    /-show/showfile/cue/028/bit  I32 F_XET
    /-show/showfile/cue/028/miditype  I32 F_XET
    /-show/showfile/cue/028/midichan  I32 F_XET
    /-show/showfile/cue/028/midipara1  I32 F_XET
    /-show/showfile/cue/028/midipara2  I32 F_XET
/-show/showfile/cue/029  <SCUE> n=0
    /-show/showfile/cue/029/numb  I32 F_XET
    /-show/showfile/cue/029/name  S32 F_XET
    /-show/showfile/cue/029/skip  I32 F_XET
    /-show/showfile/cue/029/scene  I32 F_XET
    /-show/showfile/cue/029/bit  I32 F_XET
    /-show/showfile/cue/029/miditype  I32 F_XET
    /-show/showfile/cue/029/midichan  I32 F_XET
    /-show/showfile/cue/029/midipara1  I32 F_XET
    /-show/showfile/cue/029/midipara2  I32 F_XET
/-show/showfile/cue/030  <SCUE> n=0
    /-show/showfile/cue/030/numb  I32 F_XET
    /-show/showfile/cue/030/name  S32 F_XET
    /-show/showfile/cue/030/skip  I32 F_XET
    /-show/showfile/cue/030/scene  I32 F_XET
    /-show/showfile/cue/030/bit  I32 F_XET
    /-show/showfile/cue/030/miditype  I32 F_XET
    /-show/showfile/cue/030/midichan  I32 F_XET
    /-show/showfile/cue/030/midipara1  I32 F_XET
    /-show/showfile/cue/030/midipara2  I32 F_XET
/-show/showfile/cue/031  <SCUE> n=0
    /-show/showfile/cue/031/numb  I32 F_XET
    /-show/showfile/cue/031/name  S32 F_XET
    /-show/showfile/cue/031/skip  I32 F_XET
    /-show/showfile/cue/031/scene  I32 F_XET
    /-show/showfile/cue/031/bit  I32 F_XET
    /-show/showfile/cue/031/miditype  I32 F_XET
    /-show/showfile/cue/031/midichan  I32 F_XET
    /-show/showfile/cue/031/midipara1  I32 F_XET
    /-show/showfile/cue/031/midipara2  I32 F_XET
/-show/showfile/cue/032  <SCUE> n=0
    /-show/showfile/cue/032/numb  I32 F_XET
    /-show/showfile/cue/032/name  S32 F_XET
    /-show/showfile/cue/032/skip  I32 F_XET
    /-show/showfile/cue/032/scene  I32 F_XET
    /-show/showfile/cue/032/bit  I32 F_XET
    /-show/showfile/cue/032/miditype  I32 F_XET
    /-show/showfile/cue/032/midichan  I32 F_XET
    /-show/showfile/cue/032/midipara1  I32 F_XET
    /-show/showfile/cue/032/midipara2  I32 F_XET
/-show/showfile/cue/033  <SCUE> n=0
    /-show/showfile/cue/033/numb  I32 F_XET
    /-show/showfile/cue/033/name  S32 F_XET
    /-show/showfile/cue/033/skip  I32 F_XET
    /-show/showfile/cue/033/scene  I32 F_XET
    /-show/showfile/cue/033/bit  I32 F_XET
    /-show/showfile/cue/033/miditype  I32 F_XET
    /-show/showfile/cue/033/midichan  I32 F_XET
    /-show/showfile/cue/033/midipara1  I32 F_XET
    /-show/showfile/cue/033/midipara2  I32 F_XET
/-show/showfile/cue/034  <SCUE> n=0
    /-show/showfile/cue/034/numb  I32 F_XET
    /-show/showfile/cue/034/name  S32 F_XET
    /-show/showfile/cue/034/skip  I32 F_XET
    /-show/showfile/cue/034/scene  I32 F_XET
    /-show/showfile/cue/034/bit  I32 F_XET
    /-show/showfile/cue/034/miditype  I32 F_XET
    /-show/showfile/cue/034/midichan  I32 F_XET
    /-show/showfile/cue/034/midipara1  I32 F_XET
    /-show/showfile/cue/034/midipara2  I32 F_XET
/-show/showfile/cue/035  <SCUE> n=0
    /-show/showfile/cue/035/numb  I32 F_XET
    /-show/showfile/cue/035/name  S32 F_XET
    /-show/showfile/cue/035/skip  I32 F_XET
    /-show/showfile/cue/035/scene  I32 F_XET
    /-show/showfile/cue/035/bit  I32 F_XET
    /-show/showfile/cue/035/miditype  I32 F_XET
    /-show/showfile/cue/035/midichan  I32 F_XET
    /-show/showfile/cue/035/midipara1  I32 F_XET
    /-show/showfile/cue/035/midipara2  I32 F_XET
/-show/showfile/cue/036  <SCUE> n=0
    /-show/showfile/cue/036/numb  I32 F_XET
    /-show/showfile/cue/036/name  S32 F_XET
    /-show/showfile/cue/036/skip  I32 F_XET
    /-show/showfile/cue/036/scene  I32 F_XET
    /-show/showfile/cue/036/bit  I32 F_XET
    /-show/showfile/cue/036/miditype  I32 F_XET
    /-show/showfile/cue/036/midichan  I32 F_XET
    /-show/showfile/cue/036/midipara1  I32 F_XET
    /-show/showfile/cue/036/midipara2  I32 F_XET
/-show/showfile/cue/037  <SCUE> n=0
    /-show/showfile/cue/037/numb  I32 F_XET
    /-show/showfile/cue/037/name  S32 F_XET
    /-show/showfile/cue/037/skip  I32 F_XET
    /-show/showfile/cue/037/scene  I32 F_XET
    /-show/showfile/cue/037/bit  I32 F_XET
    /-show/showfile/cue/037/miditype  I32 F_XET
    /-show/showfile/cue/037/midichan  I32 F_XET
    /-show/showfile/cue/037/midipara1  I32 F_XET
    /-show/showfile/cue/037/midipara2  I32 F_XET
/-show/showfile/cue/038  <SCUE> n=0
    /-show/showfile/cue/038/numb  I32 F_XET
    /-show/showfile/cue/038/name  S32 F_XET
    /-show/showfile/cue/038/skip  I32 F_XET
    /-show/showfile/cue/038/scene  I32 F_XET
    /-show/showfile/cue/038/bit  I32 F_XET
    /-show/showfile/cue/038/miditype  I32 F_XET
    /-show/showfile/cue/038/midichan  I32 F_XET
    /-show/showfile/cue/038/midipara1  I32 F_XET
    /-show/showfile/cue/038/midipara2  I32 F_XET
/-show/showfile/cue/039  <SCUE> n=0
    /-show/showfile/cue/039/numb  I32 F_XET
    /-show/showfile/cue/039/name  S32 F_XET
    /-show/showfile/cue/039/skip  I32 F_XET
    /-show/showfile/cue/039/scene  I32 F_XET
    /-show/showfile/cue/039/bit  I32 F_XET
    /-show/showfile/cue/039/miditype  I32 F_XET
    /-show/showfile/cue/039/midichan  I32 F_XET
    /-show/showfile/cue/039/midipara1  I32 F_XET
    /-show/showfile/cue/039/midipara2  I32 F_XET
/-show/showfile/cue/040  <SCUE> n=0
    /-show/showfile/cue/040/numb  I32 F_XET
    /-show/showfile/cue/040/name  S32 F_XET
    /-show/showfile/cue/040/skip  I32 F_XET
    /-show/showfile/cue/040/scene  I32 F_XET
    /-show/showfile/cue/040/bit  I32 F_XET
    /-show/showfile/cue/040/miditype  I32 F_XET
    /-show/showfile/cue/040/midichan  I32 F_XET
    /-show/showfile/cue/040/midipara1  I32 F_XET
    /-show/showfile/cue/040/midipara2  I32 F_XET
/-show/showfile/cue/041  <SCUE> n=0
    /-show/showfile/cue/041/numb  I32 F_XET
    /-show/showfile/cue/041/name  S32 F_XET
    /-show/showfile/cue/041/skip  I32 F_XET
    /-show/showfile/cue/041/scene  I32 F_XET
    /-show/showfile/cue/041/bit  I32 F_XET
    /-show/showfile/cue/041/miditype  I32 F_XET
    /-show/showfile/cue/041/midichan  I32 F_XET
    /-show/showfile/cue/041/midipara1  I32 F_XET
    /-show/showfile/cue/041/midipara2  I32 F_XET
/-show/showfile/cue/042  <SCUE> n=0
    /-show/showfile/cue/042/numb  I32 F_XET
    /-show/showfile/cue/042/name  S32 F_XET
    /-show/showfile/cue/042/skip  I32 F_XET
    /-show/showfile/cue/042/scene  I32 F_XET
    /-show/showfile/cue/042/bit  I32 F_XET
    /-show/showfile/cue/042/miditype  I32 F_XET
    /-show/showfile/cue/042/midichan  I32 F_XET
    /-show/showfile/cue/042/midipara1  I32 F_XET
    /-show/showfile/cue/042/midipara2  I32 F_XET
/-show/showfile/cue/043  <SCUE> n=0
    /-show/showfile/cue/043/numb  I32 F_XET
    /-show/showfile/cue/043/name  S32 F_XET
    /-show/showfile/cue/043/skip  I32 F_XET
    /-show/showfile/cue/043/scene  I32 F_XET
    /-show/showfile/cue/043/bit  I32 F_XET
    /-show/showfile/cue/043/miditype  I32 F_XET
    /-show/showfile/cue/043/midichan  I32 F_XET
    /-show/showfile/cue/043/midipara1  I32 F_XET
    /-show/showfile/cue/043/midipara2  I32 F_XET
/-show/showfile/cue/044  <SCUE> n=0
    /-show/showfile/cue/044/numb  I32 F_XET
    /-show/showfile/cue/044/name  S32 F_XET
    /-show/showfile/cue/044/skip  I32 F_XET
    /-show/showfile/cue/044/scene  I32 F_XET
    /-show/showfile/cue/044/bit  I32 F_XET
    /-show/showfile/cue/044/miditype  I32 F_XET
    /-show/showfile/cue/044/midichan  I32 F_XET
    /-show/showfile/cue/044/midipara1  I32 F_XET
    /-show/showfile/cue/044/midipara2  I32 F_XET
/-show/showfile/cue/045  <SCUE> n=0
    /-show/showfile/cue/045/numb  I32 F_XET
    /-show/showfile/cue/045/name  S32 F_XET
    /-show/showfile/cue/045/skip  I32 F_XET
    /-show/showfile/cue/045/scene  I32 F_XET
    /-show/showfile/cue/045/bit  I32 F_XET
    /-show/showfile/cue/045/miditype  I32 F_XET
    /-show/showfile/cue/045/midichan  I32 F_XET
    /-show/showfile/cue/045/midipara1  I32 F_XET
    /-show/showfile/cue/045/midipara2  I32 F_XET
/-show/showfile/cue/046  <SCUE> n=0
    /-show/showfile/cue/046/numb  I32 F_XET
    /-show/showfile/cue/046/name  S32 F_XET
    /-show/showfile/cue/046/skip  I32 F_XET
    /-show/showfile/cue/046/scene  I32 F_XET
    /-show/showfile/cue/046/bit  I32 F_XET
    /-show/showfile/cue/046/miditype  I32 F_XET
    /-show/showfile/cue/046/midichan  I32 F_XET
    /-show/showfile/cue/046/midipara1  I32 F_XET
    /-show/showfile/cue/046/midipara2  I32 F_XET
/-show/showfile/cue/047  <SCUE> n=0
    /-show/showfile/cue/047/numb  I32 F_XET
    /-show/showfile/cue/047/name  S32 F_XET
    /-show/showfile/cue/047/skip  I32 F_XET
    /-show/showfile/cue/047/scene  I32 F_XET
    /-show/showfile/cue/047/bit  I32 F_XET
    /-show/showfile/cue/047/miditype  I32 F_XET
    /-show/showfile/cue/047/midichan  I32 F_XET
    /-show/showfile/cue/047/midipara1  I32 F_XET
    /-show/showfile/cue/047/midipara2  I32 F_XET
/-show/showfile/cue/048  <SCUE> n=0
    /-show/showfile/cue/048/numb  I32 F_XET
    /-show/showfile/cue/048/name  S32 F_XET
    /-show/showfile/cue/048/skip  I32 F_XET
    /-show/showfile/cue/048/scene  I32 F_XET
    /-show/showfile/cue/048/bit  I32 F_XET
    /-show/showfile/cue/048/miditype  I32 F_XET
    /-show/showfile/cue/048/midichan  I32 F_XET
    /-show/showfile/cue/048/midipara1  I32 F_XET
    /-show/showfile/cue/048/midipara2  I32 F_XET
/-show/showfile/cue/049  <SCUE> n=0
    /-show/showfile/cue/049/numb  I32 F_XET
    /-show/showfile/cue/049/name  S32 F_XET
    /-show/showfile/cue/049/skip  I32 F_XET
    /-show/showfile/cue/049/scene  I32 F_XET
    /-show/showfile/cue/049/bit  I32 F_XET
    /-show/showfile/cue/049/miditype  I32 F_XET
    /-show/showfile/cue/049/midichan  I32 F_XET
    /-show/showfile/cue/049/midipara1  I32 F_XET
    /-show/showfile/cue/049/midipara2  I32 F_XET
/-show/showfile/cue/050  <SCUE> n=0
    /-show/showfile/cue/050/numb  I32 F_XET
    /-show/showfile/cue/050/name  S32 F_XET
    /-show/showfile/cue/050/skip  I32 F_XET
    /-show/showfile/cue/050/scene  I32 F_XET
    /-show/showfile/cue/050/bit  I32 F_XET
    /-show/showfile/cue/050/miditype  I32 F_XET
    /-show/showfile/cue/050/midichan  I32 F_XET
    /-show/showfile/cue/050/midipara1  I32 F_XET
    /-show/showfile/cue/050/midipara2  I32 F_XET
/-show/showfile/cue/051  <SCUE> n=0
    /-show/showfile/cue/051/numb  I32 F_XET
    /-show/showfile/cue/051/name  S32 F_XET
    /-show/showfile/cue/051/skip  I32 F_XET
    /-show/showfile/cue/051/scene  I32 F_XET
    /-show/showfile/cue/051/bit  I32 F_XET
    /-show/showfile/cue/051/miditype  I32 F_XET
    /-show/showfile/cue/051/midichan  I32 F_XET
    /-show/showfile/cue/051/midipara1  I32 F_XET
    /-show/showfile/cue/051/midipara2  I32 F_XET
/-show/showfile/cue/052  <SCUE> n=0
    /-show/showfile/cue/052/numb  I32 F_XET
    /-show/showfile/cue/052/name  S32 F_XET
    /-show/showfile/cue/052/skip  I32 F_XET
    /-show/showfile/cue/052/scene  I32 F_XET
    /-show/showfile/cue/052/bit  I32 F_XET
    /-show/showfile/cue/052/miditype  I32 F_XET
    /-show/showfile/cue/052/midichan  I32 F_XET
    /-show/showfile/cue/052/midipara1  I32 F_XET
    /-show/showfile/cue/052/midipara2  I32 F_XET
/-show/showfile/cue/053  <SCUE> n=0
    /-show/showfile/cue/053/numb  I32 F_XET
    /-show/showfile/cue/053/name  S32 F_XET
    /-show/showfile/cue/053/skip  I32 F_XET
    /-show/showfile/cue/053/scene  I32 F_XET
    /-show/showfile/cue/053/bit  I32 F_XET
    /-show/showfile/cue/053/miditype  I32 F_XET
    /-show/showfile/cue/053/midichan  I32 F_XET
    /-show/showfile/cue/053/midipara1  I32 F_XET
    /-show/showfile/cue/053/midipara2  I32 F_XET
/-show/showfile/cue/054  <SCUE> n=0
    /-show/showfile/cue/054/numb  I32 F_XET
    /-show/showfile/cue/054/name  S32 F_XET
    /-show/showfile/cue/054/skip  I32 F_XET
    /-show/showfile/cue/054/scene  I32 F_XET
    /-show/showfile/cue/054/bit  I32 F_XET
    /-show/showfile/cue/054/miditype  I32 F_XET
    /-show/showfile/cue/054/midichan  I32 F_XET
    /-show/showfile/cue/054/midipara1  I32 F_XET
    /-show/showfile/cue/054/midipara2  I32 F_XET
/-show/showfile/cue/055  <SCUE> n=0
    /-show/showfile/cue/055/numb  I32 F_XET
    /-show/showfile/cue/055/name  S32 F_XET
    /-show/showfile/cue/055/skip  I32 F_XET
    /-show/showfile/cue/055/scene  I32 F_XET
    /-show/showfile/cue/055/bit  I32 F_XET
    /-show/showfile/cue/055/miditype  I32 F_XET
    /-show/showfile/cue/055/midichan  I32 F_XET
    /-show/showfile/cue/055/midipara1  I32 F_XET
    /-show/showfile/cue/055/midipara2  I32 F_XET
/-show/showfile/cue/056  <SCUE> n=0
    /-show/showfile/cue/056/numb  I32 F_XET
    /-show/showfile/cue/056/name  S32 F_XET
    /-show/showfile/cue/056/skip  I32 F_XET
    /-show/showfile/cue/056/scene  I32 F_XET
    /-show/showfile/cue/056/bit  I32 F_XET
    /-show/showfile/cue/056/miditype  I32 F_XET
    /-show/showfile/cue/056/midichan  I32 F_XET
    /-show/showfile/cue/056/midipara1  I32 F_XET
    /-show/showfile/cue/056/midipara2  I32 F_XET
/-show/showfile/cue/057  <SCUE> n=0
    /-show/showfile/cue/057/numb  I32 F_XET
    /-show/showfile/cue/057/name  S32 F_XET
    /-show/showfile/cue/057/skip  I32 F_XET
    /-show/showfile/cue/057/scene  I32 F_XET
    /-show/showfile/cue/057/bit  I32 F_XET
    /-show/showfile/cue/057/miditype  I32 F_XET
    /-show/showfile/cue/057/midichan  I32 F_XET
    /-show/showfile/cue/057/midipara1  I32 F_XET
    /-show/showfile/cue/057/midipara2  I32 F_XET
/-show/showfile/cue/058  <SCUE> n=0
    /-show/showfile/cue/058/numb  I32 F_XET
    /-show/showfile/cue/058/name  S32 F_XET
    /-show/showfile/cue/058/skip  I32 F_XET
    /-show/showfile/cue/058/scene  I32 F_XET
    /-show/showfile/cue/058/bit  I32 F_XET
    /-show/showfile/cue/058/miditype  I32 F_XET
    /-show/showfile/cue/058/midichan  I32 F_XET
    /-show/showfile/cue/058/midipara1  I32 F_XET
    /-show/showfile/cue/058/midipara2  I32 F_XET
/-show/showfile/cue/059  <SCUE> n=0
    /-show/showfile/cue/059/numb  I32 F_XET
    /-show/showfile/cue/059/name  S32 F_XET
    /-show/showfile/cue/059/skip  I32 F_XET
    /-show/showfile/cue/059/scene  I32 F_XET
    /-show/showfile/cue/059/bit  I32 F_XET
    /-show/showfile/cue/059/miditype  I32 F_XET
    /-show/showfile/cue/059/midichan  I32 F_XET
    /-show/showfile/cue/059/midipara1  I32 F_XET
    /-show/showfile/cue/059/midipara2  I32 F_XET
/-show/showfile/cue/060  <SCUE> n=0
    /-show/showfile/cue/060/numb  I32 F_XET
    /-show/showfile/cue/060/name  S32 F_XET
    /-show/showfile/cue/060/skip  I32 F_XET
    /-show/showfile/cue/060/scene  I32 F_XET
    /-show/showfile/cue/060/bit  I32 F_XET
    /-show/showfile/cue/060/miditype  I32 F_XET
    /-show/showfile/cue/060/midichan  I32 F_XET
    /-show/showfile/cue/060/midipara1  I32 F_XET
    /-show/showfile/cue/060/midipara2  I32 F_XET
/-show/showfile/cue/061  <SCUE> n=0
    /-show/showfile/cue/061/numb  I32 F_XET
    /-show/showfile/cue/061/name  S32 F_XET
    /-show/showfile/cue/061/skip  I32 F_XET
    /-show/showfile/cue/061/scene  I32 F_XET
    /-show/showfile/cue/061/bit  I32 F_XET
    /-show/showfile/cue/061/miditype  I32 F_XET
    /-show/showfile/cue/061/midichan  I32 F_XET
    /-show/showfile/cue/061/midipara1  I32 F_XET
    /-show/showfile/cue/061/midipara2  I32 F_XET
/-show/showfile/cue/062  <SCUE> n=0
    /-show/showfile/cue/062/numb  I32 F_XET
    /-show/showfile/cue/062/name  S32 F_XET
    /-show/showfile/cue/062/skip  I32 F_XET
    /-show/showfile/cue/062/scene  I32 F_XET
    /-show/showfile/cue/062/bit  I32 F_XET
    /-show/showfile/cue/062/miditype  I32 F_XET
    /-show/showfile/cue/062/midichan  I32 F_XET
    /-show/showfile/cue/062/midipara1  I32 F_XET
    /-show/showfile/cue/062/midipara2  I32 F_XET
/-show/showfile/cue/063  <SCUE> n=0
    /-show/showfile/cue/063/numb  I32 F_XET
    /-show/showfile/cue/063/name  S32 F_XET
    /-show/showfile/cue/063/skip  I32 F_XET
    /-show/showfile/cue/063/scene  I32 F_XET
    /-show/showfile/cue/063/bit  I32 F_XET
    /-show/showfile/cue/063/miditype  I32 F_XET
    /-show/showfile/cue/063/midichan  I32 F_XET
    /-show/showfile/cue/063/midipara1  I32 F_XET
    /-show/showfile/cue/063/midipara2  I32 F_XET
/-show/showfile/cue/064  <SCUE> n=0
    /-show/showfile/cue/064/numb  I32 F_XET
    /-show/showfile/cue/064/name  S32 F_XET
    /-show/showfile/cue/064/skip  I32 F_XET
    /-show/showfile/cue/064/scene  I32 F_XET
    /-show/showfile/cue/064/bit  I32 F_XET
    /-show/showfile/cue/064/miditype  I32 F_XET
    /-show/showfile/cue/064/midichan  I32 F_XET
    /-show/showfile/cue/064/midipara1  I32 F_XET
    /-show/showfile/cue/064/midipara2  I32 F_XET
/-show/showfile/cue/065  <SCUE> n=0
    /-show/showfile/cue/065/numb  I32 F_XET
    /-show/showfile/cue/065/name  S32 F_XET
    /-show/showfile/cue/065/skip  I32 F_XET
    /-show/showfile/cue/065/scene  I32 F_XET
    /-show/showfile/cue/065/bit  I32 F_XET
    /-show/showfile/cue/065/miditype  I32 F_XET
    /-show/showfile/cue/065/midichan  I32 F_XET
    /-show/showfile/cue/065/midipara1  I32 F_XET
    /-show/showfile/cue/065/midipara2  I32 F_XET
/-show/showfile/cue/066  <SCUE> n=0
    /-show/showfile/cue/066/numb  I32 F_XET
    /-show/showfile/cue/066/name  S32 F_XET
    /-show/showfile/cue/066/skip  I32 F_XET
    /-show/showfile/cue/066/scene  I32 F_XET
    /-show/showfile/cue/066/bit  I32 F_XET
    /-show/showfile/cue/066/miditype  I32 F_XET
    /-show/showfile/cue/066/midichan  I32 F_XET
    /-show/showfile/cue/066/midipara1  I32 F_XET
    /-show/showfile/cue/066/midipara2  I32 F_XET
/-show/showfile/cue/067  <SCUE> n=0
    /-show/showfile/cue/067/numb  I32 F_XET
    /-show/showfile/cue/067/name  S32 F_XET
    /-show/showfile/cue/067/skip  I32 F_XET
    /-show/showfile/cue/067/scene  I32 F_XET
    /-show/showfile/cue/067/bit  I32 F_XET
    /-show/showfile/cue/067/miditype  I32 F_XET
    /-show/showfile/cue/067/midichan  I32 F_XET
    /-show/showfile/cue/067/midipara1  I32 F_XET
    /-show/showfile/cue/067/midipara2  I32 F_XET
/-show/showfile/cue/068  <SCUE> n=0
    /-show/showfile/cue/068/numb  I32 F_XET
    /-show/showfile/cue/068/name  S32 F_XET
    /-show/showfile/cue/068/skip  I32 F_XET
    /-show/showfile/cue/068/scene  I32 F_XET
    /-show/showfile/cue/068/bit  I32 F_XET
    /-show/showfile/cue/068/miditype  I32 F_XET
    /-show/showfile/cue/068/midichan  I32 F_XET
    /-show/showfile/cue/068/midipara1  I32 F_XET
    /-show/showfile/cue/068/midipara2  I32 F_XET
/-show/showfile/cue/069  <SCUE> n=0
    /-show/showfile/cue/069/numb  I32 F_XET
    /-show/showfile/cue/069/name  S32 F_XET
    /-show/showfile/cue/069/skip  I32 F_XET
    /-show/showfile/cue/069/scene  I32 F_XET
    /-show/showfile/cue/069/bit  I32 F_XET
    /-show/showfile/cue/069/miditype  I32 F_XET
    /-show/showfile/cue/069/midichan  I32 F_XET
    /-show/showfile/cue/069/midipara1  I32 F_XET
    /-show/showfile/cue/069/midipara2  I32 F_XET
/-show/showfile/cue/070  <SCUE> n=0
    /-show/showfile/cue/070/numb  I32 F_XET
    /-show/showfile/cue/070/name  S32 F_XET
    /-show/showfile/cue/070/skip  I32 F_XET
    /-show/showfile/cue/070/scene  I32 F_XET
    /-show/showfile/cue/070/bit  I32 F_XET
    /-show/showfile/cue/070/miditype  I32 F_XET
    /-show/showfile/cue/070/midichan  I32 F_XET
    /-show/showfile/cue/070/midipara1  I32 F_XET
    /-show/showfile/cue/070/midipara2  I32 F_XET
/-show/showfile/cue/071  <SCUE> n=0
    /-show/showfile/cue/071/numb  I32 F_XET
    /-show/showfile/cue/071/name  S32 F_XET
    /-show/showfile/cue/071/skip  I32 F_XET
    /-show/showfile/cue/071/scene  I32 F_XET
    /-show/showfile/cue/071/bit  I32 F_XET
    /-show/showfile/cue/071/miditype  I32 F_XET
    /-show/showfile/cue/071/midichan  I32 F_XET
    /-show/showfile/cue/071/midipara1  I32 F_XET
    /-show/showfile/cue/071/midipara2  I32 F_XET
/-show/showfile/cue/072  <SCUE> n=0
    /-show/showfile/cue/072/numb  I32 F_XET
    /-show/showfile/cue/072/name  S32 F_XET
    /-show/showfile/cue/072/skip  I32 F_XET
    /-show/showfile/cue/072/scene  I32 F_XET
    /-show/showfile/cue/072/bit  I32 F_XET
    /-show/showfile/cue/072/miditype  I32 F_XET
    /-show/showfile/cue/072/midichan  I32 F_XET
    /-show/showfile/cue/072/midipara1  I32 F_XET
    /-show/showfile/cue/072/midipara2  I32 F_XET
/-show/showfile/cue/073  <SCUE> n=0
    /-show/showfile/cue/073/numb  I32 F_XET
    /-show/showfile/cue/073/name  S32 F_XET
    /-show/showfile/cue/073/skip  I32 F_XET
    /-show/showfile/cue/073/scene  I32 F_XET
    /-show/showfile/cue/073/bit  I32 F_XET
    /-show/showfile/cue/073/miditype  I32 F_XET
    /-show/showfile/cue/073/midichan  I32 F_XET
    /-show/showfile/cue/073/midipara1  I32 F_XET
    /-show/showfile/cue/073/midipara2  I32 F_XET
/-show/showfile/cue/074  <SCUE> n=0
    /-show/showfile/cue/074/numb  I32 F_XET
    /-show/showfile/cue/074/name  S32 F_XET
    /-show/showfile/cue/074/skip  I32 F_XET
    /-show/showfile/cue/074/scene  I32 F_XET
    /-show/showfile/cue/074/bit  I32 F_XET
    /-show/showfile/cue/074/miditype  I32 F_XET
    /-show/showfile/cue/074/midichan  I32 F_XET
    /-show/showfile/cue/074/midipara1  I32 F_XET
    /-show/showfile/cue/074/midipara2  I32 F_XET
/-show/showfile/cue/075  <SCUE> n=0
    /-show/showfile/cue/075/numb  I32 F_XET
    /-show/showfile/cue/075/name  S32 F_XET
    /-show/showfile/cue/075/skip  I32 F_XET
    /-show/showfile/cue/075/scene  I32 F_XET
    /-show/showfile/cue/075/bit  I32 F_XET
    /-show/showfile/cue/075/miditype  I32 F_XET
    /-show/showfile/cue/075/midichan  I32 F_XET
    /-show/showfile/cue/075/midipara1  I32 F_XET
    /-show/showfile/cue/075/midipara2  I32 F_XET
/-show/showfile/cue/076  <SCUE> n=0
    /-show/showfile/cue/076/numb  I32 F_XET
    /-show/showfile/cue/076/name  S32 F_XET
    /-show/showfile/cue/076/skip  I32 F_XET
    /-show/showfile/cue/076/scene  I32 F_XET
    /-show/showfile/cue/076/bit  I32 F_XET
    /-show/showfile/cue/076/miditype  I32 F_XET
    /-show/showfile/cue/076/midichan  I32 F_XET
    /-show/showfile/cue/076/midipara1  I32 F_XET
    /-show/showfile/cue/076/midipara2  I32 F_XET
/-show/showfile/cue/077  <SCUE> n=0
    /-show/showfile/cue/077/numb  I32 F_XET
    /-show/showfile/cue/077/name  S32 F_XET
    /-show/showfile/cue/077/skip  I32 F_XET
    /-show/showfile/cue/077/scene  I32 F_XET
    /-show/showfile/cue/077/bit  I32 F_XET
    /-show/showfile/cue/077/miditype  I32 F_XET
    /-show/showfile/cue/077/midichan  I32 F_XET
    /-show/showfile/cue/077/midipara1  I32 F_XET
    /-show/showfile/cue/077/midipara2  I32 F_XET
/-show/showfile/cue/078  <SCUE> n=0
    /-show/showfile/cue/078/numb  I32 F_XET
    /-show/showfile/cue/078/name  S32 F_XET
    /-show/showfile/cue/078/skip  I32 F_XET
    /-show/showfile/cue/078/scene  I32 F_XET
    /-show/showfile/cue/078/bit  I32 F_XET
    /-show/showfile/cue/078/miditype  I32 F_XET
    /-show/showfile/cue/078/midichan  I32 F_XET
    /-show/showfile/cue/078/midipara1  I32 F_XET
    /-show/showfile/cue/078/midipara2  I32 F_XET
/-show/showfile/cue/079  <SCUE> n=0
    /-show/showfile/cue/079/numb  I32 F_XET
    /-show/showfile/cue/079/name  S32 F_XET
    /-show/showfile/cue/079/skip  I32 F_XET
    /-show/showfile/cue/079/scene  I32 F_XET
    /-show/showfile/cue/079/bit  I32 F_XET
    /-show/showfile/cue/079/miditype  I32 F_XET
    /-show/showfile/cue/079/midichan  I32 F_XET
    /-show/showfile/cue/079/midipara1  I32 F_XET
    /-show/showfile/cue/079/midipara2  I32 F_XET
/-show/showfile/cue/080  <SCUE> n=0
    /-show/showfile/cue/080/numb  I32 F_XET
    /-show/showfile/cue/080/name  S32 F_XET
    /-show/showfile/cue/080/skip  I32 F_XET
    /-show/showfile/cue/080/scene  I32 F_XET
    /-show/showfile/cue/080/bit  I32 F_XET
    /-show/showfile/cue/080/miditype  I32 F_XET
    /-show/showfile/cue/080/midichan  I32 F_XET
    /-show/showfile/cue/080/midipara1  I32 F_XET
    /-show/showfile/cue/080/midipara2  I32 F_XET
/-show/showfile/cue/081  <SCUE> n=0
    /-show/showfile/cue/081/numb  I32 F_XET
    /-show/showfile/cue/081/name  S32 F_XET
    /-show/showfile/cue/081/skip  I32 F_XET
    /-show/showfile/cue/081/scene  I32 F_XET
    /-show/showfile/cue/081/bit  I32 F_XET
    /-show/showfile/cue/081/miditype  I32 F_XET
    /-show/showfile/cue/081/midichan  I32 F_XET
    /-show/showfile/cue/081/midipara1  I32 F_XET
    /-show/showfile/cue/081/midipara2  I32 F_XET
/-show/showfile/cue/082  <SCUE> n=0
    /-show/showfile/cue/082/numb  I32 F_XET
    /-show/showfile/cue/082/name  S32 F_XET
    /-show/showfile/cue/082/skip  I32 F_XET
    /-show/showfile/cue/082/scene  I32 F_XET
    /-show/showfile/cue/082/bit  I32 F_XET
    /-show/showfile/cue/082/miditype  I32 F_XET
    /-show/showfile/cue/082/midichan  I32 F_XET
    /-show/showfile/cue/082/midipara1  I32 F_XET
    /-show/showfile/cue/082/midipara2  I32 F_XET
/-show/showfile/cue/083  <SCUE> n=0
    /-show/showfile/cue/083/numb  I32 F_XET
    /-show/showfile/cue/083/name  S32 F_XET
    /-show/showfile/cue/083/skip  I32 F_XET
    /-show/showfile/cue/083/scene  I32 F_XET
    /-show/showfile/cue/083/bit  I32 F_XET
    /-show/showfile/cue/083/miditype  I32 F_XET
    /-show/showfile/cue/083/midichan  I32 F_XET
    /-show/showfile/cue/083/midipara1  I32 F_XET
    /-show/showfile/cue/083/midipara2  I32 F_XET
/-show/showfile/cue/084  <SCUE> n=0
    /-show/showfile/cue/084/numb  I32 F_XET
    /-show/showfile/cue/084/name  S32 F_XET
    /-show/showfile/cue/084/skip  I32 F_XET
    /-show/showfile/cue/084/scene  I32 F_XET
    /-show/showfile/cue/084/bit  I32 F_XET
    /-show/showfile/cue/084/miditype  I32 F_XET
    /-show/showfile/cue/084/midichan  I32 F_XET
    /-show/showfile/cue/084/midipara1  I32 F_XET
    /-show/showfile/cue/084/midipara2  I32 F_XET
/-show/showfile/cue/085  <SCUE> n=0
    /-show/showfile/cue/085/numb  I32 F_XET
    /-show/showfile/cue/085/name  S32 F_XET
    /-show/showfile/cue/085/skip  I32 F_XET
    /-show/showfile/cue/085/scene  I32 F_XET
    /-show/showfile/cue/085/bit  I32 F_XET
    /-show/showfile/cue/085/miditype  I32 F_XET
    /-show/showfile/cue/085/midichan  I32 F_XET
    /-show/showfile/cue/085/midipara1  I32 F_XET
    /-show/showfile/cue/085/midipara2  I32 F_XET
/-show/showfile/cue/086  <SCUE> n=0
    /-show/showfile/cue/086/numb  I32 F_XET
    /-show/showfile/cue/086/name  S32 F_XET
    /-show/showfile/cue/086/skip  I32 F_XET
    /-show/showfile/cue/086/scene  I32 F_XET
    /-show/showfile/cue/086/bit  I32 F_XET
    /-show/showfile/cue/086/miditype  I32 F_XET
    /-show/showfile/cue/086/midichan  I32 F_XET
    /-show/showfile/cue/086/midipara1  I32 F_XET
    /-show/showfile/cue/086/midipara2  I32 F_XET
/-show/showfile/cue/087  <SCUE> n=0
    /-show/showfile/cue/087/numb  I32 F_XET
    /-show/showfile/cue/087/name  S32 F_XET
    /-show/showfile/cue/087/skip  I32 F_XET
    /-show/showfile/cue/087/scene  I32 F_XET
    /-show/showfile/cue/087/bit  I32 F_XET
    /-show/showfile/cue/087/miditype  I32 F_XET
    /-show/showfile/cue/087/midichan  I32 F_XET
    /-show/showfile/cue/087/midipara1  I32 F_XET
    /-show/showfile/cue/087/midipara2  I32 F_XET
/-show/showfile/cue/088  <SCUE> n=0
    /-show/showfile/cue/088/numb  I32 F_XET
    /-show/showfile/cue/088/name  S32 F_XET
    /-show/showfile/cue/088/skip  I32 F_XET
    /-show/showfile/cue/088/scene  I32 F_XET
    /-show/showfile/cue/088/bit  I32 F_XET
    /-show/showfile/cue/088/miditype  I32 F_XET
    /-show/showfile/cue/088/midichan  I32 F_XET
    /-show/showfile/cue/088/midipara1  I32 F_XET
    /-show/showfile/cue/088/midipara2  I32 F_XET
/-show/showfile/cue/089  <SCUE> n=0
    /-show/showfile/cue/089/numb  I32 F_XET
    /-show/showfile/cue/089/name  S32 F_XET
    /-show/showfile/cue/089/skip  I32 F_XET
    /-show/showfile/cue/089/scene  I32 F_XET
    /-show/showfile/cue/089/bit  I32 F_XET
    /-show/showfile/cue/089/miditype  I32 F_XET
    /-show/showfile/cue/089/midichan  I32 F_XET
    /-show/showfile/cue/089/midipara1  I32 F_XET
    /-show/showfile/cue/089/midipara2  I32 F_XET
/-show/showfile/cue/090  <SCUE> n=0
    /-show/showfile/cue/090/numb  I32 F_XET
    /-show/showfile/cue/090/name  S32 F_XET
    /-show/showfile/cue/090/skip  I32 F_XET
    /-show/showfile/cue/090/scene  I32 F_XET
    /-show/showfile/cue/090/bit  I32 F_XET
    /-show/showfile/cue/090/miditype  I32 F_XET
    /-show/showfile/cue/090/midichan  I32 F_XET
    /-show/showfile/cue/090/midipara1  I32 F_XET
    /-show/showfile/cue/090/midipara2  I32 F_XET
/-show/showfile/cue/091  <SCUE> n=0
    /-show/showfile/cue/091/numb  I32 F_XET
    /-show/showfile/cue/091/name  S32 F_XET
    /-show/showfile/cue/091/skip  I32 F_XET
    /-show/showfile/cue/091/scene  I32 F_XET
    /-show/showfile/cue/091/bit  I32 F_XET
    /-show/showfile/cue/091/miditype  I32 F_XET
    /-show/showfile/cue/091/midichan  I32 F_XET
    /-show/showfile/cue/091/midipara1  I32 F_XET
    /-show/showfile/cue/091/midipara2  I32 F_XET
/-show/showfile/cue/092  <SCUE> n=0
    /-show/showfile/cue/092/numb  I32 F_XET
    /-show/showfile/cue/092/name  S32 F_XET
    /-show/showfile/cue/092/skip  I32 F_XET
    /-show/showfile/cue/092/scene  I32 F_XET
    /-show/showfile/cue/092/bit  I32 F_XET
    /-show/showfile/cue/092/miditype  I32 F_XET
    /-show/showfile/cue/092/midichan  I32 F_XET
    /-show/showfile/cue/092/midipara1  I32 F_XET
    /-show/showfile/cue/092/midipara2  I32 F_XET
/-show/showfile/cue/093  <SCUE> n=0
    /-show/showfile/cue/093/numb  I32 F_XET
    /-show/showfile/cue/093/name  S32 F_XET
    /-show/showfile/cue/093/skip  I32 F_XET
    /-show/showfile/cue/093/scene  I32 F_XET
    /-show/showfile/cue/093/bit  I32 F_XET
    /-show/showfile/cue/093/miditype  I32 F_XET
    /-show/showfile/cue/093/midichan  I32 F_XET
    /-show/showfile/cue/093/midipara1  I32 F_XET
    /-show/showfile/cue/093/midipara2  I32 F_XET
/-show/showfile/cue/094  <SCUE> n=0
    /-show/showfile/cue/094/numb  I32 F_XET
    /-show/showfile/cue/094/name  S32 F_XET
    /-show/showfile/cue/094/skip  I32 F_XET
    /-show/showfile/cue/094/scene  I32 F_XET
    /-show/showfile/cue/094/bit  I32 F_XET
    /-show/showfile/cue/094/miditype  I32 F_XET
    /-show/showfile/cue/094/midichan  I32 F_XET
    /-show/showfile/cue/094/midipara1  I32 F_XET
    /-show/showfile/cue/094/midipara2  I32 F_XET
/-show/showfile/cue/095  <SCUE> n=0
    /-show/showfile/cue/095/numb  I32 F_XET
    /-show/showfile/cue/095/name  S32 F_XET
    /-show/showfile/cue/095/skip  I32 F_XET
    /-show/showfile/cue/095/scene  I32 F_XET
    /-show/showfile/cue/095/bit  I32 F_XET
    /-show/showfile/cue/095/miditype  I32 F_XET
    /-show/showfile/cue/095/midichan  I32 F_XET
    /-show/showfile/cue/095/midipara1  I32 F_XET
    /-show/showfile/cue/095/midipara2  I32 F_XET
/-show/showfile/cue/096  <SCUE> n=0
    /-show/showfile/cue/096/numb  I32 F_XET
    /-show/showfile/cue/096/name  S32 F_XET
    /-show/showfile/cue/096/skip  I32 F_XET
    /-show/showfile/cue/096/scene  I32 F_XET
    /-show/showfile/cue/096/bit  I32 F_XET
    /-show/showfile/cue/096/miditype  I32 F_XET
    /-show/showfile/cue/096/midichan  I32 F_XET
    /-show/showfile/cue/096/midipara1  I32 F_XET
    /-show/showfile/cue/096/midipara2  I32 F_XET
/-show/showfile/cue/097  <SCUE> n=0
    /-show/showfile/cue/097/numb  I32 F_XET
    /-show/showfile/cue/097/name  S32 F_XET
    /-show/showfile/cue/097/skip  I32 F_XET
    /-show/showfile/cue/097/scene  I32 F_XET
    /-show/showfile/cue/097/bit  I32 F_XET
    /-show/showfile/cue/097/miditype  I32 F_XET
    /-show/showfile/cue/097/midichan  I32 F_XET
    /-show/showfile/cue/097/midipara1  I32 F_XET
    /-show/showfile/cue/097/midipara2  I32 F_XET
/-show/showfile/cue/098  <SCUE> n=0
    /-show/showfile/cue/098/numb  I32 F_XET
    /-show/showfile/cue/098/name  S32 F_XET
    /-show/showfile/cue/098/skip  I32 F_XET
    /-show/showfile/cue/098/scene  I32 F_XET
    /-show/showfile/cue/098/bit  I32 F_XET
    /-show/showfile/cue/098/miditype  I32 F_XET
    /-show/showfile/cue/098/midichan  I32 F_XET
    /-show/showfile/cue/098/midipara1  I32 F_XET
    /-show/showfile/cue/098/midipara2  I32 F_XET
/-show/showfile/cue/099  <SCUE> n=0
    /-show/showfile/cue/099/numb  I32 F_XET
    /-show/showfile/cue/099/name  S32 F_XET
    /-show/showfile/cue/099/skip  I32 F_XET
    /-show/showfile/cue/099/scene  I32 F_XET
    /-show/showfile/cue/099/bit  I32 F_XET
    /-show/showfile/cue/099/miditype  I32 F_XET
    /-show/showfile/cue/099/midichan  I32 F_XET
    /-show/showfile/cue/099/midipara1  I32 F_XET
    /-show/showfile/cue/099/midipara2  I32 F_XET
/-show/showfile/cue/000  <SCUE> n=0
    /-show/showfile/cue/000/numb  I32 F_XET
    /-show/showfile/cue/000/name  S32 F_XET
    /-show/showfile/cue/000/skip  I32 F_XET
    /-show/showfile/cue/000/scene  I32 F_XET
    /-show/showfile/cue/000/bit  I32 F_XET
    /-show/showfile/cue/000/miditype  I32 F_XET
    /-show/showfile/cue/000/midichan  I32 F_XET
    /-show/showfile/cue/000/midipara1  I32 F_XET
    /-show/showfile/cue/000/midipara2  I32 F_XET
/-show/showfile/cue/001  <SCUE> n=0
    /-show/showfile/cue/001/numb  I32 F_XET
    /-show/showfile/cue/001/name  S32 F_XET
    /-show/showfile/cue/001/skip  I32 F_XET
    /-show/showfile/cue/001/scene  I32 F_XET
    /-show/showfile/cue/001/bit  I32 F_XET
    /-show/showfile/cue/001/miditype  I32 F_XET
    /-show/showfile/cue/001/midichan  I32 F_XET
    /-show/showfile/cue/001/midipara1  I32 F_XET
    /-show/showfile/cue/001/midipara2  I32 F_XET
/-show/showfile/cue/002  <SCUE> n=0
    /-show/showfile/cue/002/numb  I32 F_XET
    /-show/showfile/cue/002/name  S32 F_XET
    /-show/showfile/cue/002/skip  I32 F_XET
    /-show/showfile/cue/002/scene  I32 F_XET
    /-show/showfile/cue/002/bit  I32 F_XET
    /-show/showfile/cue/002/miditype  I32 F_XET
    /-show/showfile/cue/002/midichan  I32 F_XET
    /-show/showfile/cue/002/midipara1  I32 F_XET
    /-show/showfile/cue/002/midipara2  I32 F_XET
/-show/showfile/cue/003  <SCUE> n=0
    /-show/showfile/cue/003/numb  I32 F_XET
    /-show/showfile/cue/003/name  S32 F_XET
    /-show/showfile/cue/003/skip  I32 F_XET
    /-show/showfile/cue/003/scene  I32 F_XET
    /-show/showfile/cue/003/bit  I32 F_XET
    /-show/showfile/cue/003/miditype  I32 F_XET
    /-show/showfile/cue/003/midichan  I32 F_XET
    /-show/showfile/cue/003/midipara1  I32 F_XET
    /-show/showfile/cue/003/midipara2  I32 F_XET
/-show/showfile/cue/004  <SCUE> n=0
    /-show/showfile/cue/004/numb  I32 F_XET
    /-show/showfile/cue/004/name  S32 F_XET
    /-show/showfile/cue/004/skip  I32 F_XET
    /-show/showfile/cue/004/scene  I32 F_XET
    /-show/showfile/cue/004/bit  I32 F_XET
    /-show/showfile/cue/004/miditype  I32 F_XET
    /-show/showfile/cue/004/midichan  I32 F_XET
    /-show/showfile/cue/004/midipara1  I32 F_XET
    /-show/showfile/cue/004/midipara2  I32 F_XET
/-show/showfile/cue/005  <SCUE> n=0
    /-show/showfile/cue/005/numb  I32 F_XET
    /-show/showfile/cue/005/name  S32 F_XET
    /-show/showfile/cue/005/skip  I32 F_XET
    /-show/showfile/cue/005/scene  I32 F_XET
    /-show/showfile/cue/005/bit  I32 F_XET
    /-show/showfile/cue/005/miditype  I32 F_XET
    /-show/showfile/cue/005/midichan  I32 F_XET
    /-show/showfile/cue/005/midipara1  I32 F_XET
    /-show/showfile/cue/005/midipara2  I32 F_XET
/-show/showfile/cue/006  <SCUE> n=0
    /-show/showfile/cue/006/numb  I32 F_XET
    /-show/showfile/cue/006/name  S32 F_XET
    /-show/showfile/cue/006/skip  I32 F_XET
    /-show/showfile/cue/006/scene  I32 F_XET
    /-show/showfile/cue/006/bit  I32 F_XET
    /-show/showfile/cue/006/miditype  I32 F_XET
    /-show/showfile/cue/006/midichan  I32 F_XET
    /-show/showfile/cue/006/midipara1  I32 F_XET
    /-show/showfile/cue/006/midipara2  I32 F_XET
/-show/showfile/cue/007  <SCUE> n=0
    /-show/showfile/cue/007/numb  I32 F_XET
    /-show/showfile/cue/007/name  S32 F_XET
    /-show/showfile/cue/007/skip  I32 F_XET
    /-show/showfile/cue/007/scene  I32 F_XET
    /-show/showfile/cue/007/bit  I32 F_XET
    /-show/showfile/cue/007/miditype  I32 F_XET
    /-show/showfile/cue/007/midichan  I32 F_XET
    /-show/showfile/cue/007/midipara1  I32 F_XET
    /-show/showfile/cue/007/midipara2  I32 F_XET
/-show/showfile/cue/008  <SCUE> n=0
    /-show/showfile/cue/008/numb  I32 F_XET
    /-show/showfile/cue/008/name  S32 F_XET
    /-show/showfile/cue/008/skip  I32 F_XET
    /-show/showfile/cue/008/scene  I32 F_XET
    /-show/showfile/cue/008/bit  I32 F_XET
    /-show/showfile/cue/008/miditype  I32 F_XET
    /-show/showfile/cue/008/midichan  I32 F_XET
    /-show/showfile/cue/008/midipara1  I32 F_XET
    /-show/showfile/cue/008/midipara2  I32 F_XET
/-show/showfile/cue/009  <SCUE> n=0
    /-show/showfile/cue/009/numb  I32 F_XET
    /-show/showfile/cue/009/name  S32 F_XET
    /-show/showfile/cue/009/skip  I32 F_XET
    /-show/showfile/cue/009/scene  I32 F_XET
    /-show/showfile/cue/009/bit  I32 F_XET
    /-show/showfile/cue/009/miditype  I32 F_XET
    /-show/showfile/cue/009/midichan  I32 F_XET
    /-show/showfile/cue/009/midipara1  I32 F_XET
    /-show/showfile/cue/009/midipara2  I32 F_XET
/-show/showfile/cue/010  <SCUE> n=0
    /-show/showfile/cue/010/numb  I32 F_XET
    /-show/showfile/cue/010/name  S32 F_XET
    /-show/showfile/cue/010/skip  I32 F_XET
    /-show/showfile/cue/010/scene  I32 F_XET
    /-show/showfile/cue/010/bit  I32 F_XET
    /-show/showfile/cue/010/miditype  I32 F_XET
    /-show/showfile/cue/010/midichan  I32 F_XET
    /-show/showfile/cue/010/midipara1  I32 F_XET
    /-show/showfile/cue/010/midipara2  I32 F_XET
/-show/showfile/cue/011  <SCUE> n=0
    /-show/showfile/cue/011/numb  I32 F_XET
    /-show/showfile/cue/011/name  S32 F_XET
    /-show/showfile/cue/011/skip  I32 F_XET
    /-show/showfile/cue/011/scene  I32 F_XET
    /-show/showfile/cue/011/bit  I32 F_XET
    /-show/showfile/cue/011/miditype  I32 F_XET
    /-show/showfile/cue/011/midichan  I32 F_XET
    /-show/showfile/cue/011/midipara1  I32 F_XET
    /-show/showfile/cue/011/midipara2  I32 F_XET
/-show/showfile/cue/012  <SCUE> n=0
    /-show/showfile/cue/012/numb  I32 F_XET
    /-show/showfile/cue/012/name  S32 F_XET
    /-show/showfile/cue/012/skip  I32 F_XET
    /-show/showfile/cue/012/scene  I32 F_XET
    /-show/showfile/cue/012/bit  I32 F_XET
    /-show/showfile/cue/012/miditype  I32 F_XET
    /-show/showfile/cue/012/midichan  I32 F_XET
    /-show/showfile/cue/012/midipara1  I32 F_XET
    /-show/showfile/cue/012/midipara2  I32 F_XET
/-show/showfile/cue/013  <SCUE> n=0
    /-show/showfile/cue/013/numb  I32 F_XET
    /-show/showfile/cue/013/name  S32 F_XET
    /-show/showfile/cue/013/skip  I32 F_XET
    /-show/showfile/cue/013/scene  I32 F_XET
    /-show/showfile/cue/013/bit  I32 F_XET
    /-show/showfile/cue/013/miditype  I32 F_XET
    /-show/showfile/cue/013/midichan  I32 F_XET
    /-show/showfile/cue/013/midipara1  I32 F_XET
    /-show/showfile/cue/013/midipara2  I32 F_XET
/-show/showfile/cue/014  <SCUE> n=0
    /-show/showfile/cue/014/numb  I32 F_XET
    /-show/showfile/cue/014/name  S32 F_XET
    /-show/showfile/cue/014/skip  I32 F_XET
    /-show/showfile/cue/014/scene  I32 F_XET
    /-show/showfile/cue/014/bit  I32 F_XET
    /-show/showfile/cue/014/miditype  I32 F_XET
    /-show/showfile/cue/014/midichan  I32 F_XET
    /-show/showfile/cue/014/midipara1  I32 F_XET
    /-show/showfile/cue/014/midipara2  I32 F_XET
/-show/showfile/cue/015  <SCUE> n=0
    /-show/showfile/cue/015/numb  I32 F_XET
    /-show/showfile/cue/015/name  S32 F_XET
    /-show/showfile/cue/015/skip  I32 F_XET
    /-show/showfile/cue/015/scene  I32 F_XET
    /-show/showfile/cue/015/bit  I32 F_XET
    /-show/showfile/cue/015/miditype  I32 F_XET
    /-show/showfile/cue/015/midichan  I32 F_XET
    /-show/showfile/cue/015/midipara1  I32 F_XET
    /-show/showfile/cue/015/midipara2  I32 F_XET
/-show/showfile/cue/016  <SCUE> n=0
    /-show/showfile/cue/016/numb  I32 F_XET
    /-show/showfile/cue/016/name  S32 F_XET
    /-show/showfile/cue/016/skip  I32 F_XET
    /-show/showfile/cue/016/scene  I32 F_XET
    /-show/showfile/cue/016/bit  I32 F_XET
    /-show/showfile/cue/016/miditype  I32 F_XET
    /-show/showfile/cue/016/midichan  I32 F_XET
    /-show/showfile/cue/016/midipara1  I32 F_XET
    /-show/showfile/cue/016/midipara2  I32 F_XET
/-show/showfile/cue/017  <SCUE> n=0
    /-show/showfile/cue/017/numb  I32 F_XET
    /-show/showfile/cue/017/name  S32 F_XET
    /-show/showfile/cue/017/skip  I32 F_XET
    /-show/showfile/cue/017/scene  I32 F_XET
    /-show/showfile/cue/017/bit  I32 F_XET
    /-show/showfile/cue/017/miditype  I32 F_XET
    /-show/showfile/cue/017/midichan  I32 F_XET
    /-show/showfile/cue/017/midipara1  I32 F_XET
    /-show/showfile/cue/017/midipara2  I32 F_XET
/-show/showfile/cue/018  <SCUE> n=0
    /-show/showfile/cue/018/numb  I32 F_XET
    /-show/showfile/cue/018/name  S32 F_XET
    /-show/showfile/cue/018/skip  I32 F_XET
    /-show/showfile/cue/018/scene  I32 F_XET
    /-show/showfile/cue/018/bit  I32 F_XET
    /-show/showfile/cue/018/miditype  I32 F_XET
    /-show/showfile/cue/018/midichan  I32 F_XET
    /-show/showfile/cue/018/midipara1  I32 F_XET
    /-show/showfile/cue/018/midipara2  I32 F_XET
/-show/showfile/cue/019  <SCUE> n=0
    /-show/showfile/cue/019/numb  I32 F_XET
    /-show/showfile/cue/019/name  S32 F_XET
    /-show/showfile/cue/019/skip  I32 F_XET
    /-show/showfile/cue/019/scene  I32 F_XET
    /-show/showfile/cue/019/bit  I32 F_XET
    /-show/showfile/cue/019/miditype  I32 F_XET
    /-show/showfile/cue/019/midichan  I32 F_XET
    /-show/showfile/cue/019/midipara1  I32 F_XET
    /-show/showfile/cue/019/midipara2  I32 F_XET
/-show/showfile/cue/020  <SCUE> n=0
    /-show/showfile/cue/020/numb  I32 F_XET
    /-show/showfile/cue/020/name  S32 F_XET
    /-show/showfile/cue/020/skip  I32 F_XET
    /-show/showfile/cue/020/scene  I32 F_XET
    /-show/showfile/cue/020/bit  I32 F_XET
    /-show/showfile/cue/020/miditype  I32 F_XET
    /-show/showfile/cue/020/midichan  I32 F_XET
    /-show/showfile/cue/020/midipara1  I32 F_XET
    /-show/showfile/cue/020/midipara2  I32 F_XET
/-show/showfile/cue/021  <SCUE> n=0
    /-show/showfile/cue/021/numb  I32 F_XET
    /-show/showfile/cue/021/name  S32 F_XET
    /-show/showfile/cue/021/skip  I32 F_XET
    /-show/showfile/cue/021/scene  I32 F_XET
    /-show/showfile/cue/021/bit  I32 F_XET
    /-show/showfile/cue/021/miditype  I32 F_XET
    /-show/showfile/cue/021/midichan  I32 F_XET
    /-show/showfile/cue/021/midipara1  I32 F_XET
    /-show/showfile/cue/021/midipara2  I32 F_XET
/-show/showfile/cue/022  <SCUE> n=0
    /-show/showfile/cue/022/numb  I32 F_XET
    /-show/showfile/cue/022/name  S32 F_XET
    /-show/showfile/cue/022/skip  I32 F_XET
    /-show/showfile/cue/022/scene  I32 F_XET
    /-show/showfile/cue/022/bit  I32 F_XET
    /-show/showfile/cue/022/miditype  I32 F_XET
    /-show/showfile/cue/022/midichan  I32 F_XET
    /-show/showfile/cue/022/midipara1  I32 F_XET
    /-show/showfile/cue/022/midipara2  I32 F_XET
/-show/showfile/cue/023  <SCUE> n=0
    /-show/showfile/cue/023/numb  I32 F_XET
    /-show/showfile/cue/023/name  S32 F_XET
    /-show/showfile/cue/023/skip  I32 F_XET
    /-show/showfile/cue/023/scene  I32 F_XET
    /-show/showfile/cue/023/bit  I32 F_XET
    /-show/showfile/cue/023/miditype  I32 F_XET
    /-show/showfile/cue/023/midichan  I32 F_XET
    /-show/showfile/cue/023/midipara1  I32 F_XET
    /-show/showfile/cue/023/midipara2  I32 F_XET
/-show/showfile/cue/024  <SCUE> n=0
    /-show/showfile/cue/024/numb  I32 F_XET
    /-show/showfile/cue/024/name  S32 F_XET
    /-show/showfile/cue/024/skip  I32 F_XET
    /-show/showfile/cue/024/scene  I32 F_XET
    /-show/showfile/cue/024/bit  I32 F_XET
    /-show/showfile/cue/024/miditype  I32 F_XET
    /-show/showfile/cue/024/midichan  I32 F_XET
    /-show/showfile/cue/024/midipara1  I32 F_XET
    /-show/showfile/cue/024/midipara2  I32 F_XET
/-show/showfile/cue/025  <SCUE> n=0
    /-show/showfile/cue/025/numb  I32 F_XET
    /-show/showfile/cue/025/name  S32 F_XET
    /-show/showfile/cue/025/skip  I32 F_XET
    /-show/showfile/cue/025/scene  I32 F_XET
    /-show/showfile/cue/025/bit  I32 F_XET
    /-show/showfile/cue/025/miditype  I32 F_XET
    /-show/showfile/cue/025/midichan  I32 F_XET
    /-show/showfile/cue/025/midipara1  I32 F_XET
    /-show/showfile/cue/025/midipara2  I32 F_XET
/-show/showfile/cue/026  <SCUE> n=0
    /-show/showfile/cue/026/numb  I32 F_XET
    /-show/showfile/cue/026/name  S32 F_XET
    /-show/showfile/cue/026/skip  I32 F_XET
    /-show/showfile/cue/026/scene  I32 F_XET
    /-show/showfile/cue/026/bit  I32 F_XET
    /-show/showfile/cue/026/miditype  I32 F_XET
    /-show/showfile/cue/026/midichan  I32 F_XET
    /-show/showfile/cue/026/midipara1  I32 F_XET
    /-show/showfile/cue/026/midipara2  I32 F_XET
/-show/showfile/cue/027  <SCUE> n=0
    /-show/showfile/cue/027/numb  I32 F_XET
    /-show/showfile/cue/027/name  S32 F_XET
    /-show/showfile/cue/027/skip  I32 F_XET
    /-show/showfile/cue/027/scene  I32 F_XET
    /-show/showfile/cue/027/bit  I32 F_XET
    /-show/showfile/cue/027/miditype  I32 F_XET
    /-show/showfile/cue/027/midichan  I32 F_XET
    /-show/showfile/cue/027/midipara1  I32 F_XET
    /-show/showfile/cue/027/midipara2  I32 F_XET
/-show/showfile/cue/028  <SCUE> n=0
    /-show/showfile/cue/028/numb  I32 F_XET
    /-show/showfile/cue/028/name  S32 F_XET
    /-show/showfile/cue/028/skip  I32 F_XET
    /-show/showfile/cue/028/scene  I32 F_XET
    /-show/showfile/cue/028/bit  I32 F_XET
    /-show/showfile/cue/028/miditype  I32 F_XET
    /-show/showfile/cue/028/midichan  I32 F_XET
    /-show/showfile/cue/028/midipara1  I32 F_XET
    /-show/showfile/cue/028/midipara2  I32 F_XET
/-show/showfile/cue/029  <SCUE> n=0
    /-show/showfile/cue/029/numb  I32 F_XET
    /-show/showfile/cue/029/name  S32 F_XET
    /-show/showfile/cue/029/skip  I32 F_XET
    /-show/showfile/cue/029/scene  I32 F_XET
    /-show/showfile/cue/029/bit  I32 F_XET
    /-show/showfile/cue/029/miditype  I32 F_XET
    /-show/showfile/cue/029/midichan  I32 F_XET
    /-show/showfile/cue/029/midipara1  I32 F_XET
    /-show/showfile/cue/029/midipara2  I32 F_XET
/-show/showfile/cue/030  <SCUE> n=0
    /-show/showfile/cue/030/numb  I32 F_XET
    /-show/showfile/cue/030/name  S32 F_XET
    /-show/showfile/cue/030/skip  I32 F_XET
    /-show/showfile/cue/030/scene  I32 F_XET
    /-show/showfile/cue/030/bit  I32 F_XET
    /-show/showfile/cue/030/miditype  I32 F_XET
    /-show/showfile/cue/030/midichan  I32 F_XET
    /-show/showfile/cue/030/midipara1  I32 F_XET
    /-show/showfile/cue/030/midipara2  I32 F_XET
/-show/showfile/cue/031  <SCUE> n=0
    /-show/showfile/cue/031/numb  I32 F_XET
    /-show/showfile/cue/031/name  S32 F_XET
    /-show/showfile/cue/031/skip  I32 F_XET
    /-show/showfile/cue/031/scene  I32 F_XET
    /-show/showfile/cue/031/bit  I32 F_XET
    /-show/showfile/cue/031/miditype  I32 F_XET
    /-show/showfile/cue/031/midichan  I32 F_XET
    /-show/showfile/cue/031/midipara1  I32 F_XET
    /-show/showfile/cue/031/midipara2  I32 F_XET
/-show/showfile/cue/032  <SCUE> n=0
    /-show/showfile/cue/032/numb  I32 F_XET
    /-show/showfile/cue/032/name  S32 F_XET
    /-show/showfile/cue/032/skip  I32 F_XET
    /-show/showfile/cue/032/scene  I32 F_XET
    /-show/showfile/cue/032/bit  I32 F_XET
    /-show/showfile/cue/032/miditype  I32 F_XET
    /-show/showfile/cue/032/midichan  I32 F_XET
    /-show/showfile/cue/032/midipara1  I32 F_XET
    /-show/showfile/cue/032/midipara2  I32 F_XET
/-show/showfile/cue/033  <SCUE> n=0
    /-show/showfile/cue/033/numb  I32 F_XET
    /-show/showfile/cue/033/name  S32 F_XET
    /-show/showfile/cue/033/skip  I32 F_XET
    /-show/showfile/cue/033/scene  I32 F_XET
    /-show/showfile/cue/033/bit  I32 F_XET
    /-show/showfile/cue/033/miditype  I32 F_XET
    /-show/showfile/cue/033/midichan  I32 F_XET
    /-show/showfile/cue/033/midipara1  I32 F_XET
    /-show/showfile/cue/033/midipara2  I32 F_XET
/-show/showfile/cue/034  <SCUE> n=0
    /-show/showfile/cue/034/numb  I32 F_XET
    /-show/showfile/cue/034/name  S32 F_XET
    /-show/showfile/cue/034/skip  I32 F_XET
    /-show/showfile/cue/034/scene  I32 F_XET
    /-show/showfile/cue/034/bit  I32 F_XET
    /-show/showfile/cue/034/miditype  I32 F_XET
    /-show/showfile/cue/034/midichan  I32 F_XET
    /-show/showfile/cue/034/midipara1  I32 F_XET
    /-show/showfile/cue/034/midipara2  I32 F_XET
/-show/showfile/cue/035  <SCUE> n=0
    /-show/showfile/cue/035/numb  I32 F_XET
    /-show/showfile/cue/035/name  S32 F_XET
    /-show/showfile/cue/035/skip  I32 F_XET
    /-show/showfile/cue/035/scene  I32 F_XET
    /-show/showfile/cue/035/bit  I32 F_XET
    /-show/showfile/cue/035/miditype  I32 F_XET
    /-show/showfile/cue/035/midichan  I32 F_XET
    /-show/showfile/cue/035/midipara1  I32 F_XET
    /-show/showfile/cue/035/midipara2  I32 F_XET
/-show/showfile/cue/036  <SCUE> n=0
    /-show/showfile/cue/036/numb  I32 F_XET
    /-show/showfile/cue/036/name  S32 F_XET
    /-show/showfile/cue/036/skip  I32 F_XET
    /-show/showfile/cue/036/scene  I32 F_XET
    /-show/showfile/cue/036/bit  I32 F_XET
    /-show/showfile/cue/036/miditype  I32 F_XET
    /-show/showfile/cue/036/midichan  I32 F_XET
    /-show/showfile/cue/036/midipara1  I32 F_XET
    /-show/showfile/cue/036/midipara2  I32 F_XET
/-show/showfile/cue/037  <SCUE> n=0
    /-show/showfile/cue/037/numb  I32 F_XET
    /-show/showfile/cue/037/name  S32 F_XET
    /-show/showfile/cue/037/skip  I32 F_XET
    /-show/showfile/cue/037/scene  I32 F_XET
    /-show/showfile/cue/037/bit  I32 F_XET
    /-show/showfile/cue/037/miditype  I32 F_XET
    /-show/showfile/cue/037/midichan  I32 F_XET
    /-show/showfile/cue/037/midipara1  I32 F_XET
    /-show/showfile/cue/037/midipara2  I32 F_XET
/-show/showfile/cue/038  <SCUE> n=0
    /-show/showfile/cue/038/numb  I32 F_XET
    /-show/showfile/cue/038/name  S32 F_XET
    /-show/showfile/cue/038/skip  I32 F_XET
    /-show/showfile/cue/038/scene  I32 F_XET
    /-show/showfile/cue/038/bit  I32 F_XET
    /-show/showfile/cue/038/miditype  I32 F_XET
    /-show/showfile/cue/038/midichan  I32 F_XET
    /-show/showfile/cue/038/midipara1  I32 F_XET
    /-show/showfile/cue/038/midipara2  I32 F_XET
/-show/showfile/cue/039  <SCUE> n=0
    /-show/showfile/cue/039/numb  I32 F_XET
    /-show/showfile/cue/039/name  S32 F_XET
    /-show/showfile/cue/039/skip  I32 F_XET
    /-show/showfile/cue/039/scene  I32 F_XET
    /-show/showfile/cue/039/bit  I32 F_XET
    /-show/showfile/cue/039/miditype  I32 F_XET
    /-show/showfile/cue/039/midichan  I32 F_XET
    /-show/showfile/cue/039/midipara1  I32 F_XET
    /-show/showfile/cue/039/midipara2  I32 F_XET
/-show/showfile/cue/040  <SCUE> n=0
    /-show/showfile/cue/040/numb  I32 F_XET
    /-show/showfile/cue/040/name  S32 F_XET
    /-show/showfile/cue/040/skip  I32 F_XET
    /-show/showfile/cue/040/scene  I32 F_XET
    /-show/showfile/cue/040/bit  I32 F_XET
    /-show/showfile/cue/040/miditype  I32 F_XET
    /-show/showfile/cue/040/midichan  I32 F_XET
    /-show/showfile/cue/040/midipara1  I32 F_XET
    /-show/showfile/cue/040/midipara2  I32 F_XET
/-show/showfile/cue/041  <SCUE> n=0
    /-show/showfile/cue/041/numb  I32 F_XET
    /-show/showfile/cue/041/name  S32 F_XET
    /-show/showfile/cue/041/skip  I32 F_XET
    /-show/showfile/cue/041/scene  I32 F_XET
    /-show/showfile/cue/041/bit  I32 F_XET
    /-show/showfile/cue/041/miditype  I32 F_XET
    /-show/showfile/cue/041/midichan  I32 F_XET
    /-show/showfile/cue/041/midipara1  I32 F_XET
    /-show/showfile/cue/041/midipara2  I32 F_XET
/-show/showfile/cue/042  <SCUE> n=0
    /-show/showfile/cue/042/numb  I32 F_XET
    /-show/showfile/cue/042/name  S32 F_XET
    /-show/showfile/cue/042/skip  I32 F_XET
    /-show/showfile/cue/042/scene  I32 F_XET
    /-show/showfile/cue/042/bit  I32 F_XET
    /-show/showfile/cue/042/miditype  I32 F_XET
    /-show/showfile/cue/042/midichan  I32 F_XET
    /-show/showfile/cue/042/midipara1  I32 F_XET
    /-show/showfile/cue/042/midipara2  I32 F_XET
/-show/showfile/cue/043  <SCUE> n=0
    /-show/showfile/cue/043/numb  I32 F_XET
    /-show/showfile/cue/043/name  S32 F_XET
    /-show/showfile/cue/043/skip  I32 F_XET
    /-show/showfile/cue/043/scene  I32 F_XET
    /-show/showfile/cue/043/bit  I32 F_XET
    /-show/showfile/cue/043/miditype  I32 F_XET
    /-show/showfile/cue/043/midichan  I32 F_XET
    /-show/showfile/cue/043/midipara1  I32 F_XET
    /-show/showfile/cue/043/midipara2  I32 F_XET
/-show/showfile/cue/044  <SCUE> n=0
    /-show/showfile/cue/044/numb  I32 F_XET
    /-show/showfile/cue/044/name  S32 F_XET
    /-show/showfile/cue/044/skip  I32 F_XET
    /-show/showfile/cue/044/scene  I32 F_XET
    /-show/showfile/cue/044/bit  I32 F_XET
    /-show/showfile/cue/044/miditype  I32 F_XET
    /-show/showfile/cue/044/midichan  I32 F_XET
    /-show/showfile/cue/044/midipara1  I32 F_XET
    /-show/showfile/cue/044/midipara2  I32 F_XET
/-show/showfile/cue/045  <SCUE> n=0
    /-show/showfile/cue/045/numb  I32 F_XET
    /-show/showfile/cue/045/name  S32 F_XET
    /-show/showfile/cue/045/skip  I32 F_XET
    /-show/showfile/cue/045/scene  I32 F_XET
    /-show/showfile/cue/045/bit  I32 F_XET
    /-show/showfile/cue/045/miditype  I32 F_XET
    /-show/showfile/cue/045/midichan  I32 F_XET
    /-show/showfile/cue/045/midipara1  I32 F_XET
    /-show/showfile/cue/045/midipara2  I32 F_XET
/-show/showfile/cue/046  <SCUE> n=0
    /-show/showfile/cue/046/numb  I32 F_XET
    /-show/showfile/cue/046/name  S32 F_XET
    /-show/showfile/cue/046/skip  I32 F_XET
    /-show/showfile/cue/046/scene  I32 F_XET
    /-show/showfile/cue/046/bit  I32 F_XET
    /-show/showfile/cue/046/miditype  I32 F_XET
    /-show/showfile/cue/046/midichan  I32 F_XET
    /-show/showfile/cue/046/midipara1  I32 F_XET
    /-show/showfile/cue/046/midipara2  I32 F_XET
/-show/showfile/cue/047  <SCUE> n=0
    /-show/showfile/cue/047/numb  I32 F_XET
    /-show/showfile/cue/047/name  S32 F_XET
    /-show/showfile/cue/047/skip  I32 F_XET
    /-show/showfile/cue/047/scene  I32 F_XET
    /-show/showfile/cue/047/bit  I32 F_XET
    /-show/showfile/cue/047/miditype  I32 F_XET
    /-show/showfile/cue/047/midichan  I32 F_XET
    /-show/showfile/cue/047/midipara1  I32 F_XET
    /-show/showfile/cue/047/midipara2  I32 F_XET
/-show/showfile/cue/048  <SCUE> n=0
    /-show/showfile/cue/048/numb  I32 F_XET
    /-show/showfile/cue/048/name  S32 F_XET
    /-show/showfile/cue/048/skip  I32 F_XET
    /-show/showfile/cue/048/scene  I32 F_XET
    /-show/showfile/cue/048/bit  I32 F_XET
    /-show/showfile/cue/048/miditype  I32 F_XET
    /-show/showfile/cue/048/midichan  I32 F_XET
    /-show/showfile/cue/048/midipara1  I32 F_XET
    /-show/showfile/cue/048/midipara2  I32 F_XET
/-show/showfile/cue/049  <SCUE> n=0
    /-show/showfile/cue/049/numb  I32 F_XET
    /-show/showfile/cue/049/name  S32 F_XET
    /-show/showfile/cue/049/skip  I32 F_XET
    /-show/showfile/cue/049/scene  I32 F_XET
    /-show/showfile/cue/049/bit  I32 F_XET
    /-show/showfile/cue/049/miditype  I32 F_XET
    /-show/showfile/cue/049/midichan  I32 F_XET
    /-show/showfile/cue/049/midipara1  I32 F_XET
    /-show/showfile/cue/049/midipara2  I32 F_XET
/-show/showfile/cue/050  <SCUE> n=0
    /-show/showfile/cue/050/numb  I32 F_XET
    /-show/showfile/cue/050/name  S32 F_XET
    /-show/showfile/cue/050/skip  I32 F_XET
    /-show/showfile/cue/050/scene  I32 F_XET
    /-show/showfile/cue/050/bit  I32 F_XET
    /-show/showfile/cue/050/miditype  I32 F_XET
    /-show/showfile/cue/050/midichan  I32 F_XET
    /-show/showfile/cue/050/midipara1  I32 F_XET
    /-show/showfile/cue/050/midipara2  I32 F_XET
/-show/showfile/cue/051  <SCUE> n=0
    /-show/showfile/cue/051/numb  I32 F_XET
    /-show/showfile/cue/051/name  S32 F_XET
    /-show/showfile/cue/051/skip  I32 F_XET
    /-show/showfile/cue/051/scene  I32 F_XET
    /-show/showfile/cue/051/bit  I32 F_XET
    /-show/showfile/cue/051/miditype  I32 F_XET
    /-show/showfile/cue/051/midichan  I32 F_XET
    /-show/showfile/cue/051/midipara1  I32 F_XET
    /-show/showfile/cue/051/midipara2  I32 F_XET
/-show/showfile/cue/052  <SCUE> n=0
    /-show/showfile/cue/052/numb  I32 F_XET
    /-show/showfile/cue/052/name  S32 F_XET
    /-show/showfile/cue/052/skip  I32 F_XET
    /-show/showfile/cue/052/scene  I32 F_XET
    /-show/showfile/cue/052/bit  I32 F_XET
    /-show/showfile/cue/052/miditype  I32 F_XET
    /-show/showfile/cue/052/midichan  I32 F_XET
    /-show/showfile/cue/052/midipara1  I32 F_XET
    /-show/showfile/cue/052/midipara2  I32 F_XET
/-show/showfile/cue/053  <SCUE> n=0
    /-show/showfile/cue/053/numb  I32 F_XET
    /-show/showfile/cue/053/name  S32 F_XET
    /-show/showfile/cue/053/skip  I32 F_XET
    /-show/showfile/cue/053/scene  I32 F_XET
    /-show/showfile/cue/053/bit  I32 F_XET
    /-show/showfile/cue/053/miditype  I32 F_XET
    /-show/showfile/cue/053/midichan  I32 F_XET
    /-show/showfile/cue/053/midipara1  I32 F_XET
    /-show/showfile/cue/053/midipara2  I32 F_XET
/-show/showfile/cue/054  <SCUE> n=0
    /-show/showfile/cue/054/numb  I32 F_XET
    /-show/showfile/cue/054/name  S32 F_XET
    /-show/showfile/cue/054/skip  I32 F_XET
    /-show/showfile/cue/054/scene  I32 F_XET
    /-show/showfile/cue/054/bit  I32 F_XET
    /-show/showfile/cue/054/miditype  I32 F_XET
    /-show/showfile/cue/054/midichan  I32 F_XET
    /-show/showfile/cue/054/midipara1  I32 F_XET
    /-show/showfile/cue/054/midipara2  I32 F_XET
/-show/showfile/cue/055  <SCUE> n=0
    /-show/showfile/cue/055/numb  I32 F_XET
    /-show/showfile/cue/055/name  S32 F_XET
    /-show/showfile/cue/055/skip  I32 F_XET
    /-show/showfile/cue/055/scene  I32 F_XET
    /-show/showfile/cue/055/bit  I32 F_XET
    /-show/showfile/cue/055/miditype  I32 F_XET
    /-show/showfile/cue/055/midichan  I32 F_XET
    /-show/showfile/cue/055/midipara1  I32 F_XET
    /-show/showfile/cue/055/midipara2  I32 F_XET
/-show/showfile/cue/056  <SCUE> n=0
    /-show/showfile/cue/056/numb  I32 F_XET
    /-show/showfile/cue/056/name  S32 F_XET
    /-show/showfile/cue/056/skip  I32 F_XET
    /-show/showfile/cue/056/scene  I32 F_XET
    /-show/showfile/cue/056/bit  I32 F_XET
    /-show/showfile/cue/056/miditype  I32 F_XET
    /-show/showfile/cue/056/midichan  I32 F_XET
    /-show/showfile/cue/056/midipara1  I32 F_XET
    /-show/showfile/cue/056/midipara2  I32 F_XET
/-show/showfile/cue/057  <SCUE> n=0
    /-show/showfile/cue/057/numb  I32 F_XET
    /-show/showfile/cue/057/name  S32 F_XET
    /-show/showfile/cue/057/skip  I32 F_XET
    /-show/showfile/cue/057/scene  I32 F_XET
    /-show/showfile/cue/057/bit  I32 F_XET
    /-show/showfile/cue/057/miditype  I32 F_XET
    /-show/showfile/cue/057/midichan  I32 F_XET
    /-show/showfile/cue/057/midipara1  I32 F_XET
    /-show/showfile/cue/057/midipara2  I32 F_XET
/-show/showfile/cue/058  <SCUE> n=0
    /-show/showfile/cue/058/numb  I32 F_XET
    /-show/showfile/cue/058/name  S32 F_XET
    /-show/showfile/cue/058/skip  I32 F_XET
    /-show/showfile/cue/058/scene  I32 F_XET
    /-show/showfile/cue/058/bit  I32 F_XET
    /-show/showfile/cue/058/miditype  I32 F_XET
    /-show/showfile/cue/058/midichan  I32 F_XET
    /-show/showfile/cue/058/midipara1  I32 F_XET
    /-show/showfile/cue/058/midipara2  I32 F_XET
/-show/showfile/cue/059  <SCUE> n=0
    /-show/showfile/cue/059/numb  I32 F_XET
    /-show/showfile/cue/059/name  S32 F_XET
    /-show/showfile/cue/059/skip  I32 F_XET
    /-show/showfile/cue/059/scene  I32 F_XET
    /-show/showfile/cue/059/bit  I32 F_XET
    /-show/showfile/cue/059/miditype  I32 F_XET
    /-show/showfile/cue/059/midichan  I32 F_XET
    /-show/showfile/cue/059/midipara1  I32 F_XET
    /-show/showfile/cue/059/midipara2  I32 F_XET
/-show/showfile/cue/060  <SCUE> n=0
    /-show/showfile/cue/060/numb  I32 F_XET
    /-show/showfile/cue/060/name  S32 F_XET
    /-show/showfile/cue/060/skip  I32 F_XET
    /-show/showfile/cue/060/scene  I32 F_XET
    /-show/showfile/cue/060/bit  I32 F_XET
    /-show/showfile/cue/060/miditype  I32 F_XET
    /-show/showfile/cue/060/midichan  I32 F_XET
    /-show/showfile/cue/060/midipara1  I32 F_XET
    /-show/showfile/cue/060/midipara2  I32 F_XET
/-show/showfile/cue/061  <SCUE> n=0
    /-show/showfile/cue/061/numb  I32 F_XET
    /-show/showfile/cue/061/name  S32 F_XET
    /-show/showfile/cue/061/skip  I32 F_XET
    /-show/showfile/cue/061/scene  I32 F_XET
    /-show/showfile/cue/061/bit  I32 F_XET
    /-show/showfile/cue/061/miditype  I32 F_XET
    /-show/showfile/cue/061/midichan  I32 F_XET
    /-show/showfile/cue/061/midipara1  I32 F_XET
    /-show/showfile/cue/061/midipara2  I32 F_XET
/-show/showfile/cue/062  <SCUE> n=0
    /-show/showfile/cue/062/numb  I32 F_XET
    /-show/showfile/cue/062/name  S32 F_XET
    /-show/showfile/cue/062/skip  I32 F_XET
    /-show/showfile/cue/062/scene  I32 F_XET
    /-show/showfile/cue/062/bit  I32 F_XET
    /-show/showfile/cue/062/miditype  I32 F_XET
    /-show/showfile/cue/062/midichan  I32 F_XET
    /-show/showfile/cue/062/midipara1  I32 F_XET
    /-show/showfile/cue/062/midipara2  I32 F_XET
/-show/showfile/cue/063  <SCUE> n=0
    /-show/showfile/cue/063/numb  I32 F_XET
    /-show/showfile/cue/063/name  S32 F_XET
    /-show/showfile/cue/063/skip  I32 F_XET
    /-show/showfile/cue/063/scene  I32 F_XET
    /-show/showfile/cue/063/bit  I32 F_XET
    /-show/showfile/cue/063/miditype  I32 F_XET
    /-show/showfile/cue/063/midichan  I32 F_XET
    /-show/showfile/cue/063/midipara1  I32 F_XET
    /-show/showfile/cue/063/midipara2  I32 F_XET
/-show/showfile/cue/064  <SCUE> n=0
    /-show/showfile/cue/064/numb  I32 F_XET
    /-show/showfile/cue/064/name  S32 F_XET
    /-show/showfile/cue/064/skip  I32 F_XET
    /-show/showfile/cue/064/scene  I32 F_XET
    /-show/showfile/cue/064/bit  I32 F_XET
    /-show/showfile/cue/064/miditype  I32 F_XET
    /-show/showfile/cue/064/midichan  I32 F_XET
    /-show/showfile/cue/064/midipara1  I32 F_XET
    /-show/showfile/cue/064/midipara2  I32 F_XET
/-show/showfile/cue/065  <SCUE> n=0
    /-show/showfile/cue/065/numb  I32 F_XET
    /-show/showfile/cue/065/name  S32 F_XET
    /-show/showfile/cue/065/skip  I32 F_XET
    /-show/showfile/cue/065/scene  I32 F_XET
    /-show/showfile/cue/065/bit  I32 F_XET
    /-show/showfile/cue/065/miditype  I32 F_XET
    /-show/showfile/cue/065/midichan  I32 F_XET
    /-show/showfile/cue/065/midipara1  I32 F_XET
    /-show/showfile/cue/065/midipara2  I32 F_XET
/-show/showfile/cue/066  <SCUE> n=0
    /-show/showfile/cue/066/numb  I32 F_XET
    /-show/showfile/cue/066/name  S32 F_XET
    /-show/showfile/cue/066/skip  I32 F_XET
    /-show/showfile/cue/066/scene  I32 F_XET
    /-show/showfile/cue/066/bit  I32 F_XET
    /-show/showfile/cue/066/miditype  I32 F_XET
    /-show/showfile/cue/066/midichan  I32 F_XET
    /-show/showfile/cue/066/midipara1  I32 F_XET
    /-show/showfile/cue/066/midipara2  I32 F_XET
/-show/showfile/cue/067  <SCUE> n=0
    /-show/showfile/cue/067/numb  I32 F_XET
    /-show/showfile/cue/067/name  S32 F_XET
    /-show/showfile/cue/067/skip  I32 F_XET
    /-show/showfile/cue/067/scene  I32 F_XET
    /-show/showfile/cue/067/bit  I32 F_XET
    /-show/showfile/cue/067/miditype  I32 F_XET
    /-show/showfile/cue/067/midichan  I32 F_XET
    /-show/showfile/cue/067/midipara1  I32 F_XET
    /-show/showfile/cue/067/midipara2  I32 F_XET
/-show/showfile/cue/068  <SCUE> n=0
    /-show/showfile/cue/068/numb  I32 F_XET
    /-show/showfile/cue/068/name  S32 F_XET
    /-show/showfile/cue/068/skip  I32 F_XET
    /-show/showfile/cue/068/scene  I32 F_XET
    /-show/showfile/cue/068/bit  I32 F_XET
    /-show/showfile/cue/068/miditype  I32 F_XET
    /-show/showfile/cue/068/midichan  I32 F_XET
    /-show/showfile/cue/068/midipara1  I32 F_XET
    /-show/showfile/cue/068/midipara2  I32 F_XET
/-show/showfile/cue/069  <SCUE> n=0
    /-show/showfile/cue/069/numb  I32 F_XET
    /-show/showfile/cue/069/name  S32 F_XET
    /-show/showfile/cue/069/skip  I32 F_XET
    /-show/showfile/cue/069/scene  I32 F_XET
    /-show/showfile/cue/069/bit  I32 F_XET
    /-show/showfile/cue/069/miditype  I32 F_XET
    /-show/showfile/cue/069/midichan  I32 F_XET
    /-show/showfile/cue/069/midipara1  I32 F_XET
    /-show/showfile/cue/069/midipara2  I32 F_XET
/-show/showfile/cue/070  <SCUE> n=0
    /-show/showfile/cue/070/numb  I32 F_XET
    /-show/showfile/cue/070/name  S32 F_XET
    /-show/showfile/cue/070/skip  I32 F_XET
    /-show/showfile/cue/070/scene  I32 F_XET
    /-show/showfile/cue/070/bit  I32 F_XET
    /-show/showfile/cue/070/miditype  I32 F_XET
    /-show/showfile/cue/070/midichan  I32 F_XET
    /-show/showfile/cue/070/midipara1  I32 F_XET
    /-show/showfile/cue/070/midipara2  I32 F_XET
/-show/showfile/cue/071  <SCUE> n=0
    /-show/showfile/cue/071/numb  I32 F_XET
    /-show/showfile/cue/071/name  S32 F_XET
    /-show/showfile/cue/071/skip  I32 F_XET
    /-show/showfile/cue/071/scene  I32 F_XET
    /-show/showfile/cue/071/bit  I32 F_XET
    /-show/showfile/cue/071/miditype  I32 F_XET
    /-show/showfile/cue/071/midichan  I32 F_XET
    /-show/showfile/cue/071/midipara1  I32 F_XET
    /-show/showfile/cue/071/midipara2  I32 F_XET
/-show/showfile/cue/072  <SCUE> n=0
    /-show/showfile/cue/072/numb  I32 F_XET
    /-show/showfile/cue/072/name  S32 F_XET
    /-show/showfile/cue/072/skip  I32 F_XET
    /-show/showfile/cue/072/scene  I32 F_XET
    /-show/showfile/cue/072/bit  I32 F_XET
    /-show/showfile/cue/072/miditype  I32 F_XET
    /-show/showfile/cue/072/midichan  I32 F_XET
    /-show/showfile/cue/072/midipara1  I32 F_XET
    /-show/showfile/cue/072/midipara2  I32 F_XET
/-show/showfile/cue/073  <SCUE> n=0
    /-show/showfile/cue/073/numb  I32 F_XET
    /-show/showfile/cue/073/name  S32 F_XET
    /-show/showfile/cue/073/skip  I32 F_XET
    /-show/showfile/cue/073/scene  I32 F_XET
    /-show/showfile/cue/073/bit  I32 F_XET
    /-show/showfile/cue/073/miditype  I32 F_XET
    /-show/showfile/cue/073/midichan  I32 F_XET
    /-show/showfile/cue/073/midipara1  I32 F_XET
    /-show/showfile/cue/073/midipara2  I32 F_XET
/-show/showfile/cue/074  <SCUE> n=0
    /-show/showfile/cue/074/numb  I32 F_XET
    /-show/showfile/cue/074/name  S32 F_XET
    /-show/showfile/cue/074/skip  I32 F_XET
    /-show/showfile/cue/074/scene  I32 F_XET
    /-show/showfile/cue/074/bit  I32 F_XET
    /-show/showfile/cue/074/miditype  I32 F_XET
    /-show/showfile/cue/074/midichan  I32 F_XET
    /-show/showfile/cue/074/midipara1  I32 F_XET
    /-show/showfile/cue/074/midipara2  I32 F_XET
/-show/showfile/cue/075  <SCUE> n=0
    /-show/showfile/cue/075/numb  I32 F_XET
    /-show/showfile/cue/075/name  S32 F_XET
    /-show/showfile/cue/075/skip  I32 F_XET
    /-show/showfile/cue/075/scene  I32 F_XET
    /-show/showfile/cue/075/bit  I32 F_XET
    /-show/showfile/cue/075/miditype  I32 F_XET
    /-show/showfile/cue/075/midichan  I32 F_XET
    /-show/showfile/cue/075/midipara1  I32 F_XET
    /-show/showfile/cue/075/midipara2  I32 F_XET
/-show/showfile/cue/076  <SCUE> n=0
    /-show/showfile/cue/076/numb  I32 F_XET
    /-show/showfile/cue/076/name  S32 F_XET
    /-show/showfile/cue/076/skip  I32 F_XET
    /-show/showfile/cue/076/scene  I32 F_XET
    /-show/showfile/cue/076/bit  I32 F_XET
    /-show/showfile/cue/076/miditype  I32 F_XET
    /-show/showfile/cue/076/midichan  I32 F_XET
    /-show/showfile/cue/076/midipara1  I32 F_XET
    /-show/showfile/cue/076/midipara2  I32 F_XET
/-show/showfile/cue/077  <SCUE> n=0
    /-show/showfile/cue/077/numb  I32 F_XET
    /-show/showfile/cue/077/name  S32 F_XET
    /-show/showfile/cue/077/skip  I32 F_XET
    /-show/showfile/cue/077/scene  I32 F_XET
    /-show/showfile/cue/077/bit  I32 F_XET
    /-show/showfile/cue/077/miditype  I32 F_XET
    /-show/showfile/cue/077/midichan  I32 F_XET
    /-show/showfile/cue/077/midipara1  I32 F_XET
    /-show/showfile/cue/077/midipara2  I32 F_XET
/-show/showfile/cue/078  <SCUE> n=0
    /-show/showfile/cue/078/numb  I32 F_XET
    /-show/showfile/cue/078/name  S32 F_XET
    /-show/showfile/cue/078/skip  I32 F_XET
    /-show/showfile/cue/078/scene  I32 F_XET
    /-show/showfile/cue/078/bit  I32 F_XET
    /-show/showfile/cue/078/miditype  I32 F_XET
    /-show/showfile/cue/078/midichan  I32 F_XET
    /-show/showfile/cue/078/midipara1  I32 F_XET
    /-show/showfile/cue/078/midipara2  I32 F_XET
/-show/showfile/cue/079  <SCUE> n=0
    /-show/showfile/cue/079/numb  I32 F_XET
    /-show/showfile/cue/079/name  S32 F_XET
    /-show/showfile/cue/079/skip  I32 F_XET
    /-show/showfile/cue/079/scene  I32 F_XET
    /-show/showfile/cue/079/bit  I32 F_XET
    /-show/showfile/cue/079/miditype  I32 F_XET
    /-show/showfile/cue/079/midichan  I32 F_XET
    /-show/showfile/cue/079/midipara1  I32 F_XET
    /-show/showfile/cue/079/midipara2  I32 F_XET
/-show/showfile/cue/080  <SCUE> n=0
    /-show/showfile/cue/080/numb  I32 F_XET
    /-show/showfile/cue/080/name  S32 F_XET
    /-show/showfile/cue/080/skip  I32 F_XET
    /-show/showfile/cue/080/scene  I32 F_XET
    /-show/showfile/cue/080/bit  I32 F_XET
    /-show/showfile/cue/080/miditype  I32 F_XET
    /-show/showfile/cue/080/midichan  I32 F_XET
    /-show/showfile/cue/080/midipara1  I32 F_XET
    /-show/showfile/cue/080/midipara2  I32 F_XET
/-show/showfile/cue/081  <SCUE> n=0
    /-show/showfile/cue/081/numb  I32 F_XET
    /-show/showfile/cue/081/name  S32 F_XET
    /-show/showfile/cue/081/skip  I32 F_XET
    /-show/showfile/cue/081/scene  I32 F_XET
    /-show/showfile/cue/081/bit  I32 F_XET
    /-show/showfile/cue/081/miditype  I32 F_XET
    /-show/showfile/cue/081/midichan  I32 F_XET
    /-show/showfile/cue/081/midipara1  I32 F_XET
    /-show/showfile/cue/081/midipara2  I32 F_XET
/-show/showfile/cue/082  <SCUE> n=0
    /-show/showfile/cue/082/numb  I32 F_XET
    /-show/showfile/cue/082/name  S32 F_XET
    /-show/showfile/cue/082/skip  I32 F_XET
    /-show/showfile/cue/082/scene  I32 F_XET
    /-show/showfile/cue/082/bit  I32 F_XET
    /-show/showfile/cue/082/miditype  I32 F_XET
    /-show/showfile/cue/082/midichan  I32 F_XET
    /-show/showfile/cue/082/midipara1  I32 F_XET
    /-show/showfile/cue/082/midipara2  I32 F_XET
/-show/showfile/cue/083  <SCUE> n=0
    /-show/showfile/cue/083/numb  I32 F_XET
    /-show/showfile/cue/083/name  S32 F_XET
    /-show/showfile/cue/083/skip  I32 F_XET
    /-show/showfile/cue/083/scene  I32 F_XET
    /-show/showfile/cue/083/bit  I32 F_XET
    /-show/showfile/cue/083/miditype  I32 F_XET
    /-show/showfile/cue/083/midichan  I32 F_XET
    /-show/showfile/cue/083/midipara1  I32 F_XET
    /-show/showfile/cue/083/midipara2  I32 F_XET
/-show/showfile/cue/084  <SCUE> n=0
    /-show/showfile/cue/084/numb  I32 F_XET
    /-show/showfile/cue/084/name  S32 F_XET
    /-show/showfile/cue/084/skip  I32 F_XET
    /-show/showfile/cue/084/scene  I32 F_XET
    /-show/showfile/cue/084/bit  I32 F_XET
    /-show/showfile/cue/084/miditype  I32 F_XET
    /-show/showfile/cue/084/midichan  I32 F_XET
    /-show/showfile/cue/084/midipara1  I32 F_XET
    /-show/showfile/cue/084/midipara2  I32 F_XET
/-show/showfile/cue/085  <SCUE> n=0
    /-show/showfile/cue/085/numb  I32 F_XET
    /-show/showfile/cue/085/name  S32 F_XET
    /-show/showfile/cue/085/skip  I32 F_XET
    /-show/showfile/cue/085/scene  I32 F_XET
    /-show/showfile/cue/085/bit  I32 F_XET
    /-show/showfile/cue/085/miditype  I32 F_XET
    /-show/showfile/cue/085/midichan  I32 F_XET
    /-show/showfile/cue/085/midipara1  I32 F_XET
    /-show/showfile/cue/085/midipara2  I32 F_XET
/-show/showfile/cue/086  <SCUE> n=0
    /-show/showfile/cue/086/numb  I32 F_XET
    /-show/showfile/cue/086/name  S32 F_XET
    /-show/showfile/cue/086/skip  I32 F_XET
    /-show/showfile/cue/086/scene  I32 F_XET
    /-show/showfile/cue/086/bit  I32 F_XET
    /-show/showfile/cue/086/miditype  I32 F_XET
    /-show/showfile/cue/086/midichan  I32 F_XET
    /-show/showfile/cue/086/midipara1  I32 F_XET
    /-show/showfile/cue/086/midipara2  I32 F_XET
/-show/showfile/cue/087  <SCUE> n=0
    /-show/showfile/cue/087/numb  I32 F_XET
    /-show/showfile/cue/087/name  S32 F_XET
    /-show/showfile/cue/087/skip  I32 F_XET
    /-show/showfile/cue/087/scene  I32 F_XET
    /-show/showfile/cue/087/bit  I32 F_XET
    /-show/showfile/cue/087/miditype  I32 F_XET
    /-show/showfile/cue/087/midichan  I32 F_XET
    /-show/showfile/cue/087/midipara1  I32 F_XET
    /-show/showfile/cue/087/midipara2  I32 F_XET
/-show/showfile/cue/088  <SCUE> n=0
    /-show/showfile/cue/088/numb  I32 F_XET
    /-show/showfile/cue/088/name  S32 F_XET
    /-show/showfile/cue/088/skip  I32 F_XET
    /-show/showfile/cue/088/scene  I32 F_XET
    /-show/showfile/cue/088/bit  I32 F_XET
    /-show/showfile/cue/088/miditype  I32 F_XET
    /-show/showfile/cue/088/midichan  I32 F_XET
    /-show/showfile/cue/088/midipara1  I32 F_XET
    /-show/showfile/cue/088/midipara2  I32 F_XET
/-show/showfile/cue/089  <SCUE> n=0
    /-show/showfile/cue/089/numb  I32 F_XET
    /-show/showfile/cue/089/name  S32 F_XET
    /-show/showfile/cue/089/skip  I32 F_XET
    /-show/showfile/cue/089/scene  I32 F_XET
    /-show/showfile/cue/089/bit  I32 F_XET
    /-show/showfile/cue/089/miditype  I32 F_XET
    /-show/showfile/cue/089/midichan  I32 F_XET
    /-show/showfile/cue/089/midipara1  I32 F_XET
    /-show/showfile/cue/089/midipara2  I32 F_XET
/-show/showfile/cue/090  <SCUE> n=0
    /-show/showfile/cue/090/numb  I32 F_XET
    /-show/showfile/cue/090/name  S32 F_XET
    /-show/showfile/cue/090/skip  I32 F_XET
    /-show/showfile/cue/090/scene  I32 F_XET
    /-show/showfile/cue/090/bit  I32 F_XET
    /-show/showfile/cue/090/miditype  I32 F_XET
    /-show/showfile/cue/090/midichan  I32 F_XET
    /-show/showfile/cue/090/midipara1  I32 F_XET
    /-show/showfile/cue/090/midipara2  I32 F_XET
/-show/showfile/cue/091  <SCUE> n=0
    /-show/showfile/cue/091/numb  I32 F_XET
    /-show/showfile/cue/091/name  S32 F_XET
    /-show/showfile/cue/091/skip  I32 F_XET
    /-show/showfile/cue/091/scene  I32 F_XET
    /-show/showfile/cue/091/bit  I32 F_XET
    /-show/showfile/cue/091/miditype  I32 F_XET
    /-show/showfile/cue/091/midichan  I32 F_XET
    /-show/showfile/cue/091/midipara1  I32 F_XET
    /-show/showfile/cue/091/midipara2  I32 F_XET
/-show/showfile/cue/092  <SCUE> n=0
    /-show/showfile/cue/092/numb  I32 F_XET
    /-show/showfile/cue/092/name  S32 F_XET
    /-show/showfile/cue/092/skip  I32 F_XET
    /-show/showfile/cue/092/scene  I32 F_XET
    /-show/showfile/cue/092/bit  I32 F_XET
    /-show/showfile/cue/092/miditype  I32 F_XET
    /-show/showfile/cue/092/midichan  I32 F_XET
    /-show/showfile/cue/092/midipara1  I32 F_XET
    /-show/showfile/cue/092/midipara2  I32 F_XET
/-show/showfile/cue/093  <SCUE> n=0
    /-show/showfile/cue/093/numb  I32 F_XET
    /-show/showfile/cue/093/name  S32 F_XET
    /-show/showfile/cue/093/skip  I32 F_XET
    /-show/showfile/cue/093/scene  I32 F_XET
    /-show/showfile/cue/093/bit  I32 F_XET
    /-show/showfile/cue/093/miditype  I32 F_XET
    /-show/showfile/cue/093/midichan  I32 F_XET
    /-show/showfile/cue/093/midipara1  I32 F_XET
    /-show/showfile/cue/093/midipara2  I32 F_XET
/-show/showfile/cue/094  <SCUE> n=0
    /-show/showfile/cue/094/numb  I32 F_XET
    /-show/showfile/cue/094/name  S32 F_XET
    /-show/showfile/cue/094/skip  I32 F_XET
    /-show/showfile/cue/094/scene  I32 F_XET
    /-show/showfile/cue/094/bit  I32 F_XET
    /-show/showfile/cue/094/miditype  I32 F_XET
    /-show/showfile/cue/094/midichan  I32 F_XET
    /-show/showfile/cue/094/midipara1  I32 F_XET
    /-show/showfile/cue/094/midipara2  I32 F_XET
/-show/showfile/cue/095  <SCUE> n=0
    /-show/showfile/cue/095/numb  I32 F_XET
    /-show/showfile/cue/095/name  S32 F_XET
    /-show/showfile/cue/095/skip  I32 F_XET
    /-show/showfile/cue/095/scene  I32 F_XET
    /-show/showfile/cue/095/bit  I32 F_XET
    /-show/showfile/cue/095/miditype  I32 F_XET
    /-show/showfile/cue/095/midichan  I32 F_XET
    /-show/showfile/cue/095/midipara1  I32 F_XET
    /-show/showfile/cue/095/midipara2  I32 F_XET
/-show/showfile/cue/096  <SCUE> n=0
    /-show/showfile/cue/096/numb  I32 F_XET
    /-show/showfile/cue/096/name  S32 F_XET
    /-show/showfile/cue/096/skip  I32 F_XET
    /-show/showfile/cue/096/scene  I32 F_XET
    /-show/showfile/cue/096/bit  I32 F_XET
    /-show/showfile/cue/096/miditype  I32 F_XET
    /-show/showfile/cue/096/midichan  I32 F_XET
    /-show/showfile/cue/096/midipara1  I32 F_XET
    /-show/showfile/cue/096/midipara2  I32 F_XET
/-show/showfile/cue/097  <SCUE> n=0
    /-show/showfile/cue/097/numb  I32 F_XET
    /-show/showfile/cue/097/name  S32 F_XET
    /-show/showfile/cue/097/skip  I32 F_XET
    /-show/showfile/cue/097/scene  I32 F_XET
    /-show/showfile/cue/097/bit  I32 F_XET
    /-show/showfile/cue/097/miditype  I32 F_XET
    /-show/showfile/cue/097/midichan  I32 F_XET
    /-show/showfile/cue/097/midipara1  I32 F_XET
    /-show/showfile/cue/097/midipara2  I32 F_XET
/-show/showfile/cue/098  <SCUE> n=0
    /-show/showfile/cue/098/numb  I32 F_XET
    /-show/showfile/cue/098/name  S32 F_XET
    /-show/showfile/cue/098/skip  I32 F_XET
    /-show/showfile/cue/098/scene  I32 F_XET
    /-show/showfile/cue/098/bit  I32 F_XET
    /-show/showfile/cue/098/miditype  I32 F_XET
    /-show/showfile/cue/098/midichan  I32 F_XET
    /-show/showfile/cue/098/midipara1  I32 F_XET
    /-show/showfile/cue/098/midipara2  I32 F_XET
/-show/showfile/cue/099  <SCUE> n=0
    /-show/showfile/cue/099/numb  I32 F_XET
    /-show/showfile/cue/099/name  S32 F_XET
    /-show/showfile/cue/099/skip  I32 F_XET
    /-show/showfile/cue/099/scene  I32 F_XET
    /-show/showfile/cue/099/bit  I32 F_XET
    /-show/showfile/cue/099/miditype  I32 F_XET
    /-show/showfile/cue/099/midichan  I32 F_XET
    /-show/showfile/cue/099/midipara1  I32 F_XET
    /-show/showfile/cue/099/midipara2  I32 F_XET
/-show/showfile/cue/100  <SCUE> n=0
    /-show/showfile/cue/100/numb  I32 F_XET
    /-show/showfile/cue/100/name  S32 F_XET
    /-show/showfile/cue/100/skip  I32 F_XET
    /-show/showfile/cue/100/scene  I32 F_XET
    /-show/showfile/cue/100/bit  I32 F_XET
    /-show/showfile/cue/100/miditype  I32 F_XET
    /-show/showfile/cue/100/midichan  I32 F_XET
    /-show/showfile/cue/100/midipara1  I32 F_XET
    /-show/showfile/cue/100/midipara2  I32 F_XET
/-show/showfile/cue/101  <SCUE> n=0
    /-show/showfile/cue/101/numb  I32 F_XET
    /-show/showfile/cue/101/name  S32 F_XET
    /-show/showfile/cue/101/skip  I32 F_XET
    /-show/showfile/cue/101/scene  I32 F_XET
    /-show/showfile/cue/101/bit  I32 F_XET
    /-show/showfile/cue/101/miditype  I32 F_XET
    /-show/showfile/cue/101/midichan  I32 F_XET
    /-show/showfile/cue/101/midipara1  I32 F_XET
    /-show/showfile/cue/101/midipara2  I32 F_XET
/-show/showfile/cue/102  <SCUE> n=0
    /-show/showfile/cue/102/numb  I32 F_XET
    /-show/showfile/cue/102/name  S32 F_XET
    /-show/showfile/cue/102/skip  I32 F_XET
    /-show/showfile/cue/102/scene  I32 F_XET
    /-show/showfile/cue/102/bit  I32 F_XET
    /-show/showfile/cue/102/miditype  I32 F_XET
    /-show/showfile/cue/102/midichan  I32 F_XET
    /-show/showfile/cue/102/midipara1  I32 F_XET
    /-show/showfile/cue/102/midipara2  I32 F_XET
/-show/showfile/cue/103  <SCUE> n=0
    /-show/showfile/cue/103/numb  I32 F_XET
    /-show/showfile/cue/103/name  S32 F_XET
    /-show/showfile/cue/103/skip  I32 F_XET
    /-show/showfile/cue/103/scene  I32 F_XET
    /-show/showfile/cue/103/bit  I32 F_XET
    /-show/showfile/cue/103/miditype  I32 F_XET
    /-show/showfile/cue/103/midichan  I32 F_XET
    /-show/showfile/cue/103/midipara1  I32 F_XET
    /-show/showfile/cue/103/midipara2  I32 F_XET
/-show/showfile/cue/104  <SCUE> n=0
    /-show/showfile/cue/104/numb  I32 F_XET
    /-show/showfile/cue/104/name  S32 F_XET
    /-show/showfile/cue/104/skip  I32 F_XET
    /-show/showfile/cue/104/scene  I32 F_XET
    /-show/showfile/cue/104/bit  I32 F_XET
    /-show/showfile/cue/104/miditype  I32 F_XET
    /-show/showfile/cue/104/midichan  I32 F_XET
    /-show/showfile/cue/104/midipara1  I32 F_XET
    /-show/showfile/cue/104/midipara2  I32 F_XET
/-show/showfile/cue/105  <SCUE> n=0
    /-show/showfile/cue/105/numb  I32 F_XET
    /-show/showfile/cue/105/name  S32 F_XET
    /-show/showfile/cue/105/skip  I32 F_XET
    /-show/showfile/cue/105/scene  I32 F_XET
    /-show/showfile/cue/105/bit  I32 F_XET
    /-show/showfile/cue/105/miditype  I32 F_XET
    /-show/showfile/cue/105/midichan  I32 F_XET
    /-show/showfile/cue/105/midipara1  I32 F_XET
    /-show/showfile/cue/105/midipara2  I32 F_XET
/-show/showfile/cue/106  <SCUE> n=0
    /-show/showfile/cue/106/numb  I32 F_XET
    /-show/showfile/cue/106/name  S32 F_XET
    /-show/showfile/cue/106/skip  I32 F_XET
    /-show/showfile/cue/106/scene  I32 F_XET
    /-show/showfile/cue/106/bit  I32 F_XET
    /-show/showfile/cue/106/miditype  I32 F_XET
    /-show/showfile/cue/106/midichan  I32 F_XET
    /-show/showfile/cue/106/midipara1  I32 F_XET
    /-show/showfile/cue/106/midipara2  I32 F_XET
/-show/showfile/cue/107  <SCUE> n=0
    /-show/showfile/cue/107/numb  I32 F_XET
    /-show/showfile/cue/107/name  S32 F_XET
    /-show/showfile/cue/107/skip  I32 F_XET
    /-show/showfile/cue/107/scene  I32 F_XET
    /-show/showfile/cue/107/bit  I32 F_XET
    /-show/showfile/cue/107/miditype  I32 F_XET
    /-show/showfile/cue/107/midichan  I32 F_XET
    /-show/showfile/cue/107/midipara1  I32 F_XET
    /-show/showfile/cue/107/midipara2  I32 F_XET
/-show/showfile/cue/108  <SCUE> n=0
    /-show/showfile/cue/108/numb  I32 F_XET
    /-show/showfile/cue/108/name  S32 F_XET
    /-show/showfile/cue/108/skip  I32 F_XET
    /-show/showfile/cue/108/scene  I32 F_XET
    /-show/showfile/cue/108/bit  I32 F_XET
    /-show/showfile/cue/108/miditype  I32 F_XET
    /-show/showfile/cue/108/midichan  I32 F_XET
    /-show/showfile/cue/108/midipara1  I32 F_XET
    /-show/showfile/cue/108/midipara2  I32 F_XET
/-show/showfile/cue/109  <SCUE> n=0
    /-show/showfile/cue/109/numb  I32 F_XET
    /-show/showfile/cue/109/name  S32 F_XET
    /-show/showfile/cue/109/skip  I32 F_XET
    /-show/showfile/cue/109/scene  I32 F_XET
    /-show/showfile/cue/109/bit  I32 F_XET
    /-show/showfile/cue/109/miditype  I32 F_XET
    /-show/showfile/cue/109/midichan  I32 F_XET
    /-show/showfile/cue/109/midipara1  I32 F_XET
    /-show/showfile/cue/109/midipara2  I32 F_XET
/-show/showfile/cue/110  <SCUE> n=0
    /-show/showfile/cue/110/numb  I32 F_XET
    /-show/showfile/cue/110/name  S32 F_XET
    /-show/showfile/cue/110/skip  I32 F_XET
    /-show/showfile/cue/110/scene  I32 F_XET
    /-show/showfile/cue/110/bit  I32 F_XET
    /-show/showfile/cue/110/miditype  I32 F_XET
    /-show/showfile/cue/110/midichan  I32 F_XET
    /-show/showfile/cue/110/midipara1  I32 F_XET
    /-show/showfile/cue/110/midipara2  I32 F_XET
/-show/showfile/cue/111  <SCUE> n=0
    /-show/showfile/cue/111/numb  I32 F_XET
    /-show/showfile/cue/111/name  S32 F_XET
    /-show/showfile/cue/111/skip  I32 F_XET
    /-show/showfile/cue/111/scene  I32 F_XET
    /-show/showfile/cue/111/bit  I32 F_XET
    /-show/showfile/cue/111/miditype  I32 F_XET
    /-show/showfile/cue/111/midichan  I32 F_XET
    /-show/showfile/cue/111/midipara1  I32 F_XET
    /-show/showfile/cue/111/midipara2  I32 F_XET
/-show/showfile/cue/112  <SCUE> n=0
    /-show/showfile/cue/112/numb  I32 F_XET
    /-show/showfile/cue/112/name  S32 F_XET
    /-show/showfile/cue/112/skip  I32 F_XET
    /-show/showfile/cue/112/scene  I32 F_XET
    /-show/showfile/cue/112/bit  I32 F_XET
    /-show/showfile/cue/112/miditype  I32 F_XET
    /-show/showfile/cue/112/midichan  I32 F_XET
    /-show/showfile/cue/112/midipara1  I32 F_XET
    /-show/showfile/cue/112/midipara2  I32 F_XET
/-show/showfile/cue/113  <SCUE> n=0
    /-show/showfile/cue/113/numb  I32 F_XET
    /-show/showfile/cue/113/name  S32 F_XET
    /-show/showfile/cue/113/skip  I32 F_XET
    /-show/showfile/cue/113/scene  I32 F_XET
    /-show/showfile/cue/113/bit  I32 F_XET
    /-show/showfile/cue/113/miditype  I32 F_XET
    /-show/showfile/cue/113/midichan  I32 F_XET
    /-show/showfile/cue/113/midipara1  I32 F_XET
    /-show/showfile/cue/113/midipara2  I32 F_XET
/-show/showfile/cue/114  <SCUE> n=0
    /-show/showfile/cue/114/numb  I32 F_XET
    /-show/showfile/cue/114/name  S32 F_XET
    /-show/showfile/cue/114/skip  I32 F_XET
    /-show/showfile/cue/114/scene  I32 F_XET
    /-show/showfile/cue/114/bit  I32 F_XET
    /-show/showfile/cue/114/miditype  I32 F_XET
    /-show/showfile/cue/114/midichan  I32 F_XET
    /-show/showfile/cue/114/midipara1  I32 F_XET
    /-show/showfile/cue/114/midipara2  I32 F_XET
/-show/showfile/cue/115  <SCUE> n=0
    /-show/showfile/cue/115/numb  I32 F_XET
    /-show/showfile/cue/115/name  S32 F_XET
    /-show/showfile/cue/115/skip  I32 F_XET
    /-show/showfile/cue/115/scene  I32 F_XET
    /-show/showfile/cue/115/bit  I32 F_XET
    /-show/showfile/cue/115/miditype  I32 F_XET
    /-show/showfile/cue/115/midichan  I32 F_XET
    /-show/showfile/cue/115/midipara1  I32 F_XET
    /-show/showfile/cue/115/midipara2  I32 F_XET
/-show/showfile/cue/116  <SCUE> n=0
    /-show/showfile/cue/116/numb  I32 F_XET
    /-show/showfile/cue/116/name  S32 F_XET
    /-show/showfile/cue/116/skip  I32 F_XET
    /-show/showfile/cue/116/scene  I32 F_XET
    /-show/showfile/cue/116/bit  I32 F_XET
    /-show/showfile/cue/116/miditype  I32 F_XET
    /-show/showfile/cue/116/midichan  I32 F_XET
    /-show/showfile/cue/116/midipara1  I32 F_XET
    /-show/showfile/cue/116/midipara2  I32 F_XET
/-show/showfile/cue/117  <SCUE> n=0
    /-show/showfile/cue/117/numb  I32 F_XET
    /-show/showfile/cue/117/name  S32 F_XET
    /-show/showfile/cue/117/skip  I32 F_XET
    /-show/showfile/cue/117/scene  I32 F_XET
    /-show/showfile/cue/117/bit  I32 F_XET
    /-show/showfile/cue/117/miditype  I32 F_XET
    /-show/showfile/cue/117/midichan  I32 F_XET
    /-show/showfile/cue/117/midipara1  I32 F_XET
    /-show/showfile/cue/117/midipara2  I32 F_XET
/-show/showfile/cue/118  <SCUE> n=0
    /-show/showfile/cue/118/numb  I32 F_XET
    /-show/showfile/cue/118/name  S32 F_XET
    /-show/showfile/cue/118/skip  I32 F_XET
    /-show/showfile/cue/118/scene  I32 F_XET
    /-show/showfile/cue/118/bit  I32 F_XET
    /-show/showfile/cue/118/miditype  I32 F_XET
    /-show/showfile/cue/118/midichan  I32 F_XET
    /-show/showfile/cue/118/midipara1  I32 F_XET
    /-show/showfile/cue/118/midipara2  I32 F_XET
/-show/showfile/cue/119  <SCUE> n=0
    /-show/showfile/cue/119/numb  I32 F_XET
    /-show/showfile/cue/119/name  S32 F_XET
    /-show/showfile/cue/119/skip  I32 F_XET
    /-show/showfile/cue/119/scene  I32 F_XET
    /-show/showfile/cue/119/bit  I32 F_XET
    /-show/showfile/cue/119/miditype  I32 F_XET
    /-show/showfile/cue/119/midichan  I32 F_XET
    /-show/showfile/cue/119/midipara1  I32 F_XET
    /-show/showfile/cue/119/midipara2  I32 F_XET
/-show/showfile/cue/120  <SCUE> n=0
    /-show/showfile/cue/120/numb  I32 F_XET
    /-show/showfile/cue/120/name  S32 F_XET
    /-show/showfile/cue/120/skip  I32 F_XET
    /-show/showfile/cue/120/scene  I32 F_XET
    /-show/showfile/cue/120/bit  I32 F_XET
    /-show/showfile/cue/120/miditype  I32 F_XET
    /-show/showfile/cue/120/midichan  I32 F_XET
    /-show/showfile/cue/120/midipara1  I32 F_XET
    /-show/showfile/cue/120/midipara2  I32 F_XET
/-show/showfile/cue/121  <SCUE> n=0
    /-show/showfile/cue/121/numb  I32 F_XET
    /-show/showfile/cue/121/name  S32 F_XET
    /-show/showfile/cue/121/skip  I32 F_XET
    /-show/showfile/cue/121/scene  I32 F_XET
    /-show/showfile/cue/121/bit  I32 F_XET
    /-show/showfile/cue/121/miditype  I32 F_XET
    /-show/showfile/cue/121/midichan  I32 F_XET
    /-show/showfile/cue/121/midipara1  I32 F_XET
    /-show/showfile/cue/121/midipara2  I32 F_XET
/-show/showfile/cue/122  <SCUE> n=0
    /-show/showfile/cue/122/numb  I32 F_XET
    /-show/showfile/cue/122/name  S32 F_XET
    /-show/showfile/cue/122/skip  I32 F_XET
    /-show/showfile/cue/122/scene  I32 F_XET
    /-show/showfile/cue/122/bit  I32 F_XET
    /-show/showfile/cue/122/miditype  I32 F_XET
    /-show/showfile/cue/122/midichan  I32 F_XET
    /-show/showfile/cue/122/midipara1  I32 F_XET
    /-show/showfile/cue/122/midipara2  I32 F_XET
/-show/showfile/cue/123  <SCUE> n=0
    /-show/showfile/cue/123/numb  I32 F_XET
    /-show/showfile/cue/123/name  S32 F_XET
    /-show/showfile/cue/123/skip  I32 F_XET
    /-show/showfile/cue/123/scene  I32 F_XET
    /-show/showfile/cue/123/bit  I32 F_XET
    /-show/showfile/cue/123/miditype  I32 F_XET
    /-show/showfile/cue/123/midichan  I32 F_XET
    /-show/showfile/cue/123/midipara1  I32 F_XET
    /-show/showfile/cue/123/midipara2  I32 F_XET
/-show/showfile/cue/124  <SCUE> n=0
    /-show/showfile/cue/124/numb  I32 F_XET
    /-show/showfile/cue/124/name  S32 F_XET
    /-show/showfile/cue/124/skip  I32 F_XET
    /-show/showfile/cue/124/scene  I32 F_XET
    /-show/showfile/cue/124/bit  I32 F_XET
    /-show/showfile/cue/124/miditype  I32 F_XET
    /-show/showfile/cue/124/midichan  I32 F_XET
    /-show/showfile/cue/124/midipara1  I32 F_XET
    /-show/showfile/cue/124/midipara2  I32 F_XET
/-show/showfile/cue/125  <SCUE> n=0
    /-show/showfile/cue/125/numb  I32 F_XET
    /-show/showfile/cue/125/name  S32 F_XET
    /-show/showfile/cue/125/skip  I32 F_XET
    /-show/showfile/cue/125/scene  I32 F_XET
    /-show/showfile/cue/125/bit  I32 F_XET
    /-show/showfile/cue/125/miditype  I32 F_XET
    /-show/showfile/cue/125/midichan  I32 F_XET
    /-show/showfile/cue/125/midipara1  I32 F_XET
    /-show/showfile/cue/125/midipara2  I32 F_XET
/-show/showfile/cue/126  <SCUE> n=0
    /-show/showfile/cue/126/numb  I32 F_XET
    /-show/showfile/cue/126/name  S32 F_XET
    /-show/showfile/cue/126/skip  I32 F_XET
    /-show/showfile/cue/126/scene  I32 F_XET
    /-show/showfile/cue/126/bit  I32 F_XET
    /-show/showfile/cue/126/miditype  I32 F_XET
    /-show/showfile/cue/126/midichan  I32 F_XET
    /-show/showfile/cue/126/midipara1  I32 F_XET
    /-show/showfile/cue/126/midipara2  I32 F_XET
/-show/showfile/cue/127  <SCUE> n=0
    /-show/showfile/cue/127/numb  I32 F_XET
    /-show/showfile/cue/127/name  S32 F_XET
    /-show/showfile/cue/127/skip  I32 F_XET
    /-show/showfile/cue/127/scene  I32 F_XET
    /-show/showfile/cue/127/bit  I32 F_XET
    /-show/showfile/cue/127/miditype  I32 F_XET
    /-show/showfile/cue/127/midichan  I32 F_XET
    /-show/showfile/cue/127/midipara1  I32 F_XET
    /-show/showfile/cue/127/midipara2  I32 F_XET
/-show/showfile/cue/128  <SCUE> n=0
    /-show/showfile/cue/128/numb  I32 F_XET
    /-show/showfile/cue/128/name  S32 F_XET
    /-show/showfile/cue/128/skip  I32 F_XET
    /-show/showfile/cue/128/scene  I32 F_XET
    /-show/showfile/cue/128/bit  I32 F_XET
    /-show/showfile/cue/128/miditype  I32 F_XET
    /-show/showfile/cue/128/midichan  I32 F_XET
    /-show/showfile/cue/128/midipara1  I32 F_XET
    /-show/showfile/cue/128/midipara2  I32 F_XET
/-show/showfile/cue/129  <SCUE> n=0
    /-show/showfile/cue/129/numb  I32 F_XET
    /-show/showfile/cue/129/name  S32 F_XET
    /-show/showfile/cue/129/skip  I32 F_XET
    /-show/showfile/cue/129/scene  I32 F_XET
    /-show/showfile/cue/129/bit  I32 F_XET
    /-show/showfile/cue/129/miditype  I32 F_XET
    /-show/showfile/cue/129/midichan  I32 F_XET
    /-show/showfile/cue/129/midipara1  I32 F_XET
    /-show/showfile/cue/129/midipara2  I32 F_XET
/-show/showfile/cue/130  <SCUE> n=0
    /-show/showfile/cue/130/numb  I32 F_XET
    /-show/showfile/cue/130/name  S32 F_XET
    /-show/showfile/cue/130/skip  I32 F_XET
    /-show/showfile/cue/130/scene  I32 F_XET
    /-show/showfile/cue/130/bit  I32 F_XET
    /-show/showfile/cue/130/miditype  I32 F_XET
    /-show/showfile/cue/130/midichan  I32 F_XET
    /-show/showfile/cue/130/midipara1  I32 F_XET
    /-show/showfile/cue/130/midipara2  I32 F_XET
/-show/showfile/cue/131  <SCUE> n=0
    /-show/showfile/cue/131/numb  I32 F_XET
    /-show/showfile/cue/131/name  S32 F_XET
    /-show/showfile/cue/131/skip  I32 F_XET
    /-show/showfile/cue/131/scene  I32 F_XET
    /-show/showfile/cue/131/bit  I32 F_XET
    /-show/showfile/cue/131/miditype  I32 F_XET
    /-show/showfile/cue/131/midichan  I32 F_XET
    /-show/showfile/cue/131/midipara1  I32 F_XET
    /-show/showfile/cue/131/midipara2  I32 F_XET
/-show/showfile/cue/132  <SCUE> n=0
    /-show/showfile/cue/132/numb  I32 F_XET
    /-show/showfile/cue/132/name  S32 F_XET
    /-show/showfile/cue/132/skip  I32 F_XET
    /-show/showfile/cue/132/scene  I32 F_XET
    /-show/showfile/cue/132/bit  I32 F_XET
    /-show/showfile/cue/132/miditype  I32 F_XET
    /-show/showfile/cue/132/midichan  I32 F_XET
    /-show/showfile/cue/132/midipara1  I32 F_XET
    /-show/showfile/cue/132/midipara2  I32 F_XET
/-show/showfile/cue/133  <SCUE> n=0
    /-show/showfile/cue/133/numb  I32 F_XET
    /-show/showfile/cue/133/name  S32 F_XET
    /-show/showfile/cue/133/skip  I32 F_XET
    /-show/showfile/cue/133/scene  I32 F_XET
    /-show/showfile/cue/133/bit  I32 F_XET
    /-show/showfile/cue/133/miditype  I32 F_XET
    /-show/showfile/cue/133/midichan  I32 F_XET
    /-show/showfile/cue/133/midipara1  I32 F_XET
    /-show/showfile/cue/133/midipara2  I32 F_XET
/-show/showfile/cue/134  <SCUE> n=0
    /-show/showfile/cue/134/numb  I32 F_XET
    /-show/showfile/cue/134/name  S32 F_XET
    /-show/showfile/cue/134/skip  I32 F_XET
    /-show/showfile/cue/134/scene  I32 F_XET
    /-show/showfile/cue/134/bit  I32 F_XET
    /-show/showfile/cue/134/miditype  I32 F_XET
    /-show/showfile/cue/134/midichan  I32 F_XET
    /-show/showfile/cue/134/midipara1  I32 F_XET
    /-show/showfile/cue/134/midipara2  I32 F_XET
/-show/showfile/cue/135  <SCUE> n=0
    /-show/showfile/cue/135/numb  I32 F_XET
    /-show/showfile/cue/135/name  S32 F_XET
    /-show/showfile/cue/135/skip  I32 F_XET
    /-show/showfile/cue/135/scene  I32 F_XET
    /-show/showfile/cue/135/bit  I32 F_XET
    /-show/showfile/cue/135/miditype  I32 F_XET
    /-show/showfile/cue/135/midichan  I32 F_XET
    /-show/showfile/cue/135/midipara1  I32 F_XET
    /-show/showfile/cue/135/midipara2  I32 F_XET
/-show/showfile/cue/136  <SCUE> n=0
    /-show/showfile/cue/136/numb  I32 F_XET
    /-show/showfile/cue/136/name  S32 F_XET
    /-show/showfile/cue/136/skip  I32 F_XET
    /-show/showfile/cue/136/scene  I32 F_XET
    /-show/showfile/cue/136/bit  I32 F_XET
    /-show/showfile/cue/136/miditype  I32 F_XET
    /-show/showfile/cue/136/midichan  I32 F_XET
    /-show/showfile/cue/136/midipara1  I32 F_XET
    /-show/showfile/cue/136/midipara2  I32 F_XET
/-show/showfile/cue/137  <SCUE> n=0
    /-show/showfile/cue/137/numb  I32 F_XET
    /-show/showfile/cue/137/name  S32 F_XET
    /-show/showfile/cue/137/skip  I32 F_XET
    /-show/showfile/cue/137/scene  I32 F_XET
    /-show/showfile/cue/137/bit  I32 F_XET
    /-show/showfile/cue/137/miditype  I32 F_XET
    /-show/showfile/cue/137/midichan  I32 F_XET
    /-show/showfile/cue/137/midipara1  I32 F_XET
    /-show/showfile/cue/137/midipara2  I32 F_XET
/-show/showfile/cue/138  <SCUE> n=0
    /-show/showfile/cue/138/numb  I32 F_XET
    /-show/showfile/cue/138/name  S32 F_XET
    /-show/showfile/cue/138/skip  I32 F_XET
    /-show/showfile/cue/138/scene  I32 F_XET
    /-show/showfile/cue/138/bit  I32 F_XET
    /-show/showfile/cue/138/miditype  I32 F_XET
    /-show/showfile/cue/138/midichan  I32 F_XET
    /-show/showfile/cue/138/midipara1  I32 F_XET
    /-show/showfile/cue/138/midipara2  I32 F_XET
/-show/showfile/cue/139  <SCUE> n=0
    /-show/showfile/cue/139/numb  I32 F_XET
    /-show/showfile/cue/139/name  S32 F_XET
    /-show/showfile/cue/139/skip  I32 F_XET
    /-show/showfile/cue/139/scene  I32 F_XET
    /-show/showfile/cue/139/bit  I32 F_XET
    /-show/showfile/cue/139/miditype  I32 F_XET
    /-show/showfile/cue/139/midichan  I32 F_XET
    /-show/showfile/cue/139/midipara1  I32 F_XET
    /-show/showfile/cue/139/midipara2  I32 F_XET
/-show/showfile/cue/140  <SCUE> n=0
    /-show/showfile/cue/140/numb  I32 F_XET
    /-show/showfile/cue/140/name  S32 F_XET
    /-show/showfile/cue/140/skip  I32 F_XET
    /-show/showfile/cue/140/scene  I32 F_XET
    /-show/showfile/cue/140/bit  I32 F_XET
    /-show/showfile/cue/140/miditype  I32 F_XET
    /-show/showfile/cue/140/midichan  I32 F_XET
    /-show/showfile/cue/140/midipara1  I32 F_XET
    /-show/showfile/cue/140/midipara2  I32 F_XET
/-show/showfile/cue/141  <SCUE> n=0
    /-show/showfile/cue/141/numb  I32 F_XET
    /-show/showfile/cue/141/name  S32 F_XET
    /-show/showfile/cue/141/skip  I32 F_XET
    /-show/showfile/cue/141/scene  I32 F_XET
    /-show/showfile/cue/141/bit  I32 F_XET
    /-show/showfile/cue/141/miditype  I32 F_XET
    /-show/showfile/cue/141/midichan  I32 F_XET
    /-show/showfile/cue/141/midipara1  I32 F_XET
    /-show/showfile/cue/141/midipara2  I32 F_XET
/-show/showfile/cue/142  <SCUE> n=0
    /-show/showfile/cue/142/numb  I32 F_XET
    /-show/showfile/cue/142/name  S32 F_XET
    /-show/showfile/cue/142/skip  I32 F_XET
    /-show/showfile/cue/142/scene  I32 F_XET
    /-show/showfile/cue/142/bit  I32 F_XET
    /-show/showfile/cue/142/miditype  I32 F_XET
    /-show/showfile/cue/142/midichan  I32 F_XET
    /-show/showfile/cue/142/midipara1  I32 F_XET
    /-show/showfile/cue/142/midipara2  I32 F_XET
/-show/showfile/cue/143  <SCUE> n=0
    /-show/showfile/cue/143/numb  I32 F_XET
    /-show/showfile/cue/143/name  S32 F_XET
    /-show/showfile/cue/143/skip  I32 F_XET
    /-show/showfile/cue/143/scene  I32 F_XET
    /-show/showfile/cue/143/bit  I32 F_XET
    /-show/showfile/cue/143/miditype  I32 F_XET
    /-show/showfile/cue/143/midichan  I32 F_XET
    /-show/showfile/cue/143/midipara1  I32 F_XET
    /-show/showfile/cue/143/midipara2  I32 F_XET
/-show/showfile/cue/144  <SCUE> n=0
    /-show/showfile/cue/144/numb  I32 F_XET
    /-show/showfile/cue/144/name  S32 F_XET
    /-show/showfile/cue/144/skip  I32 F_XET
    /-show/showfile/cue/144/scene  I32 F_XET
    /-show/showfile/cue/144/bit  I32 F_XET
    /-show/showfile/cue/144/miditype  I32 F_XET
    /-show/showfile/cue/144/midichan  I32 F_XET
    /-show/showfile/cue/144/midipara1  I32 F_XET
    /-show/showfile/cue/144/midipara2  I32 F_XET
/-show/showfile/cue/145  <SCUE> n=0
    /-show/showfile/cue/145/numb  I32 F_XET
    /-show/showfile/cue/145/name  S32 F_XET
    /-show/showfile/cue/145/skip  I32 F_XET
    /-show/showfile/cue/145/scene  I32 F_XET
    /-show/showfile/cue/145/bit  I32 F_XET
    /-show/showfile/cue/145/miditype  I32 F_XET
    /-show/showfile/cue/145/midichan  I32 F_XET
    /-show/showfile/cue/145/midipara1  I32 F_XET
    /-show/showfile/cue/145/midipara2  I32 F_XET
/-show/showfile/cue/146  <SCUE> n=0
    /-show/showfile/cue/146/numb  I32 F_XET
    /-show/showfile/cue/146/name  S32 F_XET
    /-show/showfile/cue/146/skip  I32 F_XET
    /-show/showfile/cue/146/scene  I32 F_XET
    /-show/showfile/cue/146/bit  I32 F_XET
    /-show/showfile/cue/146/miditype  I32 F_XET
    /-show/showfile/cue/146/midichan  I32 F_XET
    /-show/showfile/cue/146/midipara1  I32 F_XET
    /-show/showfile/cue/146/midipara2  I32 F_XET
/-show/showfile/cue/147  <SCUE> n=0
    /-show/showfile/cue/147/numb  I32 F_XET
    /-show/showfile/cue/147/name  S32 F_XET
    /-show/showfile/cue/147/skip  I32 F_XET
    /-show/showfile/cue/147/scene  I32 F_XET
    /-show/showfile/cue/147/bit  I32 F_XET
    /-show/showfile/cue/147/miditype  I32 F_XET
    /-show/showfile/cue/147/midichan  I32 F_XET
    /-show/showfile/cue/147/midipara1  I32 F_XET
    /-show/showfile/cue/147/midipara2  I32 F_XET
/-show/showfile/cue/148  <SCUE> n=0
    /-show/showfile/cue/148/numb  I32 F_XET
    /-show/showfile/cue/148/name  S32 F_XET
    /-show/showfile/cue/148/skip  I32 F_XET
    /-show/showfile/cue/148/scene  I32 F_XET
    /-show/showfile/cue/148/bit  I32 F_XET
    /-show/showfile/cue/148/miditype  I32 F_XET
    /-show/showfile/cue/148/midichan  I32 F_XET
    /-show/showfile/cue/148/midipara1  I32 F_XET
    /-show/showfile/cue/148/midipara2  I32 F_XET
/-show/showfile/cue/149  <SCUE> n=0
    /-show/showfile/cue/149/numb  I32 F_XET
    /-show/showfile/cue/149/name  S32 F_XET
    /-show/showfile/cue/149/skip  I32 F_XET
    /-show/showfile/cue/149/scene  I32 F_XET
    /-show/showfile/cue/149/bit  I32 F_XET
    /-show/showfile/cue/149/miditype  I32 F_XET
    /-show/showfile/cue/149/midichan  I32 F_XET
    /-show/showfile/cue/149/midipara1  I32 F_XET
    /-show/showfile/cue/149/midipara2  I32 F_XET
/-show/showfile/cue/150  <SCUE> n=0
    /-show/showfile/cue/150/numb  I32 F_XET
    /-show/showfile/cue/150/name  S32 F_XET
    /-show/showfile/cue/150/skip  I32 F_XET
    /-show/showfile/cue/150/scene  I32 F_XET
    /-show/showfile/cue/150/bit  I32 F_XET
    /-show/showfile/cue/150/miditype  I32 F_XET
    /-show/showfile/cue/150/midichan  I32 F_XET
    /-show/showfile/cue/150/midipara1  I32 F_XET
    /-show/showfile/cue/150/midipara2  I32 F_XET
/-show/showfile/cue/151  <SCUE> n=0
    /-show/showfile/cue/151/numb  I32 F_XET
    /-show/showfile/cue/151/name  S32 F_XET
    /-show/showfile/cue/151/skip  I32 F_XET
    /-show/showfile/cue/151/scene  I32 F_XET
    /-show/showfile/cue/151/bit  I32 F_XET
    /-show/showfile/cue/151/miditype  I32 F_XET
    /-show/showfile/cue/151/midichan  I32 F_XET
    /-show/showfile/cue/151/midipara1  I32 F_XET
    /-show/showfile/cue/151/midipara2  I32 F_XET
/-show/showfile/cue/152  <SCUE> n=0
    /-show/showfile/cue/152/numb  I32 F_XET
    /-show/showfile/cue/152/name  S32 F_XET
    /-show/showfile/cue/152/skip  I32 F_XET
    /-show/showfile/cue/152/scene  I32 F_XET
    /-show/showfile/cue/152/bit  I32 F_XET
    /-show/showfile/cue/152/miditype  I32 F_XET
    /-show/showfile/cue/152/midichan  I32 F_XET
    /-show/showfile/cue/152/midipara1  I32 F_XET
    /-show/showfile/cue/152/midipara2  I32 F_XET
/-show/showfile/cue/153  <SCUE> n=0
    /-show/showfile/cue/153/numb  I32 F_XET
    /-show/showfile/cue/153/name  S32 F_XET
    /-show/showfile/cue/153/skip  I32 F_XET
    /-show/showfile/cue/153/scene  I32 F_XET
    /-show/showfile/cue/153/bit  I32 F_XET
    /-show/showfile/cue/153/miditype  I32 F_XET
    /-show/showfile/cue/153/midichan  I32 F_XET
    /-show/showfile/cue/153/midipara1  I32 F_XET
    /-show/showfile/cue/153/midipara2  I32 F_XET
/-show/showfile/cue/154  <SCUE> n=0
    /-show/showfile/cue/154/numb  I32 F_XET
    /-show/showfile/cue/154/name  S32 F_XET
    /-show/showfile/cue/154/skip  I32 F_XET
    /-show/showfile/cue/154/scene  I32 F_XET
    /-show/showfile/cue/154/bit  I32 F_XET
    /-show/showfile/cue/154/miditype  I32 F_XET
    /-show/showfile/cue/154/midichan  I32 F_XET
    /-show/showfile/cue/154/midipara1  I32 F_XET
    /-show/showfile/cue/154/midipara2  I32 F_XET
/-show/showfile/cue/155  <SCUE> n=0
    /-show/showfile/cue/155/numb  I32 F_XET
    /-show/showfile/cue/155/name  S32 F_XET
    /-show/showfile/cue/155/skip  I32 F_XET
    /-show/showfile/cue/155/scene  I32 F_XET
    /-show/showfile/cue/155/bit  I32 F_XET
    /-show/showfile/cue/155/miditype  I32 F_XET
    /-show/showfile/cue/155/midichan  I32 F_XET
    /-show/showfile/cue/155/midipara1  I32 F_XET
    /-show/showfile/cue/155/midipara2  I32 F_XET
/-show/showfile/cue/156  <SCUE> n=0
    /-show/showfile/cue/156/numb  I32 F_XET
    /-show/showfile/cue/156/name  S32 F_XET
    /-show/showfile/cue/156/skip  I32 F_XET
    /-show/showfile/cue/156/scene  I32 F_XET
    /-show/showfile/cue/156/bit  I32 F_XET
    /-show/showfile/cue/156/miditype  I32 F_XET
    /-show/showfile/cue/156/midichan  I32 F_XET
    /-show/showfile/cue/156/midipara1  I32 F_XET
    /-show/showfile/cue/156/midipara2  I32 F_XET
/-show/showfile/cue/157  <SCUE> n=0
    /-show/showfile/cue/157/numb  I32 F_XET
    /-show/showfile/cue/157/name  S32 F_XET
    /-show/showfile/cue/157/skip  I32 F_XET
    /-show/showfile/cue/157/scene  I32 F_XET
    /-show/showfile/cue/157/bit  I32 F_XET
    /-show/showfile/cue/157/miditype  I32 F_XET
    /-show/showfile/cue/157/midichan  I32 F_XET
    /-show/showfile/cue/157/midipara1  I32 F_XET
    /-show/showfile/cue/157/midipara2  I32 F_XET
/-show/showfile/cue/158  <SCUE> n=0
    /-show/showfile/cue/158/numb  I32 F_XET
    /-show/showfile/cue/158/name  S32 F_XET
    /-show/showfile/cue/158/skip  I32 F_XET
    /-show/showfile/cue/158/scene  I32 F_XET
    /-show/showfile/cue/158/bit  I32 F_XET
    /-show/showfile/cue/158/miditype  I32 F_XET
    /-show/showfile/cue/158/midichan  I32 F_XET
    /-show/showfile/cue/158/midipara1  I32 F_XET
    /-show/showfile/cue/158/midipara2  I32 F_XET
/-show/showfile/cue/159  <SCUE> n=0
    /-show/showfile/cue/159/numb  I32 F_XET
    /-show/showfile/cue/159/name  S32 F_XET
    /-show/showfile/cue/159/skip  I32 F_XET
    /-show/showfile/cue/159/scene  I32 F_XET
    /-show/showfile/cue/159/bit  I32 F_XET
    /-show/showfile/cue/159/miditype  I32 F_XET
    /-show/showfile/cue/159/midichan  I32 F_XET
    /-show/showfile/cue/159/midipara1  I32 F_XET
    /-show/showfile/cue/159/midipara2  I32 F_XET
/-show/showfile/cue/160  <SCUE> n=0
    /-show/showfile/cue/160/numb  I32 F_XET
    /-show/showfile/cue/160/name  S32 F_XET
    /-show/showfile/cue/160/skip  I32 F_XET
    /-show/showfile/cue/160/scene  I32 F_XET
    /-show/showfile/cue/160/bit  I32 F_XET
    /-show/showfile/cue/160/miditype  I32 F_XET
    /-show/showfile/cue/160/midichan  I32 F_XET
    /-show/showfile/cue/160/midipara1  I32 F_XET
    /-show/showfile/cue/160/midipara2  I32 F_XET
/-show/showfile/cue/161  <SCUE> n=0
    /-show/showfile/cue/161/numb  I32 F_XET
    /-show/showfile/cue/161/name  S32 F_XET
    /-show/showfile/cue/161/skip  I32 F_XET
    /-show/showfile/cue/161/scene  I32 F_XET
    /-show/showfile/cue/161/bit  I32 F_XET
    /-show/showfile/cue/161/miditype  I32 F_XET
    /-show/showfile/cue/161/midichan  I32 F_XET
    /-show/showfile/cue/161/midipara1  I32 F_XET
    /-show/showfile/cue/161/midipara2  I32 F_XET
/-show/showfile/cue/162  <SCUE> n=0
    /-show/showfile/cue/162/numb  I32 F_XET
    /-show/showfile/cue/162/name  S32 F_XET
    /-show/showfile/cue/162/skip  I32 F_XET
    /-show/showfile/cue/162/scene  I32 F_XET
    /-show/showfile/cue/162/bit  I32 F_XET
    /-show/showfile/cue/162/miditype  I32 F_XET
    /-show/showfile/cue/162/midichan  I32 F_XET
    /-show/showfile/cue/162/midipara1  I32 F_XET
    /-show/showfile/cue/162/midipara2  I32 F_XET
/-show/showfile/cue/163  <SCUE> n=0
    /-show/showfile/cue/163/numb  I32 F_XET
    /-show/showfile/cue/163/name  S32 F_XET
    /-show/showfile/cue/163/skip  I32 F_XET
    /-show/showfile/cue/163/scene  I32 F_XET
    /-show/showfile/cue/163/bit  I32 F_XET
    /-show/showfile/cue/163/miditype  I32 F_XET
    /-show/showfile/cue/163/midichan  I32 F_XET
    /-show/showfile/cue/163/midipara1  I32 F_XET
    /-show/showfile/cue/163/midipara2  I32 F_XET
/-show/showfile/cue/164  <SCUE> n=0
    /-show/showfile/cue/164/numb  I32 F_XET
    /-show/showfile/cue/164/name  S32 F_XET
    /-show/showfile/cue/164/skip  I32 F_XET
    /-show/showfile/cue/164/scene  I32 F_XET
    /-show/showfile/cue/164/bit  I32 F_XET
    /-show/showfile/cue/164/miditype  I32 F_XET
    /-show/showfile/cue/164/midichan  I32 F_XET
    /-show/showfile/cue/164/midipara1  I32 F_XET
    /-show/showfile/cue/164/midipara2  I32 F_XET
/-show/showfile/cue/165  <SCUE> n=0
    /-show/showfile/cue/165/numb  I32 F_XET
    /-show/showfile/cue/165/name  S32 F_XET
    /-show/showfile/cue/165/skip  I32 F_XET
    /-show/showfile/cue/165/scene  I32 F_XET
    /-show/showfile/cue/165/bit  I32 F_XET
    /-show/showfile/cue/165/miditype  I32 F_XET
    /-show/showfile/cue/165/midichan  I32 F_XET
    /-show/showfile/cue/165/midipara1  I32 F_XET
    /-show/showfile/cue/165/midipara2  I32 F_XET
/-show/showfile/cue/166  <SCUE> n=0
    /-show/showfile/cue/166/numb  I32 F_XET
    /-show/showfile/cue/166/name  S32 F_XET
    /-show/showfile/cue/166/skip  I32 F_XET
    /-show/showfile/cue/166/scene  I32 F_XET
    /-show/showfile/cue/166/bit  I32 F_XET
    /-show/showfile/cue/166/miditype  I32 F_XET
    /-show/showfile/cue/166/midichan  I32 F_XET
    /-show/showfile/cue/166/midipara1  I32 F_XET
    /-show/showfile/cue/166/midipara2  I32 F_XET
/-show/showfile/cue/167  <SCUE> n=0
    /-show/showfile/cue/167/numb  I32 F_XET
    /-show/showfile/cue/167/name  S32 F_XET
    /-show/showfile/cue/167/skip  I32 F_XET
    /-show/showfile/cue/167/scene  I32 F_XET
    /-show/showfile/cue/167/bit  I32 F_XET
    /-show/showfile/cue/167/miditype  I32 F_XET
    /-show/showfile/cue/167/midichan  I32 F_XET
    /-show/showfile/cue/167/midipara1  I32 F_XET
    /-show/showfile/cue/167/midipara2  I32 F_XET
/-show/showfile/cue/168  <SCUE> n=0
    /-show/showfile/cue/168/numb  I32 F_XET
    /-show/showfile/cue/168/name  S32 F_XET
    /-show/showfile/cue/168/skip  I32 F_XET
    /-show/showfile/cue/168/scene  I32 F_XET
    /-show/showfile/cue/168/bit  I32 F_XET
    /-show/showfile/cue/168/miditype  I32 F_XET
    /-show/showfile/cue/168/midichan  I32 F_XET
    /-show/showfile/cue/168/midipara1  I32 F_XET
    /-show/showfile/cue/168/midipara2  I32 F_XET
/-show/showfile/cue/169  <SCUE> n=0
    /-show/showfile/cue/169/numb  I32 F_XET
    /-show/showfile/cue/169/name  S32 F_XET
    /-show/showfile/cue/169/skip  I32 F_XET
    /-show/showfile/cue/169/scene  I32 F_XET
    /-show/showfile/cue/169/bit  I32 F_XET
    /-show/showfile/cue/169/miditype  I32 F_XET
    /-show/showfile/cue/169/midichan  I32 F_XET
    /-show/showfile/cue/169/midipara1  I32 F_XET
    /-show/showfile/cue/169/midipara2  I32 F_XET
/-show/showfile/cue/170  <SCUE> n=0
    /-show/showfile/cue/170/numb  I32 F_XET
    /-show/showfile/cue/170/name  S32 F_XET
    /-show/showfile/cue/170/skip  I32 F_XET
    /-show/showfile/cue/170/scene  I32 F_XET
    /-show/showfile/cue/170/bit  I32 F_XET
    /-show/showfile/cue/170/miditype  I32 F_XET
    /-show/showfile/cue/170/midichan  I32 F_XET
    /-show/showfile/cue/170/midipara1  I32 F_XET
    /-show/showfile/cue/170/midipara2  I32 F_XET
/-show/showfile/cue/171  <SCUE> n=0
    /-show/showfile/cue/171/numb  I32 F_XET
    /-show/showfile/cue/171/name  S32 F_XET
    /-show/showfile/cue/171/skip  I32 F_XET
    /-show/showfile/cue/171/scene  I32 F_XET
    /-show/showfile/cue/171/bit  I32 F_XET
    /-show/showfile/cue/171/miditype  I32 F_XET
    /-show/showfile/cue/171/midichan  I32 F_XET
    /-show/showfile/cue/171/midipara1  I32 F_XET
    /-show/showfile/cue/171/midipara2  I32 F_XET
/-show/showfile/cue/172  <SCUE> n=0
    /-show/showfile/cue/172/numb  I32 F_XET
    /-show/showfile/cue/172/name  S32 F_XET
    /-show/showfile/cue/172/skip  I32 F_XET
    /-show/showfile/cue/172/scene  I32 F_XET
    /-show/showfile/cue/172/bit  I32 F_XET
    /-show/showfile/cue/172/miditype  I32 F_XET
    /-show/showfile/cue/172/midichan  I32 F_XET
    /-show/showfile/cue/172/midipara1  I32 F_XET
    /-show/showfile/cue/172/midipara2  I32 F_XET
/-show/showfile/cue/173  <SCUE> n=0
    /-show/showfile/cue/173/numb  I32 F_XET
    /-show/showfile/cue/173/name  S32 F_XET
    /-show/showfile/cue/173/skip  I32 F_XET
    /-show/showfile/cue/173/scene  I32 F_XET
    /-show/showfile/cue/173/bit  I32 F_XET
    /-show/showfile/cue/173/miditype  I32 F_XET
    /-show/showfile/cue/173/midichan  I32 F_XET
    /-show/showfile/cue/173/midipara1  I32 F_XET
    /-show/showfile/cue/173/midipara2  I32 F_XET
/-show/showfile/cue/174  <SCUE> n=0
    /-show/showfile/cue/174/numb  I32 F_XET
    /-show/showfile/cue/174/name  S32 F_XET
    /-show/showfile/cue/174/skip  I32 F_XET
    /-show/showfile/cue/174/scene  I32 F_XET
    /-show/showfile/cue/174/bit  I32 F_XET
    /-show/showfile/cue/174/miditype  I32 F_XET
    /-show/showfile/cue/174/midichan  I32 F_XET
    /-show/showfile/cue/174/midipara1  I32 F_XET
    /-show/showfile/cue/174/midipara2  I32 F_XET
/-show/showfile/cue/175  <SCUE> n=0
    /-show/showfile/cue/175/numb  I32 F_XET
    /-show/showfile/cue/175/name  S32 F_XET
    /-show/showfile/cue/175/skip  I32 F_XET
    /-show/showfile/cue/175/scene  I32 F_XET
    /-show/showfile/cue/175/bit  I32 F_XET
    /-show/showfile/cue/175/miditype  I32 F_XET
    /-show/showfile/cue/175/midichan  I32 F_XET
    /-show/showfile/cue/175/midipara1  I32 F_XET
    /-show/showfile/cue/175/midipara2  I32 F_XET
/-show/showfile/cue/176  <SCUE> n=0
    /-show/showfile/cue/176/numb  I32 F_XET
    /-show/showfile/cue/176/name  S32 F_XET
    /-show/showfile/cue/176/skip  I32 F_XET
    /-show/showfile/cue/176/scene  I32 F_XET
    /-show/showfile/cue/176/bit  I32 F_XET
    /-show/showfile/cue/176/miditype  I32 F_XET
    /-show/showfile/cue/176/midichan  I32 F_XET
    /-show/showfile/cue/176/midipara1  I32 F_XET
    /-show/showfile/cue/176/midipara2  I32 F_XET
/-show/showfile/cue/177  <SCUE> n=0
    /-show/showfile/cue/177/numb  I32 F_XET
    /-show/showfile/cue/177/name  S32 F_XET
    /-show/showfile/cue/177/skip  I32 F_XET
    /-show/showfile/cue/177/scene  I32 F_XET
    /-show/showfile/cue/177/bit  I32 F_XET
    /-show/showfile/cue/177/miditype  I32 F_XET
    /-show/showfile/cue/177/midichan  I32 F_XET
    /-show/showfile/cue/177/midipara1  I32 F_XET
    /-show/showfile/cue/177/midipara2  I32 F_XET
/-show/showfile/cue/178  <SCUE> n=0
    /-show/showfile/cue/178/numb  I32 F_XET
    /-show/showfile/cue/178/name  S32 F_XET
    /-show/showfile/cue/178/skip  I32 F_XET
    /-show/showfile/cue/178/scene  I32 F_XET
    /-show/showfile/cue/178/bit  I32 F_XET
    /-show/showfile/cue/178/miditype  I32 F_XET
    /-show/showfile/cue/178/midichan  I32 F_XET
    /-show/showfile/cue/178/midipara1  I32 F_XET
    /-show/showfile/cue/178/midipara2  I32 F_XET
/-show/showfile/cue/179  <SCUE> n=0
    /-show/showfile/cue/179/numb  I32 F_XET
    /-show/showfile/cue/179/name  S32 F_XET
    /-show/showfile/cue/179/skip  I32 F_XET
    /-show/showfile/cue/179/scene  I32 F_XET
    /-show/showfile/cue/179/bit  I32 F_XET
    /-show/showfile/cue/179/miditype  I32 F_XET
    /-show/showfile/cue/179/midichan  I32 F_XET
    /-show/showfile/cue/179/midipara1  I32 F_XET
    /-show/showfile/cue/179/midipara2  I32 F_XET
/-show/showfile/cue/180  <SCUE> n=0
    /-show/showfile/cue/180/numb  I32 F_XET
    /-show/showfile/cue/180/name  S32 F_XET
    /-show/showfile/cue/180/skip  I32 F_XET
    /-show/showfile/cue/180/scene  I32 F_XET
    /-show/showfile/cue/180/bit  I32 F_XET
    /-show/showfile/cue/180/miditype  I32 F_XET
    /-show/showfile/cue/180/midichan  I32 F_XET
    /-show/showfile/cue/180/midipara1  I32 F_XET
    /-show/showfile/cue/180/midipara2  I32 F_XET
/-show/showfile/cue/181  <SCUE> n=0
    /-show/showfile/cue/181/numb  I32 F_XET
    /-show/showfile/cue/181/name  S32 F_XET
    /-show/showfile/cue/181/skip  I32 F_XET
    /-show/showfile/cue/181/scene  I32 F_XET
    /-show/showfile/cue/181/bit  I32 F_XET
    /-show/showfile/cue/181/miditype  I32 F_XET
    /-show/showfile/cue/181/midichan  I32 F_XET
    /-show/showfile/cue/181/midipara1  I32 F_XET
    /-show/showfile/cue/181/midipara2  I32 F_XET
/-show/showfile/cue/182  <SCUE> n=0
    /-show/showfile/cue/182/numb  I32 F_XET
    /-show/showfile/cue/182/name  S32 F_XET
    /-show/showfile/cue/182/skip  I32 F_XET
    /-show/showfile/cue/182/scene  I32 F_XET
    /-show/showfile/cue/182/bit  I32 F_XET
    /-show/showfile/cue/182/miditype  I32 F_XET
    /-show/showfile/cue/182/midichan  I32 F_XET
    /-show/showfile/cue/182/midipara1  I32 F_XET
    /-show/showfile/cue/182/midipara2  I32 F_XET
/-show/showfile/cue/183  <SCUE> n=0
    /-show/showfile/cue/183/numb  I32 F_XET
    /-show/showfile/cue/183/name  S32 F_XET
    /-show/showfile/cue/183/skip  I32 F_XET
    /-show/showfile/cue/183/scene  I32 F_XET
    /-show/showfile/cue/183/bit  I32 F_XET
    /-show/showfile/cue/183/miditype  I32 F_XET
    /-show/showfile/cue/183/midichan  I32 F_XET
    /-show/showfile/cue/183/midipara1  I32 F_XET
    /-show/showfile/cue/183/midipara2  I32 F_XET
/-show/showfile/cue/184  <SCUE> n=0
    /-show/showfile/cue/184/numb  I32 F_XET
    /-show/showfile/cue/184/name  S32 F_XET
    /-show/showfile/cue/184/skip  I32 F_XET
    /-show/showfile/cue/184/scene  I32 F_XET
    /-show/showfile/cue/184/bit  I32 F_XET
    /-show/showfile/cue/184/miditype  I32 F_XET
    /-show/showfile/cue/184/midichan  I32 F_XET
    /-show/showfile/cue/184/midipara1  I32 F_XET
    /-show/showfile/cue/184/midipara2  I32 F_XET
/-show/showfile/cue/185  <SCUE> n=0
    /-show/showfile/cue/185/numb  I32 F_XET
    /-show/showfile/cue/185/name  S32 F_XET
    /-show/showfile/cue/185/skip  I32 F_XET
    /-show/showfile/cue/185/scene  I32 F_XET
    /-show/showfile/cue/185/bit  I32 F_XET
    /-show/showfile/cue/185/miditype  I32 F_XET
    /-show/showfile/cue/185/midichan  I32 F_XET
    /-show/showfile/cue/185/midipara1  I32 F_XET
    /-show/showfile/cue/185/midipara2  I32 F_XET
/-show/showfile/cue/186  <SCUE> n=0
    /-show/showfile/cue/186/numb  I32 F_XET
    /-show/showfile/cue/186/name  S32 F_XET
    /-show/showfile/cue/186/skip  I32 F_XET
    /-show/showfile/cue/186/scene  I32 F_XET
    /-show/showfile/cue/186/bit  I32 F_XET
    /-show/showfile/cue/186/miditype  I32 F_XET
    /-show/showfile/cue/186/midichan  I32 F_XET
    /-show/showfile/cue/186/midipara1  I32 F_XET
    /-show/showfile/cue/186/midipara2  I32 F_XET
/-show/showfile/cue/187  <SCUE> n=0
    /-show/showfile/cue/187/numb  I32 F_XET
    /-show/showfile/cue/187/name  S32 F_XET
    /-show/showfile/cue/187/skip  I32 F_XET
    /-show/showfile/cue/187/scene  I32 F_XET
    /-show/showfile/cue/187/bit  I32 F_XET
    /-show/showfile/cue/187/miditype  I32 F_XET
    /-show/showfile/cue/187/midichan  I32 F_XET
    /-show/showfile/cue/187/midipara1  I32 F_XET
    /-show/showfile/cue/187/midipara2  I32 F_XET
/-show/showfile/cue/188  <SCUE> n=0
    /-show/showfile/cue/188/numb  I32 F_XET
    /-show/showfile/cue/188/name  S32 F_XET
    /-show/showfile/cue/188/skip  I32 F_XET
    /-show/showfile/cue/188/scene  I32 F_XET
    /-show/showfile/cue/188/bit  I32 F_XET
    /-show/showfile/cue/188/miditype  I32 F_XET
    /-show/showfile/cue/188/midichan  I32 F_XET
    /-show/showfile/cue/188/midipara1  I32 F_XET
    /-show/showfile/cue/188/midipara2  I32 F_XET
/-show/showfile/cue/189  <SCUE> n=0
    /-show/showfile/cue/189/numb  I32 F_XET
    /-show/showfile/cue/189/name  S32 F_XET
    /-show/showfile/cue/189/skip  I32 F_XET
    /-show/showfile/cue/189/scene  I32 F_XET
    /-show/showfile/cue/189/bit  I32 F_XET
    /-show/showfile/cue/189/miditype  I32 F_XET
    /-show/showfile/cue/189/midichan  I32 F_XET
    /-show/showfile/cue/189/midipara1  I32 F_XET
    /-show/showfile/cue/189/midipara2  I32 F_XET
/-show/showfile/cue/190  <SCUE> n=0
    /-show/showfile/cue/190/numb  I32 F_XET
    /-show/showfile/cue/190/name  S32 F_XET
    /-show/showfile/cue/190/skip  I32 F_XET
    /-show/showfile/cue/190/scene  I32 F_XET
    /-show/showfile/cue/190/bit  I32 F_XET
    /-show/showfile/cue/190/miditype  I32 F_XET
    /-show/showfile/cue/190/midichan  I32 F_XET
    /-show/showfile/cue/190/midipara1  I32 F_XET
    /-show/showfile/cue/190/midipara2  I32 F_XET
/-show/showfile/cue/191  <SCUE> n=0
    /-show/showfile/cue/191/numb  I32 F_XET
    /-show/showfile/cue/191/name  S32 F_XET
    /-show/showfile/cue/191/skip  I32 F_XET
    /-show/showfile/cue/191/scene  I32 F_XET
    /-show/showfile/cue/191/bit  I32 F_XET
    /-show/showfile/cue/191/miditype  I32 F_XET
    /-show/showfile/cue/191/midichan  I32 F_XET
    /-show/showfile/cue/191/midipara1  I32 F_XET
    /-show/showfile/cue/191/midipara2  I32 F_XET
/-show/showfile/cue/192  <SCUE> n=0
    /-show/showfile/cue/192/numb  I32 F_XET
    /-show/showfile/cue/192/name  S32 F_XET
    /-show/showfile/cue/192/skip  I32 F_XET
    /-show/showfile/cue/192/scene  I32 F_XET
    /-show/showfile/cue/192/bit  I32 F_XET
    /-show/showfile/cue/192/miditype  I32 F_XET
    /-show/showfile/cue/192/midichan  I32 F_XET
    /-show/showfile/cue/192/midipara1  I32 F_XET
    /-show/showfile/cue/192/midipara2  I32 F_XET
/-show/showfile/cue/193  <SCUE> n=0
    /-show/showfile/cue/193/numb  I32 F_XET
    /-show/showfile/cue/193/name  S32 F_XET
    /-show/showfile/cue/193/skip  I32 F_XET
    /-show/showfile/cue/193/scene  I32 F_XET
    /-show/showfile/cue/193/bit  I32 F_XET
    /-show/showfile/cue/193/miditype  I32 F_XET
    /-show/showfile/cue/193/midichan  I32 F_XET
    /-show/showfile/cue/193/midipara1  I32 F_XET
    /-show/showfile/cue/193/midipara2  I32 F_XET
/-show/showfile/cue/194  <SCUE> n=0
    /-show/showfile/cue/194/numb  I32 F_XET
    /-show/showfile/cue/194/name  S32 F_XET
    /-show/showfile/cue/194/skip  I32 F_XET
    /-show/showfile/cue/194/scene  I32 F_XET
    /-show/showfile/cue/194/bit  I32 F_XET
    /-show/showfile/cue/194/miditype  I32 F_XET
    /-show/showfile/cue/194/midichan  I32 F_XET
    /-show/showfile/cue/194/midipara1  I32 F_XET
    /-show/showfile/cue/194/midipara2  I32 F_XET
/-show/showfile/cue/195  <SCUE> n=0
    /-show/showfile/cue/195/numb  I32 F_XET
    /-show/showfile/cue/195/name  S32 F_XET
    /-show/showfile/cue/195/skip  I32 F_XET
    /-show/showfile/cue/195/scene  I32 F_XET
    /-show/showfile/cue/195/bit  I32 F_XET
    /-show/showfile/cue/195/miditype  I32 F_XET
    /-show/showfile/cue/195/midichan  I32 F_XET
    /-show/showfile/cue/195/midipara1  I32 F_XET
    /-show/showfile/cue/195/midipara2  I32 F_XET
/-show/showfile/cue/196  <SCUE> n=0
    /-show/showfile/cue/196/numb  I32 F_XET
    /-show/showfile/cue/196/name  S32 F_XET
    /-show/showfile/cue/196/skip  I32 F_XET
    /-show/showfile/cue/196/scene  I32 F_XET
    /-show/showfile/cue/196/bit  I32 F_XET
    /-show/showfile/cue/196/miditype  I32 F_XET
    /-show/showfile/cue/196/midichan  I32 F_XET
    /-show/showfile/cue/196/midipara1  I32 F_XET
    /-show/showfile/cue/196/midipara2  I32 F_XET
/-show/showfile/cue/197  <SCUE> n=0
    /-show/showfile/cue/197/numb  I32 F_XET
    /-show/showfile/cue/197/name  S32 F_XET
    /-show/showfile/cue/197/skip  I32 F_XET
    /-show/showfile/cue/197/scene  I32 F_XET
    /-show/showfile/cue/197/bit  I32 F_XET
    /-show/showfile/cue/197/miditype  I32 F_XET
    /-show/showfile/cue/197/midichan  I32 F_XET
    /-show/showfile/cue/197/midipara1  I32 F_XET
    /-show/showfile/cue/197/midipara2  I32 F_XET
/-show/showfile/cue/198  <SCUE> n=0
    /-show/showfile/cue/198/numb  I32 F_XET
    /-show/showfile/cue/198/name  S32 F_XET
    /-show/showfile/cue/198/skip  I32 F_XET
    /-show/showfile/cue/198/scene  I32 F_XET
    /-show/showfile/cue/198/bit  I32 F_XET
    /-show/showfile/cue/198/miditype  I32 F_XET
    /-show/showfile/cue/198/midichan  I32 F_XET
    /-show/showfile/cue/198/midipara1  I32 F_XET
    /-show/showfile/cue/198/midipara2  I32 F_XET
/-show/showfile/cue/199  <SCUE> n=0
    /-show/showfile/cue/199/numb  I32 F_XET
    /-show/showfile/cue/199/name  S32 F_XET
    /-show/showfile/cue/199/skip  I32 F_XET
    /-show/showfile/cue/199/scene  I32 F_XET
    /-show/showfile/cue/199/bit  I32 F_XET
    /-show/showfile/cue/199/miditype  I32 F_XET
    /-show/showfile/cue/199/midichan  I32 F_XET
    /-show/showfile/cue/199/midipara1  I32 F_XET
    /-show/showfile/cue/199/midipara2  I32 F_XET
/-show/showfile/cue/200  <SCUE> n=0
    /-show/showfile/cue/200/numb  I32 F_XET
    /-show/showfile/cue/200/name  S32 F_XET
    /-show/showfile/cue/200/skip  I32 F_XET
    /-show/showfile/cue/200/scene  I32 F_XET
    /-show/showfile/cue/200/bit  I32 F_XET
    /-show/showfile/cue/200/miditype  I32 F_XET
    /-show/showfile/cue/200/midichan  I32 F_XET
    /-show/showfile/cue/200/midipara1  I32 F_XET
    /-show/showfile/cue/200/midipara2  I32 F_XET
/-show/showfile/cue/201  <SCUE> n=0
    /-show/showfile/cue/201/numb  I32 F_XET
    /-show/showfile/cue/201/name  S32 F_XET
    /-show/showfile/cue/201/skip  I32 F_XET
    /-show/showfile/cue/201/scene  I32 F_XET
    /-show/showfile/cue/201/bit  I32 F_XET
    /-show/showfile/cue/201/miditype  I32 F_XET
    /-show/showfile/cue/201/midichan  I32 F_XET
    /-show/showfile/cue/201/midipara1  I32 F_XET
    /-show/showfile/cue/201/midipara2  I32 F_XET
/-show/showfile/cue/202  <SCUE> n=0
    /-show/showfile/cue/202/numb  I32 F_XET
    /-show/showfile/cue/202/name  S32 F_XET
    /-show/showfile/cue/202/skip  I32 F_XET
    /-show/showfile/cue/202/scene  I32 F_XET
    /-show/showfile/cue/202/bit  I32 F_XET
    /-show/showfile/cue/202/miditype  I32 F_XET
    /-show/showfile/cue/202/midichan  I32 F_XET
    /-show/showfile/cue/202/midipara1  I32 F_XET
    /-show/showfile/cue/202/midipara2  I32 F_XET
/-show/showfile/cue/203  <SCUE> n=0
    /-show/showfile/cue/203/numb  I32 F_XET
    /-show/showfile/cue/203/name  S32 F_XET
    /-show/showfile/cue/203/skip  I32 F_XET
    /-show/showfile/cue/203/scene  I32 F_XET
    /-show/showfile/cue/203/bit  I32 F_XET
    /-show/showfile/cue/203/miditype  I32 F_XET
    /-show/showfile/cue/203/midichan  I32 F_XET
    /-show/showfile/cue/203/midipara1  I32 F_XET
    /-show/showfile/cue/203/midipara2  I32 F_XET
/-show/showfile/cue/204  <SCUE> n=0
    /-show/showfile/cue/204/numb  I32 F_XET
    /-show/showfile/cue/204/name  S32 F_XET
    /-show/showfile/cue/204/skip  I32 F_XET
    /-show/showfile/cue/204/scene  I32 F_XET
    /-show/showfile/cue/204/bit  I32 F_XET
    /-show/showfile/cue/204/miditype  I32 F_XET
    /-show/showfile/cue/204/midichan  I32 F_XET
    /-show/showfile/cue/204/midipara1  I32 F_XET
    /-show/showfile/cue/204/midipara2  I32 F_XET
/-show/showfile/cue/205  <SCUE> n=0
    /-show/showfile/cue/205/numb  I32 F_XET
    /-show/showfile/cue/205/name  S32 F_XET
    /-show/showfile/cue/205/skip  I32 F_XET
    /-show/showfile/cue/205/scene  I32 F_XET
    /-show/showfile/cue/205/bit  I32 F_XET
    /-show/showfile/cue/205/miditype  I32 F_XET
    /-show/showfile/cue/205/midichan  I32 F_XET
    /-show/showfile/cue/205/midipara1  I32 F_XET
    /-show/showfile/cue/205/midipara2  I32 F_XET
/-show/showfile/cue/206  <SCUE> n=0
    /-show/showfile/cue/206/numb  I32 F_XET
    /-show/showfile/cue/206/name  S32 F_XET
    /-show/showfile/cue/206/skip  I32 F_XET
    /-show/showfile/cue/206/scene  I32 F_XET
    /-show/showfile/cue/206/bit  I32 F_XET
    /-show/showfile/cue/206/miditype  I32 F_XET
    /-show/showfile/cue/206/midichan  I32 F_XET
    /-show/showfile/cue/206/midipara1  I32 F_XET
    /-show/showfile/cue/206/midipara2  I32 F_XET
/-show/showfile/cue/207  <SCUE> n=0
    /-show/showfile/cue/207/numb  I32 F_XET
    /-show/showfile/cue/207/name  S32 F_XET
    /-show/showfile/cue/207/skip  I32 F_XET
    /-show/showfile/cue/207/scene  I32 F_XET
    /-show/showfile/cue/207/bit  I32 F_XET
    /-show/showfile/cue/207/miditype  I32 F_XET
    /-show/showfile/cue/207/midichan  I32 F_XET
    /-show/showfile/cue/207/midipara1  I32 F_XET
    /-show/showfile/cue/207/midipara2  I32 F_XET
/-show/showfile/cue/208  <SCUE> n=0
    /-show/showfile/cue/208/numb  I32 F_XET
    /-show/showfile/cue/208/name  S32 F_XET
    /-show/showfile/cue/208/skip  I32 F_XET
    /-show/showfile/cue/208/scene  I32 F_XET
    /-show/showfile/cue/208/bit  I32 F_XET
    /-show/showfile/cue/208/miditype  I32 F_XET
    /-show/showfile/cue/208/midichan  I32 F_XET
    /-show/showfile/cue/208/midipara1  I32 F_XET
    /-show/showfile/cue/208/midipara2  I32 F_XET
/-show/showfile/cue/209  <SCUE> n=0
    /-show/showfile/cue/209/numb  I32 F_XET
    /-show/showfile/cue/209/name  S32 F_XET
    /-show/showfile/cue/209/skip  I32 F_XET
    /-show/showfile/cue/209/scene  I32 F_XET
    /-show/showfile/cue/209/bit  I32 F_XET
    /-show/showfile/cue/209/miditype  I32 F_XET
    /-show/showfile/cue/209/midichan  I32 F_XET
    /-show/showfile/cue/209/midipara1  I32 F_XET
    /-show/showfile/cue/209/midipara2  I32 F_XET
/-show/showfile/cue/210  <SCUE> n=0
    /-show/showfile/cue/210/numb  I32 F_XET
    /-show/showfile/cue/210/name  S32 F_XET
    /-show/showfile/cue/210/skip  I32 F_XET
    /-show/showfile/cue/210/scene  I32 F_XET
    /-show/showfile/cue/210/bit  I32 F_XET
    /-show/showfile/cue/210/miditype  I32 F_XET
    /-show/showfile/cue/210/midichan  I32 F_XET
    /-show/showfile/cue/210/midipara1  I32 F_XET
    /-show/showfile/cue/210/midipara2  I32 F_XET
/-show/showfile/cue/211  <SCUE> n=0
    /-show/showfile/cue/211/numb  I32 F_XET
    /-show/showfile/cue/211/name  S32 F_XET
    /-show/showfile/cue/211/skip  I32 F_XET
    /-show/showfile/cue/211/scene  I32 F_XET
    /-show/showfile/cue/211/bit  I32 F_XET
    /-show/showfile/cue/211/miditype  I32 F_XET
    /-show/showfile/cue/211/midichan  I32 F_XET
    /-show/showfile/cue/211/midipara1  I32 F_XET
    /-show/showfile/cue/211/midipara2  I32 F_XET
/-show/showfile/cue/212  <SCUE> n=0
    /-show/showfile/cue/212/numb  I32 F_XET
    /-show/showfile/cue/212/name  S32 F_XET
    /-show/showfile/cue/212/skip  I32 F_XET
    /-show/showfile/cue/212/scene  I32 F_XET
    /-show/showfile/cue/212/bit  I32 F_XET
    /-show/showfile/cue/212/miditype  I32 F_XET
    /-show/showfile/cue/212/midichan  I32 F_XET
    /-show/showfile/cue/212/midipara1  I32 F_XET
    /-show/showfile/cue/212/midipara2  I32 F_XET
/-show/showfile/cue/213  <SCUE> n=0
    /-show/showfile/cue/213/numb  I32 F_XET
    /-show/showfile/cue/213/name  S32 F_XET
    /-show/showfile/cue/213/skip  I32 F_XET
    /-show/showfile/cue/213/scene  I32 F_XET
    /-show/showfile/cue/213/bit  I32 F_XET
    /-show/showfile/cue/213/miditype  I32 F_XET
    /-show/showfile/cue/213/midichan  I32 F_XET
    /-show/showfile/cue/213/midipara1  I32 F_XET
    /-show/showfile/cue/213/midipara2  I32 F_XET
/-show/showfile/cue/214  <SCUE> n=0
    /-show/showfile/cue/214/numb  I32 F_XET
    /-show/showfile/cue/214/name  S32 F_XET
    /-show/showfile/cue/214/skip  I32 F_XET
    /-show/showfile/cue/214/scene  I32 F_XET
    /-show/showfile/cue/214/bit  I32 F_XET
    /-show/showfile/cue/214/miditype  I32 F_XET
    /-show/showfile/cue/214/midichan  I32 F_XET
    /-show/showfile/cue/214/midipara1  I32 F_XET
    /-show/showfile/cue/214/midipara2  I32 F_XET
/-show/showfile/cue/215  <SCUE> n=0
    /-show/showfile/cue/215/numb  I32 F_XET
    /-show/showfile/cue/215/name  S32 F_XET
    /-show/showfile/cue/215/skip  I32 F_XET
    /-show/showfile/cue/215/scene  I32 F_XET
    /-show/showfile/cue/215/bit  I32 F_XET
    /-show/showfile/cue/215/miditype  I32 F_XET
    /-show/showfile/cue/215/midichan  I32 F_XET
    /-show/showfile/cue/215/midipara1  I32 F_XET
    /-show/showfile/cue/215/midipara2  I32 F_XET
/-show/showfile/cue/216  <SCUE> n=0
    /-show/showfile/cue/216/numb  I32 F_XET
    /-show/showfile/cue/216/name  S32 F_XET
    /-show/showfile/cue/216/skip  I32 F_XET
    /-show/showfile/cue/216/scene  I32 F_XET
    /-show/showfile/cue/216/bit  I32 F_XET
    /-show/showfile/cue/216/miditype  I32 F_XET
    /-show/showfile/cue/216/midichan  I32 F_XET
    /-show/showfile/cue/216/midipara1  I32 F_XET
    /-show/showfile/cue/216/midipara2  I32 F_XET
/-show/showfile/cue/217  <SCUE> n=0
    /-show/showfile/cue/217/numb  I32 F_XET
    /-show/showfile/cue/217/name  S32 F_XET
    /-show/showfile/cue/217/skip  I32 F_XET
    /-show/showfile/cue/217/scene  I32 F_XET
    /-show/showfile/cue/217/bit  I32 F_XET
    /-show/showfile/cue/217/miditype  I32 F_XET
    /-show/showfile/cue/217/midichan  I32 F_XET
    /-show/showfile/cue/217/midipara1  I32 F_XET
    /-show/showfile/cue/217/midipara2  I32 F_XET
/-show/showfile/cue/218  <SCUE> n=0
    /-show/showfile/cue/218/numb  I32 F_XET
    /-show/showfile/cue/218/name  S32 F_XET
    /-show/showfile/cue/218/skip  I32 F_XET
    /-show/showfile/cue/218/scene  I32 F_XET
    /-show/showfile/cue/218/bit  I32 F_XET
    /-show/showfile/cue/218/miditype  I32 F_XET
    /-show/showfile/cue/218/midichan  I32 F_XET
    /-show/showfile/cue/218/midipara1  I32 F_XET
    /-show/showfile/cue/218/midipara2  I32 F_XET
/-show/showfile/cue/219  <SCUE> n=0
    /-show/showfile/cue/219/numb  I32 F_XET
    /-show/showfile/cue/219/name  S32 F_XET
    /-show/showfile/cue/219/skip  I32 F_XET
    /-show/showfile/cue/219/scene  I32 F_XET
    /-show/showfile/cue/219/bit  I32 F_XET
    /-show/showfile/cue/219/miditype  I32 F_XET
    /-show/showfile/cue/219/midichan  I32 F_XET
    /-show/showfile/cue/219/midipara1  I32 F_XET
    /-show/showfile/cue/219/midipara2  I32 F_XET
/-show/showfile/cue/220  <SCUE> n=0
    /-show/showfile/cue/220/numb  I32 F_XET
    /-show/showfile/cue/220/name  S32 F_XET
    /-show/showfile/cue/220/skip  I32 F_XET
    /-show/showfile/cue/220/scene  I32 F_XET
    /-show/showfile/cue/220/bit  I32 F_XET
    /-show/showfile/cue/220/miditype  I32 F_XET
    /-show/showfile/cue/220/midichan  I32 F_XET
    /-show/showfile/cue/220/midipara1  I32 F_XET
    /-show/showfile/cue/220/midipara2  I32 F_XET
/-show/showfile/cue/221  <SCUE> n=0
    /-show/showfile/cue/221/numb  I32 F_XET
    /-show/showfile/cue/221/name  S32 F_XET
    /-show/showfile/cue/221/skip  I32 F_XET
    /-show/showfile/cue/221/scene  I32 F_XET
    /-show/showfile/cue/221/bit  I32 F_XET
    /-show/showfile/cue/221/miditype  I32 F_XET
    /-show/showfile/cue/221/midichan  I32 F_XET
    /-show/showfile/cue/221/midipara1  I32 F_XET
    /-show/showfile/cue/221/midipara2  I32 F_XET
/-show/showfile/cue/222  <SCUE> n=0
    /-show/showfile/cue/222/numb  I32 F_XET
    /-show/showfile/cue/222/name  S32 F_XET
    /-show/showfile/cue/222/skip  I32 F_XET
    /-show/showfile/cue/222/scene  I32 F_XET
    /-show/showfile/cue/222/bit  I32 F_XET
    /-show/showfile/cue/222/miditype  I32 F_XET
    /-show/showfile/cue/222/midichan  I32 F_XET
    /-show/showfile/cue/222/midipara1  I32 F_XET
    /-show/showfile/cue/222/midipara2  I32 F_XET
/-show/showfile/cue/223  <SCUE> n=0
    /-show/showfile/cue/223/numb  I32 F_XET
    /-show/showfile/cue/223/name  S32 F_XET
    /-show/showfile/cue/223/skip  I32 F_XET
    /-show/showfile/cue/223/scene  I32 F_XET
    /-show/showfile/cue/223/bit  I32 F_XET
    /-show/showfile/cue/223/miditype  I32 F_XET
    /-show/showfile/cue/223/midichan  I32 F_XET
    /-show/showfile/cue/223/midipara1  I32 F_XET
    /-show/showfile/cue/223/midipara2  I32 F_XET
/-show/showfile/cue/224  <SCUE> n=0
    /-show/showfile/cue/224/numb  I32 F_XET
    /-show/showfile/cue/224/name  S32 F_XET
    /-show/showfile/cue/224/skip  I32 F_XET
    /-show/showfile/cue/224/scene  I32 F_XET
    /-show/showfile/cue/224/bit  I32 F_XET
    /-show/showfile/cue/224/miditype  I32 F_XET
    /-show/showfile/cue/224/midichan  I32 F_XET
    /-show/showfile/cue/224/midipara1  I32 F_XET
    /-show/showfile/cue/224/midipara2  I32 F_XET
/-show/showfile/cue/225  <SCUE> n=0
    /-show/showfile/cue/225/numb  I32 F_XET
    /-show/showfile/cue/225/name  S32 F_XET
    /-show/showfile/cue/225/skip  I32 F_XET
    /-show/showfile/cue/225/scene  I32 F_XET
    /-show/showfile/cue/225/bit  I32 F_XET
    /-show/showfile/cue/225/miditype  I32 F_XET
    /-show/showfile/cue/225/midichan  I32 F_XET
    /-show/showfile/cue/225/midipara1  I32 F_XET
    /-show/showfile/cue/225/midipara2  I32 F_XET
/-show/showfile/cue/226  <SCUE> n=0
    /-show/showfile/cue/226/numb  I32 F_XET
    /-show/showfile/cue/226/name  S32 F_XET
    /-show/showfile/cue/226/skip  I32 F_XET
    /-show/showfile/cue/226/scene  I32 F_XET
    /-show/showfile/cue/226/bit  I32 F_XET
    /-show/showfile/cue/226/miditype  I32 F_XET
    /-show/showfile/cue/226/midichan  I32 F_XET
    /-show/showfile/cue/226/midipara1  I32 F_XET
    /-show/showfile/cue/226/midipara2  I32 F_XET
/-show/showfile/cue/227  <SCUE> n=0
    /-show/showfile/cue/227/numb  I32 F_XET
    /-show/showfile/cue/227/name  S32 F_XET
    /-show/showfile/cue/227/skip  I32 F_XET
    /-show/showfile/cue/227/scene  I32 F_XET
    /-show/showfile/cue/227/bit  I32 F_XET
    /-show/showfile/cue/227/miditype  I32 F_XET
    /-show/showfile/cue/227/midichan  I32 F_XET
    /-show/showfile/cue/227/midipara1  I32 F_XET
    /-show/showfile/cue/227/midipara2  I32 F_XET
/-show/showfile/cue/228  <SCUE> n=0
    /-show/showfile/cue/228/numb  I32 F_XET
    /-show/showfile/cue/228/name  S32 F_XET
    /-show/showfile/cue/228/skip  I32 F_XET
    /-show/showfile/cue/228/scene  I32 F_XET
    /-show/showfile/cue/228/bit  I32 F_XET
    /-show/showfile/cue/228/miditype  I32 F_XET
    /-show/showfile/cue/228/midichan  I32 F_XET
    /-show/showfile/cue/228/midipara1  I32 F_XET
    /-show/showfile/cue/228/midipara2  I32 F_XET
/-show/showfile/cue/229  <SCUE> n=0
    /-show/showfile/cue/229/numb  I32 F_XET
    /-show/showfile/cue/229/name  S32 F_XET
    /-show/showfile/cue/229/skip  I32 F_XET
    /-show/showfile/cue/229/scene  I32 F_XET
    /-show/showfile/cue/229/bit  I32 F_XET
    /-show/showfile/cue/229/miditype  I32 F_XET
    /-show/showfile/cue/229/midichan  I32 F_XET
    /-show/showfile/cue/229/midipara1  I32 F_XET
    /-show/showfile/cue/229/midipara2  I32 F_XET
/-show/showfile/cue/230  <SCUE> n=0
    /-show/showfile/cue/230/numb  I32 F_XET
    /-show/showfile/cue/230/name  S32 F_XET
    /-show/showfile/cue/230/skip  I32 F_XET
    /-show/showfile/cue/230/scene  I32 F_XET
    /-show/showfile/cue/230/bit  I32 F_XET
    /-show/showfile/cue/230/miditype  I32 F_XET
    /-show/showfile/cue/230/midichan  I32 F_XET
    /-show/showfile/cue/230/midipara1  I32 F_XET
    /-show/showfile/cue/230/midipara2  I32 F_XET
/-show/showfile/cue/231  <SCUE> n=0
    /-show/showfile/cue/231/numb  I32 F_XET
    /-show/showfile/cue/231/name  S32 F_XET
    /-show/showfile/cue/231/skip  I32 F_XET
    /-show/showfile/cue/231/scene  I32 F_XET
    /-show/showfile/cue/231/bit  I32 F_XET
    /-show/showfile/cue/231/miditype  I32 F_XET
    /-show/showfile/cue/231/midichan  I32 F_XET
    /-show/showfile/cue/231/midipara1  I32 F_XET
    /-show/showfile/cue/231/midipara2  I32 F_XET
/-show/showfile/cue/232  <SCUE> n=0
    /-show/showfile/cue/232/numb  I32 F_XET
    /-show/showfile/cue/232/name  S32 F_XET
    /-show/showfile/cue/232/skip  I32 F_XET
    /-show/showfile/cue/232/scene  I32 F_XET
    /-show/showfile/cue/232/bit  I32 F_XET
    /-show/showfile/cue/232/miditype  I32 F_XET
    /-show/showfile/cue/232/midichan  I32 F_XET
    /-show/showfile/cue/232/midipara1  I32 F_XET
    /-show/showfile/cue/232/midipara2  I32 F_XET
/-show/showfile/cue/233  <SCUE> n=0
    /-show/showfile/cue/233/numb  I32 F_XET
    /-show/showfile/cue/233/name  S32 F_XET
    /-show/showfile/cue/233/skip  I32 F_XET
    /-show/showfile/cue/233/scene  I32 F_XET
    /-show/showfile/cue/233/bit  I32 F_XET
    /-show/showfile/cue/233/miditype  I32 F_XET
    /-show/showfile/cue/233/midichan  I32 F_XET
    /-show/showfile/cue/233/midipara1  I32 F_XET
    /-show/showfile/cue/233/midipara2  I32 F_XET
/-show/showfile/cue/234  <SCUE> n=0
    /-show/showfile/cue/234/numb  I32 F_XET
    /-show/showfile/cue/234/name  S32 F_XET
    /-show/showfile/cue/234/skip  I32 F_XET
    /-show/showfile/cue/234/scene  I32 F_XET
    /-show/showfile/cue/234/bit  I32 F_XET
    /-show/showfile/cue/234/miditype  I32 F_XET
    /-show/showfile/cue/234/midichan  I32 F_XET
    /-show/showfile/cue/234/midipara1  I32 F_XET
    /-show/showfile/cue/234/midipara2  I32 F_XET
/-show/showfile/cue/235  <SCUE> n=0
    /-show/showfile/cue/235/numb  I32 F_XET
    /-show/showfile/cue/235/name  S32 F_XET
    /-show/showfile/cue/235/skip  I32 F_XET
    /-show/showfile/cue/235/scene  I32 F_XET
    /-show/showfile/cue/235/bit  I32 F_XET
    /-show/showfile/cue/235/miditype  I32 F_XET
    /-show/showfile/cue/235/midichan  I32 F_XET
    /-show/showfile/cue/235/midipara1  I32 F_XET
    /-show/showfile/cue/235/midipara2  I32 F_XET
/-show/showfile/cue/236  <SCUE> n=0
    /-show/showfile/cue/236/numb  I32 F_XET
    /-show/showfile/cue/236/name  S32 F_XET
    /-show/showfile/cue/236/skip  I32 F_XET
    /-show/showfile/cue/236/scene  I32 F_XET
    /-show/showfile/cue/236/bit  I32 F_XET
    /-show/showfile/cue/236/miditype  I32 F_XET
    /-show/showfile/cue/236/midichan  I32 F_XET
    /-show/showfile/cue/236/midipara1  I32 F_XET
    /-show/showfile/cue/236/midipara2  I32 F_XET
/-show/showfile/cue/237  <SCUE> n=0
    /-show/showfile/cue/237/numb  I32 F_XET
    /-show/showfile/cue/237/name  S32 F_XET
    /-show/showfile/cue/237/skip  I32 F_XET
    /-show/showfile/cue/237/scene  I32 F_XET
    /-show/showfile/cue/237/bit  I32 F_XET
    /-show/showfile/cue/237/miditype  I32 F_XET
    /-show/showfile/cue/237/midichan  I32 F_XET
    /-show/showfile/cue/237/midipara1  I32 F_XET
    /-show/showfile/cue/237/midipara2  I32 F_XET
/-show/showfile/cue/238  <SCUE> n=0
    /-show/showfile/cue/238/numb  I32 F_XET
    /-show/showfile/cue/238/name  S32 F_XET
    /-show/showfile/cue/238/skip  I32 F_XET
    /-show/showfile/cue/238/scene  I32 F_XET
    /-show/showfile/cue/238/bit  I32 F_XET
    /-show/showfile/cue/238/miditype  I32 F_XET
    /-show/showfile/cue/238/midichan  I32 F_XET
    /-show/showfile/cue/238/midipara1  I32 F_XET
    /-show/showfile/cue/238/midipara2  I32 F_XET
/-show/showfile/cue/239  <SCUE> n=0
    /-show/showfile/cue/239/numb  I32 F_XET
    /-show/showfile/cue/239/name  S32 F_XET
    /-show/showfile/cue/239/skip  I32 F_XET
    /-show/showfile/cue/239/scene  I32 F_XET
    /-show/showfile/cue/239/bit  I32 F_XET
    /-show/showfile/cue/239/miditype  I32 F_XET
    /-show/showfile/cue/239/midichan  I32 F_XET
    /-show/showfile/cue/239/midipara1  I32 F_XET
    /-show/showfile/cue/239/midipara2  I32 F_XET
/-show/showfile/cue/240  <SCUE> n=0
    /-show/showfile/cue/240/numb  I32 F_XET
    /-show/showfile/cue/240/name  S32 F_XET
    /-show/showfile/cue/240/skip  I32 F_XET
    /-show/showfile/cue/240/scene  I32 F_XET
    /-show/showfile/cue/240/bit  I32 F_XET
    /-show/showfile/cue/240/miditype  I32 F_XET
    /-show/showfile/cue/240/midichan  I32 F_XET
    /-show/showfile/cue/240/midipara1  I32 F_XET
    /-show/showfile/cue/240/midipara2  I32 F_XET
/-show/showfile/cue/241  <SCUE> n=0
    /-show/showfile/cue/241/numb  I32 F_XET
    /-show/showfile/cue/241/name  S32 F_XET
    /-show/showfile/cue/241/skip  I32 F_XET
    /-show/showfile/cue/241/scene  I32 F_XET
    /-show/showfile/cue/241/bit  I32 F_XET
    /-show/showfile/cue/241/miditype  I32 F_XET
    /-show/showfile/cue/241/midichan  I32 F_XET
    /-show/showfile/cue/241/midipara1  I32 F_XET
    /-show/showfile/cue/241/midipara2  I32 F_XET
/-show/showfile/cue/242  <SCUE> n=0
    /-show/showfile/cue/242/numb  I32 F_XET
    /-show/showfile/cue/242/name  S32 F_XET
    /-show/showfile/cue/242/skip  I32 F_XET
    /-show/showfile/cue/242/scene  I32 F_XET
    /-show/showfile/cue/242/bit  I32 F_XET
    /-show/showfile/cue/242/miditype  I32 F_XET
    /-show/showfile/cue/242/midichan  I32 F_XET
    /-show/showfile/cue/242/midipara1  I32 F_XET
    /-show/showfile/cue/242/midipara2  I32 F_XET
/-show/showfile/cue/243  <SCUE> n=0
    /-show/showfile/cue/243/numb  I32 F_XET
    /-show/showfile/cue/243/name  S32 F_XET
    /-show/showfile/cue/243/skip  I32 F_XET
    /-show/showfile/cue/243/scene  I32 F_XET
    /-show/showfile/cue/243/bit  I32 F_XET
    /-show/showfile/cue/243/miditype  I32 F_XET
    /-show/showfile/cue/243/midichan  I32 F_XET
    /-show/showfile/cue/243/midipara1  I32 F_XET
    /-show/showfile/cue/243/midipara2  I32 F_XET
/-show/showfile/cue/244  <SCUE> n=0
    /-show/showfile/cue/244/numb  I32 F_XET
    /-show/showfile/cue/244/name  S32 F_XET
    /-show/showfile/cue/244/skip  I32 F_XET
    /-show/showfile/cue/244/scene  I32 F_XET
    /-show/showfile/cue/244/bit  I32 F_XET
    /-show/showfile/cue/244/miditype  I32 F_XET
    /-show/showfile/cue/244/midichan  I32 F_XET
    /-show/showfile/cue/244/midipara1  I32 F_XET
    /-show/showfile/cue/244/midipara2  I32 F_XET
/-show/showfile/cue/245  <SCUE> n=0
    /-show/showfile/cue/245/numb  I32 F_XET
    /-show/showfile/cue/245/name  S32 F_XET
    /-show/showfile/cue/245/skip  I32 F_XET
    /-show/showfile/cue/245/scene  I32 F_XET
    /-show/showfile/cue/245/bit  I32 F_XET
    /-show/showfile/cue/245/miditype  I32 F_XET
    /-show/showfile/cue/245/midichan  I32 F_XET
    /-show/showfile/cue/245/midipara1  I32 F_XET
    /-show/showfile/cue/245/midipara2  I32 F_XET
/-show/showfile/cue/246  <SCUE> n=0
    /-show/showfile/cue/246/numb  I32 F_XET
    /-show/showfile/cue/246/name  S32 F_XET
    /-show/showfile/cue/246/skip  I32 F_XET
    /-show/showfile/cue/246/scene  I32 F_XET
    /-show/showfile/cue/246/bit  I32 F_XET
    /-show/showfile/cue/246/miditype  I32 F_XET
    /-show/showfile/cue/246/midichan  I32 F_XET
    /-show/showfile/cue/246/midipara1  I32 F_XET
    /-show/showfile/cue/246/midipara2  I32 F_XET
/-show/showfile/cue/247  <SCUE> n=0
    /-show/showfile/cue/247/numb  I32 F_XET
    /-show/showfile/cue/247/name  S32 F_XET
    /-show/showfile/cue/247/skip  I32 F_XET
    /-show/showfile/cue/247/scene  I32 F_XET
    /-show/showfile/cue/247/bit  I32 F_XET
    /-show/showfile/cue/247/miditype  I32 F_XET
    /-show/showfile/cue/247/midichan  I32 F_XET
    /-show/showfile/cue/247/midipara1  I32 F_XET
    /-show/showfile/cue/247/midipara2  I32 F_XET
/-show/showfile/cue/248  <SCUE> n=0
    /-show/showfile/cue/248/numb  I32 F_XET
    /-show/showfile/cue/248/name  S32 F_XET
    /-show/showfile/cue/248/skip  I32 F_XET
    /-show/showfile/cue/248/scene  I32 F_XET
    /-show/showfile/cue/248/bit  I32 F_XET
    /-show/showfile/cue/248/miditype  I32 F_XET
    /-show/showfile/cue/248/midichan  I32 F_XET
    /-show/showfile/cue/248/midipara1  I32 F_XET
    /-show/showfile/cue/248/midipara2  I32 F_XET
/-show/showfile/cue/249  <SCUE> n=0
    /-show/showfile/cue/249/numb  I32 F_XET
    /-show/showfile/cue/249/name  S32 F_XET
    /-show/showfile/cue/249/skip  I32 F_XET
    /-show/showfile/cue/249/scene  I32 F_XET
    /-show/showfile/cue/249/bit  I32 F_XET
    /-show/showfile/cue/249/miditype  I32 F_XET
    /-show/showfile/cue/249/midichan  I32 F_XET
    /-show/showfile/cue/249/midipara1  I32 F_XET
    /-show/showfile/cue/249/midipara2  I32 F_XET
/-show/showfile/cue/250  <SCUE> n=0
    /-show/showfile/cue/250/numb  I32 F_XET
    /-show/showfile/cue/250/name  S32 F_XET
    /-show/showfile/cue/250/skip  I32 F_XET
    /-show/showfile/cue/250/scene  I32 F_XET
    /-show/showfile/cue/250/bit  I32 F_XET
    /-show/showfile/cue/250/miditype  I32 F_XET
    /-show/showfile/cue/250/midichan  I32 F_XET
    /-show/showfile/cue/250/midipara1  I32 F_XET
    /-show/showfile/cue/250/midipara2  I32 F_XET
/-show/showfile/cue/251  <SCUE> n=0
    /-show/showfile/cue/251/numb  I32 F_XET
    /-show/showfile/cue/251/name  S32 F_XET
    /-show/showfile/cue/251/skip  I32 F_XET
    /-show/showfile/cue/251/scene  I32 F_XET
    /-show/showfile/cue/251/bit  I32 F_XET
    /-show/showfile/cue/251/miditype  I32 F_XET
    /-show/showfile/cue/251/midichan  I32 F_XET
    /-show/showfile/cue/251/midipara1  I32 F_XET
    /-show/showfile/cue/251/midipara2  I32 F_XET
/-show/showfile/cue/252  <SCUE> n=0
    /-show/showfile/cue/252/numb  I32 F_XET
    /-show/showfile/cue/252/name  S32 F_XET
    /-show/showfile/cue/252/skip  I32 F_XET
    /-show/showfile/cue/252/scene  I32 F_XET
    /-show/showfile/cue/252/bit  I32 F_XET
    /-show/showfile/cue/252/miditype  I32 F_XET
    /-show/showfile/cue/252/midichan  I32 F_XET
    /-show/showfile/cue/252/midipara1  I32 F_XET
    /-show/showfile/cue/252/midipara2  I32 F_XET
/-show/showfile/cue/253  <SCUE> n=0
    /-show/showfile/cue/253/numb  I32 F_XET
    /-show/showfile/cue/253/name  S32 F_XET
    /-show/showfile/cue/253/skip  I32 F_XET
    /-show/showfile/cue/253/scene  I32 F_XET
    /-show/showfile/cue/253/bit  I32 F_XET
    /-show/showfile/cue/253/miditype  I32 F_XET
    /-show/showfile/cue/253/midichan  I32 F_XET
    /-show/showfile/cue/253/midipara1  I32 F_XET
    /-show/showfile/cue/253/midipara2  I32 F_XET
/-show/showfile/cue/254  <SCUE> n=0
    /-show/showfile/cue/254/numb  I32 F_XET
    /-show/showfile/cue/254/name  S32 F_XET
    /-show/showfile/cue/254/skip  I32 F_XET
    /-show/showfile/cue/254/scene  I32 F_XET
    /-show/showfile/cue/254/bit  I32 F_XET
    /-show/showfile/cue/254/miditype  I32 F_XET
    /-show/showfile/cue/254/midichan  I32 F_XET
    /-show/showfile/cue/254/midipara1  I32 F_XET
    /-show/showfile/cue/254/midipara2  I32 F_XET
/-show/showfile/cue/255  <SCUE> n=0
    /-show/showfile/cue/255/numb  I32 F_XET
    /-show/showfile/cue/255/name  S32 F_XET
    /-show/showfile/cue/255/skip  I32 F_XET
    /-show/showfile/cue/255/scene  I32 F_XET
    /-show/showfile/cue/255/bit  I32 F_XET
    /-show/showfile/cue/255/miditype  I32 F_XET
    /-show/showfile/cue/255/midichan  I32 F_XET
    /-show/showfile/cue/255/midipara1  I32 F_XET
    /-show/showfile/cue/255/midipara2  I32 F_XET
/-show/showfile/cue/256  <SCUE> n=0
    /-show/showfile/cue/256/numb  I32 F_XET
    /-show/showfile/cue/256/name  S32 F_XET
    /-show/showfile/cue/256/skip  I32 F_XET
    /-show/showfile/cue/256/scene  I32 F_XET
    /-show/showfile/cue/256/bit  I32 F_XET
    /-show/showfile/cue/256/miditype  I32 F_XET
    /-show/showfile/cue/256/midichan  I32 F_XET
    /-show/showfile/cue/256/midipara1  I32 F_XET
    /-show/showfile/cue/256/midipara2  I32 F_XET
/-show/showfile/cue/257  <SCUE> n=0
    /-show/showfile/cue/257/numb  I32 F_XET
    /-show/showfile/cue/257/name  S32 F_XET
    /-show/showfile/cue/257/skip  I32 F_XET
    /-show/showfile/cue/257/scene  I32 F_XET
    /-show/showfile/cue/257/bit  I32 F_XET
    /-show/showfile/cue/257/miditype  I32 F_XET
    /-show/showfile/cue/257/midichan  I32 F_XET
    /-show/showfile/cue/257/midipara1  I32 F_XET
    /-show/showfile/cue/257/midipara2  I32 F_XET
/-show/showfile/cue/258  <SCUE> n=0
    /-show/showfile/cue/258/numb  I32 F_XET
    /-show/showfile/cue/258/name  S32 F_XET
    /-show/showfile/cue/258/skip  I32 F_XET
    /-show/showfile/cue/258/scene  I32 F_XET
    /-show/showfile/cue/258/bit  I32 F_XET
    /-show/showfile/cue/258/miditype  I32 F_XET
    /-show/showfile/cue/258/midichan  I32 F_XET
    /-show/showfile/cue/258/midipara1  I32 F_XET
    /-show/showfile/cue/258/midipara2  I32 F_XET
/-show/showfile/cue/259  <SCUE> n=0
    /-show/showfile/cue/259/numb  I32 F_XET
    /-show/showfile/cue/259/name  S32 F_XET
    /-show/showfile/cue/259/skip  I32 F_XET
    /-show/showfile/cue/259/scene  I32 F_XET
    /-show/showfile/cue/259/bit  I32 F_XET
    /-show/showfile/cue/259/miditype  I32 F_XET
    /-show/showfile/cue/259/midichan  I32 F_XET
    /-show/showfile/cue/259/midipara1  I32 F_XET
    /-show/showfile/cue/259/midipara2  I32 F_XET
/-show/showfile/cue/260  <SCUE> n=0
    /-show/showfile/cue/260/numb  I32 F_XET
    /-show/showfile/cue/260/name  S32 F_XET
    /-show/showfile/cue/260/skip  I32 F_XET
    /-show/showfile/cue/260/scene  I32 F_XET
    /-show/showfile/cue/260/bit  I32 F_XET
    /-show/showfile/cue/260/miditype  I32 F_XET
    /-show/showfile/cue/260/midichan  I32 F_XET
    /-show/showfile/cue/260/midipara1  I32 F_XET
    /-show/showfile/cue/260/midipara2  I32 F_XET
/-show/showfile/cue/261  <SCUE> n=0
    /-show/showfile/cue/261/numb  I32 F_XET
    /-show/showfile/cue/261/name  S32 F_XET
    /-show/showfile/cue/261/skip  I32 F_XET
    /-show/showfile/cue/261/scene  I32 F_XET
    /-show/showfile/cue/261/bit  I32 F_XET
    /-show/showfile/cue/261/miditype  I32 F_XET
    /-show/showfile/cue/261/midichan  I32 F_XET
    /-show/showfile/cue/261/midipara1  I32 F_XET
    /-show/showfile/cue/261/midipara2  I32 F_XET
/-show/showfile/cue/262  <SCUE> n=0
    /-show/showfile/cue/262/numb  I32 F_XET
    /-show/showfile/cue/262/name  S32 F_XET
    /-show/showfile/cue/262/skip  I32 F_XET
    /-show/showfile/cue/262/scene  I32 F_XET
    /-show/showfile/cue/262/bit  I32 F_XET
    /-show/showfile/cue/262/miditype  I32 F_XET
    /-show/showfile/cue/262/midichan  I32 F_XET
    /-show/showfile/cue/262/midipara1  I32 F_XET
    /-show/showfile/cue/262/midipara2  I32 F_XET
/-show/showfile/cue/263  <SCUE> n=0
    /-show/showfile/cue/263/numb  I32 F_XET
    /-show/showfile/cue/263/name  S32 F_XET
    /-show/showfile/cue/263/skip  I32 F_XET
    /-show/showfile/cue/263/scene  I32 F_XET
    /-show/showfile/cue/263/bit  I32 F_XET
    /-show/showfile/cue/263/miditype  I32 F_XET
    /-show/showfile/cue/263/midichan  I32 F_XET
    /-show/showfile/cue/263/midipara1  I32 F_XET
    /-show/showfile/cue/263/midipara2  I32 F_XET
/-show/showfile/cue/264  <SCUE> n=0
    /-show/showfile/cue/264/numb  I32 F_XET
    /-show/showfile/cue/264/name  S32 F_XET
    /-show/showfile/cue/264/skip  I32 F_XET
    /-show/showfile/cue/264/scene  I32 F_XET
    /-show/showfile/cue/264/bit  I32 F_XET
    /-show/showfile/cue/264/miditype  I32 F_XET
    /-show/showfile/cue/264/midichan  I32 F_XET
    /-show/showfile/cue/264/midipara1  I32 F_XET
    /-show/showfile/cue/264/midipara2  I32 F_XET
/-show/showfile/cue/265  <SCUE> n=0
    /-show/showfile/cue/265/numb  I32 F_XET
    /-show/showfile/cue/265/name  S32 F_XET
    /-show/showfile/cue/265/skip  I32 F_XET
    /-show/showfile/cue/265/scene  I32 F_XET
    /-show/showfile/cue/265/bit  I32 F_XET
    /-show/showfile/cue/265/miditype  I32 F_XET
    /-show/showfile/cue/265/midichan  I32 F_XET
    /-show/showfile/cue/265/midipara1  I32 F_XET
    /-show/showfile/cue/265/midipara2  I32 F_XET
/-show/showfile/cue/266  <SCUE> n=0
    /-show/showfile/cue/266/numb  I32 F_XET
    /-show/showfile/cue/266/name  S32 F_XET
    /-show/showfile/cue/266/skip  I32 F_XET
    /-show/showfile/cue/266/scene  I32 F_XET
    /-show/showfile/cue/266/bit  I32 F_XET
    /-show/showfile/cue/266/miditype  I32 F_XET
    /-show/showfile/cue/266/midichan  I32 F_XET
    /-show/showfile/cue/266/midipara1  I32 F_XET
    /-show/showfile/cue/266/midipara2  I32 F_XET
/-show/showfile/cue/267  <SCUE> n=0
    /-show/showfile/cue/267/numb  I32 F_XET
    /-show/showfile/cue/267/name  S32 F_XET
    /-show/showfile/cue/267/skip  I32 F_XET
    /-show/showfile/cue/267/scene  I32 F_XET
    /-show/showfile/cue/267/bit  I32 F_XET
    /-show/showfile/cue/267/miditype  I32 F_XET
    /-show/showfile/cue/267/midichan  I32 F_XET
    /-show/showfile/cue/267/midipara1  I32 F_XET
    /-show/showfile/cue/267/midipara2  I32 F_XET
/-show/showfile/cue/268  <SCUE> n=0
    /-show/showfile/cue/268/numb  I32 F_XET
    /-show/showfile/cue/268/name  S32 F_XET
    /-show/showfile/cue/268/skip  I32 F_XET
    /-show/showfile/cue/268/scene  I32 F_XET
    /-show/showfile/cue/268/bit  I32 F_XET
    /-show/showfile/cue/268/miditype  I32 F_XET
    /-show/showfile/cue/268/midichan  I32 F_XET
    /-show/showfile/cue/268/midipara1  I32 F_XET
    /-show/showfile/cue/268/midipara2  I32 F_XET
/-show/showfile/cue/269  <SCUE> n=0
    /-show/showfile/cue/269/numb  I32 F_XET
    /-show/showfile/cue/269/name  S32 F_XET
    /-show/showfile/cue/269/skip  I32 F_XET
    /-show/showfile/cue/269/scene  I32 F_XET
    /-show/showfile/cue/269/bit  I32 F_XET
    /-show/showfile/cue/269/miditype  I32 F_XET
    /-show/showfile/cue/269/midichan  I32 F_XET
    /-show/showfile/cue/269/midipara1  I32 F_XET
    /-show/showfile/cue/269/midipara2  I32 F_XET
/-show/showfile/cue/270  <SCUE> n=0
    /-show/showfile/cue/270/numb  I32 F_XET
    /-show/showfile/cue/270/name  S32 F_XET
    /-show/showfile/cue/270/skip  I32 F_XET
    /-show/showfile/cue/270/scene  I32 F_XET
    /-show/showfile/cue/270/bit  I32 F_XET
    /-show/showfile/cue/270/miditype  I32 F_XET
    /-show/showfile/cue/270/midichan  I32 F_XET
    /-show/showfile/cue/270/midipara1  I32 F_XET
    /-show/showfile/cue/270/midipara2  I32 F_XET
/-show/showfile/cue/271  <SCUE> n=0
    /-show/showfile/cue/271/numb  I32 F_XET
    /-show/showfile/cue/271/name  S32 F_XET
    /-show/showfile/cue/271/skip  I32 F_XET
    /-show/showfile/cue/271/scene  I32 F_XET
    /-show/showfile/cue/271/bit  I32 F_XET
    /-show/showfile/cue/271/miditype  I32 F_XET
    /-show/showfile/cue/271/midichan  I32 F_XET
    /-show/showfile/cue/271/midipara1  I32 F_XET
    /-show/showfile/cue/271/midipara2  I32 F_XET
/-show/showfile/cue/272  <SCUE> n=0
    /-show/showfile/cue/272/numb  I32 F_XET
    /-show/showfile/cue/272/name  S32 F_XET
    /-show/showfile/cue/272/skip  I32 F_XET
    /-show/showfile/cue/272/scene  I32 F_XET
    /-show/showfile/cue/272/bit  I32 F_XET
    /-show/showfile/cue/272/miditype  I32 F_XET
    /-show/showfile/cue/272/midichan  I32 F_XET
    /-show/showfile/cue/272/midipara1  I32 F_XET
    /-show/showfile/cue/272/midipara2  I32 F_XET
/-show/showfile/cue/273  <SCUE> n=0
    /-show/showfile/cue/273/numb  I32 F_XET
    /-show/showfile/cue/273/name  S32 F_XET
    /-show/showfile/cue/273/skip  I32 F_XET
    /-show/showfile/cue/273/scene  I32 F_XET
    /-show/showfile/cue/273/bit  I32 F_XET
    /-show/showfile/cue/273/miditype  I32 F_XET
    /-show/showfile/cue/273/midichan  I32 F_XET
    /-show/showfile/cue/273/midipara1  I32 F_XET
    /-show/showfile/cue/273/midipara2  I32 F_XET
/-show/showfile/cue/274  <SCUE> n=0
    /-show/showfile/cue/274/numb  I32 F_XET
    /-show/showfile/cue/274/name  S32 F_XET
    /-show/showfile/cue/274/skip  I32 F_XET
    /-show/showfile/cue/274/scene  I32 F_XET
    /-show/showfile/cue/274/bit  I32 F_XET
    /-show/showfile/cue/274/miditype  I32 F_XET
    /-show/showfile/cue/274/midichan  I32 F_XET
    /-show/showfile/cue/274/midipara1  I32 F_XET
    /-show/showfile/cue/274/midipara2  I32 F_XET
/-show/showfile/cue/275  <SCUE> n=0
    /-show/showfile/cue/275/numb  I32 F_XET
    /-show/showfile/cue/275/name  S32 F_XET
    /-show/showfile/cue/275/skip  I32 F_XET
    /-show/showfile/cue/275/scene  I32 F_XET
    /-show/showfile/cue/275/bit  I32 F_XET
    /-show/showfile/cue/275/miditype  I32 F_XET
    /-show/showfile/cue/275/midichan  I32 F_XET
    /-show/showfile/cue/275/midipara1  I32 F_XET
    /-show/showfile/cue/275/midipara2  I32 F_XET
/-show/showfile/cue/276  <SCUE> n=0
    /-show/showfile/cue/276/numb  I32 F_XET
    /-show/showfile/cue/276/name  S32 F_XET
    /-show/showfile/cue/276/skip  I32 F_XET
    /-show/showfile/cue/276/scene  I32 F_XET
    /-show/showfile/cue/276/bit  I32 F_XET
    /-show/showfile/cue/276/miditype  I32 F_XET
    /-show/showfile/cue/276/midichan  I32 F_XET
    /-show/showfile/cue/276/midipara1  I32 F_XET
    /-show/showfile/cue/276/midipara2  I32 F_XET
/-show/showfile/cue/277  <SCUE> n=0
    /-show/showfile/cue/277/numb  I32 F_XET
    /-show/showfile/cue/277/name  S32 F_XET
    /-show/showfile/cue/277/skip  I32 F_XET
    /-show/showfile/cue/277/scene  I32 F_XET
    /-show/showfile/cue/277/bit  I32 F_XET
    /-show/showfile/cue/277/miditype  I32 F_XET
    /-show/showfile/cue/277/midichan  I32 F_XET
    /-show/showfile/cue/277/midipara1  I32 F_XET
    /-show/showfile/cue/277/midipara2  I32 F_XET
/-show/showfile/cue/278  <SCUE> n=0
    /-show/showfile/cue/278/numb  I32 F_XET
    /-show/showfile/cue/278/name  S32 F_XET
    /-show/showfile/cue/278/skip  I32 F_XET
    /-show/showfile/cue/278/scene  I32 F_XET
    /-show/showfile/cue/278/bit  I32 F_XET
    /-show/showfile/cue/278/miditype  I32 F_XET
    /-show/showfile/cue/278/midichan  I32 F_XET
    /-show/showfile/cue/278/midipara1  I32 F_XET
    /-show/showfile/cue/278/midipara2  I32 F_XET
/-show/showfile/cue/279  <SCUE> n=0
    /-show/showfile/cue/279/numb  I32 F_XET
    /-show/showfile/cue/279/name  S32 F_XET
    /-show/showfile/cue/279/skip  I32 F_XET
    /-show/showfile/cue/279/scene  I32 F_XET
    /-show/showfile/cue/279/bit  I32 F_XET
    /-show/showfile/cue/279/miditype  I32 F_XET
    /-show/showfile/cue/279/midichan  I32 F_XET
    /-show/showfile/cue/279/midipara1  I32 F_XET
    /-show/showfile/cue/279/midipara2  I32 F_XET
/-show/showfile/cue/280  <SCUE> n=0
    /-show/showfile/cue/280/numb  I32 F_XET
    /-show/showfile/cue/280/name  S32 F_XET
    /-show/showfile/cue/280/skip  I32 F_XET
    /-show/showfile/cue/280/scene  I32 F_XET
    /-show/showfile/cue/280/bit  I32 F_XET
    /-show/showfile/cue/280/miditype  I32 F_XET
    /-show/showfile/cue/280/midichan  I32 F_XET
    /-show/showfile/cue/280/midipara1  I32 F_XET
    /-show/showfile/cue/280/midipara2  I32 F_XET
/-show/showfile/cue/281  <SCUE> n=0
    /-show/showfile/cue/281/numb  I32 F_XET
    /-show/showfile/cue/281/name  S32 F_XET
    /-show/showfile/cue/281/skip  I32 F_XET
    /-show/showfile/cue/281/scene  I32 F_XET
    /-show/showfile/cue/281/bit  I32 F_XET
    /-show/showfile/cue/281/miditype  I32 F_XET
    /-show/showfile/cue/281/midichan  I32 F_XET
    /-show/showfile/cue/281/midipara1  I32 F_XET
    /-show/showfile/cue/281/midipara2  I32 F_XET
/-show/showfile/cue/282  <SCUE> n=0
    /-show/showfile/cue/282/numb  I32 F_XET
    /-show/showfile/cue/282/name  S32 F_XET
    /-show/showfile/cue/282/skip  I32 F_XET
    /-show/showfile/cue/282/scene  I32 F_XET
    /-show/showfile/cue/282/bit  I32 F_XET
    /-show/showfile/cue/282/miditype  I32 F_XET
    /-show/showfile/cue/282/midichan  I32 F_XET
    /-show/showfile/cue/282/midipara1  I32 F_XET
    /-show/showfile/cue/282/midipara2  I32 F_XET
/-show/showfile/cue/283  <SCUE> n=0
    /-show/showfile/cue/283/numb  I32 F_XET
    /-show/showfile/cue/283/name  S32 F_XET
    /-show/showfile/cue/283/skip  I32 F_XET
    /-show/showfile/cue/283/scene  I32 F_XET
    /-show/showfile/cue/283/bit  I32 F_XET
    /-show/showfile/cue/283/miditype  I32 F_XET
    /-show/showfile/cue/283/midichan  I32 F_XET
    /-show/showfile/cue/283/midipara1  I32 F_XET
    /-show/showfile/cue/283/midipara2  I32 F_XET
/-show/showfile/cue/284  <SCUE> n=0
    /-show/showfile/cue/284/numb  I32 F_XET
    /-show/showfile/cue/284/name  S32 F_XET
    /-show/showfile/cue/284/skip  I32 F_XET
    /-show/showfile/cue/284/scene  I32 F_XET
    /-show/showfile/cue/284/bit  I32 F_XET
    /-show/showfile/cue/284/miditype  I32 F_XET
    /-show/showfile/cue/284/midichan  I32 F_XET
    /-show/showfile/cue/284/midipara1  I32 F_XET
    /-show/showfile/cue/284/midipara2  I32 F_XET
/-show/showfile/cue/285  <SCUE> n=0
    /-show/showfile/cue/285/numb  I32 F_XET
    /-show/showfile/cue/285/name  S32 F_XET
    /-show/showfile/cue/285/skip  I32 F_XET
    /-show/showfile/cue/285/scene  I32 F_XET
    /-show/showfile/cue/285/bit  I32 F_XET
    /-show/showfile/cue/285/miditype  I32 F_XET
    /-show/showfile/cue/285/midichan  I32 F_XET
    /-show/showfile/cue/285/midipara1  I32 F_XET
    /-show/showfile/cue/285/midipara2  I32 F_XET
/-show/showfile/cue/286  <SCUE> n=0
    /-show/showfile/cue/286/numb  I32 F_XET
    /-show/showfile/cue/286/name  S32 F_XET
    /-show/showfile/cue/286/skip  I32 F_XET
    /-show/showfile/cue/286/scene  I32 F_XET
    /-show/showfile/cue/286/bit  I32 F_XET
    /-show/showfile/cue/286/miditype  I32 F_XET
    /-show/showfile/cue/286/midichan  I32 F_XET
    /-show/showfile/cue/286/midipara1  I32 F_XET
    /-show/showfile/cue/286/midipara2  I32 F_XET
/-show/showfile/cue/287  <SCUE> n=0
    /-show/showfile/cue/287/numb  I32 F_XET
    /-show/showfile/cue/287/name  S32 F_XET
    /-show/showfile/cue/287/skip  I32 F_XET
    /-show/showfile/cue/287/scene  I32 F_XET
    /-show/showfile/cue/287/bit  I32 F_XET
    /-show/showfile/cue/287/miditype  I32 F_XET
    /-show/showfile/cue/287/midichan  I32 F_XET
    /-show/showfile/cue/287/midipara1  I32 F_XET
    /-show/showfile/cue/287/midipara2  I32 F_XET
/-show/showfile/cue/288  <SCUE> n=0
    /-show/showfile/cue/288/numb  I32 F_XET
    /-show/showfile/cue/288/name  S32 F_XET
    /-show/showfile/cue/288/skip  I32 F_XET
    /-show/showfile/cue/288/scene  I32 F_XET
    /-show/showfile/cue/288/bit  I32 F_XET
    /-show/showfile/cue/288/miditype  I32 F_XET
    /-show/showfile/cue/288/midichan  I32 F_XET
    /-show/showfile/cue/288/midipara1  I32 F_XET
    /-show/showfile/cue/288/midipara2  I32 F_XET
/-show/showfile/cue/289  <SCUE> n=0
    /-show/showfile/cue/289/numb  I32 F_XET
    /-show/showfile/cue/289/name  S32 F_XET
    /-show/showfile/cue/289/skip  I32 F_XET
    /-show/showfile/cue/289/scene  I32 F_XET
    /-show/showfile/cue/289/bit  I32 F_XET
    /-show/showfile/cue/289/miditype  I32 F_XET
    /-show/showfile/cue/289/midichan  I32 F_XET
    /-show/showfile/cue/289/midipara1  I32 F_XET
    /-show/showfile/cue/289/midipara2  I32 F_XET
/-show/showfile/cue/290  <SCUE> n=0
    /-show/showfile/cue/290/numb  I32 F_XET
    /-show/showfile/cue/290/name  S32 F_XET
    /-show/showfile/cue/290/skip  I32 F_XET
    /-show/showfile/cue/290/scene  I32 F_XET
    /-show/showfile/cue/290/bit  I32 F_XET
    /-show/showfile/cue/290/miditype  I32 F_XET
    /-show/showfile/cue/290/midichan  I32 F_XET
    /-show/showfile/cue/290/midipara1  I32 F_XET
    /-show/showfile/cue/290/midipara2  I32 F_XET
/-show/showfile/cue/291  <SCUE> n=0
    /-show/showfile/cue/291/numb  I32 F_XET
    /-show/showfile/cue/291/name  S32 F_XET
    /-show/showfile/cue/291/skip  I32 F_XET
    /-show/showfile/cue/291/scene  I32 F_XET
    /-show/showfile/cue/291/bit  I32 F_XET
    /-show/showfile/cue/291/miditype  I32 F_XET
    /-show/showfile/cue/291/midichan  I32 F_XET
    /-show/showfile/cue/291/midipara1  I32 F_XET
    /-show/showfile/cue/291/midipara2  I32 F_XET
/-show/showfile/cue/292  <SCUE> n=0
    /-show/showfile/cue/292/numb  I32 F_XET
    /-show/showfile/cue/292/name  S32 F_XET
    /-show/showfile/cue/292/skip  I32 F_XET
    /-show/showfile/cue/292/scene  I32 F_XET
    /-show/showfile/cue/292/bit  I32 F_XET
    /-show/showfile/cue/292/miditype  I32 F_XET
    /-show/showfile/cue/292/midichan  I32 F_XET
    /-show/showfile/cue/292/midipara1  I32 F_XET
    /-show/showfile/cue/292/midipara2  I32 F_XET
/-show/showfile/cue/293  <SCUE> n=0
    /-show/showfile/cue/293/numb  I32 F_XET
    /-show/showfile/cue/293/name  S32 F_XET
    /-show/showfile/cue/293/skip  I32 F_XET
    /-show/showfile/cue/293/scene  I32 F_XET
    /-show/showfile/cue/293/bit  I32 F_XET
    /-show/showfile/cue/293/miditype  I32 F_XET
    /-show/showfile/cue/293/midichan  I32 F_XET
    /-show/showfile/cue/293/midipara1  I32 F_XET
    /-show/showfile/cue/293/midipara2  I32 F_XET
/-show/showfile/cue/294  <SCUE> n=0
    /-show/showfile/cue/294/numb  I32 F_XET
    /-show/showfile/cue/294/name  S32 F_XET
    /-show/showfile/cue/294/skip  I32 F_XET
    /-show/showfile/cue/294/scene  I32 F_XET
    /-show/showfile/cue/294/bit  I32 F_XET
    /-show/showfile/cue/294/miditype  I32 F_XET
    /-show/showfile/cue/294/midichan  I32 F_XET
    /-show/showfile/cue/294/midipara1  I32 F_XET
    /-show/showfile/cue/294/midipara2  I32 F_XET
/-show/showfile/cue/295  <SCUE> n=0
    /-show/showfile/cue/295/numb  I32 F_XET
    /-show/showfile/cue/295/name  S32 F_XET
    /-show/showfile/cue/295/skip  I32 F_XET
    /-show/showfile/cue/295/scene  I32 F_XET
    /-show/showfile/cue/295/bit  I32 F_XET
    /-show/showfile/cue/295/miditype  I32 F_XET
    /-show/showfile/cue/295/midichan  I32 F_XET
    /-show/showfile/cue/295/midipara1  I32 F_XET
    /-show/showfile/cue/295/midipara2  I32 F_XET
/-show/showfile/cue/296  <SCUE> n=0
    /-show/showfile/cue/296/numb  I32 F_XET
    /-show/showfile/cue/296/name  S32 F_XET
    /-show/showfile/cue/296/skip  I32 F_XET
    /-show/showfile/cue/296/scene  I32 F_XET
    /-show/showfile/cue/296/bit  I32 F_XET
    /-show/showfile/cue/296/miditype  I32 F_XET
    /-show/showfile/cue/296/midichan  I32 F_XET
    /-show/showfile/cue/296/midipara1  I32 F_XET
    /-show/showfile/cue/296/midipara2  I32 F_XET
/-show/showfile/cue/297  <SCUE> n=0
    /-show/showfile/cue/297/numb  I32 F_XET
    /-show/showfile/cue/297/name  S32 F_XET
    /-show/showfile/cue/297/skip  I32 F_XET
    /-show/showfile/cue/297/scene  I32 F_XET
    /-show/showfile/cue/297/bit  I32 F_XET
    /-show/showfile/cue/297/miditype  I32 F_XET
    /-show/showfile/cue/297/midichan  I32 F_XET
    /-show/showfile/cue/297/midipara1  I32 F_XET
    /-show/showfile/cue/297/midipara2  I32 F_XET
/-show/showfile/cue/298  <SCUE> n=0
    /-show/showfile/cue/298/numb  I32 F_XET
    /-show/showfile/cue/298/name  S32 F_XET
    /-show/showfile/cue/298/skip  I32 F_XET
    /-show/showfile/cue/298/scene  I32 F_XET
    /-show/showfile/cue/298/bit  I32 F_XET
    /-show/showfile/cue/298/miditype  I32 F_XET
    /-show/showfile/cue/298/midichan  I32 F_XET
    /-show/showfile/cue/298/midipara1  I32 F_XET
    /-show/showfile/cue/298/midipara2  I32 F_XET
/-show/showfile/cue/299  <SCUE> n=0
    /-show/showfile/cue/299/numb  I32 F_XET
    /-show/showfile/cue/299/name  S32 F_XET
    /-show/showfile/cue/299/skip  I32 F_XET
    /-show/showfile/cue/299/scene  I32 F_XET
    /-show/showfile/cue/299/bit  I32 F_XET
    /-show/showfile/cue/299/miditype  I32 F_XET
    /-show/showfile/cue/299/midichan  I32 F_XET
    /-show/showfile/cue/299/midipara1  I32 F_XET
    /-show/showfile/cue/299/midipara2  I32 F_XET
/-show/showfile/cue/300  <SCUE> n=0
    /-show/showfile/cue/300/numb  I32 F_XET
    /-show/showfile/cue/300/name  S32 F_XET
    /-show/showfile/cue/300/skip  I32 F_XET
    /-show/showfile/cue/300/scene  I32 F_XET
    /-show/showfile/cue/300/bit  I32 F_XET
    /-show/showfile/cue/300/miditype  I32 F_XET
    /-show/showfile/cue/300/midichan  I32 F_XET
    /-show/showfile/cue/300/midipara1  I32 F_XET
    /-show/showfile/cue/300/midipara2  I32 F_XET
/-show/showfile/cue/301  <SCUE> n=0
    /-show/showfile/cue/301/numb  I32 F_XET
    /-show/showfile/cue/301/name  S32 F_XET
    /-show/showfile/cue/301/skip  I32 F_XET
    /-show/showfile/cue/301/scene  I32 F_XET
    /-show/showfile/cue/301/bit  I32 F_XET
    /-show/showfile/cue/301/miditype  I32 F_XET
    /-show/showfile/cue/301/midichan  I32 F_XET
    /-show/showfile/cue/301/midipara1  I32 F_XET
    /-show/showfile/cue/301/midipara2  I32 F_XET
/-show/showfile/cue/302  <SCUE> n=0
    /-show/showfile/cue/302/numb  I32 F_XET
    /-show/showfile/cue/302/name  S32 F_XET
    /-show/showfile/cue/302/skip  I32 F_XET
    /-show/showfile/cue/302/scene  I32 F_XET
    /-show/showfile/cue/302/bit  I32 F_XET
    /-show/showfile/cue/302/miditype  I32 F_XET
    /-show/showfile/cue/302/midichan  I32 F_XET
    /-show/showfile/cue/302/midipara1  I32 F_XET
    /-show/showfile/cue/302/midipara2  I32 F_XET
/-show/showfile/cue/303  <SCUE> n=0
    /-show/showfile/cue/303/numb  I32 F_XET
    /-show/showfile/cue/303/name  S32 F_XET
    /-show/showfile/cue/303/skip  I32 F_XET
    /-show/showfile/cue/303/scene  I32 F_XET
    /-show/showfile/cue/303/bit  I32 F_XET
    /-show/showfile/cue/303/miditype  I32 F_XET
    /-show/showfile/cue/303/midichan  I32 F_XET
    /-show/showfile/cue/303/midipara1  I32 F_XET
    /-show/showfile/cue/303/midipara2  I32 F_XET
/-show/showfile/cue/304  <SCUE> n=0
    /-show/showfile/cue/304/numb  I32 F_XET
    /-show/showfile/cue/304/name  S32 F_XET
    /-show/showfile/cue/304/skip  I32 F_XET
    /-show/showfile/cue/304/scene  I32 F_XET
    /-show/showfile/cue/304/bit  I32 F_XET
    /-show/showfile/cue/304/miditype  I32 F_XET
    /-show/showfile/cue/304/midichan  I32 F_XET
    /-show/showfile/cue/304/midipara1  I32 F_XET
    /-show/showfile/cue/304/midipara2  I32 F_XET
/-show/showfile/cue/305  <SCUE> n=0
    /-show/showfile/cue/305/numb  I32 F_XET
    /-show/showfile/cue/305/name  S32 F_XET
    /-show/showfile/cue/305/skip  I32 F_XET
    /-show/showfile/cue/305/scene  I32 F_XET
    /-show/showfile/cue/305/bit  I32 F_XET
    /-show/showfile/cue/305/miditype  I32 F_XET
    /-show/showfile/cue/305/midichan  I32 F_XET
    /-show/showfile/cue/305/midipara1  I32 F_XET
    /-show/showfile/cue/305/midipara2  I32 F_XET
/-show/showfile/cue/306  <SCUE> n=0
    /-show/showfile/cue/306/numb  I32 F_XET
    /-show/showfile/cue/306/name  S32 F_XET
    /-show/showfile/cue/306/skip  I32 F_XET
    /-show/showfile/cue/306/scene  I32 F_XET
    /-show/showfile/cue/306/bit  I32 F_XET
    /-show/showfile/cue/306/miditype  I32 F_XET
    /-show/showfile/cue/306/midichan  I32 F_XET
    /-show/showfile/cue/306/midipara1  I32 F_XET
    /-show/showfile/cue/306/midipara2  I32 F_XET
/-show/showfile/cue/307  <SCUE> n=0
    /-show/showfile/cue/307/numb  I32 F_XET
    /-show/showfile/cue/307/name  S32 F_XET
    /-show/showfile/cue/307/skip  I32 F_XET
    /-show/showfile/cue/307/scene  I32 F_XET
    /-show/showfile/cue/307/bit  I32 F_XET
    /-show/showfile/cue/307/miditype  I32 F_XET
    /-show/showfile/cue/307/midichan  I32 F_XET
    /-show/showfile/cue/307/midipara1  I32 F_XET
    /-show/showfile/cue/307/midipara2  I32 F_XET
/-show/showfile/cue/308  <SCUE> n=0
    /-show/showfile/cue/308/numb  I32 F_XET
    /-show/showfile/cue/308/name  S32 F_XET
    /-show/showfile/cue/308/skip  I32 F_XET
    /-show/showfile/cue/308/scene  I32 F_XET
    /-show/showfile/cue/308/bit  I32 F_XET
    /-show/showfile/cue/308/miditype  I32 F_XET
    /-show/showfile/cue/308/midichan  I32 F_XET
    /-show/showfile/cue/308/midipara1  I32 F_XET
    /-show/showfile/cue/308/midipara2  I32 F_XET
/-show/showfile/cue/309  <SCUE> n=0
    /-show/showfile/cue/309/numb  I32 F_XET
    /-show/showfile/cue/309/name  S32 F_XET
    /-show/showfile/cue/309/skip  I32 F_XET
    /-show/showfile/cue/309/scene  I32 F_XET
    /-show/showfile/cue/309/bit  I32 F_XET
    /-show/showfile/cue/309/miditype  I32 F_XET
    /-show/showfile/cue/309/midichan  I32 F_XET
    /-show/showfile/cue/309/midipara1  I32 F_XET
    /-show/showfile/cue/309/midipara2  I32 F_XET
/-show/showfile/cue/310  <SCUE> n=0
    /-show/showfile/cue/310/numb  I32 F_XET
    /-show/showfile/cue/310/name  S32 F_XET
    /-show/showfile/cue/310/skip  I32 F_XET
    /-show/showfile/cue/310/scene  I32 F_XET
    /-show/showfile/cue/310/bit  I32 F_XET
    /-show/showfile/cue/310/miditype  I32 F_XET
    /-show/showfile/cue/310/midichan  I32 F_XET
    /-show/showfile/cue/310/midipara1  I32 F_XET
    /-show/showfile/cue/310/midipara2  I32 F_XET
/-show/showfile/cue/311  <SCUE> n=0
    /-show/showfile/cue/311/numb  I32 F_XET
    /-show/showfile/cue/311/name  S32 F_XET
    /-show/showfile/cue/311/skip  I32 F_XET
    /-show/showfile/cue/311/scene  I32 F_XET
    /-show/showfile/cue/311/bit  I32 F_XET
    /-show/showfile/cue/311/miditype  I32 F_XET
    /-show/showfile/cue/311/midichan  I32 F_XET
    /-show/showfile/cue/311/midipara1  I32 F_XET
    /-show/showfile/cue/311/midipara2  I32 F_XET
/-show/showfile/cue/312  <SCUE> n=0
    /-show/showfile/cue/312/numb  I32 F_XET
    /-show/showfile/cue/312/name  S32 F_XET
    /-show/showfile/cue/312/skip  I32 F_XET
    /-show/showfile/cue/312/scene  I32 F_XET
    /-show/showfile/cue/312/bit  I32 F_XET
    /-show/showfile/cue/312/miditype  I32 F_XET
    /-show/showfile/cue/312/midichan  I32 F_XET
    /-show/showfile/cue/312/midipara1  I32 F_XET
    /-show/showfile/cue/312/midipara2  I32 F_XET
/-show/showfile/cue/313  <SCUE> n=0
    /-show/showfile/cue/313/numb  I32 F_XET
    /-show/showfile/cue/313/name  S32 F_XET
    /-show/showfile/cue/313/skip  I32 F_XET
    /-show/showfile/cue/313/scene  I32 F_XET
    /-show/showfile/cue/313/bit  I32 F_XET
    /-show/showfile/cue/313/miditype  I32 F_XET
    /-show/showfile/cue/313/midichan  I32 F_XET
    /-show/showfile/cue/313/midipara1  I32 F_XET
    /-show/showfile/cue/313/midipara2  I32 F_XET
/-show/showfile/cue/314  <SCUE> n=0
    /-show/showfile/cue/314/numb  I32 F_XET
    /-show/showfile/cue/314/name  S32 F_XET
    /-show/showfile/cue/314/skip  I32 F_XET
    /-show/showfile/cue/314/scene  I32 F_XET
    /-show/showfile/cue/314/bit  I32 F_XET
    /-show/showfile/cue/314/miditype  I32 F_XET
    /-show/showfile/cue/314/midichan  I32 F_XET
    /-show/showfile/cue/314/midipara1  I32 F_XET
    /-show/showfile/cue/314/midipara2  I32 F_XET
/-show/showfile/cue/315  <SCUE> n=0
    /-show/showfile/cue/315/numb  I32 F_XET
    /-show/showfile/cue/315/name  S32 F_XET
    /-show/showfile/cue/315/skip  I32 F_XET
    /-show/showfile/cue/315/scene  I32 F_XET
    /-show/showfile/cue/315/bit  I32 F_XET
    /-show/showfile/cue/315/miditype  I32 F_XET
    /-show/showfile/cue/315/midichan  I32 F_XET
    /-show/showfile/cue/315/midipara1  I32 F_XET
    /-show/showfile/cue/315/midipara2  I32 F_XET
/-show/showfile/cue/316  <SCUE> n=0
    /-show/showfile/cue/316/numb  I32 F_XET
    /-show/showfile/cue/316/name  S32 F_XET
    /-show/showfile/cue/316/skip  I32 F_XET
    /-show/showfile/cue/316/scene  I32 F_XET
    /-show/showfile/cue/316/bit  I32 F_XET
    /-show/showfile/cue/316/miditype  I32 F_XET
    /-show/showfile/cue/316/midichan  I32 F_XET
    /-show/showfile/cue/316/midipara1  I32 F_XET
    /-show/showfile/cue/316/midipara2  I32 F_XET
/-show/showfile/cue/317  <SCUE> n=0
    /-show/showfile/cue/317/numb  I32 F_XET
    /-show/showfile/cue/317/name  S32 F_XET
    /-show/showfile/cue/317/skip  I32 F_XET
    /-show/showfile/cue/317/scene  I32 F_XET
    /-show/showfile/cue/317/bit  I32 F_XET
    /-show/showfile/cue/317/miditype  I32 F_XET
    /-show/showfile/cue/317/midichan  I32 F_XET
    /-show/showfile/cue/317/midipara1  I32 F_XET
    /-show/showfile/cue/317/midipara2  I32 F_XET
/-show/showfile/cue/318  <SCUE> n=0
    /-show/showfile/cue/318/numb  I32 F_XET
    /-show/showfile/cue/318/name  S32 F_XET
    /-show/showfile/cue/318/skip  I32 F_XET
    /-show/showfile/cue/318/scene  I32 F_XET
    /-show/showfile/cue/318/bit  I32 F_XET
    /-show/showfile/cue/318/miditype  I32 F_XET
    /-show/showfile/cue/318/midichan  I32 F_XET
    /-show/showfile/cue/318/midipara1  I32 F_XET
    /-show/showfile/cue/318/midipara2  I32 F_XET
/-show/showfile/cue/319  <SCUE> n=0
    /-show/showfile/cue/319/numb  I32 F_XET
    /-show/showfile/cue/319/name  S32 F_XET
    /-show/showfile/cue/319/skip  I32 F_XET
    /-show/showfile/cue/319/scene  I32 F_XET
    /-show/showfile/cue/319/bit  I32 F_XET
    /-show/showfile/cue/319/miditype  I32 F_XET
    /-show/showfile/cue/319/midichan  I32 F_XET
    /-show/showfile/cue/319/midipara1  I32 F_XET
    /-show/showfile/cue/319/midipara2  I32 F_XET
/-show/showfile/cue/320  <SCUE> n=0
    /-show/showfile/cue/320/numb  I32 F_XET
    /-show/showfile/cue/320/name  S32 F_XET
    /-show/showfile/cue/320/skip  I32 F_XET
    /-show/showfile/cue/320/scene  I32 F_XET
    /-show/showfile/cue/320/bit  I32 F_XET
    /-show/showfile/cue/320/miditype  I32 F_XET
    /-show/showfile/cue/320/midichan  I32 F_XET
    /-show/showfile/cue/320/midipara1  I32 F_XET
    /-show/showfile/cue/320/midipara2  I32 F_XET
/-show/showfile/cue/321  <SCUE> n=0
    /-show/showfile/cue/321/numb  I32 F_XET
    /-show/showfile/cue/321/name  S32 F_XET
    /-show/showfile/cue/321/skip  I32 F_XET
    /-show/showfile/cue/321/scene  I32 F_XET
    /-show/showfile/cue/321/bit  I32 F_XET
    /-show/showfile/cue/321/miditype  I32 F_XET
    /-show/showfile/cue/321/midichan  I32 F_XET
    /-show/showfile/cue/321/midipara1  I32 F_XET
    /-show/showfile/cue/321/midipara2  I32 F_XET
/-show/showfile/cue/322  <SCUE> n=0
    /-show/showfile/cue/322/numb  I32 F_XET
    /-show/showfile/cue/322/name  S32 F_XET
    /-show/showfile/cue/322/skip  I32 F_XET
    /-show/showfile/cue/322/scene  I32 F_XET
    /-show/showfile/cue/322/bit  I32 F_XET
    /-show/showfile/cue/322/miditype  I32 F_XET
    /-show/showfile/cue/322/midichan  I32 F_XET
    /-show/showfile/cue/322/midipara1  I32 F_XET
    /-show/showfile/cue/322/midipara2  I32 F_XET
/-show/showfile/cue/323  <SCUE> n=0
    /-show/showfile/cue/323/numb  I32 F_XET
    /-show/showfile/cue/323/name  S32 F_XET
    /-show/showfile/cue/323/skip  I32 F_XET
    /-show/showfile/cue/323/scene  I32 F_XET
    /-show/showfile/cue/323/bit  I32 F_XET
    /-show/showfile/cue/323/miditype  I32 F_XET
    /-show/showfile/cue/323/midichan  I32 F_XET
    /-show/showfile/cue/323/midipara1  I32 F_XET
    /-show/showfile/cue/323/midipara2  I32 F_XET
/-show/showfile/cue/324  <SCUE> n=0
    /-show/showfile/cue/324/numb  I32 F_XET
    /-show/showfile/cue/324/name  S32 F_XET
    /-show/showfile/cue/324/skip  I32 F_XET
    /-show/showfile/cue/324/scene  I32 F_XET
    /-show/showfile/cue/324/bit  I32 F_XET
    /-show/showfile/cue/324/miditype  I32 F_XET
    /-show/showfile/cue/324/midichan  I32 F_XET
    /-show/showfile/cue/324/midipara1  I32 F_XET
    /-show/showfile/cue/324/midipara2  I32 F_XET
/-show/showfile/cue/325  <SCUE> n=0
    /-show/showfile/cue/325/numb  I32 F_XET
    /-show/showfile/cue/325/name  S32 F_XET
    /-show/showfile/cue/325/skip  I32 F_XET
    /-show/showfile/cue/325/scene  I32 F_XET
    /-show/showfile/cue/325/bit  I32 F_XET
    /-show/showfile/cue/325/miditype  I32 F_XET
    /-show/showfile/cue/325/midichan  I32 F_XET
    /-show/showfile/cue/325/midipara1  I32 F_XET
    /-show/showfile/cue/325/midipara2  I32 F_XET
/-show/showfile/cue/326  <SCUE> n=0
    /-show/showfile/cue/326/numb  I32 F_XET
    /-show/showfile/cue/326/name  S32 F_XET
    /-show/showfile/cue/326/skip  I32 F_XET
    /-show/showfile/cue/326/scene  I32 F_XET
    /-show/showfile/cue/326/bit  I32 F_XET
    /-show/showfile/cue/326/miditype  I32 F_XET
    /-show/showfile/cue/326/midichan  I32 F_XET
    /-show/showfile/cue/326/midipara1  I32 F_XET
    /-show/showfile/cue/326/midipara2  I32 F_XET
/-show/showfile/cue/327  <SCUE> n=0
    /-show/showfile/cue/327/numb  I32 F_XET
    /-show/showfile/cue/327/name  S32 F_XET
    /-show/showfile/cue/327/skip  I32 F_XET
    /-show/showfile/cue/327/scene  I32 F_XET
    /-show/showfile/cue/327/bit  I32 F_XET
    /-show/showfile/cue/327/miditype  I32 F_XET
    /-show/showfile/cue/327/midichan  I32 F_XET
    /-show/showfile/cue/327/midipara1  I32 F_XET
    /-show/showfile/cue/327/midipara2  I32 F_XET
/-show/showfile/cue/328  <SCUE> n=0
    /-show/showfile/cue/328/numb  I32 F_XET
    /-show/showfile/cue/328/name  S32 F_XET
    /-show/showfile/cue/328/skip  I32 F_XET
    /-show/showfile/cue/328/scene  I32 F_XET
    /-show/showfile/cue/328/bit  I32 F_XET
    /-show/showfile/cue/328/miditype  I32 F_XET
    /-show/showfile/cue/328/midichan  I32 F_XET
    /-show/showfile/cue/328/midipara1  I32 F_XET
    /-show/showfile/cue/328/midipara2  I32 F_XET
/-show/showfile/cue/329  <SCUE> n=0
    /-show/showfile/cue/329/numb  I32 F_XET
    /-show/showfile/cue/329/name  S32 F_XET
    /-show/showfile/cue/329/skip  I32 F_XET
    /-show/showfile/cue/329/scene  I32 F_XET
    /-show/showfile/cue/329/bit  I32 F_XET
    /-show/showfile/cue/329/miditype  I32 F_XET
    /-show/showfile/cue/329/midichan  I32 F_XET
    /-show/showfile/cue/329/midipara1  I32 F_XET
    /-show/showfile/cue/329/midipara2  I32 F_XET
/-show/showfile/cue/330  <SCUE> n=0
    /-show/showfile/cue/330/numb  I32 F_XET
    /-show/showfile/cue/330/name  S32 F_XET
    /-show/showfile/cue/330/skip  I32 F_XET
    /-show/showfile/cue/330/scene  I32 F_XET
    /-show/showfile/cue/330/bit  I32 F_XET
    /-show/showfile/cue/330/miditype  I32 F_XET
    /-show/showfile/cue/330/midichan  I32 F_XET
    /-show/showfile/cue/330/midipara1  I32 F_XET
    /-show/showfile/cue/330/midipara2  I32 F_XET
/-show/showfile/cue/331  <SCUE> n=0
    /-show/showfile/cue/331/numb  I32 F_XET
    /-show/showfile/cue/331/name  S32 F_XET
    /-show/showfile/cue/331/skip  I32 F_XET
    /-show/showfile/cue/331/scene  I32 F_XET
    /-show/showfile/cue/331/bit  I32 F_XET
    /-show/showfile/cue/331/miditype  I32 F_XET
    /-show/showfile/cue/331/midichan  I32 F_XET
    /-show/showfile/cue/331/midipara1  I32 F_XET
    /-show/showfile/cue/331/midipara2  I32 F_XET
/-show/showfile/cue/332  <SCUE> n=0
    /-show/showfile/cue/332/numb  I32 F_XET
    /-show/showfile/cue/332/name  S32 F_XET
    /-show/showfile/cue/332/skip  I32 F_XET
    /-show/showfile/cue/332/scene  I32 F_XET
    /-show/showfile/cue/332/bit  I32 F_XET
    /-show/showfile/cue/332/miditype  I32 F_XET
    /-show/showfile/cue/332/midichan  I32 F_XET
    /-show/showfile/cue/332/midipara1  I32 F_XET
    /-show/showfile/cue/332/midipara2  I32 F_XET
/-show/showfile/cue/333  <SCUE> n=0
    /-show/showfile/cue/333/numb  I32 F_XET
    /-show/showfile/cue/333/name  S32 F_XET
    /-show/showfile/cue/333/skip  I32 F_XET
    /-show/showfile/cue/333/scene  I32 F_XET
    /-show/showfile/cue/333/bit  I32 F_XET
    /-show/showfile/cue/333/miditype  I32 F_XET
    /-show/showfile/cue/333/midichan  I32 F_XET
    /-show/showfile/cue/333/midipara1  I32 F_XET
    /-show/showfile/cue/333/midipara2  I32 F_XET
/-show/showfile/cue/334  <SCUE> n=0
    /-show/showfile/cue/334/numb  I32 F_XET
    /-show/showfile/cue/334/name  S32 F_XET
    /-show/showfile/cue/334/skip  I32 F_XET
    /-show/showfile/cue/334/scene  I32 F_XET
    /-show/showfile/cue/334/bit  I32 F_XET
    /-show/showfile/cue/334/miditype  I32 F_XET
    /-show/showfile/cue/334/midichan  I32 F_XET
    /-show/showfile/cue/334/midipara1  I32 F_XET
    /-show/showfile/cue/334/midipara2  I32 F_XET
/-show/showfile/cue/335  <SCUE> n=0
    /-show/showfile/cue/335/numb  I32 F_XET
    /-show/showfile/cue/335/name  S32 F_XET
    /-show/showfile/cue/335/skip  I32 F_XET
    /-show/showfile/cue/335/scene  I32 F_XET
    /-show/showfile/cue/335/bit  I32 F_XET
    /-show/showfile/cue/335/miditype  I32 F_XET
    /-show/showfile/cue/335/midichan  I32 F_XET
    /-show/showfile/cue/335/midipara1  I32 F_XET
    /-show/showfile/cue/335/midipara2  I32 F_XET
/-show/showfile/cue/336  <SCUE> n=0
    /-show/showfile/cue/336/numb  I32 F_XET
    /-show/showfile/cue/336/name  S32 F_XET
    /-show/showfile/cue/336/skip  I32 F_XET
    /-show/showfile/cue/336/scene  I32 F_XET
    /-show/showfile/cue/336/bit  I32 F_XET
    /-show/showfile/cue/336/miditype  I32 F_XET
    /-show/showfile/cue/336/midichan  I32 F_XET
    /-show/showfile/cue/336/midipara1  I32 F_XET
    /-show/showfile/cue/336/midipara2  I32 F_XET
/-show/showfile/cue/337  <SCUE> n=0
    /-show/showfile/cue/337/numb  I32 F_XET
    /-show/showfile/cue/337/name  S32 F_XET
    /-show/showfile/cue/337/skip  I32 F_XET
    /-show/showfile/cue/337/scene  I32 F_XET
    /-show/showfile/cue/337/bit  I32 F_XET
    /-show/showfile/cue/337/miditype  I32 F_XET
    /-show/showfile/cue/337/midichan  I32 F_XET
    /-show/showfile/cue/337/midipara1  I32 F_XET
    /-show/showfile/cue/337/midipara2  I32 F_XET
/-show/showfile/cue/338  <SCUE> n=0
    /-show/showfile/cue/338/numb  I32 F_XET
    /-show/showfile/cue/338/name  S32 F_XET
    /-show/showfile/cue/338/skip  I32 F_XET
    /-show/showfile/cue/338/scene  I32 F_XET
    /-show/showfile/cue/338/bit  I32 F_XET
    /-show/showfile/cue/338/miditype  I32 F_XET
    /-show/showfile/cue/338/midichan  I32 F_XET
    /-show/showfile/cue/338/midipara1  I32 F_XET
    /-show/showfile/cue/338/midipara2  I32 F_XET
/-show/showfile/cue/339  <SCUE> n=0
    /-show/showfile/cue/339/numb  I32 F_XET
    /-show/showfile/cue/339/name  S32 F_XET
    /-show/showfile/cue/339/skip  I32 F_XET
    /-show/showfile/cue/339/scene  I32 F_XET
    /-show/showfile/cue/339/bit  I32 F_XET
    /-show/showfile/cue/339/miditype  I32 F_XET
    /-show/showfile/cue/339/midichan  I32 F_XET
    /-show/showfile/cue/339/midipara1  I32 F_XET
    /-show/showfile/cue/339/midipara2  I32 F_XET
/-show/showfile/cue/340  <SCUE> n=0
    /-show/showfile/cue/340/numb  I32 F_XET
    /-show/showfile/cue/340/name  S32 F_XET
    /-show/showfile/cue/340/skip  I32 F_XET
    /-show/showfile/cue/340/scene  I32 F_XET
    /-show/showfile/cue/340/bit  I32 F_XET
    /-show/showfile/cue/340/miditype  I32 F_XET
    /-show/showfile/cue/340/midichan  I32 F_XET
    /-show/showfile/cue/340/midipara1  I32 F_XET
    /-show/showfile/cue/340/midipara2  I32 F_XET
/-show/showfile/cue/341  <SCUE> n=0
    /-show/showfile/cue/341/numb  I32 F_XET
    /-show/showfile/cue/341/name  S32 F_XET
    /-show/showfile/cue/341/skip  I32 F_XET
    /-show/showfile/cue/341/scene  I32 F_XET
    /-show/showfile/cue/341/bit  I32 F_XET
    /-show/showfile/cue/341/miditype  I32 F_XET
    /-show/showfile/cue/341/midichan  I32 F_XET
    /-show/showfile/cue/341/midipara1  I32 F_XET
    /-show/showfile/cue/341/midipara2  I32 F_XET
/-show/showfile/cue/342  <SCUE> n=0
    /-show/showfile/cue/342/numb  I32 F_XET
    /-show/showfile/cue/342/name  S32 F_XET
    /-show/showfile/cue/342/skip  I32 F_XET
    /-show/showfile/cue/342/scene  I32 F_XET
    /-show/showfile/cue/342/bit  I32 F_XET
    /-show/showfile/cue/342/miditype  I32 F_XET
    /-show/showfile/cue/342/midichan  I32 F_XET
    /-show/showfile/cue/342/midipara1  I32 F_XET
    /-show/showfile/cue/342/midipara2  I32 F_XET
/-show/showfile/cue/343  <SCUE> n=0
    /-show/showfile/cue/343/numb  I32 F_XET
    /-show/showfile/cue/343/name  S32 F_XET
    /-show/showfile/cue/343/skip  I32 F_XET
    /-show/showfile/cue/343/scene  I32 F_XET
    /-show/showfile/cue/343/bit  I32 F_XET
    /-show/showfile/cue/343/miditype  I32 F_XET
    /-show/showfile/cue/343/midichan  I32 F_XET
    /-show/showfile/cue/343/midipara1  I32 F_XET
    /-show/showfile/cue/343/midipara2  I32 F_XET
/-show/showfile/cue/344  <SCUE> n=0
    /-show/showfile/cue/344/numb  I32 F_XET
    /-show/showfile/cue/344/name  S32 F_XET
    /-show/showfile/cue/344/skip  I32 F_XET
    /-show/showfile/cue/344/scene  I32 F_XET
    /-show/showfile/cue/344/bit  I32 F_XET
    /-show/showfile/cue/344/miditype  I32 F_XET
    /-show/showfile/cue/344/midichan  I32 F_XET
    /-show/showfile/cue/344/midipara1  I32 F_XET
    /-show/showfile/cue/344/midipara2  I32 F_XET
/-show/showfile/cue/345  <SCUE> n=0
    /-show/showfile/cue/345/numb  I32 F_XET
    /-show/showfile/cue/345/name  S32 F_XET
    /-show/showfile/cue/345/skip  I32 F_XET
    /-show/showfile/cue/345/scene  I32 F_XET
    /-show/showfile/cue/345/bit  I32 F_XET
    /-show/showfile/cue/345/miditype  I32 F_XET
    /-show/showfile/cue/345/midichan  I32 F_XET
    /-show/showfile/cue/345/midipara1  I32 F_XET
    /-show/showfile/cue/345/midipara2  I32 F_XET
/-show/showfile/cue/346  <SCUE> n=0
    /-show/showfile/cue/346/numb  I32 F_XET
    /-show/showfile/cue/346/name  S32 F_XET
    /-show/showfile/cue/346/skip  I32 F_XET
    /-show/showfile/cue/346/scene  I32 F_XET
    /-show/showfile/cue/346/bit  I32 F_XET
    /-show/showfile/cue/346/miditype  I32 F_XET
    /-show/showfile/cue/346/midichan  I32 F_XET
    /-show/showfile/cue/346/midipara1  I32 F_XET
    /-show/showfile/cue/346/midipara2  I32 F_XET
/-show/showfile/cue/347  <SCUE> n=0
    /-show/showfile/cue/347/numb  I32 F_XET
    /-show/showfile/cue/347/name  S32 F_XET
    /-show/showfile/cue/347/skip  I32 F_XET
    /-show/showfile/cue/347/scene  I32 F_XET
    /-show/showfile/cue/347/bit  I32 F_XET
    /-show/showfile/cue/347/miditype  I32 F_XET
    /-show/showfile/cue/347/midichan  I32 F_XET
    /-show/showfile/cue/347/midipara1  I32 F_XET
    /-show/showfile/cue/347/midipara2  I32 F_XET
/-show/showfile/cue/348  <SCUE> n=0
    /-show/showfile/cue/348/numb  I32 F_XET
    /-show/showfile/cue/348/name  S32 F_XET
    /-show/showfile/cue/348/skip  I32 F_XET
    /-show/showfile/cue/348/scene  I32 F_XET
    /-show/showfile/cue/348/bit  I32 F_XET
    /-show/showfile/cue/348/miditype  I32 F_XET
    /-show/showfile/cue/348/midichan  I32 F_XET
    /-show/showfile/cue/348/midipara1  I32 F_XET
    /-show/showfile/cue/348/midipara2  I32 F_XET
/-show/showfile/cue/349  <SCUE> n=0
    /-show/showfile/cue/349/numb  I32 F_XET
    /-show/showfile/cue/349/name  S32 F_XET
    /-show/showfile/cue/349/skip  I32 F_XET
    /-show/showfile/cue/349/scene  I32 F_XET
    /-show/showfile/cue/349/bit  I32 F_XET
    /-show/showfile/cue/349/miditype  I32 F_XET
    /-show/showfile/cue/349/midichan  I32 F_XET
    /-show/showfile/cue/349/midipara1  I32 F_XET
    /-show/showfile/cue/349/midipara2  I32 F_XET
/-show/showfile/cue/350  <SCUE> n=0
    /-show/showfile/cue/350/numb  I32 F_XET
    /-show/showfile/cue/350/name  S32 F_XET
    /-show/showfile/cue/350/skip  I32 F_XET
    /-show/showfile/cue/350/scene  I32 F_XET
    /-show/showfile/cue/350/bit  I32 F_XET
    /-show/showfile/cue/350/miditype  I32 F_XET
    /-show/showfile/cue/350/midichan  I32 F_XET
    /-show/showfile/cue/350/midipara1  I32 F_XET
    /-show/showfile/cue/350/midipara2  I32 F_XET
/-show/showfile/cue/351  <SCUE> n=0
    /-show/showfile/cue/351/numb  I32 F_XET
    /-show/showfile/cue/351/name  S32 F_XET
    /-show/showfile/cue/351/skip  I32 F_XET
    /-show/showfile/cue/351/scene  I32 F_XET
    /-show/showfile/cue/351/bit  I32 F_XET
    /-show/showfile/cue/351/miditype  I32 F_XET
    /-show/showfile/cue/351/midichan  I32 F_XET
    /-show/showfile/cue/351/midipara1  I32 F_XET
    /-show/showfile/cue/351/midipara2  I32 F_XET
/-show/showfile/cue/352  <SCUE> n=0
    /-show/showfile/cue/352/numb  I32 F_XET
    /-show/showfile/cue/352/name  S32 F_XET
    /-show/showfile/cue/352/skip  I32 F_XET
    /-show/showfile/cue/352/scene  I32 F_XET
    /-show/showfile/cue/352/bit  I32 F_XET
    /-show/showfile/cue/352/miditype  I32 F_XET
    /-show/showfile/cue/352/midichan  I32 F_XET
    /-show/showfile/cue/352/midipara1  I32 F_XET
    /-show/showfile/cue/352/midipara2  I32 F_XET
/-show/showfile/cue/353  <SCUE> n=0
    /-show/showfile/cue/353/numb  I32 F_XET
    /-show/showfile/cue/353/name  S32 F_XET
    /-show/showfile/cue/353/skip  I32 F_XET
    /-show/showfile/cue/353/scene  I32 F_XET
    /-show/showfile/cue/353/bit  I32 F_XET
    /-show/showfile/cue/353/miditype  I32 F_XET
    /-show/showfile/cue/353/midichan  I32 F_XET
    /-show/showfile/cue/353/midipara1  I32 F_XET
    /-show/showfile/cue/353/midipara2  I32 F_XET
/-show/showfile/cue/354  <SCUE> n=0
    /-show/showfile/cue/354/numb  I32 F_XET
    /-show/showfile/cue/354/name  S32 F_XET
    /-show/showfile/cue/354/skip  I32 F_XET
    /-show/showfile/cue/354/scene  I32 F_XET
    /-show/showfile/cue/354/bit  I32 F_XET
    /-show/showfile/cue/354/miditype  I32 F_XET
    /-show/showfile/cue/354/midichan  I32 F_XET
    /-show/showfile/cue/354/midipara1  I32 F_XET
    /-show/showfile/cue/354/midipara2  I32 F_XET
/-show/showfile/cue/355  <SCUE> n=0
    /-show/showfile/cue/355/numb  I32 F_XET
    /-show/showfile/cue/355/name  S32 F_XET
    /-show/showfile/cue/355/skip  I32 F_XET
    /-show/showfile/cue/355/scene  I32 F_XET
    /-show/showfile/cue/355/bit  I32 F_XET
    /-show/showfile/cue/355/miditype  I32 F_XET
    /-show/showfile/cue/355/midichan  I32 F_XET
    /-show/showfile/cue/355/midipara1  I32 F_XET
    /-show/showfile/cue/355/midipara2  I32 F_XET
/-show/showfile/cue/356  <SCUE> n=0
    /-show/showfile/cue/356/numb  I32 F_XET
    /-show/showfile/cue/356/name  S32 F_XET
    /-show/showfile/cue/356/skip  I32 F_XET
    /-show/showfile/cue/356/scene  I32 F_XET
    /-show/showfile/cue/356/bit  I32 F_XET
    /-show/showfile/cue/356/miditype  I32 F_XET
    /-show/showfile/cue/356/midichan  I32 F_XET
    /-show/showfile/cue/356/midipara1  I32 F_XET
    /-show/showfile/cue/356/midipara2  I32 F_XET
/-show/showfile/cue/357  <SCUE> n=0
    /-show/showfile/cue/357/numb  I32 F_XET
    /-show/showfile/cue/357/name  S32 F_XET
    /-show/showfile/cue/357/skip  I32 F_XET
    /-show/showfile/cue/357/scene  I32 F_XET
    /-show/showfile/cue/357/bit  I32 F_XET
    /-show/showfile/cue/357/miditype  I32 F_XET
    /-show/showfile/cue/357/midichan  I32 F_XET
    /-show/showfile/cue/357/midipara1  I32 F_XET
    /-show/showfile/cue/357/midipara2  I32 F_XET
/-show/showfile/cue/358  <SCUE> n=0
    /-show/showfile/cue/358/numb  I32 F_XET
    /-show/showfile/cue/358/name  S32 F_XET
    /-show/showfile/cue/358/skip  I32 F_XET
    /-show/showfile/cue/358/scene  I32 F_XET
    /-show/showfile/cue/358/bit  I32 F_XET
    /-show/showfile/cue/358/miditype  I32 F_XET
    /-show/showfile/cue/358/midichan  I32 F_XET
    /-show/showfile/cue/358/midipara1  I32 F_XET
    /-show/showfile/cue/358/midipara2  I32 F_XET
/-show/showfile/cue/359  <SCUE> n=0
    /-show/showfile/cue/359/numb  I32 F_XET
    /-show/showfile/cue/359/name  S32 F_XET
    /-show/showfile/cue/359/skip  I32 F_XET
    /-show/showfile/cue/359/scene  I32 F_XET
    /-show/showfile/cue/359/bit  I32 F_XET
    /-show/showfile/cue/359/miditype  I32 F_XET
    /-show/showfile/cue/359/midichan  I32 F_XET
    /-show/showfile/cue/359/midipara1  I32 F_XET
    /-show/showfile/cue/359/midipara2  I32 F_XET
/-show/showfile/cue/360  <SCUE> n=0
    /-show/showfile/cue/360/numb  I32 F_XET
    /-show/showfile/cue/360/name  S32 F_XET
    /-show/showfile/cue/360/skip  I32 F_XET
    /-show/showfile/cue/360/scene  I32 F_XET
    /-show/showfile/cue/360/bit  I32 F_XET
    /-show/showfile/cue/360/miditype  I32 F_XET
    /-show/showfile/cue/360/midichan  I32 F_XET
    /-show/showfile/cue/360/midipara1  I32 F_XET
    /-show/showfile/cue/360/midipara2  I32 F_XET
/-show/showfile/cue/361  <SCUE> n=0
    /-show/showfile/cue/361/numb  I32 F_XET
    /-show/showfile/cue/361/name  S32 F_XET
    /-show/showfile/cue/361/skip  I32 F_XET
    /-show/showfile/cue/361/scene  I32 F_XET
    /-show/showfile/cue/361/bit  I32 F_XET
    /-show/showfile/cue/361/miditype  I32 F_XET
    /-show/showfile/cue/361/midichan  I32 F_XET
    /-show/showfile/cue/361/midipara1  I32 F_XET
    /-show/showfile/cue/361/midipara2  I32 F_XET
/-show/showfile/cue/362  <SCUE> n=0
    /-show/showfile/cue/362/numb  I32 F_XET
    /-show/showfile/cue/362/name  S32 F_XET
    /-show/showfile/cue/362/skip  I32 F_XET
    /-show/showfile/cue/362/scene  I32 F_XET
    /-show/showfile/cue/362/bit  I32 F_XET
    /-show/showfile/cue/362/miditype  I32 F_XET
    /-show/showfile/cue/362/midichan  I32 F_XET
    /-show/showfile/cue/362/midipara1  I32 F_XET
    /-show/showfile/cue/362/midipara2  I32 F_XET
/-show/showfile/cue/363  <SCUE> n=0
    /-show/showfile/cue/363/numb  I32 F_XET
    /-show/showfile/cue/363/name  S32 F_XET
    /-show/showfile/cue/363/skip  I32 F_XET
    /-show/showfile/cue/363/scene  I32 F_XET
    /-show/showfile/cue/363/bit  I32 F_XET
    /-show/showfile/cue/363/miditype  I32 F_XET
    /-show/showfile/cue/363/midichan  I32 F_XET
    /-show/showfile/cue/363/midipara1  I32 F_XET
    /-show/showfile/cue/363/midipara2  I32 F_XET
/-show/showfile/cue/364  <SCUE> n=0
    /-show/showfile/cue/364/numb  I32 F_XET
    /-show/showfile/cue/364/name  S32 F_XET
    /-show/showfile/cue/364/skip  I32 F_XET
    /-show/showfile/cue/364/scene  I32 F_XET
    /-show/showfile/cue/364/bit  I32 F_XET
    /-show/showfile/cue/364/miditype  I32 F_XET
    /-show/showfile/cue/364/midichan  I32 F_XET
    /-show/showfile/cue/364/midipara1  I32 F_XET
    /-show/showfile/cue/364/midipara2  I32 F_XET
/-show/showfile/cue/365  <SCUE> n=0
    /-show/showfile/cue/365/numb  I32 F_XET
    /-show/showfile/cue/365/name  S32 F_XET
    /-show/showfile/cue/365/skip  I32 F_XET
    /-show/showfile/cue/365/scene  I32 F_XET
    /-show/showfile/cue/365/bit  I32 F_XET
    /-show/showfile/cue/365/miditype  I32 F_XET
    /-show/showfile/cue/365/midichan  I32 F_XET
    /-show/showfile/cue/365/midipara1  I32 F_XET
    /-show/showfile/cue/365/midipara2  I32 F_XET
/-show/showfile/cue/366  <SCUE> n=0
    /-show/showfile/cue/366/numb  I32 F_XET
    /-show/showfile/cue/366/name  S32 F_XET
    /-show/showfile/cue/366/skip  I32 F_XET
    /-show/showfile/cue/366/scene  I32 F_XET
    /-show/showfile/cue/366/bit  I32 F_XET
    /-show/showfile/cue/366/miditype  I32 F_XET
    /-show/showfile/cue/366/midichan  I32 F_XET
    /-show/showfile/cue/366/midipara1  I32 F_XET
    /-show/showfile/cue/366/midipara2  I32 F_XET
/-show/showfile/cue/367  <SCUE> n=0
    /-show/showfile/cue/367/numb  I32 F_XET
    /-show/showfile/cue/367/name  S32 F_XET
    /-show/showfile/cue/367/skip  I32 F_XET
    /-show/showfile/cue/367/scene  I32 F_XET
    /-show/showfile/cue/367/bit  I32 F_XET
    /-show/showfile/cue/367/miditype  I32 F_XET
    /-show/showfile/cue/367/midichan  I32 F_XET
    /-show/showfile/cue/367/midipara1  I32 F_XET
    /-show/showfile/cue/367/midipara2  I32 F_XET
/-show/showfile/cue/368  <SCUE> n=0
    /-show/showfile/cue/368/numb  I32 F_XET
    /-show/showfile/cue/368/name  S32 F_XET
    /-show/showfile/cue/368/skip  I32 F_XET
    /-show/showfile/cue/368/scene  I32 F_XET
    /-show/showfile/cue/368/bit  I32 F_XET
    /-show/showfile/cue/368/miditype  I32 F_XET
    /-show/showfile/cue/368/midichan  I32 F_XET
    /-show/showfile/cue/368/midipara1  I32 F_XET
    /-show/showfile/cue/368/midipara2  I32 F_XET
/-show/showfile/cue/369  <SCUE> n=0
    /-show/showfile/cue/369/numb  I32 F_XET
    /-show/showfile/cue/369/name  S32 F_XET
    /-show/showfile/cue/369/skip  I32 F_XET
    /-show/showfile/cue/369/scene  I32 F_XET
    /-show/showfile/cue/369/bit  I32 F_XET
    /-show/showfile/cue/369/miditype  I32 F_XET
    /-show/showfile/cue/369/midichan  I32 F_XET
    /-show/showfile/cue/369/midipara1  I32 F_XET
    /-show/showfile/cue/369/midipara2  I32 F_XET
/-show/showfile/cue/370  <SCUE> n=0
    /-show/showfile/cue/370/numb  I32 F_XET
    /-show/showfile/cue/370/name  S32 F_XET
    /-show/showfile/cue/370/skip  I32 F_XET
    /-show/showfile/cue/370/scene  I32 F_XET
    /-show/showfile/cue/370/bit  I32 F_XET
    /-show/showfile/cue/370/miditype  I32 F_XET
    /-show/showfile/cue/370/midichan  I32 F_XET
    /-show/showfile/cue/370/midipara1  I32 F_XET
    /-show/showfile/cue/370/midipara2  I32 F_XET
/-show/showfile/cue/371  <SCUE> n=0
    /-show/showfile/cue/371/numb  I32 F_XET
    /-show/showfile/cue/371/name  S32 F_XET
    /-show/showfile/cue/371/skip  I32 F_XET
    /-show/showfile/cue/371/scene  I32 F_XET
    /-show/showfile/cue/371/bit  I32 F_XET
    /-show/showfile/cue/371/miditype  I32 F_XET
    /-show/showfile/cue/371/midichan  I32 F_XET
    /-show/showfile/cue/371/midipara1  I32 F_XET
    /-show/showfile/cue/371/midipara2  I32 F_XET
/-show/showfile/cue/372  <SCUE> n=0
    /-show/showfile/cue/372/numb  I32 F_XET
    /-show/showfile/cue/372/name  S32 F_XET
    /-show/showfile/cue/372/skip  I32 F_XET
    /-show/showfile/cue/372/scene  I32 F_XET
    /-show/showfile/cue/372/bit  I32 F_XET
    /-show/showfile/cue/372/miditype  I32 F_XET
    /-show/showfile/cue/372/midichan  I32 F_XET
    /-show/showfile/cue/372/midipara1  I32 F_XET
    /-show/showfile/cue/372/midipara2  I32 F_XET
/-show/showfile/cue/373  <SCUE> n=0
    /-show/showfile/cue/373/numb  I32 F_XET
    /-show/showfile/cue/373/name  S32 F_XET
    /-show/showfile/cue/373/skip  I32 F_XET
    /-show/showfile/cue/373/scene  I32 F_XET
    /-show/showfile/cue/373/bit  I32 F_XET
    /-show/showfile/cue/373/miditype  I32 F_XET
    /-show/showfile/cue/373/midichan  I32 F_XET
    /-show/showfile/cue/373/midipara1  I32 F_XET
    /-show/showfile/cue/373/midipara2  I32 F_XET
/-show/showfile/cue/374  <SCUE> n=0
    /-show/showfile/cue/374/numb  I32 F_XET
    /-show/showfile/cue/374/name  S32 F_XET
    /-show/showfile/cue/374/skip  I32 F_XET
    /-show/showfile/cue/374/scene  I32 F_XET
    /-show/showfile/cue/374/bit  I32 F_XET
    /-show/showfile/cue/374/miditype  I32 F_XET
    /-show/showfile/cue/374/midichan  I32 F_XET
    /-show/showfile/cue/374/midipara1  I32 F_XET
    /-show/showfile/cue/374/midipara2  I32 F_XET
/-show/showfile/cue/375  <SCUE> n=0
    /-show/showfile/cue/375/numb  I32 F_XET
    /-show/showfile/cue/375/name  S32 F_XET
    /-show/showfile/cue/375/skip  I32 F_XET
    /-show/showfile/cue/375/scene  I32 F_XET
    /-show/showfile/cue/375/bit  I32 F_XET
    /-show/showfile/cue/375/miditype  I32 F_XET
    /-show/showfile/cue/375/midichan  I32 F_XET
    /-show/showfile/cue/375/midipara1  I32 F_XET
    /-show/showfile/cue/375/midipara2  I32 F_XET
/-show/showfile/cue/376  <SCUE> n=0
    /-show/showfile/cue/376/numb  I32 F_XET
    /-show/showfile/cue/376/name  S32 F_XET
    /-show/showfile/cue/376/skip  I32 F_XET
    /-show/showfile/cue/376/scene  I32 F_XET
    /-show/showfile/cue/376/bit  I32 F_XET
    /-show/showfile/cue/376/miditype  I32 F_XET
    /-show/showfile/cue/376/midichan  I32 F_XET
    /-show/showfile/cue/376/midipara1  I32 F_XET
    /-show/showfile/cue/376/midipara2  I32 F_XET
/-show/showfile/cue/377  <SCUE> n=0
    /-show/showfile/cue/377/numb  I32 F_XET
    /-show/showfile/cue/377/name  S32 F_XET
    /-show/showfile/cue/377/skip  I32 F_XET
    /-show/showfile/cue/377/scene  I32 F_XET
    /-show/showfile/cue/377/bit  I32 F_XET
    /-show/showfile/cue/377/miditype  I32 F_XET
    /-show/showfile/cue/377/midichan  I32 F_XET
    /-show/showfile/cue/377/midipara1  I32 F_XET
    /-show/showfile/cue/377/midipara2  I32 F_XET
/-show/showfile/cue/378  <SCUE> n=0
    /-show/showfile/cue/378/numb  I32 F_XET
    /-show/showfile/cue/378/name  S32 F_XET
    /-show/showfile/cue/378/skip  I32 F_XET
    /-show/showfile/cue/378/scene  I32 F_XET
    /-show/showfile/cue/378/bit  I32 F_XET
    /-show/showfile/cue/378/miditype  I32 F_XET
    /-show/showfile/cue/378/midichan  I32 F_XET
    /-show/showfile/cue/378/midipara1  I32 F_XET
    /-show/showfile/cue/378/midipara2  I32 F_XET
/-show/showfile/cue/379  <SCUE> n=0
    /-show/showfile/cue/379/numb  I32 F_XET
    /-show/showfile/cue/379/name  S32 F_XET
    /-show/showfile/cue/379/skip  I32 F_XET
    /-show/showfile/cue/379/scene  I32 F_XET
    /-show/showfile/cue/379/bit  I32 F_XET
    /-show/showfile/cue/379/miditype  I32 F_XET
    /-show/showfile/cue/379/midichan  I32 F_XET
    /-show/showfile/cue/379/midipara1  I32 F_XET
    /-show/showfile/cue/379/midipara2  I32 F_XET
/-show/showfile/cue/380  <SCUE> n=0
    /-show/showfile/cue/380/numb  I32 F_XET
    /-show/showfile/cue/380/name  S32 F_XET
    /-show/showfile/cue/380/skip  I32 F_XET
    /-show/showfile/cue/380/scene  I32 F_XET
    /-show/showfile/cue/380/bit  I32 F_XET
    /-show/showfile/cue/380/miditype  I32 F_XET
    /-show/showfile/cue/380/midichan  I32 F_XET
    /-show/showfile/cue/380/midipara1  I32 F_XET
    /-show/showfile/cue/380/midipara2  I32 F_XET
/-show/showfile/cue/381  <SCUE> n=0
    /-show/showfile/cue/381/numb  I32 F_XET
    /-show/showfile/cue/381/name  S32 F_XET
    /-show/showfile/cue/381/skip  I32 F_XET
    /-show/showfile/cue/381/scene  I32 F_XET
    /-show/showfile/cue/381/bit  I32 F_XET
    /-show/showfile/cue/381/miditype  I32 F_XET
    /-show/showfile/cue/381/midichan  I32 F_XET
    /-show/showfile/cue/381/midipara1  I32 F_XET
    /-show/showfile/cue/381/midipara2  I32 F_XET
/-show/showfile/cue/382  <SCUE> n=0
    /-show/showfile/cue/382/numb  I32 F_XET
    /-show/showfile/cue/382/name  S32 F_XET
    /-show/showfile/cue/382/skip  I32 F_XET
    /-show/showfile/cue/382/scene  I32 F_XET
    /-show/showfile/cue/382/bit  I32 F_XET
    /-show/showfile/cue/382/miditype  I32 F_XET
    /-show/showfile/cue/382/midichan  I32 F_XET
    /-show/showfile/cue/382/midipara1  I32 F_XET
    /-show/showfile/cue/382/midipara2  I32 F_XET
/-show/showfile/cue/383  <SCUE> n=0
    /-show/showfile/cue/383/numb  I32 F_XET
    /-show/showfile/cue/383/name  S32 F_XET
    /-show/showfile/cue/383/skip  I32 F_XET
    /-show/showfile/cue/383/scene  I32 F_XET
    /-show/showfile/cue/383/bit  I32 F_XET
    /-show/showfile/cue/383/miditype  I32 F_XET
    /-show/showfile/cue/383/midichan  I32 F_XET
    /-show/showfile/cue/383/midipara1  I32 F_XET
    /-show/showfile/cue/383/midipara2  I32 F_XET
/-show/showfile/cue/384  <SCUE> n=0
    /-show/showfile/cue/384/numb  I32 F_XET
    /-show/showfile/cue/384/name  S32 F_XET
    /-show/showfile/cue/384/skip  I32 F_XET
    /-show/showfile/cue/384/scene  I32 F_XET
    /-show/showfile/cue/384/bit  I32 F_XET
    /-show/showfile/cue/384/miditype  I32 F_XET
    /-show/showfile/cue/384/midichan  I32 F_XET
    /-show/showfile/cue/384/midipara1  I32 F_XET
    /-show/showfile/cue/384/midipara2  I32 F_XET
/-show/showfile/cue/385  <SCUE> n=0
    /-show/showfile/cue/385/numb  I32 F_XET
    /-show/showfile/cue/385/name  S32 F_XET
    /-show/showfile/cue/385/skip  I32 F_XET
    /-show/showfile/cue/385/scene  I32 F_XET
    /-show/showfile/cue/385/bit  I32 F_XET
    /-show/showfile/cue/385/miditype  I32 F_XET
    /-show/showfile/cue/385/midichan  I32 F_XET
    /-show/showfile/cue/385/midipara1  I32 F_XET
    /-show/showfile/cue/385/midipara2  I32 F_XET
/-show/showfile/cue/386  <SCUE> n=0
    /-show/showfile/cue/386/numb  I32 F_XET
    /-show/showfile/cue/386/name  S32 F_XET
    /-show/showfile/cue/386/skip  I32 F_XET
    /-show/showfile/cue/386/scene  I32 F_XET
    /-show/showfile/cue/386/bit  I32 F_XET
    /-show/showfile/cue/386/miditype  I32 F_XET
    /-show/showfile/cue/386/midichan  I32 F_XET
    /-show/showfile/cue/386/midipara1  I32 F_XET
    /-show/showfile/cue/386/midipara2  I32 F_XET
/-show/showfile/cue/387  <SCUE> n=0
    /-show/showfile/cue/387/numb  I32 F_XET
    /-show/showfile/cue/387/name  S32 F_XET
    /-show/showfile/cue/387/skip  I32 F_XET
    /-show/showfile/cue/387/scene  I32 F_XET
    /-show/showfile/cue/387/bit  I32 F_XET
    /-show/showfile/cue/387/miditype  I32 F_XET
    /-show/showfile/cue/387/midichan  I32 F_XET
    /-show/showfile/cue/387/midipara1  I32 F_XET
    /-show/showfile/cue/387/midipara2  I32 F_XET
/-show/showfile/cue/388  <SCUE> n=0
    /-show/showfile/cue/388/numb  I32 F_XET
    /-show/showfile/cue/388/name  S32 F_XET
    /-show/showfile/cue/388/skip  I32 F_XET
    /-show/showfile/cue/388/scene  I32 F_XET
    /-show/showfile/cue/388/bit  I32 F_XET
    /-show/showfile/cue/388/miditype  I32 F_XET
    /-show/showfile/cue/388/midichan  I32 F_XET
    /-show/showfile/cue/388/midipara1  I32 F_XET
    /-show/showfile/cue/388/midipara2  I32 F_XET
/-show/showfile/cue/389  <SCUE> n=0
    /-show/showfile/cue/389/numb  I32 F_XET
    /-show/showfile/cue/389/name  S32 F_XET
    /-show/showfile/cue/389/skip  I32 F_XET
    /-show/showfile/cue/389/scene  I32 F_XET
    /-show/showfile/cue/389/bit  I32 F_XET
    /-show/showfile/cue/389/miditype  I32 F_XET
    /-show/showfile/cue/389/midichan  I32 F_XET
    /-show/showfile/cue/389/midipara1  I32 F_XET
    /-show/showfile/cue/389/midipara2  I32 F_XET
/-show/showfile/cue/390  <SCUE> n=0
    /-show/showfile/cue/390/numb  I32 F_XET
    /-show/showfile/cue/390/name  S32 F_XET
    /-show/showfile/cue/390/skip  I32 F_XET
    /-show/showfile/cue/390/scene  I32 F_XET
    /-show/showfile/cue/390/bit  I32 F_XET
    /-show/showfile/cue/390/miditype  I32 F_XET
    /-show/showfile/cue/390/midichan  I32 F_XET
    /-show/showfile/cue/390/midipara1  I32 F_XET
    /-show/showfile/cue/390/midipara2  I32 F_XET
/-show/showfile/cue/391  <SCUE> n=0
    /-show/showfile/cue/391/numb  I32 F_XET
    /-show/showfile/cue/391/name  S32 F_XET
    /-show/showfile/cue/391/skip  I32 F_XET
    /-show/showfile/cue/391/scene  I32 F_XET
    /-show/showfile/cue/391/bit  I32 F_XET
    /-show/showfile/cue/391/miditype  I32 F_XET
    /-show/showfile/cue/391/midichan  I32 F_XET
    /-show/showfile/cue/391/midipara1  I32 F_XET
    /-show/showfile/cue/391/midipara2  I32 F_XET
/-show/showfile/cue/392  <SCUE> n=0
    /-show/showfile/cue/392/numb  I32 F_XET
    /-show/showfile/cue/392/name  S32 F_XET
    /-show/showfile/cue/392/skip  I32 F_XET
    /-show/showfile/cue/392/scene  I32 F_XET
    /-show/showfile/cue/392/bit  I32 F_XET
    /-show/showfile/cue/392/miditype  I32 F_XET
    /-show/showfile/cue/392/midichan  I32 F_XET
    /-show/showfile/cue/392/midipara1  I32 F_XET
    /-show/showfile/cue/392/midipara2  I32 F_XET
/-show/showfile/cue/393  <SCUE> n=0
    /-show/showfile/cue/393/numb  I32 F_XET
    /-show/showfile/cue/393/name  S32 F_XET
    /-show/showfile/cue/393/skip  I32 F_XET
    /-show/showfile/cue/393/scene  I32 F_XET
    /-show/showfile/cue/393/bit  I32 F_XET
    /-show/showfile/cue/393/miditype  I32 F_XET
    /-show/showfile/cue/393/midichan  I32 F_XET
    /-show/showfile/cue/393/midipara1  I32 F_XET
    /-show/showfile/cue/393/midipara2  I32 F_XET
/-show/showfile/cue/394  <SCUE> n=0
    /-show/showfile/cue/394/numb  I32 F_XET
    /-show/showfile/cue/394/name  S32 F_XET
    /-show/showfile/cue/394/skip  I32 F_XET
    /-show/showfile/cue/394/scene  I32 F_XET
    /-show/showfile/cue/394/bit  I32 F_XET
    /-show/showfile/cue/394/miditype  I32 F_XET
    /-show/showfile/cue/394/midichan  I32 F_XET
    /-show/showfile/cue/394/midipara1  I32 F_XET
    /-show/showfile/cue/394/midipara2  I32 F_XET
/-show/showfile/cue/395  <SCUE> n=0
    /-show/showfile/cue/395/numb  I32 F_XET
    /-show/showfile/cue/395/name  S32 F_XET
    /-show/showfile/cue/395/skip  I32 F_XET
    /-show/showfile/cue/395/scene  I32 F_XET
    /-show/showfile/cue/395/bit  I32 F_XET
    /-show/showfile/cue/395/miditype  I32 F_XET
    /-show/showfile/cue/395/midichan  I32 F_XET
    /-show/showfile/cue/395/midipara1  I32 F_XET
    /-show/showfile/cue/395/midipara2  I32 F_XET
/-show/showfile/cue/396  <SCUE> n=0
    /-show/showfile/cue/396/numb  I32 F_XET
    /-show/showfile/cue/396/name  S32 F_XET
    /-show/showfile/cue/396/skip  I32 F_XET
    /-show/showfile/cue/396/scene  I32 F_XET
    /-show/showfile/cue/396/bit  I32 F_XET
    /-show/showfile/cue/396/miditype  I32 F_XET
    /-show/showfile/cue/396/midichan  I32 F_XET
    /-show/showfile/cue/396/midipara1  I32 F_XET
    /-show/showfile/cue/396/midipara2  I32 F_XET
/-show/showfile/cue/397  <SCUE> n=0
    /-show/showfile/cue/397/numb  I32 F_XET
    /-show/showfile/cue/397/name  S32 F_XET
    /-show/showfile/cue/397/skip  I32 F_XET
    /-show/showfile/cue/397/scene  I32 F_XET
    /-show/showfile/cue/397/bit  I32 F_XET
    /-show/showfile/cue/397/miditype  I32 F_XET
    /-show/showfile/cue/397/midichan  I32 F_XET
    /-show/showfile/cue/397/midipara1  I32 F_XET
    /-show/showfile/cue/397/midipara2  I32 F_XET
/-show/showfile/cue/398  <SCUE> n=0
    /-show/showfile/cue/398/numb  I32 F_XET
    /-show/showfile/cue/398/name  S32 F_XET
    /-show/showfile/cue/398/skip  I32 F_XET
    /-show/showfile/cue/398/scene  I32 F_XET
    /-show/showfile/cue/398/bit  I32 F_XET
    /-show/showfile/cue/398/miditype  I32 F_XET
    /-show/showfile/cue/398/midichan  I32 F_XET
    /-show/showfile/cue/398/midipara1  I32 F_XET
    /-show/showfile/cue/398/midipara2  I32 F_XET
/-show/showfile/cue/399  <SCUE> n=0
    /-show/showfile/cue/399/numb  I32 F_XET
    /-show/showfile/cue/399/name  S32 F_XET
    /-show/showfile/cue/399/skip  I32 F_XET
    /-show/showfile/cue/399/scene  I32 F_XET
    /-show/showfile/cue/399/bit  I32 F_XET
    /-show/showfile/cue/399/miditype  I32 F_XET
    /-show/showfile/cue/399/midichan  I32 F_XET
    /-show/showfile/cue/399/midipara1  I32 F_XET
    /-show/showfile/cue/399/midipara2  I32 F_XET
/-show/showfile/cue/400  <SCUE> n=0
    /-show/showfile/cue/400/numb  I32 F_XET
    /-show/showfile/cue/400/name  S32 F_XET
    /-show/showfile/cue/400/skip  I32 F_XET
    /-show/showfile/cue/400/scene  I32 F_XET
    /-show/showfile/cue/400/bit  I32 F_XET
    /-show/showfile/cue/400/miditype  I32 F_XET
    /-show/showfile/cue/400/midichan  I32 F_XET
    /-show/showfile/cue/400/midipara1  I32 F_XET
    /-show/showfile/cue/400/midipara2  I32 F_XET
/-show/showfile/cue/401  <SCUE> n=0
    /-show/showfile/cue/401/numb  I32 F_XET
    /-show/showfile/cue/401/name  S32 F_XET
    /-show/showfile/cue/401/skip  I32 F_XET
    /-show/showfile/cue/401/scene  I32 F_XET
    /-show/showfile/cue/401/bit  I32 F_XET
    /-show/showfile/cue/401/miditype  I32 F_XET
    /-show/showfile/cue/401/midichan  I32 F_XET
    /-show/showfile/cue/401/midipara1  I32 F_XET
    /-show/showfile/cue/401/midipara2  I32 F_XET
/-show/showfile/cue/402  <SCUE> n=0
    /-show/showfile/cue/402/numb  I32 F_XET
    /-show/showfile/cue/402/name  S32 F_XET
    /-show/showfile/cue/402/skip  I32 F_XET
    /-show/showfile/cue/402/scene  I32 F_XET
    /-show/showfile/cue/402/bit  I32 F_XET
    /-show/showfile/cue/402/miditype  I32 F_XET
    /-show/showfile/cue/402/midichan  I32 F_XET
    /-show/showfile/cue/402/midipara1  I32 F_XET
    /-show/showfile/cue/402/midipara2  I32 F_XET
/-show/showfile/cue/403  <SCUE> n=0
    /-show/showfile/cue/403/numb  I32 F_XET
    /-show/showfile/cue/403/name  S32 F_XET
    /-show/showfile/cue/403/skip  I32 F_XET
    /-show/showfile/cue/403/scene  I32 F_XET
    /-show/showfile/cue/403/bit  I32 F_XET
    /-show/showfile/cue/403/miditype  I32 F_XET
    /-show/showfile/cue/403/midichan  I32 F_XET
    /-show/showfile/cue/403/midipara1  I32 F_XET
    /-show/showfile/cue/403/midipara2  I32 F_XET
/-show/showfile/cue/404  <SCUE> n=0
    /-show/showfile/cue/404/numb  I32 F_XET
    /-show/showfile/cue/404/name  S32 F_XET
    /-show/showfile/cue/404/skip  I32 F_XET
    /-show/showfile/cue/404/scene  I32 F_XET
    /-show/showfile/cue/404/bit  I32 F_XET
    /-show/showfile/cue/404/miditype  I32 F_XET
    /-show/showfile/cue/404/midichan  I32 F_XET
    /-show/showfile/cue/404/midipara1  I32 F_XET
    /-show/showfile/cue/404/midipara2  I32 F_XET
/-show/showfile/cue/405  <SCUE> n=0
    /-show/showfile/cue/405/numb  I32 F_XET
    /-show/showfile/cue/405/name  S32 F_XET
    /-show/showfile/cue/405/skip  I32 F_XET
    /-show/showfile/cue/405/scene  I32 F_XET
    /-show/showfile/cue/405/bit  I32 F_XET
    /-show/showfile/cue/405/miditype  I32 F_XET
    /-show/showfile/cue/405/midichan  I32 F_XET
    /-show/showfile/cue/405/midipara1  I32 F_XET
    /-show/showfile/cue/405/midipara2  I32 F_XET
/-show/showfile/cue/406  <SCUE> n=0
    /-show/showfile/cue/406/numb  I32 F_XET
    /-show/showfile/cue/406/name  S32 F_XET
    /-show/showfile/cue/406/skip  I32 F_XET
    /-show/showfile/cue/406/scene  I32 F_XET
    /-show/showfile/cue/406/bit  I32 F_XET
    /-show/showfile/cue/406/miditype  I32 F_XET
    /-show/showfile/cue/406/midichan  I32 F_XET
    /-show/showfile/cue/406/midipara1  I32 F_XET
    /-show/showfile/cue/406/midipara2  I32 F_XET
/-show/showfile/cue/407  <SCUE> n=0
    /-show/showfile/cue/407/numb  I32 F_XET
    /-show/showfile/cue/407/name  S32 F_XET
    /-show/showfile/cue/407/skip  I32 F_XET
    /-show/showfile/cue/407/scene  I32 F_XET
    /-show/showfile/cue/407/bit  I32 F_XET
    /-show/showfile/cue/407/miditype  I32 F_XET
    /-show/showfile/cue/407/midichan  I32 F_XET
    /-show/showfile/cue/407/midipara1  I32 F_XET
    /-show/showfile/cue/407/midipara2  I32 F_XET
/-show/showfile/cue/408  <SCUE> n=0
    /-show/showfile/cue/408/numb  I32 F_XET
    /-show/showfile/cue/408/name  S32 F_XET
    /-show/showfile/cue/408/skip  I32 F_XET
    /-show/showfile/cue/408/scene  I32 F_XET
    /-show/showfile/cue/408/bit  I32 F_XET
    /-show/showfile/cue/408/miditype  I32 F_XET
    /-show/showfile/cue/408/midichan  I32 F_XET
    /-show/showfile/cue/408/midipara1  I32 F_XET
    /-show/showfile/cue/408/midipara2  I32 F_XET
/-show/showfile/cue/409  <SCUE> n=0
    /-show/showfile/cue/409/numb  I32 F_XET
    /-show/showfile/cue/409/name  S32 F_XET
    /-show/showfile/cue/409/skip  I32 F_XET
    /-show/showfile/cue/409/scene  I32 F_XET
    /-show/showfile/cue/409/bit  I32 F_XET
    /-show/showfile/cue/409/miditype  I32 F_XET
    /-show/showfile/cue/409/midichan  I32 F_XET
    /-show/showfile/cue/409/midipara1  I32 F_XET
    /-show/showfile/cue/409/midipara2  I32 F_XET
/-show/showfile/cue/410  <SCUE> n=0
    /-show/showfile/cue/410/numb  I32 F_XET
    /-show/showfile/cue/410/name  S32 F_XET
    /-show/showfile/cue/410/skip  I32 F_XET
    /-show/showfile/cue/410/scene  I32 F_XET
    /-show/showfile/cue/410/bit  I32 F_XET
    /-show/showfile/cue/410/miditype  I32 F_XET
    /-show/showfile/cue/410/midichan  I32 F_XET
    /-show/showfile/cue/410/midipara1  I32 F_XET
    /-show/showfile/cue/410/midipara2  I32 F_XET
/-show/showfile/cue/411  <SCUE> n=0
    /-show/showfile/cue/411/numb  I32 F_XET
    /-show/showfile/cue/411/name  S32 F_XET
    /-show/showfile/cue/411/skip  I32 F_XET
    /-show/showfile/cue/411/scene  I32 F_XET
    /-show/showfile/cue/411/bit  I32 F_XET
    /-show/showfile/cue/411/miditype  I32 F_XET
    /-show/showfile/cue/411/midichan  I32 F_XET
    /-show/showfile/cue/411/midipara1  I32 F_XET
    /-show/showfile/cue/411/midipara2  I32 F_XET
/-show/showfile/cue/412  <SCUE> n=0
    /-show/showfile/cue/412/numb  I32 F_XET
    /-show/showfile/cue/412/name  S32 F_XET
    /-show/showfile/cue/412/skip  I32 F_XET
    /-show/showfile/cue/412/scene  I32 F_XET
    /-show/showfile/cue/412/bit  I32 F_XET
    /-show/showfile/cue/412/miditype  I32 F_XET
    /-show/showfile/cue/412/midichan  I32 F_XET
    /-show/showfile/cue/412/midipara1  I32 F_XET
    /-show/showfile/cue/412/midipara2  I32 F_XET
/-show/showfile/cue/413  <SCUE> n=0
    /-show/showfile/cue/413/numb  I32 F_XET
    /-show/showfile/cue/413/name  S32 F_XET
    /-show/showfile/cue/413/skip  I32 F_XET
    /-show/showfile/cue/413/scene  I32 F_XET
    /-show/showfile/cue/413/bit  I32 F_XET
    /-show/showfile/cue/413/miditype  I32 F_XET
    /-show/showfile/cue/413/midichan  I32 F_XET
    /-show/showfile/cue/413/midipara1  I32 F_XET
    /-show/showfile/cue/413/midipara2  I32 F_XET
/-show/showfile/cue/414  <SCUE> n=0
    /-show/showfile/cue/414/numb  I32 F_XET
    /-show/showfile/cue/414/name  S32 F_XET
    /-show/showfile/cue/414/skip  I32 F_XET
    /-show/showfile/cue/414/scene  I32 F_XET
    /-show/showfile/cue/414/bit  I32 F_XET
    /-show/showfile/cue/414/miditype  I32 F_XET
    /-show/showfile/cue/414/midichan  I32 F_XET
    /-show/showfile/cue/414/midipara1  I32 F_XET
    /-show/showfile/cue/414/midipara2  I32 F_XET
/-show/showfile/cue/415  <SCUE> n=0
    /-show/showfile/cue/415/numb  I32 F_XET
    /-show/showfile/cue/415/name  S32 F_XET
    /-show/showfile/cue/415/skip  I32 F_XET
    /-show/showfile/cue/415/scene  I32 F_XET
    /-show/showfile/cue/415/bit  I32 F_XET
    /-show/showfile/cue/415/miditype  I32 F_XET
    /-show/showfile/cue/415/midichan  I32 F_XET
    /-show/showfile/cue/415/midipara1  I32 F_XET
    /-show/showfile/cue/415/midipara2  I32 F_XET
/-show/showfile/cue/416  <SCUE> n=0
    /-show/showfile/cue/416/numb  I32 F_XET
    /-show/showfile/cue/416/name  S32 F_XET
    /-show/showfile/cue/416/skip  I32 F_XET
    /-show/showfile/cue/416/scene  I32 F_XET
    /-show/showfile/cue/416/bit  I32 F_XET
    /-show/showfile/cue/416/miditype  I32 F_XET
    /-show/showfile/cue/416/midichan  I32 F_XET
    /-show/showfile/cue/416/midipara1  I32 F_XET
    /-show/showfile/cue/416/midipara2  I32 F_XET
/-show/showfile/cue/417  <SCUE> n=0
    /-show/showfile/cue/417/numb  I32 F_XET
    /-show/showfile/cue/417/name  S32 F_XET
    /-show/showfile/cue/417/skip  I32 F_XET
    /-show/showfile/cue/417/scene  I32 F_XET
    /-show/showfile/cue/417/bit  I32 F_XET
    /-show/showfile/cue/417/miditype  I32 F_XET
    /-show/showfile/cue/417/midichan  I32 F_XET
    /-show/showfile/cue/417/midipara1  I32 F_XET
    /-show/showfile/cue/417/midipara2  I32 F_XET
/-show/showfile/cue/418  <SCUE> n=0
    /-show/showfile/cue/418/numb  I32 F_XET
    /-show/showfile/cue/418/name  S32 F_XET
    /-show/showfile/cue/418/skip  I32 F_XET
    /-show/showfile/cue/418/scene  I32 F_XET
    /-show/showfile/cue/418/bit  I32 F_XET
    /-show/showfile/cue/418/miditype  I32 F_XET
    /-show/showfile/cue/418/midichan  I32 F_XET
    /-show/showfile/cue/418/midipara1  I32 F_XET
    /-show/showfile/cue/418/midipara2  I32 F_XET
/-show/showfile/cue/419  <SCUE> n=0
    /-show/showfile/cue/419/numb  I32 F_XET
    /-show/showfile/cue/419/name  S32 F_XET
    /-show/showfile/cue/419/skip  I32 F_XET
    /-show/showfile/cue/419/scene  I32 F_XET
    /-show/showfile/cue/419/bit  I32 F_XET
    /-show/showfile/cue/419/miditype  I32 F_XET
    /-show/showfile/cue/419/midichan  I32 F_XET
    /-show/showfile/cue/419/midipara1  I32 F_XET
    /-show/showfile/cue/419/midipara2  I32 F_XET
/-show/showfile/cue/420  <SCUE> n=0
    /-show/showfile/cue/420/numb  I32 F_XET
    /-show/showfile/cue/420/name  S32 F_XET
    /-show/showfile/cue/420/skip  I32 F_XET
    /-show/showfile/cue/420/scene  I32 F_XET
    /-show/showfile/cue/420/bit  I32 F_XET
    /-show/showfile/cue/420/miditype  I32 F_XET
    /-show/showfile/cue/420/midichan  I32 F_XET
    /-show/showfile/cue/420/midipara1  I32 F_XET
    /-show/showfile/cue/420/midipara2  I32 F_XET
/-show/showfile/cue/421  <SCUE> n=0
    /-show/showfile/cue/421/numb  I32 F_XET
    /-show/showfile/cue/421/name  S32 F_XET
    /-show/showfile/cue/421/skip  I32 F_XET
    /-show/showfile/cue/421/scene  I32 F_XET
    /-show/showfile/cue/421/bit  I32 F_XET
    /-show/showfile/cue/421/miditype  I32 F_XET
    /-show/showfile/cue/421/midichan  I32 F_XET
    /-show/showfile/cue/421/midipara1  I32 F_XET
    /-show/showfile/cue/421/midipara2  I32 F_XET
/-show/showfile/cue/422  <SCUE> n=0
    /-show/showfile/cue/422/numb  I32 F_XET
    /-show/showfile/cue/422/name  S32 F_XET
    /-show/showfile/cue/422/skip  I32 F_XET
    /-show/showfile/cue/422/scene  I32 F_XET
    /-show/showfile/cue/422/bit  I32 F_XET
    /-show/showfile/cue/422/miditype  I32 F_XET
    /-show/showfile/cue/422/midichan  I32 F_XET
    /-show/showfile/cue/422/midipara1  I32 F_XET
    /-show/showfile/cue/422/midipara2  I32 F_XET
/-show/showfile/cue/423  <SCUE> n=0
    /-show/showfile/cue/423/numb  I32 F_XET
    /-show/showfile/cue/423/name  S32 F_XET
    /-show/showfile/cue/423/skip  I32 F_XET
    /-show/showfile/cue/423/scene  I32 F_XET
    /-show/showfile/cue/423/bit  I32 F_XET
    /-show/showfile/cue/423/miditype  I32 F_XET
    /-show/showfile/cue/423/midichan  I32 F_XET
    /-show/showfile/cue/423/midipara1  I32 F_XET
    /-show/showfile/cue/423/midipara2  I32 F_XET
/-show/showfile/cue/424  <SCUE> n=0
    /-show/showfile/cue/424/numb  I32 F_XET
    /-show/showfile/cue/424/name  S32 F_XET
    /-show/showfile/cue/424/skip  I32 F_XET
    /-show/showfile/cue/424/scene  I32 F_XET
    /-show/showfile/cue/424/bit  I32 F_XET
    /-show/showfile/cue/424/miditype  I32 F_XET
    /-show/showfile/cue/424/midichan  I32 F_XET
    /-show/showfile/cue/424/midipara1  I32 F_XET
    /-show/showfile/cue/424/midipara2  I32 F_XET
/-show/showfile/cue/425  <SCUE> n=0
    /-show/showfile/cue/425/numb  I32 F_XET
    /-show/showfile/cue/425/name  S32 F_XET
    /-show/showfile/cue/425/skip  I32 F_XET
    /-show/showfile/cue/425/scene  I32 F_XET
    /-show/showfile/cue/425/bit  I32 F_XET
    /-show/showfile/cue/425/miditype  I32 F_XET
    /-show/showfile/cue/425/midichan  I32 F_XET
    /-show/showfile/cue/425/midipara1  I32 F_XET
    /-show/showfile/cue/425/midipara2  I32 F_XET
/-show/showfile/cue/426  <SCUE> n=0
    /-show/showfile/cue/426/numb  I32 F_XET
    /-show/showfile/cue/426/name  S32 F_XET
    /-show/showfile/cue/426/skip  I32 F_XET
    /-show/showfile/cue/426/scene  I32 F_XET
    /-show/showfile/cue/426/bit  I32 F_XET
    /-show/showfile/cue/426/miditype  I32 F_XET
    /-show/showfile/cue/426/midichan  I32 F_XET
    /-show/showfile/cue/426/midipara1  I32 F_XET
    /-show/showfile/cue/426/midipara2  I32 F_XET
/-show/showfile/cue/427  <SCUE> n=0
    /-show/showfile/cue/427/numb  I32 F_XET
    /-show/showfile/cue/427/name  S32 F_XET
    /-show/showfile/cue/427/skip  I32 F_XET
    /-show/showfile/cue/427/scene  I32 F_XET
    /-show/showfile/cue/427/bit  I32 F_XET
    /-show/showfile/cue/427/miditype  I32 F_XET
    /-show/showfile/cue/427/midichan  I32 F_XET
    /-show/showfile/cue/427/midipara1  I32 F_XET
    /-show/showfile/cue/427/midipara2  I32 F_XET
/-show/showfile/cue/428  <SCUE> n=0
    /-show/showfile/cue/428/numb  I32 F_XET
    /-show/showfile/cue/428/name  S32 F_XET
    /-show/showfile/cue/428/skip  I32 F_XET
    /-show/showfile/cue/428/scene  I32 F_XET
    /-show/showfile/cue/428/bit  I32 F_XET
    /-show/showfile/cue/428/miditype  I32 F_XET
    /-show/showfile/cue/428/midichan  I32 F_XET
    /-show/showfile/cue/428/midipara1  I32 F_XET
    /-show/showfile/cue/428/midipara2  I32 F_XET
/-show/showfile/cue/429  <SCUE> n=0
    /-show/showfile/cue/429/numb  I32 F_XET
    /-show/showfile/cue/429/name  S32 F_XET
    /-show/showfile/cue/429/skip  I32 F_XET
    /-show/showfile/cue/429/scene  I32 F_XET
    /-show/showfile/cue/429/bit  I32 F_XET
    /-show/showfile/cue/429/miditype  I32 F_XET
    /-show/showfile/cue/429/midichan  I32 F_XET
    /-show/showfile/cue/429/midipara1  I32 F_XET
    /-show/showfile/cue/429/midipara2  I32 F_XET
/-show/showfile/cue/430  <SCUE> n=0
    /-show/showfile/cue/430/numb  I32 F_XET
    /-show/showfile/cue/430/name  S32 F_XET
    /-show/showfile/cue/430/skip  I32 F_XET
    /-show/showfile/cue/430/scene  I32 F_XET
    /-show/showfile/cue/430/bit  I32 F_XET
    /-show/showfile/cue/430/miditype  I32 F_XET
    /-show/showfile/cue/430/midichan  I32 F_XET
    /-show/showfile/cue/430/midipara1  I32 F_XET
    /-show/showfile/cue/430/midipara2  I32 F_XET
/-show/showfile/cue/431  <SCUE> n=0
    /-show/showfile/cue/431/numb  I32 F_XET
    /-show/showfile/cue/431/name  S32 F_XET
    /-show/showfile/cue/431/skip  I32 F_XET
    /-show/showfile/cue/431/scene  I32 F_XET
    /-show/showfile/cue/431/bit  I32 F_XET
    /-show/showfile/cue/431/miditype  I32 F_XET
    /-show/showfile/cue/431/midichan  I32 F_XET
    /-show/showfile/cue/431/midipara1  I32 F_XET
    /-show/showfile/cue/431/midipara2  I32 F_XET
/-show/showfile/cue/432  <SCUE> n=0
    /-show/showfile/cue/432/numb  I32 F_XET
    /-show/showfile/cue/432/name  S32 F_XET
    /-show/showfile/cue/432/skip  I32 F_XET
    /-show/showfile/cue/432/scene  I32 F_XET
    /-show/showfile/cue/432/bit  I32 F_XET
    /-show/showfile/cue/432/miditype  I32 F_XET
    /-show/showfile/cue/432/midichan  I32 F_XET
    /-show/showfile/cue/432/midipara1  I32 F_XET
    /-show/showfile/cue/432/midipara2  I32 F_XET
/-show/showfile/cue/433  <SCUE> n=0
    /-show/showfile/cue/433/numb  I32 F_XET
    /-show/showfile/cue/433/name  S32 F_XET
    /-show/showfile/cue/433/skip  I32 F_XET
    /-show/showfile/cue/433/scene  I32 F_XET
    /-show/showfile/cue/433/bit  I32 F_XET
    /-show/showfile/cue/433/miditype  I32 F_XET
    /-show/showfile/cue/433/midichan  I32 F_XET
    /-show/showfile/cue/433/midipara1  I32 F_XET
    /-show/showfile/cue/433/midipara2  I32 F_XET
/-show/showfile/cue/434  <SCUE> n=0
    /-show/showfile/cue/434/numb  I32 F_XET
    /-show/showfile/cue/434/name  S32 F_XET
    /-show/showfile/cue/434/skip  I32 F_XET
    /-show/showfile/cue/434/scene  I32 F_XET
    /-show/showfile/cue/434/bit  I32 F_XET
    /-show/showfile/cue/434/miditype  I32 F_XET
    /-show/showfile/cue/434/midichan  I32 F_XET
    /-show/showfile/cue/434/midipara1  I32 F_XET
    /-show/showfile/cue/434/midipara2  I32 F_XET
/-show/showfile/cue/435  <SCUE> n=0
    /-show/showfile/cue/435/numb  I32 F_XET
    /-show/showfile/cue/435/name  S32 F_XET
    /-show/showfile/cue/435/skip  I32 F_XET
    /-show/showfile/cue/435/scene  I32 F_XET
    /-show/showfile/cue/435/bit  I32 F_XET
    /-show/showfile/cue/435/miditype  I32 F_XET
    /-show/showfile/cue/435/midichan  I32 F_XET
    /-show/showfile/cue/435/midipara1  I32 F_XET
    /-show/showfile/cue/435/midipara2  I32 F_XET
/-show/showfile/cue/436  <SCUE> n=0
    /-show/showfile/cue/436/numb  I32 F_XET
    /-show/showfile/cue/436/name  S32 F_XET
    /-show/showfile/cue/436/skip  I32 F_XET
    /-show/showfile/cue/436/scene  I32 F_XET
    /-show/showfile/cue/436/bit  I32 F_XET
    /-show/showfile/cue/436/miditype  I32 F_XET
    /-show/showfile/cue/436/midichan  I32 F_XET
    /-show/showfile/cue/436/midipara1  I32 F_XET
    /-show/showfile/cue/436/midipara2  I32 F_XET
/-show/showfile/cue/437  <SCUE> n=0
    /-show/showfile/cue/437/numb  I32 F_XET
    /-show/showfile/cue/437/name  S32 F_XET
    /-show/showfile/cue/437/skip  I32 F_XET
    /-show/showfile/cue/437/scene  I32 F_XET
    /-show/showfile/cue/437/bit  I32 F_XET
    /-show/showfile/cue/437/miditype  I32 F_XET
    /-show/showfile/cue/437/midichan  I32 F_XET
    /-show/showfile/cue/437/midipara1  I32 F_XET
    /-show/showfile/cue/437/midipara2  I32 F_XET
/-show/showfile/cue/438  <SCUE> n=0
    /-show/showfile/cue/438/numb  I32 F_XET
    /-show/showfile/cue/438/name  S32 F_XET
    /-show/showfile/cue/438/skip  I32 F_XET
    /-show/showfile/cue/438/scene  I32 F_XET
    /-show/showfile/cue/438/bit  I32 F_XET
    /-show/showfile/cue/438/miditype  I32 F_XET
    /-show/showfile/cue/438/midichan  I32 F_XET
    /-show/showfile/cue/438/midipara1  I32 F_XET
    /-show/showfile/cue/438/midipara2  I32 F_XET
/-show/showfile/cue/439  <SCUE> n=0
    /-show/showfile/cue/439/numb  I32 F_XET
    /-show/showfile/cue/439/name  S32 F_XET
    /-show/showfile/cue/439/skip  I32 F_XET
    /-show/showfile/cue/439/scene  I32 F_XET
    /-show/showfile/cue/439/bit  I32 F_XET
    /-show/showfile/cue/439/miditype  I32 F_XET
    /-show/showfile/cue/439/midichan  I32 F_XET
    /-show/showfile/cue/439/midipara1  I32 F_XET
    /-show/showfile/cue/439/midipara2  I32 F_XET
/-show/showfile/cue/440  <SCUE> n=0
    /-show/showfile/cue/440/numb  I32 F_XET
    /-show/showfile/cue/440/name  S32 F_XET
    /-show/showfile/cue/440/skip  I32 F_XET
    /-show/showfile/cue/440/scene  I32 F_XET
    /-show/showfile/cue/440/bit  I32 F_XET
    /-show/showfile/cue/440/miditype  I32 F_XET
    /-show/showfile/cue/440/midichan  I32 F_XET
    /-show/showfile/cue/440/midipara1  I32 F_XET
    /-show/showfile/cue/440/midipara2  I32 F_XET
/-show/showfile/cue/441  <SCUE> n=0
    /-show/showfile/cue/441/numb  I32 F_XET
    /-show/showfile/cue/441/name  S32 F_XET
    /-show/showfile/cue/441/skip  I32 F_XET
    /-show/showfile/cue/441/scene  I32 F_XET
    /-show/showfile/cue/441/bit  I32 F_XET
    /-show/showfile/cue/441/miditype  I32 F_XET
    /-show/showfile/cue/441/midichan  I32 F_XET
    /-show/showfile/cue/441/midipara1  I32 F_XET
    /-show/showfile/cue/441/midipara2  I32 F_XET
/-show/showfile/cue/442  <SCUE> n=0
    /-show/showfile/cue/442/numb  I32 F_XET
    /-show/showfile/cue/442/name  S32 F_XET
    /-show/showfile/cue/442/skip  I32 F_XET
    /-show/showfile/cue/442/scene  I32 F_XET
    /-show/showfile/cue/442/bit  I32 F_XET
    /-show/showfile/cue/442/miditype  I32 F_XET
    /-show/showfile/cue/442/midichan  I32 F_XET
    /-show/showfile/cue/442/midipara1  I32 F_XET
    /-show/showfile/cue/442/midipara2  I32 F_XET
/-show/showfile/cue/443  <SCUE> n=0
    /-show/showfile/cue/443/numb  I32 F_XET
    /-show/showfile/cue/443/name  S32 F_XET
    /-show/showfile/cue/443/skip  I32 F_XET
    /-show/showfile/cue/443/scene  I32 F_XET
    /-show/showfile/cue/443/bit  I32 F_XET
    /-show/showfile/cue/443/miditype  I32 F_XET
    /-show/showfile/cue/443/midichan  I32 F_XET
    /-show/showfile/cue/443/midipara1  I32 F_XET
    /-show/showfile/cue/443/midipara2  I32 F_XET
/-show/showfile/cue/444  <SCUE> n=0
    /-show/showfile/cue/444/numb  I32 F_XET
    /-show/showfile/cue/444/name  S32 F_XET
    /-show/showfile/cue/444/skip  I32 F_XET
    /-show/showfile/cue/444/scene  I32 F_XET
    /-show/showfile/cue/444/bit  I32 F_XET
    /-show/showfile/cue/444/miditype  I32 F_XET
    /-show/showfile/cue/444/midichan  I32 F_XET
    /-show/showfile/cue/444/midipara1  I32 F_XET
    /-show/showfile/cue/444/midipara2  I32 F_XET
/-show/showfile/cue/445  <SCUE> n=0
    /-show/showfile/cue/445/numb  I32 F_XET
    /-show/showfile/cue/445/name  S32 F_XET
    /-show/showfile/cue/445/skip  I32 F_XET
    /-show/showfile/cue/445/scene  I32 F_XET
    /-show/showfile/cue/445/bit  I32 F_XET
    /-show/showfile/cue/445/miditype  I32 F_XET
    /-show/showfile/cue/445/midichan  I32 F_XET
    /-show/showfile/cue/445/midipara1  I32 F_XET
    /-show/showfile/cue/445/midipara2  I32 F_XET
/-show/showfile/cue/446  <SCUE> n=0
    /-show/showfile/cue/446/numb  I32 F_XET
    /-show/showfile/cue/446/name  S32 F_XET
    /-show/showfile/cue/446/skip  I32 F_XET
    /-show/showfile/cue/446/scene  I32 F_XET
    /-show/showfile/cue/446/bit  I32 F_XET
    /-show/showfile/cue/446/miditype  I32 F_XET
    /-show/showfile/cue/446/midichan  I32 F_XET
    /-show/showfile/cue/446/midipara1  I32 F_XET
    /-show/showfile/cue/446/midipara2  I32 F_XET
/-show/showfile/cue/447  <SCUE> n=0
    /-show/showfile/cue/447/numb  I32 F_XET
    /-show/showfile/cue/447/name  S32 F_XET
    /-show/showfile/cue/447/skip  I32 F_XET
    /-show/showfile/cue/447/scene  I32 F_XET
    /-show/showfile/cue/447/bit  I32 F_XET
    /-show/showfile/cue/447/miditype  I32 F_XET
    /-show/showfile/cue/447/midichan  I32 F_XET
    /-show/showfile/cue/447/midipara1  I32 F_XET
    /-show/showfile/cue/447/midipara2  I32 F_XET
/-show/showfile/cue/448  <SCUE> n=0
    /-show/showfile/cue/448/numb  I32 F_XET
    /-show/showfile/cue/448/name  S32 F_XET
    /-show/showfile/cue/448/skip  I32 F_XET
    /-show/showfile/cue/448/scene  I32 F_XET
    /-show/showfile/cue/448/bit  I32 F_XET
    /-show/showfile/cue/448/miditype  I32 F_XET
    /-show/showfile/cue/448/midichan  I32 F_XET
    /-show/showfile/cue/448/midipara1  I32 F_XET
    /-show/showfile/cue/448/midipara2  I32 F_XET
/-show/showfile/cue/449  <SCUE> n=0
    /-show/showfile/cue/449/numb  I32 F_XET
    /-show/showfile/cue/449/name  S32 F_XET
    /-show/showfile/cue/449/skip  I32 F_XET
    /-show/showfile/cue/449/scene  I32 F_XET
    /-show/showfile/cue/449/bit  I32 F_XET
    /-show/showfile/cue/449/miditype  I32 F_XET
    /-show/showfile/cue/449/midichan  I32 F_XET
    /-show/showfile/cue/449/midipara1  I32 F_XET
    /-show/showfile/cue/449/midipara2  I32 F_XET
/-show/showfile/cue/450  <SCUE> n=0
    /-show/showfile/cue/450/numb  I32 F_XET
    /-show/showfile/cue/450/name  S32 F_XET
    /-show/showfile/cue/450/skip  I32 F_XET
    /-show/showfile/cue/450/scene  I32 F_XET
    /-show/showfile/cue/450/bit  I32 F_XET
    /-show/showfile/cue/450/miditype  I32 F_XET
    /-show/showfile/cue/450/midichan  I32 F_XET
    /-show/showfile/cue/450/midipara1  I32 F_XET
    /-show/showfile/cue/450/midipara2  I32 F_XET
/-show/showfile/cue/451  <SCUE> n=0
    /-show/showfile/cue/451/numb  I32 F_XET
    /-show/showfile/cue/451/name  S32 F_XET
    /-show/showfile/cue/451/skip  I32 F_XET
    /-show/showfile/cue/451/scene  I32 F_XET
    /-show/showfile/cue/451/bit  I32 F_XET
    /-show/showfile/cue/451/miditype  I32 F_XET
    /-show/showfile/cue/451/midichan  I32 F_XET
    /-show/showfile/cue/451/midipara1  I32 F_XET
    /-show/showfile/cue/451/midipara2  I32 F_XET
/-show/showfile/cue/452  <SCUE> n=0
    /-show/showfile/cue/452/numb  I32 F_XET
    /-show/showfile/cue/452/name  S32 F_XET
    /-show/showfile/cue/452/skip  I32 F_XET
    /-show/showfile/cue/452/scene  I32 F_XET
    /-show/showfile/cue/452/bit  I32 F_XET
    /-show/showfile/cue/452/miditype  I32 F_XET
    /-show/showfile/cue/452/midichan  I32 F_XET
    /-show/showfile/cue/452/midipara1  I32 F_XET
    /-show/showfile/cue/452/midipara2  I32 F_XET
/-show/showfile/cue/453  <SCUE> n=0
    /-show/showfile/cue/453/numb  I32 F_XET
    /-show/showfile/cue/453/name  S32 F_XET
    /-show/showfile/cue/453/skip  I32 F_XET
    /-show/showfile/cue/453/scene  I32 F_XET
    /-show/showfile/cue/453/bit  I32 F_XET
    /-show/showfile/cue/453/miditype  I32 F_XET
    /-show/showfile/cue/453/midichan  I32 F_XET
    /-show/showfile/cue/453/midipara1  I32 F_XET
    /-show/showfile/cue/453/midipara2  I32 F_XET
/-show/showfile/cue/454  <SCUE> n=0
    /-show/showfile/cue/454/numb  I32 F_XET
    /-show/showfile/cue/454/name  S32 F_XET
    /-show/showfile/cue/454/skip  I32 F_XET
    /-show/showfile/cue/454/scene  I32 F_XET
    /-show/showfile/cue/454/bit  I32 F_XET
    /-show/showfile/cue/454/miditype  I32 F_XET
    /-show/showfile/cue/454/midichan  I32 F_XET
    /-show/showfile/cue/454/midipara1  I32 F_XET
    /-show/showfile/cue/454/midipara2  I32 F_XET
/-show/showfile/cue/455  <SCUE> n=0
    /-show/showfile/cue/455/numb  I32 F_XET
    /-show/showfile/cue/455/name  S32 F_XET
    /-show/showfile/cue/455/skip  I32 F_XET
    /-show/showfile/cue/455/scene  I32 F_XET
    /-show/showfile/cue/455/bit  I32 F_XET
    /-show/showfile/cue/455/miditype  I32 F_XET
    /-show/showfile/cue/455/midichan  I32 F_XET
    /-show/showfile/cue/455/midipara1  I32 F_XET
    /-show/showfile/cue/455/midipara2  I32 F_XET
/-show/showfile/cue/456  <SCUE> n=0
    /-show/showfile/cue/456/numb  I32 F_XET
    /-show/showfile/cue/456/name  S32 F_XET
    /-show/showfile/cue/456/skip  I32 F_XET
    /-show/showfile/cue/456/scene  I32 F_XET
    /-show/showfile/cue/456/bit  I32 F_XET
    /-show/showfile/cue/456/miditype  I32 F_XET
    /-show/showfile/cue/456/midichan  I32 F_XET
    /-show/showfile/cue/456/midipara1  I32 F_XET
    /-show/showfile/cue/456/midipara2  I32 F_XET
/-show/showfile/cue/457  <SCUE> n=0
    /-show/showfile/cue/457/numb  I32 F_XET
    /-show/showfile/cue/457/name  S32 F_XET
    /-show/showfile/cue/457/skip  I32 F_XET
    /-show/showfile/cue/457/scene  I32 F_XET
    /-show/showfile/cue/457/bit  I32 F_XET
    /-show/showfile/cue/457/miditype  I32 F_XET
    /-show/showfile/cue/457/midichan  I32 F_XET
    /-show/showfile/cue/457/midipara1  I32 F_XET
    /-show/showfile/cue/457/midipara2  I32 F_XET
/-show/showfile/cue/458  <SCUE> n=0
    /-show/showfile/cue/458/numb  I32 F_XET
    /-show/showfile/cue/458/name  S32 F_XET
    /-show/showfile/cue/458/skip  I32 F_XET
    /-show/showfile/cue/458/scene  I32 F_XET
    /-show/showfile/cue/458/bit  I32 F_XET
    /-show/showfile/cue/458/miditype  I32 F_XET
    /-show/showfile/cue/458/midichan  I32 F_XET
    /-show/showfile/cue/458/midipara1  I32 F_XET
    /-show/showfile/cue/458/midipara2  I32 F_XET
/-show/showfile/cue/459  <SCUE> n=0
    /-show/showfile/cue/459/numb  I32 F_XET
    /-show/showfile/cue/459/name  S32 F_XET
    /-show/showfile/cue/459/skip  I32 F_XET
    /-show/showfile/cue/459/scene  I32 F_XET
    /-show/showfile/cue/459/bit  I32 F_XET
    /-show/showfile/cue/459/miditype  I32 F_XET
    /-show/showfile/cue/459/midichan  I32 F_XET
    /-show/showfile/cue/459/midipara1  I32 F_XET
    /-show/showfile/cue/459/midipara2  I32 F_XET
/-show/showfile/cue/460  <SCUE> n=0
    /-show/showfile/cue/460/numb  I32 F_XET
    /-show/showfile/cue/460/name  S32 F_XET
    /-show/showfile/cue/460/skip  I32 F_XET
    /-show/showfile/cue/460/scene  I32 F_XET
    /-show/showfile/cue/460/bit  I32 F_XET
    /-show/showfile/cue/460/miditype  I32 F_XET
    /-show/showfile/cue/460/midichan  I32 F_XET
    /-show/showfile/cue/460/midipara1  I32 F_XET
    /-show/showfile/cue/460/midipara2  I32 F_XET
/-show/showfile/cue/461  <SCUE> n=0
    /-show/showfile/cue/461/numb  I32 F_XET
    /-show/showfile/cue/461/name  S32 F_XET
    /-show/showfile/cue/461/skip  I32 F_XET
    /-show/showfile/cue/461/scene  I32 F_XET
    /-show/showfile/cue/461/bit  I32 F_XET
    /-show/showfile/cue/461/miditype  I32 F_XET
    /-show/showfile/cue/461/midichan  I32 F_XET
    /-show/showfile/cue/461/midipara1  I32 F_XET
    /-show/showfile/cue/461/midipara2  I32 F_XET
/-show/showfile/cue/462  <SCUE> n=0
    /-show/showfile/cue/462/numb  I32 F_XET
    /-show/showfile/cue/462/name  S32 F_XET
    /-show/showfile/cue/462/skip  I32 F_XET
    /-show/showfile/cue/462/scene  I32 F_XET
    /-show/showfile/cue/462/bit  I32 F_XET
    /-show/showfile/cue/462/miditype  I32 F_XET
    /-show/showfile/cue/462/midichan  I32 F_XET
    /-show/showfile/cue/462/midipara1  I32 F_XET
    /-show/showfile/cue/462/midipara2  I32 F_XET
/-show/showfile/cue/463  <SCUE> n=0
    /-show/showfile/cue/463/numb  I32 F_XET
    /-show/showfile/cue/463/name  S32 F_XET
    /-show/showfile/cue/463/skip  I32 F_XET
    /-show/showfile/cue/463/scene  I32 F_XET
    /-show/showfile/cue/463/bit  I32 F_XET
    /-show/showfile/cue/463/miditype  I32 F_XET
    /-show/showfile/cue/463/midichan  I32 F_XET
    /-show/showfile/cue/463/midipara1  I32 F_XET
    /-show/showfile/cue/463/midipara2  I32 F_XET
/-show/showfile/cue/464  <SCUE> n=0
    /-show/showfile/cue/464/numb  I32 F_XET
    /-show/showfile/cue/464/name  S32 F_XET
    /-show/showfile/cue/464/skip  I32 F_XET
    /-show/showfile/cue/464/scene  I32 F_XET
    /-show/showfile/cue/464/bit  I32 F_XET
    /-show/showfile/cue/464/miditype  I32 F_XET
    /-show/showfile/cue/464/midichan  I32 F_XET
    /-show/showfile/cue/464/midipara1  I32 F_XET
    /-show/showfile/cue/464/midipara2  I32 F_XET
/-show/showfile/cue/465  <SCUE> n=0
    /-show/showfile/cue/465/numb  I32 F_XET
    /-show/showfile/cue/465/name  S32 F_XET
    /-show/showfile/cue/465/skip  I32 F_XET
    /-show/showfile/cue/465/scene  I32 F_XET
    /-show/showfile/cue/465/bit  I32 F_XET
    /-show/showfile/cue/465/miditype  I32 F_XET
    /-show/showfile/cue/465/midichan  I32 F_XET
    /-show/showfile/cue/465/midipara1  I32 F_XET
    /-show/showfile/cue/465/midipara2  I32 F_XET
/-show/showfile/cue/466  <SCUE> n=0
    /-show/showfile/cue/466/numb  I32 F_XET
    /-show/showfile/cue/466/name  S32 F_XET
    /-show/showfile/cue/466/skip  I32 F_XET
    /-show/showfile/cue/466/scene  I32 F_XET
    /-show/showfile/cue/466/bit  I32 F_XET
    /-show/showfile/cue/466/miditype  I32 F_XET
    /-show/showfile/cue/466/midichan  I32 F_XET
    /-show/showfile/cue/466/midipara1  I32 F_XET
    /-show/showfile/cue/466/midipara2  I32 F_XET
/-show/showfile/cue/467  <SCUE> n=0
    /-show/showfile/cue/467/numb  I32 F_XET
    /-show/showfile/cue/467/name  S32 F_XET
    /-show/showfile/cue/467/skip  I32 F_XET
    /-show/showfile/cue/467/scene  I32 F_XET
    /-show/showfile/cue/467/bit  I32 F_XET
    /-show/showfile/cue/467/miditype  I32 F_XET
    /-show/showfile/cue/467/midichan  I32 F_XET
    /-show/showfile/cue/467/midipara1  I32 F_XET
    /-show/showfile/cue/467/midipara2  I32 F_XET
/-show/showfile/cue/468  <SCUE> n=0
    /-show/showfile/cue/468/numb  I32 F_XET
    /-show/showfile/cue/468/name  S32 F_XET
    /-show/showfile/cue/468/skip  I32 F_XET
    /-show/showfile/cue/468/scene  I32 F_XET
    /-show/showfile/cue/468/bit  I32 F_XET
    /-show/showfile/cue/468/miditype  I32 F_XET
    /-show/showfile/cue/468/midichan  I32 F_XET
    /-show/showfile/cue/468/midipara1  I32 F_XET
    /-show/showfile/cue/468/midipara2  I32 F_XET
/-show/showfile/cue/469  <SCUE> n=0
    /-show/showfile/cue/469/numb  I32 F_XET
    /-show/showfile/cue/469/name  S32 F_XET
    /-show/showfile/cue/469/skip  I32 F_XET
    /-show/showfile/cue/469/scene  I32 F_XET
    /-show/showfile/cue/469/bit  I32 F_XET
    /-show/showfile/cue/469/miditype  I32 F_XET
    /-show/showfile/cue/469/midichan  I32 F_XET
    /-show/showfile/cue/469/midipara1  I32 F_XET
    /-show/showfile/cue/469/midipara2  I32 F_XET
/-show/showfile/cue/470  <SCUE> n=0
    /-show/showfile/cue/470/numb  I32 F_XET
    /-show/showfile/cue/470/name  S32 F_XET
    /-show/showfile/cue/470/skip  I32 F_XET
    /-show/showfile/cue/470/scene  I32 F_XET
    /-show/showfile/cue/470/bit  I32 F_XET
    /-show/showfile/cue/470/miditype  I32 F_XET
    /-show/showfile/cue/470/midichan  I32 F_XET
    /-show/showfile/cue/470/midipara1  I32 F_XET
    /-show/showfile/cue/470/midipara2  I32 F_XET
/-show/showfile/cue/471  <SCUE> n=0
    /-show/showfile/cue/471/numb  I32 F_XET
    /-show/showfile/cue/471/name  S32 F_XET
    /-show/showfile/cue/471/skip  I32 F_XET
    /-show/showfile/cue/471/scene  I32 F_XET
    /-show/showfile/cue/471/bit  I32 F_XET
    /-show/showfile/cue/471/miditype  I32 F_XET
    /-show/showfile/cue/471/midichan  I32 F_XET
    /-show/showfile/cue/471/midipara1  I32 F_XET
    /-show/showfile/cue/471/midipara2  I32 F_XET
/-show/showfile/cue/472  <SCUE> n=0
    /-show/showfile/cue/472/numb  I32 F_XET
    /-show/showfile/cue/472/name  S32 F_XET
    /-show/showfile/cue/472/skip  I32 F_XET
    /-show/showfile/cue/472/scene  I32 F_XET
    /-show/showfile/cue/472/bit  I32 F_XET
    /-show/showfile/cue/472/miditype  I32 F_XET
    /-show/showfile/cue/472/midichan  I32 F_XET
    /-show/showfile/cue/472/midipara1  I32 F_XET
    /-show/showfile/cue/472/midipara2  I32 F_XET
/-show/showfile/cue/473  <SCUE> n=0
    /-show/showfile/cue/473/numb  I32 F_XET
    /-show/showfile/cue/473/name  S32 F_XET
    /-show/showfile/cue/473/skip  I32 F_XET
    /-show/showfile/cue/473/scene  I32 F_XET
    /-show/showfile/cue/473/bit  I32 F_XET
    /-show/showfile/cue/473/miditype  I32 F_XET
    /-show/showfile/cue/473/midichan  I32 F_XET
    /-show/showfile/cue/473/midipara1  I32 F_XET
    /-show/showfile/cue/473/midipara2  I32 F_XET
/-show/showfile/cue/474  <SCUE> n=0
    /-show/showfile/cue/474/numb  I32 F_XET
    /-show/showfile/cue/474/name  S32 F_XET
    /-show/showfile/cue/474/skip  I32 F_XET
    /-show/showfile/cue/474/scene  I32 F_XET
    /-show/showfile/cue/474/bit  I32 F_XET
    /-show/showfile/cue/474/miditype  I32 F_XET
    /-show/showfile/cue/474/midichan  I32 F_XET
    /-show/showfile/cue/474/midipara1  I32 F_XET
    /-show/showfile/cue/474/midipara2  I32 F_XET
/-show/showfile/cue/475  <SCUE> n=0
    /-show/showfile/cue/475/numb  I32 F_XET
    /-show/showfile/cue/475/name  S32 F_XET
    /-show/showfile/cue/475/skip  I32 F_XET
    /-show/showfile/cue/475/scene  I32 F_XET
    /-show/showfile/cue/475/bit  I32 F_XET
    /-show/showfile/cue/475/miditype  I32 F_XET
    /-show/showfile/cue/475/midichan  I32 F_XET
    /-show/showfile/cue/475/midipara1  I32 F_XET
    /-show/showfile/cue/475/midipara2  I32 F_XET
/-show/showfile/cue/476  <SCUE> n=0
    /-show/showfile/cue/476/numb  I32 F_XET
    /-show/showfile/cue/476/name  S32 F_XET
    /-show/showfile/cue/476/skip  I32 F_XET
    /-show/showfile/cue/476/scene  I32 F_XET
    /-show/showfile/cue/476/bit  I32 F_XET
    /-show/showfile/cue/476/miditype  I32 F_XET
    /-show/showfile/cue/476/midichan  I32 F_XET
    /-show/showfile/cue/476/midipara1  I32 F_XET
    /-show/showfile/cue/476/midipara2  I32 F_XET
/-show/showfile/cue/477  <SCUE> n=0
    /-show/showfile/cue/477/numb  I32 F_XET
    /-show/showfile/cue/477/name  S32 F_XET
    /-show/showfile/cue/477/skip  I32 F_XET
    /-show/showfile/cue/477/scene  I32 F_XET
    /-show/showfile/cue/477/bit  I32 F_XET
    /-show/showfile/cue/477/miditype  I32 F_XET
    /-show/showfile/cue/477/midichan  I32 F_XET
    /-show/showfile/cue/477/midipara1  I32 F_XET
    /-show/showfile/cue/477/midipara2  I32 F_XET
/-show/showfile/cue/478  <SCUE> n=0
    /-show/showfile/cue/478/numb  I32 F_XET
    /-show/showfile/cue/478/name  S32 F_XET
    /-show/showfile/cue/478/skip  I32 F_XET
    /-show/showfile/cue/478/scene  I32 F_XET
    /-show/showfile/cue/478/bit  I32 F_XET
    /-show/showfile/cue/478/miditype  I32 F_XET
    /-show/showfile/cue/478/midichan  I32 F_XET
    /-show/showfile/cue/478/midipara1  I32 F_XET
    /-show/showfile/cue/478/midipara2  I32 F_XET
/-show/showfile/cue/479  <SCUE> n=0
    /-show/showfile/cue/479/numb  I32 F_XET
    /-show/showfile/cue/479/name  S32 F_XET
    /-show/showfile/cue/479/skip  I32 F_XET
    /-show/showfile/cue/479/scene  I32 F_XET
    /-show/showfile/cue/479/bit  I32 F_XET
    /-show/showfile/cue/479/miditype  I32 F_XET
    /-show/showfile/cue/479/midichan  I32 F_XET
    /-show/showfile/cue/479/midipara1  I32 F_XET
    /-show/showfile/cue/479/midipara2  I32 F_XET
/-show/showfile/cue/480  <SCUE> n=0
    /-show/showfile/cue/480/numb  I32 F_XET
    /-show/showfile/cue/480/name  S32 F_XET
    /-show/showfile/cue/480/skip  I32 F_XET
    /-show/showfile/cue/480/scene  I32 F_XET
    /-show/showfile/cue/480/bit  I32 F_XET
    /-show/showfile/cue/480/miditype  I32 F_XET
    /-show/showfile/cue/480/midichan  I32 F_XET
    /-show/showfile/cue/480/midipara1  I32 F_XET
    /-show/showfile/cue/480/midipara2  I32 F_XET
/-show/showfile/cue/481  <SCUE> n=0
    /-show/showfile/cue/481/numb  I32 F_XET
    /-show/showfile/cue/481/name  S32 F_XET
    /-show/showfile/cue/481/skip  I32 F_XET
    /-show/showfile/cue/481/scene  I32 F_XET
    /-show/showfile/cue/481/bit  I32 F_XET
    /-show/showfile/cue/481/miditype  I32 F_XET
    /-show/showfile/cue/481/midichan  I32 F_XET
    /-show/showfile/cue/481/midipara1  I32 F_XET
    /-show/showfile/cue/481/midipara2  I32 F_XET
/-show/showfile/cue/482  <SCUE> n=0
    /-show/showfile/cue/482/numb  I32 F_XET
    /-show/showfile/cue/482/name  S32 F_XET
    /-show/showfile/cue/482/skip  I32 F_XET
    /-show/showfile/cue/482/scene  I32 F_XET
    /-show/showfile/cue/482/bit  I32 F_XET
    /-show/showfile/cue/482/miditype  I32 F_XET
    /-show/showfile/cue/482/midichan  I32 F_XET
    /-show/showfile/cue/482/midipara1  I32 F_XET
    /-show/showfile/cue/482/midipara2  I32 F_XET
/-show/showfile/cue/483  <SCUE> n=0
    /-show/showfile/cue/483/numb  I32 F_XET
    /-show/showfile/cue/483/name  S32 F_XET
    /-show/showfile/cue/483/skip  I32 F_XET
    /-show/showfile/cue/483/scene  I32 F_XET
    /-show/showfile/cue/483/bit  I32 F_XET
    /-show/showfile/cue/483/miditype  I32 F_XET
    /-show/showfile/cue/483/midichan  I32 F_XET
    /-show/showfile/cue/483/midipara1  I32 F_XET
    /-show/showfile/cue/483/midipara2  I32 F_XET
/-show/showfile/cue/484  <SCUE> n=0
    /-show/showfile/cue/484/numb  I32 F_XET
    /-show/showfile/cue/484/name  S32 F_XET
    /-show/showfile/cue/484/skip  I32 F_XET
    /-show/showfile/cue/484/scene  I32 F_XET
    /-show/showfile/cue/484/bit  I32 F_XET
    /-show/showfile/cue/484/miditype  I32 F_XET
    /-show/showfile/cue/484/midichan  I32 F_XET
    /-show/showfile/cue/484/midipara1  I32 F_XET
    /-show/showfile/cue/484/midipara2  I32 F_XET
/-show/showfile/cue/485  <SCUE> n=0
    /-show/showfile/cue/485/numb  I32 F_XET
    /-show/showfile/cue/485/name  S32 F_XET
    /-show/showfile/cue/485/skip  I32 F_XET
    /-show/showfile/cue/485/scene  I32 F_XET
    /-show/showfile/cue/485/bit  I32 F_XET
    /-show/showfile/cue/485/miditype  I32 F_XET
    /-show/showfile/cue/485/midichan  I32 F_XET
    /-show/showfile/cue/485/midipara1  I32 F_XET
    /-show/showfile/cue/485/midipara2  I32 F_XET
/-show/showfile/cue/486  <SCUE> n=0
    /-show/showfile/cue/486/numb  I32 F_XET
    /-show/showfile/cue/486/name  S32 F_XET
    /-show/showfile/cue/486/skip  I32 F_XET
    /-show/showfile/cue/486/scene  I32 F_XET
    /-show/showfile/cue/486/bit  I32 F_XET
    /-show/showfile/cue/486/miditype  I32 F_XET
    /-show/showfile/cue/486/midichan  I32 F_XET
    /-show/showfile/cue/486/midipara1  I32 F_XET
    /-show/showfile/cue/486/midipara2  I32 F_XET
/-show/showfile/cue/487  <SCUE> n=0
    /-show/showfile/cue/487/numb  I32 F_XET
    /-show/showfile/cue/487/name  S32 F_XET
    /-show/showfile/cue/487/skip  I32 F_XET
    /-show/showfile/cue/487/scene  I32 F_XET
    /-show/showfile/cue/487/bit  I32 F_XET
    /-show/showfile/cue/487/miditype  I32 F_XET
    /-show/showfile/cue/487/midichan  I32 F_XET
    /-show/showfile/cue/487/midipara1  I32 F_XET
    /-show/showfile/cue/487/midipara2  I32 F_XET
/-show/showfile/cue/488  <SCUE> n=0
    /-show/showfile/cue/488/numb  I32 F_XET
    /-show/showfile/cue/488/name  S32 F_XET
    /-show/showfile/cue/488/skip  I32 F_XET
    /-show/showfile/cue/488/scene  I32 F_XET
    /-show/showfile/cue/488/bit  I32 F_XET
    /-show/showfile/cue/488/miditype  I32 F_XET
    /-show/showfile/cue/488/midichan  I32 F_XET
    /-show/showfile/cue/488/midipara1  I32 F_XET
    /-show/showfile/cue/488/midipara2  I32 F_XET
/-show/showfile/cue/489  <SCUE> n=0
    /-show/showfile/cue/489/numb  I32 F_XET
    /-show/showfile/cue/489/name  S32 F_XET
    /-show/showfile/cue/489/skip  I32 F_XET
    /-show/showfile/cue/489/scene  I32 F_XET
    /-show/showfile/cue/489/bit  I32 F_XET
    /-show/showfile/cue/489/miditype  I32 F_XET
    /-show/showfile/cue/489/midichan  I32 F_XET
    /-show/showfile/cue/489/midipara1  I32 F_XET
    /-show/showfile/cue/489/midipara2  I32 F_XET
/-show/showfile/cue/490  <SCUE> n=0
    /-show/showfile/cue/490/numb  I32 F_XET
    /-show/showfile/cue/490/name  S32 F_XET
    /-show/showfile/cue/490/skip  I32 F_XET
    /-show/showfile/cue/490/scene  I32 F_XET
    /-show/showfile/cue/490/bit  I32 F_XET
    /-show/showfile/cue/490/miditype  I32 F_XET
    /-show/showfile/cue/490/midichan  I32 F_XET
    /-show/showfile/cue/490/midipara1  I32 F_XET
    /-show/showfile/cue/490/midipara2  I32 F_XET
/-show/showfile/cue/491  <SCUE> n=0
    /-show/showfile/cue/491/numb  I32 F_XET
    /-show/showfile/cue/491/name  S32 F_XET
    /-show/showfile/cue/491/skip  I32 F_XET
    /-show/showfile/cue/491/scene  I32 F_XET
    /-show/showfile/cue/491/bit  I32 F_XET
    /-show/showfile/cue/491/miditype  I32 F_XET
    /-show/showfile/cue/491/midichan  I32 F_XET
    /-show/showfile/cue/491/midipara1  I32 F_XET
    /-show/showfile/cue/491/midipara2  I32 F_XET
/-show/showfile/cue/492  <SCUE> n=0
    /-show/showfile/cue/492/numb  I32 F_XET
    /-show/showfile/cue/492/name  S32 F_XET
    /-show/showfile/cue/492/skip  I32 F_XET
    /-show/showfile/cue/492/scene  I32 F_XET
    /-show/showfile/cue/492/bit  I32 F_XET
    /-show/showfile/cue/492/miditype  I32 F_XET
    /-show/showfile/cue/492/midichan  I32 F_XET
    /-show/showfile/cue/492/midipara1  I32 F_XET
    /-show/showfile/cue/492/midipara2  I32 F_XET
/-show/showfile/cue/493  <SCUE> n=0
    /-show/showfile/cue/493/numb  I32 F_XET
    /-show/showfile/cue/493/name  S32 F_XET
    /-show/showfile/cue/493/skip  I32 F_XET
    /-show/showfile/cue/493/scene  I32 F_XET
    /-show/showfile/cue/493/bit  I32 F_XET
    /-show/showfile/cue/493/miditype  I32 F_XET
    /-show/showfile/cue/493/midichan  I32 F_XET
    /-show/showfile/cue/493/midipara1  I32 F_XET
    /-show/showfile/cue/493/midipara2  I32 F_XET
/-show/showfile/cue/494  <SCUE> n=0
    /-show/showfile/cue/494/numb  I32 F_XET
    /-show/showfile/cue/494/name  S32 F_XET
    /-show/showfile/cue/494/skip  I32 F_XET
    /-show/showfile/cue/494/scene  I32 F_XET
    /-show/showfile/cue/494/bit  I32 F_XET
    /-show/showfile/cue/494/miditype  I32 F_XET
    /-show/showfile/cue/494/midichan  I32 F_XET
    /-show/showfile/cue/494/midipara1  I32 F_XET
    /-show/showfile/cue/494/midipara2  I32 F_XET
/-show/showfile/cue/495  <SCUE> n=0
    /-show/showfile/cue/495/numb  I32 F_XET
    /-show/showfile/cue/495/name  S32 F_XET
    /-show/showfile/cue/495/skip  I32 F_XET
    /-show/showfile/cue/495/scene  I32 F_XET
    /-show/showfile/cue/495/bit  I32 F_XET
    /-show/showfile/cue/495/miditype  I32 F_XET
    /-show/showfile/cue/495/midichan  I32 F_XET
    /-show/showfile/cue/495/midipara1  I32 F_XET
    /-show/showfile/cue/495/midipara2  I32 F_XET
/-show/showfile/cue/496  <SCUE> n=0
    /-show/showfile/cue/496/numb  I32 F_XET
    /-show/showfile/cue/496/name  S32 F_XET
    /-show/showfile/cue/496/skip  I32 F_XET
    /-show/showfile/cue/496/scene  I32 F_XET
    /-show/showfile/cue/496/bit  I32 F_XET
    /-show/showfile/cue/496/miditype  I32 F_XET
    /-show/showfile/cue/496/midichan  I32 F_XET
    /-show/showfile/cue/496/midipara1  I32 F_XET
    /-show/showfile/cue/496/midipara2  I32 F_XET
/-show/showfile/cue/497  <SCUE> n=0
    /-show/showfile/cue/497/numb  I32 F_XET
    /-show/showfile/cue/497/name  S32 F_XET
    /-show/showfile/cue/497/skip  I32 F_XET
    /-show/showfile/cue/497/scene  I32 F_XET
    /-show/showfile/cue/497/bit  I32 F_XET
    /-show/showfile/cue/497/miditype  I32 F_XET
    /-show/showfile/cue/497/midichan  I32 F_XET
    /-show/showfile/cue/497/midipara1  I32 F_XET
    /-show/showfile/cue/497/midipara2  I32 F_XET
/-show/showfile/cue/498  <SCUE> n=0
    /-show/showfile/cue/498/numb  I32 F_XET
    /-show/showfile/cue/498/name  S32 F_XET
    /-show/showfile/cue/498/skip  I32 F_XET
    /-show/showfile/cue/498/scene  I32 F_XET
    /-show/showfile/cue/498/bit  I32 F_XET
    /-show/showfile/cue/498/miditype  I32 F_XET
    /-show/showfile/cue/498/midichan  I32 F_XET
    /-show/showfile/cue/498/midipara1  I32 F_XET
    /-show/showfile/cue/498/midipara2  I32 F_XET
/-show/showfile/cue/499  <SCUE> n=0
    /-show/showfile/cue/499/numb  I32 F_XET
    /-show/showfile/cue/499/name  S32 F_XET
    /-show/showfile/cue/499/skip  I32 F_XET
    /-show/showfile/cue/499/scene  I32 F_XET
    /-show/showfile/cue/499/bit  I32 F_XET
    /-show/showfile/cue/499/miditype  I32 F_XET
    /-show/showfile/cue/499/midichan  I32 F_XET
    /-show/showfile/cue/499/midipara1  I32 F_XET
    /-show/showfile/cue/499/midipara2  I32 F_XET
```

### Xscene (X32Show.h, 501 entries)

```
/-show/showfile/scene  <SSCN> n=0
/-show/showfile/scene/000  <SSCN> n=0
    /-show/showfile/scene/000/name  S32 F_XET
    /-show/showfile/scene/000/notes  S32 F_XET
    /-show/showfile/scene/000/safes  P32 F_XET
    /-show/showfile/scene/000/hasdata  I32 F_XET
/-show/showfile/scene/001  <SSCN> n=0
    /-show/showfile/scene/001/name  S32 F_XET
    /-show/showfile/scene/001/notes  S32 F_XET
    /-show/showfile/scene/001/safes  P32 F_XET
    /-show/showfile/scene/001/hasdata  I32 F_XET
/-show/showfile/scene/002  <SSCN> n=0
    /-show/showfile/scene/002/name  S32 F_XET
    /-show/showfile/scene/002/notes  S32 F_XET
    /-show/showfile/scene/002/safes  P32 F_XET
    /-show/showfile/scene/002/hasdata  I32 F_XET
/-show/showfile/scene/003  <SSCN> n=0
    /-show/showfile/scene/003/name  S32 F_XET
    /-show/showfile/scene/003/notes  S32 F_XET
    /-show/showfile/scene/003/safes  P32 F_XET
    /-show/showfile/scene/003/hasdata  I32 F_XET
/-show/showfile/scene/004  <SSCN> n=0
    /-show/showfile/scene/004/name  S32 F_XET
    /-show/showfile/scene/004/notes  S32 F_XET
    /-show/showfile/scene/004/safes  P32 F_XET
    /-show/showfile/scene/004/hasdata  I32 F_XET
/-show/showfile/scene/005  <SSCN> n=0
    /-show/showfile/scene/005/name  S32 F_XET
    /-show/showfile/scene/005/notes  S32 F_XET
    /-show/showfile/scene/005/safes  P32 F_XET
    /-show/showfile/scene/005/hasdata  I32 F_XET
/-show/showfile/scene/006  <SSCN> n=0
    /-show/showfile/scene/006/name  S32 F_XET
    /-show/showfile/scene/006/notes  S32 F_XET
    /-show/showfile/scene/006/safes  P32 F_XET
    /-show/showfile/scene/006/hasdata  I32 F_XET
/-show/showfile/scene/007  <SSCN> n=0
    /-show/showfile/scene/007/name  S32 F_XET
    /-show/showfile/scene/007/notes  S32 F_XET
    /-show/showfile/scene/007/safes  P32 F_XET
    /-show/showfile/scene/007/hasdata  I32 F_XET
/-show/showfile/scene/008  <SSCN> n=0
    /-show/showfile/scene/008/name  S32 F_XET
    /-show/showfile/scene/008/notes  S32 F_XET
    /-show/showfile/scene/008/safes  P32 F_XET
    /-show/showfile/scene/008/hasdata  I32 F_XET
/-show/showfile/scene/009  <SSCN> n=0
    /-show/showfile/scene/009/name  S32 F_XET
    /-show/showfile/scene/009/notes  S32 F_XET
    /-show/showfile/scene/009/safes  P32 F_XET
    /-show/showfile/scene/009/hasdata  I32 F_XET
/-show/showfile/scene/010  <SSCN> n=0
    /-show/showfile/scene/010/name  S32 F_XET
    /-show/showfile/scene/010/notes  S32 F_XET
    /-show/showfile/scene/010/safes  P32 F_XET
    /-show/showfile/scene/010/hasdata  I32 F_XET
/-show/showfile/scene/011  <SSCN> n=0
    /-show/showfile/scene/011/name  S32 F_XET
    /-show/showfile/scene/011/notes  S32 F_XET
    /-show/showfile/scene/011/safes  P32 F_XET
    /-show/showfile/scene/011/hasdata  I32 F_XET
/-show/showfile/scene/012  <SSCN> n=0
    /-show/showfile/scene/012/name  S32 F_XET
    /-show/showfile/scene/012/notes  S32 F_XET
    /-show/showfile/scene/012/safes  P32 F_XET
    /-show/showfile/scene/012/hasdata  I32 F_XET
/-show/showfile/scene/013  <SSCN> n=0
    /-show/showfile/scene/013/name  S32 F_XET
    /-show/showfile/scene/013/notes  S32 F_XET
    /-show/showfile/scene/013/safes  P32 F_XET
    /-show/showfile/scene/013/hasdata  I32 F_XET
/-show/showfile/scene/014  <SSCN> n=0
    /-show/showfile/scene/014/name  S32 F_XET
    /-show/showfile/scene/014/notes  S32 F_XET
    /-show/showfile/scene/014/safes  P32 F_XET
    /-show/showfile/scene/014/hasdata  I32 F_XET
/-show/showfile/scene/015  <SSCN> n=0
    /-show/showfile/scene/015/name  S32 F_XET
    /-show/showfile/scene/015/notes  S32 F_XET
    /-show/showfile/scene/015/safes  P32 F_XET
    /-show/showfile/scene/015/hasdata  I32 F_XET
/-show/showfile/scene/016  <SSCN> n=0
    /-show/showfile/scene/016/name  S32 F_XET
    /-show/showfile/scene/016/notes  S32 F_XET
    /-show/showfile/scene/016/safes  P32 F_XET
    /-show/showfile/scene/016/hasdata  I32 F_XET
/-show/showfile/scene/017  <SSCN> n=0
    /-show/showfile/scene/017/name  S32 F_XET
    /-show/showfile/scene/017/notes  S32 F_XET
    /-show/showfile/scene/017/safes  P32 F_XET
    /-show/showfile/scene/017/hasdata  I32 F_XET
/-show/showfile/scene/018  <SSCN> n=0
    /-show/showfile/scene/018/name  S32 F_XET
    /-show/showfile/scene/018/notes  S32 F_XET
    /-show/showfile/scene/018/safes  P32 F_XET
    /-show/showfile/scene/018/hasdata  I32 F_XET
/-show/showfile/scene/019  <SSCN> n=0
    /-show/showfile/scene/019/name  S32 F_XET
    /-show/showfile/scene/019/notes  S32 F_XET
    /-show/showfile/scene/019/safes  P32 F_XET
    /-show/showfile/scene/019/hasdata  I32 F_XET
/-show/showfile/scene/020  <SSCN> n=0
    /-show/showfile/scene/020/name  S32 F_XET
    /-show/showfile/scene/020/notes  S32 F_XET
    /-show/showfile/scene/020/safes  P32 F_XET
    /-show/showfile/scene/020/hasdata  I32 F_XET
/-show/showfile/scene/021  <SSCN> n=0
    /-show/showfile/scene/021/name  S32 F_XET
    /-show/showfile/scene/021/notes  S32 F_XET
    /-show/showfile/scene/021/safes  P32 F_XET
    /-show/showfile/scene/021/hasdata  I32 F_XET
/-show/showfile/scene/022  <SSCN> n=0
    /-show/showfile/scene/022/name  S32 F_XET
    /-show/showfile/scene/022/notes  S32 F_XET
    /-show/showfile/scene/022/safes  P32 F_XET
    /-show/showfile/scene/022/hasdata  I32 F_XET
/-show/showfile/scene/023  <SSCN> n=0
    /-show/showfile/scene/023/name  S32 F_XET
    /-show/showfile/scene/023/notes  S32 F_XET
    /-show/showfile/scene/023/safes  P32 F_XET
    /-show/showfile/scene/023/hasdata  I32 F_XET
/-show/showfile/scene/024  <SSCN> n=0
    /-show/showfile/scene/024/name  S32 F_XET
    /-show/showfile/scene/024/notes  S32 F_XET
    /-show/showfile/scene/024/safes  P32 F_XET
    /-show/showfile/scene/024/hasdata  I32 F_XET
/-show/showfile/scene/025  <SSCN> n=0
    /-show/showfile/scene/025/name  S32 F_XET
    /-show/showfile/scene/025/notes  S32 F_XET
    /-show/showfile/scene/025/safes  P32 F_XET
    /-show/showfile/scene/025/hasdata  I32 F_XET
/-show/showfile/scene/026  <SSCN> n=0
    /-show/showfile/scene/026/name  S32 F_XET
    /-show/showfile/scene/026/notes  S32 F_XET
    /-show/showfile/scene/026/safes  P32 F_XET
    /-show/showfile/scene/026/hasdata  I32 F_XET
/-show/showfile/scene/027  <SSCN> n=0
    /-show/showfile/scene/027/name  S32 F_XET
    /-show/showfile/scene/027/notes  S32 F_XET
    /-show/showfile/scene/027/safes  P32 F_XET
    /-show/showfile/scene/027/hasdata  I32 F_XET
/-show/showfile/scene/028  <SSCN> n=0
    /-show/showfile/scene/028/name  S32 F_XET
    /-show/showfile/scene/028/notes  S32 F_XET
    /-show/showfile/scene/028/safes  P32 F_XET
    /-show/showfile/scene/028/hasdata  I32 F_XET
/-show/showfile/scene/029  <SSCN> n=0
    /-show/showfile/scene/029/name  S32 F_XET
    /-show/showfile/scene/029/notes  S32 F_XET
    /-show/showfile/scene/029/safes  P32 F_XET
    /-show/showfile/scene/029/hasdata  I32 F_XET
/-show/showfile/scene/030  <SSCN> n=0
    /-show/showfile/scene/030/name  S32 F_XET
    /-show/showfile/scene/030/notes  S32 F_XET
    /-show/showfile/scene/030/safes  P32 F_XET
    /-show/showfile/scene/030/hasdata  I32 F_XET
/-show/showfile/scene/031  <SSCN> n=0
    /-show/showfile/scene/031/name  S32 F_XET
    /-show/showfile/scene/031/notes  S32 F_XET
    /-show/showfile/scene/031/safes  P32 F_XET
    /-show/showfile/scene/031/hasdata  I32 F_XET
/-show/showfile/scene/032  <SSCN> n=0
    /-show/showfile/scene/032/name  S32 F_XET
    /-show/showfile/scene/032/notes  S32 F_XET
    /-show/showfile/scene/032/safes  P32 F_XET
    /-show/showfile/scene/032/hasdata  I32 F_XET
/-show/showfile/scene/033  <SSCN> n=0
    /-show/showfile/scene/033/name  S32 F_XET
    /-show/showfile/scene/033/notes  S32 F_XET
    /-show/showfile/scene/033/safes  P32 F_XET
    /-show/showfile/scene/033/hasdata  I32 F_XET
/-show/showfile/scene/034  <SSCN> n=0
    /-show/showfile/scene/034/name  S32 F_XET
    /-show/showfile/scene/034/notes  S32 F_XET
    /-show/showfile/scene/034/safes  P32 F_XET
    /-show/showfile/scene/034/hasdata  I32 F_XET
/-show/showfile/scene/035  <SSCN> n=0
    /-show/showfile/scene/035/name  S32 F_XET
    /-show/showfile/scene/035/notes  S32 F_XET
    /-show/showfile/scene/035/safes  P32 F_XET
    /-show/showfile/scene/035/hasdata  I32 F_XET
/-show/showfile/scene/036  <SSCN> n=0
    /-show/showfile/scene/036/name  S32 F_XET
    /-show/showfile/scene/036/notes  S32 F_XET
    /-show/showfile/scene/036/safes  P32 F_XET
    /-show/showfile/scene/036/hasdata  I32 F_XET
/-show/showfile/scene/037  <SSCN> n=0
    /-show/showfile/scene/037/name  S32 F_XET
    /-show/showfile/scene/037/notes  S32 F_XET
    /-show/showfile/scene/037/safes  P32 F_XET
    /-show/showfile/scene/037/hasdata  I32 F_XET
/-show/showfile/scene/038  <SSCN> n=0
    /-show/showfile/scene/038/name  S32 F_XET
    /-show/showfile/scene/038/notes  S32 F_XET
    /-show/showfile/scene/038/safes  P32 F_XET
    /-show/showfile/scene/038/hasdata  I32 F_XET
/-show/showfile/scene/039  <SSCN> n=0
    /-show/showfile/scene/039/name  S32 F_XET
    /-show/showfile/scene/039/notes  S32 F_XET
    /-show/showfile/scene/039/safes  P32 F_XET
    /-show/showfile/scene/039/hasdata  I32 F_XET
/-show/showfile/scene/040  <SSCN> n=0
    /-show/showfile/scene/040/name  S32 F_XET
    /-show/showfile/scene/040/notes  S32 F_XET
    /-show/showfile/scene/040/safes  P32 F_XET
    /-show/showfile/scene/040/hasdata  I32 F_XET
/-show/showfile/scene/041  <SSCN> n=0
    /-show/showfile/scene/041/name  S32 F_XET
    /-show/showfile/scene/041/notes  S32 F_XET
    /-show/showfile/scene/041/safes  P32 F_XET
    /-show/showfile/scene/041/hasdata  I32 F_XET
/-show/showfile/scene/042  <SSCN> n=0
    /-show/showfile/scene/042/name  S32 F_XET
    /-show/showfile/scene/042/notes  S32 F_XET
    /-show/showfile/scene/042/safes  P32 F_XET
    /-show/showfile/scene/042/hasdata  I32 F_XET
/-show/showfile/scene/043  <SSCN> n=0
    /-show/showfile/scene/043/name  S32 F_XET
    /-show/showfile/scene/043/notes  S32 F_XET
    /-show/showfile/scene/043/safes  P32 F_XET
    /-show/showfile/scene/043/hasdata  I32 F_XET
/-show/showfile/scene/044  <SSCN> n=0
    /-show/showfile/scene/044/name  S32 F_XET
    /-show/showfile/scene/044/notes  S32 F_XET
    /-show/showfile/scene/044/safes  P32 F_XET
    /-show/showfile/scene/044/hasdata  I32 F_XET
/-show/showfile/scene/045  <SSCN> n=0
    /-show/showfile/scene/045/name  S32 F_XET
    /-show/showfile/scene/045/notes  S32 F_XET
    /-show/showfile/scene/045/safes  P32 F_XET
    /-show/showfile/scene/045/hasdata  I32 F_XET
/-show/showfile/scene/046  <SSCN> n=0
    /-show/showfile/scene/046/name  S32 F_XET
    /-show/showfile/scene/046/notes  S32 F_XET
    /-show/showfile/scene/046/safes  P32 F_XET
    /-show/showfile/scene/046/hasdata  I32 F_XET
/-show/showfile/scene/047  <SSCN> n=0
    /-show/showfile/scene/047/name  S32 F_XET
    /-show/showfile/scene/047/notes  S32 F_XET
    /-show/showfile/scene/047/safes  P32 F_XET
    /-show/showfile/scene/047/hasdata  I32 F_XET
/-show/showfile/scene/048  <SSCN> n=0
    /-show/showfile/scene/048/name  S32 F_XET
    /-show/showfile/scene/048/notes  S32 F_XET
    /-show/showfile/scene/048/safes  P32 F_XET
    /-show/showfile/scene/048/hasdata  I32 F_XET
/-show/showfile/scene/049  <SSCN> n=0
    /-show/showfile/scene/049/name  S32 F_XET
    /-show/showfile/scene/049/notes  S32 F_XET
    /-show/showfile/scene/049/safes  P32 F_XET
    /-show/showfile/scene/049/hasdata  I32 F_XET
/-show/showfile/scene/050  <SSCN> n=0
    /-show/showfile/scene/050/name  S32 F_XET
    /-show/showfile/scene/050/notes  S32 F_XET
    /-show/showfile/scene/050/safes  P32 F_XET
    /-show/showfile/scene/050/hasdata  I32 F_XET
/-show/showfile/scene/051  <SSCN> n=0
    /-show/showfile/scene/051/name  S32 F_XET
    /-show/showfile/scene/051/notes  S32 F_XET
    /-show/showfile/scene/051/safes  P32 F_XET
    /-show/showfile/scene/051/hasdata  I32 F_XET
/-show/showfile/scene/052  <SSCN> n=0
    /-show/showfile/scene/052/name  S32 F_XET
    /-show/showfile/scene/052/notes  S32 F_XET
    /-show/showfile/scene/052/safes  P32 F_XET
    /-show/showfile/scene/052/hasdata  I32 F_XET
/-show/showfile/scene/053  <SSCN> n=0
    /-show/showfile/scene/053/name  S32 F_XET
    /-show/showfile/scene/053/notes  S32 F_XET
    /-show/showfile/scene/053/safes  P32 F_XET
    /-show/showfile/scene/053/hasdata  I32 F_XET
/-show/showfile/scene/054  <SSCN> n=0
    /-show/showfile/scene/054/name  S32 F_XET
    /-show/showfile/scene/054/notes  S32 F_XET
    /-show/showfile/scene/054/safes  P32 F_XET
    /-show/showfile/scene/054/hasdata  I32 F_XET
/-show/showfile/scene/055  <SSCN> n=0
    /-show/showfile/scene/055/name  S32 F_XET
    /-show/showfile/scene/055/notes  S32 F_XET
    /-show/showfile/scene/055/safes  P32 F_XET
    /-show/showfile/scene/055/hasdata  I32 F_XET
/-show/showfile/scene/056  <SSCN> n=0
    /-show/showfile/scene/056/name  S32 F_XET
    /-show/showfile/scene/056/notes  S32 F_XET
    /-show/showfile/scene/056/safes  P32 F_XET
    /-show/showfile/scene/056/hasdata  I32 F_XET
/-show/showfile/scene/057  <SSCN> n=0
    /-show/showfile/scene/057/name  S32 F_XET
    /-show/showfile/scene/057/notes  S32 F_XET
    /-show/showfile/scene/057/safes  P32 F_XET
    /-show/showfile/scene/057/hasdata  I32 F_XET
/-show/showfile/scene/058  <SSCN> n=0
    /-show/showfile/scene/058/name  S32 F_XET
    /-show/showfile/scene/058/notes  S32 F_XET
    /-show/showfile/scene/058/safes  P32 F_XET
    /-show/showfile/scene/058/hasdata  I32 F_XET
/-show/showfile/scene/059  <SSCN> n=0
    /-show/showfile/scene/059/name  S32 F_XET
    /-show/showfile/scene/059/notes  S32 F_XET
    /-show/showfile/scene/059/safes  P32 F_XET
    /-show/showfile/scene/059/hasdata  I32 F_XET
/-show/showfile/scene/060  <SSCN> n=0
    /-show/showfile/scene/060/name  S32 F_XET
    /-show/showfile/scene/060/notes  S32 F_XET
    /-show/showfile/scene/060/safes  P32 F_XET
    /-show/showfile/scene/060/hasdata  I32 F_XET
/-show/showfile/scene/061  <SSCN> n=0
    /-show/showfile/scene/061/name  S32 F_XET
    /-show/showfile/scene/061/notes  S32 F_XET
    /-show/showfile/scene/061/safes  P32 F_XET
    /-show/showfile/scene/061/hasdata  I32 F_XET
/-show/showfile/scene/062  <SSCN> n=0
    /-show/showfile/scene/062/name  S32 F_XET
    /-show/showfile/scene/062/notes  S32 F_XET
    /-show/showfile/scene/062/safes  P32 F_XET
    /-show/showfile/scene/062/hasdata  I32 F_XET
/-show/showfile/scene/063  <SSCN> n=0
    /-show/showfile/scene/063/name  S32 F_XET
    /-show/showfile/scene/063/notes  S32 F_XET
    /-show/showfile/scene/063/safes  P32 F_XET
    /-show/showfile/scene/063/hasdata  I32 F_XET
/-show/showfile/scene/064  <SSCN> n=0
    /-show/showfile/scene/064/name  S32 F_XET
    /-show/showfile/scene/064/notes  S32 F_XET
    /-show/showfile/scene/064/safes  P32 F_XET
    /-show/showfile/scene/064/hasdata  I32 F_XET
/-show/showfile/scene/065  <SSCN> n=0
    /-show/showfile/scene/065/name  S32 F_XET
    /-show/showfile/scene/065/notes  S32 F_XET
    /-show/showfile/scene/065/safes  P32 F_XET
    /-show/showfile/scene/065/hasdata  I32 F_XET
/-show/showfile/scene/066  <SSCN> n=0
    /-show/showfile/scene/066/name  S32 F_XET
    /-show/showfile/scene/066/notes  S32 F_XET
    /-show/showfile/scene/066/safes  P32 F_XET
    /-show/showfile/scene/066/hasdata  I32 F_XET
/-show/showfile/scene/067  <SSCN> n=0
    /-show/showfile/scene/067/name  S32 F_XET
    /-show/showfile/scene/067/notes  S32 F_XET
    /-show/showfile/scene/067/safes  P32 F_XET
    /-show/showfile/scene/067/hasdata  I32 F_XET
/-show/showfile/scene/068  <SSCN> n=0
    /-show/showfile/scene/068/name  S32 F_XET
    /-show/showfile/scene/068/notes  S32 F_XET
    /-show/showfile/scene/068/safes  P32 F_XET
    /-show/showfile/scene/068/hasdata  I32 F_XET
/-show/showfile/scene/069  <SSCN> n=0
    /-show/showfile/scene/069/name  S32 F_XET
    /-show/showfile/scene/069/notes  S32 F_XET
    /-show/showfile/scene/069/safes  P32 F_XET
    /-show/showfile/scene/069/hasdata  I32 F_XET
/-show/showfile/scene/070  <SSCN> n=0
    /-show/showfile/scene/070/name  S32 F_XET
    /-show/showfile/scene/070/notes  S32 F_XET
    /-show/showfile/scene/070/safes  P32 F_XET
    /-show/showfile/scene/070/hasdata  I32 F_XET
/-show/showfile/scene/071  <SSCN> n=0
    /-show/showfile/scene/071/name  S32 F_XET
    /-show/showfile/scene/071/notes  S32 F_XET
    /-show/showfile/scene/071/safes  P32 F_XET
    /-show/showfile/scene/071/hasdata  I32 F_XET
/-show/showfile/scene/072  <SSCN> n=0
    /-show/showfile/scene/072/name  S32 F_XET
    /-show/showfile/scene/072/notes  S32 F_XET
    /-show/showfile/scene/072/safes  P32 F_XET
    /-show/showfile/scene/072/hasdata  I32 F_XET
/-show/showfile/scene/073  <SSCN> n=0
    /-show/showfile/scene/073/name  S32 F_XET
    /-show/showfile/scene/073/notes  S32 F_XET
    /-show/showfile/scene/073/safes  P32 F_XET
    /-show/showfile/scene/073/hasdata  I32 F_XET
/-show/showfile/scene/074  <SSCN> n=0
    /-show/showfile/scene/074/name  S32 F_XET
    /-show/showfile/scene/074/notes  S32 F_XET
    /-show/showfile/scene/074/safes  P32 F_XET
    /-show/showfile/scene/074/hasdata  I32 F_XET
/-show/showfile/scene/075  <SSCN> n=0
    /-show/showfile/scene/075/name  S32 F_XET
    /-show/showfile/scene/075/notes  S32 F_XET
    /-show/showfile/scene/075/safes  P32 F_XET
    /-show/showfile/scene/075/hasdata  I32 F_XET
/-show/showfile/scene/076  <SSCN> n=0
    /-show/showfile/scene/076/name  S32 F_XET
    /-show/showfile/scene/076/notes  S32 F_XET
    /-show/showfile/scene/076/safes  P32 F_XET
    /-show/showfile/scene/076/hasdata  I32 F_XET
/-show/showfile/scene/077  <SSCN> n=0
    /-show/showfile/scene/077/name  S32 F_XET
    /-show/showfile/scene/077/notes  S32 F_XET
    /-show/showfile/scene/077/safes  P32 F_XET
    /-show/showfile/scene/077/hasdata  I32 F_XET
/-show/showfile/scene/078  <SSCN> n=0
    /-show/showfile/scene/078/name  S32 F_XET
    /-show/showfile/scene/078/notes  S32 F_XET
    /-show/showfile/scene/078/safes  P32 F_XET
    /-show/showfile/scene/078/hasdata  I32 F_XET
/-show/showfile/scene/079  <SSCN> n=0
    /-show/showfile/scene/079/name  S32 F_XET
    /-show/showfile/scene/079/notes  S32 F_XET
    /-show/showfile/scene/079/safes  P32 F_XET
    /-show/showfile/scene/079/hasdata  I32 F_XET
/-show/showfile/scene/080  <SSCN> n=0
    /-show/showfile/scene/080/name  S32 F_XET
    /-show/showfile/scene/080/notes  S32 F_XET
    /-show/showfile/scene/080/safes  P32 F_XET
    /-show/showfile/scene/080/hasdata  I32 F_XET
/-show/showfile/scene/081  <SSCN> n=0
    /-show/showfile/scene/081/name  S32 F_XET
    /-show/showfile/scene/081/notes  S32 F_XET
    /-show/showfile/scene/081/safes  P32 F_XET
    /-show/showfile/scene/081/hasdata  I32 F_XET
/-show/showfile/scene/082  <SSCN> n=0
    /-show/showfile/scene/082/name  S32 F_XET
    /-show/showfile/scene/082/notes  S32 F_XET
    /-show/showfile/scene/082/safes  P32 F_XET
    /-show/showfile/scene/082/hasdata  I32 F_XET
/-show/showfile/scene/083  <SSCN> n=0
    /-show/showfile/scene/083/name  S32 F_XET
    /-show/showfile/scene/083/notes  S32 F_XET
    /-show/showfile/scene/083/safes  P32 F_XET
    /-show/showfile/scene/083/hasdata  I32 F_XET
/-show/showfile/scene/084  <SSCN> n=0
    /-show/showfile/scene/084/name  S32 F_XET
    /-show/showfile/scene/084/notes  S32 F_XET
    /-show/showfile/scene/084/safes  P32 F_XET
    /-show/showfile/scene/084/hasdata  I32 F_XET
/-show/showfile/scene/085  <SSCN> n=0
    /-show/showfile/scene/085/name  S32 F_XET
    /-show/showfile/scene/085/notes  S32 F_XET
    /-show/showfile/scene/085/safes  P32 F_XET
    /-show/showfile/scene/085/hasdata  I32 F_XET
/-show/showfile/scene/086  <SSCN> n=0
    /-show/showfile/scene/086/name  S32 F_XET
    /-show/showfile/scene/086/notes  S32 F_XET
    /-show/showfile/scene/086/safes  P32 F_XET
    /-show/showfile/scene/086/hasdata  I32 F_XET
/-show/showfile/scene/087  <SSCN> n=0
    /-show/showfile/scene/087/name  S32 F_XET
    /-show/showfile/scene/087/notes  S32 F_XET
    /-show/showfile/scene/087/safes  P32 F_XET
    /-show/showfile/scene/087/hasdata  I32 F_XET
/-show/showfile/scene/088  <SSCN> n=0
    /-show/showfile/scene/088/name  S32 F_XET
    /-show/showfile/scene/088/notes  S32 F_XET
    /-show/showfile/scene/088/safes  P32 F_XET
    /-show/showfile/scene/088/hasdata  I32 F_XET
/-show/showfile/scene/089  <SSCN> n=0
    /-show/showfile/scene/089/name  S32 F_XET
    /-show/showfile/scene/089/notes  S32 F_XET
    /-show/showfile/scene/089/safes  P32 F_XET
    /-show/showfile/scene/089/hasdata  I32 F_XET
/-show/showfile/scene/090  <SSCN> n=0
    /-show/showfile/scene/090/name  S32 F_XET
    /-show/showfile/scene/090/notes  S32 F_XET
    /-show/showfile/scene/090/safes  P32 F_XET
    /-show/showfile/scene/090/hasdata  I32 F_XET
/-show/showfile/scene/091  <SSCN> n=0
    /-show/showfile/scene/091/name  S32 F_XET
    /-show/showfile/scene/091/notes  S32 F_XET
    /-show/showfile/scene/091/safes  P32 F_XET
    /-show/showfile/scene/091/hasdata  I32 F_XET
/-show/showfile/scene/092  <SSCN> n=0
    /-show/showfile/scene/092/name  S32 F_XET
    /-show/showfile/scene/092/notes  S32 F_XET
    /-show/showfile/scene/092/safes  P32 F_XET
    /-show/showfile/scene/092/hasdata  I32 F_XET
/-show/showfile/scene/093  <SSCN> n=0
    /-show/showfile/scene/093/name  S32 F_XET
    /-show/showfile/scene/093/notes  S32 F_XET
    /-show/showfile/scene/093/safes  P32 F_XET
    /-show/showfile/scene/093/hasdata  I32 F_XET
/-show/showfile/scene/094  <SSCN> n=0
    /-show/showfile/scene/094/name  S32 F_XET
    /-show/showfile/scene/094/notes  S32 F_XET
    /-show/showfile/scene/094/safes  P32 F_XET
    /-show/showfile/scene/094/hasdata  I32 F_XET
/-show/showfile/scene/095  <SSCN> n=0
    /-show/showfile/scene/095/name  S32 F_XET
    /-show/showfile/scene/095/notes  S32 F_XET
    /-show/showfile/scene/095/safes  P32 F_XET
    /-show/showfile/scene/095/hasdata  I32 F_XET
/-show/showfile/scene/096  <SSCN> n=0
    /-show/showfile/scene/096/name  S32 F_XET
    /-show/showfile/scene/096/notes  S32 F_XET
    /-show/showfile/scene/096/safes  P32 F_XET
    /-show/showfile/scene/096/hasdata  I32 F_XET
/-show/showfile/scene/097  <SSCN> n=0
    /-show/showfile/scene/097/name  S32 F_XET
    /-show/showfile/scene/097/notes  S32 F_XET
    /-show/showfile/scene/097/safes  P32 F_XET
    /-show/showfile/scene/097/hasdata  I32 F_XET
/-show/showfile/scene/098  <SSCN> n=0
    /-show/showfile/scene/098/name  S32 F_XET
    /-show/showfile/scene/098/notes  S32 F_XET
    /-show/showfile/scene/098/safes  P32 F_XET
    /-show/showfile/scene/098/hasdata  I32 F_XET
/-show/showfile/scene/099  <SSCN> n=0
    /-show/showfile/scene/099/name  S32 F_XET
    /-show/showfile/scene/099/notes  S32 F_XET
    /-show/showfile/scene/099/safes  P32 F_XET
    /-show/showfile/scene/099/hasdata  I32 F_XET
```

### Xsnippet (X32Show.h, 701 entries)

```
/-show/showfile/snippet  <SSNP> n=0
/-show/showfile/snippet/000  <SSNP> n=0
    /-show/showfile/snippet/000/name  S32 F_XET
    /-show/showfile/snippet/000/eventtyp  P32 F_XET
    /-show/showfile/snippet/000/channels  P32 F_XET
    /-show/showfile/snippet/000/auxbuses  P32 F_XET
    /-show/showfile/snippet/000/maingrps  P32 F_XET
    /-show/showfile/snippet/000/hasdata  I32 F_XET
/-show/showfile/snippet/001  <SSNP> n=0
    /-show/showfile/snippet/001/name  S32 F_XET
    /-show/showfile/snippet/001/eventtyp  P32 F_XET
    /-show/showfile/snippet/001/channels  P32 F_XET
    /-show/showfile/snippet/001/auxbuses  P32 F_XET
    /-show/showfile/snippet/001/maingrps  P32 F_XET
    /-show/showfile/snippet/001/hasdata  I32 F_XET
/-show/showfile/snippet/002  <SSNP> n=0
    /-show/showfile/snippet/002/name  S32 F_XET
    /-show/showfile/snippet/002/eventtyp  P32 F_XET
    /-show/showfile/snippet/002/channels  P32 F_XET
    /-show/showfile/snippet/002/auxbuses  P32 F_XET
    /-show/showfile/snippet/002/maingrps  P32 F_XET
    /-show/showfile/snippet/002/hasdata  I32 F_XET
/-show/showfile/snippet/003  <SSNP> n=0
    /-show/showfile/snippet/003/name  S32 F_XET
    /-show/showfile/snippet/003/eventtyp  P32 F_XET
    /-show/showfile/snippet/003/channels  P32 F_XET
    /-show/showfile/snippet/003/auxbuses  P32 F_XET
    /-show/showfile/snippet/003/maingrps  P32 F_XET
    /-show/showfile/snippet/003/hasdata  I32 F_XET
/-show/showfile/snippet/004  <SSNP> n=0
    /-show/showfile/snippet/004/name  S32 F_XET
    /-show/showfile/snippet/004/eventtyp  P32 F_XET
    /-show/showfile/snippet/004/channels  P32 F_XET
    /-show/showfile/snippet/004/auxbuses  P32 F_XET
    /-show/showfile/snippet/004/maingrps  P32 F_XET
    /-show/showfile/snippet/004/hasdata  I32 F_XET
/-show/showfile/snippet/005  <SSNP> n=0
    /-show/showfile/snippet/005/name  S32 F_XET
    /-show/showfile/snippet/005/eventtyp  P32 F_XET
    /-show/showfile/snippet/005/channels  P32 F_XET
    /-show/showfile/snippet/005/auxbuses  P32 F_XET
    /-show/showfile/snippet/005/maingrps  P32 F_XET
    /-show/showfile/snippet/005/hasdata  I32 F_XET
/-show/showfile/snippet/006  <SSNP> n=0
    /-show/showfile/snippet/006/name  S32 F_XET
    /-show/showfile/snippet/006/eventtyp  P32 F_XET
    /-show/showfile/snippet/006/channels  P32 F_XET
    /-show/showfile/snippet/006/auxbuses  P32 F_XET
    /-show/showfile/snippet/006/maingrps  P32 F_XET
    /-show/showfile/snippet/006/hasdata  I32 F_XET
/-show/showfile/snippet/007  <SSNP> n=0
    /-show/showfile/snippet/007/name  S32 F_XET
    /-show/showfile/snippet/007/eventtyp  P32 F_XET
    /-show/showfile/snippet/007/channels  P32 F_XET
    /-show/showfile/snippet/007/auxbuses  P32 F_XET
    /-show/showfile/snippet/007/maingrps  P32 F_XET
    /-show/showfile/snippet/007/hasdata  I32 F_XET
/-show/showfile/snippet/008  <SSNP> n=0
    /-show/showfile/snippet/008/name  S32 F_XET
    /-show/showfile/snippet/008/eventtyp  P32 F_XET
    /-show/showfile/snippet/008/channels  P32 F_XET
    /-show/showfile/snippet/008/auxbuses  P32 F_XET
    /-show/showfile/snippet/008/maingrps  P32 F_XET
    /-show/showfile/snippet/008/hasdata  I32 F_XET
/-show/showfile/snippet/009  <SSNP> n=0
    /-show/showfile/snippet/009/name  S32 F_XET
    /-show/showfile/snippet/009/eventtyp  P32 F_XET
    /-show/showfile/snippet/009/channels  P32 F_XET
    /-show/showfile/snippet/009/auxbuses  P32 F_XET
    /-show/showfile/snippet/009/maingrps  P32 F_XET
    /-show/showfile/snippet/009/hasdata  I32 F_XET
/-show/showfile/snippet/010  <SSNP> n=0
    /-show/showfile/snippet/010/name  S32 F_XET
    /-show/showfile/snippet/010/eventtyp  P32 F_XET
    /-show/showfile/snippet/010/channels  P32 F_XET
    /-show/showfile/snippet/010/auxbuses  P32 F_XET
    /-show/showfile/snippet/010/maingrps  P32 F_XET
    /-show/showfile/snippet/010/hasdata  I32 F_XET
/-show/showfile/snippet/011  <SSNP> n=0
    /-show/showfile/snippet/011/name  S32 F_XET
    /-show/showfile/snippet/011/eventtyp  P32 F_XET
    /-show/showfile/snippet/011/channels  P32 F_XET
    /-show/showfile/snippet/011/auxbuses  P32 F_XET
    /-show/showfile/snippet/011/maingrps  P32 F_XET
    /-show/showfile/snippet/011/hasdata  I32 F_XET
/-show/showfile/snippet/012  <SSNP> n=0
    /-show/showfile/snippet/012/name  S32 F_XET
    /-show/showfile/snippet/012/eventtyp  P32 F_XET
    /-show/showfile/snippet/012/channels  P32 F_XET
    /-show/showfile/snippet/012/auxbuses  P32 F_XET
    /-show/showfile/snippet/012/maingrps  P32 F_XET
    /-show/showfile/snippet/012/hasdata  I32 F_XET
/-show/showfile/snippet/013  <SSNP> n=0
    /-show/showfile/snippet/013/name  S32 F_XET
    /-show/showfile/snippet/013/eventtyp  P32 F_XET
    /-show/showfile/snippet/013/channels  P32 F_XET
    /-show/showfile/snippet/013/auxbuses  P32 F_XET
    /-show/showfile/snippet/013/maingrps  P32 F_XET
    /-show/showfile/snippet/013/hasdata  I32 F_XET
/-show/showfile/snippet/014  <SSNP> n=0
    /-show/showfile/snippet/014/name  S32 F_XET
    /-show/showfile/snippet/014/eventtyp  P32 F_XET
    /-show/showfile/snippet/014/channels  P32 F_XET
    /-show/showfile/snippet/014/auxbuses  P32 F_XET
    /-show/showfile/snippet/014/maingrps  P32 F_XET
    /-show/showfile/snippet/014/hasdata  I32 F_XET
/-show/showfile/snippet/015  <SSNP> n=0
    /-show/showfile/snippet/015/name  S32 F_XET
    /-show/showfile/snippet/015/eventtyp  P32 F_XET
    /-show/showfile/snippet/015/channels  P32 F_XET
    /-show/showfile/snippet/015/auxbuses  P32 F_XET
    /-show/showfile/snippet/015/maingrps  P32 F_XET
    /-show/showfile/snippet/015/hasdata  I32 F_XET
/-show/showfile/snippet/016  <SSNP> n=0
    /-show/showfile/snippet/016/name  S32 F_XET
    /-show/showfile/snippet/016/eventtyp  P32 F_XET
    /-show/showfile/snippet/016/channels  P32 F_XET
    /-show/showfile/snippet/016/auxbuses  P32 F_XET
    /-show/showfile/snippet/016/maingrps  P32 F_XET
    /-show/showfile/snippet/016/hasdata  I32 F_XET
/-show/showfile/snippet/017  <SSNP> n=0
    /-show/showfile/snippet/017/name  S32 F_XET
    /-show/showfile/snippet/017/eventtyp  P32 F_XET
    /-show/showfile/snippet/017/channels  P32 F_XET
    /-show/showfile/snippet/017/auxbuses  P32 F_XET
    /-show/showfile/snippet/017/maingrps  P32 F_XET
    /-show/showfile/snippet/017/hasdata  I32 F_XET
/-show/showfile/snippet/018  <SSNP> n=0
    /-show/showfile/snippet/018/name  S32 F_XET
    /-show/showfile/snippet/018/eventtyp  P32 F_XET
    /-show/showfile/snippet/018/channels  P32 F_XET
    /-show/showfile/snippet/018/auxbuses  P32 F_XET
    /-show/showfile/snippet/018/maingrps  P32 F_XET
    /-show/showfile/snippet/018/hasdata  I32 F_XET
/-show/showfile/snippet/019  <SSNP> n=0
    /-show/showfile/snippet/019/name  S32 F_XET
    /-show/showfile/snippet/019/eventtyp  P32 F_XET
    /-show/showfile/snippet/019/channels  P32 F_XET
    /-show/showfile/snippet/019/auxbuses  P32 F_XET
    /-show/showfile/snippet/019/maingrps  P32 F_XET
    /-show/showfile/snippet/019/hasdata  I32 F_XET
/-show/showfile/snippet/020  <SSNP> n=0
    /-show/showfile/snippet/020/name  S32 F_XET
    /-show/showfile/snippet/020/eventtyp  P32 F_XET
    /-show/showfile/snippet/020/channels  P32 F_XET
    /-show/showfile/snippet/020/auxbuses  P32 F_XET
    /-show/showfile/snippet/020/maingrps  P32 F_XET
    /-show/showfile/snippet/020/hasdata  I32 F_XET
/-show/showfile/snippet/021  <SSNP> n=0
    /-show/showfile/snippet/021/name  S32 F_XET
    /-show/showfile/snippet/021/eventtyp  P32 F_XET
    /-show/showfile/snippet/021/channels  P32 F_XET
    /-show/showfile/snippet/021/auxbuses  P32 F_XET
    /-show/showfile/snippet/021/maingrps  P32 F_XET
    /-show/showfile/snippet/021/hasdata  I32 F_XET
/-show/showfile/snippet/022  <SSNP> n=0
    /-show/showfile/snippet/022/name  S32 F_XET
    /-show/showfile/snippet/022/eventtyp  P32 F_XET
    /-show/showfile/snippet/022/channels  P32 F_XET
    /-show/showfile/snippet/022/auxbuses  P32 F_XET
    /-show/showfile/snippet/022/maingrps  P32 F_XET
    /-show/showfile/snippet/022/hasdata  I32 F_XET
/-show/showfile/snippet/023  <SSNP> n=0
    /-show/showfile/snippet/023/name  S32 F_XET
    /-show/showfile/snippet/023/eventtyp  P32 F_XET
    /-show/showfile/snippet/023/channels  P32 F_XET
    /-show/showfile/snippet/023/auxbuses  P32 F_XET
    /-show/showfile/snippet/023/maingrps  P32 F_XET
    /-show/showfile/snippet/023/hasdata  I32 F_XET
/-show/showfile/snippet/024  <SSNP> n=0
    /-show/showfile/snippet/024/name  S32 F_XET
    /-show/showfile/snippet/024/eventtyp  P32 F_XET
    /-show/showfile/snippet/024/channels  P32 F_XET
    /-show/showfile/snippet/024/auxbuses  P32 F_XET
    /-show/showfile/snippet/024/maingrps  P32 F_XET
    /-show/showfile/snippet/024/hasdata  I32 F_XET
/-show/showfile/snippet/025  <SSNP> n=0
    /-show/showfile/snippet/025/name  S32 F_XET
    /-show/showfile/snippet/025/eventtyp  P32 F_XET
    /-show/showfile/snippet/025/channels  P32 F_XET
    /-show/showfile/snippet/025/auxbuses  P32 F_XET
    /-show/showfile/snippet/025/maingrps  P32 F_XET
    /-show/showfile/snippet/025/hasdata  I32 F_XET
/-show/showfile/snippet/026  <SSNP> n=0
    /-show/showfile/snippet/026/name  S32 F_XET
    /-show/showfile/snippet/026/eventtyp  P32 F_XET
    /-show/showfile/snippet/026/channels  P32 F_XET
    /-show/showfile/snippet/026/auxbuses  P32 F_XET
    /-show/showfile/snippet/026/maingrps  P32 F_XET
    /-show/showfile/snippet/026/hasdata  I32 F_XET
/-show/showfile/snippet/027  <SSNP> n=0
    /-show/showfile/snippet/027/name  S32 F_XET
    /-show/showfile/snippet/027/eventtyp  P32 F_XET
    /-show/showfile/snippet/027/channels  P32 F_XET
    /-show/showfile/snippet/027/auxbuses  P32 F_XET
    /-show/showfile/snippet/027/maingrps  P32 F_XET
    /-show/showfile/snippet/027/hasdata  I32 F_XET
/-show/showfile/snippet/028  <SSNP> n=0
    /-show/showfile/snippet/028/name  S32 F_XET
    /-show/showfile/snippet/028/eventtyp  P32 F_XET
    /-show/showfile/snippet/028/channels  P32 F_XET
    /-show/showfile/snippet/028/auxbuses  P32 F_XET
    /-show/showfile/snippet/028/maingrps  P32 F_XET
    /-show/showfile/snippet/028/hasdata  I32 F_XET
/-show/showfile/snippet/029  <SSNP> n=0
    /-show/showfile/snippet/029/name  S32 F_XET
    /-show/showfile/snippet/029/eventtyp  P32 F_XET
    /-show/showfile/snippet/029/channels  P32 F_XET
    /-show/showfile/snippet/029/auxbuses  P32 F_XET
    /-show/showfile/snippet/029/maingrps  P32 F_XET
    /-show/showfile/snippet/029/hasdata  I32 F_XET
/-show/showfile/snippet/030  <SSNP> n=0
    /-show/showfile/snippet/030/name  S32 F_XET
    /-show/showfile/snippet/030/eventtyp  P32 F_XET
    /-show/showfile/snippet/030/channels  P32 F_XET
    /-show/showfile/snippet/030/auxbuses  P32 F_XET
    /-show/showfile/snippet/030/maingrps  P32 F_XET
    /-show/showfile/snippet/030/hasdata  I32 F_XET
/-show/showfile/snippet/031  <SSNP> n=0
    /-show/showfile/snippet/031/name  S32 F_XET
    /-show/showfile/snippet/031/eventtyp  P32 F_XET
    /-show/showfile/snippet/031/channels  P32 F_XET
    /-show/showfile/snippet/031/auxbuses  P32 F_XET
    /-show/showfile/snippet/031/maingrps  P32 F_XET
    /-show/showfile/snippet/031/hasdata  I32 F_XET
/-show/showfile/snippet/032  <SSNP> n=0
    /-show/showfile/snippet/032/name  S32 F_XET
    /-show/showfile/snippet/032/eventtyp  P32 F_XET
    /-show/showfile/snippet/032/channels  P32 F_XET
    /-show/showfile/snippet/032/auxbuses  P32 F_XET
    /-show/showfile/snippet/032/maingrps  P32 F_XET
    /-show/showfile/snippet/032/hasdata  I32 F_XET
/-show/showfile/snippet/033  <SSNP> n=0
    /-show/showfile/snippet/033/name  S32 F_XET
    /-show/showfile/snippet/033/eventtyp  P32 F_XET
    /-show/showfile/snippet/033/channels  P32 F_XET
    /-show/showfile/snippet/033/auxbuses  P32 F_XET
    /-show/showfile/snippet/033/maingrps  P32 F_XET
    /-show/showfile/snippet/033/hasdata  I32 F_XET
/-show/showfile/snippet/034  <SSNP> n=0
    /-show/showfile/snippet/034/name  S32 F_XET
    /-show/showfile/snippet/034/eventtyp  P32 F_XET
    /-show/showfile/snippet/034/channels  P32 F_XET
    /-show/showfile/snippet/034/auxbuses  P32 F_XET
    /-show/showfile/snippet/034/maingrps  P32 F_XET
    /-show/showfile/snippet/034/hasdata  I32 F_XET
/-show/showfile/snippet/035  <SSNP> n=0
    /-show/showfile/snippet/035/name  S32 F_XET
    /-show/showfile/snippet/035/eventtyp  P32 F_XET
    /-show/showfile/snippet/035/channels  P32 F_XET
    /-show/showfile/snippet/035/auxbuses  P32 F_XET
    /-show/showfile/snippet/035/maingrps  P32 F_XET
    /-show/showfile/snippet/035/hasdata  I32 F_XET
/-show/showfile/snippet/036  <SSNP> n=0
    /-show/showfile/snippet/036/name  S32 F_XET
    /-show/showfile/snippet/036/eventtyp  P32 F_XET
    /-show/showfile/snippet/036/channels  P32 F_XET
    /-show/showfile/snippet/036/auxbuses  P32 F_XET
    /-show/showfile/snippet/036/maingrps  P32 F_XET
    /-show/showfile/snippet/036/hasdata  I32 F_XET
/-show/showfile/snippet/037  <SSNP> n=0
    /-show/showfile/snippet/037/name  S32 F_XET
    /-show/showfile/snippet/037/eventtyp  P32 F_XET
    /-show/showfile/snippet/037/channels  P32 F_XET
    /-show/showfile/snippet/037/auxbuses  P32 F_XET
    /-show/showfile/snippet/037/maingrps  P32 F_XET
    /-show/showfile/snippet/037/hasdata  I32 F_XET
/-show/showfile/snippet/038  <SSNP> n=0
    /-show/showfile/snippet/038/name  S32 F_XET
    /-show/showfile/snippet/038/eventtyp  P32 F_XET
    /-show/showfile/snippet/038/channels  P32 F_XET
    /-show/showfile/snippet/038/auxbuses  P32 F_XET
    /-show/showfile/snippet/038/maingrps  P32 F_XET
    /-show/showfile/snippet/038/hasdata  I32 F_XET
/-show/showfile/snippet/039  <SSNP> n=0
    /-show/showfile/snippet/039/name  S32 F_XET
    /-show/showfile/snippet/039/eventtyp  P32 F_XET
    /-show/showfile/snippet/039/channels  P32 F_XET
    /-show/showfile/snippet/039/auxbuses  P32 F_XET
    /-show/showfile/snippet/039/maingrps  P32 F_XET
    /-show/showfile/snippet/039/hasdata  I32 F_XET
/-show/showfile/snippet/040  <SSNP> n=0
    /-show/showfile/snippet/040/name  S32 F_XET
    /-show/showfile/snippet/040/eventtyp  P32 F_XET
    /-show/showfile/snippet/040/channels  P32 F_XET
    /-show/showfile/snippet/040/auxbuses  P32 F_XET
    /-show/showfile/snippet/040/maingrps  P32 F_XET
    /-show/showfile/snippet/040/hasdata  I32 F_XET
/-show/showfile/snippet/041  <SSNP> n=0
    /-show/showfile/snippet/041/name  S32 F_XET
    /-show/showfile/snippet/041/eventtyp  P32 F_XET
    /-show/showfile/snippet/041/channels  P32 F_XET
    /-show/showfile/snippet/041/auxbuses  P32 F_XET
    /-show/showfile/snippet/041/maingrps  P32 F_XET
    /-show/showfile/snippet/041/hasdata  I32 F_XET
/-show/showfile/snippet/042  <SSNP> n=0
    /-show/showfile/snippet/042/name  S32 F_XET
    /-show/showfile/snippet/042/eventtyp  P32 F_XET
    /-show/showfile/snippet/042/channels  P32 F_XET
    /-show/showfile/snippet/042/auxbuses  P32 F_XET
    /-show/showfile/snippet/042/maingrps  P32 F_XET
    /-show/showfile/snippet/042/hasdata  I32 F_XET
/-show/showfile/snippet/043  <SSNP> n=0
    /-show/showfile/snippet/043/name  S32 F_XET
    /-show/showfile/snippet/043/eventtyp  P32 F_XET
    /-show/showfile/snippet/043/channels  P32 F_XET
    /-show/showfile/snippet/043/auxbuses  P32 F_XET
    /-show/showfile/snippet/043/maingrps  P32 F_XET
    /-show/showfile/snippet/043/hasdata  I32 F_XET
/-show/showfile/snippet/044  <SSNP> n=0
    /-show/showfile/snippet/044/name  S32 F_XET
    /-show/showfile/snippet/044/eventtyp  P32 F_XET
    /-show/showfile/snippet/044/channels  P32 F_XET
    /-show/showfile/snippet/044/auxbuses  P32 F_XET
    /-show/showfile/snippet/044/maingrps  P32 F_XET
    /-show/showfile/snippet/044/hasdata  I32 F_XET
/-show/showfile/snippet/045  <SSNP> n=0
    /-show/showfile/snippet/045/name  S32 F_XET
    /-show/showfile/snippet/045/eventtyp  P32 F_XET
    /-show/showfile/snippet/045/channels  P32 F_XET
    /-show/showfile/snippet/045/auxbuses  P32 F_XET
    /-show/showfile/snippet/045/maingrps  P32 F_XET
    /-show/showfile/snippet/045/hasdata  I32 F_XET
/-show/showfile/snippet/046  <SSNP> n=0
    /-show/showfile/snippet/046/name  S32 F_XET
    /-show/showfile/snippet/046/eventtyp  P32 F_XET
    /-show/showfile/snippet/046/channels  P32 F_XET
    /-show/showfile/snippet/046/auxbuses  P32 F_XET
    /-show/showfile/snippet/046/maingrps  P32 F_XET
    /-show/showfile/snippet/046/hasdata  I32 F_XET
/-show/showfile/snippet/047  <SSNP> n=0
    /-show/showfile/snippet/047/name  S32 F_XET
    /-show/showfile/snippet/047/eventtyp  P32 F_XET
    /-show/showfile/snippet/047/channels  P32 F_XET
    /-show/showfile/snippet/047/auxbuses  P32 F_XET
    /-show/showfile/snippet/047/maingrps  P32 F_XET
    /-show/showfile/snippet/047/hasdata  I32 F_XET
/-show/showfile/snippet/048  <SSNP> n=0
    /-show/showfile/snippet/048/name  S32 F_XET
    /-show/showfile/snippet/048/eventtyp  P32 F_XET
    /-show/showfile/snippet/048/channels  P32 F_XET
    /-show/showfile/snippet/048/auxbuses  P32 F_XET
    /-show/showfile/snippet/048/maingrps  P32 F_XET
    /-show/showfile/snippet/048/hasdata  I32 F_XET
/-show/showfile/snippet/049  <SSNP> n=0
    /-show/showfile/snippet/049/name  S32 F_XET
    /-show/showfile/snippet/049/eventtyp  P32 F_XET
    /-show/showfile/snippet/049/channels  P32 F_XET
    /-show/showfile/snippet/049/auxbuses  P32 F_XET
    /-show/showfile/snippet/049/maingrps  P32 F_XET
    /-show/showfile/snippet/049/hasdata  I32 F_XET
/-show/showfile/snippet/050  <SSNP> n=0
    /-show/showfile/snippet/050/name  S32 F_XET
    /-show/showfile/snippet/050/eventtyp  P32 F_XET
    /-show/showfile/snippet/050/channels  P32 F_XET
    /-show/showfile/snippet/050/auxbuses  P32 F_XET
    /-show/showfile/snippet/050/maingrps  P32 F_XET
    /-show/showfile/snippet/050/hasdata  I32 F_XET
/-show/showfile/snippet/051  <SSNP> n=0
    /-show/showfile/snippet/051/name  S32 F_XET
    /-show/showfile/snippet/051/eventtyp  P32 F_XET
    /-show/showfile/snippet/051/channels  P32 F_XET
    /-show/showfile/snippet/051/auxbuses  P32 F_XET
    /-show/showfile/snippet/051/maingrps  P32 F_XET
    /-show/showfile/snippet/051/hasdata  I32 F_XET
/-show/showfile/snippet/052  <SSNP> n=0
    /-show/showfile/snippet/052/name  S32 F_XET
    /-show/showfile/snippet/052/eventtyp  P32 F_XET
    /-show/showfile/snippet/052/channels  P32 F_XET
    /-show/showfile/snippet/052/auxbuses  P32 F_XET
    /-show/showfile/snippet/052/maingrps  P32 F_XET
    /-show/showfile/snippet/052/hasdata  I32 F_XET
/-show/showfile/snippet/053  <SSNP> n=0
    /-show/showfile/snippet/053/name  S32 F_XET
    /-show/showfile/snippet/053/eventtyp  P32 F_XET
    /-show/showfile/snippet/053/channels  P32 F_XET
    /-show/showfile/snippet/053/auxbuses  P32 F_XET
    /-show/showfile/snippet/053/maingrps  P32 F_XET
    /-show/showfile/snippet/053/hasdata  I32 F_XET
/-show/showfile/snippet/054  <SSNP> n=0
    /-show/showfile/snippet/054/name  S32 F_XET
    /-show/showfile/snippet/054/eventtyp  P32 F_XET
    /-show/showfile/snippet/054/channels  P32 F_XET
    /-show/showfile/snippet/054/auxbuses  P32 F_XET
    /-show/showfile/snippet/054/maingrps  P32 F_XET
    /-show/showfile/snippet/054/hasdata  I32 F_XET
/-show/showfile/snippet/055  <SSNP> n=0
    /-show/showfile/snippet/055/name  S32 F_XET
    /-show/showfile/snippet/055/eventtyp  P32 F_XET
    /-show/showfile/snippet/055/channels  P32 F_XET
    /-show/showfile/snippet/055/auxbuses  P32 F_XET
    /-show/showfile/snippet/055/maingrps  P32 F_XET
    /-show/showfile/snippet/055/hasdata  I32 F_XET
/-show/showfile/snippet/056  <SSNP> n=0
    /-show/showfile/snippet/056/name  S32 F_XET
    /-show/showfile/snippet/056/eventtyp  P32 F_XET
    /-show/showfile/snippet/056/channels  P32 F_XET
    /-show/showfile/snippet/056/auxbuses  P32 F_XET
    /-show/showfile/snippet/056/maingrps  P32 F_XET
    /-show/showfile/snippet/056/hasdata  I32 F_XET
/-show/showfile/snippet/057  <SSNP> n=0
    /-show/showfile/snippet/057/name  S32 F_XET
    /-show/showfile/snippet/057/eventtyp  P32 F_XET
    /-show/showfile/snippet/057/channels  P32 F_XET
    /-show/showfile/snippet/057/auxbuses  P32 F_XET
    /-show/showfile/snippet/057/maingrps  P32 F_XET
    /-show/showfile/snippet/057/hasdata  I32 F_XET
/-show/showfile/snippet/058  <SSNP> n=0
    /-show/showfile/snippet/058/name  S32 F_XET
    /-show/showfile/snippet/058/eventtyp  P32 F_XET
    /-show/showfile/snippet/058/channels  P32 F_XET
    /-show/showfile/snippet/058/auxbuses  P32 F_XET
    /-show/showfile/snippet/058/maingrps  P32 F_XET
    /-show/showfile/snippet/058/hasdata  I32 F_XET
/-show/showfile/snippet/059  <SSNP> n=0
    /-show/showfile/snippet/059/name  S32 F_XET
    /-show/showfile/snippet/059/eventtyp  P32 F_XET
    /-show/showfile/snippet/059/channels  P32 F_XET
    /-show/showfile/snippet/059/auxbuses  P32 F_XET
    /-show/showfile/snippet/059/maingrps  P32 F_XET
    /-show/showfile/snippet/059/hasdata  I32 F_XET
/-show/showfile/snippet/060  <SSNP> n=0
    /-show/showfile/snippet/060/name  S32 F_XET
    /-show/showfile/snippet/060/eventtyp  P32 F_XET
    /-show/showfile/snippet/060/channels  P32 F_XET
    /-show/showfile/snippet/060/auxbuses  P32 F_XET
    /-show/showfile/snippet/060/maingrps  P32 F_XET
    /-show/showfile/snippet/060/hasdata  I32 F_XET
/-show/showfile/snippet/061  <SSNP> n=0
    /-show/showfile/snippet/061/name  S32 F_XET
    /-show/showfile/snippet/061/eventtyp  P32 F_XET
    /-show/showfile/snippet/061/channels  P32 F_XET
    /-show/showfile/snippet/061/auxbuses  P32 F_XET
    /-show/showfile/snippet/061/maingrps  P32 F_XET
    /-show/showfile/snippet/061/hasdata  I32 F_XET
/-show/showfile/snippet/062  <SSNP> n=0
    /-show/showfile/snippet/062/name  S32 F_XET
    /-show/showfile/snippet/062/eventtyp  P32 F_XET
    /-show/showfile/snippet/062/channels  P32 F_XET
    /-show/showfile/snippet/062/auxbuses  P32 F_XET
    /-show/showfile/snippet/062/maingrps  P32 F_XET
    /-show/showfile/snippet/062/hasdata  I32 F_XET
/-show/showfile/snippet/063  <SSNP> n=0
    /-show/showfile/snippet/063/name  S32 F_XET
    /-show/showfile/snippet/063/eventtyp  P32 F_XET
    /-show/showfile/snippet/063/channels  P32 F_XET
    /-show/showfile/snippet/063/auxbuses  P32 F_XET
    /-show/showfile/snippet/063/maingrps  P32 F_XET
    /-show/showfile/snippet/063/hasdata  I32 F_XET
/-show/showfile/snippet/064  <SSNP> n=0
    /-show/showfile/snippet/064/name  S32 F_XET
    /-show/showfile/snippet/064/eventtyp  P32 F_XET
    /-show/showfile/snippet/064/channels  P32 F_XET
    /-show/showfile/snippet/064/auxbuses  P32 F_XET
    /-show/showfile/snippet/064/maingrps  P32 F_XET
    /-show/showfile/snippet/064/hasdata  I32 F_XET
/-show/showfile/snippet/065  <SSNP> n=0
    /-show/showfile/snippet/065/name  S32 F_XET
    /-show/showfile/snippet/065/eventtyp  P32 F_XET
    /-show/showfile/snippet/065/channels  P32 F_XET
    /-show/showfile/snippet/065/auxbuses  P32 F_XET
    /-show/showfile/snippet/065/maingrps  P32 F_XET
    /-show/showfile/snippet/065/hasdata  I32 F_XET
/-show/showfile/snippet/066  <SSNP> n=0
    /-show/showfile/snippet/066/name  S32 F_XET
    /-show/showfile/snippet/066/eventtyp  P32 F_XET
    /-show/showfile/snippet/066/channels  P32 F_XET
    /-show/showfile/snippet/066/auxbuses  P32 F_XET
    /-show/showfile/snippet/066/maingrps  P32 F_XET
    /-show/showfile/snippet/066/hasdata  I32 F_XET
/-show/showfile/snippet/067  <SSNP> n=0
    /-show/showfile/snippet/067/name  S32 F_XET
    /-show/showfile/snippet/067/eventtyp  P32 F_XET
    /-show/showfile/snippet/067/channels  P32 F_XET
    /-show/showfile/snippet/067/auxbuses  P32 F_XET
    /-show/showfile/snippet/067/maingrps  P32 F_XET
    /-show/showfile/snippet/067/hasdata  I32 F_XET
/-show/showfile/snippet/068  <SSNP> n=0
    /-show/showfile/snippet/068/name  S32 F_XET
    /-show/showfile/snippet/068/eventtyp  P32 F_XET
    /-show/showfile/snippet/068/channels  P32 F_XET
    /-show/showfile/snippet/068/auxbuses  P32 F_XET
    /-show/showfile/snippet/068/maingrps  P32 F_XET
    /-show/showfile/snippet/068/hasdata  I32 F_XET
/-show/showfile/snippet/069  <SSNP> n=0
    /-show/showfile/snippet/069/name  S32 F_XET
    /-show/showfile/snippet/069/eventtyp  P32 F_XET
    /-show/showfile/snippet/069/channels  P32 F_XET
    /-show/showfile/snippet/069/auxbuses  P32 F_XET
    /-show/showfile/snippet/069/maingrps  P32 F_XET
    /-show/showfile/snippet/069/hasdata  I32 F_XET
/-show/showfile/snippet/070  <SSNP> n=0
    /-show/showfile/snippet/070/name  S32 F_XET
    /-show/showfile/snippet/070/eventtyp  P32 F_XET
    /-show/showfile/snippet/070/channels  P32 F_XET
    /-show/showfile/snippet/070/auxbuses  P32 F_XET
    /-show/showfile/snippet/070/maingrps  P32 F_XET
    /-show/showfile/snippet/070/hasdata  I32 F_XET
/-show/showfile/snippet/071  <SSNP> n=0
    /-show/showfile/snippet/071/name  S32 F_XET
    /-show/showfile/snippet/071/eventtyp  P32 F_XET
    /-show/showfile/snippet/071/channels  P32 F_XET
    /-show/showfile/snippet/071/auxbuses  P32 F_XET
    /-show/showfile/snippet/071/maingrps  P32 F_XET
    /-show/showfile/snippet/071/hasdata  I32 F_XET
/-show/showfile/snippet/072  <SSNP> n=0
    /-show/showfile/snippet/072/name  S32 F_XET
    /-show/showfile/snippet/072/eventtyp  P32 F_XET
    /-show/showfile/snippet/072/channels  P32 F_XET
    /-show/showfile/snippet/072/auxbuses  P32 F_XET
    /-show/showfile/snippet/072/maingrps  P32 F_XET
    /-show/showfile/snippet/072/hasdata  I32 F_XET
/-show/showfile/snippet/073  <SSNP> n=0
    /-show/showfile/snippet/073/name  S32 F_XET
    /-show/showfile/snippet/073/eventtyp  P32 F_XET
    /-show/showfile/snippet/073/channels  P32 F_XET
    /-show/showfile/snippet/073/auxbuses  P32 F_XET
    /-show/showfile/snippet/073/maingrps  P32 F_XET
    /-show/showfile/snippet/073/hasdata  I32 F_XET
/-show/showfile/snippet/074  <SSNP> n=0
    /-show/showfile/snippet/074/name  S32 F_XET
    /-show/showfile/snippet/074/eventtyp  P32 F_XET
    /-show/showfile/snippet/074/channels  P32 F_XET
    /-show/showfile/snippet/074/auxbuses  P32 F_XET
    /-show/showfile/snippet/074/maingrps  P32 F_XET
    /-show/showfile/snippet/074/hasdata  I32 F_XET
/-show/showfile/snippet/075  <SSNP> n=0
    /-show/showfile/snippet/075/name  S32 F_XET
    /-show/showfile/snippet/075/eventtyp  P32 F_XET
    /-show/showfile/snippet/075/channels  P32 F_XET
    /-show/showfile/snippet/075/auxbuses  P32 F_XET
    /-show/showfile/snippet/075/maingrps  P32 F_XET
    /-show/showfile/snippet/075/hasdata  I32 F_XET
/-show/showfile/snippet/076  <SSNP> n=0
    /-show/showfile/snippet/076/name  S32 F_XET
    /-show/showfile/snippet/076/eventtyp  P32 F_XET
    /-show/showfile/snippet/076/channels  P32 F_XET
    /-show/showfile/snippet/076/auxbuses  P32 F_XET
    /-show/showfile/snippet/076/maingrps  P32 F_XET
    /-show/showfile/snippet/076/hasdata  I32 F_XET
/-show/showfile/snippet/077  <SSNP> n=0
    /-show/showfile/snippet/077/name  S32 F_XET
    /-show/showfile/snippet/077/eventtyp  P32 F_XET
    /-show/showfile/snippet/077/channels  P32 F_XET
    /-show/showfile/snippet/077/auxbuses  P32 F_XET
    /-show/showfile/snippet/077/maingrps  P32 F_XET
    /-show/showfile/snippet/077/hasdata  I32 F_XET
/-show/showfile/snippet/078  <SSNP> n=0
    /-show/showfile/snippet/078/name  S32 F_XET
    /-show/showfile/snippet/078/eventtyp  P32 F_XET
    /-show/showfile/snippet/078/channels  P32 F_XET
    /-show/showfile/snippet/078/auxbuses  P32 F_XET
    /-show/showfile/snippet/078/maingrps  P32 F_XET
    /-show/showfile/snippet/078/hasdata  I32 F_XET
/-show/showfile/snippet/079  <SSNP> n=0
    /-show/showfile/snippet/079/name  S32 F_XET
    /-show/showfile/snippet/079/eventtyp  P32 F_XET
    /-show/showfile/snippet/079/channels  P32 F_XET
    /-show/showfile/snippet/079/auxbuses  P32 F_XET
    /-show/showfile/snippet/079/maingrps  P32 F_XET
    /-show/showfile/snippet/079/hasdata  I32 F_XET
/-show/showfile/snippet/080  <SSNP> n=0
    /-show/showfile/snippet/080/name  S32 F_XET
    /-show/showfile/snippet/080/eventtyp  P32 F_XET
    /-show/showfile/snippet/080/channels  P32 F_XET
    /-show/showfile/snippet/080/auxbuses  P32 F_XET
    /-show/showfile/snippet/080/maingrps  P32 F_XET
    /-show/showfile/snippet/080/hasdata  I32 F_XET
/-show/showfile/snippet/081  <SSNP> n=0
    /-show/showfile/snippet/081/name  S32 F_XET
    /-show/showfile/snippet/081/eventtyp  P32 F_XET
    /-show/showfile/snippet/081/channels  P32 F_XET
    /-show/showfile/snippet/081/auxbuses  P32 F_XET
    /-show/showfile/snippet/081/maingrps  P32 F_XET
    /-show/showfile/snippet/081/hasdata  I32 F_XET
/-show/showfile/snippet/082  <SSNP> n=0
    /-show/showfile/snippet/082/name  S32 F_XET
    /-show/showfile/snippet/082/eventtyp  P32 F_XET
    /-show/showfile/snippet/082/channels  P32 F_XET
    /-show/showfile/snippet/082/auxbuses  P32 F_XET
    /-show/showfile/snippet/082/maingrps  P32 F_XET
    /-show/showfile/snippet/082/hasdata  I32 F_XET
/-show/showfile/snippet/083  <SSNP> n=0
    /-show/showfile/snippet/083/name  S32 F_XET
    /-show/showfile/snippet/083/eventtyp  P32 F_XET
    /-show/showfile/snippet/083/channels  P32 F_XET
    /-show/showfile/snippet/083/auxbuses  P32 F_XET
    /-show/showfile/snippet/083/maingrps  P32 F_XET
    /-show/showfile/snippet/083/hasdata  I32 F_XET
/-show/showfile/snippet/084  <SSNP> n=0
    /-show/showfile/snippet/084/name  S32 F_XET
    /-show/showfile/snippet/084/eventtyp  P32 F_XET
    /-show/showfile/snippet/084/channels  P32 F_XET
    /-show/showfile/snippet/084/auxbuses  P32 F_XET
    /-show/showfile/snippet/084/maingrps  P32 F_XET
    /-show/showfile/snippet/084/hasdata  I32 F_XET
/-show/showfile/snippet/085  <SSNP> n=0
    /-show/showfile/snippet/085/name  S32 F_XET
    /-show/showfile/snippet/085/eventtyp  P32 F_XET
    /-show/showfile/snippet/085/channels  P32 F_XET
    /-show/showfile/snippet/085/auxbuses  P32 F_XET
    /-show/showfile/snippet/085/maingrps  P32 F_XET
    /-show/showfile/snippet/085/hasdata  I32 F_XET
/-show/showfile/snippet/086  <SSNP> n=0
    /-show/showfile/snippet/086/name  S32 F_XET
    /-show/showfile/snippet/086/eventtyp  P32 F_XET
    /-show/showfile/snippet/086/channels  P32 F_XET
    /-show/showfile/snippet/086/auxbuses  P32 F_XET
    /-show/showfile/snippet/086/maingrps  P32 F_XET
    /-show/showfile/snippet/086/hasdata  I32 F_XET
/-show/showfile/snippet/087  <SSNP> n=0
    /-show/showfile/snippet/087/name  S32 F_XET
    /-show/showfile/snippet/087/eventtyp  P32 F_XET
    /-show/showfile/snippet/087/channels  P32 F_XET
    /-show/showfile/snippet/087/auxbuses  P32 F_XET
    /-show/showfile/snippet/087/maingrps  P32 F_XET
    /-show/showfile/snippet/087/hasdata  I32 F_XET
/-show/showfile/snippet/088  <SSNP> n=0
    /-show/showfile/snippet/088/name  S32 F_XET
    /-show/showfile/snippet/088/eventtyp  P32 F_XET
    /-show/showfile/snippet/088/channels  P32 F_XET
    /-show/showfile/snippet/088/auxbuses  P32 F_XET
    /-show/showfile/snippet/088/maingrps  P32 F_XET
    /-show/showfile/snippet/088/hasdata  I32 F_XET
/-show/showfile/snippet/089  <SSNP> n=0
    /-show/showfile/snippet/089/name  S32 F_XET
    /-show/showfile/snippet/089/eventtyp  P32 F_XET
    /-show/showfile/snippet/089/channels  P32 F_XET
    /-show/showfile/snippet/089/auxbuses  P32 F_XET
    /-show/showfile/snippet/089/maingrps  P32 F_XET
    /-show/showfile/snippet/089/hasdata  I32 F_XET
/-show/showfile/snippet/090  <SSNP> n=0
    /-show/showfile/snippet/090/name  S32 F_XET
    /-show/showfile/snippet/090/eventtyp  P32 F_XET
    /-show/showfile/snippet/090/channels  P32 F_XET
    /-show/showfile/snippet/090/auxbuses  P32 F_XET
    /-show/showfile/snippet/090/maingrps  P32 F_XET
    /-show/showfile/snippet/090/hasdata  I32 F_XET
/-show/showfile/snippet/091  <SSNP> n=0
    /-show/showfile/snippet/091/name  S32 F_XET
    /-show/showfile/snippet/091/eventtyp  P32 F_XET
    /-show/showfile/snippet/091/channels  P32 F_XET
    /-show/showfile/snippet/091/auxbuses  P32 F_XET
    /-show/showfile/snippet/091/maingrps  P32 F_XET
    /-show/showfile/snippet/091/hasdata  I32 F_XET
/-show/showfile/snippet/092  <SSNP> n=0
    /-show/showfile/snippet/092/name  S32 F_XET
    /-show/showfile/snippet/092/eventtyp  P32 F_XET
    /-show/showfile/snippet/092/channels  P32 F_XET
    /-show/showfile/snippet/092/auxbuses  P32 F_XET
    /-show/showfile/snippet/092/maingrps  P32 F_XET
    /-show/showfile/snippet/092/hasdata  I32 F_XET
/-show/showfile/snippet/093  <SSNP> n=0
    /-show/showfile/snippet/093/name  S32 F_XET
    /-show/showfile/snippet/093/eventtyp  P32 F_XET
    /-show/showfile/snippet/093/channels  P32 F_XET
    /-show/showfile/snippet/093/auxbuses  P32 F_XET
    /-show/showfile/snippet/093/maingrps  P32 F_XET
    /-show/showfile/snippet/093/hasdata  I32 F_XET
/-show/showfile/snippet/094  <SSNP> n=0
    /-show/showfile/snippet/094/name  S32 F_XET
    /-show/showfile/snippet/094/eventtyp  P32 F_XET
    /-show/showfile/snippet/094/channels  P32 F_XET
    /-show/showfile/snippet/094/auxbuses  P32 F_XET
    /-show/showfile/snippet/094/maingrps  P32 F_XET
    /-show/showfile/snippet/094/hasdata  I32 F_XET
/-show/showfile/snippet/095  <SSNP> n=0
    /-show/showfile/snippet/095/name  S32 F_XET
    /-show/showfile/snippet/095/eventtyp  P32 F_XET
    /-show/showfile/snippet/095/channels  P32 F_XET
    /-show/showfile/snippet/095/auxbuses  P32 F_XET
    /-show/showfile/snippet/095/maingrps  P32 F_XET
    /-show/showfile/snippet/095/hasdata  I32 F_XET
/-show/showfile/snippet/096  <SSNP> n=0
    /-show/showfile/snippet/096/name  S32 F_XET
    /-show/showfile/snippet/096/eventtyp  P32 F_XET
    /-show/showfile/snippet/096/channels  P32 F_XET
    /-show/showfile/snippet/096/auxbuses  P32 F_XET
    /-show/showfile/snippet/096/maingrps  P32 F_XET
    /-show/showfile/snippet/096/hasdata  I32 F_XET
/-show/showfile/snippet/097  <SSNP> n=0
    /-show/showfile/snippet/097/name  S32 F_XET
    /-show/showfile/snippet/097/eventtyp  P32 F_XET
    /-show/showfile/snippet/097/channels  P32 F_XET
    /-show/showfile/snippet/097/auxbuses  P32 F_XET
    /-show/showfile/snippet/097/maingrps  P32 F_XET
    /-show/showfile/snippet/097/hasdata  I32 F_XET
/-show/showfile/snippet/098  <SSNP> n=0
    /-show/showfile/snippet/098/name  S32 F_XET
    /-show/showfile/snippet/098/eventtyp  P32 F_XET
    /-show/showfile/snippet/098/channels  P32 F_XET
    /-show/showfile/snippet/098/auxbuses  P32 F_XET
    /-show/showfile/snippet/098/maingrps  P32 F_XET
    /-show/showfile/snippet/098/hasdata  I32 F_XET
/-show/showfile/snippet/099  <SSNP> n=0
    /-show/showfile/snippet/099/name  S32 F_XET
    /-show/showfile/snippet/099/eventtyp  P32 F_XET
    /-show/showfile/snippet/099/channels  P32 F_XET
    /-show/showfile/snippet/099/auxbuses  P32 F_XET
    /-show/showfile/snippet/099/maingrps  P32 F_XET
    /-show/showfile/snippet/099/hasdata  I32 F_XET
```

### Xmisc (X32Misc.h, 311 entries)

```
/-usb  <USB> n=0
    /-usb/path  S32 F_XET
    /-usb/title  S32 F_XET
    /-usb/maxpos  I32 F_XET
    /-usb/dirpos  I32 F_XET
    /-usb/dir/000/name  S32 F_XET
    /-usb/dir/001/name  S32 F_XET
    /-usb/dir/002/name  S32 F_XET
    /-usb/dir/003/name  S32 F_XET
    /-usb/dir/004/name  S32 F_XET
    /-usb/dir/005/name  S32 F_XET
    /-usb/dir/006/name  S32 F_XET
    /-usb/dir/007/name  S32 F_XET
    /-usb/dir/008/name  S32 F_XET
    /-usb/dir/009/name  S32 F_XET
    /-usb/dir/010/name  S32 F_XET
    /-usb/dir/011/name  S32 F_XET
    /-usb/dir/012/name  S32 F_XET
    /-usb/dir/013/name  S32 F_XET
    /-usb/dir/014/name  S32 F_XET
    /-usb/dir/015/name  S32 F_XET
    /-usb/dir/016/name  S32 F_XET
    /-usb/dir/017/name  S32 F_XET
    /-usb/dir/018/name  S32 F_XET
    /-usb/dir/019/name  S32 F_XET
    /-usb/dir/020/name  S32 F_XET
    /-usb/dir/021/name  S32 F_XET
    /-usb/dir/022/name  S32 F_XET
    /-usb/dir/023/name  S32 F_XET
    /-usb/dir/024/name  S32 F_XET
    /-usb/dir/025/name  S32 F_XET
    /-usb/dir/026/name  S32 F_XET
    /-usb/dir/027/name  S32 F_XET
    /-usb/dir/028/name  S32 F_XET
    /-usb/dir/029/name  S32 F_XET
    /-usb/dir/030/name  S32 F_XET
    /-usb/dir/031/name  S32 F_XET
    /-usb/dir/032/name  S32 F_XET
    /-usb/dir/033/name  S32 F_XET
    /-usb/dir/034/name  S32 F_XET
    /-usb/dir/035/name  S32 F_XET
    /-usb/dir/036/name  S32 F_XET
    /-usb/dir/037/name  S32 F_XET
    /-usb/dir/038/name  S32 F_XET
    /-usb/dir/039/name  S32 F_XET
    /-usb/dir/040/name  S32 F_XET
    /-usb/dir/041/name  S32 F_XET
    /-usb/dir/042/name  S32 F_XET
    /-usb/dir/043/name  S32 F_XET
    /-usb/dir/044/name  S32 F_XET
    /-usb/dir/045/name  S32 F_XET
    /-usb/dir/046/name  S32 F_XET
    /-usb/dir/047/name  S32 F_XET
    /-usb/dir/048/name  S32 F_XET
    /-usb/dir/049/name  S32 F_XET
    /-usb/dir/050/name  S32 F_XET
    /-usb/dir/051/name  S32 F_XET
    /-usb/dir/052/name  S32 F_XET
    /-usb/dir/053/name  S32 F_XET
    /-usb/dir/054/name  S32 F_XET
    /-usb/dir/055/name  S32 F_XET
    /-usb/dir/056/name  S32 F_XET
    /-usb/dir/057/name  S32 F_XET
    /-usb/dir/058/name  S32 F_XET
    /-usb/dir/059/name  S32 F_XET
    /-usb/dir/060/name  S32 F_XET
    /-usb/dir/061/name  S32 F_XET
    /-usb/dir/062/name  S32 F_XET
    /-usb/dir/063/name  S32 F_XET
    /-usb/dir/064/name  S32 F_XET
    /-usb/dir/065/name  S32 F_XET
    /-usb/dir/066/name  S32 F_XET
    /-usb/dir/067/name  S32 F_XET
    /-usb/dir/068/name  S32 F_XET
    /-usb/dir/069/name  S32 F_XET
    /-usb/dir/070/name  S32 F_XET
    /-usb/dir/071/name  S32 F_XET
    /-usb/dir/072/name  S32 F_XET
    /-usb/dir/073/name  S32 F_XET
    /-usb/dir/074/name  S32 F_XET
    /-usb/dir/075/name  S32 F_XET
    /-usb/dir/076/name  S32 F_XET
    /-usb/dir/077/name  S32 F_XET
    /-usb/dir/078/name  S32 F_XET
    /-usb/dir/079/name  S32 F_XET
    /-usb/dir/080/name  S32 F_XET
    /-usb/dir/081/name  S32 F_XET
    /-usb/dir/082/name  S32 F_XET
    /-usb/dir/083/name  S32 F_XET
    /-usb/dir/084/name  S32 F_XET
    /-usb/dir/085/name  S32 F_XET
    /-usb/dir/086/name  S32 F_XET
    /-usb/dir/087/name  S32 F_XET
    /-usb/dir/088/name  S32 F_XET
    /-usb/dir/089/name  S32 F_XET
    /-usb/dir/090/name  S32 F_XET
    /-usb/dir/091/name  S32 F_XET
    /-usb/dir/092/name  S32 F_XET
    /-usb/dir/093/name  S32 F_XET
    /-usb/dir/094/name  S32 F_XET
    /-usb/dir/095/name  S32 F_XET
    /-usb/dir/096/name  S32 F_XET
    /-usb/dir/097/name  S32 F_XET
    /-usb/dir/098/name  S32 F_XET
    /-usb/dir/099/name  S32 F_XET
    /-usb/dir/100/name  S32 F_XET
    /-usb/dir/101/name  S32 F_XET
    /-usb/dir/102/name  S32 F_XET
    /-usb/dir/103/name  S32 F_XET
    /-usb/dir/104/name  S32 F_XET
    /-usb/dir/105/name  S32 F_XET
    /-usb/dir/106/name  S32 F_XET
    /-usb/dir/107/name  S32 F_XET
    /-usb/dir/108/name  S32 F_XET
    /-usb/dir/109/name  S32 F_XET
    /-usb/dir/110/name  S32 F_XET
    /-usb/dir/111/name  S32 F_XET
    /-usb/dir/112/name  S32 F_XET
    /-usb/dir/113/name  S32 F_XET
    /-usb/dir/114/name  S32 F_XET
    /-usb/dir/115/name  S32 F_XET
    /-usb/dir/116/name  S32 F_XET
    /-usb/dir/117/name  S32 F_XET
    /-usb/dir/118/name  S32 F_XET
    /-usb/dir/119/name  S32 F_XET
    /-usb/dir/120/name  S32 F_XET
    /-usb/dir/121/name  S32 F_XET
    /-usb/dir/122/name  S32 F_XET
    /-usb/dir/123/name  S32 F_XET
    /-usb/dir/124/name  S32 F_XET
    /-usb/dir/125/name  S32 F_XET
    /-usb/dir/126/name  S32 F_XET
    /-usb/dir/127/name  S32 F_XET
    /-usb/dir/128/name  S32 F_XET
    /-usb/dir/129/name  S32 F_XET
    /-usb/dir/130/name  S32 F_XET
    /-usb/dir/131/name  S32 F_XET
    /-usb/dir/132/name  S32 F_XET
    /-usb/dir/133/name  S32 F_XET
    /-usb/dir/134/name  S32 F_XET
    /-usb/dir/135/name  S32 F_XET
    /-usb/dir/136/name  S32 F_XET
    /-usb/dir/137/name  S32 F_XET
    /-usb/dir/138/name  S32 F_XET
    /-usb/dir/139/name  S32 F_XET
    /-usb/dir/140/name  S32 F_XET
    /-usb/dir/141/name  S32 F_XET
    /-usb/dir/142/name  S32 F_XET
    /-usb/dir/143/name  S32 F_XET
    /-usb/dir/144/name  S32 F_XET
    /-usb/dir/145/name  S32 F_XET
    /-usb/dir/146/name  S32 F_XET
    /-usb/dir/147/name  S32 F_XET
    /-usb/dir/148/name  S32 F_XET
    /-usb/dir/149/name  S32 F_XET
    /-usb/dir/150/name  S32 F_XET
    /-usb/dir/151/name  S32 F_XET
    /-usb/dir/152/name  S32 F_XET
    /-usb/dir/153/name  S32 F_XET
    /-usb/dir/154/name  S32 F_XET
    /-usb/dir/155/name  S32 F_XET
    /-usb/dir/156/name  S32 F_XET
    /-usb/dir/157/name  S32 F_XET
    /-usb/dir/158/name  S32 F_XET
    /-usb/dir/159/name  S32 F_XET
    /-usb/dir/160/name  S32 F_XET
    /-usb/dir/161/name  S32 F_XET
    /-usb/dir/162/name  S32 F_XET
    /-usb/dir/163/name  S32 F_XET
    /-usb/dir/164/name  S32 F_XET
    /-usb/dir/165/name  S32 F_XET
    /-usb/dir/166/name  S32 F_XET
    /-usb/dir/167/name  S32 F_XET
    /-usb/dir/168/name  S32 F_XET
    /-usb/dir/169/name  S32 F_XET
    /-usb/dir/170/name  S32 F_XET
    /-usb/dir/171/name  S32 F_XET
    /-usb/dir/172/name  S32 F_XET
    /-usb/dir/173/name  S32 F_XET
    /-usb/dir/174/name  S32 F_XET
    /-usb/dir/175/name  S32 F_XET
    /-usb/dir/176/name  S32 F_XET
    /-usb/dir/177/name  S32 F_XET
    /-usb/dir/178/name  S32 F_XET
    /-usb/dir/179/name  S32 F_XET
    /-usb/dir/180/name  S32 F_XET
    /-usb/dir/181/name  S32 F_XET
    /-usb/dir/182/name  S32 F_XET
    /-usb/dir/183/name  S32 F_XET
    /-usb/dir/184/name  S32 F_XET
    /-usb/dir/185/name  S32 F_XET
    /-usb/dir/186/name  S32 F_XET
    /-usb/dir/187/name  S32 F_XET
    /-usb/dir/188/name  S32 F_XET
    /-usb/dir/189/name  S32 F_XET
    /-usb/dir/190/name  S32 F_XET
    /-usb/dir/191/name  S32 F_XET
    /-usb/dir/192/name  S32 F_XET
    /-usb/dir/193/name  S32 F_XET
    /-usb/dir/194/name  S32 F_XET
    /-usb/dir/195/name  S32 F_XET
    /-usb/dir/196/name  S32 F_XET
    /-usb/dir/197/name  S32 F_XET
    /-usb/dir/198/name  S32 F_XET
    /-usb/dir/199/name  S32 F_XET
    /-usb/dir/200/name  S32 F_XET
    /-usb/dir/201/name  S32 F_XET
    /-usb/dir/202/name  S32 F_XET
    /-usb/dir/203/name  S32 F_XET
    /-usb/dir/204/name  S32 F_XET
    /-usb/dir/205/name  S32 F_XET
    /-usb/dir/206/name  S32 F_XET
    /-usb/dir/207/name  S32 F_XET
    /-usb/dir/208/name  S32 F_XET
    /-usb/dir/209/name  S32 F_XET
    /-usb/dir/210/name  S32 F_XET
    /-usb/dir/211/name  S32 F_XET
    /-usb/dir/212/name  S32 F_XET
    /-usb/dir/213/name  S32 F_XET
    /-usb/dir/214/name  S32 F_XET
    /-usb/dir/215/name  S32 F_XET
    /-usb/dir/216/name  S32 F_XET
    /-usb/dir/217/name  S32 F_XET
    /-usb/dir/218/name  S32 F_XET
    /-usb/dir/219/name  S32 F_XET
    /-usb/dir/220/name  S32 F_XET
    /-usb/dir/221/name  S32 F_XET
    /-usb/dir/222/name  S32 F_XET
    /-usb/dir/223/name  S32 F_XET
    /-usb/dir/224/name  S32 F_XET
    /-usb/dir/225/name  S32 F_XET
    /-usb/dir/226/name  S32 F_XET
    /-usb/dir/227/name  S32 F_XET
    /-usb/dir/228/name  S32 F_XET
    /-usb/dir/229/name  S32 F_XET
    /-usb/dir/230/name  S32 F_XET
    /-usb/dir/231/name  S32 F_XET
    /-usb/dir/232/name  S32 F_XET
    /-usb/dir/233/name  S32 F_XET
    /-usb/dir/234/name  S32 F_XET
    /-usb/dir/235/name  S32 F_XET
    /-usb/dir/236/name  S32 F_XET
    /-usb/dir/237/name  S32 F_XET
    /-usb/dir/238/name  S32 F_XET
    /-usb/dir/239/name  S32 F_XET
    /-usb/dir/240/name  S32 F_XET
    /-usb/dir/241/name  S32 F_XET
    /-usb/dir/242/name  S32 F_XET
    /-usb/dir/243/name  S32 F_XET
    /-usb/dir/244/name  S32 F_XET
    /-usb/dir/245/name  S32 F_XET
    /-usb/dir/246/name  S32 F_XET
    /-usb/dir/247/name  S32 F_XET
    /-usb/dir/248/name  S32 F_XET
    /-usb/dir/249/name  S32 F_XET
    /-usb/dir/250/name  S32 F_XET
    /-usb/dir/251/name  S32 F_XET
    /-usb/dir/252/name  S32 F_XET
    /-usb/dir/253/name  S32 F_XET
    /-usb/dir/254/name  S32 F_XET
    /-usb/dir/255/name  S32 F_XET
    /-usb/dir/256/name  S32 F_XET
    /undo/time  S32 F_XET
    /insert/aux/1  I32 F_XET
    /insert/aux/2  I32 F_XET
    /insert/aux/3  I32 F_XET
    /insert/aux/4  I32 F_XET
    /insert/aux/5  I32 F_XET
    /insert/aux/6  I32 F_XET
/-ha  <HA> n=0
/-ha/00  <HA> n=0
    /-ha/00/index  I32 F_XET
    /-ha/01/index  I32 F_XET
    /-ha/02/index  I32 F_XET
    /-ha/03/index  I32 F_XET
    /-ha/04/index  I32 F_XET
    /-ha/05/index  I32 F_XET
    /-ha/06/index  I32 F_XET
    /-ha/07/index  I32 F_XET
    /-ha/08/index  I32 F_XET
    /-ha/09/index  I32 F_XET
    /-ha/10/index  I32 F_XET
    /-ha/11/index  I32 F_XET
    /-ha/12/index  I32 F_XET
    /-ha/13/index  I32 F_XET
    /-ha/14/index  I32 F_XET
    /-ha/15/index  I32 F_XET
    /-ha/16/index  I32 F_XET
    /-ha/17/index  I32 F_XET
    /-ha/18/index  I32 F_XET
    /-ha/19/index  I32 F_XET
    /-ha/20/index  I32 F_XET
    /-ha/21/index  I32 F_XET
    /-ha/22/index  I32 F_XET
    /-ha/23/index  I32 F_XET
    /-ha/24/index  I32 F_XET
    /-ha/25/index  I32 F_XET
    /-ha/26/index  I32 F_XET
    /-ha/27/index  I32 F_XET
    /-ha/28/index  I32 F_XET
    /-ha/29/index  I32 F_XET
    /-ha/30/index  I32 F_XET
    /-ha/31/index  I32 F_XET
    /-ha/32/index  I32 F_XET
    /-ha/33/index  I32 F_XET
    /-ha/34/index  I32 F_XET
    /-ha/35/index  I32 F_XET
    /-ha/36/index  I32 F_XET
    /-ha/37/index  I32 F_XET
    /-ha/38/index  I32 F_XET
    /-ha/39/index  I32 F_XET
```

### Xlibsc (X32Libs.h, 602 entries)

```
/-libs  <SLIBS> n=0
/-libs/ch  <SLIBS> n=0
/-libs/ch/001  <SLIBS> n=0
    /-libs/ch/001/pos  I32 F_XET
    /-libs/ch/001/name  S32 F_XET
    /-libs/ch/001/type  I32 F_XET
    /-libs/ch/001/flags  P32 F_XET
    /-libs/ch/001/hasdata  I32 F_XET
/-libs/ch/002  <SLIBS> n=0
    /-libs/ch/002/pos  I32 F_XET
    /-libs/ch/002/name  S32 F_XET
    /-libs/ch/002/type  I32 F_XET
    /-libs/ch/002/flags  P32 F_XET
    /-libs/ch/002/hasdata  I32 F_XET
/-libs/ch/003  <SLIBS> n=0
    /-libs/ch/003/pos  I32 F_XET
    /-libs/ch/003/name  S32 F_XET
    /-libs/ch/003/type  I32 F_XET
    /-libs/ch/003/flags  P32 F_XET
    /-libs/ch/003/hasdata  I32 F_XET
/-libs/ch/004  <SLIBS> n=0
    /-libs/ch/004/pos  I32 F_XET
    /-libs/ch/004/name  S32 F_XET
    /-libs/ch/004/type  I32 F_XET
    /-libs/ch/004/flags  P32 F_XET
    /-libs/ch/004/hasdata  I32 F_XET
/-libs/ch/005  <SLIBS> n=0
    /-libs/ch/005/pos  I32 F_XET
    /-libs/ch/005/name  S32 F_XET
    /-libs/ch/005/type  I32 F_XET
    /-libs/ch/005/flags  P32 F_XET
    /-libs/ch/005/hasdata  I32 F_XET
/-libs/ch/006  <SLIBS> n=0
    /-libs/ch/006/pos  I32 F_XET
    /-libs/ch/006/name  S32 F_XET
    /-libs/ch/006/type  I32 F_XET
    /-libs/ch/006/flags  P32 F_XET
    /-libs/ch/006/hasdata  I32 F_XET
/-libs/ch/007  <SLIBS> n=0
    /-libs/ch/007/pos  I32 F_XET
    /-libs/ch/007/name  S32 F_XET
    /-libs/ch/007/type  I32 F_XET
    /-libs/ch/007/flags  P32 F_XET
    /-libs/ch/007/hasdata  I32 F_XET
/-libs/ch/008  <SLIBS> n=0
    /-libs/ch/008/pos  I32 F_XET
    /-libs/ch/008/name  S32 F_XET
    /-libs/ch/008/type  I32 F_XET
    /-libs/ch/008/flags  P32 F_XET
    /-libs/ch/008/hasdata  I32 F_XET
/-libs/ch/009  <SLIBS> n=0
    /-libs/ch/009/pos  I32 F_XET
    /-libs/ch/009/name  S32 F_XET
    /-libs/ch/009/type  I32 F_XET
    /-libs/ch/009/flags  P32 F_XET
    /-libs/ch/009/hasdata  I32 F_XET
/-libs/ch/010  <SLIBS> n=0
    /-libs/ch/010/pos  I32 F_XET
    /-libs/ch/010/name  S32 F_XET
    /-libs/ch/010/type  I32 F_XET
    /-libs/ch/010/flags  P32 F_XET
    /-libs/ch/010/hasdata  I32 F_XET
/-libs/ch/011  <SLIBS> n=0
    /-libs/ch/011/pos  I32 F_XET
    /-libs/ch/011/name  S32 F_XET
    /-libs/ch/011/type  I32 F_XET
    /-libs/ch/011/flags  P32 F_XET
    /-libs/ch/011/hasdata  I32 F_XET
/-libs/ch/012  <SLIBS> n=0
    /-libs/ch/012/pos  I32 F_XET
    /-libs/ch/012/name  S32 F_XET
    /-libs/ch/012/type  I32 F_XET
    /-libs/ch/012/flags  P32 F_XET
    /-libs/ch/012/hasdata  I32 F_XET
/-libs/ch/013  <SLIBS> n=0
    /-libs/ch/013/pos  I32 F_XET
    /-libs/ch/013/name  S32 F_XET
    /-libs/ch/013/type  I32 F_XET
    /-libs/ch/013/flags  P32 F_XET
    /-libs/ch/013/hasdata  I32 F_XET
/-libs/ch/014  <SLIBS> n=0
    /-libs/ch/014/pos  I32 F_XET
    /-libs/ch/014/name  S32 F_XET
    /-libs/ch/014/type  I32 F_XET
    /-libs/ch/014/flags  P32 F_XET
    /-libs/ch/014/hasdata  I32 F_XET
/-libs/ch/015  <SLIBS> n=0
    /-libs/ch/015/pos  I32 F_XET
    /-libs/ch/015/name  S32 F_XET
    /-libs/ch/015/type  I32 F_XET
    /-libs/ch/015/flags  P32 F_XET
    /-libs/ch/015/hasdata  I32 F_XET
/-libs/ch/016  <SLIBS> n=0
    /-libs/ch/016/pos  I32 F_XET
    /-libs/ch/016/name  S32 F_XET
    /-libs/ch/016/type  I32 F_XET
    /-libs/ch/016/flags  P32 F_XET
    /-libs/ch/016/hasdata  I32 F_XET
/-libs/ch/017  <SLIBS> n=0
    /-libs/ch/017/pos  I32 F_XET
    /-libs/ch/017/name  S32 F_XET
    /-libs/ch/017/type  I32 F_XET
    /-libs/ch/017/flags  P32 F_XET
    /-libs/ch/017/hasdata  I32 F_XET
/-libs/ch/018  <SLIBS> n=0
    /-libs/ch/018/pos  I32 F_XET
    /-libs/ch/018/name  S32 F_XET
    /-libs/ch/018/type  I32 F_XET
    /-libs/ch/018/flags  P32 F_XET
    /-libs/ch/018/hasdata  I32 F_XET
/-libs/ch/019  <SLIBS> n=0
    /-libs/ch/019/pos  I32 F_XET
    /-libs/ch/019/name  S32 F_XET
    /-libs/ch/019/type  I32 F_XET
    /-libs/ch/019/flags  P32 F_XET
    /-libs/ch/019/hasdata  I32 F_XET
/-libs/ch/020  <SLIBS> n=0
    /-libs/ch/020/pos  I32 F_XET
    /-libs/ch/020/name  S32 F_XET
    /-libs/ch/020/type  I32 F_XET
    /-libs/ch/020/flags  P32 F_XET
    /-libs/ch/020/hasdata  I32 F_XET
/-libs/ch/021  <SLIBS> n=0
    /-libs/ch/021/pos  I32 F_XET
    /-libs/ch/021/name  S32 F_XET
    /-libs/ch/021/type  I32 F_XET
    /-libs/ch/021/flags  P32 F_XET
    /-libs/ch/021/hasdata  I32 F_XET
/-libs/ch/022  <SLIBS> n=0
    /-libs/ch/022/pos  I32 F_XET
    /-libs/ch/022/name  S32 F_XET
    /-libs/ch/022/type  I32 F_XET
    /-libs/ch/022/flags  P32 F_XET
    /-libs/ch/022/hasdata  I32 F_XET
/-libs/ch/023  <SLIBS> n=0
    /-libs/ch/023/pos  I32 F_XET
    /-libs/ch/023/name  S32 F_XET
    /-libs/ch/023/type  I32 F_XET
    /-libs/ch/023/flags  P32 F_XET
    /-libs/ch/023/hasdata  I32 F_XET
/-libs/ch/024  <SLIBS> n=0
    /-libs/ch/024/pos  I32 F_XET
    /-libs/ch/024/name  S32 F_XET
    /-libs/ch/024/type  I32 F_XET
    /-libs/ch/024/flags  P32 F_XET
    /-libs/ch/024/hasdata  I32 F_XET
/-libs/ch/025  <SLIBS> n=0
    /-libs/ch/025/pos  I32 F_XET
    /-libs/ch/025/name  S32 F_XET
    /-libs/ch/025/type  I32 F_XET
    /-libs/ch/025/flags  P32 F_XET
    /-libs/ch/025/hasdata  I32 F_XET
/-libs/ch/026  <SLIBS> n=0
    /-libs/ch/026/pos  I32 F_XET
    /-libs/ch/026/name  S32 F_XET
    /-libs/ch/026/type  I32 F_XET
    /-libs/ch/026/flags  P32 F_XET
    /-libs/ch/026/hasdata  I32 F_XET
/-libs/ch/027  <SLIBS> n=0
    /-libs/ch/027/pos  I32 F_XET
    /-libs/ch/027/name  S32 F_XET
    /-libs/ch/027/type  I32 F_XET
    /-libs/ch/027/flags  P32 F_XET
    /-libs/ch/027/hasdata  I32 F_XET
/-libs/ch/028  <SLIBS> n=0
    /-libs/ch/028/pos  I32 F_XET
    /-libs/ch/028/name  S32 F_XET
    /-libs/ch/028/type  I32 F_XET
    /-libs/ch/028/flags  P32 F_XET
    /-libs/ch/028/hasdata  I32 F_XET
/-libs/ch/029  <SLIBS> n=0
    /-libs/ch/029/pos  I32 F_XET
    /-libs/ch/029/name  S32 F_XET
    /-libs/ch/029/type  I32 F_XET
    /-libs/ch/029/flags  P32 F_XET
    /-libs/ch/029/hasdata  I32 F_XET
/-libs/ch/030  <SLIBS> n=0
    /-libs/ch/030/pos  I32 F_XET
    /-libs/ch/030/name  S32 F_XET
    /-libs/ch/030/type  I32 F_XET
    /-libs/ch/030/flags  P32 F_XET
    /-libs/ch/030/hasdata  I32 F_XET
/-libs/ch/031  <SLIBS> n=0
    /-libs/ch/031/pos  I32 F_XET
    /-libs/ch/031/name  S32 F_XET
    /-libs/ch/031/type  I32 F_XET
    /-libs/ch/031/flags  P32 F_XET
    /-libs/ch/031/hasdata  I32 F_XET
/-libs/ch/032  <SLIBS> n=0
    /-libs/ch/032/pos  I32 F_XET
    /-libs/ch/032/name  S32 F_XET
    /-libs/ch/032/type  I32 F_XET
    /-libs/ch/032/flags  P32 F_XET
    /-libs/ch/032/hasdata  I32 F_XET
/-libs/ch/033  <SLIBS> n=0
    /-libs/ch/033/pos  I32 F_XET
    /-libs/ch/033/name  S32 F_XET
    /-libs/ch/033/type  I32 F_XET
    /-libs/ch/033/flags  P32 F_XET
    /-libs/ch/033/hasdata  I32 F_XET
/-libs/ch/034  <SLIBS> n=0
    /-libs/ch/034/pos  I32 F_XET
    /-libs/ch/034/name  S32 F_XET
    /-libs/ch/034/type  I32 F_XET
    /-libs/ch/034/flags  P32 F_XET
    /-libs/ch/034/hasdata  I32 F_XET
/-libs/ch/035  <SLIBS> n=0
    /-libs/ch/035/pos  I32 F_XET
    /-libs/ch/035/name  S32 F_XET
    /-libs/ch/035/type  I32 F_XET
    /-libs/ch/035/flags  P32 F_XET
    /-libs/ch/035/hasdata  I32 F_XET
/-libs/ch/036  <SLIBS> n=0
    /-libs/ch/036/pos  I32 F_XET
    /-libs/ch/036/name  S32 F_XET
    /-libs/ch/036/type  I32 F_XET
    /-libs/ch/036/flags  P32 F_XET
    /-libs/ch/036/hasdata  I32 F_XET
/-libs/ch/037  <SLIBS> n=0
    /-libs/ch/037/pos  I32 F_XET
    /-libs/ch/037/name  S32 F_XET
    /-libs/ch/037/type  I32 F_XET
    /-libs/ch/037/flags  P32 F_XET
    /-libs/ch/037/hasdata  I32 F_XET
/-libs/ch/038  <SLIBS> n=0
    /-libs/ch/038/pos  I32 F_XET
    /-libs/ch/038/name  S32 F_XET
    /-libs/ch/038/type  I32 F_XET
    /-libs/ch/038/flags  P32 F_XET
    /-libs/ch/038/hasdata  I32 F_XET
/-libs/ch/039  <SLIBS> n=0
    /-libs/ch/039/pos  I32 F_XET
    /-libs/ch/039/name  S32 F_XET
    /-libs/ch/039/type  I32 F_XET
    /-libs/ch/039/flags  P32 F_XET
    /-libs/ch/039/hasdata  I32 F_XET
/-libs/ch/030  <SLIBS> n=0
    /-libs/ch/040/pos  I32 F_XET
    /-libs/ch/040/name  S32 F_XET
    /-libs/ch/040/type  I32 F_XET
    /-libs/ch/040/flags  P32 F_XET
    /-libs/ch/040/hasdata  I32 F_XET
/-libs/ch/041  <SLIBS> n=0
    /-libs/ch/041/pos  I32 F_XET
    /-libs/ch/041/name  S32 F_XET
    /-libs/ch/041/type  I32 F_XET
    /-libs/ch/041/flags  P32 F_XET
    /-libs/ch/041/hasdata  I32 F_XET
/-libs/ch/042  <SLIBS> n=0
    /-libs/ch/042/pos  I32 F_XET
    /-libs/ch/042/name  S32 F_XET
    /-libs/ch/042/type  I32 F_XET
    /-libs/ch/042/flags  P32 F_XET
    /-libs/ch/042/hasdata  I32 F_XET
/-libs/ch/043  <SLIBS> n=0
    /-libs/ch/043/pos  I32 F_XET
    /-libs/ch/043/name  S32 F_XET
    /-libs/ch/043/type  I32 F_XET
    /-libs/ch/043/flags  P32 F_XET
    /-libs/ch/043/hasdata  I32 F_XET
/-libs/ch/044  <SLIBS> n=0
    /-libs/ch/044/pos  I32 F_XET
    /-libs/ch/044/name  S32 F_XET
    /-libs/ch/044/type  I32 F_XET
    /-libs/ch/044/flags  P32 F_XET
    /-libs/ch/044/hasdata  I32 F_XET
/-libs/ch/045  <SLIBS> n=0
    /-libs/ch/045/pos  I32 F_XET
    /-libs/ch/045/name  S32 F_XET
    /-libs/ch/045/type  I32 F_XET
    /-libs/ch/045/flags  P32 F_XET
    /-libs/ch/045/hasdata  I32 F_XET
/-libs/ch/046  <SLIBS> n=0
    /-libs/ch/046/pos  I32 F_XET
    /-libs/ch/046/name  S32 F_XET
    /-libs/ch/046/type  I32 F_XET
    /-libs/ch/046/flags  P32 F_XET
    /-libs/ch/046/hasdata  I32 F_XET
/-libs/ch/047  <SLIBS> n=0
    /-libs/ch/047/pos  I32 F_XET
    /-libs/ch/047/name  S32 F_XET
    /-libs/ch/047/type  I32 F_XET
    /-libs/ch/047/flags  P32 F_XET
    /-libs/ch/047/hasdata  I32 F_XET
/-libs/ch/048  <SLIBS> n=0
    /-libs/ch/048/pos  I32 F_XET
    /-libs/ch/048/name  S32 F_XET
    /-libs/ch/048/type  I32 F_XET
    /-libs/ch/048/flags  P32 F_XET
    /-libs/ch/048/hasdata  I32 F_XET
/-libs/ch/049  <SLIBS> n=0
    /-libs/ch/049/pos  I32 F_XET
    /-libs/ch/049/name  S32 F_XET
    /-libs/ch/049/type  I32 F_XET
    /-libs/ch/049/flags  P32 F_XET
    /-libs/ch/049/hasdata  I32 F_XET
/-libs/ch/040  <SLIBS> n=0
    /-libs/ch/050/pos  I32 F_XET
    /-libs/ch/050/name  S32 F_XET
    /-libs/ch/050/type  I32 F_XET
    /-libs/ch/050/flags  P32 F_XET
    /-libs/ch/050/hasdata  I32 F_XET
/-libs/ch/051  <SLIBS> n=0
    /-libs/ch/051/pos  I32 F_XET
    /-libs/ch/051/name  S32 F_XET
    /-libs/ch/051/type  I32 F_XET
    /-libs/ch/051/flags  P32 F_XET
    /-libs/ch/051/hasdata  I32 F_XET
/-libs/ch/052  <SLIBS> n=0
    /-libs/ch/052/pos  I32 F_XET
    /-libs/ch/052/name  S32 F_XET
    /-libs/ch/052/type  I32 F_XET
    /-libs/ch/052/flags  P32 F_XET
    /-libs/ch/052/hasdata  I32 F_XET
/-libs/ch/053  <SLIBS> n=0
    /-libs/ch/053/pos  I32 F_XET
    /-libs/ch/053/name  S32 F_XET
    /-libs/ch/053/type  I32 F_XET
    /-libs/ch/053/flags  P32 F_XET
    /-libs/ch/053/hasdata  I32 F_XET
/-libs/ch/054  <SLIBS> n=0
    /-libs/ch/054/pos  I32 F_XET
    /-libs/ch/054/name  S32 F_XET
    /-libs/ch/054/type  I32 F_XET
    /-libs/ch/054/flags  P32 F_XET
    /-libs/ch/054/hasdata  I32 F_XET
/-libs/ch/055  <SLIBS> n=0
    /-libs/ch/055/pos  I32 F_XET
    /-libs/ch/055/name  S32 F_XET
    /-libs/ch/055/type  I32 F_XET
    /-libs/ch/055/flags  P32 F_XET
    /-libs/ch/055/hasdata  I32 F_XET
/-libs/ch/056  <SLIBS> n=0
    /-libs/ch/056/pos  I32 F_XET
    /-libs/ch/056/name  S32 F_XET
    /-libs/ch/056/type  I32 F_XET
    /-libs/ch/056/flags  P32 F_XET
    /-libs/ch/056/hasdata  I32 F_XET
/-libs/ch/057  <SLIBS> n=0
    /-libs/ch/057/pos  I32 F_XET
    /-libs/ch/057/name  S32 F_XET
    /-libs/ch/057/type  I32 F_XET
    /-libs/ch/057/flags  P32 F_XET
    /-libs/ch/057/hasdata  I32 F_XET
/-libs/ch/058  <SLIBS> n=0
    /-libs/ch/058/pos  I32 F_XET
    /-libs/ch/058/name  S32 F_XET
    /-libs/ch/058/type  I32 F_XET
    /-libs/ch/058/flags  P32 F_XET
    /-libs/ch/058/hasdata  I32 F_XET
/-libs/ch/059  <SLIBS> n=0
    /-libs/ch/059/pos  I32 F_XET
    /-libs/ch/059/name  S32 F_XET
    /-libs/ch/059/type  I32 F_XET
    /-libs/ch/059/flags  P32 F_XET
    /-libs/ch/059/hasdata  I32 F_XET
/-libs/ch/060  <SLIBS> n=0
    /-libs/ch/060/pos  I32 F_XET
    /-libs/ch/060/name  S32 F_XET
    /-libs/ch/060/type  I32 F_XET
    /-libs/ch/060/flags  P32 F_XET
    /-libs/ch/060/hasdata  I32 F_XET
/-libs/ch/061  <SLIBS> n=0
    /-libs/ch/061/pos  I32 F_XET
    /-libs/ch/061/name  S32 F_XET
    /-libs/ch/061/type  I32 F_XET
    /-libs/ch/061/flags  P32 F_XET
    /-libs/ch/061/hasdata  I32 F_XET
/-libs/ch/062  <SLIBS> n=0
    /-libs/ch/062/pos  I32 F_XET
    /-libs/ch/062/name  S32 F_XET
    /-libs/ch/062/type  I32 F_XET
    /-libs/ch/062/flags  P32 F_XET
    /-libs/ch/062/hasdata  I32 F_XET
/-libs/ch/063  <SLIBS> n=0
    /-libs/ch/063/pos  I32 F_XET
    /-libs/ch/063/name  S32 F_XET
    /-libs/ch/063/type  I32 F_XET
    /-libs/ch/063/flags  P32 F_XET
    /-libs/ch/063/hasdata  I32 F_XET
/-libs/ch/064  <SLIBS> n=0
    /-libs/ch/064/pos  I32 F_XET
    /-libs/ch/064/name  S32 F_XET
    /-libs/ch/064/type  I32 F_XET
    /-libs/ch/064/flags  P32 F_XET
    /-libs/ch/064/hasdata  I32 F_XET
/-libs/ch/065  <SLIBS> n=0
    /-libs/ch/065/pos  I32 F_XET
    /-libs/ch/065/name  S32 F_XET
    /-libs/ch/065/type  I32 F_XET
    /-libs/ch/065/flags  P32 F_XET
    /-libs/ch/065/hasdata  I32 F_XET
/-libs/ch/066  <SLIBS> n=0
    /-libs/ch/066/pos  I32 F_XET
    /-libs/ch/066/name  S32 F_XET
    /-libs/ch/066/type  I32 F_XET
    /-libs/ch/066/flags  P32 F_XET
    /-libs/ch/066/hasdata  I32 F_XET
/-libs/ch/067  <SLIBS> n=0
    /-libs/ch/067/pos  I32 F_XET
    /-libs/ch/067/name  S32 F_XET
    /-libs/ch/067/type  I32 F_XET
    /-libs/ch/067/flags  P32 F_XET
    /-libs/ch/067/hasdata  I32 F_XET
/-libs/ch/068  <SLIBS> n=0
    /-libs/ch/068/pos  I32 F_XET
    /-libs/ch/068/name  S32 F_XET
    /-libs/ch/068/type  I32 F_XET
    /-libs/ch/068/flags  P32 F_XET
    /-libs/ch/068/hasdata  I32 F_XET
/-libs/ch/069  <SLIBS> n=0
    /-libs/ch/069/pos  I32 F_XET
    /-libs/ch/069/name  S32 F_XET
    /-libs/ch/069/type  I32 F_XET
    /-libs/ch/069/flags  P32 F_XET
    /-libs/ch/069/hasdata  I32 F_XET
/-libs/ch/070  <SLIBS> n=0
    /-libs/ch/070/pos  I32 F_XET
    /-libs/ch/070/name  S32 F_XET
    /-libs/ch/070/type  I32 F_XET
    /-libs/ch/070/flags  P32 F_XET
    /-libs/ch/070/hasdata  I32 F_XET
/-libs/ch/071  <SLIBS> n=0
    /-libs/ch/071/pos  I32 F_XET
    /-libs/ch/071/name  S32 F_XET
    /-libs/ch/071/type  I32 F_XET
    /-libs/ch/071/flags  P32 F_XET
    /-libs/ch/071/hasdata  I32 F_XET
/-libs/ch/072  <SLIBS> n=0
    /-libs/ch/072/pos  I32 F_XET
    /-libs/ch/072/name  S32 F_XET
    /-libs/ch/072/type  I32 F_XET
    /-libs/ch/072/flags  P32 F_XET
    /-libs/ch/072/hasdata  I32 F_XET
/-libs/ch/073  <SLIBS> n=0
    /-libs/ch/073/pos  I32 F_XET
    /-libs/ch/073/name  S32 F_XET
    /-libs/ch/073/type  I32 F_XET
    /-libs/ch/073/flags  P32 F_XET
    /-libs/ch/073/hasdata  I32 F_XET
/-libs/ch/074  <SLIBS> n=0
    /-libs/ch/074/pos  I32 F_XET
    /-libs/ch/074/name  S32 F_XET
    /-libs/ch/074/type  I32 F_XET
    /-libs/ch/074/flags  P32 F_XET
    /-libs/ch/074/hasdata  I32 F_XET
/-libs/ch/075  <SLIBS> n=0
    /-libs/ch/075/pos  I32 F_XET
    /-libs/ch/075/name  S32 F_XET
    /-libs/ch/075/type  I32 F_XET
    /-libs/ch/075/flags  P32 F_XET
    /-libs/ch/075/hasdata  I32 F_XET
/-libs/ch/076  <SLIBS> n=0
    /-libs/ch/076/pos  I32 F_XET
    /-libs/ch/076/name  S32 F_XET
    /-libs/ch/076/type  I32 F_XET
    /-libs/ch/076/flags  P32 F_XET
    /-libs/ch/076/hasdata  I32 F_XET
/-libs/ch/077  <SLIBS> n=0
    /-libs/ch/077/pos  I32 F_XET
    /-libs/ch/077/name  S32 F_XET
    /-libs/ch/077/type  I32 F_XET
    /-libs/ch/077/flags  P32 F_XET
    /-libs/ch/077/hasdata  I32 F_XET
/-libs/ch/078  <SLIBS> n=0
    /-libs/ch/078/pos  I32 F_XET
    /-libs/ch/078/name  S32 F_XET
    /-libs/ch/078/type  I32 F_XET
    /-libs/ch/078/flags  P32 F_XET
    /-libs/ch/078/hasdata  I32 F_XET
/-libs/ch/079  <SLIBS> n=0
    /-libs/ch/079/pos  I32 F_XET
    /-libs/ch/079/name  S32 F_XET
    /-libs/ch/079/type  I32 F_XET
    /-libs/ch/079/flags  P32 F_XET
    /-libs/ch/079/hasdata  I32 F_XET
/-libs/ch/080  <SLIBS> n=0
    /-libs/ch/080/pos  I32 F_XET
    /-libs/ch/080/name  S32 F_XET
    /-libs/ch/080/type  I32 F_XET
    /-libs/ch/080/flags  P32 F_XET
    /-libs/ch/080/hasdata  I32 F_XET
/-libs/ch/081  <SLIBS> n=0
    /-libs/ch/081/pos  I32 F_XET
    /-libs/ch/081/name  S32 F_XET
    /-libs/ch/081/type  I32 F_XET
    /-libs/ch/081/flags  P32 F_XET
    /-libs/ch/081/hasdata  I32 F_XET
/-libs/ch/082  <SLIBS> n=0
    /-libs/ch/082/pos  I32 F_XET
    /-libs/ch/082/name  S32 F_XET
    /-libs/ch/082/type  I32 F_XET
    /-libs/ch/082/flags  P32 F_XET
    /-libs/ch/082/hasdata  I32 F_XET
/-libs/ch/083  <SLIBS> n=0
    /-libs/ch/083/pos  I32 F_XET
    /-libs/ch/083/name  S32 F_XET
    /-libs/ch/083/type  I32 F_XET
    /-libs/ch/083/flags  P32 F_XET
    /-libs/ch/083/hasdata  I32 F_XET
/-libs/ch/084  <SLIBS> n=0
    /-libs/ch/084/pos  I32 F_XET
    /-libs/ch/084/name  S32 F_XET
    /-libs/ch/084/type  I32 F_XET
    /-libs/ch/084/flags  P32 F_XET
    /-libs/ch/084/hasdata  I32 F_XET
/-libs/ch/085  <SLIBS> n=0
    /-libs/ch/085/pos  I32 F_XET
    /-libs/ch/085/name  S32 F_XET
    /-libs/ch/085/type  I32 F_XET
    /-libs/ch/085/flags  P32 F_XET
    /-libs/ch/085/hasdata  I32 F_XET
/-libs/ch/086  <SLIBS> n=0
    /-libs/ch/086/pos  I32 F_XET
    /-libs/ch/086/name  S32 F_XET
    /-libs/ch/086/type  I32 F_XET
    /-libs/ch/086/flags  P32 F_XET
    /-libs/ch/086/hasdata  I32 F_XET
/-libs/ch/087  <SLIBS> n=0
    /-libs/ch/087/pos  I32 F_XET
    /-libs/ch/087/name  S32 F_XET
    /-libs/ch/087/type  I32 F_XET
    /-libs/ch/087/flags  P32 F_XET
    /-libs/ch/087/hasdata  I32 F_XET
/-libs/ch/088  <SLIBS> n=0
    /-libs/ch/088/pos  I32 F_XET
    /-libs/ch/088/name  S32 F_XET
    /-libs/ch/088/type  I32 F_XET
    /-libs/ch/088/flags  P32 F_XET
    /-libs/ch/088/hasdata  I32 F_XET
/-libs/ch/089  <SLIBS> n=0
    /-libs/ch/089/pos  I32 F_XET
    /-libs/ch/089/name  S32 F_XET
    /-libs/ch/089/type  I32 F_XET
    /-libs/ch/089/flags  P32 F_XET
    /-libs/ch/089/hasdata  I32 F_XET
/-libs/ch/090  <SLIBS> n=0
    /-libs/ch/090/pos  I32 F_XET
    /-libs/ch/090/name  S32 F_XET
    /-libs/ch/090/type  I32 F_XET
    /-libs/ch/090/flags  P32 F_XET
    /-libs/ch/090/hasdata  I32 F_XET
/-libs/ch/091  <SLIBS> n=0
    /-libs/ch/091/pos  I32 F_XET
    /-libs/ch/091/name  S32 F_XET
    /-libs/ch/091/type  I32 F_XET
    /-libs/ch/091/flags  P32 F_XET
    /-libs/ch/091/hasdata  I32 F_XET
/-libs/ch/092  <SLIBS> n=0
    /-libs/ch/092/pos  I32 F_XET
    /-libs/ch/092/name  S32 F_XET
    /-libs/ch/092/type  I32 F_XET
    /-libs/ch/092/flags  P32 F_XET
    /-libs/ch/092/hasdata  I32 F_XET
/-libs/ch/093  <SLIBS> n=0
    /-libs/ch/093/pos  I32 F_XET
    /-libs/ch/093/name  S32 F_XET
    /-libs/ch/093/type  I32 F_XET
    /-libs/ch/093/flags  P32 F_XET
    /-libs/ch/093/hasdata  I32 F_XET
/-libs/ch/094  <SLIBS> n=0
    /-libs/ch/094/pos  I32 F_XET
    /-libs/ch/094/name  S32 F_XET
    /-libs/ch/094/type  I32 F_XET
    /-libs/ch/094/flags  P32 F_XET
    /-libs/ch/094/hasdata  I32 F_XET
/-libs/ch/095  <SLIBS> n=0
    /-libs/ch/095/pos  I32 F_XET
    /-libs/ch/095/name  S32 F_XET
    /-libs/ch/095/type  I32 F_XET
    /-libs/ch/095/flags  P32 F_XET
    /-libs/ch/095/hasdata  I32 F_XET
/-libs/ch/096  <SLIBS> n=0
    /-libs/ch/096/pos  I32 F_XET
    /-libs/ch/096/name  S32 F_XET
    /-libs/ch/096/type  I32 F_XET
    /-libs/ch/096/flags  P32 F_XET
    /-libs/ch/096/hasdata  I32 F_XET
/-libs/ch/097  <SLIBS> n=0
    /-libs/ch/097/pos  I32 F_XET
    /-libs/ch/097/name  S32 F_XET
    /-libs/ch/097/type  I32 F_XET
    /-libs/ch/097/flags  P32 F_XET
    /-libs/ch/097/hasdata  I32 F_XET
/-libs/ch/098  <SLIBS> n=0
    /-libs/ch/098/pos  I32 F_XET
    /-libs/ch/098/name  S32 F_XET
    /-libs/ch/098/type  I32 F_XET
    /-libs/ch/098/flags  P32 F_XET
    /-libs/ch/098/hasdata  I32 F_XET
/-libs/ch/099  <SLIBS> n=0
    /-libs/ch/099/pos  I32 F_XET
    /-libs/ch/099/name  S32 F_XET
    /-libs/ch/099/type  I32 F_XET
    /-libs/ch/099/flags  P32 F_XET
    /-libs/ch/099/hasdata  I32 F_XET
/-libs/ch/100  <SLIBS> n=0
    /-libs/ch/100/pos  I32 F_XET
    /-libs/ch/100/name  S32 F_XET
    /-libs/ch/100/type  I32 F_XET
    /-libs/ch/100/flags  P32 F_XET
    /-libs/ch/100/hasdata  I32 F_XET
```

### Xlibsr (X32Libs.h, 601 entries)

```
/-libs/r  <SLIBS> n=0
/-libs/r/001  <SLIBS> n=0
    /-libs/r/001/pos  I32 F_XET
    /-libs/r/001/name  S32 F_XET
    /-libs/r/001/type  I32 F_XET
    /-libs/r/001/flags  P32 F_XET
    /-libs/r/001/hasdata  I32 F_XET
/-libs/r/002  <SLIBS> n=0
    /-libs/r/002/pos  I32 F_XET
    /-libs/r/002/name  S32 F_XET
    /-libs/r/002/type  I32 F_XET
    /-libs/r/002/flags  P32 F_XET
    /-libs/r/002/hasdata  I32 F_XET
/-libs/r/003  <SLIBS> n=0
    /-libs/r/003/pos  I32 F_XET
    /-libs/r/003/name  S32 F_XET
    /-libs/r/003/type  I32 F_XET
    /-libs/r/003/flags  P32 F_XET
    /-libs/r/003/hasdata  I32 F_XET
/-libs/r/004  <SLIBS> n=0
    /-libs/r/004/pos  I32 F_XET
    /-libs/r/004/name  S32 F_XET
    /-libs/r/004/type  I32 F_XET
    /-libs/r/004/flags  P32 F_XET
    /-libs/r/004/hasdata  I32 F_XET
/-libs/r/005  <SLIBS> n=0
    /-libs/r/005/pos  I32 F_XET
    /-libs/r/005/name  S32 F_XET
    /-libs/r/005/type  I32 F_XET
    /-libs/r/005/flags  P32 F_XET
    /-libs/r/005/hasdata  I32 F_XET
/-libs/r/006  <SLIBS> n=0
    /-libs/r/006/pos  I32 F_XET
    /-libs/r/006/name  S32 F_XET
    /-libs/r/006/type  I32 F_XET
    /-libs/r/006/flags  P32 F_XET
    /-libs/r/006/hasdata  I32 F_XET
/-libs/r/007  <SLIBS> n=0
    /-libs/r/007/pos  I32 F_XET
    /-libs/r/007/name  S32 F_XET
    /-libs/r/007/type  I32 F_XET
    /-libs/r/007/flags  P32 F_XET
    /-libs/r/007/hasdata  I32 F_XET
/-libs/r/008  <SLIBS> n=0
    /-libs/r/008/pos  I32 F_XET
    /-libs/r/008/name  S32 F_XET
    /-libs/r/008/type  I32 F_XET
    /-libs/r/008/flags  P32 F_XET
    /-libs/r/008/hasdata  I32 F_XET
/-libs/r/009  <SLIBS> n=0
    /-libs/r/009/pos  I32 F_XET
    /-libs/r/009/name  S32 F_XET
    /-libs/r/009/type  I32 F_XET
    /-libs/r/009/flags  P32 F_XET
    /-libs/r/009/hasdata  I32 F_XET
/-libs/r/010  <SLIBS> n=0
    /-libs/r/010/pos  I32 F_XET
    /-libs/r/010/name  S32 F_XET
    /-libs/r/010/type  I32 F_XET
    /-libs/r/010/flags  P32 F_XET
    /-libs/r/010/hasdata  I32 F_XET
/-libs/r/011  <SLIBS> n=0
    /-libs/r/011/pos  I32 F_XET
    /-libs/r/011/name  S32 F_XET
    /-libs/r/011/type  I32 F_XET
    /-libs/r/011/flags  P32 F_XET
    /-libs/r/011/hasdata  I32 F_XET
/-libs/r/012  <SLIBS> n=0
    /-libs/r/012/pos  I32 F_XET
    /-libs/r/012/name  S32 F_XET
    /-libs/r/012/type  I32 F_XET
    /-libs/r/012/flags  P32 F_XET
    /-libs/r/012/hasdata  I32 F_XET
/-libs/r/013  <SLIBS> n=0
    /-libs/r/013/pos  I32 F_XET
    /-libs/r/013/name  S32 F_XET
    /-libs/r/013/type  I32 F_XET
    /-libs/r/013/flags  P32 F_XET
    /-libs/r/013/hasdata  I32 F_XET
/-libs/r/014  <SLIBS> n=0
    /-libs/r/014/pos  I32 F_XET
    /-libs/r/014/name  S32 F_XET
    /-libs/r/014/type  I32 F_XET
    /-libs/r/014/flags  P32 F_XET
    /-libs/r/014/hasdata  I32 F_XET
/-libs/r/015  <SLIBS> n=0
    /-libs/r/015/pos  I32 F_XET
    /-libs/r/015/name  S32 F_XET
    /-libs/r/015/type  I32 F_XET
    /-libs/r/015/flags  P32 F_XET
    /-libs/r/015/hasdata  I32 F_XET
/-libs/r/016  <SLIBS> n=0
    /-libs/r/016/pos  I32 F_XET
    /-libs/r/016/name  S32 F_XET
    /-libs/r/016/type  I32 F_XET
    /-libs/r/016/flags  P32 F_XET
    /-libs/r/016/hasdata  I32 F_XET
/-libs/r/017  <SLIBS> n=0
    /-libs/r/017/pos  I32 F_XET
    /-libs/r/017/name  S32 F_XET
    /-libs/r/017/type  I32 F_XET
    /-libs/r/017/flags  P32 F_XET
    /-libs/r/017/hasdata  I32 F_XET
/-libs/r/018  <SLIBS> n=0
    /-libs/r/018/pos  I32 F_XET
    /-libs/r/018/name  S32 F_XET
    /-libs/r/018/type  I32 F_XET
    /-libs/r/018/flags  P32 F_XET
    /-libs/r/018/hasdata  I32 F_XET
/-libs/r/019  <SLIBS> n=0
    /-libs/r/019/pos  I32 F_XET
    /-libs/r/019/name  S32 F_XET
    /-libs/r/019/type  I32 F_XET
    /-libs/r/019/flags  P32 F_XET
    /-libs/r/019/hasdata  I32 F_XET
/-libs/r/020  <SLIBS> n=0
    /-libs/r/020/pos  I32 F_XET
    /-libs/r/020/name  S32 F_XET
    /-libs/r/020/type  I32 F_XET
    /-libs/r/020/flags  P32 F_XET
    /-libs/r/020/hasdata  I32 F_XET
/-libs/r/021  <SLIBS> n=0
    /-libs/r/021/pos  I32 F_XET
    /-libs/r/021/name  S32 F_XET
    /-libs/r/021/type  I32 F_XET
    /-libs/r/021/flags  P32 F_XET
    /-libs/r/021/hasdata  I32 F_XET
/-libs/r/022  <SLIBS> n=0
    /-libs/r/022/pos  I32 F_XET
    /-libs/r/022/name  S32 F_XET
    /-libs/r/022/type  I32 F_XET
    /-libs/r/022/flags  P32 F_XET
    /-libs/r/022/hasdata  I32 F_XET
/-libs/r/023  <SLIBS> n=0
    /-libs/r/023/pos  I32 F_XET
    /-libs/r/023/name  S32 F_XET
    /-libs/r/023/type  I32 F_XET
    /-libs/r/023/flags  P32 F_XET
    /-libs/r/023/hasdata  I32 F_XET
/-libs/r/024  <SLIBS> n=0
    /-libs/r/024/pos  I32 F_XET
    /-libs/r/024/name  S32 F_XET
    /-libs/r/024/type  I32 F_XET
    /-libs/r/024/flags  P32 F_XET
    /-libs/r/024/hasdata  I32 F_XET
/-libs/r/025  <SLIBS> n=0
    /-libs/r/025/pos  I32 F_XET
    /-libs/r/025/name  S32 F_XET
    /-libs/r/025/type  I32 F_XET
    /-libs/r/025/flags  P32 F_XET
    /-libs/r/025/hasdata  I32 F_XET
/-libs/r/026  <SLIBS> n=0
    /-libs/r/026/pos  I32 F_XET
    /-libs/r/026/name  S32 F_XET
    /-libs/r/026/type  I32 F_XET
    /-libs/r/026/flags  P32 F_XET
    /-libs/r/026/hasdata  I32 F_XET
/-libs/r/027  <SLIBS> n=0
    /-libs/r/027/pos  I32 F_XET
    /-libs/r/027/name  S32 F_XET
    /-libs/r/027/type  I32 F_XET
    /-libs/r/027/flags  P32 F_XET
    /-libs/r/027/hasdata  I32 F_XET
/-libs/r/028  <SLIBS> n=0
    /-libs/r/028/pos  I32 F_XET
    /-libs/r/028/name  S32 F_XET
    /-libs/r/028/type  I32 F_XET
    /-libs/r/028/flags  P32 F_XET
    /-libs/r/028/hasdata  I32 F_XET
/-libs/r/029  <SLIBS> n=0
    /-libs/r/029/pos  I32 F_XET
    /-libs/r/029/name  S32 F_XET
    /-libs/r/029/type  I32 F_XET
    /-libs/r/029/flags  P32 F_XET
    /-libs/r/029/hasdata  I32 F_XET
/-libs/r/030  <SLIBS> n=0
    /-libs/r/030/pos  I32 F_XET
    /-libs/r/030/name  S32 F_XET
    /-libs/r/030/type  I32 F_XET
    /-libs/r/030/flags  P32 F_XET
    /-libs/r/030/hasdata  I32 F_XET
/-libs/r/031  <SLIBS> n=0
    /-libs/r/031/pos  I32 F_XET
    /-libs/r/031/name  S32 F_XET
    /-libs/r/031/type  I32 F_XET
    /-libs/r/031/flags  P32 F_XET
    /-libs/r/031/hasdata  I32 F_XET
/-libs/r/032  <SLIBS> n=0
    /-libs/r/032/pos  I32 F_XET
    /-libs/r/032/name  S32 F_XET
    /-libs/r/032/type  I32 F_XET
    /-libs/r/032/flags  P32 F_XET
    /-libs/r/032/hasdata  I32 F_XET
/-libs/r/033  <SLIBS> n=0
    /-libs/r/033/pos  I32 F_XET
    /-libs/r/033/name  S32 F_XET
    /-libs/r/033/type  I32 F_XET
    /-libs/r/033/flags  P32 F_XET
    /-libs/r/033/hasdata  I32 F_XET
/-libs/r/034  <SLIBS> n=0
    /-libs/r/034/pos  I32 F_XET
    /-libs/r/034/name  S32 F_XET
    /-libs/r/034/type  I32 F_XET
    /-libs/r/034/flags  P32 F_XET
    /-libs/r/034/hasdata  I32 F_XET
/-libs/r/035  <SLIBS> n=0
    /-libs/r/035/pos  I32 F_XET
    /-libs/r/035/name  S32 F_XET
    /-libs/r/035/type  I32 F_XET
    /-libs/r/035/flags  P32 F_XET
    /-libs/r/035/hasdata  I32 F_XET
/-libs/r/036  <SLIBS> n=0
    /-libs/r/036/pos  I32 F_XET
    /-libs/r/036/name  S32 F_XET
    /-libs/r/036/type  I32 F_XET
    /-libs/r/036/flags  P32 F_XET
    /-libs/r/036/hasdata  I32 F_XET
/-libs/r/037  <SLIBS> n=0
    /-libs/r/037/pos  I32 F_XET
    /-libs/r/037/name  S32 F_XET
    /-libs/r/037/type  I32 F_XET
    /-libs/r/037/flags  P32 F_XET
    /-libs/r/037/hasdata  I32 F_XET
/-libs/r/038  <SLIBS> n=0
    /-libs/r/038/pos  I32 F_XET
    /-libs/r/038/name  S32 F_XET
    /-libs/r/038/type  I32 F_XET
    /-libs/r/038/flags  P32 F_XET
    /-libs/r/038/hasdata  I32 F_XET
/-libs/r/039  <SLIBS> n=0
    /-libs/r/039/pos  I32 F_XET
    /-libs/r/039/name  S32 F_XET
    /-libs/r/039/type  I32 F_XET
    /-libs/r/039/flags  P32 F_XET
    /-libs/r/039/hasdata  I32 F_XET
/-libs/r/030  <SLIBS> n=0
    /-libs/r/040/pos  I32 F_XET
    /-libs/r/040/name  S32 F_XET
    /-libs/r/040/type  I32 F_XET
    /-libs/r/040/flags  P32 F_XET
    /-libs/r/040/hasdata  I32 F_XET
/-libs/r/041  <SLIBS> n=0
    /-libs/r/041/pos  I32 F_XET
    /-libs/r/041/name  S32 F_XET
    /-libs/r/041/type  I32 F_XET
    /-libs/r/041/flags  P32 F_XET
    /-libs/r/041/hasdata  I32 F_XET
/-libs/r/042  <SLIBS> n=0
    /-libs/r/042/pos  I32 F_XET
    /-libs/r/042/name  S32 F_XET
    /-libs/r/042/type  I32 F_XET
    /-libs/r/042/flags  P32 F_XET
    /-libs/r/042/hasdata  I32 F_XET
/-libs/r/043  <SLIBS> n=0
    /-libs/r/043/pos  I32 F_XET
    /-libs/r/043/name  S32 F_XET
    /-libs/r/043/type  I32 F_XET
    /-libs/r/043/flags  P32 F_XET
    /-libs/r/043/hasdata  I32 F_XET
/-libs/r/044  <SLIBS> n=0
    /-libs/r/044/pos  I32 F_XET
    /-libs/r/044/name  S32 F_XET
    /-libs/r/044/type  I32 F_XET
    /-libs/r/044/flags  P32 F_XET
    /-libs/r/044/hasdata  I32 F_XET
/-libs/r/045  <SLIBS> n=0
    /-libs/r/045/pos  I32 F_XET
    /-libs/r/045/name  S32 F_XET
    /-libs/r/045/type  I32 F_XET
    /-libs/r/045/flags  P32 F_XET
    /-libs/r/045/hasdata  I32 F_XET
/-libs/r/046  <SLIBS> n=0
    /-libs/r/046/pos  I32 F_XET
    /-libs/r/046/name  S32 F_XET
    /-libs/r/046/type  I32 F_XET
    /-libs/r/046/flags  P32 F_XET
    /-libs/r/046/hasdata  I32 F_XET
/-libs/r/047  <SLIBS> n=0
    /-libs/r/047/pos  I32 F_XET
    /-libs/r/047/name  S32 F_XET
    /-libs/r/047/type  I32 F_XET
    /-libs/r/047/flags  P32 F_XET
    /-libs/r/047/hasdata  I32 F_XET
/-libs/r/048  <SLIBS> n=0
    /-libs/r/048/pos  I32 F_XET
    /-libs/r/048/name  S32 F_XET
    /-libs/r/048/type  I32 F_XET
    /-libs/r/048/flags  P32 F_XET
    /-libs/r/048/hasdata  I32 F_XET
/-libs/r/049  <SLIBS> n=0
    /-libs/r/049/pos  I32 F_XET
    /-libs/r/049/name  S32 F_XET
    /-libs/r/049/type  I32 F_XET
    /-libs/r/049/flags  P32 F_XET
    /-libs/r/049/hasdata  I32 F_XET
/-libs/r/040  <SLIBS> n=0
    /-libs/r/050/pos  I32 F_XET
    /-libs/r/050/name  S32 F_XET
    /-libs/r/050/type  I32 F_XET
    /-libs/r/050/flags  P32 F_XET
    /-libs/r/050/hasdata  I32 F_XET
/-libs/r/051  <SLIBS> n=0
    /-libs/r/051/pos  I32 F_XET
    /-libs/r/051/name  S32 F_XET
    /-libs/r/051/type  I32 F_XET
    /-libs/r/051/flags  P32 F_XET
    /-libs/r/051/hasdata  I32 F_XET
/-libs/r/052  <SLIBS> n=0
    /-libs/r/052/pos  I32 F_XET
    /-libs/r/052/name  S32 F_XET
    /-libs/r/052/type  I32 F_XET
    /-libs/r/052/flags  P32 F_XET
    /-libs/r/052/hasdata  I32 F_XET
/-libs/r/053  <SLIBS> n=0
    /-libs/r/053/pos  I32 F_XET
    /-libs/r/053/name  S32 F_XET
    /-libs/r/053/type  I32 F_XET
    /-libs/r/053/flags  P32 F_XET
    /-libs/r/053/hasdata  I32 F_XET
/-libs/r/054  <SLIBS> n=0
    /-libs/r/054/pos  I32 F_XET
    /-libs/r/054/name  S32 F_XET
    /-libs/r/054/type  I32 F_XET
    /-libs/r/054/flags  P32 F_XET
    /-libs/r/054/hasdata  I32 F_XET
/-libs/r/055  <SLIBS> n=0
    /-libs/r/055/pos  I32 F_XET
    /-libs/r/055/name  S32 F_XET
    /-libs/r/055/type  I32 F_XET
    /-libs/r/055/flags  P32 F_XET
    /-libs/r/055/hasdata  I32 F_XET
/-libs/r/056  <SLIBS> n=0
    /-libs/r/056/pos  I32 F_XET
    /-libs/r/056/name  S32 F_XET
    /-libs/r/056/type  I32 F_XET
    /-libs/r/056/flags  P32 F_XET
    /-libs/r/056/hasdata  I32 F_XET
/-libs/r/057  <SLIBS> n=0
    /-libs/r/057/pos  I32 F_XET
    /-libs/r/057/name  S32 F_XET
    /-libs/r/057/type  I32 F_XET
    /-libs/r/057/flags  P32 F_XET
    /-libs/r/057/hasdata  I32 F_XET
/-libs/r/058  <SLIBS> n=0
    /-libs/r/058/pos  I32 F_XET
    /-libs/r/058/name  S32 F_XET
    /-libs/r/058/type  I32 F_XET
    /-libs/r/058/flags  P32 F_XET
    /-libs/r/058/hasdata  I32 F_XET
/-libs/r/059  <SLIBS> n=0
    /-libs/r/059/pos  I32 F_XET
    /-libs/r/059/name  S32 F_XET
    /-libs/r/059/type  I32 F_XET
    /-libs/r/059/flags  P32 F_XET
    /-libs/r/059/hasdata  I32 F_XET
/-libs/r/060  <SLIBS> n=0
    /-libs/r/060/pos  I32 F_XET
    /-libs/r/060/name  S32 F_XET
    /-libs/r/060/type  I32 F_XET
    /-libs/r/060/flags  P32 F_XET
    /-libs/r/060/hasdata  I32 F_XET
/-libs/r/061  <SLIBS> n=0
    /-libs/r/061/pos  I32 F_XET
    /-libs/r/061/name  S32 F_XET
    /-libs/r/061/type  I32 F_XET
    /-libs/r/061/flags  P32 F_XET
    /-libs/r/061/hasdata  I32 F_XET
/-libs/r/062  <SLIBS> n=0
    /-libs/r/062/pos  I32 F_XET
    /-libs/r/062/name  S32 F_XET
    /-libs/r/062/type  I32 F_XET
    /-libs/r/062/flags  P32 F_XET
    /-libs/r/062/hasdata  I32 F_XET
/-libs/r/063  <SLIBS> n=0
    /-libs/r/063/pos  I32 F_XET
    /-libs/r/063/name  S32 F_XET
    /-libs/r/063/type  I32 F_XET
    /-libs/r/063/flags  P32 F_XET
    /-libs/r/063/hasdata  I32 F_XET
/-libs/r/064  <SLIBS> n=0
    /-libs/r/064/pos  I32 F_XET
    /-libs/r/064/name  S32 F_XET
    /-libs/r/064/type  I32 F_XET
    /-libs/r/064/flags  P32 F_XET
    /-libs/r/064/hasdata  I32 F_XET
/-libs/r/065  <SLIBS> n=0
    /-libs/r/065/pos  I32 F_XET
    /-libs/r/065/name  S32 F_XET
    /-libs/r/065/type  I32 F_XET
    /-libs/r/065/flags  P32 F_XET
    /-libs/r/065/hasdata  I32 F_XET
/-libs/r/066  <SLIBS> n=0
    /-libs/r/066/pos  I32 F_XET
    /-libs/r/066/name  S32 F_XET
    /-libs/r/066/type  I32 F_XET
    /-libs/r/066/flags  P32 F_XET
    /-libs/r/066/hasdata  I32 F_XET
/-libs/r/067  <SLIBS> n=0
    /-libs/r/067/pos  I32 F_XET
    /-libs/r/067/name  S32 F_XET
    /-libs/r/067/type  I32 F_XET
    /-libs/r/067/flags  P32 F_XET
    /-libs/r/067/hasdata  I32 F_XET
/-libs/r/068  <SLIBS> n=0
    /-libs/r/068/pos  I32 F_XET
    /-libs/r/068/name  S32 F_XET
    /-libs/r/068/type  I32 F_XET
    /-libs/r/068/flags  P32 F_XET
    /-libs/r/068/hasdata  I32 F_XET
/-libs/r/069  <SLIBS> n=0
    /-libs/r/069/pos  I32 F_XET
    /-libs/r/069/name  S32 F_XET
    /-libs/r/069/type  I32 F_XET
    /-libs/r/069/flags  P32 F_XET
    /-libs/r/069/hasdata  I32 F_XET
/-libs/r/070  <SLIBS> n=0
    /-libs/r/070/pos  I32 F_XET
    /-libs/r/070/name  S32 F_XET
    /-libs/r/070/type  I32 F_XET
    /-libs/r/070/flags  P32 F_XET
    /-libs/r/070/hasdata  I32 F_XET
/-libs/r/071  <SLIBS> n=0
    /-libs/r/071/pos  I32 F_XET
    /-libs/r/071/name  S32 F_XET
    /-libs/r/071/type  I32 F_XET
    /-libs/r/071/flags  P32 F_XET
    /-libs/r/071/hasdata  I32 F_XET
/-libs/r/072  <SLIBS> n=0
    /-libs/r/072/pos  I32 F_XET
    /-libs/r/072/name  S32 F_XET
    /-libs/r/072/type  I32 F_XET
    /-libs/r/072/flags  P32 F_XET
    /-libs/r/072/hasdata  I32 F_XET
/-libs/r/073  <SLIBS> n=0
    /-libs/r/073/pos  I32 F_XET
    /-libs/r/073/name  S32 F_XET
    /-libs/r/073/type  I32 F_XET
    /-libs/r/073/flags  P32 F_XET
    /-libs/r/073/hasdata  I32 F_XET
/-libs/r/074  <SLIBS> n=0
    /-libs/r/074/pos  I32 F_XET
    /-libs/r/074/name  S32 F_XET
    /-libs/r/074/type  I32 F_XET
    /-libs/r/074/flags  P32 F_XET
    /-libs/r/074/hasdata  I32 F_XET
/-libs/r/075  <SLIBS> n=0
    /-libs/r/075/pos  I32 F_XET
    /-libs/r/075/name  S32 F_XET
    /-libs/r/075/type  I32 F_XET
    /-libs/r/075/flags  P32 F_XET
    /-libs/r/075/hasdata  I32 F_XET
/-libs/r/076  <SLIBS> n=0
    /-libs/r/076/pos  I32 F_XET
    /-libs/r/076/name  S32 F_XET
    /-libs/r/076/type  I32 F_XET
    /-libs/r/076/flags  P32 F_XET
    /-libs/r/076/hasdata  I32 F_XET
/-libs/r/077  <SLIBS> n=0
    /-libs/r/077/pos  I32 F_XET
    /-libs/r/077/name  S32 F_XET
    /-libs/r/077/type  I32 F_XET
    /-libs/r/077/flags  P32 F_XET
    /-libs/r/077/hasdata  I32 F_XET
/-libs/r/078  <SLIBS> n=0
    /-libs/r/078/pos  I32 F_XET
    /-libs/r/078/name  S32 F_XET
    /-libs/r/078/type  I32 F_XET
    /-libs/r/078/flags  P32 F_XET
    /-libs/r/078/hasdata  I32 F_XET
/-libs/r/079  <SLIBS> n=0
    /-libs/r/079/pos  I32 F_XET
    /-libs/r/079/name  S32 F_XET
    /-libs/r/079/type  I32 F_XET
    /-libs/r/079/flags  P32 F_XET
    /-libs/r/079/hasdata  I32 F_XET
/-libs/r/080  <SLIBS> n=0
    /-libs/r/080/pos  I32 F_XET
    /-libs/r/080/name  S32 F_XET
    /-libs/r/080/type  I32 F_XET
    /-libs/r/080/flags  P32 F_XET
    /-libs/r/080/hasdata  I32 F_XET
/-libs/r/081  <SLIBS> n=0
    /-libs/r/081/pos  I32 F_XET
    /-libs/r/081/name  S32 F_XET
    /-libs/r/081/type  I32 F_XET
    /-libs/r/081/flags  P32 F_XET
    /-libs/r/081/hasdata  I32 F_XET
/-libs/r/082  <SLIBS> n=0
    /-libs/r/082/pos  I32 F_XET
    /-libs/r/082/name  S32 F_XET
    /-libs/r/082/type  I32 F_XET
    /-libs/r/082/flags  P32 F_XET
    /-libs/r/082/hasdata  I32 F_XET
/-libs/r/083  <SLIBS> n=0
    /-libs/r/083/pos  I32 F_XET
    /-libs/r/083/name  S32 F_XET
    /-libs/r/083/type  I32 F_XET
    /-libs/r/083/flags  P32 F_XET
    /-libs/r/083/hasdata  I32 F_XET
/-libs/r/084  <SLIBS> n=0
    /-libs/r/084/pos  I32 F_XET
    /-libs/r/084/name  S32 F_XET
    /-libs/r/084/type  I32 F_XET
    /-libs/r/084/flags  P32 F_XET
    /-libs/r/084/hasdata  I32 F_XET
/-libs/r/085  <SLIBS> n=0
    /-libs/r/085/pos  I32 F_XET
    /-libs/r/085/name  S32 F_XET
    /-libs/r/085/type  I32 F_XET
    /-libs/r/085/flags  P32 F_XET
    /-libs/r/085/hasdata  I32 F_XET
/-libs/r/086  <SLIBS> n=0
    /-libs/r/086/pos  I32 F_XET
    /-libs/r/086/name  S32 F_XET
    /-libs/r/086/type  I32 F_XET
    /-libs/r/086/flags  P32 F_XET
    /-libs/r/086/hasdata  I32 F_XET
/-libs/r/087  <SLIBS> n=0
    /-libs/r/087/pos  I32 F_XET
    /-libs/r/087/name  S32 F_XET
    /-libs/r/087/type  I32 F_XET
    /-libs/r/087/flags  P32 F_XET
    /-libs/r/087/hasdata  I32 F_XET
/-libs/r/088  <SLIBS> n=0
    /-libs/r/088/pos  I32 F_XET
    /-libs/r/088/name  S32 F_XET
    /-libs/r/088/type  I32 F_XET
    /-libs/r/088/flags  P32 F_XET
    /-libs/r/088/hasdata  I32 F_XET
/-libs/r/089  <SLIBS> n=0
    /-libs/r/089/pos  I32 F_XET
    /-libs/r/089/name  S32 F_XET
    /-libs/r/089/type  I32 F_XET
    /-libs/r/089/flags  P32 F_XET
    /-libs/r/089/hasdata  I32 F_XET
/-libs/r/090  <SLIBS> n=0
    /-libs/r/090/pos  I32 F_XET
    /-libs/r/090/name  S32 F_XET
    /-libs/r/090/type  I32 F_XET
    /-libs/r/090/flags  P32 F_XET
    /-libs/r/090/hasdata  I32 F_XET
/-libs/r/091  <SLIBS> n=0
    /-libs/r/091/pos  I32 F_XET
    /-libs/r/091/name  S32 F_XET
    /-libs/r/091/type  I32 F_XET
    /-libs/r/091/flags  P32 F_XET
    /-libs/r/091/hasdata  I32 F_XET
/-libs/r/092  <SLIBS> n=0
    /-libs/r/092/pos  I32 F_XET
    /-libs/r/092/name  S32 F_XET
    /-libs/r/092/type  I32 F_XET
    /-libs/r/092/flags  P32 F_XET
    /-libs/r/092/hasdata  I32 F_XET
/-libs/r/093  <SLIBS> n=0
    /-libs/r/093/pos  I32 F_XET
    /-libs/r/093/name  S32 F_XET
    /-libs/r/093/type  I32 F_XET
    /-libs/r/093/flags  P32 F_XET
    /-libs/r/093/hasdata  I32 F_XET
/-libs/r/094  <SLIBS> n=0
    /-libs/r/094/pos  I32 F_XET
    /-libs/r/094/name  S32 F_XET
    /-libs/r/094/type  I32 F_XET
    /-libs/r/094/flags  P32 F_XET
    /-libs/r/094/hasdata  I32 F_XET
/-libs/r/095  <SLIBS> n=0
    /-libs/r/095/pos  I32 F_XET
    /-libs/r/095/name  S32 F_XET
    /-libs/r/095/type  I32 F_XET
    /-libs/r/095/flags  P32 F_XET
    /-libs/r/095/hasdata  I32 F_XET
/-libs/r/096  <SLIBS> n=0
    /-libs/r/096/pos  I32 F_XET
    /-libs/r/096/name  S32 F_XET
    /-libs/r/096/type  I32 F_XET
    /-libs/r/096/flags  P32 F_XET
    /-libs/r/096/hasdata  I32 F_XET
/-libs/r/097  <SLIBS> n=0
    /-libs/r/097/pos  I32 F_XET
    /-libs/r/097/name  S32 F_XET
    /-libs/r/097/type  I32 F_XET
    /-libs/r/097/flags  P32 F_XET
    /-libs/r/097/hasdata  I32 F_XET
/-libs/r/098  <SLIBS> n=0
    /-libs/r/098/pos  I32 F_XET
    /-libs/r/098/name  S32 F_XET
    /-libs/r/098/type  I32 F_XET
    /-libs/r/098/flags  P32 F_XET
    /-libs/r/098/hasdata  I32 F_XET
/-libs/r/099  <SLIBS> n=0
    /-libs/r/099/pos  I32 F_XET
    /-libs/r/099/name  S32 F_XET
    /-libs/r/099/type  I32 F_XET
    /-libs/r/099/flags  P32 F_XET
    /-libs/r/099/hasdata  I32 F_XET
/-libs/r/100  <SLIBS> n=0
    /-libs/r/100/pos  I32 F_XET
    /-libs/r/100/name  S32 F_XET
    /-libs/r/100/type  I32 F_XET
    /-libs/r/100/flags  P32 F_XET
    /-libs/r/100/hasdata  I32 F_XET
```

### Xlibsf (X32Libs.h, 601 entries)

```
/-libs/fx  <SLIBS> n=0
/-libs/fx/001  <SLIBS> n=0
    /-libs/fx/001/pos  I32 F_XET
    /-libs/fx/001/name  S32 F_XET
    /-libs/fx/001/type  I32 F_XET
    /-libs/fx/001/flags  P32 F_XET
    /-libs/fx/001/hasdata  I32 F_XET
/-libs/fx/002  <SLIBS> n=0
    /-libs/fx/002/pos  I32 F_XET
    /-libs/fx/002/name  S32 F_XET
    /-libs/fx/002/type  I32 F_XET
    /-libs/fx/002/flags  P32 F_XET
    /-libs/fx/002/hasdata  I32 F_XET
/-libs/fx/003  <SLIBS> n=0
    /-libs/fx/003/pos  I32 F_XET
    /-libs/fx/003/name  S32 F_XET
    /-libs/fx/003/type  I32 F_XET
    /-libs/fx/003/flags  P32 F_XET
    /-libs/fx/003/hasdata  I32 F_XET
/-libs/fx/004  <SLIBS> n=0
    /-libs/fx/004/pos  I32 F_XET
    /-libs/fx/004/name  S32 F_XET
    /-libs/fx/004/type  I32 F_XET
    /-libs/fx/004/flags  P32 F_XET
    /-libs/fx/004/hasdata  I32 F_XET
/-libs/fx/005  <SLIBS> n=0
    /-libs/fx/005/pos  I32 F_XET
    /-libs/fx/005/name  S32 F_XET
    /-libs/fx/005/type  I32 F_XET
    /-libs/fx/005/flags  P32 F_XET
    /-libs/fx/005/hasdata  I32 F_XET
/-libs/fx/006  <SLIBS> n=0
    /-libs/fx/006/pos  I32 F_XET
    /-libs/fx/006/name  S32 F_XET
    /-libs/fx/006/type  I32 F_XET
    /-libs/fx/006/flags  P32 F_XET
    /-libs/fx/006/hasdata  I32 F_XET
/-libs/fx/007  <SLIBS> n=0
    /-libs/fx/007/pos  I32 F_XET
    /-libs/fx/007/name  S32 F_XET
    /-libs/fx/007/type  I32 F_XET
    /-libs/fx/007/flags  P32 F_XET
    /-libs/fx/007/hasdata  I32 F_XET
/-libs/fx/008  <SLIBS> n=0
    /-libs/fx/008/pos  I32 F_XET
    /-libs/fx/008/name  S32 F_XET
    /-libs/fx/008/type  I32 F_XET
    /-libs/fx/008/flags  P32 F_XET
    /-libs/fx/008/hasdata  I32 F_XET
/-libs/fx/009  <SLIBS> n=0
    /-libs/fx/009/pos  I32 F_XET
    /-libs/fx/009/name  S32 F_XET
    /-libs/fx/009/type  I32 F_XET
    /-libs/fx/009/flags  P32 F_XET
    /-libs/fx/009/hasdata  I32 F_XET
/-libs/fx/010  <SLIBS> n=0
    /-libs/fx/010/pos  I32 F_XET
    /-libs/fx/010/name  S32 F_XET
    /-libs/fx/010/type  I32 F_XET
    /-libs/fx/010/flags  P32 F_XET
    /-libs/fx/010/hasdata  I32 F_XET
/-libs/fx/011  <SLIBS> n=0
    /-libs/fx/011/pos  I32 F_XET
    /-libs/fx/011/name  S32 F_XET
    /-libs/fx/011/type  I32 F_XET
    /-libs/fx/011/flags  P32 F_XET
    /-libs/fx/011/hasdata  I32 F_XET
/-libs/fx/012  <SLIBS> n=0
    /-libs/fx/012/pos  I32 F_XET
    /-libs/fx/012/name  S32 F_XET
    /-libs/fx/012/type  I32 F_XET
    /-libs/fx/012/flags  P32 F_XET
    /-libs/fx/012/hasdata  I32 F_XET
/-libs/fx/013  <SLIBS> n=0
    /-libs/fx/013/pos  I32 F_XET
    /-libs/fx/013/name  S32 F_XET
    /-libs/fx/013/type  I32 F_XET
    /-libs/fx/013/flags  P32 F_XET
    /-libs/fx/013/hasdata  I32 F_XET
/-libs/fx/014  <SLIBS> n=0
    /-libs/fx/014/pos  I32 F_XET
    /-libs/fx/014/name  S32 F_XET
    /-libs/fx/014/type  I32 F_XET
    /-libs/fx/014/flags  P32 F_XET
    /-libs/fx/014/hasdata  I32 F_XET
/-libs/fx/015  <SLIBS> n=0
    /-libs/fx/015/pos  I32 F_XET
    /-libs/fx/015/name  S32 F_XET
    /-libs/fx/015/type  I32 F_XET
    /-libs/fx/015/flags  P32 F_XET
    /-libs/fx/015/hasdata  I32 F_XET
/-libs/fx/016  <SLIBS> n=0
    /-libs/fx/016/pos  I32 F_XET
    /-libs/fx/016/name  S32 F_XET
    /-libs/fx/016/type  I32 F_XET
    /-libs/fx/016/flags  P32 F_XET
    /-libs/fx/016/hasdata  I32 F_XET
/-libs/fx/017  <SLIBS> n=0
    /-libs/fx/017/pos  I32 F_XET
    /-libs/fx/017/name  S32 F_XET
    /-libs/fx/017/type  I32 F_XET
    /-libs/fx/017/flags  P32 F_XET
    /-libs/fx/017/hasdata  I32 F_XET
/-libs/fx/018  <SLIBS> n=0
    /-libs/fx/018/pos  I32 F_XET
    /-libs/fx/018/name  S32 F_XET
    /-libs/fx/018/type  I32 F_XET
    /-libs/fx/018/flags  P32 F_XET
    /-libs/fx/018/hasdata  I32 F_XET
/-libs/fx/019  <SLIBS> n=0
    /-libs/fx/019/pos  I32 F_XET
    /-libs/fx/019/name  S32 F_XET
    /-libs/fx/019/type  I32 F_XET
    /-libs/fx/019/flags  P32 F_XET
    /-libs/fx/019/hasdata  I32 F_XET
/-libs/fx/020  <SLIBS> n=0
    /-libs/fx/020/pos  I32 F_XET
    /-libs/fx/020/name  S32 F_XET
    /-libs/fx/020/type  I32 F_XET
    /-libs/fx/020/flags  P32 F_XET
    /-libs/fx/020/hasdata  I32 F_XET
/-libs/fx/021  <SLIBS> n=0
    /-libs/fx/021/pos  I32 F_XET
    /-libs/fx/021/name  S32 F_XET
    /-libs/fx/021/type  I32 F_XET
    /-libs/fx/021/flags  P32 F_XET
    /-libs/fx/021/hasdata  I32 F_XET
/-libs/fx/022  <SLIBS> n=0
    /-libs/fx/022/pos  I32 F_XET
    /-libs/fx/022/name  S32 F_XET
    /-libs/fx/022/type  I32 F_XET
    /-libs/fx/022/flags  P32 F_XET
    /-libs/fx/022/hasdata  I32 F_XET
/-libs/fx/023  <SLIBS> n=0
    /-libs/fx/023/pos  I32 F_XET
    /-libs/fx/023/name  S32 F_XET
    /-libs/fx/023/type  I32 F_XET
    /-libs/fx/023/flags  P32 F_XET
    /-libs/fx/023/hasdata  I32 F_XET
/-libs/fx/024  <SLIBS> n=0
    /-libs/fx/024/pos  I32 F_XET
    /-libs/fx/024/name  S32 F_XET
    /-libs/fx/024/type  I32 F_XET
    /-libs/fx/024/flags  P32 F_XET
    /-libs/fx/024/hasdata  I32 F_XET
/-libs/fx/025  <SLIBS> n=0
    /-libs/fx/025/pos  I32 F_XET
    /-libs/fx/025/name  S32 F_XET
    /-libs/fx/025/type  I32 F_XET
    /-libs/fx/025/flags  P32 F_XET
    /-libs/fx/025/hasdata  I32 F_XET
/-libs/fx/026  <SLIBS> n=0
    /-libs/fx/026/pos  I32 F_XET
    /-libs/fx/026/name  S32 F_XET
    /-libs/fx/026/type  I32 F_XET
    /-libs/fx/026/flags  P32 F_XET
    /-libs/fx/026/hasdata  I32 F_XET
/-libs/fx/027  <SLIBS> n=0
    /-libs/fx/027/pos  I32 F_XET
    /-libs/fx/027/name  S32 F_XET
    /-libs/fx/027/type  I32 F_XET
    /-libs/fx/027/flags  P32 F_XET
    /-libs/fx/027/hasdata  I32 F_XET
/-libs/fx/028  <SLIBS> n=0
    /-libs/fx/028/pos  I32 F_XET
    /-libs/fx/028/name  S32 F_XET
    /-libs/fx/028/type  I32 F_XET
    /-libs/fx/028/flags  P32 F_XET
    /-libs/fx/028/hasdata  I32 F_XET
/-libs/fx/029  <SLIBS> n=0
    /-libs/fx/029/pos  I32 F_XET
    /-libs/fx/029/name  S32 F_XET
    /-libs/fx/029/type  I32 F_XET
    /-libs/fx/029/flags  P32 F_XET
    /-libs/fx/029/hasdata  I32 F_XET
/-libs/fx/030  <SLIBS> n=0
    /-libs/fx/030/pos  I32 F_XET
    /-libs/fx/030/name  S32 F_XET
    /-libs/fx/030/type  I32 F_XET
    /-libs/fx/030/flags  P32 F_XET
    /-libs/fx/030/hasdata  I32 F_XET
/-libs/fx/031  <SLIBS> n=0
    /-libs/fx/031/pos  I32 F_XET
    /-libs/fx/031/name  S32 F_XET
    /-libs/fx/031/type  I32 F_XET
    /-libs/fx/031/flags  P32 F_XET
    /-libs/fx/031/hasdata  I32 F_XET
/-libs/fx/032  <SLIBS> n=0
    /-libs/fx/032/pos  I32 F_XET
    /-libs/fx/032/name  S32 F_XET
    /-libs/fx/032/type  I32 F_XET
    /-libs/fx/032/flags  P32 F_XET
    /-libs/fx/032/hasdata  I32 F_XET
/-libs/fx/033  <SLIBS> n=0
    /-libs/fx/033/pos  I32 F_XET
    /-libs/fx/033/name  S32 F_XET
    /-libs/fx/033/type  I32 F_XET
    /-libs/fx/033/flags  P32 F_XET
    /-libs/fx/033/hasdata  I32 F_XET
/-libs/fx/034  <SLIBS> n=0
    /-libs/fx/034/pos  I32 F_XET
    /-libs/fx/034/name  S32 F_XET
    /-libs/fx/034/type  I32 F_XET
    /-libs/fx/034/flags  P32 F_XET
    /-libs/fx/034/hasdata  I32 F_XET
/-libs/fx/035  <SLIBS> n=0
    /-libs/fx/035/pos  I32 F_XET
    /-libs/fx/035/name  S32 F_XET
    /-libs/fx/035/type  I32 F_XET
    /-libs/fx/035/flags  P32 F_XET
    /-libs/fx/035/hasdata  I32 F_XET
/-libs/fx/036  <SLIBS> n=0
    /-libs/fx/036/pos  I32 F_XET
    /-libs/fx/036/name  S32 F_XET
    /-libs/fx/036/type  I32 F_XET
    /-libs/fx/036/flags  P32 F_XET
    /-libs/fx/036/hasdata  I32 F_XET
/-libs/fx/037  <SLIBS> n=0
    /-libs/fx/037/pos  I32 F_XET
    /-libs/fx/037/name  S32 F_XET
    /-libs/fx/037/type  I32 F_XET
    /-libs/fx/037/flags  P32 F_XET
    /-libs/fx/037/hasdata  I32 F_XET
/-libs/fx/038  <SLIBS> n=0
    /-libs/fx/038/pos  I32 F_XET
    /-libs/fx/038/name  S32 F_XET
    /-libs/fx/038/type  I32 F_XET
    /-libs/fx/038/flags  P32 F_XET
    /-libs/fx/038/hasdata  I32 F_XET
/-libs/fx/039  <SLIBS> n=0
    /-libs/fx/039/pos  I32 F_XET
    /-libs/fx/039/name  S32 F_XET
    /-libs/fx/039/type  I32 F_XET
    /-libs/fx/039/flags  P32 F_XET
    /-libs/fx/039/hasdata  I32 F_XET
/-libs/fx/030  <SLIBS> n=0
    /-libs/fx/040/pos  I32 F_XET
    /-libs/fx/040/name  S32 F_XET
    /-libs/fx/040/type  I32 F_XET
    /-libs/fx/040/flags  P32 F_XET
    /-libs/fx/040/hasdata  I32 F_XET
/-libs/fx/041  <SLIBS> n=0
    /-libs/fx/041/pos  I32 F_XET
    /-libs/fx/041/name  S32 F_XET
    /-libs/fx/041/type  I32 F_XET
    /-libs/fx/041/flags  P32 F_XET
    /-libs/fx/041/hasdata  I32 F_XET
/-libs/fx/042  <SLIBS> n=0
    /-libs/fx/042/pos  I32 F_XET
    /-libs/fx/042/name  S32 F_XET
    /-libs/fx/042/type  I32 F_XET
    /-libs/fx/042/flags  P32 F_XET
    /-libs/fx/042/hasdata  I32 F_XET
/-libs/fx/043  <SLIBS> n=0
    /-libs/fx/043/pos  I32 F_XET
    /-libs/fx/043/name  S32 F_XET
    /-libs/fx/043/type  I32 F_XET
    /-libs/fx/043/flags  P32 F_XET
    /-libs/fx/043/hasdata  I32 F_XET
/-libs/fx/044  <SLIBS> n=0
    /-libs/fx/044/pos  I32 F_XET
    /-libs/fx/044/name  S32 F_XET
    /-libs/fx/044/type  I32 F_XET
    /-libs/fx/044/flags  P32 F_XET
    /-libs/fx/044/hasdata  I32 F_XET
/-libs/fx/045  <SLIBS> n=0
    /-libs/fx/045/pos  I32 F_XET
    /-libs/fx/045/name  S32 F_XET
    /-libs/fx/045/type  I32 F_XET
    /-libs/fx/045/flags  P32 F_XET
    /-libs/fx/045/hasdata  I32 F_XET
/-libs/fx/046  <SLIBS> n=0
    /-libs/fx/046/pos  I32 F_XET
    /-libs/fx/046/name  S32 F_XET
    /-libs/fx/046/type  I32 F_XET
    /-libs/fx/046/flags  P32 F_XET
    /-libs/fx/046/hasdata  I32 F_XET
/-libs/fx/047  <SLIBS> n=0
    /-libs/fx/047/pos  I32 F_XET
    /-libs/fx/047/name  S32 F_XET
    /-libs/fx/047/type  I32 F_XET
    /-libs/fx/047/flags  P32 F_XET
    /-libs/fx/047/hasdata  I32 F_XET
/-libs/fx/048  <SLIBS> n=0
    /-libs/fx/048/pos  I32 F_XET
    /-libs/fx/048/name  S32 F_XET
    /-libs/fx/048/type  I32 F_XET
    /-libs/fx/048/flags  P32 F_XET
    /-libs/fx/048/hasdata  I32 F_XET
/-libs/fx/049  <SLIBS> n=0
    /-libs/fx/049/pos  I32 F_XET
    /-libs/fx/049/name  S32 F_XET
    /-libs/fx/049/type  I32 F_XET
    /-libs/fx/049/flags  P32 F_XET
    /-libs/fx/049/hasdata  I32 F_XET
/-libs/fx/040  <SLIBS> n=0
    /-libs/fx/050/pos  I32 F_XET
    /-libs/fx/050/name  S32 F_XET
    /-libs/fx/050/type  I32 F_XET
    /-libs/fx/050/flags  P32 F_XET
    /-libs/fx/050/hasdata  I32 F_XET
/-libs/fx/051  <SLIBS> n=0
    /-libs/fx/051/pos  I32 F_XET
    /-libs/fx/051/name  S32 F_XET
    /-libs/fx/051/type  I32 F_XET
    /-libs/fx/051/flags  P32 F_XET
    /-libs/fx/051/hasdata  I32 F_XET
/-libs/fx/052  <SLIBS> n=0
    /-libs/fx/052/pos  I32 F_XET
    /-libs/fx/052/name  S32 F_XET
    /-libs/fx/052/type  I32 F_XET
    /-libs/fx/052/flags  P32 F_XET
    /-libs/fx/052/hasdata  I32 F_XET
/-libs/fx/053  <SLIBS> n=0
    /-libs/fx/053/pos  I32 F_XET
    /-libs/fx/053/name  S32 F_XET
    /-libs/fx/053/type  I32 F_XET
    /-libs/fx/053/flags  P32 F_XET
    /-libs/fx/053/hasdata  I32 F_XET
/-libs/fx/054  <SLIBS> n=0
    /-libs/fx/054/pos  I32 F_XET
    /-libs/fx/054/name  S32 F_XET
    /-libs/fx/054/type  I32 F_XET
    /-libs/fx/054/flags  P32 F_XET
    /-libs/fx/054/hasdata  I32 F_XET
/-libs/fx/055  <SLIBS> n=0
    /-libs/fx/055/pos  I32 F_XET
    /-libs/fx/055/name  S32 F_XET
    /-libs/fx/055/type  I32 F_XET
    /-libs/fx/055/flags  P32 F_XET
    /-libs/fx/055/hasdata  I32 F_XET
/-libs/fx/056  <SLIBS> n=0
    /-libs/fx/056/pos  I32 F_XET
    /-libs/fx/056/name  S32 F_XET
    /-libs/fx/056/type  I32 F_XET
    /-libs/fx/056/flags  P32 F_XET
    /-libs/fx/056/hasdata  I32 F_XET
/-libs/fx/057  <SLIBS> n=0
    /-libs/fx/057/pos  I32 F_XET
    /-libs/fx/057/name  S32 F_XET
    /-libs/fx/057/type  I32 F_XET
    /-libs/fx/057/flags  P32 F_XET
    /-libs/fx/057/hasdata  I32 F_XET
/-libs/fx/058  <SLIBS> n=0
    /-libs/fx/058/pos  I32 F_XET
    /-libs/fx/058/name  S32 F_XET
    /-libs/fx/058/type  I32 F_XET
    /-libs/fx/058/flags  P32 F_XET
    /-libs/fx/058/hasdata  I32 F_XET
/-libs/fx/059  <SLIBS> n=0
    /-libs/fx/059/pos  I32 F_XET
    /-libs/fx/059/name  S32 F_XET
    /-libs/fx/059/type  I32 F_XET
    /-libs/fx/059/flags  P32 F_XET
    /-libs/fx/059/hasdata  I32 F_XET
/-libs/fx/060  <SLIBS> n=0
    /-libs/fx/060/pos  I32 F_XET
    /-libs/fx/060/name  S32 F_XET
    /-libs/fx/060/type  I32 F_XET
    /-libs/fx/060/flags  P32 F_XET
    /-libs/fx/060/hasdata  I32 F_XET
/-libs/fx/061  <SLIBS> n=0
    /-libs/fx/061/pos  I32 F_XET
    /-libs/fx/061/name  S32 F_XET
    /-libs/fx/061/type  I32 F_XET
    /-libs/fx/061/flags  P32 F_XET
    /-libs/fx/061/hasdata  I32 F_XET
/-libs/fx/062  <SLIBS> n=0
    /-libs/fx/062/pos  I32 F_XET
    /-libs/fx/062/name  S32 F_XET
    /-libs/fx/062/type  I32 F_XET
    /-libs/fx/062/flags  P32 F_XET
    /-libs/fx/062/hasdata  I32 F_XET
/-libs/fx/063  <SLIBS> n=0
    /-libs/fx/063/pos  I32 F_XET
    /-libs/fx/063/name  S32 F_XET
    /-libs/fx/063/type  I32 F_XET
    /-libs/fx/063/flags  P32 F_XET
    /-libs/fx/063/hasdata  I32 F_XET
/-libs/fx/064  <SLIBS> n=0
    /-libs/fx/064/pos  I32 F_XET
    /-libs/fx/064/name  S32 F_XET
    /-libs/fx/064/type  I32 F_XET
    /-libs/fx/064/flags  P32 F_XET
    /-libs/fx/064/hasdata  I32 F_XET
/-libs/fx/065  <SLIBS> n=0
    /-libs/fx/065/pos  I32 F_XET
    /-libs/fx/065/name  S32 F_XET
    /-libs/fx/065/type  I32 F_XET
    /-libs/fx/065/flags  P32 F_XET
    /-libs/fx/065/hasdata  I32 F_XET
/-libs/fx/066  <SLIBS> n=0
    /-libs/fx/066/pos  I32 F_XET
    /-libs/fx/066/name  S32 F_XET
    /-libs/fx/066/type  I32 F_XET
    /-libs/fx/066/flags  P32 F_XET
    /-libs/fx/066/hasdata  I32 F_XET
/-libs/fx/067  <SLIBS> n=0
    /-libs/fx/067/pos  I32 F_XET
    /-libs/fx/067/name  S32 F_XET
    /-libs/fx/067/type  I32 F_XET
    /-libs/fx/067/flags  P32 F_XET
    /-libs/fx/067/hasdata  I32 F_XET
/-libs/fx/068  <SLIBS> n=0
    /-libs/fx/068/pos  I32 F_XET
    /-libs/fx/068/name  S32 F_XET
    /-libs/fx/068/type  I32 F_XET
    /-libs/fx/068/flags  P32 F_XET
    /-libs/fx/068/hasdata  I32 F_XET
/-libs/fx/069  <SLIBS> n=0
    /-libs/fx/069/pos  I32 F_XET
    /-libs/fx/069/name  S32 F_XET
    /-libs/fx/069/type  I32 F_XET
    /-libs/fx/069/flags  P32 F_XET
    /-libs/fx/069/hasdata  I32 F_XET
/-libs/fx/070  <SLIBS> n=0
    /-libs/fx/070/pos  I32 F_XET
    /-libs/fx/070/name  S32 F_XET
    /-libs/fx/070/type  I32 F_XET
    /-libs/fx/070/flags  P32 F_XET
    /-libs/fx/070/hasdata  I32 F_XET
/-libs/fx/071  <SLIBS> n=0
    /-libs/fx/071/pos  I32 F_XET
    /-libs/fx/071/name  S32 F_XET
    /-libs/fx/071/type  I32 F_XET
    /-libs/fx/071/flags  P32 F_XET
    /-libs/fx/071/hasdata  I32 F_XET
/-libs/fx/072  <SLIBS> n=0
    /-libs/fx/072/pos  I32 F_XET
    /-libs/fx/072/name  S32 F_XET
    /-libs/fx/072/type  I32 F_XET
    /-libs/fx/072/flags  P32 F_XET
    /-libs/fx/072/hasdata  I32 F_XET
/-libs/fx/073  <SLIBS> n=0
    /-libs/fx/073/pos  I32 F_XET
    /-libs/fx/073/name  S32 F_XET
    /-libs/fx/073/type  I32 F_XET
    /-libs/fx/073/flags  P32 F_XET
    /-libs/fx/073/hasdata  I32 F_XET
/-libs/fx/074  <SLIBS> n=0
    /-libs/fx/074/pos  I32 F_XET
    /-libs/fx/074/name  S32 F_XET
    /-libs/fx/074/type  I32 F_XET
    /-libs/fx/074/flags  P32 F_XET
    /-libs/fx/074/hasdata  I32 F_XET
/-libs/fx/075  <SLIBS> n=0
    /-libs/fx/075/pos  I32 F_XET
    /-libs/fx/075/name  S32 F_XET
    /-libs/fx/075/type  I32 F_XET
    /-libs/fx/075/flags  P32 F_XET
    /-libs/fx/075/hasdata  I32 F_XET
/-libs/fx/076  <SLIBS> n=0
    /-libs/fx/076/pos  I32 F_XET
    /-libs/fx/076/name  S32 F_XET
    /-libs/fx/076/type  I32 F_XET
    /-libs/fx/076/flags  P32 F_XET
    /-libs/fx/076/hasdata  I32 F_XET
/-libs/fx/077  <SLIBS> n=0
    /-libs/fx/077/pos  I32 F_XET
    /-libs/fx/077/name  S32 F_XET
    /-libs/fx/077/type  I32 F_XET
    /-libs/fx/077/flags  P32 F_XET
    /-libs/fx/077/hasdata  I32 F_XET
/-libs/fx/078  <SLIBS> n=0
    /-libs/fx/078/pos  I32 F_XET
    /-libs/fx/078/name  S32 F_XET
    /-libs/fx/078/type  I32 F_XET
    /-libs/fx/078/flags  P32 F_XET
    /-libs/fx/078/hasdata  I32 F_XET
/-libs/fx/079  <SLIBS> n=0
    /-libs/fx/079/pos  I32 F_XET
    /-libs/fx/079/name  S32 F_XET
    /-libs/fx/079/type  I32 F_XET
    /-libs/fx/079/flags  P32 F_XET
    /-libs/fx/079/hasdata  I32 F_XET
/-libs/fx/080  <SLIBS> n=0
    /-libs/fx/080/pos  I32 F_XET
    /-libs/fx/080/name  S32 F_XET
    /-libs/fx/080/type  I32 F_XET
    /-libs/fx/080/flags  P32 F_XET
    /-libs/fx/080/hasdata  I32 F_XET
/-libs/fx/081  <SLIBS> n=0
    /-libs/fx/081/pos  I32 F_XET
    /-libs/fx/081/name  S32 F_XET
    /-libs/fx/081/type  I32 F_XET
    /-libs/fx/081/flags  P32 F_XET
    /-libs/fx/081/hasdata  I32 F_XET
/-libs/fx/082  <SLIBS> n=0
    /-libs/fx/082/pos  I32 F_XET
    /-libs/fx/082/name  S32 F_XET
    /-libs/fx/082/type  I32 F_XET
    /-libs/fx/082/flags  P32 F_XET
    /-libs/fx/082/hasdata  I32 F_XET
/-libs/fx/083  <SLIBS> n=0
    /-libs/fx/083/pos  I32 F_XET
    /-libs/fx/083/name  S32 F_XET
    /-libs/fx/083/type  I32 F_XET
    /-libs/fx/083/flags  P32 F_XET
    /-libs/fx/083/hasdata  I32 F_XET
/-libs/fx/084  <SLIBS> n=0
    /-libs/fx/084/pos  I32 F_XET
    /-libs/fx/084/name  S32 F_XET
    /-libs/fx/084/type  I32 F_XET
    /-libs/fx/084/flags  P32 F_XET
    /-libs/fx/084/hasdata  I32 F_XET
/-libs/fx/085  <SLIBS> n=0
    /-libs/fx/085/pos  I32 F_XET
    /-libs/fx/085/name  S32 F_XET
    /-libs/fx/085/type  I32 F_XET
    /-libs/fx/085/flags  P32 F_XET
    /-libs/fx/085/hasdata  I32 F_XET
/-libs/fx/086  <SLIBS> n=0
    /-libs/fx/086/pos  I32 F_XET
    /-libs/fx/086/name  S32 F_XET
    /-libs/fx/086/type  I32 F_XET
    /-libs/fx/086/flags  P32 F_XET
    /-libs/fx/086/hasdata  I32 F_XET
/-libs/fx/087  <SLIBS> n=0
    /-libs/fx/087/pos  I32 F_XET
    /-libs/fx/087/name  S32 F_XET
    /-libs/fx/087/type  I32 F_XET
    /-libs/fx/087/flags  P32 F_XET
    /-libs/fx/087/hasdata  I32 F_XET
/-libs/fx/088  <SLIBS> n=0
    /-libs/fx/088/pos  I32 F_XET
    /-libs/fx/088/name  S32 F_XET
    /-libs/fx/088/type  I32 F_XET
    /-libs/fx/088/flags  P32 F_XET
    /-libs/fx/088/hasdata  I32 F_XET
/-libs/fx/089  <SLIBS> n=0
    /-libs/fx/089/pos  I32 F_XET
    /-libs/fx/089/name  S32 F_XET
    /-libs/fx/089/type  I32 F_XET
    /-libs/fx/089/flags  P32 F_XET
    /-libs/fx/089/hasdata  I32 F_XET
/-libs/fx/090  <SLIBS> n=0
    /-libs/fx/090/pos  I32 F_XET
    /-libs/fx/090/name  S32 F_XET
    /-libs/fx/090/type  I32 F_XET
    /-libs/fx/090/flags  P32 F_XET
    /-libs/fx/090/hasdata  I32 F_XET
/-libs/fx/091  <SLIBS> n=0
    /-libs/fx/091/pos  I32 F_XET
    /-libs/fx/091/name  S32 F_XET
    /-libs/fx/091/type  I32 F_XET
    /-libs/fx/091/flags  P32 F_XET
    /-libs/fx/091/hasdata  I32 F_XET
/-libs/fx/092  <SLIBS> n=0
    /-libs/fx/092/pos  I32 F_XET
    /-libs/fx/092/name  S32 F_XET
    /-libs/fx/092/type  I32 F_XET
    /-libs/fx/092/flags  P32 F_XET
    /-libs/fx/092/hasdata  I32 F_XET
/-libs/fx/093  <SLIBS> n=0
    /-libs/fx/093/pos  I32 F_XET
    /-libs/fx/093/name  S32 F_XET
    /-libs/fx/093/type  I32 F_XET
    /-libs/fx/093/flags  P32 F_XET
    /-libs/fx/093/hasdata  I32 F_XET
/-libs/fx/094  <SLIBS> n=0
    /-libs/fx/094/pos  I32 F_XET
    /-libs/fx/094/name  S32 F_XET
    /-libs/fx/094/type  I32 F_XET
    /-libs/fx/094/flags  P32 F_XET
    /-libs/fx/094/hasdata  I32 F_XET
/-libs/fx/095  <SLIBS> n=0
    /-libs/fx/095/pos  I32 F_XET
    /-libs/fx/095/name  S32 F_XET
    /-libs/fx/095/type  I32 F_XET
    /-libs/fx/095/flags  P32 F_XET
    /-libs/fx/095/hasdata  I32 F_XET
/-libs/fx/096  <SLIBS> n=0
    /-libs/fx/096/pos  I32 F_XET
    /-libs/fx/096/name  S32 F_XET
    /-libs/fx/096/type  I32 F_XET
    /-libs/fx/096/flags  P32 F_XET
    /-libs/fx/096/hasdata  I32 F_XET
/-libs/fx/097  <SLIBS> n=0
    /-libs/fx/097/pos  I32 F_XET
    /-libs/fx/097/name  S32 F_XET
    /-libs/fx/097/type  I32 F_XET
    /-libs/fx/097/flags  P32 F_XET
    /-libs/fx/097/hasdata  I32 F_XET
/-libs/fx/098  <SLIBS> n=0
    /-libs/fx/098/pos  I32 F_XET
    /-libs/fx/098/name  S32 F_XET
    /-libs/fx/098/type  I32 F_XET
    /-libs/fx/098/flags  P32 F_XET
    /-libs/fx/098/hasdata  I32 F_XET
/-libs/fx/099  <SLIBS> n=0
    /-libs/fx/099/pos  I32 F_XET
    /-libs/fx/099/name  S32 F_XET
    /-libs/fx/099/type  I32 F_XET
    /-libs/fx/099/flags  P32 F_XET
    /-libs/fx/099/hasdata  I32 F_XET
/-libs/fx/100  <SLIBS> n=0
    /-libs/fx/100/pos  I32 F_XET
    /-libs/fx/100/name  S32 F_XET
    /-libs/fx/100/type  I32 F_XET
    /-libs/fx/100/flags  P32 F_XET
    /-libs/fx/100/hasdata  I32 F_XET
```

