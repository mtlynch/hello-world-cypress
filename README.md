# hello-world-cypress

## Overview

This repo shows a basic example of using Cypress and Docker Compose to create simple end-to-end tests for any web application. This example uses a Go application, but you can reuse the pattern in this repository for any web application that can run in Docker.

For more information, see the blog post, ["Easy End-to-End Testing with Cypress."](https://mtlynch.io/painless-web-app-testing)

## Run the test app

The example application is called Sentimentalyzer, a very rudimentary text sentiment analyzer. To run it, enter the following commands:

```bash
docker build --tag sentimentalyzer .
docker run \
  --interactive \
  --tty \
  --env PORT=8123 \
  --publish 8123:8123 \
  sentimentalyzer
```

The app will be running on [localhost:8123](http://localhost:8123).

## Run end-to-end tests

To execute the end-to-end tests for Sentimentalyzer, enter the followinng commands:

```bash
cd e2e
docker-compose up --exit-code-from cypress
```

When the command completes, you will see test output on the console and a video of the test run will appear in the folder `e2e/cypress/integration/videos`.

## Other branches

This repo contains several branches to demonstrate different Cypress scenarios:

| Scenario | Branch |
|----------|---------|
| [Basic Cypress example](https://github.com/mtlynch/hello-world-cypress) | [`master`](https://github.com/mtlynch/hello-world-cypress) |
| [Using Cypress with Chrome browser](https://github.com/mtlynch/hello-world-cypress/tree/chrome) | [`chrome`](https://github.com/mtlynch/hello-world-cypress/tree/chrome) |
| [Running Cypress from within Circle CI](https://github.com/mtlynch/hello-world-cypress/tree/circle) | [`circle`](https://github.com/mtlynch/hello-world-cypress/tree/circle) |
| [Running Cypress from within Travis CI](https://github.com/mtlynch/hello-world-cypress/tree/travis) | [`travis`](https://github.com/mtlynch/hello-world-cypress/tree/travis) |
| [Running Cypress in interactive mode](https://github.com/mtlynch/hello-world-cypress/tree/interactive) | [`interactive`](https://github.com/mtlynch/hello-world-cypress/tree/interactive) |
