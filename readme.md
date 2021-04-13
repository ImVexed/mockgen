# Mockgen
A CLI utility to generate arbitrarily large realistic mock data sets

## Background
After experimenting with other tools in the space, I frequently ran into issues with generating more than 1M rows of test data, so I made my own.

Currently, the tool generates postgres flavored SQL, but maybe in the future that can be expanded to multiple SQL flavors or even different formats (xml,json, etc).

## Installation

```bash
go install github.com/ImVexed/mockgen@latest
```

or download directly from [releases](https://github.com/ImVexed/mockgen/releases)
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

- xml
	-  type, string, Type of XML, single or array
	-  rootelement, string, Root element wrapper name
	-  recordelement, string, Record element for each record row
	-  rowcount, int, Number of rows in JSON array
	-  fields, []Field, Fields containing key name and function to run in json format
	-  indent, bool, Whether or not to add indents and newlines
- int8
- shuffleints
	-  ints, []int, Delimited separated integers
- verb
- loglevel
- timezoneregion
- hipsterword
- hipsterparagraph
	-  paragraphcount, int, Number of paragraphs
	-  sentencecount, int, Number of sentences in a paragraph
	-  wordcount, int, Number of words in a sentence
	-  paragraphseparator, string, String value to add between paragraphs
- weighted
	-  options, []string, Array of any values
	-  weights, []float, Array of weights
- quote
- safariuseragent
- httpmethod
- jobtitle
- emojitag
- int32
- streetnumber
- question
- languageabbreviation
- int64
- generate
	-  str, string, String value to generate from
- ipv6address
- day
- timezonefull
- creditcardcvv
- bs
- csv
	-  rowcount, int, Number of rows in JSON array
	-  fields, []Field, Fields containing key name and function to run in json format
	-  delimiter, string, Separator in between row values
- latituderange
	-  min, float, Minimum range
	-  max, float, Maximum range
- cartype
- cat
- dessert
- zip
- bitcoinaddress
- domainname
- achrouting
- stateabr
- phrase
- price
	-  min, float, Minimum price value
	-  max, float, Maximum price value
- programminglanguage
- imageurl
	-  width, int, Image width in px
	-  height, int, Image height in px
- animaltype
- farmanimal
- nameprefix
- bool
- paragraph
	-  paragraphcount, int, Number of paragraphs
	-  sentencecount, int, Number of sentences in a paragraph
	-  wordcount, int, Number of words in a sentence
	-  paragraphseparator, string, String value to add between paragraphs
- monthstring
- creditcardnumber
	-  types, []string, A select number of types you want to use when generating a credit card number
	-  bins, []string, Optional list of prepended bin numbers to pick from
	-  gaps, bool, Whether or not to have gaps in number
- appauthor
- countryabr
- cartransmissiontype
- dog
- int16
- float64
- regex
	-  str, string, Regex RE2 syntax string
- minute
- weekday
- currencylong
- appname
- address
- beeralcohol
- nanosecond
- beerhop
- adverb
- url
- second
- hour
- appversion
- password
	-  lower, bool, Whether or not to add lower case characters
	-  upper, bool, Whether or not to add upper case characters
	-  numeric, bool, Whether or not to add numeric characters
	-  special, bool, Whether or not to add special characters
	-  space, bool, Whether or not to add spaces
	-  length, int, Number of characters in password
- carfueltype
- domainsuffix
- creditcardexp
- hackerverb
- hackeringverb
- filemimetype
- emojidescription
- carmodel
- loremipsumparagraph
	-  paragraphcount, int, Number of paragraphs
	-  sentencecount, int, Number of sentences in a paragraph
	-  wordcount, int, Number of words in a sentence
	-  paragraphseparator, string, String value to add between paragraphs
- float32range
	-  min, float, Minimum float32 value
	-  max, float, Maximum float32 value
- color
- creditcard
- company
- uint64
- lettern
	-  count, uint, Number of digits to generate
- dinner
- carmaker
- person
- adjective
- uuid
- month
- year
- uint32
- street
- phoneformatted
- timezoneoffset
- currency
- randomstring
	-  strs, []string, Delimited separated strings
- vegetable
- breakfast
- loremipsumword
- httpstatuscodesimple
- loremipsumsentence
	-  wordcount, int, Number of words in a sentence
- ipv4address
- hackerabbreviation
- hackernoun
- lexify
	-  str, string, String value to replace #'s
- beeribu
- email
- uint8
- letter
- sentence
	-  wordcount, int, Number of words in a sentence
- emojialias
- achaccount
- imagejpeg
	-  width, int, Image width in px
	-  height, int, Image height in px
- imagepng
	-  width, int, Image width in px
	-  height, int, Image height in px
- number
	-  min, int, Minimum integer value
	-  max, int, Maximum integer value
- digitn
	-  count, uint, Number of digits to generate
- petname
- namesuffix
- flipacoin
- noun
- float32
- username
- streetsuffix
- companysuffix
- json
	-  type, string, Type of JSON, object or array
	-  rowcount, int, Number of rows in JSON array
	-  fields, []Field, Fields containing key name and function to run in json format
	-  indent, bool, Whether or not to add indents and newlines
- emojicategory
- animal
- snack
- beerblg
- safecolor
- firstname
- language
- latitude
- beername
- lastname
- gender
- httpstatuscode
- joblevel
- hackeradjective
- fileextension
- streetprefix
- beermalt
- date
	-  format, string, Date time string format output
- gamertag
- fruit
- name
- rgbcolor
- beerstyle
- creditcardtype
- buzzword
- emoji
- longitude
- longituderange
	-  min, float, Minimum range
	-  max, float, Maximum range
- teams
	-  people, []string, Array of people
	-  teams, []string, Array of teams
- operauseragent
- hipstersentence
	-  wordcount, int, Number of words in a sentence
- float64range
	-  min, float, Minimum float64 value
	-  max, float, Maximum float64 value
- city
- car
- ssn
- timezone
- digit
- country
- streetname
- preposition
- word
- hexcolor
- useragent
- chromeuseragent
- bitcoinprivatekey
- state
- beeryeast
- uint16
- job
- jobdescriptor
- timezoneabv
- currencyshort
- hackerphrase
- numerify
	-  str, string, String value to replace #'s
- shufflestrings
	-  strs, []string, Delimited separated strings
- lunch
- phone
- firefoxuseragent

