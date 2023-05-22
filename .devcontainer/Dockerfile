FROM ghcr.io/cirruslabs/flutter:3.10.0

WORKDIR /home/

COPY . .
RUN bash ./setup.sh

ENV PATH="/root/.cargo/bin:$PATH"
