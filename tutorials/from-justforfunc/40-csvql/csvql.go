package csvql

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/src-d/go-mysql-serveer.v0/sql"
)

type Database struct {
	tables map[string]sql.Table
}

func NewDatabase(path string) (*Database, error) {
	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.Wrapf(err, "could not read %s", path)
	}
	tables := make(map[string]sql.Table)
	for _, fi := range fis {
		name := strings.ToLower(fi.Name())
		if fi.IsDir() || filepath.Ext(name) != ".csv" {
			continue
		}
		t, err := newTable(filepath.Join(path, name))
	}
}
