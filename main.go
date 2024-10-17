package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Patient struct {
	ID           int
	FirstName    string
	LastName     string
	BirthDate    string
	Diagnosis    string
	PlansTherapy []PlanTherapy
}

type PlanTherapy struct {
	ID          int
	StartDate   string
	FinishDate  string
	Description string
	SideEffect  string
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

func main() {
	mux := http.NewServeMux()
	// Пример данных о пациентах
	patients := []Patient{
		{
			ID:        1,
			FirstName: "Иван",
			LastName:  "Иванов",
			BirthDate: "1980-05-01",
			Diagnosis: "Рак легких",
			PlansTherapy: []PlanTherapy{
				{ID: 1, StartDate: "2024-02-01", FinishDate: "2024-03-01", Description: "Химиотерапия", SideEffect: "Усталость"},
				{ID: 2, StartDate: "2024-03-15", FinishDate: "2024-04-15", Description: "Лучевая терапия", SideEffect: "Тошнота"},
			},
		},
		{
			ID:        2,
			FirstName: "Мария",
			LastName:  "Петрова",
			BirthDate: "1975-10-10",
			Diagnosis: "Рак молочной железы",
			PlansTherapy: []PlanTherapy{
				{ID: 1, StartDate: "2024-01-15", FinishDate: "2024-02-15", Description: "Хирургическое вмешательство", SideEffect: "Боль"},
			},
		},
		{
			ID:        3,
			FirstName: "Алексей",
			LastName:  "Сидоров",
			BirthDate: "1965-03-20",
			Diagnosis: "Рак простаты",
			PlansTherapy: []PlanTherapy{
				{ID: 1, StartDate: "2024-01-25", FinishDate: "2024-03-25", Description: "Гормональная терапия", SideEffect: "Головные боли"},
				{ID: 2, StartDate: "2024-04-01", FinishDate: "2024-05-01", Description: "Симптоматическая терапия", SideEffect: "Отечность"},
			},
		},
	}

	// Настройка маршрутизации
	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./src/static/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, patients)
	})

	mux.Handle("/src/static/", http.StripPrefix("/src/static/", http.FileServer(http.Dir("./src/static"))))

	// Используем middleware для включения CORS
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", enableCORS(mux))
}
