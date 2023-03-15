build:
	go build -o ./doxapi

docker-build:
	docker build -t doxapi .

docker-run:
	docker run -it --rm -p 30999:30999 --name DoxapiContainer doxapi

run: docker-build docker-run