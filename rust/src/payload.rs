use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct Payload {
    field1: String,
    field2: String,
    field3: String,
    field4: String,
    field5: String,
    detail: Vec<PayloadDetail>,
}

#[derive(Serialize, Deserialize)]
pub struct PayloadDetail {
    field6: f64,
    field7: f64,
    field8: f64,
    field9: f64,
    field10: f64,
}
