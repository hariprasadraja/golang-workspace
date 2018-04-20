package mgoQueryBuilder

// This file provides utility tools to parse the query params from the request.
// It provides a standardised approach to read the data from request query.

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"regexp"
	"strings"
	"log"
)

// keyOperations is to perform operations on designated key from QueryParamsFilters
// DBKey maps the key to database path or key
// Allowed Values are optional. It is to restrict user request to allow defined values.
// DoBefore is an optional function which is to perform value specific operations.
// eg: If value is a strings which must be treated as bson.ObjectId, we can do such a operations with these data
type keyOperations struct {
	DBPath        string
	AllowedValues []string
	DoBefore      func(string) (interface{}, error)
}

// QueryParamsFilters contains key and it's operations to be performed before converting it to db query.
type QueryParamsFilters map[string]keyOperations

// formValidPair returns a proper key and value to make db query.
// It verifies key and val in parameters to allowedkeys map and returns a valid key and value.
func formValidPair(allowedKeys map[string]keyOperations, key string, val string) (string, interface{}, error) {
	key, val = strings.TrimSpace(key), strings.TrimSpace(val)
	allowedKey, ok := allowedKeys[key]
	if !ok {
		return allowedKey.DBPath,val, fmt.Errorf("key %s is not allowed", key)
	}

	valid := true
	if len(allowedKey.AllowedValues) > 0 {
		valid = false
		for _, value := range allowedKey.AllowedValues {
			if value == val {
				valid = true
			}
		}
	}

	if !valid {
		return key, val, fmt.Errorf("Invalid Values, Only %+v are allowed", allowedKey.AllowedValues)
	}

	if allowedKey.DoBefore != nil {
		val, err := allowedKey.DoBefore(val)
		return allowedKey.DBPath, val, err
	}

	return allowedKey.DBPath, val, nil
}

func makeOptionalQuery(filterParams string, allowedKeys map[string]keyOperations, query bson.M) error {
	pairs := strings.Split(filterParams, ",")
	optionalQuery := make([]bson.M, 0)
	for _, pair := range pairs {
		val := strings.Split(pair, ":")
		if len(val) > 0 {
			key, value, err := formValidPair(allowedKeys, val[0], val[1])
			if err != nil {
				return err
			}

			optionalQuery = append(optionalQuery, bson.M{key: value})
		}
	}

	query["$or"] = optionalQuery
	return nil
}

func makeMandatoryQuery(filterParams string, allowedKeys map[string]keyOperations, query bson.M) (error) {
	pairs := strings.Split(filterParams, ",")
	for _, pair := range pairs {
		val := strings.Split(pair, ":")
		if len(val) > 0 {
			key, value, err := formValidPair(allowedKeys, val[0], val[1])
			if err != nil {
				return err
			}

			query[key] = value
		}
	}

	return nil
}

// ParseRequestForFilterByParams returns a database query based on the inputs parsed from "filterBy" params from the request.
// eg: http://localhost:9000/callcenter/orders?filterBy={optional:{email:hariprasad@benseron.com},mandatory:{"store":"cafe"}}
// After a successful parsing, it will return a query like
//  {
//     "$or" : [{"emailId" : "hariprasad@csmails@gmail.com"}],  // optional filters.
//     "store":"cafe",          // mandatory filters.
//  }
// It is mandatory to have either optional or mandatory fields.
func ParseRequestForFilterByParams(r *http.Request, allowedKeys map[string]keyOperations) (bson.M, error) {
	query := make(bson.M, 0)
	filterBy := r.URL.Query().Get("filterBy")

	// ****** Optional Filters ******
	optionalFilterRegx, _ := regexp.Compile("optional:[^}]*}")
	optionalData := optionalFilterRegx.FindString(filterBy)
	if optionalData != "" {
		optionalData = strings.TrimPrefix(strings.TrimSuffix(optionalData, "}"), "optional:{")
		optionalData = strings.TrimSpace(optionalData)
		if optionalData == "" {
			return query, errors.New("Optional filters must contain any values but got empty.")
		}

		err := makeOptionalQuery(optionalData, allowedKeys, query)
		if err != nil {
			return query, err
		}

		log.Print("Optional Query: ",query)
	}

	// ******  Mandatory Filters ******
	mandatoryFilterRegx, _ := regexp.Compile("mandatory:[^}]*}")
	mandatoryData := mandatoryFilterRegx.FindString(filterBy)
	if mandatoryData != "" {
		mandatoryData = strings.TrimPrefix(strings.TrimSuffix(mandatoryData, "}"), "mandatory:{")
		mandatoryData = strings.TrimSpace(mandatoryData)
		if mandatoryData == "" {
			return query, errors.New("mandatory filters must contain any values but got empty.")
		}

		err := makeMandatoryQuery(mandatoryData, allowedKeys, query)
		if err != nil {
			return query, err
		}

		log.Print("mandatory Query: ",query)
	}

	if filterBy != "" && optionalData == "" && mandatoryData == "" {
		return query, errors.New("unable to read filterBy, need a valid value.")
	}

	return query, nil
}