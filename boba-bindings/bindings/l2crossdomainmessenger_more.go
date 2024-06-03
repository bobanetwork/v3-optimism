// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/bobanetwork/boba/boba-bindings/solc"
)

const L2CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1002,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1003,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_1_0_1600\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1004,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_51_0_20\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_52_0_1568\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1006,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_101_0_1\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1007,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_102_0_1568\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1008,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_151_0_32\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_152_0_1568\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1010,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1011,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1012,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1013,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":1014,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_uint240\"},{\"astId\":1015,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1016,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"otherMessenger\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_contract(CrossDomainMessenger)1018\"},{\"astId\":1017,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"208\",\"type\":\"t_array(t_uint256)43_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)43_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[43]\",\"numberOfBytes\":\"1376\",\"base\":\"t_uint256\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_contract(CrossDomainMessenger)1018\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L2CrossDomainMessengerDeployedBin = "0x60806040526004361061016a5760003560e01c806383a74074116100cb578063b1b1b2091161007f578063d764ad0b11610059578063d764ad0b146103c7578063db505d80146103da578063ecc704281461040757600080fd5b8063b1b1b20914610357578063b28ade2514610387578063c4d66de8146103a757600080fd5b80639fce812c116100b05780639fce812c146102fc578063a4e7f8bd14610327578063a7119869146102fc57600080fd5b806383a74074146102e55780638cbeeef21461020957600080fd5b80634c1d6a69116101225780635644cfdf116101075780635644cfdf146102755780635c975abb1461028b5780636e296e45146102ab57600080fd5b80634c1d6a691461020957806354fd4d501461021f57600080fd5b80632828d7e8116101535780632828d7e8146101b75780633dbb202b146101cc5780633f827a5a146101e157600080fd5b8063028f85f71461016f5780630c568498146101a2575b600080fd5b34801561017b57600080fd5b50610184601081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b3480156101ae57600080fd5b50610184603f81565b3480156101c357600080fd5b50610184604081565b6101df6101da36600461191b565b61046c565b005b3480156101ed57600080fd5b506101f6600181565b60405161ffff9091168152602001610199565b34801561021557600080fd5b50610184619c4081565b34801561022b57600080fd5b506102686040518060400160405280600581526020017f322e312e3000000000000000000000000000000000000000000000000000000081525081565b60405161019991906119ed565b34801561028157600080fd5b5061018461138881565b34801561029757600080fd5b5060005b6040519015158152602001610199565b3480156102b757600080fd5b506102c061079f565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610199565b3480156102f157600080fd5b5061018462030d4081565b34801561030857600080fd5b5060cf5473ffffffffffffffffffffffffffffffffffffffff166102c0565b34801561033357600080fd5b5061029b610342366004611a00565b60ce6020526000908152604090205460ff1681565b34801561036357600080fd5b5061029b610372366004611a00565b60cb6020526000908152604090205460ff1681565b34801561039357600080fd5b506101846103a2366004611a48565b610886565b3480156103b357600080fd5b506101df6103c2366004611b28565b6108f6565b6101df6103d5366004611b45565b610af5565b3480156103e657600080fd5b5060cf546102c09073ffffffffffffffffffffffffffffffffffffffff1681565b34801561041357600080fd5b5061045e60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b604051908152602001610199565b6104746113f6565b1561050c57341561050c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603d60248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f642076616c7565207769746820637573746f6d2067617320746f6b656e00000060648201526084015b60405180910390fd5b60cf54604080516020601f86018190048102820181019092528481526106749273ffffffffffffffffffffffffffffffffffffffff169161056a91908790879081908401838280828437600092019190915250879250610886915050565b347fd764ad0b000000000000000000000000000000000000000000000000000000006105d660cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c6040516024016105f29796959493929190611c14565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611435565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a3385856106f960cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b8660405161070b959493929190611c73565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff215301610869576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f742073657400000000000000000000006064820152608401610503565b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b6000611388619c4080603f6108a2604063ffffffff8816611cf0565b6108ac9190611d20565b601087516108ba9190611cf0565b6108c79062030d40611d6e565b6108d19190611d6e565b6108db9190611d6e565b6108e59190611d6e565b6108ef9190611d6e565b9392505050565b6000547501000000000000000000000000000000000000000000900460ff1615808015610941575060005460017401000000000000000000000000000000000000000090910460ff16105b806109735750303b158015610973575060005474010000000000000000000000000000000000000000900460ff166001145b6109ff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610503565b600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790558015610a8557600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1675010000000000000000000000000000000000000000001790555b610a8e826114c3565b8015610af157600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b60f087901c60028110610bb0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206f722031206d657373616765732061726520737570706f7274656460648201527f20617420746869732074696d6500000000000000000000000000000000000000608482015260a401610503565b8061ffff16600003610ca5576000610c01878986868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508f92506115ff915050565b600081815260cb602052604090205490915060ff1615610ca3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a206c65676163792077697460448201527f6864726177616c20616c72656164792072656c617965640000000000000000006064820152608401610503565b505b6000610ceb898989898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061161e92505050565b9050610d3460cf54337fffffffffffffffffffffffffeeeeffffffffffffffffffffffffffffffffeeef0173ffffffffffffffffffffffffffffffffffffffff90811691161490565b15610d6c57853414610d4857610d48611d9a565b600081815260ce602052604090205460ff1615610d6757610d67611d9a565b610ebe565b3415610e20576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a401610503565b600081815260ce602052604090205460ff16610ebe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c61796564000000000000000000000000000000006064820152608401610503565b610ec787611641565b15610f7a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a401610503565b600081815260cb602052604090205460ff1615611019576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c61796564000000000000000000006064820152608401610503565b61103a8561102b611388619c40611d6e565b67ffffffffffffffff16611696565b1580611060575060cc5473ffffffffffffffffffffffffffffffffffffffff1661dead14155b1561117957600081815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555182917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff3201611172576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d657373616765000000000000000000000000000000000000006064820152608401610503565b50506113d1565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a16179055600061120a88619c405a6111cd9190611dc9565b8988888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506116b492505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055905080156112c057600082815260cb602052604090205460ff161561125d5761125d611d9a565b600082815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a26113cd565b600082815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff32016113cd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d657373616765000000000000000000000000000000000000006064820152608401610503565b5050505b50505050505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6000806114016116ce565b5073ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee141592915050565b6040517fc2b3e5ac0000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000169063c2b3e5ac90849061148b90889088908790600401611de0565b6000604051808303818588803b1580156114a457600080fd5b505af11580156114b8573d6000803e3d6000fd5b505050505050505050565b6000547501000000000000000000000000000000000000000000900460ff1661156e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610503565b60cc5473ffffffffffffffffffffffffffffffffffffffff166115b85760cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead1790555b60cf80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600061160d8585858561175c565b805190602001209050949350505050565b600061162e8787878787876117f5565b8051906020012090509695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff8216301480611690575073ffffffffffffffffffffffffffffffffffffffff8216734200000000000000000000000000000000000016145b92915050565b600080603f83619c4001026040850201603f5a021015949350505050565b600080600080845160208601878a8af19695505050505050565b60008073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16634397dfef6040518163ffffffff1660e01b81526004016040805180830381865afa15801561172f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117539190611e28565b90939092509050565b6060848484846040516024016117759493929190611e68565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcbd4ece9000000000000000000000000000000000000000000000000000000001790529050949350505050565b606086868686868660405160240161181296959493929190611eb2565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b73ffffffffffffffffffffffffffffffffffffffff811681146118b657600080fd5b50565b60008083601f8401126118cb57600080fd5b50813567ffffffffffffffff8111156118e357600080fd5b6020830191508360208285010111156118fb57600080fd5b9250929050565b803563ffffffff8116811461191657600080fd5b919050565b6000806000806060858703121561193157600080fd5b843561193c81611894565b9350602085013567ffffffffffffffff81111561195857600080fd5b611964878288016118b9565b9094509250611977905060408601611902565b905092959194509250565b6000815180845260005b818110156119a85760208185018101518683018201520161198c565b818111156119ba576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006108ef6020830184611982565b600060208284031215611a1257600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060408385031215611a5b57600080fd5b823567ffffffffffffffff80821115611a7357600080fd5b818501915085601f830112611a8757600080fd5b813581811115611a9957611a99611a19565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611adf57611adf611a19565b81604052828152886020848701011115611af857600080fd5b826020860160208301376000602084830101528096505050505050611b1f60208401611902565b90509250929050565b600060208284031215611b3a57600080fd5b81356108ef81611894565b600080600080600080600060c0888a031215611b6057600080fd5b873596506020880135611b7281611894565b95506040880135611b8281611894565b9450606088013593506080880135925060a088013567ffffffffffffffff811115611bac57600080fd5b611bb88a828b016118b9565b989b979a50959850939692959293505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a0830152611c6660c083018486611bcb565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff86168152608060208201526000611ca3608083018688611bcb565b905083604083015263ffffffff831660608301529695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681851681830481118215151615611d1757611d17611cc1565b02949350505050565b600067ffffffffffffffff80841680611d62577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600067ffffffffffffffff808316818516808303821115611d9157611d91611cc1565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600082821015611ddb57611ddb611cc1565b500390565b73ffffffffffffffffffffffffffffffffffffffff8416815267ffffffffffffffff83166020820152606060408201526000611e1f6060830184611982565b95945050505050565b60008060408385031215611e3b57600080fd5b8251611e4681611894565b602084015190925060ff81168114611e5d57600080fd5b809150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525060806040830152611ea16080830185611982565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a0830152611efd60c0830184611982565b9897505050505050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2CrossDomainMessengerStorageLayoutJSON), L2CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2CrossDomainMessenger"] = L2CrossDomainMessengerStorageLayout
	deployedBytecodes["L2CrossDomainMessenger"] = L2CrossDomainMessengerDeployedBin
}
