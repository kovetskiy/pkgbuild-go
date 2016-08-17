package pkgbuild

import (
	"fmt"
	"strings"

	"github.com/reconquest/faces"
	"github.com/reconquest/faces/commands/bash"
)

var (
	shell *bash.Bash
)

func getShell() (*bash.Bash, error) {
	if shell == nil {
		var err error
		shell, err = faces.NewBash()
		if err != nil {
			return nil, err
		}
	}

	return shell, nil
}

type PKGBUILD struct {
	Path string
}

func Open(path string) (*PKGBUILD, error) {
	return &PKGBUILD{
		Path: path,
	}, nil
}

func (pkgbuild *PKGBUILD) GetDepends() ([]string, error) {
	shell, err := getShell()
	if err != nil {
		return nil, err
	}

	stdout, _, err := shell.Eval(fmt.Sprintf(
		`source %q && echo "${depends[@]}"`,
		pkgbuild.Path,
	))
	if err != nil {
		return nil, err
	}

	depends := strings.Fields(strings.TrimSpace(string(stdout)))

	return depends, nil
}
