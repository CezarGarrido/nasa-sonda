### Build binary from official Go image

FROM golang:1.14-alpine AS build

COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o /nasa-sonda .

### Put the binary onto Heroku image
FROM heroku/heroku:16
COPY --from=build /nasa-sonda /nasa-sonda
CMD ["/nasa-sonda"]