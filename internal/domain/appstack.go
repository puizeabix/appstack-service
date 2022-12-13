package domain

type AppStack struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description,omitempty" bson:"description"`
	Icon        string `json:"icon,omitempty" bson:"icon"`
	Tenant      string `json:"tenant" bson:"tenant"`
}
