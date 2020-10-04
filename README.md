# Questionmate

Questionmate is a library and API for building configurable online surveys. Surveys are modeled in a proprietary DSL:

```
1: Which programming languages do you use?
type: multi
options:
    1: modern (e.g. Go, Rust, Swift)
    - sustainability: 3 
    2: standard (e.g. Java, C/C++, JavaScript, PHP)
    - sustainability: 2
    4: expiring (Cobol, Delphi)
    - sustainability: 1
```

A question consists of its id (1), it's text, a type (single or multi) and a list of options (modern, ...). An option can have one or more targets (sustainability) associated with a credit value (3). A target can be followed by a list of statements that explain why a selected option impacts the associated target.

## Credits

Users can select one or more of the question's options, depending on the question type. A selected option pays into its associated targets. For instance, option 1 *modern* pays 3 credits into the target *sustainability*.

## Question order

Questions are served in a natural order defined by their ids, thus question 2 follows question 1. 

## Architecture

Questionmate's design is based on hexagonal architecture ideas. Thus, it focuses on its domain model and rules but makes no assumptions about how this model is used within a specific context, e.g. the web.

Instead of forcing programmers to use *Questionmate* within a specific context examples are given about how to use the library wrapped as REST services deployed to AWS lambda.

## API

**POST http://questionmate.ralfwirdemann.de/{questionaire}/questions**

Returns the next question according to the answers the user has given so far contained in the request body as follows:

```json
{
  "answers": [
    {
      "question_id": 1,
      "value": 1
    }
  ]
}
```

HTTP/1.1 200 Ok
```json
{
  "id": 2,
  "text": "How often do incidents occur without first changing the software?",
  "type": "single",
  "options": [
    {
      "value": 1,
      "text": "almost never",
      "targets": {
        "robustness": {
          "value": 10
        }
      }
    },
    ...
  ]
}
```

## License

* [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)