package infra

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/CommonBerry/goman/internal/core"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func NewPostgresDataBase(ctx context.Context) (*PostgresDataBase, error) {
	_ = godotenv.Load()

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is required")
	}

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database URL: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return &PostgresDataBase{pool: pool}, nil
}

type PostgresDataBase struct {
	pool *pgxpool.Pool
}

// Aliases
// CreateAlias creates an alias in the database.
func (db *PostgresDataBase) CreateAlias(ctx context.Context, alias *core.Alias) error {
	query := `INSERT INTO aliases (id, name, alias, path)
			  VALUES ($1, $2, $3, $4)
	`

	alias.ID = uuid.New()
	_, err := db.pool.Exec(ctx, query, alias.ID, alias.Name, alias.Alias, alias.Path)
	if err != nil {
		return fmt.Errorf("failed to insert alias: %w", err)
	}

	return nil
}

// ListAliases lists the database aliases.
func (db *PostgresDataBase) ListAliases(ctx context.Context) ([]*core.Alias, error) {
	query := `
		SELECT id, name, alias, path
		FROM aliases
	`

	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query aliases: %w", err)
	}

	defer rows.Close()

	var aliases []*core.Alias

	for rows.Next() {
		a := &core.Alias{}

		err := rows.Scan(&a.ID, &a.Name, &a.Alias, &a.Path)
		if err != nil {
			return nil, fmt.Errorf("failed to scan aliases row: %w", err)
		}

		aliases = append(aliases, a)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return aliases, nil
}

// GetAliasByName retrieves an alias by name from the database.
func (db *PostgresDataBase) GetAliasByName(ctx context.Context, name string) (*core.Alias, error) {
	query := `SELECT id, name, alias, path FROM aliases WHERE name = $1`

	item := &core.Alias{}

	err := db.pool.QueryRow(ctx, query, name).Scan(&item.ID, &item.Name, &item.Alias, &item.Path)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("alias not found")
		}
		return nil, err
	}

	return item, nil
}

// UpdadeAlias
func (db *PostgresDataBase) UpdateAlias(ctx context.Context, id string, alias *core.Alias) error {
	query := `
		UPDATE aliases
		SET name = $1, alias = $2, path = $3
		WHERE id = $4
	`

	_, err := db.pool.Exec(ctx, query, alias.Name, alias.Alias, alias.Path, alias.ID)
	if err != nil {
		return fmt.Errorf("failed to update alias: %w", err)
	}

	return nil
}

// DeleteAlias
func (db *PostgresDataBase) DeleteAlias(ctx context.Context, id string) error {
	query := `DELETE FROM aliases WHERE id = $1`

	_, err := db.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete alias: %w", err)
	}

	return nil
}

// Templates
// CreateTemplate creates an template in the database
func (db *PostgresDataBase) CreateTemplate(ctx context.Context, template *core.Template) error {
	query := `INSERT INTO templates (id, name, description, author, version, repository, init_commands)
			  VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	template.ID = uuid.New()
	_, err := db.pool.Exec(ctx, query, template.ID, template.Name, template.Description, template.Author, template.Version, template.Repository, template.InitCommands)
	if err != nil {
		return fmt.Errorf("failed to insert template: %w", err)
	}

	return nil
}

// ListTemplates lists the database templates.
func (db *PostgresDataBase) ListTemplates(ctx context.Context) ([]*core.Template, error) {
	query := `
		SELECT id, name, description, author, version, repository, init_commands
		FROM templates
	`

	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query templates: %w", err)
	}

	defer rows.Close()

	var templates []*core.Template

	for rows.Next() {
		t := &core.Template{}

		err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.Author, &t.Version, &t.Repository, &t.InitCommands)
		if err != nil {
			return nil, fmt.Errorf("failed to scan templates row: %w", err)
		}

		templates = append(templates, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return templates, nil
}

// GetTemplateByName
func (db *PostgresDataBase) GetTemplateByName(ctx context.Context, name string) (*core.Template, error) {
	query := `SELECT id, name, description, author, version, repository, init_commands FROM templates WHERE name = $1`

	item := &core.Template{}

	err := db.pool.QueryRow(ctx, query, name).Scan(&item.ID, &item.Name, &item.Description, &item.Author, &item.Version, &item.Repository, &item.InitCommands)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("template not found")
		}
		return nil, err
	}

	return item, nil
}

// UpdateTemplate
func (db *PostgresDataBase) UpdateTemplate(ctx context.Context, id string, template *core.Template) error {
	query := `
		UPDATE templates
		SET name = $1, description = $2, author = $3, version = $4, repository = $5, init_commands = $6
		WHERE id = $7
	`

	_, err := db.pool.Exec(ctx, query, template.Name, template.Description, template.Author, template.Version, template.Repository, template.InitCommands, template.ID)
	if err != nil {
		return fmt.Errorf("failed to update template: %w", err)
	}

	return nil
}

// DeleteTemplate
func (db *PostgresDataBase) DeleteTemplate(ctx context.Context, id string) error {
	query := `DELETE FROM templates WHERE id = $1`

	_, err := db.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete template: %w", err)
	}

	return nil
}

func (db *PostgresDataBase) Close() {
	db.pool.Close()
}
