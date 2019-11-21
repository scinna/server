package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/oxodao/scinna/dal"
	"github.com/oxodao/scinna/services"
	"github.com/oxodao/scinna/utils"
)

func main() {

	fmt.Println("Scinna Server - V1")

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using currently exported vars")
	}
	fmt.Println("- Env var loaded")

	db := utils.LoadDatabase()
	defer db.Close()
	fmt.Println("- Connected to database")

	prv := services.New(db)

	u, err := dal.GetUser(prv, "admin")
	fmt.Println(u.ToString())
	fmt.Println(err)

	p, err := dal.GetPicture(prv, 1)
	fmt.Println(p.ToString())
	fmt.Println(err)

	fmt.Println()
	fmt.Println("Les photos de admin sont: ")
	fmt.Println()

	ps, err := dal.GetPicturesFromUser(prv, "admin")
	fmt.Println(err)
	for i := 0; i < len(ps); i++ {
		fmt.Println(ps[i].ToString())
	}

}
