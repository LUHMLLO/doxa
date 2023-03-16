FROM node:latest
RUN npm install -g pnpm
WORKDIR /web
COPY package.json pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile
COPY . .
RUN pnpm build
EXPOSE 4173
RUN pnpm preview



