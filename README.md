# Web App Template
> Go API with React frontend template

TODO

## Running with Docker
Create a file in the root of this directory named `.env.docker` and populate it based on `.env.template`.
```sh
# run just the Postgres and PGAdmin container
docker-compose --profile postgres up

# build this container docker build
docker build -t webapptemplate .

# run this container
docker run -it --env-file ./.env.docker --network=web-app-template_backend -p 8000:8000 webapptemplate

# OR
docker compose --profile all up
```