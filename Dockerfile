ARG OS=linux

FROM golang:alpine AS build
WORKDIR /grpc
COPY . .
RUN go mod download &&  CGO_ENABLED=0 GOOS=${OS} go build -a -installsuffix cgo -o server ./server/ && \
    CGO_ENABLED=0 GOOS=${OS} go build -a -installsuffix cgo -o client ./client/

FROM golang:alpine
COPY --from=build /grpc/server/server /grpc/server
COPY --from=build /grpc/client/client /grpc/client
CMD ["/grpc/server"]