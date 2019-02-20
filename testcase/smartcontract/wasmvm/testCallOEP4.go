package wasmvm
//
//import (
//	"github.com/ontio/ontology-test/testframework"
//	"fmt"
//	"github.com/ontio/ontology/smartcontract/service/wasmvm"
//	"time"
//)
//
//func TestWasmCallOEP4(ctx *testframework.TestFrameworkContext) bool {
//	timeoutSec := 30 * time.Second
//
//	testFile := filePath + "/" + "callOEP4.wasm"
//	signer,_ := ctx.GetDefaultAccount()
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
//
//	ctx.LogInfo("=====================invoke balanceOf before==============================")
//	res,err := PreExecWasmContract(ctx,
//		addr,
//		"balanceOfOEP4",
//		wasmvm.Raw,byte(1),[]interface{}{"AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb"})
//
//	tmp ,err:= res.Result.ToString()
//	fmt.Printf("balanceOf AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb is %v\n",tmp)
//	ctx.LogInfo("=====================invoke balanceOf end==============================")
//
//	ctx.LogInfo("=====================invoke balanceOf==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"balanceOfOEP4",
//		wasmvm.Raw,byte(1),[]interface{}{"Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT"})
//
//	tmp ,err= res.Result.ToString()
//	fmt.Printf("balanceOf Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT is %v\n",tmp)
//	ctx.LogInfo("=====================invoke balanceOf end==============================")
//
//
//	ctx.LogInfo("=====================invoke transfer==============================")
//	txhash,err = InvokeWasmContract(ctx,
//		signer,
//		addr,
//		"transferOEP4",
//		wasmvm.Raw,byte(1),[]interface{}{"Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT","AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb",int64(300)})
//
//	_, err = ctx.Ont.WaitForGenerateBlock(timeoutSec)
//	if err != nil {
//		return false
//	}
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
//	ctx.LogInfo("=====================invoke transfer end==============================")
//
//	ctx.LogInfo("=====================invoke balanceOf after==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"balanceOfOEP4",
//		wasmvm.Raw,byte(1),[]interface{}{"AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb"})
//
//	tmp ,err= res.Result.ToString()
//	fmt.Printf("balanceOf AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb is %v\n",tmp)
//	ctx.LogInfo("=====================invoke balanceOf end==============================")
//
//	ctx.LogInfo("=====================invoke balanceOf==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"balanceOfOEP4",
//		wasmvm.Raw,byte(1),[]interface{}{"Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT"})
//
//	tmp ,err= res.Result.ToString()
//	fmt.Printf("balanceOf Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT is %v\n",tmp)
//	ctx.LogInfo("=====================invoke balanceOf end==============================")
//
//	return true
//
//}
