FROM golang as builder

# Move all of the required source files from the host machine to the docker image to be built
COPY ./proto /go/src/github.com/jpreese/istio-mesh-poc/cf-factservice/proto
COPY ./main.go .

# This pulls all of the dependencies needed to compile the application. This is similar to npm install
RUN go get -d -v

# This gets us closer to a more static binary. With CGO disabled, Go will build in such a way that
# it won't depend on anything at all from the system, allowing us to use the scratch image (which is nothing)
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/main

# This begins a multistage build. Multistaging allows us to have smaller images, as we don't need to include
# the source files required for building the image. we just need the binary itself. (Think release and debug builds)

# scratch is a base that includes nothing
FROM scratch

# Copy JUST the main.exe binary from the builder layer above.
COPY --from=builder /bin/main .

# Annotate that 80 should be exposed and run the main.exe binary when the container is run
ENTRYPOINT ["./main"]
EXPOSE 80
