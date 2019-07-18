# whichday

Determine the day of the week for a date from November, 1st 1582 to December, 31st 9999 

Based on method #2 from [wikibooks](https://fr.wikibooks.org/wiki/Curiosit%C3%A9s_math%C3%A9matiques/Trouver_le_jour_de_la_semaine_avec_une_date_donn%C3%A9e)

## Get Involved

* Fork from https://github.com/marcpascal/whichday
* Local git clone in your github environment\

## Current state of whichday

Command line support only with user interaction.

## Future of whichday

* Operations on date like delta days between 2 dates.
* GUI for better experience
* Local store of data to quickly populate the GUI
* Determine the moon phase
* Addition of a biorythm

## Building whichday

``` bash
make build
```

and add to your path

``` bash
mv whichday /usr/local/bin
```
## Execution

```bash
# Open a terminal
$ whichday --help
```
