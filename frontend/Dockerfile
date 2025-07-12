# Use a specific Node.js version or a newer version like 20.x if required
FROM node:20-alpine

# Set working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json (if available) for consistent installs
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Optional: Run npm audit and fix vulnerabilities (without breaking changes)
RUN npm audit fix

# Optional: Use `--force` to fix vulnerabilities even if breaking changes are involved
# RUN npm audit fix --force

# Copy the rest of your application files
COPY . .

# Optional: Run build if required for production
RUN npm run build

# Expose the port your app runs on
EXPOSE 4200

# Start the application
CMD ["npm", "start"]
