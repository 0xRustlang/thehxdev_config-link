# config-link

1. Clone this repository to `/root/`.
1. Put `clash.yaml`, `surfboard.conf` and `xray.txt` files (With exact same name) in `/root/config-link/` (Repository directory).
1. Run the app with `./app` command.
1. Server will start on port `3000`.
1. Then import files to Clash/Surfboard app:

```text
# Clash
http://SERVER_IP:3000/clash.yaml

# Surfboard
http://SERVER_IP:3000/surfboard.conf

# Any V2ray/Xray client that supports subscription
http://SERVER_IP:3000/xray.txt
```
