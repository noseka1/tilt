package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquash(t *testing.T) {
	cmds := ToUnixShellCmds([]string{"echo a", "echo b", "echo c"})
	cmds = TrySquash(cmds)
	assert.Equal(t, []Cmd{Cmd{
		Argv: []string{
			"sh",
			"-exc",
			"echo a;\necho b;\necho c",
		},
	}}, cmds)
}

func TestSquashFail(t *testing.T) {
	cmds := []Cmd{ToUnixShellCmd("echo a"), Cmd{Argv: []string{"echo"}}, ToUnixShellCmd("echo c")}
	cmds2 := TrySquash(cmds)
	assert.Equal(t, cmds, cmds2)
}

func TestSquashPartial(t *testing.T) {
	cmds := []Cmd{ToUnixShellCmd("echo a"), ToUnixShellCmd("echo c"), Cmd{Argv: []string{"echo"}}}
	cmds2 := TrySquash(cmds)
	assert.Equal(t, []Cmd{
		Cmd{
			Argv: []string{
				"sh",
				"-exc",
				"echo a;\necho c",
			},
		},
		cmds[2],
	}, cmds2)
}
