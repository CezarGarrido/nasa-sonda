package sonda_test

import (
	"testing"

	sonda "github.com/CezarGarrido/nasa-sonda/sonda"
	"github.com/stretchr/testify/assert"
)

type commandTest struct {
	commands []string
}

// Testes para comandos inválidos
func TestInvalidCommands(t *testing.T) {

	tests := []commandTest{
		{
			commands: []string{"GD", "M", "M", "M"},
		},
		{
			commands: []string{"M", "M", "M", "M", "M"},
		},
		{
			commands: []string{"GD", "M"},
		},
	}

	for _, tt := range tests {
		probe := sonda.NewProbe()
		err := probe.Run(tt.commands)
		assert.NotNil(t, err)
	}

}

func TestValidCommands(t *testing.T) {

	tests := []commandTest{
		{
			commands: []string{"GE", "M", "M", "M"},
		},
		{
			commands: []string{"M", "M", "M", "M"},
		},
		{
			commands: []string{"GE", "M"},
		},
	}

	for _, tt := range tests {
		probe := sonda.NewProbe()
		err := probe.Run(tt.commands)
		assert.Nil(t, err)
	}

}
