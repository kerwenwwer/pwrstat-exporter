#!/bin/bash

/etc/init.d/pwrstatd start

# disable automations
pwrstat -lowbatt -capacity 0 -shutdown off -active off -runtime 0
pwrstat -pwrfail -active off -shutdown off

/app/pwrstat-exporter
