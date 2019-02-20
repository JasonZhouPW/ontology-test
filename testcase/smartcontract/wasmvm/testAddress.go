package wasmvm
//
//import (
//	"github.com/ontio/ontology-test/testframework"
//	"fmt"
//	"github.com/ontio/ontology/smartcontract/service/wasmvm"
//)
//
//func TestWasmAddressTest(ctx *testframework.TestFrameworkContext) bool {
//
//	testFile := filePath + "/" + "testAddress.wasm"
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
//	ctx.LogInfo("=====================invoke selfAddr before==============================")
//	res,err := PreExecWasmContract(ctx,
//		addr,
//		"selfAddr",
//		wasmvm.Raw,byte(1),[]interface{}{})
//
//	fmt.Printf("result is %v\n",res)
//	if res != nil{
//		tmp ,_:= res.Result.ToString()
//		fmt.Printf("res is %v\n",tmp)
//	}
//
//	ctx.LogInfo("=====================invoke selfAddr end==============================")
//
//	ctx.LogInfo("=====================invoke callingAddr before==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"callerAddr",
//		wasmvm.Raw,byte(1),[]interface{}{})
//
//	fmt.Printf("result is %v\n",res)
//	if res != nil{
//		tmp ,_:= res.Result.ToString()
//		fmt.Printf("res is %v\n",tmp)
//	}
//
//	ctx.LogInfo("=====================invoke callingAddr end==============================")
//
//	ctx.LogInfo("=====================invoke entryAddr before==============================")
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"entryAddr",
//		wasmvm.Raw,byte(1),[]interface{}{})
//
//	fmt.Printf("result is %v\n",res)
//	if res != nil{
//		tmp ,_:= res.Result.ToString()
//		fmt.Printf("res is %v\n",tmp)
//	}
//
//	ctx.LogInfo("=====================invoke entryAddr end==============================")
//
//	return true
//}