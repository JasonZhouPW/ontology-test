package deploy_invoke

import (
	"fmt"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"io/ioutil"
	"strconv"
	"time"
)

func DEXTest(ctx *testframework.TestFrameworkContext) bool {
	avmfile := "test_data/ONTDEX.avm"

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
	account2, err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
	account4, err := ctx.GetAccount("ALerVnMj3eNk9xe8BnQJtoWvwGmY3x4KMi")

	priceMultiple := 1000000000
	ctx.LogInfo("=================Deploy===============================")

	if err != nil {
		ctx.LogError("Dice GetDefaultAccount error:%s", err)
		return false
	}

	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		true,
		codeHash,
		"DEXGuess",
		"1.0",
		"",
		"",
		"",
	)

	if err != nil {
		ctx.LogError("Dice DeploySmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}
	ctx.LogInfo("=================Deploy end===========================")

	var txHash common.Uint256

	ctx.LogInfo("--------------------testing init--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"init", []interface{}{}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing init end--------------------")

	ctx.LogInfo("--------------------testing add sell ONT order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"addOrder", []interface{}{account2.Address[:], 1, 10 * priceMultiple, 1, 2 * priceMultiple,account4.Address[:],account4.Address[:]}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	sellorderid := 0
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)

		if notify.ContractAddress == codeAddress.ToHexString() {
			state := notify.States.([]interface{})

			n, err := strconv.ParseUint(state[1].(string), 16, 32)
			if err != nil {
				panic(err)
			}

			//bi:= common.BigIntFromNeoBytes([]byte(state[1].(string)))
			sellorderid = int(n)
		}

	}
	fmt.Printf("sellorderid :%d\n",sellorderid)

	ctx.LogInfo("--------------------testing add order end--------------------")

	ctx.LogInfo("--------------------testing getorder prex --------------------")
	obj, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getOrder", []interface{}{sellorderid}})
	if err != nil {
		ctx.LogError("unboundong NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	bs, err := obj.Result.ToString()
	if err != nil {
		ctx.LogError("unboundong PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("order is %s\n", bs)
	ctx.LogInfo("--------------------testing getorder prex end--------------------")

	ctx.LogInfo("--------------------testing getorderOwner prex --------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getOrderOwner", []interface{}{sellorderid}})
	if err != nil {
		ctx.LogError("unboundong NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	bytes, err := obj.Result.ToByteArray()
	if err != nil {
		ctx.LogError("unboundong PrepareInvokeContract error:%s", err)

		return false
	}

	addr, err := common.AddressParseFromBytes(bytes)
	if err != nil {
		ctx.LogError("unboundong PrepareInvokeContract error:%s", err)

		return false
	}
	//
	fmt.Printf("order owner is %s\n", addr.ToBase58())
	ctx.LogInfo("--------------------testing getorder prex end--------------------")

	account3, err := ctx.GetAccount("AK98G45DhmPXg4TFPG1KjftvkEaHbU8SHM")

	ctx.LogInfo("--------------------testing add buy ONT order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account3,
		codeAddress,
		[]interface{}{"addOrder", []interface{}{account3.Address[:], 1, 200 * priceMultiple, 0, 2 * priceMultiple,account4.Address[:],account4.Address[:]}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	buyorderid := 0
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
		if notify.ContractAddress == codeAddress.ToHexString() {
			state := notify.States.([]interface{})
			n, err := strconv.ParseUint(state[1].(string), 16, 32)
			if err != nil {
				panic(err)
			}

			//bi:= common.BigIntFromNeoBytes([]byte(state[1].(string)))
			buyorderid = int(n)
		}
	}
	fmt.Printf("buyorderid :%d\n",buyorderid)

	ctx.LogInfo("--------------------testing add order end--------------------")

	ctx.LogInfo("--------------------testing match order--------------------")
	fmt.Printf("matching:buyorderid:%d,sellorderid:%d\n", buyorderid, sellorderid)
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"matchOrder", []interface{}{buyorderid, sellorderid, 10 * priceMultiple, 2 * priceMultiple, 1}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing match end--------------------")

	ctx.LogInfo("--------------------testing add sell ONT order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"addOrder", []interface{}{account2.Address[:], 1, 100 * priceMultiple, 1, 2 * priceMultiple,"",""}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
		if notify.ContractAddress == codeAddress.ToHexString() {
			state := notify.States.([]interface{})
			n, err := strconv.ParseUint(state[1].(string), 16, 32)
			if err != nil {
				panic(err)
			}

			//bi:= common.BigIntFromNeoBytes([]byte(state[1].(string)))
			sellorderid = int(n)
		}
	}
	ctx.LogInfo("--------------------testing add order end--------------------")

	ctx.LogInfo("--------------------testing cancel sell ONT order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"cancelOrder", []interface{}{account2.Address[:], sellorderid}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing cancel order end--------------------")

	ctx.LogInfo("--------------------testing add buy ONT order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account3,
		codeAddress,
		[]interface{}{"addOrder", []interface{}{account3.Address[:], 1, 200 * priceMultiple, 0, 2 * priceMultiple}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
		if notify.ContractAddress == codeAddress.ToHexString() {
			state := notify.States.([]interface{})
			n, err := strconv.ParseUint(state[1].(string), 16, 32)
			if err != nil {
				panic(err)
			}

			//bi:= common.BigIntFromNeoBytes([]byte(state[1].(string)))
			buyorderid = int(n)
		}
	}
	ctx.LogInfo("--------------------testing add order end--------------------")

	ctx.LogInfo("--------------------testing cancel buy ONT order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account3,
		codeAddress,
		[]interface{}{"cancelOrder", []interface{}{account3.Address[:], buyorderid}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing cancel order end--------------------")
	ctx.LogInfo("====================testing OEP4 case ==================")
	ctx.LogInfo("--------------------testing add pair--------------------")
	addr, _ = common.AddressFromBase58("APPWgNbWvUdQjQxeN7RduYweH3caaM1LM1")
	ong, _ := common.AddressFromBase58("AFmseVrdL9f9oyCzZefL9tG6UbvhfRZMHJ")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"addPair", []interface{}{"SYMBOL", addr[:], 8, "ONG", ong[:], 9, 1000}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing addpair end--------------------")

	ctx.LogInfo("--------------------testing set fee--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"setTradeFee", []interface{}{5000}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing setfee end--------------------")

	ctx.LogInfo("--------------------testing add sell OEP4 order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"addOrder", []interface{}{account2.Address[:], 2, priceMultiple / 10, 1, priceMultiple / 100}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
		if notify.ContractAddress == codeAddress.ToHexString() {
			state := notify.States.([]interface{})
			n, err := strconv.ParseUint(state[1].(string), 16, 32)
			if err != nil {
				panic(err)
			}

			//bi:= common.BigIntFromNeoBytes([]byte(state[1].(string)))
			sellorderid = int(n)
		}
	}
	ctx.LogInfo("--------------------testing add order end--------------------")

	ctx.LogInfo("--------------------testing add buy OEP4 order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account3,
		codeAddress,
		[]interface{}{"addOrder", []interface{}{account3.Address[:], 2, 200 * priceMultiple, 0, priceMultiple / 100}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
		if notify.ContractAddress == codeAddress.ToHexString() {
			state := notify.States.([]interface{})
			n, err := strconv.ParseUint(state[1].(string), 16, 32)
			if err != nil {
				panic(err)
			}

			//bi:= common.BigIntFromNeoBytes([]byte(state[1].(string)))
			buyorderid = int(n)
		}
	}
	ctx.LogInfo("--------------------testing add order end--------------------")

	ctx.LogInfo("--------------------testing match order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"matchOrder", []interface{}{buyorderid, sellorderid, 100 * priceMultiple, priceMultiple / 100, 1}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing match end--------------------")

	bakAccount, _ := ctx.GetAccount("AKmowTi8NcAMjZrg7ZNtSQUtnEgdaC65wG")
	ctx.LogInfo("--------------------testing withdrawAssets order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		bakAccount,
		codeAddress,
		[]interface{}{"withdrawAssets", []interface{}{}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing withdrawAssets end--------------------")

	ctx.LogInfo("--------------------testing unboundong --------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getUnboundONG", []interface{}{}})
	if err != nil {
		ctx.LogError("unboundong NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	balance, err := obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("unboundong PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("unboundong is %d\n", balance)
	ctx.LogInfo("--------------------testing poolBalance end--------------------")

	ctx.LogInfo("--------------------testing withdrawUnboundOng order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"withUnboundONG", []interface{}{}})
	if err != nil {
		ctx.LogError("withUnboundONG invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("withUnboundONG WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("withUnboundONG GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("withUnboundONG failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing withUnboundONG end--------------------")

	ctx.LogInfo("--------------------testing getOrder --------------------")

	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getOrder", []interface{}{1}})
	if err != nil {
		ctx.LogError("withUnboundONG invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("withUnboundONG WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("withUnboundONG GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("withUnboundONG failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing getOrder end--------------------")

	ctx.LogInfo("--------------------testing getorder prex --------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getOrder", []interface{}{1}})
	if err != nil {
		ctx.LogError("unboundong NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	str, err := obj.Result.ToString()
	if err != nil {
		ctx.LogError("unboundong PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("order is %s\n", str)
	ctx.LogInfo("--------------------testing getorder prex end--------------------")

	ctx.LogInfo("--------------------testing add sell OEP4 order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"addOrder", []interface{}{account2.Address[:], 2, 300 * priceMultiple, 1, priceMultiple / 100}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
		if notify.ContractAddress == codeAddress.ToHexString() {
			state := notify.States.([]interface{})
			n, err := strconv.ParseUint(state[1].(string), 16, 32)
			if err != nil {
				panic(err)
			}

			//bi:= common.BigIntFromNeoBytes([]byte(state[1].(string)))
			sellorderid = int(n)
		}
	}
	ctx.LogInfo("--------------------testing add order end--------------------")

	ctx.LogInfo("--------------------testing deletePair start--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"deletePair", []interface{}{2}})
	if err != nil {
		ctx.LogError("withUnboundONG invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("withUnboundONG WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("withUnboundONG GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("withUnboundONG failed invoked exec state return 0")
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing deletePair end--------------------")

	//ctx.LogInfo("--------------------test migrate start----------------------")
	//newcode := "0122c56b05322e302e306a00527ac4681953797374656d2e53746f726167652e476574436f6e746578746a51527ac422416434706a7a3262716570345268517255417a4d755a4a6b424333714a31745a755468204f6e746f6c6f67792e52756e74696d652e426173653538546f416464726573736a52527ac422414c6572566e4d6a33654e6b39786538426e514a746f577677476d593378344b4d6968204f6e746f6c6f67792e52756e74696d652e426173653538546f416464726573736a53527ac422414b6d6f775469384e63414d6a5a7267375a4e74535155746e45676461433635774768204f6e746f6c6f67792e52756e74696d652e426173653538546f416464726573736a54527ac4682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e67536372697074486173686a55527ac4056f726465726a56527ac404696e69746a57527ac4096f6e675470616972736a58527ac4096f6e745470616972736a59527ac404706169726a5a527ac40a70616972636f756e74736a5b527ac4076163636f756e746a5c527ac4086c6f636b616363746a5d527ac40b6f72646572636f756e74736a5e527ac4006a5f527ac4516a60527ac4006a0111527ac4516a0112527ac40400ca9a3b6a0113527ac40400ca9a3b6a0114527ac407666565726174656a0115527ac40340420f6a0116527ac41400000000000000000000000000000000000000016a0117527ac41400000000000000000000000000000000000000026a0118527ac46c011fc56b6a00527ac46a51527ac46a52527ac46a51c304696e69747d9c7c75641800006a53527ac46a53c36a00c3651f0a6c75666203006a51c307616464506169727d9c7c7564a8006a52c3c0567d9e7c75640a00006c75666203006a00c352c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f1006a53527ac46a52c355c3516a53c3936a53527ac46a52c354c3516a53c3936a53527ac46a52c353c3516a53c3936a53527ac46a52c352c3516a53c3936a53527ac46a52c351c3516a53c3936a53527ac46a52c300c3516a53c3936a53527ac46a53c36a00c365f90d6c75666203006a51c30a64656c657465506169727d9c7c75643a006a52c3c0567d9e7c75640a00006c7566620300006a53527ac46a52c300c3516a53c3936a53527ac46a53c36a00c365da0f6c75666203006a51c3086164644f726465727d9c7c756476006a52c3c0557d9e7c75640a00006c7566620300006a53527ac46a52c354c3516a53c3936a53527ac46a52c353c3516a53c3936a53527ac46a52c352c3516a53c3936a53527ac46a52c351c3516a53c3936a53527ac46a52c300c3516a53c3936a53527ac46a53c36a00c36585106c75666203006a51c30b63616e63656c4f726465727d9c7c756449006a52c3c0527d9e7c75640a00006c7566620300006a53527ac46a52c351c3516a53c3936a53527ac46a52c300c3516a53c3936a53527ac46a53c36a00c36592136c75666203006a51c30a6d617463684f726465727d9c7c756476006a52c3c0557d9e7c75640a00006c7566620300006a53527ac46a52c354c3516a53c3936a53527ac46a52c353c3516a53c3936a53527ac46a52c352c3516a53c3936a53527ac46a52c351c3516a53c3936a53527ac46a52c300c3516a53c3936a53527ac46a53c36a00c365c1156c75666203006a51c30b73657454726164654665657d9c7c75643a006a52c3c0517d9e7c75640a00006c7566620300006a53527ac46a52c300c3516a53c3936a53527ac46a53c36a00c365030b6c75666203006a51c30b67657454726164654665657d9c7c75641800006a53527ac46a53c36a00c365990b6c75666203006a51c30f6d696772617465436f6e74726163747d9c7c756494006a52c3c0577d9e7c75640a00006c7566620300006a53527ac46a52c356c3516a53c3936a53527ac46a52c355c3516a53c3936a53527ac46a52c354c3516a53c3936a53527ac46a52c353c3516a53c3936a53527ac46a52c352c3516a53c3936a53527ac46a52c351c3516a53c3936a53527ac46a52c300c3516a53c3936a53527ac46a53c36a00c3655c1e6c75666203006a51c3086765744f726465727d9c7c75643a006a52c3c0517d9e7c75640a00006c7566620300006a53527ac46a52c300c3516a53c3936a53527ac46a53c36a00c3654b076c75666203006a51c30d6765744f726465724f776e65727d9c7c75643a006a52c3c0517d9e7c75640a00006c7566620300006a53527ac46a52c300c3516a53c3936a53527ac46a53c36a00c36512096c75666203006a51c30e6765744f72646572436f756e74737d9c7c75641800006a53527ac46a53c36a00c36535096c75666203006a51c30e77697468647261774173736574737d9c7c75641800006a53527ac46a53c36a00c365b2256c75666203006a51c30d676574556e626f756e644f4e477d9c7c75641800006a53527ac46a53c36a00c365151e6c75666203006a51c30e77697468556e626f756e644f4e477d9c7c75641800006a53527ac46a53c36a00c3654c1e6c7566620300006c75660111c56b6a00527ac46a51527ac46a51c300947600a0640c00c16a52527ac4620e007562030000c56a52527ac46a52c3c0517d9c7c75641c00006a53527ac46a52c300c36a54527ac4516a55527ac4625c006a52c3c0527d9c7c756421006a52c300c36a53527ac46a52c351c36a54527ac4516a55527ac4616232006a52c3c0537d9c7c756424006a52c300c36a53527ac46a52c351c36a54527ac46a52c352c36a55527ac462050000f100c176c96a56527ac46a53c36a57527ac46a57c36a54c37d9f7c756419006a56c36a57c3c86a57c36a55c3936a57527ac462e0ff6a56c36c75665cc56b6a00527ac46a51527ac46a52527ac46a53527ac4620300006a54527ac46a52c36a55527ac46a55c3c06a56527ac46a54c36a56c39f6430006a55c36a54c3c36a57527ac46a54c351936a54527ac46a53c36a57c37d877c75640a00516c756662030062ccff006c75660118c56b6a00527ac46a51527ac46a52527ac46a51c3519c640600620b006a53527ac46209005a6a53527ac46a52c3c06a54527ac4516a55527ac4006a56527ac40130013101320133013401350136013701380139012d5bc176c96a57527ac401610162016301640165016656c176c96a58527ac401410142014301440145014656c176c96a59527ac46a53c35a7d9e7c7576640c00756a53c3607d9e7c7564080000f1620300006a5a527ac4006a5b527ac46a54c3516a5bc3936a5b527ac400516a5bc3936a5b527ac46a5bc36a00c365defd76c96a5c527ac46a5cc3c06a5d527ac46a5ac36a5dc39f64b0016a5cc36a5ac3c36a5e527ac46a5ac351936a5a527ac46a52c36a5ec351936a5ec37b6b766b946c6c52727f6a5f527ac46a5ec3007d9e7c7576640d00756a5fc3012d7d9c7c7564080000f16203006a5ec3007d9c7c7576640d00756a5fc3012d7d9c7c75641300006a56c3946a56527ac4623c01622a016a53c35a7d9c7c75644300006a5b527ac46a5fc3516a5bc3936a5b527ac46a57c3516a5bc3936a5b527ac46a5bc36a00c365fefdf16a56c36a5fc30130946a55c395936a56527ac462df006a53c3607d9c7c7564d200006a5b527ac46a5fc3516a5bc3936a5b527ac46a57c3516a5bc3936a5b527ac46a5bc36a00c365b3fd6419006a56c36a5fc30130946a55c395936a56527ac4628d00006a5b527ac46a5fc3516a5bc3936a5b527ac46a58c3516a5bc3936a5b527ac46a5bc36a00c36571fd641b006a56c36a5fc30161945a936a55c395936a56527ac4624900006a5b527ac46a5fc3516a5bc3936a5b527ac46a59c3516a5bc3936a5b527ac46a5bc36a00c3652dfd641b006a56c36a5fc30141945a936a55c395936a56527ac462050000f162050000f16a55c36a53c3956a55527ac4624cfe6a56c36c756656c56b6a00527ac46a51527ac46a52527ac4620300006a53527ac45a516a53c3936a53527ac46a52c3516a53c3936a53527ac46a53c36a00c36506006c75665ec56b6a00527ac46a51527ac46a52527ac46a53527ac46203006a53c35a7d9e7c7576640c00756a53c3607d9e7c7564080000f16203006a52c300936a52527ac4006a54527ac4c77601307c007bc47601317c517bc47601327c527bc47601337c537bc47601347c547bc47601357c557bc47601367c567bc47601377c577bc47601387c587bc47601397c597bc47601617c5a7bc47601627c5b7bc47601637c5c7bc47601647c5d7bc47601657c5e7bc47601667c5f7bc46a55527ac46a52c36a56527ac46a56c3007d9e7c75642e006a56c36a53c3976a57527ac46a55c36a57c3c36a54c37e6a54527ac46a56c36a53c3966a56527ac462cdff6a54c36c756656c56b6a00527ac46a51527ac46203006a00c352c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f16a00c357c36a00c351c3681253797374656d2e53746f726167652e476574640a00006c756662030004747275656a00c357c36a00c351c3681253797374656d2e53746f726167652e507574006a52527ac459516a52c3936a52527ac46a00c30118c3516a52c3936a52527ac4034f4e47516a52c3936a52527ac400516a52c3936a52527ac46a00c30117c3516a52c3936a52527ac4034f4e54516a52c3936a52527ac46a52c36a00c365b503750288136a00c30115c36a00c351c3681253797374656d2e53746f726167652e507574516c75665ac56b6a00527ac46a51527ac46a52527ac4620300006a53527ac46a52c3516a53c3936a53527ac46a53c36a00c365f9186a54527ac46a54c391640a00006c756662d201086765744f726465726a52c36a54c306706169726964c36a54c3056f776e6572c36a54c306616d6f756e74c36a54c30474797065c36a54c3057072696365c36a54c3096465616c7072696365c36a54c306737461747573c359c176c9681553797374656d2e52756e74696d652e4e6f746966796a54c3056f776e6572c368204f6e746f6c6f67792e52756e74696d652e41646472657373546f4261736535386a55527ac4006a53527ac46a52c3516a53c3936a53527ac46a53c36a00c365b9fc006a53527ac46a54c306706169726964c3516a53c3936a53527ac46a53c36a00c36595fc6a55c3006a53527ac46a54c306616d6f756e74c3516a53c3936a53527ac46a53c36a00c3656efc006a53527ac46a54c30474797065c3516a53c3936a53527ac46a53c36a00c3654cfc006a53527ac46a54c3057072696365c3516a53c3936a53527ac46a53c36a00c36529fc006a53527ac46a54c3096465616c7072696365c3516a53c3936a53527ac46a53c36a00c36502fc006a53527ac46a54c306737461747573c3516a53c3936a53527ac46a53c36a00c365defb58c176c96a56527ac4006a53527ac4012c516a53c3936a53527ac46a56c3516a53c3936a53527ac46a53c36a00c365c0216c75666c756658c56b6a00527ac46a51527ac46a52527ac4620300006a53527ac46a52c3516a53c3936a53527ac46a53c36a00c365e3166a54527ac46a54c391640a00006c75666203006a54c3056f776e6572c36c756655c56b6a00527ac46a51527ac46203006a00c35ec36a00c351c3681253797374656d2e53746f726167652e4765746c756657c56b6a00527ac46a51527ac46a52527ac46203006a00c352c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f16a52c3007da27c7576640f00756a52c30340420f7d9f7c75f16a00c30115c36a00c351c3681253797374656d2e53746f726167652e4765746a53527ac46a52c36a00c30115c36a00c351c3681253797374656d2e53746f726167652e5075746a52c36a53c30673657446656553c1681553797374656d2e52756e74696d652e4e6f74696679516c756655c56b6a00527ac46a51527ac46203006a00c30115c36a00c351c3681253797374656d2e53746f726167652e4765746c75665ec56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46a55527ac46a56527ac46a57527ac46203006a00c352c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f1006a58527ac46a54c3516a58c3936a58527ac46a52c3516a58c3936a58527ac46a53c3516a58c3936a58527ac46a58c36a00c365ef1ef1006a58527ac46a57c3516a58c3936a58527ac46a55c3516a58c3936a58527ac46a56c3516a58c3936a58527ac46a58c36a00c365b81ef16a00c35bc36a00c351c3681253797374656d2e53746f726167652e47657451936a59527ac4006a58527ac4006a58527ac46a59c3516a58c3936a58527ac46a58c36a00c3652cf9516a58c3936a58527ac46a00c35ac3516a58c3936a58527ac46a58c36a00c365c1156a5a527ac46a5ac36a00c351c3681253797374656d2e53746f726167652e47657491f1c7766a53c37c0862617365616464727bc4766a54c37c0b62617365646563696d616c7bc4766a56c37c0971756f7465616464727bc4766a57c37c0c71756f7465646563696d616c7bc46a5b527ac46a59c36a00c35bc36a00c351c3681253797374656d2e53746f726167652e5075746a5bc3681853797374656d2e52756e74696d652e53657269616c697a656a5ac36a00c351c3681253797374656d2e53746f726167652e5075746a57c36a56c36a55c36a54c36a53c36a52c36a59c3076164645061697258c1681553797374656d2e52756e74696d652e4e6f74696679516c756659c56b6a00527ac46a51527ac46a52527ac46203006a00c352c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f10131681253797374656d2e52756e74696d652e4c6f67006a53527ac46a52c3516a53c3936a53527ac46a53c36a00c365ad136a54527ac46a54c3086261736561646472c36a55527ac4006a53527ac4006a53527ac46a52c3516a53c3936a53527ac46a53c36a00c36562f7516a53c3936a53527ac46a00c35ac3516a53c3936a53527ac46a53c36a00c365f7136a56527ac46a56c36a00c351c3681553797374656d2e53746f726167652e44656c657465006a53527ac46a55c3516a53c3936a53527ac46a53c36a00c3658b1a756a52c30a64656c6574655061697252c1681553797374656d2e52756e74696d652e4e6f74696679516c7566011bc56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46a55527ac46a56527ac46203006a52c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f16a56c3007da07c75f16a54c3007da07c75f1006a57527ac46a53c3516a57c3936a57527ac46a57c36a00c3656d126a58527ac46a58c3086261736561646472c36a59527ac46a58c30971756f746561646472c36a5a527ac46a59c36a5b527ac46a58c30b62617365646563696d616cc36a5c527ac46a54c36a5d527ac4516a5e527ac4006a5f527ac46a55c36a00c30111c37d9c7c75645f006a54c36a56c3956a5d527ac46a00c30113c36a5e527ac46a5ac36a5b527ac46a00c30115c36a00c351c3681253797374656d2e53746f726167652e4765746a5f527ac46a58c30c71756f7465646563696d616cc36a5c527ac4620300006a57527ac46a5dc3006a57527ac46a5cc3516a57c3936a57527ac45a516a57c3936a57527ac46a57c36a00c3656415956a5ec36a00c30114c395966a5dc36a5fc395006a57527ac46a5cc3516a57c3936a57527ac45a516a57c3936a57527ac46a57c36a00c3652a15956a00c30116c36a5ec3956a00c30114c3959693516a57c3936a57527ac46a5bc3516a57c3936a57527ac46a52c3516a57c3936a57527ac46a57c36a00c3653314f16a00c35ec36a60527ac46a60c36a00c351c3681253797374656d2e53746f726167652e4765746a0111527ac46a0111c351936a0112527ac46a0111c351936a60c36a00c351c3681253797374656d2e53746f726167652e507574c7766a53c37c067061697269647bc4766a52c37c056f776e65727bc4766a54c37c06616d6f756e747bc4766a55c37c04747970657bc4766a56c37c0570726963657bc476007c096465616c70726963657bc4766a00c35fc37c067374617475737bc46a0113527ac46a0113c3681853797374656d2e52756e74696d652e53657269616c697a65006a57527ac4006a57527ac46a0112c3516a57c3936a57527ac46a57c36a00c365e5f3516a57c3936a57527ac46a00c356c3516a57c3936a57527ac46a57c36a00c3657a106a00c351c3681253797374656d2e53746f726167652e5075746a00c35fc3006a56c36a55c36a54c36a53c36a52c36a0112c3086164644f7264657259c1681553797374656d2e52756e74696d652e4e6f74696679516c75660117c56b6a00527ac46a51527ac46a52527ac46a53527ac46203006a52c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f1006a54527ac46a53c3516a54c3936a54527ac46a54c36a00c365800e6a55527ac46a55c3056f776e6572c36a56527ac46a52c36a56c37d9c7c75f16a55c306737461747573c36a00c360c37d9e7c75f16a55c306616d6f756e74c36a57527ac46a55c3057072696365c36a58527ac46a55c306706169726964c36a59527ac4006a54527ac46a59c3516a54c3936a54527ac46a54c36a00c365a60e6a5a527ac46a5ac3086261736561646472c36a5b527ac46a5ac30971756f746561646472c36a5c527ac46a5bc36a5d527ac46a5ac30b62617365646563696d616cc36a5e527ac4006a5f527ac46a55c30474797065c36a00c30111c37d9c7c75645b006a57c36a58c3956a00c30113c3966a57527ac46a5cc36a5d527ac46a00c30115c36a00c351c3681253797374656d2e53746f726167652e4765746a5f527ac46a5ac30c71756f7465646563696d616cc36a5e527ac4620300006a54527ac46a57c3006a54527ac46a5ec3516a54c3936a54527ac45a516a54c3936a54527ac46a54c36a00c365a911956a00c30114c3966a57c36a5fc395006a54527ac46a5ec3516a54c3936a54527ac45a516a54c3936a54527ac46a54c36a00c3657311956a00c30116c36a00c30114c3959693516a54c3936a54527ac46a5dc3516a54c3936a54527ac46a52c3516a54c3936a54527ac46a54c36a00c365d811f1006a54527ac46a53c3516a54c3936a54527ac46a00c356c3516a54c3936a54527ac46a54c36a00c365c10d6a60527ac46a60c36a00c351c3681553797374656d2e53746f726167652e44656c6574656a53c36a59c36a52c30b63616e63656c4f7264657254c1681553797374656d2e52756e74696d652e4e6f74696679516c75660124c56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46a55527ac46a56527ac46203006a00c352c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f1006a57527ac46a52c3516a57c3936a57527ac46a57c36a00c365b80b6a58527ac4006a57527ac46a53c3516a57c3936a57527ac46a57c36a00c365970b6a59527ac46a54c3007da07c75916460000b616d6f756e7420697320306a56c36a55c36a54c36a53c36a52c36a59c306706169726964c36a58c306706169726964c30b4d617463684661696c656459c1681553797374656d2e52756e74696d652e4e6f74696679006c75666203006a58c306737461747573c36a00c360c37d9e7c7591646a00156275796f7264657220737461747573206572726f726a56c36a55c36a54c36a53c36a52c36a59c306706169726964c36a58c306706169726964c30b4d617463684661696c656459c1681553797374656d2e52756e74696d652e4e6f74696679006c75666203006a59c306737461747573c36a00c360c37d9e7c7591646b001673656c6c6f7264657220737461747573206572726f726a56c36a55c36a54c36a53c36a52c36a59c306706169726964c36a58c306706169726964c30b4d617463684661696c656459c1681553797374656d2e52756e74696d652e4e6f74696679006c75666203006a58c30474797065c36a00c30111c37d9c7c7576641700756a59c30474797065c36a00c30112c37d9c7c7591646500106f726465722074797065206572726f726a56c36a55c36a54c36a53c36a52c36a59c306706169726964c36a58c306706169726964c30b4d617463684661696c656459c1681553797374656d2e52756e74696d652e4e6f74696679006c75666203006a58c306706169726964c36a59c306706169726964c37d9c7c7591646d00186f7264657220706169726964206973206e6f742073616d656a56c36a55c36a54c36a53c36a52c36a59c306706169726964c36a58c306706169726964c30b4d617463684661696c656459c1681553797374656d2e52756e74696d652e4e6f74696679006c75666203006a58c306616d6f756e74c36a54c37da27c7576641600756a59c306616d6f756e74c36a54c37da27c7591646700126d6174636820616d6f756e74206572726f726a56c36a55c36a54c36a53c36a52c36a59c306706169726964c36a58c306706169726964c30b4d617463684661696c656459c1681553797374656d2e52756e74696d652e4e6f74696679006c75666203006a58c306616d6f756e74c36a54c3946a58c306616d6f756e747bc46a59c306616d6f756e74c36a54c3946a59c306616d6f756e747bc4006a57527ac46a58c306706169726964c3516a57c3936a57527ac46a57c36a00c365cb086a5a527ac46a5ac3916463000e70616972206e6f742065786973746a56c36a55c36a54c36a53c36a52c36a59c306706169726964c36a58c306706169726964c30b4d617463684661696c656459c1681553797374656d2e52756e74696d652e4e6f74696679006c75666203006a5ac3086261736561646472c36a5b527ac46a5ac30b62617365646563696d616cc36a5c527ac46a5ac30971756f746561646472c36a5d527ac46a5ac30c71756f7465646563696d616cc36a5e527ac46a00c30115c36a00c351c3681253797374656d2e53746f726167652e4765746a5f527ac4006a60527ac46a5fc3007da07c75643d016a54c3006a57527ac46a5ec3516a57c3936a57527ac45a516a57c3936a57527ac46a57c36a00c3659e0b956a55c3956a5fc3956a00c30113c36a00c30116c3956a00c30114c395966a60527ac4006a57527ac4526a60c395516a57c3936a57527ac46a5dc3516a57c3936a57527ac46a00c353c3516a57c3936a57527ac46a57c36a00c365e30bf16a55c36a58c3057072696365c37d9f7c75649e006a54c3006a57527ac46a5ec3516a57c3936a57527ac45a516a57c3936a57527ac46a57c36a00c365020b956a58c3057072696365c36a55c394956a5fc3956a00c30113c36a00c30116c3956a00c30114c395966a0111527ac4006a57527ac46a0111c3516a57c3936a57527ac46a5dc3516a57c3936a57527ac46a58c3056f776e6572c3516a57c3936a57527ac46a57c36a00c365370bf1620300620300006a57527ac46a54c3006a57527ac46a5cc3516a57c3936a57527ac45a516a57c3936a57527ac46a57c36a00c3655e0a956a00c30114c396516a57c3936a57527ac46a5bc3516a57c3936a57527ac46a58c3056f776e6572c3516a57c3936a57527ac46a57c36a00c365c40af1006a57527ac46a54c3006a57527ac46a5ec3516a57c3936a57527ac45a516a57c3936a57527ac46a57c36a00c365f109956a55c3956a00c30113c36a00c30114c395966a60c394516a57c3936a57527ac46a5dc3516a57c3936a57527ac46a59c3056f776e6572c3516a57c3936a57527ac46a57c36a00c365480af16a58c306616d6f756e74c3007d9c7c756425006a55c36a58c3096465616c70726963657bc4516a58c3067374617475737bc46203006a59c306616d6f756e74c3007d9c7c756425006a55c36a59c3096465616c70726963657bc4516a59c3067374617475737bc4620300006a57527ac4006a57527ac46a52c3516a57c3936a57527ac46a57c36a00c36519e9516a57c3936a57527ac46a00c356c3516a57c3936a57527ac46a57c36a00c365ae056a0112527ac4006a57527ac4006a57527ac46a53c3516a57c3936a57527ac46a57c36a00c365cfe8516a57c3936a57527ac46a00c356c3516a57c3936a57527ac46a57c36a00c36564056a0113527ac46a59c3681853797374656d2e52756e74696d652e53657269616c697a656a0113c36a00c351c3681253797374656d2e53746f726167652e5075746a58c3681853797374656d2e52756e74696d652e53657269616c697a656a0112c36a00c351c3681253797374656d2e53746f726167652e5075746a56c36a55c36a54c36a59c3056f776e6572c36a53c36a58c3056f776e6572c36a52c36a58c306706169726964c3096465616c4f7264657259c1681553797374656d2e52756e74696d652e4e6f746966796a58c306737461747573c36a58c3096465616c7072696365c36a58c3057072696365c36a58c30474797065c36a58c306616d6f756e74c36a58c306706169726964c36a58c3056f776e6572c36a52c30b7570646174654f7264657259c1681553797374656d2e52756e74696d652e4e6f746966796a59c306737461747573c36a59c3096465616c7072696365c36a59c3057072696365c36a59c30474797065c36a59c306616d6f756e74c36a59c306706169726964c36a59c3056f776e6572c36a53c30b7570646174654f7264657259c1681553797374656d2e52756e74696d652e4e6f74696679516c75665dc56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46a55527ac46a56527ac46a57527ac46a58527ac46203006a00c352c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f16a58c36a57c36a56c36a55c36a54c36a53c36a52c368194f6e746f6c6f67792e436f6e74726163742e4d6967726174656a59527ac46a59c3f10f6d696772617465436f6e74726163746a00c352c3681653797374656d2e52756e74696d652e47657454696d6553c176c9681553797374656d2e52756e74696d652e4e6f74696679516c756657c56b6a00527ac46a51527ac46203006a00c355c36a00c30117c352c66b6a00527ac46a51527ac46c6a52527ac46a52c309616c6c6f77616e63656a00c30118c30068164f6e746f6c6f67792e4e61746976652e496e766f6b656a53527ac46a53c36c75665ac56b6a00527ac46a51527ac4620300006a52527ac46a52c36a00c3657fff6a53527ac46a53c3007d9c7c75640a00006c75666203006a53c36a00c353c36a00c30117c36a00c355c354c66b6a00527ac46a51527ac46a52527ac46a53527ac46c6a54527ac46a54c30c7472616e7366657246726f6d6a00c30118c30068164f6e746f6c6f67792e4e61746976652e496e766f6b656a55527ac46a55c376640d00756a55c301017d9c7c75643e00187769746864726177206f6e67207375636365737366756c2151c176c9681553797374656d2e52756e74696d652e4e6f74696679516c7566623700147769746864726177206f6e67206661696c65642151c176c9681553797374656d2e52756e74696d652e4e6f74696679006c75666c756658c56b6a00527ac46a51527ac46a52527ac4620300006a53527ac4006a53527ac46a52c3516a53c3936a53527ac46a53c36a00c3655fe4516a53c3936a53527ac46a00c356c3516a53c3936a53527ac46a53c36a00c365f4006a54527ac46a54c36a00c351c3681253797374656d2e53746f726167652e4765746a55527ac46a55c3f16a55c3681a53797374656d2e52756e74696d652e446573657269616c697a656c756658c56b6a00527ac46a51527ac46a52527ac4620300006a53527ac4006a53527ac46a52c3516a53c3936a53527ac46a53c36a00c365bae3516a53c3936a53527ac46a00c35ac3516a53c3936a53527ac46a53c36a00c3654f006a54527ac46a54c36a00c351c3681253797374656d2e53746f726167652e4765746a55527ac46a55c3f16a55c3681a53797374656d2e52756e74696d652e446573657269616c697a656c756657c56b6a00527ac46a51527ac46a52527ac46a53527ac46203006a52c3015f7e6a53c37e6c75665ac56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46203006a53c36a00c30118c37d9c7c7576631100756a53c36a00c30117c37d9c7c75f16a52c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f16a54c3007da07c75f16a54c36a00c355c36a52c353c66b6a00527ac46a51527ac46a52527ac46c6a55527ac46a55c351c176c9087472616e736665726a53c30068164f6e746f6c6f67792e4e61746976652e496e766f6b656a56527ac46a56c376640d00756a56c301017d9c7c75f1516c75665ac56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46203006a53c36a00c30118c37d9c7c7576631100756a53c36a00c30117c37d9c7c75f16a54c3007da07c75f16a54c36a52c36a00c355c353c66b6a00527ac46a51527ac46a52527ac46c6a55527ac46a55c351c176c9087472616e736665726a53c30068164f6e746f6c6f67792e4e61746976652e496e766f6b656a56527ac46a56c376640d00756a56c301017d9c7c75f1516c756659c56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46203006a52c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f16a54c3007da07c75f16a52c36a00c355c36a54c353c176c9087472616e736665726a53c36700000000000000000000000000000000000000006a55527ac46a55c36c756659c56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46203006a54c3007da07c75f16a00c355c36a52c36a54c353c176c9087472616e736665726a53c36700000000000000000000000000000000000000006a55527ac46a55c36c756659c56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46203006a53c36a00c30118c37d9c7c7576631100756a53c36a00c30117c37d9c7c75643f00006a55527ac46a54c3516a55c3936a55527ac46a53c3516a55c3936a55527ac46a52c3516a55c3936a55527ac46a55c36a00c36520fd6c7566620300006a55527ac46a54c3516a55c3936a55527ac46a53c3516a55c3936a55527ac46a52c3516a55c3936a55527ac46a55c36a00c36569fe6c75665ec56b6a00527ac46a51527ac46a52527ac46a53527ac46203006a53c3007d9c7c75640a00516c7566620300516a54527ac4006a55527ac4006a56527ac46a53c3516a56c3936a56527ac400516a56c3936a56527ac46a56c36a00c365b2db6a57527ac46a57c3c06a58527ac46a55c36a58c39f6428006a57c36a55c3c36a59527ac46a55c351936a55527ac46a54c36a52c3956a54527ac462d4ff6a54c36c756659c56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46203006a53c36a00c30118c37d9c7c7576631100756a53c36a00c30117c37d9c7c75643f00006a55527ac46a54c3516a55c3936a55527ac46a53c3516a55c3936a55527ac46a52c3516a55c3936a55527ac46a55c36a00c3659bfc6c7566620300006a55527ac46a54c3516a55c3936a55527ac46a53c3516a55c3936a55527ac46a52c3516a55c3936a55527ac46a55c36a00c36595fd6c75665fc56b6a00527ac46a51527ac46203006a00c354c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f16a00c35bc36a00c351c3681253797374656d2e53746f726167652e4765746a52527ac400c176c96a53527ac4006a54527ac4006a55527ac46a52c35193516a55c3936a55527ac451516a55c3936a55527ac46a55c36a00c3652bda6a56527ac46a56c3c06a57527ac46a54c36a57c39f648c016a56c36a54c3c36a58527ac46a54c351936a54527ac4006a55527ac4006a55527ac46a58c3516a55c3936a55527ac46a55c36a00c365cddd516a55c3936a55527ac46a00c35ac3516a55c3936a55527ac46a55c36a00c36562fa6a00c351c3681253797374656d2e53746f726167652e4765746a59527ac46a59c3640b016a59c3681a53797374656d2e52756e74696d652e446573657269616c697a656a5a527ac46a5ac3086261736561646472c36a5b527ac46a5bc36a53c37d0078c0787c9f7664140075527952795279c3876b51936c64eaff516b7575756c7c7591643400006a55527ac46a5bc3516a55c3936a55527ac46a55c36a00c3659300517d9c7c75640d006a53c36a5bc3c86203006203006a5ac30971756f746561646472c36a5c527ac46a5cc36a53c37d0078c0787c9f7664140075527952795279c3876b51936c64eaff516b7575756c7c7591643400006a55527ac46a5cc3516a55c3936a55527ac46a55c36a00c3652200517d9c7c75640d006a53c36a5cc3c86203006203006203006270fe516c75665cc56b6a00527ac46a51527ac46a52527ac46203006a52c36a00c30118c37d9c7c7576631100756a52c36a00c30117c37d9c7c7564d1006a00c355c351c66b6a00527ac46c6a53527ac46a53c30962616c616e63654f666a52c30068164f6e746f6c6f67792e4e61746976652e496e766f6b656a54527ac46a54c3007da07c75647e00006a55527ac46a54c3516a55c3936a55527ac46a52c3516a55c3936a55527ac46a00c354c3516a55c3936a55527ac46a55c36a00c3656ef991643f0013776974686472617741737365744661696c65646a52c36a54c353c176c9681553797374656d2e52756e74696d652e4e6f74696679006c7566620300620300516c75666203006a00c355c351c176c90962616c616e63654f666a52c36700000000000000000000000000000000000000006a54527ac46a54c3007da07c756478006a00c355c36a00c354c36a54c353c176c9087472616e736665726a52c367000000000000000000000000000000000000000091643f0013776974686472617741737365744661696c65646a52c36a54c353c176c9681553797374656d2e52756e74696d652e4e6f74696679006c7566620300620300516c75665ec56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac4620300006a55527ac4006a56527ac46a52c36a00c30118c37d9c7c7576631100756a52c36a00c30117c37d9c7c75640a00516c756662540000c176c907646563696d616c6a52c36700000000000000000000000000000000000000006a55527ac400c176c90673796d626f6c6a52c36700000000000000000000000000000000000000006a56527ac46a55c36a54c37d9e7c7576630e00756a53c36a56c37d9e7c75640a00006c7566620700516c75666c75665dc56b6a00527ac46a51527ac46a52527ac46a53527ac46203006a52c300c36a54527ac4006a55527ac4006a56527ac46a52c3c0516a56c3936a56527ac451516a56c3936a56527ac46a56c36a00c365acd56a57527ac46a57c3c06a58527ac46a55c36a58c39f6430006a57c36a55c3c36a59527ac46a55c351936a55527ac46a54c36a53c37e6a52c36a59c3c37e6a54527ac462ccff6a54c36c7566"
	//
	//newcodeaddr,err := utils.GetContractAddress(newcode)
	//if err != nil{
	//	ctx.LogError("AddressParseFromBytes error: %s", err)
	//	return false
	//}
	//ctx.LogInfo("new address is " + newcodeaddr.ToBase58())
	//ctx.LogInfo("new address is " + newcodeaddr.ToHexString())
	//
	//b,err := hex.DecodeString(newcode)
	//
	//txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
	//	signer,
	//	codeAddress,
	//	[]interface{}{"migrateContract", []interface{}{b,true,"ONTDEX","1.0.1","test","aa@bb.com","ont dex"}})
	//if err != nil {
	//	ctx.LogError("withUnboundONG invest error: %s", err)
	//}
	//
	////WaitForGenerateBlock
	//_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	//if err != nil {
	//	ctx.LogError("withUnboundONG WaitForGenerateBlock error: %s", err)
	//	return false
	//}
	//
	////GetEventLog, to check the result of invoke
	//events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	//if err != nil {
	//	ctx.LogError("withUnboundONG GetSmartContractEvent error:%s", err)
	//	return false
	//}
	//if events.State == 0 {
	//	ctx.LogError("withUnboundONG failed invoked exec state return 0")
	//	return false
	//}
	//for _,notify:= range events.Notify{
	//	ctx.LogInfo("%+v", notify)
	//}
	//
	//ctx.LogInfo("--------------------test migrate end----------------------")
	//
	//ctx.LogInfo("--------------------testing getorder prex --------------------")
	//obj, err =ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(newcodeaddr, []interface{}{"getOrder", []interface{}{1}})
	//if err != nil {
	//	ctx.LogError("unboundong NewNeoVMSInvokeTransaction error:%s", err)
	//
	//	return false
	//}
	//
	//str ,err = obj.Result.ToString()
	//if err != nil{
	//	ctx.LogError("unboundong PrepareInvokeContract error:%s", err)
	//
	//	return false
	//}
	//
	////
	//fmt.Printf("order is %s\n",str)
	//ctx.LogInfo("--------------------testing getorder prex end--------------------")

	return true
}
