# loquacious

Command line application that returns the number of tweets/retweets made daily by users

## Setup

### Building the application

To build the application, move into its main directory and run

```bash
$ go build ./cmd/loquacious.go
```

This will create an executable into the main directory, otherwise you can just use

```bash
$ go run ./cmd/loquacious.go
```

if you don't want to build the application.

### Adding the client and secret

In order to use this application, you'll need to apply for and create a new Twitter Application.
More information can be found here about this part https://developer.twitter.com/en/apply-for-access.

Once the application is created, you should have a client id and a client secret, which you will need to add to the application.
To do that run the command:

```bash
$ loquacious -client-id aaaaaaaaaaidaaaaaaaaaaaaa -client-secret bbbbbbbbbbbsecretbbbbbbbbb
```

This needs to be done only once, since the client id and client secret will be saved in your home directory under ```/.loquacious/lauth.yaml```.
The file can be seen and accessed by other people on your computer, so since it holds information that can be used to access data from your Twitter account you should be careful and not share it.

## Usage

### Help

To get more information about how to use the application use ```-h``` or ```-help```
```bash
$ loquacious -help
```

### Adding new users

You can add specific users by their Twitter handle, whose daily tweet count you want to see.
Only users whose account is not protected/private are available for adding.
```bash
$ loquacious -add FooBar
```

### Getting the daily tweet counts

Getting the daily tweet counts for the added users can be achieved simply by running the application

```bashg
$ loquacious
```
and by default the tweet counts for each user will be displayed on console.

```bash
$ loquacious
Today is 03/12/2020 (Thu)
FooBar: 0 tweets
FooBar1: 4 tweets
FooBar2: 5 tweets
```

### Removing users

To remove a user run

```bash 
$ loquacious -remove FooBar
```
