FROM node:20.11.1-alpine3.19

WORKDIR /usr/src/admin-portal

COPY ./frontend/admin-portal/package.json .

RUN npm install

COPY ./frontend/admin-portal .

RUN npm run build
RUN npm install -g serve

CMD ["serve","-p","3000","build/"]
