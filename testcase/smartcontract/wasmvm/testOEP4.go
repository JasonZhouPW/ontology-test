package wasmvm

import (
	"github.com/ontio/ontology-test/testframework"
	"fmt"
	"github.com/ontio/ontology/smartcontract/service/wasmvm"
	"time"
)


func TestWasmOEP4(ctx *testframework.TestFrameworkContext) bool {
	testFile := filePath + "/" + "OEP4.wasm"
	signer,_ := ctx.GetDefaultAccount()

	txhash,addr,err := DeployWasmJsonContract(ctx,signer,testFile,"testContract","1")
	if err != nil{
		fmt.Printf("deploy failed:%s\n",err.Error())
		return false
	}

	fmt.Printf("the txHash is %s\n",txhash.ToHexString())
	fmt.Printf("contract address is %s\n", addr.ToBase58())

	ctx.LogInfo("=====================invoke init==============================")
	txhash,err = InvokeWasmContract(ctx,
		signer,
		addr,
		"init",
		wasmvm.Raw,byte(1),[]interface{}{})

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
	fmt.Printf("events.Notify:%v",events.Notify)
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("=====================invoke init end==============================")

	ctx.LogInfo("=====================invoke name==============================")
	res,err := PreExecWasmContract(ctx,
		addr,
		"name",
		wasmvm.Raw,byte(1),[]interface{}{})

	tmp ,err:= res.Result.ToString()
	fmt.Printf("name is %v\n",tmp)

	ctx.LogInfo("=====================invoke name end==============================")

	ctx.LogInfo("=====================invoke symbol==============================")
	res,err = PreExecWasmContract(ctx,
		addr,
		"symbol",
		wasmvm.Raw,byte(1),[]interface{}{})

	tmp ,err= res.Result.ToString()
	fmt.Printf("symbol is %v\n",tmp)
	ctx.LogInfo("=====================invoke symbol end==============================")

	ctx.LogInfo("=====================invoke decimals==============================")
	res,err = PreExecWasmContract(ctx,
		addr,
		"decimals",
		wasmvm.Raw,byte(1),[]interface{}{})

	tmp ,err= res.Result.ToString()
	fmt.Printf("decimals is %v\n",tmp)
	ctx.LogInfo("=====================invoke decimals end==============================")

	ctx.LogInfo("=====================invoke totalSupply==============================")
	res,err = PreExecWasmContract(ctx,
		addr,
		"totalSupply",
		wasmvm.Raw,byte(1),[]interface{}{})

	tmp ,err= res.Result.ToString()
	fmt.Printf("totalSupply is %v\n",tmp)
	ctx.LogInfo("=====================invoke totalSupply end==============================")

	ctx.LogInfo("=====================invoke balanceOf==============================")
	res,err = PreExecWasmContract(ctx,
		addr,
		"balanceOf",
		wasmvm.Raw,byte(1),[]interface{}{"Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT"})

	tmp ,err= res.Result.ToString()
	fmt.Printf("balanceOf Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT is %v\n",tmp)
	ctx.LogInfo("=====================invoke balanceOf end==============================")

	ctx.LogInfo("=====================invoke transfer==============================")
	txhash,err = InvokeWasmContract(ctx,
		signer,
		addr,
		"transfer",
		wasmvm.Raw,byte(1),[]interface{}{"Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT","AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb",int64(300)})

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
	fmt.Printf("events.Notify:%v",events.Notify)
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("=====================invoke transfer end ==============================")

	ctx.LogInfo("=====================invoke balanceOf==============================")
	res,err = PreExecWasmContract(ctx,
		addr,
		"balanceOf",
		wasmvm.Raw,byte(1),[]interface{}{"AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb"})

	tmp ,err= res.Result.ToString()
	fmt.Printf("balanceOf AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb is %v\n",tmp)
	ctx.LogInfo("=====================invoke balanceOf end==============================")

	ctx.LogInfo("=====================invoke balanceOf==============================")
	res,err = PreExecWasmContract(ctx,
		addr,
		"balanceOf",
		wasmvm.Raw,byte(1),[]interface{}{"Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT"})

	tmp ,err= res.Result.ToString()
	fmt.Printf("balanceOf Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT is %v\n",tmp)
	ctx.LogInfo("=====================invoke balanceOf end==============================")


	return true
}
