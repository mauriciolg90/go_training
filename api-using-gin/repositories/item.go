package repositories

import (
    "database/sql"

    "../models"
)

func GetItems(dbPtr *sql.DB) ([]*models.Item, error) {
    const query = `
        SELECT
            id,
            title,
            description
        FROM
            items
    `
    rows, err := dbPtr.Query(query)
    // If error, then return a nil slice
    if err != nil {
        return nil, err
    }
    // Create an slice and iter over the rows
    items := make([]*models.Item, 0)
    for rows.Next() {
        // Save the object fields
        var item models.Item
        err = rows.Scan(&item.Id, &item.Title, &item.Description)
        if err != nil {
            return nil, err
        }
        items = append(items, &item)
    }
    return items, err
}

func GetItem(dbPtr *sql.DB, id int) (*models.Item, error) {
    const query = `
        SELECT
            id,
            title,
            description
        FROM
            items
        WHERE
            id = ?
    `
    // Save the object fields
    var item models.Item
    err := dbPtr.QueryRow(query, id).Scan(&item.Id, &item.Title, &item.Description)
    return &item, err
}

func CreateItem(dbPtr *sql.DB, title, description string) error {
    const query = `
        INSERT INTO items (
            title,
            description
        ) VALUES (
            ?,
            ?
        )
    `
    // Build the query
    _, err := dbPtr.Exec(query, title, description)
    return err
}

func UpdateItem(dbPtr *sql.DB, id int, title, description string) error {
    const query = `
        UPDATE items SET
            title = ?,
            description = ?
        WHERE
            id = ?
    `
    // Build the query
    _, err := dbPtr.Exec(query, title, description, id)
    return err
}

func DeleteItem(dbPtr *sql.DB, id int) error {
    const query = `
        DELETE FROM
            items
        WHERE
            id = ?
    `
    // Build the query
    _, err := dbPtr.Exec(query, id)
    return err
}
