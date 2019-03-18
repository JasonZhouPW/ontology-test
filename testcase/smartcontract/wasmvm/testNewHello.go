package wasmvm

//
//import (
//	"github.com/ontio/ontology-test/testframework"
//	"fmt"
//	"bytes"
//	"github.com/ontio/ontology/common/serialization"
//	"time"
//)
//
//func TestNewHello(ctx *testframework.TestFrameworkContext) bool {
//	testFile := filePath + "/" + "helloworld2.wasm"
//	signer,_ := ctx.GetDefaultAccount()
//	timeoutSec := 30 * time.Second
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
//	ctx.LogInfo("=====================invoke name==============================")
//	res,err := PreExecWasmContract(ctx,
//		addr,
//		"hello2",
//		byte(1),[]interface{}{"Jordan"})
//
//	if err != nil{
//		fmt.Printf("invoke name failed:%s\n",err.Error())
//		return false
//	}
//
//	bs,err := res.Result.ToByteArray()
//
//	fmt.Printf("res is %v\n",bs)
//
//	tmp ,err:= serialization.ReadString(bytes.NewBuffer(bs))
//	fmt.Printf("return is %v\n",tmp)
//
//	ctx.LogInfo("=====================invoke name end==============================")
//
//	ctx.LogInfo("=====================invoke save==============================")
//	txhash,err = InvokeWasmContract(ctx,
//		signer,
//		addr,
//		"save",
//		byte(1),[]interface{}{"mykey","myvalue"})
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
//	ctx.LogInfo("=====================invoke save end==============================")
//
//	ctx.LogInfo("=====================invoke get==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"get",
//		byte(1),[]interface{}{"mykey"})
//
//	if err != nil{
//		fmt.Printf("invoke name failed:%s\n",err.Error())
//		return false
//	}
//
//	bs,err = res.Result.ToByteArray()
//
//	fmt.Printf("res is %v\n",bs)
//
//	tmp ,err= serialization.ReadString(bytes.NewBuffer(bs))
//	fmt.Printf("return is %v\n",tmp)
//
//	ctx.LogInfo("=====================invoke get end==============================")
//
//
//	return true
//}
