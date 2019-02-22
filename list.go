package igdb

//go:generate gomodifytags -file $GOFILE -struct List -add-tags json -w

type List struct {
	ID           int    `json:"id"`
	CreatedAt    int    `json:"created_at"`
	Description  string `json:"description"`
	EntriesCount int    `json:"entries_count"`
	ListEntries  []int  `json:"list_entries"`
	ListTags     []int  `json:"list_tags"`
	ListedGames  []int  `json:"listed_games"`
	Name         string `json:"name"`
	Numbering    bool   `json:"numbering"`
	Private      bool   `json:"private"`
	SimilarLists []int  `json:"similar_lists"`
	Slug         string `json:"slug"`
	UpdatedAt    int    `json:"updated_at"`
	URL          string `json:"url"`
	User         int    `json:"user"`
}
