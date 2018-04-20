package wasmvm

import (
	"github.com/ontio/ontology/common"
	"fmt"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/account"
	"math/big"
	"github.com/ontio/ontology/smartcontract/service/wasmvm"
	"time"
)



func TestAssetRawContract(ctx *testframework.TestFrameworkContext) bool {
	wasmWallet := "/home/zhoupw/work/go/src/github.com/ontio/ontology/wallet.dat"
	wasmWalletPwd := "123456"

	wallet, err := ctx.Ont.OpenWallet(wasmWallet, wasmWalletPwd)
	if err != nil {
		ctx.LogError("OpenWallet:%s error:%s", wasmWallet, err)
		return false
	}

	admin, err := wallet.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestAssetContract wallet.GetDefaultAccount error:%s", err)
		return false
	}


	txHash, err := DeployWasmJsonContract(ctx,admin,filePath + "/assetraw.wasm","tcoinRaw","1.0")

	if err != nil {
		ctx.LogError("TestAssetContract deploy error:%s", err)
		return false
	}

	ctx.LogInfo("TestAssetContract deploy TxHash:%x", txHash)

	address ,err := GetWasmContractAddress(filePath + "/assetraw.wasm")
	if err != nil{
		ctx.LogError("TestAssetContract GetWasmContractAddress error:%s", err)
		return false
	}

	txHash,err = invokeRawInit(ctx,admin,address)
	if err != nil {
		ctx.LogError("TestAssetContract invokeInit error:%s", err)
		return false
	}


	ctx.LogInfo("invokeContract: %x\n", txHash)
	ctx.LogInfo("TestAssetContract invokeInit success")
	notifies, err := ctx.Ont.Rpc.GetSmartContractEvent(txHash)
	if err != nil {
		ctx.LogError("TestAssetContract init invokeInit error:%s", err)
		return false
	}
	ctx.LogInfo("TestAssetContract invoke notify %s", notifies)
	fmt.Println("============result is===============")
	bs ,_:= common.HexToBytes(notifies[0].States[0].(string))

	fmt.Printf("+==========%s\n",string(bs))


	txHash,err = invokeTotalRawSupply(ctx,admin,address)
	if err != nil {
		ctx.LogError("TestAssetContract invokeTotalSupply error:%s", err)
		return false
	}


	ctx.LogInfo("invokeContract: %x\n", txHash)
	ctx.LogInfo("TestAssetContract invokeTotalSupply success")
	notifies, err = ctx.Ont.Rpc.GetSmartContractEvent(txHash)
	if err != nil {
		ctx.LogError("TestAssetContract init invokeTotalSupply error:%s", err)
		return false
	}
	ctx.LogInfo("TestAssetContract invokeTotalSupply notify %s", notifies)
	fmt.Println("============invokeTotalSupply result is===============")
	bs ,_= common.HexToBytes(notifies[0].States[0].(string))

	fmt.Printf("+==========%s\n",string(bs))

	txHash,err = invokeRawBalanceOf(ctx,admin,address,"00000001")
	if err != nil {
		ctx.LogError("TestAssetContract invokeBalanceOf error:%s", err)
		return false
	}


	ctx.LogInfo("invokeContract: %x\n", txHash)
	ctx.LogInfo("TestAssetContract invokeBalanceOf success")
	notifies, err = ctx.Ont.Rpc.GetSmartContractEvent(txHash)
	if err != nil {
		ctx.LogError("TestAssetContract init invokeBalanceOf error:%s", err)
		return false
	}
	ctx.LogInfo("TestAssetContract invokeBalanceOf notify %s", notifies)
	fmt.Println("============invokeBalanceOf 00000001 result is===============")
	bs ,_= common.HexToBytes(notifies[0].States[0].(string))

	fmt.Printf("+==========%s\n",string(bs))

	txHash,err = invokeRawTransfer(ctx,admin,address,"00000001","00000002",int64(20000))
	if err != nil {
		ctx.LogError("TestAssetContract invokeBalanceOf error:%s", err)
		return false
	}


	ctx.LogInfo("invokeContract: %x\n", txHash)
	ctx.LogInfo("TestAssetContract invokeBalanceOf success")
	notifies, err = ctx.Ont.Rpc.GetSmartContractEvent(txHash)
	if err != nil {
		ctx.LogError("TestAssetContract init invokeBalanceOf error:%s", err)
		return false
	}
	ctx.LogInfo("TestAssetContract invokeBalanceOf notify %s", notifies)
	fmt.Println("============invokeTransfer result is===============")
	bs ,_= common.HexToBytes(notifies[0].States[0].(string))

	fmt.Printf("+==========%s\n",string(bs))


	txHash,err = invokeRawBalanceOf(ctx,admin,address,"00000001")
	if err != nil {
		ctx.LogError("TestAssetContract invokeBalanceOf error:%s", err)
		return false
	}


	ctx.LogInfo("invokeContract: %x\n", txHash)
	ctx.LogInfo("TestAssetContract invokeBalanceOf success")
	notifies, err = ctx.Ont.Rpc.GetSmartContractEvent(txHash)
	if err != nil {
		ctx.LogError("TestAssetContract init invokeBalanceOf error:%s", err)
		return false
	}
	ctx.LogInfo("TestAssetContract invokeBalanceOf notify %s", notifies)
	fmt.Println("============00000001 result is===============")
	bs ,_= common.HexToBytes(notifies[0].States[0].(string))

	fmt.Printf("+==========%s\n",string(bs))


	txHash,err = invokeRawBalanceOf(ctx,admin,address,"00000002")
	if err != nil {
		ctx.LogError("TestAssetContract invokeBalanceOf error:%s", err)
		return false
	}

	ctx.LogInfo("invokeContract: %x\n", txHash)
	ctx.LogInfo("TestAssetContract invokeBalanceOf success")
	notifies, err = ctx.Ont.Rpc.GetSmartContractEvent(txHash)
	if err != nil {
		ctx.LogError("TestAssetContract init invokeBalanceOf error:%s", err)
		return false
	}
	ctx.LogInfo("TestAssetContract invokeBalanceOf notify %s", notifies)
	fmt.Println("============00000002 result is===============")
	bs ,_= common.HexToBytes(notifies[0].States[0].(string))

	fmt.Printf("+==========%s\n",string(bs))


	return true
}


