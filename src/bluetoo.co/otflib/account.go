package otflib

type Provider struct {
	Name string // "wunderlist" | "todoist"
}

// ---------------------------------------
// Account
// ---------------------------------------

type Account struct {
	Provider Provider
	Lists []*List
	Inbox *List
}

func NewAccount() *Account {
	return &Account{}
}

func AddList(list *List, account *Account) {
	account.Lists = append(account.Lists, list)
}




