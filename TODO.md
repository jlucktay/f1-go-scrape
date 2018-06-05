# TODO

- pipeline pattern
- take base URL as input (`flags` package)
- also take element selector pattern as another arg?
- send elements down the pipe (channel)
- take elements off channel and fetch
- rinse/repeat

## Pipeline layout

1. Starting URL (`string`)
1. `*goquery.Document`
1. ...?
1. image URL (`string`)
1. size-checker
- May need to reverse the order here, since we need to download the image in order to be able to check the size
1. downloader

Get:

- image links
- other pages which might contain images
- stay on the same domain

Branching paths in the pipe?
From the document, get 1) `<a>` and 2) `<img>` then send `<a>` down one channel and `<img>` down another.

Need a master map to maintain a list of all crawled URLs.

Do a HEAD request on an image URL and see if we can get the dimensions?
Then have a minimum threshold to only download nice big high-res images.
> Nope, this is not a thing. Look at `examples/image-resolution/`.
