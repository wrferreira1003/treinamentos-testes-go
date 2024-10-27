package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/wrferreira1003/SQLC/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

type CourseDb struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDb(dbConn *sql.DB) *CourseDb {
	return &CourseDb{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

// Executa uma transação e retorna um erro se houver algum problema ou se der erro no commit
func (c *CourseDb) CallTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, errRb)
		}
		return err
	}
	return tx.Commit()
}

func (c *CourseDb) CreateCourseAndCategory(ctx context.Context, course CourseParams, category CategoryParams) error {
	err := c.CallTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
		if err != nil {
			return err
		}
		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			Price:       course.Price,
			CategoryID:  category.ID,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()

	dbq, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		log.Fatal(err)
	}
	defer dbq.Close()

	queries := db.New(dbq)

	//Executar uma transação se de ok damos um commit, se der erro damos um rollback

	// courseArgs := CourseParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Go",
	// 	Description: sql.NullString{String: "Go is a programming language", Valid: true},
	// 	Price:       10.99,
	// }

	// categoryArgs := CategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Programming",
	// 	Description: sql.NullString{String: "Programming category", Valid: true},
	// }

	// courseDb := NewCourseDb(dbq)
	// err = courseDb.CreateCourseAndCategory(ctx, courseArgs, categoryArgs)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	courses, err := queries.ListCourses(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, course := range courses {
		fmt.Println("Course:", course.Name, "ID:", course.ID, "Price:", course.Price, "Description:", course.Description.String, "Category:", course.CategoryName.String)
	}
}
