# Introduction

This repo adds custom Go server upon the existent https://github.com/TykTechnologies/tyk-gateway-docker.

## Steps

1. Create your Go server as here: https://github.com/buraksekili/tyk-test-custom/blob/master/main.go
2. Create container of your server.
    1. Your container will be created based on [this Dockerfile](https://github.com/buraksekili/tyk-test-custom/blob/master/Dockerfile.golang).
    2. Compose file creates your container which referred to as [your-go-server](https://github.com/buraksekili/tyk-test-custom/blob/master/docker-compose.yml#L26) along with Tyk gateway
3. Any requests coming into the host, on the port that Tyk is configured to run on, that go to [this path](https://github.com/buraksekili/tyk-test-custom/blob/b7813fb6dc0b79ebfb31ec8ef56d860f5523fbef/apps/hello-world.json#L23) will be proxied upstream target described [here](https://github.com/buraksekili/tyk-test-custom/blob/b7813fb6dc0b79ebfb31ec8ef56d860f5523fbef/apps/hello-world.json#L24)

## Genarate keys
```bash
curl localhost:8080/tyk/keys -X POST --header "x-tyk-authorization: foo" -d '
{
  "quota_max": 0,
  "rate": 2,
  "per": 5,
  "org_id": "demo-org",
  "access_rights": {
      "demo-api-gateway-id": {
          "api_name": "Tyk Test API",
          "api_id": "demo-api-gateway-id",
          "versions": [
              "Default"
          ],
          "allowed_urls": [],
          "limit": null,
          "allowance_scope": ""
      }
    }
}'
```
And store it as, for instance, `$secretkey`.
## Example

```bash
✗ curl -iSs -H "Authorization: $secretkey2" localhost:8080/test-api/api
HTTP/1.1 200 OK
Access-Control-Allow-Headers: Content-Type,access-control-allow-origin, access-control-allow-headers
Access-Control-Allow-Origin: *
Content-Length: 23
Content-Type: application/json
Date: Mon, 06 Dec 2021 14:29:29 GMT
X-Ratelimit-Limit: 0
X-Ratelimit-Remaining: 0
X-Ratelimit-Reset: 1638799276

{"document":"new doc"}
```