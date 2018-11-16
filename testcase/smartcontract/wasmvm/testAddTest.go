package wasmvm

import (
	"github.com/ontio/ontology-test/testframework"
	"fmt"
	"github.com/ontio/ontology/smartcontract/service/wasmvm"
	"time"
)



func TestWasmAddTest1(ctx *testframework.TestFrameworkContext) bool {

	testFile := filePath + "/" + "test1.wasm"
	signer,_ := ctx.GetDefaultAccount()

	txhash,addr,err := DeployWasmJsonContract(ctx,signer,testFile,"testContract","1")
	if err != nil{
		fmt.Printf("deploy failed:%s\n",err.Error())
		return false
	}

	fmt.Printf("the txHash is %s\n",txhash.ToHexString())
	fmt.Printf("contract address is %s\n", addr.ToBase58())

	ctx.LogInfo("=====================invoke add==============================")
	txhash,err = InvokeWasmContract(ctx,
							signer,
							addr,
		"add",
		wasmvm.Json,byte(1),[]interface{}{1,2})

	_, err = ctx.Ont.WaitForGenerateBlock(30 * time.Second)
	if err != nil {
		return false
	}

	events, err := ctx.Ont.GetSmartContractEvent(txhash.ToHexString())
	if err != nil {
		ctx.LogError("TestOEP5Py GetSmartContractEvent error:%s", err)
		return false
	}
	fmt.Printf("event is %v\n", events)
	if events.State == 0 {
		ctx.LogError("TestOEP5Py failed invoked exec state return 0")
		return false
	}
	fmt.Printf("events.Notify:%v",events.Notify)
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}

	ctx.LogInfo("=====================invoke add end==============================")


	ctx.LogInfo("=====================invoke put==============================")
	txhash,err = InvokeWasmContract(ctx,
		signer,
		addr,
		"put",
		wasmvm.Json,byte(1),[]interface{}{"mykey","myValue"})

	_, err = ctx.Ont.WaitForGenerateBlock(30 * time.Second)
	if err != nil {
		return false
	}

	events, err = ctx.Ont.GetSmartContractEvent(txhash.ToHexString())
	if err != nil {
		ctx.LogError("TestOEP5Py GetSmartContractEvent error:%s", err)
		return false
	}
	fmt.Printf("event is %v\n", events)
	if events.State == 0 {
		ctx.LogError("TestOEP5Py failed invoked exec state return 0")
		return false
	}
	fmt.Printf("events.Notify:%v",events.Notify)
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}

	ctx.LogInfo("=====================invoke put end==============================")

	ctx.LogInfo("=====================invoke get==============================")
	txhash,err = InvokeWasmContract(ctx,
		signer,
		addr,
		"get",
		wasmvm.Json,byte(1),[]interface{}{"mykey"})

	_, err = ctx.Ont.WaitForGenerateBlock(30 * time.Second)
	if err != nil {
		return false
	}

	events, err = ctx.Ont.GetSmartContractEvent(txhash.ToHexString())
	if err != nil {
		ctx.LogError("TestOEP5Py GetSmartContractEvent error:%s", err)
		return false
	}
	fmt.Printf("event is %v\n", events)
	if events.State == 0 {
		ctx.LogError("TestOEP5Py failed invoked exec state return 0")
		return false
	}
	fmt.Printf("events.Notify:%v",events.Notify)
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}

	ctx.LogInfo("=====================invoke get end==============================")


	ctx.LogInfo("=====================invoke get preexec==============================")
	res,err := PreExecWasmContract(ctx,
		addr,
		"get",
		wasmvm.Json,byte(1),[]interface{}{"mykey"})

	tmp ,err:= res.Result.ToString()
	fmt.Printf("res is %v\n",tmp)

	ctx.LogInfo("=====================invoke get end==============================")


	return true

}
