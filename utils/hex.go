// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hex implements hexadecimal encoding and decoding.
// This is a modified version for custom use

package utils

const (
	hexTable = "0123456789abcdef"
)

var reverseHexTable [256]byte

func init() {
	for i := range reverseHexTable {
		reverseHexTable[i] = 0xff
	}
	for i, c := range hexTable {
		reverseHexTable[c] = byte(i)
	}
}

func Encode(b []byte) string {
	buf := make([]byte, len(b)*2+2)
	buf[0], buf[1] = '0', 'x'
	for i, v := range b {
		buf[2+i*2], buf[2+i*2+1] = hexTable[v/16], hexTable[v%16]
	}
	return string(buf[:])
}

func Decode(s string) ([]byte, error) {
	if len(s) == 0 {
		return nil, ErrEmptyString.Error()
	}
	if s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
		s = s[2:]
	}
	buf := []byte(s)
	for i := 0; i < len(buf)/2; i++ {
		p, q := buf[i*2], buf[i*2+1]
		a, b := reverseHexTable[p], reverseHexTable[q]
		if a > 15 {
			return buf[:len(buf)/2-1], ErrUnknownChar.Errorf(p)
		}
		if b > 15 {
			return buf[:len(buf)/2-1], ErrUnknownChar.Errorf(q)
		}
		buf[i] = (a * 16) | b
	}
	if len(buf)%2 == 1 {
		if reverseHexTable[buf[len(buf)-1]] > 15 {
			return buf[:len(buf)/2], ErrUnknownChar.Errorf(buf[len(buf)-1])
		}
		return buf[:len(buf)/2], ErrOddLength.Error()
	}
	return buf[:len(buf)/2], nil
}
