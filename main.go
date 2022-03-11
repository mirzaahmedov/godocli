package main

import (
	"log"
	"os"

	"github.com/mirzaahmedov/todo/cli"
	"github.com/mirzaahmedov/todo/db"
)

func main(){
  var err error
  if 2 > len(os.Args) {
    os.Exit(2)
  }
  storage, err := db.NewStorage("/home/mirzaahmedov/.config/todo/data.json")
  if err != nil {
    log.Fatal(err)
  }
  app, err := cli.NewApp(storage)
  if err != nil {
    log.Fatal(err)
  }
  err = app.Run()
  if err != nil {
    log.Fatal(err)
  }
}
