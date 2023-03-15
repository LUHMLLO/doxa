build:
	go build -o ./doxapi

docker-build:
	docker build -t doxapi .

docker-run:
	docker run -it --rm -p 3000:3000 --name DoxapiContainer -d doxapi

run: docker-build docker-run