package main

import (
	"context"
	"database/sql"
	"log"
)

// App struct with DB connection
type App struct {
	ctx context.Context
	db  *sql.DB
}

// ✅ NewApp now accepts `db` and stores it
func NewApp(db *sql.DB) *App {
	return &App{db: db}
}

// startup is called when the app starts
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Task struct
type todos struct {
	Id     int
	Task   string
	Status bool
}

// ✅ Corrected Insert function
func (a *App) Insert(c *todos) (int64, error) {
	if a.db == nil {
		log.Println("❌ Database is NOT initialized")
		return 0, sql.ErrConnDone
	}

	sqlQuery := `INSERT INTO todos (task, status) VALUES (?, ?);`
	result, err := a.db.Exec(sqlQuery, c.Task, false) // Default `status` to false
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (a *App) ListTodos() ([]todos, error) {
	if a.db == nil {
		log.Println("❌ Database is NOT initialized")
		return nil, sql.ErrConnDone
	}

	var tasks []todos
	rows, err := a.db.Query("SELECT id, task, status FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task todos
		if err := rows.Scan(&task.Id, &task.Task, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (a *App) DeleteTask(id int) error {
	_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}

func (a *App) ToggleStatus(id int) error {
	_, err := db.Exec("UPDATE todos SET status = NOT status WHERE id = ?", id)
	return err
}

// func CreateTable(db *sql.DB) (sql.Result, error) {
// 	sql := `CREATE TABLE IF NOT EXISTS todos (
// 		id INTEGER PRIMARY KEY,
// 		task     TEXT NOT NULL,
// 		status Boolean DEFAULT FALSE
//     );`

// 	return db.Exec(sql)
// }
