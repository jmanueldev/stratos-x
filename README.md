# STRATOS-X
STRATOS is a distributed HPC orchestration platform designed for
large-scale compute clusters.

Features

- mission-based scheduling
- adaptive compute topology
- distributed telemetry
- dynamic node grouping
- dataset staging

Architecture

API Layer
Scheduler
Topology Engine
Telemetry Network
Node Agents

Languages

Go -> control plane
Rust -> node agents
Python -> client SDK

Next-generation HPC orchestration with:

- Mission-based scheduling
- GPU / RDMA aware node selection
- Distributed telemetry & checkpointing
- Planetary-scale simulation
- Web-based topology editor
- Reinforcement-learning AI optimizer

## Usage

1. Start API: `make run-api`
2. Run node agent: `make run-agent`
3. Submit missions: `python sdk/client.py missions/example.yaml`
4. Launch web dashboard: `npm start` inside `web-dashboard/`

# Bash

- docker-compose up
- cargo run   # node agent
- npm start   # dashboard
- python sdk/client.py missions/example.yaml
