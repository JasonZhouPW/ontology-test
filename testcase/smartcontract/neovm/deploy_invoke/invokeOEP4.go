package deploy_invoke

import (
	"fmt"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"time"
)

func TestOEP4Py(ctx *testframework.TestFrameworkContext) bool {

	//avmfile := "test_data/OEP4Sample.avm"
	//
	//code, err := ioutil.ReadFile(avmfile)
	//if err != nil {
	//	return false
	//}
	codeHash := "0132c56b6a00527ac46a51527ac46a00c304696e69749c6409006501096c7566616a00c3046e616d659c64090065d8086c7566616a00c30673796d626f6c9c64090065b1086c7566616a00c308646563696d616c739c640900658b086c7566616a00c30b746f74616c537570706c799c640900651d086c7566616a00c30962616c616e63654f669c6424006a51c3c0519e640700006c7566616a51c300c36a52527ac46a52c36576076c7566616a00c3087472616e736665729c6440006a51c3c0539e640700006c7566616a51c300c36a53527ac46a51c351c36a54527ac46a51c352c36a55527ac46a53c36a54c36a55c35272657d056c7566616a00c30d7472616e736665724d756c74699c640c006a51c365c5046c7566616a00c30c7472616e7366657246726f6d9c645f006a51c3c0549e640700006c7566616a51c300c36a56527ac46a51c351c36a53527ac46a51c352c36a54527ac46a51c353c36a55527ac46a56c36a53c36a54c36a55c3537951795572755172755279527954727552727565fd006c7566616a00c307617070726f76659c6440006a51c3c0539e640700006c7566616a51c300c36a57527ac46a51c351c36a56527ac46a51c352c36a55527ac46a57c36a56c36a55c3527265f9026c7566616a00c309616c6c6f77616e63659c6432006a51c3c0529e640700006c7566616a51c300c36a57527ac46a51c351c36a56527ac46a57c36a56c37c650b006c756661006c756658c56b6a00527ac46a51527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a52527ac401026a53527ac46a53c36a00c37e6a51c37e6a54527ac46a52c36a54c37c681253797374656d2e53746f726167652e476574616c7566011fc56b6a00527ac46a51527ac46a52527ac46a53527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a54527ac401016a55527ac401026a56527ac46a00c3c001149e6317006a51c3c001149e630d006a52c3c001149e641a00611461646472657373206c656e677468206572726f72f0616a00c3681b53797374656d2e52756e74696d652e436865636b5769746e65737361009c640700006c7566616a55c36a51c37e6a57527ac46a54c36a57c37c681253797374656d2e53746f726167652e476574616a58527ac46a53c36a58c3a0630b006a53c3009f64080061006c7566616a56c36a51c37e6a00c37e6a59527ac46a54c36a59c37c681253797374656d2e53746f726167652e476574616a5a527ac46a55c36a52c37e6a5b527ac46a53c36a5ac3a0640700006c7566616a53c36a5ac39c6449006a54c36a59c37c681553797374656d2e53746f726167652e44656c657465616a54c36a57c36a58c36a53c3945272681253797374656d2e53746f726167652e50757461624c00616a54c36a59c36a5ac36a53c3945272681253797374656d2e53746f726167652e507574616a54c36a57c36a58c36a53c3945272681253797374656d2e53746f726167652e50757461616a54c36a5bc37c681253797374656d2e53746f726167652e476574616a5c527ac46a54c36a5bc36a5cc36a53c3935272681253797374656d2e53746f726167652e507574616a51c36a52c36a53c35272087472616e7366657254c1681553797374656d2e52756e74696d652e4e6f74696679516c75660111c56b6a00527ac46a51527ac46a52527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a53527ac401026a54527ac46a51c3c001149e630d006a00c3c001149e641a00611461646472657373206c656e677468206572726f72f0616a00c3681b53797374656d2e52756e74696d652e436865636b5769746e65737361009c640700006c7566616a52c36a00c365ba02a0630b006a52c3009f64080061006c7566616a54c36a00c37e6a51c37e6a55527ac46a53c36a55c36a52c35272681253797374656d2e53746f726167652e507574616a00c36a51c36a52c3527208617070726f76616c54c1681553797374656d2e52756e74696d652e4e6f74696679516c756659c56b6a00527ac4006a52527ac46a00c3c06a53527ac4616a52c36a53c39f6473006a00c36a52c3c36a51527ac46a52c351936a52527ac46a51c3c0539e6420001b7472616e736665724d756c746920706172616d73206572726f722ef0616a51c300c36a51c351c36a51c352c35272652900009c64a2ff157472616e736665724d756c7469206661696c65642ef06288ff616161516c75660117c56b6a00527ac46a51527ac46a52527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a53527ac401016a54527ac46a51c3c001149e630d006a00c3c001149e641a00611461646472657373206c656e677468206572726f72f0616a00c3681b53797374656d2e52756e74696d652e436865636b5769746e65737361009c630b006a52c3009f64080061006c7566616a54c36a00c37e6a55527ac46a53c36a55c37c681253797374656d2e53746f726167652e476574616a56527ac46a52c36a56c3a0640700006c7566616a52c36a56c39c6425006a53c36a55c37c681553797374656d2e53746f726167652e44656c65746561622800616a53c36a55c36a56c36a52c3945272681253797374656d2e53746f726167652e50757461616a54c36a51c37e6a57527ac46a53c36a57c37c681253797374656d2e53746f726167652e476574616a58527ac46a53c36a57c36a58c36a52c3935272681253797374656d2e53746f726167652e507574616a00c36a51c36a52c35272087472616e7366657254c1681553797374656d2e52756e74696d652e4e6f74696679516c756658c56b6a00527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a51527ac401016a52527ac46a00c3c001149e6419001461646472657373206c656e677468206572726f72f0616a51c36a52c36a00c37e7c681253797374656d2e53746f726167652e476574616c756655c56b681953797374656d2e53746f726167652e476574436f6e74657874616a00527ac40b546f74616c537570706c796a51527ac46a00c36a51c37c681253797374656d2e53746f726167652e476574616c756654c56b586a00527ac46a00c36c756654c56b034d59546a00527ac46a00c36c756654c56b074d79546f6b656e6a00527ac46a00c36c75660113c56b681953797374656d2e53746f726167652e476574436f6e74657874616a00527ac40400e1f5056a51527ac422416434706a7a3262716570345268517255417a4d755a4a6b424333714a31745a75547514e98f4998d837fcdd44a50561f7f32140c7c6c2606a52527ac40400ca9a3b6a53527ac401016a54527ac40b546f74616c537570706c796a55527ac46a52c3c001149e6432000e4f776e657220696c6c6567616c2151c176c9681553797374656d2e52756e74696d652e4e6f7469667961006c7566616a00c36a55c37c681253797374656d2e53746f726167652e4765746164340014416c726561647920696e697469616c697a656421681553797374656d2e52756e74696d652e4e6f7469667961006c7566616a53c36a51c3956a56527ac46a00c36a55c36a56c35272681253797374656d2e53746f726167652e507574616a00c36a54c36a52c37e6a56c35272681253797374656d2e53746f726167652e50757461006a52c36a56c35272087472616e7366657254c1681553797374656d2e52756e74696d652e4e6f74696679516c7566006c75665ec56b6a00527ac46a51527ac46a51c36a00c3946a52527ac46a52c3c56a53527ac4006a54527ac46a00c36a55527ac461616a00c36a51c39f6433006a54c36a55c3936a56527ac46a56c36a53c36a54c37bc46a54c351936a54527ac46a55c36a54c3936a00527ac462c8ff6161616a53c36c7566"

	codeAddress, _ := utils.GetContractAddress(codeHash)

	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToHexString())
	ctx.LogInfo("=====CodeAddress base58===%s", codeAddress.ToBase58())

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestOEP4Py GetDefaultAccount error:%s", err)
		return false
	}

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

	ctx.LogInfo("--------------------testing init--------------------")
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"init", []interface{}{}})
	if err != nil {
		ctx.LogError("TestOEP4Py init error: %s", err)
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
	if events.State == 0 {
		ctx.LogError("TestOEP4Py failed invoked exec state return 0")
		return false
	}

	ctx.LogInfo("--------------------testing init end--------------------")

	ctx.LogInfo("--------------------testing totalSupply--------------------")
	obj, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"totalSupply", []interface{}{}})

	totalSupply, err := obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("total supply is %d\n", totalSupply)
	ctx.LogInfo("--------------------testing totalSupply end--------------------")

	ctx.LogInfo("--------------------testing name--------------------")

	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"name", []interface{}{}})

	name, err := obj.Result.ToString()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("name is %s\n", name)
	ctx.LogInfo("--------------------testing name end--------------------")

	ctx.LogInfo("--------------------testing symbol--------------------")

	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"symbol", []interface{}{}})

	symbol, err := obj.Result.ToString()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("symbol is %s\n", symbol)
	ctx.LogInfo("--------------------testing symbol end--------------------")

	ctx.LogInfo("--------------------testing decimals--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"decimals", []interface{}{}})

	decimal, err := obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("decimal is %d\n", decimal)
	ctx.LogInfo("--------------------testing decimals end--------------------")

	ctx.LogInfo("--------------------testing balanceOf owner--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"balanceOf", []interface{}{signer.Address[:]}})
	if err != nil {
		ctx.LogError("TestOEP4Py NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	balance, err := obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("balance is %d\n", balance)
	ctx.LogInfo("--------------------testing balanceOf owner end--------------------")

	ctx.LogInfo("--------------------testing transfer ---------------------------")
	tmpaddr ,_ := common.AddressFromBase58("AW6cKyfTTfMQPwG7NH1FjZ7ar8QLENRQwS")
	account2, err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
	if err != nil {
		ctx.LogError("get account AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb failed")
		return false
	}


	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"transfer", []interface{}{signer.Address[:], tmpaddr[:], 10000000000000,}})
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
	if events.State == 0 {
		ctx.LogError("TestOEP4Py failed invoked exec state return 0")
		return false
	}
	notify := events.Notify[0]
	ctx.LogInfo("%+v", notify)
	invokeState := notify.States.([]interface{})
	ctx.LogInfo(invokeState)

	method, _ := common.HexToBytes(invokeState[0].(string))
	addFromTmp, _ := common.HexToBytes(invokeState[1].(string))
	addFrom, _ := common.AddressParseFromBytes(addFromTmp)

	addToTmp, _ := common.HexToBytes(invokeState[2].(string))
	addTo, _ := common.AddressParseFromBytes(addToTmp)
	tmp, _ := common.HexToBytes(invokeState[3].(string))
	amount := common.BigIntFromNeoBytes(tmp)
	ctx.LogInfo("states[method:%s,from:%s,to:%s,value:%d]", method, addFrom.ToBase58(), addTo.ToBase58(), amount.Int64())

	ctx.LogInfo("--------------------testing transfer end---------------------------")

	ctx.LogInfo("--------------------testing balanceOf owner--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"balanceOf", []interface{}{signer.Address[:]}})
	if err != nil {
		ctx.LogError("TestOEP4Py NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	balance, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("balance is %d\n", balance)
	ctx.LogInfo("--------------------testing balanceOf owner end--------------------")

	ctx.LogInfo("--------------------testing balanceOf acct2--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"balanceOf", []interface{}{account2.Address[:]}})
	if err != nil {
		ctx.LogError("TestOEP4Py NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	balance, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("balance is %d\n", balance)
	ctx.LogInfo("--------------------testing balanceOf acct2 end--------------------")

	ctx.LogInfo("--------------------testing approve ---------------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"approve", []interface{}{signer.Address[:], account2.Address[:], 60000000000}})
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
	if events.State == 0 {
		ctx.LogError("TestOEP4Py failed invoked exec state return 0")
		return false
	}
	notify = events.Notify[0]
	ctx.LogInfo("%+v", notify)
	invokeState = notify.States.([]interface{})
	ctx.LogInfo(invokeState)

	method, _ = common.HexToBytes(invokeState[0].(string))
	addFromTmp, _ = common.HexToBytes(invokeState[1].(string))
	addFrom, _ = common.AddressParseFromBytes(addFromTmp)

	addToTmp, _ = common.HexToBytes(invokeState[2].(string))
	addTo, _ = common.AddressParseFromBytes(addToTmp)
	tmp, _ = common.HexToBytes(invokeState[3].(string))
	amount = common.BigIntFromNeoBytes(tmp)
	ctx.LogInfo("states[method:%s,from:%s,to:%s,value:%d]", method, addFrom.ToBase58(), addTo.ToBase58(), amount.Int64())

	ctx.LogInfo("--------------------testing approve end---------------------------")

	ctx.LogInfo("--------------------testing allowance signer--------------------")

	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"allowance", []interface{}{signer.Address[:], account2.Address[:]}})
	if err != nil {
		ctx.LogError("TestOEP4Py NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	allowance, err := obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("allowance is %d\n", allowance)
	ctx.LogInfo("--------------------testing allowance signer end--------------------")


	ctx.LogInfo("--------------------testing approve to contract---------------------------")

	acct5 ,_  := common.AddressFromBase58("ASNFJ3cxYrgrLnWPaMFWam2QSpUQrqJ9ue")

	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"approve", []interface{}{signer.Address[:], acct5[:], 6000000000000}})
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
	if events.State == 0 {
		ctx.LogError("TestOEP4Py failed invoked exec state return 0")
		return false
	}
	notify = events.Notify[0]
	ctx.LogInfo("%+v", notify)
	invokeState = notify.States.([]interface{})
	ctx.LogInfo(invokeState)

	method, _ = common.HexToBytes(invokeState[0].(string))
	addFromTmp, _ = common.HexToBytes(invokeState[1].(string))
	addFrom, _ = common.AddressParseFromBytes(addFromTmp)

	addToTmp, _ = common.HexToBytes(invokeState[2].(string))
	addTo, _ = common.AddressParseFromBytes(addToTmp)
	tmp, _ = common.HexToBytes(invokeState[3].(string))
	amount = common.BigIntFromNeoBytes(tmp)
	ctx.LogInfo("states[method:%s,from:%s,to:%s,value:%d]", method, addFrom.ToBase58(), addTo.ToBase58(), amount.Int64())

	ctx.LogInfo("--------------------testing approve end---------------------------")
	ctx.LogInfo("--------------------testing allowance of acct5--------------------")
	fmt.Printf("singer:%v\n", signer.Address[:])
	fmt.Printf("acct5:%v\n", acct5[:])
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"allowance", []interface{}{signer.Address[:], acct5[:]}})
	if err != nil {
		ctx.LogError("TestOEP4Py NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	allowance, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("allowance is %d\n", allowance)
	ctx.LogInfo("--------------------testing allowance signer end--------------------")


	ctx.LogInfo("--------------------testing transfer from ---------------------------")

	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"transferFrom", []interface{}{account2.Address[:], signer.Address[:], account2.Address[:], 30000000000}})
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
	if events.State == 0 {
		ctx.LogError("TestOEP4Py failed invoked exec state return 0")
		return false
	}
	notify = events.Notify[0]
	ctx.LogInfo("%+v", notify)
	invokeState = notify.States.([]interface{})
	ctx.LogInfo(invokeState)

	method, _ = common.HexToBytes(invokeState[0].(string))
	addFromTmp, _ = common.HexToBytes(invokeState[1].(string))
	addFrom, _ = common.AddressParseFromBytes(addFromTmp)

	addToTmp, _ = common.HexToBytes(invokeState[2].(string))
	addTo, _ = common.AddressParseFromBytes(addToTmp)
	tmp, _ = common.HexToBytes(invokeState[3].(string))
	amount = common.BigIntFromNeoBytes(tmp)
	ctx.LogInfo("states[method:%s,from:%s,to:%s,value:%d]", method, addFrom.ToBase58(), addTo.ToBase58(), amount.Int64())

	ctx.LogInfo("--------------------testing transfer from  end---------------------------")

	return true
}
