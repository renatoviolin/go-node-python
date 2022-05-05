use warp::Filter;

use crate::payload::Payload;

pub fn hello() -> impl warp::Reply {
    "hello world"
}

pub fn single_json() -> Box<dyn warp::Reply> {
    let payload = serde_json::from_str::<Payload>(
        r#"{
        "field1": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
        "field2": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
        "field4": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
        "field3": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
        "field5": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
        "detail": [
            {
                "field6": 123.345,
                "field7": 435345.34,
                "field8": 15345.323,
                "field9": 3232315400,
                "field10": 323235.3
            },
            {
                "field6": 123.345,
                "field7": 435345.34,
                "field8": 15345.323,
                "field9": 3232315400,
                "field10": 323235.3
            },
            {
                "field6": 123.345,
                "field7": 435345.34,
                "field8": 15345.323,
                "field9": 3232315400,
                "field10": 323235.3
            },
            {
                "field6": 123.345,
                "field7": 435345.34,
                "field8": 15345.323,
                "field9": 3232315400,
                "field10": 323235.3
            }
        ]
    }"#,
    );

    match payload {
        Ok(payload) => Box::new(warp::reply::json(&payload)),
        Err(error) => Box::new(warp::reply::with_status(
            format!("JSON error: {error:?}"),
            warp::http::StatusCode::BAD_REQUEST,
        )),
    }
}

pub async fn run(port: u16) {
    let hello = warp::get().and(warp::path!("hello").map(hello));
    let single_json = warp::get().and(warp::path!("single_json")).map(single_json);

    let routes = hello.or(single_json);

    println!("Warp server listening on http://localhost:{port}");
    warp::serve(routes).run(([127, 0, 0, 1], port)).await;
}
