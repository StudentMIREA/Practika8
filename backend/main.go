package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Product struct {
	ID          int
	Name        string
	Image       string
	Cost        float64
	Describtion string
	Favorite    bool
	ShopCart    bool
	Count       int
}

type User struct {
	ID    int
	Image string
	Name  string
	Phone string
	Mail  string
}

var products = []Product{
	{
		ID:          1,
		Name:        "Кружка \"Гусь\"",
		Image:       "https://ir.ozone.ru/s3/multimedia-1-3/wc1000/7110855183.jpg",
		Cost:        600,
		Describtion: "Материал: Полистоун\nРазмер: 380мл\nМожно мыть в посудомоечной машине.",
		Favorite:    true,
		ShopCart:    true,
		Count:       0,
	},
	{
		ID:          2,
		Name:        "Кружка керамическая",
		Image:       "https://ir.ozone.ru/s3/multimedia-1-j/wc1000/7075750735.jpg",
		Cost:        500,
		Describtion: "Материал: Керамика\nРазмер: 400мл\nМожно мыть в посудомоечной машине.",
		Favorite:    true,
		ShopCart:    true,
		Count:       0,
	},
	{
		ID:          3,
		Name:        "Кружка Доляна Клубника",
		Image:       "https://ir.ozone.ru/s3/multimedia-h/wc1000/6785147933.jpg",
		Cost:        600,
		Describtion: "Материал: Керамика, Силикон\nРазмер: 350мл\nМожно мыть в посудомоечной машине.",
		Favorite:    false,
		ShopCart:    false,
		Count:       0,
	},
	{
		ID:          4,
		Name:        "Кружка хамелеон Чеширский Кот",
		Image:       "https://ir.ozone.ru/s3/multimedia-0/wc1000/6294798312.jpg",
		Cost:        400,
		Describtion: "Материал: Керамика\nРазмер: 300мл\nОсобенность этой кружки в том, что она меняет свой цвет при нагревании. Можно мыть в посудомоечной машине.",
		Favorite:    false,
		ShopCart:    false,
		Count:       0,
	},
	{
		ID:          5,
		Name:        "Кружка металлическая \"Череп\"",
		Image:       "https://ir.ozone.ru/s3/multimedia-h/wc1000/6304198541.jpg",
		Cost:        400,
		Describtion: "Материал: Нержавеющая сталь\nРазмер: 400мл\nНельзя мыть в посудомоечной машине.",
		Favorite:    false,
		ShopCart:    false,
		Count:       0,
	},
	{
		ID:          6,
		Name:        "Кружка керамическая Доляна Twist",
		Image:       "https://ir.ozone.ru/s3/multimedia-1-j/wc1000/7090710499.jpg",
		Cost:        400,
		Describtion: "Материал: Керамика\nРазмер: 350мл\nНельзя мыть в посудомоечной машине.",
		Favorite:    false,
		ShopCart:    false,
		Count:       0,
	},
	{
		ID:          7,
		Name:        "Кружка фарфоровая \"Шины\"",
		Image:       "https://ir.ozone.ru/s3/multimedia-1-9/wc1000/7092772317.jpg",
		Cost:        500,
		Describtion: "Материал: Керамика\nРазмер: 460мл\nМожно мыть в посудомоечной машине.",
		Favorite:    false,
		ShopCart:    false,
		Count:       0,
	},
	{
		ID:          8,
		Name:        "Кружка с двойными стенками \"Cердце\"",
		Image:       "https://ir.ozone.ru/s3/multimedia-j/wc1000/6808645387.jpg",
		Cost:        400,
		Describtion: "Материал: Стекло, Боросиликатное стекло\nРазмер: 250мл\nМожно мыть в посудомоечной машине.",
		Favorite:    false,
		ShopCart:    false,
		Count:       0,
	},
	{
		ID:          9,
		Name:        "Starbucks Кружка",
		Image:       "https://ir.ozone.ru/s3/multimedia-1-8/wc1000/7107265016.jpg",
		Cost:        3600,
		Describtion: "Материал: Керамика\nРазмер: 355мл\nВращающееся, Двойные стенки. Можно мыть в посудомоечной машине.",
		Favorite:    false,
		ShopCart:    false,
		Count:       0,
	},
	{
		ID:          10,
		Name:        "Кружка керамическая Mushroom",
		Image:       "https://ir.ozone.ru/s3/multimedia-m/wc1000/6632228578.jpg",
		Cost:        1600,
		Describtion: "Материал: Керамика\nРазмер: 300мл\nКружка керамическая Mushroom Гриб из Супер Марио. Можно мыть в посудомоечной машине.",
		Favorite:    false,
		ShopCart:    false,
		Count:       0,
	},
}

