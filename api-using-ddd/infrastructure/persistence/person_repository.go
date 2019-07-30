package persistence

import (
    "database/sql"
    "errors"
    "log"

    _ "github.com/go-sql-driver/mysql"
    "github.com/mauriciolg90/go_training/api-using-ddd/domain/entity"
)

// Returns a pointer representing the database connection
func GetDB(driver, user, pass, host, schema string) *sql.DB {
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

// Returns all persons stored in repository
func (r *PersonRepositoryImpl) GetAll() ([]*entity.Person, error) {
    // SQL query
    const query = `SELECT * FROM persons`
    rows, err := r.DB.Query(query)

    if err != nil {
        return nil, err
    }

    // Get all persons iterating the rows
    persons := []*entity.Person{}
    for rows.Next() {
        // Save the entity
        var person entity.Person
        err = rows.Scan(&person.ID, &person.Name)
        if err != nil {
            return nil, err
        }
        persons = append(persons, &person)
    }

    return persons, nil
}

// Returns a person with the given id
func (r *PersonRepositoryImpl) Get(id int64) (*entity.Person, error) {
    var person entity.Person

    // SQL query
    const query = `SELECT * FROM persons WHERE id=?`
    err := r.DB.QueryRow(query, id).Scan(&person.ID, &person.Name)

    return &person, err
}

// Adds a new person with the given name
func (r *PersonRepositoryImpl) Add(person *entity.Person) error {
    if person == nil {
        return errors.New("Nil person!!")
    }

    // SQL query (ID is autoincrement!!)
    const query = `INSERT INTO persons (name) VALUES (?)`
    _, err := r.DB.Exec(query, person.Name)

    return err
}
