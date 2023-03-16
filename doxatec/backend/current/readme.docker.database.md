# Crea la db en docker

```ps
docker run --name DoxaDB -e POSTGRES_DB=doxatec -e POSTGRES_USER=doxadmin -e POSTGRES_PASSWORD=d@x@dm1n -p 5432:5432 -d postgres:latest
```
