package tests

import (
	sharedfrp "github.com/fatedier/frp/sharefrp/mobile"
	"testing"
)

const frpcIni = `
	[common]
   server_addr = 127.0.0.1
   server_port = 7000

   [ssh]
   type = tcp
   local_ip = 127.0.0.1
   local_port = 22
   remote_port = 6000
`

func TestRunClient(t *testing.T) {
	_ = sharedfrp.RunClient("default", frpcIni, "")
}

const frpsIni = `

	[common]
	bind_port = 7000

`

func TestRunServer(t *testing.T) {
	_ = sharedfrp.RunServer(frpsIni)
}
