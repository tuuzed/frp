package statik

import (
	"github.com/rakyll/statik/fs"
)

func Init() {
	fs.Register(data)
}