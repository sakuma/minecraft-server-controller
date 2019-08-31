package main

import (
	"testing"
	"time"
)

func TestIsHoliday(t *testing.T)  {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	var tt time.Time
	var got bool

	tt = time.Date(2019, 9, 01, 5, 59, 59, 0, jst)
	got = isHoliday(tt)
	if got != true {
		t.Fatal("got: false, should be true")
	}

	tt = time.Date(2019, 9, 2, 8, 59, 59, 0, jst)
	got = isHoliday(tt)
	if got != false {
		t.Fatal("got: true, should be false")
	}
}
func TestIsActive(t *testing.T) {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	var tt time.Time
	var got bool

	tt = time.Date(2018, 5, 19, 5, 59, 59, 0, jst)
	got = isActive(tt)
	if got != false {
		t.Fatal("got: true, should be false")
	}

	tt = time.Date(2018, 5, 19, 6, 00, 00, 0, jst)
	got = isActive(tt)
	if got != true {
		t.Fatal("got: false, should be true")
	}

	tt = time.Date(2018, 5, 19, 21, 59, 59, 0, jst)
	got = isActive(tt)
	if got != true {
		t.Fatal("got: false, should be true")
	}

	tt = time.Date(2018, 5, 19, 23, 59, 59, 0, jst)
	got = isActive(tt)
	if got != false {
		t.Fatal("got: false, should be true")
	}

	tt = time.Date(2018, 5, 20, 00, 00, 00, 0, jst)
	got = isActive(tt)
	if got != false {
		t.Fatal("got: false, should be true")
	}
}
