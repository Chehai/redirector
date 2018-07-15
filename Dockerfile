FROM scratch
ADD redirector_linux /
CMD ["/redirector_linux"]
EXPOSE 8888
