package util

import (
	"fmt"
	"strings"
	"time"
)

func CheckDate(commandMsg string, now time.Time) time.Time {
	msg := strings.Split(commandMsg, " ")

	if len(msg) > 1 {
		now = checkThisDate(msg[0], now)
		now = checkWeekday(msg[1], now)
	} else {
		now = checkThisDate(msg[0], now)
	}
	return now
}

func checkThisDate(msg string, now time.Time) time.Time {
	//TODO: 
	// case "[0-9]달 [후,뒤]":
	// 	now = now.AddDate( 0,1,0)
	// case "[0-9]달 전":
	// 	now = now.AddDate( 0,-1,0)
	// 같은 경우 [0-9]를 판별하기
	// case "/*[요일]":
	// 	now = checkWeekday(msg, now)
	// 어떤 요일이라고 하는 지 판별하기
	
	switch msg {
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
		case "지난주":
			now = now.AddDate( 0,0,-7)
		case "내년":
			now = now.AddDate( 1,0,0)
		case "작년":
			now = now.AddDate( -1,0,0)
		case "내후년":
			now = now.AddDate( 2,0,0)
		case "재작년":
			now = now.AddDate( -2,0,0)
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
			fmt.Println("적절하지 않은 명령어 :", msg)
	}

	return now
}

func checkWeekday(msg string, now time.Time) time.Time {
	switch msg {
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
	default :
		fmt.Println("적절하지 않은 요일")
	}
	return now
}