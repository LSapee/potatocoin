package db

import (
	"github.com/LSapee/potatocoin/utils"
	"github.com/boltdb/bolt"
)

const (
	dbName       = "blockchain.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
)

var db *bolt.DB

// DB생성
func DB() *bolt.DB {
	if db == nil {
		//init db
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		utils.HandleErr(err)
		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = tx.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		db = dbPointer
	}
	return db
}
