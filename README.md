# 💊 Pillbox

## What is this?

This is a CLI app written in Golang, that can be used to set medications reminders.
I hope you *never* use it, but if you are sick and you are taking meds, this
might help you remember when to take your meds. At least it helped me.

### ❗️ Important ❗️

Currently this works only on OSX, since it heavily relies on `launchd` for
management of the daemon. When I get the OSX support clean and shiny I will take
a look at GNU/Linux and Windows.

## Installation

```
go get github.com/fteem/pbox
```

That should do it.

## Usage

## Installation

To install the `pillboxd` daemon, you need to run:

```
pillbox install
```

This will create a new `launchd` agent for your user, which will manage the state
of the daemon.
