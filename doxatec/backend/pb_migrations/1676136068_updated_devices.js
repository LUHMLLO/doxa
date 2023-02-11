migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ufw2q8ibsg6zjo7")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "qa8pffh0",
    "name": "temp",
    "type": "number",
    "required": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ufw2q8ibsg6zjo7")

  // remove
  collection.schema.removeField("qa8pffh0")

  return dao.saveCollection(collection)
})
