FROM golang:1.20.4

WORKDIR /app

COPY . ./
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /logistic-svc

EXPOSE 50052

CMD [ "/logistic-svc" ]