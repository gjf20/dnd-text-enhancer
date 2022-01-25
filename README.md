# DnD Text Enhancer
This is an application intended to assist Dungeon Masters create descriptive content that uses a variety of vocabulary.  The application includes a set of examples that fit this theme.  The enhancer swaps identified nouns, adjectives, and adverbs for a random synonym from Merriam-Webster's Collegiate Thesaurus.

## Usage 
### Running from a pre-built binary
Ensure you have an API key from https://dictionaryapi.com/products/api-collegiate-thesaurus.  

    1. Determine your computer's Operating System and Architecture
        1. Operating system
            1. Unix-like: `uname -r` in Terminal
            1. Windows: simply "Windows"
        1. Architecture
            1. Unix-like systems: `uname -m` in Terminal
            1. Windows: `set processor` in Command Prompt
    1. Check the ./builds directory for the appropriate binary file (enhancer-<OS>-<Arch>)
    1. Run that binary file with `./<binary-filename> <api-key>


### Running directly from the source code
If the necessary binary for your Operating System and Architecture is not present in the builds directory, you can run the program directly on your machine.  Ensure you have [go v1.17 installed](https://go.dev/doc/install).

    1. Open a command prompt or terminal 
    1. Navigate to the directory where this README is located
    1. Run `go run main.go <api-key>`

The program will then prompt you for text to enhance.  Enter the desired text and use 'Enter' or a newline to submit.  

Optionally, enter words that you would like to be retained in the enhanced text. These words will not be replaced by synonyms.

The input text will then be processed by a Natural Language Processing library to identify nouns, adjectives, and adverbs.  These words will be replaced by a random synonym from Merriam-Webster's Collegiate Thesaurus.

After the synonym replacement, the enhanced text will be output.

Then users will have the option to edit any of the words that were replaced.  The user can choose a different synonym from the thesaurus or the original word. 

While in edit mode, the user will see the enhanced text annotated with brackets of indicies (ex. "Nice[0] day[1]").  The user can input an index of the word they would like to edit or 'done' to exit edit mode.  On selecting a valid index, the synonyms and original word will be listed and the user can again input an index to select a new replacement.

## Limitations

Please note that, under free useage, the Merriam-Webster's Collegiate Thesaurus API only permits 1000 requests per day per key.

The natural language processor has limitiations.  For example, the word 'tall' is sometimes interpreted as a "Determiner" like in "tall tree" and will not be considered an adjective.  Determiners are not replaced in order to conserve the number of requests to the Thesaurus API (most determiners (ex. 'a') do not have proper synonyms).

## Examples

Below you can find some example text to try out yourself:


## Testing

Tests can be run from the directory this README is located in by running:

`go test ./... -cover`