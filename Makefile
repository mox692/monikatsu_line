include secret.mk

ENV_LOCAL_FILE := env.local
ENV_LOCAL = $(shell cat $(ENV_LOCAL_FILE))


####################################
# cloudSQLの作成
####################################
.PHONY: create_cloud_sql
create_cloud_sql:
	gcloud beta sql instances create ${SQL_INSTANCE_ID} \
	       --tier=db-f1-micro \
           --no-assign-ip \
		   --region asia-northeast1 \
		   --network default

.PHONY: setup_cloud_sql
setup_cloud_sql:

	# バケットへダンプファイルをpush
	# gsutil cp ${shell pwd}/database/init/line_user.sql gs://${BA}

	# dbの作成
	gcloud sql databases create ${DATABASE_NAME} --instance=${SQL_INSTANCE_ID}

.PHONY: setup_cloud_sql_user
setup_cloud_sql_user:

	# デフォルト ユーザー アカウントの構成
	gcloud sql users set-password root \
    --host=${DATABASE_HOST} --instance=${SQL_INSTANCE_ID} --prompt-for-password

	# userの作成
	gcloud sql users create test \
    --host=${DATABASE_HOST} --instance=${DATABASE_USER_NAME} --password=${DATABASE_USER_PASSWORD}

	
	# バケットのサービス アカウントに storage.objectAdmin IAM ロールを付与
	# gsutil iam ch serviceAccount:p1058761262135-17xlrw@gcp-sa-cloud-sql.iam.gserviceaccount.com:objectAdmin \
  	# gs://monikatsu-bucket


	# dataのダンプ (権限問題でできない。。)
	gcloud sql import sql ${SQL_INSTANCE_ID} gs://${BUCKET_NAME}/${FILE_NAME} \
                            --database=${DATABASE_NAME}

####################################
# VPCネイティブのクラスタ作成
# ref: https://cloud.google.com/kubernetes-engine/docs/how-to/alias-ips?hl=ja
####################################
.PHONY: create_cluster
create_cluster:
	gcloud container clusters create ${CLUSTER_NAME} \
	--num-nodes=1 \
    --region=asia-northeast1 \
    --enable-ip-alias \
    --create-subnetwork name="" \

	gcloud container clusters get-credentials ${CLUSTER_NAME} --region asia-northeast1 \


####################################
# secretの作成
####################################
# .PHONY: create_secret	
# create_secret:
# 	kubectl create secret generic ${DB_SECRET} \
#   	--from-literal=username=${DB_USER}\
#   	--from-literal=password=${DB_PASSWOERD} \
#   	--from-literal=database=${DB_NAME}
#	kubectl create secret generic ${PRIVATE_IP_SECRET} \
#		--from-literal=db_host=${PRIVATE_IP_ADDRESS}


####################################
# imageのbuildとdockerhubへのpush
####################################
.PHONY: build 
deployment:
	docker build -t motoyukikimura/monikatsu-server .
	docker push motoyukikimura/monikatsu-server

####################################
# GKEへのdeployとserviceのデプロイ
####################################
.PHONY: deployment
deployment:
	gcloud builds submit --tag gcr.io/${ENV_PRO_GCP_PROJECT}/${ENV_PRO_GCP_PROJECT} .
	kubectl apply -f deployment.yaml
	kubectl apply -f service.yaml

####################################
# GKEクラスターの削除
####################################
# ! deploymentの環境変数がcloudSQLと一致しているか確認
.PHONY: delete_cluster
delete_cluster:
	kubectl delete -f service.yaml
	kubectl delete -f deployment.yaml
	gcloud container clusters delete ${CLUSTER_NAME} --region asia-northeast1

### テスト
.PHONY: conn-test
conn-test:
	curl -X ${CLUSTER_LB_IP}/hello
	curl -X ${CLUSTER_LB_IP}/insert_line_user


####################################
# localの起動色々
####################################
.PHONY: up-db
up-db:
	$(ENV_LOCAL) docker-compose -f docker/docker-compose.dev.db.yml -p local up -d

.PHONY: down-db
down-db:
	docker-compose -f docker/docker-compose.dev.db.yml -p local down

.PHONY: up-redis
up-redis:
	docker-compose -f docker/docker-compose.dev.redis.yml -p local up -d

.PHONY: down-redis
down-redis:
	docker-compose -f docker/docker-compose.dev.redis.yml -p local down

.PHONY: up-app
up-app:
	$(ENV_LOCAL) docker-compose -f docker/docker-compose.dev.app.yml -p local up

.PHONY: down-app
down-app:
	docker-compose -f docker/docker-compose.dev.app.yml -p local down

.PHONY: up-session
up-session:
	$(ENV_LOCAL) docker-compose -f docker/docker-compose.dev.session.yml -p local up

.PHONY: down-session
down-session:
	docker-compose -f docker/docker-compose.dev.session.yml -p local down