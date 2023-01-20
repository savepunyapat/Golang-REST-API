package routes

import (
	"context"
	db "example/user/hello/app"
	collection "example/user/hello/collections"
	model "example/user/hello/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateData(c *gin.Context) {
	fmt.Print("Starting create data...")
	var DB = db.ConnectDB()
	var movieCollection = collection.GetCollection(DB, "mock1")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	movie := new(model.Movie)
	defer cancel()

	if err := c.BindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		fmt.Println("err1")
		log.Fatal(err)
		return
	}

	moviePayload := model.Movie{
		ID:    primitive.NewObjectID(),
		Title: movie.Title,
		Name:  movie.Name,
	}

	result, err := movieCollection.InsertOne(ctx, moviePayload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}})
}

func UpdateData(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = db.ConnectDB()
	var movieCollection = collection.GetCollection(DB, "mock1")

	movieId := c.Param("movieID")
	var movie model.Movie

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(movieId)

	if err := c.BindJSON(&movie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	edited := bson.M{"Title": movie.Title, "Name": movie.Name}

	result, err := movieCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": edited})

	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if result.MatchedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data doesn't exist"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "data updated successfully!", "Data": res})
}

func ReadOne(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = db.ConnectDB()
	var movieCollection = collection.GetCollection(DB, "mock1")
	movieId := c.Param("movieID")
	var result model.Movie
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(movieId)
	err := movieCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&result)
	res := map[string]interface{}{"data": result}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success!", "Data": res})
}

func DeletePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = db.ConnectDB()
	movieId := c.Param("movieID")

	var movieCollection = collection.GetCollection(DB, "mock1")
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(movieId)
	result, err := movieCollection.DeleteOne(ctx, bson.M{"id": objId})
	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No data to delete"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Article deleted successfully", "Data": res})
}

func ReadAll(c *gin.Context) {

	var DB = db.ConnectDB()
	var movieCollection = collection.GetCollection(DB, "mock1")
	filter := bson.D{{"Comm", 210}}
	result, err := movieCollection.Find(context.TODO(), filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	res := map[string]interface{}{"data": result}

	c.JSON(http.StatusCreated, gin.H{"message": "Read All Data", "Data": res})
}
