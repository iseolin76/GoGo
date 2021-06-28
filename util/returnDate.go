package util

import "time"

func ReturnDate(msg string) string {
	now := time.Now();
		switch msg{
		case "오늘":
		case "내일":
			now = now.AddDate( 0,0,1)
		case "모레":
			now = now.AddDate( 0,0,2)
		case "글피":
			now = now.AddDate( 0,0,3)
		case "그글피":
			now = now.AddDate( 0,0,4)
		case "어제":
			now = now.AddDate( 0,0,-1)
		case "그제":
			now = now.AddDate( 0,0,-2)
		case "그그제":
			now = now.AddDate( 0,0,-3)
		case "다음주":
			now = now.AddDate( 0,0,7)
		case "저번주":
			now = now.AddDate( 0,0,-7)
		case "월요일":
			now = now.AddDate( 0,0,int(time.Monday-now.Weekday()))
		case "화요일":
			now = now.AddDate( 0,0,int(time.Tuesday-now.Weekday()))
		case "수요일":
			now = now.AddDate( 0,0,int(time.Wednesday-now.Weekday()))
		case "목요일":
			now = now.AddDate( 0,0,int(time.Thursday-now.Weekday()))
		case "금요일":
			now = now.AddDate( 0,0,int(time.Friday-now.Weekday()))
		case "토요일":
			now = now.AddDate( 0,0,int(time.Saturday-now.Weekday()))
		case "일요일":
			now = now.AddDate( 0,0,int(time.Sunday-now.Weekday()))
		default:
		}
	return now.Format("20060102")
}