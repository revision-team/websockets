package schema

type WsMessage struct {
	Ws      string
	Message []byte
}

type WsValues struct {
	Topic   string
	Payload struct {
		TenantName string
		Payload    struct {
			ProjectId int64 `json:"project_id"`
			Symbol    string
			Currency  string
			// Levels    struct {
			// 	Buy []struct {
			// 		Quantity int64
			// 		Price    float64
			// 	}
			// 	Sell []struct {
			// 		Quantity int64
			// 		Price    float64
			// 	}
			// }
		}
	}
}
