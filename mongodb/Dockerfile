# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# It starts from the mongo base image.
FROM mongo

# Add Maintainer Info.
LABEL maintainer="Ícaro Ribeiro <icaroribeiro@hotmail.com>"

# Add the content of the scripts/ directory to the image.
# All scripts in docker-entrypoint-initdb.d are automatically executed during container startup.
COPY ./scripts/ /docker-entrypoint-initdb.d/