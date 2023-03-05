package main
//
//func checkLogin(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		// Get the session
//		session, err := store.Get(r, "session-name")
//		if err != nil {
//			http.Error(w, "Error getting session", http.StatusInternalServerError)
//			return
//		}
//
//		// Check if the user is logged in
//		email, ok := session.Values["email"].(string)
//		if ok {
//			// User is logged in, set "loggedIn" variable to true
//			ctx := context.WithValue(r.Context(), "loggedIn", true)
//			next.ServeHTTP(w, r.WithContext(ctx))
//			return
//		}
//
//		// User is not logged in, set "loggedIn" variable to false
//		ctx := context.WithValue(r.Context(), "loggedIn", false)
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}
//
