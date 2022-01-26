package command

import (
	"io/ioutil"
	"log"
	"task1/core"
)

type CP struct {
}

func (c CP) Copy(data core.InputData) {

	srcF := data.Result[1]
	dstF := data.Result[2]
	data2, err := ioutil.ReadFile(srcF)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(dstF, data2, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
