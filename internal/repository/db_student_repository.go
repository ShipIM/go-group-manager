package repository

import (
	"database/sql"
	"fmt"

	"github.com/ShipIM/go-group-manager/internal/domain/entity"
)

type DbStudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *DbStudentRepository {
	return &DbStudentRepository{db}
}

func (r *DbStudentRepository) CreateStudent(student entity.Student) (entity.Student, error) {
	var created entity.Student

	query := `
		INSERT INTO student (name, surname, patronymic, age, group_name) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, surname, patronymic, age, group_name`

	row := r.db.QueryRow(query, student.Name, student.Surname, student.Patronymic,
		student.Age, student.GroupName)

	err := row.Scan(&created.Id, &created.Name, &created.Surname, &created.Patronymic,
		&created.Age, &created.GroupName)
	if err != nil {
		return student, err
	}

	return created, nil
}

func (r *DbStudentRepository) GetStudentById(id int) (entity.Student, error) {
	var student entity.Student

	query := `
		SELECT id, name, surname, patronymic, age, group_name
		FROM student
	 	WHERE id = $1`

	row := r.db.QueryRow(query, id)

	err := row.Scan(&student.Id, &student.Name, &student.Surname, &student.Patronymic,
		&student.Age, &student.GroupName)
	if err != nil {
		if err == sql.ErrNoRows {
			return student, fmt.Errorf("no student found with the id %d", id)
		}

		return student, err
	}

	return student, nil
}

func (r *DbStudentRepository) FindAllStudents() ([]entity.Student, error) {
	var students []entity.Student

	query := `
		SELECT id, name, surname, patronymic, age, group_name
		FROM student`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var student entity.Student
		if err := rows.Scan(&student.Id, &student.Name, &student.Surname, &student.Patronymic,
			&student.Age, &student.GroupName); err != nil {
			return nil, err
		}

		students = append(students, student)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (r *DbStudentRepository) UpdateStudent(student entity.Student) error {
	query := `
		UPDATE student
		SET name = $1, surname = $2, patronymic = $3, age = $4, group_name = $5
		WHERE id = $6`

	result, err := r.db.Exec(query, student.Name, student.Surname, student.Patronymic, student.Age,
		student.GroupName, student.Id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("no student found with the id %d", student.Id)
	}

	return nil
}

func (r *DbStudentRepository) DeleteStudentById(id int) error {
	query := `
		DELETE FROM student
		WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("no student found with the id %d", id)
	}

	return nil
}
