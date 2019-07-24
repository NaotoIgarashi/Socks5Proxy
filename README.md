# HTTP Access via Socks5 Proxy

You can specify a proxy server for each HTTP request. In this sample, http://inet-ip.info/ip endpoint will return different access source IP address for each proxy.

# Getting Started

## Create Proxy Servers

Create Linux server on different sites. (e.g. Azure, AWS, Other region...)

## Setting for All Proxy Severs 

```shell
sudo vi /etc/ssh/sshd_config

# Enable "AllowTcpForwarding yes" setting

sudo systemctl restart sshd
```

## Proxy Client setting

ssh to each ssh server as Socks5 Proxy

```sh
ssh <username>@<proxy server ip or url> -N -f -D <portnumber>
```

```sh
ssh ec2-user@ec2-18-182-66-111.ap-northeast-1.compute.amazonaws.com -N -f -D 10001
```

## Program Setting

Set the proxy port in config.json

```json
{
	"url": "http://inet-ip.info/ip",
	"proxyPorts": [
        "10001",
        "10002"
    ]
}
```

## Run

```sh
go run main.go
```

# Notes
When execute from shell
```sh
curl --proxy socks5://localhost:10001 http://inet-ip.info/ip
```