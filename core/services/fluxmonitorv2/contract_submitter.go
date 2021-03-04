package fluxmonitorv2

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/flux_aggregator_wrapper"
)

//go:generate mockery --name ContractSubmitter --output ./mocks/ --case=underscore

// ContractSubmitter defines an interface to submit an eth tx.
type ContractSubmitter interface {
	Submit(roundID *big.Int, submission *big.Int) error
}

// FluxAggregatorContractSubmitter submits the polled answer in an eth tx.
type FluxAggregatorContractSubmitter struct {
	flux_aggregator_wrapper.FluxAggregatorInterface
	keyStore KeyStoreInterface
}

// NewFluxAggregatorContractSubmitter constructs a new NewFluxAggregatorContractSubmitter
func NewFluxAggregatorContractSubmitter(contract flux_aggregator_wrapper.FluxAggregatorInterface, keyStore KeyStoreInterface) *FluxAggregatorContractSubmitter {
	return &FluxAggregatorContractSubmitter{
		FluxAggregatorInterface: contract,
		keyStore:                keyStore,
	}
}

// Submit submits the answer by writing a EthTx for the bulletprooftxmanager to
// pick up
func (c *FluxAggregatorContractSubmitter) Submit(roundID *big.Int, submission *big.Int) error {
	// fromAddress, err := c.orm.GetRoundRobinAddress()
	// roundIDArg := utils.EVMWordUint64(uint64(roundID))

	_, err := c.FluxAggregatorInterface.Submit(&bind.TransactOpts{
		// From: fromAddress,
	}, roundID, submission)

	// bind.NewTransactor()

	return err

	// return nil
}

// func (c *BoundContract) Transact(opts *TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	// Otherwise pack up the parameters and invoke the contract
// 	input, err := c.abi.Pack(method, params...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// todo(rjl493456442) check the method is payable or not,
// 	// reject invalid transaction at the first place
// 	return c.transact(opts, &c.address, input)
// }

// TransactOpts is the collection of authorization data required to create a
// valid Ethereum transaction.
// type TransactOpts struct {
// 	From   common.Address // Ethereum account to send the transaction from
// 	Nonce  *big.Int       // Nonce to use for the transaction execution (nil = use pending state)
// 	Signer SignerFn       // Method to use for signing the transaction (mandatory)

// 	Value    *big.Int // Funds to transfer along the transaction (nil = 0 = no funds)
// 	GasPrice *big.Int // Gas price to use for the transaction execution (nil = gas price oracle)
// 	GasLimit uint64   // Gas limit to set for the transaction execution (0 = estimate)

// 	Context context.Context // Network context to support cancellation and timeouts (nil = no timeout)
// }

// roundIDArg := utils.EVMWordUint64(uint64(roundID))

// _, err := fluxAggregatorABI.Pack("submit", roundIDArg, submission)
// if err != nil {
// 	return errors.Wrap(err, "abi.Pack failed")
// }

// Write an Eth TX to the DB for the bulletproof transaction manager to execute
