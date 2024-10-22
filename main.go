package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

// func remove(slice []int, index int) []int {
//     return append(slice[:s], slice[s+1:]...)

// }

func remove(todos []Todo, index int) []Todo {
	return append(todos[:index], todos[index + 1:]...)

}

func convertSTRtoINT(str string) int {
	numId, err := strconv.Atoi(str)
	if (err != nil) {
		return -1

	}

	return numId

}

func main() {
	app := fiber.New()
	
	todos := [] Todo{}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading ENV")

	}

	PORT := os.Getenv("PORT")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello world"})

	})

	app.Post("/api/todos/new", func (c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err

		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required" })

		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		fmt.Println(todos)

		return c.Status(201).JSON(todo)
	})

	app.Patch("/api/todos/update", func (c *fiber.Ctx) error {
		fmt.Println("WE ARE UPDATEING")
		id := c.Query("id")

		fmt.Println(id)

		numId := convertSTRtoINT(id)

		fmt.Println(numId)

		newTodo := &Todo{}
		newTodo.ID = numId

		if err := c.BodyParser(newTodo); err != nil {
			return err

		}

		if newTodo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"msg": "HELL NAHH BODY IS REQUIERD"})

		}

		todoIndexToUpdate := numId - 1
		
		if todoIndexToUpdate < 0 || todoIndexToUpdate >= len(todos) {
			return c.Status(404).JSON(fiber.Map{ "error": "No TODO found to update" })

		}

		todos[todoIndexToUpdate] = *newTodo

		fmt.Println(todos)

		return c.Status(200).JSON(fiber.Map{ "msg": "Updated TODO sucssfully" })

	})

	app.Delete("/api/todos/delete", func (c *fiber.Ctx) error {
		var id string = c.Query("id")

		var todoIndexToDelete int = convertSTRtoINT(id) - 1
		
		if todoIndexToDelete < 0 || todoIndexToDelete >= len(todos) {
			return c.Status(404).JSON(fiber.Map{ "error": "No TODO found to delete" })

		}

		todos = remove(todos, todoIndexToDelete)

		return c.Status(204).JSON(fiber.Map{"msg": "Todo deleted sussfully"})

	})


	app.Listen(":" + PORT)

}




































// TODO: solve ID unsorted bug



