# ------------------------------------------------------------ #
# Deploy by CloudFormation
# ------------------------------------------------------------ #

.PHONY: cfn-deploy-stack
cfn-deploy-stack:
	aws cloudformation deploy \
	--stack-name ${APP_NAME}-stack \
	--tags Name=${APP_NAME}-stack \
	--capabilities CAPABILITY_NAMED_IAM \
	--template-file ./deployments/cfn_template.yml \
	--parameter-overrides AppEnv=${APP_ENV} AppName=${APP_NAME} DockerImage=${DOCKER_REPO_URI}:${DOCKER_IMAGE_TAG} \
	${CHANGESET_OPTION}

.PHONY: cfn-delete-stack
cfn-delete-stack:
	aws cloudformation delete-stack --stack-name ${APP_NAME}-stack

.PHONY: aws-logs-tail
aws-logs-tail:
	aws logs tail --follow /ecs/${APP_NAME}