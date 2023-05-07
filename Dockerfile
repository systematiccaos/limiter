FROM alpine:latest

WORKDIR /app
RUN adduser --home /app --disabled-password app
ADD limiter .
RUN chmod +x /app/limiter
USER app

CMD /app/limiter