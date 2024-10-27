package database

import (
	"database/sql"

	"github.com/google/uuid"
)

// Course representa um curso
type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

// NewCourse cria uma nova instância de Course
func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

// Create cria um novo curso
func (c *Course) Create(name, description, categoryID string) (*Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)", id, name, description, categoryID)
	if err != nil {
		return nil, err
	}

	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
	}, nil
}

// FindAll encontra todos os cursos
func (c *Course) FindAll() ([]Course, error) {
	// Encontra todos os cursos
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Cria um slice para armazenar os cursos
	courses := []Course{}

	// Itera sobre as linhas retornadas pelo banco de dados
	for rows.Next() {
		var id, name, description, categoryID string
		err = rows.Scan(&id, &name, &description, &categoryID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, Course{
			ID:          id,
			Name:        name,
			Description: description,
			CategoryID:  categoryID,
		})
	}
	// Verifica se ocorreu algum erro durante a iteração
	if err = rows.Err(); err != nil {
		return nil, err
	}
	// Retorna os cursos encontrados
	return courses, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	// Encontra todos os cursos de uma categoria específica
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Cria um slice para armazenar os cursos
	courses := []Course{}

	// Itera sobre as linhas retornadas pelo banco de dados
	for rows.Next() {
		var id, name, description, categoryID string
		err = rows.Scan(&id, &name, &description, &categoryID)
		if err != nil {
			return nil, err
		}
		// Adiciona o curso ao slice de cursos
		courses = append(courses, Course{
			ID:          id,
			Name:        name,
			Description: description,
			CategoryID:  categoryID,
		})
	}
	// Verifica se ocorreu algum erro durante a iteração
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}
