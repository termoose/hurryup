package display

import (
	"fmt"
	"hurryup/internal/speed"
)

func PromFile(tester speed.Tester) error {
	data := speed.Data{ServerData: tester.GetServerData(), TesterData: tester.GetTesterData()}

	output := fmt.Sprintf("hurryup_download_speed %.1f\n"+
		"hurryup_upload_speed %.1f\n"+
		"hurryup_latency %d\n"+
		"hurryup_location %s\n"+
		"hurryup_country %s",
		data.DownloadRate, data.UploadRate, data.Latency, data.Name, data.Country)

	fmt.Println(output)

	return nil
}
