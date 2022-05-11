package testing

import (
	"github.com/nttcom/eclcloud/v3/ecl/network/v2/common_function_pool"
)

const ListResponse = `
{
	"common_function_pools": [
		{
			"description": "Common Function Pool 1",
			"id": "c57066cc-9553-43a6-90de-asfdfesfffff",
			"name": "CF_Pool1"
		},
		{
			"description": "Common Function Pool 2",
			"id": "fesg66cc-9553-43a6-90de-c8472fdsafedf",
			"name": "CF_Pool2"
		}
	]
}
`

const GetResponse = `
{
  "common_function_pool": {
    "description": "Common Function Pool Description",
    "id": "c57066cc-9553-43a6-90de-c847231bc70b",
    "name": "CF_Pool1"
  }
}
`

var CommonFunctionPool1 = common_function_pool.CommonFunctionPool{
	Description: "Common Function Pool 1",
	ID:          "c57066cc-9553-43a6-90de-asfdfesfffff",
	Name:        "CF_Pool1",
}

var CommonFunctionPool2 = common_function_pool.CommonFunctionPool{
	Description: "Common Function Pool 2",
	ID:          "fesg66cc-9553-43a6-90de-c8472fdsafedf",
	Name:        "CF_Pool2",
}

var CommonFunctionDetail = common_function_pool.CommonFunctionPool{
	Description: "Common Function Pool Description",
	ID:          "c57066cc-9553-43a6-90de-c847231bc70b",
	Name:        "CF_Pool1",
}

var ExpectedCommonFunctionPoolSlice = []common_function_pool.CommonFunctionPool{CommonFunctionPool1, CommonFunctionPool2}

const ListResponseDuplicatedNames = `
{
	"common_function_pools": [
		{
			"description": "Common Function Pool Description 1",
			"id": "c57066cc-9553-43a6-90de-asfdfesfffff",
			"name": "CF_Pool1"
		},
		{
			"description": "Common Function Pool Description 2",
			"id": "fesg66cc-9553-43a6-90de-c8472fdsafedf",
			"name": "CF_Pool1"
		}
	]
}
`
