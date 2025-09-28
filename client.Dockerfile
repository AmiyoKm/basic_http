FROM node:24-alpine AS builder

WORKDIR /app
COPY client/package*.json ./
RUN npm i

COPY client/ ./
RUN npm run build

FROM nginx:1.25-alpine
COPY --from=builder /app/dist /usr/share/nginx/html
COPY client/nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
