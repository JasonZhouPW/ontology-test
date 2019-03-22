package wasmvm

import "github.com/ontio/ontology-test/testframework"

var filePath = "/home/zhoupw/work/go/src/github.com/ontio/ontology-test/test_data"

func TestWasmVM() {
	//testframework.TFramework.RegTestCase("TestWasmAddTest1", TestWasmAddTest1)
	//testframework.TFramework.RegTestCase("TestWasmOEP4", TestWasmOEP4)
	//testframework.TFramework.RegTestCase("TestStrarr", TestStrarr)
	//testframework.TFramework.RegTestCase("TestCallWasmOEP4", TestWasmCallOEP4)
	//
	//testframework.TFramework.RegTestCase("TestWasmAddressTest", TestWasmAddressTest)
	//testframework.TFramework.RegTestCase("TestWasmCallAddressTest", TestWasmCallAddressTest)
	//testframework.TFramework.RegTestCase("TestWasmCallAddressTest3", TestWasmCallAddressTest3)

	//testframework.TFramework.RegTestCase("TestFloat", TestFloat)
	//testframework.TFramework.RegTestCase("TestWasmIntFlow", TestWasmIntFlow)
	//testframework.TFramework.RegTestCase("TestWasmTransfer", TestWasmTransfer)

	//testframework.TFramework.RegTestCase("TestWasmRawContract", TestWasmRawContract)
	//
	//testframework.TFramework.RegTestCase("TestCallWasmJsonContract", TestCallWasmJsonContract)
	//
	////ICO-Test
	//testframework.TFramework.RegTestCase("TestCallNativeContract", TestCallNativeContract)
	//testframework.TFramework.RegTestCase("TestCallICOContract", TestICOContract)
	//testframework.TFramework.RegTestCase("TestICOContractCollect", TestICOContractCollect)
	//////domain-test
	//testframework.TFramework.RegTestCase("TestDomainContract", TestDomainContract)
	//testframework.TFramework.RegTestCase("TestDomainContract_invoke", TestDomainContract_Invoke)
	//testframework.TFramework.RegTestCase("TestDomainContract_invoke2", TestDomainContract_Invoke2)
	//testframework.TFramework.RegTestCase("TestDomainContract_invoke3", TestDomainContract_Invoke3)
	//
	////call neovm test
	//testframework.TFramework.RegTestCase("TestCallNeoContract", TestCallNeoContract)
	////test blockapi
	//testframework.TFramework.RegTestCase("TestBlockApi", TestBlockApi)
	//
	////test headerapi
	//testframework.TFramework.RegTestCase("TestBlockApi", TestHeaderApi)
	////test transactionapi
	//testframework.TFramework.RegTestCase("TestTransactionApi", TestTransactionApi)
	//testframework.TFramework.RegTestCase("TestNewOEP4", TestNewOEP4)
	//testframework.TFramework.RegTestCase("TestNewHello", TestNewHello)

	//rust contracts
	testframework.TFramework.RegTestCase("TestOEP4C", TestOEP4C)
	//testframework.TFramework.RegTestCase("TestRustOEP4", TestRustOEP4)
	//testframework.TFramework.RegTestCase("TestCWasmAPI", TestCWasmAPI)
	//testframework.TFramework.RegTestCase("TestCWasmAPI2", TestCWasmAPI2)
	//testframework.TFramework.RegTestCase("TestRustOEP4Performance", TestRustOEP4Performance)
	//testframework.TFramework.RegTestCase("TestRustApi", TestRustApi)
}
