package ttacross

import (
	"testing"

	"github.com/prebid/prebid-server/adapters/adapterstest"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "ttacross", NewTtxAcrossBidder("http://ssc.33across.com/api/v1/hb"))
}
