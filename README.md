# Web App Template
> Go API with React frontend template

TODO

## Running with Docker
```sh
# run the postgres container
docker-compose --profile postgres up

# build this container docker build
docker build -t webapptemplate .

# run this container
docker run -it --env-file ./.env --network=web-app-template_backend -p 8000:8000 webapptemplate
```