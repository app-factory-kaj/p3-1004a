# greeter — PRD

## Problem Statement

Teams building services on this platform need a minimal, known-good reference
for a single-endpoint Go HTTP service — something small enough to stand up
quickly and verify the platform's conventions end-to-end, without any real
business logic getting in the way.

## Solution

Greeter is a small Go HTTP service exposing one endpoint, `GET /hello`, that
returns a JSON greeting for a name supplied as a query parameter. It exists as
a lightweight, conventions-following reference service rather than a
production business application.

## Actors

- **API Consumer** — any client or integration that calls the `/hello`
endpoint. No sign-in or identity is involved; the caller is not a named
product user, just a technical consumer of the API.

## User Stories

1. As an API Consumer, I want to call `GET /hello?name=X` and receive a JSON
greeting addressed to `X`, so that I can verify the service responds
correctly for a given name.
2. As an API Consumer, I want `GET /hello` to still return a valid JSON
greeting when I omit the `name` parameter, so that the endpoint never
fails just because I left it out.
3. As an API Consumer, I want `GET /hello?name=` (an explicitly empty value)
to return the same default greeting as omitting `name` entirely, so that
an empty value is never treated as an error.
4. As an API Consumer, I want `GET /hello` to accept a very long `name`
value and return a greeting addressed to it in full, so that the endpoint
never fails just because the name is unusually long.
5. As an API Consumer, I want `GET /hello` to accept a `name` containing
non-ASCII/unicode characters (e.g. accented letters, CJK script, emoji) and
return a greeting addressed to it correctly, so that the endpoint works for
names outside the ASCII range.
6. As an API Consumer, I want every response from the service to carry a
`Content-Type: application/json` header, so that I can parse it
programmatically without guessing its format.
7. As an API Consumer, I want to call `GET /health` and get a simple status
response, so that I can verify the service is up before integrating against
it.
8. As an API Consumer, I want `GET /hello` to ignore any query parameters
other than `name`, so that an unrelated or misspelled parameter never
changes or breaks the response.

## Product Decisions

- When `name` is missing or empty, the service returns a default greeting
(e.g. addressed to "World") rather than an error. *assumed*
- The `/hello` endpoint is open — no sign-in, no API key — reachable by any
client. *assumed*
- The response is JSON, carrying at minimum the greeting message.
- An explicitly empty `name` value (`?name=`) is treated the same as an
omitted `name` — both fall back to the default greeting. *assumed*
- A very long `name` value is accepted and echoed in full; the service
enforces no maximum length. *assumed*
- A `name` containing unicode characters is passed through as-is (valid
UTF-8), with no normalization, transliteration, or rejection. *assumed*
- Every response, success or error, carries `Content-Type: application/json`
and a stable, documented shape. *assumed*
- Error responses use the single shape `{"error": "message"}`, per the
`app-factory-kaj/e2e-reference` conventions. *assumed*
- The service is stateless and holds no in-memory session data, so any
number of instances may run behind a load balancer with no coordination
between them. *assumed*
- The service listens on a port read from the `PORT` environment variable,
defaulting to `8080` when unset, per the `e2e-reference` conventions.
*assumed*
- `GET /health` responds quickly (no downstream calls, no I/O) so it is
cheap to poll for liveness/readiness checks. *assumed*
- Unrecognized query parameters on `GET /hello` are silently ignored rather
than rejected. *assumed*
- The service is built on the Go standard library (`net/http`) with no web
framework dependency, per the `e2e-reference` conventions. *assumed*
- No request or response data is logged or persisted anywhere, since names
may be arbitrary user input and the service keeps no records. *assumed*
- A single request is handled independently of any other; there is no
rate limit, quota, or per-caller state to maintain. *assumed*

## Out of Scope

- Persistence of any kind — the service is stateless.
- Any additional endpoints beyond `/hello`.
- Authentication, authorization, or per-caller rate limiting.
- Internationalization or localization of the greeting text.

## Open Questions

1. Is there a maximum request size or `name` length the hosting gateway
enforces upstream, independent of what this service itself accepts?
Deferred — the user will decide later if it becomes relevant.

## Further Notes

The implementation should follow the conventions demonstrated in
`app-factory-kaj/e2e-reference`, per the original project brief. The
`specs/requirements/references/ref.txt` reference note attached to this
project contained no additional product requirements beyond that pointer.