package models

import (
	"github.com/gofiber/fiber/v2"
	"todolist/database"
)

type Todo struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetTodos(c *fiber.Ctx) error {
	db := database.DBConn
	var todos []Todo
	db.Find(&todos)

	return c.JSON(todos)
}

func GetTodoById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var todo Todo
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not find todo",
			"data":    err,
		})
	}

	return c.Status(200).JSON(&todo)
}

func CreatTodo(c *fiber.Ctx) error {
	db := database.DBConn

	todo := new(Todo)
	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Check your input",
			"data":    err,
		})
	}

	err = db.Create(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create todo",
			"data":    err,
		})
	}

	return c.Status(200).JSON(&todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	type UpdatedTodo struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	id := c.Params("id")
	db := database.DBConn

	var todo Todo
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not find todo",
			"data":    err,
		})
	}

	var updatedTodo UpdatedTodo
	err = c.BodyParser(&updatedTodo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Check your input",
			"data":    err,
		})
	}

	todo.Title = updatedTodo.Title
	todo.Completed = updatedTodo.Completed
	db.Save(&todo)

	return c.JSON(&todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var todo Todo
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not find todo",
			"data":    err,
		})
	}

	db.Delete(&todo)

	return c.Status(200).JSON(fiber.Map{
		"message": "Delete todo success",
	})
}
