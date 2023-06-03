FROM golang:1.20.4

WORKDIR /app

COPY . ./
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-svc

EXPOSE 50051

CMD [ "/auth-svc" ]