## Install Dependencies

### Install [go](https://go.dev/doc/install)

### Install [docker](https://docs.docker.com/engine/install/)

### Install `yarn`
<details>
<summary>
Expand
</summary>

```
npm install --global yarn
```

</details>

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


# Changing the form field

View documentation on field types [Form Fields](pkg/views/forms/README.md)

**WARNING: This is potentially incomplete**

## Database

Add a database migration
TODO: Add more details

## Backend

`internal/constants/individual.go`
For field `Foo`:
- Add constant `FormParamIndividualFoo`
- Add constant `FormParamGetIndividualFoo`
- Add constant `DBColumnIndividualFoo` and add to `IndividualDBColumns` array
- Add constant `FileColumnIndividualFoo` and add to `IndividualFileColumns` array
- Update `IndividualDBToFileMap` and `IndividualFileToDBMap` maps

`internal/api/individual.go`
Add the field to the individual struct, eg:
```
type Individual struct {
    ...
    Foo string `db:"foo"`
    ...
}
```
`GetFieldValue`: add/remove/change the field in the switch statement, eg:
```
case constants.DBColumnIndividualFoo:
		return i.Foo, nil
```
`Normalize` If required, add a normalise step for the field, eg:
```
individual.Foo = trimString(individual.Foo)
```

`internal/api/individual_tabular.go`
`unmarshalTabularData`: add/remove/change the field in the switch statement, eg:
```
case constants.FormParamGetIndividualFoo:
    i.Foo = cols[idx]
```
`marshalTabularData` If the field is not a string, add a formatting step to the switch statement, eg:
```
case constants.FileColumnIndividualFoo:
  row[j] = strconv.FormatInt(value, 10)
```
If a column of the given type already exists, add the field to the switch statement, eg:
```
case constants.FileColumnIndividualIsMinor, constants.FileColumnIndividualPresentsProtectionConcerns, constats.FileColumnIndividualFoo:
	row[j] = strconv.FormatBool(value.(bool))
```

## Frontend

