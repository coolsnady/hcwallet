// Copyright (c) 2016 The coolsnady developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package ticketbuyer

import (
	"bytes"
	"time"

	"github.com/coolsnady/hxd/dcrjson"
	"github.com/coolsnady/hxd/dcrutil"
	"github.com/coolsnady/hxwallet/wallet"
)

// ownTicketsInMempool finds all the tickets owned by the user in the
// daemon mempool. It searches for the ticket address if it is specified,
// and otherwise uses getstakeinfo to determine this number.
func (t *TicketPurchaser) ownTicketsInMempool() (int, error) {
	tickets := 0

	// Voting address is specified and may not belong to our own
	// wallet. Search the mempool directly for the number of tickets.
	if t.votingAddress != nil {
		tiHashes, err := t.dcrdChainSvr.GetRawMempool(dcrjson.GRMTickets)
		if err != nil {
			return 0, err
		}

		// Fetch each ticket and check the address it pays out to.
		for i := range tiHashes {
			raw, err := t.dcrdChainSvr.GetRawTransactionVerbose(tiHashes[i])
			if err != nil {
				return 0, err
			}

			// Tickets can only pay to a single address. Assume that
			// the address is on the right network.
			addrStr := raw.Vout[0].ScriptPubKey.Addresses[0]
			addr, err := dcrutil.DecodeAddress(addrStr)
			if err != nil {
				return 0, err
			}
			if bytes.Equal(addr.ScriptAddress(),
				t.votingAddress.ScriptAddress()) {
				tickets++
			}
		}

		return tickets, nil
	}

	// The ticket address is generated by the wallet
	// and is assumed to be owned by the wallet. Use
	// getstakeinfo to figure out the number of tickets.
	//
	// It can take a little while for the wallet to sync,
	// so loop this and recheck to see if we've got the
	// next block attached yet.
	var curStakeInfo *wallet.StakeInfoData
	var err error
	for i := 0; i < stakeInfoReqTries; i++ {
		curStakeInfo, err = t.wallet.StakeInfo(t.dcrdChainSvr)
		if err != nil {
			log.Tracef("Failed to fetch stake information "+
				"on attempt %v: %v", i, err.Error())
			time.Sleep(stakeInfoReqTryDelay)
			continue
		}
		if err == nil {
			break
		}
	}
	if err != nil {
		return 0, err
	}

	return int(curStakeInfo.OwnMempoolTix), nil
}

// allTicketsInMempool fetches the number of tickets currently in the memory
// pool.
func (t *TicketPurchaser) allTicketsInMempool() (int, error) {
	tfi, err := t.dcrdChainSvr.TicketFeeInfo(&zeroUint32, &zeroUint32)
	if err != nil {
		return 0, err
	}

	return int(tfi.FeeInfoMempool.Number), nil
}
