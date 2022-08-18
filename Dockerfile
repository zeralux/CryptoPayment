# 這是給CICD執行的
FROM golang:1.17.13-alpine as base

# build
FROM base as build
RUN apk add --update --no-cache make
WORKDIR /build
COPY . /build
RUN make build

# service
FROM base
WORKDIR /app
COPY --from=build /build/bin/*.* .
EXPOSE 8080
CMD [ "sh", "-c", "./main" ]


