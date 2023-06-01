package routeUtil

import (
	"context"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"time"
)

//tcp with http

func Wsock() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := websocket.Accept(w, r, nil)
		if err != nil {
			http.Error(w, "could not handshake websocket", http.StatusBadRequest)
		}
		defer func(ctx *websocket.Conn, code websocket.StatusCode, reason string) {
			err := ctx.Close(code, reason)
			if err != nil {

			}
		}(ctx, websocket.StatusInternalError, "the sky is falling")
		//webrtc reference
		//tcp websocket synack handshake websocket webrtc tcp websocket handhskae udap udcp webocket connection
		c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var rt interface{}
		err = wsjson.Read(c, ctx, &rt)
		if err != nil {
			err := ctx.Close(websocket.StatusInternalError, "read")
			if err != nil {
				return
			}
		}

		log.Println("received: ", rt)
		err = ctx.Close(websocket.StatusNormalClosure, "")
		if err != nil {
			return
		}
	}
}
