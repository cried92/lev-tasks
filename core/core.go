package core

type InputData struct {
	Result []string
}

func NewID() InputData {
	return InputData{}
}

type MemComm map[string]string

func (m *MemComm) Unset() {
	*m = map[string]string{}
}

func (i *InputData) SetRes(res []string) {
	for _, data := range res {
		i.Result = append(i.Result, data)
	}
}

func (i *InputData) UnsetRes() {
	i.Result = nil
}
