# mlsa event 3

a tiny app written in Go (Gin) to demonstrate a use case for queueing systems in data intensive applications.

## key features

This go app:

- uses redis as a backend data store to persist user info
- leverages queues (azure service bus) to handle sending of onboarding emails

## available endpoints and errors

- `/` needs jwt authorization else redirects to `/createUser`

- `/health` endpoint to ensure the server is up and running

- `/createUser` endpoint handles creating user accounts and processing onboarding emails to be sent via the azure service bus

## theory
Right before we jump in, it would be good to give a quick overview of how web services work with a request response model (a synchronous process). let's consider youtube from the point of a creator, you upload a video, wait for it to be processed (thumbnails are created and stored in different sizes, the video is compressed, stored and optimized to be viewed over different speeds), and it gives you a feedback as to whether the process was successful or not. Now this is great if you're working on a small scale, but when you have a ton of users you'll tend to think a bit more about your service architecture with respect to latency. Like in this case, right after the user is done uploading the video there's really no more input required from them, and so, having them wait for the entire process to be completed may not be the best solution business wise.

what can we do? we can decouple the service so that right after uploads all of the video's meta data is sent over a queue for the rest of the processing to take place (the thumbnail generation and all that) and the web server can send an immediate response to the user. this way the user gets a faster experience and if something were to fail we could just requeue the process and, no data is lost and we do not need to trouble the user.[this is introductory, we will not discuss queue overloading]

With web applications, our primary aims are reliability, scalability, and dependability. Message queues are one of the concepts that allow developers to ensure this.

### what is a message?
A message is a string of bytes that is meaningful to the applications that use it. It is a piece of information to be processed (aka a task or job). it could be information to be persisted to disk, meta-data needed for additional processing, or in our case, form data to be used to complete an onboarding flow.

### what is a queue?
A queue is a line of messages waiting to be processed. A producer sends messages unto the queue and a consumer collects that message to be processed. FIFO (first in, first out) order

### what is a message queue?
It is an asynchronous concept employed in inter-process communication (or inter-microservice communication) that allows for tasks or jobs to be sent via a queue as messages waiting to be processed. A producer sends messages unto the queue and a consumer collects that message to be processed.

### types of queueing systems
- point to point: message goes from the queue to only one service
- pub/sub systems: message can be sent to all interested parties

### when to adopt queueing systems:
Like with system design in general, there's no holy grail approach to building systems. but here are situations where adopting a message queue may be helpful:
- indeterministic: db persistance, adopting third party services
- resource hungry backend: crunching tons of geo data
- long running: backing

### benefits of message queues
- decoupling: see usability
- scalability: 
- speed up processing:

## additional reading
- [12 Factor App - Treat backing services as attached resources](https://12factor.net/backing-services)
- [The what, why and how of message queues by Dakshraj Sharma](https://sharmarajdaksh.github.io/blog/the-what-why-and-how-of-message-queues)
- [The Big Little Guide to Message Queues by Sudhir Jonathan](https://sudhir.io/the-big-little-guide-to-message-queues)
- [Industry review of Sudhir's article](https://news.ycombinator.com/item?id=25591492)
- [RabbitMQ Use Cases for Message Queues](https://www.cloudamqp.com/blog/rabbitmq-use-cases-explaining-message-queues-and-when-to-use-them.html)
- [VIDEO - Message Queues by Hussein Nasser](https://www.youtube.com/watch?v=W4_aGb_MOls)
- [VIDEO - Message Queues by Jamil Spain from IBM](https://www.youtube.com/watch?v=xErwDaOc-Gs)
- [Azure Service Bus](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-messaging-overview)