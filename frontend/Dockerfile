# frontend/Dockerfile
FROM node:18-alpine AS build

WORKDIR /app

# Copy package files and install dependencies
COPY package*.json ./
RUN npm install

# Copy source code and build
COPY . .
RUN npm run build

# Serve with Nginx
FROM nginx:alpine

# Copy built frontend from previous stage
COPY --from=build /app/dist /usr/share/nginx/html

# Copy custom Nginx config if needed
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
