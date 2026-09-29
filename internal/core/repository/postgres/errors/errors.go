package core_postgres_errors

import (
	"fmt"
)

const PgForeignKeyViolation = "23503"

var ErrFKViolation = fmt.Errorf("violates foreign key")
