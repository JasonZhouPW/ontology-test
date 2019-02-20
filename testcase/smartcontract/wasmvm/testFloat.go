package wasmvm
//
//import (
//	"fmt"
//	"github.com/ontio/ontology/smartcontract/service/wasmvm"
//	"github.com/ontio/ontology-test/testframework"
//)
//
//func TestFloat(ctx *testframework.TestFrameworkContext) bool {
//	testFile := filePath + "/" + "testFloat.wasm"
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
//		"init",
//		wasmvm.Raw,byte(1),[]interface{}{int64(8)})
//	//wasmvm.Raw,byte(1),[]interface{}{int64(12),int64(2)})
//
//	fmt.Printf("result is %v\n",res)
//	if res != nil{
//		tmp ,_:= res.Result.ToString()
//		fmt.Printf("res is %v\n",tmp)
//	}
//
//	ctx.LogInfo("=====================invoke multiple end==============================")
//
//	return true
//
//}
