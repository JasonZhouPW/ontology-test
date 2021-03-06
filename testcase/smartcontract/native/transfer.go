/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package native

import (
	"github.com/ontio/ontology-test/testframework"
	"math/big"
	"time"
)

//TestOntTransfer test native transfer case
func TestOntTransfer(ctx *testframework.TestFrameworkContext) bool {
	admin, err := ctx.Wallet.GetDefaultAccount()
	if err != nil {
		ctx.LogError("Wallet.GetDefaultAccount error:%s", err)
		return false
	}

	wallet, err := ctx.Ont.CreateWallet("./wallet_test.dat", "wangbing")
	if err != nil {
		ctx.LogError("CreateWallet ./wallet_test.dat error:%s", err)
		return false
	}

	user, err := wallet.GetDefaultAccount()
	if err != nil {
		ctx.LogError("Wallet.CreateAccount error:%s", err)
		return false
	}

	adminBalanceBefore, err := ctx.Ont.Rpc.GetBalance(admin.Address)
	if err != nil {
		ctx.LogError("Rpc.GetBalance error:%s", err)
		return false
	}

	if adminBalanceBefore.Ont.Cmp(new(big.Int)) == 0 {
		ctx.LogWarn("TestOntTransfer failed. Balance of admin is 0")
		return false
	}
	ctx.LogInfo("adminBalanceBefore %d", adminBalanceBefore.Ont.Int64())

	userBalanceBefore, err := ctx.Ont.Rpc.GetBalance(user.Address)
	if err != nil {
		ctx.LogError("Rpc.GetBalance error:%s", err)
		return false
	}
	ctx.LogInfo("userBalanceBefore %d", userBalanceBefore.Ont.Int64())

	amount := new(big.Int).SetInt64(100)
	_, err = ctx.Ont.Rpc.Transfer("ONT", admin, user, amount)
	if err != nil {
		ctx.LogError("Rpc.Transfer error:%s", err)
		return false
	}

	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30 * time.Second,1)
	if err != nil {
		ctx.LogError("WaitForGenerateBlock error:%s", err)
		return false
	}

	adminBalanceAfter, err := ctx.Ont.Rpc.GetBalance(admin.Address)
	if err != nil {
		if err != nil {
			ctx.LogError("Rpc.GetBalance error:%s", err)
			return false
		}
	}
	ctx.LogInfo("adminBalanceAfter :%d", adminBalanceAfter.Ont.Int64())

	userBalanceAfter, err := ctx.Ont.Rpc.GetBalance(user.Address)
	if err != nil {
		ctx.LogError("Rpc.GetBalance error:%s", err)
		return false
	}
	ctx.LogInfo("userBalanceAfter :%d", userBalanceAfter.Ont.Int64())

	//Assert admin balance
	adminRes := new(big.Int).Sub(adminBalanceBefore.Ont, amount)
	if adminRes.Cmp(adminBalanceAfter.Ont) != 0 {
		ctx.LogError("TestOntTransfer failed. Admin balance after transfer %d != %d", adminBalanceAfter.Ont.Int64(), adminRes.Int64())
		return false
	}

	//Assert user balance
	userRes := new(big.Int).Add(userBalanceBefore.Ont, amount)
	if userRes.Cmp(userBalanceAfter.Ont) != 0 {
		ctx.LogError("TestOntTransfer failed. User balance after transfer %d != %d", userBalanceAfter.Ont.Int64(), userRes.Int64())
		return false
	}
	return true
}
