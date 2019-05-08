package main

import (
	"fmt"

	"./publisher"
	"./receiver"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "docker"
	dbname   = "demo"
)

func main() {
	go receiver.Main()
	// fmt.Println("give time for receiver to settle")
	// fmt.Scanln() // wait for Enter Key
	publisher.Init()
	fmt.Println("Press the Enter Key to wait!")
	fmt.Scanln() // wait for Enter Key
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	// // Connect to the DB, panic if failed
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	fmt.Println(`Could not connect to db`)
	// 	panic(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	publisher.PublishAMessage("Database connected!")
	fmt.Println("Let the receiver have time to listen and output")
	fmt.Scanln() // wait for Enter Key
}
