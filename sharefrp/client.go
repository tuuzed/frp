package main

import (
	"C"
	sharedfrp "github.com/fatedier/frp/sharefrp/mobile"
)

//export LoadClientAssets
func LoadClientAssets() {
	sharedfrp.LoadClientAssets()
}

//export RunClient
func RunClient(name, iniContent, iniPath *C.char) C.int {
	err := sharedfrp.RunClient(C.GoString(name), C.GoString(iniContent), C.GoString(iniPath))
	if err == nil {
		return C.int(1)
	} else {
		return C.int(0)
	}
}

//export GetClientStatus
func GetClientStatus(name *C.char) C.int {
	return C.int(sharedfrp.GetClientStatus(C.GoString(name)))
}

//export CloseClient
func CloseClient(name *C.char) C.int {
	err := sharedfrp.CloseClient(C.GoString(name))
	if err == nil {
		return C.int(1)
	} else {
		return C.int(0)
	}
}

//export ReloadClientConf
func ReloadClientConf(name, iniContent *C.char) C.int {
	err := sharedfrp.ReloadClientConf(C.GoString(name), C.GoString(iniContent))
	if err == nil {
		return C.int(1)
	} else {
		return C.int(0)
	}
}
