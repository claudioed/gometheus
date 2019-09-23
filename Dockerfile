##Builder Image
FROM golang:1.12-stretch as builder
ENV GO111MODULE=on
COPY . /staging
WORKDIR /staging
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/application

#Run Image
FROM scratch
COPY --from=builder /staging/bin/application application
EXPOSE 9090
ENTRYPOINT ["./application"]