func invokeRawInit(ctx *testframework.TestFrameworkContext, acc *account.Account,address common.Address) (common.Uint256, error) {
	method := "init"
	txHash,err := ctx.Ont.Rpc.InvokeWasmVMSmartContract(acc,new(big.Int),address,method, wasmvm.Raw,1,nil)
	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30 * time.Second)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("WaitForGenerateBlock error:%s", err)
	}
	return txHash, nil
}

func invokeTotalRawSupply(ctx *testframework.TestFrameworkContext, acc *account.Account,address common.Address) (common.Uint256, error) {
	method := "totalSupply"
	txHash,err := ctx.Ont.Rpc.InvokeWasmVMSmartContract(acc,new(big.Int),address,method, wasmvm.Raw,1,nil)
	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30 * time.Second)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("WaitForGenerateBlock error:%s", err)
	}
	return txHash, nil
}

func invokeRawBalanceOf(ctx *testframework.TestFrameworkContext, acc *account.Account,address common.Address,accountaddress string) (common.Uint256, error) {
	method := "balanceOf"
	params := make([]interface{},1)
	params[0] = accountaddress

	txHash,err := ctx.Ont.Rpc.InvokeWasmVMSmartContract(acc,new(big.Int),address,method, wasmvm.Raw,1,params)
	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30 * time.Second)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("WaitForGenerateBlock error:%s", err)
	}
	return txHash, nil
}

func invokeRawTransfer(ctx *testframework.TestFrameworkContext, acc *account.Account,address common.Address,from,to string,amount int64) (common.Uint256, error) {
	method := "transfer"
	params := make([]interface{},3)
	params[0] = from
	params[1] = to
	params[2] = amount

	txHash,err := ctx.Ont.Rpc.InvokeWasmVMSmartContract(acc,new(big.Int),address,method, wasmvm.Raw,1,params)
	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30 * time.Second)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("WaitForGenerateBlock error:%s", err)
	}
	return txHash, nil
}
