package deploy_invoke

import (
	"github.com/ontio/ontology-test/testframework"
	"io/ioutil"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology-go-sdk/utils"
	"time"
	"fmt"
	"bytes"
	"github.com/ontio/ontology/smartcontract/service/neovm"
)

func TestOEP5Py(ctx *testframework.TestFrameworkContext) bool {


	avmfile := "test_data/OEP5Sample.avm"

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
		ctx.LogError("TestOEP5Py GetDefaultAccount error:%s", err)
		return false
	}
	ctx.LogInfo("-------------------deploy start ---------------------------")
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
		ctx.LogError("TestOEP5Py DeploySmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOEP5Py WaitForGenerateBlock error: %s", err)
		return false
	}

	ctx.LogInfo("-------------------deploy end ---------------------------")
	ctx.LogInfo("-------------------call init start -----------------------")
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"init", []interface{}{}})
	if err != nil {
		ctx.LogError("TestOEP5Py init error: %s", err)
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
	ctx.LogInfo("-------------------call init end -----------------------")
	ctx.LogInfo("--------------------testing Name--------------------")
	obj,err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"name", []interface{}{}})

	name ,err := obj.Result.ToString()
	if err != nil{
		ctx.LogError("TestOEP5Py PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("name is %s\n",name)
	ctx.LogInfo("--------------------testing Name end--------------------")

	ctx.LogInfo("--------------------testing symbol--------------------")
	obj,err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"symbol", []interface{}{}})

	symbol ,err := obj.Result.ToString()
	if err != nil{
		ctx.LogError("TestOEP5Py PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("symbol is %s\n",symbol)
	ctx.LogInfo("--------------------testing symbol end--------------------")

	ctx.LogInfo("--------------------testing queryAssetIDByIndex--------------------")
	obj,err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"queryAssetIDByIndex", []interface{}{1}})

	assetID ,err := obj.Result.ToByteArray()
	if err != nil{
		ctx.LogError("TestOEP5Py PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("assetID is %s\n",common.ToHexString(assetID))
	ctx.LogInfo("--------------------testing queryAssetIDByIndex end--------------------")

	ctx.LogInfo("--------------------testing queryAssetCount--------------------")
	obj,err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"queryAssetCount", []interface{}{}})

	assetCount ,err := obj.Result.ToInteger()
	if err != nil{
		ctx.LogError("TestOEP5Py PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("assetCount is %d\n",assetCount.Int64())
	ctx.LogInfo("--------------------testing queryAssetCount end--------------------")

	ctx.LogInfo("--------------------testing buy asset ---------------------------")

	account2,err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
	if err != nil{
		ctx.LogError("get account AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb failed")
		return false
	}

	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"buyAsset", []interface{}{account2.Address[:],assetID}})
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
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestOEP5Py GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestOEP5Py failed invoked exec state return 0")
		return false
	}

	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}

	ctx.LogInfo("--------------------testing buy end---------------------------")
	ctx.LogInfo("--------------------testing ownerOf--------------------")
	obj,err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"ownerOf", []interface{}{assetID}})

	owner ,err := obj.Result.ToByteArray()
	if err != nil{
		ctx.LogError("TestOEP5Py PrepareInvokeContract error:%s", err)

		return false
	}

	tmpaddr,err := common.AddressParseFromBytes(owner)

	fmt.Printf("owner is %s\n",tmpaddr.ToBase58())
	ctx.LogInfo("--------------------testing ownerOf end--------------------")

	ctx.LogInfo("--------------------testing queryAssetByID--------------------")
	obj,err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"queryAssetByID", []interface{}{assetID}})

	bs ,err := obj.Result.ToByteArray()
	if err != nil{
		ctx.LogError("TestOEP5Py PrepareInvokeContract error:%s", err)

		return false
	}

	bf := bytes.NewBuffer(bs)
	stacks,err := neovm.DeserializeStackItem(bf)
	if err != nil{
		ctx.LogError("TestOEP5Py PrepareInvokeContract error:%s", err)

		return false
	}
	smap,err := stacks.GetMap()
	if err != nil{
		ctx.LogError("TestOEP5Py PrepareInvokeContract error:%s", err)

		return false
	}


	id, err:= getMapvalue(smap, "ID").GetByteArray()
	if err != nil{
		ctx.LogError("TestOEP5Py PrepareInvokeContract error:%s", err)
		return false
	}

	fmt.Printf("id is %s\n",common.ToHexString(id))

	namebs, err:= getMapvalue(smap, "Name").GetByteArray()
	if err != nil{
		ctx.LogError("TestOEP5Py PrepareInvokeContract error:%s", err)
		return false
	}
	fmt.Printf("name is %s\n",string(namebs))

	image, err:= getMapvalue(smap, "Image").GetByteArray()
	if err != nil{
		ctx.LogError("TestOEP5Py PrepareInvokeContract error:%s", err)
		return false
	}
	fmt.Printf("images is %s\n",image)

	tp, err:= getMapvalue(smap, "Type").GetByteArray()
	if err != nil{
		ctx.LogError("TestOEP5Py PrepareInvokeContract error:%s", err)
		return false
	}
	fmt.Printf("type is %s\n",tp)

	ctx.LogInfo("--------------------testing queryAssetByID end--------------------")



	ctx.LogInfo("--------------------testing transfer ---------------------------")

	account3,err := ctx.GetAccount("AK98G45DhmPXg4TFPG1KjftvkEaHbU8SHM")
	if err != nil{
		ctx.LogError("get account AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb failed")
		return false
	}

	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"transfer", []interface{}{account3.Address[:],assetID}})
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
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestOEP5Py GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestOEP5Py failed invoked exec state return 0")
		return false
	}
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}

	ctx.LogInfo("--------------------testing transfer end---------------------------")
	ctx.LogInfo("--------------------testing ownerOf--------------------")
	obj,err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"ownerOf", []interface{}{assetID}})

	owner ,err = obj.Result.ToByteArray()
	if err != nil{
		ctx.LogError("TestOEP5Py PrepareInvokeContract error:%s", err)

		return false
	}

	tmpaddr,err = common.AddressParseFromBytes(owner)

	fmt.Printf("owner is %s\n",tmpaddr.ToBase58())
	ctx.LogInfo("--------------------testing ownerOf end--------------------")


	ctx.LogInfo("--------------------testing approve ---------------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account3,
		codeAddress,
		[]interface{}{"approve", []interface{}{account2.Address[:],assetID}})
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
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestOEP5Py GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestOEP5Py failed invoked exec state return 0")
		return false
	}
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}


	ctx.LogInfo("--------------------testing approve end---------------------------")


	ctx.LogInfo("--------------------testing getApproved--------------------")
	obj,err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getApproved", []interface{}{assetID}})

	owner ,err = obj.Result.ToByteArray()
	if err != nil{
		ctx.LogError("TestOEP5Py PrepareInvokeContract error:%s", err)

		return false
	}

	tmpaddr,err = common.AddressParseFromBytes(owner)

	fmt.Printf("approved account: is %s\n",tmpaddr.ToBase58())
	ctx.LogInfo("--------------------testing getApproved end--------------------")

	ctx.LogInfo("--------------------testing takeOwnership  ---------------------------")

	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"takeOwnership", []interface{}{ account2.Address[:], assetID}})
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
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestOEP5Py GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestOEP5Py failed invoked exec state return 0")
		return false
	}
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}

	ctx.LogInfo("--------------------testing transfer from  end---------------------------")


	ctx.LogInfo("--------------------testing ownerOf--------------------")
	obj,err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"ownerOf", []interface{}{assetID}})

	owner ,err = obj.Result.ToByteArray()
	if err != nil{
		ctx.LogError("TestOEP5Py PrepareInvokeContract error:%s", err)

		return false
	}

	tmpaddr,err = common.AddressParseFromBytes(owner)

	fmt.Printf("owner is %s\n",tmpaddr.ToBase58())
	ctx.LogInfo("--------------------testing ownerOf end--------------------")

	ctx.LogInfo("--------------------testing withdraw  ---------------------------")

	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"withdraw", []interface{}{ signer.Address[:]}})
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
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestOEP5Py GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestOEP5Py failed invoked exec state return 0")
		return false
	}
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}

	ctx.LogInfo("--------------------testing withdraw end---------------------------")



	return true
}