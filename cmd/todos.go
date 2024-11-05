package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/AteevMishra/todo-CLI/db"
	"github.com/AteevMishra/todo-CLI/models"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(getCmd)
	RootCmd.AddCommand(addCmd)
	RootCmd.AddCommand(updateCmd)
	RootCmd.AddCommand(deleteCmd)
}

// Command to fetch all todos
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get all todos",
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := db.GetAllTodos()
		if err != nil {
			log.Fatalf("Error fetching todos: %v", err)
		}
		allTodos, _ := json.MarshalIndent(todos, "", "  ")
		fmt.Println(todos)
		fmt.Println("hulk")
		fmt.Println(allTodos)
		fmt.Println("spider")
		fmt.Println(string(allTodos))
	},
}

// Command to add a new todo
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Run: func(cmd *cobra.Command, args []string) {
		body, _ := cmd.Flags().GetString("body")
		if body == "" {
			fmt.Println("Please provide a body for the todo item using --body")
			return
		}

		todo := models.Todo{Body: body, Completed: false}
		err := db.AddTodo(&todo)
		if err != nil {
			log.Fatalf("Failed to add todo item: %v", err)
		}
		fmt.Println("Todo added:", todo)
	},
}

// Command to mark a todo as completed
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a todo as completed",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		err := db.UpdateTodoStatus(id)
		if err != nil {
			log.Fatalf("Failed to update todo: %v", err)
		}
		fmt.Println("Todo updated as completed:", id)
	},
}

// Command to delete a todo by ID
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a todo by ID",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		err := db.DeleteTodoByID(id)
		if err != nil {
			log.Fatalf("Failed to delete todo: %v", err)
		}
		fmt.Println("Todo deleted:", id)
	},
}

func init() {
	addCmd.Flags().String("body", "", "Body of the todo item")
	updateCmd.Flags().String("id", "", "ID of the todo item to mark as completed")
	deleteCmd.Flags().String("id", "", "ID of the todo item to delete")
}