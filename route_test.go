package restclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	TestRoute    = NewRoute("/channels/{channel.id}/messages/{message.id}/reactions/{emoji}/@me", "wait")
	APITestRoute = NewAPIRoute(PUT, "/channels/{channel.id}/messages/{message.id}/reactions/{emoji}/@me", "wait")
	CDNTestRoute = NewCDNRoute("/emojis/{emote.id}.", []FileExtension{PNG, GIF}, "size")
)

func TestRoute_Compile(t *testing.T) {
	queryParams := map[string]interface{}{
		"wait": true,
	}
	compiledRoute, err := TestRoute.Compile(queryParams, "test1", "test2", "test3")
	assert.NoError(t, err)
	assert.Equal(t, APIBaseRoute+"/channels/test1/messages/test2/reactions/test3/@me?wait=true", compiledRoute.URL())
}

func TestAPIRoute_Compile(t *testing.T) {
	queryParams := map[string]interface{}{
		"wait": true,
	}
	compiledRoute, err := APITestRoute.Compile(queryParams, "test1", "test2", "test3")
	assert.NoError(t, err)
	assert.Equal(t, APIBaseRoute+"/channels/test1/messages/test2/reactions/test3/@me?wait=true", compiledRoute.URL())
}

func TestCDNRoute_Compile(t *testing.T) {
	queryParams := map[string]interface{}{
		"size": 256,
	}
	compiledRoute, err := CDNTestRoute.Compile(queryParams, PNG, "test1")
	assert.NoError(t, err)
	assert.Equal(t, CDNBaseRoute+"/emojis/test1.png?size=256", compiledRoute.URL())

	compiledRoute, err = CDNTestRoute.Compile(queryParams, GIF, "test1")
	assert.NoError(t, err)
	assert.Equal(t, CDNBaseRoute+"/emojis/test1.gif?size=256", compiledRoute.URL())
}

func TestCustomRoute_Compile(t *testing.T) {
	testAPI := NewCustomRoute(GET, "https://test.de/{test}")

	compiledRoute, err := testAPI.Compile(nil, "test")
	assert.NoError(t, err)
	assert.Equal(t, "https://test.de/test", compiledRoute.URL())
}
