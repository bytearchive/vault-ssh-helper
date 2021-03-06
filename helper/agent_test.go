package helper

import (
	"net"
	"testing"
)

func TestBelongsToCIDR(t *testing.T) {
	t.Parallel()
	testIP := net.ParseIP("10.50.100.101")
	testCIDR := "0.0.0.0/0"
	belongs, err := belongsToCIDR(testIP, testCIDR)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if !belongs {
		t.Fatalf("bad: expected:true, actual:false")
	}

	testCIDR = "192.168.0.1/16"
	belongs, err = belongsToCIDR(testIP, testCIDR)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if belongs {
		t.Fatalf("bad: expected:false, actual:true")
	}

	testCIDR = "invalid"
	_, err = belongsToCIDR(testIP, testCIDR)
	if err == nil {
		t.Fatalf("expected error")
	}
}
