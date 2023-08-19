# Simple remote clipboard

Simply send back data to the local system clipboard when working on the server via SSH/telnet.

It has two modes:

- Client mode
  Setup on the SSH server, it reads data from stdin and sends them to the server.
- Server mode 
  Setup on the local machine, it listens for the data from SSH server and sends it to the system clipboard.

## Usage

### Server mode

```bash
xcopy server -l 0.0.0.0 -p 9001
```

This will start the server mode and listen on port 9001 for data.

### Client mode

```bash
cat README.md | xcopy client -l <SERVER_IP> -p 9001
```

The SSH remote port forwarding can be set if the SSH connection to the server is available.

Create a remote port forwarding when establishing SSH connection:

```bash
ssh -R 9001:127.0.0.1:9001 <USER>@<SERVER_IP>
```

Or config it in `.ssh/config`, make it available on every connection:

```text
Host <SERVER_NAME>
    HostName <SERVER_IP>
    User <USER>>
    Port 22
    RemoteForward 9001 127.0.0.1:9001
```

After remote port forwarding is established, on the SSH server, you can just send to the local port 9001, which is default:

```bash
cat README.md | xcopy client
```
