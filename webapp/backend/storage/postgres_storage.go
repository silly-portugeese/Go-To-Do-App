package storage

import (
	"context"
	"fmt"
	"log"
	"strings"
	"todo-webapp/backend/models"
	"github.com/jackc/pgx/v5"
)

type postgresStore struct {
	conn *pgx.Conn
}

func NewPostgres(username, password, host, port, dbName string) (postgresStore, error) {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbName)
	fmt.Println(connStr)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return postgresStore{}, err
	}

	return postgresStore{conn: conn}, nil
}

func (pgs postgresStore) Close() {
	pgs.conn.Close(context.Background())
}

// --- Interface Methods ---
// They are exposed to the rest of the application.
// Each one sends a task to the RequestManager so the request can be processed

func (pgs postgresStore) FindAll() []models.ToDo {
	rows, err := pgs.conn.Query(context.Background(), "SELECT id, task, status FROM todos")

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
	err := pgs.conn.QueryRow(context.Background(), "SELECT id, task, status FROM todos WHERE id=$1::int", id).Scan(&item.Id, &item.Task, &item.Status)

	if err != nil {
		log.Printf("Scan failed: %v\n", err)
		return item, err
	}

	return item, nil
}

func (pgs postgresStore) Create(task string, status models.Status) models.ToDo {
	var item models.ToDo
	err := pgs.conn.QueryRow(context.Background(), "INSERT INTO todos (task, status) VALUES ($1, $2) RETURNING id, task, status", task, status).Scan(&item.Id, &item.Task, &item.Status)
	
	if err != nil {
		log.Printf("Scan failed: %v\n", err)
	}

	return item
}

func (pgs postgresStore) Update(id int, task *string, status *models.Status) (models.ToDo, error) {

	var setClauses []string
	args := pgx.NamedArgs{"id": id}

	if task != nil {
		setClauses = append(setClauses, "task = @task")
		args["task"] = *task
	}

	if status != nil {
		setClauses = append(setClauses, "status = @status")
		args["status"] = *status
	}

	setClause := strings.Join(setClauses, ", ")
	query := fmt.Sprintf("UPDATE todos SET %s WHERE id = @id RETURNING id, task, status", setClause)

	var item models.ToDo
	err := pgs.conn.QueryRow(context.Background(), query, args).Scan(&item.Id, &item.Task, &item.Status)
	if err != nil {
		log.Printf("Scan failed: %v\n", err)
		return item, err
	}
	return item, nil
}

func (pgs postgresStore) Delete(id int) error {

	_, err := pgs.conn.Exec(context.Background(), "DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		log.Printf("Delete failed: %v\n", err)
		return err
	}

	return nil
}
