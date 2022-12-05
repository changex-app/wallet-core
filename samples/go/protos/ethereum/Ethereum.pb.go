// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: Ethereum.proto

package ethereum

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	common "tw/protos/common"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Transaction type
type TransactionMode int32

const (
	// Legacy transaction, pre-EIP2718/EIP1559; for fee gasPrice/gasLimit is used
	TransactionMode_Legacy TransactionMode = 0
	// Enveloped transaction EIP2718 (with type 0x2), fee is according to EIP1559 (base fee, inclusion fee, ...)
	TransactionMode_Enveloped TransactionMode = 1
)

// Enum value maps for TransactionMode.
var (
	TransactionMode_name = map[int32]string{
		0: "Legacy",
		1: "Enveloped",
	}
	TransactionMode_value = map[string]int32{
		"Legacy":    0,
		"Enveloped": 1,
	}
)

func (x TransactionMode) Enum() *TransactionMode {
	p := new(TransactionMode)
	*p = x
	return p
}

func (x TransactionMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionMode) Descriptor() protoreflect.EnumDescriptor {
	return file_Ethereum_proto_enumTypes[0].Descriptor()
}

func (TransactionMode) Type() protoreflect.EnumType {
	return &file_Ethereum_proto_enumTypes[0]
}

func (x TransactionMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionMode.Descriptor instead.
func (TransactionMode) EnumDescriptor() ([]byte, []int) {
	return file_Ethereum_proto_rawDescGZIP(), []int{0}
}

// Transaction (transfer, smart contract call, ...)
type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Payload transfer
	//
	// Types that are assignable to TransactionOneof:
	//	*Transaction_Transfer_
	//	*Transaction_Erc20Transfer
	//	*Transaction_Erc20Approve
	//	*Transaction_Erc721Transfer
	//	*Transaction_Erc1155Transfer
	//	*Transaction_ContractGeneric_
	TransactionOneof isTransaction_TransactionOneof `protobuf_oneof:"transaction_oneof"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ethereum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_Ethereum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_Ethereum_proto_rawDescGZIP(), []int{0}
}

func (m *Transaction) GetTransactionOneof() isTransaction_TransactionOneof {
	if m != nil {
		return m.TransactionOneof
	}
	return nil
}

func (x *Transaction) GetTransfer() *Transaction_Transfer {
	if x, ok := x.GetTransactionOneof().(*Transaction_Transfer_); ok {
		return x.Transfer
	}
	return nil
}

func (x *Transaction) GetErc20Transfer() *Transaction_ERC20Transfer {
	if x, ok := x.GetTransactionOneof().(*Transaction_Erc20Transfer); ok {
		return x.Erc20Transfer
	}
	return nil
}

func (x *Transaction) GetErc20Approve() *Transaction_ERC20Approve {
	if x, ok := x.GetTransactionOneof().(*Transaction_Erc20Approve); ok {
		return x.Erc20Approve
	}
	return nil
}

func (x *Transaction) GetErc721Transfer() *Transaction_ERC721Transfer {
	if x, ok := x.GetTransactionOneof().(*Transaction_Erc721Transfer); ok {
		return x.Erc721Transfer
	}
	return nil
}

func (x *Transaction) GetErc1155Transfer() *Transaction_ERC1155Transfer {
	if x, ok := x.GetTransactionOneof().(*Transaction_Erc1155Transfer); ok {
		return x.Erc1155Transfer
	}
	return nil
}

func (x *Transaction) GetContractGeneric() *Transaction_ContractGeneric {
	if x, ok := x.GetTransactionOneof().(*Transaction_ContractGeneric_); ok {
		return x.ContractGeneric
	}
	return nil
}

type isTransaction_TransactionOneof interface {
	isTransaction_TransactionOneof()
}

type Transaction_Transfer_ struct {
	Transfer *Transaction_Transfer `protobuf:"bytes,1,opt,name=transfer,proto3,oneof"`
}

type Transaction_Erc20Transfer struct {
	Erc20Transfer *Transaction_ERC20Transfer `protobuf:"bytes,2,opt,name=erc20_transfer,json=erc20Transfer,proto3,oneof"`
}

type Transaction_Erc20Approve struct {
	Erc20Approve *Transaction_ERC20Approve `protobuf:"bytes,3,opt,name=erc20_approve,json=erc20Approve,proto3,oneof"`
}

type Transaction_Erc721Transfer struct {
	Erc721Transfer *Transaction_ERC721Transfer `protobuf:"bytes,4,opt,name=erc721_transfer,json=erc721Transfer,proto3,oneof"`
}

type Transaction_Erc1155Transfer struct {
	Erc1155Transfer *Transaction_ERC1155Transfer `protobuf:"bytes,5,opt,name=erc1155_transfer,json=erc1155Transfer,proto3,oneof"`
}

type Transaction_ContractGeneric_ struct {
	ContractGeneric *Transaction_ContractGeneric `protobuf:"bytes,6,opt,name=contract_generic,json=contractGeneric,proto3,oneof"`
}

func (*Transaction_Transfer_) isTransaction_TransactionOneof() {}

func (*Transaction_Erc20Transfer) isTransaction_TransactionOneof() {}

func (*Transaction_Erc20Approve) isTransaction_TransactionOneof() {}

func (*Transaction_Erc721Transfer) isTransaction_TransactionOneof() {}

func (*Transaction_Erc1155Transfer) isTransaction_TransactionOneof() {}

func (*Transaction_ContractGeneric_) isTransaction_TransactionOneof() {}

// Input data necessary to create a signed transaction.
// Legacy and EIP2718/EIP1559 transactions supported, see TransactionMode.
type SigningInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chain identifier (uint256, serialized little endian)
	ChainId []byte `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// Nonce (uint256, serialized little endian)
	Nonce []byte `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Transaction version selector: Legacy or enveloped, has impact on fee structure.
	// Default is Legacy (value 0)
	TxMode TransactionMode `protobuf:"varint,3,opt,name=tx_mode,json=txMode,proto3,enum=TW.Ethereum.Proto.TransactionMode" json:"tx_mode,omitempty"`
	// Gas price (uint256, serialized little endian)
	// Relevant for legacy transactions only (disregarded for enveloped/EIP1559)
	GasPrice []byte `protobuf:"bytes,4,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	// Gas limit (uint256, serialized little endian)
	GasLimit []byte `protobuf:"bytes,5,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	// Maximum optional inclusion fee (aka tip) (uint256, serialized little endian)
	// Relevant for enveloped/EIP1559 transactions only, tx_mode=Enveloped, (disregarded for legacy)
	MaxInclusionFeePerGas []byte `protobuf:"bytes,6,opt,name=max_inclusion_fee_per_gas,json=maxInclusionFeePerGas,proto3" json:"max_inclusion_fee_per_gas,omitempty"`
	// Maximum fee (uint256, serialized little endian)
	// Relevant for enveloped/EIP1559 transactions only, tx_mode=Enveloped, (disregarded for legacy)
	MaxFeePerGas []byte `protobuf:"bytes,7,opt,name=max_fee_per_gas,json=maxFeePerGas,proto3" json:"max_fee_per_gas,omitempty"`
	// Recipient's address.
	ToAddress string `protobuf:"bytes,8,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	// The secret private key used for signing (32 bytes).
	PrivateKey []byte `protobuf:"bytes,9,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// The payload transaction
	Transaction *Transaction `protobuf:"bytes,10,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *SigningInput) Reset() {
	*x = SigningInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ethereum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningInput) ProtoMessage() {}

