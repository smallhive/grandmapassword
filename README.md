# grandmapassword #

## Requirements

- go 1.19+

## Solution

The solution has 3 steps

- Loading dictionary
- Generation pairs of words
- Generating password

### Dictionary

The application uses words dictionary to find the best password. Right now it supports only file dictionary (example in
words.txt).
New types may be added in easy way, you just need to implement `word.Loader` interface

### Pairs

Generating pairs of word allows us to control of resulting password length in required range in the next step

### Generation

Iterating sums of word pairs length, we are taking pairs which has total symbol length in range 20-24. These pairs we
combine
in one word and calculate total difficulty of password variant. After that we sort all variants by difficulty and take
the simplest one

### Run

Application have possibility to take dictionary as a parameter:

> By default App uses `words_5k.txt`

```shell
$ ./build/bin/grandmapassword words_10k.txt

$ ./build/bin/grandmapassword words_100k.txt
```

## Output description

Solution result is the string with next format:
```shell
dictionary 4798
generatePairs 173.34465ms
generateVariants 1.884765ms
Password wwwweeddeferreddress l=20,d=15
```

You should read this like:

- `dictionary: <N>` - actual size of words dictionary. All short letters already excluded
- `generatePairs` - spent time for generating pairs
- `generateVariants` - spent time for generating passwords from generated pairs
- `l (length)=<N>` - the password length
- `d (difficulty)=<N>` - difficulty means amount of finger moving from first password letter to last. It also includes
  moves between words

### Measurements

```shell
$ ./build/bin/grandmapassword 
dictionary 4798
generatePairs 173.34465ms
generateVariants 1.884765ms
Password wwwweeddeferreddress l=20,d=15 

$ ./build/bin/grandmapassword words_10k.txt 
dictionary 9578
generatePairs 691.574905ms
generateVariants 2.434074ms
Password weeddeerreeddeferred l=20,d=14

$ ./build/bin/grandmapassword words_100k.txt 
dictionary 92422
generatePairs 47.30321564s
generateVariants 22.46303ms
Password wedderreedeessweewee l=20,d=13
```

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
