package helper

import (
	"time"
	"strconv"
)

func EnglishTimeConvertToIndonesia(t time.Time, format string) string {
	var DayName string
	var Day 	string
	var Month 	string
	var Year 	string
	var HM 		string

	loc, _ := time.LoadLocation("Asia/Jakarta")
	t 		= t.In(loc)

	//Day Name in String
	switch d:= t.Weekday().String(); d {
		case "Monday":
			DayName = "Senin"
			break
		case "Tuesday":
			DayName = "Selasa"
			break
		case "Wednesday":
			DayName = "Rabu"
			break
		case "Thursday":
			DayName = "Kamis"
			break
		case "Friday":
			DayName = "Jumat"
			break
		case "Saturday":
			DayName = "Sabtu"
			break
		case "Sunday":
			DayName = "Minggu"
			break
	}

	// Day
	Day = strconv.Itoa(t.Day())

	// Month Name in String
	switch m := t.Month().String(); m {
	case "January":
		Month = "Januari"
		break
	case "February":
		Month = "Febuari"
		break
	case "March":
		Month = "Maret"
		break
	case "April":
		Month = "April"
		break
	case "May":
		Month = "Mei"
		break
	case "June":
		Month = "Juni"
		break
	case "July":
		Month = "Juli"
		break
	case "August":
		Month = "Agustus"
		break
	case "September":
		Month = "September"
		break
	case "October":
		Month = "Oktober"
		break
	case "November":
		Month = "November"
		break
	case "December":
		Month = "Desember"
		break
	}

	return t.format(format)
}