ffmpeg -y -i examples/example.mp4 \
  -preset slow -g 48 -sc_threshold 0 \
  -map 0:0 -map 0:1
  -s:v:0 640x360 -c:v:0 libx264 -b:v:0 2000k \
  -s:v:1 960x540 -c:v:1 libx264 -b:v:1 365k \
  -c:a copy -var_stream_map "v:0,a:0 v:1,a:1" -master_pl_name master.m3u8 -f hls
  -hls_time 10 -hls_list_size 0 \
  -hls_segment_filename "v%v/fileSequence%d.ts" v%v/prog_index.m3u8


4.2 ffmpeg only:
ffmpeg -i examples/example.mp4 -y \
-preset slow -g 48 -sc_threshold 0 \
-map 0:v -map 0:a -map 0:v -map 0:a \
-s:v:0 854x480 -c:v:0 h264 -b:v:0 500k \
-s:v:1 1280x720 -c:v:1 h264 -b:v:1 1500k \
-c:a aac -ar 48000 -c:v h264 \
-profile:v:0 baseline -crf 20 \
-profile:v:1 baseline \
-pix_fmt yuv420p \
-var_stream_map "v:0,a:0,name:sd v:1,a:1,name:hd" \
-strict -2 \
-vsync 2 \
-f hls \
-master_pl_name master.m3u8 \
-hls_list_size 0 \
-hls_time 10 \
-hls_segment_filename "streams/v%v/segments_%03d.ts" \
streams/index%v.m3u8

ffmpeg -re -i "https://00e9e64bac679dfe2a8c222fa5fb1ce26ee90c67d965f97c98-apidata.googleusercontent.com/download/storage/v1/b/wmt-video-test/o/test%2Fexample.mp4?qk=AD5uMEtKhhgU1uO7jcvChbqcjKa7MkvI68IzAcr--FByF0tlycl1xy_3GOB4h2dbtVy4bWkXttILlvah4dbsC_USdU-Old1wJrDBDIxNPueNVDMRlQOCDaDG-TfWB7sLo7_h5Jf2-z-tSiVwkGjgwFIZpGYGglsF9O2Qz3nGbpQheXr_-Mp5tTBeRZE03L7o-QTYuVzG6yzzmQbisu_RY7zAlEBQXQQ5Bd_VAS8DZgDHdFp6V8XWFeVXV2vEj3A5_9-fjFyzg8_YV9p8Yo061zZr5fLGlU4iTeCmBS0peryMZC-1-ZSlh3nkSFpVgtxdfiuAZVpUEOUCqNjfiLom7V5UcBhbQcHWGcL0r3lJENn1zS9e-RhbN2WJiybZ1KGbzx_3qkH9N9GToQPoD9qUbBEQHNAv9A-0i-D29PBOEE-4qEPVKQAebPvzR0c4f0bCJ96EhoGs8Xm5CV8LuW-PwARVPhO3QYnW4zhik9Cbc0hScKXosgsyIOAk5eFNnh1aIgpm7hMJjhIwLcR7vf3cIvZ5EGLU2w_QoaBEYO4PnnoHNjaTPnTU-7opEN552DdzdS1oRxfpUJAbjBNgghijzVj3AvW3UTE5KkRj2qOuIsJ2MepRMLcpD0TBTe58faU6aUn7oB7wJIU1seZYopMnIlFoztTD3k1FVk48jb6e964rCgGsXSDlQ5e9jlVxn4sY4uvDfNGSDCSfRjL07XEE_Gb-4ziTRR6ce-GAO_TL6QeWntqLkRi-913sKQnVrxlCgZSvipztN-1AjLlTOJ3vCiAvHxQBVIhOdSRSjPdE_TvOA8BaYaJ_v-s" -y \
-preset slow -g 48 -sc_threshold 0 \
-map 0:v -map 0:a -map 0:v -map 0:a \
-s:v:0 854x480 -c:v:0 h264 -b:v:0 500k \
-s:v:1 1280x720 -c:v:1 h264 -b:v:1 1500k \
-c:a aac -ar 48000 -c:v h264 \
-profile:v:0 baseline -crf 20 \
-profile:v:1 baseline \
-pix_fmt yuv420p \
-var_stream_map "v:0,a:0,name:sd v:1,a:1,name:hd" \
-strict -2 \
-vsync 2 \
-f hls \
-master_pl_name master.m3u8 \
-hls_list_size 0 \
-hls_time 10 \
-hls_segment_filename "streams/v%v/segments_%03d.ts" \
streams/%v.m3u8