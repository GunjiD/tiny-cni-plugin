package main

import (
	"fmt"
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/containernetworking/plugins/pkg/testutils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

//var _ = Describe("tiny cni plugin test", func() {
//	var targetNs ns.NetNS
//
//	// 前処理
//	BeforeEach(func() {
//		var err error
//		targetNs, err = testutils.NewNS()
//		Expect(err).NotTo(HaveOccurred())
//	})
//
//	// 後処理
//	ginkgo.AfterEach(func() {
//		targetNs.Close()
//	})
//
//	// スペック（テストケース）
//	// スペックの集合をスイートと呼ぶ
//	Context("AddCmd", func() {
//		When("failed", func() {
//			It("prev result is nil", func() {
//				const IFNAME = "eth0"
//				conf := `{
//	"cniVersion": "1.0.0",
//	"name": "tiny-cni-plugin-test",
//	"type": "sample",
//	"prevResult": {
//		"cniVersion": "0.3.0",
//		"interfaces": [
//			{
//				"name": "%s",
//				"sandbox": "%s"
//			}
//		],
//		"ips": [
//			{
//				"version": "4",
//				"address": "10.0.0.2/24",
//				"gateway": "10.0.0.1",
//				"interface": 0
//			}
//		],
//		"routes": []
//}`
//				conf = fmt.Sprintf(conf, IFNAME, targetNs.Path())
//				args := &skel.CmdArgs{
//					ContainerID: "dummy",
//					Netns:       targetNs.Path(),
//					IfName:      IFNAME,
//					StdinData:   []byte(conf),
//				}
//				_, _, err := testutils.CmdAddWithArgs(args, func() error { return cmdAdd(args) })
//				Expect(err).NotTo(MatchError("anotherAwesomeArg must be specified"))
//			})
//		})
//	})
//})

// Add
// Check
// Delete
var _ = Describe("sample test", func() {
	var targetNs ns.NetNS

	BeforeEach(func() {
		var err error
		targetNs, err = testutils.NewNS()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		targetNs.Close()
	})

	It("Works with a 0.3.0 config", func() {
		ifname := "eth0"
		conf := `{
	"cniVersion": "0.3.0",
	"name": "cni-plugin-sample-test",
	"type": "sample",
	"anotherAwesomeArg": "awesome",
	"prevResult": {
		"cniVersion": "0.3.0",
		"interfaces": [
			{
				"name": "%s",
				"sandbox": "%s"
			}
		],
		"ips": [
			{
				"version": "4",
				"address": "10.0.0.2/24",
				"gateway": "10.0.0.1",
				"interface": 0
			}
		],
		"routes": []
	}
}`
		conf = fmt.Sprintf(conf, ifname, targetNs.Path())
		args := &skel.CmdArgs{
			ContainerID: "dummy",
			Netns:       targetNs.Path(),
			IfName:      ifname,
			StdinData:   []byte(conf),
		}
		_, _, err := testutils.CmdAddWithArgs(args, func() error { return cmdAdd(args) })
		Expect(err).NotTo(HaveOccurred())
	})

	It("fails an invalid config", func() {
		conf := `{
	"cniVersion": "0.3.0",
	"name": "cni-plugin-sample-test",
	"type": "sample",
	"prevResult": {
		"interfaces": [
			{
				"name": "eth0",
				"sandbox": "/var/run/netns/test"
			}
		],
		"ips": [
			{
				"version": "4",
				"address": "10.0.0.2/24",
				"gateway": "10.0.0.1",
				"interface": 0
			}
		],
		"routes": []
	}
}`

		args := &skel.CmdArgs{
			ContainerID: "dummy",
			Netns:       targetNs.Path(),
			IfName:      "eth0",
			StdinData:   []byte(conf),
		}
		_, _, err := testutils.CmdAddWithArgs(args, func() error { return cmdAdd(args) })
		Expect(err).To(MatchError("anotherAwesomeArg must be specified"))
	})

	It("fails with CNI spec versions that don't support plugin chaining", func() {
		conf := `{
	"cniVersion": "0.2.0",
	"name": "cni-plugin-sample-test",
	"type": "sample",
	"anotherAwesomeArg": "foo"
}`

		args := &skel.CmdArgs{
			ContainerID: "dummy",
			Netns:       targetNs.Path(),
			IfName:      "eth0",
			StdinData:   []byte(conf),
		}
		_, _, err := testutils.CmdAddWithArgs(args, func() error { return cmdAdd(args) })
		Expect(err).To(MatchError("must be called as chained plugin"))
	})
})
