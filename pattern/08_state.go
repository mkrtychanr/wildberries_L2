package pattern

import "fmt"

type State interface {
	RespondLockButton()
}

type MacBook struct {
	locked       State
	unlocked     State
	currentState State
}

type LockedState struct{}

type UnlockedState struct{}

func (s *LockedState) RespondLockButton() {
	fmt.Println("Trynna to unlock")
}

func (s *UnlockedState) RespondLockButton() {
	fmt.Println("Trynna to lock")
}
