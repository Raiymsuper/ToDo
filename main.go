package main

import (
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

type App struct{}

func NewApp() *App {
	return &App{}
}

// Прокидываем функции в Wails
func (a *App) AddTask(title string) error {
	return AddTask(title)
}

func (a *App) GetTasks() ([]map[string]interface{}, error) {
	return GetTasks()
}

func main() {
	err := initDB("sqlite.db")
	if err != nil {
		fmt.Println("❌ Ошибка при инициализации БД:", err)
		return
	}
	defer closeDB()

	err = wails.Run(&options.App{
		Title:  "ToDo",
		Width:  1024,
		Height: 768,
		Assets: assets,
		Bind: []interface{}{
			NewApp(),
		},
	})

	if err != nil {
		fmt.Println("❌ Ошибка запуска приложения:", err)
	}
}
