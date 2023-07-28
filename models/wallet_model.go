package models

import "time"

type Wallet struct {
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index" json:"deleted_at"`
  Address      string     `gorm:"type:varchar(126)" json:"address"`
  OwnerAddress string     `gorm:"type:varchar(126)" json:"owner_address"`
  Salt         int        `gorm:"type:int" json:"salt"`
  NetworkID    string     `gorm:"type:varchar(125)" json:"network_id"`
}
