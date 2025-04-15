package main

import (
	"dat-package/redisclient"
	"dat-package/startup"
)

func main() {
	redisclient.InitRedis()
	startup.UpdateRedisCacheOnStartup()

	// Continue with the rest of your service startup
}
