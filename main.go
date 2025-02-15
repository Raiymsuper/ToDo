package main

import (
	"database/sql"
	"embed"
	"log"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	_ "github.com/glebarez/go-sqlite"
)

//go:embed all:frontend/dist
var assets embed.FS
var db *sql.DB

func InitDB() {
	// Определяем путь к файлу базы данных в корне проекта
	wd, _ := os.Getwd()
	dbPath := filepath.Join(wd, "tasks.db")

	// Проверяем, существует ли база данных
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Fatal("Database file not found:", dbPath)
	}

	// Открываем существующую базу данных (НЕ создаём новую)
	var err error
	db, err = sql.Open("sqlite", dbPath+"?_foreign_keys=on")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
}
func main() {
	InitDB() // Initialize database first

	app := NewApp(db) // Pass `db` to App

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "todo",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
	defer db.Close()
}
