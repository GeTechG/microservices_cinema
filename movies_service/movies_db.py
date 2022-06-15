import uuid

movies = []

def valid_movie(movie):
    not_empty = movie["name"] != None and movie["description"] != None and movie["min_age"] != None
    is_type = type(movie["name"]) is str and type(movie["description"]) is str and type(movie["min_age"]) is int
    return not_empty and is_type

def new_movie(request_movie):
    return {
        "id": str(uuid.uuid4()),
        "name": request_movie["name"],
        "description": request_movie["description"],
        "min_age": request_movie["min_age"],
    }

def add_new_movie(request_movie):
    if valid_movie(request_movie):
        movie = new_movie(request_movie)
        movies.append(movie)
        return movie["id"]

def get_movies():
    return movies

def get_movie_by_uuid(uuid_movie):
    for movie in movies:
        if movie["id"] == uuid_movie:
            return movie
