# monikatsu_server

### 概要

- LINE サーバーからのリクエストを具体的に処理する package です。main.go から JudgeEvent が呼ばれ、これがこの package のエントリポイントになります。
- LINE から送られたイベントによって後続の処理を切り分けます。

- はじめ routing.go で全ての処理を捌きますが、その後イベント毎に別々の処理が走ります(file 分けた)

- LINE の token を Redis に保存し、セッションとして利用します。この部分は session package に依存しています。セッションは会話のコンテキストを判別するのに利用されます。

- 特定の会話処理によっては、line サーバーからのリクエストを db へ保存する操作も含まれます。

### todo

- session server と gRPC で通信できる事。session 保存、取り出しができる事。
- https://app.diagrams.net/#G1GuzbG3DYwBPoyNBTnjiMk-KnH4dsLaFD に習って、会話 bot を作成する
- LINE サーバーからのリクエスト、session、レスポンス、エラー発生時、それぞれのタイミングで log を生成させる
