package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
)

var (
    tailObject *tail.Tail
	tailChan *tail.Line
)

func Init(fileName string) (err error) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	tailObject, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail.TailFile err: ", err)
		return
	}
	fmt.Println("Init success: ", tailObject)
    return
}

func GetTailChan() (tailChan <-chan *tail.Line) {
	return tailObject.Lines
}