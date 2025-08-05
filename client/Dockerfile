# --- Build stage ---
FROM node:20-alpine AS builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

ENV VITE_API_EXPENSES_BASE="http://localhost:9000/api"

COPY . .
RUN npm run build

# --- Production stage ---
FROM caddy:alpine

# Copy built static files to Caddy's web root
COPY --from=builder /app/dist /usr/share/caddy

# Copy Caddyfile if you want to customize Caddy config (optional)
# COPY Caddyfile /etc/caddy/Caddyfile

# Expose port 3000 for the frontend
EXPOSE 80

# By default, Caddy serves /usr/share/caddy on port 80.
# To serve on port 3000, set CADDY_LISTEN environment variable.
ENV CADDY_LISTEN=:80

# Entrypoint is already set by caddy:alpine
