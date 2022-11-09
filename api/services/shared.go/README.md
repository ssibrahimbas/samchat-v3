# Samchat V3 / API / Core

Core package of the chat application.

## Dependencies

- [ ] [mongodb](https://mongodb.org)
- [ ] [jetstream](https://docs.nats.io/jetstream/)

## Run Tests

```bash
go test ./*/* -cover -coverprofile=coverage.out
```

## See Coverage

```bash
go tool cover -html=coverage.out
```
