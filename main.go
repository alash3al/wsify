package main

import (
	"fmt"
	"github.com/bclicn/color"
)

func main() {
	// handling any panic here ...
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(color.BRed("[!] Panic: ") + color.BLightYellow(err.(string)))
		}
	}()

	// parsing the command line flags
	fmt.Println(color.BGreen("[*] Welcome to WSIFY"), color.BCyan(VERSION))
	InitFlags()

	// start the pub/sub server
	fmt.Println(color.BGreen("[*] Listening for connections on address"), color.BCyan(*FLAG_HTTP_ADDR), color.BGreen(" ..."))
	if err := InitWsServer(*FLAG_HTTP_ADDR); err != nil {
		fmt.Println(color.BRed("[!] Error: ") + color.BLightYellow(err.Error()))
		return
	}
}
