package pkgbuild

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/kovetskiy/executil"
)

type PKGBUILD struct {
	Path string
}

func Open(path string) (*PKGBUILD, error) {
	return &PKGBUILD{
		Path: path,
	}, nil
}

func (pkgbuild *PKGBUILD) GetDepends() ([]string, error) {
	cmd := exec.Command(
		"bash", "-c", fmt.Sprintf(
			`source %q && echo "${depends[@]}"`,
			pkgbuild.Path,
		),
	)

	stdout, _, err := executil.Run(cmd)
	if err != nil {
		return nil, err
	}

	depends := strings.Fields(strings.TrimSpace(string(stdout)))

	return depends, nil
}
