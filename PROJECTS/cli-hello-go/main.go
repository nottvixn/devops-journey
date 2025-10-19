package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello, Andra! ✅ Your Go tool is running on Linux (WSL).")
	fmt.Print("Type your name then press Enter: ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Nice to meet you, %s", name)
}
