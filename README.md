# L4 PACKET MONITORING

```
Monitoring port connectivity by sending packet.

Usage:
  packet-mon [command]

Available Commands:
  help        Help about any command
  start       Start Packet Mon Service

Flags:
      --config string   config file (default is $HOME/.xxx.yaml)
```

## TCP Packet Sender

### TCP Port Checking

- Send SYN from SRC:PORT to DST:PORT
- Send SYN-ACK from DST:PORT to SRC:PORT

### UDP Port Checking

- Send from SRC:PORT to DST:PORT

## Packer Receiver & Uploader (1 min upload once)

- Add change filter and monitoring packet.
- Upload unique signatured packet info to central server.

## Config Downloader (1 min fetch once)

- Download config for local.

## Frequency

- Production: 1 packet/min
- Dev: 1 packet/sec

## TODO

- [o] Cobra for the command line. update config frequency and report frequency should use params
- [o] Daemonize
- [o] Systemd script
