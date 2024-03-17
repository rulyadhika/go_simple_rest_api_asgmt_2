## Installation

You can clone the repository with this command, or download this [zip](https://github.com/rulyadhika/go_simple_rest_api_asgmt_2/archive/refs/heads/main.zip) file.

```bash
> git clone https://github.com/rulyadhika/go_simple_rest_api_asgmt_2
```

## Configuration
1. Change terminal directory to go_simple_rest_api_asgmt_2 folder
```bash
> cd go_simple_rest_api_asgmt_2
```

2. Run this command
```bash
> go mod download
```

3. Duplicate .env.example file and rename it to .env . Or you can run this command
```bash
> copy .env.example .env
```

4. Configure your server_port and database in .env file

5. Run local development server
```bash
> go run main.go
```