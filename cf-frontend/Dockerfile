FROM node:11.6.0-alpine AS builder
COPY . ./cf-frontend
WORKDIR /cf-frontend
RUN npm install
RUN npx webpack main.js

FROM nginx:1.15.8-alpine
COPY --from=builder /cf-frontend /usr/share/nginx/html
EXPOSE 80
