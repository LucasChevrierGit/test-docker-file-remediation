FROM ubuntu:22.04                                                                                                                                                                                                       
RUN apt-get update && apt-get install -y \
    curl \
    openssl \
    && rm -rf /var/lib/apt/lists/*                                                                                                                                                                                      
COPY . /app                                                                                                                                                                                                             
WORKDIR /app                                                                                                                                                                                                            
CMD ["./start.sh"]                                                                                                                                                                                                      
