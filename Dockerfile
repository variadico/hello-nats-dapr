FROM scratch
COPY ./hello-nats-dapr /hello-nats-dapr
CMD ["/hello-nats-dapr"]
