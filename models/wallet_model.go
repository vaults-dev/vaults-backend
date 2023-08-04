package models

type Wallet struct {
  TimeStamps
  Address      string     `gorm:"type:varchar(126);uniqueIndex:idx_address_network_id" json:"address"`
  OwnerAddress string     `gorm:"type:varchar(126)" json:"owner_address"`
  Salt         int        `gorm:"type:int" json:"salt"`
  NetworkID    string     `gorm:"type:varchar(125);uniqueIndex:idx_address_network_id" json:"network_id"`
}
