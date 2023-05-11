package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/veyselaksin/Golang-Tutorial/database"
	"github.com/veyselaksin/Golang-Tutorial/models"
	"gorm.io/gorm"
)

type User struct {
	// This is not the model User, see this as the serializer
	UUID      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type IUser interface {
	CreateUser(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
}

type user struct {
	db *gorm.DB
}

var _ IUser = &user{}

func NewUser(db *gorm.DB) IUser {
	return &user{db: db}
}

func CreateResponseUser(userModel models.User) User {
	return User{
		UUID:      userModel.UUID.String(),
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
}

func (u user) CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	a, _ := u.db.DB()

	var max int64 = 2000
	for i := 1; i < int(max); i++ {
		user.UUID, _ = uuid.NewUUID()
		fmt.Println("Connection Open Connection : ", a.Stats().OpenConnections)
		fmt.Println("Connection Max Connection : ", a.Stats().MaxOpenConnections)
		fmt.Println("Idle Connection Size : ", a.Stats().Idle)
		fmt.Println("Max Idle Closed : ", a.Stats().MaxIdleClosed)
		fmt.Println("Max Life Time Closed : ", a.Stats().MaxLifetimeClosed)
		fmt.Println("Wait Count", a.Stats().WaitCount)
		fmt.Println()
		u.db.Create(&user)
	}

	// fmt.Println("Connection Open Connection 1: ", a.Stats().OpenConnections)
	// fmt.Println("Connection Max Connection 1: ", a.Stats().MaxOpenConnections)
	// fmt.Println("Idle Connection Size 1: ", a.Stats().Idle)
	// fmt.Println("Max Idle Closed 1: ", a.Stats().MaxIdleClosed)

	// u.db.Save(&user)
	// fmt.Println("Connection Open Connection 2: ", a.Stats().OpenConnections)
	// fmt.Println("Connection Max Connection 2: ", a.Stats().MaxOpenConnections)
	// fmt.Println("Idle Connection Size 2: ", a.Stats().Idle)
	// fmt.Println("Max Idle Closed 2: ", a.Stats().MaxIdleClosed)

	// tx := u.db.Begin()
	// tx.SavePoint("sp")

	// flag := true
	// if flag {
	// 	u.db.Find(&user)
	// }

	// fmt.Println("Connection Open Connection 3: ", a.Stats().OpenConnections)
	// fmt.Println("Connection Max Connection 3: ", a.Stats().MaxOpenConnections)
	// fmt.Println("Idle Connection Size 3: ", a.Stats().Idle)
	// fmt.Println("Max Idle Closed 3: ", a.Stats().MaxIdleClosed)

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func (u user) GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.ConnectDB().Find(&users)

	responseUsers := []User{}

	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}
