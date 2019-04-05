# This is a similar approach as already highlighted in cf-factserve's Dockerfile
# We build our webapplication using a node image, and then copy the artifact from
# that build to a base Python image which lets us spin up a really simple http server
# with a single command

FROM node:11.6.0-alpine AS builder
COPY . ./cf-frontend
WORKDIR /cf-frontend
RUN npm install
RUN npx webpack main.js

## Now that the website is built, copy the artifacts to the to-be webserver
FROM python:2.7
WORKDIR /cf-frontend

COPY --from=builder /cf-frontend/ .
ENTRYPOINT [ "python" ]
CMD [ "-m", "SimpleHTTPServer", "80" ]
EXPOSE 80
