db = connect( 'mongodb://mongo1:27017,mongo2:27018,mongo3:27019/sndb?replicaSet=rs0' );

db.createUser({
    user: "msgpostservice",
    pwd: "msgpostservice",
    roles: [ { role: "readWrite", db: "sndb" } ]
})

db.getUsers()

db.createCollection("posts")
db.createCollection("messages")
db.getCollectionNames()

db.posts.insertMany(
 {
    "account_id": "11A45F0A",
    "content": "junk1"
 },
 {
    "account_id": "11A45F0B",
    "content": "junk2"
 }
)

db.messages.insertMany(
  {
    "sender_id": "11A45F0D",
    "receiver_id": "11A45F0A",
    "content": "THE FIRST MSG"
  },
  {
    "sender_id": "11A45F0A",
    "receiver_id": "11A45F0D",
    "content": "THE SECOND MSG"
  }
 )

 db.posts.createIndex( { "account_id": 1 } )
 db.messages.createIndex( { "receiver_id": 1 } )

 db.posts.getIndexes()
 db.messages.getIndexes()