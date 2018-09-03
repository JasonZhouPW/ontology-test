package deploy_invoke

import (
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"io/ioutil"
	"github.com/ontio/ontology-go-sdk/utils"
	"fmt"
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

	tx1, err := ctx.Ont.Rpc.DeploySmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
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
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestInvokeContractPy WaitForGenerateBlock error: %s", err)
		return false
	}

	tx, err := ctx.Ont.Rpc.NewNeoVMSInvokeTransaction(ctx.GetGasPrice(), ctx.GetGasLimit(),codeAddress, []interface{}{"test", 1,2,3,4})
	if err != nil {
		ctx.LogError("TestInvokeContractPy NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}
	err = ctx.Ont.Rpc.SignToTransaction(tx, signer)
	if err != nil {
		ctx.LogError("TestInvokeContractPy SignToTransaction error:%s", err)

		return false
	}


	obj,err:=ctx.Ont.Rpc.PrepareInvokeContract(tx)
	if err != nil {
		ctx.LogError("TestInvokeContractPy PrepareInvokeContract error:%s", err)

		return false
	}
	fmt.Println(obj)





	return true
}
