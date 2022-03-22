# Pwrstat exporter
Cyberpower UPS Linux daemon (pwrstat) exporter for prometheus 

## Installation
Must have Linux PowerPanel application from CyberPower already downloaded (https://www.cyberpowersystems.com/product/software/powerpanel-for-linux/)
And make sure the ``pwrstat.service`` has been started: `service pwrstat status`

### Build from source
**Please install golang 1.16** 

Install dep.
```bash
git clone https://github.com/kerwenwwer/pwrstat-exporter.git
cd pwrstat-exporter
go build && mv pwrstat-exporter /usr/local/bin/
```

## Usage
Since that ``pwrstat`` require sudo permission so:
```bash
sudo pwrstat-exporter 
```
Args
```bash
sudo pwrstat-exporter --web.listen-address 8088 --web.telemetry-path /metrics
```

## Systemd service
Create the service config
`nano /etc/systemd/system/pwrstat-exporter.service`

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

`systemctl daemon-reload`

`systemctl enable pwrstat-exporter`

`service pwrstat-exporter start`

## Dashboard
You can import the dashboard using ``grafana-dashboard.json``
![grafana](/image/grafana.png)

