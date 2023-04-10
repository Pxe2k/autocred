FROM debian:buster-slim as prod
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates \
    # start deps needed for wkhtmltopdf
    curl \
    libxrender1 \
    libjpeg62-turbo \
    fontconfig \
    libxtst6 \
    xfonts-75dpi \
    xfonts-base \
    xz-utils && \
    # stop deps needed for wkhtmltopdf
    rm -rf /var/lib/apt/lists/*

RUN curl "https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.buster_amd64.deb" -L -o "wkhtmltopdf.deb"
RUN dpkg -i wkhtmltopdf.deb

RUN chmod -R 777 /usr/local/bin
