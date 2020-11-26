
ENV_PROD_FILE := env.prod
ENV_PROD = $(shell cat $(ENV_PROD_FILE))
ENV_PRO_GCP_PROJECT = monikatsu-project

.PHONY: deployment
deployment:
	gcloud builds submit --tag gcr.io/${ENV_PRO_GCP_PROJECT}/monikatsu-project .
	kubectl apply -f deployment.yaml
	kubectl apply -f service.yaml

