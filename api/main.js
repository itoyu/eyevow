const Koa = require('koa');
const Router = require('koa-router');
const bodyParser = require('koa-bodyparser');
const shortid = require('shortid');
const fs = require('fs').promises;

class UserRepo {
  constructor() {
    this.mem = [];
  }

  async get(id) {
    const f = this.mem.filter(u => u.id == id);
    return f.length != 0 ? f[0] : null;
  }

  async post(name) {
    const id = shortid.generate();
    this.mem.push({
      id,
      name
    });
  }
}

class VowRepo {
  constructor() {
    this.mem = [];
  }

  async ofUser(user_id) {
    for (let v of this.mem) {
      if (v.user_id === user_id) {
        return v;
      }
    }
    return null;
  }

  async post(user_id, text, eyevow) {
    const id = shortid.generate();

    const v = {
      id, user_id,
      text,
      eyevow,
      archived: false, cheer_count: 0
    };

    this.mem.push(v);

    return v;
  }
}

const db = {
  users: new UserRepo(),
  vow: new VowRepo()
};

const app = new Koa();

app.use(bodyParser());
app.use(async (ctx, next) => {
  ctx.json = ctx.request.body;
  await next();
});

const router = new Router({prefix: '/api'});

router.get('/', async ctx => {
  ctx.body = `<!DOCTYPE html>
    <html>
      <head>
        <title>ReDoc</title>
        <!-- needed for adaptive design -->
        <meta charset="utf-8"/>
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">
        <!--
        ReDoc doesn't change outer page styles
        -->
        <style>
          body {
            margin: 0;
            padding: 0;
          }
        </style>
      </head>
      <body>
        <redoc spec-url='/api/api.yml'></redoc>
        <script src="https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js"> </script>
      </body>
    </html>
  `;
});

router.get('/api.yml', async ctx => {
  ctx.body = await fs.readFile('./api.yml');
});

router.post('/signup/twitter', async ctx => {

});

router.post('/signup/line', async ctx => {

});

router.post('/signin/twitter', async ctx => {

});

router.post('/signin/line', async ctx => {

});

router.get('/vows/', async ctx => {
  ctx.body = {vows: [
    {
      text: "エモすぎる楽曲を作って世界一のロックバンドを結成する！"
    },
    {
      text: "たくさん勉強して宇宙の星々を研究する学者になります。たくさん勉強して宇宙の星々を研究する学者になります。たくさん勉強して宇宙の星々を研究する学者になります。"
    }
  ]};
});

router.get('/vow', async ctx => {
  ctx.body = {vow: await db.vow.ofUser('test')};
});

router.put('/vow', async ctx => {
  const {
    text,
    eyevow
  } = ctx.json;

  ctx.body = {
    vow: await db.vow.post('test', text, eyevow)
  };
});

app.use(router.routes());

app.listen(process.env.PORT || 3000);