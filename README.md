<div align="center">
  <img src="https://user-images.githubusercontent.com/11808903/112536517-e07bcb00-8dad-11eb-9931-10ad4fe5c1d9.png" width="200"/>

  <h1>shamir</h1>

  <p>Split and combine secrets using <a href="https://en.wikipedia.org/wiki/Shamir%27s_Secret_Sharing">Shamir's Secret Sharing</a> algorithm</p>

  <a href="https://github.com/slashbinslashnoname/shamir/releases/latest">
    <img src="https://img.shields.io/github/release/incipher/shamir.svg?style=for-the-badge" />
  </a>
</div>

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Description](#description)
- [Background](#background)
- [Installation](#installation)
- [Usage](#usage)
  - [Interactive](#interactive)
  - [Non-interactive](#non-interactive)
- [License](#license)

## Description

Featuring UNIX-style composability, this command-line tool facilitates splitting and combining secrets using [HashiCorp Vault's implementation](https://github.com/hashicorp/vault/blob/main/shamir/shamir.go) of [Shamir's Secret Sharing](https://en.wikipedia.org/wiki/Shamir%27s_Secret_Sharing) algorithm.

## Background

Formulated by [Adi Shamir](https://en.wikipedia.org/wiki/Adi_Shamir) (the S in [RSA](<https://en.wikipedia.org/wiki/RSA_(cryptosystem)>)) in his 1979 paper [“How to share a secret”](http://web.mit.edu/6.857/OldStuff/Fall03/ref/Shamir-HowToShareASecret.pdf), Shamir's Secret Sharing is an algorithm that allows you to split a secret (e.g. a [symmetric encryption](https://en.wikipedia.org/wiki/Symmetric-key_algorithm) key) into $n$ shares, which can be combined later to reconstruct that secret.

![Diagram](./doc/assets/diagram.png)

Not all shares need to be present for a successful reconstruction, but actually any subset thereof with a size greater than or equal to the minimum threshold $k$, where $2 \le k \le n$. The algorithm mathematically guarantees that knowledge of $k - 1$ shares reveals absolutely no information about the original secret.

## Installation

| Platform     | Package manager             | Command                              |
| ------------ | --------------------------- | ------------------------------------ |
| Linux, macOS | [Homebrew](https://brew.sh) | `$ brew install incipher/tap/shamir` |

## Usage

### Interactive

![A GIF showing how to use shamir interactively](./doc/assets/interactive-usage.gif)

### Non-interactive

```
$ echo "SayHelloToMyLittleFriend" | shamir split -n 5 -k 3 > shares.txt
Secret: ************************
```

```
$ head -n 3 shares.txt | shamir combine -k 3
SayHelloToMyLittleFriend
```

## License

<a href="https://creativecommons.org/publicdomain/zero/1.0/">CC0</a>
