package sqlstore

import (
	"database/sql"
	"fmt"
	"sberapi/internal/model"
)

type DirectionRepository struct {
	db *sql.DB
}

func (r *DirectionRepository) Create(dir *model.Direction) error {
	err := r.db.QueryRow(
		"INSERT INTO directions(name)"+
			"VALUES ($1) RETURNING id",
		dir.Name).Scan(&dir.ID)

	if err != nil {
		return err
	}
	return nil

}

func (r *DirectionRepository) Find(id uint64) (*model.Direction, error) {
	dir := &model.Direction{}
	if err := r.db.QueryRow("SELECT * FROM directions WHERE id = $1", id).Scan(
		&dir.ID,
		&dir.Name,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Direction with that id is not found")
		}
		return nil, err
	}
	return dir, nil
}

func (r *DirectionRepository) All() ([]model.Direction, error) {
	dirs := make([]model.Direction, 0, 10)

	rows, err := r.db.Query("SELECT id, name FROM directions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		dir := model.Direction{}
		err := rows.Scan(&dir.ID, &dir.Name)
		if err != nil {
			return nil, err
		}
		dirs = append(dirs, dir)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dirs, nil
}
