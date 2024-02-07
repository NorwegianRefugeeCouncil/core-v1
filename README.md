## Install Dependencies

### Install [go](https://go.dev/doc/install)

### Install [docker](https://docs.docker.com/engine/install/)
On Linux, also ensure your user [has permissions to run docker commands](https://docs.docker.com/engine/install/linux-postinstall/#manage-docker-as-a-non-root-user)

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

### Generate TLS certificates
<details>
<summary>Expand
</summary>

```
make certs
```
</details>

## Getting Started

### Start docker-compose file

```
make up
```

### Start the server
First time run `make bootstrap` to generate css files.

```
go run . serve --listen-address=:8080 --db-driver=postgres --db-dsn=postgres://postgres:postgres@localhost:5432/core?sslmode=disable --log-level=debug
# or
make serve
```

Go to [https://localhost:10000](https://localhost:10000)

The available usernames and passwords for logging in during development are in
the `/deploy/oidc-users.json` file.

### Generate test participants
Run `go run . mock-data --count=X` where `X` is the amount of participants to generate. This will create a csv file that then can be uploaded to the system.

# Changing the form field

View documentation on field types [Form Fields](pkg/views/forms/README.md)

**WARNING: This is potentially incomplete**

## Database

Add a database migration

Create a new file in `internal/db/migrations/postgres/` in the format `NNN_description.up.sql` increasing the number from the previous migration.

Reference the created migration file `migrations` array in `internal/db/migrations.go`

## Backend

`internal/constants/individual.go`
For field `Foo`:
- Add constant `FormParamsGetIndividualFoo`
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

`internal/api/individual_list_options.go`
`internal/api/individual_list_options_encoder.go`
`internal/api/individual_list_options_decoder.go`
Add the field to make it searchable

`internal/api/validation/individual.go`
Add validation for the field if needed. E.g. max length for text fields

## Frontend

`web/templates/individuals.gohtml`
Add new column headers and body to display new field in the list/table in the expected position:

```
{{template "columnHeader" (dict
  "Options" .Options
  "Sortable" true
  "SortKey" "foo"
  "Label" (translate "foo")
  "Title" "Foo"
)}}
```

```
<!-- Foo -->
  <td>
    <div>
      {{.Foo}}
    </div>
  </td>
<!-- End Foo -->
```

`internal/views/individual_form.go`
Add a function call to the `fieldBuilders` to include the new field in the form in the expected position.

`web/templates/searchForm.gohtml`
Add a div block for search

### Translations
Create translation keys for the new field in the form/list and file header for all supported locales.

