FROM node:latest
RUN curl -f https://get.pnpm.io/v6.16.js | node - add --global pnpm
COPY .npmrc package.json pnpm-lock.yaml .pnpmfile.cjs ./
RUN pnpm install --frozen-lockfile --prod
COPY . .
EXPOSE 4173
CMD [ "node", "server.js" ]