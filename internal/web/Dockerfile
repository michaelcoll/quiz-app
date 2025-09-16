FROM node:22-alpine AS dependency-base

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

WORKDIR /app

COPY package.json .
COPY pnpm-lock.yaml .

RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm i --frozen-lockfile --shamefully-hoist

FROM dependency-base AS production-base

COPY . .
RUN pnpm run build

FROM node:22-alpine AS production

COPY --from=production-base /app/.output /app/.output

# start the app
CMD [ "node", "/app/.output/server/index.mjs" ]