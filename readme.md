WIP

# GCSFfmpeg

Inspired by [Goffmpeg](https://github.com/xfrr/goffmpeg) package

# Dependencies
- [FFmpeg](https://www.ffmpeg.org/)
- [FFProbe](https://www.ffmpeg.org/ffprobe.html)
- [GCS](https://cloud.google.com/storage/)

# Supported platforms

 - OS X
 - Linux
 - Windows

# Getting started

```shell
go get github.com/Nks/go-gcs-ffmpeg
```

or run

```shell
./gcs-transcoder -source=test/example.mp4 -output=test/stream -storage=your-bucket-name
```

# Examples

Check `main.go` for the examples.

# FAQ

1. CORS settings for GCS:

Create cors.json file:
```
[
  {
    "origin": [
      "https://example.com"
    ],
    "responseHeader": [
      "Content-Type"
    ],
    "method": [
      "GET",
      "HEAD",
      "DELETE"
    ],
    "maxAgeSeconds": 3600
  }
]
```

Run command: `gsutil cors set cors.json gs://your-gcs-bucket`

2. Why HLS files downloading instead of playing?

- Check Content-Type for your files. You should set content type for the m3u8 and ts files. .m3u8 files should be `application/vnd.apple.mpegurl` and .ts files should be `video/mp2t`

