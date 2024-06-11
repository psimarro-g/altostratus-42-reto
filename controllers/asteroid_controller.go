package controllers

import (
	"context"
	"net/http"
	"time"

	"altostratus-42-reto/configs"
	"altostratus-42-reto/models"
	"altostratus-42-reto/responses"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var AsteroidCollection *mongo.Collection = configs.GetCollection(configs.DB, "asteroids")
var validate = validator.New()

func CreateAsteroid(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var asteroid models.Asteroid
	defer cancel()

	//validate request
	if err := c.Bind(&asteroid); err != nil {
		return c.JSON(http.StatusBadRequest, responses.AsteroidResponse{Status: http.StatusBadRequest, Message: "Error", Data: &echo.Map{"data": err.Error()}})
	}

	//validate required fields
	if validationErr := validate.Struct(&asteroid); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.AsteroidResponse{Status: http.StatusBadRequest, Message: "Error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	newAsteroid := models.Asteroid{
		ID:            primitive.NewObjectID(),
		Name:          asteroid.Name,
		Diameter:      asteroid.Diameter,
		DiscoveryDate: asteroid.DiscoveryDate,
		Observations:  asteroid.Observations,
		//Distances:     asteroid.Distances,
	}

	result, err := AsteroidCollection.InsertOne(ctx, newAsteroid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.AsteroidResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, responses.AsteroidResponse{Status: http.StatusCreated, Message: "Success", Data: &echo.Map{"data": result}})
}
