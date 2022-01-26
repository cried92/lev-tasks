package command

import (
	"fmt"
	"path"
	"task1/core"
)

type CD struct {
	ID     core.InputData
	curDir string
}

func (c CD) ChangeDirectory(data core.InputData, dir string) {

	fmt.Println(len(data.Result))
	if (len(data.Result)) == 1 {
		fmt.Println(dir)
	}

	value := data.Result[1]

	switch value {
	case "..":
		fmt.Println("move to prev (higher) level")
		prevDir := path.Dir(dir)
		fmt.Println(prevDir)
	case "...":
		fmt.Println("move to Home directory")
	default:
		fmt.Println(dir)

		curDir := dir + "/" + data.Result[1]
		fmt.Println(curDir)
		//curDir := []string
		//dir := append(curDir, i.result[1])
		//dir := strings.Join((curDir), "/")
		//fmt.Println(dir)
	}

}
