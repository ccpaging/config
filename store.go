package config

type Store interface {
	// load retrieves the configuration stored.
	load() (map[string]interface{}, error)

	// save replaces the current configuration in its entirety.
	save(map[string]interface{}) error

	// Existed returns true if the file was persisted.
	Existed() (bool, error)

	// Remove removes persisted configuration file.
	Remove() error

	// String describes the backing store for the config.
	String() string

	// Close cleans up resources associated with the store.
	Close() error
}

func NewStore(dsn string) Store {
	return newFileStore(dsn)
}
