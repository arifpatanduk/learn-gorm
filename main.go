package main

import (
	"errors"
	"fmt"
	"learn-gorm/database"
	"learn-gorm/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	// createUser("john@gmail.com")
	// getUserById(1)
	// updateUserById(1, "anjoy@gmail.com")

	// createProduct(1, "UHT", "ultra milk")
	// getUsersWithProducts()
	// deleteProductById(1)
	getAllUser()
}

func createUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}

	fmt.Println("New user data:", User)
}

func getAllUser() {
	db := database.GetDB()
	users := []models.User{}

	err := db.Find(&users).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}

		print("Error finding user:", err)
	}

	fmt.Println("User data:")
	for _, user := range users {
		fmt.Printf("%+v \n", user)
	}

}

func getUserById (id uint) {
	db := database.GetDB()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}

		print("Error finding user:", err)
	}

	fmt.Println("User data : %v \n", user)
}

func updateUserById (id uint, email string) {
	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error updating user data:", err)
		return
	}

	fmt.Println("Update user's email : %v \n", user.Email)
}

func createProduct(userId uint, brand string, name string) {
	db := database.GetDB()

	Product := models.Product{
		UserID: userId,
		Brand: brand,
		Name: name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating product data:", err.Error())
		return
	}

	fmt.Println("New product data:", Product)

}

func getUsersWithProducts() {
	db := database.GetDB()

	users := models.User{}
	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user data with products:", err.Error())
		return
	}

	fmt.Println("User Datas with Products")
	fmt.Println("%+v", users)
}

func deleteProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}
	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error deleting product:", err.Error())
		return
	}

	fmt.Printf("Product with id %d has been deleted", id)
}