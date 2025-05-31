FROM alpine:latest

RUN mkdir /backend

EXPOSE 8282

COPY backend /backend

WORKDIR /backend

CMD [ "./backend" ]
