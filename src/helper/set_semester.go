package helper

import (
	"strconv"
	"time"
)

func Set_semester(t time.Time) string {
	year := t.Year()
	month := t.Month()
	a := year - 2000
	Semester := "HK" + strconv.Itoa(a)
	switch {
	case month >= 9 && month <= 12:
		Semester += "1"
	case month >= 1 && month <= 4:
		Semester += "2"
	case month >= 5 && month <= 8:
		Semester += "3"
	}

	return Semester

}
