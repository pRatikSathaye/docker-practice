from flask import Flask
app = Flask(__name__)


@app.route("/")
def sayHello():
  return "Hello Python..."

if __name__ == '__main__':
  app.run()