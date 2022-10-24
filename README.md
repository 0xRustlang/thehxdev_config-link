# config-link

1. Put all of your xray/v2ray configurations to `clash.yaml` (For Clash) and `surfboard.conf` (For Surfboard)
1. Run the app with `./app` command.
1. Server will start on port `3000`.
1. Then import files to Clash/Surfboard app:

```text
# Clash
http://SERVER_IP:3000/clash.yaml

# Surfboard
http://SERVER_IP:3000/surfboard.conf
```
