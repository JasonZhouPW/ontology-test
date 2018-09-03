package native

import (
	"github.com/ontio/ontology-test/testcase/smartcontract/native/ontid"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology-test/testcase/smartcontract/native/auth"
)

func TestNative() {
	testframework.TFramework.RegTestCase("TestOntTransfer", TestOntTransfer)
	testframework.TFramework.RegTestCase("TestWithdrawONG", TestWithdrawONG)
	testframework.TFramework.RegTestCase("TestGlobalParam", TestGlobalParam)
	testframework.TFramework.RegTestCase("TestAuth", auth.TestAuthContract)
	ontid.TestNativeOntID()
}
