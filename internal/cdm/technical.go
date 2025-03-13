package cdm

//TIP <p>Order represents the Canonical Data Model for an order</p>

type Envelope struct {
	FlowId   string `json:"flow_id"`
	FlowName string `json:"flow_name"`
}

type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
