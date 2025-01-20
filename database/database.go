package database

import (
    "database/sql"
    "log"
    "time"
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
            task TEXT,
            status TEXT DEFAULT 'active',
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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
    result, err := db.Exec("INSERT INTO todos (task, status, created_at, updated_at) VALUES (?, ?, ?, ?)",
        todo.Task, todo.Status, time.Now(), time.Now())
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}

func GetTodos() ([]models.Todo, error) {
    rows, err := db.Query("SELECT id, task, status, created_at, updated_at FROM todos")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var todos []models.Todo
    for rows.Next() {
        var todo models.Todo
        err := rows.Scan(&todo.ID, &todo.Task, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt)
        if err != nil {
            return nil, err
        }
        todos = append(todos, todo)
    }
    return todos, nil
}

func GetTodo(id int) (models.Todo, error) {
    var todo models.Todo
    err := db.QueryRow("SELECT id, task, status, created_at, updated_at FROM todos WHERE id = ?", id).
        Scan(&todo.ID, &todo.Task, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt)
    if err != nil {
        return todo, err
    }
    return todo, nil
}

func UpdateTodo(id int, todo models.Todo) error {
    _, err := db.Exec("UPDATE todos SET task = ?, status = ?, updated_at = ? WHERE id = ?",
        todo.Task, todo.Status, time.Now(), id)
    return err
}

func DeleteTodo(id int) error {
    _, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
    return err
}