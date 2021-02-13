package sonda

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// Direction
type Direction string

const (
	Esquerda Direction = "E" // E - Esquerda
	Direita            = "D" // D - Direita
	Cima               = "C" // C - Cima
	Baixo              = "B" // B - Baixo
)

// Probe
type Probe struct {
	X                 int       `json:"x"`
	Y                 int       `json:"y"`
	Direction         Direction `json:"face"`
	SequenceMovements string    `json:"sequence_movements"`
	countX            int
	countY            int
	lastMove          string
}

// NewProbe:
func NewProbe() *Probe {
	return &Probe{
		X:         0,
		Y:         0,
		Direction: Direita,
		countX:    0,
		countY:    0,
	}

}

func (probe *Probe) ResetCounters() {
	probe.countX = 0
	probe.countY = 0
}

// Move - movimentar. Para cada comando M a sonda se move uma posição na direção à qual sua face está apontada.
func (probe *Probe) IsValidPosition() bool {
	if (probe.Direction == Baixo || probe.Direction == Esquerda) && (probe.X > 0 || probe.Y > 0) {
		return false
	}
	return (probe.X >= 0 && probe.X <= 4) && (probe.Y >= 0 && probe.Y <= 4)
}

func (probe *Probe) Move() error {

	switch probe.Direction {
	case Cima:
		probe.MoveY()
	case Direita:
		probe.MoveX()
	case Baixo:
		probe.MoveY()
	case Esquerda:
		probe.MoveX()
	}

	if !probe.IsValidPosition() {
		return errors.New(fmt.Sprintf("Um movimento inválido foi detectado, infelizmente a sonda ainda não possui a habilidade de navegar na area {\"x\": %d, \"y\": %d}", probe.X, probe.Y))
	}
	return nil
}

// Print -
func (probe *Probe) Print() {
	b, _ := json.MarshalIndent(probe, "", "  ")
	fmt.Println(string(b))
}

// GE - girar 90 graus à esquerda
func (probe *Probe) GE() {
	switch probe.Direction {
	case Cima:
		probe.Direction = Esquerda
	case Direita:
		probe.Direction = Cima
	case Baixo:
		probe.Direction = Direita
	case Esquerda:
		probe.Direction = Baixo
	}
	if probe.lastMove == "M" {
		if probe.countY > 0 {
			probe.SequenceMovements += "se moveu " + strconv.Itoa(probe.countY) + " casas no eixo y, "
		}
		if probe.countX > 0 {
			probe.SequenceMovements += "se moveu " + strconv.Itoa(probe.countX) + " casas no eixo y, "
		}
	}
	probe.SequenceMovements += "girou para esquerda, "

}

// GD - girar 90 graus à direta
func (probe *Probe) GD() {
	switch probe.Direction {
	case Cima:
		probe.Direction = Direita
	case Direita:
		probe.Direction = Baixo
	case Baixo:
		probe.Direction = Esquerda
	case Esquerda:
		probe.Direction = Cima
	}
	if probe.lastMove == "M" {
		if probe.countY > 0 {
			probe.SequenceMovements += "se moveu " + strconv.Itoa(probe.countY) + " casas no eixo y, "
		}
		if probe.countX > 0 {
			probe.SequenceMovements += "se moveu " + strconv.Itoa(probe.countX) + " casas no eixo y, "
		}
	}
	probe.SequenceMovements += "girou para direita, "
}

func (probe *Probe) MoveX() {
	probe.X += 1
	probe.countX++
}

func (probe *Probe) MoveY() {
	probe.Y += 1
	probe.countY++
}

// Restart-
func (probe *Probe) Restart() *Probe {
	probe.X = 0
	probe.Y = 0
	probe.lastMove = ""
	probe.ResetCounters()
	return probe

}

func (probe *Probe) Run(commands []string) (err error) {
	count := len(commands)
	probe.SequenceMovements = "a sonda "
	for i, command := range commands {
		err = probe.runCommand(count, i, command)
		if err != nil {
			probe.Restart()
			break
		}
	}

	return err
}

func (probe *Probe) runCommand(count int, i int, command string) (err error) {

	if command == "GE" {
		probe.GE()
		probe.lastMove = command
		probe.ResetCounters()
		return
	}
	if command == "GD" {
		probe.GD()
		probe.lastMove = command
		probe.ResetCounters()
		return
	}
	if command == "M" {
		err = probe.Move()
		if count-i == 1 {
			if probe.countY > 0 {
				probe.SequenceMovements += "e andou mais " + strconv.Itoa(probe.countY) + " casas no eixo y"
			}
			if probe.countX > 0 {
				probe.SequenceMovements += "e andou mais " + strconv.Itoa(probe.countX) + " casas no eixo y"
			}
		}
		probe.lastMove = command
		return
	}

	err = errors.New("Comando inválido:" + command)
	return
}
