package models

const (
	TargetTypeArticle int = 1
	TargetTypeTag     int = 2
)

type OperationLog struct {
	Model
	UserId        uint `gorm:"column:user_id" json:"user_id" gorm:"index:user_index"`
	TargetType    int  `gorm:"column:target_type" json:"target_type" gorm:"index:target_index"`
	TargetId      uint `gorm:"column:target_id" json:"target_id" gorm:"index:target_index"`
	OperationType uint `gorm:"column:operation_type" json:"operation_type"`
}

func GetLogByTarget(limit int, offset int, query map[string]interface{}) (log *[]OperationLog, err error) {
	err = nil
	DB.Model(&OperationLog{}).Limit(limit).Offset(offset).Where(query).Find(log)

	return
}

func CreateOperationLog(log *OperationLog) error {
	return DB.Model(&OperationLog{}).Create(log).Error
}

func DeleteOperationLog(log *OperationLog) (err error) {
	return DB.Model(&OperationLog{}).Delete(log).Error
}