func (x *SigningInput) ProtoReflect() protoreflect.Message {
	mi := &file_Ethereum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigningInput.ProtoReflect.Descriptor instead.
func (*SigningInput) Descriptor() ([]byte, []int) {
	return file_Ethereum_proto_rawDescGZIP(), []int{1}
}

func (x *SigningInput) GetChainId() []byte {
	if x != nil {
		return x.ChainId
	}
	return nil
}

func (x *SigningInput) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *SigningInput) GetTxMode() TransactionMode {
	if x != nil {
		return x.TxMode
	}
	return TransactionMode_Legacy
}

func (x *SigningInput) GetGasPrice() []byte {
	if x != nil {
		return x.GasPrice
	}
	return nil
}

func (x *SigningInput) GetGasLimit() []byte {
	if x != nil {
		return x.GasLimit
	}
	return nil
}

func (x *SigningInput) GetMaxInclusionFeePerGas() []byte {
	if x != nil {
		return x.MaxInclusionFeePerGas
	}
	return nil
}

func (x *SigningInput) GetMaxFeePerGas() []byte {
	if x != nil {
		return x.MaxFeePerGas
	}
	return nil
}

func (x *SigningInput) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *SigningInput) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *SigningInput) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

// Result containing the signed and encoded transaction.
type SigningOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Signed and encoded transaction bytes.
	Encoded []byte `protobuf:"bytes,1,opt,name=encoded,proto3" json:"encoded,omitempty"`
	// The V, R, S components of the resulting signature, (each uint256, serialized little endian)
	V []byte `protobuf:"bytes,2,opt,name=v,proto3" json:"v,omitempty"`
	R []byte `protobuf:"bytes,3,opt,name=r,proto3" json:"r,omitempty"`
	S []byte `protobuf:"bytes,4,opt,name=s,proto3" json:"s,omitempty"`
	// The payload part, supplied in the input or assembled from input parameters
	Data []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	// error code, 0 is ok, other codes will be treated as errors
	Error common.SigningError `protobuf:"varint,6,opt,name=error,proto3,enum=TW.Common.Proto.SigningError" json:"error,omitempty"`
	// error code description
	ErrorMessage string `protobuf:"bytes,7,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *SigningOutput) Reset() {
	*x = SigningOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ethereum_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningOutput) ProtoMessage() {}

func (x *SigningOutput) ProtoReflect() protoreflect.Message {
	mi := &file_Ethereum_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigningOutput.ProtoReflect.Descriptor instead.
func (*SigningOutput) Descriptor() ([]byte, []int) {
	return file_Ethereum_proto_rawDescGZIP(), []int{2}
}

func (x *SigningOutput) GetEncoded() []byte {
	if x != nil {
		return x.Encoded
	}
	return nil
}

func (x *SigningOutput) GetV() []byte {
	if x != nil {
		return x.V
	}
	return nil
}

func (x *SigningOutput) GetR() []byte {
	if x != nil {
		return x.R
	}
	return nil
}

func (x *SigningOutput) GetS() []byte {
	if x != nil {
		return x.S
	}
	return nil
}

func (x *SigningOutput) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SigningOutput) GetError() common.SigningError {
	if x != nil {
		return x.Error
	}
	return common.SigningError(0)
}

func (x *SigningOutput) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

// Native coin transfer transaction
type Transaction_Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Amount to send in wei (uint256, serialized little endian)
	Amount []byte `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// Optional payload data
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Transaction_Transfer) Reset() {
	*x = Transaction_Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ethereum_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction_Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction_Transfer) ProtoMessage() {}

