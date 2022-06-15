use rocket::response::status::{BadRequest};
use rocket::serde::json::Json;

use crate::entities::movie::Movie;
use crate::entities::user::User;
use crate::utils;

#[post("/get_recommendation", data = "<user_id>")]
pub async fn get_recommendation(user_id: String) -> Result<Json<Vec<Movie>>, BadRequest<String>> {
    let user_data = reqwest::Client::new().post("http://localhost:8080/api/get_user")
        .body(user_id)
        .send().await.map_err(|e| BadRequest(Some(e.to_string())))?
        .json::<User>().await;

    if user_data.is_ok() {
        let user = user_data.unwrap();
        let user_age: u16 = utils::between_dates(user.date_birthday.date(), chrono::Utc::today()).years as u16;
        let movies_data = reqwest::get("http://localhost:5000/v1/get_movies")
            .await.map_err(|e| BadRequest(Some(e.to_string())))?
            .json::<Vec<Movie>>().await;
        println!("{:?}", user_age);
        if movies_data.is_ok() {
            let movies = movies_data.unwrap();
            let movies: Vec<Movie> = movies.
                into_iter().filter(|movie| user_age >= movie.min_age).collect();
            return Ok(Json(movies));
        }
        return Err(BadRequest(Some(movies_data.unwrap_err().to_string())))
    }
    Err(BadRequest(Some(user_data.unwrap_err().to_string())))
}