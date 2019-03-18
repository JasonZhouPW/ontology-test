package deploy_invoke

import (
	"time"

	"fmt"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"io/ioutil"
)

func TestDomainSmartContractPy(ctx *testframework.TestFrameworkContext) bool {

	avmfile := "test_data/DomainAuction.avm"

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
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
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
		ctx.LogError("TestDomainSmartContract DeploySmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}

	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"register", []interface{}{signer.Address[:], []byte("www.g.com")}})
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
	invokeState := notify.States.(string)
	ctx.LogInfo(invokeState)
	s, _ := common.HexToBytes(invokeState)
	ctx.LogInfo("%s", s)

	ctx.LogInfo("==============query begin=================")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"query", []interface{}{[]byte("www.g.com")}})
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
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}
	notify = events.Notify[0]
	ctx.LogInfo("%+v", notify)
	invokeState = notify.States.(string)
	ctx.LogInfo(invokeState)
	addr, _ := common.AddressParseFromBytes(s)
	ctx.LogInfo("%s", addr.ToBase58())
	ctx.LogInfo("==============query end=================")

	ctx.LogInfo("==============sell begin====================")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"sell", []interface{}{signer.Address[:], []byte("www.g.com"), 500}})
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
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}
	notify = events.Notify[0]
	ctx.LogInfo("%+v", notify)
	invokeState = notify.States.(string)
	ctx.LogInfo(invokeState)
	s, _ = common.HexToBytes(invokeState)
	ctx.LogInfo("%s", s)
	ctx.LogInfo("==============sell end====================")

	ctx.LogInfo("==============query begin=================")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"query", []interface{}{[]byte("www.g.com")}})
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
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}
	notify = events.Notify[0]
	ctx.LogInfo("%+v", notify)
	invokeState = notify.States.(string)
	ctx.LogInfo(invokeState)
	s, _ = common.HexToBytes(invokeState)
	addr, _ = common.AddressParseFromBytes(s)
	ctx.LogInfo("%s", addr.ToBase58())
	ctx.LogInfo("==============query end=================")

	ctx.LogInfo("==============buy begin=================")
	account2, err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
	if err != nil {
		ctx.LogError("get account AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb failed")
		return false
	}
	//txHash, err = ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
	//	account2,
	//	codeAddress,
	//	[]interface{}{"transfer", []interface{}{account2.Address[:],500}})
	//if err != nil {
	//	ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
	//}
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"buy", []interface{}{account2.Address[:], []byte("www.g.com"), 500}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}

	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}

	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo(notify)

		//invokeState := notify.States.(string)
		//ctx.LogInfo(invokeState)
		//s,_  :=common.HexToBytes(invokeState)
		//ctx.LogInfo("%s", s)
	}

	ctx.LogInfo("==============buy end=================")

	ctx.LogInfo("==============done start=================")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"done", []interface{}{signer.Address[:], []byte("www.g.com")}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
	}
	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}

	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}

	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo(notify)

		switch notify.States.(type) {
		case []interface{}:
			ctx.LogInfo(notify.States)
		case interface{}:
			invokeState := notify.States.(string)
			ctx.LogInfo(invokeState)
			s, _ := common.HexToBytes(invokeState)
			ctx.LogInfo("%s", s)
		}

	}

	ctx.LogInfo("==============done end=================")

	return true
}
