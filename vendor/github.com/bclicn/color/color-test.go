package color

import (
	"fmt"
)

// github.com/bclicn/color
// colorized output for Mac & Linux terminal
// version: 1.0.0
// author:  Beichen Li, bclicn@gmail.com (banned), relidin@126.com, 2016-11-4
// see:     http://misc.flogisoft.com/bash/tip_colors_and_formatting
// usage:
// For official code layout
// 	$ go get github.com/bclicn/color
// 	# in your project
// 	import "github.com/bclicn/color"
// 	func main() {
//		color.Test()
//	}

func ColorTest (){

	const HEAD = " "
	const TAIL = " "

	// regular
	fmt.Println(HEAD + Black("black") 		+ TAIL)
	fmt.Println(HEAD + Red("red") 			+ TAIL)
	fmt.Println(HEAD + Green("green") 		+ TAIL)
	fmt.Println(HEAD + Yellow("yellow") 		+ TAIL)
	fmt.Println(HEAD + Blue("blue") 		+ TAIL)
	fmt.Println(HEAD + Purple("purple") 		+ TAIL)
	fmt.Println(HEAD + Cyan("cyan") 		+ TAIL)
	fmt.Println(HEAD + LightGray("light gray") 	+ TAIL)
	fmt.Println(HEAD + DarkGray("dark gray") 	+ TAIL)
	fmt.Println(HEAD + LightRed("light red") 	+ TAIL)
	fmt.Println(HEAD + LightGreen("light green") 	+ TAIL)
	fmt.Println(HEAD + LightYellow("light yellow") 	+ TAIL)
	fmt.Println(HEAD + LightBlue("light blue") 	+ TAIL)
	fmt.Println(HEAD + LightPurple("light purple") 	+ TAIL)
	fmt.Println(HEAD + LightCyan("light cyan") 	+ TAIL)
	fmt.Println(HEAD + White("white") 		+ TAIL)

	// bold
	fmt.Println(HEAD + BBlack("bold black")			+ TAIL)
	fmt.Println(HEAD + BRed("bold red")			+ TAIL)
	fmt.Println(HEAD + BGreen("bold green")			+ TAIL)
	fmt.Println(HEAD + BYellow("bold yellow")		+ TAIL)
	fmt.Println(HEAD + BBlue("bold blue")			+ TAIL)
	fmt.Println(HEAD + BPurple("bold purple")		+ TAIL)
	fmt.Println(HEAD + BCyan("bold cyan")			+ TAIL)
	fmt.Println(HEAD + BLightGray("bold light gray")	+ TAIL)
	fmt.Println(HEAD + BDarkGray("bold dark gray")		+ TAIL)
	fmt.Println(HEAD + BLightRed("bold light red")		+ TAIL)
	fmt.Println(HEAD + BLightGreen("bold light green")	+ TAIL)
	fmt.Println(HEAD + BLightYellow("bold light yellow")	+ TAIL)
	fmt.Println(HEAD + BLightBlue("bold light blue")	+ TAIL)
	fmt.Println(HEAD + BLightPurple("bold light purple")	+ TAIL)
	fmt.Println(HEAD + BLightCyan("bold light cyan")	+ TAIL)
	fmt.Println(HEAD + BWhite("bold white")			+ TAIL)

	// background
	fmt.Println(HEAD + GBlack("background black") 			+ TAIL)
	fmt.Println(HEAD + GRed("background red") 			+ TAIL)
	fmt.Println(HEAD + GGreen("background green") 			+ TAIL)
	fmt.Println(HEAD + GYellow("background yellow") 		+ TAIL)
	fmt.Println(HEAD + GBlue("background blue") 			+ TAIL)
	fmt.Println(HEAD + GPurple("background purple") 		+ TAIL)
	fmt.Println(HEAD + GCyan("background cyan") 			+ TAIL)
	fmt.Println(HEAD + GLightGray("background light gray") 		+ TAIL)
	fmt.Println(HEAD + GDarkGray("background dark gray") 		+ TAIL)
	fmt.Println(HEAD + GLightRed("background light red") 		+ TAIL)
	fmt.Println(HEAD + GLightGreen("background light green") 	+ TAIL)
	fmt.Println(HEAD + GLightYellow("background light yellow") 	+ TAIL)
	fmt.Println(HEAD + GLightBlue("background light blue") 		+ TAIL)
	fmt.Println(HEAD + GLightPurple("background light purple") 	+ TAIL)
	fmt.Println(HEAD + GLightCyan("background light cyan") 		+ TAIL)
	fmt.Println(HEAD + GWhite("background white") 			+ TAIL)

	// special
	fmt.Println("A " + Bold("bold") + " text")
	fmt.Println("This is a " + Dim("dimmed") + " text")
	fmt.Println("Add a " + Underline("underline"))
	fmt.Println("Use " + Invert("invert") + " to highlight your text")
	fmt.Println("Your password is:" + Hide("myPass"))
	fmt.Println("OMG I'm " + Blink("blinking") + " !!!")	// blinking works only on mac
}