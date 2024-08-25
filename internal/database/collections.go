package database

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

// SaveGeneratedNumbers saves the generated numbers in the database
func SaveGeneratedNumbers(app *pocketbase.PocketBase, n []int) error {
	// Encode the numbers to JSON
	numberJson, err := json.Marshal(n)
	if err != nil {
		return err
	}

	// Get the collection
	collection, err := app.Dao().FindCollectionByNameOrId("numbers")
	if err != nil {
		return err
	}

	// Createa new record
	record := models.NewRecord(collection)
	record.Set("numbers", numberJson)

	if err := app.Dao().SaveRecord(record); err != nil {
		return err
	}
	return nil
}
