package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "q8lu8udakya8d8u",
			"created": "2024-08-25 12:25:02.003Z",
			"updated": "2024-08-25 12:25:02.003Z",
			"name": "numbers",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "0rpcayct",
					"name": "numbers",
					"type": "json",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSize": 2000000
					}
				}
			],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("q8lu8udakya8d8u")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
