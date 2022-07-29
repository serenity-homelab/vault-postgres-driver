# vault-postgres-driver

Custom driver to handle vault cred rotation


### SetVaultFileName
`SetVaultFileName(string fn)` - Call to set the sidecar file name where the database creds are located. Default `database.json`

### Creating your DSN
This driver with replace the string `$1` in your dsn with the Username, and `$2` with the Password.
So create a dsn similar to this `host=127.0.0.1 port=5432 user=$1 password=$2 dbname=example sslmode=disable`

```go
package main

import (
  "database/sql"
  "log"
  
  _ "github.com/serenity-homelab/vault-postgres-driver"
)

func main() {
	psqlInfo := "host=127.0.0.1 port=5432 user=$1 password=$2 dbname=example sslmode=disable"

	db, err := sql.Open("vault-postgres-driver", psqlInfo)

	if err != nil {
		log.Panic(err)
	}

  log.Info("successful connection")
}
```