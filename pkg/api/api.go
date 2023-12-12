package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg/v10"
	"log"
	"net/http"
	"product-backend/pkg/db/models"
)

func StartAPI(pgdb *pg.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger, middleware.WithValue("DB", pgdb))

	r.Route("/product", func(r chi.Router) {
		r.Post("/", createProduct)
		r.Get("/", getProducts)
	})

	r.Route("/task", func(r chi.Router) {
		r.Post("/", createTask)
		r.Get("/", getTasks)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("up and running"))
	})

	return r
}

type createTaskRequest struct {
	Name     string `json:"name"`
	Duration int64  `json:"duration"`
}

type TaskResponse struct {
	Success bool         `json:"success"`
	Error   string       `json:"error"`
	Task    *models.Task `json:"task"`
}

func createTask(w http.ResponseWriter, r *http.Request) {
	req := &createTaskRequest{}
	err := json.NewDecoder(r.Body).Decode(req)

	if err != nil {
		res := &TaskResponse{
			Success: false,
			Error:   err.Error(),
			Task:    nil,
		}
		err = json.NewEncoder(w).Encode(res)

		if err != nil {
			log.Printf("error sending response %v\n", err)
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pgdb, ok := r.Context().Value("DB").(*pg.DB)

	if !ok {
		res := &TaskResponse{
			Success: false,
			Error:   "could not get the DB from context",
			Task:    nil,
		}
		err = json.NewEncoder(w).Encode(res)

		if err != nil {
			log.Printf("error sending response %v\n", err)
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task, err := models.CreateTask(pgdb, &models.Task{
		Name:     req.Name,
		Duration: req.Duration,
	})
	if err != nil {
		res := &TaskResponse{
			Success: false,
			Error:   err.Error(),
			Task:    nil,
		}
		err = json.NewEncoder(w).Encode(res)

		if err != nil {
			log.Printf("error sending response %v\n", err)
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := &TaskResponse{
		Success: true,
		Error:   "",
		Task:    task,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding after creating task %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type TasksResponse struct {
	Success bool           `json:"success"`
	Error   string         `json:"error"`
	Tasks   []*models.Task `json:"tasks"`
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &TasksResponse{
			Success: false,
			Error:   "could not get DB from context",
			Tasks:   nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tasks, err := models.GetTasks(pgdb)
	if err != nil {
		res := &TasksResponse{
			Success: false,
			Error:   err.Error(),
			Tasks:   nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := &TasksResponse{
		Success: true,
		Error:   "",
		Tasks:   tasks,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding tasks: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}


type createProductRequest struct {
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Weight      float64 `json:"weight"`
	Description string  `json:"description"`
}

type ProductResponse struct {
	Success   bool            `json:"success"`
	Error     string          `json:"error"`
	Product   *models.Product `json:"product"`
}

type ProductsResponse struct {
	Success  bool             `json:"success"`
	Error    string           `json:"error"`
	Products []*models.Product `json:"products"`
}


func createProduct(w http.ResponseWriter, r *http.Request) {
	req := &createProductRequest{}
	err := json.NewDecoder(r.Body).Decode(req)

	if err != nil {
		res := &ProductResponse{
			Success: false,
			Error:   err.Error(),
			Product: nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &ProductResponse{
			Success: false,
			Error:   "could not get the DB from context",
			Product: nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := models.CreateProduct(pgdb, &models.Product{
		Code:        req.Code,
		Name:        req.Name,
		Weight:      req.Weight,
		Description: req.Description,
	})
	if err != nil {
		res := &ProductResponse{
			Success: false,
			Error:   err.Error(),
			Product: nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := &ProductResponse{
		Success: true,
		Error:   "",
		Product: product,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding after creating product %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &ProductsResponse{
			Success: false,
			Error:   "could not get DB from context",
			Products: nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	products, err := models.GetProducts(pgdb)
	if err != nil {
		res := &ProductsResponse{
			Success: false,
			Error:   err.Error(),
			Products: nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := &ProductsResponse{
		Success: true,
		Error:   "",
		Products: products,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding products: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
