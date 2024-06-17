package curl

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/spf13/cobra"
)

func Cmd(newClient func() api.Client) *cobra.Command {
	method := http.MethodGet
	form := map[string]string{}
	data := ""
	curlCmd := &cobra.Command{
		Use:  "curl",
		Long: `calls an arbitrary Synology API based on user input, and returns the response as pretty-printed JSON`,
		Example: `synoctl curl 'api=SYNO.FileStation.Mount.List&method=get&version=1'
synoctl curl -X POST --data 'api=SYNO.FileStation.Mount.List&method=get&version=1'
synoctl curl -F 'api=SYNO.FileStation.Mount.List' -F 'method=get' -F 'version=1'
`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			requestFunc := api.GET[interface{}]
			var setter api.ValueTransformer
			if len(args) > 0 {
				s, err := setterFromRawString(args[0])
				if err != nil {
					return err
				}
				setter = s
			} else if data != "" {
				s, err := setterFromRawString(data)
				if err != nil {
					return err
				}
				setter = s
			} else {
				setter = setterFromStringMap(form)
			}

			if method == http.MethodPost {
				requestFunc = api.POST[interface{}]
			}

			result, err := requestFunc(newClient(), setter)
			if err != nil {
				return err
			}
			encoder := json.NewEncoder(os.Stdout)
			encoder.SetIndent("", "  ")
			return encoder.Encode(&result)
		},
	}

	curlCmd.Flags().StringVarP(&method, "request", "X", http.MethodGet, "")
	curlCmd.Flags().StringToStringVarP(&form, "form", "F", form, "")
	curlCmd.Flags().StringVar(&data, "data", "", "")
	curlCmd.MarkFlagsMutuallyExclusive("data", "form")

	return curlCmd
}

func setterFromRawString(s string) (api.ValueTransformer, error) {
	values, err := url.ParseQuery(s)
	if err != nil {
		return nil, err
	}
	return func(query url.Values) {
		for k, v := range values {
			for _, s := range v {
				query.Set(k, s)
			}
		}
	}, nil
}

func setterFromStringMap(m map[string]string) api.ValueTransformer {
	return func(query url.Values) {
		for k, v := range m {
			query.Set(k, v)
		}
	}
}
