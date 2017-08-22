package main

import (
	"fmt"
	"testing"
)

func TestDial(t *testing.T) {

	expire, err := Dial("google.com", "google.com", "443", false)

	if err != nil {
		fmt.Println(err)
		t.Errorf("google.com's cert should not have error")
	}

	fmt.Println(expire)
}
