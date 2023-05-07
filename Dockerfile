FROM alpine:latest

WORKDIR /app
RUN adduser -q --home /app app
USER app
ADD limiter .

CMD /app/limiter