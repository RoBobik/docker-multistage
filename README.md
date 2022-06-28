# Example app for Docker multi-stage build

This is a simple web app which shows some random quotes.

The app has 2 parts which live in their respective directories:

- `frontend` -- Svelte app which fetches a random quote from the backend.
- `backend` -- Go app which serves static assets (the frontend) and offers a single API endpoint `/api/quotes/radnom`, which returns a random quote from a fixed list There is an artificial delay of 1 second when querying the API.

An example Dockerfile is provided to demonstrate how a Docker multi-stage build could be implemented.
