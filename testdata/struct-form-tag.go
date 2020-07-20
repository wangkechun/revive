package fixtures

type GetProjectRequest struct {
	ID   uint
	Name string `form:"namex"` // MATCH /tag form and json should exist simultaneously/
	Desc string `form:"desc" json:"desc"`
	Desc2 string `form:"desc2" json:"descxxx"` // MATCH /tag form and json should equal/
}
