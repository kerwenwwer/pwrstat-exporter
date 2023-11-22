# Pwrstat Exporter
A Prometheus exporter for CyberPower UPS Linux daemon (pwrstat).

## Overview
The Pwrstat Exporter enables Prometheus to monitor data from CyberPower Uninterruptible Power Supply (UPS) systems running on Linux. It uses the pwrstat Linux daemon for data acquisition.

# Deployment
## Docker Deployment
Prerequisites
Docker installed on your system.
### Installation
Run the following command to install the Pwrstat Exporter using Docker:

```bash
docker run \
  --name pwrstat-exporter \
  --device /dev/bus/usb:/dev/bus/usb \
  --device /dev/usb/hiddev0:/dev/usb/hiddev0 \
  --privileged \
  --restart unless-stopped \
  -p 8088:8088 \
  -d cardif99/pwrstat-exporter:latest
```

Alternatively, refer to the provided [docker-compose.yaml](https://github.com/kerwenwwer/pwrstat-exporter/blob/main/docker-compose.yaml) for a Docker Compose setup.

## Building from Source
Prerequisites
* Golang version 1.16 or higher.
### Installation
Clone the repository and build the executable:

```bash
git clone https://github.com/kerwenwwer/pwrstat-exporter.git
cd pwrstat-exporter
go build && mv pwrstat-exporter /usr/local/bin/
```

### Usage
The ``pwrstat`` command requires sudo privileges:

```bash
sudo pwrstat-exporter 
```
To specify arguments:
```bash
sudo pwrstat-exporter --web.listen-address 8088 --web.telemetry-path /metrics
```

## Systemd Service Integration
Configuration
Create a systemd service configuration file at `/etc/systemd/system/pwrstat-exporter.service`:

```
[Unit]
Description=pwrstat-exporter

[Service]
TimeoutStartSec=0
ExecStart=/usr/local/bin/pwrstat-exporter

[Install]
WantedBy=multi-user.target
```

Reload the systemd daemon, enable the service at startup, and start the service:

```bash
sudo systemctl daemon-reload
sudo systemctl enable pwrstat-exporter
sudo service pwrstat-exporter start
```

## Grafana Integration
A custom Grafana dashboard is available for visualizing the data. Import the dashboard using the [grafana-dashboard.json](https://github.com/kerwenwwer/pwrstat-exporter/blob/main/grafana-dashboard.json) file.
![grafana](/image/grafana.png)
