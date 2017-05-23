# aws-describe-prj

## AWS用環境変数セット済みの前提

AWS_ACCESS_KEY_ID=xxxxxxxxxx

AWS_SECRET_ACCESS_KEY=xxxxxxxxxxxxxxxxxxxx

AWS_REGION=ap-northeast-1

## 機能

##### ・AWS機能別にサブコマンド実行

##### ・テンプレート形式(template配下)による出力内容制御(※現在はMarkDown形式のみ用意)

##### ・一覧表示項目に手動で足したい情報を設定ファイル(config配下)に記載可能

##### ・表示したくない分をフィルタリング可能(設定ファイル(config配下)に記載)

## コマンド

### get SQS information

$ awsdescribe sqs

[output example]

SQS

| usecase | environment | queue_name | url |
| :--- | :--- | :--- | :--- |
| 開発環境用キュー１ | develop | example-develop-queue01 | https://sqs.ap-northeast-1.amazonaws.com/0123456789/example-develop-queue01 |
| 本番環境用キュー１ | production | example-queue01 | https://sqs.ap-northeast-1.amazonaws.com/0123456789/example-queue01 |
| ステージング環境用キュー１ | staging | example-staging-queue01 | https://sqs.ap-northeast-1.amazonaws.com/0123456789/example-staging-queue01 |


### get EC2 information

$ awsdescribe ec2

[output example]

EC2

| usecase | environment | instance name | instance type | [public] hostname | [public] IP | [private] hostname | [private] IP | state |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 開発環境用インスタンス | develop | PublisherService | t2.micro | ec2-9999999.ap-northeast-1.compute.amazonaws.com | 99.999.999.99 | ip-9999999.ap-northeast-1.compute.internal | 99.999.999.99 | running |
|  |  |  | t2.micro |  |  | ip-192-168-1-4.ap-northeast-1.compute.internal | 192.168.1.4 | stopped |


### get RDS information

$ awsdescribe rds

[output example]

RDS

| usecase | environment | instance type | DB name | address | port | engine | version | username | state |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 開発環境用 | develop | db.t2.micro |  | example.xxxxxxxxxx.ap-northeast-1.rds.amazonaws.com | 5432 | postgres | 9.6.2 | dummyuser | available |
| ステージング環境用 | staging | db.t2.micro | testdb | gginstance.xxxxxxxx.ap-northeast-1.rds.amazonaws.com | 3306 | mysql | 5.6.27 | testuser | available |


## 改修予定

##### ・ソート機能を追加

##### ・MarkDown形式以外の出力形式（例：CSV、HTML）に対応

##### ・標準出力以外の出力先に対応

##### ・「all」サブコマンド、ないし、サブコマンドなしの際に全AWSサブコマンドを実行する機能を追加

##### ・デフォルトの出力形式、出力先の決定、及び、起動オプションで選択できるよう変更