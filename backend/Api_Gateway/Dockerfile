FROM nginx:latest

COPY ./api_gateway.conf /etc/nginx/conf.d/api_gateway.conf
COPY ./certificates  /etc/nginx/certificates
RUN rm /etc/nginx/conf.d/default.conf
