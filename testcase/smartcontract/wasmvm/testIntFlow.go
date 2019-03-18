package wasmvm

//
//import (
//	"github.com/ontio/ontology-test/testframework"
//	"fmt"
//	"github.com/ontio/ontology/smartcontract/service/wasmvm"
//)
//
//func TestWasmIntFlow(ctx *testframework.TestFrameworkContext) bool {
//	testFile := filePath + "/" + "testIntFlow.wasm"
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
//	ctx.LogInfo("=====================invoke multiple before==============================")
//	res,err := PreExecWasmContract(ctx,
//		addr,
//		"multiple",
//		wasmvm.Raw,byte(1),[]interface{}{int64(-9223372036854775807),int64(100000)})
//		//wasmvm.Raw,byte(1),[]interface{}{int64(12),int64(2)})
//
//	fmt.Printf("result is %v\n",res)
//	if res != nil{
//		tmp ,_:= res.Result.ToString()
//		fmt.Printf("res is %v\n",tmp)
//	}
//
//	ctx.LogInfo("=====================invoke multiple end==============================")
//
//	ctx.LogInfo("=====================invoke add before==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"add",
//		wasmvm.Raw,byte(1),[]interface{}{int64(9223372036854775807),int64(100000)})
//	//wasmvm.Raw,byte(1),[]interface{}{int64(12),int64(2)})
//
//	fmt.Printf("result is %v\n",res)
//	if res != nil{
//		tmp ,_:= res.Result.ToString()
//		fmt.Printf("res is %v\n",tmp)
//	}
//	ctx.LogInfo("=====================invoke add end==============================")
//
//	ctx.LogInfo("=====================invoke add before==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"add",
//		wasmvm.Raw,byte(1),[]interface{}{int64(-1),int64(-100)})
//	//wasmvm.Raw,byte(1),[]interface{}{int64(12),int64(2)})
//
//	fmt.Printf("result is %v\n",res)
//	if res != nil{
//		tmp ,_:= res.Result.ToString()
//		fmt.Printf("res is %v\n",tmp)
//	}
//	ctx.LogInfo("=====================invoke add end==============================")
//	return true
//
//}
