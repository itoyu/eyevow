const Koa = require('koa');
const Mongo = require('mongodb').MongoClient;

const app = new Koa();

app.boot = async function() {
  app.mongo = await new Mongo("mongodb://localhost:27017").connect();
  app.db = app.mongo.db('eyevow');
  return app;
};

app.shutdown = async function() {
  app.mongo.close();
}

module.exports = app;