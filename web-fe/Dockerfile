# ─────────────────────────────────────────
# 1) BUILD STAGE
# ─────────────────────────────────────────
FROM node:20-alpine AS build

WORKDIR /app

# Install dependencies based on package.json + package-lock.json
COPY package*.json ./
RUN npm ci

# Copy project files
COPY . .

# Build
RUN npm run build


# ─────────────────────────────────────────
# 2) RUNTIME STAGE (nginx)
# ─────────────────────────────────────────
FROM nginx:alpine

# Remove default nginx site config
RUN rm /etc/nginx/conf.d/default.conf

# Copy our custom SPA config
COPY nginx.conf /etc/nginx/conf.d/vue.conf

# Copy built files from previous stage
COPY --from=build /app/dist /usr/share/nginx/html

# Expose HTTP
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
