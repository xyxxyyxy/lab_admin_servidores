from time import gmtime, strftime
from flask import Flask


app = Flask(__name__)


@app.route('/clock')
def getTime():
    return '{}\n'.format(strftime("%Y-%m-%d %H:%M:%S", gmtime()))

if __name__ == "__main__":
    app.run(host="0.0.0.0", debug=True)