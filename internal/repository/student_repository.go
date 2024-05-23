package repository

import (
	"database/sql"
	"fmt"

	"github.com/ShipIM/go-group-manager/internal/domain/entity"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{db}
}

func (r *StudentRepository) CreateStudent(student entity.Student) (entity.Student, error) {
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

func (r *StudentRepository) GetStudentById(id int) (entity.Student, error) {
	var student entity.Student

	query := `
		SELECT id, name, surname, patronymic, age, group_name
	 	WHERE id = $1
		FROM student`

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

func (r *StudentRepository) FindAllStudents(name, surname, patronymic, groupName string, age int) ([]entity.Student, error) {
	var students []entity.Student

	query := `
		SELECT id, name, surname, patronymic, age, group_name
		FROM student
		WHERE ($1 IS NULL OR name = $1)
		AND ($2 IS NULL OR surname = $2)
		AND ($3 IS NULL OR patronymic = $3)
		AND ($4 IS NULL OR age = $4)
		AND ($5 IS NULL OR group_name = $5)`

	rows, err := r.db.Query(query, name, surname, patronymic, age, groupName)
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

func (r *StudentRepository) UpdateStudent(student entity.Student) error {
	query := `
		UPDATE student
		SET name = $1, surname = $2, patronymic = $3, age = $4, group_name = $5
		WHERE id = $6`

	result, err := r.db.Exec(query, student.Name, student.Surname, student.Patronymic, student.Age,
		student.GroupName)
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

func (r *StudentRepository) DeleteStudentById(id int) error {
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
