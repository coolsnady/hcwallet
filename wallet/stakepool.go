// Copyright (c) 2016 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wallet

import (
	"errors"

	hxutil "github.com/coolsnady/hxd/hxutil"
	"github.com/coolsnady/hxwallet/wallet/udb"
	"github.com/coolsnady/hxwallet/walletdb"
)

// StakePoolUserInfo returns the stake pool user information for a user
// identified by their P2SH voting address.
func (w *Wallet) StakePoolUserInfo(userAddress hxutil.Address) (*udb.StakePoolUser, error) {
	switch userAddress.(type) {
	case *hxutil.AddressPubKeyHash: // ok
	case *hxutil.AddressScriptHash: // ok
	default:
		return nil, errors.New("stake pool user address must be P2PKH or P2SH")
	}

	var user *udb.StakePoolUser
	err := walletdb.View(w.db, func(tx walletdb.ReadTx) error {
		stakemgrNs := tx.ReadBucket(wstakemgrNamespaceKey)
		var err error
		user, err = w.StakeMgr.StakePoolUserInfo(stakemgrNs, userAddress)
		return err
	})
	return user, err
}
