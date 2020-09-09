# IPS (pronounced yips)

Command line utility for IP addresses.

## Usage

### cidr

The `-cidr` flag writes all IPs (in ascending order) belonging to the given CIDR value.

#### examples

Render all addresses in the prompt

```
ips -cidr 192.168.0.0/24
```

Output the result to a file

```
ips -cidr 192.168.0.0/24 > addresses.txt
```
