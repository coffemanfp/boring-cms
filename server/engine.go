package server

// Engine defines the contract for an engine that can run a server.
type Engine interface {
	// Run starts the server on the specified addresses.
	// It returns an error if the server fails to start.
	Run(addr ...string) error
}
