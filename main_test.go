package main

import (
	"testing"

	"github.com/google/uuid"
)

func TestHandleGetFooRR(t *testing.T) {
	// This ID must be created first, can't be hard-coded
	id := uuid.MustParse("89faf3cd-fc6e-4e87-b930-00c182cafb05")
	as := Account{}
	accData, err := Get(as, id)

	if err != nil {
		t.Error("error should not be nil")
	}

	if (AccountData{}) == accData {
		t.Error("account data should not be nil")
	}
}
