# Project Setup Guide

This guide will walk you through setting up the project using Docker and PostgreSQL.

## Prerequisites

- Docker installed on your system. You can download and install Docker from [here](https://www.docker.com/products/docker-desktop).

## Steps to Setup

1. **Build Docker Image**:

   Build the Docker image using the following command:

   ```bash
   docker build -t dbms-postgres .

2. **Run Docker Container**:

   Run the Docker container in detached mode, mapping port 5432 of the container to port 5432 of the host system. This allows you to access PostgreSQL running inside the container from your host machine.

   ```bash
   docker run -d -p 5432:5432 dbms-postgres

3. **Access PostgreSQL**:

    You can access PostgreSQL using `psql` command-line tool. Run the following command:

   ```bash
   psql -h 127.0.0.1 -U userpharmacy -d dbpharmacy

4. **Run API local**:

   After the whole steps, you can access API run the makefile `web` and after accessing to `localhost:8080` on browser