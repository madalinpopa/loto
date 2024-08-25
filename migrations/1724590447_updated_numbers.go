package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("q8lu8udakya8d8u")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("")

		collection.UpdateRule = types.Pointer("")

		// add
		new_results := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "mcwfhp5a",
			"name": "results",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_results); err != nil {
			return err
		}
		collection.Schema.AddField(new_results)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("q8lu8udakya8d8u")
		if err != nil {
			return err
		}

		collection.CreateRule = nil

		collection.UpdateRule = nil

		// remove
		collection.Schema.RemoveField("mcwfhp5a")

		return dao.SaveCollection(collection)
	})
}
