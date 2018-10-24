FROM bash

RUN ls

ADD data_service_runner /

RUN ls

# CMD ["/data_service_runner"]
CMD tail -f /dev/null