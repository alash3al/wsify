package main

import (
	"fmt"
)

func main() {
	// handle panics
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[!!] Panic")
			fmt.Println(err)
		}
	}()

	// Step 1
	{
		fmt.Println("Welcome to WSIFY", VERSION)
		InitFlags()
	}
	// Step 2
	{
		fmt.Println("[*] listening for connections on address", *FLAG_HTTP_ADDR, " ...")
		if err := InitWsServer(*FLAG_HTTP_ADDR); err != nil {
			fmt.Println("[!] please solve the following issue")
			fmt.Println(err)
			return
		}
	}
}