// обработчик для GET-запроса, возвращает список продуктов
func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовки для правильного формата JSON
	w.Header().Set("Content-Type", "application/json")
	// Преобразуем список заметок в JSON
	json.NewEncoder(w).Encode(products)
}

// обработчик для POST-запроса, добавляет продукт
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Received new Product: %+v\n", newProduct)

	newProduct.ID = len(products) + 1
	products = append(products, newProduct)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)
}

//Добавление маршрута для получения одного продукта

func getProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Ищем продукт с данным ID
	for _, Product := range products {
		if Product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Product)
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

// удаление продукта по id
func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Ищем и удаляем продукт с данным ID
	for i, Product := range products {
		if Product.ID == id {
			// Удаляем продукт из среза
			products = append(products[:i], products[i+1:]...)
			w.WriteHeader(http.StatusNoContent) // Успешное удаление, нет содержимого
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

// Обновление продукта по id
func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/update/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Декодируем обновлённые данные продукта
	var updatedProduct Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	fmt.Println(updatedProduct.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Ищем продукт для обновления
	for i, Product := range products {
		if Product.ID == id {
			products[i].Name = updatedProduct.Name
			products[i].Image = updatedProduct.Image
			products[i].Cost = updatedProduct.Cost
			products[i].Describtion = updatedProduct.Describtion
			products[i].Favorite = updatedProduct.Favorite
			products[i].ShopCart = updatedProduct.ShopCart
			products[i].Count = updatedProduct.Count

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products[i])
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

func getFavoriteProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовки для правильного формата JSON
	w.Header().Set("Content-Type", "application/json")
	// Преобразуем список заметок в JSON
	var favoriteProducts = []Product{}

	for _, Product := range products {
		if Product.Favorite == true {
			w.Header().Set("Content-Type", "application/json")
			favoriteProducts = append(favoriteProducts, Product)
		}
	}

	json.NewEncoder(w).Encode(favoriteProducts)
}

func getShopCartProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовки для правильного формата JSON
	w.Header().Set("Content-Type", "application/json")
	// Преобразуем список заметок в JSON
	var shopCartProducts = []Product{}

	for _, Product := range products {
		if Product.ShopCart == true {
			w.Header().Set("Content-Type", "application/json")
			shopCartProducts = append(shopCartProducts, Product)
		}
	}

	json.NewEncoder(w).Encode(shopCartProducts)
}

func main() {
	http.HandleFunc("/products", getProductsHandler)           // Получить все продукты
	http.HandleFunc("/products/create", createProductHandler)  // Создать продукт
	http.HandleFunc("/products/", getProductByIDHandler)       // Получить продукт по ID
	http.HandleFunc("/products/update/", updateProductHandler) // Обновить продукт
	http.HandleFunc("/products/delete/", deleteProductHandler) // Удалить продукт

	http.HandleFunc("/favorite_products", getFavoriteProductsHandler)
	http.HandleFunc("/shop_cart_products", getShopCartProductsHandler)

	fmt.Println("Server is running on port 8080!")
	http.ListenAndServe(":8080", nil)
}
