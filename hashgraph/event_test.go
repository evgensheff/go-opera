/*
Copyright 2017 Mosaic Networks Ltd

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package hashgraph

import (
	"reflect"
	"testing"
	"time"

	"github.com/arrivets/go-swirlds/crypto"
)

func createDummyEventBody() EventBody {
	body := EventBody{}
	body.Transactions = [][]byte{[]byte("abc"), []byte("def")}
	body.Parents = []string{"self", "other"}
	body.Creator = []byte("public key")
	body.Timestamp = time.Now()
	return body
}

func TestMarshallBody(t *testing.T) {
	body := createDummyEventBody()

	raw, err := body.Marshal()
	if err != nil {
		t.Fatalf("Error marshalling EventBody: %s", err)
	}

	newBody := new(EventBody)
	if err := newBody.Unmarshal(raw); err != nil {
		t.Fatalf("Error unmarshalling EventBody: %s", err)
	}

	if !reflect.DeepEqual(body.Transactions, newBody.Transactions) {
		t.Fatalf("Payloads do not match. Expected %#v, got %#v", body.Transactions, newBody.Transactions)
	}
	if !reflect.DeepEqual(body.Parents, newBody.Parents) {
		t.Fatalf("Parents do not match. Expected %#v, got %#v", body.Parents, newBody.Parents)
	}
	if !reflect.DeepEqual(body.Creator, newBody.Creator) {
		t.Fatalf("Creators do not match. Expected %#v, got %#v", body.Creator, newBody.Creator)
	}
	if body.Timestamp != newBody.Timestamp {
		t.Fatalf("Timestamps do not match. Expected %#v, got %#v", body.Timestamp, newBody.Timestamp)
	}

}

func TestSignEvent(t *testing.T) {
	privateKey, _ := crypto.GenerateECDSAKey()
	publicKeyBytes := crypto.FromECDSAPub(&privateKey.PublicKey)

	body := createDummyEventBody()
	body.Creator = publicKeyBytes

	event := Event{Body: body}
	if err := event.Sign(privateKey); err != nil {
		t.Fatalf("Error signing Event: %s", err)
	}

	res, err := event.Verify()
	if err != nil {
		t.Fatalf("Error verifying signature: %s", err)
	}
	if !res {
		t.Fatalf("Verify returned false")
	}
}

func TestMarshallEvent(t *testing.T) {
	privateKey, _ := crypto.GenerateECDSAKey()
	publicKeyBytes := crypto.FromECDSAPub(&privateKey.PublicKey)

	body := createDummyEventBody()
	body.Creator = publicKeyBytes

	event := Event{Body: body}
	if err := event.Sign(privateKey); err != nil {
		t.Fatalf("Error signing Event: %s", err)
	}

	raw, err := event.Marshal()
	if err != nil {
		t.Fatalf("Error marshalling Event: %s", err)
	}

	newEvent := new(Event)
	if err := newEvent.Unmarshal(raw); err != nil {
		t.Fatalf("Error unmarshalling Event: %s", err)
	}

	if !reflect.DeepEqual(*newEvent, event) {
		t.Fatalf("Events are not deeply equal")
	}
}