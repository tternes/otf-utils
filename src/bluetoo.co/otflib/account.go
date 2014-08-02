package otflib

type Provider struct {
	Name string // "wunderlist" | "todoist"
}

// ---------------------------------------
// Account
// ---------------------------------------

type Account struct {
	Provider Provider
	Lists []List
	Inbox *List
}

func AddList(list *List, account *Account) {
	account.Lists = append(account.Lists, *list)
}




