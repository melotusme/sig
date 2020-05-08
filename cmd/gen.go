/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

var (
	catalog        string
	sn             string
	platform       string
	ip             string
	privateKeyPath string
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate credentials",
	Long:  `generate credentials`,
	Run: func(cmd *cobra.Command, args []string) {
		privateKeyBytes, err := ioutil.ReadFile(privateKeyPath)
		if err != nil {
			log.Fatalln(err)
		}
		block, _ := pem.Decode(privateKeyBytes)
		privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			log.Fatalln(err)
		}
		kws := KWS{
			Catalog:  catalog,
			Sn:       sn,
			Platform: platform,
			Ip:       ip,
		}
		marshal, err := json.Marshal(kws)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(marshal))
		dst := base64.StdEncoding.EncodeToString(marshal)
		fmt.Println(dst)
		hash := sha256.New()
		hash.Write([]byte(dst))
		encryptedData, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash.Sum(nil))

		fmt.Println(fmt.Sprintf("%s.%s", dst, base64.StdEncoding.EncodeToString(encryptedData)))
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	genCmd.Flags().StringVar(&catalog, "catalog", "webservice.abc", "catalog")
	genCmd.Flags().StringVar(&sn, "sn", "kuaiyun", "service name")
	genCmd.Flags().StringVar(&platform, "platform", "IS", "platform")
	genCmd.Flags().StringVar(&ip, "ip", "192.168.199.247", "ip")
	genCmd.Flags().StringVar(&privateKeyPath, "privateKeyPath", "ca.key", "private key path")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
