package model

import (
	"errors"

	"github.com/nahidhasan98/kgc-crud/config"
)

type Student struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Reg         string `json:"reg"`
	Session     string `json:"session"`
	Roll        string `json:"roll"`
	PassingYear string `json:"passingYear"`
	AvatarURL   string `json:"avatarURL"`
}

func GetAllStudents() ([]Student, error) {
	// taking DB
	DB := config.GetDB()
	defer DB.Close()

	rows, err := DB.Query("SELECT * FROM student ORDER BY name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var temp Student
		err = rows.Scan(&temp.ID, &temp.Name, &temp.Email, &temp.Phone, &temp.Reg, &temp.Session, &temp.Roll, &temp.PassingYear, &temp.AvatarURL)
		if err != nil {
			return nil, err
		}
		students = append(students, temp)
	}

	return students, nil
}

func GetSingleStudentByID(studentID int) (*Student, error) {
	// taking DB
	DB := config.GetDB()
	defer DB.Close()

	rows, err := DB.Query(`SELECT * FROM student WHERE id = ?`, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var student Student
	for rows.Next() {
		err = rows.Scan(&student.ID, &student.Name, &student.Email, &student.Phone, &student.Reg, &student.Session, &student.Roll, &student.PassingYear, &student.AvatarURL)
		if err != nil {
			return nil, err
		}
	}

	if student.ID == 0 {
		return nil, errors.New("student id not found")
	}

	return &student, nil
}

func GetSingleStudentByEmail(email string) (*Student, error) {
	// taking DB
	DB := config.GetDB()
	defer DB.Close()

	rows, err := DB.Query(`SELECT * FROM student WHERE email = ?`, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var student Student
	for rows.Next() {
		err = rows.Scan(&student.ID, &student.Name, &student.Email, &student.Phone, &student.Reg, &student.Session, &student.Roll, &student.PassingYear, &student.AvatarURL)
		if err != nil {
			return nil, err
		}
	}

	if student.ID == 0 {
		return nil, errors.New("email not found")
	}

	return &student, nil
}
func GetSingleStudentByReg(reg string) (*Student, error) {
	// taking DB
	DB := config.GetDB()
	defer DB.Close()

	rows, err := DB.Query(`SELECT * FROM student WHERE reg = ?`, reg)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var student Student
	for rows.Next() {
		err = rows.Scan(&student.ID, &student.Name, &student.Email, &student.Phone, &student.Reg, &student.Session, &student.Roll, &student.PassingYear, &student.AvatarURL)
		if err != nil {
			return nil, err
		}
	}

	if student.ID == 0 {
		return nil, errors.New("registration number not found")
	}

	return &student, nil
}
func GetSingleStudentByRoll(roll string) (*Student, error) {
	// taking DB
	DB := config.GetDB()
	defer DB.Close()

	rows, err := DB.Query(`SELECT * FROM student WHERE roll = ?`, roll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var student Student
	for rows.Next() {
		err = rows.Scan(&student.ID, &student.Name, &student.Email, &student.Phone, &student.Reg, &student.Session, &student.Roll, &student.PassingYear, &student.AvatarURL)
		if err != nil {
			return nil, err
		}
	}

	if student.ID == 0 {
		return nil, errors.New("roll number not found")
	}

	return &student, nil
}

func CreateStudent(studentData *Student) error {
	// taking DB
	DB := config.GetDB()
	defer DB.Close()

	_, err := DB.Exec(`INSERT INTO student (name, email, phone, reg, session, roll, passingYear, avatarURL) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, studentData.Name, studentData.Email, studentData.Phone, studentData.Reg, studentData.Session, studentData.Roll, studentData.PassingYear, studentData.AvatarURL)
	if err != nil {
		return err
	}

	return nil
}

func UpdateStudent(studentID int, studentData *Student) error {
	// taking DB
	DB := config.GetDB()
	defer DB.Close()

	res, err := DB.Exec(`UPDATE student SET name = ?, email = ?, phone = ?, reg = ?, session = ?, roll = ?, passingYear = ? WHERE id = ?`, studentData.Name, studentData.Email, studentData.Phone, studentData.Reg, studentData.Session, studentData.Roll, studentData.PassingYear, studentID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no document found with this id")
	}

	return nil
}

func DeleteStudent(studentID int) error {
	// taking DB
	DB := config.GetDB()
	defer DB.Close()

	res, err := DB.Exec(`DELETE FROM student WHERE id = ?`, studentID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no document found with this id")
	}

	return nil
}
