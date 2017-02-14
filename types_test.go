package ethrpc

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHexIntUnmarshal(t *testing.T) {
	test := struct {
		ID hexInt `json:"id"`
	}{}

	data := []byte(`{"id": "0x1cc348"}`)
	err := json.Unmarshal(data, &test)

	require.Nil(t, err)
	require.Equal(t, hexInt(1885000), test.ID)
}

func TestHexBigUnmarshal(t *testing.T) {
	test := struct {
		ID hexBig `json:"id"`
	}{}

	data := []byte(`{"id": "0x51248487c7466b7062d"}`)
	err := json.Unmarshal(data, &test)

	require.Nil(t, err)
	b := big.Int{}
	b.SetString("23949082357483433297453", 10)

	require.Equal(t, hexBig(b), test.ID)
}

func TestLogUnmarshal(t *testing.T) {
	data := []byte(`{
        "address": "0xd10e3be2bc8f959bc8c41cf65f60de721cf89adf",
        "topics": ["0x78e4fc71ff7e525b3b4660a76336a2046232fd9bba9c65abb22fa3d07d6e7066"],
        "data": "0x0000000000000000000000000000000000000000000000000000000000000000",
        "blockNumber": "0x7f2cd",
        "blockHash": "0x3757b6efd7f82e3a832f0ec229b2fa36e622033ae7bad76b95763055a69374f7",
        "transactionIndex": "0x1",
        "transactionHash": "0xecd8a21609fa852c08249f6c767b7097481da34b9f8d2aae70067918955b4e69",
        "logIndex": "0x6",
        "removed": false
    }`)

	log := new(Log)
	err := json.Unmarshal(data, log)

	require.Nil(t, err)
	require.Equal(t, "0xd10e3be2bc8f959bc8c41cf65f60de721cf89adf", log.Address)
	require.Equal(t, []string{"0x78e4fc71ff7e525b3b4660a76336a2046232fd9bba9c65abb22fa3d07d6e7066"}, log.Topics)
	require.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000000", log.Data)
	require.Equal(t, 520909, log.BlockNumber)
	require.Equal(t, "0x3757b6efd7f82e3a832f0ec229b2fa36e622033ae7bad76b95763055a69374f7", log.BlockHash)
	require.Equal(t, 1, log.TransactionIndex)
	require.Equal(t, "0xecd8a21609fa852c08249f6c767b7097481da34b9f8d2aae70067918955b4e69", log.TransactionHash)
	require.Equal(t, 6, log.LogIndex)
	require.Equal(t, false, log.Removed)
}

func TestTransactionReceiptUnmarshal(t *testing.T) {
	data := []byte(`{
        "blockHash": "0x3757b6efd7f82e3a832f0ec229b2fa36e622033ae7bad76b95763055a69374f7",
        "blockNumber": "0x7f2cd",
        "contractAddress": null,
        "cumulativeGasUsed": "0x13356",
        "gasUsed": "0x6384",
        "logs": [{
            "address": "0xd10e3be2bc8f959bc8c41cf65f60de721cf89adf",
            "topics": ["0x78e4fc71ff7e525b3b4660a76336a2046232fd9bba9c65abb22fa3d07d6e7066"],
            "data": "0x0000000000000000000000000000000000000000000000000000000000000000",
            "blockNumber": "0x7f2cd",
            "blockHash": "0x3757b6efd7f82e3a832f0ec229b2fa36e622033ae7bad76b95763055a69374f7",
            "transactionIndex": "0x1",
            "transactionHash": "0xecd8a21609fa852c08249f6c767b7097481da34b9f8d2aae70067918955b4e69",
            "logIndex": "0x6",
            "removed": false
        }],
        "logsBloom": "0x00000000000000000000000000000000000000000000000000000000000020000000000000000000000000040000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000",
        "root": "0xe367ea197d629892e7b25ea246fba93cd8ae053d468cc5997a816cc85d660321",
        "transactionHash": "0xecd8a21609fa852c08249f6c767b7097481da34b9f8d2aae70067918955b4e69",
        "transactionIndex": "0x1"
    }`)

	receipt := new(TransactionReceipt)
	err := json.Unmarshal(data, receipt)

	require.Nil(t, err)
	require.Equal(t, 1, len(receipt.Logs))
	require.Equal(t, "0x3757b6efd7f82e3a832f0ec229b2fa36e622033ae7bad76b95763055a69374f7", receipt.BlockHash)
	require.Equal(t, 520909, receipt.BlockNumber)
	require.Equal(t, "", receipt.ContractAddress)
	require.Equal(t, 78678, receipt.CumulativeGasUsed)
	require.Equal(t, 25476, receipt.GasUsed)
	require.Equal(t, "0x00000000000000000000000000000000000000000000000000000000000020000000000000000000000000040000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000", receipt.LogsBloom)
	require.Equal(t, "0xe367ea197d629892e7b25ea246fba93cd8ae053d468cc5997a816cc85d660321", receipt.Root)
	require.Equal(t, "0xecd8a21609fa852c08249f6c767b7097481da34b9f8d2aae70067918955b4e69", receipt.TransactionHash)
	require.Equal(t, 1, receipt.TransactionIndex)

	require.Equal(t, "0xd10e3be2bc8f959bc8c41cf65f60de721cf89adf", receipt.Logs[0].Address)
	require.Equal(t, []string{"0x78e4fc71ff7e525b3b4660a76336a2046232fd9bba9c65abb22fa3d07d6e7066"}, receipt.Logs[0].Topics)
	require.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000000", receipt.Logs[0].Data)
	require.Equal(t, 520909, receipt.Logs[0].BlockNumber)
	require.Equal(t, "0x3757b6efd7f82e3a832f0ec229b2fa36e622033ae7bad76b95763055a69374f7", receipt.Logs[0].BlockHash)
	require.Equal(t, 1, receipt.Logs[0].TransactionIndex)
	require.Equal(t, "0xecd8a21609fa852c08249f6c767b7097481da34b9f8d2aae70067918955b4e69", receipt.Logs[0].TransactionHash)
	require.Equal(t, 6, receipt.Logs[0].LogIndex)
	require.Equal(t, false, receipt.Logs[0].Removed)
}