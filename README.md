# Go-URLShortener

This project is used to create shortener URL for given URL using REST API written in Go.


## Local Deployment

Install Docker before runing this project from `https://docs.docker.com/engine/install/ubuntu/`

To run this project run

```bash
  docker-compose up
```

enter `localhost:5000/URLShortner` endpoint in postman to use the REST API.


## Unit testing

To run in continuous unit testing mode:

```
docker-compose run urlshortener test
```

To view coverage report in browser, run:

```
go tool cover -html=cover.out
```
