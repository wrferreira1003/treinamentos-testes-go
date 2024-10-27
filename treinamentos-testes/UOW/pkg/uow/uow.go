package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// RepositoryFactory é uma função que cria um repositório
type RepositoryFactory func(tx *sql.Tx) interface{}

type UowInterface interface {
	Register(name string, fc RepositoryFactory)                          // Registra um repositorio
	GetRepository(ctx context.Context, name string) (interface{}, error) // Retorna um repositorio
	Do(ctx context.Context, fn func(uow *Uow) error) error               // Executa a função passada como parametro dentro de uma transação
	CommitOrRollback() error                                             // Commita ou Rollbacka a transação
	Rollback() error                                                     // Rollbacka a transação
	Unregister(name string)                                              // Remove um repositorio registrado
}

type Uow struct {
	Db           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory // Repositorios registrados
}

// NewUow cria uma nova instancia de Uow
func NewUow(ctx context.Context, db *sql.DB) *Uow {
	return &Uow{
		Db:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

// Register registra um repositorio
func (u *Uow) Register(name string, fc RepositoryFactory) {
	u.Repositories[name] = fc
}

// Unregister remove um repositorio registrado
func (u *Uow) Unregister(name string) {
	delete(u.Repositories, name)
}

// GetRepository retorna um repositorio
func (u *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	// Verifica se já existe uma transação, caso não exista cria uma
	if u.Tx == nil {
		tx, err := u.Db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		u.Tx = tx
	}
	// Retorna o repositorio com a transação
	repo := u.Repositories[name](u.Tx) // Cria o repositorio
	return repo, nil                   // Retorna o repositorio
}

// Do executa a função passada como parametro dentro de uma transação
func (u *Uow) Do(ctx context.Context, fn func(uow *Uow) error) error {
	// Verifica se já existe uma transação, caso exista retorna um erro
	if u.Tx != nil {
		return fmt.Errorf("transaction already started")
	}

	// Inicia uma transação
	tx, err := u.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	u.Tx = tx // Define a transação

	err = fn(u) // Executa a função passada como parametro dentro da transação
	if err != nil {
		erroRB := u.Rollback() // Rollbacka a transação
		if erroRB != nil {
			return fmt.Errorf("original error: %w, rollback error: %w", err, erroRB)
		}
		return err
	}
	// Commita ou Rollbacka a transação
	return u.CommitOrRollback()
}

// CommitOrRollback commita ou rollbacka a transação
func (u *Uow) CommitOrRollback() error {
	err := u.Tx.Commit()
	if err != nil {
		// Rollbacka a transação
		erroRB := u.Rollback()
		if erroRB != nil {
			return fmt.Errorf("commit error: %w, rollback error: %w", err, erroRB)
		}
		return err
	}
	u.Tx = nil // Define a transação como nil
	return nil
}

// Rollback rollbacka a transação
func (u *Uow) Rollback() error {
	if u.Tx == nil {
		return errors.New("transaction not started")
	}
	// Rollbacka a transação
	err := u.Tx.Rollback()
	if err != nil {
		return fmt.Errorf("rollback failed: %w", err)
	}
	u.Tx = nil // Define a transação como nil
	return nil
}
