package rules

import (
	"strings"

	"github.com/Gadzhimurad17/link-checker/internal/infrastructure/dto"
)

const (
	httpPrefix  = "http://"
	httpsPrefix = "https://"
)

func ValidateLinks(links dto.LinkRequest) *dto.LinkRequest {
	linksURL := links.Links

	for index, link := range linksURL {
		if !strings.HasPrefix(link, httpPrefix) &&
			!strings.HasPrefix(link, httpsPrefix) {
			linksURL[index] = httpsPrefix + link
		}
	}

	return &dto.LinkRequest{
		Links: linksURL,
	}
}
