package database

type Classify struct {
	ClassId     int    `json:"class_id" yaml:"class_id"`
	ClassParent int    `json:"class_parent" yaml:"class_parent"`
	CnName      string `json:"cn_name" yaml:"cn_name"`
	EnName      string `json:"en_name" yaml:"en_name"`
	Status      int    `json:"status" yaml:"status"`
	Sort        int    `json:"sort" yaml:"sort"`
}

func (Classify) TableName() string {
	return "classify"
}
