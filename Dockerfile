FROM alpine:latest

WORKDIR /app
RUN useradd app
RUN usermod -d /app
USER app
ADD limiter .

CMD /app/limiter