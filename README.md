# autocorrect-with-go 
Welcome!

Autocorrect-with-go is a simple auto-correction tool that can manipulate text in acoordance to common grammar rules. Autocorrect-with-go 's features include (but not limited to) capitalization, transforming text between uppercase and lowercase, and transforming hexadecimal or binary repesentations to decimal digits. That said, there are a few rules to follow to reap the benefits of these features.

## Features
<b> Replace Hexadecimal Numbers: </b> Replace every instance of (hex) with the decimal equivalent of the preceding hexadecimal number.

* *For Example:*
<pre>
input: "1E (hex) files were added"
output: "30 files were added"
</pre>

**Replace Binary Numbers:** Replace every instance of (bin) with the decimal equivalent of the preceding binary number.

* *For Example:*
<pre>
input: "It has been 10 (bin) years"
output: "It has been 2 years"
</pre>

**Uppercase Conversion:** Convert the preceding word to uppercase for every instance of (up).

* *For Example:*
<pre>
input: "Ready, set, go (up) !"
output: "Ready, set, GO !"
</pre>

**Lowercase Conversion:** Convert the preceding word to lowercase for every instance of (low).

* *For Example:*
<pre>
input: "I should stop SHOUTING (low)"
output: "I should stop shouting"
</pre>

**Capitalization:** Capitalize the preceding word for every instance of (cap). Optionally, capitalize a specified number of words after the (cap) tag.

* *For Example:*
<pre>
input: "Welcome to the Brooklyn bridge (cap)"
output: "Welcome to the Brooklyn Bridge"
</pre>

**Punctuation Adjustment:** Ensure proper spacing around punctuation marks such as ., ,, !, ?, :, and ;. Handle special cases like multiple punctuation marks and ellipses.

* *For Example:*
<pre>
input: "I was sitting over there ,and then BAMM !!"
output: "I was sitting over there, and then BAMM!!"
</pre>

**Single Quotes Handling:** Place single quotation marks properly around words, ensuring no spaces between the quotes and the words.
* *For Example:*
<pre>
input: "I am exactly how they describe me: ' awesome '"
output: "I am exactly how they describe me: 'awesome'"
</pre>

**Conditional 'A' Conversion:** Replace instances of 'a' with 'an' if the next word begins with a vowel or 'h'.
* *For Example:*
<pre>
input: "There it was. A amazing rock!"
output: "There it was. An amazing rock!"
</pre>

## Usage

To use this tool, run the following command:

* ***Terminal***

<pre>
$ go run . input_file.txt output_file.txt
</pre>

Replace input_file.txt with the name of the input file containing the text to be modified, and output_file.txt with the name of the output file where the modified text will be saved.
* ***Examples***

Here are some examples of how to use the tool:

* ***Terminal***
<pre>
$ go run . sample.txt result.txt
</pre>

## Dependencies

This tool requires no external dependencies. It is a standalone executable written in Go.

## Author

David Jesse Odhiambo

Aprentice Software Developer, Zone01 Kisumu.
