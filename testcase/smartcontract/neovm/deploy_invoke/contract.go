package deploy_invoke

import "github.com/ontio/ontology-go-sdk/utils"

//Smart contract code in hex string after compiled
var contractCode = "54c56b6c766b00527ac46c766b51527ac4616c766b00c30548656c6c6f876c766b52527ac46c766b52c364120061616521006c766b53527ac4620e00006c766b53527ac46203006c766b53c3616c756651c56b6161681953797374656d2e53746f726167652e476574436f6e746578740548656c6c6f05576f726c64615272681253797374656d2e53746f726167652e50757461610568656c6c6f05776f726c64017b615272087472616e7366657254c1681553797374656d2e52756e74696d652e4e6f7469667961033132336c766b00527ac46203006c766b00c3616c7566"

//Address of smart contract
var contractCodeAddress, _ = utils.GetContractAddress(contractCode)

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using Neo.SmartContract.Framework.Services.System;
using System;
using System.ComponentModel;
using System.Numerics;

namespace NeoContract3
{
    public class Contract1 : SmartContract
    {
        public delegate void TransferDel(byte[] from, byte[] to, BigInteger value);

        [DisplayName("transfer")]
        public static event TransferDel Transferred;

        public static Object Main(string operation, params object[] args)
        {
            if (operation == "Hello")
            {
                return Hello();
            }
            return false;

        }

        public static string Hello()
        {
            Storage.Put(Storage.CurrentContext, "Hello", "World");
            Transferred("hello".AsByteArray(), "world".AsByteArray(), 123);
            return "123";
        }
    }
}
*/
