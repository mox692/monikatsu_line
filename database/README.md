### 概要

- monikatsu は LINE を通して user からの入力を db に保存します。そのための db 関連のソースコードがこのディレクトリにあります。

```
connection.go
dbとの接続を行うSetupDB関数を定義。os.Getenvで環境変数を取得しており、編集の設定方法は別所に依存してる

repository
dbが返したdataを特定のモデルに変換して返す関数たちをこの辺に置いている。
これをdb package内に作ってしまっているのはイケテナイので、今後改修予定。

init
コンテナ起動時 or cloudSQLのダンプファイルとして使いたいリソースをはじめに設定しておきたいshemaを定義しておく。

```
