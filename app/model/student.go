package model

import (
	"errors"
	"fmt"

	"github.com/nahidhasan98/kgc-crud/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	ID          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Email       string `json:"email" bson:"email"`
	Phone       string `json:"phone" bson:"phone"`
	Reg         string `json:"reg" bson:"reg"`
	Session     string `json:"session" bson:"session"`
	Roll        string `json:"roll" bson:"roll"`
	PassingYear string `json:"passingYear" bson:"passingYear"`
	AvatarURL   string `json:"avatarURL" bson:"avatarURL"`
}

func GetAllStudents() ([]Student, error) {
	// connecting to DB
	DB, ctx, cancel := config.DBConnect()
	defer cancel()
	defer DB.Client().Disconnect(ctx)

	// taking DB collection/table to a variable
	studentCollection := DB.Collection("student")

	var students []Student
	opts := options.Find()
	opts.SetSort(bson.D{{Key: "name", Value: 1}})
	cursor, err := studentCollection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var temp Student
		err := cursor.Decode(&temp)
		if err != nil {
			fmt.Println(err)
		}

		students = append(students, temp)
	}

	return students, nil
}

func GetSingleStudentByID(studentID string) (*Student, error) {
	// connecting to DB
	DB, ctx, cancel := config.DBConnect()
	defer cancel()
	defer DB.Client().Disconnect(ctx)

	// taking DB collection/table to a variable
	studentCollection := DB.Collection("student")

	var student Student
	err := studentCollection.FindOne(ctx, bson.M{"_id": studentID}).Decode(&student)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("no document found with this id")
	} else if err != nil {
		return nil, err
	}

	return &student, nil
}
func GetSingleStudentByReg(reg string) (*Student, error) {
	// connecting to DB
	DB, ctx, cancel := config.DBConnect()
	defer cancel()
	defer DB.Client().Disconnect(ctx)

	// taking DB collection/table to a variable
	studentCollection := DB.Collection("student")

	var student Student
	err := studentCollection.FindOne(ctx, bson.M{"reg": reg}).Decode(&student)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("no document found with this registration number")
	} else if err != nil {
		return nil, err
	}

	return &student, nil
}
func GetSingleStudentByRoll(roll string) (*Student, error) {
	// connecting to DB
	DB, ctx, cancel := config.DBConnect()
	defer cancel()
	defer DB.Client().Disconnect(ctx)

	// taking DB collection/table to a variable
	studentCollection := DB.Collection("student")

	var student Student
	err := studentCollection.FindOne(ctx, bson.M{"roll": roll}).Decode(&student)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("no document found with this class roll")
	} else if err != nil {
		return nil, err
	}

	return &student, nil
}
func GetSingleStudentByEmail(email string) (*Student, error) {
	// connecting to DB
	DB, ctx, cancel := config.DBConnect()
	defer cancel()
	defer DB.Client().Disconnect(ctx)

	// taking DB collection/table to a variable
	studentCollection := DB.Collection("student")

	var student Student
	err := studentCollection.FindOne(ctx, bson.M{"email": email}).Decode(&student)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("no document found with this email")
	} else if err != nil {
		return nil, err
	}

	return &student, nil
}

func CreateStudent(studentData *Student) error {
	// connecting to DB
	DB, ctx, cancel := config.DBConnect()
	defer cancel()
	defer DB.Client().Disconnect(ctx)

	// taking DB collection/table to a variable
	studentCollection := DB.Collection("student")

	// getting original password for this user from DB
	_, err := studentCollection.InsertOne(ctx, studentData)
	if err != nil {
		return err
	}

	return nil
}

func UpdateStudent(studentID string, studentData *Student) error {
	// connecting to DB
	DB, ctx, cancel := config.DBConnect()
	defer cancel()
	defer DB.Client().Disconnect(ctx)

	// taking DB collection/table to a variable
	studentCollection := DB.Collection("student")

	updateField := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: studentData.Name},
			{Key: "session", Value: studentData.Session},
			{Key: "reg", Value: studentData.Reg},
			{Key: "roll", Value: studentData.Roll},
			{Key: "passingYear", Value: studentData.PassingYear},
			{Key: "phone", Value: studentData.Phone},
			{Key: "email", Value: studentData.Email},
		}},
	}
	result, err := studentCollection.UpdateOne(ctx, bson.M{"_id": studentID}, updateField)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("no document found with this id")
	}

	return nil
}

func DeleteStudent(studentID string) error {
	// connecting to DB
	DB, ctx, cancel := config.DBConnect()
	defer cancel()
	defer DB.Client().Disconnect(ctx)

	// taking DB collection/table to a variable
	studentCollection := DB.Collection("student")

	result, err := studentCollection.DeleteOne(ctx, bson.M{"_id": studentID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no document found with this id")
	}

	return nil
}
