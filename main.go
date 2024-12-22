package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"todolist/database"
	"todolist/models"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	initDatabase()

	app.Get("/", helloWorld)
	setupRoutes(app)
	err := app.Listen(":8080")
	if err != nil {
		return
	}
}

func initDatabase() {
	var err error

	dsn := "host=127.0.0.1 user=postgres password=admin dbname=goTodo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	fmt.Println("Successfully connected to database")

	err = database.DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Migrated DB")
}

func setupRoutes(app *fiber.App) {
	baseUrl := "/api/v1"

	app.Get(baseUrl+"/todos", models.GetTodos)
	app.Get(baseUrl+"/todos/:id", models.GetTodoById)
	app.Post(baseUrl+"/todos", models.CreatTodo)
	app.Put(baseUrl+"/todos/:id", models.UpdateTodo)
	app.Delete(baseUrl+"/todos/:id", models.DeleteTodo)
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
