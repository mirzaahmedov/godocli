package cli

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mirzaahmedov/todo/model"
)

type Storage interface {
  AddTodo(*model.Todo) error
  PrintTodos()
  DeleteTodo(int) error
  CheckTodo(int, bool) error
}
type App struct {
  storage Storage
}

func NewApp(st Storage) (*App, error) {
  if st == nil {
    return nil, errors.New("No Storage Provided")
  }
  return &App{ storage: st }, nil
}
func (a *App) Run() error {
  addCmd := flag.NewFlagSet("add", flag.ExitOnError)
  deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
  checkCmd := flag.NewFlagSet("check", flag.ExitOnError)
  statusCmd := flag.NewFlagSet("print", flag.ExitOnError)

  addCmd.Usage = func() {
    fmt.Println("todo add [Name of your todo] - adds new item to the list")
  }
  deleteCmd.Usage = func() {
    fmt.Println("todo delete [index of todo] - deletes item from list")
  }
  checkCmd.Usage = func() {
    fmt.Println("todo check [index of todo] - check as done todo item")
  }
  statusCmd.Usage = func() {
    fmt.Println("todo print - prints all items in todo list")
  }

  switch os.Args[1] {
    case "add":
      addCmd.Parse(os.Args[2:])
      a.storage.AddTodo(&model.Todo{ Name: addCmd.Arg(0), Date: time.Now() })
    case "delete":
      deleteCmd.Parse(os.Args[2:])
      id, err := strconv.Atoi(deleteCmd.Arg(0))
      if err != nil {
        return err
      }
      err = a.storage.DeleteTodo(id)
      if err != nil {
        return err
      }
    case "check":
      undo := checkCmd.Bool("undo", false, "undo to in the list")
      checkCmd.Parse(os.Args[2:])
      id, err := strconv.Atoi(checkCmd.Arg(0))
      if err != nil {
        return err
      }
      err = a.storage.CheckTodo(id, !*undo)
      if err != nil {
        return err
      }
    case "status":
      statusCmd.Parse(os.Args[2:])
      a.storage.PrintTodos()
    default:
      return errors.New("Command Not Found")
  }
  return nil
}

func (a *App) ParseCommand() string {
  return os.Args[1]
}
