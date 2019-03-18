package deploy_invoke

import (
	"bytes"
	"fmt"
	"github.com/ontio/ontology-crypto/ec"
	"github.com/ontio/ontology-crypto/keypair"
	"github.com/ontio/ontology-crypto/sm2"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/smartcontract/service/neovm"
	"github.com/ontio/ontology/vm/neovm/types"
	"io/ioutil"
	"time"
)

func TestCryptoMessage(ctx *testframework.TestFrameworkContext) bool {

	avmfile := "test_data/cryptoMessage.avm"

	code, err := ioutil.ReadFile(avmfile)
	if err != nil {
		return false
	}
	codeHash := common.ToHexString(code)

	codeAddress, _ := utils.GetContractAddress(codeHash)

	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToHexString())
	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToBase58())
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestCryptoMessage GetDefaultAccount error:%s", err)
		return false
	}

	sm2Acct, err := ctx.GetAccount("AQRsgNZ3xXz5dGLLXib23vdbwgbqqNm7cw")
	if err != nil {
		ctx.LogError("get account AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb failed")
		return false
	}
	//tmpkey1 := sm2Acct.PublicKey.(*ec.PublicKey)
	//
	//msg1,err:= sm2.Encrypt(tmpkey1.PublicKey,[]byte("中文,abcdefg"))
	//
	//tmpPrikey1 := sm2Acct.PrivateKey.(*ec.PrivateKey)
	//
	//decryptMsg1,err := sm2.Decrypt(tmpPrikey1.PrivateKey,msg1)
	//
	//fmt.Printf("=========%s\n",decryptMsg1)

	//account2,err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
	//if err != nil{
	//	ctx.LogError("get account AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb failed")
	//	return false
	//}
	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		true,
		codeHash,
		"TestCryptoMessage",
		"1.0",
		"",
		"",
		"",
	)

	if err != nil {
		ctx.LogError("TestCryptoMessage DeploySmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestCryptoMessage WaitForGenerateBlock error: %s", err)
		return false
	}

	ctx.LogInfo("============test register start===========")

	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		sm2Acct,
		codeAddress,
		[]interface{}{"register", []interface{}{sm2Acct.Address[:], keypair.SerializePublicKey(sm2Acct.PublicKey)}})
	if err != nil {
		ctx.LogError("TestCryptoMessage InvokeNeoVMSmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestCryptoMessage WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestCryptoMessage GetSmartContractEvent error:%s", err)
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("============test register end===========")

	ctx.LogInfo("============test sendMessasge start===========")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"sendMessage", []interface{}{signer.Address[:], sm2Acct.Address[:], 0, "中文,abcdefg"}})
	if err != nil {
		ctx.LogError("TestCryptoMessage InvokeNeoVMSmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestCryptoMessage WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestCryptoMessage GetSmartContractEvent error:%s", err)
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}

	ctx.LogInfo("============test sendMessasge end===========")

	ctx.LogInfo("============test getMessageCount start===========")
	obj, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getMessageCount", []interface{}{sm2Acct.Address[:]}})
	if err != nil {
		ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

		return false
	}

	count, err := obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

		return false
	}
	fmt.Printf("message count is %d\n", count)

	ctx.LogInfo("============test getMessageCount  end===========")

	ctx.LogInfo("============test getMessage start ==============")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getMessage", []interface{}{sm2Acct.Address[:], count}})

	bs, err := obj.Result.ToByteArray()
	if err != nil {
		ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("%v\n", bs)
	fmt.Printf("%s\n", bs)

	bf := bytes.NewBuffer(bs)
	stacks, err := neovm.DeserializeStackItem(bf)
	if err != nil {
		ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

		return false
	}
	smap, err := stacks.GetMap()
	if err != nil {
		ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

		return false
	}

	for k, v := range smap {

		key, err := k.GetByteArray()
		if err != nil {
			ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

			return false
		}

		fmt.Printf("key is %s\n", key)
		if string(key) == "ENCRYPT" {
			value, err := v.GetBigInteger()
			if err != nil {
				ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

				return false
			}
			fmt.Printf("ENCRYPT is %d\n", value)
		} else if string(key) == "FROM" {
			value, err := v.GetByteArray()
			if err != nil {
				ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

				return false
			}
			tmpaddr, err := common.AddressParseFromBytes(value)
			if err != nil {
				ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

				return false
			}
			fmt.Printf("FROM is %s\n", tmpaddr.ToBase58())
		} else if string(key) == "MESSAGE" {
			value, err := v.GetByteArray()
			if err != nil {
				ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

				return false
			}
			fmt.Printf("MESSAGE is %s\n", value)
		} else if string(key) == "TIMESTAMP" {
			value, err := v.GetBigInteger()
			if err != nil {
				ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

				return false
			}
			fmt.Printf("TIMESTAMP is %d\n", value)
		}

	}

	ctx.LogInfo("============test getMessage end ==============")

	ctx.LogInfo("============test sendMessasge encrypt start===========")
	//get pubkey first
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getPubkey", []interface{}{sm2Acct.Address[:]}})
	if err != nil {
		ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

		return false
	}

	pubkeybytes, err := obj.Result.ToByteArray()
	if err != nil {
		ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("pubkeybytes is %v\n", pubkeybytes)

	//encrypt a message with pubkey

	pubkey, err := keypair.DeserializePublicKey(pubkeybytes)
	if err != nil {
		ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

		return false
	}

	tmpkey := pubkey.(*ec.PublicKey)

	msg, err := sm2.Encrypt(tmpkey.PublicKey, []byte("中文测试,测试文本2"))

	//send this message
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"sendMessage", []interface{}{signer.Address[:], sm2Acct.Address[:], 1, msg}})
	if err != nil {
		ctx.LogError("TestCryptoMessage InvokeNeoVMSmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestCryptoMessage WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestCryptoMessage GetSmartContractEvent error:%s", err)
		return false
	}
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("============test sendMessasge encrypt end===========")

	ctx.LogInfo("============test getMessageCount start===========")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getMessageCount", []interface{}{sm2Acct.Address[:]}})
	if err != nil {
		ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

		return false
	}

	count, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

		return false
	}
	fmt.Printf("message count is %d\n", count)

	ctx.LogInfo("============test getMessageCount  end===========")

	ctx.LogInfo("============test getMessage encrypt start ==============")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getMessage", []interface{}{sm2Acct.Address[:], count}})

	bs, err = obj.Result.ToByteArray()
	if err != nil {
		ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("%v\n", bs)
	fmt.Printf("%s\n", bs)

	bf = bytes.NewBuffer(bs)
	stacks, err = neovm.DeserializeStackItem(bf)
	if err != nil {
		ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

		return false
	}
	smap, err = stacks.GetMap()
	if err != nil {
		ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

		return false
	}

	for k, v := range smap {

		key, err := k.GetByteArray()
		if err != nil {
			ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

			return false
		}

		fmt.Printf("key is %s\n", key)
		if string(key) == "ENCRYPT" {
			value, err := v.GetBigInteger()
			if err != nil {
				ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

				return false
			}
			fmt.Printf("ENCRYPT is %d\n", value)
		} else if string(key) == "FROM" {
			value, err := v.GetByteArray()
			if err != nil {
				ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

				return false
			}
			tmpaddr, err := common.AddressParseFromBytes(value)
			if err != nil {
				ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

				return false
			}
			fmt.Printf("FROM is %s\n", tmpaddr.ToBase58())
		} else if string(key) == "MESSAGE" {
			value, err := v.GetByteArray()
			if err != nil {
				ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

				return false
			}
			fmt.Printf("MESSAGE is %s\n", value)

			tmpPrikey := sm2Acct.PrivateKey.(*ec.PrivateKey)

			decryptMsg, err := sm2.Decrypt(tmpPrikey.PrivateKey, value)
			fmt.Printf("decrypt message is %s\n", decryptMsg)

		} else if string(key) == "TIMESTAMP" {
			value, err := v.GetBigInteger()
			if err != nil {
				ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

				return false
			}
			fmt.Printf("TIMESTAMP is %d\n", value)
		}

	}
	ctx.LogInfo("============test getMessage encrypt end ==============")

	ctx.LogInfo("============test getRangeMessage start ==============")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getRangeMessage", []interface{}{sm2Acct.Address[:], 1, 5}})

	bs, err = obj.Result.ToByteArray()
	if err != nil {
		ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("%v\n", bs)
	fmt.Printf("%s\n", bs)

	bf = bytes.NewBuffer(bs)
	stacks, err = neovm.DeserializeStackItem(bf)
	if err != nil {
		ctx.LogError("TestCryptoMessage stacks error:%s", err)

		return false
	}
	sarr, err := stacks.GetArray()
	if err != nil {
		ctx.LogError("TestCryptoMessage sarr error:%s", err)

		return false
	}

	for i, m := range sarr {
		fmt.Printf("====message index:%d====\n", i)
		bytesMap, err := m.GetByteArray()
		if err != nil {
			ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

			return false
		}
		tmp, err := neovm.DeserializeStackItem(bytes.NewBuffer(bytesMap))
		if err != nil {
			ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

			return false
		}
		smap, err = tmp.GetMap()
		if err != nil {
			ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

			return false
		}

		encrypt, err := getMapvalue(smap, "ENCRYPT").GetBigInteger()
		if err != nil {
			ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

			return false
		}
		fmt.Printf("encrypt is %d\n", encrypt.Int64())

		from, err := getMapvalue(smap, "FROM").GetByteArray()
		if err != nil {
			ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

			return false
		}
		tmpaddr, err := common.AddressParseFromBytes(from)
		if err != nil {
			ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

			return false
		}
		fmt.Printf("from address is %s\n", tmpaddr.ToBase58())

		msg, err := getMapvalue(smap, "MESSAGE").GetByteArray()
		if err != nil {
			ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

			return false
		}

		if int(encrypt.Int64()) == 1 {
			tmpPrikey := sm2Acct.PrivateKey.(*ec.PrivateKey)
			decryptMsg, err := sm2.Decrypt(tmpPrikey.PrivateKey, msg)
			if err != nil {
				ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

				return false
			}
			fmt.Printf("encrypt message is %s\n", decryptMsg)
		} else {
			fmt.Printf("message is %s\n", msg)
		}

		timstamp, err := getMapvalue(smap, "TIMESTAMP").GetBigInteger()
		if err != nil {
			ctx.LogError("TestCryptoMessage PrepareInvokeContract error:%s", err)

			return false
		}
		fmt.Printf("timestamp is %d\n", timstamp.Int64())

	}

	ctx.LogInfo("============test getRangeMessage end ==============")

	return true
}

func getMapvalue(smap map[types.StackItems]types.StackItems, key string) types.StackItems {

	for k, v := range smap {
		skey, err := k.GetByteArray()
		if err != nil {
			fmt.Printf("error:%s\n", err.Error())
			return nil
		}
		if string(skey) == key {
			return v
		}
	}
	return nil

}
