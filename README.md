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
docker-compose up \
  --build \
  --abort-on-container-exit \
  --exit-code-from cypress
```