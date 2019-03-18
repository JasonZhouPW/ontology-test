package deploy_invoke

import (
	"fmt"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"io/ioutil"
	"time"
)

func TestJsonMarshal(ctx *testframework.TestFrameworkContext) bool {

	avmfile := "test_data/testJsonMarshal.avm"

	code, err := ioutil.ReadFile(avmfile)
	if err != nil {
		return false
	}
	codeHash := common.ToHexString(code)

	codeAddress, _ := utils.GetContractAddress(codeHash)

	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToHexString())
	ctx.LogInfo("=====CodeAddress base58===%s", codeAddress.ToBase58())

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestJsonMarshal GetDefaultAccount error:%s", err)
		return false
	}

	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		true,
		codeHash,
		"TestJsonMarshal",
		"1.0",
		"",
		"",
		"",
	)

	if err != nil {
		ctx.LogError("TestJsonMarshal DeploySmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestJsonMarshal WaitForGenerateBlock error: %s", err)
		return false
	}

	ctx.LogInfo("--------------------testing jsonmarshal--------------------")
	obj, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"testJson", []interface{}{}})

	json, err := obj.Result.ToString()
	if err != nil {
		ctx.LogError("jsonmarshal PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("jsonmarshal is %s\n", json)
	ctx.LogInfo("--------------------testing jsonmarshal end--------------------")

	return true
}
