FROM node:16.18.0-buster

ARG NODE_ENV=production
ENV NODE_ENV $NODE_ENV

ARG PORT=19006
ENV PORT $PORT
EXPOSE $PORT 19000 19001 19002

ENV NPM_CONFIG_PREFIX=/home/node/.npm-global
ENV PATH /home/node/.npm-global/bin:$PATH
RUN npm i --unsafe-perm --allow-root -g npm@latest expo-cli@latest

WORKDIR /opt/front
ENV PATH /opt/font/.bin:$PATH
COPY ./package.json ./package-lock.json ./
COPY . ./
RUN npm install --legacy-peer-deps

WORKDIR /opt/front/app

# COPY ./front .

# ENTRYPOINT ["npm", "run"]
CMD ["npm", "run", "start"]