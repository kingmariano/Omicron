package database
import (
	"sync"
	"encoding/json"
	"os"
	"errors"
)

type ContactDB struct{
	path string
	mu   *sync.RWMutex
}
type DBStructure struct {
	Contacts map[int]Contact `json:"contacts"`
}
func NewDB(path string) (*ContactDB, error) {
	db := &ContactDB{
		path: path,
		mu:   &sync.RWMutex{},
	}
	err := db.ensureDB()
	return db, err
}
func (db *ContactDB) ensureDB() error {
	_, err := os.ReadFile(db.path)
	if errors.Is(err, os.ErrNotExist) {
		return db.CreateDB()
	}
	return err
}
func (db *ContactDB) CreateDB() error {
	dbStructure := DBStructure{
		Contacts: map[int]Contact{},
	}
	return db.WriteDB(dbStructure)
}
func (db *ContactDB) loadDB() (DBStructure, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	dbStructure := DBStructure{}
	dat, err := os.ReadFile(db.path)
	if errors.Is(err, os.ErrNotExist) {
		return dbStructure, err
	}
	err = json.Unmarshal(dat, &dbStructure)
	if err != nil {
		return dbStructure, err
	}
	return dbStructure, nil
}


func (db *ContactDB) WriteDB(dbstructure DBStructure) error{
	db.mu.Lock()
	defer db.mu.Unlock()
	dat, err := json.Marshal(dbstructure)
	if err != nil {
		return err
	}
	err = os.WriteFile(db.path, dat, 0600)
	if err != nil {
		return err
	}
	return nil
}