package main

import (
	"database/sql"
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	_ "github.com/glebarez/go-sqlite"
)

//go:embed all:frontend/dist
var assets embed.FS
var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite", "tasks.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
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
