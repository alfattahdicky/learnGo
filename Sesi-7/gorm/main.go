package main

import (
	"errors"
	"fmt"
	"gorm/database"
	"gorm/models"

	"gorm.io/gorm"
)


func main()  {
	database.StartDB()

	createUser("dicky")

	// getUserById(1)

	// updateUserById(1, "azkafaiz@gmail.com")

	// createProduct(1, "QWE", "RERE")

	// getUsersWithProduct()
	// deleteProductById(1)
}

func createUser(customerName string)  {
	db := database.GetDb()
	
	Order := models.Order{
		Id: 1,
		CustomerName: customerName,
	}

	err := db.Create(&Order).Error

	if err != nil {
		fmt.Println("Error creating user data", err)

		return
	}

	fmt.Println("New User Data", Order)
}

func getUserById(id uint)  {
	db := database.GetDb()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User Data not found")
			return
		}
		print("Error finding user:", err)
	}

	fmt.Printf("User Data: %+v \n", user)
}

func updateUserById(id uint, email string)  {
	db := database.GetDb()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email : email}).Error

	if err != nil {
		fmt.Println("Error updating user data: ", err)
		return
	}

	fmt.Printf("Update user's email %+v \n", user.Email)
}

func createProduct(userId uint, brand string, name string)  {
	db := database.GetDb()

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

	fmt.Println("New Product Data", Product)
}

// Eager Loading => join statement dari Gorm menggunakan method preload
func getUsersWithProduct()  {
	db := database.GetDb()

	users := models.User{}

	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user datas with products", err.Error())
		return
	}

	fmt.Println("User Datas With Products")
	fmt.Printf("%+v", users)

}

func deleteProductById(id uint)  {
	db := database.GetDb()

	product := models.Product{}

	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil  {
		fmt.Println("Error deleting product:", err.Error())
		return
	}

	fmt.Printf("Product with id %d has been successfully deleted \n", id)
}