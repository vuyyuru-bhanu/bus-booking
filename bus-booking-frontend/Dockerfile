# Dockerfile for Angular app
FROM node:18 AS build

WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build --prod

FROM nginx:alpine
COPY --from=build /app/dist/bus-booking-app /usr/share/nginx/html
EXPOSE 80

