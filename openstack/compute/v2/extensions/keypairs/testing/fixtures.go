package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/YoungPioneers/huaweicloud-sdk-go/openstack/compute/v2/extensions/keypairs"
	th "github.com/YoungPioneers/huaweicloud-sdk-go/testhelper"
	"github.com/YoungPioneers/huaweicloud-sdk-go/testhelper/client"
)

// ListOutput is a sample response to a List call.
const ListOutput = `
{
	"keypairs": [
		{
			"keypair": {
				"fingerprint": "15:b0:f8:b3:f9:48:63:71:cf:7b:5b:38:6d:44:2d:4a",
				"name": "firstkey",
				"public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQC+Eo/RZRngaGTkFs7I62ZjsIlO79KklKbMXi8F+KITD4bVQHHn+kV+4gRgkgCRbdoDqoGfpaDFs877DYX9n4z6FrAIZ4PES8TNKhatifpn9NdQYWA+IkU8CuvlEKGuFpKRi/k7JLos/gHi2hy7QUwgtRvcefvD/vgQZOVw/mGR9Q== Generated by Nova\n"
			}
		},
		{
			"keypair": {
				"fingerprint": "35:9d:d0:c3:4a:80:d3:d8:86:f1:ca:f7:df:c4:f9:d8",
				"name": "secondkey",
				"public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQC9mC3WZN9UGLxgPBpP7H5jZMc6pKwOoSgre8yun6REFktn/Kz7DUt9jaR1UJyRzHxITfCfAIgSxPdGqB/oF1suMyWgu5i0625vavLB5z5kC8Hq3qZJ9zJO1poE1kyD+htiTtPWJ88e12xuH2XB/CZN9OpEiF98hAagiOE0EnOS5Q== Generated by Nova\n"
			}
		}
	]
}
`

// GetOutput is a sample response to a Get call.
const GetOutput = `
{
	"keypair": {
		"public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQC+Eo/RZRngaGTkFs7I62ZjsIlO79KklKbMXi8F+KITD4bVQHHn+kV+4gRgkgCRbdoDqoGfpaDFs877DYX9n4z6FrAIZ4PES8TNKhatifpn9NdQYWA+IkU8CuvlEKGuFpKRi/k7JLos/gHi2hy7QUwgtRvcefvD/vgQZOVw/mGR9Q== Generated by Nova\n",
		"name": "firstkey",
		"fingerprint": "15:b0:f8:b3:f9:48:63:71:cf:7b:5b:38:6d:44:2d:4a"
	}
}
`

// CreateOutput is a sample response to a Create call.
const CreateOutput = `
{
	"keypair": {
		"fingerprint": "35:9d:d0:c3:4a:80:d3:d8:86:f1:ca:f7:df:c4:f9:d8",
		"name": "createdkey",
		"private_key": "-----BEGIN RSA PRIVATE KEY-----\nMIICXAIBAAKBgQC9mC3WZN9UGLxgPBpP7H5jZMc6pKwOoSgre8yun6REFktn/Kz7\nDUt9jaR1UJyRzHxITfCfAIgSxPdGqB/oF1suMyWgu5i0625vavLB5z5kC8Hq3qZJ\n9zJO1poE1kyD+htiTtPWJ88e12xuH2XB/CZN9OpEiF98hAagiOE0EnOS5QIDAQAB\nAoGAE5XO1mDhORy9COvsg+kYPUhB1GsCYxh+v88wG7HeFDKBY6KUc/Kxo6yoGn5T\nTjRjekyi2KoDZHz4VlIzyZPwFS4I1bf3oCunVoAKzgLdmnTtvRNMC5jFOGc2vUgP\n9bSyRj3S1R4ClVk2g0IDeagko/jc8zzLEYuIK+fbkds79YECQQDt3vcevgegnkga\ntF4NsDmmBPRkcSHCqrANP/7vFcBQN3czxeYYWX3DK07alu6GhH1Y4sHbdm616uU0\nll7xbDzxAkEAzAtN2IyftNygV2EGiaGgqLyo/tD9+Vui2qCQplqe4jvWh/5Sparl\nOjmKo+uAW+hLrLVMnHzRWxbWU8hirH5FNQJATO+ZxCK4etXXAnQmG41NCAqANWB2\nB+2HJbH2NcQ2QHvAHUm741JGn/KI/aBlo7KEjFRDWUVUB5ji64BbUwCsMQJBAIku\nLGcjnBf/oLk+XSPZC2eGd2Ph5G5qYmH0Q2vkTx+wtTn3DV+eNsDfgMtWAJVJ5t61\ngU1QSXyhLPVlKpnnxuUCQC+xvvWjWtsLaFtAsZywJiqLxQzHts8XLGZptYJ5tLWV\nrtmYtBcJCN48RrgQHry/xWYeA4K/AFQpXfNPgprQ96Q=\n-----END RSA PRIVATE KEY-----\n",
		"public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQC9mC3WZN9UGLxgPBpP7H5jZMc6pKwOoSgre8yun6REFktn/Kz7DUt9jaR1UJyRzHxITfCfAIgSxPdGqB/oF1suMyWgu5i0625vavLB5z5kC8Hq3qZJ9zJO1poE1kyD+htiTtPWJ88e12xuH2XB/CZN9OpEiF98hAagiOE0EnOS5Q== Generated by Nova\n",
		"user_id": "fake"
	}
}
`

