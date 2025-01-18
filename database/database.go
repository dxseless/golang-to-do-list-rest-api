package database

import (
    "database/sql"
    "log"
    "todo-api/models"
    _ "modernc.org/sqlite"
)

var db *sql.DB

func InitDB() {
    var err error
    db, err = sql.Open("sqlite", "./todos.db") 
    if err != nil {
        log.Fatal(err)
    }

    createTable := `
        CREATE TABLE IF NOT EXISTS todos (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            task TEXT
        );
    `
    _, err = db.Exec(createTable)
    if err != nil {
        log.Fatal(err)
    }
}

func GetDB() *sql.DB {
    return db
}

func CreateTodo(todo models.Todo) (int64, error) {
    result, err := db.Exec("INSERT INTO todos (task) VALUES (?)", todo.Task)
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}

func GetTodos() ([]models.Todo, error) {
    rows, err := db.Query("SELECT id, task FROM todos")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var todos []models.Todo
    for rows.Next() {
        var todo models.Todo
        err := rows.Scan(&todo.ID, &todo.Task)
        if err != nil {
            return nil, err
        }
        todos = append(todos, todo)
    }
    return todos, nil
}

func GetTodo(id int) (models.Todo, error) {
    var todo models.Todo
    err := db.QueryRow("SELECT id, task FROM todos WHERE id = ?", id).Scan(&todo.ID, &todo.Task)
    if err != nil {
        return todo, err
    }
    return todo, nil
}

func UpdateTodo(id int, todo models.Todo) error {
    _, err := db.Exec("UPDATE todos SET task = ? WHERE id = ?", todo.Task, id)
    return err
}

func DeleteTodo(id int) error {
    _, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
    return err
}