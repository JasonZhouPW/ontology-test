package wasmvm

import (
	"github.com/ontio/ontology-test/testframework"
)

func TestWasmVM() {
	//testframework.TFramework.RegTestCase("TestWasmJsonContract", TestWasmJsonContract)
	//testframework.TFramework.RegTestCase("TestWasmRawContract", TestWasmRawContract)
	//testframework.TFramework.RegTestCase("TestCallWasmJsonContract", TestCallWasmJsonContract)
	//testframework.TFramework.RegTestCase("TestAssetContract", TestAssetContract)
	//testframework.TFramework.RegTestCase("TestAssetRawContract", TestAssetRawContract)
	//testframework.TFramework.RegTestCase("TestCallNativeContract", TestCallNativeContract)
	testframework.TFramework.RegTestCase("TestCallNativeContractJson", TestCallNativeContractJson)
}
