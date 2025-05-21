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

## Dockerイメージのタグ付けとプッシュ方法

1. Dockerイメージにタグを付ける  
   ```
   docker tag <ローカルイメージ名>:<タグ> <リポジトリURL>/<リポジトリ名>:<タグ>
   ```
   例:  
   ```
   docker tag sandbox-go:latest <your-account-id>.dkr.ecr.<region>.amazonaws.com/<your-repo>:latest
   ```

2. Dockerイメージをリポジトリにプッシュする  
   ```
   docker push <リポジトリURL>/<リポジトリ名>:<タグ>
   ```
   例:  
   ```
   docker push <your-account-id>.dkr.ecr.<region>.amazonaws.com/<your-repo>:latest
   ```

※ `<your-account-id>`, `<region>`, `<your-repo>` などはご自身の環境に合わせて置き換えてください。

## 補足

- `go.mod`がない場合は、以下で作成できます:
  ```sh
  go mod init sandbox-go
  ```
- ECS/ECRへのデプロイはAWS公式ドキュメントを参照してください。