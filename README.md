# EgoKey
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FDiti%2Fegokey.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FDiti%2Fegokey?ref=badge_shield)


An OpenPGP key generation program which aims at generating “pretty keys”.

Given the highly-secure nature of [OpenPGP](https://tools.ietf.org/html/rfc4880),
the only way to get such a result is by [brute force](https://en.wikipedia.org/wiki/Brute-force_attack).
EgoKey will continuously generate keypairs and discard the uninteresting ones.

## Important disclaimer

EgoKey is probably NOT secure enough to be used in sensitive settings, such as if
you are a Debian maintainer use it to sign your releases. Only use it if what you
have to protect doesn’t matter too much.

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FDiti%2Fegokey.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FDiti%2Fegokey?ref=badge_large)