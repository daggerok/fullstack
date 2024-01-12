from flask import Flask, render_template
from pybars import Compiler

app = Flask(__name__)
compiler = Compiler()


@app.route('/')
def get_index_page():
    context = {'name': 'Max'}
    return render_template('index.hbs', **context)


if __name__ == '__main__':
    app.run(debug=True)
