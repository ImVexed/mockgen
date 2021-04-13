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

```
sentence  | wordcount - Number of words in a sentence
rgbcolor 
petname 
country 
generate  | str - String value to generate from
domainname 
xml  | type - Type of XML, single or array | rootelement - Root element wrapper name | recordelement - Record element for each record row | rowcount - Number of rows in JSON array | fields - Fields containing key name and function to run in json format | indent - Whether or not to add indents and newlines
int8 
int32 
state 
cartype 
day 
creditcardnumber  | types - A select number of types you want to use when generating a credit card number | bins - Optional list of prepended bin numbers to pick from | gaps - Whether or not to have gaps in number
bitcoinaddress 
shufflestrings  | strs - Delimited separated strings
hackerabbreviation 
programminglanguage 
json  | type - Type of JSON, object or array | rowcount - Number of rows in JSON array | fields - Fields containing key name and function to run in json format | indent - Whether or not to add indents and newlines
shuffleints  | ints - Delimited separated integers
car 
carfueltype 
carmodel 
firstname 
hexcolor 
httpmethod 
animal 
snack 
username 
achrouting 
emojidescription 
digitn  | count - Number of digits to generate
animaltype 
streetname 
streetsuffix 
longitude 
emojialias 
vegetable 
streetnumber 
beername 
firefoxuseragent 
hipstersentence  | wordcount - Number of words in a sentence
uint32 
zip 
beerhop 
domainsuffix 
buzzword 
languageabbreviation 
appname 
latituderange  | min - Minimum range | max - Maximum range
paragraph  | paragraphcount - Number of paragraphs | sentencecount - Number of sentences in a paragraph | wordcount - Number of words in a sentence | paragraphseparator - String value to add between paragraphs
operauseragent 
jobtitle 
hipsterword 
beermalt 
beerblg 
number  | min - Minimum integer value | max - Maximum integer value
uint64 
lexify  | str - String value to replace #'s
address 
name 
phone 
adjective 
loremipsumparagraph  | paragraphcount - Number of paragraphs | sentencecount - Number of sentences in a paragraph | wordcount - Number of words in a sentence | paragraphseparator - String value to add between paragraphs
uuid 
chromeuseragent 
timezone 
emojicategory 
emojitag 
lunch 
beerstyle 
hackeradjective 
imagepng  | width - Image width in px | height - Image height in px
beeryeast 
beeralcohol 
word 
flipacoin 
color 
safariuseragent 
fileextension 
question 
ipv6address 
company 
hipsterparagraph  | paragraphcount - Number of paragraphs | sentencecount - Number of sentences in a paragraph | wordcount - Number of words in a sentence | paragraphseparator - String value to add between paragraphs
float32 
float64range  | min - Minimum float64 value | max - Maximum float64 value
digit 
quote 
regex  | str - Regex RE2 syntax string
creditcardexp 
dog 
adverb 
nanosecond 
float64 
farmanimal 
breakfast 
lastname 
loremipsumsentence  | wordcount - Number of words in a sentence
achaccount 
hackerverb 
int64 
numerify  | str - String value to replace #'s
street 
cartransmissiontype 
teams  | people - Array of people | teams - Array of teams
noun 
bool 
letter 
ssn 
job 
csv  | rowcount - Number of rows in JSON array | fields - Fields containing key name and function to run in json format | delimiter - Separator in between row values
float32range  | min - Minimum float32 value | max - Maximum float32 value
weighted  | options - Array of any values | weights - Array of weights
longituderange  | min - Minimum range | max - Maximum range
carmaker 
nameprefix 
email 
second 
minute 
month 
timezoneabv 
dinner 
person 
verb 
preposition 
safecolor 
date  | format - Date time string format output
currencyshort 
currencylong 
creditcardcvv 
hackerphrase 
hackernoun 
imagejpeg  | width - Image width in px | height - Image height in px
httpstatuscode 
httpstatuscodesimple 
currency 
language 
randomstring  | strs - Delimited separated strings
cat 
gender 
phoneformatted 
creditcardtype 
hackeringverb 
lettern  | count - Number of digits to generate
countryabr 
streetprefix 
uint16 
appversion 
companysuffix 
bs 
int16 
appauthor 
beeribu 
monthstring 
bitcoinprivatekey 
emoji 
dessert 
city 
stateabr 
namesuffix 
url 
timezoneoffset 
creditcard 
joblevel 
fruit 
latitude 
weekday 
jobdescriptor 
filemimetype 
gamertag 
password  | lower - Whether or not to add lower case characters | upper - Whether or not to add upper case characters | numeric - Whether or not to add numeric characters | special - Whether or not to add special characters | space - Whether or not to add spaces | length - Number of characters in password
loglevel 
year 
imageurl  | width - Image width in px | height - Image height in px
phrase 
loremipsumword 
ipv4address 
useragent 
hour 
timezonefull 
timezoneregion 
price  | min - Minimum price value | max - Maximum price value
uint8 
```
