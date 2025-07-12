FROM node:20-alpine AS builder

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .
RUN npm run build -- --configuration production

# Stage 2: Serve with http-server
FROM node:20-alpine

RUN npm install -g http-server

WORKDIR /app

COPY --from=builder /app/dist/bus-booking-frontend .

EXPOSE 4200

ENTRYPOINT ["http-server", ".", "-p", "4200"]
