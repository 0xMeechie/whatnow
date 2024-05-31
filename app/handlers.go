package app

import (
	"context"
	"net/http"

	"github.com/0xMeechie/whatnow/views"
)

func Home(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	page := views.Home()
	page.Render(ctx, w)
}
