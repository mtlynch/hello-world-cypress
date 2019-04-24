# hello-world-cypress

[![CircleCI](https://circleci.com/gh/mtlynch/hello-world-cypress/tree/circle.svg?style=svg)](https://circleci.com/gh/mtlynch/hello-world-cypress/tree/circle)

## Overview

This branch demonstrates how to run [Cypress](https://cypress.io) end-to-end tests in [Circle CI](https://circleci.com). The file [`.circleci/config.yml`](https://github.com/mtlynch/hello-world-cypress/blob/circle/.circleci/config.yml) defines a basic Circle CI configuration.

One excellent feature of Circle CI is that it makes the screenshot and video outputs from your Cypress tests available right in the Circle CI web dashboard:

![circle-ci-artifact](https://user-images.githubusercontent.com/7783288/56669425-20484e00-667f-11e9-8424-7c05b2e507ee.jpg)

For the basic demo without continuous integration, see the [`master` branch](https://github.com/mtlynch/hello-world-cypress).

## Run the test app

The example application is called Sentimentalyzer, a very rudimentary text sentiment analyzer. To run it, enter the following commands:

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
