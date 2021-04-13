# Mockgen
A CLI utility to generate arbitrarily large realistic mock data sets

## Background
After experimenting with other tools in the space, I frequently ran into issues with generating more than 1M rows of test data, so I made my own.

Currently, the tool generates postgres flavored SQL, but maybe in the future that can be expanded to multiple flavores or even formats.

## Installation

```bash
go install github.com/ImVexed/mockgen@latest
```

TODO: Make an actions stage to build a static binary

## Usage

```bash
Usage: mockgen sql <table> <count> <columns> ...

Generate SQL mock data

Arguments:
  <table>          Table to insert into
  <count>          Ammount of rows to generate
  <columns> ...    Columns to create (name={firstname} id={uuid} bio={sentence:30})

Flags:
  -h, --help    Show context-sensitive help.

  -b, --bulk    Bulk insert
```

```bash
mockgen sql -b message 1000000 room_id=1 user_id=2 "author_name={firstname} {lastname}" content={sentence:10} > messages.sql
```
generates:
```sql
INSERT INTO "message" ("room_id","user_id","author_name","content") VALUES
(1,2,'Whitney Predovic','Design map wonder bread agree celebrate motion fashion examine pull.')
(1,2,'Quentin Maggio','Recover cover child sport imagine justify of sort purchase stop.')
(1,2,'Esta Romaguera','Case prepare charge restore challenge unemployment future north involve teach.')
(1,2,'Cathryn Vandervort','Achieve marriage violence seem persuade solution independence curriculum promote limit.')
...
```


## Fake Data Types
Under the hood we use [GoFakeIt](https://github.com/brianvoe/gofakeit), see their readme for more info

- address
    - number
    - street_suffix
    - city
    - state_abr
    - country
    - country_abr
    - street_prefix
    - street_name
    - state
    - zip
- animal
    - petname
    - animal
    - type
    - farm
    - cat
    - dog
- beer
    - yeast
    - malt
    - style
    - name
    - hop
- car
    - maker
    - model
    - type
    - fuel_type
    - transmission_type
- color
    - safe
    - full
- company
    - name
    - suffix
    - buzzwords
    - bs
- computer
    - linux_processor
    - mac_processor
    - windows_platform
- currency
    - short
    - long
- emoji
    - tag
    - emoji
    - description
    - category
    - alias
- file
    - mime_type
    - extension
- food
    - dinner
    - snack
    - dessert
    - fruit
    - vegetable
    - breakfast
    - lunch
- hacker
    - ingverb
    - phrase
    - abbreviation
    - adjective
    - noun
    - verb
- hipster
    - word
- internet
    - browser
    - domain_suffix
    - http_method
- job
    - descriptor
    - level
    - title
- language
    - short
    - long
    - programming
- log_level
    - general
    - syslog
    - apache
- lorem
    - word
- person
    - first
    - last
    - phone
    - prefix
    - suffix
- timezone
    - offset
    - abr
    - text
    - full
    - region
- word
    - adverb
    - preposition
    - adjective
    - phrase
    - noun
    - verb
