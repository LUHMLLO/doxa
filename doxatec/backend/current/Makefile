build:
	go build -o ./doxapi

docker-build:
	docker build -t doxapi .

docker-run:
	docker run --name Doxapi -p 3000:3000 --network=host -it --rm -d doxapi

run: docker-build docker-run