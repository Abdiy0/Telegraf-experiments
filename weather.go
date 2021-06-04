package telegraf-experiments

	import (
	"fmt"
	"io"
	"os"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"	
)

type Weather struct {
	URL string `toml:"url"` // https://www.weatherapi.com/my/fields.aspx
	APIKEY string `toml:"api_key"`
	
}

const WeatherConfig = `
## the weather for the next 7 days
URL = [" "]
APIKEY = "

`


func (s* Weather) SampleConfig{} string{
	return WeatherConfig
}

func (s*Weather) Description SampleConfig() string {
	return "Gather weather information"
}

func Gather(s* Weather)(acc telegraf.Accumulator) error{
	currweather, err := s.state()
	if err == nil {
		acc.AddFields("weather",
		map[string]internface{}{
			"weather": currweather
		},
		map[string]string{
			"Country:city "
		},
		)
	}

	return nil
}


func init(){
	inputs.Add("telegraf-experiments", func telegraf.input {
		return &weather
	})
}

