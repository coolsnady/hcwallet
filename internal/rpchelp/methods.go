// Copyright (c) 2015 The btcsuite developers
// Copyright (c) 2015-2017 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

//+build !generate

package rpchelp

import "github.com/coolsnady/hxd/hxjson"

// Common return types.
var (
	returnsBool        = []interface{}{(*bool)(nil)}
	returnsNumber      = []interface{}{(*float64)(nil)}
	returnsString      = []interface{}{(*string)(nil)}
	returnsStringArray = []interface{}{(*[]string)(nil)}
	returnsLTRArray    = []interface{}{(*[]hxjson.ListTransactionsResult)(nil)}
)

// Methods contains all methods and result types that help is generated for,
// for every locale.
var Methods = []struct {
	Method      string
	ResultTypes []interface{}
}{
	{"accountaddressindex", []interface{}{(*int)(nil)}},
	{"accountsyncaddressindex", nil},
	{"addmultisigaddress", returnsString},
	{"consolidate", returnsString},
	{"createmultisig", []interface{}{(*hxjson.CreateMultiSigResult)(nil)}},
	{"dumpprivkey", returnsString},
	{"getaccount", returnsString},
	{"getaccountaddress", returnsString},
	{"getaddressesbyaccount", returnsStringArray},
	{"getbalance", append(returnsNumber, returnsNumber[0])},
	{"getbestblockhash", returnsString},
	{"getblockcount", returnsNumber},
	{"getinfo", []interface{}{(*hxjson.InfoWalletResult)(nil)}},
	{"getmasterpubkey", []interface{}{(*string)(nil)}},
	{"getmultisigoutinfo", []interface{}{(*hxjson.GetMultisigOutInfoResult)(nil)}},
	{"getnewaddress", returnsString},
	{"getrawchangeaddress", returnsString},
	{"getreceivedbyaccount", returnsNumber},
	{"getreceivedbyaddress", returnsNumber},
	{"gettickets", []interface{}{(*hxjson.GetTicketsResult)(nil)}},
	{"gettransaction", []interface{}{(*hxjson.GetTransactionResult)(nil)}},
	{"getvotechoices", []interface{}{(*hxjson.GetVoteChoicesResult)(nil)}},
	{"help", append(returnsString, returnsString[0])},
	{"importprivkey", nil},
	{"importscript", nil},
	{"keypoolrefill", nil},
	{"listaccounts", []interface{}{(*map[string]float64)(nil)}},
	{"listlockunspent", []interface{}{(*[]hxjson.TransactionInput)(nil)}},
	{"listreceivedbyaccount", []interface{}{(*[]hxjson.ListReceivedByAccountResult)(nil)}},
	{"listreceivedbyaddress", []interface{}{(*[]hxjson.ListReceivedByAddressResult)(nil)}},
	{"listsinceblock", []interface{}{(*hxjson.ListSinceBlockResult)(nil)}},
	{"listtransactions", returnsLTRArray},
	{"listunspent", []interface{}{(*hxjson.ListUnspentResult)(nil)}},
	{"lockunspent", returnsBool},
	{"redeemmultisigout", []interface{}{(*hxjson.RedeemMultiSigOutResult)(nil)}},
	{"redeemmultisigouts", []interface{}{(*hxjson.RedeemMultiSigOutResult)(nil)}},
	{"rescanwallet", nil},
	{"revoketickets", nil},
	{"sendfrom", returnsString},
	{"sendmany", returnsString},
	{"sendtoaddress", returnsString},
	{"sendtomultisig", returnsString},
	{"settxfee", returnsBool},
	{"setvotechoice", nil},
	{"signmessage", returnsString},
	{"signrawtransaction", []interface{}{(*hxjson.SignRawTransactionResult)(nil)}},
	{"signrawtransactions", []interface{}{(*hxjson.SignRawTransactionsResult)(nil)}},
	{"validateaddress", []interface{}{(*hxjson.ValidateAddressWalletResult)(nil)}},
	{"verifymessage", returnsBool},
	{"version", []interface{}{(*map[string]hxjson.VersionResult)(nil)}},
	{"walletlock", nil},
	{"walletpassphrase", nil},
	{"walletpassphrasechange", nil},
	{"createnewaccount", nil},
	{"exportwatchingwallet", returnsString},
	{"getbestblock", []interface{}{(*hxjson.GetBestBlockResult)(nil)}},
	{"getunconfirmedbalance", returnsNumber},
	{"listaddresstransactions", returnsLTRArray},
	{"listalltransactions", returnsLTRArray},
	{"renameaccount", nil},
	{"walletislocked", returnsBool},
	{"walletinfo", []interface{}{(*hxjson.WalletInfoResult)(nil)}},

	// TODO Alphabetize
	{"purchaseticket", returnsString},
	{"sendtossrtx", returnsString},
	{"sendtosstx", returnsString},
	{"sendtossgen", returnsString},
	{"generatevote", []interface{}{(*hxjson.GenerateVoteResult)(nil)}},
	{"getstakeinfo", []interface{}{(*hxjson.GetStakeInfoResult)(nil)}},
	{"getticketfee", returnsNumber},
	{"setticketfee", returnsBool},
	{"getwalletfee", returnsNumber},
	{"addticket", nil},
	{"listscripts", []interface{}{(*hxjson.ListScriptsResult)(nil)}},
	{"stakepooluserinfo", []interface{}{(*hxjson.StakePoolUserInfoResult)(nil)}},
	{"ticketsforaddress", returnsBool},
}

// HelpDescs contains the locale-specific help strings along with the locale.
var HelpDescs = []struct {
	Locale   string // Actual locale, e.g. en_US
	GoLocale string // Locale used in Go names, e.g. EnUS
	Descs    map[string]string
}{
	{"en_US", "EnUS", helpDescsEnUS}, // helpdescs_en_US.go
}
