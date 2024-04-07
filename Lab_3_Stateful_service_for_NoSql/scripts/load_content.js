db = connect( 'mongodb://mongo1:27017,mongo2:27018,mongo3:27019/?replicaSet=rs0' );

db.createUser({
    user: "msgpostservice",
    pwd: "msgpostservice",
    roles: [ { role: "readWrite", db: "sndb" } ]
  })