package serrors

// PostgresError maps a standardized name accross all the software to the postgres error code
var PostgresError map[string]string = map[string]string{
	"AlreadyExisting": "unique_violation",
	"InvalidUID":      "invalid_text_representation",
}
