package command

import (
	"fmt"
	"io/ioutil"
	"log"
)

type LS struct {
}

// работает везде
func (l LS) List() {
	files, err := ioutil.ReadDir(curDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}
