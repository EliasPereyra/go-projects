# RabbitMQ

RabbitMQ is an open source independent implementation of AMQP, written in Erlang.

[AMQP](https://en.wikipedia.org/wiki/Advanced_Message_Queuing_Protocol) (Advanced Message Queuing Protocol) is an open standard protocol for managing messages, queueing, routing, etc.

The AMQP's model defines a series of entities:
  - **Exchangers**
  - **User**
  - **Connection**
  - **Channel**

All the entities for sending and receiving messages are declared inside a channel.

Each entity is named, it must be unique. The client will use the name as the identifier of the service. It must begin with a digit, a letter or an underscore.

## Installation

You can install the [Go RabbitMQL Client library](https://github.com/rabbitmq/amqp091-go), maintained by the RabbitMQ core team.
