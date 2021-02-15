FROM golang:1.15
ARG VERSION
ARG COMMIT
ARG BRANCH

WORKDIR /build
COPY . .
RUN GOOS=linux go build -ldflags "-X github.com/zetsub0u/objcache/cmd.version=$VERSION -X github.com/zetsub0u/objcache/cmd.commit=$COMMIT -X github.com/zetsub0u/objcache/cmd.branch=$BRANCH" -mod vendor -o bin/objcache

FROM ubuntu:18.04
COPY --from=0 /build/bin/objcache /usr/local/bin/objcache
ENTRYPOINT ["/bin/bash"]
