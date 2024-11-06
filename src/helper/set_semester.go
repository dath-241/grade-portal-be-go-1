package helper

import (
	"strconv"
	"time"
)

type Semester struct {
	CUREENT string
	NEXT    string
	PREV    string
}

func Set_semester() Semester {
	t := time.Now()
	year := t.Year()
	month := t.Month()
	a := year - 2000
	current := "HK"
	next := "HK"
	prev := "HK"
	switch {
	case month >= 9 && month <= 12:
		current += strconv.Itoa(a) + "1"
		prev += strconv.Itoa(a-1) + "3"
		next += strconv.Itoa(a) + "2"
	case month >= 1 && month <= 4:
		current += strconv.Itoa(a-1) + "2"
		prev += strconv.Itoa(a-1) + "1"
		next += strconv.Itoa(a-1) + "3"
	case month >= 5 && month <= 8:
		current += strconv.Itoa(a-1) + "3"
		prev += strconv.Itoa(a-1) + "2"
		next += strconv.Itoa(a) + "1"
	}
	return Semester{current, next, prev}
}
