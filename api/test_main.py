from unittest import main, TestCase
from bson.objectid import ObjectId
from main import app
from db import db

def cleanup():
  print("cleanup")
  db.users.delete_many({})
  db.vows.delete_many({})

class BaseTestCase(TestCase):
  def setUp(self):
    cleanup()
    self.client = app.test_client()

class MainTestCase(BaseTestCase):
  def test_get_vow(self):
    pass
  def test_get_vows(self):
    u = ObjectId()

    db.users.insert_many([
      {
        "_id": u
      }
    ])
    db.vows.insert_many([{
      "text": "test",
      "user": u,
      "cheer_count": 0,
      "archived": False
    }])

    rs = self.client.get('/vows')
    print(rs.json)

if __name__ == '__main__':
    main()