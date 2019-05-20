# REST APIs for [Trivy](https://github.com/knqyf263/trivy)

## Usage

### Run API server

```console
docker run --name trivy -d --rm -p 9000:9000 pottava/trivy:0.0.16
```

### Consume APIs

get repositories

```console
curl -si -X GET -H 'Content-Type:application/json' \
  http://localhost:9000/api/v1/images/python:3.4.10-alpine3.9/vulnerabilities
```
