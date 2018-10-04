package ttacross

import (
	"encoding/json"
	"fmt"
	"github.com/mxmCherry/openrtb"
	"github.com/prebid/prebid-server/adapters"
	"github.com/prebid/prebid-server/errortypes"
	"net/http"
)

type TtAcrossAdapter struct {
	endpoint string
}

func (a *TtAcrossAdapter) MakeRequests(request *openrtb.BidRequest) ([]*adapters.RequestData, []error) {
	fmt.Println("MakeRequestsd~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	jsonData, err := json.Marshal(request)

	headers := http.Header{}
	headers.Add("Content-Type", "application/json;charset=utf-8")
	reqData := []*adapters.RequestData{{
		Method:  "POST",
		Uri:     a.endpoint,
		Body:    jsonData,
		Headers: headers,
	}}

	return reqData, []error{err}
}

func (a *TtAcrossAdapter) MakeBids(internalRequest *openrtb.BidRequest, externalRequest *adapters.RequestData, response *adapters.ResponseData) (*adapters.BidderResponse, []error) {
	fmt.Println("make bid~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	if response.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	if response.StatusCode == http.StatusBadRequest {
		return nil, []error{&errortypes.BadInput{
			Message: fmt.Sprintf("Unexpected status code: %d. Run with request.debug = 1 for more info", response.StatusCode),
		}}
	}

	if response.StatusCode != http.StatusOK {
		return nil, []error{&errortypes.BadServerResponse{
			Message: fmt.Sprintf("Unexpected status code: %d. Run with request.debug = 1 for more info", response.StatusCode),
		}}
	}

	var bidResp openrtb.BidResponse
	if err := json.Unmarshal(response.Body, &bidResp); err != nil {
		return nil, []error{err}
	}

	bidResponse := adapters.NewBidderResponseWithBidsCapacity(5)

	for _, sb := range bidResp.SeatBid {
		for i := range sb.Bid {
			bidResponse.Bids = append(bidResponse.Bids, &adapters.TypedBid{
				Bid:     &sb.Bid[i],
				BidType: "banner",
			})
		}
	}
	return bidResponse, nil

}

func NewTtxAcrossBidder(endpoint string) *TtAcrossAdapter {
	return &TtAcrossAdapter{
		endpoint: endpoint,
	}
}
