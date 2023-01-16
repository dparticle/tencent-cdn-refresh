REPOSITORY = registry.cn-hangzhou.aliyuncs.com/particle/tencent-cdn-refresh
TAG = latest

build:
	docker build -t $(REPOSITORY):$(TAG) .

buildx:
	docker buildx build --platform linux/arm,linux/arm64,linux/amd64 -t $(REPOSITORY):$(TAG) . --push