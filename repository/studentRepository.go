package repository

import (
	"context"
	"database/sql"
	"db-api/models"
	"log"
)

var _ StudentRepository = (*studentRepo)(nil)

type StudentRepository interface {
	GetAll() ([]models.Student, error)
	GetByID(ID int) (*models.Student, error)
	Delete(ID int) error
	Create(student *models.Student) error
	Update(id int, student *models.Student) error
}

// repository veriye erişen katmandır. tablodan gelen satırları go structlarına dönüştürür. sql sorguları buradadır
// başka bil db kulanıldığında repositoryi baştan yazmak yeterlidir ör postgresql
type studentRepo struct {
	ctx context.Context
	db  *sql.DB
}

func NewStudentRepository(ctx context.Context, db *sql.DB) StudentRepository {
	return &studentRepo{ctx: ctx, db: db}
}

func (s *studentRepo) GetAll() ([]models.Student, error) {
	query := ` SELECT students.id, students.name, students.department_id, departments.name, students.grade_average from students
	INNER JOIN departments ON students.department_id = departments.id `

	rows, err := s.db.QueryContext(s.ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student

	for rows.Next() {
		var std models.Student

		if err := rows.Scan(&std.ID, &std.Name, &std.DepartmentID, &std.Department, &std.GradeAverage); err != nil {
			return nil, err
		}
		students = append(students, std)
	}
	if rows.Err() != nil {
		return nil, err
	}

	log.Println("fetching all students", students)

	return students, nil
}

func (s *studentRepo) GetByID(id int) (*models.Student, error) {
	query := ` SELECT students.id, students.name, students.department_id, departments.name, students.grade_average FROM students
	INNER JOIN departments ON students.department_id = departments.id WHERE students.id = $1`
	var std models.Student
	err := s.db.QueryRow(query, id).Scan(&std.ID, &std.Name, &std.DepartmentID, &std.Department, &std.GradeAverage)
	if err != nil {
		return nil, err
	}
	return &std, nil
}

func (s *studentRepo) Create(student *models.Student) error {
	query := `INSERT INTO students (name, department_id, grade_average) VALUES ($1, $2, $3) RETURNING id`
	err := s.db.QueryRowContext(s.ctx, query, student.Name, student.DepartmentID, student.GradeAverage).Scan(&student.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *studentRepo) Update(id int, student *models.Student) error {
	query := `UPDATE students SET name = $1, department_id = $2, grade_average = $3 WHERE id = $4`
	_, err := s.db.Exec(query, student.Name, student.DepartmentID, student.GradeAverage, id)
	return err
}

func (s *studentRepo) Delete(id int) error {
	query := `DELETE FROM students WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}
