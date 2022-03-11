package model

import "time"

type Todo struct {
  Name string`json:"name"`
  Done bool`json:"done"`
  Date time.Time`json:"time"`
}
