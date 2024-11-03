package models

import (
	"context"
	"log"
	"time"
)

func (m *mysqlRepository) AllDogBreeds() ([]*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
			SELECT id, breed, weight_low_lbs, weight_high_lbs,
			cast(((weight_low_lbs + weight_high_lbs) / 2) as unsigned) as average_weight,
			lifespan, coalesce(details, ''),
			coalesce(alternate_names, ''),
			coalesce(geographic_origin, '')
			FROM dog_breeds ORDER BY breed
			`
	var breeds []*DogBreed

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b DogBreed

		err := rows.Scan(
			&b.ID,
			&b.Breed,
			&b.WeightLowLbs,
			&b.WeightHighLbs,
			&b.WeightAverageLbs,
			&b.Lifespan,
			&b.Details,
			&b.AlternateNames,
			&b.GeographicOrigin,
		)
		if err != nil {
			return nil, err
		}

		breeds = append(breeds, &b)
	}

	return breeds, nil
}

func (m *mysqlRepository) GetBreedByName(breed string) (*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
			SELECT id, breed, weight_low_lbs, weight_high_lbs,
			cast(((weight_low_lbs + weight_high_lbs) / 2) as unsigned) as average_weight,
			lifespan, coalesce(details, ''),
			coalesce(alternate_names, ''),
			coalesce(geographic_origin, '')
			FROM dog_breeds WHERE breed = ?
			`
	row := m.DB.QueryRowContext(ctx, query, breed)

	var dogBreed DogBreed
	err := row.Scan(
		&dogBreed.ID,
		&dogBreed.Breed,
		&dogBreed.WeightLowLbs,
		&dogBreed.WeightHighLbs,
		&dogBreed.WeightAverageLbs,
		&dogBreed.Lifespan,
		&dogBreed.Details,
		&dogBreed.AlternateNames,
		&dogBreed.GeographicOrigin,
	)
	if err != nil {
		log.Println("Error getting breed by name: ", err)
		return nil, err
	}

	return &dogBreed, nil
}

func (m *mysqlRepository) GetDofOfMonthByID(id int) (*DogOfMonth, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, image, video FROM dog_of_month WHERE id = ?`

	row := m.DB.QueryRowContext(ctx, query, id)

	var dog DogOfMonth
	err := row.Scan(
		&dog.ID,
		&dog.Image,
		&dog.Video,
	)
	if err != nil {
		log.Println("Error on getting dog of month by ID: ", err)
		return nil, err
	}

	return &dog, nil
}
