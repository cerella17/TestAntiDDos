Anti-DDoS Project

Overview

This project is a backend solution to mitigate and monitor Distributed Denial of Service (DDoS) attacks. It leverages rate-limiting middleware, IP blacklisting, and metrics collection to protect an API endpoint. Designed with simplicity and security in mind, this project is ideal for testing and learning about anti-DDoS measures.

Features

Rate Limiting: Limits the number of requests per minute per IP address.

IP Blacklisting: Automatically blocks IPs exceeding the request threshold.

Request Queue: Queues excessive requests for future processing.

Real-Time Metrics: Integrates Prometheus and Grafana for monitoring blocked IPs and processed requests.

Proxy Compatibility: Supports real client IPs through X-Forwarded-For headers.

Technologies Used

Backend

Go: A high-performance, statically typed programming language for the backend.

Gin: A lightweight HTTP framework for routing and middleware.

Middleware

Custom Rate Limiter: Tracks and limits requests using in-memory counters.

IP Blocking: Uses iptables to block malicious IPs.

Queue Management

Redis: A fast, in-memory data store for handling the request queue.

Monitoring

Prometheus: Collects metrics for blocked IPs and queued requests.

Grafana: Visualizes Prometheus metrics with real-time dashboards.

Web Server

NGINX: Acts as a reverse proxy to forward traffic to the Go backend and handle real client IPs.

System Architecture

NGINX: Forwards incoming requests to the Go backend and ensures X-Forwarded-For headers are passed.

Gin Middleware:

Validates client IPs.

Applies rate limiting.

Queues excessive requests to Redis.

Blocks IPs via iptables if necessary.

Prometheus: Collects and exposes metrics about blocked IPs and processed requests.

Grafana: Displays metrics on a real-time dashboard.

Project Structure

├── main.go                # Entry point for the backend
├── middleware             # Contains rate limiting logic
│   └── rate_limiter.go
├── ip_blocker             # Handles IP blocking
│   └── ip_blocker.go
├── queue                  # Manages request queueing
│   └── queue.go
├── go.mod                 # Go module dependencies
├── go.sum                 # Go module checksums
└── README.md              # Project documentation

Design Decisions

Why Go?

High performance and concurrency support.

Simple and efficient memory management.

Strong ecosystem with libraries like Gin.

Why Redis for Queue Management?

Handles high-throughput data with low latency.

Reliable queueing mechanism for deferred processing.

Why Prometheus and Grafana?

Prometheus provides efficient metrics collection.

Grafana delivers powerful, customizable visualizations.

Why NGINX?

Handles real client IPs via X-Forwarded-For.

Provides reverse proxying and scalability.

Installation

Prerequisites

Go installed (>= 1.18).

Redis installed and running.

Prometheus and Grafana set up.

NGINX installed and configured.

Steps

Clone the repository:

git clone https://github.com/your-username/anti-ddos-project.git
cd anti-ddos-project

Install dependencies:

go mod tidy

Start the backend:

go run main.go

Configure NGINX:
Update the NGINX configuration file:

server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header Host $host;
    }
}

Restart NGINX:

sudo systemctl restart nginx

Set up Prometheus:
Add the following to your Prometheus configuration file:

scrape_configs:
  - job_name: 'anti-ddos'
    static_configs:
      - targets: ['localhost:8080']

Restart Prometheus to apply changes.

Access Grafana:
Import dashboards using Prometheus as the data source.

Usage

Endpoints

/api/data:

Example API endpoint protected by the rate limiter.

Returns:

{
    "message": "Benvenuto al backend Go!"
}

/metrics:

Exposes Prometheus metrics, including:

blocked_ips_total: Number of blocked IPs.

processed_request_total: Number of queued requests processed.

How to Test

Simulate a DDoS Attack

Use tools like hping3 or Apache JMeter:

hping3:

sudo hping3 -S -p 8080 -i u1000 <server-ip>

Apache JMeter:
Set up a Thread Group with multiple users and HTTP requests to /api/data.

Monitor Metrics

Access Prometheus:

URL: http://localhost:9090

Visualize metrics in Grafana:

URL: http://localhost:3000

Contributing

Feel free to open issues or submit pull requests to enhance this project.

License

This project is licensed under the MIT License. See the LICENSE file for details.

Acknowledgements

Gin Framework

Prometheus

Grafana

Redis

