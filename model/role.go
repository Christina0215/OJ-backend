package model

type Role struct {
	ID     int     `gorm:"primary_key;unique"`
	Alias  string  `gorm:"unique:not null"`
	Name   string  `gorm:"unique;not null"`

	User   []User  `gorm:"ForeignKey:ID;AssociationForeignKey:RoleId"`
}

func (role *Role) GetData() map[string]interface{}  {
	return map[string]interface{}{
		"roleId": role.ID,
		"alias":  role.Alias,
		"name":   role.Name,
	}
}
