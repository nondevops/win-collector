package ports

import (
	"time"

	"github.com/n9e/win-collector/stra"
)

func Detect() {
	detect()
	go loopDetect()
}

func loopDetect() {
	for {
		time.Sleep(time.Second * 10)
		detect()
	}
}

func detect() {
	ps := stra.GetPortCollects()
	DelNoPortCollect(ps)
	AddNewPortCollect(ps)
}
