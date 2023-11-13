package vr_sros

import (
	"bytes"
	"context"
	_ "embed"
	"strings"
	"text/template"

	"github.com/hairyhenderson/gomplate/v3"
	"github.com/hairyhenderson/gomplate/v3/data"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

// importing Default Config template at compile time
//
//go:embed ssh_keys.go.tpl
var SROSSSHKeysTemplate string

// mapSSHPubKeys goes over s.sshPubKeys and puts the supported keys to the corresponding
// slices associated with the supported SSH key algorithms.
// supportedSSHKeyAlgos key is a SSH key algorithm and the value is a pointer to the slice
// that is used to store the keys of the corresponding algorithm family.
// Two slices are used to store RSA and ECDSA keys separately.
// The slices are modified in place by reference, so no return values are needed.
func (s *vrSROS) mapSSHPubKeys(supportedSSHKeyAlgos map[string]*[]string) {
	for _, k := range s.sshPubKeys {
		sshKeys, ok := supportedSSHKeyAlgos[k.Type()]
		if !ok {
			log.Debugf("unsupported SSH Key Algo %q, skipping key", k.Type())
			continue
		}

		// extract the fields
		// <keytype> <key> <comment>
		keyFields := strings.Fields(string(ssh.MarshalAuthorizedKey(k)))

		*sshKeys = append(*sshKeys, keyFields[1])
	}
}

// SROSTemplateData holds ssh keys for template generation.
type SROSTemplateData struct {
	SSHPubKeysRSA   []string
	SSHPubKeysECDSA []string
}

// configureSSHPublicKeys cofigures public keys extracted from clab host
// on SR OS node using SSH.
func (s *vrSROS) configureSSHPublicKeys(
	ctx context.Context, addr, platformName,
	username, password string, pubKeys []ssh.PublicKey) error {
	tplData := SROSTemplateData{}

	// a map of supported SSH key algorithms and the template slices
	// the keys should be added to.
	// In mapSSHPubKeys we map supported SSH key algorithms to the template slices.
	supportedSSHKeyAlgos := map[string]*[]string{
		ssh.KeyAlgoRSA:      &tplData.SSHPubKeysRSA,
		ssh.KeyAlgoECDSA521: &tplData.SSHPubKeysECDSA,
		ssh.KeyAlgoECDSA384: &tplData.SSHPubKeysECDSA,
		ssh.KeyAlgoECDSA256: &tplData.SSHPubKeysECDSA,
	}

	s.mapSSHPubKeys(supportedSSHKeyAlgos)

	t, err := template.New("SSHKeys").Funcs(
		gomplate.CreateFuncs(context.Background(), new(data.Data))).Parse(SROSSSHKeysTemplate)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, tplData)
	if err != nil {
		return err
	}

	err = s.applyPartialConfig(ctx, s.Cfg.MgmtIPv4Address, scrapliPlatformName,
		defaultCredentials.GetUsername(), defaultCredentials.GetPassword(),
		buf,
	)

	return err
}