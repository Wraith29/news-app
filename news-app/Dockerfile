FROM node:latest AS build

ENV BUN_INSTALL="/usr"

RUN curl -fsSL https://bun.sh/install | bash

WORKDIR /usr/src/app

COPY package.json bun.lockb .

RUN bun install

COPY . .

RUN bun run build

FROM nginx:latest

COPY --from=build /usr/src/app/dist/news-app/browser /etc/nginx/news-app
