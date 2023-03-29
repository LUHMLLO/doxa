build:
	go build -o ./build/doxapi

docker-build:
	docker build -t ./build/doxapi .

docker-run:
	docker run --name Doxapi -p 3000:3000 --network=host -it --rm -d ./build/doxapi

run: docker-build docker-run