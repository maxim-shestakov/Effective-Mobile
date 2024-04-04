package postgresql

import (
	"log"
	"fmt"

	st "Effective-Mobile/internal/structures"

	l "Effective-Mobile/internal/dbconn"

	_ "github.com/lib/pq"
)

// AddOwner inserts a new owner into the database.
//
// Parameter: owner of type st.Owner.
// Return type: error.
func AddOwner(owner st.Owner) error {
	_, err := l.Db.Exec("INSERT INTO owners (name, surname, patronymic) VALUES ($1, $2, $3)", owner.Name, owner.Surname, owner.Patronymic)
	if err != nil {
		log.Println(err)
	}
	return err
}


// AddCar inserts a car into the database.
//
// Parameter: car of type st.Car.
// Return type: error.
func AddCar(car st.Car) error {
	_, err := l.Db.Exec("INSERT INTO cars (reqnum, mark, model, year, owner_id) VALUES ($1, $2, $3, $4, $5)", car.Regnum, car.Mark, car.Model, car.Year, car.OwnerID)
	if err != nil {
		log.Println(err)
	}
	return err
}

// GetCars retrieves all cars from the database.
//
// No parameters.
// Returns a slice of st.Car structs and an error.
func GetCars() ([]st.Car, error) {
	rows, err := l.Db.Query("SELECT * FROM cars")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	cars := []st.Car{}
	for rows.Next() {
		var car st.Car
		if err := rows.Scan(&car.ID, &car.Regnum, &car.Mark, &car.Model, &car.Year, &car.OwnerID); err != nil {
			log.Println(err)
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}


// UpdateCar updates the information of a car in the database.
//
// Parameter:
//   car st.Car: the car struct containing the information to be updated.
// Return:
//   error: an error if the update operation fails.
func UpdateCar(car st.Car) error {
	args := []interface{}{}
	counter := 1
	query := "UPDATE cars SET"
	if car.Regnum!= "" { //When regnum is not empty we should update it
		query += fmt.Sprintf(" regnum=$%d,", counter)
		args = append(args, car.Regnum)
		counter++
	}
	if car.Mark != "" {
		query += fmt.Sprintf(" mark=$%d,", counter)
		args = append(args, car.Mark)
		counter++
	}
	if car.Model != "" {
		query += fmt.Sprintf(" model=$%d,", counter)
		args = append(args, car.Model)
		counter++
	}
	if car.Year != 0 {
		query += fmt.Sprintf(" year=$%d,", counter)
		args = append(args, car.Year)
		counter++
	}
	
	if counter == 1 {
		log.Println("no fields to update")
		return nil
	}

	query = query[:len(query)-1]
	query += fmt.Sprintf(" WHERE id=$%d", counter)
	args = append(args, car.ID)
	_, err := l.Db.Exec(query, args...)
	if err != nil {
		log.Println("problem with updating information about film", err)
		return err
	}
	return nil
}


// DeleteCar deletes a car from the database based on the provided id.
//
// Parameter: id string
// Return type: error
func DeleteCar(id string) error {
	_, err := l.Db.Exec("DELETE FROM cars WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}