# Initial Login
 MXR comes with internal bridge running by default. So teh detector can be eccessted in teh external port.

 ## External IP Address
 The external IP address is 192.168.1.10

 ## MXR Login Credentials

Default login is:
username: pi
password: raspberry

## SSH interface
Device is preconfigured to accept ssh connections on port 22

It can be connect on external port

Host: 192.168.1.10
Use the credential provided above

## Serial Connection

Device can be connected via serial copnnection using teh credential above.
BAUD rate is 115200 bps.

# Power on and off detector

To turn on the detector power use:

``` 
/mxr/power/mxr -on
```

To turn off the detector ower use:

```
/mxr/power/mxr -off
```


# Pico Sence COnnection
PicoSence can be connected to the device.
Set teh device IP address as 192.168.1.10 in the settings.txt

# Setup

## Dual ethernet

For the second ethernet intsall a driver as follows:

```
sudo apt-get install linux-header
wget https://files.waveshare.com/upload/e/ee/CM4-DUAL-ETH-MINI-Example.zip
unzip CM4-DUAL-ETH-MINI-Example.zip -d ./CM4-DUAL-ETH-MINI-Example
cd CM4-DUAL-ETH-MINI-Example/Driver/
tar vjxf r8168-8.050.03.tar.bz2
cd r8168-8.050.03
sudo ./autorun.sh
```



## Static IP addresses
```
sudo nmtui
```

```
sudo apt-get install dhcpcd5
nano /etc/dhcpcd.conf
```

```
...

interface eth0
static ip_address=192.168.1.10/24
static routers=192.168.1.254
static domain_name_servers=8.8.8.8

interface eth1
static ip_address=192.168.0.10/24
static domain_name_servers=8.8.8.8

...
```
# Ethernet Bridge

Ethernet bridge code is written in GO. It forwards UDP trafic between PC connected to poer eth0 and detector connected to eth1. The bridge is configures an a serce describe below.

bridge can be run as:
```
/mxr/bridge/mxbridge [-t]
```
-t enables the traffic log on screen

## Compile
```
go build -o mxbridge main.go
```
## Service
```
cp /mxr/bridge/mxbridge.service  /etc/systemd/system/mxbridge.service
systemctl enable mxbridge.service  
systemctl start mxbridge.service 
```
### Start Bridge
```
service mxbridge strart
```
### Stop Bridge
```
service mxbridge stop
```

# Startup Code

On startup the code automaticallt turns on teh detection power

```
sudo nano /etc/rc.local
```
add
```
su - pi -c /mxr/power/mxr
```



starts mxbridge.service


