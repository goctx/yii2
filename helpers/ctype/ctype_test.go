package ctype

import "testing"

func TestAlNum(t *testing.T) {
	if !AlNum("a1") {
		t.Error("errror")
		return
	}
	if AlNum("a1-") {
		t.Error("error2")
		return
	}
}
