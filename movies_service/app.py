from flask import Flask
from flask import request
import movies_db
import json
from flask import jsonify

app = Flask(__name__)

@app.route("/v1/new_movie", methods=['POST'])
def new_movie():
    new_movie = json.loads(request.data.decode("utf-8"))
    new_movie_uuid = movies_db.add_new_movie(new_movie)
    if (new_movie_uuid == None):
        return "Invalid request", 400
    return new_movie_uuid


@app.route("/v1/get_movie_by_uuid", methods=['POST'])
def get_movie_by_uuid():
    uuid = request.data.decode("utf-8")
    movie = movies_db.get_movie_by_uuid(uuid)
    if (movie == None):
        return "Invalid uuid", 400
    return jsonify(movie)
