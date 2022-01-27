package core

type InputData struct {
	Result []string
}

func NewID() InputData {
	return InputData{}
}

func (i *InputData) SetRes(res []string) {
	for _, data := range res {
		i.Result = append(i.Result, data)
	}
}

func (i *InputData) UnsetRes() {
	i.Result = nil
}
