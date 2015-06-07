FROM scratch

ADD https://github.com/nickschuch/karma/releases/download/0.0.1/karma-linux-amd64 /karma
RUN chmod a+x /karma

EXPOSE 80

ENTRYPOINT ["/karma"]
CMD ["--help"]
