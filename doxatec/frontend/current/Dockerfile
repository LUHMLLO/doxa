FROM node:latest
WORKDIR /api

RUN npm install -g pnpm
COPY .npmrc package.json pnpm-lock.yaml ./

RUN pnpm install --frozen-lockfile --prod
COPY . .

RUN pnpm i
RUN pnpm run build
EXPOSE 4173
CMD ["pnpm", "run", "preview"]
