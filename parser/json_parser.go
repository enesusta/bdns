package parser

import (
	"encoding/json"

	"github.com/enesusta/bdns/model"
)

func ParseEntities(content []byte) ([]model.DnsEntity, error) {
	var dnsEntities []model.DnsEntity

	err := json.Unmarshal(content, &dnsEntities)

	return dnsEntities, err
}
