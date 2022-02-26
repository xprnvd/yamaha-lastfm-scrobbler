

#[tokio::main]
async fn get_artist() -> Result<(), Box<dyn std::error::Error>> {
    // Build the client using the builder pattern
    let client = reqwest::Client::builder()
        .build()?;

    // Perform the actual execution of the network request
    let res = client
        .get("http://192.168.0.243:80/YamahaExtendedControl/v1/netusb/getPlayInfo")
        .send()
        .await?;

    // Parse the response body as Json in this case
    let ip = res
        .json::<Ip>()
        .await?;

    println!("{:?}", ip);
    Ok(())
}
