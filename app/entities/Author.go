package entities

type Author struct {
	ID   int    `json:"id" gorm:"primary_key;auto_increment;not_null;"`
	Name string `json:"name"`
	Age  int    `json:"age"`

	Books []*Book `json:"books"`
}
