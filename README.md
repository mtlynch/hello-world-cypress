# hello-world-cypress

## Overview

This branch demonstrates how to run Cypress end-to-end tests in using Google Chrome.

Testing with Chrome requires only a small change to the [`e2e/docker-compose.yml` file](https://github.com/mtlynch/hello-world-cypress/blob/chrome/e2e/docker-compose.yml):

```diff
  cypress:
-    image: "mtlynch/cypress:3.2.0"
+    image: "mtlynch/cypress:3.2.0-chrome69"
+    command: ["--browser", "chrome"]
```

It uses the `mtlynch/cypress:3.2.0-chrome69` Docker image, which has the Chrome browser pre-installed, and it appends `--browser chrome` to the [entrypoint command](https://github.com/mtlynch/docker-cypress/blob/09a44838c13bc9b7fe7badd48f4cbef9654f78c7/Dockerfile#L11) of the image, making the complete entrypoint:

```bash
cypress run --browser chrome
```

**Note**: Cypress does not save video recordings when running under Chrome, but it does save screenshots on failure.

For the basic demo using Electron, Cypress' default browser, see the [`master` branch](https://github.com/mtlynch/hello-world-cypress).

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
