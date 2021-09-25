package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	now := time.Now()
	nextMeetingTime := getNextMeetingTime()
	fmt.Printf("now: %s", now)
	fmt.Printf("nextMeetingTime: %s", nextMeetingTime)
	fmt.Printf("opening browser")
	//openBrowser("http://www.google.com")
}

func getNextMeetingTime() time.Time {
	return time.Now().Add(time.Second * 5)
}

func openBrowser(url string) {
	exec.Command("open", url).Start()
}
