#!/bin/bash

/etc/init.d/pwrstatd start

# disable automations
sleep 2
/usr/sbin/pwrstat -lowbatt -capacity 0 -shutdown off -active off -runtime 0
/usr/sbin/pwrstat -pwrfail -active off -shutdown off

/app/pwrstat-exporter
