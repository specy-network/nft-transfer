package types

// IBC transfer events
const (
	EventTypeTimeout      = "timeout"
	EventTypePacket       = "non_fungible_token_packet"
	EventTypeTransfer     = "ibc_nft_transfer"
	EventTypeReceive      = "ibc_nft_receive"
	EventTypeChannelClose = "channel_closed"
	EventTypeClassTrace   = "class_trace"

	AttributeValueCategory = "ics721"

	AttributeKeySender          = "sender"
	AttributeKeyReceiver        = "receiver"
	AttributeKeyClassID         = "classID"
	AttributeKeyTokenIDs        = "tokenIDs"
	AttributeKeyTokenID         = "tokenID"
	AttributeKeyAckSuccess      = "success"
	AttributeKeyAckError        = "error"
	AttributeKeyTraceHash       = "trace_hash"
	AttributeKeyContractAddress = "contract_address"
)
