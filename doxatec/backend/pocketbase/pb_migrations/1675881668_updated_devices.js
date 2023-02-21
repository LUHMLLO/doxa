migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ufw2q8ibsg6zjo7")

  collection.listRule = ""
  collection.viewRule = ""
  collection.createRule = "owner = @request.auth.id"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ufw2q8ibsg6zjo7")

  collection.listRule = null
  collection.viewRule = null
  collection.createRule = null

  return dao.saveCollection(collection)
})
