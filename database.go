package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite" // Альтернативный SQLite драйвер
)

var DB *sql.DB

// Инициализация базы данных
func initDB(dbPath string) error {
	var err error
	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("не удалось подключиться к базе данных: %v", err)
	}

	query := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	);`
	_, err = DB.Exec(query)
	if err != nil {
		return fmt.Errorf("ошибка при создании таблицы: %v", err)
	}

	fmt.Println("✅ База данных успешно инициализирована!")
	return nil
}

// Закрытие соединения с БД
func closeDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("🛑 Соединение с БД закрыто")
	}
}

// Добавление задачи
func AddTask(title string) error {
	_, err := DB.Exec("INSERT INTO tasks (title) VALUES (?)", title)
	return err
}

// Получение списка задач
func GetTasks() ([]map[string]interface{}, error) {
	rows, err := DB.Query("SELECT id, title FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []map[string]interface{}
	for rows.Next() {
		var id int
		var title string
		rows.Scan(&id, &title)
		tasks = append(tasks, map[string]interface{}{"id": id, "title": title})
	}
	return tasks, nil
}
