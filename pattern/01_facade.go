package pattern

import (
	"fmt"
	"time"
)

type Computer interface {
	On()
	Off()
}

type ComputerDevice interface {
	Start()
	Down()
}

type Monitor struct{}

type Keyboard struct{}

type Mouse struct{}

type Laptop struct {
	Monitor
	Mouse
	Keyboard
}

func (m Mouse) Start() {
	fmt.Println("The mouse starts...")
	time.Sleep(1 * time.Second)
	fmt.Println("Mouse is ready")
}

func (m Mouse) Down() {
	fmt.Println("The mouse is turns off")
	time.Sleep(1 * time.Second)
	fmt.Println("The mouse is down")
}

func (m Monitor) Start() {
	fmt.Println("The monitor starts...")
	time.Sleep(1 * time.Second)
	fmt.Println("Monitor is ready")
}

func (m Monitor) Down() {
	fmt.Println("The monitor is turns off")
	time.Sleep(1 * time.Second)
	fmt.Println("The monitor is down")
}

func (k Keyboard) Start() {
	fmt.Println("The keyboard starts...")
	time.Sleep(1 * time.Second)
	fmt.Println("keyboard is ready")
}

func (k Keyboard) Down() {
	fmt.Println("The keyboard is turns off")
	time.Sleep(1 * time.Second)
	fmt.Println("The keyboard is down")
}

func (L Laptop) On() {
	L.Monitor.Start()
	L.Mouse.Start()
	L.Keyboard.Start()
	fmt.Println("Laptop is ready")
}

func (L Laptop) Down() {
	L.Monitor.Down()
	L.Mouse.Down()
	L.Keyboard.Down()
	fmt.Println("Laptop is down")
}

// Паттерн фасад применяется когда у нас очень много подклассов (в случае golang – структур),
// но мы хотим сокрыть все простым интерфейсом, а сложную работу сокрыть в подклассах

// Плюсы: простой в использовании интерфейс
// Минусы: сложный в разработке, есть вероятность превращения в божественный объект
