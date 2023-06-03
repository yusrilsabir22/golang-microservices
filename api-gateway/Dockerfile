FROM golang:1.20.4

WORKDIR /app

COPY . ./
RUN go mod download

FROM quay.io/goswagger/swagger as swagger
COPY --from=auth-svc:latest /opt/app/docs/auth_swagger.json ./docs/auth_swagger.json
COPY --from=logistic-svc:latest /opt/app/docs/logistic_swagger.json ./docs/logistic_swagger.json


# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-gateway

EXPOSE 3000

CMD [ "/api-gateway" ]