package pkgbuild

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/kovetskiy/lorg"
	"github.com/reconquest/faces"
	"github.com/stretchr/testify/assert"
)

func TestOpen_ReturnsPkgbuild(t *testing.T) {
	test := assert.New(t)

	file, err := ioutil.TempFile(os.TempDir(), "pkgbuild-go")
	if err != nil {
		panic(err)
	}

	pkgbuild, err := Open(file.Name())
	test.NoError(err)
	test.EqualValues(file.Name(), pkgbuild.Path)
}

func TestGetDepends_ReturnsDependenciesList(t *testing.T) {
	test := assert.New(t)

	file, err := ioutil.TempFile(os.TempDir(), "pkgbuild-go")
	if err != nil {
		panic(err)
	}

	pkgbuild, err := Open(file.Name())
	test.NoError(err)

	file.WriteString(`
depends=(go git)
`)
	logger := lorg.NewLog()
	logger.SetLevel(lorg.LevelDebug)

	faces.SetLogger(logger)

	depends, err := pkgbuild.GetDepends()
	test.NoError(err)
	test.EqualValues([]string{"go", "git"}, depends)
}