// ImportOutput is a sample response to a Create call that provides its own public key.
const ImportOutput = `
{
	"keypair": {
		"fingerprint": "1e:2c:9b:56:79:4b:45:77:f9:ca:7a:98:2c:b0:d5:3c",
		"name": "importedkey",
		"public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDx8nkQv/zgGgB4rMYmIf+6A4l6Rr+o/6lHBQdW5aYd44bd8JttDCE/F/pNRr0lRE+PiqSPO8nDPHw0010JeMH9gYgnnFlyY3/OcJ02RhIPyyxYpv9FhY+2YiUkpwFOcLImyrxEsYXpD/0d3ac30bNH6Sw9JD9UZHYcpSxsIbECHw== Generated by Nova",
		"user_id": "fake"
	}
}
`

// FirstKeyPair is the first result in ListOutput.
var FirstKeyPair = keypairs.KeyPair{
	Name:        "firstkey",
	Fingerprint: "15:b0:f8:b3:f9:48:63:71:cf:7b:5b:38:6d:44:2d:4a",
	PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQC+Eo/RZRngaGTkFs7I62ZjsIlO79KklKbMXi8F+KITD4bVQHHn+kV+4gRgkgCRbdoDqoGfpaDFs877DYX9n4z6FrAIZ4PES8TNKhatifpn9NdQYWA+IkU8CuvlEKGuFpKRi/k7JLos/gHi2hy7QUwgtRvcefvD/vgQZOVw/mGR9Q== Generated by Nova\n",
}

// SecondKeyPair is the second result in ListOutput.
var SecondKeyPair = keypairs.KeyPair{
	Name:        "secondkey",
	Fingerprint: "35:9d:d0:c3:4a:80:d3:d8:86:f1:ca:f7:df:c4:f9:d8",
	PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQC9mC3WZN9UGLxgPBpP7H5jZMc6pKwOoSgre8yun6REFktn/Kz7DUt9jaR1UJyRzHxITfCfAIgSxPdGqB/oF1suMyWgu5i0625vavLB5z5kC8Hq3qZJ9zJO1poE1kyD+htiTtPWJ88e12xuH2XB/CZN9OpEiF98hAagiOE0EnOS5Q== Generated by Nova\n",
}

// ExpectedKeyPairSlice is the slice of results that should be parsed from ListOutput, in the expected
// order.
var ExpectedKeyPairSlice = []keypairs.KeyPair{FirstKeyPair, SecondKeyPair}

