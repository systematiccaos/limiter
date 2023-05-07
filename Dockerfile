FROM alpine:latest

WORKDIR /app
RUN adduser --home /app app
USER app
ADD limiter .

CMD /app/limiter