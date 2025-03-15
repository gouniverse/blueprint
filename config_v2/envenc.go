package config_v2

import (
	"crypto/sha256"
	"fmt"

	"github.com/gouniverse/envenc"
)

/**
 * ENVENC_KEY_PUBLIC: Public Key for Envenc Vault Encryption
 *
 * This constant stores the public key used to encrypt your envenc vault file.
 * * **Important Security Information:**
 * - This public key, combined with a corresponding private key, is used
 *   as input to a secure **one-way** hashing function to derive the final
 *   encryption key.
 * - Both the private and public keys must be at least 32-character strings,
 *   composed of randomly generated characters, numbers, and symbols.
 * - **DO NOT store the actual final key anywhere.** It should be generated dynamically when needed.
 * - **DO NOT directly commit the actual PRIVATE key to version control.** Use environment variables or secure configuration management.
 * - Replace "YOUR_PUBLIC_KEY" with your actual 32-character public key.
 * - The associated private key must be kept extremely secure.
 * - Ensure that the random number generator used to create the keys is cryptographically secure (CSPRNG).
 * - **Ideally, the public key should be obfuscated. See envenc for more details.**
 *
 * Example:
 * const ENVENC_KEY_PUBLIC = "aBcD123$456!eFgH789%iJkL0mNoPqRsTuVwXyZ"; // Replace with your actual key
 */
const ENVENC_KEY_PUBLIC = "YOUR_PUBLIC_KEY"

// deriveEnvEncKey: Derives the final envenc encryption key.
//
// This function generates the final encryption key used for securing the envenc vault file.
// It combines a private key with a public key, performs a secure hashing
// operation, and returns the resulting hash as the encryption key.
//
// Business Logic:
//  1. Public Key Retrieval and Deobfuscation:
//     - The function retrieves the pre-configured public key (ENVENC_KEY_PUBLIC).
//     - It then deobfuscates the public key using the envenc.Deobfuscate function.
//     This step adds a layer of security against reverse engineering.
//  2. Temporary Key Generation:
//     - The function concatenates the public key with the private key.
//     - Key Concatenation Order: The order of concatenation (public key + private key)
//     is vital and MUST be consistent across all key generation and validation processes.
//
// 3. Secure Hashing:
//   - The function calculates the SHA256 hash of the temporary key.
//   - SHA256 is a robust one-way hashing algorithm. This ensures that the original keys cannot
//     be recovered from the generated hash, providing strong cryptographic security.
//
// 4. Final Key Formatting:
//   - The resulting hash (a byte array) is converted into a hexadecimal string representation.
//
// 5. Key Return:
//   - The function returns the hexadecimal string, which is the final encryption key.
//
// Parameters:
// - envencPrivateKey (string): The private key used in the encryption key derivation.
// This key MUST be kept strictly confidential and handled with extreme care.
//
// Returns:
// - string: The final envenc encryption key as a hexadecimal string.
// - error: Returns an error if the public key deobfuscation fails.
//
// Security Considerations:
// - Private Key Protection: The `envencPrivateKey` is the most sensitive piece of information.
// It should never be stored in plain text or committed to version control. Use secure environment
// variables or configuration management systems.
// - Public Key Obfuscation: The public key is deobfuscated to prevent it from being easily
// extracted from compiled applications. While not as sensitive as the private key, obfuscation
// adds an extra layer of security.
// - One-Way Hashing: The use of SHA256 ensures that the key derivation process is one-way.
// It is computationally infeasible to derive the original private and public keys from the
// generated hash.
// - Key Generation Dynamics: The final encryption key is generated dynamically each time it is
// needed. It should not be stored persistently.
// - CSPRNG: Ensure the private and public keys are generated using a cryptographically secure
// pseudorandom number generator (CSPRNG).
// - Zeroize tempKey: The tempKey variable should be overwritten as soon as the hash is generated.
//
// Example:
// privateKey := "your_private_key"
// finalKey, err := deriveEnvEncKey(privateKey)
// if err != nil {
// // Handle error
// }
// // Use finalKey for encryption
func deriveEnvEncKey(envencPrivateKey string) (string, error) {
	envencPublicKey := ENVENC_KEY_PUBLIC

	if envencPublicKey == "" {
		return "", fmt.Errorf("envenc public key is empty")
	}

	if envencPrivateKey == "" {
		return "", fmt.Errorf("envenc private key is empty")
	}

	if len(envencPublicKey) < 32 {
		return "", fmt.Errorf("envenc public key is too short (should be at least 32 characters)")
	}

	if len(envencPrivateKey) < 32 {
		return "", fmt.Errorf("envenc private key is too short (should be at least 32 characters)")
	}

	envencPublicKey, err := envenc.Deobfuscate(envencPublicKey)

	if err != nil {
		return "", fmt.Errorf("failed to deobfuscate public key: %w", err)
	}

	tempKey := envencPublicKey + envencPrivateKey

	hash := sha256.Sum256([]byte(tempKey))
	realKey := fmt.Sprintf("%x", hash)
	tempKey = "" //Zeroize tempKey

	return realKey, nil
}
