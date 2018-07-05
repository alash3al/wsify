package main

import (
	"fmt"

	"github.com/bclicn/color"
)

func debug(txt string) {
	if !*FlagDebug {
		return
	}
	fmt.Println(color.BCyan("[DEBUG]: ") + color.BLightYellow(txt))
}
