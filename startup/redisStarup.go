package startup

import (
	"context"
	"dat-package/data"
	"dat-package/redisclient"
	"encoding/json"
	"log"
)

func UpdateRedisCacheOnStartup() {
	components := data.GetPredefinedComponents()
	for _, component := range components {
		key := "package:" + component.Name
		value, err := json.Marshal(component)
		if err != nil {
			log.Printf("Error marshalling component %s: %v", component.Name, err)
			continue
		}
		err = redisclient.Rdb.Set(context.Background(), key, value, 0).Err()
		if err != nil {
			log.Printf("Error setting Redis key %s: %v", key, err)
		} else {
			log.Printf("Cached component %s in Redis.", component.Name)
		}
	}
}
