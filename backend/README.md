# doxavet

This project was created using create-payload-app using the ts-todo template.

## How to Use

`yarn dev` will start up your application and reload on any changes.

### Docker

If you have docker and docker-compose installed, you can run `docker-compose up`

To build the docker image, run `docker build -t my-tag .`

Ensure you are passing all needed environment variables when starting up your container via `--env-file` or setting them with your deployment.

The 3 typical env vars will be `MONGODB_URI`, `PAYLOAD_SECRET`, and `PAYLOAD_CONFIG_PATH`

`docker run --env-file .env -p 3000:3000 my-tag`

```go
type Client struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Mail           string `json:"mail"`
	Phone          int    `json:"phone"`
	Secret         string `json:"secret"`
	ContractNumber int    `json:"contract_number"`
}

type Device struct {
	ID     int    `json:"id"`
	Name   string `json:"Name"`
	Client Client `json:"client"`
	Date   string `json:"date"`
}

type SmartFridge struct {
	Device       Device  `json:"device"`
	TempSuperior float64 `json:"temp_superior"`
	TempMedium   float64 `json:"temp_medium"`
	TempInferior float64 `json:"temp_inferior"`
}

type SmartVaccinator struct {
	Device       Device `json:"device"`
	Vaccine      string `json:"vaccine"`
	Lot          string `json:"lot"`
	Vaccinations string `json:"vaccinations"`
}
```
