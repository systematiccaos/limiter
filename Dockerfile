FROM alpine:latest

WORKDIR /app
USER app
ADD limiter .

CMD /app/limiter