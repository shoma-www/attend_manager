FROM node:14-alpine

WORKDIR /app

RUN apk update && \
    npm install -g npm

EXPOSE 3000

CMD ["npm", "run", "dev"]
