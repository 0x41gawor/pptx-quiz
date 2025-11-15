build: 
	@cd cmd/; \
	go build -o ../bin/pptx-quiz
run: build
	@./bin/pptx-quiz
docker-build: 
	docker build -t pptx-be:latest .
docker-run: 
	docker run -d --name pptx-be \
		--add-host=host.docker.internal:host-gateway \
		-p 3456:3456 \
		--env NOME \
		--env AGANDSKODE \
		--env PRISTAV \
		--env DB_HOST=host.docker.internal \
		pptx-be:latest