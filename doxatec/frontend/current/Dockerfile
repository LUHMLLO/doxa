FROM node:latest
WORKDIR /frontend

RUN npm install -g pnpm vite
COPY .npmrc package.json pnpm-lock.yaml ./
RUN pnpm install
COPY . .
RUN pnpm run build
EXPOSE 4173
CMD ["pnpm", "run", "preview"]
