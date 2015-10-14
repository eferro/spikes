import redis
import random
import datetime

r =redis.Redis(host='localhost', port=6379)
print r


t1 = datetime.datetime.now()
print t1
for i in range(2000, 10000):
	r.hset("test", "key_{}".format(i), "{},{},{}".format(
		random.choice([True, False]),
		random.choice(['lab:gpon1', 'lab:gpon2', 'lab:gpon3']),
		random.choice([True, False])
	))
t2 = datetime.datetime.now()
print "hset", t2, t2-t1

m = r.hgetall("test")
t3 = datetime.datetime.now()
for k,v in m.iteritems():

    "false, lab:gpon1, true"
    (bool,false), str, lab:gpon1)

	(a, b, c) = [t(s) for t,s in zip((bool, str, bool), v.split(','))]
	print (a, b, c)

print type(m)
print t3

