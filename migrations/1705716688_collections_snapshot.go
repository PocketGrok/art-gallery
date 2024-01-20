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
		jsonData := `[
			{
				"id": "_pb_users_auth_",
				"created": "2023-04-07 17:35:05.763Z",
				"updated": "2023-04-07 17:35:05.764Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "users_name",
						"name": "name",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "users_avatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif",
								"image/webp"
							],
							"thumbs": null,
							"maxSelect": 1,
							"maxSize": 5242880,
							"protected": false
						}
					}
				],
				"indexes": [],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": true,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"onlyVerified": false,
					"requireEmail": false
				}
			},
			{
				"id": "80lkdeds0yxgt0e",
				"created": "2023-04-07 17:50:01.389Z",
				"updated": "2023-04-09 20:31:38.595Z",
				"name": "people",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "vliihqhn",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "nwe9gjjuew8trhp",
				"created": "2023-04-07 17:50:47.940Z",
				"updated": "2023-04-09 20:27:48.450Z",
				"name": "characters",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "rqnujrfd",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "aswfopgz",
						"name": "owner",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "80lkdeds0yxgt0e",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					}
				],
				"indexes": [],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "fjqye679n86wycc",
				"created": "2023-04-07 17:53:20.433Z",
				"updated": "2023-04-09 20:31:51.595Z",
				"name": "websites",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "6t4owmut",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "wwxtggez",
						"name": "url",
						"type": "url",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"exceptDomains": [],
							"onlyDomains": []
						}
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "9zwa6bdunvmkr70",
				"created": "2023-04-07 17:54:55.580Z",
				"updated": "2023-04-09 20:31:23.821Z",
				"name": "accounts",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "omocniuq",
						"name": "username",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "7qg2rtdc",
						"name": "website",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "fjqye679n86wycc",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "951ljx4t",
						"name": "owner",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "80lkdeds0yxgt0e",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_cGlDgON` + "`" + ` ON ` + "`" + `accounts` + "`" + ` (\n  ` + "`" + `username` + "`" + `,\n  ` + "`" + `website` + "`" + `\n)"
				],
				"listRule": null,
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "3zgcmsuamx2h4qj",
				"created": "2023-04-07 18:01:03.738Z",
				"updated": "2024-01-18 03:02:53.613Z",
				"name": "commissions",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "0vptvath",
						"name": "description",
						"type": "text",
						"required": false,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "cxrxbbsp",
						"name": "artists",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "80lkdeds0yxgt0e",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "bjrvbxzu",
						"name": "commissioners",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "80lkdeds0yxgt0e",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "nquvuzjl",
						"name": "opened",
						"type": "date",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "gyaavrjo",
						"name": "completed",
						"type": "date",
						"required": false,
						"presentable": true,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "5v3pd7ws",
						"name": "ref",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					}
				],
				"indexes": [],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "cffiujgxwn838ps",
				"created": "2023-04-07 18:08:08.006Z",
				"updated": "2023-08-03 01:51:55.960Z",
				"name": "images",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "0eaqev18",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "b8qsu4qb",
						"name": "file",
						"type": "file",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [],
							"thumbs": [
								"100x100",
								"200x200",
								"350x500f",
								"700x1000f"
							],
							"maxSelect": 1,
							"maxSize": 209715200,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "pnvsv78x",
						"name": "completed",
						"type": "date",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "sfafrsqp",
						"name": "commission",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "3zgcmsuamx2h4qj",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": [
								"id",
								"description",
								"completed"
							]
						}
					},
					{
						"system": false,
						"id": "iuslkanj",
						"name": "artists",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "80lkdeds0yxgt0e",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "qx1x18s7",
						"name": "characters",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "nwe9gjjuew8trhp",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "grlvjxzi",
						"name": "nsfw",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "gns0nhj0",
						"name": "priority",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					}
				],
				"indexes": [],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "bl9jbxx7fkpff5s",
				"created": "2023-04-07 18:10:25.391Z",
				"updated": "2023-04-09 20:31:03.719Z",
				"name": "posts",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "ume1s324",
						"name": "account",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "9zwa6bdunvmkr70",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "urzembza",
						"name": "url",
						"type": "url",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"exceptDomains": null,
							"onlyDomains": null
						}
					},
					{
						"system": false,
						"id": "q9szttub",
						"name": "posted",
						"type": "date",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "3jrvc3pw",
						"name": "artworks",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "cffiujgxwn838ps",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": []
						}
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "5dc78m0knf8wxif",
				"created": "2023-05-02 17:49:40.630Z",
				"updated": "2024-01-18 03:02:53.614Z",
				"name": "sets",
				"type": "view",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "wlsdfbt6",
						"name": "commission",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "3zgcmsuamx2h4qj",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": [
								"id",
								"description",
								"completed"
							]
						}
					},
					{
						"system": false,
						"id": "xqitrwey",
						"name": "images",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 2000000
						}
					},
					{
						"system": false,
						"id": "vp63lcai",
						"name": "completed",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 2000000
						}
					},
					{
						"system": false,
						"id": "gjorsxeu",
						"name": "nsfw",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 2000000
						}
					},
					{
						"system": false,
						"id": "dedt0ady",
						"name": "description",
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
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {
					"query": "SELECT\n  (ROW_NUMBER() OVER()) as id,\n  commission,\n  JSON_GROUP_ARRAY(DISTINCT images.id) as images,\n  IIF(commissions.completed != \"\", commissions.completed, MAX(images.completed)) as completed,\n  MIN(images.nsfw) as nsfw,\n  IIF(commissions.description, commissions.description, images.name) as description\nFROM images\nLEFT JOIN commissions on commissions.id = images.commission\nGROUP BY IIF(images.commission != \"\", images.commission, images.id)\nORDER BY completed DESC"
				}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
