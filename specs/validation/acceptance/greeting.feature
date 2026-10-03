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

  @story-3
  Rule: An explicitly empty name is treated the same as an omitted name

    Scenario: Greeting with an explicitly empty name
      Given the greeter service is running
      When an API Consumer calls GET /hello with name ""
      Then the response is a JSON greeting addressed to "World"

  @story-4
  Rule: A very long name is accepted and echoed in full

    Scenario: Greeting a caller with a very long name
      Given the greeter service is running
      When an API Consumer calls GET /hello with a name that is 2000 characters long, all the letter "a"
      Then the response is a JSON greeting addressed to that same 2000-character name

  @story-5
  Rule: A unicode name is accepted and echoed correctly

    Scenario: Greeting a caller with an accented name
      Given the greeter service is running
      When an API Consumer calls GET /hello with name "José"
      Then the response is a JSON greeting addressed to "José"

    Scenario: Greeting a caller with a CJK name
      Given the greeter service is running
      When an API Consumer calls GET /hello with name "田中"
      Then the response is a JSON greeting addressed to "田中"

    Scenario: Greeting a caller with an emoji in their name
      Given the greeter service is running
      When an API Consumer calls GET /hello with name "Ada🚀"
      Then the response is a JSON greeting addressed to "Ada🚀"
