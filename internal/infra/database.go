// Package infra is the package that handles the technical infrastructure of the Goman API, such as the repository database.
package infra

import (
	"errors"
	"slices"

	"github.com/CommonBerry/goman/internal/core"
)

// DataBase In-memory, TODO - Implement PostgreSQL here later.
type DataBase struct {
	aliases   []*core.Alias
	templates []*core.Template
}

// Aliases

func (db *DataBase) GetAliasByName(name string) (*core.Alias, error) {
	for i := range db.aliases {
		if db.aliases[i].Name == name {
			return db.aliases[i], nil
		}
	}

	return nil, errors.New("alias not found")
}

func (db *DataBase) ListAliases() []*core.Alias {
	return db.aliases
}

func (db *DataBase) CreateAlias(alias *core.Alias) error {
	// TODO - Tratar mais tipos de erro depois
	if alias.Name == "" {
		return errors.New("the name cannot be empty")
	}

	db.aliases = append(db.aliases, alias)

	return nil
}

func (db *DataBase) UpdateAlias(oldName string, alias *core.Alias) error {
	for i := range db.aliases {
		if db.aliases[i].Name == oldName {
			db.aliases[i] = alias
			return nil
		}
	}

	return errors.New("alias not found")
}

func (db *DataBase) DeleteAlias(name string) error {
	for i := range db.aliases {
		if db.aliases[i].Name == name {
			db.aliases = slices.Delete(db.aliases, i, i+1)
			return nil
		}
	}

	return errors.New("alias not found")
}

// Templates

func (db *DataBase) GetTemplateByName(name string) (*core.Template, error) {
	for i := range db.templates {
		if db.templates[i].Name == name {
			return db.templates[i], nil
		}
	}

	return nil, errors.New("template not found")
}

func (db *DataBase) ListTemplates() []*core.Template {
	return db.templates
}

func (db *DataBase) CreateTemplate(template *core.Template) error {
	// TODO - Tratar mais erros aqui támbem
	if template.Name == "" {
		return errors.New("the name cannot be empty")
	}

	db.templates = append(db.templates, template)

	return nil
}

func (db *DataBase) UpdateTemplate(oldName string, template *core.Template) error {
	for i := range db.templates {
		if db.templates[i].Name == oldName {
			db.templates[i] = template
			return nil
		}
	}
	return errors.New("template not found")
}

func (db *DataBase) DeleteTemplate(name string) error {
	for i := range db.templates {
		if db.templates[i].Name == name {
			db.templates = slices.Delete(db.templates, i, i+1)
			return nil
		}
	}

	return errors.New("template not found")
}
