# DnD Text Enhancer
This is an application intended to assist Dungeon Masters create descriptive content that uses a variety of vocabulary.  The application includes a set of examples that fit this theme.  The enhancer swaps identified nouns, adjectives, and adverbs for a random synonym from Merriam-Webster's Collegiate Thesaurus.

I chose to create this project becuase, while I like to improvise new settings and descriptions, my vocabulary can get repetitive.  This project helps to automate that process of using more varied language by using Natural Language Processing.  To combat the sometimes awkward output, I have included functionality to edit the resulting text - enabling me to pick and choose the right word within the context of the sentence.  

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
        1. If your operating system / architecture does not have a corresponding build file, please see the section about running directly from the source code.
    1. Run that binary file:
        1. Use `chmod +x <filename>` to make that file an executable (if not already)
        1. Execute the file with `./<binary-filename> <api-key>


### Running directly from the source code
If the necessary binary for your Operating System and Architecture is not present in the builds directory, you can run the program directly on your machine.  Ensure you have [go v1.17 installed](https://go.dev/doc/install).

    1. Open a command prompt or terminal 
    1. Navigate to the directory where this README is located
    1. Run `go run main.go <api-key>`


### Program Flow
The program will then prompt you for text to enhance.  Enter the desired text and use 'Enter' or a newline to submit.  

Optionally, enter words that you would like to be retained in the enhanced text. These words will not be replaced by synonyms.

The input text will then be processed by a Natural Language Processing library to identify nouns, adjectives, and adverbs.  These words will be replaced by a random synonym from Merriam-Webster's Collegiate Thesaurus.

After the synonym replacement, the enhanced text will be output.

Then users will have the option to edit any of the words that were replaced.  The user can choose a different synonym from the thesaurus or the original word. 

While in edit mode, the user will see the enhanced text annotated with brackets of indicies (ex. "Nice[0] day[1]").  The user can input an index of the word they would like to edit or the cancel index to exit edit mode.  On selecting a valid index, the synonyms and original word will be listed and the user can again input an index to select a new replacement.

## What I'd Do Differently in Production

I would opt to use a method of collecting user input that has a better user experience.  Ideally, something like [promptui](https://github.com/manifoldco/promptui) that is interactive.  The reason I chose not to use promptui was for compatibility issues with my Windows Command Prompt - promptui was duplicating lines of output that muddled the user experience.

Additionally, I would implement automated integration testing to ensure that the user interface was behaving as expected.  These tests would cover both the common/userInput.go file and main.go as well as the Http Client for the Thesaurus interface in thesaurus.go.

If I were to deploy this program to a server, then I would utilize a secret sharing mechanism in order to automatically load the Thesaurus API key rather than passing it as a command line argument.  Loading the secret within the program would help to obfuscate that key rather than leaving it more exposed as an argument.

Finally, in production - and with more time - I would want to increase the feature set.  Integrating a Dungeons and Dragons API (like http://www.dnd5eapi.co/) that allows me to get Spell effects and Item Stats would enable me to add an Item-Generation feature that could randomly generate creative item names that match the stat effects.  For example: a "Potion of Melting" could be a potion that deals 1d4 Acid damage and a "Sword of Brawniness" could be a sword that give a character +1 to their Strength stat.

## Limitations

Please note that, under free useage, the Merriam-Webster's Collegiate Thesaurus API only permits 1000 requests per day per key.

The natural language processor has limitiations.  For example, the word 'tall' is sometimes interpreted as a "Determiner" like in "tall tree" and will not be considered an adjective.  Determiners are not replaced in order to conserve the number of requests to the Thesaurus API (most determiners (ex. 'a') do not have proper synonyms).

## Examples

Below you can find some example text to try out yourself:

```
The tree looms over the vacant field.  Dottering peasants comb the field for artifacts long forgotten.

A tavern with a blue door can be seen through the mist, lit by a single lantern hanging over the entrance.  Through a small window, you can see the movement of shadows within.

The tomb is dark and made up of a large room with stone walls.  In the center of the room, there is a large dais presenting a sarcophagus - as if inviting you to come closer.
```

## Testing

Tests can be run from the directory this README is located in by running:

`go test ./... -cover`