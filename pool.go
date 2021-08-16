package bytespool

import (
	. "github.com/itsabgr/go-q"
)

var bytes Q

var MinBytesLen = 16
func Put(b []byte) {
	if len(b) <= MinBytesLen{
		return
	}
	bytes.Push(b)
}
func Get(max int)  []byte{
	IByte,found := bytes.Pull()
	if !found{
		return make([]byte,max)
	}
	b := IByte.([]byte)
	if len(b) > max {
		Put(b[max:])
		return b[:max]
	}
	return b
}