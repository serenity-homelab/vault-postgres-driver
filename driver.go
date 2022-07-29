package vpd

import (
	"database/sql"
	"database/sql/driver"
	"strings"

	pq "github.com/lib/pq"
	sidecar "github.com/serenity-homelab/sidecar"
)

type VaultPostgresDriver struct {
	*pq.Driver
}

var filename string = "database.json"

// logic to fetch password from vault file
func updateDsn(dsn string) (string, error) {

	creds, err := sidecar.GetDatabaseCreds(filename)
	if err != nil {
		return "", err
	}

	dsn = strings.ReplaceAll(dsn, "$1", creds.Username)
	dsn = strings.ReplaceAll(dsn, "$2", creds.Password)

	return dsn, nil
}

func (d VaultPostgresDriver) Open(dsn string) (driver.Conn, error) {

	updateddsn, err := updateDsn(dsn)

	if err != nil {
		return nil, err
	}

	// Pass down the dsn with password to postgres] driver's open function
	return d.Driver.Open(updateddsn)

}

func SetVaultFileName(fn string) {
	filename = fn
}

// When initialised will register the driver in sql package
func init() {
	sql.Register("vault-postgres-driver", &VaultPostgresDriver{&pq.Driver{}})
}
