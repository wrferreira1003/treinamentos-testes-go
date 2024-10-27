package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/wrferreira1003/SQLC/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	dbq, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		log.Fatal(err)
	}
	defer dbq.Close()

	queries := db.New(dbq)

	// // Create a new category
	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Golang",
	// 	Description: sql.NullString{String: "Golang course", Valid: true},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // update a category
	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:          "676fb5e6-13d6-4841-b30a-6b3821d1c502",
	// 	Name:        "Golang",
	// 	Description: sql.NullString{String: "Golang course updated", Valid: true},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// list all categories
	categories, err := queries.ListCategories(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Delete a category
	err = queries.DeleteCategory(ctx, "676fb5e6-13d6-4841-b30a-6b3821d1c502")
	if err != nil {
		log.Fatal(err)
	}

	// list all categories
	for _, category := range categories {
		fmt.Println(category)
	}
}
