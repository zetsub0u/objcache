# Object Storage REST API

## Summary
ObjCache provides a way to store and get objects

## Build Instructions
### Local Binary
The application is written in go and has been tested with **go v1.15** only.
To build just run `make` for a linux binary or `make build-darwin` for an osx binary.
You can check the version of the binary by running ` ./bin/objcache version`
```
objcache f37656ec6334-dirty darwin amd64 go1.15.5 (commit f37656ec6334, branch master)
```
### Docker Image
To build  docker image containing the application run `make docker-image`, this will generate a local image
called `zetsub0u/objcache` and tagged with the git describe at build time.

## Running it
* `make run` or `make run-darwin` to start the server directly
or
* `make docker-run` to run it inside a container

## Endpoints
* `GET /object`: returns an object from the store
* `PUT /object/:obj`: receives an object (int) in the url
* `POST /object`: receives a json object in the body. (ex.`{"obj": 5}`)
* `GET /swagger/index.html`: auto generated swagger docs ui: http://localhost:8080/swagger/index.html

