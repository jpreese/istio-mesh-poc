FROM golang as builder

COPY ./protos /go/src/github.com/jpreese/istio-mesh-poc/cf-factservice/protos
COPY ./main.go .

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/main

FROM scratch
COPY --from=builder /bin/main .

ENTRYPOINT ["./main"]
