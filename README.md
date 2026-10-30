# go-ai-router

**Goal:** A multi-model LLM router/proxy in Go. Right now we are at the simplest end-to-end working version with a single endpoint and a single provider (OpenAI). Token budgets, rate limiting, failover, caching and Grafana metrics will be added later.

**Prerequisite:** Go 1.22+ (`go version`). Method-based route patterns like `"GET /healthz"` arrived with this version.

---

## Current state

```
go-ai-router/
├── cmd/go-ai-router/main.go
├── internal/
│   ├── config/config.go
│   ├── provider/{provider.go, openai.go}
│   ├── server/server.go
│   └── router/ ratelimit/ cache/ metrics/   (empty)
├── configs/config.yaml
├── .env (not committed to git), .gitignore, Makefile, go.mod, go.sum
```

**Next steps:**

- **Step 7:** Anthropic provider and simple routing between the two providers
- **Step 8:** Failover and circuit breaker
- **Step 9:** Rate limiter (`golang.org/x/time/rate`)
- **Step 10:** Token budget tracking
- **Step 11:** Caching (in-memory LRU, then Redis)
- **Step 12:** Prometheus metrics and Grafana dashboard

---