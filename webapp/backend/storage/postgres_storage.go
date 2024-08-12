package storage

import (
	"context"
	"fmt"
	"log"
	"strings"
	"todo-webapp/backend/models"
	"github.com/jackc/pgx/v5"
)

type PostgresStore struct {
	conn *pgx.Conn
}

func NewPostgres(username, password, host, port, dbName string) (PostgresStore, error) {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbName)
	fmt.Println(connStr)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return PostgresStore{}, err
	}

	return PostgresStore{conn: conn}, nil
}

func (pgs PostgresStore) Close() {
	pgs.conn.Close(context.Background())
}

// --- Interface Methods ---
// They are exposed to the rest of the application.
// Each one sends a task to the RequestManager so the request can be processed

func (pgs PostgresStore) FindAll() []models.ToDo {
	rows, err := pgs.conn.Query(context.Background(), "SELECT * FROM mytodos")

	if err != nil {
		log.Printf("Query failed: %v\n", err)
		return []models.ToDo{}
	}

	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ToDo])
	if err != nil {
		log.Printf("CollectRows failed: %v\n", err)
		return items
	}

	return items
}

func (pgs PostgresStore) FindById(id int) (models.ToDo, error) {

	row, err := pgs.conn.Query(context.Background(), "SELECT * FROM mytodos WHERE id=$1", id)

	if err != nil {
		log.Printf("Query failed: %v\n", err)
		return models.ToDo{}, err
	}

	item, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.ToDo])
	if err != nil {
		log.Printf("CollectOneRow failed: %v\n", err)
		return item, err
	}

	return item, nil
}

func (pgs PostgresStore) Create(task string, status models.Status) models.ToDo {

	row, err := pgs.conn.Query(context.Background(), "INSERT INTO mytodos (task, status) VALUES ($1, $2) RETURNING id, task, status", task, status)
	if err != nil {
		log.Printf("Query failed: %v\n", err)
		return models.ToDo{}
	}

	item, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.ToDo])

	if err != nil {
		log.Printf("CollectOneRow failed: %v\n", err)
		return models.ToDo{}
	}

	return item
}

func (pgs PostgresStore) Update(id int, task *string, status *models.Status) (models.ToDo, error) {

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
	query := fmt.Sprintf("UPDATE mytodos SET %s WHERE id = @id RETURNING id, task, status", setClause)
	 
	row, err := pgs.conn.Query(context.Background(), query, args)
	if err != nil {
		log.Printf("Query failed: %v\n", err)
		return models.ToDo{}, err
	}

	item, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.ToDo])
	if err != nil {
		log.Printf("CollectOneRow failed: %v\n", err)
		return models.ToDo{}, err
	}

	return item, nil
}

func (pgs PostgresStore) Delete(id int) error {

	_, err := pgs.conn.Exec(context.Background(), "DELETE FROM mytodos WHERE id=$1", id)
	if err != nil {
		log.Printf("Delete failed: %v\n", err)
		return err
	}

	return nil
}
