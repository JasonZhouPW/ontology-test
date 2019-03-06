package wasmvm
//
//import (
//	"github.com/ontio/ontology-test/testframework"
//	"fmt"
//	"neo-go-compiler-bak/serialization"
//	"bytes"
//)
//
//func TestNewOEP4(ctx *testframework.TestFrameworkContext) bool {
//	testFile := filePath + "/" + "helloworld.wasm"
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
//
//
//	ctx.LogInfo("=====================invoke name==============================")
//	res,err := PreExecWasmContract(ctx,
//		addr,
//		"hello",
//		byte(1),[]interface{}{})
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
//	fmt.Printf("name is %v\n",tmp)
//
//	ctx.LogInfo("=====================invoke name end==============================")
//
//	return true
//
//}