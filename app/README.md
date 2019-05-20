# How to develop

Resolve dependencies

```console
docker-compose -f app/tools.yml run --rm deps ensure
```

Run the web server

```console
docker-compose -f app/runtime.yml up
```
