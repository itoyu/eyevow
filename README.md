# eyevow project

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Run your tests
```
npm run test
```

### Lints and fixes files
```
npm run lint
```

## 開発環境でのTwitterログイン

1. https://eyevow.work.suichu.net/id/ にアクセス
2. https://eyevow.work.suichu.net/signed?token=xxx の ホスト部分をローカルホストに変更してアクセス
3. localStorageの"USER"キーにトークンが保存されます

## リリース

```
yarn release
```

本番環境が更新されます