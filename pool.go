package bytespool

import (
	. "github.com/itsabgr/go-q"
)

var bytes Q

//MinBytesLen is minimum bytes len to add to pool
var MinBytesLen = 16

//Put adds b to pull
func Put(b []byte) {
	if len(b) <= MinBytesLen {
		return
	}
	bytes.Push(b)
}

//Get return bytes[:max]
func Get(max int) []byte {
	IByte, found := bytes.Pull()
	if !found {
		return make([]byte, max)
	}
	b := IByte.([]byte)
	if len(b) > max {
		Put(b[max:])
		return b[:max]
	}
	return b
}
