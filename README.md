# Gematria

This package is a command line utility (CLI) that is written in Go that
provides you with a powerful interactive menu interface to work with
the [https://github.com/andreimerlescu/gematria](gematria) package.

## What Is Gematria?

If you think of the alphabet as an array of characters or bytes, the
index position of the character is kind of like what Gematria is. 

If `A=1` `P=16` `L=12` `E=5`, the word `APPLE` would be `1 16 16 12 5`.
And when you express this as **Simple Gematria**, you refer to the
sum of the character's positional indexes, so `APPLE=50`.

The phrase `AMERICA=50` as well. Given`A=1` `M=13` `E=5` `R=18` `I=9` 
`C=3` `A=1`, you can see `1 13 5 18 9 3 1 = 50`, therefore in **Simple
Gematria**, `APPLE` == `AMERICA` through this numerological system.

# Installation

You can begin using the GEMATRIA utility on your system by installing this Go package. 

```bash
go install github.com/andreimerlescu/cli-gematria@latest
```

Then you can begin using the program by running: 

```bash
[~]$ cli-gematria 
Welcome to Gematria Calculator! 🎉
Type '?' for menu, 'exit' or 'quit' to exit.
```

## Example Usage

```bash
Welcome to Gematria Calculator! 🎉
Type '?' for menu, 'exit' or 'quit' to exit.
> i love yahuah
🔢 Gematria result for 'i love yahuah'
        Jewish = 1402   English = 762   Simple = 127    Mystery = 3676  Majestic = 381  Eights = 889    
> america
🔢 Gematria result for 'america'
        Jewish = 129    English = 300   Simple = 50     Mystery = 2834  Majestic = 150  Eights = 325    
> apple
🔢 Gematria result for 'apple'
        Jewish = 146    English = 300   Simple = 50     Mystery = 1080  Majestic = 148  Eights = 339    
> andrei
🔢 Gematria result for 'andrei'
        Jewish = 139    English = 306   Simple = 51     Mystery = 2264  Majestic = 153  Eights = 338    
> michael
🔢 Gematria result for 'michael'
        Jewish = 76     English = 306   Simple = 51     Mystery = 2648  Majestic = 153  Eights = 328    
> i am saint andrei
🔢 Gematria result for 'i am saint andrei'
        Jewish = 419    English = 822   Simple = 137    Mystery = 7209  Majestic = 411  Eights = 934    
> i am saint michael
🔢 Gematria result for 'i am saint michael'
        Jewish = 356    English = 822   Simple = 137    Mystery = 7593  Majestic = 411  Eights = 924    
> ?   
Available options:
+--------+---------------+
| OPTION |  DESCRIPTION  |
+--------+---------------+
| 1      | ❓ Help       |
| 2      | 🚪 Close Menu |
| 3      | 🚪 Exit App   |
+--------+---------------+
Select an option (1-3): 1
Help information:
- Type any text to calculate its gematria.
- Type '?' to show this menu.
- Type 'exit' or 'quit' to exit.
> ?   
Available options:
+--------+---------------+
| OPTION |  DESCRIPTION  |
+--------+---------------+
| 1      | ❓ Help       |
| 2      | 🚪 Close Menu |
| 3      | 🚪 Exit App   |
+--------+---------------+
Select an option (1-3): 2
> ?
Available options:
+--------+---------------+
| OPTION |  DESCRIPTION  |
+--------+---------------+
| 1      | ❓ Help       |
| 2      | 🚪 Close Menu |
| 3      | 🚪 Exit App   |
+--------+---------------+
Select an option (1-3): 3
Exiting... 👋


```
