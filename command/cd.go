package command

import (
	"fmt"
	"log"
	"os"
	"strings"
	"task1/core"
)

type CD struct {
	ID  core.InputData
	dir string
}

var curDir string

func (c CD) ChangeDirectory(data core.InputData) {

	if c.dir != "" {
		curDir = c.dir
	}
	orDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	if (len(data.Result)) == 1 {
		if curDir == "" {
			curDir = orDir
			c.SetDir(curDir)
			fmt.Println(curDir)
			return
		} else {
			fmt.Println(curDir)
			c.SetDir(curDir)
			return
		}
	}

	value := data.Result[1]

	switch value {
	case "":
		//fmt.Println("show current directory")
		fmt.Println(curDir)
		c.SetDir(curDir)
	case "..":
		//fmt.Println("move to prev (higher) level")
		dir := strings.Split(curDir, "/")
		dir = dir[:len(dir)-1]
		curDir = strings.Join(dir, "/")
		fmt.Println(curDir)
		c.SetDir(curDir)
	default:
		newCurDir := curDir + "/" + data.Result[1]
		fmt.Println(newCurDir)
		curDir = newCurDir
		c.SetDir(curDir)
	}
}

func (c *CD) SetDir(g string) {
	c.dir = g
}
