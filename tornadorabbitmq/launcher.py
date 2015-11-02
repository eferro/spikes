#!/usr/bin/env python

import os
import logging
import tornadoredis
import pika
import tornado.ioloop
import tornado.options
from tornado import web


c = tornadoredis.Client()
c.connect()


class AsyncAmqpConsumer(object):
    EXCHANGE_TYPE = 'topic'

    def __init__(self, amqp_url, exchange, routing_key, queue, processor):
        self._url = amqp_url
        self._exchange = exchange
        self._routing_key = routing_key
        self._queue = queue
        self._processor = processor

        self._connection = None
        self._channel = None
        self._closing = False
        self._consumer_tag = None

    def connect(self):
        logging.info('Connecting to %s', self._url)
        return pika.adapters.TornadoConnection(pika.URLParameters(self._url),
                                          self.on_connection_open)

    def close_connection(self):
        logging.info('Closing connection')
        self._connection.close()

    def add_on_connection_close_callback(self):
        logging.info('Adding connection close callback')
        self._connection.add_on_close_callback(self.on_connection_closed)

    def on_connection_closed(self, connection, reply_code, reply_text):
        self._channel = None
        if self._closing:
            self._connection.ioloop.stop()
        else:
            logging.warning('Connection closed, reopening in 5 seconds: (%s) %s',
                           reply_code, reply_text)
            self._connection.add_timeout(5, self.reconnect)

    def on_connection_open(self, unused_connection):
        logging.info('Connection opened')
        self.add_on_connection_close_callback()
        self.open_channel()

    def reconnect(self):
        if not self._closing:
            self._connection = self.connect()

    def add_on_channel_close_callback(self):
        logging.info('Adding channel close callback')
        self._channel.add_on_close_callback(self.on_channel_closed)

    def on_channel_closed(self, channel, reply_code, reply_text):
        logging.warning('Channel %i was closed: (%s) %s',
                       channel, reply_code, reply_text)
        self._connection.close()

    def on_channel_open(self, channel):
        logging.info('Channel opened')
        self._channel = channel
        self.add_on_channel_close_callback()
        self.setup_exchange(self._exchange)

    def setup_exchange(self, exchange_name):
        logging.info('Declaring exchange %s', exchange_name)
        self._channel.exchange_declare(self.on_exchange_declareok,
                                       exchange_name,
                                       self.EXCHANGE_TYPE,
                                       durable=True)

    def on_exchange_declareok(self, unused_frame):
        logging.info('Exchange declared')
        self.setup_queue(self._queue)

    def setup_queue(self, queue_name):
        logging.info('Declaring queue %s', queue_name)
        self._channel.queue_declare(self.on_queue_declareok, queue_name)

    def on_queue_declareok(self, method_frame):
        logging.info('Binding %s to %s with %s',
                    self._exchange, self._queue, self._routing_key)
        self._channel.queue_bind(self.on_bindok, self._queue,
                                 self._exchange, self._routing_key)

    def add_on_cancel_callback(self):
        logging.info('Adding consumer cancellation callback')
        self._channel.add_on_cancel_callback(self.on_consumer_cancelled)

    def on_consumer_cancelled(self, method_frame):
        logging.info('Consumer was cancelled remotely, shutting down: %r',
                    method_frame)
        if self._channel:
            self._channel.close()

    def acknowledge_message(self, delivery_tag):
        self._channel.basic_ack(delivery_tag)

    def on_message(self, unused_channel, basic_deliver, properties, body):
        logging.info("RabbitMQ message received: %s", body)
        self._processor.process(body, basic_deliver.delivery_tag)
        self.acknowledge_message(basic_deliver.delivery_tag)

    def on_cancelok(self, unused_frame):
        logging.info('RabbitMQ acknowledged the cancellation of the consumer')
        self.close_channel()

    def stop_consuming(self):
        if self._channel:
            logging.info('Sending a Basic.Cancel RPC command to RabbitMQ')
            self._channel.basic_cancel(self.on_cancelok, self._consumer_tag)

    def start_consuming(self):
        logging.info('Issuing consumer related RPC commands')
        self.add_on_cancel_callback()
        self._consumer_tag = self._channel.basic_consume(self.on_message,
                                                         self._queue)

    def on_bindok(self, unused_frame):
        logging.info('Queue bound')
        self.start_consuming()

    def close_channel(self):
        logging.info('Closing the channel')
        self._channel.close()

    def open_channel(self):
        logging.info('Creating a new channel')
        self._connection.channel(on_open_callback=self.on_channel_open)

    def run(self):
        self._connection = self.connect()
        self._connection.ioloop.start()

    def stop(self):
        logging.info('Stopping')
        self._closing = True
        self.stop_consuming()
        self._connection.ioloop.start()
        logging.info('Stopped')


class MessageProcessor(object):
    def process(self, message, delivery_tag):
        print "Process", message, delivery_tag
        yield tornado.gen.Task(c.incr, 'num_events')
        
class MainHandler(web.RequestHandler):
    def initialize(self):
        print "initialize"

    def get(self):
        print "INI1"
        value = yield tornado.gen.Task(c.get, 'num_events')
        print "INI2", value, type(value)
        #self.write("Hello, world %d" % value)
        self.write("Hello, world")
        
        
address = os.getenv('SERVER_ADDRESS', '127.0.0.1')
port = os.getenv('SERVER_PORT', 3000)


if __name__ == '__main__':
    tornado.options.parse_command_line()

    application = web.Application([ (r'/', MainHandler, dict()), ], debug=True)
    application.listen(port, address, xheaders=True)
    
    consumer = AsyncAmqpConsumer(
        os.environ['BROKER_URI'],
        os.environ['EXCHANGE'],
        os.environ['ROUTING_KEY'],
        os.environ['QUEUE'],
        MessageProcessor())
    try:

        consumer.run()        
        tornado.ioloop.IOLoop.instance().start()
    except KeyboardInterrupt:
        consumer.stop()
