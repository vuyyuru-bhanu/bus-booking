FROM node:18-alpine

WORKDIR /app

# Copy only package.json (ignore lock file)
COPY package.json ./

# Install dependencies and generate package-lock.json inside the container
RUN npm install

# Copy the rest of your Angular project
COPY . .

# Build the Angular app
RUN npm run build

EXPOSE 4200

# Start the dev server
CMD ["npm", "start"]
