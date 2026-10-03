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

## Product Decisions

- When `name` is missing or empty, the service returns a default greeting
(e.g. addressed to "World") rather than an error. *assumed*
- The `/hello` endpoint is open — no sign-in, no API key — reachable by any
client. *assumed*
- The response is JSON, carrying at minimum the greeting message.

## Out of Scope

- Persistence of any kind — the service is stateless.
- Any additional endpoints beyond `/hello`.
- Authentication, authorization, or per-caller rate limiting.
- Internationalization or localization of the greeting text.

## Open Questions

None at this time.

## Further Notes

The implementation should follow the conventions demonstrated in
`app-factory-kaj/e2e-reference`, per the original project brief. The
`specs/requirements/references/ref.txt` reference note attached to this
project contained no additional product requirements beyond that pointer.