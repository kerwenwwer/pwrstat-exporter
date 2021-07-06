# Pwrstat exporter
Cyberpower UPS Linux daemon (pwrstat) exporter for prometheus 

## Installation
Must have Linux PowerPanel application from CyberPower already downloaded (https://www.cyberpowersystems.com/product/software/powerpanel-for-linux/)
And make sure the ``pwrstat.service`` has been staring.

### Build from source
**Please install golang 1.16** 

Install dep.
```bash
sudo apt install -y libsensors4-dev
git clone https://github.com/kerwenwwer/sensor-exporter.git
cd sensor-exporter/sensor-exporter
go build && mv sensor-exporter /usr/loca/bin/
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