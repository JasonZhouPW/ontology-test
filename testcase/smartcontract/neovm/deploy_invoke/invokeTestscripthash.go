package deploy_invoke

import (
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology-go-sdk/utils"
	"time"
	"io/ioutil"
	"fmt"
)

func Testscripthash(ctx *testframework.TestFrameworkContext) bool {


	avmfile := "test_data/toscripthash.avm"

	code, err := ioutil.ReadFile(avmfile)
	if err != nil {
		return false
	}
	codeHash := common.ToHexString(code)

	codeAddress, _ := utils.GetContractAddress(codeHash)

	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToHexString())
	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToBase58())
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestGoBytesEq GetDefaultAccount error:%s", err)
		return false
	}

	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		true,
		codeHash,
		"Testtoscripthash",
		"1.0",
		"",
		"",
		"",
	)

	if err != nil {
		ctx.LogError("Testscripthash DeploySmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Testscripthash WaitForGenerateBlock error: %s", err)
		return false
	}


	obj, err :=ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"test", []interface{}{}})
	if err != nil {
		ctx.LogError("Testscripthash NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	res ,err := obj.Result.ToString()
	if err != nil{
		ctx.LogError("Testscripthash PrepareInvokeContract error:%s", err)

		return false
	}
	bs,err := common.HexToBytes(res)
	if err != nil{
		ctx.LogError("Testscripthash PrepareInvokeContract error:%s", err)

		return false
	}

	add,err:=common.AddressParseFromBytes(bs)
	if err != nil{
		ctx.LogError("TestOEP4Py parse error:%s", err)
	}
	fmt.Println(add.ToBase58())


	return true
}
