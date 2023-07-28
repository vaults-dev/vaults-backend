package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vaults-dev/vaults-backend/graph/model"
	"github.com/vaults-dev/vaults-backend/web3/smart_contracts"
	"gorm.io/gorm"
)

const VAULTS_WALLET_FACTORY_CONTRACT_ADDRESS = "0xD86aF383e8871554359bfBDAE307CF55A9088306"

// TODO: This is a test private key, we need to generate a new one for production
const EOA_PRIVATE_KEY = "84659b93fdde90cf67c4411684df0d53b19c0e0904b2f36e14ef564ed068d3bb"

type WalletRepository struct {
  db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
  return &WalletRepository{db}
}

func (r *WalletRepository) CreateWallet(input model.CreateWalletPayload) (*model.Wallet, error) {
  // TODO: RPC url should be in database
  // the RPC url should be part of Networks table 
  // The network is fetched using input.NetworkID

  rpcUrl := "https://eth-sepolia.blastapi.io/4c042713-b83d-4a24-a6bf-41e20b08b216"
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	factoryContractAddress := common.HexToAddress(VAULTS_WALLET_FACTORY_CONTRACT_ADDRESS)
	walletFactory, err := smart_contracts.NewVaultsWalletFactory(factoryContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(EOA_PRIVATE_KEY)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

  ownerAddress := common.HexToAddress(input.OwnerAddress)
  salt := big.NewInt(int64(input.Salt))

	// Create Account
	tx, err := walletFactory.CreateAccount(auth, ownerAddress, salt)
	if err != nil {
    return nil, err
	}
	// Get the account address created
	createdAccountAddress, err := walletFactory.GetAddress(nil, ownerAddress, salt)
	if err != nil {
		log.Fatal(err)
	}
  fmt.Println("createdAccountAddress.Hex()", createdAccountAddress.Hex())

	maxRetries := 5
	retries := 0
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
    txReceipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
    if err != nil {
			if retries < maxRetries {
				retries++
				continue
			}
    }
    if txReceipt == nil {
      continue
    }
		if txReceipt.Status == 1 {
      // Transaction was successful
      err = r.db.Create(&model.Wallet{
        Address:     createdAccountAddress.Hex(),
        OwnerAddress: input.OwnerAddress,
        Salt:        input.Salt,
        NetworkID:   input.NetworkID,
      }).Error
      if err != nil {
        return nil, err
      }

      return &model.Wallet{
        Address:     createdAccountAddress.Hex(),
        OwnerAddress: input.OwnerAddress,
        Salt:        input.Salt,
        NetworkID:   input.NetworkID,
      }, nil
    }
    if txReceipt.Status == 0 {
      // Transaction failed
      return nil, errors.New("create wallet transaction failed")
    }
  }

  return nil, errors.New("create wallet transaction failed")
}
