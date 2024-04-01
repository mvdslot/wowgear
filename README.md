# Utility to optimize World of Warcraft gear given an inventory and priotization os stats

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

wowgear -inv=<path to inventory file> - stats=<path to stats file> [-hitcap=<hitcap>] [-overrides=<stat1=value1[,stat2=value2]...>]

## Examples

* A warlock tanking the twin emps might value shadow resistance higher than the value in ths stats file:
wowgear -inv=warlock_inv.json -stats=warlock_stats.json -overrides=rs=10

* A hunter in PvP might want to cap hit rating at 6:
wowgear -inv=hunter_inv.json -stats=hunter_stats.json -hitcap=6

Sample inventory and stats files for warlock and hunter are provided, you can edit them to your own preferences, or might create different versions for raid, Pvp, solo etc.
