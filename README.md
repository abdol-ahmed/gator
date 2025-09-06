# Welcome to the Blog Aggregator

This CLI tool allows users to add and manage RSS feeds, storing the collected posts in a PostgreSQL database. Users can follow or unfollow feeds and view summaries of aggregated posts—complete with links—directly in their terminal. The project targets users familiar with Go and SQL, and its key learning goals are:
- Integrating a Go application with PostgreSQL
- Practicing SQL skills (using sqlc and goose for typesafe migrations)
- Writing a long-running Go service for continuous RSS feed collection and storage
It’s designed to help you track your favorite blogs, news, and podcasts efficiently with aggregated updates.