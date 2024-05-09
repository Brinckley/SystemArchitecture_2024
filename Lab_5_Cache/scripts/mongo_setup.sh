#!/bin/bash
sleep 10

mongosh --host mongo1:27017 <<EOF
var cfg = {
  "_id": "rs0",
  "version": 1,
  "members": [
    {
      "_id": 0,
      "host": "mongo1:27017",
      "priority": 2
    },
    {
      "_id": 1,
      "host": "mongo2:27017",
      "priority": 0
    },
    {
      "_id": 2,
      "host": "mongo3:27017",
      "priority": 0
    }
  ]
};
rs.initiate(cfg);

db = connect( 'mongodb://mongo1:27017,mongo2:27018,mongo3:27019/sndb?replicaSet=rs0' );

db.createUser({
    user: "userTry",
    pwd: "userTry",
    roles: [ { role: "readWrite", db: "sndb" } ]
});
db.getUsers();

use sndb;

db.messages.insertMany([{account_id: "AAAAAAAAAAAAA", content: "junk1"},
                       {account_id: "AAAAAAAAAAAAA", content: "junk2"}]);

db.messages.insertMany([{sender_id: "AAAAAAAAAAAAA", receiver_id: "AAAAAAAAAAAAA", content: "THE FIRST MSG"},
                       {sender_id: "AAAAAAAAAAAAA", receiver_id: "AAAAAAAAAAAAA", content: "THE SECOND MSG"}]);

db.getCollectionNames();

db.posts.createIndex( { "account_id": 1 } );
db.messages.createIndex( { "receiver_id": 1 } );

db.posts.getIndexes();
db.messages.getIndexes();
EOF