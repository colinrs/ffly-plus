package models

// CustomeModel a basic GoLang struct which includes the following fields: ID, CreatedAt, UpdatedAt, DeletedAt
// It may be embedded into your model or you may build your own model without it
//    type User struct {
//      CustomeModel
//    }
type CustomeModel struct {
	ID        uint   `gorm:"primarykey"`
	CreatedAt uint64 `json:"created_at"`
	UpdatedAt uint64 `json:"update_at"`
	DeletedAt uint64 `json:"deleted_at"`
	IsDelete  bool   `json:"is_delete"`
}

// migration 执行数据迁移
func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&User{})
}
