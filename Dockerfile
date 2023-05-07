FROM alpine:latest

WORKDIR /app
RUN adduser --home /app --disabled-password app
USER app
ADD limiter .

CMD /app/limiter