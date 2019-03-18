package deploy_invoke

import (
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"io/ioutil"
	"time"
)

func TestGoBytesEq(ctx *testframework.TestFrameworkContext) bool {

	avmfile := "test_data/testBytesEq.avm"

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
		"TestGoBytesEq",
		"1.0",
		"",
		"",
		"",
	)

	if err != nil {
		ctx.LogError("TestGoBytesEq DeploySmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGoBytesEq WaitForGenerateBlock error: %s", err)
		return false
	}

	ctx.LogInfo("============test eq start===========")
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"isEq", []interface{}{signer.Address[:], signer.Address[:]}})
	if err != nil {
		ctx.LogError("TestOEP4Py InvokeNeoVMSmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOEP4Py WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestOEP4Py GetSmartContractEvent error:%s", err)
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}

	ctx.LogInfo("============test eq end===========")
	ctx.LogInfo("============test eq2 start===========")
	account2, err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
	if err != nil {
		ctx.LogError("get account AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb failed")
		return false
	}
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"isEq", []interface{}{signer.Address[:], account2.Address[:]}})
	if err != nil {
		ctx.LogError("TestOEP4Py InvokeNeoVMSmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOEP4Py WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestOEP4Py GetSmartContractEvent error:%s", err)
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}

	ctx.LogInfo("============test eq2 end===========")

	return true
}
