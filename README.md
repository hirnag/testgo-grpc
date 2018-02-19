## はじめに

GO言語でgRPC APIを作成しました。
シンプルにクライアントサイドからサーバサイドにgRPCリクエストを投げ、
サーバサイドから返却されたレスポンスを画面に表示するサンプルです。
フレームワークはEchoを使っています。
<br><br>

> 使用方法

最初にDockerをインストールしてください。
docker-compose.ymlが存在するディレクトリに移動し、
docker-compose upすれば、サーバサイド、フロントエンド共に起動します。

```sh
docker-compose up
```
<br>

> 機能

ブラウザから以下にアクセスすると、
gRPCでサーバサイドに通信した結果がブラウザに表示されます。

*Request*
```
localhost:8080/send/[入力値]
```
*Response*
```
Hello, [入力値]!

```

<br><br>
> 参考

変更などを加えた後の再起動は、cacheなし起動を推奨します。
```sh
docker-compose build --no-cache
docker-compose up
```



