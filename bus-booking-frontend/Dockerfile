# Use Node.js official image for building
FROM node:18 AS build

# Set working directory
WORKDIR /app

# Install dependencies
COPY package*.json ./
RUN npm install

# Copy application files
COPY . .

# Build the Angular application
RUN npm run build -- --configuration production

# Use NGINX to serve the application
FROM nginx:alpine

# Copy built files from the previous stage
COPY --from=build /app/dist/bus-booking-app /usr/share/nginx/html

# Expose port 80
EXPOSE 80

# Start NGINX
CMD ["nginx", "-g", "daemon off;"]
