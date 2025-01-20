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

func GetTodos(filter map[string]string) ([]models.Todo, error) {
    query := "SELECT id, task, status, created_at, updated_at FROM todos"
    var args []interface{}
    var conditions []string

    if status, ok := filter["status"]; ok {
        conditions = append(conditions, "status = ?")
        args = append(args, status)
    }

    if createdAfter, ok := filter["created_after"]; ok {
        conditions = append(conditions, "created_at >= ?")
        args = append(args, createdAfter)
    }

    if createdBefore, ok := filter["created_before"]; ok {
        conditions = append(conditions, "created_at <= ?")
        args = append(args, createdBefore)
    }

    if updatedAfter, ok := filter["updated_after"]; ok {
        conditions = append(conditions, "updated_at >= ?")
        args = append(args, updatedAfter)
    }

    if updatedBefore, ok := filter["updated_before"]; ok {
        conditions = append(conditions, "updated_at <= ?")
        args = append(args, updatedBefore)
    }

    if len(conditions) > 0 {
        query += " WHERE " + strings.Join(conditions, " AND ")
    }

    rows, err := db.Query(query, args...)
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