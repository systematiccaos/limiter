FROM alpine:latest

WORKDIR /app
USER app
ADD build/limiter .

CMD /app/limiter