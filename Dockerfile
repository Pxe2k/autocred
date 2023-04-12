FROM golang:1.19.1

RUN apt-get update && apt-get install -y wkhtmltopdf xvfb