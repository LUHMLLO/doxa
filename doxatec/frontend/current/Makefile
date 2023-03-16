docker-build:
	docker build -t doxatec-frontend .

docker-run:
	docker run --name DoxatecFrontend -p 4173:4173 --network=host -it --rm -d doxatec-frontend

run: docker-build docker-run
