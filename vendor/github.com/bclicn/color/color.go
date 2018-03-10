// github.com/bclicn/color
// colorized output for Mac & Linux terminal
// version: 1.0.0
// author:  Beichen Li, bclicn@gmail.com (banned), relidin@126.com, 2016-9-23
// see:     http://misc.flogisoft.com/bash/tip_colors_and_formatting
// usage:
// Official layout
// 	$ go get github.com/bclicn/color
// 	# in your project
// 	import "github.com/bclicn/color"
// 	func main() {
//		fmt.Println(color.Red("I'm red!"))
//		color.Test()
//	}
// Quick usage
//	# put color.go and yourScript.go under a same folder
//	# change package name of both scripts to main
//	package main
//	# call color functions from yourScript by
//	fmt.Println(Red("I'm Red!"))
//	color.Test()
//	# run your script by
//	$ go run color.go yourScript.go
//	# build your script by
//	$ go build color.go yourScript.go

package color

import (
	"strconv"
)

const (
	// common
	reset 		= "\033[0m" 	// auto reset the rest of text to default color
	normal		= 0
	bold 		= 1 		// increase this value if you want bolder text
	// special
	dim		= 2
	underline 	= 4
	blink		= 5
	reverse		= 7
	hidden		= 8
	// color
	black 		= 30		// default = 39
	red		= 31
	green		= 32
	yellow		= 33
	blue		= 34
	purple		= 35		// purple = magenta
	cyan		= 36
	lightGray 	= 37
	darkGray 	= 90
	lightRed 	= 91
	lightGreen 	= 92
	lightYellow 	= 93
	lightBlue	= 94
	lightPurple 	= 95
	lightCyan	= 96
	white		= 97
)

// you can use custom color code and font size by calling this function
func Render (colorCode int, fontSize int, content string) string{
	return "\033[" + strconv.Itoa(fontSize) + ";" + strconv.Itoa(colorCode) + "m" + content + reset
}

// black text (use this with caution since most geeks use dark console)
func Black (txt string) string{
	return Render(black, normal, txt)
}

// red text
func Red (txt string) string {
	return Render(red, normal, txt)
}

// green text
func Green (txt string) string{
	return Render(green, normal, txt)
}

// yellow text
func Yellow (txt string) string{
	return Render(yellow, normal, txt)
}

// blue text
func Blue (txt string) string{
	return Render(blue, normal, txt)
}

// purple text
func Purple (txt string) string{
	return Render(purple, normal, txt)
}

// cyan text
func Cyan (txt string) string{
	return Render(cyan, normal, txt)
}

// light gray text
func LightGray (txt string) string{
	return Render(lightGray, normal, txt)
}

// dark gray text
func DarkGray (txt string) string{
	return Render(darkGray, normal, txt)
}

// light red text
func LightRed (txt string) string{
	return Render(lightRed, normal, txt)
}

// light green text
func LightGreen (txt string) string{
	return Render(lightGreen, normal, txt)
}

// light yellow text
func LightYellow (txt string) string{
	return Render(lightYellow, normal, txt)
}

// light blue text
func LightBlue (txt string) string{
	return Render(lightBlue, normal, txt)
}

// light purple text
func LightPurple (txt string) string{
	return Render(lightPurple, normal, txt)
}

// light cyan text
func LightCyan (txt string) string{
	return Render(lightCyan, normal, txt)
}

// white text
func White (txt string) string{
	return Render(white, normal, txt)
}

// black text (use this with caution since most geeks use dark console)
func BBlack (txt string) string{
	return Render(black, bold, txt)
}

// bold red
func BRed (txt string) string {
	return Render(red, bold, txt)
}

// bold green
func BGreen (txt string) string{
	return Render(green, bold, txt)
}

// bold yellow
func BYellow (txt string) string{
	return Render(yellow, bold, txt)
}

// bold blue
func BBlue (txt string) string{
	return Render(blue, bold, txt)
}

// bold purple
func BPurple (txt string) string{
	return Render(purple,  bold, txt)
}

// bold cyan
func BCyan (txt string) string{
	return Render(cyan, bold, txt)
}

// bold light gray
func BLightGray (txt string) string{
	return Render(lightGray, bold, txt)
}

// bold dark gray
func BDarkGray (txt string) string{
	return Render(darkGray, bold, txt)
}

// bold light red
func BLightRed (txt string) string{
	return Render(lightRed, bold, txt)
}

// bold light green
func BLightGreen (txt string) string{
	return Render(lightGreen, bold, txt)
}

// bold light yellow
func BLightYellow (txt string) string{
	return Render(lightYellow, bold, txt)
}

// bold light blue
func BLightBlue (txt string) string{
	return Render(lightBlue,  bold, txt)
}

// bold light purple
func BLightPurple (txt string) string{
	return Render(lightPurple,  bold, txt)
}

// bold light cyan
func BLightCyan (txt string) string{
	return Render(lightCyan,  bold, txt)
}

// bold white
func BWhite (txt string) string{
	return Render(white, bold, txt)
}

// black background (use this with caution since most of geeks use dark console)
func GBlack (txt string) string{
	return Render(black + 1, normal, txt)
}

// red background
func GRed (txt string) string {
	return Render(red + 1, normal, txt)
}

// green background
func GGreen (txt string) string{
	return Render(green + 1, normal, txt)
}

// yellow background
func GYellow (txt string) string{
	return Render(yellow + 1, normal, txt)
}

// blue background
func GBlue (txt string) string{
	return Render(blue + 1, normal, txt)
}

// purple background
func GPurple (txt string) string{
	return Render(purple + 1, normal, txt)
}

// cyan background
func GCyan (txt string) string{
	return Render(cyan + 1, normal, txt)
}

// light gray background
func GLightGray (txt string) string{
	return Render(lightGray + 1, normal, txt)
}

// dark gray background
func GDarkGray (txt string) string{
	return Render(darkGray + 1, normal, txt)
}

// light red background
func GLightRed (txt string) string{
	return Render(lightRed + 1, normal, txt)
}

// light green background
func GLightGreen (txt string) string{
	return Render(lightGreen + 1, normal, txt)
}

// light yellow background
func GLightYellow (txt string) string{
	return Render(lightYellow + 1, normal, txt)
}

// blue background
func GLightBlue (txt string) string{
	return Render(lightBlue + 1, normal, txt)
}

// light purple background
func GLightPurple (txt string) string{
	return Render(lightPurple + 1, normal, txt)
}

// light cyan background
func GLightCyan (txt string) string{
	return Render(lightCyan + 1, normal, txt)
}

// give text a white background
func GWhite (txt string) string{
	return Render(white + 1, normal, txt)
}

// bold text
func Bold (txt string) string{
	return Render(bold, normal, txt)
}

// dimmed text
func Dim (txt string) string{
	return Render(dim, normal,  txt)
}

// underlined text
func Underline (txt string) string{
	return Render(underline, 0 , txt)
}

// make given text blink, not supported by all consoles
func Blink (txt string) string{
	return Render(blink, normal,  txt)
}

// invert the color of text and its background
func Invert (txt string) string{
	return Render(reverse, normal,  txt)
}

// hide given text, useful for password input
func Hide (txt string) string{
	return Render(hidden, normal,  txt)
}

// test all functions
func Test() {
	ColorTest()
}