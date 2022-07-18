FROM node:16-buster as build

WORKDIR /app/

COPY package.json .
COPY package-lock.json .

RUN npm install

COPY ./ ./

RUN npm run build

FROM nginx:1.23.0 as server

COPY --from=build /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf