// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2019 Intel Corporation

package cnca

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	//"strconv"
	"strings"

	y2j "github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"k8s.io/klog"
)

// paSubscribeCmd represents subscribe command for policy-authorization
var paSubscribeCmd = &cobra.Command{
	Use: "subscribe",
	Short: "Subscribe to NGC AF PCF app-session notification" +
		"using YAML configuration file",
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Read file from the filename provided in command
		data, err := readInputData(cmd)
		if err != nil {
			fmt.Println(err)
			return
		}

		var c Header
		if err = yaml.Unmarshal(data, &c); err != nil {
			fmt.Println(err)
			return
		}

		if c.Kind != "ngc_policy_authorization" {
			fmt.Println(errors.New("`kind` missing or unknown in YAML file"))
			return
		}

		if args[0] != "" {
			var evSubscRespData []byte
			appSessionID := args[0]
			var s AFEventsSubscReqData
			if err = yaml.Unmarshal(data, &s); err != nil {
				fmt.Println(err)
				return
			}

			paEvSubscReqData := getpaEvSubscReqData(s)

			var evSubsc []byte
			evSubsc, err = json.Marshal(paEvSubscReqData)
			if err != nil {
				fmt.Println(err)
				return
			}
			evSubscRespData, err = AFSubscribePaAppSession(args[0], evSubsc)
			if err != nil {
				klog.Info(err)
				if err.Error() == "HTTP failure: 500" && evSubscRespData != nil {
					//print response data EventsSubscReqData and EventsNotification
				}
				return
			}
			fmt.Printf("App Session %s patched\n", args[0])
			return
		}
		fmt.Println(errors.New("Invalid input(s)"))
	},
}

func init() {

	const help = `Subscribe to NGC AF PCF app-session notification
using YAML configuration file

Usage:
  cnca policy-authorization subscribe <appSessionID> -f <config_patch.yml>

Example:
  cnca policy-authorization subscribe <appSessionID> -f <config_patch.yml>

Flags:
  -h, --help       help
  -f, --filename   YAML configuration file
`

	//add policy-authorization `subscribe` command
	paCmd.AddCommand(paSubscribeCmd)
	paSubscribeCmd.Flags().StringP("filename", "f", "", "YAML configuration file")
	_ = paSubscribeCmd.MarkFlagRequired("filename")
	paSubscribeCmd.SetHelpTemplate(help)
}

func getpaEvSubscReqData(inputPaEvSubscReqData AFEventsSubscReqData) EventsSubscReqData {
	var paEvSubscReqData EventsSubscReqData

	paEvSubscReqData.Events = make([]EventSubscription, len(inputPaEvSubscReqData.Events))
	for i, inputEvents := range inputPaEvSubscReqData.Events {
		var events EventSubscription

		events.Event = Event(inputEvents.Event)
		events.NotifMethod = NotifMethod(inputEvents.NotifMethod)

		paEvSubscReqData.Events[i] = events
	}

	//NotifURI
	paEvSubscReqData.NotifURI = inputEvSubsc.NotifURI

	//UsgThres
	if inputPaEvSubscReqData.UsgThres != nil {
		usgThres := UsageThreshold(*inputPaEvSubscReqData.UsgThres)
		paEvSubscReqData.UsgThres = &usgThres
	}
	paEvSubscReqData.EvSubsc = &evSubsc
}
