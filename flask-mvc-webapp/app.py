from flask import Flask, render_template
from gevent.pywsgi import WSGIServer


app = Flask(__name__, static_url_path='')


@app.route('/')
def get_index_page():
    context = {'name': 'Max'}
    return render_template('index.hbs', **context)


if __name__ == '__main__':
    http_server = WSGIServer(('0.0.0.0', 3000), app)
    http_server.serve_forever()
