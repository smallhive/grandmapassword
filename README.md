# grandmapassword #

## Requirements

- go 1.19+

## Solution

The solution has 3 steps

- Loading dictionary
- Generation pairs of words
- Generating password

### Dictionary

The application uses words dictionary to find the best password. Right now it supports only file dictionary (example in words.txt).
New types may be added in easy way, you just need to implement `word.Loader` interface

### Pairs

Generating pairs of word allows us to control of resulting password length in required range in the next step

### Generation

Iterating sums of word pairs length, we are taking pairs which has total symbol length in range 20-24. These pairs we combine
in one word and calculate total difficulty of password variant. After that we sort all variants by difficulty and take
the simplest one

## Output description

Solution result is the string with next format: `length:23  difficulty:39  phrasesreservessasseeds`

You should read this like:
- `length:<N>` - the password length
- `difficulty:<N>` - difficulty means amount of finger moving from first password letter to last. It also includes 
moves between words

## Build

```shell
make build
```

## Run

```shell
make run
```

## Test

```shell
make test
```

## Linter

```shell
make lint
```
