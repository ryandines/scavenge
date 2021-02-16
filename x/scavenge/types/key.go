package types

const (
	// ModuleName is the name of the module
	ModuleName = "scavenge"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)

const (
	// ScavengePrefix is for the hash
	ScavengePrefix = "scavenge-value-"
	// ScavengeCountPrefix is for the hash
	ScavengeCountPrefix = "scavenge-count-"
)

const (
	// CommitPrefix is for the hash
	CommitPrefix = "commit-value-"
	// CommitCountPrefix is for the hash
	CommitCountPrefix = "commit-count-"
)
