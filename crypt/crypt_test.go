package crypt

import (
	"crypto/rand"
	"reflect"
	"testing"
)

func TestEncrypt(t *testing.T) {

	//32文字
	key := "t0JozH+6Ev7/msBIcocYllPEakZPL0UD"

	for i := 0; i < 10; i++ {
		buf := make([]byte, 1000000)

		if _, err := rand.Read(buf); err != nil {
			panic(err)
		}

		enc := Encrypt(buf, key)
		dec := Decrypt(enc, key)

		if !reflect.DeepEqual(buf, dec) {
			t.Error("Must be equal")
		}
	}
}
