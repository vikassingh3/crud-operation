package controller

import (
	"os"
	"path/database"
	"path/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllData(c *fiber.Ctx) error {

	ConnectDb := database.MI.DB.Collection(os.Getenv("DATABASE_COLLECTION"))

	queary := bson.D{{}}

	curser, _ := ConnectDb.Find(c.Context(), queary)

	//fmt.Printf("%T", curser)

	var data []models.User = make([]models.User, 0)

	curser.All(c.Context(), &data)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": data,
	})

}

func CreateData(c *fiber.Ctx) error {

	ConnectDb := database.MI.DB.Collection(os.Getenv("DATABASE_COLLECTION"))

	data := new(models.User)

	c.BodyParser(&data)

	data.ID = nil
	data.Created_At = time.Now()

	result, _ := ConnectDb.InsertOne(c.Context(), data)

	///////////////////////////////////////////////

	content := &models.User{}

	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	ConnectDb.FindOne(c.Context(), query).Decode(content)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"meassage": "successfull",
		"data": fiber.Map{
			"data": content,
		},
	})

}

func GetOneData(c *fiber.Ctx) error {

	ConnectDb := database.MI.DB.Collection(os.Getenv("DATABASE_COLLECTION"))

	getId := c.Params("id")

	id, _ := primitive.ObjectIDFromHex(getId)

	data := &models.User{}

	query := bson.D{{Key: "_id", Value: id}}

	ConnectDb.FindOne(c.Context(), query).Decode(data)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"data": data,
		},
	})

}

func DeleteOneData(c *fiber.Ctx) error {

	ConnectDb := database.MI.DB.Collection(os.Getenv("DATABASE_COLLECTION"))

	getId := c.Params("id")

	id, _ := primitive.ObjectIDFromHex(getId)

	query := bson.D{{Key: "_id", Value: id}}

	ConnectDb.FindOneAndDelete(c.Context(), query).Err()

	return c.SendStatus(fiber.StatusNoContent)

}

func UpdateData(c *fiber.Ctx) error {

	ConnectDB := database.MI.DB.Collection(os.Getenv("DATABASE_COLLECTION"))

	getId := c.Params("id")

	id, _ := primitive.ObjectIDFromHex(getId)

	data := new(models.User)
	c.BodyParser(&data)

	query := bson.D{{Key: "id", Value: id}}

	var DataUpdate bson.D

	if data.Name != nil {
		DataUpdate = append(DataUpdate, bson.E{Key: "name", Value: data.Name})
	}

	if data.Designation != nil {
		DataUpdate = append(DataUpdate, bson.E{Key: "designation", Value: data.Designation})

	}

	if data.Age != nil {
		DataUpdate = append(DataUpdate, bson.E{Key: "age", Value: data.Age})

	}

	update := bson.D{{Key: "$key", Value: DataUpdate}}

	ConnectDB.FindOneAndUpdate(c.Context(), query, update).Err()

	data1 := &models.User{}

	ConnectDB.FindOne(c.Context(), query).Decode(data1)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfull",
		"data": fiber.Map{
			"data": data1,
		},
	})

}
