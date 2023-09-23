package test_test

import (
	"bitcoin/address"
	"bitcoin/constant"
	"bitcoin/keypair"
	"bitcoin/scripts"
	"math/big"
	"strings"
	"testing"
)

func TestP2WSH(t *testing.T) {
	// Your test logic for the Add function goes here
	sk1, _ := keypair.NewECPrivateFromWIF("cTALNpTpRbbxTCJ2A5Vq88UxT44w1PE2cYqiB3n4hRvzyCev1Wwo")
	sk2, _ := keypair.NewECPrivateFromWIF("cRvyLwCPLU88jsyj94L7iJjQX5C2f8koG4G2gevN4BeSGcEvfKe9")
	p2wshScript := scripts.NewScript("OP_2", sk1.GetPublic().ToHex(true),
		sk2.GetPublic().ToHex(true),
		"OP_2",
		"OP_CHECKMULTISIG")
	p2wshAddr, _ := address.P2WSHAddresssFromScript(p2wshScript)

	p2pkhAddr := sk1.GetPublic().ToAddress(true)

	txin1 := scripts.NewDefaultTxInput("6e9a0692ed4b3328909d66d41531854988dc39edba5df186affaefda91824e69", 0)
	txout1 := scripts.NewTxOutput(big.NewInt(970000), p2wshAddr.Program().ToScriptPubKey())
	txinSpend := scripts.NewDefaultTxInput("6233aca9f2d6165da2d7b4e35d73b039a22b53f58ce5af87dddee7682be937ea", 0)
	txout2 := scripts.NewTxOutput(big.NewInt(960000), p2pkhAddr.Program().ToScriptPubKey())
	p2wshRedeemScript := p2wshScript
	txinSpendAmount := big.NewInt(970000)

	txin1Multiple := scripts.NewDefaultTxInput("24d949f8c77d7fc0cd09c8d5fccf7a0249178c16170c738da19f6c4b176c9f4b", 0)
	txin2Multiple := scripts.NewDefaultTxInput("65f4d69c91a8de54dc11096eaa315e84ef91a389d1d1c17a691b72095100a3a4", 0)
	txin3Multiple := scripts.NewDefaultTxInput("6c8fc6453a2a3039c2b5b55dcc59587e8b0afa52f92607385b5f4c7e84f38aa2", 0)
	// txin2MultipleAmount := big.NewInt(690000)
	// txin3MultipleAmount := big.NewInt(790000)
	output1Multiple := scripts.NewTxOutput(big.NewInt(100000), p2wshAddr.Program().ToScriptPubKey())
	output2Multiple := scripts.NewTxOutput(big.NewInt(100000), sk1.GetPublic().ToSegwitAddress(true).Program().ToScriptPubKey())
	output3Multiple := scripts.NewTxOutput(big.NewInt(1770000), p2pkhAddr.Program().ToScriptPubKey())
	spendP2pkhResult :=
		"02000000000101ea37e92b68e7dedd87afe58cf5532ba239b0735de3b4d7a25d16d6f2a9ac33620000000000ffffffff0100a60e00000000001976a914fd337ad3bf81e086d96a68e1f8d6a0a510f8c24a88ac040047304402205c88b6c247c6b59e1cc48493b66629b6c011d97b99ecf991b595e891542cf1a802204fa0e3c238818a65adc87a0b2511ba780e4b57ff6c1ba6b27815b1dca7b72c1c01473044022012840e38d61972f32208c23a05c73952cc36503112b0c2250fc8428b1e9c5fe4022051758dc7ce32567e2b71efb9df6dc161c9ec4bc0c2e8116c4228d27810cdb4d70147522102d82c9860e36f15d7b72aa59e29347f951277c21cd4d34822acdeeadbcff8a5462103a2fef1829e0742b89c218c51898d9e7cb9d51201ba2bf9d9e9214ebb6af3270852ae00000000"
	createSendToP2pkhResult :=
		"0200000001694e8291daeffaaf86f15dbaed39dc8849853115d4669d9028334bed92069a6e000000006a473044022038516db4e67c9217b871c690c09f60a57235084f888e23b8ac77ba01d0cba7ae022027a811be50cf54718fc6b88ea900bfa9c8d3e218208fef0e185163e3a47d9a08012102d82c9860e36f15d7b72aa59e29347f951277c21cd4d34822acdeeadbcff8a546ffffffff0110cd0e00000000002200203956f9730cf7275000f4e3faf5db0505b216222c1f7ca1bdfb81a877003fcb9300000000"
	multipleInputMultipleOuputResult := "020000000001034b9f6c174b6c9fa18d730c17168c1749027acffcd5c809cdc07f7dc7f849d924000000006a47304402206932c93458a6ebb85f9fd6f69666cd383a3b8c8d517a096501438840d90493070220544d996a737ca9affda3573635b09e215be1ffddbee9b1260fc3d85d61d90ae5012102d82c9860e36f15d7b72aa59e29347f951277c21cd4d34822acdeeadbcff8a546ffffffffa4a3005109721b697ac1d1d189a391ef845e31aa6e0911dc54dea8919cd6f4650000000000ffffffffa28af3847e4c5f5b380726f952fa0a8b7e5859cc5db5b5c239302a3a45c68f6c0000000000ffffffff03a0860100000000002200203956f9730cf7275000f4e3faf5db0505b216222c1f7ca1bdfb81a877003fcb93a086010000000000160014fd337ad3bf81e086d96a68e1f8d6a0a510f8c24a10021b00000000001976a914fd337ad3bf81e086d96a68e1f8d6a0a510f8c24a88ac00040047304402206503d3610d916835412449f262c8623146503d6f58c9b0343e8d1670b906c4da02200b2b8db13ddc9f157bb95e74c28d273adce49944307aa6a041dba1ed7c528d610147304402207ea74eff48e56f2c0d9afb70b2a90ebf6fcd3ce1e084350f3c061f88dde5eff402203c841f7bf969d04b383ebb1dee4118724bfc9da0260b10f64a0ba7ef3a8d43f00147522102d82c9860e36f15d7b72aa59e29347f951277c21cd4d34822acdeeadbcff8a5462103a2fef1829e0742b89c218c51898d9e7cb9d51201ba2bf9d9e9214ebb6af3270852ae024730440220733fcbd21517a1559e9561668e480ffd0a24b62520cfa16ca7689b20f7f82be402204f053a27f19e0bd1346676c74c65e9e452515bc6510ab307ac3a3fb6d3c89ca7012102d82c9860e36f15d7b72aa59e29347f951277c21cd4d34822acdeeadbcff8a54600000000"
	t.Run("spend_p2pkh_to_p2wsh", func(t *testing.T) {
		tx := scripts.NewBtcTransaction([]*scripts.TxInput{txin1},
			[]*scripts.TxOutput{txout1},
			false)
		digest := tx.GetTransactionDigest(0, p2pkhAddr.Program().ToScriptPubKey(), constant.SIGHASH_ALL)
		signature := sk1.SingInput(digest, constant.SIGHASH_ALL)
		tx.SetScriptSig(0, scripts.NewScript(signature, sk1.GetPublic().ToHex(true)))
		if !strings.EqualFold(tx.Serialize(), createSendToP2pkhResult) {
			t.Errorf("Expected %v, but got %v", createSendToP2pkhResult, tx.Serialize())
		}

	})
	t.Run("spend_from_p2wsh_multisig_to_p2pkh", func(t *testing.T) {
		tx := scripts.NewBtcTransaction([]*scripts.TxInput{txinSpend},
			[]*scripts.TxOutput{txout2},
			true)
		digest := tx.GetTransactionSegwitDigit(0, p2wshRedeemScript, txinSpendAmount)
		sig1 := sk1.SingInput(digest)
		sig2 := sk2.SingInput(digest)
		pk := p2wshRedeemScript.ToHex()
		witness := scripts.NewTxWitnessInput("", sig1, sig2, pk)
		tx.Witnesses = append(tx.Witnesses, witness)
		if !strings.EqualFold(tx.Serialize(), spendP2pkhResult) {
			t.Errorf("Expected %v, but got %v", spendP2pkhResult, tx.Serialize())
		}

	})

	t.Run("spend_multiple_inputs_with_multiple_outputs", func(t *testing.T) {
		tx := scripts.NewBtcTransaction([]*scripts.TxInput{txin1Multiple, txin2Multiple, txin3Multiple},
			[]*scripts.TxOutput{output1Multiple, output2Multiple, output3Multiple},
			true)
		digest := tx.GetTransactionDigest(0, p2pkhAddr.Program().ToScriptPubKey(), constant.SIGHASH_ALL)
		sig1 := sk1.SingInput(digest)
		// utxo inputs 0 is not segwit and we must set secript sig for input 0
		scriptSig := scripts.NewScript(sig1, sk1.GetPublic().ToHex(true))
		tx.SetScriptSig(0, scriptSig)
		// we must set empty witnesses for this input index beacuse input is not segwit but transaction is segwit
		witnesss1 := scripts.NewTxWitnessInput()
		tx.Witnesses = append(tx.Witnesses, witnesss1)
		// next input is pay-to-witness-script-hash we must get transaction segwit deigest from index 1
		segwitDigest1 := tx.GetTransactionSegwitDigit(1, p2wshRedeemScript, big.NewInt(690000))
		// this script is multi sig we must sign with 2 address p2wsh(2-2)
		sigP2sh1 := sk1.SingInput(segwitDigest1)
		sigP2sh2 := sk2.SingInput(segwitDigest1)
		witnesss2 := scripts.NewTxWitnessInput("", sigP2sh1, sigP2sh2, p2wshRedeemScript.ToHex())
		// add witness to transaction
		tx.Witnesses = append(tx.Witnesses, witnesss2)

		segwitDigest2 := tx.GetTransactionSegwitDigit(2, p2pkhAddr.Program().ToScriptPubKey(), big.NewInt(790000))
		sig3 := sk1.SingInput(segwitDigest2)
		witnesss3 := scripts.NewTxWitnessInput(sig3, sk1.GetPublic().ToHex(true))
		// add witness to transaction
		tx.Witnesses = append(tx.Witnesses, witnesss3)
		if !strings.EqualFold(tx.Serialize(), multipleInputMultipleOuputResult) {
			t.Errorf("Expected %v, but got %v", multipleInputMultipleOuputResult, tx.Serialize())
		}

	})
}
