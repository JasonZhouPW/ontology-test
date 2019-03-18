package deploy_invoke

import (
	"fmt"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"io/ioutil"
	"time"
)

func TestInvokeContractPy(ctx *testframework.TestFrameworkContext) bool {

	avmfile := "test_data/AddTest1.avm"

	code, err := ioutil.ReadFile(avmfile)
	if err != nil {
		return false
	}
	codeHash := common.ToHexString(code)

	codeAddress, _ := utils.GetContractAddress(codeHash)
	fmt.Println(codeAddress)
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestInvokeContractPy GetDefaultAccount error:%s", err)
		return false
	}

	tx1, err := ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		true,
		codeHash,
		"TestDomainSmartContract",
		"1.0",
		"",
		"",
		"",
	)

	if err != nil {
		ctx.LogError("TestInvokeContractPy DeploySmartContract error: %s", err)
	}

	fmt.Println(tx1.ToHexString())

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestInvokeContractPy WaitForGenerateBlock error: %s", err)
		return false
	}

	obj, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{1, 2, 3, 4})

	fmt.Println(obj)
	bs, err := obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract 1 error:%s", err)

		return false
	}

	//bs,err := common.HexToBytes(res)
	//if err != nil{
	//	ctx.LogError("TestLottery PrepareInvokeContract 2 error:%s", err)
	//
	//	return false
	//}
	fmt.Printf("total supply is %d\n", bs)

	return true
}
