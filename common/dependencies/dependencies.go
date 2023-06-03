package dependencies

type Config struct {
	Port int
}

var Configs = map[string]Config{
	"strava-service":   {Port: 666},
	"database-service": {Port: 666},
	"http_gatewayTLS":  {Port: 443},
	"http_gateway":     {Port: 80},
	"runner":           {Port: 666},
	"populate-db":      {Port: 666},
}
