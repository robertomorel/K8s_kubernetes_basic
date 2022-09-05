# docker build -t robertomorel/hello-go
# docker run --rm -p 80:80 robertomorel/hello-go

FROM golang:1.15

COPY . .

RUN go build -o server .

CMD ["./server"]