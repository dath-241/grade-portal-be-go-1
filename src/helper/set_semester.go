package helper

import (
	"strconv"
	"time"
)

func Set_semester(t time.Time) string {
	year := t.Year()
	month := t.Month()
	var Semester string
	switch {
	case month >= 9 && month <= 12:
		Semester = "HK" + strconv.Itoa(year-2000) + "1"
	case month >= 1 && month <= 4:
		Semester = "HK" + strconv.Itoa(year-2001) + "2"
	case month >= 5 && month <= 8:
		Semester = "HK" + strconv.Itoa(year-2001) + "3"
	}

	return Semester

}
