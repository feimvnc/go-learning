src https://www.youtube.com/watch?v=zICaTPBkupY&list=PL5dTjWUk_cPYztKD7WxVFluHvpBNM28N9&index=4
Golang And MongoDB REST API


#bson format 
BSON is a binary serialization format used to store documents and make remote procedure calls in MongoDB. 
The BSON specification is located at bsonspec.org .

# run mongo db docker image 
docker pull mongo
docker run --name mongo -d mongo:latest 
# or publish with localhost port 
# MongoDB instance will be accessible on mongodb://localhost:27017
docker run -p 27017:27017 --name mongo -d mongo 
# or persist with local volume 
docker run -p 27017:27017 --name mongo -v data-vol:/data/db -d mongo 
docker run -d --name mongodb -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=admin -p 27017:27017 mongo

# if you want to mount local volume 
#docker run -d --name mongodb -v ~/data:/data/db -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=admin -p 27017:27017 mongo


docker logs mongo
docker exec -it mongo bash
# docker rm mongo    # remove docker container 

# get mongodb shell 
docker exec -ti mongodb -- bash

# after start get a shell into mongo container 
# The mongo shell is removed from MongoDB 6.0. The replacement is mongosh.
root@bb36822b465b:/# mongosh -u "admin" -p admin --authenticationDatabase "admin"
Current Mongosh Log ID:	6379be8facd66eef18afef0e
Connecting to:		mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.6.0
Using MongoDB:		6.0.3
Using Mongosh:		1.6.0

# show database
test>  show dbs
admin   40.00 KiB
config  12.00 KiB
local   40.00 KiB

#You can create a new Database by using “use Database_Name” 
test> use demo;
switched to db demo
demo> show dbs;
# create collections 
demo> db.createCollection("demo")
{ ok: 1 }
# create docs 
demo> db.demo.insertMany([ {name: "mango", origin: "localhost", port: 27017} ])
{
  acknowledged: true,
  insertedIds: { '0': ObjectId("6379bf99adaee12b0256382b") }
}
# search docs 
demo> db.demo.find().pretty()


#
appDB> db.users.drop()
true
appDB> db.users.find()

# create check 
appDB> db.users.find()
[ { _id: '2', name: 'yyy', gender: 'female', age: 41 } ]

# update check 
appDB> db.users.find()
[ { _id: '2', name: 'zzz', gender: 'female', age: 41 } ]

# delete check 
appDB> db.users.find()