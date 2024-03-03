# Build Your Own JSON Parser

This challenge is a part of the Coding Challenges by John Crickett, which helps software engineers level up their skills through deliberate practice.

Link to the JSON Parser Challenge: [Click Here](https://codingchallenges.fyi/challenges/challenge-json-parser/)

## Table Of Contents
- [Description](#description)
- [Usage](#usage)
- [Learnings](#learnings)
- [Challenges](#challenges)
- [Contribution](#contribution)
- [About me](#about-me)

## Description

JSON (which stands for JavaScript Object Notation) is a lightweight data-interchange format, which is widely used for transmitting data over the Internet.

Everyday we are doing some development where we have a need to built JSON data and then parse it. JSON parsers are used everywhere, from your IDE to your browser.

Building a JSON parser is an easy way to learn about parsing techniques which are useful for everything from parsing simple data formats through to building a fully featured compiler for a programming language.

<b>This challenge is to build your own JSON parser.</b>

The JSON Parser does the parsing and validation of data in a format resembling JSON. It starts by breaking down the input data into individual tokens through lexical analysis. Then, it proceeds to parse and validate the data structure, ensuring that it conforms to the expected JSON-like format. It ultimately determines whether the input data is "VALID JSON" or "INVALID JSON" based on the correctness of its structure.

## Usage

Execute the following to create an executable file.

```shell
go build main.go
```

Run the following command to test the validity of JSON stored in the `filName.json`

```shell
cat <fileName.json> | ./main
```

To test all the tests, we need to run `run_tests.zsh`. You can also copy the contents of the file `run_tests.zsh` to `.sh`. You need to give executable permission, before running it. To give the permissions, run `chmod +x <filename>`. Once you have the permissions, run the following commmand.

```shell
./run_tests.zsh
```

## Learnings

There were lots of learning while building JSON Parser in GoLang. Some key learnings are:

- Building JSON Parser was a great excuse to revise the concepts I learnt in the subject <b>Compiler Design</b> back in my College.
- In order to come up with one logic and design of building my Solution, I went through many existing solutions, both in GoLang and other Langauges. This helped me to learn many things, both in a programatic way and logical way. 
- Observing the coding techniques of others helped me discover many parts of programming in GoLang.
- I learnt about how to use `Constants` in GoLang, practiced more on `Structs`, `Arrays` and `Slices`.
- One thing which I found interesting was to declare an `Array` in a way it looked like `Map`, but it wasn't.
- I had a good practice working with files and stdin inputs, and iterating through the contents for generating tokens.
- I learnt more about the formats of strings and numbers allowed in JSON format, when I got few tests wrong.

## Challenges

I came accross many challenges while building JSON Parser. Building Parser was not as easy as I thought after making the lexer. I had some good brainstorming time figuring out how to build the Parser.

The most challenging part was to fix the JSON Parser when I was hit by some test failures. Writing the correct validator for strings and numbers took some time. Then came the last failing test, related to the depth of nesting, but I found a simpler way to do it eventually.

## Contribution

butions are welcome! If you find any bugs or want to enhance the solution, feel free to open issues or submit pull requests. Please make sure to follow the coding standards and guidelines.

Happy coding! If you have any questions or need assistance, don't hesitate to reach out.

## About Me

<b>Harsh Bardolia</b>
-   [Github](https://github.com/HarshBardolia01)
-   [LinkedIn](https://www.linkedin.com/in/harsh-bardolia-0a0407217/)

## License

[MIT](/LICENSE) © Harsh Bardolia