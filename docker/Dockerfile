FROM node:19-alpine

RUN apk add git
RUN git clone https://gitea.arsenm.dev/Arsen6331/lure-web
RUN cd lure-web && \
    npm i && \
    LURE_WEB_API_URL='http://api.lure.arsenm.dev' npm run build

WORKDIR /lure-web
ENTRYPOINT node build