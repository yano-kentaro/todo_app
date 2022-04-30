package main

import (
	"fmt"
	"log"

	"github.com/yano-kentaro/todo_app/config"
)

func main() {
	fmt.Print("================================出力結果================================\n\n")

	fmt.Println(config.Config)
	log.Println("test")

	fmt.Print("\n========================================================================\n")
}
