package schema

import (
	"bytes"
	"github.com/google/jsonapi"
	"time"
)

type Sample struct {
	TitleId   string          `jsonapi:"attr,titleId"`
	SharedId  string          `jsonapi:"attr,sharedId"`
	CreatedAt time.Time       `jsonapi:"attr,createdAt,iso8601"`
	Matches   []*MatchSummary `jsonapi:"relation,matches"`
}

func ParseSamples(jsonStr string) (*Sample, error) {
	in := bytes.NewReader([]byte(jsonStr))
	sample := new(Sample)
	err := jsonapi.UnmarshalPayload(in, sample)
	if err != nil {
		return nil, err
	}

	return sample, nil
}