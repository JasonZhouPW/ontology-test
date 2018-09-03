package deploy_invoke

import (
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology-go-sdk/utils"
	"time"
	"fmt"
	"github.com/ontio/ontology/common"
)


func TestInvokeContract(ctx *testframework.TestFrameworkContext) bool {
	var testcode = "017bc56b6c766b00527ac46c766b51527ac461681953797374656d2e53746f726167652e476574436f6e74657874616c766b52527ac461682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e6753637269707448617368616c766b53527ac46c766b00c30571756572799c64a8006c766b51c361c0519e643e00177175657279206172677320636f756e74206572726f722151c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b51c300c36c766b54527ac46c766b52c36c766b54c37c61681253797374656d2e53746f726167652e476574616c766b55527ac46c766b55c351c161681553797374656d2e52756e74696d652e4e6f746966796162030051616c75666c766b00c30872656769737465729c6488016c766b51c361c0529e6441001a7265676973746572206172677320636f756e74206572726f722151c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b51c300c36c766b54527ac46c766b51c351c36c766b56527ac46c766b56c361681b53797374656d2e52756e74696d652e436865636b5769746e65737361009c643c00154e6f7420612076616c69646520616464726573732151c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b52c36c766b54c37c61681253797374656d2e53746f726167652e476574616c766b57c39c645f006c766b52c36c766b54c36c766b56c3527261681253797374656d2e53746f726167652e5075746111726567697374657220737563636565642151c161681553797374656d2e52756e74696d652e4e6f746966796162030051616c756613616c726561647920726567697374657265642151c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b00c30473656c6c9c640b026c766b51c361c0539e643d001673656c6c206172677320636f756e74206572726f722151c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b51c300c36c766b54527ac46c766b51c351c36c766b56527ac46c766b51c352c36c766b58527ac46c766b56c361681b53797374656d2e52756e74696d652e436865636b5769746e65737361009c643c00154e6f7420612076616c69646520616464726573732151c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b52c36c766b54c37c61681253797374656d2e53746f726167652e476574616c766b55527ac46c766b55c36c766b56c37c616564096c766b59527ac46c766b59c3009c6432000b4e6f74206f776e6572212051c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b52c36c766b54c36c766b53c3527261681253797374656d2e53746f726167652e507574616c766b52c30f4f726967696e616c5f4f776e65725f6c766b54c37e6c766b56c3527261681253797374656d2e53746f726167652e507574616c766b52c30650726963655f6c766b54c37e6c766b58c3527261681253797374656d2e53746f726167652e507574610d53656c6c20737563636565642151c161681553797374656d2e52756e74696d652e4e6f746966796162030051616c75666c766b00c3036275799c64b0036c766b51c361c0539e643c0015627579206172677320636f756e74206572726f722151c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b51c300c36c766b54527ac46c766b51c351c36c766b56527ac46c766b51c352c36c766b58527ac46c766b56c361681b53797374656d2e52756e74696d652e436865636b5769746e65737361009c643c00154e6f7420612076616c69646520616464726573732151c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b52c36c766b54c37c61681253797374656d2e53746f726167652e476574616c766b55527ac46c766b55c36c766b53c37c616550076c766b59527ac46c766b59c3009c643a001375726c206973206e6f7420696e2073616c652051c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b52c30650726963655f6c766b54c37e7c61681253797374656d2e53746f726167652e476574616c766b5a527ac46c766b5ac36c766b58c3a2644800215072696365206973206c6f776572207468616e2063757272656e7420707269636551c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b52c30354505f6c766b54c37e7c61681253797374656d2e53746f726167652e476574616c766b5b527ac46c766b5bc36c766b57c39e645a006c766b53c36c766b5bc36c766b5ac352726165d606009c64400019726566756e6420746f207072656275796572206661696c656451c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b56c361681b53797374656d2e52756e74696d652e436865636b5769746e65737361009c643b0014436865636b5769746e65737320206661696c656451c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b56c36c766b53c36c766b58c3527261651e06519c6466006c766b52c30650726963655f6c766b54c37e6c766b58c3527261681253797374656d2e53746f726167652e507574616c766b52c30354505f6c766b54c37e6c766b56c3527261681253797374656d2e53746f726167652e5075746162030051616c7566157472616e73666572206f6e74206661696c6564212051c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75660d6275792073756363656564212051c161681553797374656d2e52756e74696d652e4e6f746966796162030051616c75666c766b00c30d7175657279546f7050726963659c6414016c766b51c361c0519e6446001f7175657279546f705072696365206172677320636f756e74206572726f722151c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b51c300c36c766b54527ac46c766b52c30650726963655f6c766b54c37e7c61681253797374656d2e53746f726167652e476574616c766b5c527ac46c766b5cc36c766b57c39c641f000052c161681553797374656d2e52756e74696d652e4e6f74696679616c766b52c30354505f6c766b54c37e7c61681253797374656d2e53746f726167652e476574616c766b5d527ac46c766b5dc36c766b5cc352c161681553797374656d2e52756e74696d652e4e6f746966796162030051616c75666c766b00c3046465616c9c643c036c766b51c361c0529e643d00166465616c206172677320636f756e74206572726f722151c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b51c300c36c766b54527ac46c766b51c351c36c766b56527ac46c766b52c30f4f726967696e616c5f4f776e65725f6c766b54c37e7c61681253797374656d2e53746f726167652e476574616c766b5e527ac46c766b5ec36c766b57c39c6446001f6465616c20676574204f726967696e616c5f4f776e6572206661696c65642151c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b5ec36c766b56c37c61657a02009c6441001a6465616c206e6f7420746865206f726967696e206f776e65722151c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b52c30650726963655f6c766b54c37e7c61681253797374656d2e53746f726167652e476574616c766b58527ac46c766b58c36c766b57c39c643d00166465616c20676574207072696365206661696c65642151c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b52c30354505f6c766b54c37e7c61681253797374656d2e53746f726167652e476574616c766b5f527ac46c766b5fc36c766b57c39c643d00166465616c20676574206275796572206661696c65642151c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b53c36c766b5fc36c766b58c352726165d601009c6441001a6465616c207472616e73666572206f6e7420206661696c65642151c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75666c766b52c30950726963655f75726c7c61681553797374656d2e53746f726167652e44656c657465616c766b52c30f4f726967696e616c5f4f776e65725f6c766b54c37e7c61681553797374656d2e53746f726167652e44656c657465616c766b52c30354505f6c766b54c37e7c61681553797374656d2e53746f726167652e44656c657465616c766b52c36c766b54c36c766b5fc3527261681253797374656d2e53746f726167652e5075746162030051616c7566156e6f7420737570706f72746564206d6574686f642151c161681553797374656d2e52756e74696d652e4e6f746966796162030000616c75665ac56b6c766b00527ac46c766b51527ac46c766b00c36c766b52c39c6419006c766b51c36c766b52c39c640b0062030051616c75666c766b00c361c06c766b51c361c09e640b0062030000616c7566006c766b53527ac46c766b53c36c766b00c361c09f6435006c766b00c36c766b53c3c36c766b51c36c766b53c3c39e640b0062030000616c75666c766b53c38b6c766b53527ac462c1ff62030051616c75665ac56b6c766b00527ac46c766b51527ac46c766b52527ac41400000000000000000000000000000000000000016c766b53527ac46153c66b6c6c766b00c3517a6b6c766b00527ac46c6c766b51c3517a6b6c766b51527ac46c6c766b52c3517a6b6c766b52527ac46c6c766b54527ac4516c766b55527ac46c766b54c351c1087472616e736665726c766b53c36c766b55c36168164f6e746f6c6f67792e4e61746976652e496e766f6b65616c766b56527ac46c766b56c36c766b57c39e6418006c766b56c300c301019c640b0062030051616c756662030000616c756652c56b6c766b00527ac462030000616c756651c56b6c766b00527ac46c756655c56b6c766b00527ac46c766b51527ac46c766b52527ac46c766b53527ac46203006c766b54c3616c756651c56b6c766b00527ac46c766b51527ac46203006c766b52c3616c756653c56b6c766b00527ac46c766b51527ac46c766b52527ac46c756652c56b6c766b00527ac46c766b51527ac46c756651c56b6203006c766b00c3616c756651c56b6203006c766b00c3616c756651c56b6203006c766b52c3616c7566"
	codeAddress, _ := utils.GetContractAddress(testcode)
	fmt.Printf("contract addr:%s\n",codeAddress.ToBase58())
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	tx1, err := ctx.Ont.Rpc.DeploySmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		true,
		testcode,
		"TestDomainSmartContract",
		"1.0",
		"",
		"",
		"",
	)

	if err != nil {
		ctx.LogError("TestDomainSmartContract DeploySmartContract error: %s", err)
	}

	fmt.Println(tx1.ToHexString())

	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}

	//t,err:=common.Uint256FromHexString("973bb206682debcf06b370043c504a10723b435881340536bcbf538c11ae9aed")
	//if err != nil {
	//	ctx.LogError("TestDomainSmartContract Uint256FromHexString error:%s", err)
	//
	//	return false
	//}

	//addr ,_:= common.AddressFromBase58("Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT")

	ctx.LogInfo("------Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT preExecuse register 'www.g.cn' start------")
	tx, err := ctx.Ont.Rpc.NewNeoVMSInvokeTransaction(ctx.GetGasPrice(), ctx.GetGasLimit(),codeAddress, []interface{}{"register", []interface{}{"www.g.cn",signer.Address[:]}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}
	err = ctx.Ont.Rpc.SignToTransaction(tx, signer)
	if err != nil {
		ctx.LogError("TestDomainSmartContract SignToTransaction error:%s", err)

		return false
	}


	obj,err:=ctx.Ont.Rpc.PrepareInvokeContract(tx)
	if err != nil {
		ctx.LogError("TestDomainSmartContract PrepareInvokeContract error:%s", err)

		return false
	}
	fmt.Println(obj)
	ctx.LogInfo("------Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT preExecuse register 'www.g.cn' end------")

	ctx.LogInfo("------Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT  register 'www.g.cn' start------")
	thx,err:= ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"register", []interface{}{"www.g.cn",signer.Address[:]}})


	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}


	//svalue, err := ctx.Ont.Rpc.GetStorage(codeAddress, []byte("www.g.cn"))
	//if err != nil {
	//	ctx.LogError("TestDomainSmartContract GetStorageItem key:hello error: %s", err)
	//	return false
	//}
	//
	//ctx.LogInfo("==svalue = %v", string(svalue))

	//GetEventLog, to check the result of invoke
	events, err := ctx.Ont.Rpc.GetSmartContractEvent(thx)
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

	invokeState := notify.States.([]interface{})
	ctx.LogInfo(invokeState)
	s,_  :=common.HexToBytes(invokeState[0].(string))
	ctx.LogInfo("%s", s)
	ctx.LogInfo("------Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT  register 'www.g.cn' end------")


	//first query belongs to register
	ctx.LogInfo("------query 'www.g.cn' owner start------")

	thx,err= ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"query", []interface{}{"www.g.cn"}})


	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}
	events, err = ctx.Ont.Rpc.GetSmartContractEvent(thx)
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}
	notify = events.Notify[0]
	ctx.LogInfo("query %+v", notify)

	invokeState = notify.States.([]interface{})
	ctx.LogInfo(invokeState)
	s,_  =common.HexToBytes(invokeState[0].(string))
	ctx.LogInfo("query %v", s)
	retaddr,_ := common.AddressParseFromBytes(s)
	ctx.LogInfo("owner of 'www.g.cn' is " + retaddr.ToBase58())
	ctx.LogInfo("------query 'www.g.cn' owner end------")


	//sell domain
	ctx.LogInfo("------Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT sell 'www.g.cn'  start------")

	thx,err= ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"sell", []interface{}{"www.g.cn",signer.Address[:],500}})


	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}
	events, err = ctx.Ont.Rpc.GetSmartContractEvent(thx)
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}
	notify = events.Notify[0]
	ctx.LogInfo("sell %+v", notify)

	invokeState = notify.States.([]interface{})
	ctx.LogInfo(invokeState)
	ctx.LogInfo("------Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT sell 'www.g.cn'  end------")



	//second query belongs to contract
	ctx.LogInfo("------query 'www.g.cn' owner start------")

	thx,err= ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"query", []interface{}{"www.g.cn"}})


	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}
	events, err = ctx.Ont.Rpc.GetSmartContractEvent(thx)
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}
	notify = events.Notify[0]
	ctx.LogInfo("query %+v", notify)

	invokeState = notify.States.([]interface{})
	ctx.LogInfo(invokeState)
	s,_  =common.HexToBytes(invokeState[0].(string))
	ctx.LogInfo("query %v", s)
	retaddr,_ = common.AddressParseFromBytes(s)
	ctx.LogInfo(retaddr.ToBase58())
	ctx.LogInfo("------query 'www.g.cn' owner end------")


	ctx.LogInfo("------AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb  register 'www.baidu.com' start------")

	account2,err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
	if err != nil{
		ctx.LogError("get account AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb failed")
		return false
	}

	thx,err= ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"register", []interface{}{"www.baidu.com",account2.Address[:]}})


	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}


	events, err = ctx.Ont.Rpc.GetSmartContractEvent(thx)
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

	invokeState = notify.States.([]interface{})
	ctx.LogInfo(invokeState)
	s,_  =common.HexToBytes(invokeState[0].(string))
	ctx.LogInfo("%s", s)
	ctx.LogInfo("------AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb  register 'www.baidu.com' end------")


	//first query belongs to register
	ctx.LogInfo("------query 'www.baidu.com' owner start------")

	thx,err= ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"query", []interface{}{"www.baidu.com"}})


	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}
	events, err = ctx.Ont.Rpc.GetSmartContractEvent(thx)
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}
	notify = events.Notify[0]
	ctx.LogInfo("query %+v", notify)

	invokeState = notify.States.([]interface{})
	ctx.LogInfo(invokeState)
	s,_  =common.HexToBytes(invokeState[0].(string))
	ctx.LogInfo("query %v", s)
	retaddr,_ = common.AddressParseFromBytes(s)
	ctx.LogInfo(retaddr.ToBase58())
	ctx.LogInfo("------query 'www.baidu.com' owner end------")


	//queryTopPrice
	ctx.LogInfo("------query topprice of  'www.g.cn' start------")

	thx,err= ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"queryTopPrice", []interface{}{"www.g.cn"}})


	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}
	events, err = ctx.Ont.Rpc.GetSmartContractEvent(thx)
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}
	notify = events.Notify[0]
	ctx.LogInfo("query %+v", notify)

	invokeState = notify.States.([]interface{})
	ctx.LogInfo(invokeState)
	s,_  =common.HexToBytes(invokeState[0].(string))
	bi := common.BigIntFromNeoBytes(s)
	ctx.LogInfo("query current price of 'www.g.cn' is %d", bi)
	fmt.Printf("invokeState[1] is %v\n",invokeState[1])

	if len(invokeState[1].(string)) > 0{
		s2,_ := common.HexToBytes(invokeState[1].(string))


		retaddr,_ = common.AddressParseFromBytes(s2)

		ctx.LogInfo("query top price buyer of 'www.g.cn' is %s", retaddr.ToBase58())
		ctx.LogInfo("------query topprice of  'www.g.cn' end------")

	}

	//buy
	ctx.LogInfo("------AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb buy  'www.g.cn' start------")
	thx,err= ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"buy", []interface{}{"www.g.cn",account2.Address[:],int(bi.Int64()) + 10}})
	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}
	events, err = ctx.Ont.Rpc.GetSmartContractEvent(thx)
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}

	for _,n := range events.Notify{
		ctx.LogInfo("query %+v", n)

		invokeState = notify.States.([]interface{})
		ctx.LogInfo(invokeState)
		s,_  =common.HexToBytes(invokeState[0].(string))
		ctx.LogInfo("%s",s)
	}

	ctx.LogInfo("------AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb buy  'www.g.cn' end------")



	//deal the auction
	ctx.LogInfo("------Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT deal the 'www.g.cn' auction start")
	thx,err= ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"deal", []interface{}{"www.g.cn",signer.Address[:]}})

	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}
	events, err = ctx.Ont.Rpc.GetSmartContractEvent(thx)
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}
	notify = events.Notify[0]
	ctx.LogInfo("query %+v", notify)
	invokeState = notify.States.([]interface{})
	ctx.LogInfo(invokeState)
	s,_  =common.HexToBytes(invokeState[0].(string))
	ctx.LogInfo("%s",s)
	ctx.LogInfo("------Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT deal the 'www.g.cn' auction end ")


	//finally ,the domain should belong to AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb
	ctx.LogInfo("------query 'www.g.cn' owner start------")

	thx,err= ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"query", []interface{}{"www.g.cn"}})


	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}
	events, err = ctx.Ont.Rpc.GetSmartContractEvent(thx)
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}
	notify = events.Notify[0]
	ctx.LogInfo("query %+v", notify)

	invokeState = notify.States.([]interface{})
	ctx.LogInfo(invokeState)
	s,_  =common.HexToBytes(invokeState[0].(string))
	ctx.LogInfo("query %v", s)
	retaddr,_ = common.AddressParseFromBytes(s)
	ctx.LogInfo(retaddr.ToBase58())
	ctx.LogInfo("------query 'www.g.cn' owner end------")


	return true
}