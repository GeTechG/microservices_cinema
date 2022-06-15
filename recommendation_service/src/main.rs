mod routes;
mod entities;
mod utils;

#[macro_use] extern crate rocket;

use crate::routes::get_recommendation;

#[rocket::main]
async fn main() -> Result<(), ()> {
    rocket::build()
        .mount("/v1", routes![get_recommendation])
        .launch().await;

    Ok(())
}