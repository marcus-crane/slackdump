// Code generated by "stringer -type=datafmt -trimprefix=F"; DO NOT EDIT.

package convertcmd

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Fdump-0]
	_ = x[Fexport-1]
	_ = x[Fchunk-2]
}

const _datafmt_name = "dumpexportchunk"

var _datafmt_index = [...]uint8{0, 4, 10, 15}

func (i datafmt) String() string {
	if i >= datafmt(len(_datafmt_index)-1) {
		return "datafmt(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _datafmt_name[_datafmt_index[i]:_datafmt_index[i+1]]
}
