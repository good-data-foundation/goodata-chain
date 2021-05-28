package interfaces

import (
	"fmt"
	reflect "reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	// GoodDataSrvAddr ...
	// TODO: Need be updated for prodction
	GoodDataSrvAddr = common.HexToAddress("2222222222222222222222222222222222222222")
)
var goodDataSrvAbi = `
[
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "uuid",
        "type": "bytes"
      },
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "publicKey",
        "type": "bytes"
      }
    ],
    "name": "DORegistered",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "queryUUID",
        "type": "bytes"
      },
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "doUUIDs",
        "type": "bytes"
      }
    ],
    "name": "DOsOfQuery",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "uuid",
        "type": "bytes"
      },
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "payload",
        "type": "bytes"
      }
    ],
    "name": "LogAdded",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "uuid",
        "type": "bytes"
      },
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "publicKey",
        "type": "bytes"
      }
    ],
    "name": "MPCRegistered",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "queryUUID",
        "type": "bytes"
      },
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "mpcUUIDs",
        "type": "bytes"
      }
    ],
    "name": "MPCsOfQuery",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "predictionUUID",
        "type": "bytes"
      }
    ],
    "name": "PredictionFinished",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "queryUUID",
        "type": "bytes"
      },
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "predictionUUID",
        "type": "bytes"
      }
    ],
    "name": "PredictionSubmitted",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "uuid",
        "type": "bytes"
      },
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "publicKey",
        "type": "bytes"
      }
    ],
    "name": "QCRegistered",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "queryUUID",
        "type": "bytes"
      }
    ],
    "name": "QueryFinished",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "queryUUID",
        "type": "bytes"
      },
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "doUUID",
        "type": "bytes"
      }
    ],
    "name": "QueryFinishedInOneDO",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "queryUUID",
        "type": "bytes"
      },
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "qcUUID",
        "type": "bytes"
      },
      {
        "indexed": false,
        "internalType": "int64",
        "name": "dataSet",
        "type": "int64"
      }
    ],
    "name": "QuerySubmitted",
    "type": "event"
  },
  {
    "inputs": [
      {
        "internalType": "bytes",
        "name": "uuid",
        "type": "bytes"
      },
      {
        "internalType": "bytes",
        "name": "payload",
        "type": "bytes"
      }
    ],
    "name": "addLog",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bytes",
        "name": "queryUUID",
        "type": "bytes"
      },
      {
        "internalType": "bytes",
        "name": "doUUIDs",
        "type": "bytes"
      }
    ],
    "name": "assignQueryToDOs",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bytes",
        "name": "queryUUID",
        "type": "bytes"
      },
      {
        "internalType": "bytes",
        "name": "mpcUUIDs",
        "type": "bytes"
      }
    ],
    "name": "assignQueryToMPCs",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bytes",
        "name": "predictionUUID",
        "type": "bytes"
      }
    ],
    "name": "finishPrediction",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bytes",
        "name": "queryUUID",
        "type": "bytes"
      }
    ],
    "name": "finishQuery",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bytes",
        "name": "queryUUID",
        "type": "bytes"
      },
      {
        "internalType": "bytes",
        "name": "doUUID",
        "type": "bytes"
      }
    ],
    "name": "finishQueryInOneDO",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bytes",
        "name": "uuid",
        "type": "bytes"
      },
      {
        "internalType": "bytes",
        "name": "publicKey",
        "type": "bytes"
      }
    ],
    "name": "registerDO",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bytes",
        "name": "uuid",
        "type": "bytes"
      },
      {
        "internalType": "bytes",
        "name": "publicKey",
        "type": "bytes"
      }
    ],
    "name": "registerMPC",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bytes",
        "name": "uuid",
        "type": "bytes"
      },
      {
        "internalType": "bytes",
        "name": "publicKey",
        "type": "bytes"
      }
    ],
    "name": "registerQC",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bytes",
        "name": "queryUUID",
        "type": "bytes"
      },
      {
        "internalType": "bytes",
        "name": "predictionUUID",
        "type": "bytes"
      }
    ],
    "name": "submitPrediction",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bytes",
        "name": "queryUUID",
        "type": "bytes"
      },
      {
        "internalType": "bytes",
        "name": "qcUUID",
        "type": "bytes"
      },
      {
        "internalType": "int64",
        "name": "dataSet",
        "type": "int64"
      },
      {
        "internalType": "uint256",
        "name": "fees",
        "type": "uint256"
      }
    ],
    "name": "submitQuery",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }
]
`

var (
	GdSrvAbi, _ = abi.JSON(strings.NewReader(goodDataSrvAbi))
)

// GetEthWSClient returns ethclient with websocket
func GetEthWSClient() *ethclient.Client {
	conn, err := ethclient.Dial("ws://localhost:8546")
	if err != nil {
		fmt.Printf("Failed to connect ws port, err: %s\n", err)
		return nil
	}

	return conn
}

// ValidateParamsType validate type of params
// returns nil if valid
// returns non-nil if invalid
func ValidateParamsType(actuals []interface{}, expects []interface{}) error {
	if len(actuals) != len(expects) {
		return fmt.Errorf("params length not equal, expect %d, actual: %d", len(expects), len(actuals))
	}
	i := 0
	for ; i < len(actuals); i++ {
		t1 := reflect.TypeOf(actuals[i])
		t2 := reflect.TypeOf(expects[i])
		if t1.String() != t2.String() {
			return fmt.Errorf("params[%d] type not equal, expect: %T, actual: %T", i, expects[i], actuals[i])
		}
	}
	return nil
}
