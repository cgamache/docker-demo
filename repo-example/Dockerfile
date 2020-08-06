FROM node:12.4-alpine

WORKDIR /app

COPY package*.json ./
COPY src/ ./src/

RUN npm ci --only=production

EXPOSE 8080

CMD [ "npm", "run", "start" ]