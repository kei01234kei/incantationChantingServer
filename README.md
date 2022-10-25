# incantationChantingServer

岐阜大学3年生の情報工学実験Ⅲで使うコードをまとめます。

## 開発環境の整備

このプロジェクトでは開発環境にDockerを使います。

## クイックスタート

### ローカルで開発環境を整備

```bash
go mod tidy
go run src/main.go
```

```bash
# このコマンドを実行すると{"message":"Hello World"}という値が帰ってきます。
curl http://localhost:8000/test
```

### Dockerで開発環境を整備

```bash
docker build -t incantatio-chanting-server .
docker run -itd -p 8000:8000 incantatio-chanting-server
curl http://localhost:8000/test
```

### サーバにファイルをアップロードする

```bash
# このコマンドでは /usr/local 配下のtestファイルをサーバにアップロードしています。
curl -X POST http://localhost:8000/test-upload-file \
-F "file=@/usr/local/test" \
-H "Content-Type: multipart/form-data"
```

アップロードしたファイルは下のコマンドを実行することでダウンロードすることができます。

```bash
# サーバにアップロードしたtestファイルをダウンロードします。
wget http://localhost:8000/test-get-file/test
```
