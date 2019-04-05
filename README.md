# hello-world-cypress

## Run end-to-end tests

```bash
cd e2e
docker-compose up \
  --build \
  --abort-on-container-exit \
  --exit-code-from cypress
```