package service

import "testing"

func TestXXX(t *testing.T) {
	for i := uint64(0); i < 3; i++ {
		t.Logf("\nTestCase id=%+v\n", i)
		if err := xxx(i); err != nil {
			t.Logf("TestXXX Err=%+v", err)
		}
	}
}
