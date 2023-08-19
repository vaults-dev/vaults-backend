package graph

import "github.com/vaults-dev/vaults-backend/repositories"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	WalletRepository *repositories.WalletRepository
}
