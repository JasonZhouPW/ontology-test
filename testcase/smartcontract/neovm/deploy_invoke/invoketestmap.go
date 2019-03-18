package deploy_invoke

import (
	"fmt"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"io/ioutil"
	"time"
)

func TestInvoketestmap(ctx *testframework.TestFrameworkContext) bool {

	avmfile := "test_data/testmap.avm"

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

	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
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

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestInvokeContractPy WaitForGenerateBlock error: %s", err)
		return false
	}

	ctx.LogInfo("--------------------testing add ---------------------------")
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"add", []interface{}{11111111222222222}})
	if err != nil {
		ctx.LogError("TestOEP5Py InvokeNeoVMSmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOEP5Py WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestOEP5Py GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestOEP5Py failed invoked exec state return 0")
		return false
	}

	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)

		state := notify.States.(string)
		ctx.LogInfo(state)
		bs, _ := common.HexToBytes(state)
		res := common.BigIntFromNeoBytes(bs)
		ctx.LogInfo(res.Int64())
	}

	ctx.LogInfo("--------------------testing add end ---------------------------")

	ctx.LogInfo("--------------------testing get--------------------")
	obj, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"get", []interface{}{}})

	balance, err := obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestOEP5Py PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("get is %d\n", balance.Int64())
	ctx.LogInfo("--------------------testing get end--------------------")

	return true
}
