package utils

import (
	"regexp"
	"strings"
)

// Reserved is a map of reserved words
var Reserved = map[string]bool{
	"AES128":            true,
	"AES256":            true,
	"ALL":               true,
	"ALLOWOVERWRITE":    true,
	"ANALYSE":           true,
	"ANALYZE":           true,
	"AND":               true,
	"ANY":               true,
	"ARRAY":             true,
	"AS":                true,
	"ASC":               true,
	"AUTHORIZATION":     true,
	"BACKUP":            true,
	"BETWEEN":           true,
	"BINARY":            true,
	"BLANKSASNULL":      true,
	"BOTH":              true,
	"BYTEDICT":          true,
	"CASE":              true,
	"CAST":              true,
	"CHECK":             true,
	"COLLATE":           true,
	"COLUMN":            true,
	"CONSTRAINT":        true,
	"CREATE":            true,
	"CREDENTIALS":       true,
	"CROSS":             true,
	"CURRENT_DATE":      true,
	"CURRENT_TIME":      true,
	"CURRENT_TIMESTAMP": true,
	"CURRENT_USER":      true,
	"CURRENT_USER_ID":   true,
	"DEFAULT":           true,
	"DEFERRABLE":        true,
	"DEFLATE":           true,
	"DEFRAG":            true,
	"DELTA":             true,
	"DELTA32K":          true,
	"DESC":              true,
	"DISABLE":           true,
	"DISTINCT":          true,
	"DO":                true,
	"ELSE":              true,
	"EMPTYASNULL":       true,
	"ENABLE":            true,
	"ENCODE":            true,
	"ENCRYPT":           true,
	"ENCRYPTION":        true,
	"END":               true,
	"EXCEPT":            true,
	"EXPLICIT":          true,
	"FALSE":             true,
	"FOR":               true,
	"FOREIGN":           true,
	"FREEZE":            true,
	"FROM":              true,
	"FULL":              true,
	"GLOBALDICT256":     true,
	"GLOBALDICT64K":     true,
	"GRANT":             true,
	"GROUP":             true,
	"GZIP":              true,
	"HAVING":            true,
	"IDENTITY":          true,
	"IGNORE":            true,
	"ILIKE":             true,
	"IN":                true,
	"INITIALLY":         true,
	"INNER":             true,
	"INTERSECT":         true,
	"INTO":              true,
	"IS":                true,
	"ISNULL":            true,
	"JOIN":              true,
	"LEADING":           true,
	"LEFT":              true,
	"LIKE":              true,
	"LIMIT":             true,
	"LOCALTIME":         true,
	"LOCALTIMESTAMP":    true,
	"LUN":               true,
	"LUNS":              true,
	"LZO":               true,
	"LZOP":              true,
	"MINUS":             true,
	"MOSTLY13":          true,
	"MOSTLY32":          true,
	"MOSTLY8":           true,
	"NATURAL":           true,
	"NEW":               true,
	"NOT":               true,
	"NOTNULL":           true,
	"NULL":              true,
	"NULLS":             true,
	"OFF":               true,
	"OFFLINE":           true,
	"OFFSET":            true,
	"OLD":               true,
	"ON":                true,
	"ONLY":              true,
	"OPEN":              true,
	"OR":                true,
	"ORDER":             true,
	"OUTER":             true,
	"OVERLAPS":          true,
	"PARALLEL":          true,
	"PARTITION":         true,
	"PERCENT":           true,
	"PLACING":           true,
	"PRIMARY":           true,
	"RAW":               true,
	"READRATIO":         true,
	"RECOVER":           true,
	"REFERENCES":        true,
	"REJECTLOG":         true,
	"RESORT":            true,
	"RESTORE":           true,
	"RIGHT":             true,
	"SELECT":            true,
	"SESSION_USER":      true,
	"SIMILAR":           true,
	"SOME":              true,
	"SYSDATE":           true,
	"SYSTEM":            true,
	"TABLE":             true,
	"TAG":               true,
	"TDES":              true,
	"TEXT255":           true,
	"TEXT32K":           true,
	"THEN":              true,
	"TO":                true,
	"TOP":               true,
	"TRAILING":          true,
	"TRUE":              true,
	"TRUNCATECOLUMNS":   true,
	"UNION":             true,
	"UNIQUE":            true,
	"USER":              true,
	"USING":             true,
	"VERBOSE":           true,
	"WALLET":            true,
	"WHEN":              true,
	"WHERE":             true,
	"WITH":              true,
	"WITHIN":            true,
	"WITHOUT":           true,
}

// var ident = regexp.MustCompile(`(?i)^[a-z_][a-z0-9_$]*$`)

// special chars [~!@#$%^&*()-_+={}[]|\/:;"'<>,.?]
var specialChars = regexp.MustCompile(`[~!@#$%^&*()+={}[]|\/:;"'<>,.?]`)

// EscapeSpecial escapes special characters in a string
func EscapeSpecial(s string) string {
	matches := specialChars.FindAllStringSubmatch(s, -1)

	for _, match := range matches {
		//nolint
		s = strings.Replace(s, match[0], ``, -1)
	}

	return s
}

// IdentNeedsQuotes checks if the given identifier requires quoting
func IdentNeedsQuotes(s string) string {
	if Reserved[strings.ToUpper(s)] {
		s = QuoteIdent(s)
		return s
	}

	return s
}

// QuoteIdent quotes the given identifier string.
func QuoteIdent(s string) string {
	s = strings.Replace(s, `"`, `""`, -1) //nolint
	return `'` + s + `'`
}

func UniqueArrayOfString(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
