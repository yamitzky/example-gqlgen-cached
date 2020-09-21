# gqlgen well-cached example

Well-cached gqlgen server using APQ and Cache-Control

## Abstract

GraphQL API is a little bit harder to cache than REST API. This example repository shows several techniques to improve cache.

I am developing a GraphQL API for iOS/Android apps and confirmed it by curl. Some approaches do not work well for Apollo and GraphQL client libraries.

## Improvement #1: Persisted Query

https://github.com/yamitzky/example-gqlgen-cached/commit/5917721c4cad555ccdac3eef24f325247ca1ea1e

## Improvement #2: Cache-Control Header

https://github.com/yamitzky/example-gqlgen-cached/commit/f4e956fb36086aeee11fb52a6e978af8a294a9ac

## Improvement #3: Resolver-level cache

https://github.com/yamitzky/example-gqlgen-cached/commit/05c6e5e69ed5d4b0de20cf0c3a433315bcfabd77
