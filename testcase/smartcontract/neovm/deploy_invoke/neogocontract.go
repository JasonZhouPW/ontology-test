package deploy_invoke

import (
	"github.com/ontio/ontology-test/testframework"
	"io/ioutil"
	"github.com/ontio/ontology/common"
	types "github.com/ontio/ontology/smartcontract/types"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/account"
	"time"
	"fmt"
)

func TestInvokeNeogoContract(ctx *testframework.TestFrameworkContext) bool {

	path := "test_data"
	file := path + "/" + "test.avm"
	code, err := ioutil.ReadFile(file)
	if err != nil {
		return false
	}
	codeHash := common.ToHexString(code)
	address :=  utils.GetNeoVMContractAddress(codeHash)
	ctx.LogInfo("address is %s\n",address.ToHexString())
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestInvokeNeogoContract GetDefaultAccount error:%s", err)
		return false
	}

	txHash,err := DeployNeogoContract(ctx,signer,codeHash,"TestNeogoContract","1.0")
	if err != nil {
		ctx.LogError("TestInvokeNeogoContract deploy error:%s", err)
		return false
	}
	ctx.LogInfo("TestInvokeNeogoContract deploy TxHash:%x", txHash)

	params := []interface{}{"Register", []interface{}{100, 300}}
	txHash,err = InvokeNeogoContract(ctx,signer,address,params)
	if err != nil {
		ctx.LogError("TestInvokeNeogoContract deploy error:%s", err)
		return false
	}

	notifies, err := ctx.Ont.Rpc.GetSmartContractEvent(txHash)
	if err != nil {
		ctx.LogError("TestInvokeNeogoContract InvokeNeogoContract error:%s", err)
		return false
	}

	if len(notifies) < 1{
		ctx.LogError("TestInvokeNeogoContract InvokeNeogoContract return notifies count error!")
		return false
	}
	ctx.LogInfo("==========TestInvokeNeogoContract InvokeNeogoContract ============")
	for i ,n := range notifies{
		ctx.LogInfo(fmt.Sprintf("notify %d is %v",i, n))
	}
	return true

}

func DeployNeogoContract(ctx *testframework.TestFrameworkContext,signer *account.Account,codeHash string,name string,ver string)  (common.Uint256, error) {
	txHash, err := ctx.Ont.Rpc.DeploySmartContract(
		0, 0,
		signer,
		types.NEOVM,
		true,
		codeHash,
		name,
		ver,
		"",
		"",
		"",
	)

	if err != nil {
		ctx.LogError("TestDeploySmartContract DeploySmartContract error:%s", err)
		return common.Uint256{},err
	}
	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDeploySmartContract WaitForGenerateBlock error:%s", err)
		return common.Uint256{},err
	}
	return txHash,nil
}

func InvokeNeogoContract(ctx *testframework.TestFrameworkContext,signer *account.Account,address common.Address,params []interface{}) (common.Uint256, error){
	txHash, err := ctx.Ont.Rpc.InvokeNeoVMSmartContract(0, 0, signer, 0, address, params)
	if err != nil {
		ctx.LogError("TestInvokeSmartContract InvokeNeoVMSmartContract error:%s", err)
		return common.Uint256{},err
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestInvokeSmartContract WaitForGenerateBlock error:%s", err)
		return common.Uint256{},err
	}

	return txHash,nil

}