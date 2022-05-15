# atlas-domain-navigation-pf

- each observed sector of the galaxy has unique numeric SectorID assigned to it
- each sector will have at least one DNS deployed
- each sector has different number of drones deployed at any given moment
- it’s future, but not that far, so drones will still use JSON REST API

#### Curl Command

- POST /dns/v1/locate-data-bank
````aidl
curl --location --request POST 'http://localhost:3001/dns/v1/locate-data-bank' \
--header 'Content-Type: application/json' \
--data-raw '{
    "x": 123.12,
    "y": 456.56,
    "z": 789.89,
    "vel": 20.0
}'
````
- GET /dns/v1/status
```aidl
curl --location --request GET 'http://localhost:3001/dns/v1/status'
```

#### Libraries

````aidl
	github.com/go-chi/chi v1.5.4
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.11.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.4.0
	github.com/spf13/viper v1.11.0
````

#### docker command
````aidl
docker-compose up --build
````

#### Environment Variables

| Environment Variables | Description | Value |
| :---: | :---: | :---: |
| PORT | the port on which the app runs | 3001 |
| VERSION |api version | v1 |
| SECTOR_ID | The sector id of the location where the app is deployed | 1 |