// Code generated by "stringer -type=LineJoin"; DO NOT EDIT.

package gi

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

const _LineJoin_name = "LineJoinMiterLineJoinMiterClipLineJoinRoundLineJoinBevelLineJoinArcsLineJoinArcsClipLineJoinN"

var _LineJoin_index = [...]uint8{0, 13, 30, 43, 56, 68, 84, 93}

func (i LineJoin) String() string {
	if i < 0 || i >= LineJoin(len(_LineJoin_index)-1) {
		return "LineJoin(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LineJoin_name[_LineJoin_index[i]:_LineJoin_index[i+1]]
}

func (i *LineJoin) FromString(s string) error {
	for j := 0; j < len(_LineJoin_index)-1; j++ {
		if s == _LineJoin_name[_LineJoin_index[j]:_LineJoin_index[j+1]] {
			*i = LineJoin(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: LineJoin")
}