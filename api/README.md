# Eyevow API

# API仕様書の確認

http://eyevow.work.suichu.net/api/

## 接続チェック

### サインアップ(デバッグ用)

```
curl -X POST http://eyevow.work.suichu.net/api/signup
```

```
{
  "token": "xyyzzz",
  ...
}
```

### 自分情報の確認

```
curl -H "Authorization: Bearer xyyzzz" http://eyevow.work.suichu.net/api/user
```

```
{
  "user": {...}
}
```

### 誓う

``curl -X POST -H "Authorization: Bearer xyyzzz" -d '{"text": "wish!"}' http://eyevow.work.suichu.net/api/vows
```

