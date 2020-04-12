# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# It starts from the node base image.
FROM node

# Add maintainer info.
LABEL maintainer="Ícaro Ribeiro <icaroribeiro@hotmail.com>"

# Set the current working directory inside the container.
WORKDIR /app

# Copy the source from the current directory to the working directory inside the container.
COPY . .

# This command installs a package, and any packages that it depends on. 
RUN npm install

# This command runs the npm build command when we install a new version of node.
RUN npm rebuild

# Command to run the application.
CMD ["node", "cmd/server/index.js"]