# HTTPS Certificate Generator for Local Development

This Go program generates self-signed SSL/TLS certificates for local web development purposes.

## Features

- Generates self-signed SSL/TLS certificates and private keys
- Supports both RSA and ECDSA key types
- Allows customization of certificate validity period
- Generates certificates for localhost and common local development domains
- Provides clear output showing where certificates were saved
- Includes proper error handling and user-friendly messages
- Makes the certificates compatible with major browsers for local development

## Usage

Run the program and follow the interactive prompts:

```bash
cd session/z/certgen
go run generate_cert.go
```

### Configuration Options

1. **Key Type**: Choose between `rsa` or `ecdsa`
   - RSA: 2048-bit key by default
   - ECDSA: 256-bit key by default

2. **Key Size**: 
   - RSA: 2048, 4096 bits (minimum 512)
   - ECDSA: 256, 384, 521 bits

3. **Validity Period**: Number of days the certificate should be valid (default: 365 days)

4. **Output Directory**: Directory where certificate files will be saved (default: ./certs)

## Generated Files

The program creates two files in the specified output directory:

- `localhost.crt`: The SSL/TLS certificate file
- `localhost.key`: The private key file

## Example Usage

### Generate RSA certificate (default):
```bash
go run generate_cert.go
# Accept defaults by pressing Enter for all prompts
```

### Generate ECDSA certificate:
```bash
go run generate_cert.go
# Enter: ecdsa
# Enter: 384
# Enter: 30 (30 days validity)
# Enter: ./certs_ecdsa (custom output directory)
```

## Certificate Details

The generated certificate includes:
- Common Name: "Local Development Certificate"
- Organization: "Local Development"
- Domains: localhost, 127.0.0.1, ::1
- Key Usage: Digital signature and key encipherment
- Extended Key Usage: Server authentication

## Browser Compatibility

The certificates are generated with standard parameters that make them compatible with major browsers for local development. However, since they are self-signed, browsers may display a security warning. You can typically proceed through the warning to use the certificate locally.

## Requirements

- Go 1.16 or higher
- No external dependencies

## License

This program is provided as-is for local development use.