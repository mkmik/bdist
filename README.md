# bdist

[![Go Report Card](https://goreportcard.com/badge/github.com/mkmik/bdist)](https://goreportcard.com/report/github.com/mkmik/bdist)

```
$ bdist -d /www/patches foo-1.0 foo-2.0
$ ls /www/patches
f52e22a6ef467d206e884794e39880956990228f735836c8375ea3d4ea38c074-to-0eb56015fdd2281fdbe3e804bbfd82739695e335d63dabecdff09ced11851aae.bpatch
```

Then, once somebody has `foo-1.0` and wants to upgrade to `foo-2.0`, all they have to do is to fetch the sha256 signature for foo-2.0, e.g.:

```
$ sha256 foo
f52e22a6ef467d206e884794e39880956990228f735836c8375ea3d4ea38c074
$ curl https://github.com/mkmik/foo/releases/download/v2.0/foo.sha256
0eb56015fdd2281fdbe3e804bbfd82739695e335d63dabecdff09ced11851aae foo
$ curl -o foo-1.0-to-2.0.patch https://someserver/f52e22a6ef467d206e884794e39880956990228f735836c8375ea3d4ea38c074-to-0eb56015fdd2281fdbe3e804bbfd82739695e335d63dabecdff09ced11851aae.bpatch
$ bspatch foo foo-2.0 foo-1.0-to-2.0.patch
```

Usually this logic would be embedded in the `foo` binary itself.

## Contributing

PRs accepted
