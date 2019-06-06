FROM golang:alpine
RUN mkdir /app 
ADD . /app/
WORKDIR /app
ENV PORT=3001
RUN go build -o main .
CMD ["./main"]