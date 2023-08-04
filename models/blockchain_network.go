package models

type BlockchainNetwork struct {
	TimeStamps
	ID               string                          `json:"id" gorm:"primaryKey;type:varchar(125)"`
	ChainID          int                             `json:"chain_id" gorm:"uniqueIndex:idx_chain_id_network_type"`
	Name             string                          `json:"name" gorm:"varchar(125)"`
	IsTestNet        bool                            `json:"is_test_net"`
	LogoUrl          string                          `json:"logo_url" gorm:"varchar(255)"`
	RpcUrl           string                          `json:"rpc_url" gorm:"varchar(255)"`
	BlockExplorerUrl string                          `json:"block_explorer_url" gorm:"varchar(255)"`
}
