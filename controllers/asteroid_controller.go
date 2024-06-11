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
	"go.mongodb.org/mongo-driver/bson"
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
		Distances:     asteroid.Distances,
	}

	result, err := AsteroidCollection.InsertOne(ctx, newAsteroid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.AsteroidResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, responses.AsteroidResponse{Status: http.StatusCreated, Message: "Success", Data: &echo.Map{"data": result}})
}

func GetAsteroids(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var asteroids []models.Asteroid
	defer cancel()

	results, err := AsteroidCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.AsteroidResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &echo.Map{"data": err.Error()}})
	}

	//reading from db
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleAsteroid models.Asteroid
		if err = results.Decode(&singleAsteroid); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.AsteroidResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &echo.Map{"data": err.Error()}})
		}

		asteroids = append(asteroids, singleAsteroid)
	}

	return c.JSON(http.StatusOK, responses.AsteroidResponse{Status: http.StatusOK, Message: "Success", Data: &echo.Map{"data": asteroids}})
}

func GetAsteroid(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	asteroidID := c.Param("id")
	var asteroid models.Asteroid
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(asteroidID)

	err := AsteroidCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&asteroid)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.AsteroidResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.AsteroidResponse{Status: http.StatusOK, Message: "Success", Data: &echo.Map{"data": asteroid}})
}

func UpdateAsteroid(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	asteroidID := c.Param("id")
	var asteroid models.Asteroid
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(asteroidID)

	//validate request
	if err := c.Bind(&asteroid); err != nil {
		return c.JSON(http.StatusBadRequest, responses.AsteroidResponse{Status: http.StatusBadRequest, Message: "Error", Data: &echo.Map{"data": err.Error()}})
	}

	//validate required fields
	if validationErr := validate.Struct(&asteroid); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.AsteroidResponse{Status: http.StatusBadRequest, Message: "Error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"name": asteroid.Name, "diameter": asteroid.Diameter, "discovery_date": asteroid.DiscoveryDate, "observations": asteroid.Observations, "distances": asteroid.Distances}

	result, err := AsteroidCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.AsteroidResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &echo.Map{"data": err.Error()}})
	}

	//get upddated asteroid details
	var updatedAsteroid models.Asteroid
	if result.MatchedCount == 1 {
		err := AsteroidCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedAsteroid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.AsteroidResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &echo.Map{"data": err.Error()}})
		}
	}

	return c.JSON(http.StatusOK, responses.AsteroidResponse{Status: http.StatusOK, Message: "Success", Data: &echo.Map{"data": updatedAsteroid}})
}

func DeleteAsteroid(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	asteroidID := c.Param("id")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(asteroidID)

	result, err := AsteroidCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.AsteroidResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &echo.Map{"data": err.Error()}})
	}

	if result.DeletedCount == 0 {
		return c.JSON(http.StatusNotFound, responses.AsteroidResponse{Status: http.StatusNotFound, Message: "Error", Data: &echo.Map{"data": "Asteroid with specified ID not found"}})
	}

	return c.JSON(http.StatusOK, responses.AsteroidResponse{Status: http.StatusOK, Message: "Success", Data: &echo.Map{"data": "Asteroid deleted successfully"}})
}
