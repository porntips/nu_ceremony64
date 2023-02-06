#stage 1
FROM node:16-alpine AS builder
WORKDIR /app
COPY . .
RUN npm install
RUN npm run build --prod --build-optimizer=false --base-href
#stage 2
FROM nginx:alpine
# COPY --from=builder /app/dist/ng-docker /usr/share/nginx/html
COPY --from=builder /app/dist/app/ /usr/share/nginx/html
COPY /nginx-custom.conf /etc/nginx/conf.d/default.conf
