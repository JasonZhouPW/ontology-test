package deploy_invoke

import (
	"github.com/ontio/ontology-test/testframework"
	"io/ioutil"
	"fmt"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology-go-sdk/utils"
	"time"
	"github.com/ontio/ontology-go-sdk"
)

func TestONS(ctx *testframework.TestFrameworkContext) bool {

	avmfile := "test_data/ONS.avm"

	code, err := ioutil.ReadFile(avmfile)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	codeHash := common.ToHexString(code)

	codeAddress, _ := utils.GetContractAddress(codeHash)

	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToHexString())
	ctx.LogInfo("=====CodeAddress base58===%s", codeAddress.ToBase58())
	signer, err := ctx.GetDefaultAccount()
	//account4, err := ctx.GetAccount("ALerVnMj3eNk9xe8BnQJtoWvwGmY3x4KMi")
	//accountOrg,_ := ctx.GetAccount("AUyHN4iVcVFAKxAa18EryG9zNAB4tYXUt6")

	ontidPrefix := "did:ont:"

	ctx.LogInfo("=================Deploy===============================")

	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
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
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOEP4Py WaitForGenerateBlock error: %s", err)
		return false
	}
	ctx.LogInfo("=================register top domain===============================")

	flag := false
	if flag {
		ctx.LogInfo("--------------------testing registerDomain--------------------")
		timeont := time.Now().Unix() + 3600
		txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"registerDomain", []interface{}{"ont",ontidPrefix+signer.Address.ToBase58(),1,timeont}})
		if err != nil {
		ctx.LogError("registerDomain error: %s", err)
		}

		//WaitForGenerateBlock
		_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
		if err != nil {
		ctx.LogError(" WaitForGenerateBlock error: %s", err)
			return false
		}

		//GetEventLog, to check the result of invoke
		events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
		if err != nil {
		ctx.LogError("registerDomain GetSmartContractEvent error:%s", err)
			return false
		}
		if events.State == 0 {
		ctx.LogError("registerDomain failed invoked exec state return 0")
			return false
		}
		for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
		}
		ctx.LogInfo("--------------------testing registerDomain end--------------------")
	}


	ctx.LogInfo("--------------------testing ownerOf ont--------------------")
	obj, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"ownerOf", []interface{}{"ont"}})

	owner, err := obj.Result.ToString()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("owner of ont is %s\n", owner)
	ctx.LogInfo("--------------------testing ownerOf ont end--------------------")
	ctx.LogInfo("--------------------testing validTo ont--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"validTo", []interface{}{"ont"}})

	validto, err := obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("validTo of ont is %d\n", validto)
	ctx.LogInfo("--------------------testing ownerOf ont end--------------------")
	ctx.LogInfo("--------------------testing isDomainValid ont--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"isDomainValid", []interface{}{"ont"}})

	isDomainValid, err := obj.Result.ToBool()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("isDomainValid of ont is %v\n", isDomainValid)
	ctx.LogInfo("--------------------testing ownerOf ont end--------------------")
	flag = true
	if flag {
		ctx.LogInfo("--------------------testing updateValidPeriod--------------------")
		txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
			signer,
			codeAddress,
			[]interface{}{"updateValidPeriod", []interface{}{"ont",1,-1}})
		if err != nil {
			ctx.LogError("updateValidPeriod error: %s", err)
		}

		//WaitForGenerateBlock
		_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
		if err != nil {
			ctx.LogError(" WaitForGenerateBlock error: %s", err)
			return false
		}

		//GetEventLog, to check the result of invoke
		events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
		if err != nil {
			ctx.LogError("registerDomain GetSmartContractEvent error:%s", err)
			return false
		}
		if events.State == 0 {
			ctx.LogError("registerDomain failed invoked exec state return 0")
			return false
		}
		for _, notify := range events.Notify {
			ctx.LogInfo("%+v", notify)
		}
		ctx.LogInfo("--------------------testing updateValidPeriod end--------------------")
	}

	flag = false
	if flag{
		ctx.LogInfo("--------------------testing reg ontid--------------------")
		controller := &ontology_go_sdk.Controller{
			ID:"did:ont:"+signer.Address.ToBase58(),
			PublicKey:signer.PublicKey,
			PrivateKey:signer.PrivateKey,
			SigScheme:signer.SigScheme,
		}
		txhash, err := ctx.Ont.Native.OntId.RegIDWithPublicKey(ctx.GetGasPrice(), ctx.GetGasLimit(),signer,"did:ont:"+signer.Address.ToBase58(),controller)
		if err != nil {
			ctx.LogError("updateValidPeriod error: %s", err)
		}

		//WaitForGenerateBlock
		_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
		if err != nil {
			ctx.LogError(" WaitForGenerateBlock error: %s", err)
			return false
		}
		events, err := ctx.Ont.GetSmartContractEvent(txhash.ToHexString())
		if err != nil {
			ctx.LogError("registerDomain GetSmartContractEvent error:%s", err)
			return false
		}
		if events.State == 0 {
			ctx.LogError("registerDomain failed invoked exec state return 0")
			return false
		}
		for _, notify := range events.Notify {
			ctx.LogInfo("%+v", notify)
		}
		ctx.LogInfo("--------------------testing reg ontid end--------------------")

	}



	ctx.LogInfo("--------------------testing bindvalue--------------------")
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"bindValue", []interface{}{"ont",1,"1","somevalue"}})
	if err != nil {
		ctx.LogError("updateValidPeriod error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError(" WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("registerDomain GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("registerDomain failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing bindvalue end--------------------")
	ctx.LogInfo("--------------------testing valueOf ont--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"valueOf", []interface{}{"ont"}})

	value, err := obj.Result.ToString()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("value of ont is %s\n", value)
	ctx.LogInfo("--------------------testing valueOf ont end--------------------")


	ctx.LogInfo("--------------------testing subdomain test start--------------------")
	account2, _ := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")

	flag = false
	if flag {
		ctx.LogInfo("--------------------testing registerDomain--------------------")
		timeont := time.Now().Unix() + 3600
		txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
			signer,
			codeAddress,
			[]interface{}{"registerDomain", []interface{}{"test.ont",ontidPrefix+account2.Address.ToBase58(),1,timeont}})
		if err != nil {
			ctx.LogError("registerDomain error: %s", err)
		}

		//WaitForGenerateBlock
		_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
		if err != nil {
			ctx.LogError(" WaitForGenerateBlock error: %s", err)
			return false
		}

		//GetEventLog, to check the result of invoke
		events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
		if err != nil {
			ctx.LogError("registerDomain GetSmartContractEvent error:%s", err)
			return false
		}
		if events.State == 0 {
			ctx.LogError("registerDomain failed invoked exec state return 0")
			return false
		}
		for _, notify := range events.Notify {
			ctx.LogInfo("%+v", notify)
		}
		ctx.LogInfo("--------------------testing registerDomain end--------------------")
	}

	ctx.LogInfo("--------------------testing subdomain test end--------------------")
	ctx.LogInfo("--------------------testing ownerOf ont--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"ownerOf", []interface{}{"test.ont"}})

	owner, err = obj.Result.ToString()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("owner of test.ont is %s\n", owner)
	ctx.LogInfo("--------------------testing ownerOf ont end--------------------")
	ctx.LogInfo("--------------------testing validTo ont--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"validTo", []interface{}{"test.ont"}})

	validto, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("validTo of test.ont is %d\n", validto)
	ctx.LogInfo("--------------------testing ownerOf test.ont end--------------------")
	ctx.LogInfo("--------------------testing isDomainValid test.ont--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"isDomainValid", []interface{}{"test.ont"}})

	isDomainValid, err = obj.Result.ToBool()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("isDomainValid of ont is %v\n", isDomainValid)
	ctx.LogInfo("--------------------testing isDomainValid ont end--------------------")
	flag = true
	if flag {
		ctx.LogInfo("--------------------testing updateValidPeriod--------------------")
		txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
			signer,
			codeAddress,
			[]interface{}{"updateValidPeriod", []interface{}{"test.ont",1,-1}})
		if err != nil {
			ctx.LogError("updateValidPeriod error: %s", err)
		}

		//WaitForGenerateBlock
		_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
		if err != nil {
			ctx.LogError(" WaitForGenerateBlock error: %s", err)
			return false
		}

		//GetEventLog, to check the result of invoke
		events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
		if err != nil {
			ctx.LogError("registerDomain GetSmartContractEvent error:%s", err)
			return false
		}
		if events.State == 0 {
			ctx.LogError("registerDomain failed invoked exec state return 0")
			return false
		}
		for _, notify := range events.Notify {
			ctx.LogInfo("%+v", notify)
		}
		ctx.LogInfo("--------------------testing updateValidPeriod end--------------------")
	}
	ctx.LogInfo("--------------------testing validTo ont--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"validTo", []interface{}{"test.ont"}})

	validto, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("validTo of test.ont is %d\n", validto)
	ctx.LogInfo("--------------------testing ownerOf test.ont end--------------------")
	flag = false
	if flag{
		ctx.LogInfo("--------------------testing reg ontid--------------------")
		controller := &ontology_go_sdk.Controller{
			ID:"did:ont:"+account2.Address.ToBase58(),
			PublicKey:account2.PublicKey,
			PrivateKey:account2.PrivateKey,
			SigScheme:account2.SigScheme,
		}
		txhash, err := ctx.Ont.Native.OntId.RegIDWithPublicKey(ctx.GetGasPrice(), ctx.GetGasLimit(),account2,"did:ont:"+account2.Address.ToBase58(),controller)
		if err != nil {
			ctx.LogError("updateValidPeriod error: %s", err)
		}

		//WaitForGenerateBlock
		_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
		if err != nil {
			ctx.LogError(" WaitForGenerateBlock error: %s", err)
			return false
		}
		events, err := ctx.Ont.GetSmartContractEvent(txhash.ToHexString())
		if err != nil {
			ctx.LogError("registerDomain GetSmartContractEvent error:%s", err)
			return false
		}
		if events.State == 0 {
			ctx.LogError("registerDomain failed invoked exec state return 0")
			return false
		}
		for _, notify := range events.Notify {
			ctx.LogInfo("%+v", notify)
		}
		ctx.LogInfo("--------------------testing reg ontid end--------------------")

	}
	ctx.LogInfo("--------------------testing bindvalue test.ont--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"bindValue", []interface{}{"test.ont",1,"1","somevalue_test"}})
	if err != nil {
		ctx.LogError("updateValidPeriod error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError(" WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("registerDomain GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("registerDomain failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing bindvalue test.ont end--------------------")
	ctx.LogInfo("--------------------testing valueOf ont--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"valueOf", []interface{}{"test.ont"}})

	value, err = obj.Result.ToString()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("value of ont is %s\n", value)
	ctx.LogInfo("--------------------testing valueOf ont end--------------------")

	ctx.LogInfo("--------------------testing 2nd level sub domain--------------------")
	account4, err := ctx.GetAccount("ALerVnMj3eNk9xe8BnQJtoWvwGmY3x4KMi")

	flag = false
	if flag {
		ctx.LogInfo("--------------------testing registerDomain--------------------")
		timeont := time.Now().Unix() + 3600
		txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
			account2,
			codeAddress,
			[]interface{}{"registerDomain", []interface{}{"abc.test.ont",ontidPrefix+account4.Address.ToBase58(),1,timeont}})
		if err != nil {
			ctx.LogError("registerDomain error: %s", err)
		}

		//WaitForGenerateBlock
		_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
		if err != nil {
			ctx.LogError(" WaitForGenerateBlock error: %s", err)
			return false
		}

		//GetEventLog, to check the result of invoke
		events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
		if err != nil {
			ctx.LogError("registerDomain GetSmartContractEvent error:%s", err)
			return false
		}
		if events.State == 0 {
			ctx.LogError("registerDomain failed invoked exec state return 0")
			return false
		}
		for _, notify := range events.Notify {
			ctx.LogInfo("%+v", notify)
		}
		ctx.LogInfo("--------------------testing registerDomain end--------------------")
	}
	ctx.LogInfo("--------------------testing 2nd level sub domain end--------------------")
	ctx.LogInfo("--------------------testing ownerOf abc.test.ont--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"ownerOf", []interface{}{"abc.test.ont"}})

	owner, err = obj.Result.ToString()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("owner of test.ont is %s\n", owner)
	ctx.LogInfo("--------------------testing ownerOf abc.test.ont end--------------------")
	ctx.LogInfo("--------------------testing validTo abc.test.ont--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"validTo", []interface{}{"abc.test.ont"}})

	validto, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("validTo of test.ont is %d\n", validto)
	ctx.LogInfo("--------------------testing ownerOf abc.test.ont end--------------------")
	ctx.LogInfo("--------------------testing isDomainValid abc.test.ont--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"isDomainValid", []interface{}{"abc.test.ont"}})

	isDomainValid, err = obj.Result.ToBool()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("isDomainValid of ont is %v\n", isDomainValid)
	ctx.LogInfo("--------------------testing isDomainValid ont end--------------------")
	flag = false
	if flag{
		ctx.LogInfo("--------------------testing reg ontid--------------------")
		controller := &ontology_go_sdk.Controller{
			ID:"did:ont:"+account4.Address.ToBase58(),
			PublicKey:account4.PublicKey,
			PrivateKey:account4.PrivateKey,
			SigScheme:account4.SigScheme,
		}
		txhash, err := ctx.Ont.Native.OntId.RegIDWithPublicKey(ctx.GetGasPrice(), ctx.GetGasLimit(),account4,"did:ont:"+account4.Address.ToBase58(),controller)
		if err != nil {
			ctx.LogError("updateValidPeriod error: %s", err)
		}

		//WaitForGenerateBlock
		_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
		if err != nil {
			ctx.LogError(" WaitForGenerateBlock error: %s", err)
			return false
		}
		events, err := ctx.Ont.GetSmartContractEvent(txhash.ToHexString())
		if err != nil {
			ctx.LogError("registerDomain GetSmartContractEvent error:%s", err)
			return false
		}
		if events.State == 0 {
			ctx.LogError("registerDomain failed invoked exec state return 0")
			return false
		}
		for _, notify := range events.Notify {
			ctx.LogInfo("%+v", notify)
		}
		ctx.LogInfo("--------------------testing reg ontid end--------------------")

	}
	ctx.LogInfo("--------------------testing bindvalue abc.test.ont--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account4,
		codeAddress,
		[]interface{}{"bindValue", []interface{}{"abc.test.ont",1,"1","somevalue_test_abc"}})
	if err != nil {
		ctx.LogError("updateValidPeriod error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError(" WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("registerDomain GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("registerDomain failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing bindvalue abc.test.ont end--------------------")
	ctx.LogInfo("--------------------testing valueOf ont--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"valueOf", []interface{}{"abc.test.ont"}})

	value, err = obj.Result.ToString()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("value of abc.test.ont is %s\n", value)
	ctx.LogInfo("--------------------testing valueOf ont end--------------------")


	flag = false
	if flag {
		ctx.LogInfo("--------------------testing registerDomain--------------------")
		timeont := time.Now().Unix() + 3600
		txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
			account4,
			codeAddress,
			[]interface{}{"registerDomain", []interface{}{"dde.abc.test.ont",ontidPrefix+account4.Address.ToBase58(),1,timeont}})
		if err != nil {
			ctx.LogError("registerDomain error: %s", err)
		}

		//WaitForGenerateBlock
		_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
		if err != nil {
			ctx.LogError(" WaitForGenerateBlock error: %s", err)
			return false
		}

		//GetEventLog, to check the result of invoke
		events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
		if err != nil {
			ctx.LogError("registerDomain GetSmartContractEvent error:%s", err)
			return false
		}
		if events.State == 0 {
			ctx.LogError("registerDomain failed invoked exec state return 0")
			return false
		}
		for _, notify := range events.Notify {
			ctx.LogInfo("%+v", notify)
		}
		ctx.LogInfo("--------------------testing registerDomain end--------------------")
	}
	ctx.LogInfo("--------------------testing ownerOf abc.test.ont--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"ownerOf", []interface{}{"dde.abc.test.ont"}})

	owner, err = obj.Result.ToString()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("owner of test.ont is %s\n", owner)
	ctx.LogInfo("--------------------testing ownerOf abc.test.ont end--------------------")
	ctx.LogInfo("--------------------testing validTo abc.test.ont--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"validTo", []interface{}{"dde.abc.test.ont"}})

	validto, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("validTo of test.ont is %d\n", validto)
	ctx.LogInfo("--------------------testing ownerOf abc.test.ont end--------------------")
	ctx.LogInfo("--------------------testing isDomainValid abc.test.ont--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"isDomainValid", []interface{}{"dde.abc.test.ont"}})

	isDomainValid, err = obj.Result.ToBool()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("isDomainValid of ont is %v\n", isDomainValid)
	ctx.LogInfo("--------------------testing isDomainValid ont end--------------------")

	ctx.LogInfo("-------------------------test transfer-------------------")
	flag = true
	if flag{
		account3, err := ctx.GetAccount("AK98G45DhmPXg4TFPG1KjftvkEaHbU8SHM")
		txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
			account4,
			codeAddress,
			[]interface{}{"transfer", []interface{}{"dde.abc.test.ont",1,ontidPrefix+account3.Address.ToBase58()}})
		if err != nil {
			ctx.LogError("registerDomain error: %s", err)
		}

		//WaitForGenerateBlock
		_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
		if err != nil {
			ctx.LogError(" WaitForGenerateBlock error: %s", err)
			return false
		}

		//GetEventLog, to check the result of invoke
		events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
		if err != nil {
			ctx.LogError("registerDomain GetSmartContractEvent error:%s", err)
			return false
		}
		if events.State == 0 {
			ctx.LogError("registerDomain failed invoked exec state return 0")
			return false
		}
		for _, notify := range events.Notify {
			ctx.LogInfo("%+v", notify)
		}

	}
	ctx.LogInfo("--------------------testing ownerOf abc.test.ont--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"ownerOf", []interface{}{"dde.abc.test.ont"}})

	owner, err = obj.Result.ToString()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("owner of test.ont is %s\n", owner)
	ctx.LogInfo("--------------------testing ownerOf abc.test.ont end--------------------")


	ctx.LogInfo("--------------------testing getDomains--------------------")
	did1:=ontidPrefix+signer.Address.ToBase58()
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getDomains", []interface{}{did1}})

	domains, err := obj.Result.ToString()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("did1:%s, domains is %s\n",did1, domains)
	ctx.LogInfo("--------------------testing getDomains end--------------------")

	ctx.LogInfo("--------------------testing getDomains--------------------")
	did1 =ontidPrefix+account2.Address.ToBase58()
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getDomains", []interface{}{did1}})

	domains, err = obj.Result.ToString()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("did1:%s, domains is %s\n",did1, domains)
	ctx.LogInfo("--------------------testing getDomains end--------------------")



	ctx.LogInfo("--------------------testing getDomains--------------------")
	did1 =ontidPrefix+account4.Address.ToBase58()
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getDomains", []interface{}{did1}})

	domains, err = obj.Result.ToString()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("did1:%s, domains is %s\n",did1, domains)
	ctx.LogInfo("--------------------testing getDomains end--------------------")


	ctx.LogInfo("--------------------testing delete domain--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account4,
		codeAddress,
		[]interface{}{"deleteDomain", []interface{}{"dde.abc.test.ont",1}})
	if err != nil {
		ctx.LogError("registerDomain error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError(" WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("registerDomain GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("registerDomain failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing delete Domain end--------------------")
	ctx.LogInfo("--------------------testing isDomainValid dde.abc.test.ont--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"isDomainValid", []interface{}{"dde.abc.test.ont"}})

	isDomainValid, err = obj.Result.ToBool()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("isDomainValid of ont is %v\n", isDomainValid)
	ctx.LogInfo("--------------------testing isDomainValid dde.abc.test.ont end--------------------")


	return true
}

