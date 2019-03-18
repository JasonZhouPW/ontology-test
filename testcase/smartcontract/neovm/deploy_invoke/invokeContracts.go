package deploy_invoke

import (
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"io/ioutil"
	"time"
)

func TestContractsAPI(ctx *testframework.TestFrameworkContext) bool {

	avmfile := "test_data/contract_api.avm"

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
		ctx.LogError("TestContractsAPI GetDefaultAccount error:%s", err)
		return false
	}
	/*	ctx.LogInfo("-------------------deploy start ---------------------------")
		_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
			signer,
			true,
			codeHash,
			"TestOEP5Py",
			"1.0",
			"",
			"",
			"",
		)

		if err != nil {
			ctx.LogError("TestContractsAPI DeploySmartContract error: %s", err)
		}

		//WaitForGenerateBlock
		_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
		if err != nil {
			ctx.LogError("TestContractsAPI WaitForGenerateBlock error: %s", err)
			return false
		}

		ctx.LogInfo("-------------------deploy end ---------------------------")*/
	ctx.LogInfo("-------------------call destroy start -----------------------")
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"Destroy", []interface{}{}})
	if err != nil {
		ctx.LogError("destroy init error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("destroy WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("destroy GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("destroy failed invoked exec state return 0")
		return false
	}
	ctx.LogInfo("-------------------call destroy end -----------------------")
	return true
}
