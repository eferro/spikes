from flask import Flask
import socket

version = 1.1
app = Flask(__name__)

@app.route("/")
def hello():
    return "Hello World! from {} {}".format(socket.gethostname(), version)

if __name__ == "__main__":
    app.run(host='0.0.0.0', debug=True)
