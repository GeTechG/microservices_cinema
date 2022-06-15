use rocket::serde::uuid::Uuid;


#[derive(serde::Serialize, serde::Deserialize, Debug)]
pub struct Movie {
    pub id: Uuid,
    pub name: String,
    pub description: String,
    pub min_age: u16
}