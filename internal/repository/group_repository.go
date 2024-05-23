package repository

import (
	"database/sql"
	"fmt"

	"github.com/ShipIM/go-group-manager/internal/domain/entity"
)

type GroupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{db}
}

func (r *GroupRepository) CreateGroup(group entity.Group) (entity.Group, error) {
	var created entity.Group

	query := `
		INSERT INTO _group (name, course, grade) 
		VALUES ($1, $2, $3)
		RETURNING id, name, course, grade`

	row := r.db.QueryRow(query, group.Name, group.Course, group.Grade)

	err := row.Scan(&created.Name, &created.Course, &created.Grade)
	if err != nil {
		return group, err
	}

	return created, nil
}

func (r *GroupRepository) GetGroupByName(name string) (entity.Group, error) {
	var group entity.Group

	query := `
		SELECT name, course, grade FROM _group
	 	WHERE name = $1
		FROM _group`

	row := r.db.QueryRow(query, name)

	err := row.Scan(&group.Name, &group.Course, &group.Grade)
	if err != nil {
		if err == sql.ErrNoRows {
			return group, fmt.Errorf("no group found with the name %s", name)
		}

		return group, err
	}

	return group, nil
}

func (r *GroupRepository) FindAllGroups(course *int, grade *string) ([]entity.Group, error) {
	var groups []entity.Group

	query := `
		SELECT name, course, grade
		FROM _group
		WHERE ($1 IS NULL OR course = $1)
		AND ($2 IS NULL OR grade = $2)`

	rows, err := r.db.Query(query, course, grade)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var group entity.Group
		if err := rows.Scan(&group.Name, &group.Course, &group.Grade); err != nil {
			return nil, err
		}

		groups = append(groups, group)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func (r *GroupRepository) UpdateGroup(group entity.Group) error {
	query := `
		UPDATE _group
		SET course = $1, grade = $2
		WHERE name = $3`

	result, err := r.db.Exec(query, group.Course, group.Grade, group.Name)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("no group found with the name %s", group.Name)
	}

	return nil
}

func (r *GroupRepository) DeleteGroupByName(name string) error {
	query := `
		DELETE FROM _group
		WHERE name = $1`

	result, err := r.db.Exec(query, name)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("no group found with the name %s", name)
	}

	return nil
}
