package wasmvm
//
//import (
//	"github.com/ontio/ontology-test/testframework"
//	"fmt"
//	"github.com/ontio/ontology/smartcontract/service/wasmvm"
//	"time"
//)
//
//
//func TestWasmOEP4(ctx *testframework.TestFrameworkContext) bool {
//	timeoutSec := 30 * time.Second
//	testFile := filePath + "/" + "OEP4.wasm"
//	signer,_ := ctx.GetDefaultAccount()
//	//account2,_ := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
//	//account3, err := ctx.GetAccount("AK98G45DhmPXg4TFPG1KjftvkEaHbU8SHM")
//
//	txhash,addr,err := DeployWasmJsonContract(ctx,signer,testFile,"testContract","1")
//	if err != nil{
//		fmt.Printf("deploy failed:%s\n",err.Error())
//		return false
//	}
//
//	fmt.Printf("the txHash is %s\n",txhash.ToHexString())
//	fmt.Printf("contract address is %s\n", addr.ToBase58())
//
//	ctx.LogInfo("=====================invoke init==============================")
//	txhash,err = InvokeWasmContract(ctx,
//		signer,
//		addr,
//		"init",
//		wasmvm.Raw,byte(1),[]interface{}{})
//	if err != nil{
//		fmt.Println("init error:" + err.Error())
//		return false
//	}
//
//	_, err = ctx.Ont.WaitForGenerateBlock(timeoutSec)
//	if err != nil {
//		return false
//	}
//
//	ctx.LogInfo("txhash is %s\n",txhash.ToHexString())
//
//	events, err := ctx.Ont.GetSmartContractEvent(txhash.ToHexString())
//	if err != nil {
//		ctx.LogError("TestWasmOEP4 GetSmartContractEvent error:%s", err)
//		return false
//	}
//	fmt.Printf("event is %v\n", events)
//	if events.State == 0 {
//		ctx.LogError("TestWasmOEP4 failed invoked exec state return 0")
//		return false
//	}
//	fmt.Printf("events.Notify:%v",events.Notify)
//	for _,notify:= range events.Notify{
//		ctx.LogInfo("%+v", notify)
//	}
//	ctx.LogInfo("=====================invoke init end==============================")
//
//	ctx.LogInfo("=====================invoke name==============================")
//	res,err := PreExecWasmContract(ctx,
//		addr,
//		"name",
//		wasmvm.Raw,byte(1),[]interface{}{})
//
//	if err != nil{
//		fmt.Printf("invoke name failed:%s\n",err.Error())
//		return false
//	}
//
//	tmp ,err:= res.Result.ToString()
//	fmt.Printf("name is %v\n",tmp)
//
//	ctx.LogInfo("=====================invoke name end==============================")
//
//	ctx.LogInfo("=====================invoke symbol==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"symbol",
//		wasmvm.Raw,byte(1),[]interface{}{})
//
//	tmp ,err= res.Result.ToString()
//	fmt.Printf("symbol is %v\n",tmp)
//	ctx.LogInfo("=====================invoke symbol end==============================")
//
//	ctx.LogInfo("=====================invoke decimals==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"decimals",
//		wasmvm.Raw,byte(1),[]interface{}{})
//
//	tmp ,err= res.Result.ToString()
//	fmt.Printf("decimals is %v\n",tmp)
//	ctx.LogInfo("=====================invoke decimals end==============================")
//
//	ctx.LogInfo("=====================invoke totalSupply==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"totalSupply",
//		wasmvm.Raw,byte(1),[]interface{}{})
//
//	tmp ,err= res.Result.ToString()
//	fmt.Printf("totalSupply is %v\n",tmp)
//	ctx.LogInfo("=====================invoke totalSupply end==============================")
//
//	ctx.LogInfo("=====================invoke balanceOf==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"balanceOf",
//		wasmvm.Raw,byte(1),[]interface{}{"Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT"})
//
//	tmp ,err= res.Result.ToString()
//	fmt.Printf("balanceOf Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT is %v\n",tmp)
//	ctx.LogInfo("=====================invoke balanceOf end==============================")
//
//	ctx.LogInfo("=====================invoke transfer==============================")
//	txhash,err = InvokeWasmContract(ctx,
//		signer,
//		addr,
//		"transfer",
//		wasmvm.Raw,byte(1),[]interface{}{"Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT","AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb",int64(300)})
//
//	_, err = ctx.Ont.WaitForGenerateBlock(timeoutSec)
//	if err != nil {
//		return false
//	}
//
//	events, err = ctx.Ont.GetSmartContractEvent(txhash.ToHexString())
//	if err != nil {
//		ctx.LogError("TestWasmOEP4 GetSmartContractEvent error:%s", err)
//		return false
//	}
//	fmt.Printf("event is %v\n", events)
//	if events.State == 0 {
//		ctx.LogError("TestWasmOEP4 failed invoked exec state return 0")
//		return false
//	}
//	fmt.Printf("events.Notify:%v",events.Notify)
//	for _,notify:= range events.Notify{
//		ctx.LogInfo("%+v", notify)
//	}
//	ctx.LogInfo("=====================invoke transfer end ==============================")
//
//	ctx.LogInfo("=====================invoke balanceOf==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"balanceOf",
//		wasmvm.Raw,byte(1),[]interface{}{"AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb"})
//
//	tmp ,err= res.Result.ToString()
//	fmt.Printf("balanceOf AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb is %v\n",tmp)
//	ctx.LogInfo("=====================invoke balanceOf end==============================")
//
//	ctx.LogInfo("=====================invoke balanceOf==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"balanceOf",
//		wasmvm.Raw,byte(1),[]interface{}{"Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT"})
//
//	tmp ,err= res.Result.ToString()
//	fmt.Printf("balanceOf Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT is %v\n",tmp)
//	ctx.LogInfo("=====================invoke balanceOf end==============================")
//
//
//	ctx.LogInfo("=====================invoke approve==============================")
//	txhash,err = InvokeWasmContract(ctx,
//		signer,
//		addr,
//		"approve",
//		wasmvm.Raw,byte(1),[]interface{}{"Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT","AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb",int64(800)})
//
//	_, err = ctx.Ont.WaitForGenerateBlock(timeoutSec)
//	if err != nil {
//		return false
//	}
//
//	events, err = ctx.Ont.GetSmartContractEvent(txhash.ToHexString())
//	if err != nil {
//		ctx.LogError("TestWasmOEP4 GetSmartContractEvent error:%s", err)
//		return false
//	}
//	fmt.Printf("event is %v\n", events)
//	if events.State == 0 {
//		ctx.LogError("TestWasmOEP4 failed invoked exec state return 0")
//		return false
//	}
//	fmt.Printf("events.Notify:%v",events.Notify)
//	for _,notify:= range events.Notify{
//		ctx.LogInfo("%+v", notify)
//	}
//	ctx.LogInfo("=====================invoke approve end ==============================")
//
//	ctx.LogInfo("=====================invoke allownance==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"allowance",
//		wasmvm.Raw,byte(1),[]interface{}{"Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT","AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb"})
//
//	tmp ,err= res.Result.ToString()
//	fmt.Printf("allowance Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT is %v\n",tmp)
//	ctx.LogInfo("=====================invoke allownance end==============================")
//
//	ctx.LogInfo("=====================invoke transferFrom==============================")
//	txhash,err = InvokeWasmContract(ctx,
//		account2,
//		addr,
//		"transferFrom",
//		wasmvm.Raw,byte(1),[]interface{}{"AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb","Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT","AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb",int64(800)})
//
//	_, err = ctx.Ont.WaitForGenerateBlock(timeoutSec)
//	if err != nil {
//		return false
//	}
//
//	events, err = ctx.Ont.GetSmartContractEvent(txhash.ToHexString())
//	if err != nil {
//		ctx.LogError("TestWasmOEP4 GetSmartContractEvent error:%s", err)
//		return false
//	}
//	fmt.Printf("event is %v\n", events)
//	if events.State == 0 {
//		ctx.LogError("TestWasmOEP4 failed invoked exec state return 0")
//		return false
//	}
//	fmt.Printf("events.Notify:%v",events.Notify)
//	for _,notify:= range events.Notify{
//		ctx.LogInfo("%+v", notify)
//	}
//	ctx.LogInfo("=====================invoke approve end ==============================")
//
//	ctx.LogInfo("=====================invoke allownance==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"allowance",
//		wasmvm.Raw,byte(1),[]interface{}{"Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT","AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb"})
//
//	tmp ,err= res.Result.ToString()
//	fmt.Printf("allowance Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT is %v\n",tmp)
//	ctx.LogInfo("=====================invoke allownance end==============================")
//
//	ctx.LogInfo("=====================invoke balanceOf==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"balanceOf",
//		wasmvm.Raw,byte(1),[]interface{}{"AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb"})
//
//	tmp ,err= res.Result.ToString()
//	fmt.Printf("balanceOf AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb is %v\n",tmp)
//	ctx.LogInfo("=====================invoke balanceOf end==============================")
//
//
//	ctx.LogInfo("=====================invoke transferMulti==============================")
//	argarr := []interface{}{[]string{"Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT","AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb","500"},[]string{"Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT","AK98G45DhmPXg4TFPG1KjftvkEaHbU8SHM","600"}}
//
//	txhash,err = InvokeWasmContract(ctx,
//		signer,
//		addr,
//		"transferMulti",
//		wasmvm.Raw,byte(1),[]interface{}{argarr})
//
//	_, err = ctx.Ont.WaitForGenerateBlock(timeoutSec)
//	if err != nil {
//		return false
//	}
//
//	events, err = ctx.Ont.GetSmartContractEvent(txhash.ToHexString())
//	if err != nil {
//		ctx.LogError("TestWasmOEP4 GetSmartContractEvent error:%s", err)
//		return false
//	}
//	fmt.Printf("event is %v\n", events)
//	if events.State == 0 {
//		ctx.LogError("TestWasmOEP4 failed invoked exec state return 0")
//		return false
//	}
//	fmt.Printf("events.Notify:%v",events.Notify)
//	for _,notify:= range events.Notify{
//		ctx.LogInfo("%+v", notify)
//	}
//	ctx.LogInfo("=====================invoke transferMulti end ==============================")
//
//	ctx.LogInfo("=====================invoke balanceOf==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"balanceOf",
//		wasmvm.Raw,byte(1),[]interface{}{"AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb"})
//
//	tmp ,err= res.Result.ToString()
//	fmt.Printf("balanceOf AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb is %v\n",tmp)
//	ctx.LogInfo("=====================invoke balanceOf end==============================")
//
//	ctx.LogInfo("=====================invoke balanceOf==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"balanceOf",
//		wasmvm.Raw,byte(1),[]interface{}{"AK98G45DhmPXg4TFPG1KjftvkEaHbU8SHM"})
//
//	tmp ,err= res.Result.ToString()
//	fmt.Printf("balanceOf AK98G45DhmPXg4TFPG1KjftvkEaHbU8SHM is %v\n",tmp)
//	ctx.LogInfo("=====================invoke balanceOf end==============================")
//
//
//	return true
//}
