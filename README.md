# Utility to optimize World of Warcraft gear given an inventory and prioritization of stats

## Prerequisites

* Go 1.22 (download from <https://go.dev/doc/install>)
* Git (download from <https://git-scm.com/downloads>)

## Use

### Get the sources

* git clone <https://github.com/mvdslot/wowgear>

### Create the executable

build.cmd

### Create inventory and stats files

You can use the ones provided for warlock and hunter as examples, just copy and modify.

### Run the executable

wowgear -inv="path to inventory file" - stats="path to stats file" [-hitcap=value] [-overrides=<stat1=value1[+stat2=value2]...>]

## Examples

* A warlock tanking the twin emps might value shadow resistance and stamina higher than the values configured in the stats file:
* wowgear -inv=warlock_inv.yaml -stats=warlock_stats.yaml -overrides=rs=2+sta=1.5
*
* A hunter in PvP might want to cap hit rating at 6:
* wowgear -inv=hunter_inv.yaml -stats=hunter_stats.yaml -hitcap=6

Sample inventory and stats files for warlock and hunter are provided, you can edit them to your own preferences, or might create different versions for raid, PvP, solo etc.
