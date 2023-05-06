package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"
	)
type SmartContract struct {

	contractapi.Contract

}



// Meter describes basic details of what makes up a simple meter

//Insert struct field in alphabetic order => to achieve determinism across languages

// golang keeps the order when marshal to json but doesn't order automatically

type Meter struct {

	ID             string `json:"ID"`
	
	H_pk          string `json:"H_pk"`
	
	Sec_param          string `json:"Sec_param"`

}



// InitLedger adds a base set of meters to the ledger

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	meters := []Meter{

		{ID: "meter1", Sec_param: "12345132523452134", H_pk: "qwer12rf4q34r123rq23r2q4r"},

		{ID: "meter2", Sec_param: "63465634562341234", H_pk: "fwrtf2qweq34fadw34feawef"},

		{ID: "meter3", Sec_param: "4235234523454215325", H_pk: "f32wrefq23ewfq24wreq2rwfa"},

		{ID: "meter4", Sec_param: "412341234r53q24234", H_pk: "f234qrq23rfq23rfq23rfaw3r"},

		// {ID: "meter5", Sec_param: "black"5, H_pk: "Adriana"},

		// {ID: "meter6", Sec_param: "white"5, H_pk: "Michel"},

	}



	for _, meter := range meters {

		meterJSON, err := json.Marshal(meter)

		if err != nil {

			return err

		}



		err = ctx.GetStub().PutState(meter.ID, meterJSON)

		if err != nil {

			return fmt.Errorf("failed to put to world state. %v", err)

		}

	}



	return nil

}



// AddMeter issues a new meter to the world state with given details.

func (s *SmartContract) AddMeter(ctx contractapi.TransactionContextInterface, id string, sec_param string, h_pk string) error {

	exists, err := s.MeterExists(ctx, id)

	if err != nil {

		return err

	}

	if exists {

		return fmt.Errorf("the meter %s already exists", id)

	}



	meter := Meter{

		ID:             id,

		Sec_param:          sec_param,

		H_pk:          h_pk,


	}

	meterJSON, err := json.Marshal(meter)

	if err != nil {

		return err

	}
	// fmt.println("message")


	return ctx.GetStub().PutState(id, meterJSON)

}



// ReadMeter returns the meter stored in the world state with given id.

func (s *SmartContract) ReadMeter(ctx contractapi.TransactionContextInterface, id string) (*Meter, error) {

	meterJSON, err := ctx.GetStub().GetState(id)

	if err != nil {

		return nil, fmt.Errorf("failed to read from world state: %v", err)

	}

	if meterJSON == nil {

		return nil, fmt.Errorf("the meter %s does not exist", id)

	}



	var meter Meter

	err = json.Unmarshal(meterJSON, &meter)

	if err != nil {

		return nil, err

	}



	return &meter, nil

}



// UpdateMeter updates an existing meter in the world state with provided parameters.

func (s *SmartContract) UpdateMeter(ctx contractapi.TransactionContextInterface, id string, sec_param string, h_pk string) error {

	exists, err := s.MeterExists(ctx, id)

	if err != nil {

		return err

	}

	if !exists {

		return fmt.Errorf("the meter %s does not exist", id)

	}



	// overwriting original meter with new meter

	meter := Meter{

		ID:             id,

		Sec_param:          sec_param,


		H_pk:          h_pk,


	}

	meterJSON, err := json.Marshal(meter)

	if err != nil {

		return err

	}



	return ctx.GetStub().PutState(id, meterJSON)

}



// RemoveMeter deletes an given meter from the world state.

func (s *SmartContract) RemoveMeter(ctx contractapi.TransactionContextInterface, id string) error {

	exists, err := s.MeterExists(ctx, id)

	if err != nil {

		return err

	}

	if !exists {

		return fmt.Errorf("the meter %s does not exist", id)

	}



	return ctx.GetStub().DelState(id)

}



// MeterExists returns true when meter with given ID exists in world state

func (s *SmartContract) MeterExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {

	meterJSON, err := ctx.GetStub().GetState(id)

	if err != nil {

		return false, fmt.Errorf("failed to read from world state: %v", err)

	}



	return meterJSON != nil, nil

}



// TransferMeter updates the h_pk field of meter with given id in world state, and returns the old h_pk.





// GetAllMeters returns all meters found in world state

func (s *SmartContract) GetAllMeters(ctx contractapi.TransactionContextInterface) ([]*Meter, error) {

	// range query with empty string for startSec_param and endSec_param does an

	// open-ended query of all meters in the chaincode namespace.

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")

	if err != nil {

		return nil, err

	}

	defer resultsIterator.Close()



	var meters []*Meter

	for resultsIterator.HasNext() {

		queryResponse, err := resultsIterator.Next()

		if err != nil {

			return nil, err

		}



		var meter Meter

		err = json.Unmarshal(queryResponse.Value, &meter)

		if err != nil {

			return nil, err

		}

		meters = append(meters, &meter)

	}



	return meters, nil

}
func main() {
    assetChaincode, err := contractapi.NewChaincode(&SmartContract{})
    if err != nil {
      log.Panicf("Error creating asset-transfer-basic chaincode: %v", err)
    }

    if err := assetChaincode.Start(); err != nil {
      log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
    }
  }