FROM postgres:15.2-alpine
COPY db/migrations/20240404140043_cars.sql /docker-entrypoint-initdb.d/