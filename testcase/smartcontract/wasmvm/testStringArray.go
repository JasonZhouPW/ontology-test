package wasmvm
//
//import (
//	"fmt"
//	"github.com/ontio/ontology-test/testframework"
//	"github.com/ontio/ontology/smartcontract/service/wasmvm"
//)
//
//func TestStrarr(ctx *testframework.TestFrameworkContext) bool {
//	testFile := filePath + "/" + "teststringArray.wasm"
//	signer, _ := ctx.GetDefaultAccount()
//
//	txhash, addr, err := DeployWasmJsonContract(ctx, signer, testFile, "testContract", "1")
//	if err != nil {
//		fmt.Printf("deploy failed:%s\n", err.Error())
//		return false
//	}
//
//	fmt.Printf("the txHash is %s\n", txhash.ToHexString())
//	fmt.Printf("contract address is %s\n", addr.ToBase58())
//
//	ctx.LogInfo("=====================invoke name==============================")
//	strarr := []string{"a","b","c"}
//	res,err := PreExecWasmContract(ctx,
//		addr,
//		"testArray",
//		wasmvm.Raw,byte(1),[]interface{}{strarr})
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
//
//	ctx.LogInfo("=====================invoke name==============================")
//	//argarr := []interface{}{[]string{"a","b","c"},[]string{"d","e","f"}}
//	argarr := []interface{}{[]string{"Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT","AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb","500"},[]string{"Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT","AK98G45DhmPXg4TFPG1KjftvkEaHbU8SHM","600"}}
//
//	res,err = PreExecWasmContract(ctx,
//		addr,
//		"testNestedArray",
//		wasmvm.Raw,byte(1),[]interface{}{argarr})
//
//	if err != nil{
//		fmt.Printf("invoke name failed:%s\n",err.Error())
//		return false
//	}
//
//	tmp ,err = res.Result.ToString()
//	fmt.Printf("name is %v\n",tmp)
//
//	ctx.LogInfo("=====================invoke name end==============================")
//
//
//	return true
//
//}