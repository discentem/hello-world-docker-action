FROM golang:1.18-alpine as GOBUILD
# FROM ubuntu:latest as GOBUILD

# This avoids an error about gopath
WORKDIR /app

COPY $(WORKSPACE) .

# RUN echo $(ls) >> pwd.txt

# RUN cat pwd.txt

RUN CGO_ENABLED=0 go build -o /usr/bin/main

FROM golang:1.18-alpine as RUNPANTRI

COPY --from=GOBUILD /usr/bin/main /usr/bin/main

ENTRYPOINT ["/usr/bin/main"]