package command

import (
	"log"
	"os"
	"task1/core"
)

type DEL struct {
}

func (d DEL) Delete(data core.InputData) {
	del := os.Remove(data.Result[1])
	if del != nil {
		log.Fatal(del)
	}
}
