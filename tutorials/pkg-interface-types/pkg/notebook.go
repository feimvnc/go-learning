package pkg

import "fmt"

type Notebook struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewNotebook() Computer {
	return Notebook{
		Type:    NotebookType,
		Core:    8,
		Memory:  16,
		Monitor: true,
	}
}

func (pc Notebook) GetType() string {
	return pc.Type
}

func (pc Notebook) PrintDetails() {
	fmt.Printf("%s core:[%d] Mem:[%d] Monitor:[%v]\n", pc.Type, pc.Core, pc.Memory, pc.Monitor)
}
