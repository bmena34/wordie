# Wordie API.

API project using Golang and Redis.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- You have installed Docker.
- You have installed Docker-Compose
- You have installed the Redis-CLI

## Installation

1. Clone the repository:

2. Add your API key as an environment variable to your docker-compose:
   ```sh
   API_KEY=<your_api_key_here>
   ```

3. Build the Docker containers (this will also start the containers):
   ```sh
   docker-compose up --build
   ```

## Usage

1. Start the Docker containers (if build isn't required):

   ```sh
   docker-compose up
   ```

2. Start Redis Session (in a different terminal window):

   ```sh
   redis-cli
   ```

3. How to set test values in Redis using `migration.sh`:

- This will add all Category/Word pairs with a key starting at 1.
- Add a 2 column csv to the root level of the project with the first column being category and second being the word.
- Change Script permissions.

   ```sh
   chmod +x migration.sh
   ```

- The Script accepts the csv as and argument. Example of how to run:

   ```sh
   ./migration.sh <example.csv>
   ```

4. API currently has 2 endpoints:

   - Status OK check:

   ```sh
   http://localhost:3000
   ```

   - Word Endpoint (GET) returns json with category and word keys and respected values. This is retrieved by the key created by the order of the CSV:

   ```sh
   http://localhost:3000/word/{id}
   ```

5. Authentication currently handled by API-TOKEN.

6. More to come...
