// Package core
package core

type Template struct {
	Name        string
	Description string
	Author      string
	Version     string
	Repository  string
	InitComands []string
}

type Alias struct {
	Name  string
	Alias string
	Path  string
}
