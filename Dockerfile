FROM debian:8.7
MAINTAINER  Andrew Minkin <minkin.andrew@gmail.com>
RUN  apt-get update \
  && apt-get install -y --no-install-recommends ca-certificates locales \
  && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
RUN localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG en_US.utf8
COPY comedian /
COPY goose /
COPY en.toml / 
COPY ru.toml /
COPY entrypoint.sh /
ADD migrations /migrations
ENTRYPOINT ["/entrypoint.sh"]`