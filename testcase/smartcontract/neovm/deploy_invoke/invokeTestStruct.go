package deploy_invoke

import (
	"fmt"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"io/ioutil"
	"time"
)

func TestStructPy(ctx *testframework.TestFrameworkContext) bool {

	account2, _ := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")

	avmfile := "test_data/testStruct.avm"

	code, err := ioutil.ReadFile(avmfile)
	if err != nil {
		return false
	}
	codeHash := common.ToHexString(code)

	codeAddress, _ := utils.GetContractAddress(codeHash)
	fmt.Println("contract address:" + codeAddress.ToBase58())

	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToHexString())
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestStructPy GetDefaultAccount error:%s", err)
		return false
	}

	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		true,
		codeHash,
		"TestStructPy",
		"1.0",
		"",
		"",
		"",
	)

	if err != nil {
		ctx.LogError("TestStructPy DeploySmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestStructPy WaitForGenerateBlock error: %s", err)
		return false
	}

	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"transfer", []interface{}{signer.Address[:], account2.Address[:], 500}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}
	notify := events.Notify[0]
	ctx.LogInfo("%+v", notify)
	//invokeState := notify.States.(string)
	//ctx.LogInfo(invokeState)
	//s,_  :=common.HexToBytes(invokeState)
	//ctx.LogInfo("%s", s)

	return true
}
