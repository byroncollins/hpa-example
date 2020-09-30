FROM golang:1.15.2-alpine
ENV XDG_CACHE_HOME /tmp/.cache
LABEL maintainer="Byron Collins <byronical@gmail.com>"

COPY main.go /main.go
ENTRYPOINT ["go", "run", "/main.go"]