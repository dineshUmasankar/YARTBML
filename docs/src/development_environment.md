---
title: "YARTBML Language Documentation"
author: [Joseph Porrino, Dinesh Umasankar, Katherine Banis, Paul Jensen]
date: "2024-03-15"
subject: "Markdown"
keywords: [Markdown, Example]
lang: "en"
...

# Development Environment and Runtime

This document outlines the tools used for development and runtime of the YARTBML langauge. 

## IDE

The IDE we will be using is Visual Studio Code. VSCode integrates our code editor, compliler, and version control system all in one place for easy builds and control of our langauge.

## Version Control System

We will be using Git as our version control system, and GitHub to host our repository in the cloud. 

## Testing

For testing we will be using Go's built in testing package.

## Documentation Tools

We will be using a github action that utilizes DocGen and Pandoc to automatically convert our markdown documents to PDF format.

## Runtime

The YARTBML langauge utilizes the Go runtime environment. The runtime environment gets linked to the YARTBML language when the project is built to an executable. The Go runtime includes Goroutines for concurrent code execution, a garbage collector, memory allocation, and a scheduler.

