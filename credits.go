package main

// CreditCategory codes
type CreditCategory int

// Credits is
type Credits struct {
	ID                    ID             `json:"id"`
	Name                  string         `json:"name"`
	Slug                  string         `json:"slug"`
	URL                   URL            `json:"url"`
	CreatedAt             int            `json:"created_at"` //unix epoch
	UpdatedAt             int            `json:"updated_at"` //unix epoch
	Game                  ID             `json:"game"`
	Category              CreditCategory `json:"category"`
	Company               ID             `json:"company"`
	Position              int            `json:"position"`
	Person                ID             `json:"person"`
	Character             ID             `json:"character"`
	Title                 ID             `json:"title"`
	Country               CountryCode    `json:"country"`
	CreditedName          string         `json:"credited_name"`
	CharacterCreditedName string         `json:"character_credited_name"`
}
