package bytespool

import (
	"github.com/itsabgr/atomic2"
	. "github.com/itsabgr/go-q"
)

var pool Q

//MinBytesLen is minimum bytes len to add to pool
var MinBytesLen uint = 16

//MaxBytesLen is maximum bytes len to add to pool
var MaxBytesLen uint = 536870912 //512MB
//MaxPoolLen is maximum pool len
var MaxPoolLen uint = 4294967296 //4GB

//Push adds b to pull
func Push(b []byte) {
	if uint(len(b)) <= MinBytesLen {
		return
	}
	if uint(len(b)) > MaxBytesLen {
		return
	}
	if uint(poolLen.Get()+uintptr(len(b))) > MaxPoolLen {
		return
	}
	poolLen.Add(uintptr(len(b)))
	pool.Push(b)
}

var poolLen = atomic2.Uintptr(0)

//Len returns total pool bytes len
func Len() uint {
	return uint(poolLen.Get())
}

//Pull return bytes[:max]
func Pull(max uint) []byte {
	IByte, found := pool.Pull()
	if !found {
		return make([]byte, max)
	}
	b := IByte.([]byte)
	poolLen.Sub(uintptr(len(b)))
	if uint(len(b)) > max {
		Push(b[max:])
		return b[:max]
	}
	return b
}
