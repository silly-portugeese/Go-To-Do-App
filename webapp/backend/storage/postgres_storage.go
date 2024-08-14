package storage

import (
	"context"
	"fmt"
	"log"
	"strings"
	"todo-webapp/backend/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)


type postgresStore struct {
	dbpool *pgxpool.Pool
}

func NewPostgres(username, password, host, port, dbName string) (postgresStore, error) {

	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbName)

	dbpool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Fatal("Error while creating pool")
	}

	// conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return postgresStore{}, err
	}

	return postgresStore{dbpool: dbpool}, nil
}

func (pgs postgresStore) Close() {
	pgs.dbpool.Close()
}

// --- Interface Methods ---
// They are exposed to the rest of the application.
// Each one sends a task to the RequestManager so the request can be processed

func (pgs postgresStore) FindAll() []models.ToDo {
	
	query := "SELECT id, task, status, user_id FROM todos"
	rows, err := pgs.dbpool.Query(context.Background(), query)

	if err != nil {
		log.Printf("Query failed: %v\n", err)
		return []models.ToDo{}
	}

	items, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.ToDo])
	if err != nil {
		log.Printf("CollectRows failed: %v\n", err)
		return items
	}

	return items
}

func (pgs postgresStore) FindById(id int) (models.ToDo, error) {

	var item models.ToDo

	query := "SELECT (id, task, status, user_id) FROM todos WHERE id=$1::int"
	err := pgs.dbpool.QueryRow(context.Background(), query, id).Scan(&item)

	// query :=  "SELECT id, task, status FROM todos WHERE id=$1::int"
	// err := pgs.conn.QueryRow(context.Background(), query, id).Scan(&item.Id, &item.Task, &item.Status)

	if err != nil {
		log.Printf("Scan failed: %v\n", err)
		return item, err
	}

	return item, nil
}

func (pgs postgresStore) Create(params models.TodoCreateParams) models.ToDo {

	var item models.ToDo
	// TODO: validation
	// TODO: check if user exists
	query := "INSERT INTO todos (task, status, user_id) VALUES ($1, $2, $3) RETURNING (id, task, status, user_id)"
	err := pgs.dbpool.QueryRow(context.Background(), query, params.Task, params.Status, params.UserId).Scan(&item)

	// query := "INSERT INTO todos (task, status) VALUES ($1, $2) RETURNING id, task, status"
	// err := pgs.conn.QueryRow(context.Background(), query, task, status).Scan(&item.Id, &item.Task, &item.Status)

	if err != nil {
		log.Printf("Scan failed: %v\n", err)
	}

	return item
}

func (pgs postgresStore) Update(id int, params models.TodoUpdateParams) (models.ToDo, error) {

	var setClauses []string
	args := pgx.NamedArgs{"id": id}

	if params.Task != nil {
		setClauses = append(setClauses, "task = @task")
		args["task"] = *params.Task
	}

	if params.Status != nil {
		setClauses = append(setClauses, "status = @status")
		args["status"] = *params.Status
	}

	setClause := strings.Join(setClauses, ", ")

	var item models.ToDo
	query := fmt.Sprintf("UPDATE todos SET %s WHERE id = @id RETURNING (id, task, status, user_id)", setClause)
	err := pgs.dbpool.QueryRow(context.Background(), query, args).Scan(&item)

	// query := fmt.Sprintf("UPDATE todos SET %s WHERE id = @id RETURNING id, task, status", setClause)
	// err := pgs.conn.QueryRow(context.Background(), query, args).Scan(&item.Id, &item.Task, &item.Status)

	if err != nil {
		log.Printf("Scan failed: %v\n", err)
		return item, err
	}
	return item, nil
}

func (pgs postgresStore) Delete(id int) error {

	query := "DELETE FROM todos WHERE id=$1"
	_, err := pgs.dbpool.Exec(context.Background(), query, id)
	if err != nil {
		log.Printf("Delete failed: %v\n", err)
		return err
	}

	return nil
}
