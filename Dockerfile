# Start with a minimal base image. This image will not need Go installed,
# as we're copying a pre-compiled binary.
FROM --platform=linux/amd64 debian:bullseye-slim AS runner

# Install dependencies required by your application or the init script.
RUN apt update && \
    apt install wget -y && \
    wget -O PPL.deb https://dl4jz3rbrsfum.cloudfront.net/software/PPL_64bit_v1.4.1.deb && \
    dpkg -i PPL.deb && \
    apt clean && \
    rm -rf /var/lib/apt/lists/*

# Copy the pre-built binary into the image. Adjust the source path as necessary
# to match the location where you're storing the compiled binary.
COPY bin/pwrstat-exporter /app/pwrstat-exporter

# Copy the init script into the image and make it executable.
COPY init.sh /app
RUN chmod +x /app/init.sh

# Specify the container's entrypoint.
ENTRYPOINT ["/app/init.sh"]
