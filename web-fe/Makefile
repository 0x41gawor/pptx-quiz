APP_NAME=pptx-fe
TAG=latest
PORT=58174

run:
	npm run dev -- --host

docker-build:
	docker build -t $(APP_NAME):$(TAG) .

docker-run:
	docker run -d -p $(PORT):80 --name $(APP_NAME) $(APP_NAME):$(TAG)