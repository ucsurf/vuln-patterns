package p
func h(r *http.Request) {
	// ruleid: templ-safeurl-dynamic
	u := templ.SafeURL(r.URL.Query().Get("next"))
	// ruleid: templ-safeurl-dynamic
	j := template.JS(r.FormValue("cb"))
	// ok: templ-safeurl-dynamic
	k := templ.SafeURL("/home")
	_, _, _ = u, j, k
}
