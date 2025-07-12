# Stage 1: Build the Angular app
FROM node:20-alpine AS builder

WORKDIR /app

COPY ../frontend/package*.json ./
RUN npm install

COPY ../frontend .
RUN npm run build -- --configuration production

# Stage 2: Serve with nginx
FROM nginx:alpine

COPY ./nginx/nginx.conf /etc/nginx/conf.d/default.conf
# Copy build output to nginx html directory
COPY --from=builder /app/dist/bus-booking-frontend /usr/share/nginx/html

EXPOSE 80