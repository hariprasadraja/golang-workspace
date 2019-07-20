package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Format(time.RFC3339))
	fmt.Println(now.Format(time.RFC822))
	fmt.Println(now.Format(time.RFC1123Z))
	fmt.Println(now.Format(time.RFC850))

	//seconds :=now.Unix()
	//fmt.Println(seconds)
	//nanoseconds := now.UnixNano()
	//fmt.Println(nanoseconds)

	//
	//then := time.Date(
	//	2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	//fmt.Println(then)
	//fmt.Println(then.Weekday())
	//fmt.Println(then.After(now))
	//fmt.Println(then.Before(now))
	//fmt.Println(then.Equal(now))
	//timediffrence := then.Sub(now)
	//fmt.Println(timediffrence)

}
