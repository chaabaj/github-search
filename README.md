# github-search
Github repository search API with statistics

This server application use the GitHub API for searching public repositories and make some basics statistics

Usage :

```bash
 export GITHUB_AUTH_TOKEN=YOUR_TOKEN
./github-search
```

Since we are using the GitHub Dev API v3 and the application do a lot of request.
The rate limit of API call is easily reached with no Authorization token
So you need to put in the environment variable GITHUB_AUTH_TOKEN one of your generated token
We only need the right to read the public repository

You can check at this address if you don't know how to generate a GitHub API Authorization token

https://github.com/blog/1509-personal-api-tokens

After go to your browser at http://localhost:8080 and search what you want :)
