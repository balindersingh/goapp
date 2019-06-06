### Build binary from official Go image
FROM golang:stretch as build
COPY . /app
WORKDIR /app
RUN go build -o /golang-examples .

### Put the binary onto Heroku image
FROM heroku/heroku:16
COPY --from=build /golang-examples /golang-examples
CMD ["/golang-examples"]