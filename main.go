package main

import (
	"fmt"
	"github.com/Nks/go-gcs-ffmpeg/models"
	"github.com/Nks/go-gcs-ffmpeg/services"
	"github.com/Nks/go-gcs-ffmpeg/utils"
	"log"
	"path"
)

func main() {
	tempDir := new(utils.TemporaryDirectory)
	tempDir, err := tempDir.CreateTempDirectory()

	if err != nil {
		log.Fatal("unable create temporary directory")
	}

	defer tempDir.CleanUp()

	params, err := models.ParseParams()

	if err != nil {
		log.Fatal(err)
	}

	client := new(services.GcsClient)

	err = client.CreateClient()

	if err != nil {
		log.Fatal(err)
	}

	url, err := client.CreateSelfSignedUrlForFile(params.Storage, params.Source, params.ServiceAccount)

	if err != nil {
		log.Fatal(err)
	}

	commands := []string{
		"-preset", "ultrafast", "-g", "48", "-sc_threshold", "0", "-hide_banner",
		"-map", "0:v", "-map", "0:a", "-map", "0:v", "-map", "0:a",
		"-s:v:0", "854x480", "-c:v:0", "h264", "-b:v:0", "500k",
		"-s:v:1", "1280x720", "-c:v:1", "h264", "-b:v:1", "1500k",
		"-c:a", "aac", "-ar", "8000", "-c:v", "h264",
		"-profile:v:0", "baseline", "-crf", "20",
		"-profile:v:1", "baseline",
		"-pix_fmt", "yuv420p",
		"-var_stream_map", "v:0,a:0,name:sd v:1,a:1,name:hd",
		"-strict", "-2", "-vsync", "2", "-f", "hls",
		"-master_pl_name", "master.m3u8",
		"-hls_list_size", "0", "-hls_time", "10",
		"-hls_segment_filename", path.Join(tempDir.GetTempPath(), "v%v/segments_%03d.ts"),
		path.Join(tempDir.GetTempPath(), "%v.m3u8"),
	}

	err = runTranscoder(url, commands)

	if err != nil {
		log.Fatal(err)
	}

	err = client.UploadStreamToGcs(tempDir.GetTempPath(), params, true)

	if err != nil {
		log.Fatal(err)
	}
}

func runTranscoder(inputPath string, command []string) error {
	trans := new(services.Transcoder)
	err := trans.Initialize(inputPath, command)

	if err != nil {
		return err
	}

	fmt.Println("Starting transcoding process")

	done := trans.Run(true)

	progress := trans.Output()

	for msg := range progress {
		fmt.Println("progress: ", msg.Progress)
	}

	err = <-done

	if err != nil {
		return fmt.Errorf("unable transcode video %s", err)
	}

	return nil
}
