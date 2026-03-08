# stratos-x
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
