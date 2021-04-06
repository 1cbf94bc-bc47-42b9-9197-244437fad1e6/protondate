# Protondate.
![Twitter Follow](https://img.shields.io/twitter/follow/chaosd0c?style=social)

Protondate is a Go based CLI application that returns the potential signup date for a given protonmail email address.

## How does it work?

Protonmail creates PGP keys for users that sign up to their platform. This allows us to query their [keyserver](https://api.protonmail.ch/pks/lookup?op=get&search=example@example.com), download the PGP key, parse it and extract the creation date/time to potentially find the date they signed up for their account. Or renewed their PGP key.

## Usage

```
Protondate 0.0.1 by Doctor Chaos

Usage: protondate <email address>
  -filename string
    	The filename the public key should be saved as (default "public_key.pgp")
  -save
    	Save the public key to disk
```

## Example

```
./protondate example@example.com

The protonmail account belonging to "example@example.com" was potentially created on 2021-01-01 00:00:00 +0100 CET
```

## Installation

```
git clone https://github.com/1cbf94bc-bc47-42b9-9197-244437fad1e6/protondate
cd protondate
go build
./protondate
```
