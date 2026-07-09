package core

import "context"

type IDataBase interface {
	GetAliasByName(ctx context.Context, name string) (*Alias, error)
	ListAliases(ctx context.Context) ([]*Alias, error)
	CreateAlias(ctx context.Context, alias *Alias) error
	UpdateAlias(ctx context.Context, id string, alias *Alias) error
	DeleteAlias(ctx context.Context, id string) error

	GetTemplateByName(ctx context.Context, name string) (*Template, error)
	ListTemplates(ctx context.Context) ([]*Template, error)
	CreateTemplate(ctx context.Context, template *Template) error
	UpdateTemplate(ctx context.Context, id string, template *Template) error
	DeleteTemplate(ctx context.Context, id string) error

	Close()
}
