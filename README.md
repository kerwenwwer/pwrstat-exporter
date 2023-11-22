# Pwrstat exporter
Cyberpower UPS Linux daemon (pwrstat) exporter for prometheus 

## Deploy

### Docker

#### Prerequisites
- Docker

#### Installation

To install pwrstat exporter using Docker, simply run the following command.

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

or you can check the [docker-compose.yaml](https://github.com/kerwenwwer/pwrstat-exporter/blob/main/docker-compose.yaml) file

### Build from source

#### Prerequisites
- Golang 1.16

#### Installation
```bash
git clone https://github.com/kerwenwwer/pwrstat-exporter.git
cd pwrstat-exporter
go build && mv pwrstat-exporter /usr/local/bin/
```

#### Usage
Since that ``pwrstat`` require sudo permission so:
```bash
sudo pwrstat-exporter 
```
or with arguments
```bash
sudo pwrstat-exporter --web.listen-address 8088 --web.telemetry-path /metrics
```

#### Systemd service
Create the service config
`/etc/systemd/system/pwrstat-exporter.service`

```
[Unit]
Description=pwrstat-exporter

[Service]
TimeoutStartSec=0
ExecStart=/usr/local/bin/pwrstat-exporter

[Install]
WantedBy=multi-user.target
```
Restart the systemd daemon, set the service to start on boot, and start the service manually for the first time. 

```bash
sudo systemctl daemon-reload
```
```bash
sudo systemctl enable pwrstat-exporter
```
```bash
sudo service pwrstat-exporter start
```

## Grafana dashboard
You can import the dashboard using the [grafana-dashboard.json](https://github.com/kerwenwwer/pwrstat-exporter/blob/main/grafana-dashboard.json)

![grafana](/image/grafana.png)
