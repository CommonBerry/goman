// Package core
package core

import "github.com/google/uuid"

type Template struct {
	ID           uuid.UUID
	Name         string
	Description  string
	Author       string
	Version      string
	Repository   string
	InitCommands []string
}

type Alias struct {
	ID    uuid.UUID
	Name  string
	Alias string
	Path  string
}
