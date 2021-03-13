

GCP_PROJECT_ID = monikatsu-project
PROTO_PATH = session/session.proto 
ENDPOINT_CONFIG_FILE = api_config.yaml
CLUSTER_NAME = k8


.PHONY: pb
pb:
	protoc \
	    --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		${PROTO_PATH}


.PHONY: pbout
pbout:
	protoc \
		--include_imports \
		--include_source_info \
		--descriptor_set_out out.pb \
		${PROTO_PATH}

# endpointサービスをGCP上に展開
.PHONY: endpoint
endpoint:
	gcloud endpoints services deploy out.pb ${ENDPOINT_CONFIG_FILE}

# dockerimageをGCRにpush
.PHONY: pushimage
pushimage:
	gcloud builds submit --tag gcr.io/${GCP_PROJECT_ID}/manage-session:1.0 .

.PHONY: create_cluster
create_cluster:
	gcloud container clusters create ${CLUSTER_NAME} \
	--num-nodes=2 \
    --region=asia-northeast1 \
    --enable-ip-alias \
    --create-subnetwork name=mysubnet \

# serviceとdeployを定義
.PHONY: deployapp
deployapp:
	kubectl apply -f service.yaml
	kubectl apply -f deployment.yaml