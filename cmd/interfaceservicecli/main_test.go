// Copyright 2019 Intel Corporation. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main_test

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	cli "github.com/otcshare/edgecontroller/cmd/interfaceservicecli"
	pb "github.com/otcshare/edgecontroller/pb/ela"
	"io/ioutil"
	"os"
)

var _ = Describe("CLI tests", func() {

	BeforeEach(func() {
		cli.Cfg.Endpoint = ""
		cli.Cfg.ServiceName = ""
		cli.Cfg.Cmd = ""
		cli.Cfg.Val = ""
		cli.Cfg.CertsDir = "./certs"
	})

	AfterEach(func() {
		Iserv.getAllReturnNi = nil
		Iserv.getAllReturnErr = nil
	})

	Context("start Cli without command", func() {
		It("should return help print", func() {
			saveStd := os.Stdout
			read, write, _ := os.Pipe()
			os.Stdout = write

			err := cli.StartCli()
			Expect(err).NotTo(HaveOccurred())

			write.Close()
			out, _ := ioutil.ReadAll(read)
			os.Stdout = saveStd
			outString := string(out[:])
			Expect(outString).To(Equal(HelpOut))
		})
	})

	Context("'help' command", func() {
		It("should return help print", func() {
			saveStd := os.Stdout
			read, write, _ := os.Pipe()
			os.Stdout = write

			cli.Cfg.Cmd = "help"
			err := cli.StartCli()
			Expect(err).NotTo(HaveOccurred())

			write.Close()
			out, _ := ioutil.ReadAll(read)
			os.Stdout = saveStd
			outString := string(out[:])
			Expect(outString).To(Equal(HelpOut))
		})
	})

	Context("unrecognized command", func() {
		It("should return 'Unrecognized action' + help print", func() {
			saveStd := os.Stdout
			read, write, _ := os.Pipe()
			os.Stdout = write

			cli.Cfg.Cmd = "test123"
			err := cli.StartCli()
			Expect(err).NotTo(HaveOccurred())

			write.Close()
			out, _ := ioutil.ReadAll(read)
			os.Stdout = saveStd
			outString := string(out[:])
			Expect(outString).To(Equal(WarningOut))
		})
	})

	Context("'attach' command on existing interface", func() {
		It("should call 'BulkUpdate'", func() {
			cli.Cfg.Endpoint = Iserv.Endpoint
			cli.Cfg.ServiceName = "localhost"

			cli.Cfg.Cmd = "attach"
			cli.Cfg.Val = "5201:54:00.0"

			Ni := &pb.NetworkInterface{
				Driver:     0,
				Id:         "1",
				MacAddress: "5201:54:00.0",
			}

			Iserv.getAllReturnNi = &pb.NetworkInterfaces{
				NetworkInterfaces: []*pb.NetworkInterface{Ni},
			}

			err := cli.StartCli()
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("'attach' command on unknown interface", func() {
		It("should dont call 'BulkUpdate'", func() {
			cli.Cfg.Endpoint = Iserv.Endpoint
			cli.Cfg.ServiceName = "localhost"

			cli.Cfg.Cmd = "attach"
			cli.Cfg.Val = "5222:54:00.0"

			Ni := &pb.NetworkInterface{
				Driver:     0,
				Id:         "1",
				MacAddress: "5201:54:00.0",
			}

			Iserv.getAllReturnNi = &pb.NetworkInterfaces{
				NetworkInterfaces: []*pb.NetworkInterface{Ni},
			}

			err := cli.StartCli()
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("'detach' command on existing interface", func() {
		It("should call 'BulkUpdate'", func() {
			cli.Cfg.Endpoint = Iserv.Endpoint
			cli.Cfg.ServiceName = "localhost"

			cli.Cfg.Cmd = "detach"
			cli.Cfg.Val = "5201:54:00.0"

			Ni := &pb.NetworkInterface{
				Driver:     1,
				Id:         "1",
				MacAddress: "5201:54:00.0",
			}

			Iserv.getAllReturnNi = &pb.NetworkInterfaces{
				NetworkInterfaces: []*pb.NetworkInterface{Ni},
			}

			err := cli.StartCli()
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("'get' command", func() {
		It("should return interfaces", func() {
			cli.Cfg.Endpoint = Iserv.Endpoint
			cli.Cfg.ServiceName = "localhost"

			cli.Cfg.Cmd = "get"

			Ni := &pb.NetworkInterface{
				Driver:     0,
				Id:         "1",
				MacAddress: "5200:54:00.0",
			}

			Ni2 := &pb.NetworkInterface{
				Driver:     1,
				Id:         "2",
				MacAddress: "5201:54:00.0",
			}

			Iserv.getAllReturnNi = &pb.NetworkInterfaces{
				NetworkInterfaces: []*pb.NetworkInterface{Ni, Ni2},
			}

			err := cli.StartCli()
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("'get' command with no interfaces", func() {
		It("should return error", func() {
			cli.Cfg.Endpoint = Iserv.Endpoint
			cli.Cfg.ServiceName = "localhost"

			cli.Cfg.Cmd = "get"

			Iserv.getAllReturnNi = &pb.NetworkInterfaces{
				NetworkInterfaces: []*pb.NetworkInterface{},
			}

			err := cli.StartCli()
			Expect(err)
		})
	})

	Context("'get' command with 'GetAll' error", func() {
		It("should return error", func() {
			cli.Cfg.Endpoint = Iserv.Endpoint
			cli.Cfg.ServiceName = "localhost"

			cli.Cfg.Cmd = "get"

			Iserv.getAllReturnNi = &pb.NetworkInterfaces{
				NetworkInterfaces: []*pb.NetworkInterface{},
			}

			Iserv.getAllReturnErr = errors.New("")

			err := cli.StartCli()
			Expect(err)
		})
	})
})
