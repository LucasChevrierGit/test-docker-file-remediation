  FROM debian:12.2                                                                                                                                                 

  RUN apt-get update && apt-get install -y \
      openssl \
      curl \
      bash

  COPY . /app