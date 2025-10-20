package repository

import (
	"database/sql"
	"hutchison-test/model"
	"log"
)

type DogsRepositoryI interface {
	Create(dog *model.Dog) (*model.Dog, error)
	ListAll() ([]model.Dog, error)
	GetByID(id uint) (*model.Dog, error)
	DeleteByID(id uint) error
	EditByID(id uint, dog *model.Dog) (*model.Dog, error)
}

type DogsRepository struct {
	Db *sql.DB
}

func (r DogsRepository) Create(dog *model.Dog) (*model.Dog, error) {
	res, err := r.Db.Exec("INSERT INTO dogs (breed, variants) VALUES (?, ?)", dog.Breed, dog.Variants)
	if err != nil {
		return nil, err
	}
	if id, err := res.LastInsertId(); err == nil {
		rows, err := r.Db.Query("SELECT * FROM dogs WHERE id = ?", id)
		if err != nil {
			return nil, err
		}

		var createdDog model.Dog
		for rows.Next() {
			var dog model.Dog
			err := rows.Scan(&dog.ID, &dog.Breed, &dog.Variants, &dog.CreatedAt, &dog.UpdatedAt, &dog.DeletedAt)

			if err != nil {
				return nil, err
			}
			createdDog = dog

		}
		return &createdDog, nil
	}

	return nil, err
}
func (r DogsRepository) ListAll() ([]*model.Dog, error) {
	rows, err := r.Db.Query("SELECT * FROM dogs WHERE deleted_at is NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dogsList []*model.Dog
	for rows.Next() {
		var dog model.Dog
		err := rows.Scan(&dog.ID, &dog.Breed, &dog.Variants, &dog.CreatedAt, &dog.UpdatedAt, &dog.DeletedAt)

		if err != nil {
			return nil, err
		}

		dogsList = append(dogsList, &dog)
	}
	return dogsList, nil
}

func (r DogsRepository) GetByID(id uint) (*model.Dog, error) {
	rows, err := r.Db.Query("SELECT * FROM dogs WHERE id = ? AND deleted_at is NULL", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foundDog model.Dog
	for rows.Next() {
		var dog model.Dog
		err := rows.Scan(&dog.ID, &dog.Breed, &dog.Variants, &dog.CreatedAt, &dog.UpdatedAt, &dog.DeletedAt)

		if err != nil {
			return nil, err
		}
		foundDog = dog

	}
	return &foundDog, nil
}
func (r DogsRepository) DeleteByID(id uint) error {
	_, err := r.Db.Exec("UPDATE dogs SET deleted_at= CURRENT_TIMESTAMP WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
func (r DogsRepository) EditByID(id uint, dog *model.Dog) (*model.Dog, error) {
	var variants interface{}
	if dog.Variants == nil || *dog.Variants == `[""]` {
		variants = nil
	} else {
		variants = *dog.Variants
	}
	log.Printf("Variants %v", variants)
	_, err := r.Db.Exec("UPDATE dogs SET breed = ?, variants = ? WHERE id = ?", dog.Breed, variants, id)
	if err != nil {
		return dog, err
	}
	return dog, nil
}
