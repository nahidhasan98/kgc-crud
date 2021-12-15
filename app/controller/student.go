package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nahidhasan98/kgc-crud/app/model"
)

func GetAllStudents(c *gin.Context) {
	students, err := model.GetAllStudents()
	if err != nil {
		c.JSON(http.StatusOK, &model.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &model.Response{
		Status: "success",
		Data:   students,
	})
}
func GetSingleStudent(c *gin.Context) {
	studentID := c.Param("id")
	student, err := model.GetSingleStudentByID(studentID)
	if err != nil {
		c.JSON(http.StatusOK, &model.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &model.Response{
		Status: "success",
		Data:   []model.Student{*student},
	})
}

func checkEmptyField(c *gin.Context, student model.Student) *model.Response {
	if student.Name == "" {
		return &model.Response{
			Status:  "error",
			Message: "no name provided",
			Type:    "name",
		}
	}
	if student.Session == "" {
		return &model.Response{
			Status:  "error",
			Message: "no session provided",
			Type:    "session",
		}
	}
	if student.Reg == "" {
		return &model.Response{
			Status:  "error",
			Message: "no registration number provided",
			Type:    "reg",
		}
	}
	if student.Roll == "" {
		return &model.Response{
			Status:  "error",
			Message: "no class roll provided",
			Type:    "roll",
		}
	}
	if student.PassingYear == "" {
		return &model.Response{
			Status:  "error",
			Message: "no passing year provided",
			Type:    "passingYear",
		}
	}
	return nil
}

func checkForExistence(student model.Student) *model.Response {
	// checking the reg
	s, _ := model.GetSingleStudentByReg(student.Reg)
	if s != nil {
		// (s!=nil) => means got a student with this requested updatable reg num
		// (student.ID == "") => means this existence check called from add student func
		// so this will return err
		//
		// next one, if first part false then execution come to 2nd part of [OR-||] - so this is true [student.ID !=""]
		// and (student.ID != "") => means this existence check called from update student func, thus has an id
		// so updatable student id should be same as the found student id (because their reg num same)
		// but if updatable student id is not same as existing student id,
		// then it is clear that updatable student want to use existing other's reg number
		// so return error
		if (student.ID == "") || (student.ID != s.ID) {
			return &model.Response{
				Status:  "error",
				Message: "registration number already exist",
				Type:    "reg",
			}
		}
	}
	// checking the roll
	s, _ = model.GetSingleStudentByRoll(student.Roll)
	if s != nil {
		if (student.ID == "") || (student.ID != s.ID) {
			return &model.Response{
				Status:  "error",
				Message: "roll number already exist",
				Type:    "roll",
			}
		}
	}
	// checking the email
	s, _ = model.GetSingleStudentByEmail(student.Email)
	if s != nil {
		fmt.Println(s.ID)
		if (student.ID == "") || (student.ID != s.ID) {
			return &model.Response{
				Status:  "error",
				Message: "email already exist",
				Type:    "email",
			}
		}
	}
	return nil
}

func CreateStudent(c *gin.Context) {
	var student model.Student
	err := c.BindJSON(&student)
	if err != nil {
		c.JSON(http.StatusOK, &model.Response{
			Status:  "error",
			Message: "invalid JSON object",
		})
		return
	}

	checkResponse := checkEmptyField(c, student) // checking for form value/json field empty or not
	if checkResponse != nil {
		c.JSON(http.StatusOK, checkResponse)
		return
	}

	checkResponse = checkForExistence(student) // checking if this student already exists or not in DB
	fmt.Println(checkResponse)
	if checkResponse != nil {
		c.JSON(http.StatusOK, checkResponse)
		return
	}

	student.ID = uuid.NewString()
	err = model.CreateStudent(&student)
	if err != nil {
		c.JSON(http.StatusOK, &model.Response{
			Status:  "error",
			Message: "couldn't store data to database. error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &model.Response{
		Status:  "success",
		Message: "student successfully added",
	})
}

func UpdateStudent(c *gin.Context) {
	var student model.Student
	err := c.BindJSON(&student)
	if err != nil {
		c.JSON(http.StatusOK, &model.Response{
			Status:  "error",
			Message: "invalid JSON object",
		})
		return
	}

	checkResponse := checkEmptyField(c, student) // checking for form value/json field empty or not
	if checkResponse != nil {
		c.JSON(http.StatusOK, checkResponse)
		return
	}

	checkResponse = checkForExistence(student) // checking if this student already exists or not in DB
	if checkResponse != nil {
		c.JSON(http.StatusOK, checkResponse)
		return
	}

	studentID := c.Param("id")
	err = model.UpdateStudent(studentID, &student)
	if err != nil {
		c.JSON(http.StatusOK, &model.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &model.Response{
		Status:  "success",
		Message: "data successfully updated",
	})
}

func DeleteStudent(c *gin.Context) {
	studentID := c.Param("id")
	err := model.DeleteStudent(studentID)
	if err != nil {
		c.JSON(http.StatusOK, &model.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &model.Response{
		Status:  "success",
		Message: "data successfully deleted",
	})
}
