package dependencies

type Config struct {
	Port int
}

var Configs = map[string]Config{
	"strava-service": {Port: 666},
	"runner":         {Port: 80},
}
