const app = require('./app');

app.boot().then(() => {
  return app.shutdown();
}).catch(console.error);