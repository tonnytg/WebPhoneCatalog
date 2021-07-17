# pull official base image
FROM node:latest

# set working directory
WORKDIR /app

# add app
COPY ./webphonecatalog/ .
RUN npm install --silent
RUN npm install react-scripts@3.4.1 -g --silent

# start app
CMD ["npm", "start"]