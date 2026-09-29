package bridge

// stripState is the cached state of one REAPER track as it is (or was) shown on
// the desk. The bridge needs it because the console only ever shows one bank of
// channels: without a cache, switching away from a bank and back would lose the
// fader positions of the tracks that are no longer visible.
type stripState struct {
	fader float32
	pan   float32
	mute  bool
	solo  bool
	name  string
	sends [MaxSends]float32
	// vu is the last meter level REAPER reported for the track (the loudest side
	// when it sends /vu/L and /vu/R). It is only kept to restore the meters of a
	// bank the desk switches to, so that the bars do not stay at the level of the
	// tracks that were shown before.
	vu    float32
	vuL   float32 // /track/N/vu/L
	vuR   float32 // /track/N/vu/R
	valid bool
}

// stateCache remembers the desk-visible state of every track it has seen, keyed
// by absolute REAPER track number.
type stateCache struct {
	tracks map[int]*stripState
}

func newStateCache() *stateCache {
	return &stateCache{tracks: make(map[int]*stripState)}
}

// get returns the (created on demand) state of a track.
func (c *stateCache) get(track int) *stripState {
	s := c.tracks[track]
	if s == nil {
		s = &stripState{}
		c.tracks[track] = s
	}
	return s
}

// peek returns the state of a track, nil when the track was never seen.
func (c *stateCache) peek(track int) *stripState {
	return c.tracks[track]
}

// Note that the cache keeps stale entries for tracks that were deleted in
// REAPER; they are only ever read to restore the desk, so a stale value is
// harmless (REAPER corrects it as soon as the track reports again).
func (c *stateCache) forget(track int) {
	delete(c.tracks, track)
}
