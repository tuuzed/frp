package frps

import (
	"github.com/rakyll/statik/fs"
)

func Load() {
	fs.Register(data)
}