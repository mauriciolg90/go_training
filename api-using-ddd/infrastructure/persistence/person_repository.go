package persistence

import (
    "database/sql"
    "errors"
    "log"

    _ "github.com/go-sql-driver/mysql"
    "../../domain/entity"
)

// Returns a pointer to the database connection
func SetupDB(driver, user, pass, host, schema string) *sql.DB {
    db, err := sql.Open(driver, user + ":" + pass + "@" + host + "/" + schema)

    if err != nil {
        log.Fatal("Error creating a connection on database!!\n", err)
    }

    // Test connection
    err = db.Ping()
    if err != nil {
        log.Fatal("Error connecting database!!\n", err)
    }

    return db
}

// It's the implementation of PersonRepository
type PersonRepositoryImpl struct {
    DB *sql.DB
}

// Returns all persons stored in database
func (r *PersonRepositoryImpl) GetAll() ([]*entity.Person, error) {
    const query = `SELECT * FROM persons`
    rows, err := r.DB.Query(query)

    if err != nil {
        return nil, err
    }

    // Get all persons iterating the rows
    persons := []*entity.Person{}
    for rows.Next() {
        // Save the entities
        var person entity.Person
        err = rows.Scan(&person.ID, &person.Name)
        if err != nil {
            return nil, err
        }
        persons = append(persons, &person)
    }

    return persons, nil
}

// Returns a person from the database with the id
func (r *PersonRepositoryImpl) Get(id int64) (*entity.Person, error) {
    /*if r.DB[id] == nil {
        return nil, errors.New("Person not found")
    }*/

    //return r.DB[id], nil
    return nil, nil
}

// Inserts a person to database
func (r *PersonRepositoryImpl) Save(person *entity.Person) error {
    if person == nil {
        return errors.New("Nil person!!")
    }

    //person.ID = int64(len(r.DB) + 1)
    //r.DB[person.ID] = person
    return nil
}
