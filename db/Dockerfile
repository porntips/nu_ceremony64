# start with base image
FROM mariadb:10.10.2

# Custom config
COPY ./database/custom.cnf /etc/mysql/conf.d/custom.cnf
# import data into container
# All scripts in docker-entrypoint-initdb.d/ are automatically executed during container startup
COPY ./database/*.sql /docker-entrypoint-initdb.d/

ENV TZ=Asia/Bangkok