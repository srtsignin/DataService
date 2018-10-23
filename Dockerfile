FROM scratch

COPY . .
RUN ls

CMD /service.exe