FROM node:11.6.0-alpine AS builder
COPY . ./cf-frontend
WORKDIR /cf-frontend
RUN npm install
RUN npx webpack main.js

FROM python:2.7
WORKDIR /cf-frontend

COPY --from=builder /cf-frontend/ .
ENTRYPOINT [ "python" ]
CMD [ "-m", "SimpleHTTPServer", "80" ]
EXPOSE 80
