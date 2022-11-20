# Dockerfile for angular with build and run stages
# Path: ng-client/Dockerfile
FROM node:16.13.0-alpine3.14 AS builder

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install

COPY . .

RUN npm run build

FROM nginx:1.21.3-alpine AS runner

COPY --from=builder /app/dist/angular-admin /usr/share/nginx/html

COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 3000

CMD ["nginx", "-g", "daemon off;"]