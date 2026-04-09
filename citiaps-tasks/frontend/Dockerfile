# Build stage
FROM node:20-alpine AS builder

WORKDIR /app

# Copiar package.json y package-lock.json
COPY package*.json ./

# Instalar dependencias (este layer se cachea)
RUN npm ci

# Copiar código fuente
COPY . .

# Build Nuxt
RUN npm run build

# Runtime stage
FROM node:20-alpine

WORKDIR /app

# Copiar package.json para tener express/server disponible
COPY package*.json ./

# Instalar solo dependencias de producción
RUN npm ci --only=production

# Copiar aplicación compilada
COPY --from=builder /app/.output ./.output

EXPOSE 3000

ENV NUXT_HOST=0.0.0.0 NUXT_PORT=3000

CMD ["node", ".output/server/index.mjs"]
