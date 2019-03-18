package deploy_invoke

import (
	"fmt"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"io/ioutil"
	"time"
)

func TestMulitparam(ctx *testframework.TestFrameworkContext) bool {

	avmfile := "test_data/testMultiParam.avm"

	code, err := ioutil.ReadFile(avmfile)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	codeHash := common.ToHexString(code)

	codeAddress, _ := utils.GetContractAddress(codeHash)

	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToHexString())
	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToBase58())
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestMulitparam GetDefaultAccount error:%s", err)
		return false
	}

	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		true,
		codeHash,
		"TestMulitparam",
		"1.0",
		"",
		"",
		"",
	)

	if err != nil {
		ctx.LogError("TestMulitparam DeploySmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestMulitparam WaitForGenerateBlock error: %s", err)
		return false
	}

	ctx.LogInfo("============test  start===========")
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"test", []interface{}{"a", "b", "c", "d"}})
	if err != nil {
		ctx.LogError("TestMulitparam InvokeNeoVMSmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestMulitparam WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestMulitparam GetSmartContractEvent error:%s", err)
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}

	ctx.LogInfo("============test  end===========")

	return true
}