func (x *Transaction_Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_Ethereum_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction_Transfer.ProtoReflect.Descriptor instead.
func (*Transaction_Transfer) Descriptor() ([]byte, []int) {
	return file_Ethereum_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Transaction_Transfer) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Transaction_Transfer) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// ERC20 token transfer transaction
type Transaction_ERC20Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// destination address (string)
	To string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	// Amount to send (uint256, serialized little endian)
	Amount []byte `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Transaction_ERC20Transfer) Reset() {
	*x = Transaction_ERC20Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ethereum_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction_ERC20Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction_ERC20Transfer) ProtoMessage() {}

func (x *Transaction_ERC20Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_Ethereum_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction_ERC20Transfer.ProtoReflect.Descriptor instead.
func (*Transaction_ERC20Transfer) Descriptor() ([]byte, []int) {
	return file_Ethereum_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Transaction_ERC20Transfer) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Transaction_ERC20Transfer) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

// ERC20 approve transaction
type Transaction_ERC20Approve struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Target of the approval
	Spender string `protobuf:"bytes,1,opt,name=spender,proto3" json:"spender,omitempty"`
	// Amount to send (uint256, serialized little endian)
	Amount []byte `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Transaction_ERC20Approve) Reset() {
	*x = Transaction_ERC20Approve{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ethereum_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction_ERC20Approve) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction_ERC20Approve) ProtoMessage() {}

