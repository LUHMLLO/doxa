migrate((db) => {
  const collection = new Collection({
    "id": "ufw2q8ibsg6zjo7",
    "created": "2023-02-08 18:16:41.697Z",
    "updated": "2023-02-08 18:16:41.697Z",
    "name": "devices",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "yiapoamr",
        "name": "name",
        "type": "text",
        "required": true,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      },
      {
        "system": false,
        "id": "c9vr57tv",
        "name": "owner",
        "type": "relation",
        "required": true,
        "unique": false,
        "options": {
          "collectionId": "_pb_users_auth_",
          "cascadeDelete": false,
          "maxSelect": 1,
          "displayFields": [
            "username"
          ]
        }
      }
    ],
    "listRule": null,
    "viewRule": null,
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {}
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("ufw2q8ibsg6zjo7");

  return dao.deleteCollection(collection);
})
