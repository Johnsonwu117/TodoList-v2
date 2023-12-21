package member

type Table struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	UserPassword string `json:"userpassword" gorm:"column:userpassword"`
	UserEmail    string `json:"useremail"`
	Identify     string `json:"identify"`
}

type Base struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	UserPassword string `json:"userpassword"`
	UserEmail    string `json:"useremail"`
	Identify     string `json:"identify"`
}

type Create struct {
	Username     string `json:"username"`
	UserPassword string `json:"userpassword" gorm:"column:userpassword"`
	UserEmail    string `json:"useremail" gorm:"column:useremail"`
	Identify     string `json:"identify"`
}

func (t *Table) TableName() string {
	return "members"
}

func (t *Create) TableName() string {
	return "members"
}
