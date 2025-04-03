package main

type Contact struct {
	ID         int
	Username   string
	GivenName  string
	FamilyName string
	Phone      []struct {
		TypeID      int
		CountryCode int
		Operator    int
		Number      int
	}
	Email     []string
	Birthdate string
}

type Group struct {
	ID          int
	Title       string
	Description string
	Contacts    []int
}
