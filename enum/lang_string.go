// Code generated by "stringer -type Lang ./enum/lang.go"; DO NOT EDIT.

package enum

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Go-1]
	_ = x[Python-2]
	_ = x[Java-3]
	_ = x[PHP-4]
}

const _Lang_name = "GoPythonJavaPHP"

var _Lang_index = [...]uint8{0, 2, 8, 12, 15}

func (i Lang) String() string {
	i -= 1
	if i < 0 || i >= Lang(len(_Lang_index)-1) {
		return "Lang(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Lang_name[_Lang_index[i]:_Lang_index[i+1]]
}
