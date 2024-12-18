package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:100;not null"`
	Email string `gorm:"size:100;unique;not null"`
}

type RequestPayload struct {
	Message string `json:"message"`
}

type ResponsePayload struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	// Убедитесь, что запрос POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Чтение тела запроса
	var reqData RequestPayload
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		sendErrorResponse(w, "Некорректное JSON-сообщение")
		return
	}

	// Проверка поля "message"
	if reqData.Message == "" {
		sendErrorResponse(w, "Некорректное JSON-сообщение")
		return
	}

	// Логирование сообщения
	fmt.Printf("Получено сообщение: %s\n", reqData.Message)

	// Успешный ответ
	sendSuccessResponse(w, "Данные успешно приняты")
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	// Убедитесь, что запрос GET
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Возврат тестового ответа
	sendSuccessResponse(w, "Это тестовый GET-ответ")
}

func sendSuccessResponse(w http.ResponseWriter, message string) {
	response := ResponsePayload{
		Status:  "success",
		Message: message,
	}
	sendJSONResponse(w, http.StatusOK, response)
}

func sendErrorResponse(w http.ResponseWriter, message string) {
	response := ResponsePayload{
		Status:  "fail",
		Message: message,
	}
	sendJSONResponse(w, http.StatusBadRequest, response)
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, data ResponsePayload) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func connectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=1234 dbname=userdatago port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Автоматическое создание таблиц
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	fmt.Println("Подключение к базе данных успешно")
	return db
}

func createUser(db *gorm.DB, name, email string) {
	user := User{Name: name, Email: email}
	result := db.Create(&user)
	if result.Error != nil {
		log.Printf("Ошибка создания пользователя: %v", result.Error)
		return
	}
	fmt.Printf("Пользователь создан: %v\n", user)
}

func getUserByID(db *gorm.DB, id uint) {
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		log.Printf("Ошибка получения пользователя: %v", result.Error)
		return
	}
	fmt.Printf("Получен пользователь: %v\n", user)
}
func updateUserByID(db *gorm.DB, id uint, newName, newEmail string) {
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		log.Printf("Ошибка получения пользователя для обновления: %v", result.Error)
		return
	}

	user.Name = newName
	user.Email = newEmail
	db.Save(&user)

	fmt.Printf("Пользователь обновлен: %v\n", user)
}

func deleteUserByID(db *gorm.DB, id uint) {
	result := db.Delete(&User{}, id)
	if result.Error != nil {
		log.Printf("Error deleting user: %v", result.Error)
		return
	}
	if result.RowsAffected == 0 {
		fmt.Printf("No user found with ID %d\n", id)
		return
	}
	fmt.Printf("User with ID %d deleted\n", id)
}

func listUsers(db *gorm.DB) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		log.Printf("Ошибка получения списка пользователей: %v", result.Error)
		return
	}
	fmt.Printf("Список пользователей: %v\n", users)
}

func main() {
	db := connectDB()

	deleteUserByID(db, 1)
	deleteUserByID(db, 2)
	listUsers(db)

	http.HandleFunc("/post", handlePost)
	http.HandleFunc("/get", handleGet)
	http.Handle("/", http.FileServer(http.Dir("./")))

	fmt.Println("Сервер запущен на http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
