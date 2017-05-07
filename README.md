# aws-describe-prj

## must set env

AWS_ACCESS_KEY_ID=xxxxxxxxxx

AWS_SECRET_ACCESS_KEY=xxxxxxxxxxxxxxxxxxxx

AWS_REGION=ap-northeast-1

## function

### get SQS information

$ awsdescribe sqs


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

　


