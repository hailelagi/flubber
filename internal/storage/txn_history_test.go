package storage

import (
	"context"
	"testing"

	"github.com/anishathalye/porcupine"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type testWalTxn struct {
	op    string
	key   string
	value string
}

func TestTxnCommit(t *testing.T) {
	// todo: mock wal
	/*
		config.SetupTestConfig()
		config := config.GetStorageConfig()
		store := InitObjectStoreClient(config)
	*/

	wal := NewFSWal(&s3.Client{}, "test-bucket", "test")

	kvModel := porcupine.Model{
		Init: func() interface{} {
			return map[string]string{}
		},
		Step: func(stateInt, inputInt, outputInt interface{}) (bool, interface{}) {
			input := inputInt.(testWalTxn)
			state := stateInt.(map[string]string)
			output := outputInt.(map[string]string)
			ctx := context.Background()

			txn := NewTxn(ctx, wal)
			if err := txn.Begin(); err != nil {
				return false, state
			}

			switch input.op {
			case "set":
				err := txn.Put(0, []byte(input.value))
				if err != nil {
					return false, state
				}
				err = txn.Commit()
				if err != nil {
					return false, state
				}
				newState := make(map[string]string)
				for k, v := range state {
					newState[k] = v
				}
				newState[input.key] = input.value
				return true, newState

			case "get":
				entry, err := txn.Get(0)
				if err != nil {
					return false, state
				}
				readCorrectValue := string(entry.buffer) == output[input.key]
				return readCorrectValue, state

			default:
				return false, state
			}
		},
		Equal: func(aInt, bInt interface{}) bool {
			a := aInt.(map[string]string)
			b := bInt.(map[string]string)
			if len(a) != len(b) {
				return false
			}
			for k, v := range a {
				if b[k] != v {
					return false
				}
			}
			return true
		},
	}

	ops := []porcupine.Operation{
		{ClientId: 0, Input: testWalTxn{"set", "a", "100"}, Call: 0, Output: map[string]string{}, Return: 2},
		{ClientId: 1, Input: testWalTxn{"set", "a", "200"}, Call: 1, Output: map[string]string{}, Return: 3},
		{ClientId: 0, Input: testWalTxn{"get", "a", "0"}, Call: 4, Output: map[string]string{"a": "100"}, Return: 6},
		{ClientId: 1, Input: testWalTxn{"get", "a", "0"}, Call: 5, Output: map[string]string{"a": "100"}, Return: 7},
	}

	res, info := porcupine.CheckOperationsVerbose(kvModel, ops, 0)
	if res != porcupine.Ok {
		t.Errorf("fail basic operations: %v", info)
	}
}
