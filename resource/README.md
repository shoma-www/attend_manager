## mysql

MySQLの設定ファイルとマイグレーションファイルを配置している


### migration

- migrateコマンドのinstall
```sh
curl -sSL https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ bionic main" > /etc/apt/sources.list.d/migrate.list
apt-get update
apt-get install -y migrate
```

- fileのテンプレートを生成する
```sh
make migrate-file file=create_cbjpgk
# 2ファイル出力される
# 0000001_{file_name}.up.sql
# 0000001_{file_name}.down.sql
```
  - up: migrationのup時に実行される
  - down: migrationのdown(clean)時に実行される

- migrationの実行
```sh
make migrate-up
make migrate-down
```

- DBのDataのDumpを取得
```sh
make data-dump
```