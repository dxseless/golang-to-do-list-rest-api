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
    database.InitDB()

    r := mux.NewRouter()

    r.HandleFunc("/todos", getTodos).Methods("GET")
    r.HandleFunc("/todos/{id}", getTodo).Methods("GET")
    r.HandleFunc("/todos", createTodo).Methods("POST")
    r.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
    r.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

    log.Println("Сервер запущен на :8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
    filter := make(map[string]string)
    if status := r.URL.Query().Get("status"); status != "" {
        filter["status"] = status
    }
    if createdAfter := r.URL.Query().Get("created_after"); createdAfter != "" {
        filter["created_after"] = createdAfter
    }
    if createdBefore := r.URL.Query().Get("created_before"); createdBefore != "" {
        filter["created_before"] = createdBefore
    }
    if updatedAfter := r.URL.Query().Get("updated_after"); updatedAfter != "" {
        filter["updated_after"] = updatedAfter
    }
    if updatedBefore := r.URL.Query().Get("updated_before"); updatedBefore != "" {
        filter["updated_before"] = updatedBefore
    }

    todos, err := database.GetTodos(filter)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}

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

func createTodo(w http.ResponseWriter, r *http.Request) {
    var todo models.Todo
    _ = json.NewDecoder(r.Body).Decode(&todo)

    if todo.Status == "" {
        todo.Status = "active"
    }

    id, err := database.CreateTodo(todo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    todo.ID = int(id)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todo)
}

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