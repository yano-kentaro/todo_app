package main

import (
	"fmt"
	// "log"

	"github.com/yano-kentaro/todo_app/app/models"
	// "github.com/yano-kentaro/todo_app/config"
)

func main() {
	fmt.Print("================================出力結果================================\n\n")

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "qwe123"
	// fmt.Println(u)
	// u.CreateUser()

	u, _ := models.GetUser(10)
	fmt.Println(u)

	fmt.Print("\n========================================================================\n")
}
