package request

type PointsRequest struct {
	Verify Verify
	Points Points
}

type Points struct {
	Value int `json:"value"`
}
