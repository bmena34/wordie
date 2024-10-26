# Wordie

Wordie is a project designed to help users manage and analyze words efficiently.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- You have installed Docker.
- You have installed Docker-Compose

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/wordie.git
   cd wordie
   ```

2. Add your API key as an environment variable:
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

2. Set test values in Redis using `HSET`:
   ```sh
   redis-cli
   HSET test_key_1 field1 "value1" field2 "value2"
   HSET test_key_2 field1 "value1" field2 "value2"
   ```
