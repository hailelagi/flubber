package storage

import (
	"context"
	"testing"

	"github.com/anishathalye/porcupine"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type testWalOperation struct {
	op    string // "begin", "put", "delete", "commit", "rollback"
	id    uint64
	value []byte
}

func TestWalTxBasic(t *testing.T) {
	// todo: mock wal
	/*
		config.SetupTestConfig()
		config := config.GetStorageConfig()
		store := InitObjectStoreClient(config)
	*/

	wal := NewFSWal(&s3.Client{}, "test-bucket", "test")

	walModel := porcupine.Model{
		Init: func() interface{} {
			return map[uint64][]byte{}
		},
		Step: func(stateInt, inputInt, outputInt interface{}) (bool, interface{}) {
			input := inputInt.(testWalOperation)
			state := stateInt.(map[uint64][]byte)
			ctx := context.Background()

			txn := NewWalTxn(ctx, wal)

			switch input.op {
			case "begin":
				err := txn.Begin()
				return err == nil, state

			case "put":
				err := txn.Put(input.id, input.value)
				if err != nil {
					return false, state
				}
				return true, state

			case "delete":
				err := txn.Delete(input.id)
				if err != nil {
					return false, state
				}
				return true, state

			case "commit":
				err := txn.Commit()
				if err != nil {
					return false, state
				}
				newState := make(map[uint64][]byte)
				for k, v := range state {
					newState[k] = v
				}
				newState[input.id] = input.value
				return true, newState

			case "rollback":
				err := txn.Rollback()
				return err == nil, state

			default:
				return false, state
			}
		},
		Equal: func(aInt, bInt interface{}) bool {
			a := aInt.(map[uint64][]byte)
			b := bInt.(map[uint64][]byte)
			if len(a) != len(b) {
				return false
			}
			for k, v := range a {
				if string(b[k]) != string(v) {
					return false
				}
			}
			return true
		},
	}

	ops := []porcupine.Operation{
		{ClientId: 0, Input: testWalOperation{"begin", 1, nil}, Call: 0, Output: map[uint64][]byte{}, Return: 1},
		{ClientId: 0, Input: testWalOperation{"put", 1, []byte("value1")}, Call: 1, Output: map[uint64][]byte{}, Return: 2},
		{ClientId: 0, Input: testWalOperation{"commit", 1, []byte("value1")}, Call: 2, Output: map[uint64][]byte{}, Return: 3},
		{ClientId: 1, Input: testWalOperation{"begin", 2, nil}, Call: 3, Output: map[uint64][]byte{}, Return: 4},
		{ClientId: 1, Input: testWalOperation{"put", 2, []byte("value2")}, Call: 4, Output: map[uint64][]byte{}, Return: 5},
		{ClientId: 1, Input: testWalOperation{"commit", 2, []byte("value2")}, Call: 5, Output: map[uint64][]byte{}, Return: 6},
	}

	res, info := porcupine.CheckOperationsVerbose(walModel, ops, 0)
	if res != porcupine.Ok {
		t.Errorf("wal ops not linearizable: %v", info)
	}
}

func TestConcurrentWALOperations(t *testing.T) {
	// todo: mock wal
	/*
		config.SetupTestConfig()
		config := config.GetStorageConfig()
		store := InitObjectStoreClient(config)
	*/

	wal := NewFSWal(&s3.Client{}, "test-bucket", "test")

	walModel := porcupine.Model{
		Init: func() interface{} {
			return map[uint64][]byte{}
		},
		Step: func(stateInt, inputInt, outputInt interface{}) (bool, interface{}) {
			input := inputInt.(testWalOperation)
			state := stateInt.(map[uint64][]byte)
			ctx := context.Background()
			txn := NewWalTxn(ctx, wal)

			switch input.op {
			case "put":
				err := txn.Put(input.id, input.value)
				if err != nil {
					return false, state
				}
				err = txn.Commit()
				if err != nil {
					return false, state
				}
				newState := make(map[uint64][]byte)
				for k, v := range state {
					newState[k] = v
				}
				newState[input.id] = input.value
				return true, newState

			default:
				return false, state
			}
		},
		Equal: func(aInt, bInt interface{}) bool {
			a := aInt.(map[uint64][]byte)
			b := bInt.(map[uint64][]byte)
			if len(a) != len(b) {
				return false
			}
			for k, v := range a {
				if string(b[k]) != string(v) {
					return false
				}
			}
			return true
		},
	}

	ops := []porcupine.Operation{
		{ClientId: 0, Input: testWalOperation{"put", 1, []byte("value1")}, Call: 0, Output: map[uint64][]byte{}, Return: 2},
		{ClientId: 1, Input: testWalOperation{"put", 1, []byte("value2")}, Call: 1, Output: map[uint64][]byte{}, Return: 3},
	}

	res, info := porcupine.CheckOperationsVerbose(walModel, ops, 0)
	if res != porcupine.Ok {
		t.Errorf("Concurrent WAL operations are not linearizable: %v", info)
	}
}
