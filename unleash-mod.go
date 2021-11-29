package main

import (
	"net/http"

	"github.com/Unleash/unleash-client-go/v3"
)

func main() {
	unleash.Initialize(
		unleash.WithListener(&unleash.DebugListener{}),
		unleash.WithAppName("my-application"),
		unleash.WithUrl("https://aperture.section.io/ops/featuretoggle/api"),
		unleash.WithCustomHeaders(http.Header{"Authorization": {"https://aperture.section.io/ops/featuretoggle/api"}}),
	)
	unleash.IsEnabled("section.io.hyvorTalk")
}
