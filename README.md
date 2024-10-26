# Wordie API.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- You have installed Docker.
- You have installed Docker-Compose
- You have installed the Redis-CLI

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/wordie.git
   cd wordie
   ```

2. Add your API key as an environment variable to you docker-compose:
   API_KEY=your_api_key_here

3. Build the Docker containers:
   ```sh
   docker-compose --build
   ```

## Usage

1. Start the Docker containers:

   ```sh
   docker-compose up
   ```

2. Start Redis Session
    ```sh
    redis-cli
    ```

3. Set test values in Redis using `HSET`:
   ```sh
   HSET test_key_1 field1 "value1" field2 "value2"
   HSET test_key_2 field1 "value1" field2 "value2"
   ```
4. Sell all test values using `HGETALL`:
    ```sh
    HGETALL <test_key>
    ```
