package models

type RegForm struct {
	Raw string   `form:"raw_text"`
	Reg string   `form:"reg_text"`
	Res []string `form:"-"`
}
