from pymongo import MongoClient

users = [
  {
    "_id": "e_taro" 
  },
  {
    "_id": "y_taro"
  },
  {
    "_id": "v_taro"
  },
  {
    "_id": "o_taro"
  },
  {
    "_id": "w_taro"
  }
]

vows = [
  {
    "user": "e_taro",
    "text": "01 I going to make a go of this project!",
    "cheer_count": 10,
    "archived": True
  },
  {
    "id": 2,
    "user": "y_taro",
    "text": "02 I going to make a go of this project!",
    "cheer_count": 20,
    "archived": False
  },
  {
    "id": 3,
    "user": "e_taro",
    "text": "03 I going to make a go of this project!",
    "cheer_count": 30,
    "archived": True
  },
  {
    "id": 4,
    "user": "v_taro",
    "text": "04 I going to make a go of this project!",
    "cheer_count": 0,
    "archived": True
  },
  {
    "id": 5,
    "user": "o_taro",
    "text": "05 I going to make a go of this project!",
    "cheer_count": 0,
    "archived": False
  },
  {
    "id": 6,
    "user": "w_taro",
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