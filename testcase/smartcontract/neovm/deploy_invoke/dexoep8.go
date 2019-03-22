package deploy_invoke

import (
	"io/ioutil"
	"fmt"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology-go-sdk/utils"
	"time"
	"github.com/ontio/ontology-test/testframework"
	"strconv"
)

func DEXOep8Test(ctx *testframework.TestFrameworkContext) bool {

	ctx.LogInfo("=========== Deploy a sample OEP8 contract ===============")
	//deploy oep8 sample first
	oep8code:="0154c56b6a00527ac46a51527ac46a00c3046e616d659c6424006a51c3c0519e640700006c7566616a51c300c36a52527ac46a52c36545156c7566616a00c30673796d626f6c9c6424006a51c3c0519e640700006c7566616a51c300c36a52527ac46a52c365c2146c7566616a00c30b746f74616c537570706c799c6424006a51c3c0519e640700006c7566616a51c300c36a52527ac46a52c36535146c7566616a00c30962616c616e63654f669c6432006a51c3c0529e640700006c7566616a51c300c36a53527ac46a51c351c36a52527ac46a53c36a52c37c6594136c7566616a00c3087472616e736665729c645f006a51c3c0549e640700006c7566616a51c300c36a54527ac46a51c351c36a55527ac46a51c352c36a52527ac46a51c353c36a56527ac46a54c36a55c36a52c36a56c35379517955727551727552795279547275527275651d116c7566616a00c30d7472616e736665724d756c74699c640c006a51c36535106c7566616a00c307617070726f76659c645f006a51c3c0549e640700006c7566616a51c300c36a57527ac46a51c351c36a58527ac46a51c352c36a52527ac46a51c353c36a56527ac46a57c36a58c36a52c36a56c3537951795572755172755279527954727552727565c60e6c7566616a00c30c617070726f76654d756c74699c640c006a51c365e20d6c7566616a00c309616c6c6f77616e63659c6440006a51c3c0539e640700006c7566616a51c300c36a57527ac46a51c351c36a58527ac46a51c352c36a52527ac46a57c36a58c36a52c35272651f0d6c7566616a00c30c7472616e7366657246726f6d9c646c006a51c3c0559e640700006c7566616a51c300c36a58527ac46a51c351c36a54527ac46a51c352c36a55527ac46a51c353c36a52527ac46a51c354c36a56527ac46a58c36a54c36a55c36a52c36a56c35479517956727551727553795279557275527275656e096c7566616a00c3117472616e7366657246726f6d4d756c74699c640c006a51c36575086c7566616a00c304696e69749c6417006a51c3c0009e640700006c7566616548076c7566616a00c30a62616c616e6365734f669c6424006a51c3c0519e640700006c7566616a51c300c36a53527ac46a53c365dd006c7566616a00c30e746f74616c42616c616e63654f669c6424006a51c3c0519e640700006c7566616a51c300c36a53527ac46a53c36581016c7566616a00c3046d696e749c6432006a51c3c0529e640700006c7566616a51c300c36a52527ac46a51c351c36a59527ac46a52c36a59c37c650c056c756661006c756657c56b6a00527ac4044e616d656a51527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a00c36a51c37c6552077c681253797374656d2e53746f726167652e47657461640700516c756661006c7566006c75665bc56b6a00527ac4034f4e450354574f05544852454504464f5552044649564555c176c96a51527ac40742616c616e63656a52527ac4005152535455c176c96a53527ac400c176c96a54527ac4006a57527ac46a53c3c06a58527ac4616a57c36a58c39f646f006a53c36a57c3c36a55527ac46a57c351936a57527ac46a51c36a55c3c36a56527ac46a54c3681953797374656d2e53746f726167652e476574436f6e74657874616a56c36a52c37c657a066a00c37c6573067c681253797374656d2e53746f726167652e47657461c8628cff6161616a54c36c75665bc56b6a00527ac4034f4e450354574f05544852454504464f5552044649564555c176c96a51527ac40742616c616e63656a52527ac4005152535455c176c96a53527ac4006a54527ac4006a57527ac46a53c3c06a58527ac4616a57c36a58c39f6477006a53c36a57c3c36a55527ac46a57c351936a57527ac46a51c36a55c3c36a56527ac46a54c3681953797374656d2e53746f726167652e476574436f6e74657874616a56c36a52c37c65a1056a00c37c659a057c681253797374656d2e53746f726167652e476574617c6594106a54527ac46284ff6161616a54c36c75660117c56b22416434706a7a3262716570345268517255417a4d755a4a6b424333714a31745a75547514e98f4998d837fcdd44a50561f7f32140c7c6c2606a00527ac4034f4e450354574f05544852454504464f5552044649564555c176c96a51527ac4044e616d656a52527ac40653796d626f6c6a53527ac40742616c616e63656a54527ac40b546f74616c537570706c796a55527ac4005152535455c176c96a56527ac40e546f6b656e4e616d6546697273740f546f6b656e4e616d655365636f6e640e546f6b656e4e616d6554686972640f546f6b656e4e616d65466f757274680e546f6b656e4e616d65466966746855c176c96a57527ac403544e4603544e5303544e4803544e4f03544e4955c176c96a58527ac403a0860103400d0303e0930403801a060320a10755c176c96a59527ac4006a5f527ac46a56c3c06a60527ac4616a5fc36a60c39f6493016a56c36a5fc3c36a5a527ac46a5fc351936a5f527ac46a57c36a5ac3c36a5b527ac46a58c36a5ac3c36a5c527ac46a59c36a5ac3c36a5d527ac46a51c36a5ac3c36a5e527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a5ec36a52c37c65b4036a5bc35272681253797374656d2e53746f726167652e50757461681953797374656d2e53746f726167652e476574436f6e74657874616a5ec36a53c37c6574036a5cc35272681253797374656d2e53746f726167652e50757461681953797374656d2e53746f726167652e476574436f6e74657874616a5ec36a55c37c6534036a5dc35272681253797374656d2e53746f726167652e50757461681953797374656d2e53746f726167652e476574436f6e74657874616a5ec36a54c37c65f4026a00c37c65ed026a5dc35272681253797374656d2e53746f726167652e50757461006a00c36a5ec36a5dc35379517955727551727552795279547275527275087472616e7366657255c1681553797374656d2e52756e74696d652e4e6f746966796268fe616161516c756660c56b6a00527ac46a51527ac422416434706a7a3262716570345268517255417a4d755a4a6b424333714a31745a75547514e98f4998d837fcdd44a50561f7f32140c7c6c2606a52527ac40742616c616e63656a53527ac40b546f74616c537570706c796a54527ac46a52c3681b53797374656d2e52756e74696d652e436865636b5769746e6573736165820d756a00c3656efa65780d756a00c3657e0b6a55527ac46a51c300a065640d756a51c36a55c37c65e70c6a56527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a00c36a54c37c65a8016a56c35272681253797374656d2e53746f726167652e50757461681953797374656d2e53746f726167652e476574436f6e74657874616a00c36a53c37c6568016a52c37c6561016a51c36a52c36a00c37c65870a7c65640c5272681253797374656d2e53746f726167652e50757461006a52c36a00c36a51c35379517955727551727552795279547275527275087472616e7366657255c1681553797374656d2e52756e74696d652e4e6f74696679516c75665cc56b22416434706a7a3262716570345268517255417a4d755a4a6b424333714a31745a75547514e98f4998d837fcdd44a50561f7f32140c7c6c2606a00527ac40b496e697469616c697a65646a51527ac46a00c365db0b75681953797374656d2e53746f726167652e476574436f6e74657874616a51c37c681253797374656d2e53746f726167652e47657461635f0065f5fa6a52527ac46a52c3519c644200681953797374656d2e53746f726167652e476574436f6e74657874616a51c304545255455272681253797374656d2e53746f726167652e50757461516c7566610a696e6974206572726f72f061006c756655c56b6a00527ac46a51527ac46a00c3015f7e6a51c37e6c756659c56b6a00527ac4006a52527ac46a00c3c06a53527ac4616a52c36a53c39f64b0006a00c36a52c3c36a51527ac46a52c351936a52527ac46a51c3c0559e642c00277472616e7366657246726f6d4d756c7469206661696c6564202d20696e707574206572726f7221f0616a51c300c36a51c351c36a51c352c36a51c353c36a51c354c35479517956727551727553795279557275527275653e00009c647aff2a7472616e7366657246726f6d4d756c7469206661696c6564202d207472616e73666572206572726f7221f0624bff616161516c75660120c56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac40742616c616e63656a55527ac407417070726f76656a56527ac46a00c365140a756a00c3653e0a756a51c365370a756a52c365300a756a53c36533f7653d0a756a53c36a55c37c65acfe6a51c37c65a5fe6a57527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a57c37c681253797374656d2e53746f726167652e476574616a58527ac46a58c36a54c3a265e209756a54c300a065d909756a53c36a55c37c6548fe6a52c37c6541fe6a59527ac46a53c36a56c37c6532fe6a51c37c652bfe6a00c37c6524fe6a5a527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a5ac37c681253797374656d2e53746f726167652e476574616a5b527ac46a54c36a5bc3a06437002f796f7520617265206e6f7420616c6c6f77656420746f20776974686472617720746f6f206d616e7920746f6b656e73f0620901616a54c36a5bc39c647d00681953797374656d2e53746f726167652e476574436f6e74657874616a5ac37c681553797374656d2e53746f726167652e44656c65746561681953797374656d2e53746f726167652e476574436f6e74657874616a58c36a54c37c6530087c681553797374656d2e53746f726167652e44656c6574656162840061681953797374656d2e53746f726167652e476574436f6e74657874616a5ac36a5bc36a54c37c65ea075272681253797374656d2e53746f726167652e50757461681953797374656d2e53746f726167652e476574436f6e74657874616a57c36a58c36a54c37c65aa075272681253797374656d2e53746f726167652e5075746161681953797374656d2e53746f726167652e476574436f6e74657874616a59c37c681253797374656d2e53746f726167652e476574616a5c527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a59c36a5cc36a54c37c6551075272681253797374656d2e53746f726167652e507574616a51c36a52c36a53c36a54c35379517955727551727552795279547275527275087472616e7366657255c1681553797374656d2e52756e74696d652e4e6f74696679516c756658c56b6a00527ac46a51527ac46a52527ac407417070726f76656a53527ac46a52c36a53c37c65b7fb6a00c37c65b0fb6a51c37c65a9fb6a54527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a54c37c681253797374656d2e53746f726167652e476574616c756659c56b6a00527ac4006a52527ac46a00c3c06a53527ac4616a52c36a53c39f64a0006a00c36a52c3c36a51527ac46a52c351936a52527ac46a51c3c0549e64270022617070726f76654d756c7469206661696c6564202d20696e707574206572726f7221f0616a51c300c36a51c351c36a51c352c36a51c353c35379517955727551727552795279547275527275653800009c6484ff24617070726f76654d756c7469206661696c6564202d20617070726f7665206572726f7221f0625bff616161516c75660112c56b6a00527ac46a51527ac46a52527ac46a53527ac407417070726f76656a54527ac46a00c365b905756a00c365e305756a51c365dc05756a52c365dff265e905756a00c36a52c37c658b036a55527ac46a55c36a53c3a265cf05756a53c300a065c605756a52c36a54c37c6535fa6a00c37c652efa6a51c37c6527fa6a56527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a56c36a53c35272681253797374656d2e53746f726167652e507574616a00c36a51c36a52c36a53c3537951795572755172755279527954727552727508617070726f76616c55c1681553797374656d2e52756e74696d652e4e6f74696679516c756659c56b6a00527ac4006a52527ac46a00c3c06a53527ac4616a52c36a53c39f64a3006a00c36a52c3c36a51527ac46a52c351936a52527ac46a51c3c0549e642800237472616e736665724d756c7469206661696c6564202d20696e707574206572726f7221f0616a51c300c36a51c351c36a51c352c36a51c353c35379517955727551727552795279547275527275653a00009c6483ff267472616e736665724d756c7469206661696c6564202d207472616e73666572206572726f7221f06258ff616161516c75660118c56b6a00527ac46a51527ac46a52527ac46a53527ac40742616c616e63656a54527ac46a00c365ed03756a52c36521f1652b04756a00c3650d04756a51c3650604756a52c36a54c37c658cf86a55527ac46a55c36a00c37c657df86a56527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a56c37c681253797374656d2e53746f726167652e476574616a57527ac46a53c36a57c3a0630b006a53c300a164080061006c7566616a53c36a57c39c643e00681953797374656d2e53746f726167652e476574436f6e74657874616a56c37c681553797374656d2e53746f726167652e44656c6574656162440061681953797374656d2e53746f726167652e476574436f6e74657874616a56c36a57c36a53c37c65a9025272681253797374656d2e53746f726167652e50757461616a55c36a51c37c6595f76a58527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a58c37c681253797374656d2e53746f726167652e476574616a59527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a58c36a59c36a53c37c6541025272681253797374656d2e53746f726167652e507574616a00c36a51c36a52c36a53c35379517955727551727552795279547275527275087472616e7366657255c1681553797374656d2e52756e74696d652e4e6f74696679516c756656c56b6a00527ac46a51527ac40742616c616e63656a52527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a51c36a52c37c6590f66a00c37c6589f67c681253797374656d2e53746f726167652e476574616c756655c56b6a00527ac40b546f74616c537570706c796a51527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a00c36a51c37c6531f67c681253797374656d2e53746f726167652e476574616c756655c56b6a00527ac40653796d626f6c6a51527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a00c36a51c37c65def57c681253797374656d2e53746f726167652e476574616c756655c56b6a00527ac4044e616d656a51527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a00c36a51c37c658df57c681253797374656d2e53746f726167652e476574616c756657c56b6a00527ac46a51527ac46a51c300a065e500756a00c36a51c3966a52527ac46a52c36c756659c56b6a00527ac46a51527ac46a00c3009c640700006c7566616a00c36a51c3956a52527ac46a52c36a00c3966a51c39c659e00756a52c36c756656c56b6a00527ac46a51527ac46a00c36a51c3a2658000756a00c36a51c3946c756657c56b6a00527ac46a51527ac46a00c36a51c3936a52527ac46a52c36a00c3a2655200756a52c36c756655c56b6a00527ac46a00c3681b53797374656d2e52756e74696d652e436865636b5769746e65737361651f0075516c756655c56b6a00527ac46a00c3c001149c65080075516c756656c56b6a00527ac46a00c36307006509007561516c756653c56b09f4f4f3f3f2f2f1f100f0006c75665ec56b6a00527ac46a51527ac46a51c36a00c3946a52527ac46a52c3c56a53527ac4006a54527ac46a00c36a55527ac461616a00c36a51c39f6433006a54c36a55c3936a56527ac46a56c36a53c36a54c37bc46a54c351936a54527ac46a55c36a54c3936a00527ac462c8ff6161616a53c36c7566"

	oep8codeaddr , _ := utils.GetContractAddress(oep8code)
	ctx.LogInfo("=====CodeAddress===%s", oep8codeaddr.ToHexString())
	ctx.LogInfo("=====CodeAddress base58===%s", oep8codeaddr.ToBase58())
	priceMultiple := 1000000000
	account2, err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")

	signer, err := ctx.GetDefaultAccount()
	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		true,
		oep8code,
		"TestOEP4Py",
		"1.0",
		"",
		"",
		"",
	)

	if err != nil {
		ctx.LogError("DEXOep8Test DeploySmartContract error: %s", err)
	}
	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("DEXOep8Test WaitForGenerateBlock error: %s", err)
		return false
	}
	ctx.LogInfo("=========== Deploy a sample OEP8 contract end ===============")

	ctx.LogInfo("--------------------testing OEP8 init--------------------")
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		oep8codeaddr,
		[]interface{}{"init", []interface{}{}})
	if err != nil {
		ctx.LogError("DEXOep8Test init error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("DEXOep8Test WaitForGenerateBlock error: %s", err)
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

	ctx.LogInfo("--------------------testing OEP8 init end--------------------")

	ctx.LogInfo("--------------------testing balanceOf owner--------------------")
	obj, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(oep8codeaddr, []interface{}{"balanceOf", []interface{}{signer.Address[:],"ONE"}})
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

	ctx.LogInfo("--------------------testing transfer OEP8 ---------------------------")
	if err != nil {
		ctx.LogError("get account AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb failed")
		return false
	}

	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		oep8codeaddr,
		[]interface{}{"transfer", []interface{}{signer.Address[:], account2.Address[:], "ONE",1}})
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
	ctx.LogInfo("--------------------testing balanceOf acct2 --------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(oep8codeaddr, []interface{}{"balanceOf", []interface{}{account2.Address[:],"ONE"}})
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



	done := false
	if done {
		return true
	}


	avmfile := "test_data/ONTDEX_OEP8.avm"

	code, err := ioutil.ReadFile(avmfile)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	codeHash := common.ToHexString(code)

	codeAddress, _ := utils.GetContractAddress(codeHash)

	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToHexString())
	ctx.LogInfo("=====CodeAddress base58===%s", codeAddress.ToBase58())
	account4, err := ctx.GetAccount("ALerVnMj3eNk9xe8BnQJtoWvwGmY3x4KMi")

	//priceMultiple := 1000000000
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
	ctx.LogInfo("--------------------testing init end--------------------")

	ctx.LogInfo("--------------------testing add pair--------------------")
	//addr, _ = common.AddressFromBase58("APPWgNbWvUdQjQxeN7RduYweH3caaM1LM1")
	ong, _ := common.AddressFromBase58("AFmseVrdL9f9oyCzZefL9tG6UbvhfRZMHJ")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"addPair", []interface{}{oep8codeaddr[:], "ONE", "TokenNameFirst", "ONG", ong[:], 9, 1000000000}})
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


	ctx.LogInfo("--------------------testing add sell OEP8 order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"addOrder", []interface{}{signer.Address[:], 1, 10 , 1, 2 * priceMultiple,account4.Address[:],account4.Address[:]}})
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

			sellorderid = int(n)
		}

	}
	fmt.Printf("sellorderid :%d\n",sellorderid)

	ctx.LogInfo("--------------------testing add order end--------------------")

	ctx.LogInfo("--------------------testing getorder prex --------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getOrder", []interface{}{sellorderid}})
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
	//account3, err := ctx.GetAccount("AK98G45DhmPXg4TFPG1KjftvkEaHbU8SHM")

	ctx.LogInfo("--------------------testing add buy OEP8 order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"addOrder", []interface{}{account2.Address[:], 1, 2, 0, 2 * priceMultiple,account4.Address[:],account4.Address[:]}})
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
	ctx.LogInfo("--------------------testing getorder prex --------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getOrder", []interface{}{buyorderid}})
	if err != nil {
		ctx.LogError("unboundong NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	bs, err = obj.Result.ToString()
	if err != nil {
		ctx.LogError("unboundong PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("order is %s\n", bs)
	ctx.LogInfo("--------------------testing getorder prex end--------------------")

	ctx.LogInfo("========================Deploy OEP4 for platform coin =======================")
	oep4codeHash := "0132c56b6a00527ac46a51527ac46a00c304696e69749c6409006501096c7566616a00c3046e616d659c64090065d8086c7566616a00c30673796d626f6c9c64090065b1086c7566616a00c308646563696d616c739c640900658b086c7566616a00c30b746f74616c537570706c799c640900651d086c7566616a00c30962616c616e63654f669c6424006a51c3c0519e640700006c7566616a51c300c36a52527ac46a52c36576076c7566616a00c3087472616e736665729c6440006a51c3c0539e640700006c7566616a51c300c36a53527ac46a51c351c36a54527ac46a51c352c36a55527ac46a53c36a54c36a55c35272657d056c7566616a00c30d7472616e736665724d756c74699c640c006a51c365c5046c7566616a00c30c7472616e7366657246726f6d9c645f006a51c3c0549e640700006c7566616a51c300c36a56527ac46a51c351c36a53527ac46a51c352c36a54527ac46a51c353c36a55527ac46a56c36a53c36a54c36a55c3537951795572755172755279527954727552727565fd006c7566616a00c307617070726f76659c6440006a51c3c0539e640700006c7566616a51c300c36a57527ac46a51c351c36a56527ac46a51c352c36a55527ac46a57c36a56c36a55c3527265f9026c7566616a00c309616c6c6f77616e63659c6432006a51c3c0529e640700006c7566616a51c300c36a57527ac46a51c351c36a56527ac46a57c36a56c37c650b006c756661006c756658c56b6a00527ac46a51527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a52527ac401026a53527ac46a53c36a00c37e6a51c37e6a54527ac46a52c36a54c37c681253797374656d2e53746f726167652e476574616c7566011fc56b6a00527ac46a51527ac46a52527ac46a53527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a54527ac401016a55527ac401026a56527ac46a00c3c001149e6317006a51c3c001149e630d006a52c3c001149e641a00611461646472657373206c656e677468206572726f72f0616a00c3681b53797374656d2e52756e74696d652e436865636b5769746e65737361009c640700006c7566616a55c36a51c37e6a57527ac46a54c36a57c37c681253797374656d2e53746f726167652e476574616a58527ac46a53c36a58c3a0630b006a53c3009f64080061006c7566616a56c36a51c37e6a00c37e6a59527ac46a54c36a59c37c681253797374656d2e53746f726167652e476574616a5a527ac46a55c36a52c37e6a5b527ac46a53c36a5ac3a0640700006c7566616a53c36a5ac39c6449006a54c36a59c37c681553797374656d2e53746f726167652e44656c657465616a54c36a57c36a58c36a53c3945272681253797374656d2e53746f726167652e50757461624c00616a54c36a59c36a5ac36a53c3945272681253797374656d2e53746f726167652e507574616a54c36a57c36a58c36a53c3945272681253797374656d2e53746f726167652e50757461616a54c36a5bc37c681253797374656d2e53746f726167652e476574616a5c527ac46a54c36a5bc36a5cc36a53c3935272681253797374656d2e53746f726167652e507574616a51c36a52c36a53c35272087472616e7366657254c1681553797374656d2e52756e74696d652e4e6f74696679516c75660111c56b6a00527ac46a51527ac46a52527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a53527ac401026a54527ac46a51c3c001149e630d006a00c3c001149e641a00611461646472657373206c656e677468206572726f72f0616a00c3681b53797374656d2e52756e74696d652e436865636b5769746e65737361009c640700006c7566616a52c36a00c365ba02a0630b006a52c3009f64080061006c7566616a54c36a00c37e6a51c37e6a55527ac46a53c36a55c36a52c35272681253797374656d2e53746f726167652e507574616a00c36a51c36a52c3527208617070726f76616c54c1681553797374656d2e52756e74696d652e4e6f74696679516c756659c56b6a00527ac4006a52527ac46a00c3c06a53527ac4616a52c36a53c39f6473006a00c36a52c3c36a51527ac46a52c351936a52527ac46a51c3c0539e6420001b7472616e736665724d756c746920706172616d73206572726f722ef0616a51c300c36a51c351c36a51c352c35272652900009c64a2ff157472616e736665724d756c7469206661696c65642ef06288ff616161516c75660117c56b6a00527ac46a51527ac46a52527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a53527ac401016a54527ac46a51c3c001149e630d006a00c3c001149e641a00611461646472657373206c656e677468206572726f72f0616a00c3681b53797374656d2e52756e74696d652e436865636b5769746e65737361009c630b006a52c3009f64080061006c7566616a54c36a00c37e6a55527ac46a53c36a55c37c681253797374656d2e53746f726167652e476574616a56527ac46a52c36a56c3a0640700006c7566616a52c36a56c39c6425006a53c36a55c37c681553797374656d2e53746f726167652e44656c65746561622800616a53c36a55c36a56c36a52c3945272681253797374656d2e53746f726167652e50757461616a54c36a51c37e6a57527ac46a53c36a57c37c681253797374656d2e53746f726167652e476574616a58527ac46a53c36a57c36a58c36a52c3935272681253797374656d2e53746f726167652e507574616a00c36a51c36a52c35272087472616e7366657254c1681553797374656d2e52756e74696d652e4e6f74696679516c756658c56b6a00527ac4681953797374656d2e53746f726167652e476574436f6e74657874616a51527ac401016a52527ac46a00c3c001149e6419001461646472657373206c656e677468206572726f72f0616a51c36a52c36a00c37e7c681253797374656d2e53746f726167652e476574616c756655c56b681953797374656d2e53746f726167652e476574436f6e74657874616a00527ac40b546f74616c537570706c796a51527ac46a00c36a51c37c681253797374656d2e53746f726167652e476574616c756654c56b586a00527ac46a00c36c756654c56b034d59546a00527ac46a00c36c756654c56b074d79546f6b656e6a00527ac46a00c36c75660113c56b681953797374656d2e53746f726167652e476574436f6e74657874616a00527ac40400e1f5056a51527ac422416434706a7a3262716570345268517255417a4d755a4a6b424333714a31745a75547514e98f4998d837fcdd44a50561f7f32140c7c6c2606a52527ac40400ca9a3b6a53527ac401016a54527ac40b546f74616c537570706c796a55527ac46a52c3c001149e6432000e4f776e657220696c6c6567616c2151c176c9681553797374656d2e52756e74696d652e4e6f7469667961006c7566616a00c36a55c37c681253797374656d2e53746f726167652e4765746164340014416c726561647920696e697469616c697a656421681553797374656d2e52756e74696d652e4e6f7469667961006c7566616a53c36a51c3956a56527ac46a00c36a55c36a56c35272681253797374656d2e53746f726167652e507574616a00c36a54c36a52c37e6a56c35272681253797374656d2e53746f726167652e50757461006a52c36a56c35272087472616e7366657254c1681553797374656d2e52756e74696d652e4e6f74696679516c7566006c75665ec56b6a00527ac46a51527ac46a51c36a00c3946a52527ac46a52c3c56a53527ac4006a54527ac46a00c36a55527ac461616a00c36a51c39f6433006a54c36a55c3936a56527ac46a56c36a53c36a54c37bc46a54c351936a54527ac46a55c36a54c3936a00527ac462c8ff6161616a53c36c7566"
	oep4codeAddress, _ := utils.GetContractAddress(oep4codeHash)

	ctx.LogInfo("=====CodeAddress===%s", oep4codeAddress.ToHexString())
	ctx.LogInfo("=====CodeAddress base58===%s", oep4codeAddress.ToBase58())


	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		true,
		oep4codeHash,
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
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		oep4codeAddress,
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
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestOEP4Py GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestOEP4Py failed invoked exec state return 0")
		return false
	}
	ctx.LogInfo("--------------------testing init end--------------------")

	ctx.LogInfo("--------------------testing approve to contract---------------------------")


	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		oep4codeAddress,
		[]interface{}{"approve", []interface{}{signer.Address[:], codeAddress[:], 6000000000000}})
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
	fmt.Printf("codeAddress:%v\n", codeAddress.ToBase58())
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(oep4codeAddress, []interface{}{"allowance", []interface{}{signer.Address[:], codeAddress[:]}})
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




	ctx.LogInfo("========================Deploy OEP4 for platform coin end =======================")


	ctx.LogInfo("--------------------testing match order--------------------")
	fmt.Printf("matching:buyorderid:%d,sellorderid:%d\n", buyorderid, sellorderid)
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"matchOrder", []interface{}{buyorderid, sellorderid, 2 , 2 * priceMultiple, 1}})
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

	ctx.LogInfo("--------------------testing balanceOf OEP8 account2--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(oep8codeaddr, []interface{}{"balanceOf", []interface{}{account2.Address[:],"ONE"}})
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


	ctx.LogInfo("--------------------testing add sell OEP8 order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"addOrder", []interface{}{signer.Address[:], 1, 10 , 1, 2 * priceMultiple,account4.Address[:],account4.Address[:]}})
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

			sellorderid = int(n)
		}

	}
	fmt.Printf("sellorderid :%d\n",sellorderid)

	ctx.LogInfo("--------------------testing add order end--------------------")

	ctx.LogInfo("--------------------testing cancel sell ONT order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"cancelOrder", []interface{}{signer.Address[:], sellorderid}})
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

	ctx.LogInfo("--------------------testing add buy OEP8 order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"addOrder", []interface{}{account2.Address[:], 1, 2, 0, 2 * priceMultiple,account4.Address[:],account4.Address[:]}})
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
	fmt.Printf("buyorderid :%d\n",buyorderid)

	ctx.LogInfo("--------------------testing add order end--------------------")
	ctx.LogInfo("--------------------testing cancel sell ONT order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"cancelOrder", []interface{}{account2.Address[:], buyorderid}})
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


	return true
}
