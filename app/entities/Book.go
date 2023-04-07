package entities

type Book struct {
	ID        int    `json:"id" gorm:"primary_key;auto_increment;not_null;"`
	Name      string `json:"name"`
	Vendor    string `json:"vendor"`
	PageCount int    `json:"page_count"`

	AuthorID *int    `json:"author_id"`
	Author   *Author `json:"author"`
}
