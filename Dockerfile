FROM alpine:latest

WORKDIR /app
RUN adduser --home /app --disabled-password app
ADD limiter .
RUN chmod +x /app/limiter
RUN chown -R /app app
USER app

CMD /app/limiter