FROM golang:1.16-alpine as builder

ARG PROJECT_WORKING_DIR
WORKDIR ${PROJECT_WORKING_DIR}

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
RUN apk --no-cache add make

COPY ./src/go.mod ./
RUN if [ -e ./src/go.sum ]; then \
        COPY ./src/go.sum ./; \
    fi

RUN go mod download
COPY ./src ./

RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder ${PROJECT_WORKING_DIR} /app
EXPOSE 1323
CMD ["/app"]