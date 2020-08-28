package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"prototest/goods.ru/prototest/goods_response"
)

func main() {
	rh := &goods_response.ResponseHead{
		Result: 1,
		Errors: []*goods_response.Error{
			&goods_response.Error{
				Code:        400,
				Description: "ddd",
			},
			&goods_response.Error{
				Code:        403,
				Description: "ddd",
			},
		},
	}

	rb := &goods_response.ResponseBody{
		Meta: &goods_response.Meta{
			Num:  12,
			Test: "Test string",
		},
	}

	resp := &goods_response.Response{
		ResponseHead: rh,
		ResponseBody: rb,
	}

	m, _ := json.Marshal(resp)
	x, _ := xml.MarshalIndent(resp, "  ", "    ")

	fmt.Printf("json: %+v\n\n\n", string(m))
	fmt.Printf("xml: %+v\n\n\n", string(x))
}
