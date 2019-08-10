
# Pub/Sub Minimalbeispiel mit Go

[:uk: English README](./README.md)

Dieses Repository enthält ein Minimalbeispiel für die Implementierung des [Publish-/Subscribe-Musters][pubsub] mit Go und RabbitMQ.

Diese Code-Beispiele ergänzen meinen kürzlichen Artikel über das Publish-/Subscribe-Muster im [T3N-Magazin](https://t3n.de/magazin/).

Dieses Repository besteht aus zwei Software-Komponenten: Einem _Publisher_ und einem _Subscriber_.

1. Der *Publisher* veröffentlicht Ereignisse (hier in Form eines JSON-kodierten Strings) an einen Event Broker (in diesem Fall [RabbitMQ][rmq]). In diesem Beispiel veröffentlicht der Publisher alle fünf Sekunden eine `UserCreatedMessage` (es soll ja schließlich nur eine Demo des Konzepts sein).
1. Der *Subscriber* abonniert die `UserCreatedMessage`-Nachrichten, die über den Message Brokker veröffentlicht werden und verarbeitet sie (indem sie auf der Befehlszeile ausgegeben werden -- nochmal: das hier ist nur eine Demo).

## Los geht's

Um dieses Beispiel zu starten wird [Docker und `docker-compose`](https://docs.docker.com/install/) benötigt. Alle Komponenten werden wie folgt gestartet:

```
$ docker-compose up
```

## Mehr zu tun

1. Weitere Subscriber können wie folgt gestartet werden:

        $ docker-compose scale subscriber=3

    In diesem Fall sollte man sehen können, dass die veröffentlichten Nachrichten reihum an jeweils einen der drei Subscriber verteilt werden.
    
1. Deaktivieren des Subscribers:

        $ docker-compoes stop subscriber

    In diesem Fall wird der Publisher immer noch weiter neue Events produzieren. Diese Nachrichten werden in der Message Queue des Subscribers gespeichert und verarbeitet, wenn der Subscriber wieder neu gestartet wird.

[pubsub]: https://www.enterpriseintegrationpatterns.com/patterns/messaging/PublishSubscribeChannel.html
[rmq]: https://www.rabbitmq.com/