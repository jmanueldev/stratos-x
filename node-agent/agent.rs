use sysinfo::{System, SystemExt, CpuExt};
use tokio::time::{sleep, Duration};
use serde::{Serialize};
use reqwest::Client;

#[derive(Serialize)]
struct Telemetry {
    node_id: String,
    cpu: f32,
    memory: u64,
}

#[tokio::main]
async fn main() {
    let node_id = "node-1";
    let client = Client::new();
    let mut sys = System::new_all();

    loop {
        sys.refresh_all();
        let cpu = sys.global_cpu_info().cpu_usage();
        let memory = sys.used_memory();

        let data = Telemetry {
            node_id: node_id.to_string(),
            cpu,
            memory,
        };

        let _ = client.post("http://localhost:8080/telemetry")
            .json(&data)
            .send()
            .await;

        sleep(Duration::from_secs(2)).await;
    }
}
