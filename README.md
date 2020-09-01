# httpcat

[![BMC Donate](https://img.shields.io/badge/BMC-Donate-orange)](https://www.buymeacoffee.com/vFa5wfRq6)

a HTTP debug service, prints details of request and returns pre-defined response

## Get

`docker pull guoyk/httpcat`

## Env

* `PORT` port to listen, default to `80`
* `RESPONSE_CODE` response code, default to `200`
* `RESPONSE_TYPE` response content type, default to `text/plain; charset=utf-8`
* `RESPONSE_BODY` response body, default to `OK`

## Credits

Guo Y.K., MIT License
