package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// NewCategory cria uma nova categoria
func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (Category, error) {
	// Cria uma nova categoria
	id := uuid.New().String()

	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)",
		id, name, description)
	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	// Encontra todas as categorias
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Cria um slice para armazenar as categorias
	categories := []Category{}

	// Itera sobre as linhas retornadas pelo banco de dados
	for rows.Next() {
		var id, name, description string
		err = rows.Scan(&id, &name, &description)
		if err != nil {
			return nil, err
		}
		// Adiciona a categoria ao slice
		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}

	// Retorna o slice de categorias
	return categories, nil
}

func (c *Category) FindByCourseID(courseID string) (Category, error) {
	// Encontra uma categoria pelo ID
	row := c.db.QueryRow("SELECT c.id, c.name, c.description FROM categories c INNER JOIN courses co ON c.id = co.category_id WHERE co.id = $1", courseID)

	// Cria uma variável para armazenar a categoria encontrada
	var category Category

	// Escaneia os dados da linha para a variável
	err := row.Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return Category{}, err
	}

	// Retorna a categoria encontrada
	return category, nil
}
