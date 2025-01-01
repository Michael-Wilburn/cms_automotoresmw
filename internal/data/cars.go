package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Car struct {
	Id         uuid.UUID
	Online     bool
	CarType    string
	Brand      string
	Model      string
	Year       int32
	Kilometers int64
	CarDomain  string
	Price      float64
	InfoPrice  float64
	Currency   string
	ChasisCode string
	MotorCode  string
	Images     []string
	Codia      string
}

type CarModel struct {
	DB *sql.DB
}

func (c CarModel) Insert(car *Car) (uuid.UUID, error) {

	query := `INSERT INTO cars (online, car_type, brand, model, year, kilometers, car_domain, price, info_price, currency, chasis_code, motor_code, images_urls, codia) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id;`

	args := []interface{}{
		car.Online,           // $1
		car.CarType,          // $2
		car.Brand,            // $3
		car.Model,            // $4
		car.Year,             // $5
		car.Kilometers,       // $6
		car.CarDomain,        // $7
		car.Price,            // $8
		car.InfoPrice,        // $9
		car.Currency,         // $10
		car.ChasisCode,       // $11
		car.MotorCode,        // $12
		pq.Array([]string{}), // $13 (convertir slice de strings a array PostgreSQL)
		car.Codia,            // $14
	}

	return car.Id, c.DB.QueryRow(query, args...).Scan(&car.Id)
}

func (c CarModel) Get(id uuid.UUID) (*Car, error) {
	query := `SELECT * FROM cars WHERE id = $1`

	var car Car

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := c.DB.QueryRowContext(ctx, query, id).Scan(
		&car.Id,
		&car.Online,
		&car.CarType,
		&car.Brand,
		&car.Model,
		&car.Year,
		&car.Kilometers,
		&car.CarDomain,
		&car.Price,
		&car.InfoPrice,
		&car.Currency,
		&car.ChasisCode,
		&car.MotorCode,
		pq.Array(&car.Images),
		&car.Codia,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &car, nil
}

func (m CarModel) Update(car *Car) error {
	return nil
}

func (c CarModel) Delete(id uuid.UUID) error {
	return nil
}
