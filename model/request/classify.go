package request

type ClassifyList struct {
	ClassId     int    `json:"id" yaml:"id"`
	ClassParent int    `json:"class_parent" yaml:"class_parent"`
	CnName      string `json:"cn_name" yaml:"cn_name"`
	EnName      string `json:"en_name" yaml:"en_name"`
	Status      int    `json:"status" yaml:"status"`
	Sort        int    `json:"sort" yaml:"sort"`
	PageInfo
}
