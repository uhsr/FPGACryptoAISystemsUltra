# FPGACryptoAISystemsUltra

## Description



## Features

- Leverages pipelined Advanced Encryption Standard (AES) cores on the FPGA for high-throughput encryption and decryption.
- Implements side-channel attack countermeasures, including masking and shuffling, within the cryptographic cores.
- Integrates an Automatic Identification System (AIS) data parser and decoder optimized for low-latency processing on the FPGA.
- Utilizes ultra-low latency direct memory access (DMA) controllers to transfer AIS data between the FPGA and the host processor.
- Employs a dynamically reconfigurable architecture to adapt to different cryptographic algorithms and AIS message types.
- Incorporates a hardware-based true random number generator (TRNG) compliant with NIST SP 800-90B for key generation.
- Provides a secure boot mechanism with FPGA bitstream authentication and encryption to prevent tampering.
- Offers a comprehensive API for integrating the FPGA-based crypto-AIS system with existing software applications.
## Installation

```bash
go get github.com/uhsr/FPGACryptoAISystemsUltra
```

## Usage

```go
package main

import (
    "fpgacryptoaisystemsultra/internal/fpgacryptoaisystemsultra"
)

func main() {
    app := fpgacryptoaisystemsultra.NewApp(false)
    app.Run()
}
```

## Contributing

We welcome contributions! Here's how to get started:

1. Fork this repository
2. Create a new branch for your feature (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -am 'Add some awesome feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.
