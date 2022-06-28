## Install Dependencies

### Install `mkcert`

<details>
<summary>
On macOS
</summary>

```
brew install mkcert
brew install nss # if you use Firefox
```

</details>

<details>
<summary>
On Linux
</summary>

#### Install [homebrew on Linux](https://docs.brew.sh/Homebrew-on-Linux)
#### Install `nss-tools`

```
sudo apt install libnss3-tools
# or
sudo yum install nss-tools
# or 
sudo pacman -S nss
# or
sudo zypper install mozilla-nss-tools
# or
sudo dnf install nss-tools
```
#### Install mkcert
```
brew install mkcert
```

</details>

### Install `envoy`

<details>
<summary>
Expand
</summary>

[Envoy Proxy Website](https://www.envoyproxy.io/)

```
brew install envoy
```
</details>

### Generate TLS certificates
<details>
<summary>Expand
</summary>

```
mkcert -install
mkcert -cert-file certs/tls.crt -key-file certs/tls.key core.dev "*.core.dev" localhost 127.0.0.1 ::1 
```
</details>

## Getting Started

### Start docker-compose file

```
make up
```

### Start the envoy proxy

```
make proxy
```

### Start the server

```
go run . serve --listen-address=:8080 --db-driver=postgres --db-dsn=postgres://postgres:postgres@localhost:5432/core?sslmode=disable --log-level=debug
# or
make serve
```

Go to [https://localhost:10000](https://localhost:10000)

The available usernames and passwords for logging in during development are in
the `/deploy/oidc-users.json` file.

