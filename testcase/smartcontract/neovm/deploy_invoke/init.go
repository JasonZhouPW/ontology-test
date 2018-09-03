package deploy_invoke

import (
	"github.com/ontio/ontology-test/testframework"
)

func TestDeployInvoke() {
	//testframework.TFramework.RegTestCase("TestDeploySmartContract", TestDeploySmartContract)
	//testframework.TFramework.RegTestCase("TestInvokeSmartContract", TestInvokeSmartContract)
	//testframework.TFramework.RegTestCase("TestDomainSmartContract", TestDomainSmartContract)
	//testframework.TFramework.RegTestCase("TestInvokeContract", TestInvokeContract)
	//testframework.TFramework.RegTestCase("TestInvokeContractPy", TestInvokeContractPy)
	//testframework.TFramework.RegTestCase("TestDomainSmartContractPy", TestDomainSmartContractPy)
	//testframework.TFramework.RegTestCase("TestOEP4Py", TestOEP4Py)
	//testframework.TFramework.RegTestCase("TestStructPy", TestStructPy)
	testframework.TFramework.RegTestCase("TestLottery", TestLottery)
}
