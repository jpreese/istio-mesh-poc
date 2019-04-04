FROM golang as builder

COPY ./proto /go/src/github.com/jpreese/istio-mesh-poc/cf-factservice/proto
COPY ./main.go .

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/main

FROM scratch
COPY --from=builder /bin/main .

EXPOSE 8000

ENTRYPOINT ["./main"]
