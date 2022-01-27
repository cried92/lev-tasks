package command

import (
	"fmt"
	"path"
	"task1/core"
)

type CD struct {
	ID  core.InputData
	dir string
}

func (c CD) ChangeDirectory(data core.InputData, dir string) {

	if (len(data.Result)) == 1 {
		curDir := dir
		fmt.Println(curDir)
		return
	}

	value := data.Result[1]

	switch value {
	case "..":
		//fmt.Println("move to prev (higher) level")
		curDir := path.Dir(dir)
		fmt.Println(curDir)
	default:
		newDir := dir + "/" + data.Result[1]
		curDir := newDir
		fmt.Println(curDir)
	}
}
