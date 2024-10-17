package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Patient struct {
	ID        int
	FirstName string
	LastName  string
	Diagnosis string
	LastVisit string
	Status    string
}

// Middleware to enable CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Измените "*" на конкретные источники в продакшене
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, hx-request, hx-current-url")

		// Если это preflight request, возвращаем 200 и ничего не обрабатываем
		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func sendAlertHandler(w http.ResponseWriter, r *http.Request) {
	// Пример обработчика
	fmt.Fprintln(w, "Alert sent!")
}

func main() {
	mux := http.NewServeMux()
	// Пример данных о пациентах
	patients := []Patient{
		{ID: 1, FirstName: "Иван", LastName: "Иванов", Diagnosis: "Рак легких", LastVisit: "2024-01-10", Status: "В процессе лечения"},
		{ID: 2, FirstName: "Мария", LastName: "Петрова", Diagnosis: "Рак молочной железы", LastVisit: "2024-01-05", Status: "Завершено"},
		{ID: 3, FirstName: "Алексей", LastName: "Сидоров", Diagnosis: "Рак простаты", LastVisit: "2024-01-15", Status: "В процессе лечения"},
	}

	// Настройка маршрутизации
	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./src/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, patients)
	})
	mux.Handle("/src/static/", http.StripPrefix("/src/static/", http.FileServer(http.Dir("src/static"))))

	// Используем middleware для включения CORS
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", enableCORS(mux))
}
