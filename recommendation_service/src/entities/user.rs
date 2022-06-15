use chrono::{DateTime, Utc};
use rocket::serde::uuid::Uuid;

#[derive(serde::Deserialize, Debug)]
#[serde(crate = "rocket::serde")]
pub struct User {
    pub id: Uuid,
    pub login: String,
    pub date_birthday: DateTime<Utc>
}