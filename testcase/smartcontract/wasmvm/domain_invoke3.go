package wasmvm

import (
	"github.com/ontio/ontology-test/testframework"
	"fmt"
)



func TestDomainContract_Invoke3(ctx *testframework.TestFrameworkContext) bool {
	wasmWallet := "wallet.dat"
	wasmWalletPwd := "123456"

	wallet, err := ctx.Ont.OpenWallet(wasmWallet, wasmWalletPwd)
	if err != nil {
		ctx.LogError("OpenWallet:%s error:%s", wasmWallet, err)
		return false
	}

	admin, err := wallet.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainContract_Invoke3 wallet.GetDefaultAccount error:%s", err)
		return false
	}
	address, err := GetWasmContractAddress(filePath + "/domain.wasm")


	//current Price
	txHash,err := invokeDomainCurrentPrice(ctx,admin,address,"www.goodthings.com")
	if err != nil {
		ctx.LogError("TestDomainContract_Invoke3 invokeDomainCurrentPrice error:%s", err)
		return false
	}

	notifies, err := ctx.Ont.Rpc.GetSmartContractEvent(txHash)
	if err != nil {
		ctx.LogError("TestDomainContract_Invoke3 invokeDomainCurrentPrice error:%s", err)
		return false
	}

	if len(notifies) < 1{
		ctx.LogError("TestDomainContract_Invoke3 invokeDomainCurrentPrice return notifies count error!")
		return false
	}
	ctx.LogInfo("==========TestDomainContract_Invoke3 invokeDomainCurrentPrice ============")
	for i ,n := range notifies{
		ctx.LogInfo(fmt.Sprintf("notify %d is %v",i, n))
	}

	//deal
	txHash,err = invokeDomainDeal(ctx,admin,address,"TA4tBPFEn7Amutm7QWTBYesEHE5sbWZKsB","www.goodthings.com")
	if err != nil {
		ctx.LogError("TestDomainContract_Invoke3 invokeDomainDeal error:%s", err)
		return false
	}

	notifies, err = ctx.Ont.Rpc.GetSmartContractEvent(txHash)
	if err != nil {
		ctx.LogError("TestDomainContract_Invoke3 invokeDomainDeal error:%s", err)
		return false
	}

	if len(notifies) < 1{
		ctx.LogError("TestDomainContract_Invoke3 invokeDomainDeal return notifies count error!")
		return false
	}
	ctx.LogInfo("==========TestDomainContract_Invoke3 invokeDomainDeal ============")
	for i ,n := range notifies{
		ctx.LogInfo(fmt.Sprintf("notify %d is %v",i, n))
	}

	//query
	txHash,err = invokeDomainQuery(ctx,admin,address,"www.goodthings.com")
	if err != nil {
		ctx.LogError("TestDomainContract_Invoke3 invokeDomainQuery error:%s", err)
		return false
	}

	notifies, err = ctx.Ont.Rpc.GetSmartContractEvent(txHash)
	if err != nil {
		ctx.LogError("TestDomainContract_Invoke3 invokeDomainQuery error:%s", err)
		return false
	}

	if len(notifies) < 1{
		ctx.LogError("TestDomainContract_Invoke3 invokeDomainQuery return notifies count error!")
		return false
	}
	ctx.LogInfo("==========TestDomainContract_Invoke3 invokeDomainQuery ============")
	for i ,n := range notifies{
		ctx.LogInfo(fmt.Sprintf("notify %d is %v",i, n))
	}


	return true
}

