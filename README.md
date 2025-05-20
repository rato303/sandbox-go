# sandbox-go

ECSデプロイ向けのシンプルなGoアプリケーションです。

## ビルド方法

```sh
docker build -t sandbox-go .
```

## 実行方法

```sh
docker run -p 8080:8080 sandbox-go
```

ブラウザまたはcurlで http://localhost:8080 にアクセスしてください。

## 補足

- `go.mod`がない場合は、以下で作成できます:
  ```sh
  go mod init sandbox-go
  ```
- ECS/ECRへのデプロイはAWS公式ドキュメントを参照してください。