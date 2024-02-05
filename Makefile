test:
	go test ./... -cover -coverprofile=test_result/coverage.out
	go tool cover -html=test_result/coverage.out -o test_result/cover.html
build:
	go build -o server .
run:
	./server
clean:
	rm -f server
build-local:
	docker build -t rw-fiber . 
build-dev:
	docker buildx build --push --tag inyourtime/ecommerce-be:dev --platform=linux/amd64 .	