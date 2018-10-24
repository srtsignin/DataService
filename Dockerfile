FROM bash

RUN ls
ADD data-service /
RUN ls


CMD ["/data-service"]