// CreatedKeyPair is the parsed result from CreatedOutput.
var CreatedKeyPair = keypairs.KeyPair{
	Name:        "createdkey",
	Fingerprint: "35:9d:d0:c3:4a:80:d3:d8:86:f1:ca:f7:df:c4:f9:d8",
	PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQC9mC3WZN9UGLxgPBpP7H5jZMc6pKwOoSgre8yun6REFktn/Kz7DUt9jaR1UJyRzHxITfCfAIgSxPdGqB/oF1suMyWgu5i0625vavLB5z5kC8Hq3qZJ9zJO1poE1kyD+htiTtPWJ88e12xuH2XB/CZN9OpEiF98hAagiOE0EnOS5Q== Generated by Nova\n",
	PrivateKey:  "-----BEGIN RSA PRIVATE KEY-----\nMIICXAIBAAKBgQC9mC3WZN9UGLxgPBpP7H5jZMc6pKwOoSgre8yun6REFktn/Kz7\nDUt9jaR1UJyRzHxITfCfAIgSxPdGqB/oF1suMyWgu5i0625vavLB5z5kC8Hq3qZJ\n9zJO1poE1kyD+htiTtPWJ88e12xuH2XB/CZN9OpEiF98hAagiOE0EnOS5QIDAQAB\nAoGAE5XO1mDhORy9COvsg+kYPUhB1GsCYxh+v88wG7HeFDKBY6KUc/Kxo6yoGn5T\nTjRjekyi2KoDZHz4VlIzyZPwFS4I1bf3oCunVoAKzgLdmnTtvRNMC5jFOGc2vUgP\n9bSyRj3S1R4ClVk2g0IDeagko/jc8zzLEYuIK+fbkds79YECQQDt3vcevgegnkga\ntF4NsDmmBPRkcSHCqrANP/7vFcBQN3czxeYYWX3DK07alu6GhH1Y4sHbdm616uU0\nll7xbDzxAkEAzAtN2IyftNygV2EGiaGgqLyo/tD9+Vui2qCQplqe4jvWh/5Sparl\nOjmKo+uAW+hLrLVMnHzRWxbWU8hirH5FNQJATO+ZxCK4etXXAnQmG41NCAqANWB2\nB+2HJbH2NcQ2QHvAHUm741JGn/KI/aBlo7KEjFRDWUVUB5ji64BbUwCsMQJBAIku\nLGcjnBf/oLk+XSPZC2eGd2Ph5G5qYmH0Q2vkTx+wtTn3DV+eNsDfgMtWAJVJ5t61\ngU1QSXyhLPVlKpnnxuUCQC+xvvWjWtsLaFtAsZywJiqLxQzHts8XLGZptYJ5tLWV\nrtmYtBcJCN48RrgQHry/xWYeA4K/AFQpXfNPgprQ96Q=\n-----END RSA PRIVATE KEY-----\n",
	UserID:      "fake",
}

// ImportedKeyPair is the parsed result from ImportOutput.
var ImportedKeyPair = keypairs.KeyPair{
	Name:        "importedkey",
	Fingerprint: "1e:2c:9b:56:79:4b:45:77:f9:ca:7a:98:2c:b0:d5:3c",
	PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDx8nkQv/zgGgB4rMYmIf+6A4l6Rr+o/6lHBQdW5aYd44bd8JttDCE/F/pNRr0lRE+PiqSPO8nDPHw0010JeMH9gYgnnFlyY3/OcJ02RhIPyyxYpv9FhY+2YiUkpwFOcLImyrxEsYXpD/0d3ac30bNH6Sw9JD9UZHYcpSxsIbECHw== Generated by Nova",
	UserID:      "fake",
}

// HandleListSuccessfully configures the test server to respond to a List request.
func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-keypairs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

// HandleGetSuccessfully configures the test server to respond to a Get request for "firstkey".
func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-keypairs/firstkey", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

// HandleCreateSuccessfully configures the test server to respond to a Create request for a new
// keypair called "createdkey".
func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-keypairs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{ "keypair": { "name": "createdkey" } }`)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

// HandleImportSuccessfully configures the test server to respond to an Import request for an
// existing keypair called "importedkey".
func HandleImportSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-keypairs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `
			{
				"keypair": {
					"name": "importedkey",
					"public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDx8nkQv/zgGgB4rMYmIf+6A4l6Rr+o/6lHBQdW5aYd44bd8JttDCE/F/pNRr0lRE+PiqSPO8nDPHw0010JeMH9gYgnnFlyY3/OcJ02RhIPyyxYpv9FhY+2YiUkpwFOcLImyrxEsYXpD/0d3ac30bNH6Sw9JD9UZHYcpSxsIbECHw== Generated by Nova"
				}
			}
		`)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ImportOutput)
	})
}

// HandleDeleteSuccessfully configures the test server to respond to a Delete request for a
// keypair called "deletedkey".
func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-keypairs/deletedkey", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusAccepted)
	})
}
