# Minimal Pub/Sub example with Go

[:de: German README](./README.de.md)

This repository contains a minimal example on how to implement the [Publish-/Subscribe Pattern][pubsub] with Go and RabbitMQ.

These code examples complement my recent article on the Publish-/Subscribe Pattern in the [T3N Magazine](https://t3n.de/magazin/).

This repo consists of two pieceses of software: a *Publisher* and a *Subscriber*.

1. The *Publisher* publishes events (here in the form of a JSON-encoded string) to an event broker (in this case, [RabbitMQ][rmq]). In this example, the publisher simply pushes a `UserCreatedMessage` every five seconds (this is a demo of the concept, after all).
1. The *Subscriber* subscribes to the `UserCreatedMessage` published at the message broker and processes these messages (by printing them to the command line -- again, this is a demo).

## Getting started

Running the example requires [Docker and `docker-compose`](https://docs.docker.com/install/). Start everything with:

```
$ docker-compose up
```

## Things to do

1. Start more subscribers:

        $ docker-compose scale subscriber=3

    You should now see that events are distributed in a round-robin manner across all three subscribers.

1. Shut off the subscriber:

        $ docker-compoes stop subscriber

    In this case, the publisher will still continue to produce new events. These events will be stored in the subscriber's message queue and be processed when you start the subscriber again.

[pubsub]: https://www.enterpriseintegrationpatterns.com/patterns/messaging/PublishSubscribeChannel.html
[rmq]: https://www.rabbitmq.com/
