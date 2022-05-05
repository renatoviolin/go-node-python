mod payload;
mod warp;


#[tokio::main]
async fn main() {
    let warp_server = tokio::spawn(async move {
        crate::warp::run(7100).await
    });
    
    // add other servers here if you want, then add them to the `join!` below.

    let (warp_result,) = tokio::join!(warp_server);

    if let Err(warp_err) = warp_result {
        println!("Warp server failed: {}", warp_err);
    }
}
