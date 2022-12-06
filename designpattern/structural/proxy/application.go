package proxy

type Application struct {
}

func (a *Application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "OK"
	}

	if url == "/create/user" && method == "GET" {
		return 201, "User Created"
	}

	return 404, "Not Ok"
}