func (x *Transaction_ERC20Approve) ProtoReflect() protoreflect.Message {
	mi := &file_Ethereum_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction_ERC20Approve.ProtoReflect.Descriptor instead.
func (*Transaction_ERC20Approve) Descriptor() ([]byte, []int) {
	return file_Ethereum_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Transaction_ERC20Approve) GetSpender() string {
	if x != nil {
		return x.Spender
	}
	return ""
}

func (x *Transaction_ERC20Approve) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

// ERC721 NFT transfer transaction
type Transaction_ERC721Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source address
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// Destination address
	To string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	// ID of the token (uint256, serialized little endian)
	TokenId []byte `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
}

func (x *Transaction_ERC721Transfer) Reset() {
	*x = Transaction_ERC721Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ethereum_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction_ERC721Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction_ERC721Transfer) ProtoMessage() {}

func (x *Transaction_ERC721Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_Ethereum_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction_ERC721Transfer.ProtoReflect.Descriptor instead.
func (*Transaction_ERC721Transfer) Descriptor() ([]byte, []int) {
	return file_Ethereum_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Transaction_ERC721Transfer) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Transaction_ERC721Transfer) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Transaction_ERC721Transfer) GetTokenId() []byte {
	if x != nil {
		return x.TokenId
	}
	return nil
}

// ERC1155 NFT transfer transaction
type Transaction_ERC1155Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source address
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// Destination address
	To string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	// ID of the token (uint256, serialized little endian)
	TokenId []byte `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	// The amount of tokens being transferred (uint256, serialized little endian)
	Value []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Data  []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Transaction_ERC1155Transfer) Reset() {
	*x = Transaction_ERC1155Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ethereum_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction_ERC1155Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction_ERC1155Transfer) ProtoMessage() {}

func (x *Transaction_ERC1155Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_Ethereum_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction_ERC1155Transfer.ProtoReflect.Descriptor instead.
func (*Transaction_ERC1155Transfer) Descriptor() ([]byte, []int) {
	return file_Ethereum_proto_rawDescGZIP(), []int{0, 4}
}

func (x *Transaction_ERC1155Transfer) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Transaction_ERC1155Transfer) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Transaction_ERC1155Transfer) GetTokenId() []byte {
	if x != nil {
		return x.TokenId
	}
	return nil
}

