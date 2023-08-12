package config

// Config is an interface that defines a method to retrieve configuration information.
type Config interface {
	Get() ConfigInfo
}

// ConfigInfo holds various configuration settings.
type ConfigInfo struct {
	Server               server               `yaml:"server"` // Server configuration
	PostgreSQLProperties postgreSQLProperties `yaml:"psql"`   // PostgreSQL database properties
}

// server represents server configuration settings.
type server struct {
	Port           int      `yaml:"port"`            // Port the server should listen on
	Host           string   `yaml:"host"`            // Host address for the server
	AllowedOrigins []string `yaml:"allowed_origins"` // List of allowed origins for CORS
	SecretKey      string   `yaml:"secret_key"`      // Secret key for JWT signing
	JWTLifespan    int      `yaml:"jwt_lifespan"`    // Lifespan of JWT tokens
}

// postgreSQLProperties holds properties for connecting to a PostgreSQL database.
type postgreSQLProperties struct {
	URL      string `yaml:"url"`      // Full URL PostgreSQL connection
	User     string `yaml:"user"`     // Username for PostgreSQL connection
	Password string `yaml:"password"` // Password for PostgreSQL connection
	Name     string `yaml:"name"`     // Name of the database
	Host     string `yaml:"host"`     // Host address of the PostgreSQL server
	Port     int    `yaml:"port"`     // Port number for PostgreSQL connection
}
