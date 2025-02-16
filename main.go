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
	execPath, err := os.Executable()
	if err != nil {
		log.Fatal("Failed to get executable path:", err)
	}

	execDir := filepath.Dir(execPath)

	dbPath := filepath.Join(execDir, "tasks.db")

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Println("Database not found, copying from project root...")

		projectDbPath := "tasks.db"
		err = copyFile(projectDbPath, dbPath)
		if err != nil {
			log.Fatal("Failed to copy database file:", err)
		}
	}

	// Открываем БД
	db, err = sql.Open("sqlite", dbPath+"?_foreign_keys=on")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
}

func copyFile(src, dst string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, input, 0644)
}

func main() {
	InitDB()

	app := NewApp(db)

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
