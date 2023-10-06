// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/bobanetwork/v3-anchorage/boba-bindings/solc"
)

const L1StandardBridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"spacer_0_2_20\",\"offset\":2,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1003,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"spacer_1_0_20\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1005,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_contract(CrossDomainMessenger)1007\"},{\"astId\":1006,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_array(t_uint256)46_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)46_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[46]\",\"numberOfBytes\":\"1472\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(CrossDomainMessenger)1007\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1StandardBridgeStorageLayout = new(solc.StorageLayout)

var L1StandardBridgeDeployedBin = "0x6080604052600436106101635760003560e01c806387087623116100c0578063a9f9e67511610074578063c4d66de811610059578063c4d66de8146104c5578063c89701a214610421578063e11013dd146104e557600080fd5b8063a9f9e67514610492578063b1a1a882146104b257600080fd5b806391c49bf8116100a557806391c49bf814610421578063927ede2d146104545780639a2ac6d51461047f57600080fd5b806387087623146103bb5780638f601f66146103db57600080fd5b8063540abf731161011757806358a997f6116100fc57806358a997f6146103475780637f46ddb214610367578063838b25201461039b57600080fd5b8063540abf73146102d157806354fd4d50146102f157600080fd5b80631532ec34116101485780631532ec34146102545780631635f5fd146102675780633cb747bf1461027a57600080fd5b80630166a07a1461022157806309fc88431461024157600080fd5b3661021c57333b156101fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084015b60405180910390fd5b61021a333362030d40604051806020016040528060008152506104f8565b005b600080fd5b34801561022d57600080fd5b5061021a61023c36600461248b565b61050b565b61021a61024f36600461253c565b6108d2565b61021a61026236600461258f565b6109a9565b61021a61027536600461258f565b6109bd565b34801561028657600080fd5b506003546102a79073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156102dd57600080fd5b5061021a6102ec366004612602565b610e33565b3480156102fd57600080fd5b5061033a6040518060400160405280600581526020017f312e342e3000000000000000000000000000000000000000000000000000000081525081565b6040516102c891906126ef565b34801561035357600080fd5b5061021a610362366004612702565b610e78565b34801561037357600080fd5b506102a77f000000000000000000000000000000000000000000000000000000000000000081565b3480156103a757600080fd5b5061021a6103b6366004612602565b610f4c565b3480156103c757600080fd5b5061021a6103d6366004612702565b610f91565b3480156103e757600080fd5b506104136103f6366004612785565b600260209081526000928352604080842090915290825290205481565b6040519081526020016102c8565b34801561042d57600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006102a7565b34801561046057600080fd5b5060035473ffffffffffffffffffffffffffffffffffffffff166102a7565b61021a61048d3660046127be565b611065565b34801561049e57600080fd5b5061021a6104ad36600461248b565b6110a7565b61021a6104c036600461253c565b6110b6565b3480156104d157600080fd5b5061021a6104e0366004612821565b611187565b61021a6104f33660046127be565b6112d1565b6105058484348585611314565b50505050565b60035473ffffffffffffffffffffffffffffffffffffffff16331480156105fa5750600354604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116931691636e296e459160048083019260209291908290030181865afa1580156105be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e2919061283e565b73ffffffffffffffffffffffffffffffffffffffff16145b6106ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101f3565b6106b5876114fa565b15610803576106c4878761155c565b610776576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a4016101f3565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590528816906340c10f1990604401600060405180830381600087803b1580156107e657600080fd5b505af11580156107fa573d6000803e3d6000fd5b50505050610885565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a168352929052205461084190849061288a565b73ffffffffffffffffffffffffffffffffffffffff8089166000818152600260209081526040808320948c168352939052919091209190915561088590858561167c565b6108c9878787878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061175092505050565b50505050505050565b333b15610961576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101f3565b6109a43333348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061131492505050565b505050565b6109b685858585856109bd565b5050505050565b60035473ffffffffffffffffffffffffffffffffffffffff1633148015610aac5750600354604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116931691636e296e459160048083019260209291908290030181865afa158015610a70573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a94919061283e565b73ffffffffffffffffffffffffffffffffffffffff16145b610b5e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101f3565b823414610bed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f5374616e646172644272696467653a20616d6f756e742073656e7420646f657360448201527f206e6f74206d6174636820616d6f756e7420726571756972656400000000000060648201526084016101f3565b3073ffffffffffffffffffffffffffffffffffffffff851603610c92576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f207360448201527f656c66000000000000000000000000000000000000000000000000000000000060648201526084016101f3565b60035473ffffffffffffffffffffffffffffffffffffffff90811690851603610d3d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f206d60448201527f657373656e67657200000000000000000000000000000000000000000000000060648201526084016101f3565b610d7f85858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506117de92505050565b6000610d9c855a8660405180602001604052806000815250611851565b905080610e2b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e736665722066616960448201527f6c6564000000000000000000000000000000000000000000000000000000000060648201526084016101f3565b505050505050565b6108c987873388888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061186b92505050565b333b15610f07576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101f3565b610e2b86863333888888888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611bb392505050565b6108c987873388888888888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611bb392505050565b333b15611020576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101f3565b610e2b86863333888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061186b92505050565b61050533858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506104f892505050565b6108c98787878787878761050b565b333b15611145576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101f3565b6109a433338585858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506104f892505050565b600054600390610100900460ff161580156111a9575060005460ff8083169116105b611235576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016101f3565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff83161761010017905561126f82611bc2565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b6105053385348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061131492505050565b8234146113a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f5374616e646172644272696467653a206272696467696e6720455448206d757360448201527f7420696e636c7564652073756666696369656e74204554482076616c7565000060648201526084016101f3565b6113af85858584611ca0565b60035460405173ffffffffffffffffffffffffffffffffffffffff90911690633dbb202b9085907f0000000000000000000000000000000000000000000000000000000000000000907f1635f5fd000000000000000000000000000000000000000000000000000000009061142e908b908b9086908a906024016128a1565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e086901b90921682526114c1929188906004016128ea565b6000604051808303818588803b1580156114da57600080fd5b505af11580156114ee573d6000803e3d6000fd5b50505050505050505050565b6000611526827f1d1d8b6300000000000000000000000000000000000000000000000000000000611d13565b806115565750611556827fec4fc8e300000000000000000000000000000000000000000000000000000000611d13565b92915050565b6000611588837f1d1d8b6300000000000000000000000000000000000000000000000000000000611d13565b15611631578273ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115fc919061283e565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050611556565b8273ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115d8573d6000803e3d6000fd5b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526109a49084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611d36565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b38686866040516117c89392919061292f565b60405180910390a4610e2b868686868686611e42565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f2ac69ee804d9a7a0984249f508dfab7cb2534b465b6ce1580f99a38ba9c5e631848460405161183d92919061296d565b60405180910390a361050584848484611eca565b600080600080845160208601878a8af19695505050505050565b611874876114fa565b156119c257611883878761155c565b611935576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a4016101f3565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015260248201859052881690639dc29fac90604401600060405180830381600087803b1580156119a557600080fd5b505af11580156119b9573d6000803e3d6000fd5b50505050611a56565b6119e473ffffffffffffffffffffffffffffffffffffffff8816863086611f37565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a1683529290522054611a22908490612986565b73ffffffffffffffffffffffffffffffffffffffff8089166000908152600260209081526040808320938b16835292905220555b611a64878787878786611f95565b60035460405173ffffffffffffffffffffffffffffffffffffffff90911690633dbb202b907f0000000000000000000000000000000000000000000000000000000000000000907f0166a07a0000000000000000000000000000000000000000000000000000000090611ae5908b908d908c908c908c908b9060240161299e565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e085901b9092168252611b78929187906004016128ea565b600060405180830381600087803b158015611b9257600080fd5b505af1158015611ba6573d6000803e3d6000fd5b5050505050505050505050565b6108c98787878787878761186b565b600054610100900460ff16611c59576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016101f3565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f35d79ab81f2b2017e19afb5c5571778877782d7a8786f5907f93b0f4702f4f238484604051611cff92919061296d565b60405180910390a361050584848484612023565b6000611d1e83612082565b8015611d2f5750611d2f83836120e6565b9392505050565b6000611d98826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166121b59092919063ffffffff16565b8051909150156109a45780806020019051810190611db691906129f9565b6109a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101f3565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd868686604051611eba9392919061292f565b60405180910390a4505050505050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d8484604051611f2992919061296d565b60405180910390a350505050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526105059085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016116ce565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d039686868660405161200d9392919061292f565b60405180910390a4610e2b8686868686866121cc565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af58484604051611f2992919061296d565b60006120ae827f01ffc9a7000000000000000000000000000000000000000000000000000000006120e6565b801561155657506120df827fffffffff000000000000000000000000000000000000000000000000000000006120e6565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d9150600051905082801561219e575060208210155b80156121aa5750600081115b979650505050505050565b60606121c48484600085612244565b949350505050565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf868686604051611eba9392919061292f565b6060824710156122d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101f3565b73ffffffffffffffffffffffffffffffffffffffff85163b612354576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101f3565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161237d9190612a1b565b60006040518083038185875af1925050503d80600081146123ba576040519150601f19603f3d011682016040523d82523d6000602084013e6123bf565b606091505b50915091506121aa828286606083156123d9575081611d2f565b8251156123e95782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f391906126ef565b73ffffffffffffffffffffffffffffffffffffffff8116811461243f57600080fd5b50565b60008083601f84011261245457600080fd5b50813567ffffffffffffffff81111561246c57600080fd5b60208301915083602082850101111561248457600080fd5b9250929050565b600080600080600080600060c0888a0312156124a657600080fd5b87356124b18161241d565b965060208801356124c18161241d565b955060408801356124d18161241d565b945060608801356124e18161241d565b93506080880135925060a088013567ffffffffffffffff81111561250457600080fd5b6125108a828b01612442565b989b979a50959850939692959293505050565b803563ffffffff8116811461253757600080fd5b919050565b60008060006040848603121561255157600080fd5b61255a84612523565b9250602084013567ffffffffffffffff81111561257657600080fd5b61258286828701612442565b9497909650939450505050565b6000806000806000608086880312156125a757600080fd5b85356125b28161241d565b945060208601356125c28161241d565b935060408601359250606086013567ffffffffffffffff8111156125e557600080fd5b6125f188828901612442565b969995985093965092949392505050565b600080600080600080600060c0888a03121561261d57600080fd5b87356126288161241d565b965060208801356126388161241d565b955060408801356126488161241d565b94506060880135935061265d60808901612523565b925060a088013567ffffffffffffffff81111561250457600080fd5b60005b8381101561269457818101518382015260200161267c565b838111156105055750506000910152565b600081518084526126bd816020860160208601612679565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611d2f60208301846126a5565b60008060008060008060a0878903121561271b57600080fd5b86356127268161241d565b955060208701356127368161241d565b94506040870135935061274b60608801612523565b9250608087013567ffffffffffffffff81111561276757600080fd5b61277389828a01612442565b979a9699509497509295939492505050565b6000806040838503121561279857600080fd5b82356127a38161241d565b915060208301356127b38161241d565b809150509250929050565b600080600080606085870312156127d457600080fd5b84356127df8161241d565b93506127ed60208601612523565b9250604085013567ffffffffffffffff81111561280957600080fd5b61281587828801612442565b95989497509550505050565b60006020828403121561283357600080fd5b8135611d2f8161241d565b60006020828403121561285057600080fd5b8151611d2f8161241d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561289c5761289c61285b565b500390565b600073ffffffffffffffffffffffffffffffffffffffff8087168352808616602084015250836040830152608060608301526128e060808301846126a5565b9695505050505050565b73ffffffffffffffffffffffffffffffffffffffff8416815260606020820152600061291960608301856126a5565b905063ffffffff83166040830152949350505050565b73ffffffffffffffffffffffffffffffffffffffff8416815282602082015260606040820152600061296460608301846126a5565b95945050505050565b8281526040602082015260006121c460408301846126a5565b600082198211156129995761299961285b565b500190565b600073ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a08301526129ed60c08301846126a5565b98975050505050505050565b600060208284031215612a0b57600080fd5b81518015158114611d2f57600080fd5b60008251612a2d818460208701612679565b919091019291505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1StandardBridgeStorageLayoutJSON), L1StandardBridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1StandardBridge"] = L1StandardBridgeStorageLayout
	deployedBytecodes["L1StandardBridge"] = L1StandardBridgeDeployedBin
}
