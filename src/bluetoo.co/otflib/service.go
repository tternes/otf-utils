package otflib

type Service interface {
	
	// Populates account data (lists, tasks)
	LoadServiceAccount() bool
	
	// Returns otf account object (populated by LoadServiceAccount())
	GetAccount() *Account
	
	// Adds the list and child tasks to the service
	AddListToService(l *List) error
	
	// Removes the list and child tasks from the service
	RemoveListFromService(l *List) error
}