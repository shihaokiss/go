package main

import (
	"fmt"
	"time"
	"github.com/hpcloud/tail"
)


func tailDemo() {
	fmt.Println("tailDemo start.")
	fileName := "./BackupService.log"
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail.TailFile err: ", err)
		return
	}

    var (
		li *tail.Line
		ok bool
	)
    for {
		li, ok = <-tails.Lines
		fmt.Println("tails.Lines: ", li, ok)
		if !ok {
			fmt.Println("file close.")
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line: ", li.Text)
	}
}

func main() {
	tailDemo()
}