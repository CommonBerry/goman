package core

type IDataBase interface {
	GetAliasByName(name string) (*Alias, error)
	ListAliases() []*Alias
	CreateAlias(alias *Alias) error
	UpdateAlias(oldName string, alias *Alias) error
	DeleteAlias(name string) error

	GetTemplateByName(name string) (*Template, error)
	ListTemplates() []*Template
	CreateTemplate(template *Template) error
	UpdateTemplate(oldName string, template *Template) error
	DeleteTemplate(name string) error
}
