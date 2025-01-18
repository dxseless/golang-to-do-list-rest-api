package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "todo-api/database"
    "todo-api/models"
    "github.com/gorilla/mux"
)

func main() {
    // Инициализация базы данных
    database.InitDB()

    // Создание маршрутизатора
    r := mux.NewRouter()

    // Маршруты
    r.HandleFunc("/todos", getTodos).Methods("GET")
    r.HandleFunc("/todos/{id}", getTodo).Methods("GET")
    r.HandleFunc("/todos", createTodo).Methods("POST")
    r.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
    r.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

    // Запуск сервера
    log.Println("Сервер запущен на :8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}

// Обработчик для получения всех задач
func getTodos(w http.ResponseWriter, r *http.Request) {
    todos, err := database.GetTodos()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}

// Обработчик для получения задачи по ID
func getTodo(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    todo, err := database.GetTodo(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todo)
}

// Обработчик для создания новой задачи
func createTodo(w http.ResponseWriter, r *http.Request) {
    var todo models.Todo
    _ = json.NewDecoder(r.Body).Decode(&todo)

    id, err := database.CreateTodo(todo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    todo.ID = int(id)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todo)
}

// Обработчик для обновления задачи
func updateTodo(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var todo models.Todo
    _ = json.NewDecoder(r.Body).Decode(&todo)

    err := database.UpdateTodo(id, todo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

// Обработчик для удаления задачи
func deleteTodo(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    err := database.DeleteTodo(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}