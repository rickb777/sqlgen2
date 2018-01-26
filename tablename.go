package sqlgen2

// TableName holds a two-part name. The prefix part is optional.
type TableName struct {
	// Prefix on the table name. It can be used as the schema name, in which case
	// it should include the trailing dot. Or it can be any prefix as needed.
	Prefix string

	// The principal name of the table.
	Name string
}

// String gets the full table name.
func (tn TableName) String() string {
	return tn.Prefix + tn.Name
}

// PrefixWithoutDot return the prefix; if this ends with a dot, the dot is removed.
func (tn TableName) PrefixWithoutDot() string {
	last := len(tn.Prefix)-1
	if last > 0 && tn.Prefix[last] == '.' {
		return tn.Prefix[:last]
	}
	return tn.Prefix
}
