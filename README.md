# aws-describe-prj

## must set env

AWS_ACCESS_KEY_ID=xxxxxxxxxx

AWS_SECRET_ACCESS_KEY=xxxxxxxxxxxxxxxxxxxx

AWS_REGION=ap-northeast-1

## function

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

| type | environment | instance type | [public] hostname | [public] IP | [private] hostname | [private] IP | state |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 開発環境用インスタンス | develop | t2.micro | ec2-99-99-9-999.ap-northeast-1.compute.amazonaws.com | 99.99.9.999 | ip-999-99-99-999.ap-northeast-1.compute.internal | 999.99.99.999 | running |
|  |  | t2.micro |  |  | ip-999-9-9-999.ap-northeast-1.compute.internal | 999.99.9.999 | stopped |



#####  ----------------------------------------------------

## aws-cli

pip install awscli --upgrade --user

　

## EC2

aws ec2 describe-instances

aws ec2 describe-network-interfaces

aws ec2 describe-route-tables

aws ec2 describe-security-groups

aws ec2 describe-vpc-endpoints

　

## DynamoDB

aws dynamodb list-tables

aws dynamodb describe-table --table-name m_movie

　

## ECR

aws ecr describe-repositories

　

## ElastiCache

aws elasticache describe-cache-clusters

　

## ELB

aws elb describe-load-balancers

　

## RDS

aws rds describe-db-clusters

aws rds describe-db-instances

　

## redshift

aws redshift describe-clusters

　

## route53

aws route53 list-hosted-zones

aws route53 list-health-checks

　

## S3

aws s3 ls

## SQS

aws sqs list-queues

　


