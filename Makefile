build: 
	@cd cmd/; \
	go build -o ../bin/pptx-quiz
run: build
	@./bin/pptx-quiz