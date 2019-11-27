from pymongo import MongoClient
from aiohttp import web
routes = web.RouteTableDef()

db = MongoClient().eyevow

def user_schema(doc):
    return {
        "id": str(doc["_id"])
    }

def vow_schema(doc):
    return dict(
        {k: doc[k] for k in ['text', 'cheer_count', 'archived']},
        id=str(doc['_id']),
        user=user_schema(db.users.find_one(doc['user']))
    )

def vows_schema(cursor):
   return [vow_schema(doc) for doc in cursor]

@routes.get('/')
async def hello(request):
    return web.json_response({"text": "Hello, world"})

@routes.get('/vows/')
async def vows(request):
    vows = db.vows.find()
    return web.json_response({"vows": vows_schema(vows)})

app = web.Application()
app.add_routes(routes)

if __name__ == '__main__':
    web.run_app(app)