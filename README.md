# Datafy.ID Go Stuff

Monorepo berisi kumpulan hal-hal menarik terkait [Go](https://go.dev/).

Requirements:

- [Bazel](https://bazel.build/) - Install Bazel v5 & jdk-11
- [Gazelle](https://github.com/bazelbuild/bazel-gazelle)


TLDR;

```sh
cd ./hello
gazelle update
gazelle update-repos -from_file go.mod

# Build
cd ..
bazel build //hello:hello

# Cleanup
bazel clean --expunge
```
