FROM node:18-alpine

WORKDIR /app

# Only copy package.json to avoid stale/corrupt lock
COPY package.json ./

# Clean install without lock file
RUN npm install

# Now copy the rest
COPY . .

# Optional: If build needed
RUN npm run build

EXPOSE 4200

CMD ["npm", "start"]
