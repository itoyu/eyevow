from pymongo import MongoClient
from bson.objectid import ObjectId

userIds = [
  ObjectId()
]

users = [
  {
    "_id": userIds[0]
  }
]

vows = [
  {
    "user": userIds[0],
    "text": "01 I going to make a go of this project!",
    "cheer_count": 10,
    "archived": True
  },
  {
    "id": 2,
    "user": userIds[0],
    "text": "02 I going to make a go of this project!",
    "cheer_count": 20,
    "archived": False
  },
  {
    "id": 3,
    "user": userIds[0],
    "text": "03 I going to make a go of this project!",
    "cheer_count": 30,
    "archived": True
  },
  {
    "id": 4,
    "user": userIds[0],
    "text": "04 I going to make a go of this project!",
    "cheer_count": 0,
    "archived": True
  },
  {
    "id": 5,
    "user": userIds[0],
    "text": "05 I going to make a go of this project!",
    "cheer_count": 0,
    "archived": False
  },
  {
    "id": 6,
    "user": userIds[0],
    "text": "06 I going to make a go of this project!",
    "cheer_count": 60,
    "archived": False
  }
]

db = MongoClient().eyevow
db.users.remove()
db.vows.remove()

db.users.insert_many(users)
db.vows.insert_many(vows)