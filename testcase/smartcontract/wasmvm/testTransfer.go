package wasmvm

import (
	"github.com/ontio/ontology-test/testframework"
	"fmt"
	"github.com/ontio/ontology/smartcontract/service/wasmvm"
	"time"
)


func TestWasmTransfer(ctx *testframework.TestFrameworkContext) bool {

	testFile := filePath + "/" + "transfer.wasm"
	signer,_ := ctx.GetDefaultAccount()

	txhash,addr,err := DeployWasmJsonContract(ctx,signer,testFile,"testContract","1")
	if err != nil{
		fmt.Printf("deploy failed:%s\n",err.Error())
		return false
	}

	fmt.Printf("the txHash is %s\n",txhash.ToHexString())
	fmt.Printf("contract address is %s\n", addr.ToBase58())

	ctx.LogInfo("=====================invoke transferONT==============================")
	txhash,err = InvokeWasmContract(ctx,
		signer,
		addr,
		"transferONT",
		wasmvm.Raw,byte(1),[]interface{}{"Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT","AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb",int64(100)})

	_, err = ctx.Ont.WaitForGenerateBlock(30 * time.Second)
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
	fmt.Printf("events.Notify:%v\n",events.Notify)
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("=====================invoke transferONT end==============================")

	ctx.LogInfo("=====================invoke transferONG==============================")
	txhash,err = InvokeWasmContract(ctx,
		signer,
		addr,
		"transferONG",
		wasmvm.Raw,byte(1),[]interface{}{"Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT","AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb",int64(100)})

	_, err = ctx.Ont.WaitForGenerateBlock(30 * time.Second)
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
	fmt.Printf("events.Notify:%v\n",events.Notify)
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("=====================invoke transferONG end==============================")


	return true

}