func (x *Transaction_ERC1155Transfer) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Transaction_ERC1155Transfer) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Generic smart contract transaction
type Transaction_ContractGeneric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Amount to send in wei (uint256, serialized little endian)
	Amount []byte `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// Contract call payload data
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Transaction_ContractGeneric) Reset() {
	*x = Transaction_ContractGeneric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ethereum_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction_ContractGeneric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction_ContractGeneric) ProtoMessage() {}

func (x *Transaction_ContractGeneric) ProtoReflect() protoreflect.Message {
	mi := &file_Ethereum_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction_ContractGeneric.ProtoReflect.Descriptor instead.
func (*Transaction_ContractGeneric) Descriptor() ([]byte, []int) {
	return file_Ethereum_proto_rawDescGZIP(), []int{0, 5}
}

func (x *Transaction_ContractGeneric) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Transaction_ContractGeneric) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_Ethereum_proto protoreflect.FileDescriptor

var file_Ethereum_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x54, 0x57, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe7, 0x07, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x45, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x54, 0x57, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x48, 0x00, 0x52, 0x08,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x55, 0x0a, 0x0e, 0x65, 0x72, 0x63, 0x32,
	0x30, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x54, 0x57, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x45, 0x52, 0x43, 0x32, 0x30, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x48, 0x00,
	0x52, 0x0d, 0x65, 0x72, 0x63, 0x32, 0x30, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12,
	0x52, 0x0a, 0x0d, 0x65, 0x72, 0x63, 0x32, 0x30, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x54, 0x57, 0x2e, 0x45, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x52, 0x43, 0x32, 0x30, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x65, 0x72, 0x63, 0x32, 0x30, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x12, 0x58, 0x0a, 0x0f, 0x65, 0x72, 0x63, 0x37, 0x32, 0x31, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x54,
	0x57, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x52, 0x43,
	0x37, 0x32, 0x31, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0e, 0x65,
	0x72, 0x63, 0x37, 0x32, 0x31, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x5b, 0x0a,
	0x10, 0x65, 0x72, 0x63, 0x31, 0x31, 0x35, 0x35, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x54, 0x57, 0x2e, 0x45, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x52, 0x43, 0x31, 0x31, 0x35, 0x35, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0f, 0x65, 0x72, 0x63, 0x31, 0x31,
	0x35, 0x35, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x5b, 0x0a, 0x10, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x54, 0x57, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x69, 0x63, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x1a, 0x36, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x37, 0x0a, 0x0d, 0x45, 0x52, 0x43, 0x32, 0x30, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x40, 0x0a, 0x0c, 0x45, 0x52, 0x43, 0x32,
	0x30, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x4f, 0x0a, 0x0e, 0x45, 0x52,
	0x43, 0x37, 0x32, 0x31, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x1a, 0x7a, 0x0a, 0x0f, 0x45,
	0x52, 0x43, 0x31, 0x31, 0x35, 0x35, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3d, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x13, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22, 0x99, 0x03, 0x0a, 0x0c,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x3b, 0x0a,
	0x07, 0x74, 0x78, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22,
	0x2e, 0x54, 0x57, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f,
	0x64, 0x65, 0x52, 0x06, 0x74, 0x78, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61,
	0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x67,
	0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x38, 0x0a, 0x19, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x67, 0x61,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x15, 0x6d, 0x61, 0x78, 0x49, 0x6e, 0x63, 0x6c,
	0x75, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x50, 0x65, 0x72, 0x47, 0x61, 0x73, 0x12, 0x25,
	0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x67, 0x61,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x46, 0x65, 0x65, 0x50,
	0x65, 0x72, 0x47, 0x61, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x40, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x54, 0x57, 0x2e,
	0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc1, 0x01, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e,
	0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01,
	0x76, 0x12, 0x0c, 0x0a, 0x01, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x72, 0x12,
	0x0c, 0x0a, 0x01, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x33, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x54, 0x57, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x2c, 0x0a, 0x0f, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0a,
	0x0a, 0x06, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x64, 0x10, 0x01, 0x42, 0x17, 0x0a, 0x15, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6a, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Ethereum_proto_rawDescOnce sync.Once
	file_Ethereum_proto_rawDescData = file_Ethereum_proto_rawDesc
)

func file_Ethereum_proto_rawDescGZIP() []byte {
	file_Ethereum_proto_rawDescOnce.Do(func() {
		file_Ethereum_proto_rawDescData = protoimpl.X.CompressGZIP(file_Ethereum_proto_rawDescData)
	})
	return file_Ethereum_proto_rawDescData
}

var file_Ethereum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Ethereum_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_Ethereum_proto_goTypes = []interface{}{
	(TransactionMode)(0),                // 0: TW.Ethereum.Proto.TransactionMode
	(*Transaction)(nil),                 // 1: TW.Ethereum.Proto.Transaction
	(*SigningInput)(nil),                // 2: TW.Ethereum.Proto.SigningInput
	(*SigningOutput)(nil),               // 3: TW.Ethereum.Proto.SigningOutput
	(*Transaction_Transfer)(nil),        // 4: TW.Ethereum.Proto.Transaction.Transfer
	(*Transaction_ERC20Transfer)(nil),   // 5: TW.Ethereum.Proto.Transaction.ERC20Transfer
	(*Transaction_ERC20Approve)(nil),    // 6: TW.Ethereum.Proto.Transaction.ERC20Approve
	(*Transaction_ERC721Transfer)(nil),  // 7: TW.Ethereum.Proto.Transaction.ERC721Transfer
	(*Transaction_ERC1155Transfer)(nil), // 8: TW.Ethereum.Proto.Transaction.ERC1155Transfer
	(*Transaction_ContractGeneric)(nil), // 9: TW.Ethereum.Proto.Transaction.ContractGeneric
	(common.SigningError)(0),            // 10: TW.Common.Proto.SigningError
}
var file_Ethereum_proto_depIdxs = []int32{
	4,  // 0: TW.Ethereum.Proto.Transaction.transfer:type_name -> TW.Ethereum.Proto.Transaction.Transfer
	5,  // 1: TW.Ethereum.Proto.Transaction.erc20_transfer:type_name -> TW.Ethereum.Proto.Transaction.ERC20Transfer
	6,  // 2: TW.Ethereum.Proto.Transaction.erc20_approve:type_name -> TW.Ethereum.Proto.Transaction.ERC20Approve
	7,  // 3: TW.Ethereum.Proto.Transaction.erc721_transfer:type_name -> TW.Ethereum.Proto.Transaction.ERC721Transfer
	8,  // 4: TW.Ethereum.Proto.Transaction.erc1155_transfer:type_name -> TW.Ethereum.Proto.Transaction.ERC1155Transfer
	9,  // 5: TW.Ethereum.Proto.Transaction.contract_generic:type_name -> TW.Ethereum.Proto.Transaction.ContractGeneric
	0,  // 6: TW.Ethereum.Proto.SigningInput.tx_mode:type_name -> TW.Ethereum.Proto.TransactionMode
	1,  // 7: TW.Ethereum.Proto.SigningInput.transaction:type_name -> TW.Ethereum.Proto.Transaction
	10, // 8: TW.Ethereum.Proto.SigningOutput.error:type_name -> TW.Common.Proto.SigningError
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_Ethereum_proto_init() }
func file_Ethereum_proto_init() {
	if File_Ethereum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Ethereum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Ethereum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigningInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Ethereum_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigningOutput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Ethereum_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction_Transfer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Ethereum_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction_ERC20Transfer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Ethereum_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction_ERC20Approve); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Ethereum_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction_ERC721Transfer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Ethereum_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction_ERC1155Transfer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Ethereum_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction_ContractGeneric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_Ethereum_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Transaction_Transfer_)(nil),
		(*Transaction_Erc20Transfer)(nil),
		(*Transaction_Erc20Approve)(nil),
		(*Transaction_Erc721Transfer)(nil),
		(*Transaction_Erc1155Transfer)(nil),
		(*Transaction_ContractGeneric_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Ethereum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Ethereum_proto_goTypes,
		DependencyIndexes: file_Ethereum_proto_depIdxs,
		EnumInfos:         file_Ethereum_proto_enumTypes,
		MessageInfos:      file_Ethereum_proto_msgTypes,
	}.Build()
	File_Ethereum_proto = out.File
	file_Ethereum_proto_rawDesc = nil
	file_Ethereum_proto_goTypes = nil
	file_Ethereum_proto_depIdxs = nil
}
