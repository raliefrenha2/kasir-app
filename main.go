package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

var products = []Product{
	{ID: 1, Name: "Indomie Kari Ayam", Description: "Indomie Kari Ayam", Price: 3000, Stock: 100},
	{ID: 2, Name: "Indomie Goreng", Description: "Indomie Goreng", Price: 3000, Stock: 100},
	{ID: 3, Name: "Indomie Soto", Description: "Indomie Soto", Price: 3000, Stock: 100},
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var categories = []Category{
	{ID: 1, Name: "Food", Description: "Kategori Makanan"},
	{ID: 2, Name: "Drink", Description: "Kategori Minuman"},
	{ID: 3, Name: "Snack", Description: "Kategori Cemilan"},
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProductByID(w, r)
		} else if r.Method == "PUT" {
			updateProduct(w, r)
		} else if r.Method == "DELETE" {
			deleteProduct(w, r)
		}
	})

	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products)
		} else if r.Method == "POST" {
			var newProduct Product
			err := json.NewDecoder(r.Body).Decode(&newProduct)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			newProduct.ID = len(products) + 1
			products = append(products, newProduct)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(newProduct)
		}
	})

	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getCategoryByID(w, r)
		} else if r.Method == "PUT" {
			updateCategory(w, r)
		} else if r.Method == "DELETE" {
			deleteCategory(w, r)
		}
	})

	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categories)
		} else if r.Method == "POST" {
			var newCategory Category
			err := json.NewDecoder(r.Body).Decode(&newCategory)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			newCategory.ID = len(categories) + 1
			categories = append(categories, newCategory)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(newCategory)
		}
	})

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}

}

func getProductByID(w http.ResponseWriter, r *http.Request) {
	// Parse ID from URL
	// URL : /api/products/1
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Find product by ID
	for _, product := range products {
		if product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	// Product not found
	http.Error(w, "Product not found", http.StatusNotFound)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	// Parse ID from URL
	// URL : /api/products/1
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// get data from body
	var updatedProduct Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find product by ID
	// for i, product := range products {
	// 	if product.ID == id {
	// 		var updatedProduct Product
	// 		err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	// 		if err != nil {
	// 			http.Error(w, err.Error(), http.StatusBadRequest)
	// 			return
	// 		}

	// 		updatedProduct.ID = id
	// 		products[i] = updatedProduct
	// 		w.Header().Set("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode(updatedProduct)
	// 		return
	// 	}
	// }

	// loop using index
	for i := range products {
		if products[i].ID == id {
			updatedProduct.ID = id
			products[i] = updatedProduct

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedProduct)
			return
		}
	}

	// Product not found
	http.Error(w, "Product not found", http.StatusNotFound)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")

	// convert string to int
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Find product by ID
	for i := range products {
		if products[i].ID == id {

			// remove product from slice
			products = append(products[:i], products[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Product deleted successfully",
			})
			return
		}
	}

	// Product not found
	http.Error(w, "Product not found", http.StatusNotFound)
}

func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	// Parse ID from URL
	// URL : /api/categories/1
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// Find category by ID
	for _, category := range categories {
		if category.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(category)
			return
		}
	}

	// Category not found
	http.Error(w, "Category not found", http.StatusNotFound)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	// Parse ID from URL
	// URL : /api/categories/1
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// get data from body
	var updatedCategory Category
	err = json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find category by ID
	for i := range categories {
		if categories[i].ID == id {
			updatedCategory.ID = id
			categories[i] = updatedCategory

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedCategory)
			return
		}
	}

	// Category not found
	http.Error(w, "Category not found", http.StatusNotFound)
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

	// convert string to int
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// Find category by ID
	for i := range categories {
		if categories[i].ID == id {

			// remove category from slice
			categories = append(categories[:i], categories[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Category deleted successfully",
			})
			return
		}
	}

	// Category not found
	http.Error(w, "Category not found", http.StatusNotFound)
}
