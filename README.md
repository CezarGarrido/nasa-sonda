# nasa-sonda

## Build Heroku

- heroku login
- heroku create
- git push heroku main

## Docker

- sudo docker build -t sonda .
- sudo docker image ls
- sudo docker container run -p 8889:8889 sonda
- `docker-compose build` or `docker-compose up --build`

## Tests

-- go test -v ./... -cover