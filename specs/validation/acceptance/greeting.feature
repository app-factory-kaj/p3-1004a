Feature: Greeting

  @story-1
  Rule: Calling /hello with a name returns a greeting addressed to that name

    Scenario: Greeting a named caller
      Given the greeter service is running
      When an API Consumer calls GET /hello with name "Ada"
      Then the response is a JSON greeting addressed to "Ada"

  @story-2
  Rule: Calling /hello without a name still returns a valid greeting

    Scenario: Greeting with no name supplied
      Given the greeter service is running
      When an API Consumer calls GET /hello with no name parameter
      Then the response is a JSON greeting addressed to "World"

    Scenario: Greeting with an empty name
      Given the greeter service is running
      When an API Consumer calls GET /hello with name ""
      Then the response is a JSON greeting addressed to "World"
