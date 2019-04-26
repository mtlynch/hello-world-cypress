# hello-world-cypress

## Run the test app

```bash
docker build --tag sentimentalyzer .
docker run \
  --interactive \
  --tty \
  --env PORT=8999 \
  --publish 8999:8999 \
  sentimentalyzer
```

The app will be running on [localhost:8999](http://localhost:8999).

## Run end-to-end tests

```bash
cd e2e
docker-compose up --exit-code-from cypress
```

## Other branches

This repo contains several branches to demonstrate different Cypress scenarios:

| Scenario | Branch |
|----------|---------|
| [Basic Cypress example](https://github.com/mtlynch/hello-world-cypress) | [`master`](https://github.com/mtlynch/hello-world-cypress) |
| [Using Cypress with Chrome browser](https://github.com/mtlynch/hello-world-cypress/tree/chrome) | [`chrome`](https://github.com/mtlynch/hello-world-cypress/tree/chrome) |
| [Running Cypress from within Circle CI](https://github.com/mtlynch/hello-world-cypress/tree/circle) | [`circle`](https://github.com/mtlynch/hello-world-cypress/tree/circle) |
| [Running Cypress from within Travis CI](https://github.com/mtlynch/hello-world-cypress/tree/travis) | [`travis`](https://github.com/mtlynch/hello-world-cypress/tree/travis) |
