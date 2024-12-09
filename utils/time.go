package utils

import (
	"time"
)

type AbsoluteTime int64

// nanotime implements for generate nanotime int64
func nanotime() int64                                     { return time.Now().UnixNano() }
func Now() AbsoluteTime                                   { return AbsoluteTime(nanotime()) }
func (at AbsoluteTime) Add(d time.Duration) AbsoluteTime  { return at + AbsoluteTime(d) }
func (at AbsoluteTime) Sub(t2 AbsoluteTime) time.Duration { return time.Duration(at - t2) }

type System struct{}

func (s System) Now() AbsoluteTime {
	return Now()
}

func (s System) Sleep(d time.Duration) {
	time.Sleep(d)
}
