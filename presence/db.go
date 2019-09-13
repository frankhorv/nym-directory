package presence

import (
	"time"

	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/directory-server/models"
)

// Db holds presence information
type Db interface {
	Add()
	Get()
	List()
}

type db struct {
	mixNodes map[string]models.MixNodePresence
}

func newPresenceDb() *db {
	return &db{
		mixNodes: map[string]models.MixNodePresence{},
	}
}

func (db db) Add(presence models.MixNodePresence) {
	db.killOldsters()
	db.mixNodes[presence.PubKey] = presence
}

func (db db) get(key string) models.MixNodePresence {
	return db.mixNodes[key]
}

func (db db) List() map[string]models.MixNodePresence {
	db.killOldsters()
	return db.mixNodes
}

func (db db) killOldsters() {
	for key := range db.mixNodes {
		presence := db.mixNodes[key]
		if presence.LastSeen < cutoff() {
			delete(db.mixNodes, key)
		}
	}
}

func cutoff() int64 {
	return timemock.Now().Add(time.Duration(-5 * time.Second)).Unix()
}
