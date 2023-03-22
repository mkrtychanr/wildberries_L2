package pattern

import "fmt"

const (
	ServType     string = "serv"
	NotebookType string = "notebook"
)

type StoreObject interface {
	GetType() string
	PrintDetails()
}

type Notebook struct {
	Display  string
	Keyboard string
	Trackpad string
}

type Serv struct {
	CPU    string
	Memory int
}

func NewServ() Serv {
	return Serv{
		CPU:    "Intel",
		Memory: 256,
	}
}

func NewNotebook() Notebook {
	return Notebook{
		Display:  "HP",
		Keyboard: "FullSize",
		Trackpad: "Crap",
	}
}

func (n Notebook) GetType() string {
	return "notebook"
}

func (n Notebook) PrintDetails() {
	fmt.Printf("Display %s, Keyboard %s, Trackpad %s\n", n.Display, n.Keyboard, n.Trackpad)
}

func (s Serv) GetType() string {
	return "serv"
}

func (s Serv) PrintDetails() {
	fmt.Printf("CPU %s, Mem %d\n", s.CPU, s.Memory)
}

func New(typeName string) StoreObject {
	switch typeName {
	default:
		fmt.Printf("Несуществующий тип %s\n", typeName)
		return nil
	case ServType:
		return NewServ()
	case NotebookType:
		return NewNotebook()
	}
}
