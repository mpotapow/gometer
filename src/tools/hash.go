package tools

import (
	"crypto/md5"
	"encoding/hex"
	"sync"
)

var (
	once     sync.Once
	instance *Hash
)

// Hash ...
type Hash struct {
}

// GetHashInstance ...
func GetHashInstance() *Hash {
	once.Do(func() {
		instance = &Hash{}
	})
	return instance
}

// Make ...
func (h *Hash) Make(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Check ...
func (h *Hash) Check(text string, hash string) bool {
	return h.Make(text) == hash
}
