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
            task TEXT,
            status TEXT DEFAULT 'active'
        );
    `
    _, err = db.Exec(createTable)
    if err != nil {
        log.Fatal(err)
    }

	_, err = db.Exec("ALTER TABLE todos ADD COLUMN category TEXT;")
    if err != nil {
        log.Println("Поле category уже существует или произошла ошибка:", err)
    }
}

func GetDB() *sql.DB {
    return db
}

func CreateTodo(todo models.Todo) (int64, error) {
    result, err := db.Exec("INSERT INTO todos (task, status, category) VALUES (?, ?, ?)", todo.Task, todo.Status, todo.Category)
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}

func GetTodos(category string) ([]models.Todo, error) {
    var rows *sql.Rows
    var err error

    if category != "" {
        rows, err = db.Query("SELECT id, task, status, category FROM todos WHERE category = ?", category)
    } else {
        rows, err = db.Query("SELECT id, task, status, category FROM todos")
    }

    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var todos []models.Todo
    for rows.Next() {
        var todo models.Todo
        err := rows.Scan(&todo.ID, &todo.Task, &todo.Status, &todo.Category)
        if err != nil {
            return nil, err
        }
        todos = append(todos, todo)
    }
    return todos, nil
}

func GetTodo(id int) (models.Todo, error) {
    var todo models.Todo
    err := db.QueryRow("SELECT id, task, status FROM todos WHERE id = ?", id).Scan(&todo.ID, &todo.Task, &todo.Status)
    if err != nil {
        return todo, err
    }
    return todo, nil
}

func UpdateTodo(id int, todo models.Todo) error {
    _, err := db.Exec("UPDATE todos SET task = ?, status = ? WHERE id = ?", todo.Task, todo.Status, id)
    return err
}

func UpdateTodoStatus(id int, status string) error {
    _, err := db.Exec("UPDATE todos SET status = ? WHERE id = ?", status, id)
    return err
}

func DeleteTodo(id int) error {
    _, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
    return err
}