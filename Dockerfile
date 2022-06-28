# STAGE #1 - Build backend
FROM golang:1.18-alpine AS builder-backend
WORKDIR /build

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend ./
RUN CGO_ENABLED=0 go build -o /quotes


# STAGE #2 - Build frontend
FROM node:18.4-alpine AS builder-frontend
WORKDIR /build

COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci

COPY frontend ./
RUN npm run build


# STAGE #3 - Build final image
FROM scratch

COPY --from=builder-backend /quotes /
COPY --from=builder-frontend /build/dist /public

CMD ["/quotes"]
