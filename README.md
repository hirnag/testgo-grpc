## はじめに

フレームワークechoで作成したテスト用rest apiです。
最初にDockerをインストールしてください。
下記でdockerを起動し、以下にアクセスすれば、実行結果が返却されます。
- localhost:8080/hello
- localhost:8080/calc/[数値]

```sh
docker-compose up
```


>2回目以降の起動は、以下コマンドでcacheなしの起動をして下さい
```sh
docker-compose build --no-cache
docker-compose up
```



