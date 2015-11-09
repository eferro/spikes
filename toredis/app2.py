import tornado.httpserver
import tornado.web
import tornado.ioloop
import tornado.gen
import logging
from toredis import Client

logging.basicConfig(level=logging.DEBUG)

log = logging.getLogger('app')


class MainHandler(tornado.web.RequestHandler):

    @tornado.web.asynchronous
    @tornado.gen.engine
    def get(self):
        redis = Client()
        foo = yield tornado.gen.Task(redis.get, 'foo')
        bar = yield tornado.gen.Task(redis.get, 'bar')
        zar = yield tornado.gen.Task(redis.get, 'zar')        
        
        self.set_header('Content-Type', 'text/html')
        self.render("template.html", title="Simple demo",  foo=foo, bar=bar, zar=zar)


application = tornado.web.Application([
    (r'/', MainHandler),
])

@tornado.gen.engine
def create_initial_data():
    redis = Client()
    redis.connect('localhost')
    result = yield tornado.gen.Task(redis.set, 'foo', 10)
    print "EFA1", result
    result = yield tornado.gen.Task(redis.set, 'bar', 20)
    print "EFA2", result    
    result = yield tornado.gen.Task(redis.set, 'zar', 30)
    print "EFA3", result    

if __name__ == '__main__':
    create_initial_data()
    http_server = tornado.httpserver.HTTPServer(application)
    http_server.listen(8888)
    print 'Demo is runing at 0.0.0.0:8888\nQuit the demo with CONTROL-C'
    tornado.ioloop.IOLoop.instance().start()
