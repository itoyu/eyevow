from pymongo import MongoClient
from flask import Flask, jsonify
from db import db

def build_user(doc):
    return {
        "id": str(doc["_id"])
    }

def build_vow(doc):
    return dict(
        {k: doc[k] for k in ['text', 'cheer_count', 'archived']},
        id=str(doc['_id']),
        user=build_user(db.users.find_one(doc['user']))
    )

def render_vows(cs):
    return [build_vow(d) for d in cs]

app = Flask(__name__)

@app.route('/vow')
def get_vow(request):
    pass

@app.route('/vows')
def get_vows():
    vows = db.vows.find()
    return jsonify(render_vows(vows))

if __name__ == '__main__':
    app.debug = True
    app.run()