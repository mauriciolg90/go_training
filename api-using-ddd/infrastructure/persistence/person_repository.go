package persistence

import (
    "errors"
    "../../domain/entity"
)

// It's the implementation of PersonRepository
type PersonRepositoryImpl struct {
    db map[int64]*entity.Person // FIXME, aca tengo que meter un ptr a una base de datos real
}

// Returns a person from the database with the id
func (r *PersonRepositoryImpl) Get(id int64) (*entity.Person, error) {
    if r.db == nil {
        return nil, errors.New("Database error")
    }

    if r.db[id] == nil {
        return nil, errors.New("Person not found")
    }

    return r.db[id], nil
}

// Returns all persons stored in database
func (r *PersonRepositoryImpl) GetAll() ([]*entity.Person, error) {
    if r.db == nil {
        return nil, errors.New("database error")
    }

    persons := []*entity.Person{}
    for _, person := range r.db {
        persons = append(persons, person)
    }

    return persons, nil
}

// Inserts a person to database
func (r *PersonRepositoryImpl) Save(person *entity.Person) error {
    if person == nil {
        return errors.New("Nil person")
    }
    if r.db == nil {
        return errors.New("Database error")
    }

    person.ID = int64(len(r.db) + 1)
    r.db[person.ID] = person
    return nil
}
