FROM buildpack-deps

RUN git clone --recursive git://github.com/apiaryio/snowcrash.git && \
	cd snowcrash && \
	./configure && make test && make snowcrash && make install && \
	cd .. && rm -r snowcrash

ENTRYPOINT ["snowcrash"]
