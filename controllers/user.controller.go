package controllers

import (
	"api/configs"
	"api/models"
	"api/responses"
	"context"
	"log"
	"net/http"
	"time"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var user models.User

		// validate body request
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Success: false, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		//use the validator library to validate required fields
        if validationErr := validate.Struct(&user); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Success: false, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }

		newUser := models.User {
            Name:     user.Name,
            Location: user.Location,
            Title:    user.Title,
        }

		userCollection := configs.GetCollection(configs.DB, "users")
		result, err := userCollection.InsertOne(ctx, newUser)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Success: false, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }
		c.JSON(http.StatusCreated, responses.UserResponse{Success: true, Message: "success", Data: map[string]interface{}{"user": result}})
	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		userCollection := configs.GetCollection(configs.DB, "users")
		var users []bson.M

		curror, err := userCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Success: false, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
			return
		}
		if err = curror.All(ctx, &users); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, responses.UserResponse{Success: true, Message: "success", Data: map[string]interface{}{"users": users}})
	}
}

func GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		
		id := c.Param("id")
		objId, _ := primitive.ObjectIDFromHex(id)

		userCollection := configs.GetCollection(configs.DB, "users")
		// var user models.User
		var user bson.M
		if err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Success: false, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		fmt.Println("result is", user)
        c.JSON(http.StatusOK, responses.UserResponse{Success: true, Message: "success", Data: map[string]interface{}{"user": user}})
	}
}


