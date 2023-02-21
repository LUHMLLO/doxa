migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ufw2q8ibsg6zjo7")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "y32t88l5",
    "name": "category",
    "type": "select",
    "required": true,
    "unique": false,
    "options": {
      "maxSelect": 1,
      "values": [
        "SmartFridge",
        "SmartVaccinator"
      ]
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ufw2q8ibsg6zjo7")

  // remove
  collection.schema.removeField("y32t88l5")

  return dao.saveCollection(collection)
})
