package main

// write you struct here and add declared struct into registry.go
type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string
}
type Test struct {
	Name    string   `json:"name"`
	Emails  []string `json:"email"`
	Authors []User   `json:"authors"`
}
