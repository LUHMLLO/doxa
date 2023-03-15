build:
	go build -o ./doxapi

docker-build:
	docker build -t doxapi .

docker-run:
	docker run -it --rm -p 8081:8081 --name DoxapiContainer doxapi

run: docker-build docker-run