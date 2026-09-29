package bridge

import (
	"testing"

	"x32emu/osc"
)

// recordingSink collects what the bridge sends to the console meters.
type recordingSink struct {
	strips []string  // "section/number" of every call
	levels []float32 // level of every call
	master [2]float32
	mmade  int
}

func (r *recordingSink) sink() MeterSink {
	return MeterSink{
		Track: func(section string, number int, v float32) {
			r.strips = append(r.strips, section+"/"+itoa(number))
			r.levels = append(r.levels, v)
		},
		Master: func(left, right float32) {
			r.master[0], r.master[1] = left, right
			r.mmade++
		},
	}
}

func itoa(i int) string {
	if i < 10 {
		return string(rune('0' + i))
	}
	return string(rune('0'+i/10)) + string(rune('0'+i%10))
}

// REAPER's track meters are resolved to the console strip the track is shown on.
func TestMeterSinkTrack(t *testing.T) {
	rec := &recordingSink{}
	b := newTestBridge(t, nil)
	b.SetMeters(rec.sink())

	b.fromReaper(osc.Build("/track/3/vu", "f", osc.Arg{Tag: 'f', F: 0.75}))
	if len(rec.strips) != 1 || rec.strips[0] != "ch/3" || rec.levels[0] != 0.75 {
		t.Fatalf("sink calls = %v %v", rec.strips, rec.levels)
	}
	// the L/R pair of the same track becomes one bar: the loudest side wins
	b.fromReaper(osc.Build("/track/3/vu/L", "f", osc.Arg{Tag: 'f', F: 0.2}))
	b.fromReaper(osc.Build("/track/3/vu/R", "f", osc.Arg{Tag: 'f', F: 0.6}))
	if len(rec.strips) != 3 || rec.levels[2] != 0.6 {
		t.Fatalf("levels = %v, want the last one to be 0.6", rec.levels)
	}
	// ... and the mono value of the next bundle wins again
	b.fromReaper(osc.Build("/track/3/vu", "f", osc.Arg{Tag: 'f', F: 0.1}))
	if rec.levels[len(rec.levels)-1] != 0.1 {
		t.Fatalf("levels = %v", rec.levels)
	}
	// a track outside the mapped ranges is cached but not fed
	b.fromReaper(osc.Build("/track/99/vu", "f", osc.Arg{Tag: 'f', F: 0.9}))
	if len(rec.strips) != 4 {
		t.Fatalf("an unmapped track was fed: %v", rec.strips)
	}
	if b.state.peek(99) == nil {
		t.Fatal("the level of an unmapped track must still be cached")
	}
}

// The aux inputs, effect returns and buses of the desk have their own numbers.
func TestMeterSinkSections(t *testing.T) {
	rec := &recordingSink{}
	b := newTestBridge(t, func(c *Config) {
		c.AuxMin, c.AuxMax = 33, 40
		c.FxrMin, c.FxrMax = 41, 48
		c.BusMin, c.BusMax = 49, 64
	})
	b.SetMeters(rec.sink())
	for track, want := range map[int]string{33: "auxin/1", 40: "auxin/8", 41: "fxrtn/1", 64: "bus/16"} {
		b.fromReaper(osc.Build("/track/"+itoa(track)+"/vu", "f", osc.Arg{Tag: 'f', F: 0.5}))
		if got := rec.strips[len(rec.strips)-1]; got != want {
			t.Fatalf("track %d -> %s, want %s", track, got, want)
		}
	}
}

// The main bus meters are fed as two levels.
func TestMeterSinkMaster(t *testing.T) {
	rec := &recordingSink{}
	b := newTestBridge(t, nil)
	b.SetMeters(rec.sink())
	b.fromReaper(osc.Build("/master/vu/L", "f", osc.Arg{Tag: 'f', F: 0.25}))
	b.fromReaper(osc.Build("/master/vu/R", "f", osc.Arg{Tag: 'f', F: 0.5}))
	if rec.mmade != 2 || rec.master[0] != 0.25 || rec.master[1] != 0.5 {
		t.Fatalf("master = %v (%d calls)", rec.master, rec.mmade)
	}
	// the mono value feeds both channels
	b.fromReaper(osc.Build("/master/vu", "f", osc.Arg{Tag: 'f', F: 0.75}))
	if rec.master[0] != 0.75 || rec.master[1] != 0.75 {
		t.Fatalf("master = %v", rec.master)
	}
}

// Switching banks must refresh the meters of the new bank: the levels of the old
// bank would otherwise stay on screen until REAPER reports again.
func TestMeterSinkBankRecall(t *testing.T) {
	rec := &recordingSink{}
	b := newTestBridge(t, func(c *Config) { c.TrkMin, c.TrkMax = 1, 64 })
	b.SetMeters(rec.sink())
	b.fromReaper(osc.Build("/track/40/vu", "f", osc.Arg{Tag: 'f', F: 0.4}))
	rec.strips = nil
	b.cfg.BankOffset = 1
	b.recallBank()
	found := false
	for i, s := range rec.strips {
		if s == "ch/8" && rec.levels[i] == 0.4 {
			found = true
		}
	}
	if !found {
		t.Fatalf("the bank recall did not restore the cached level: %v %v", rec.strips, rec.levels)
	}
}
