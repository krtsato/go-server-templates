# ------------------------------------------------------------ #
# Docker for ECR
# ------------------------------------------------------------ #

.PHONY: docker-build
docker-build:
	make go-build
	docker build -f ${DOCKERFILE_PATH} -t ${DOCKER_REPO_URI}:${DOCKER_IMAGE_TAG} --build-arg APP_ENV=${APP_ENV} .

.PHONY: docker-push
docker-push:
	make ecr-login
	docker push ${DOCKER_REPO_URI}:${DOCKER_IMAGE_TAG}

.PHONY: docker-image-rm
docker-image-rm:
	docker image rm -f ${DOCKER_REPO_URI}:${DOCKER_IMAGE_TAG}

.PHONY: ecr-delete-repo
ecr-delete-repo:
	aws ecr delete-repository --repository-name ${APP_NAME} --force