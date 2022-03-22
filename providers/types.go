package providers

import (
	"math/big"
	"time"
)

type TenderlySimulate struct {
	NetworkId          string        `json:"network_id"`
	BlockNumber        uint64        `json:"block_number"`
	TransactionIndex   int           `json:"transaction_index"`
	From               string        `json:"from"`
	Input              string        `json:"input"`
	To                 string        `json:"to"`
	Gas                uint64        `json:"gas"`
	GasPrice           string        `json:"gas_price"`
	Value              string        `json:"value"`
	AccessList         []interface{} `json:"access_list"`
	GenerateAccessList bool          `json:"generate_access_list"`
	Save               bool          `json:"save"`
	Source             string        `json:"source"`
}

type TenderlySimulateRes struct {
	Transaction         TenderlySimulateTransaction `json:"transaction"`
	Simulation          TenderlySimulationRes       `json:"simulation"`
	Contracts           []string                    `json:"contracts"`
	GeneratedAccessList []interface{}               `json:"generated_access_list"`
}

type TenderlySimulateTransaction struct {
	Hash              string                  `json:"hash"`
	BlockHash         string                  `json:"block_hash"`
	BlockNumber       *big.Int                `json:"block_number"`
	From              string                  `json:"from"`
	Gas               uint64                  `json:"gas"`
	GasPrice          *big.Int                `json:"gas_price"`
	GasFeeCap         *big.Int                `json:"gas_fee_cap"`
	GasTipCap         *big.Int                `json:"gas_tip_cap"`
	CumulativeGasUsed *big.Int                `json:"cumulative_gas_used"`
	GasUsed           *big.Int                `json:"gas_used"`
	EffectiveGasPrice *big.Int                `json:"effective_gas_price"`
	Input             string                  `json:"input"`
	Nonce             uint64                  `json:"nonce"`
	To                string                  `json:"to"`
	Index             int                     `json:"index"`
	Value             string                  `json:"value"`
	AccessList        []string                `json:"access_list"`
	Status            bool                    `json:"status"`
	Addresses         []string                `json:"addresses"`
	ContractIds       []string                `json:"contract_ids"`
	NetworkId         string                  `json:"network_id"`
	Timestamp         time.Time               `json:"timestamp"`
	FunctionSelector  string                  `json:"function_selector"`
	TransactionInfo   TenderlyTransactionInfo `json:"transaction_info"`
	Method            string                  `json:"method"`
	DecodeInput       interface{}             `json:"decode_input"`
}

type TenderlyTransactionInfo struct {
	ContractId      string              `json:"contract_id"`
	BlockNumber     *big.Int            `json:"block_number"`
	TransactionId   string              `json:"transaction_id"`
	ContractAddress string              `json:"contract_address"`
	Method          interface{}         `json:"method"`
	Parameters      interface{}         `json:"parameters"`
	IntrinsicGas    *big.Int            `json:"intrinsic_gas"`
	RefundGas       *big.Int            `json:"refund_gas"`
	CallTrace       TenderlyCallTrace   `json:"call_trace"`
	StackTrace      interface{}         `json:"stack_trace"`
	Logs            []TenderlyLog       `json:"logs"`
	StateDiff       []TenderlyStateDiff `json:"state_diff"`
	RawStateDiff    interface{}         `json:"raw_state_diff"`
	ConsoleLogs     interface{}         `json:"console_logs"`
	CreatedAt       time.Time           `json:"created_at"`
}

type TenderlyCallTrace struct {
	Hash             string              `json:"hash"`
	ContractName     string              `json:"contract_name"`
	FunctionPc       int                 `json:"function_pc"`
	FunctionOp       string              `json:"function_op"`
	AbsolutePosition int                 `json:"absolute_position"`
	CallerPc         int                 `json:"caller_pc"`
	CallerOp         string              `json:"caller_op"`
	CallType         string              `json:"call_type"`
	From             string              `json:"from"`
	FromBalance      string              `json:"from_balance"`
	To               string              `json:"to"`
	ToBalance        string              `json:"to_balance"`
	Value            string              `json:"value"`
	BlockTimestamp   time.Time           `json:"block_timestamp"`
	Gas              int                 `json:"gas"`
	GasUsed          int                 `json:"gas_used"`
	IntrinsicGas     int                 `json:"intrinsic_gas"`
	Input            string              `json:"input"`
	StateDiff        []TenderlyStateDiff `json:"state_diff"`
	Logs             []TenderlyLog       `json:"logs"`
	Output           string              `json:"output"`
	DecodedOutput    interface{}         `json:"decoded_output"`
	NetworkID        string              `json:"network_id"`
	Calls            interface{}         `json:"calls"`
}

type TenderlyLog struct {
	Name      string         `json:"name"`
	Anonymous bool           `json:"anonymous"`
	Inputs    interface{}    `json:"inputs"`
	Raw       TenderlyLogRaw `json:"raw"`
}

type TenderlyLogRaw struct {
	Address string   `json:"address"`
	Topics  []string `json:"topics"`
	Data    string   `json:"data"`
}

type TenderlyStateDiff struct {
	Soltype  interface{}            `json:"soltype"`
	Original interface{}            `json:"original"`
	Dirty    interface{}            `json:"dirty"`
	Raw      []TenderlyStateDiffRaw `json:"raw"`
}

type TenderlyStateDiffRaw struct {
	Address  string `json:"address"`
	Key      string `json:"key"`
	Original string `json:"original"`
	Dirty    string `json:"dirty"`
}

type TenderlySimulationRes struct {
	Id               string        `json:"id"`
	ProjectId        string        `json:"project_id"`
	OwnerId          string        `json:"owner_id"`
	NetworkId        string        `json:"network_id"`
	BlockNumber      *big.Int      `json:"block_number"`
	TransactionIndex int           `json:"transaction_index"`
	From             string        `json:"from"`
	To               string        `json:"to"`
	Input            string        `json:"input"`
	Gas              uint64        `json:"gas"`
	GasPrice         string        `json:"gas_price"`
	Value            string        `json:"value"`
	Status           bool          `json:"status"`
	AccessList       []interface{} `json:"access_list"`
	QueueOrigin      string        `json:"queue_origin"`
	CreatedAt        time.Time     `json:"created_at"`
}
