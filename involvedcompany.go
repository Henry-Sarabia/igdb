package igdb

//go:generate gomodifytags -file $GOFILE -struct InvolvedCompany -add-tags json -w

// InvolvedCompany represents a company involved in the development
// of a particular video game.
// For more information visit: https://api-docs.igdb.com/#involved-company
type InvolvedCompany struct {
	Company    int  `json:"company"`
	CreatedAt  int  `json:"created_at"`
	Developer  bool `json:"developer"`
	Game       int  `json:"game"`
	Porting    bool `json:"porting"`
	Publisher  bool `json:"publisher"`
	Supporting bool `json:"supporting"`
	UpdatedAt  int  `json:"updated_at"`
}
