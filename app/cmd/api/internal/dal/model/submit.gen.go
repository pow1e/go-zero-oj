// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSubmit = "submit"

// Submit mapped from table <submit>
type Submit struct {
	ID              int32          `gorm:"column:id;primaryKey" json:"id"`
	Identity        string         `gorm:"column:identity" json:"identity"`
	ProblemIdentity string         `gorm:"column:problem_identity;comment:问题的唯一标识" json:"problem_identity"`                      // 问题的唯一标识
	UserIdentity    string         `gorm:"column:user_identity;comment:用户的唯一标识" json:"user_identity"`                            // 用户的唯一标识
	Path            string         `gorm:"column:path;comment:代码路径" json:"path"`                                                 // 代码路径
	Status          bool           `gorm:"column:status;not null;comment:0表示待判断，1表示答案正确，2表示答案错误，3表示运行超时，4表示运行超内存" json:"status"` // 0表示待判断，1表示答案正确，2表示答案错误，3表示运行超时，4表示运行超内存
	CreatedAt       time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"`                                     // 创建时间
	UpdatedAt       time.Time      `gorm:"column:updated_at;comment:修改时间" json:"updated_at"`                                     // 修改时间
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间(软删除)" json:"deleted_at"`                                // 删除时间(软删除)
}

// TableName Submit's table name
func (*Submit) TableName() string {
	return TableNameSubmit
}
