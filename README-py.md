# Douban API

## Overview

This repository contains the Douban API implementation using FastAPI.

## Branches

- **py**: Contains the FastAPI implementation of the Douban API.
- **go**: An empty branch for future development.

## Getting Started

### Prerequisites

- Python 3.8 or higher
- Poetry

### Cloning the Repository

1. Clone the repository:

   ```bash
   git clone https://github.com/Fromsko/douban_api.git
   cd douban_api
   ```

2. Checkout the `py` branch:

   ```bash
   git checkout py
   ```

### Setting Up the Environment

1. Install dependencies using Poetry:

   ```bash
   poetry install
   ```

2. Activate the virtual environment:

   ```bash
   poetry shell
   ```

### Running the Application

1. Start the FastAPI application:

   ```bash
   uvicorn douban_api.web.application:app --reload
   ```

2. Access the API documentation:

   - Open your browser and go to [http://127.0.0.1:8000/docs](http://127.0.0.1:8000/docs) for Swagger UI.
   - Alternatively, go to [http://127.0.0.1:8000/redoc](http://127.0.0.1:8000/redoc) for ReDoc.

## Contributing

Please refer to the [CONTRIBUTING.md](CONTRIBUTING.md) file for guidelines on how to contribute to the project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
