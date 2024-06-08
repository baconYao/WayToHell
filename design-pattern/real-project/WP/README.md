# WP

WP is a video convertor which leverage the ffmpeg to help us transform video to MP4 or HLS format. This tool has been implemented with Worker Pool Patterns.

## Usage

```bash
$ cd app
$ go run .
```

### output example

By default, we give tool 4 4 worders to handle 4 jobs. The output is like below that one example shows `false` result because of `bad.txt` file be used, others shows `true`.

```bash
New: creating worker pool
vd.Run: starting worker pool by running workers
newVideoWorker: creating video worker id 1
w.start(): starting worker id 1
newVideoWorker: creating video worker id 2
w.start(): starting worker id 2
newVideoWorker: creating video worker id 3
w.start(): starting worker id 3
newVideoWorker: creating video worker id 4
w.start(): starting worker id 4
NewVideo: New video created: 1 ./input/puppy1.mp4
NewVideo: New video created: 2 ./input/bad.txt
NewVideo: New video created: 3 ./input/puppy2.mp4
NewVideo: New video created: 4 ./input/puppy2.mp4
vd.dispatch: sending job 1 to worker job queue
vd.dispatch: sending job 2 to worker job queue
vd.dispatch: sending job 3 to worker job queue
vd.dispatch: sending job 4 to worker job queue
w.processVideoJob: starting encode on video 4
w.processVideoJob: starting encode on video 2
w.processVideoJob: starting encode on video 1
w.processVideoJob: starting encode on video 3
v.sendToNotifyChan: sending message to notifyChan for video id 2
i: 1 msg: {2 false encode failed for 2: error executing ([-i ./input/bad.txt -print_format json -show_format -show_streams -show_error]) | error: exit status 1 | message: {
    "error": {
        "code": -1094995529,
        "string": "Invalid data found when processing input"
    }
}
 ffprobe version 6.0 Copyright (c) 2007-2023 the FFmpeg developers
  built with Apple clang version 14.0.3 (clang-1403.0.22.14.1)
  configuration: --prefix=/opt/homebrew/Cellar/ffmpeg/6.0_2 --enable-shared --enable-pthreads --enable-version3 --cc=clang --host-cflags= --host-ldflags= --enable-ffplay --enable-gnutls --enable-gpl --enable-libaom --enable-libaribb24 --enable-libbluray --enable-libdav1d --enable-libjxl --enable-libmp3lame --enable-libopus --enable-librav1e --enable-librist --enable-librubberband --enable-libsnappy --enable-libsrt --enable-libsvtav1 --enable-libtesseract --enable-libtheora --enable-libvidstab --enable-libvmaf --enable-libvorbis --enable-libvpx --enable-libwebp --enable-libx264 --enable-libx265 --enable-libxml2 --enable-libxvid --enable-lzma --enable-libfontconfig --enable-libfreetype --enable-frei0r --enable-libass --enable-libopencore-amrnb --enable-libopencore-amrwb --enable-libopenjpeg --enable-libspeex --enable-libsoxr --enable-libzmq --enable-libzimg --disable-libjack --disable-indev=jack --enable-videotoolbox --enable-audiotoolbox --enable-neon
  libavutil      58.  2.100 / 58.  2.100
  libavcodec     60.  3.100 / 60.  3.100
  libavformat    60.  3.100 / 60.  3.100
  libavdevice    60.  1.100 / 60.  1.100
  libavfilter     9.  3.100 /  9.  3.100
  libswscale      7.  1.100 /  7.  1.100
  libswresample   4. 10.100 /  4. 10.100
  libpostproc    57.  1.100 / 57.  1.100
./input/bad.txt: Invalid data found when processing input
 }
v.sendToNotifyChan: sending message to notifyChan for video id 1
i: 2 msg: {1 true video id 1 processed and saved as ./output/puppy1.mp4 puppy1.mp4}
v.sendToNotifyChan: sending message to notifyChan for video id 4
i: 3 msg: {4 true video id 4 processed and saved as ./output/puppy2.mp4 puppy2.mp4}
v.sendToNotifyChan: sending message to notifyChan for video id 3
i: 4 msg: {3 true video id 3 processed and saved as ./output/jf1hNXRfD3.m3u8 jf1hNXRfD3.m3u8}
Done!
```
