package models

import "database/sql"

type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
}

type mysqlRepository struct {
	DB *sql.DB
}

func NewMysqlRepository(conn *sql.DB) Repository {
	return &mysqlRepository{
		DB: conn,
	}
}

type testRepository struct {
	DB *sql.DB
}

func NewTestRepository(conn *sql.DB) Repository {
	return &testRepository{
		DB: nil,
	}
}
