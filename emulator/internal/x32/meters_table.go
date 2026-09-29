package x32

// Xmeters is the /meters command table (Xmeters[] in X32.c). Entries are F_GET
// only: the meter frames are built by prepMeter(), not by functParams().
var Xmeters = Table{
	{"/meters/0", TI32, FGET, 0, nil, 0, 0, nil},
	{"/meters/1", TI32, FGET, 0, nil, 0, 0, nil},
	{"/meters/2", TI32, FGET, 0, nil, 0, 0, nil},
	{"/meters/3", TI32, FGET, 0, nil, 0, 0, nil},
	{"/meters/4", TI32, FGET, 0, nil, 0, 0, nil},
	{"/meters/5", TI32, FGET, 0, nil, 0, 0, nil},
	{"/meters/6", TI32, FGET, 0, nil, 0, 0, nil},
	{"/meters/7", TI32, FGET, 0, nil, 0, 0, nil},
	{"/meters/8", TI32, FGET, 0, nil, 0, 0, nil},
	{"/meters/9", TI32, FGET, 0, nil, 0, 0, nil},
	{"/meters/10", TI32, FGET, 0, nil, 0, 0, nil},
	{"/meters/11", TI32, FGET, 0, nil, 0, 0, nil},
	{"/meters/12", TI32, FGET, 0, nil, 0, 0, nil},
	{"/meters/13", TI32, FGET, 0, nil, 0, 0, nil},
	{"/meters/14", TI32, FGET, 0, nil, 0, 0, nil},
	{"/meters/15", TI32, FGET, 0, nil, 0, 0, nil},
	{"/meters/16", TI32, FGET, 0, nil, 0, 0, nil},
}
