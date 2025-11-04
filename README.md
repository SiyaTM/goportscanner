# goportscanner
# GoPortScanner 🔍

A minimal, educational TCP port scanner written in Go.
This repo contains a simple single-file scanner (`scanner.go`) you can run in VS Code or the terminal to learn about networking, concurrency, and basic system programming concepts in Go.

---

## Features

* Parse a single port or a port range (e.g. `80` or `20-1024`)
* Attempt TCP connects with a timeout to detect open ports
* Small and easy to read — great as a learning/portfolio piece

---

## Usage

### Run directly with `go run`

From the folder containing `scanner.go`:

```bash
go run scanner.go
```

The example `scanner.go` file uses these defaults:

* `host` = `scanme.nmap.org`
* `portRange` = `20-1024`
* per-port timeout = `500ms`

You can edit those variables at the top of the file, or replace them with command-line flags if you extend the file later.

### Build a binary

```bash
go build -o goscanner scanner.go
./goscanner
```

---

## Example output

```
Scanning scanme.nmap.org on ports 20-1024...
Port 22 is OPEN
Port 80 is OPEN
Port 443 is OPEN
```

---

## How it works (high level)

1. `parsePortRange` — accepts `"80"` or `"20-1024"` and returns a list of ports to scan.
2. `scanPort` — uses `net.DialTimeout("tcp", address, timeout)` to attempt a TCP connection. Successful connect → port is open; otherwise closed/filtered.
3. `main` — iterates ports and prints open ones.

This directly demonstrates `net` networking, simple parsing, and timeout-based error handling.

---

## Suggested next steps / roadmap

Make small incremental commits so the repo shows progress:

* Add CLI flags (`-host`, `-ports`, `-timeout`) using the `flag` package.
* Use goroutines + channels to scan ports concurrently and control concurrency with a worker pool.
* Add UDP scanning (careful — UDP requires different handling).
* Add banner grabbing: read the first bytes from an open TCP connection to infer service.
* Output results to JSON/CSV.
* Add unit tests for `parsePortRange`.

---

## Notes & ethics

This tool is for **educational use only**. Only scan hosts you own or have explicit permission to test. Unauthorized scanning may be illegal or violate terms of service.

---

## License

MIT — feel free to reuse and adapt. If you publish artifacts, consider adding a `LICENSE` file.

---

If you want, I can:

* Turn `scanner.go` into a proper CLI with flags and a README usage section for those flags, or
* Create the concurrent worker-pool version and update the README with examples and benchmark notes. Which would you prefer?
