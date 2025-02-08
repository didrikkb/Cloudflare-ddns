# Cloudflare-ddns
Configure the settings in config.conf using the format below.
config.conf needs to be in the same directory as the executable.
## Configuration

mail: mail@example.com<br />
key: <apiKey><br />
zone: <zoneID><br />
<br />
record<br />
name: www.example.com<br />
type: <A|AAAA><br />
proxied: <true|false> #Optional, default is false<br />
ttl: 60 #Optional, default is 300<br />
comment:Some comment 1 #Optional<br />
end<br />
<br />
record<br />
name: dns.example.com<br />
type: <A|AAAA><br />
proxied: <true|false> #Optional, default is false<br />
ttl: 3600 #Optional, default is 300<br />
comment:Some comment 2 #Optional<br />
end<br />