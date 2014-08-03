package otflib

type Service interface {
	
	// Populates account data (lists, tasks)
	LoadServiceAccount() bool
	
	// Returns otf account object (populated by LoadServiceAccount())
	GetAccount() *Account
}