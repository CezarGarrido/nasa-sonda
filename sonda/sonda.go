// The probe package implements a framework to manipulate a probe in a rectangular area.
// The position of the probe is represented by its x and y axis, and the direction it is pointed by the initial letter,
// with valid directions being:
//  E - Left
//  D - Right
//  C - Up
//  B - Low
// The probe accepts three commands:
//  GE - rotate left 90 degrees
//  GD - rotate 90 degrees to the right
//  M - move. For each M command, the probe moves one position in the direction that its face is pointed.

package sonda

import (
	"errors"
	"fmt"
	"strconv"
)

// Direction: Represents the directions the probe can move and rotate
type Direction string

// The possible directions listed and mapped in PT-BR
const (
	LEFT   Direction = "E" // PT-BR: E - Esquerda
	RIGHT            = "D" // PT-BR: D - Direita
	TOP              = "C" // PT-BR: C - Cima
	BOTTOM           = "B" // PT-BR: B - Baixo
)

// Probe: Structure that represents the Probe, with position and direction and movement information
type Probe struct {
	X                 int       `json:"x"`                  // Axis X
	Y                 int       `json:"y"`                  // Axis y
	Direction         Direction `json:"face"`               // Direction
	SequenceMovements string    `json:"sequence_movements"` // Description of the movement that the probe does
	countX            int       // Amount that the probe moves in X in a complete sequence
	countY            int       // Amount that the probe moves in Y in a complete sequence
	lastCommand       string    // The penultimate command the probe took
}

// NewProbe : Creates and returns a new probe with the initial state
func NewProbe() *Probe {
	return &Probe{
		X:         0,
		Y:         0,
		Direction: RIGHT,
		countX:    0,
		countY:    0,
	}

}

// Reset the quantity counters
func (probe *Probe) ResetCounters() {
	probe.countX = 0
	probe.countY = 0
}

// IsValidPosition : Validates the current state of the probe, whether the movement it made is valid or not
// Simple comparisons are made to calculate whether a position is valid
func (probe *Probe) IsValidPosition() bool {
	if (probe.Direction == BOTTOM || probe.Direction == LEFT) && (probe.X > 0 || probe.Y > 0) {
		return false
	}
	return (probe.X >= 0 && probe.X <= 4) && (probe.Y >= 0 && probe.Y <= 4)
}

// Move : For each M command, the probe moves one position in the direction that its face is pointed.
// Returns error if it is an invalid state
func (probe *Probe) Move() error {

	lastX := probe.X
	lastY := probe.Y

	switch probe.Direction {
	case TOP:
		probe.MoveY()
	case RIGHT:
		probe.MoveX()
	case BOTTOM:
		probe.MoveY()
	case LEFT:
		probe.MoveX()
	}

	if !probe.IsValidPosition() {
		probe.X = lastX
		probe.Y = lastY
		return errors.New(fmt.Sprintf("Um movimento inválido foi detectado, infelizmente a sonda ainda não possui a habilidade de navegar nessa direção"))
	}
	return nil
}

// Rotate90DegLeft : Rotate 90 degrees to the left
func (probe *Probe) Rotate90DegLeft() {
	switch probe.Direction {
	case TOP:
		probe.Direction = LEFT
	case RIGHT:
		probe.Direction = TOP
	case BOTTOM:
		probe.Direction = RIGHT
	case LEFT:
		probe.Direction = BOTTOM
	}
	if probe.lastCommand == "M" {
		if probe.countY > 0 {
			probe.SequenceMovements += "se moveu " + strconv.Itoa(probe.countY) + " casas no eixo y, "
		}
		if probe.countX > 0 {
			probe.SequenceMovements += "se moveu " + strconv.Itoa(probe.countX) + " casas no eixo y, "
		}
	}
	probe.SequenceMovements += "girou para esquerda, "

}

// Rotate90DegRight : Rotate 90 degrees to the right
func (probe *Probe) Rotate90DegRight() {
	switch probe.Direction {
	case TOP:
		probe.Direction = RIGHT
	case RIGHT:
		probe.Direction = BOTTOM
	case BOTTOM:
		probe.Direction = LEFT
	case LEFT:
		probe.Direction = TOP
	}
	if probe.lastCommand == "M" {
		if probe.countY > 0 {
			probe.SequenceMovements += "se moveu " + strconv.Itoa(probe.countY) + " casas no eixo y, "
		}
		if probe.countX > 0 {
			probe.SequenceMovements += "se moveu " + strconv.Itoa(probe.countX) + " casas no eixo y, "
		}
	}
	probe.SequenceMovements += "girou para direita, "
}

// MoveX : A move in x
func (probe *Probe) MoveX() {
	probe.X += 1
	probe.countX++
}

// MoveY : A move in y
func (probe *Probe) MoveY() {
	probe.Y += 1
	probe.countY++
}

// Restart : Sends the probe to the initial state
func (probe *Probe) Restart() *Probe {
	probe.X = 0
	probe.Y = 0
	probe.lastCommand = ""
	probe.Direction = RIGHT
	probe.SequenceMovements = ""
	probe.ResetCounters()
	return probe
}

// Run : Receive a list of commands,
// And for each command it executes a function, returning an error if the probe goes into an unconscious state
// The probe accepts three commands:
//  GE - rotate left 90 degrees
//  GD - rotate 90 degrees to the right
//  M - move. For each M command, the probe moves one position in the direction that its face is pointed.

func (probe *Probe) Run(commands []string) (err error) {
	count := len(commands)
	probe.SequenceMovements = "a sonda "
	for i, command := range commands {
		err = probe.runCommand(count, i, command)
		if err != nil {
			break
		}
	}

	return err
}

func (probe *Probe) runCommand(count int, i int, command string) (err error) {

	if command == "GE" {
		probe.Rotate90DegLeft()
		probe.lastCommand = command
		probe.ResetCounters()
		return
	}
	if command == "GD" {
		probe.Rotate90DegRight()
		probe.lastCommand = command
		probe.ResetCounters()
		return
	}
	if command == "M" {
		err = probe.Move()

		if count-i == 1 {
			if probe.countY > 0 {
				if probe.lastCommand == "M" {
					probe.SequenceMovements += "andou " + strconv.Itoa(probe.countY) + " casas no eixo y"
				} else {
					probe.SequenceMovements += "e andou mais " + strconv.Itoa(probe.countY) + " casas no eixo y"
				}
			}
			if probe.countX > 0 {
				if probe.lastCommand == "M" {
					probe.SequenceMovements += "andou " + strconv.Itoa(probe.countX) + " casas no eixo x"
				} else {
					probe.SequenceMovements += "e andou mais " + strconv.Itoa(probe.countX) + " casas no eixo x"
				}

			}
		}
		probe.lastCommand = command
		return
	}

	err = errors.New("Comando inválido:" + command)
	return
}
