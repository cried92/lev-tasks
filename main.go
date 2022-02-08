package main

import (
	"bufio"
	"task1/command"
	"task1/core"

	"os"
	"strings"
)

func main() {
	commands := []string{"cd", "ls", "sz", "del", "cp"}

	sc := bufio.NewScanner(os.Stdin)
	inpD := core.NewID()

	for sc.Scan() {
		input := sc.Text()
		result := strings.Split(input, " ")
		inpD.SetRes(result)

		switch result[0] {
		case commands[0]:
			cdD := command.CD{}
			cdD.ChangeDirectory(inpD)

		case commands[1]:
			ls := command.LS{}
			ls.List()

		case commands[2]:
			sz := command.SZ{}
			sz.Size(inpD)

		case commands[3]:
			del := command.DEL{}
			del.Delete(inpD)

		case commands[4]:
			cp := command.CP{}
			cp.Copy(inpD)
		}

		inpD.UnsetRes()
	}
}
