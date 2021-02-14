package sonda_test

import (
	"testing"

	sonda "github.com/CezarGarrido/nasa-sonda/sonda"
	"github.com/stretchr/testify/assert"
)

type CommandTest struct {
	commands []string
}

// Testes para comandos inválidos
func TestInvalidCommands(t *testing.T) {

	tests := []CommandTest{
		{
			commands: []string{"GD", "M", "M", "M"},
		},
		{
			commands: []string{"M", "M", "M", "M", "M"},
		},
		{
			commands: []string{"GD", "M"},
		},
		{
			commands: []string{"T", "K"},
		},
		{
			commands: []string{"OLA", ""},
		},
		{
			commands: []string{"GE", "M", "GE", "M"},
		},
	}

	for _, tt := range tests {
		probe := sonda.NewProbe()
		err := probe.Run(tt.commands)
		assert.NotNil(t, err)
	}

}

func TestValidCommands(t *testing.T) {

	tests := []CommandTest{
		{
			commands: []string{"GE", "M", "M", "M"},
		},
		{
			commands: []string{"M", "M", "M", "M"},
		},
		{
			commands: []string{"GE", "M"},
		},
		{
			commands: []string{"GE", "GD"},
		},
		{
			commands: []string{"GE", "M", "M", "M", "GD", "M", "M"},
		},
		{
			commands: []string{"GE", "M", "M", "M", "GD", "GE", "M"},
		},
		{
			commands: []string{"GE", "M", "GD", "M"},
		},
		{
			commands: []string{"GE", "GE", "GE", "GE"},
		},
		{
			commands: []string{"GD", "GD", "GD", "GD"},
		},
		{
			commands: []string{"GE", "M", "GE"},
		},
	}

	for _, tt := range tests {
		probe := sonda.NewProbe()
		err := probe.Run(tt.commands)
		assert.Nil(t, err)
	}

}

func TestIsLastValidPosition(t *testing.T) {

	type TestPosition struct {
		X         int
		Y         int
		direction sonda.Direction
		expected  bool
	}

	var tests = []TestPosition{
		{0, 5, sonda.Esquerda, false},
		{0, 4, sonda.Esquerda, false},
		{2, 2, sonda.Esquerda, false},
		{0, 5, sonda.Baixo, false},
		{0, 4, sonda.Baixo, false},
		{2, 2, sonda.Baixo, false},
		{-0, -5, sonda.Baixo, false},
		{-10, -52, sonda.Direita, false},
		{0, 5, sonda.Esquerda, false},
	}

	for _, tt := range tests {
		probe := sonda.NewProbe()
		probe.X = tt.X
		probe.Y = tt.Y
		probe.Direction = tt.direction
		isValid := probe.IsValidPosition()
		assert.Equal(t, tt.X, probe.X)
		assert.Equal(t, tt.Y, probe.Y)
		assert.Equal(t, tt.expected, isValid)
	}

}
