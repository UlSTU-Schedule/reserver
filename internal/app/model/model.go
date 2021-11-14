package model

import (
	"github.com/jmoiron/sqlx/types"
	"time"
)

// Schedule ...
type Schedule struct {
	ID         int
	Name       string
	UpdateTime time.Time
	Info       types.JSONText
}
