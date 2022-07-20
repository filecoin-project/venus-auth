package storage

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/venus-auth/core"
	"time"
)

var MockToken = "test-token-01"
var MockKeyPair = KeyPair{
	Name:       "test-token-01",
	Perm:       "admin",
	Secret:     "d6234bf3f14a568a9c8315a6ee4f474e380beb2b65a64e6ba0142df72b454f4e",
	Extra:      "",
	Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiemwtdG9rZW4iLCJwZXJtIjoiYWRtaW4iLCJleHQiOiIifQ.DQ-ETWoEnNrpGKCikwZax6YUzdQIkhT0pHOTSta8770",
	CreateTime: time.Now(),
	IsDeleted:  0,
}

type MockStore struct{}

func (m MockStore) Get(token Token) (*KeyPair, error) {
	if token == Token(MockToken) {
		return &MockKeyPair, nil
	}
	return nil, nil
}

func (m MockStore) GetTokenRecord(token Token) (*KeyPair, error) {
	return nil, nil
}
func (m MockStore) ByName(name string) ([]*KeyPair, error) {
	return nil, nil
}
func (m MockStore) Put(kp *KeyPair) error {
	return nil
}
func (m MockStore) Delete(token Token) error {
	return nil
}
func (m MockStore) Has(token Token) (bool, error) {
	return false, nil
}
func (m MockStore) List(skip, limit int64) ([]*KeyPair, error) {
	return nil, nil
}
func (m MockStore) UpdateToken(kp *KeyPair) error {
	return nil
}

// user
func (m MockStore) HasUser(name string) (bool, error) {
	return false, nil
}
func (m MockStore) GetUser(name string) (*User, error) {
	return nil, nil
}

// GetUserRecord return a user, whether deleted or not
func (m MockStore) GetUserRecord(name string) (*User, error) {
	return nil, nil
}
func (m MockStore) HasMiner(maddr address.Address) (bool, error) {
	return false, nil
}
func (m MockStore) PutUser(*User) error {
	return nil
}
func (m MockStore) UpdateUser(*User) error {
	return nil
}
func (m MockStore) ListUsers(skip, limit int64, state int, sourceType core.SourceType, code core.KeyCode) ([]*User, error) {
	return nil, nil
}
func (m MockStore) DeleteUser(name string) error {
	return nil
}

// rate limit
func (m MockStore) GetRateLimits(name, id string) ([]*UserRateLimit, error) {
	return nil, nil
}
func (m MockStore) PutRateLimit(limit *UserRateLimit) (string, error) {
	return "", nil
}
func (m MockStore) DelRateLimit(name, id string) error {
	return nil
}

// miner
// first returned bool, 'miner' is created(true) or updated(false)
func (m MockStore) UpsertMiner(miner address.Address, userName string) (bool, error) {
	return false, nil
}

// first returned bool, if miner exists(true) or false
func (m MockStore) DelMiner(miner address.Address) (bool, error) {
	return false, nil
}
func (m MockStore) GetUserByMiner(miner address.Address) (*User, error) {
	return nil, nil
}
func (m MockStore) ListMiners(user string) ([]*Miner, error) {
	return nil, nil
}

func (m MockStore) Version() (uint64, error) {
	return 0, nil
}
func (m MockStore) MigrateToV1() error {
	return nil
}
