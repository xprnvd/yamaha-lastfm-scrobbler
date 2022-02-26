use serde::{Deserialize, Serialize};
use hyper::body::HttpBody as _;
use hyper::Client;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
    // define http client
    let client = Client::new();
    // Parse an `http::Uri`...
    let uri = "http://192.168.0.243:80/YamahaExtendedControl/v1/netusb/getPlayInfo".parse()?;

    // Await the response...
    let mut resp = client.get(uri).await?;

    println!("Response: {}", resp.status());

    // Read the response body...
    while let Some(chunk) = resp.body_mut().data().await {
        let aaaa = &chunk?;
        let s = String::from_utf8_lossy(aaaa);
        println!("{:?}", s);

    }
    
    Ok(())
}
