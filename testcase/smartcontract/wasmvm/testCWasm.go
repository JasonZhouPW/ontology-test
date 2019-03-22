package wasmvm

import (
	"time"
	"fmt"
	"github.com/ontio/ontology-test/testframework"
	"encoding/binary"
	"github.com/ontio/ontology/common"
)

func TestCOEP4(ctx *testframework.TestFrameworkContext) bool {
	testFile := filePath + "/" + "cOEP4.wasm"
	signer, _ := ctx.GetDefaultAccount()
	timeoutSec := 30 * time.Second
	txhash, addr, err := DeployWasmJsonContract(ctx, signer, testFile, "testContract", "1")
	if err != nil {
		fmt.Printf("deploy failed:%s\n", err.Error())
		return false
	}
	ctx.LogInfo("==contract address is %s", addr.ToBase58())
	account2, err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
	if err != nil {
		ctx.LogError("get account AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb failed")
		return false
	}


	ctx.LogInfo("=====================invoke transfer==============================")
	amountbs := make([]byte, 16)

	binary.LittleEndian.PutUint64(amountbs, uint64(500))
	//copy(parambs[:16], amountbs)
	//fmt.Printf("parambs %v\n", parambs)


	amount, _ := common.Uint256ParseFromBytes(amountbs)

	txhash, err = InvokeWasmContract(ctx,
		signer,
		addr,
		"transfer",
		[]interface{}{signer.Address, account2.Address, amount})

	_, err = ctx.Ont.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		return false
	}

	events, err := ctx.Ont.GetSmartContractEvent(txhash.ToHexString())
	if err != nil {
		ctx.LogError("TestWasmOEP4 GetSmartContractEvent error:%s", err)
		return false
	}
	fmt.Printf("event is %v\n", events)
	if events.State == 0 {
		ctx.LogError("TestWasmOEP4 failed invoked exec state return 0")
		return false
	}
	fmt.Printf("events.Notify:%v\n", events.Notify)
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("=====================invoke transfer end==============================")


	ctx.LogInfo("=====================invoke approve==============================")

	txhash, err = InvokeWasmContract(ctx,
		signer,
		addr,
		"approve",
		[]interface{}{signer.Address, account2.Address, amount})

	_, err = ctx.Ont.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		return false
	}

	events, err = ctx.Ont.GetSmartContractEvent(txhash.ToHexString())
	if err != nil {
		ctx.LogError("TestWasmOEP4 GetSmartContractEvent error:%s", err)
		return false
	}
	fmt.Printf("event is %v\n", events)
	if events.State == 0 {
		ctx.LogError("TestWasmOEP4 failed invoked exec state return 0")
		return false
	}
	fmt.Printf("events.Notify:%v\n", events.Notify)
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("=====================invoke approve end==============================")


	ctx.LogInfo("=====================invoke transfer from==============================")

	txhash, err = InvokeWasmContract(ctx,
		account2,
		addr,
		"transferfrom",
		[]interface{}{account2.Address,signer.Address, account2.Address, amount})

	_, err = ctx.Ont.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		return false
	}

	events, err = ctx.Ont.GetSmartContractEvent(txhash.ToHexString())
	if err != nil {
		ctx.LogError("TestWasmOEP4 GetSmartContractEvent error:%s", err)
		return false
	}
	fmt.Printf("event is %v\n", events)
	if events.State == 0 {
		ctx.LogError("TestWasmOEP4 failed invoked exec state return 0")
		return false
	}
	fmt.Printf("events.Notify:%v\n", events.Notify)
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("=====================invoke transfer from end==============================")

	return true

}