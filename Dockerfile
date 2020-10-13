FROM scratch

COPY ./server /
COPY ./templates/ /templates/

WORKDIR /

CMD [ "/server" ]
