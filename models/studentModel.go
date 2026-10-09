package models

type Student struct {
	ID           int     `db:"id" json:"id"`
	Name         string  `db:"name" json:"name"`
	DepartmentID int     `db:"department_id" json:"department_id"`
	Department   string  `db:"department"    json:"department"`
	GradeAverage float32 `db:"grade_average" json:"grade_average"`
}
