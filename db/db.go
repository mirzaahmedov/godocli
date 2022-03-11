package db

import (
	"strconv"

	"github.com/mirzaahmedov/todo/model"
	"github.com/olekukonko/tablewriter"

	"encoding/json"
	"errors"
	"os"
)

type Storage struct {
  path string
  Todos []model.Todo 
}

func NewStorage(path string) (*Storage, error) {
  var todos []model.Todo
  if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
    file, err := os.Create(path)
    if err != nil {
      return nil, err
    }
    _, err = file.Write([]byte("[]"))
    if err != nil {
      return nil, err
    }
    defer file.Close()
    return &Storage{ path: path, Todos: []model.Todo{} }, nil
  }
  bytes, err := os.ReadFile(path)
  if err != nil {
    return nil, err
  }
  err = json.Unmarshal(bytes, &todos)
  if err != nil {
    return nil, err
  }
  return &Storage{ path: path, Todos: todos }, nil
}


func (s *Storage) save() error {
  bytes, err := json.Marshal(s.Todos)
  if err != nil {
    return err
  }
  err = os.WriteFile(s.path, bytes, 0644)
  if err != nil {
    return err
  }
  return nil
}


func (s *Storage) PrintTodos() {
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"ID", "Name", "Date", "Done"})
  table.SetBorder(false)                                
  for index, todo := range s.Todos {
    table.Append([]string{ strconv.Itoa(index), todo.Name, todo.Date.Format("Mon Jan _2 15:04:05 2006"), strconv.FormatBool(todo.Done) })
  }
  table.Render()
}

func (s *Storage) AddTodo(t *model.Todo) error {
  s.Todos = append(s.Todos, *t)
  err := s.save()
  if err != nil {
    return err
  }
  return nil
}

func (s *Storage) FindTodo(index int) (*model.Todo, error) { 
  if index >= len(s.Todos) {
    return nil, errors.New("Not Found 404")
  }
  return &s.Todos[index], nil
}  

func (s *Storage) DeleteTodo(index int) error {
  if index >= len(s.Todos) {
    return errors.New("Not Found 404")
  }
  s.Todos = append(s.Todos[:index], s.Todos[index + 1:]...)
  s.save()
  return nil
}  

func (s *Storage) CheckTodo(index int, value bool) error {
  if index >= len(s.Todos) {
    return errors.New("Not Found")
  }
  s.Todos[index].Done = value
  s.save()
  return nil
}
