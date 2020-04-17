# hello-world-cypress

## Overview

This branch demonstrates how to run Cypress end-to-end tests using Firefox.

Testing with Firefox requires only a small change to the [`e2e/docker-compose.yml` file](https://github.com/mtlynch/hello-world-cypress/blob/chrome/e2e/docker-compose.yml):

```diff
  cypress:
     image: "cypress/cypress:4.4.0"
+    command: ["--browser", "firefox"]
```

It uses the `cypress/included:4.4.0` Docker image, which has the Firefox browser pre-installed, and it appends `--browser firefox` to the entrypoint command of the image, making the complete entrypoint:

```bash
cypress run --browser firefox
```

For the basic demo using Electron, Cypress' default browser, see the [`master` branch](https://github.com/mtlynch/hello-world-cypress).

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
