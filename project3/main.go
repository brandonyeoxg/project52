package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/brandonyeoxg/project52/project3/statemachine"
)

type Elevator struct {
	currentLevel  uint
	maxLevel      uint
	levelsPressed []uint
}

func NewElevator(maxLevel uint) Elevator {
	return Elevator{
		currentLevel:  1,
		maxLevel:      maxLevel,
		levelsPressed: make([]uint, 0, maxLevel),
	}
}

func (e Elevator) Validate() error {
	if e.currentLevel <= 0 || e.currentLevel > e.maxLevel {
		return errors.New("unsupported levels")
	}
	return nil
}

func selectLevels(ctx context.Context, elevator Elevator, levels []uint) {
	elevator.levelsPressed = levels
	if err := elevator.Validate(); err != nil {
		log.Println("Elevator is invalid because err:", err)
		return
	}

	elevator, err := statemachine.Run(context.Background(), IdleState, elevator)
	if err != nil {
		log.Printf("Encountered some issues running the elevator err: %s", err.Error())
	}

	log.Println("Elevator final level is", elevator.currentLevel)
}

func IdleState(ctx context.Context, elevator Elevator) (Elevator, statemachine.State[Elevator], error) {
	if err := elevator.Validate(); err != nil {
		return elevator, nil, err
	}
	log.Println("Elevator is idling at level", elevator.currentLevel)
	// check if there are buttons pressed
	if len(elevator.levelsPressed) == 0 {
		return elevator, nil, nil
	}

	nextLevel := elevator.levelsPressed[0]
	log.Printf("Elevator level %d button pressed", nextLevel)

	if nextLevel > elevator.currentLevel {
		return elevator, MoveUpState, nil
	}

	if nextLevel < elevator.currentLevel {
		return elevator, MoveDownState, nil
	}

	log.Printf("Elevator level is the same as current level")
	elevator.levelsPressed = elevator.levelsPressed[1:]
	return elevator, IdleState, nil
}

func MoveUpState(ctx context.Context, elevator Elevator) (Elevator, statemachine.State[Elevator], error) {
	if len(elevator.levelsPressed) == 0 {
		return elevator, nil, errors.New("no button has been pressed")
	}
	newLevel := elevator.levelsPressed[0]
	if newLevel > elevator.maxLevel {
		return elevator, nil, fmt.Errorf("unable to go up to the new level: %d", newLevel)
	}
	log.Println("Elevator moving up to level", newLevel)
	elevator.levelsPressed = elevator.levelsPressed[1:]
	elevator.currentLevel = newLevel
	return elevator, IdleState, nil
}

func MoveDownState(ctx context.Context, elevator Elevator) (Elevator, statemachine.State[Elevator], error) {
	if len(elevator.levelsPressed) == 0 {
		return elevator, nil, errors.New("no button has been pressed")
	}
	newLevel := elevator.levelsPressed[0]
	if newLevel < 1 {
		return elevator, nil, fmt.Errorf("unable to go down to the new level: %d", newLevel)
	}
	log.Println("Elevator down to level", newLevel)
	elevator.levelsPressed = elevator.levelsPressed[1:]
	elevator.currentLevel = newLevel
	return elevator, IdleState, nil
}

func main() {
	elevator := NewElevator(14)
	selectLevels(context.Background(), elevator, []uint{1, 2, 3, 4, 5, 6, 7})
}
