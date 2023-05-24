package api

import (
	"context"
	"net/http"
	"strconv"
)

func editorOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, cA := authToken(r)
		w.WriteHeader(cA)
		w.Write([]byte(strconv.Itoa(cA) + ": " + http.StatusText(cA)))
		if claims != nil {
			ctx := context.WithValue(r.Context(), "claims", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
