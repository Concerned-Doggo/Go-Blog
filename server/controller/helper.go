package controller

import (
	"context"
	"log"

	"github.com/Concerned-Doggo/Go-React-Blog/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)


func checkNilErr(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func createBlog(blog model.Blog) (bson.ObjectID){

    createResult, err := collection.InsertOne(context.Background(), blog)
    checkNilErr(err, "unable to create blog")

    return createResult.InsertedID.(bson.ObjectID)
}

func getRecentBlogs(limit int) []model.Blog {
	var blogs []model.Blog

	cursor, err := collection.Find(context.Background(), bson.D{})
	checkNilErr(err, "unable to find blogs")
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		// just for getting for say 10 or 20 recent blogs written
		// if there are less blogs no problem
		if limit == len(blogs) {
			break
		}
		var blog model.Blog

		err := cursor.Decode(&blog)
		checkNilErr(err, "unable to decode blog")

		blogs = append(blogs, blog)
	}

	return blogs
}

func getBlogById(blogId string) model.Blog {
	_id, err := bson.ObjectIDFromHex(blogId)
	checkNilErr(err, "unable to convert string to object id")

	var blog model.Blog
	err = collection.FindOne(context.Background(), bson.M{"_id": _id}).Decode(&blog)
	checkNilErr(err, "unable to find blog")

	return blog
}

func updateBlog(blogId string, blog model.Blog) (int64) {
	_id, err := bson.ObjectIDFromHex(blogId)
	checkNilErr(err, "unable to convert string to object id")

	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{
		"title": blog.Title,
		"body":  blog.Body,
	}}

    updatedResult, err := collection.UpdateOne(context.Background(), filter, update)
    checkNilErr(err, "unable to update blog")

    return updatedResult.ModifiedCount
}

func deleteBlogById(blogId string) (int64) {
    _id, err := bson.ObjectIDFromHex(blogId)
    checkNilErr(err, "unable to convert string to object id")

    deleteResult, err := collection.DeleteOne(context.Background(), bson.M{"_id":_id})
    checkNilErr(err, "unable to delete one blog")

    return deleteResult.DeletedCount
}


