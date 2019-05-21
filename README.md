# REST APIs for [Trivy](https://github.com/knqyf263/trivy)

[![CircleCI](https://circleci.com/gh/pottava/trivy-restapi.svg?style=svg)](https://circleci.com/gh/pottava/trivy-restapi)

[![pottava/trivy](http://dockeri.co/image/pottava/trivy)](https://hub.docker.com/r/pottava/trivy/)

Supported tags and respective `Dockerfile` links:  
・latest ([versions/0.x/Dockerfile](https://github.com/pottava/trivy-restapi/blob/master/versions/0.x/Dockerfile))  
・0.0.16 ([versions/0.x/Dockerfile](https://github.com/pottava/trivy-restapi/blob/master/versions/0.x/Dockerfile))  

## Usage

### Run the API server

```bash
$ docker run --name trivy -d --rm -p 9000:9000 \
    -v "${HOME}/Library/Caches":/root/.cache/ \
    pottava/trivy:0.0.16
```

### Consume APIs

get repositories ([API spec](https://raw.githubusercontent.com/pottava/trivy-restapi/master/spec.yaml))

```bash
$ curl -s -X GET -H 'Content-Type:application/json' \
  "http://localhost:9000/api/v1/images/python:3.4.10-alpine3.9/vulnerabilities"
{
  "Count": 1,
  "Vulnerabilities": [
    {
      "Description": "ChaCha20-Poly1305 is ...",
      "FixedVersion": "1.1.1b-r1",
      "InstalledVersion": "1.1.1a-r1",
      "PkgName": "openssl",
      "References": [
        "https://www.openssl.org/news/secadv/20190306.txt",
        "..."
      ],
      "Severity": "MEDIUM",
      "Title": "openssl: ChaCha20-Poly1305 with long nonces",
      "VulnerabilityID": "CVE-2019-1543"
    }
  ]
}
```
