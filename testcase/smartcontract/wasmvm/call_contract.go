package wasmvm

import (
	"fmt"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/account"
	"io/ioutil"
	"errors"
	"github.com/ontio/ontology/smartcontract/types"
	"time"
	"github.com/ontio/ontology/smartcontract/service/wasmvm"
)

func TestCallWasmJsonContract(ctx *testframework.TestFrameworkContext) bool{
	wasmWallet := "wallet.dat"
	wasmWalletPwd := "123456"
	wallet, err := ctx.Ont.OpenWallet(wasmWallet, wasmWalletPwd)
	if err != nil {
		ctx.LogError("OpenWallet:%s error:%s", wasmWallet, err)
		return false
	}

	admin, err := wallet.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestCallWasmJsonContract wallet.GetDefaultAccount error:%s", err)
		return false
	}

	txHash, err := deployCallWasmJsonContract(ctx, admin)
	if err != nil {
		ctx.LogError("TestCallWasmJsonContract deploy error:%s", err)
		return false
	}

	ctx.LogInfo("TestCallWasmJsonContract deploy TxHash:%x", txHash)

	address,err := GetWasmContractAddress(filePath+"/callContract.wasm")
	if err != nil {
		ctx.LogError("TestCallWasmJsonContract GetWasmContractAddress error:%s", err)
		return false
	}
	txHash,err = invokeCallContractAddValue(ctx,admin,address)
	if err != nil {
		ctx.LogError("TestCallWasmJsonContract invokeCallContractAddValue error:%s", err)
		return false
	}

	notifies, err := ctx.Ont.Rpc.GetSmartContractEvent(txHash)
	if err != nil {
		ctx.LogError("TestCallWasmJsonContract invokeCallContractAddValue GetSmartContractEvent error:%s", err)
		return false
	}

	if len(notifies) < 1{
		ctx.LogError("TestWasmJsonContract invokeCallContractAddValue return notifies count error!")
		return false
	}
	ctx.LogInfo("==========TestWasmJsonContract invokeCallContractAddValue ============")
	for i ,n := range notifies{
		ctx.LogInfo(fmt.Sprintf("notify %d is %v",i, n))
	}


	txHash,err = invokeCallContractGetValue(ctx,admin,address)
	if err != nil {
		ctx.LogError("TestCallWasmJsonContract invokeCallContractGetValue error:%s", err)
		return false
	}

	notifies, err = ctx.Ont.Rpc.GetSmartContractEvent(txHash)
	if err != nil {
		ctx.LogError("TestCallWasmJsonContract invokeCallContractGetValue GetSmartContractEvent error:%s", err)
		return false
	}

	if len(notifies) < 1{
		ctx.LogError("TestWasmJsonContract invokeCallContractGetValue return notifies count error!")
		return false
	}
	ctx.LogInfo("==========TestWasmJsonContract invokeCallContractAddValue ============")
	for i ,n := range notifies{
		ctx.LogInfo(fmt.Sprintf("notify %d is %v",i, n))
	}



	return true
}

func deployCallWasmJsonContract(ctx *testframework.TestFrameworkContext, signer *account.Account) (common.Uint256, error){

	code, err := ioutil.ReadFile(filePath + "/" + "callContract.wasm")
	if err != nil {
		return common.Uint256{}, errors.New("")
	}

	codeHash := common.ToHexString(code)

	txHash, err := ctx.Ont.Rpc.DeploySmartContract(0,0,
		signer,
		types.WASMVM,
		true,
		codeHash,
		"cwjc",
		"1.0",
		"test",
		"",
		"",
	)

	if err != nil {
		return common.Uint256{}, fmt.Errorf("TestNep5Contract DeploySmartContract error:%s", err)
	}
	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30 * time.Second)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("WaitForGenerateBlock error:%s", err)
	}
	return txHash, nil
}

func invokeCallContractGetValue(ctx *testframework.TestFrameworkContext, acc *account.Account,address common.Address) (common.Uint256, error) {
	method := "getValue"
	params := make([]interface{},1)
	params[0] = "TestKey"
	txHash,err := ctx.Ont.Rpc.InvokeWasmVMSmartContract(0,0,acc,1,address,method, wasmvm.Json,params)
	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30 * time.Second)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("WaitForGenerateBlock error:%s", err)
	}
	return txHash, nil
}

func invokeCallContractAddValue(ctx *testframework.TestFrameworkContext, acc *account.Account,address common.Address) (common.Uint256, error) {
	method := "putValue"
	params := make([]interface{},2)
	params[0] = "TestKey"
	params[1] = "Hello world again!"
	txHash,err := ctx.Ont.Rpc.InvokeWasmVMSmartContract(0,0,acc,1,address,method, wasmvm.Json,params)
	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30 * time.Second)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("WaitForGenerateBlock error:%s", err)
	}
	return txHash, nil
}

