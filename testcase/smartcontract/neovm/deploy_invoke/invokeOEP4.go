package deploy_invoke

import (
	"github.com/ontio/ontology-test/testframework"
	"io/ioutil"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology-go-sdk/utils"
	"fmt"
	"time"
)

func TestOEP4Py(ctx *testframework.TestFrameworkContext) bool {


	avmfile := "test_data/OEP4Sample.avm"

	code, err := ioutil.ReadFile(avmfile)
	if err != nil {
		return false
	}
	codeHash := common.ToHexString(code)

	codeAddress, _ := utils.GetContractAddress(codeHash)

	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToHexString())
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestOEP4Py GetDefaultAccount error:%s", err)
		return false
	}

	_, err = ctx.Ont.Rpc.DeploySmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		true,
		codeHash,
		"TestOEP4Py",
		"1.0",
		"",
		"",
		"",
	)

	if err != nil {
		ctx.LogError("TestOEP4Py DeploySmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOEP4Py WaitForGenerateBlock error: %s", err)
		return false
	}

	//-------------------init-------------------
	ctx.LogInfo("--------------------testing init--------------------")
	txHash, err := ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"init", []interface{}{}})
	if err != nil {
		ctx.LogError("TestOEP4Py init error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOEP4Py WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err := ctx.Ont.Rpc.GetSmartContractEvent(txHash)
	if err != nil {
		ctx.LogError("TestOEP4Py GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestOEP4Py failed invoked exec state return 0")
		return false
	}

	ctx.LogInfo("--------------------testing init end--------------------")


	ctx.LogInfo("--------------------testing totalSupply--------------------")
	tx, err := ctx.Ont.Rpc.NewNeoVMSInvokeTransaction(ctx.GetGasPrice(), ctx.GetGasLimit(),codeAddress, []interface{}{"totalSupply", []interface{}{}})
	if err != nil {
		ctx.LogError("TestOEP4Py NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}
	err = ctx.Ont.Rpc.SignToTransaction(tx, signer)
	if err != nil {
		ctx.LogError("TestOEP4Py SignToTransaction error:%s", err)

		return false
	}


	obj,err:=ctx.Ont.Rpc.PrepareInvokeContract(tx)
	if err != nil {
		ctx.LogError("TestOEP4Py PrepareInvokeContract error:%s", err)

		return false
	}

	bs,err:= common.HexToBytes(obj.Result.(string))
	if err != nil{
		ctx.LogError("TestOEP4Py PrepareInvokeContract error:%s", err)

		return false
	}
	totalSupply := common.BigIntFromNeoBytes(bs)

	fmt.Printf("total supply is %d\n",totalSupply.Int64())
	ctx.LogInfo("--------------------testing totalSupply end--------------------")

	ctx.LogInfo("--------------------testing name--------------------")
	tx, err = ctx.Ont.Rpc.NewNeoVMSInvokeTransaction(ctx.GetGasPrice(), ctx.GetGasLimit(),codeAddress, []interface{}{"name", []interface{}{}})
	if err != nil {
		ctx.LogError("TestOEP4Py NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}
	err = ctx.Ont.Rpc.SignToTransaction(tx, signer)
	if err != nil {
		ctx.LogError("TestOEP4Py SignToTransaction error:%s", err)

		return false
	}


	obj,err=ctx.Ont.Rpc.PrepareInvokeContract(tx)
	if err != nil {
		ctx.LogError("TestOEP4Py PrepareInvokeContract error:%s", err)

		return false
	}

	bs,err= common.HexToBytes(obj.Result.(string))
	if err != nil{
		ctx.LogError("TestOEP4Py PrepareInvokeContract error:%s", err)

		return false
	}
	name := string(bs)
	//
	fmt.Printf("name is %s\n",name)
	ctx.LogInfo("--------------------testing name end--------------------")

	ctx.LogInfo("--------------------testing symbol--------------------")
	tx, err = ctx.Ont.Rpc.NewNeoVMSInvokeTransaction(ctx.GetGasPrice(), ctx.GetGasLimit(),codeAddress, []interface{}{"symbol", []interface{}{}})
	if err != nil {
		ctx.LogError("TestOEP4Py NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}
	err = ctx.Ont.Rpc.SignToTransaction(tx, signer)
	if err != nil {
		ctx.LogError("TestOEP4Py SignToTransaction error:%s", err)

		return false
	}


	obj,err=ctx.Ont.Rpc.PrepareInvokeContract(tx)
	if err != nil {
		ctx.LogError("TestOEP4Py PrepareInvokeContract error:%s", err)

		return false
	}

	bs,err= common.HexToBytes(obj.Result.(string))
	if err != nil{
		ctx.LogError("TestOEP4Py PrepareInvokeContract error:%s", err)

		return false
	}
	symbol := string(bs)
	//
	fmt.Printf("symbol is %s\n",symbol)
	ctx.LogInfo("--------------------testing symbol end--------------------")

	ctx.LogInfo("--------------------testing decimal--------------------")
	tx, err = ctx.Ont.Rpc.NewNeoVMSInvokeTransaction(ctx.GetGasPrice(), ctx.GetGasLimit(),codeAddress, []interface{}{"decimal", []interface{}{}})
	if err != nil {
		ctx.LogError("TestOEP4Py NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}
	err = ctx.Ont.Rpc.SignToTransaction(tx, signer)
	if err != nil {
		ctx.LogError("TestOEP4Py SignToTransaction error:%s", err)

		return false
	}


	obj,err=ctx.Ont.Rpc.PrepareInvokeContract(tx)
	if err != nil {
		ctx.LogError("TestOEP4Py PrepareInvokeContract error:%s", err)

		return false
	}

	bs,err= common.HexToBytes(obj.Result.(string))
	if err != nil{
		ctx.LogError("TestOEP4Py PrepareInvokeContract error:%s", err)

		return false
	}
	decimal := common.BigIntFromNeoBytes(bs)
	//
	fmt.Printf("decimal is %d\n",decimal.Int64())
	ctx.LogInfo("--------------------testing decimal end--------------------")


	ctx.LogInfo("--------------------testing balanceOf owner--------------------")
	tx, err = ctx.Ont.Rpc.NewNeoVMSInvokeTransaction(ctx.GetGasPrice(), ctx.GetGasLimit(),codeAddress, []interface{}{"balanceOf", []interface{}{signer.Address[:]}})
	if err != nil {
		ctx.LogError("TestOEP4Py NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}
	err = ctx.Ont.Rpc.SignToTransaction(tx, signer)
	if err != nil {
		ctx.LogError("TestOEP4Py SignToTransaction error:%s", err)

		return false
	}


	obj,err=ctx.Ont.Rpc.PrepareInvokeContract(tx)
	if err != nil {
		ctx.LogError("TestOEP4Py PrepareInvokeContract error:%s", err)

		return false
	}

	bs,err= common.HexToBytes(obj.Result.(string))
	if err != nil{
		ctx.LogError("TestOEP4Py PrepareInvokeContract error:%s", err)

		return false
	}
	balance := common.BigIntFromNeoBytes(bs)
	//
	fmt.Printf("balance is %d\n",balance.Int64())
	ctx.LogInfo("--------------------testing balanceOf owner end--------------------")


	ctx.LogInfo("--------------------testing transfer ---------------------------")

	account2,err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
	if err != nil{
		ctx.LogError("get account AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb failed")
		return false
	}

	txHash, err = ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"transfer", []interface{}{signer.Address[:], account2.Address[:],50000000000}})
	if err != nil {
		ctx.LogError("TestOEP4Py InvokeNeoVMSmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOEP4Py WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.Rpc.GetSmartContractEvent(txHash)
	if err != nil {
		ctx.LogError("TestOEP4Py GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestOEP4Py failed invoked exec state return 0")
		return false
	}
	notify := events.Notify[0]
	ctx.LogInfo("%+v", notify)
	invokeState := notify.States.([]interface{})
	ctx.LogInfo(invokeState)

	method,_  :=common.HexToBytes(invokeState[0].(string))
	addFromTmp,_:= common.HexToBytes(invokeState[1].(string))
	addFrom,_ := common.AddressParseFromBytes(addFromTmp)

	addToTmp,_:= common.HexToBytes(invokeState[2].(string))
	addTo,_ := common.AddressParseFromBytes(addToTmp)
	tmp,_:= common.HexToBytes(invokeState[3].(string))
	amount := common.BigIntFromNeoBytes(tmp)
	ctx.LogInfo("states[method:%s,from:%s,to:%s,value:%d]", method,addFrom.ToBase58(),addTo.ToBase58(),amount.Int64())


	ctx.LogInfo("--------------------testing transfer end---------------------------")


	ctx.LogInfo("--------------------testing balanceOf owner--------------------")
	tx, err = ctx.Ont.Rpc.NewNeoVMSInvokeTransaction(ctx.GetGasPrice(), ctx.GetGasLimit(),codeAddress, []interface{}{"balanceOf", []interface{}{signer.Address[:]}})
	if err != nil {
		ctx.LogError("TestOEP4Py NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}
	err = ctx.Ont.Rpc.SignToTransaction(tx, signer)
	if err != nil {
		ctx.LogError("TestOEP4Py SignToTransaction error:%s", err)

		return false
	}


	obj,err=ctx.Ont.Rpc.PrepareInvokeContract(tx)
	if err != nil {
		ctx.LogError("TestOEP4Py PrepareInvokeContract error:%s", err)

		return false
	}

	bs,err= common.HexToBytes(obj.Result.(string))
	if err != nil{
		ctx.LogError("TestOEP4Py PrepareInvokeContract error:%s", err)

		return false
	}
	balance = common.BigIntFromNeoBytes(bs)
	//
	fmt.Printf("balance is %d\n",balance.Int64())
	ctx.LogInfo("--------------------testing balanceOf owner end--------------------")


	ctx.LogInfo("--------------------testing balanceOf acct2--------------------")
	tx, err = ctx.Ont.Rpc.NewNeoVMSInvokeTransaction(ctx.GetGasPrice(), ctx.GetGasLimit(),codeAddress, []interface{}{"balanceOf", []interface{}{account2.Address[:]}})
	if err != nil {
		ctx.LogError("TestOEP4Py NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}
	err = ctx.Ont.Rpc.SignToTransaction(tx, signer)
	if err != nil {
		ctx.LogError("TestOEP4Py SignToTransaction error:%s", err)

		return false
	}


	obj,err=ctx.Ont.Rpc.PrepareInvokeContract(tx)
	if err != nil {
		ctx.LogError("TestOEP4Py PrepareInvokeContract error:%s", err)

		return false
	}

	bs,err= common.HexToBytes(obj.Result.(string))
	if err != nil{
		ctx.LogError("TestOEP4Py PrepareInvokeContract error:%s", err)

		return false
	}
	balance = common.BigIntFromNeoBytes(bs)
	//
	fmt.Printf("balance is %d\n",balance.Int64())
	ctx.LogInfo("--------------------testing balanceOf acct2 end--------------------")



	ctx.LogInfo("--------------------testing approve ---------------------------")

	txHash, err = ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"approve", []interface{}{signer.Address[:], account2.Address[:],60000000000}})
	if err != nil {
		ctx.LogError("TestOEP4Py InvokeNeoVMSmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOEP4Py WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.Rpc.GetSmartContractEvent(txHash)
	if err != nil {
		ctx.LogError("TestOEP4Py GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestOEP4Py failed invoked exec state return 0")
		return false
	}
	notify = events.Notify[0]
	ctx.LogInfo("%+v", notify)
	invokeState = notify.States.([]interface{})
	ctx.LogInfo(invokeState)

	method,_  =common.HexToBytes(invokeState[0].(string))
	addFromTmp,_= common.HexToBytes(invokeState[1].(string))
	addFrom,_ = common.AddressParseFromBytes(addFromTmp)

	addToTmp,_= common.HexToBytes(invokeState[2].(string))
	addTo,_ = common.AddressParseFromBytes(addToTmp)
	tmp,_= common.HexToBytes(invokeState[3].(string))
	amount = common.BigIntFromNeoBytes(tmp)
	ctx.LogInfo("states[method:%s,from:%s,to:%s,value:%d]", method,addFrom.ToBase58(),addTo.ToBase58(),amount.Int64())

	ctx.LogInfo("--------------------testing approve end---------------------------")



	ctx.LogInfo("--------------------testing allowance signer--------------------")
	tx, err = ctx.Ont.Rpc.NewNeoVMSInvokeTransaction(ctx.GetGasPrice(), ctx.GetGasLimit(),codeAddress, []interface{}{"allowance", []interface{}{signer.Address[:],account2.Address[:]}})
	if err != nil {
		ctx.LogError("TestOEP4Py NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}
	err = ctx.Ont.Rpc.SignToTransaction(tx, signer)
	if err != nil {
		ctx.LogError("TestOEP4Py SignToTransaction error:%s", err)

		return false
	}


	obj,err=ctx.Ont.Rpc.PrepareInvokeContract(tx)
	if err != nil {
		ctx.LogError("TestOEP4Py PrepareInvokeContract error:%s", err)

		return false
	}

	bs,err= common.HexToBytes(obj.Result.(string))
	if err != nil{
		ctx.LogError("TestOEP4Py PrepareInvokeContract error:%s", err)

		return false
	}
	balance = common.BigIntFromNeoBytes(bs)
	//
	fmt.Printf("allowance is %d\n",balance.Int64())
	ctx.LogInfo("--------------------testing allowance signer end--------------------")


	ctx.LogInfo("--------------------testing transfer from ---------------------------")

	txHash, err = ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"transferFrom", []interface{}{ account2.Address[:],signer.Address[:],account2.Address[:],30000000000}})
	if err != nil {
		ctx.LogError("TestOEP4Py InvokeNeoVMSmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOEP4Py WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.Rpc.GetSmartContractEvent(txHash)
	if err != nil {
		ctx.LogError("TestOEP4Py GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestOEP4Py failed invoked exec state return 0")
		return false
	}
	notify = events.Notify[0]
	ctx.LogInfo("%+v", notify)
	invokeState = notify.States.([]interface{})
	ctx.LogInfo(invokeState)

	method,_  =common.HexToBytes(invokeState[0].(string))
	addFromTmp,_= common.HexToBytes(invokeState[1].(string))
	addFrom,_ = common.AddressParseFromBytes(addFromTmp)

	addToTmp,_= common.HexToBytes(invokeState[2].(string))
	addTo,_ = common.AddressParseFromBytes(addToTmp)
	tmp,_= common.HexToBytes(invokeState[3].(string))
	amount = common.BigIntFromNeoBytes(tmp)
	ctx.LogInfo("states[method:%s,from:%s,to:%s,value:%d]", method,addFrom.ToBase58(),addTo.ToBase58(),amount.Int64())

	ctx.LogInfo("--------------------testing transfer from  end---------------------------")




	return true
}
