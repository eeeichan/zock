# zock
ZOZOをハックします。

1. `docker-compose build`
2. `docker-compose images` #確認
3. `docker-compose up -d`


## ファイル実行(仮)
```
docker-compose exec app go run main.go
```
or
```
docker-compose exec app /bin/bash # コンテナのシェル起動
go run main.go
```
