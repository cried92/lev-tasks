package command

import (
	"fmt"
	"os"
	"task1/core"
)

type SZ struct {
	data []string
}

// работает только в пределах текущего каталога
func (s SZ) Size(data core.InputData) {

	file, err := os.Open(data.Result[1])
	if err != nil {
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}
	fmt.Println(stat.Name(), stat.Size(), "bytes")
}
